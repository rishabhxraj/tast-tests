// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package setup

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"chromiumos/tast/errors"
	"chromiumos/tast/local/chrome/internal/config"
	"chromiumos/tast/local/chrome/internal/driver"
	"chromiumos/tast/local/chrome/internal/extension"
	"chromiumos/tast/local/session"
	"chromiumos/tast/local/upstart"
	"chromiumos/tast/testing"
	"chromiumos/tast/timing"
)

// Obfuscated username is generated by SHA1, which has a digest length of 20 bytes.
// https://source.corp.google.com/chromeos_public/src/platform2/libbrillo/brillo/cryptohome.cc;l=83
var obfuscatedUsernameRegexp = regexp.MustCompile(`^[\da-f]{40}$`)

// RestartChromeForTesting restarts the ui job, asks session_manager to enable Chrome testing,
// and waits for Chrome to listen on its debugging port.
func RestartChromeForTesting(ctx context.Context, cfg *config.Config, exts *extension.Files) error {
	ctx, st := timing.Start(ctx, "restart")
	defer st.End()

	if err := restartSession(ctx, cfg); err != nil {
		return err
	}

	sm, err := session.NewSessionManager(ctx)
	if err != nil {
		return err
	}

	// Remove the file where Chrome will write its debugging port after it's restarted.
	if err := driver.PrepareForRestart(); err != nil {
		return err
	}

	testing.ContextLog(ctx, "Asking session_manager to enable Chrome testing")
	args := []string{
		"--remote-debugging-port=0",                  // Let Chrome choose its own debugging port.
		"--disable-logging-redirect",                 // Disable redirection of Chrome logging into cryptohome.
		"--ash-disable-system-sounds",                // Disable system startup sound.
		"--enable-experimental-extension-apis",       // Allow Chrome to use the Chrome Automation API.
		"--redirect-libassistant-logging",            // Redirect libassistant logging to /var/log/chrome/.
		"--no-first-run",                             // Prevent showing up offer pages, e.g. google.com/chromebooks.
		"--cros-region=" + cfg.Region(),              // Force the region.
		"--cros-regions-mode=hide",                   // Ignore default values in VPD.
		"--enable-oobe-test-api",                     // Enable OOBE helper functions for authentication.
		"--disable-hid-detection-on-oobe",            // Skip OOBE check for keyboard/mouse on chromeboxes/chromebases.
		"--force-hwid-check-result-for-test=success", // Forcefully ignore incorrect hardware IDs on devices.
		"--keep-login-events-for-testing",            // Keep LoginEventRecorder data for later retrieval by tests.
		"--force-color-profile=srgb",                 // Force chrome to treat the display as sRGB. See b/221643955 for details.
		"--force-raster-color-profile=srgb",          // Force rendering to run in the sRGB color space. See b/221643955 for details.
	}
	if !cfg.EnableRestoreTabs() {
		args = append(args, "--no-startup-window") // Do not start up chrome://newtab by default to avoid unexpected patterns (doodle etc.)
	}
	if cfg.HideCrashRestoreBubble() {
		args = append(args, "--hide-crash-restore-bubble") // Do not show "Chrome did not shut down correctly" bubble
	}

	if cfg.SkipOOBEAfterLogin() {
		args = append(args, "--oobe-skip-postlogin")
	}

	if !cfg.InstallWebApp() {
		args = append(args, "--disable-features=DefaultWebAppInstallation")
	}

	if cfg.VKEnabled() {
		args = append(args, "--enable-virtual-keyboard")
	}

	if cfg.ForceLaunchBrowser() {
		args = append(args, "--force-launch-browser")
	}

	// Enable verbose logging on some enrollment related files.
	if cfg.EnableLoginVerboseLogs() {
		args = append(args,
			"--vmodule="+strings.Join([]string{
				"*auto_enrollment_check_screen*=1",
				"*enrollment_screen*=1",
				"*login_display_host_common*=1",
				"*wizard_controller*=1",
				"*auto_enrollment_controller*=1"}, ","))
	}

	if cfg.LoginMode() != config.GAIALogin && cfg.EnrollMode() != config.GAIAEnroll {
		args = append(args, "--disable-gaia-services")
	}

	// Enable verbose logging on gaia_auth_fetcher to help debug some login failures. See crbug.com/1166530
	if cfg.LoginMode() == config.GAIALogin {
		args = append(args, "--vmodule=gaia_auth_fetcher=1")
		if cfg.UseSandboxGaia() {
			args = append(args, "--gaia-url=https://accounts.sandbox.google.com")
		}
	}
	if cfg.SkipForceOnlineSignInForTesting() {
		args = append(args, "--skip-force-online-signin-for-testing")
	}

	args = append(args, exts.ChromeArgs()...)
	if cfg.PolicyEnabled() {
		args = append(args, "--profile-requires-policy=true")
	} else {
		args = append(args, "--profile-requires-policy=false")
	}
	if cfg.RealtimeReportingAddr() != "" {
		args = append(args, "--realtime-reporting-url="+cfg.RealtimeReportingAddr())
	}
	if cfg.EncryptedReportingAddr() != "" {
		args = append(args, "--encrypted-reporting-url="+cfg.EncryptedReportingAddr())
	}
	if cfg.DMSAddr() != "" {
		args = append(args, "--device-management-url="+cfg.DMSAddr())

		// crbug.com/1225937
		// When using a fake DMServer, the tokens provided to the device are not
		// valid for reporting and trigger errors on the reporting server.
		if cfg.RealtimeReportingAddr() == "" {
			args = append(args, "--realtime-reporting-url=http://invalid.domain.name:54321/")
		}
		if cfg.EncryptedReportingAddr() == "" {
			args = append(args, "--encrypted-reporting-url=http://invalid.domain.name:54321/")
		}
	}
	if cfg.DisablePolicyKeyVerification() {
		args = append(args, "--disable-policy-key-verification")
	}
	switch cfg.ARCMode() {
	case config.ARCDisabled:
		// Make sure ARC is never enabled.
		args = append(args, "--arc-availability=none")
	case config.ARCEnabled:
		args = append(args,
			// Disable ARC opt-in verification to test ARC with mock GAIA accounts.
			"--disable-arc-opt-in-verification",
			// Always start ARC to avoid unnecessarily stopping mini containers.
			"--arc-start-mode=always-start-with-no-play-store")
	case config.ARCSupported:
		// Allow ARC being enabled on the device to test ARC with real gaia accounts.
		args = append(args, "--arc-availability=officially-supported")
	}
	if cfg.ARCMode() == config.ARCEnabled || cfg.ARCMode() == config.ARCSupported {
		args = append(args,
			// Do not sync the locale with ARC.
			"--arc-disable-locale-sync",
			// Do not update Play Store automatically.
			"--arc-play-store-auto-update=off",
			// Make 1 Android pixel always match 1 Chrome devicePixel.
			"--force-remote-shell-scale=")
		if !cfg.RestrictARCCPU() {
			args = append(args,
				// Disable CPU restrictions to let tests run faster
				"--disable-arc-cpu-restriction")
		}
	}

	if cfg.ARCMode() == config.ARCEnabled && cfg.ARCUseHugePages() == true {
		args = append(args,
			// Enable huge pages for guest memory
			"--arcvm-use-hugepages")
	}

	if fs := cfg.EnableFeatures(); len(fs) != 0 {
		args = append(args, "--enable-features="+strings.Join(fs, ","))
	}
	if fs := cfg.DisableFeatures(); len(fs) != 0 {
		args = append(args, "--disable-features="+strings.Join(fs, ","))
	}

	// Lacros features and additional args used to launch lacros-chrome should be delimited by
	// '####' and passed in from ash-chrome as a single argument with --lacros-chrome-additional-args.
	// See browser_manager.cc in Chrome source.
	// Example:
	//   --lacros-chrome-additional-args="--enable-features=Feature1,Feature2####--disable-features=Feature3####--foo=bar"
	// will result in multiple arguments passed to lacros-chrome:
	//   --enable-features=Feature1,Feature2
	//   --disable-features=Feature3
	//   --foo=bar
	var largs []string
	if fs := cfg.LacrosEnableFeatures(); len(fs) != 0 {
		largs = append(largs, "--enable-features="+strings.Join(fs, ","))
	}
	if fs := cfg.LacrosDisableFeatures(); len(fs) != 0 {
		largs = append(largs, "--disable-features="+strings.Join(fs, ","))
	}
	if as := cfg.LacrosExtraArgs(); len(as) != 0 {
		largs = append(largs, as...)
	}
	if len(largs) != 0 {
		args = append(args, "--lacros-chrome-additional-args="+strings.Join(largs, "####"))
	}

	// TODO(b/207576612): Remove this explicit override once all tests have migrated
	// to use the new Files app SWA version.
	if cfg.EnableFilesAppSWA() {
		args = append(args, "--enable-features=FilesSWA")
	} else {
		args = append(args, "--disable-features=FilesSWA")
	}

	args = append(args, cfg.ExtraArgs()...)
	var envVars []string
	if cfg.BreakpadTestMode() {
		envVars = append(envVars,
			"CHROME_HEADLESS=",
			"BREAKPAD_DUMP_LOCATION=/home/chronos/crash") // Write crash dumps outside cryptohome.
	}

	ctx, st = timing.Start(ctx, "enable_chrome_testing")
	defer st.End()

	_, err = sm.EnableChromeTestingAndWait(ctx, true, args, envVars)
	return err
}

