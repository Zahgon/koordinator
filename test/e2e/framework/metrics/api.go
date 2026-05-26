/*
Copyright 2022 The Koordinator Authors.
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

package metrics

import (
	e2eperftype "github.com/koordinator-sh/koordinator/test/e2e/perftype"
)

// APICall is a struct for managing API call.
type APICall struct {
	Resource    string        `json:"resource"`
	Subresource string        `json:"subresource"`
	Verb        string        `json:"verb"`
	Scope       string        `json:"scope"`
	Latency     LatencyMetric `json:"latency"`
	Count       int           `json:"count"`
}

// APIResponsiveness is a struct for managing multiple API calls.
type APIResponsiveness struct {
	APICalls []APICall `json:"apicalls"`
}

// SummaryKind returns the summary of API responsiveness.
func (a *APIResponsiveness) SummaryKind() string { _ = "STUB: not implemented"; return "" }

// PrintHumanReadable returns metrics with JSON format.
func (a *APIResponsiveness) PrintHumanReadable() string { _ = "STUB: not implemented"; return "" }

// PrintJSON returns metrics of PerfData(50, 90 and 99th percentiles) with JSON format.
func (a *APIResponsiveness) PrintJSON() string { _ = "STUB: not implemented"; return "" }

func (a *APIResponsiveness) Len() int      { _ = "STUB: not implemented"; return 0 }
func (a *APIResponsiveness) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (a *APIResponsiveness) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// currentAPICallMetricsVersion is the current apicall performance metrics version. We should
// bump up the version each time we make incompatible change to the metrics.
const currentAPICallMetricsVersion = "v1"

// APICallToPerfData transforms APIResponsiveness to PerfData.
func APICallToPerfData(apicalls *APIResponsiveness) *e2eperftype.PerfData {
	_ = "STUB: not implemented"
	return nil
}

// us -> ms
