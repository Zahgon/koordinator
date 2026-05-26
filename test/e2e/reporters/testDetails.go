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
	"io"

	"github.com/onsi/ginkgo/v2/config"
	"github.com/onsi/ginkgo/v2/types"
)

// DetailsReporter is a ginkgo reporter which dumps information regarding the tests which is difficult to get
// via an AST app. This allows us to leverage the existing ginkgo logic to walk the tests and such while then following
// up with an custom app which leverages AST to generate conformance documentation.
type DetailsReporter struct {
	Writer io.Writer
}

// NewDetailsReporterWithWriter returns a reporter which will write the SpecSummary objects as tests
// complete to the given writer.
func NewDetailsReporterWithWriter(w io.Writer) *DetailsReporter {
	_ = "STUB: not implemented"
	return nil
}

// NewDetailsReporterFile returns a reporter which will create the file given and dump the specs
// to it as they complete.
func NewDetailsReporterFile(filename string) *DetailsReporter {
	_ = "STUB: not implemented"
	return nil
}

// SpecSuiteWillBegin is implemented as a noop to satisfy the reporter interface for ginkgo.
func (reporter *DetailsReporter) SpecSuiteWillBegin(cfg config.GinkgoConfigType, summary *types.SuiteSummary) {
	_ = "STUB: not implemented"

	// SpecSuiteDidEnd is implemented as a noop to satisfy the reporter interface for ginkgo.
	return
}

func (reporter *DetailsReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	_ = "STUB: not implemented"

	// SpecDidComplete is invoked by Ginkgo each time a spec is completed (including skipped specs).
	return
}

func (reporter *DetailsReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	_ = "STUB: not implemented"
	return
}

// Printing newline between records for easier viewing in various tools.

// SpecWillRun is implemented as a noop to satisfy the reporter interface for ginkgo.
func (reporter *DetailsReporter) SpecWillRun(specSummary *types.SpecSummary) {
	_ = "STUB: not implemented"

	// BeforeSuiteDidRun is implemented as a noop to satisfy the reporter interface for ginkgo.
	return
}

func (reporter *DetailsReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	_ = "STUB: not implemented"

	// AfterSuiteDidRun is implemented as a noop to satisfy the reporter interface for ginkgo.
	return
}

func (reporter *DetailsReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	_ = "STUB: not implemented"
	return
}
