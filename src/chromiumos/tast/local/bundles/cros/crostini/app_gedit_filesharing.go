// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"chromiumos/tast/ctxutil"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/apps"
	"chromiumos/tast/local/chrome"
	"chromiumos/tast/local/chrome/ash"
	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/chrome/uiauto/faillog"
	"chromiumos/tast/local/chrome/uiauto/filesapp"
	"chromiumos/tast/local/chrome/uiauto/nodewith"
	"chromiumos/tast/local/chrome/uiauto/role"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/sharedfolders"
	"chromiumos/tast/local/crostini/ui/terminalapp"
	"chromiumos/tast/local/input"
	"chromiumos/tast/local/uidetection"
	"chromiumos/tast/local/vm"
	"chromiumos/tast/testing"
)

var (
	tmpFilename              = "testfile.txt"
	tmpFileCrosDownloadsPath = filepath.Join(filesapp.DownloadPath, tmpFilename)
	tmpFileCrostiniMountPath = filepath.Join(sharedfolders.MountPathDownloads, tmpFilename)
	tmpFileContents          = "This is a text string in a text file in the Downloads folder."
	geditContextMenuItem     = "Open with Text Editor"
	viewContextMenuItem      = "View"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         AppGeditFilesharing,
		LacrosStatus: testing.LacrosVariantUnknown,
		Desc:         "Test gedit in Terminal window",
		Contacts: []string{
			"ashpakov@google.com", // until Oct 2022
			"cros-containers-dev@google.com",
		},
		Attr:         []string{"group:mainline", "informational"},
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

func AppGeditFilesharing(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cr := s.FixtValue().(crostini.FixtureData).Chrome
	keyboard := s.FixtValue().(crostini.FixtureData).KB
	cont := s.FixtValue().(crostini.FixtureData).Cont

	// Use a shortened context for test operations to reserve time for cleanup.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, 5*time.Second)
	defer cancel()
	defer func(ctx context.Context) {
		if err := os.Remove(tmpFileCrosDownloadsPath); err != nil {
			s.Logf("Cleanup: failed to remove file %s on cleanup: %v", tmpFileCrosDownloadsPath, err)
		}
	}(cleanupCtx)

	defer faillog.DumpUITreeOnError(cleanupCtx, s.OutDir(), s.HasError, tconn)

	// Create a temp text file in the /Downloads folder to use in this test.
	if err := ioutil.WriteFile(tmpFileCrosDownloadsPath, []byte(tmpFileContents), 0644); err != nil {
		s.Fatal("Failed to create text file in Downloads folder: ", err)
	}

	filesApp, err := filesapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Files app: ", err)
	}

	// Open tmp file with Gedit.
	err = uiauto.Combine("open tmp file with Gedit",
		filesApp.OpenDownloads(),
		filesApp.ClickContextMenuItem(tmpFilename, filesapp.OpenWith, geditContextMenuItem),
	)(ctx)
	if err != nil {
		s.Fatal("Failed to open tmp file in the Downloads folder: ", err)
	}

	// Launch terminal so we can run commands in the container.
	terminalApp, err := terminalapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Crostini Terminal: ", err)
	}

	// Restart crostini in the end in case any error in the middle and gedit is not closed.
	// This also closes the Terminal window.
	restartIfError := true
	defer func() {
		if restartIfError {
			if err := terminalApp.RestartCrostini(keyboard, cont, cr.NormalizedUser())(cleanupCtx); err != nil {
				s.Log("Failed to restart crostini: ", err)
			}
		}
	}()

	geditWindow := nodewith.NameContaining(tmpFilename).Role(role.Window).First()
	filesAppShelfButton := nodewith.Name(apps.Files.Name).ClassName("ash/ShelfAppButton")
	ui := uiauto.New(tconn)
	ud := uidetection.NewDefault(tconn)

	err = checkFilesharingBeforeRestart(
		ctx, cont, tconn, ui, ud, filesApp, keyboard, geditWindow, filesAppShelfButton)
	if err != nil {
		s.Fatal("Failed test cases before restart of Crostini: ", err)
	}

	// Restart Crostini. This will start linux and the terminal and then close the terminal, but not Linux.
	s.Log("Restarting Crostini for the rest of this test")
	if err := terminalApp.RestartCrostini(keyboard, cont, cr.NormalizedUser())(ctx); err != nil {
		s.Log("Failed to restart Crostini: ", err)
	}

	err = checkFilesharingWorksAfterRestart(
		ctx, cont, tconn, filesApp, keyboard, ui, ud, geditWindow)
	if err != nil {
		s.Fatal("Failed tests to check whether file sharing  works correctly after restart of Crostini: ", err)
	}

	restartIfError = false
}

