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

const (
	// Cluster Autoscaler metrics names
	caFunctionMetric      = "cluster_autoscaler_function_duration_seconds_bucket"
	caFunctionMetricLabel = "function"
)

// ComponentCollection is metrics collection of components.
type ComponentCollection Collection

func (m *ComponentCollection) filterMetrics() { _ = "STUB: not implemented"; return }

// PrintHumanReadable returns e2e metrics with JSON format.
func (m *ComponentCollection) PrintHumanReadable() string { _ = "STUB: not implemented"; return "" }

// PrettyPrintJSON converts metrics to JSON format.
// TODO: This function should be replaced with framework.PrettyPrintJSON after solving
// circulary dependency between core framework and this metrics subpackage.
func PrettyPrintJSON(metrics interface{}) string { _ = "STUB: not implemented"; return "" }

// PrintJSON returns e2e metrics with JSON format.
func (m *ComponentCollection) PrintJSON() string { _ = "STUB: not implemented"; return "" }

// SummaryKind returns the summary of e2e metrics.
func (m *ComponentCollection) SummaryKind() string { _ = "STUB: not implemented"; return "" }

// ComputeClusterAutoscalerMetricsDelta computes the change in cluster
// autoscaler metrics.
func (m *ComponentCollection) ComputeClusterAutoscalerMetricsDelta(before Collection) {
	_ = "STUB: not implemented"
	return
}
