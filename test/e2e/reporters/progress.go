/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reporters

import (
	"net/http"

	"github.com/onsi/ginkgo/v2/config"
	"github.com/onsi/ginkgo/v2/types"
)

// ProgressReporter is a ginkgo reporter which tracks the total number of tests to be run/passed/failed/skipped.
// As new tests are completed it updates the values and prints them to stdout and optionally, sends the updates
// to the configured URL.
type ProgressReporter struct {
	LastMsg string `json:"msg"`

	TestsTotal     int `json:"total"`
	TestsCompleted int `json:"completed"`
	TestsSkipped   int `json:"skipped"`
	TestsFailed    int `json:"failed"`

	Failures []string `json:"failures,omitempty"`

	progressURL string
	client      *http.Client
}

// NewProgressReporter returns a progress reporter which posts updates to the given URL.
func NewProgressReporter(progressReportURL string) *ProgressReporter {
	_ = "STUB: not implemented"
	return nil
}

// SpecSuiteWillBegin is invoked by ginkgo when the suite is about to start and is the first point in which we can
// antipate the number of tests which will be run.
func (reporter *ProgressReporter) SpecSuiteWillBegin(cfg config.GinkgoConfigType, summary *types.SuiteSummary) {
	_ = "STUB: not implemented"
	return
}

// SpecSuiteDidEnd is the last method invoked by Ginkgo after all the specs are run.
func (reporter *ProgressReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	_ = "STUB: not implemented"
	return
}

// SpecDidComplete is invoked by Ginkgo each time a spec is completed (including skipped specs).
func (reporter *ProgressReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	_ = "STUB: not implemented"
	return
}

// sendUpdates serializes the current progress and prints it to stdout and also posts it to the configured endpoint if set.
func (reporter *ProgressReporter) sendUpdates() { _ = "STUB: not implemented"; return }

func (reporter *ProgressReporter) postProgressToURL(b []byte) {
	_ = "STUB: not implemented"
	// If a progressURL and client is set/available then POST to it. Noop otherwise.
	return
}

func (reporter *ProgressReporter) serialize() []byte { _ = "STUB: not implemented"; return nil }

// SpecWillRun is implemented as a noop to satisfy the reporter interface for ginkgo.
func (reporter *ProgressReporter) SpecWillRun(specSummary *types.SpecSummary) {
	_ = "STUB: not implemented"

	// BeforeSuiteDidRun is implemented as a noop to satisfy the reporter interface for ginkgo.
	return
}

func (reporter *ProgressReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	_ = "STUB: not implemented"

	// AfterSuiteDidRun is implemented as a noop to satisfy the reporter interface for ginkgo.
	return
}

func (reporter *ProgressReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	_ = "STUB: not implemented"
	return
}
