// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crash

import (
	"context"
	"io/ioutil"
	"os"
	"strings"

	"chromiumos/tast/common/testexec"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/crash"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:     MountFailure,
		Desc:     "Verify mount and umount failures are logged as expected",
		Contacts: []string{"sarthakkukreti@google.com", "cros-telemetry@google.com"},
		Attr:     []string{"group:mainline"},
	})
}

// mountFailureCrashLog describes debug logs of interest during mount failure.
type mountFailureCrashLog struct {
	// path is the location where the log is stored. Empty path implies that the log isn't generated by this test.
	path string
	// contents contains fake contents to store in the log before the test.
	contents string
	// wantHeader is the expected header that is present if the log collection was successful.
	wantHeader string
}

// mountFailureCrash describes types of mount()/umount() crashes covered by this test, how to invoke these, and expected log commands.
type mountFailureCrash struct {
	// name denotes the crash name.
	name string
	// crashReporterArgs are the command line arguments used to invoke crash-reporter.
	crashReporterArgs []string
	// wantLogCommands are the expected log commands run by crash-reporter.
	wantLogCommands []string
}

// logCommandMap maps (u)mount failure related log commands to the respective crash logs collected by these commands.
type logCommandMap map[string]mountFailureCrashLog

// Description of logs in /run that are collected by the mount failure collector.
// Please see crash-reporter/crash_reporter_logs.conf for the usage of the equivalent commands in crash-reporter.
var mountFailureLogMap = logCommandMap{
	"shutdown_umount_failure_state": {"/run/shutdown_umount_failure.log", "log_shutdown_umount_failure", "===shutdown umount() failure logs==="},
	"dumpe2fs_stateful":             {"/run/dumpe2fs_stateful.log", "log_dumpe2fs_stateful", "===dumpe2fs (stateful partition)==="},
	"dumpe2fs_encstateful":          {"/run/mount_encrypted/dumpe2fs.log", "log_dumpe2fs_encstateful", "===dumpe2fs (/dev/mapper/encstateful)==="},
	"mount-encrypted":               {"/run/mount_encrypted/mount-encrypted.log", "log_mount-encrypted", "===mount-encrypted==="},
	"umount-encrypted":              {"/run/mount_encrypted/umount-encrypted.log", "log_umount-encrypted", "===umount-encrypted==="},
	// Ramoops and dmesg are collected directly: check if the header is present in the file.
	"console-ramoops": {"", "", "===ramoops==="},
	"kernel-warning":  {"", "", "===dmesg==="},
}

// Description of mount failures that are covered by this test and the logging each crash collection covers.
var mountFailures = []mountFailureCrash{
	{"mount_failure_stateful", []string{"--mount_failure", "--mount_device=stateful"}, []string{"console-ramoops", "kernel-warning", "dumpe2fs_stateful"}},
	{"mount_failure_encstateful", []string{"--mount_failure", "--mount_device=encstateful"}, []string{"console-ramoops", "kernel-warning", "dumpe2fs_encstateful", "mount-encrypted"}},
	{"umount_failure_stateful", []string{"--testonly_send_all", "--umount_failure", "--mount_device=stateful"}, []string{"shutdown_umount_failure_state", "umount-encrypted"}},
}

func saveMountFailureLogs() error {
	for _, log := range mountFailureLogMap {
		if log.path == "" {
			continue
		}
		if err := ioutil.WriteFile(log.path, []byte(log.contents), 0644); err != nil && !os.IsNotExist(err) {
			return errors.Wrapf(err, "failed to write %s", log.path)
		}
	}
	return nil
}

func cleanupMountFailureLogs() error {
	for _, log := range mountFailureLogMap {
		if log.path == "" {
			continue
		}
		if err := os.Remove(log.path); err != nil {
			return errors.Wrapf(err, "failed to remove %s", log.path)
		}
	}
	return nil
}

func reportMountFailures(ctx context.Context) error {
	for _, mf := range mountFailures {
		cmd := testexec.CommandContext(ctx, "/sbin/crash_reporter", mf.crashReporterArgs...)
		if err := cmd.Run(testexec.DumpLogOnError); err != nil {
			return errors.Wrapf(err, "crash_reporter command failed: %s", mf.name)
		}
	}
	return nil
}

func expectedFilesRegexes() []string {
	var res []string
	for _, mf := range mountFailures {
		res = append(res, mf.name+`\.\d{8}\.\d{6}\.\d+\.0\.log`, mf.name+`\.\d{8}\.\d{6}\.\d+\.0\.meta`)
	}
	return res
}

func validateCrashLogs(ctx context.Context, outDir string, files map[string][]string) error {
	for _, mf := range mountFailures {
		logFileRegex := mf.name + `\.\d{8}\.\d{6}\.\d+\.0\.log`

		if len(files[logFileRegex]) != 1 {
			crash.MoveFilesToOut(ctx, outDir, files[logFileRegex]...)
			return errors.Errorf("multiple log files (%v) within the same regex bucket: %s", files[logFileRegex], mf.name)
		}

		f := files[logFileRegex][0]
		contents, err := ioutil.ReadFile(f)
		if err != nil {
			return errors.Wrap(err, "couldn't read log file")
		}

		for _, cmd := range mf.wantLogCommands {
			if !strings.Contains(string(contents), mountFailureLogMap[cmd].wantHeader) {
				crash.MoveFilesToOut(ctx, outDir, f)
				return errors.Errorf("header not found: %s", mountFailureLogMap[cmd].wantHeader)
			}
			if !strings.Contains(string(contents), mountFailureLogMap[cmd].contents) {
				crash.MoveFilesToOut(ctx, outDir, f)
				return errors.Errorf("contents not found: %s", mountFailureLogMap[cmd].contents)
			}
		}
	}

	return nil
}

func MountFailure(ctx context.Context, s *testing.State) {
	if err := crash.SetUpCrashTest(ctx, crash.WithMockConsent()); err != nil {
		s.Fatal("SetUpCrashTest failed: ", err)
	}

	// Teardown on exiting the test.
	defer crash.TearDownCrashTest(ctx)

	err := saveMountFailureLogs()

	// Cleanup logs on exit.
	defer cleanupMountFailureLogs()

	if err != nil {
		s.Fatal("Failed to set up debug logs for mount failure collector: ", err)
	}

	// Run crash_reporter to generate the crashes.
	if err := reportMountFailures(ctx); err != nil {
		s.Fatal("Failed to report mount failure collectors: ", err)
	}

	s.Log("Waiting for crash files")

	wantFileRegs := expectedFilesRegexes()

	if len(wantFileRegs) == 0 {
		s.Fatal("No regexes to test against")
	}

	crashDirs, err := crash.GetDaemonStoreCrashDirs(ctx)
	if err != nil {
		s.Fatal("Couldn't get daemon store dirs: ", err)
	}
	// We might not be logged in, so also allow system crash dir.
	crashDirs = append(crashDirs, crash.SystemCrashDir)
	files, err := crash.WaitForCrashFiles(ctx, crashDirs, wantFileRegs)
	if err != nil {
		s.Fatal("Couldn't find expected files: ", err)
	}

	if err := validateCrashLogs(ctx, s.OutDir(), files); err != nil {
		s.Error("Failed to validate contents of the crash reporters: ", err)
	}

	if err := crash.RemoveAllFiles(ctx, files); err != nil {
		s.Log("Couldn't clean up files: ", err)
	}
}
