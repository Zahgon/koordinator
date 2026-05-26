/*
Copyright 2022 The Koordinator Authors.

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

package workloadauditor

import (
	"time"

	"github.com/spf13/pflag"
)

// WorkloadAuditorConfig holds configuration for the workloadAuditorImpl.
var (
	WorkloadAuditorEnabled      = false
	WorkloadAuditorMetricLabels = "" // comma-separated metricLabel=podLabelKey pairs

	// Anomaly detection thresholds
	VictimRescheduleDuration = 10 * time.Second
	VictimDeletionDuration   = 30 * time.Second
	VictimDeletingRetries    = 3
	SchedulingEventInterval  = 5 * time.Minute
)

// AddFlags registers the workloadAuditorImpl command-line flags.
func AddFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type WorkloadAuditorConfig struct {
	Enabled                  bool
	VictimRescheduleDuration time.Duration
	VictimDeletionDuration   time.Duration
	VictimDeletingRetries    int
	SchedulingEventInterval  time.Duration
	MetricLabelNames         []string // ["priority", "gpu", "quota_name", ...]
	PodLabelKeys             []string // parallel to MetricLabelNames[1:]
}

// parseMetricLabels parses the comma-separated metricLabel=podLabelKey pairs.
// Returns the full label name list (always starting with "priority") and the
// parallel pod-label-key list (for custom labels only).
func parseMetricLabels(raw string) (names []string, podKeys []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DefaultWorkloadAuditorConfig returns the configuration built from package-level vars.
func DefaultWorkloadAuditorConfig() WorkloadAuditorConfig {
	_ = "STUB: not implemented"
	return *new(WorkloadAuditorConfig)
}
