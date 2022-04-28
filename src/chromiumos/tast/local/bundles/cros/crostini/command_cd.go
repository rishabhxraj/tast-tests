// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/terminalapp"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         CommandCd,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test command cd in Terminal window",
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

func CommandCd(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cr := s.FixtValue().(crostini.FixtureData).Chrome
	keyboard := s.FixtValue().(crostini.FixtureData).KB
	cont := s.FixtValue().(crostini.FixtureData).Cont

	userName := strings.Split(cr.NormalizedUser(), "@")[0]

	// Open Terminal app.
	terminalApp, err := terminalapp.Launch(ctx, tconn)
	if err != nil {
		s.Fatal("Failed to open Terminal app: ", err)
	}

	outputFile := "test.txt"
	folderName := "testFolder"
	if err := uiauto.Combine("run command cd",
		// Create a test folder.
		terminalApp.RunCommand(keyboard, fmt.Sprintf("mkdir %s", folderName)),
		// Cd to the newly created folder.
		terminalApp.RunCommand(keyboard, fmt.Sprintf("cd %s", folderName)),
		// Run pwd to check the path has changed.
		terminalApp.RunCommand(keyboard, fmt.Sprintf("pwd > %s", outputFile)))(ctx); err != nil {
		s.Fatal("Failed to test command cd: ", err)
	}

	// Check the content of the test file.
	if err := cont.CheckFileContent(ctx, filepath.Join("/home", userName, folderName, outputFile), fmt.Sprintf("/home/%s/%s\n", userName, folderName)); err != nil {
		s.Fatal("Cd failed to take user into the newly created folder: ", err)
	}
}