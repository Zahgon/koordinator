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

	corev1 "k8s.io/api/core/v1"
)

// PodIsGated determines whether a pod is considered gated.
// Other packages may override this function to customize the gating check.
var PodIsGated = func(pod *corev1.Pod) bool {
	return len(pod.Spec.SchedulingGates) > 0
}

type WorkloadAuditor interface {
	Enabled() bool
	AddPod(pod *corev1.Pod)
	DeletePod(pod *corev1.Pod)
	RecordPod(pod *corev1.Pod, recordType RecordType, message string)
	RecordPodGating(pod *corev1.Pod, gated bool)
	RecordAttemptPod(pod *corev1.Pod)
	RecordPodScheduleResult(pod *corev1.Pod, recordType RecordType, message string)

	AddGangGroup(gangGroupID string)
	DeleteGangGroup(gangGroupID string)
	RecordGangGroup(gangGroupID string, pod *corev1.Pod, recordType RecordType, message string)
	RecordGangGating(gangGroupID string, pod *corev1.Pod, gated bool)
	RecordGangScheduleResult(gangGroupID string, recordType RecordType, message string)

	RecordDiagnosis(pod *corev1.Pod, questionKey string, recordType RecordType, message string)
}

// workloadAuditorImpl tracks the scheduling lifecycle of workloads.
// It is a singleton at the FrameworkExtenderFactory level.
// records is a sync.Map for lock-free reads; each WorkloadRecord has its own mu
// for field-level operations, so different workloads never contend.
type workloadAuditorImpl struct {
	records sync.Map // workloadKey -> *WorkloadRecord
	Config  WorkloadAuditorConfig
}

func (w *workloadAuditorImpl) Enabled() bool { _ = "STUB: not implemented"; return false }

// getRecord returns the WorkloadRecord for the given key, or nil if not found.
func (w *workloadAuditorImpl) getRecord(key string) *WorkloadRecord {
	_ = "STUB: not implemented"
	return nil
}

func (w *workloadAuditorImpl) AddGangGroup(gangGroupID string) { _ = "STUB: not implemented"; return }

func (w *workloadAuditorImpl) DeleteGangGroup(gangGroupID string) {
	_ = "STUB: not implemented"
	return
}

func (w *workloadAuditorImpl) RecordGangGroup(gangGroupID string, pod *corev1.Pod, recordType RecordType, message string) {
	_ = "STUB: not implemented"
	return
}

func (w *workloadAuditorImpl) RecordDiagnosis(pod *corev1.Pod, questionKey string, recordType RecordType, message string) {
	_ = "STUB: not implemented"
	return
}

func (w *workloadAuditorImpl) AddPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (w *workloadAuditorImpl) RecordPod(pod *corev1.Pod, recordType RecordType, message string) {
	_ = "STUB: not implemented"
	return
}

func (w *workloadAuditorImpl) RecordPodGating(pod *corev1.Pod, gated bool) {
	_ = "STUB: not implemented"
	return
}

func (w *workloadAuditorImpl) RecordGangGating(gangGroupID string, pod *corev1.Pod, gated bool) {
	_ = "STUB: not implemented"
	return
}

func (w *workloadAuditorImpl) DeletePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (w *workloadAuditorImpl) RecordAttemptPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (w *workloadAuditorImpl) RecordPodScheduleResult(pod *corev1.Pod, recordType RecordType, message string) {
	_ = "STUB: not implemented"
	return
}

// When recording a ScheduleFailure, skip if a Diagnosis or other Failure
// has already been recorded for the current attempt.

// On successful schedule, remove the workload record.

func (w *workloadAuditorImpl) RecordGangScheduleResult(gangKey string, recordType RecordType, message string) {
	_ = "STUB: not implemented"
	return
}

// On successful gang schedule, remove the gang record.

// appendRecord increments the record-type count, logs the event, and runs anomaly detection.
func (w *workloadAuditorImpl) appendRecord(wr *WorkloadRecord, recordType RecordType, message string) {
	_ = "STUB: not implemented"
	return
}

// NewWorkloadAuditor creates a new workloadAuditorImpl reading config from package-level vars.
func NewWorkloadAuditor() WorkloadAuditor { _ = "STUB: not implemented"; return *new(WorkloadAuditor) }

// tryExtractLabels populates the WorkloadRecord's label values from the given pod,
// but only on the first call with a non-nil pod.
func (w *workloadAuditorImpl) tryExtractLabels(wr *WorkloadRecord, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// extractLabels reads metric label values from a pod.
// The first value is always the priority class; subsequent values come from pod labels
// as configured by PodLabelKeys.
func (w *workloadAuditorImpl) extractLabels(pod *corev1.Pod) []string {
	_ = "STUB: not implemented"
	return nil
}

// formatLabelDetail returns a compact string of label values for ALERT logs.
func formatLabelDetail(names []string, values []string) string {
	_ = "STUB: not implemented"
	return ""
}

// priority printed as bare value
