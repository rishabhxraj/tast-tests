// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"time"

	"chromiumos/tast/common/testexec"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         NestedVM,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test that /dev/kvm is present and basic functionality works",
		Contacts:     []string{"sidereal@google.com", "cros-containers-dev@google.com"},
		SoftwareDeps: []string{"chrome", "vm_host", "untrusted_vm", "amd64"},
		Attr:         []string{"group:mainline", "informational"},
		// The build_image script strips out .c files, so we
		// have to use a different extension here.
		Data: []string{"kvm_test.c-bypass-mask"},
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

func NestedVM(ctx context.Context, s *testing.State) {
	cont := s.FixtValue().(crostini.FixtureData).Cont

	if err := cont.Command(ctx, "ls", "/dev/kvm").Run(testexec.DumpLogOnError); err != nil {
		s.Fatal("Failed to find /dev/kvm in VM: ", err)
	}

	if err := cont.PushFile(ctx, s.DataPath("kvm_test.c-bypass-mask"), "kvm_test.c"); err != nil {
		s.Fatal("Failed to push test program to VM: ", err)
	}

	if err := cont.Command(ctx, "gcc", "kvm_test.c").Run(testexec.DumpLogOnError); err != nil {
		s.Fatal("Failed to build test program: ", err)
	}

	if output, err := cont.Command(ctx, "./a.out").CombinedOutput(testexec.DumpLogOnError); err != nil {
		s.Error("Failed to run test program: ", err)
		s.Error("Test program output: ", string(output))
	}
}
