// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is generated by `tools/generate_automation2_constants.py`.

// Package state describes characteristics of a chrome.automation AutomationNode.
package state

// State describes characteristics of a chrome.automation AutomationNode.
type State string

// As defined in https://chromium.googlesource.com/chromium/src/+/refs/heads/main/extensions/common/api/automation.idl
const (
	AutofillAvailable State = "autofillAvailable"
	Collapsed         State = "collapsed"
	Default           State = "default"
	Editable          State = "editable"
	Expanded          State = "expanded"
	Focusable         State = "focusable"
	Focused           State = "focused"
	Horizontal        State = "horizontal"
	Hovered           State = "hovered"
	Ignored           State = "ignored"
	Invisible         State = "invisible"
	Linked            State = "linked"
	Multiline         State = "multiline"
	Multiselectable   State = "multiselectable"
	Offscreen         State = "offscreen"
	Protected         State = "protected"
	Required          State = "required"
	RichlyEditable    State = "richlyEditable"
	Vertical          State = "vertical"
	Visited           State = "visited"
)
