// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package crostini

import (
	"context"
	"fmt"
	"time"

	"chromiumos/tast/ctxutil"
	"chromiumos/tast/errors"
	"chromiumos/tast/local/chrome"
	"chromiumos/tast/local/chrome/uiauto"
	"chromiumos/tast/local/chrome/uiauto/nodewith"
	"chromiumos/tast/local/chrome/uiauto/restriction"
	"chromiumos/tast/local/chrome/uiauto/role"
	"chromiumos/tast/local/crostini"
	cui "chromiumos/tast/local/crostini/ui"
	"chromiumos/tast/local/crostini/ui/settings"
	"chromiumos/tast/local/crostini/ui/terminalapp"
	"chromiumos/tast/local/input"
	"chromiumos/tast/testing"
)

/*
The rules for valid users are:
-- Less than 32 characters, so if the default user name is more than 32 characters, it should be truncated.
   For example, if the default username is abcdefghijklmnopqrstuvwxyzabcdefg, then it will be truncated as abcdefghijklmnopqrstuvwxyzabcdef.
-- Not one of the reserved usernames: refer to var reservedUsers.
	Please also refer container default usernames to check whether there is update.
-- The username must start with a lowercase letter or underscore. some invalid examples: 9user, Tus.
-- The username must only contain lowercase letters, digits, underscores, and hyphens. Some invalid examples: uTfdsa, us9)x.
*/

var validUsersMap = map[string][]string{
	"users_start_with_letter":     {"a", "tester00", "whoisthis", "myuser"},
	"users_start_with_underscore": {"_", "_tast", "_9", "__"},
	"users_contain_reserved_user": {"daemonvsdaemon", "binbin", "ssys", "synccnys",
		"gamesmoregames", "manmanmen", "lplplp", "mailmail9", "newss", "uucpcp",
		"proxyfail", "www-datawww-", "backupbackup", "listli", "ircir",
		"gnatsss", "nobodynobodybutyou", "_aptapt", "systemd-timesyncsystemd-timesync",
		"_systemd-network", "systemd5-resolve",
		"systemd-bus-proxyx", "messagebus-", "sshdfas", "rtkit99",
		"pul--se", "android-rootroot", "chronos-access_", "android-everybodybody", "rootroot"},
	"long_users": {"________________________________",
		"rootrootrootrootrootrootrootroot",
		"_-------------------------------",
		"_123456789012345678901234567890-",
		"t-9_somethingelse---48395-______"},
}

func init() {
	testing.AddTest(&testing.Test{
		Func:         CheckValidUsers,
		LacrosStatus: testing.LacrosVariantUnneeded,
		Desc:         "Test the installation could proceed with valid users",
		Contacts:     []string{"jinrongwu@google.com", "cros-containers-dev@google.com"},
		Attr:         []string{"group:mainline", "informational"},
		SoftwareDeps: []string{"chrome", "vm_host"},
		Params: []testing.Param{
			// Parameters generated by check_valid_users_test.go. DO NOT EDIT.
			{
				Name:              "stable",
				ExtraData:         []string{crostini.GetContainerMetadataArtifact("buster", false), crostini.GetContainerRootfsArtifact("buster", false)},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniStable,
				Timeout:           20 * time.Minute,
			}, {
				Name:              "unstable",
				ExtraAttr:         []string{"informational"},
				ExtraData:         []string{crostini.GetContainerMetadataArtifact("buster", false), crostini.GetContainerRootfsArtifact("buster", false)},
				ExtraSoftwareDeps: []string{"dlc"},
				ExtraHardwareDeps: crostini.CrostiniUnstable,
				Timeout:           20 * time.Minute,
			},
		},
	})
}

func CheckValidUsers(ctx context.Context, s *testing.State) {
	cr, err := chrome.New(ctx)
	if err != nil {
		s.Fatal("Failed to start Chrome: ", err)
	}
	tconn, err := cr.TestAPIConn(ctx)
	if err != nil {
		s.Fatal("Failed to create Test API connection: ", err)
	}

	kb, err := input.Keyboard(ctx)
	if err != nil {
		s.Fatal("Failed to get keyboard: ", err)
	}
	defer kb.Close()

	// Use a shortened context for test operations to reserve time for cleanup.
	cleanupCtx := ctx
	ctx, cancel := ctxutil.Shorten(ctx, crostini.PostTimeout)
	defer cancel()
	defer crostini.RunCrostiniPostTest(cleanupCtx, crostini.PreData{Chrome: cr, TestAPIConn: tconn, Container: nil, Keyboard: kb, Post: &crostini.PostTestData{}})

	for k, users := range validUsersMap {
		testName := fmt.Sprintf("test_%s", k)
		s.Run(ctx, testName, func(ctx context.Context, s *testing.State) {
			cr, err = chrome.New(ctx)
			if err != nil {
				s.Fatal("Failed to start Chrome: ", err)
			}
			tconn, err = cr.TestAPIConn(ctx)
			if err != nil {
				s.Fatal("Failed to creat test API connection: ", err)
			}

			if err := checkUsersAndInstall(ctx, tconn, cr, kb, users); err != nil {
				s.Fatal("Failed to check user and install crostini: ", err)
			}
		})
	}

}

// checkUsersAndInstall checks a group of valid users,
// and makes sure that there is no error msg, and button Install is enabled.
// It also goes through the installation process with the last username in users.
func checkUsersAndInstall(ctx context.Context, tconn *chrome.TestConn, cr *chrome.Chrome, kb *input.KeyboardEventWriter, users []string) error {
	if err := settings.OpenInstaller(ctx, tconn, cr); err != nil {
		return errors.Wrap(err, "failed to open Crostini installer")
	}

	ui := uiauto.New(tconn)

	userNameField := nodewith.Name("Username").Role(role.TextField).Ancestor(cui.InstallWindow)
	installButton := nodewith.Name("Install").Role(role.Button)

	// alertNode is the node for error message when the input user name is invalid.
	// The node does not exist if the user name is valid.
	alertNode := nodewith.Role(role.Alert).Ancestor(cui.InstallWindow)

	for _, userName := range users {
		typedUserName := nodewith.Name(userName).Role(role.StaticText).Ancestor(userNameField)
		if err := uiauto.Combine("check user name",
			ui.LeftClick(userNameField),
			kb.AccelAction("Ctrl+A"),
			kb.AccelAction("Backspace"),
			kb.TypeAction(userName),
			ui.WaitUntilExists(typedUserName),
			ui.Gone(alertNode),
			ui.CheckRestriction(installButton, restriction.None))(ctx); err != nil {
			return errors.Wrapf(err, "failed to check user %s", userName)
		}

	}

	// Install with the last user.
	if err := uiauto.Combine("",
		ui.LeftClick(installButton),
		ui.WithTimeout(8*time.Minute).WaitUntilGone(cui.InstallWindow))(ctx); err != nil {
		return errors.Wrapf(err, "failed to install Crostini with user %q", users[len(users)-1])
	}

	// Find Terminal window.
	terminalApp, err := terminalapp.Find(ctx, tconn)
	if err != nil {
		return errors.Wrap(err, "failed to find terminal after installing Crostini")
	}
	defer terminalApp.Close()(ctx)

	// Check user name in the Terminal app.
	terminalWindow := nodewith.Name(users[len(users)-1] + "@penguin: ~").Role(role.RootWebArea)
	if err := ui.WaitUntilExists(terminalWindow)(ctx); err != nil {
		return errors.Wrapf(err, "failed to find username %s in terminal", users[len(users)-1])
	}

	return nil
}
