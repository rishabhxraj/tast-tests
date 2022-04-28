// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"time"

	"github.com/google/go-cmp/cmp"

	"chromiumos/tast/ctxutil"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/chrome"
	"chromiumos/tast/local/chrome/uiauto/filesapp"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/settings"
	"chromiumos/tast/local/crostini/ui/sharedfolders"
	"chromiumos/tast/local/crostini/ui/terminalapp"
	"chromiumos/tast/local/vm"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         ShareFilesRestart,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test shared folders are persistent after restarting Crostini",
		Contacts:     []string{"jinrongwu@google.com", "cros-containers-dev@google.com"},
		Attr:         []string{"group:mainline", "informational"},
		SoftwareDeps: []string{"chrome", "vm_host"},
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
func ShareFilesRestart(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cont := s.FixtValue().(crostini.FixtureData).Cont
	cr := s.FixtValue().(crostini.FixtureData).Chrome
	keyboard := s.FixtValue().(crostini.FixtureData).KB

	// Use a shortened context for unshare all folders.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, 5*time.Second)
	defer cancel()

	// Open Files app.
	filesApp, err := filesapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Files app: ", err)
	}

	sharedFolders := sharedfolders.NewSharedFolders(tconn)
	// Clean up shared folders in the end.
	defer func() {
		if err := sharedFolders.UnshareAll(cont, cr)(cleanupCtx); err != nil {
			s.Error("Failed to unshare all folders: ", err)
		}
	}()

	if err := sharedFolders.ShareMyFilesOK(ctx, filesApp)(ctx); err != nil {
		s.Fatal("Failed to share My files: ", err)
	}

	if err := checkResults(ctx, tconn, cont, cr); err != nil {
		s.Fatal("Faied to verify results after sharing My files: ", err)
	}

	// Restart Crostini.
	terminalApp, err := terminalapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to lauch terminal: ", err)
	}
	if err := terminalApp.RestartCrostini(keyboard, cont, cr.NormalizedUser())(ctx); err != nil {
		s.Fatal("Failed to restart crostini: ", err)
	}

	// Check the shared folders again after restart Crostini.
	if err := checkResults(ctx, tconn, cont, cr); err != nil {
		s.Fatal("Faied to verify results after restarting Crostini: ", err)
	}
}

func checkResults(ctx context.Context, tconn *chrome.TestConn, cont *vm.Container, cr *chrome.Chrome) error {
	// Check shared folders on the Settings app.
	st, err := settings.OpenLinuxSettings(ctx, tconn, cr, settings.ManageSharedFolders)
	if err != nil {
		return errors.Wrap(err, "failed to open Manage shared folders")
	}
	defer st.Close(ctx)

	shared, err := st.GetSharedFolders(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to find the shared folders list")
	}
	want := []string{sharedfolders.MyFiles}
	if diff := cmp.Diff(want, shared); diff != "" {
		return errors.Errorf("failed to verify shared folders list, got %s, want %s", shared, want)
	}

	// Check the file list in the container.
	if err := testing.Poll(ctx, func(ctx context.Context) error {
		list, err := cont.GetFileList(ctx, sharedfolders.MountPath)
		if err != nil {
			return err
		}
		want := []string{"fonts", sharedfolders.MountFolderMyFiles}
		if diff := cmp.Diff(want, list); diff != "" {
			return errors.Errorf("failed to verify file list in /mnt/chromeos, got %s, want %s", list, want)
		}
		return nil
	}, &testing.PollOptions{Timeout: 5 * time.Second}); err != nil {
		return errors.Wrap(err, "failed to verify file list in container")
	}

	return nil
}