// restartSession stops the "ui" job, clears policy files and the user's cryptohome if requested,
// and restarts the job.
func restartSession(ctx context.Context, cfg *config.Config) error {
	ctx, st := timing.Start(ctx, "restart_ui")
	defer st.End()

	ctx, cancel := context.WithTimeout(ctx, upstart.UIRestartTimeout)
	defer cancel()

	testing.ContextLog(ctx, "Stopping ui job")
	if err := upstart.StopJob(ctx, "ui"); err != nil {
		return err
	}

	if err := clearUserData(ctx, cfg); err != nil {
		return err
	}

	fixLogSymlink(ctx)

	testing.ContextLog(ctx, "Starting ui job")
	return upstart.EnsureJobRunning(ctx, "ui")
}

func clearUserData(ctx context.Context, cfg *config.Config) error {
	if cfg.KeepState() {
		return nil
	}

	testing.ContextLog(ctx, "Clearing user data")
	ctx, st := timing.Start(ctx, "clear_user_data")
	defer st.End()

	const chronosDir = "/home/chronos"
	const shadowDir = "/home/.shadow"

	if !cfg.KeepOwnership() {
		// This always fails because /home/chronos is a mount point, but all files
		// under the directory should be removed.
		os.RemoveAll(chronosDir)
		fis, err := ioutil.ReadDir(chronosDir)
		if err != nil {
			return err
		}
		// Retry cleanup of remaining files. Don't fail if removal reports an error.
		for _, left := range fis {
			if err := os.RemoveAll(filepath.Join(chronosDir, left.Name())); err != nil {
				testing.ContextLogf(ctx, "Failed to clear %s; failed to remove %q: %v", chronosDir, left.Name(), err)
			} else {
				testing.ContextLogf(ctx, "Failed to clear %s; %q needed repeated removal", chronosDir, left.Name())
			}
		}
	}

	// Delete files from shadow directory.
	shadowFiles, err := ioutil.ReadDir(shadowDir)
	if err != nil {
		return errors.Wrapf(err, "failed to read directory %q", shadowDir)
	}
	for _, file := range shadowFiles {
		// Should not remove folders other than the user profiles.
		if !file.IsDir() || !obfuscatedUsernameRegexp.MatchString(file.Name()) {
			continue
		}
		// Only look for chronos file with names matching u-*.
		chronosName := filepath.Join(chronosDir, "u-"+file.Name())
		shadowName := filepath.Join(shadowDir, file.Name())
		// Remove the shadow directory if it does not have a corresponding chronos directory.
		if _, err := os.Stat(chronosName); err != nil && os.IsNotExist(err) {
			if err := os.RemoveAll(shadowName); err != nil {
				testing.ContextLogf(ctx, "Failed to remove %q: %v", shadowName, err)
			}
		}
	}

	if !cfg.KeepOwnership() {
		// Delete policy files to clear the device's ownership state since the account
		// whose cryptohome we'll delete may be the owner: http://crbug.com/897278
		if err := session.ClearDeviceOwnership(ctx); err != nil {
			return err
		}
	}
	return nil
}

