// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"chromiumos/tast/common/testexec"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/chrome/uiauto/filesapp"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/sharedfolders"
	"chromiumos/tast/local/cryptohome"
	"chromiumos/tast/local/vm"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         NoAccessToDownloads,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Run a test to make sure Linux does not have access to downloads on Chrome using a pre-built crostini image",
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

func NoAccessToDownloads(ctx context.Context, s *testing.State) {
	cont := s.FixtValue().(crostini.FixtureData).Cont
	cr := s.FixtValue().(crostini.FixtureData).Chrome
	tconn := s.FixtValue().(crostini.FixtureData).Tconn

	if err := checkHomeDirInContainerEmpty(ctx, cont); err != nil {
		s.Fatal("Home directory in container is not empty by default: ", err)
	}

	if err := cont.CheckFileDoesNotExistInDir(ctx, sharedfolders.MountPath, sharedfolders.MountFolderMyFiles); err != nil {
		s.Fatal("MyFiles is unexpectedly listed in /mnt/chromeos in container by default: ", err)
	}

	// Create a file in Downloads.
	const fileName = "test.txt"
	ownerID, err := cryptohome.UserHash(ctx, cr.NormalizedUser())
	if err != nil {
		s.Fatal("Failed to get user hash: ", err)
	}
	filePath := filepath.Join("/home/user", ownerID, "Downloads", fileName)
	if err := ioutil.WriteFile(filePath, []byte("teststring"), 0644); err != nil {
		for _, dir := range []string{"/home/user", filepath.Join("/home/user", ownerID), filepath.Join("/home/user", ownerID, "Downloads")} {
			if stat, err := os.Stat(dir); err == nil && stat.IsDir() {
				s.Logf("%s exists", dir)
			} else {
				s.Logf("%s does not exist", dir)
				result, err := testexec.CommandContext(ctx, "mount").Output(testexec.DumpLogOnError)
				if err != nil {
					s.Fatal("Failed to run command mount")
				}
				s.Logf("Result of command mount is: %s", string(result))
				// Open Files app.
				filesApp, err := filesapp.Launch(ctx, tconn)
				if err != nil {
					s.Fatal("Failed to open Files app: ", err)
				}
				defer filesApp.Close(ctx)
				if err := filesApp.OpenDownloads()(ctx); err != nil {
					s.Fatal("Failed to open Downloads: ", err)
				}
				break
			}
		}
		s.Fatal("Failed to create a file in Downloads: ", err)
	}
	defer os.Remove(filePath)

	if err := checkHomeDirInContainerEmpty(ctx, cont); err != nil {
		s.Fatal("Home directory in container is not empty after creating a file in Downloads in Chrome: ", err)
	}

	if err := cont.CheckFileDoesNotExistInDir(ctx, sharedfolders.MountPath, sharedfolders.MountFolderMyFiles); err != nil {
		s.Fatal("MyFiles is unexpectedly listed in /mnt/chromeos in container: ", err)
	}
}

func checkHomeDirInContainerEmpty(ctx context.Context, cont *vm.Container) error {
	// Get file list in home directory in container.
	fileList, err := cont.GetFileList(ctx, ".")
	if err != nil {
		return errors.Wrap(err, "failed to list the content of home directory in container")
	}
	if len(fileList) != 0 {
		return errors.Errorf("home directory unexpectedly contains some files %s", fileList)
	}
	return nil
}
