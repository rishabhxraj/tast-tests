// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"os"
	"time"

	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/settings"
	"chromiumos/tast/local/disk"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         ResizeSpaceConstrained,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test resizing disk of Crostini from the Settings with constrained host disk space",
		Contacts:     []string{"nverne@google.com", "cros-containers-dev@google.com"},
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

func ResizeSpaceConstrained(ctx context.Context, s *testing.State) {
	pre := s.FixtValue().(crostini.FixtureData)
	cr := pre.Chrome
	tconn := pre.Tconn
	keyboard := pre.KB
	cont := pre.Cont

	// Open the Linux settings.
	st, err := settings.OpenLinuxSettings(ctx, tconn, cr)
	if err != nil {
		s.Fatal("Failed to open Linux Settings: ", err)
	}
	const GB uint64 = 1 << 30
	targetDiskSizeBytes := []uint64{20 * GB, 10 * GB, 5 * GB, 1 * GB, 5 * GB, 10 * GB}
	currSizeStr, err := st.GetDiskSize(ctx)
	if err != nil {
		s.Fatal("Failed to get current disk size: ", err)
	}
	currSizeBytes, err := settings.ParseDiskSize(currSizeStr)
	if err != nil {
		s.Fatalf("Failed to parse disk size string %s: %v", currSizeStr, err)
	}
	const fillPath = "/home/user/"
	for _, tBytes := range targetDiskSizeBytes {
		freeSpace, err := disk.FreeSpace(fillPath)
		if err != nil {
			s.Fatalf("Failed to read free space in %s: %v", fillPath, err)
		}
		if freeSpace < tBytes {
			s.Logf("Not enough free space to run test. Have %v, need %v", freeSpace, tBytes)
			continue
		}
		fillFile, err := disk.FillUntil(fillPath, tBytes)
		if err != nil {
			s.Fatal("Failed to fill disk space: ", err)
		}

		s.Logf("Resizing from %v to %v", currSizeBytes, tBytes)
		sizeOnSlider, sizeInCont, err := st.Resize(ctx, keyboard, tBytes, currSizeBytes)
		if err != nil {
			s.Fatal("Failed to resize back to the default value: ", err)
		}
		if err := st.VerifyResizeResults(ctx, cont, sizeOnSlider, sizeInCont); err != nil {
			s.Fatal("Failed to verify resize results: ", err)
		}

		currSizeBytes = tBytes
		if err = os.Remove(fillFile); err != nil {
			s.Fatalf("Failed to remove fill file %s: %v", fillFile, err)
		}
	}

}
