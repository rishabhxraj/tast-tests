// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"time"

	"chromiumos/tast/ctxutil"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/chrome"
	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/chrome/uiauto/faillog"
	"chromiumos/tast/local/chrome/uiauto/nodewith"
	"chromiumos/tast/local/chrome/uiauto/role"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/terminalapp"
	"chromiumos/tast/local/input"
	"chromiumos/tast/local/screenshot"
	"chromiumos/tast/local/vm"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         AppGedit,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test gedit in Terminal window",
		Contacts:     []string{"jinrongwu@google.com", "cros-containers-dev@google.com"},
		Attr:         []string{"group:mainline", "informational"},
		Vars:         screenshot.ScreenDiffVars,
		SoftwareDeps: []string{"chrome", "vm_host"},
		Params: []testing.Param{
			// Parameters generated by params_test.go. DO NOT EDIT.
			{
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniAppTest,
				Fixture:           "crostiniBusterLargeContainer",
				Timeout:           15 * time.Minute,
			},
		},
	})
}
func AppGedit(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cr := s.FixtValue().(crostini.FixtureData).Chrome
	keyboard := s.FixtValue().(crostini.FixtureData).KB
	cont := s.FixtValue().(crostini.FixtureData).Cont

	// Use a shortened context for test operations to reserve time for cleanup.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, 5*time.Second)
	defer cancel()

	// Open Terminal app.
	terminalApp, err := terminalapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Terminal app: ", err)
	}

	restartIfError := true

	defer func() {
		// Restart crostini in the end in case any error in the middle and gedit is not closed.
		// This also closes the Terminal window.
		if restartIfError {
			if err := terminalApp.RestartCrostini(keyboard, cont, cr.NormalizedUser())(cleanupCtx); err != nil {
				s.Log("Failed to restart crostini: ", err)
			}
		}
	}()

	defer faillog.DumpUITreeOnError(cleanupCtx, s.OutDir(), s.HasError, tconn)

	d, err := screenshot.NewDifferFromChrome(ctx, s, cr, screenshot.Config{DefaultOptions: screenshot.Options{WindowWidthDP: 652, WindowHeightDP: 484}})
	if err != nil {
		s.Fatal("Failed to start screen differ: ", err)
	}
	defer d.DieOnFailedDiffs()

	// Create a file using gedit in Terminal.
	if err := testCreateFileWithGedit(ctx, terminalApp, keyboard, tconn, cont, d); err != nil {
		s.Fatal("Failed to create file with gedit in Terminal: ", err)
	}

	restartIfError = false

}

func testCreateFileWithGedit(ctx context.Context, terminalApp *terminalapp.TerminalApp, keyboard *input.KeyboardEventWriter, tconn *chrome.TestConn, cont *vm.Container, d screenshot.Differ) error {
	const (
		testFile   = "test.txt"
		testString = "This is a test string"
	)

	ui := uiauto.New(tconn)
	appWindow := nodewith.NameStartingWith(testFile).Role(role.Window).First()
	if err := uiauto.Combine("Create file with Gedit",
		// Launch Gedit.
		terminalApp.RunCommand(keyboard, "gedit "+testFile),
		// Focus on the Gedit window and input string.
		ui.LeftClick(appWindow),
		keyboard.TypeAction(testString),
		// Press ctrl+S to save the file.
		keyboard.AccelAction("ctrl+S"),
		// Take screenshot.
		crostini.TakeAppScreenshot("gedit"),
		// Screendiff test. Retrying 4 times, every 600 millis as cursor blinks about
		// once a second, and blinking causes diffs to fail.
		d.DiffWindow(ctx, "gedit", screenshot.Retries(4), screenshot.RetryInterval(time.Millisecond*600)),
		// Press ctrl+W twice to exit window.
		keyboard.AccelAction("ctrl+W"),
		keyboard.AccelAction("ctrl+W"),
		// Check window close.
		ui.WaitUntilGone(appWindow),
	)(ctx); err != nil {
		return err
	}

	// Check the content of the test file.
	if err := cont.CheckFileContent(ctx, testFile, testString+"\n"); err != nil {
		return errors.Wrap(err, "failed to verify the content of the test file")
	}

	return nil
}
