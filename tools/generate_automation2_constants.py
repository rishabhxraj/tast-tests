#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# Copyright 2021 The Chromium OS Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""Tool for generating go code from automation.js

Used to create src/chromiumos/tast/local/chrome/ui2/nodewith/*/constants.go.

Usage example:
# ./generate_automation_constants.py \
  /path/to/chromium/src/third_party/closure_compiler/externs/automation.js \
  /path/to/chromeos/src/platform/tast-tests/src/chromiumos/tast/local/chrome/ui2/

Make sure to apply gofmt to the output of this script.

TODO(hirokisato): Currently this script reads closure compiler's definition js
file, because it's easy to parse. But the automation IDL file is the source of
truth.
"""

import os
import sys
import re

HEADER = """\
// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This file is generated by `tools/generate_automation2_constants.py`.

"""


def to_camel_case(snake_case_str):
  """Converts a SNAKE_CASE string into a CamelCase string."""
  return ''.join(s.lower().title() for s in snake_case_str.split('_'))


def generate_definitions(
    lines, defined_type_name, go_type_name, description, output_none=False):
  """Generates Golang type definitions.

  Args:
    lines: List of strings that is read from a js file.
    defined_type_name: Enum type name in a source file which is parsed here.
    go_type_name: Enum type name to be generated as Golang definition.
    description: A string that is used to generate a Go doc comment.
    output_none: Whether or not to output an empty string as the "None" value.
  """
  item_pattern = re.compile(r'\s*(\w*):\s\'(\w*)\'')
  defs = []
  reading = False
  for l in lines:
    if reading:
      match = item_pattern.match(l)
      if not match:
        reading = False
        break
      defs.append((to_camel_case(match.group(1)), match.group(2)))
    elif l.startswith('chrome.automation.%s' % defined_type_name):
      reading = True

  out = '// %s describes %s.\n' % (go_type_name, description)
  out += 'type %s string\n\n' % go_type_name
  out += '// As defined in https://chromium.googlesource.com/chromium/src/+/refs/heads/main/extensions/common/api/automation.idl\n'
  out += 'const (\n'

  for r in defs:
    out += '\t%s %s = "%s"\n' % (r[0], go_type_name, r[1])
  if output_none:
    out += '\t%s %s = ""\n' % ("None", go_type_name)
  out += ')\n'
  return out


def generate_package(
    lines,
    parent_folder,
    defined_type_name,
    go_type_name,
    description,
    output_none=False):
  """Generates a Golang package for the constants.

  Args:
    lines: List of strings that is read from a js file.
    parent_folder: A string with the path to the folder to generate the package in.
    defined_type_name: Enum type name in a source file which is parsed here.
    go_type_name: Enum type name to be generated as Golang definition.
    description: A string that is used to generate a Go doc comment.
    output_none: Whether or not to output an empty string as the "None" value.
  """
  pkg_name = go_type_name.lower()
  pkg_folder = os.path.join(parent_folder, pkg_name)
  if not os.path.exists(pkg_folder):
    os.makedirs(pkg_folder)

  out = HEADER
  out += '// Package %s describes %s.\n' % (pkg_name, description)
  out += 'package %s\n\n' % pkg_name
  out += generate_definitions(
      lines, defined_type_name, go_type_name, description, output_none)
  f = open(os.path.join(pkg_folder, 'constants.go'), 'w')
  f.write(out)
  f.close()


def main(argv):
  f = open(argv[1], 'r')
  lines = f.readlines()
  f.close()

  folder = argv[2]
  generate_package(
      lines, folder, 'StateType', 'State',
      'characteristics of a chrome.automation AutomationNode')
  generate_package(
      lines, folder, 'RoleType', 'Role',
      'the purpose of a chrome.automation AutomationNode')
  generate_package(
      lines, folder, 'EventType', 'Event',
      'the type of a chrome.automation AutomationEvent')
  generate_package(
      lines, folder, 'Restriction', 'Restriction',
      'the restriction state of a chrome.automation AutomationNode', True)


if __name__ == '__main__':
  main(sys.argv)