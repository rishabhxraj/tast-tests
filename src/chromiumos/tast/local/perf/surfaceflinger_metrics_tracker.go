// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package perf

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"chromiumos/tast/errors"
	"chromiumos/tast/local/arc"
)

// frame is a struct used to store frame information from a single SurfaceFlinger
// latency entry.
type frame struct {
	// When the app started to draw, in ns.
	draw int
	// The vsync immediately preceding SF submitting the frame to the h/w, in ns.
	vsync int
	// Timestamp immediately after SF submitted that frame to the h/w, in ns.
	submit int
}

// SurfaceFlingerMetrics contains information needed to call SurfaceFlinger and
// data returned by SurfaceFlinger.
type SurfaceFlingerMetrics struct {
	frames []frame
	// Contains the last saved timestamp. Used to filter out duplicates.
	lastTimestamp int
	surfaceName   string
	arc           *arc.ARC
	collecting    chan bool
	collectingErr chan error
}

// NewSurfaceFlingerMetrics returns a new instance of SurfaceFlinger.
func NewSurfaceFlingerMetrics(surfaceName string, a *arc.ARC) *SurfaceFlingerMetrics {
	f := &SurfaceFlingerMetrics{
		frames:        nil,
		lastTimestamp: -1,
		// TODO(b/230500493): change surfaceName to be obtained via some other method.
		surfaceName: surfaceName,
		arc:         a,
	}
	return f
}

// Start starts the SurfaceFlinger dumpsys for reporting latencies.
func (f *SurfaceFlingerMetrics) Start(ctx context.Context) error {
	if f.collecting != nil {
		return errors.New("SurfaceFlinger start failed since it was already started")
	}
	f.collecting = make(chan bool)
	f.collectingErr = make(chan error, 1)

	// TODO(b/230396035): Change the interval to be obtained based on the device's screen refresh rate.
	const screenRefreshInterval = 500 * time.Millisecond

	// Start running SurfaceFlinger.
	go func() {
		ticker := time.NewTicker(screenRefreshInterval)
		defer ticker.Stop()

		for {
			select {
			case <-f.collecting:
				close(f.collectingErr)
				return
			case <-ticker.C:
				err := f.recentFrames(ctx)
				if err != nil {
					f.collectingErr <- errors.Wrap(err, "failed to get recent frames")
					return
				}
			case <-ctx.Done():
				f.collectingErr <- ctx.Err()
				return
			}
		}
	}()

	return nil
}

// Stop stops the SurfaceFlinger dumpsys, calculates FPS and average latency, and returns them.
func (f *SurfaceFlingerMetrics) Stop(ctx context.Context) (fps, latency float64, err error) {
	if f.collecting == nil {
		return 0, 0, errors.New("SurfaceFlinger stop failed since it was never started")
	}
	// Stop SurfaceFlinger routine.
	close(f.collecting)
	select {
	case err = <-f.collectingErr:
		if err != nil {
			return 0, 0, err
		}
	case <-ctx.Done():
		return 0, 0, ctx.Err()
	}
	fps, latency, err = f.calculateMetrics()
	return fps, latency, err
}

// recentFrames saves all recent frames generated by SurfaceFlinger to the
// SurfaceFlingerMetrics struct.
func (f *SurfaceFlingerMetrics) recentFrames(ctx context.Context) error {
	/*
	   adb shell dumpsys SurfaceFlinger --latency <window name> prints some
	   information about the last 127 frames displayed in that window.
	   The data returned looks like this:
	   16954612
	   7657467895508   7657482691352   7657493499756
	   7657484466553   7657499645964   7657511077881
	   7657500793457   7657516600576   7657527404785
	   (...)

	   The first line is the refresh period (here 16.95 ms), it is followed
	   by 127 lines w/ 3 timestamps in nanosecond each:
	   A) when the app started to draw
	   B) the vsync immediately preceding SF submitting the frame to the h/w
	   C) timestamp immediately after SF submitted that frame to the h/w

	   We use the special "SurfaceView" window name because the statistics for
	   the activity's main window are not updated when the main web content is
	   composited into a SurfaceView.
	*/
	out, err := f.arc.SurfaceFlingerLatencyCommand(ctx, f.surfaceName).Output()
	if err != nil {
		return errors.Wrap(err, "failed to execute SurfaceFlinger start command")
	}
	lines := strings.Split(string(out), "\n")
	// Store all recent frames into the SurfaceFlingerMetrics object.
	for _, line := range lines {
		entries := strings.Split(line, "\t")

		// If the line does not contain three entries or if the first entry is "0", skip;
		// there are various situations in which an entry could be "0", e.g: when the
		// SF stats are cleared and immediately this function is called.
		if len(entries) != 3 || entries[0] == "0" {
			continue
		}

		// Parse each line into ints.
		var draw, vsync, submit int
		var parseError error
		if draw, parseError = strconv.Atoi(entries[0]); err != nil {
			return errors.Wrap(parseError, "failed to parse draw timestamp")
		}
		if vsync, parseError = strconv.Atoi(entries[1]); err != nil {
			return errors.Wrap(parseError, "failed to parse vsync timestamp")
		}
		if submit, parseError = strconv.Atoi(entries[2]); err != nil {
			return errors.Wrap(parseError, "failed to parse submit timestamp")
		}

		// A max integer denotes that a frame is still pending; continue if encountered.
		if draw == math.MaxInt || vsync == math.MaxInt || submit == math.MaxInt {
			continue
		}

		// If the frame is recent, append.
		if draw > f.lastTimestamp {
			f.frames = append(f.frames, frame{draw, vsync, submit})
		}
	}
	if len(f.frames) > 0 {
		// Reassign the latest timestamp.
		f.lastTimestamp = f.frames[len(f.frames)-1].draw
	}
	return nil
}

// calculateMetrics uses the collected frame data to calculate FPS and average latency respectively.
func (f *SurfaceFlingerMetrics) calculateMetrics() (fps, latency float64, err error) {
	// We need to normalize to seconds (divide by 1 billion).
	numFrames := len(f.frames)
	if numFrames < 2 {
		return 0, 0, errors.Errorf("Only %v frames recorded", numFrames)
	}
	// Calculate FPS.
	dt := float64(f.frames[numFrames-1].vsync-f.frames[0].vsync) / float64(time.Second)
	fps = float64(numFrames-1) / dt

	// Calculate average latency.
	latencies := 0.0
	for _, f := range f.frames {
		latencies += float64(f.submit-f.draw) / float64(time.Second)
	}
	latency = latencies / float64(numFrames)

	return fps, latency, nil
}