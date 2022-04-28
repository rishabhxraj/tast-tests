// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"chromiumos/tast/ctxutil"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/bundles/cros/crostini/listset"
	"chromiumos/tast/local/chrome"
	"chromiumos/tast/local/chrome/uiauto/filesapp"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/settings"
	"chromiumos/tast/local/crostini/ui/sharedfolders"
	"chromiumos/tast/local/drivefs"
	"chromiumos/tast/local/vm"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         ShareDrive,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test sharing Google Drive with Crostini",
		Contacts:     []string{"jinrongwu@google.com", "cros-containers-dev@google.com"},
		Attr:         []string{"group:mainline", "informational"},
		SoftwareDeps: []string{"chrome", "vm_host", "chrome_internal", "drivefs"},
		Params: []testing.Param{
			// Parameters generated by params_test.go. DO NOT EDIT.
			{
				Name:              "buster_stable_gaia",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBusterGaia",
				Timeout:           7 * time.Minute,
			}, {
				Name:              "buster_unstable_gaia",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBusterGaia",
				Timeout:           7 * time.Minute,
			}, {
				Name:              "bullseye_stable_gaia",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBullseyeGaia",
				Timeout:           7 * time.Minute,
			}, {
				Name:              "bullseye_unstable_gaia",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBullseyeGaia",
				Timeout:           7 * time.Minute,
			},
		},
	})
}

func ShareDrive(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cont := s.FixtValue().(crostini.FixtureData).Cont
	cr := s.FixtValue().(crostini.FixtureData).Chrome

	// Use a shortened context for test operations to reserve time for cleanup.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, 5*time.Second)
	defer cancel()

	sharedFolders := sharedfolders.NewSharedFolders(tconn)
	// Clean up shared folders in the end.
	defer func() {
		if err := sharedFolders.UnshareAll(cont, cr)(cleanupCtx); err != nil {
			s.Error("Failed to unshare all folders: ", err)
		}
	}()

	// Open Files app.
	filesApp, err := filesapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Files app: ", err)
	}

	if err := testing.Poll(ctx, func(ctx context.Context) error {
		if err := sharedFolders.ShareDriveOK(ctx, filesApp)(ctx); err != nil {
			if errClose := filesApp.Close(cleanupCtx); errClose != nil {
				return errors.Wrap(errClose, "failed to close Files app")
			}
			var errOpen error
			if filesApp, errOpen = filesapp.Launch(ctx, tconn); errOpen != nil {
				return errors.Wrap(errOpen, "failed to open Files app")
			}
		}
		return err
	}, &testing.PollOptions{Timeout: 5 * time.Minute, Interval: time.Minute}); err != nil {
		s.Fatal("Failed to share Google Drive: ", err)
	}

	if err := checkDriveResults(ctx, tconn, cont, cr); err != nil {
		s.Fatal("Failed to verify sharing results: ", err)
	}

	const (
		testFile   = "testD.sh"
		echoString = "hello"
		testString = "#!/bin/sh\necho -n " + echoString
	)
	testFolder := fmt.Sprintf("testFolderD_%d", rand.Intn(1000000000))

	mountPath, err := drivefs.WaitForDriveFs(ctx, cr.NormalizedUser())
	if err != nil {
		s.Fatal("Failed waiting for DriveFS to start: ", err)
	}
	drivefsRoot := filepath.Join(mountPath, "root")
	folderPath := filepath.Join(drivefsRoot, testFolder)

	// Delete newly created files in the end.
	defer os.RemoveAll(folderPath)

	// Add a file and a folder in Drive.
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		s.Fatal("Failed to create test folder in Drive: ", err)
	}
	if err := ioutil.WriteFile(filepath.Join(folderPath, testFile), []byte(testString), 0755); err != nil {
		s.Fatal("Failed to create file in Drive: ", err)
	}
	// With the line above the file is not executable. Chmod is necessary here.
	if err := os.Chmod(filepath.Join(folderPath, testFile), 0755); err != nil {
		s.Fatal("Failed to chmod for test file in Drive: ", err)
	}

	// Check file list in the container.
	filesInCont, err := cont.GetFileList(ctx, sharedfolders.MountPathMyDrive)
	if err != nil {
		s.Fatal("Failed to get file list of /mnt/chromeos/MyFiles/MyDrive: ", err)
	}
	list, err := ioutil.ReadDir(drivefsRoot)
	if err != nil {
		s.Fatal("Failed to read files in Drive: ", err)
	}
	var filesInDrive []string
	for _, f := range list {
		filesInDrive = append(filesInDrive, f.Name())
	}
	if err := listset.CheckListsMatch(filesInCont, filesInDrive...); err != nil {
		s.Fatal("Failed to verify the files list in container: ", err)
	}

	shared, err := cont.GetFileList(ctx, filepath.Join(sharedfolders.MountPathMyDrive, testFolder))
	if err != nil {
		s.Fatalf("Failed to get file list of /mnt/chromeos/MyFiles/MyDrive/%s: %s", testFolder, err)
	}
	if want := []string{testFile}; !reflect.DeepEqual(shared, want) {
		s.Fatalf("Failed to verify shared folders list, got %s, want %s", shared, want)
	}

	sharedFilePath := filepath.Join(sharedfolders.MountPathMyDrive, testFolder, testFile)

	// Check the content of the test file in the container.
	if err := cont.CheckFileContent(ctx, sharedFilePath, testString); err != nil {
		s.Fatal("Failed to verify the content of the test file: ", err)
	}

	// Check the permission.
	result, err := cont.Command(ctx, "ls", "-l", sharedFilePath).Output()
	if err != nil {
		s.Fatal("Failed to run ls on the test file in the container: ", err)
	}
	permission := strings.Split(string(result), " ")[0]
	expected := "-rwxr-x---"
	if permission != expected {
		s.Fatalf("Failed to verify the permission of shared file, got %s, want %s", permission, expected)
	}

	// Run the test file in shared folder.
	err = cont.Command(ctx, sharedFilePath).Run()
	if err != nil {
		s.Fatal("Was unable to run " + sharedFilePath)
	}

	// Copy file to home dir and run it.
	if err := cont.Command(ctx, "cp", sharedFilePath, ".").Run(); err != nil {
		s.Fatalf("Failed to copy %s to home directory: %s", sharedFilePath, err)
	}
	// Run the test file to make sure it is executable in home directory.
	result, err = cont.Command(ctx, "./"+testFile).Output()
	if err != nil {
		s.Fatalf("Failed to run %s in home directory: %s", testFile, err)
	}
	if string(result) != echoString {
		s.Fatalf("Failed to verify the output of the test file, got %s, want %s", string(result), echoString)
	}
}

