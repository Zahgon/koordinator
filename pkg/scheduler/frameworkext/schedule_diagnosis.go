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

package frameworkext

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/workloadauditor"
)

var (
	dumpDiagnosis         = false
	dumpDiagnosisBlocking = false
	diagnosisQueueSize    = 1000
	diagnosisWorkerCount  = 10 // Default number of workers
)

var customDiagnosisProcessor map[string]func(diagnosis *Diagnosis)

func RegisterCustomDiagnosisProcessor(name string, processor func(diagnosis *Diagnosis)) {
	_ = "STUB: not implemented"
	return
}

// DumpDiagnosisSetter set dumpDiagnosis
func DumpDiagnosisSetter(val string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func DumpDiagnosisBlockingSetter(val string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func DumpDiagnosis(diagnosis *Diagnosis) string { _ = "STUB: not implemented"; return "" }

// Handle blocking mode

// For blocking mode, we still process synchronously

// For non-blocking mode, enqueue for asynchronous processing

func GetDiagnosis(state fwktype.CycleState) *Diagnosis { _ = "STUB: not implemented"; return nil }

// just for test

var nowFunc = metav1.Now

const (
	diagnosisStateKey = extension.SchedulingDomainPrefix + "/diagnosis"
)

func InitDiagnosis(state fwktype.CycleState, pod *corev1.Pod) { _ = "STUB: not implemented"; return }

var (
	_ fwktype.StateData = &Diagnosis{}
)

func (d *Diagnosis) Clone() fwktype.StateData {
	_ = "STUB: not implemented"

	// Diagnosis Help diagnose the journey of the Pod in SchedulePod and PostFilter.
	return *new(fwktype.StateData)
}

type Diagnosis struct {
	Timestamp            metav1.Time `json:"timestamp"`
	QuestionedKey        string      `json:"questionedKey,omitempty"`
	TargetPod            *corev1.Pod `json:"-"`
	NominatedNode        string      `json:"nominatedNode,omitempty"`
	PreFilterMessage     string      `json:"preFilterMessage,omitempty"`
	TopologyKeyToExplain string      `json:"topologyKeyToExplain,omitempty"`
	IsRootCausePod       bool        `json:"isRootCausePod"`
	// maybe modify fwktype.Status to cover addedNominatedPods, corresponding resourceView(such as requested and total) when failed
	ScheduleDiagnosis   *ScheduleDiagnosis   `json:"scheduleDiagnosis"`
	PreemptionDiagnosis *PreemptionDiagnosis `json:"preemptionDiagnosis"`

	// Suggestion is the scheduler's advice to external components (e.g. descheduler, autoscaler)
	// after a job/pod scheduling failure. For example, it may suggest resubmitting the pod or job,
	// waiting for victims to be deleted, or deleting the PVC bound to the pod.
	// Use SetSuggestion to set this field; once set, it cannot be overwritten.
	suggestionMu sync.RWMutex        `json:"-"`
	Suggestion   *ScheduleSuggestion `json:"suggestion,omitempty"`

	AuditType    workloadauditor.RecordType `json:"-"`
	AuditMessage string                     `json:"-"`
}

// SetSuggestion atomically sets the Suggestion field. It returns true if the suggestion
// was successfully set, or false if it was already set by a prior caller.
func (d *Diagnosis) SetSuggestion(suggestion *ScheduleSuggestion) bool {
	_ = "STUB: not implemented"
	return false
}

// GetSuggestion returns the current Suggestion under a read lock.
func (d *Diagnosis) GetSuggestion() *ScheduleSuggestion { _ = "STUB: not implemented"; return nil }

type ScheduleDiagnosis struct {
	SchedulingMode SchedulingMode `json:"-"`
	// AlreadyWaitForBoundPods and AlreadyWaitForBound only meaningful when PodSchedulingMode
	AlreadyWaitForBoundPods []*corev1.Pod              `json:"-"`
	AlreadyWaitForBound     int                        `json:"alreadyWaitForBound"`
	NodeOfferSlot           map[string]int             `json:"nodeOfferSlot,omitempty"`
	NodeToStatusMap         map[string]*fwktype.Status `json:"-"`
	// NodeFailedDetails
	NodeFailedDetails v1alpha1.NodeFailedDetails `json:"nodeFailedDetails,omitempty"`
}

// ScheduleSuggestion represents the scheduler's suggestion to external components
// when a job/pod scheduling failure occurs.
type ScheduleSuggestion struct {
	// Type indicates the category of the suggestion.
	Type SuggestionType `json:"type,omitempty"`
	// Message provides a human-readable explanation of the suggestion.
	Message string `json:"message,omitempty"`
}

// SuggestionType defines the type of scheduler suggestion.
type SuggestionType string

const (
	// SuggestionEvictWorkloadSelf indicates the current Job/Pod is no longer schedulable
	// and suggests deleting the Pod and resubmitting it.
	SuggestionEvictWorkloadSelf SuggestionType = "EvictWorkloadSelf"
	// SuggestionWaitingVictimReleased indicates the scheduler is waiting for victims to be deleted
	// before the Job/Pod can be scheduled.
	SuggestionWaitingVictimReleased SuggestionType = "WaitingVictimReleased"
	// SuggestionDeleteConflictPVC indicates the Pod's bound PVC is preventing scheduling
	// and suggests deleting the PVC so the Pod can be rescheduled.
	SuggestionDeleteConflictPVC SuggestionType = "DeleteConflictPVC"
)

type SchedulingMode string

const (
	PodSchedulingMode SchedulingMode = "Pod"
	JobSchedulingMode SchedulingMode = "Job"
)

type PreemptionDiagnosis struct {
	DryRunFilterDiagnosis *ScheduleDiagnosis `json:"dryRunFilterDiagnosis"`
	OtherDiagnosis        interface{}        `json:"otherDiagnosis"`
}

// DiagnosisQueue is a queue for handling diagnosis logs asynchronously
type DiagnosisQueue struct {
	queue chan *Diagnosis
	once  sync.Once
}

// Global diagnosis queue instance
var diagnosisQueue = &DiagnosisQueue{}

// StartWorker starts the worker goroutines for processing diagnosis logs
func (dq *DiagnosisQueue) StartWorker() { _ = "STUB: not implemented"; return }

// worker processes diagnosis logs from the queue
func (dq *DiagnosisQueue) worker() { _ = "STUB: not implemented"; return }

// processDiagnosis handles the actual logging of diagnosis information
func (dq *DiagnosisQueue) processDiagnosis(diagnosis *Diagnosis) string {
	_ = "STUB: not implemented"
	// Process NodeFailedDetails if empty
	return ""
}

func convertStatusMapToFailedDetail(statusMap map[string]*fwktype.Status) v1alpha1.NodeFailedDetails {
	_ = "STUB: not implemented"
	return *new(v1alpha1.NodeFailedDetails)
}

// Enqueue adds a diagnosis to the queue for asynchronous processing
func (dq *DiagnosisQueue) Enqueue(diagnosis *Diagnosis) { _ = "STUB: not implemented"; return }

// If the queue is full, drop the diagnosis to prevent blocking
