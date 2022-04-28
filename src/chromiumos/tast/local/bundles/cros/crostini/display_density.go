// Copyright 2019 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"strings"
	"time"

	"chromiumos/tast/common/testexec"
	"chromiumos/tast/ctxutil"
	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/chrome/uiauto/faillog"
	"chromiumos/tast/local/coords"
	"chromiumos/tast/local/crostini"
	"chromiumos/tast/shutil"
	"chromiumos/tast/testing"
)

func init() {
	testing.AddTest(&testing.Test{
		Func:         DisplayDensity,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Runs a crostini application from the terminal in high/low DPI modes and compares sizes",
		Contacts:     []string{"smbarber@chromium.org", "cros-containers-dev@google.com"},
		Attr:         []string{"group:mainline", "informational"},
		SoftwareDeps: []string{"chrome", "vm_host"},
		Params: []testing.Param{
			// Parameters generated by display_density_test.go. DO NOT EDIT.
			{
				Name:              "wayland_buster_stable",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBuster",
				Timeout:           7 * time.Minute,
				Val:               crostini.WaylandDemoConfig(),
			}, {
				Name:              "wayland_buster_unstable",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBuster",
				Timeout:           7 * time.Minute,
				Val:               crostini.WaylandDemoConfig(),
			}, {
				Name:              "wayland_bullseye_stable",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBullseye",
				Timeout:           7 * time.Minute,
				Val:               crostini.WaylandDemoConfig(),
			}, {
				Name:              "wayland_bullseye_unstable",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBullseye",
				Timeout:           7 * time.Minute,
				Val:               crostini.WaylandDemoConfig(),
			}, {
				Name:              "x11_buster_stable",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBuster",
				Timeout:           7 * time.Minute,
				Val:               crostini.X11DemoConfig(),
			}, {
				Name:              "x11_buster_unstable",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBuster",
				Timeout:           7 * time.Minute,
				Val:               crostini.X11DemoConfig(),
			}, {
				Name:              "x11_bullseye_stable",
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Fixture:           "crostiniBullseye",
				Timeout:           7 * time.Minute,
				Val:               crostini.X11DemoConfig(),
			}, {
				Name:              "x11_bullseye_unstable",
				ExtraAttr:         []string{"informational"},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Fixture:           "crostiniBullseye",
				Timeout:           7 * time.Minute,
				Val:               crostini.X11DemoConfig(),
			},
		},
	})
}

func DisplayDensity(ctx context.Context, s *testing.State) {
	pre := s.FixtValue().(crostini.FixtureData)
	tconn := pre.Tconn
	cont := pre.Cont
	keyboard := pre.KB
	conf := s.Param().(crostini.DemoConfig)

	// Use a shortened context for test operations to reserve time for cleanup.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, 5*time.Second)
	defer cancel()

	type density int

	const (
		lowDensity density = iota
		highDensity
	)

	if err := uiauto.StartRecordFromKB(ctx, tconn, keyboard); err != nil {
		s.Log("Failed to start recording from keyboard: ", err)
	}

	defer uiauto.StopRecordFromKBAndSaveOnError(cleanupCtx, tconn, s.HasError, s.OutDir())
	defer faillog.DumpUITreeOnError(cleanupCtx, s.OutDir(), s.HasError, tconn)

	demoWindowSize := func(densityConfiguration density) (coords.Size, error) {
		windowName := conf.Name
		var subCommandArgs []string
		if densityConfiguration == lowDensity {
			windowName = windowName + "_low_density"
			// TODO(hollingum): Find a better way to pass environment vars to a container command (rather than invoking sh).
			subCommandArgs = append(subCommandArgs, "DISPLAY=${DISPLAY_LOW_DENSITY}", "WAYLAND_DISPLAY=${WAYLAND_DISPLAY_LOW_DENSITY}")
		}
		subCommandArgs = append(subCommandArgs, conf.AppPath, "--width=100", "--height=100", "--title="+windowName)

		cmd := cont.Command(ctx, "sh", "-c", strings.Join(subCommandArgs, " "))
		s.Logf("Running %q", shutil.EscapeSlice(cmd.Args))
		if err := cmd.Start(); err != nil {
			return coords.Size{}, err
		}
		defer cmd.Wait(testexec.DumpLogOnError)
		defer cmd.Kill()

		var sz coords.Size
		var err error
		if sz, err = crostini.PollWindowSize(ctx, tconn, windowName, 10*time.Second); err != nil {
			return coords.Size{}, err
		}
		s.Logf("Window %q size is %v", windowName, sz)

		s.Logf("Closing %q with keypress", windowName)
		err = keyboard.Accel(ctx, "Enter")

		return sz, err
	}

	sizeHighDensity, err := demoWindowSize(highDensity)
	if err != nil {
		s.Fatal("Failed getting high-density window size: ", err)
	}

	sizeLowDensity, err := demoWindowSize(lowDensity)
	if err != nil {
		s.Fatal("Failed getting low-density window size: ", err)
	}

	if err := crostini.VerifyWindowDensities(ctx, tconn, sizeHighDensity, sizeLowDensity); err != nil {
		s.Fatal("Failed during window density comparison: ", err)
	}
}
