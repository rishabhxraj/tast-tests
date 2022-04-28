// Copyright 2018 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package graphics

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"chromiumos/tast/errors"
)

// deqpParser holds state to parse log files generated by DEQP tests. This is
// used by the ParseDEQPOutput function.
type deqpParser struct {
	stats     map[string]uint // accumulates statistics as "outcome": # tests
	nonFailed []string        // accumulates the non-failed tests
	line      string          // current line of input being read
	lastLine  bool            // true if line is the last line of input
	start     bool            // true at the start of a test case
	rawXML    strings.Builder // accumulates XML code for the current test case
	test      string          // name of the current test case
	outcome   string          // outcome of the current test case
	complete  bool            // true when rawXML has complete XML code
	bad       bool            // true when a recoverable parsing error occurs
}

// Markers that we need to look for in the DEQP log files.
const (
	beginResult = "#beginTestCaseResult"
	endResult   = "#endTestCaseResult"
	termResult  = "#terminateTestCaseResult"
)

// Set of outcomes not considered to be a failure. This is a port of the
// TEST_RESULT_FILTER list in
// autotest/files/client/site_tests/graphics_dEQP/graphics_dEQP.py.
var nonFailOutcomes = map[string]struct{}{
	"pass":                 {},
	"notsupported":         {},
	"internalerror":        {},
	"qualitywarning":       {},
	"compatibilitywarning": {},
	"skipped":              {},
}

// handleBeginResult handles a line starting with #beginTestCaseResult.
func (d *deqpParser) handleBeginResult() error {
	// If we see another begin before an end/terminate then something is wrong.
	if d.start {
		return errors.Errorf("unexpected %v", beginResult)
	}

	// Derive the test name from #beginTestCaseResult.
	if s := strings.TrimSpace(strings.TrimPrefix(d.line, beginResult)); len(s) > 0 {
		d.test = s
		d.start = true
		return nil
	}
	return errors.Errorf("%v is not followed by test name", beginResult)
}

// handleEndResult handles a line starting with #endTestCaseResult.
func (d *deqpParser) handleEndResult() error {
	// We shouldn't see an end without a begin.
	if !d.start {
		return errors.Errorf("unexpected %v", endResult)
	}
	d.complete = true
	return nil
}

// handleTermResult handles a line starting with #terminateTestCaseResult. This
// is the case when the test terminated early. The XML could be incomplete and
// should not be parsed.
func (d *deqpParser) handleTermResult() error {
	// We shouldn't see a terminate without a begin.
	if !d.start {
		return errors.Errorf("unexpected %v", termResult)
	}

	// Get the cause for early termination.
	if s := strings.TrimSpace(strings.TrimPrefix(d.line, termResult)); len(s) > 0 {
		d.outcome = s
		d.bad = true
		return nil
	}

	// #terminateTestCaseResult is not accompanied by a cause. If this is the
	// last line, let's assume that Tast killed the test due to a timeout and
	// make this error recoverable. Otherwise, report an unrecoverable error.
	if d.lastLine {
		d.outcome = "parsefailed"
		d.bad = true
		return nil
	}
	return errors.Errorf("missing cause for %v", termResult)
}

// handleXML parses the XML present in d.rawXML and performs some smoke checks.
func (d *deqpParser) handleXML() error {
	// Structure to parse the XML into. Note that it's necessary to capitalize
	// the first letter of each field so that xml.Unmarshal works.
	r := struct {
		XMLName  xml.Name `xml:"TestCaseResult"`
		CasePath string   `xml:",attr"`
		Result   []struct {
			StatusCode string `xml:",attr"`
		}
	}{}

	// Parse and perform smoke checks.
	if err := xml.Unmarshal([]byte(d.rawXML.String()), &r); err != nil {
		return err
	}
	if r.CasePath != d.test {
		return errors.Errorf("bad CasePath: %q", r.CasePath)
	}
	if len(r.Result) != 1 {
		return errors.Errorf("%v <Result> elements found", len(r.Result))
	}
	s := strings.TrimSpace(r.Result[0].StatusCode)
	if len(s) == 0 {
		return errors.Errorf("bad StatusCode: %q", s)
	}
	d.outcome = s
	return nil
}