// fixLogSymlink fixes the symlink at /var/log/chrome/chrome.
// This is a workaround for a bug in Chrome that causes the symlink to
// occasionally become a regular file and Chrome logs to stop rotating
// (b/187795771).
// This function must be called while the ui job is stopped.
func fixLogSymlink(ctx context.Context) {
	const (
		logDir          = "/var/log/chrome"
		symlinkName     = "chrome"
		timestampLayout = "20060102-150405"
	)

	if err := func() error {
		// Check if /var/log/chrome/chrome is a symlink.
		st, err := os.Lstat(filepath.Join(logDir, symlinkName))
		if os.IsNotExist(err) {
			return nil
		}
		if err != nil {
			return err
		}
		if st.Mode()&os.ModeType == os.ModeSymlink {
			return nil
		}

		// If the file is not a symlink, rename it so that a new symlink is created
		// later when starting the ui job.
		newName := fmt.Sprintf("%s_%s_b187795771", symlinkName, time.Now().Format(timestampLayout))
		testing.ContextLogf(ctx, "Warning: %s is not a symlink; renaming it to %s to recover log rotation (b/187795771)", filepath.Join(logDir, symlinkName), newName)
		if err := os.Rename(filepath.Join(logDir, symlinkName), filepath.Join(logDir, newName)); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		// Don't make the whole process fail even if we could not fix
		// the symlink.
		testing.ContextLogf(ctx, "Warning: failed to fix Chrome log symlink: %v; continuing anyway", err)
	}
}
