// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"time"

	"chromiumos/tast/local/chrome/vmc"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/local/crostini/ui/settings"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         RemoveOk,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test uninstalling Crostini via the Settings app",
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

func RemoveOk(ctx context.Context, s *testing.State) {
	tconn := s.FixtValue().(crostini.FixtureData).Tconn
	cr := s.FixtValue().(crostini.FixtureData).Chrome

	// Open the Linux settings.
	st, err := settings.OpenLinuxSettings(ctx, tconn, cr)
	if err != nil {
		s.Fatal("Failed to open Linux Settings: ", err)
	}

	// Click button remove and click Delete on the confirmation dialog.
	if err := st.Remove()(ctx); err != nil {
		s.Fatal("Failed to remove Linux: ", err)
	}

	// Verify "vmc list" outputs 0.
	hash, err := vmc.UserIDHash(ctx)
	if err != nil {
		s.Fatal("Failed to get CROS_USER_ID_HASH: ", err)
	}
	result, err := vmc.Command(ctx, hash, "list").Output()
	if err != nil {
		s.Fatal("Failed to run 'vmc list': ", err)
	}

	const expectedResult = "Total Size (bytes): 0\n"
	if string(result) != expectedResult {
		s.Fatalf("Failed to verify the result of 'vmc list', got %s, want %s", string(result), expectedResult)
	}
}
