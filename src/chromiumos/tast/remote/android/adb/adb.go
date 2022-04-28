// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package adb enables controlling android devices from remote test host via ADB.
package adb

import (
	"context"

	"chromiumos/tast/common/android/adb"
	"chromiumos/tast/common/testexec"
	"chromiumos/tast/errors"
	"chromiumos/tast/testing"
)

// LaunchServer installs vendor keys and relaunches the local ADB sever.
// The server must be relaunched to load the vendor keys.
func LaunchServer(ctx context.Context) error {
	testing.ContextLog(ctx, "Installing ADB vendor keys")
	if err := adb.InstallVendorKeys(); err != nil {
		return err
	}

	testing.ContextLog(ctx, "Killing existing ADB server process(es)")
	if err := adb.KillADBLocalServer(ctx); err != nil {
		return errors.Wrap(err, "failed to kill ADB local server")
	}

	testing.ContextLog(ctx, "Starting ADB server")
	if err := adb.Command(ctx, "start-server").Run(testexec.DumpLogOnError); err != nil {
		return errors.Wrap(err, "failed starting ADB local server")
	}
	return nil
}