func checkDriveResults(ctx context.Context, tconn *chrome.TestConn, cont *vm.Container, cr *chrome.Chrome) error {
	// Check the shared folders on Settings.
	s, err := settings.OpenLinuxSettings(ctx, tconn, cr, settings.ManageSharedFolders)
	if err != nil {
		return errors.Wrap(err, "failed to find Manage shared folders")
	}
	defer s.Close(ctx)

	if shared, err := s.GetSharedFolders(ctx); err != nil {
		return errors.Wrap(err, "failed to find the shared folders list")
	} else if want := []string{sharedfolders.SharedDrive}; !reflect.DeepEqual(shared, want) {
		return errors.Errorf("failed to verify shared folders list, got %s, want %s", shared, want)
	}

	// Check the file list in the container.
	if err := testing.Poll(ctx, func(ctx context.Context) error {
		if list, err := cont.GetFileList(ctx, sharedfolders.MountPath); err != nil {
			return err
		} else if want := []string{"fonts", sharedfolders.MountFolderGoogleDrive}; !reflect.DeepEqual(list, want) {
			return errors.Errorf("failed to verify file list in /mnt/chromeos, got %s, want %s", list, want)
		}

		if list, err := cont.GetFileList(ctx, sharedfolders.MountPathGoogleDrive); err != nil {
			return err
		} else if want := []string{sharedfolders.MountFolderMyDrive}; !reflect.DeepEqual(list, want) {
			return errors.Errorf("failed to verify file list in /mnt/chromeos/GoogleDrive, got %s, want %s", list, want)
		}
		return nil
	}, &testing.PollOptions{Timeout: 5 * time.Second}); err != nil {
		return errors.Wrap(err, "failed to verify file list in container")
	}

	return nil
}
