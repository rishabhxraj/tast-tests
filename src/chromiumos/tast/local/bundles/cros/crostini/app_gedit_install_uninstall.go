// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"chromiumos/tast/common/testexec"
	"chromiumos/tast/ctxutil"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/chrome/uiauto/faillog"
	"chromiumos/tast/local/chrome/uiauto/launcher"
	"chromiumos/tast/local/chrome/uiauto/nodewith"
	"chromiumos/tast/local/chrome/uiauto/role"
	"chromiumos/tast/local/chrome/vmc"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/settings"
	"chromiumos/tast/local/crostini/ui/terminalapp"
	"chromiumos/tast/local/screenshot"
	"chromiumos/tast/local/uidetection"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         AppGeditInstallUninstall,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Install Gedit, check rendering, icons, saving and uninstall behavior",
		Contacts:     []string{"zubinpratap@google.com", "cros-containers-dev@google.com"},
		Attr:         []string{"group:mainline", "informational"},
		Vars:         screenshot.ScreenDiffVars,
		SoftwareDeps: []string{"chrome", "vm_host"},
		Data:         []string{"logo_gedit.png"},
		Params: []testing.Param{
			// Parameters generated by params_test.go. DO NOT EDIT.
			{
				Name:              "buster_stable",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBuster",
				Timeout:           7 * time.Minute,
			}, {
				Name:              "buster_unstable",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBuster",
				Timeout:           7 * time.Minute,
			}, {
				Name:              "bullseye_stable",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBullseye",
				Timeout:           7 * time.Minute,
			}, {
				Name:              "bullseye_unstable",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBullseye",
				Timeout:           7 * time.Minute,
			},
		},
	})
}
func AppGeditInstallUninstall(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	keyboard := s.FixtValue().(crostini.FixtureData).KB
	cont := s.FixtValue().(crostini.FixtureData).Cont

	// Use a shortened context for test operations to reserve time for cleanup.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, 5*time.Second)
	defer cancel()
	defer faillog.DumpUITreeOnError(cleanupCtx, s.OutDir(), s.HasError, tconn)

	// Check VM size before Gedit install.
	sizeBytesBeforeGedit, err := getVMSizeInBytes(ctx)
	if err != nil {
		s.Fatal("Failed to get VM size before installing Gedit: ", err)
	}

	// Install Gedit.
	s.Log("Installing Gedit")
	if err := cont.Command(ctx, "sudo", "apt-get", "update").Run(testexec.DumpLogOnError); err != nil {
		s.Fatal("Failed to run apt-update: ", err)
	}
	if err := cont.Command(ctx, "sudo", "apt-get", "-y", "install", "gedit").Run(testexec.DumpLogOnError); err != nil {
		s.Fatal("Failed to install Gedit: ", err)
	}

	// Check VM size after Gedit install.
	sizeBytesAfterGedit, err := getVMSizeInBytes(ctx)
	if err != nil {
		s.Fatal("Failed to get VM size after installing Gedit: ", err)
	}
	if err := assertGBSizeUnchanged(ctx, sizeBytesBeforeGedit, sizeBytesAfterGedit); err != nil {
		s.Fatal("VM size has changed unexpectedly after installing gedit: ", err)
	}

	// Launch and test Gedit.
	s.Log("Launching Gedit from launcher")
	if err := launcher.SearchAndLaunchWithQuery(tconn, keyboard, "t", "Text Editor")(ctx); err != nil {
		s.Fatal("Failed to launch gedit: ", err)
	}

	ud := uidetection.NewDefault(tconn)
	ui := uiauto.New(tconn)
	geditIcon := uidetection.CustomIcon(s.DataPath("logo_gedit.png"))
	if err := ud.WaitUntilExists(geditIcon)(ctx); err != nil {
		s.Fatal("Failed to find the Gedit icon in the shelf")
	}

	closeGeditContextMenuItem := uidetection.Word("Close")

	if err := uiauto.Combine("edit file and save",
		keyboard.TypeAction("Hello, gedit!"),
		keyboard.AccelAction("ctrl+S"), // Bring up the save window
		ud.LeftClick(uidetection.Word("Save")),
		ud.RightClick(geditIcon),
		ud.WaitUntilExists(closeGeditContextMenuItem),
		ud.LeftClick(closeGeditContextMenuItem),
	)(ctx); err != nil {
		s.Fatal("Failed to edit and save the Gedit file: ", err)
	}

	s.Log("Uninstall Gedit from Terminal")
	ta, err := terminalapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to launch terminal: ", err)
	}

	if err = ta.RunCommand(keyboard, "sudo apt-get -y remove gedit")(ctx); err != nil {
		s.Fatal("Failed to run command in Terminal window: ", err)
	}

	// Check VM size again.
	sizeBytesAfterGeditRemoved, err := getVMSizeInBytes(ctx)
	if err != nil {
		s.Fatal("Failed to get VM size after removing Gedit: ", err)
	}
	if err := assertGBSizeUnchanged(ctx, sizeBytesBeforeGedit, sizeBytesAfterGeditRemoved); err != nil {
		s.Fatal("VM size has changed unexpectedly after removing gedit: ", err)
	}

	// Close terminal.
	terminalNode := nodewith.NameRegex(regexp.MustCompile(`\@penguin\: `)).Role(role.Window).ClassName("BrowserFrame")
	leaveButton := uidetection.Word("Leave").Nth(1)
	if err := uiauto.Combine("close Terminal window",
		ta.WaitForPrompt(),             // Wait until Gedit uninstall streams finish printing.
		ta.ClickShelfMenuItem("Close"), // Closing terminal from the shelf throws up an alert dialogue.
		ud.LeftClick(leaveButton),      // Dialogue has "Cancel" and "Leave buttons"
		ui.WithTimeout(time.Minute).WaitUntilGone(terminalNode),
	)(ctx); err != nil {
		s.Fatal("Failed to close terminal: ", err)
	}
}

func getVMSizeInBytes(ctx context.Context) (int64, error) {
	hash, err := vmc.UserIDHash(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get CROS_USER_ID_HASH")
		// s.Fatal("Failed to get CROS_USER_ID_HASH: ", err)
	}
	cmdOut, err := vmc.Command(ctx, hash, "list").Output()
	if err != nil {
		return 0, errors.Wrap(err, "failed to run 'vmc list'")
		// s.Fatal("Failed to run 'vmc list': ", err)
	}
	testing.ContextLog(ctx, "Here is the result of 'vmc list': ", string(cmdOut))

	// 'vmc list' produces an output like
	// 'termina (10737434624 bytes, min shrinkable size 2723151872 bytes, raw)
	// Total Size (bytes): 10737434624'
	// Note the first and third ints are the same. We extract all digits
	// from the string with the regex pattern.
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(string(cmdOut), -1)

	// Tests that the output has the expected structure - first and third sizes must be the same.
	if matches[0] != matches[2] {
		return 0, errors.Errorf("expected termina VM sizes to match, but got two values - %s and %s", matches[0], matches[2])
	}

	return strconv.ParseInt(matches[0], 10, 64)
}

func assertGBSizeUnchanged(ctx context.Context, before, after int64) error {
	// Convert to GB and floating point.
	beforeSizeGB := float64(before / settings.SizeGB)
	afterSizeGB := float64(after / settings.SizeGB)

	if beforeSizeGB != afterSizeGB {
		return errors.Errorf("Termina VM of unexpected size- want %fGB but got %fGB", beforeSizeGB, afterSizeGB)
	}

	return nil
}