func checkFilesharingBeforeRestart(
	ctx context.Context,
	cont *vm.Container,
	tconn *chrome.TestConn,
	ui *uiauto.Context,
	ud *uidetection.Context,
	filesApp *filesapp.FilesApp,
	keyboard *input.KeyboardEventWriter,
	geditWindow,
	filesAppShelfButton *nodewith.Finder) error {

	// Check that file is mounted in the container.
	if err := crostini.VerifyFileInContainer(ctx, cont, tmpFileCrostiniMountPath); err != nil {
		return errors.Wrap(err, "temp file missing from crostini container")
	}

	if err := closeGedit(ctx, keyboard, ui, geditWindow); err != nil {
		return errors.Wrap(err, "failed to close Gedit")
	}

	// Check file remains mounted in container.
	if err := crostini.VerifyFileInContainer(ctx, cont, tmpFileCrostiniMountPath); err != nil {
		return errors.Wrap(err, "failed to find tmp file in crostini container after linux app is closed and before VM rebooted")
	}

	// Bring Downloads folder to foreground, open the file with Chrome (a non-linux app), validate file contents.
	if _, err := ash.BringWindowToForeground(ctx, tconn, "Files - Downloads"); err != nil {
		return errors.Wrap(err, "failed to bring the Filesapp Download window to foreground")
	}

	err := uiauto.Combine("validate contents of text file opened with Chrome (a non-linux app)",
		filesApp.ClickContextMenuItem(tmpFilename, filesapp.OpenWith, viewContextMenuItem),
		ud.WaitUntilExists(uidetection.TextBlock(strings.Split(tmpFileContents, " "))),
	)(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to right click on text file in Downloads folder and open with default Text app")
	}

	return nil
}

func checkFilesharingWorksAfterRestart(
	ctx context.Context,
	cont *vm.Container,
	tconn *chrome.TestConn,
	filesApp *filesapp.FilesApp,
	keyboard *input.KeyboardEventWriter,
	ui *uiauto.Context,
	ud *uidetection.Context,
	geditWindow *nodewith.Finder) error {

	// Check that file is no longer mounted in the container after Linux is restarted.
	if err := crostini.VerifyFileNotInContainer(ctx, cont, tmpFileCrostiniMountPath); err != nil {
		return errors.Wrap(err, "error verifying that file is no longer in container")
	}

	// Find and close the Chrome (a non-linux app) window.
	w, err := ash.FindWindow(ctx, tconn, func(w *ash.Window) bool { return w.Title == fmt.Sprintf("Chrome - %s", tmpFilename) })
	if err != nil {
		return errors.Wrap(err, "failed to find Text App Window")
	}
	if err = w.CloseWindow(ctx, tconn); err != nil {
		return errors.Wrap(err, "failed to close Text App Window")
	}

	// Re-open file with Gedit and validate contents.
	err = uiauto.Combine("open downloads folder, open text file with Gedit, validate contents",
		filesApp.OpenDownloads(),
		filesApp.ClickContextMenuItem(tmpFilename, filesapp.OpenWith, geditContextMenuItem),
		ui.WaitUntilExists(geditWindow),
		ud.WaitUntilExists(uidetection.TextBlock(strings.Split(tmpFileContents, " "))),
	)(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to re-open Gedit to check file contents")
	}

	if err = closeGedit(ctx, keyboard, ui, geditWindow); err != nil {
		return errors.Wrap(err, "failed to close Gedit after Crostini restart")
	}
	return nil
}

func closeGedit(ctx context.Context, keyboard *input.KeyboardEventWriter, ui *uiauto.Context, geditWindow *nodewith.Finder) error {
	return uiauto.Combine("close out of Gedit",
		// Focus on the Gedit window.
		ui.LeftClick(geditWindow),
		// Close out of Gedit.
		keyboard.AccelAction("ctrl+W"),
		keyboard.AccelAction("ctrl+W"),
	)(ctx)
}