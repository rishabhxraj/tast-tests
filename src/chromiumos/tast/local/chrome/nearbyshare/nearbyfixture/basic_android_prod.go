// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package nearbyfixture

import (
	"time"

	nearbycommon "chromiumos/tast/common/cros/nearbyshare"
	"chromiumos/tast/local/chrome/nearbyshare/nearbysnippet"
	"chromiumos/tast/testing"
)

// addProdAndroidFixtures registers fixtures for basic CrOS<->Android sharing tests.
// The Android phone will be signed in with a GAIA account that is not a member of any Nearby user groups,
// so it will use the production version of Nearby Share.
func addProdAndroidFixtures() {
	// Fixtures for in-contacts sharing tests.
	testing.AddFixture(&testing.Fixture{
		Name: "nearbyShareDataUsageOfflineAllContactsProd",
		Desc: "Nearby Share enabled on CrOS and Android configured with 'Data Usage' set to 'Offline',  'Visibility' set to 'All Contacts'",
		Impl: NewNearbyShareFixture(fixtureOptions{
			crosDataUsage:              nearbycommon.DataUsageOffline,
			crosVisibility:             nearbycommon.VisibilityAllContacts,
			androidDataUsage:           nearbysnippet.DataUsageOffline,
			androidVisibility:          nearbysnippet.VisibilityAllContacts,
			crosSelectAndroidAsContact: false,
		}),
		Contacts: []string{
			"chromeos-sw-engprod@google.com",
		},
		Parent:          "nearbyShareGAIALoginAndroidAccountProd",
		SetUpTimeout:    2 * time.Minute,
		ResetTimeout:    resetTimeout,
		TearDownTimeout: resetTimeout,
		PreTestTimeout:  resetTimeout,
		PostTestTimeout: resetTimeout,
	})

	// Fixtures for high-visibility sharing tests.
	testing.AddFixture(&testing.Fixture{
		Name:   "nearbyShareDataUsageOnlineNoOneProd",
		Desc:   "Nearby Share enabled on CrOS and Android configured with 'Data Usage' set to 'Online' and 'Visibility' set to 'No One'",
		Parent: "nearbyShareGAIALoginProd",
		Impl: NewNearbyShareFixture(fixtureOptions{
			crosDataUsage:              nearbycommon.DataUsageOnline,
			crosVisibility:             nearbycommon.VisibilityNoOne,
			androidDataUsage:           nearbysnippet.DataUsageOnline,
			androidVisibility:          nearbysnippet.VisibilityNoOne,
			crosSelectAndroidAsContact: false,
		}),
		Contacts: []string{
			"chromeos-sw-engprod@google.com",
		},
		SetUpTimeout:    2 * time.Minute,
		ResetTimeout:    resetTimeout,
		TearDownTimeout: resetTimeout,
		PreTestTimeout:  resetTimeout,
		PostTestTimeout: resetTimeout,
	})

	testing.AddFixture(&testing.Fixture{
		Name: "nearbyShareDataUsageOfflineNoOneProd",
		Desc: "Nearby Share enabled on CrOS and Android configured with 'Data Usage' set to 'Offline' and 'Visibility' set to 'No One'",
		Impl: NewNearbyShareFixture(fixtureOptions{
			crosDataUsage:              nearbycommon.DataUsageOffline,
			crosVisibility:             nearbycommon.VisibilityNoOne,
			androidDataUsage:           nearbysnippet.DataUsageOffline,
			androidVisibility:          nearbysnippet.VisibilityNoOne,
			crosSelectAndroidAsContact: false,
		}),
		Contacts: []string{
			"chromeos-sw-engprod@google.com",
		},
		Parent:          "nearbyShareGAIALoginProd",
		SetUpTimeout:    2 * time.Minute,
		ResetTimeout:    resetTimeout,
		TearDownTimeout: resetTimeout,
		PreTestTimeout:  resetTimeout,
		PostTestTimeout: resetTimeout,
	})
}