// handleRemainingInput is expected to be called after going through all the
// lines in the file. If by that time, d.start == true, this method produces a
// recoverable error for the test that never ended. This can happen, e.g., if
// Tast killed the test due to a timeout or a crash occurred.
func (d *deqpParser) handleRemainingInput() {
	if d.start {
		d.outcome = "parsefailed"
		d.collectOutcome()
	}
}

// collectOutcome adds the current test to the d.nonFailed slice if it's
// considered to not have failed. It also counts it in the d.stats map.
func (d *deqpParser) collectOutcome() {
	if !DEQPOutcomeIsFailure(d.outcome) {
		d.nonFailed = append(d.nonFailed, d.test)
	}
	d.stats[strings.ToLower(d.outcome)]++
}

// prepareForNextTestCase resets the state necessary to move onto the next test
// case.
func (d *deqpParser) prepareForNextTestCase() {
	d.start = false
	d.rawXML.Reset()
	d.test = ""
	d.outcome = ""
	d.complete = false
	d.bad = false
}

// parse is the most important method for a deqpParser. It reads the passed file
// line by line to parse it and eventually produce a d.stats map which counts
// the number of tests for each outcome and a d.nonFailed slice that contains
// the names of the tests considered to not have failed.
//
// This is a (hopefully improved) port of the functionality of the parsing
// function defined in
// autotest/files/client/site_tests/graphics_dEQP/graphics_dEQP.py: the version
// here tends to be more conservative with parsing errors since they could
// indicate a problem with the DEQP output.
//
// TODO(andrescj): another possible error check is to make sure a test doesn't
// appear twice.
func (d *deqpParser) parse(p string) error {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, l := range lines {
		d.line = l
		d.lastLine = (i == len(lines)-1)

		switch {
		case strings.HasPrefix(d.line, beginResult):
			if err := d.handleBeginResult(); err != nil {
				return errors.Wrapf(err, "line %v", i+1)
			}
		case strings.HasPrefix(d.line, endResult):
			if err := d.handleEndResult(); err != nil {
				return errors.Wrapf(err, "line %v", i+1)
			}
		case strings.HasPrefix(d.line, termResult):
			if err := d.handleTermResult(); err != nil {
				return errors.Wrapf(err, "line %v", i+1)
			}
		case d.start:
			// We're currently collecting XML. Note that we don't need to add a
			// newline after each read line: the XML parser is ok with that.
			d.rawXML.WriteString(d.line)
		}

		// Check if we're done with the test case.
		if d.complete || d.bad {
			if d.complete {
				if err := d.handleXML(); err != nil {
					return errors.Wrapf(err, "test %v", d.test)
				}
			}
			d.collectOutcome()
			d.prepareForNextTestCase()
		}
	}

	d.handleRemainingInput()
	return nil
}

// DEQPOutcomeIsFailure decides if an outcome found in the output of a DEQP test
// is considered a failure.
func DEQPOutcomeIsFailure(s string) bool {
	_, isNonFail := nonFailOutcomes[strings.ToLower(s)]
	return !isNonFail
}

// ParseDEQPOutput parses the given DEQP log file to extract the number of tests
// per outcome (returned in the stats map) and the names of the non-failed tests
// (returned in the nonFailed slice). An error is returned if an unrecoverable
// error occurs, i.e., an error that suggests problems with the DEQP output.
//
// The returned stats map might look something like
//	"pass": 3
//	"fail": 1
//
// This means that 3 tests passed and 1 failed. The keys are always lowercase.
// When a recoverable error occurs, it is reported in the stats map with the
// reserved outcome "parsefailed".
//
// This parser expects the format explained in
// https://android.googlesource.com/platform/external/deqp/+/deqp-dev/doc/qpa_file_format.txt
// but only cares about the #beginTestCaseResult ... #endTestCaseResult or
// #beginTestCaseResult ... #terminateTestCaseResult sections.
func ParseDEQPOutput(p string) (stats map[string]uint, nonFailed []string, err error) {
	d := deqpParser{stats: make(map[string]uint)}
	if err := d.parse(p); err != nil {
		return nil, nil, err
	}
	return d.stats, d.nonFailed, nil
}
