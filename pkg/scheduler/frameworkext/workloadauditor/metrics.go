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
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/component-base/metrics"
)

const (
	outcomeScheduled     = "scheduled"
	outcomeDeleted       = "deleted"
	outcomeGangScheduled = "gangScheduled"
	outcomeGangDeleted   = "gangDeleted"
)

var (
	initOnce             sync.Once
	configuredLabelNames sets.Set[string]

	// RepeatedPreemptionTotal repeated preemption counter.
	RepeatedPreemptionTotal *metrics.CounterVec

	// PreemptionInvalidationsTotal preemption invalidation counter.
	PreemptionInvalidationsTotal *metrics.CounterVec

	// VictimRescheduleDurationSeconds histogram of time from victimAllDeleted to next scheduling result.
	VictimRescheduleDurationSeconds *metrics.HistogramVec

	// PreemptionVictimDeletingRetries histogram of preemptVictimDeleting retry count per preemption cycle.
	PreemptionVictimDeletingRetries *metrics.HistogramVec

	// PreemptionToVictimDeletedSeconds histogram of time from preemptNominated to victimAllDeleted (normal cycle completion).
	PreemptionToVictimDeletedSeconds *metrics.HistogramVec

	// PreemptionCycleInterruptedSeconds histogram of time from preemptNominated to cycle interruption
	// (next preemptNominated, scheduled, preemptFailure, scheduleFailure, or workload deletion).
	// These durations may be affected by scheduling backlog and are less accurate.
	PreemptionCycleInterruptedSeconds *metrics.HistogramVec

	// SchedulingEventsBeforeOutcome histogram of scheduling event counts before workload outcome.
	SchedulingEventsBeforeOutcome *metrics.HistogramVec

	// SchedulingEventIntervalSeconds histogram of time between consecutive scheduling events
	// within a single dequeue-attempt round.
	SchedulingEventIntervalSeconds *metrics.HistogramVec

	// RecordMethodDurationSeconds histogram of RecordAttemptPod / RecordDiagnosis call duration.
	RecordMethodDurationSeconds *metrics.HistogramVec
)

// InitMetrics creates and registers all workload auditor metrics with the given label names.
// Must be called once after flag parsing (typically from NewWorkloadAuditor).
func InitMetrics(labelNames []string) { _ = "STUB: not implemented"; return }

// 0.1s … ~52428s (~14.5h)

// 1µs … ~0.26s

// DeleteMetricsByLabel deletes all metric time series where labelName equals labelValue.
// This is a no-op if labelName is not among the configured metric labels or metrics are not initialized.
func DeleteMetricsByLabel(labelName, labelValue string) { _ = "STUB: not implemented"; return }
