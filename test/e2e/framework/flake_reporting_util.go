/*
Copyright 2018 The Kubernetes Authors.

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

package framework

import (
	"sync"
)

// FlakeReport is a struct for managing the flake report.
type FlakeReport struct {
	lock       sync.RWMutex
	Flakes     []string `json:"flakes"`
	FlakeCount int      `json:"flakeCount"`
}

// NewFlakeReport returns a new flake report.
func NewFlakeReport() *FlakeReport { _ = "STUB: not implemented"; return nil }

func buildDescription(optionalDescription ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// RecordFlakeIfError records the error (if non-nil) as a flake along with an optional description.
// This can be used as a replacement of framework.ExpectNoError() for non-critical errors that can
// be considered as 'flakes' to avoid causing failures in tests.
func (f *FlakeReport) RecordFlakeIfError(err error, optionalDescription ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetFlakeCount returns the flake count.
func (f *FlakeReport) GetFlakeCount() int { _ = "STUB: not implemented"; return 0 }

// PrintHumanReadable returns string of flake report.
func (f *FlakeReport) PrintHumanReadable() string { _ = "STUB: not implemented"; return "" }

// PrintJSON returns the summary of frake report with JSON format.
func (f *FlakeReport) PrintJSON() string { _ = "STUB: not implemented"; return "" }

// SummaryKind returns the summary of flake report.
func (f *FlakeReport) SummaryKind() string { _ = "STUB: not implemented"; return "" }
