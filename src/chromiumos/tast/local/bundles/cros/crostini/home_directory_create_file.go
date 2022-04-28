// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"time"

	"chromiumos/tast/common/testexec"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/apps"
	"chromiumos/tast/local/chrome"
	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/chrome/uiauto/filesapp"
	"chromiumos/tast/local/chrome/uiauto/nodewith"
	"chromiumos/tast/local/chrome/uiauto/role"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/input"
	"chromiumos/tast/local/vm"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         HomeDirectoryCreateFile,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test creating a file/folder in Linux files and container using a pre-built crostini image",
		Contacts:     []string{"jinrongwu@chromium.org"},
		Attr:         []string{"group:mainline"},
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

func HomeDirectoryCreateFile(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cont := s.FixtValue().(crostini.FixtureData).Cont
	kb := s.FixtValue().(crostini.FixtureData).KB

	// Open Files app.
	filesApp, err := filesapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Files app: ", err)
	}

	if err := filesApp.OpenLinuxFiles()(ctx); err != nil {
		s.Fatal("Failed to open Linux files: ", err)
	}

	if err := testCreateFolderFromLinuxFiles(ctx, filesApp, cont, kb); err != nil {
		s.Fatal("Failed to test creating files in Linux files: ", err)
	}

	if err := testCreateFileFromContainer(ctx, tconn, filesApp, cont); err != nil {
		s.Fatal("Failed to test creating files in container: ", err)
	}
}

func testCreateFolderFromLinuxFiles(ctx context.Context, filesApp *filesapp.FilesApp, cont *vm.Container, kb *input.KeyboardEventWriter) error {
	const dirName = "test_folder"

	// Files app doesn't have a way to directly create a file, but
	// we can create a folder, which is just as good.
	if err := filesApp.CreateFolder(kb, dirName)(ctx); err != nil {
		return errors.Wrapf(err, "failed to create new folder %q", dirName)
	}

	// Check that the file now exists in the container.
	if err := filesApp.WaitForFile(dirName)(ctx); err != nil {
		return errors.Wrapf(err, "creation of folder %q did not propagate to container", dirName)
	}
	return nil
}

func testCreateFileFromContainer(ctx context.Context, tconn *chrome.TestConn, filesApp *filesapp.FilesApp, cont *vm.Container) error {
	const fileName = "testfile.txt"

	// Create file in container.
	if err := cont.Command(ctx, "touch", fileName).Run(testexec.DumpLogOnError); err != nil {
		return errors.Wrap(err, "failed to create test file in the container")
	}

	refresh := nodewith.Name("Refresh").Role(role.Button).Ancestor(filesapp.WindowFinder(apps.Files.ID))
	if err := uiauto.New(tconn).LeftClickUntil(refresh, filesApp.FileExists(fileName))(ctx); err != nil {
		// Sometimes refresh does not work. Close and reopen Files app instead.
		testing.ContextLogf(ctx, "Failed to find the new file: %s, try to relaunch Files app", err)
		filesApp, err := filesapp.Relaunch(ctx, tconn, filesApp)
		if err != nil {
			return errors.Wrap(err, "failed to relaunch Files app")
		}
		if err := filesApp.OpenLinuxFiles()(ctx); err != nil {
			return errors.Wrap(err, "failed to open Linux files")
		}
		return filesApp.FileExists(fileName)(ctx)
	}
	return nil
}
