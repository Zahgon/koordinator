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

package quotaevaluate

import (
	"sync"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/utils/clock"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

// podResources are the set of resources managed by quota associated with pods.
var podResources = []corev1.ResourceName{
	corev1.ResourceCPU,
	corev1.ResourceMemory,
	corev1.ResourceEphemeralStorage,
	corev1.ResourceRequestsCPU,
	corev1.ResourceRequestsMemory,
	corev1.ResourceRequestsEphemeralStorage,

	// batch resource
	extension.BatchCPU,
	extension.BatchMemory,

	// mid resource
	extension.MidCPU,
	extension.MidMemory,

	// gpu resource
	extension.ResourceGPU,
	extension.ResourceNvidiaGPU,
	extension.ResourceGPUShared,
	extension.ResourceGPUMemoryRatio,
}

type Attributes struct {
	QuotaNamespace string
	QuotaName      string
	Operation      admissionv1.Operation
	Pod            *corev1.Pod
}

// Evaluator is used to see if quota constraints are satisfied.
type Evaluator interface {
	// Evaluate takes an operation and checks to see if quota constraints are satisfied.  It returns an error if they are not.
	// The default implementation processes related operations in chunks when possible.
	Evaluate(a *Attributes) error
}

type quotaEvaluator struct {
	quotaAccessor QuotaAccessor

	queue      *workqueue.Type
	workLock   sync.Mutex
	work       map[string][]*admissionWaiter
	dirtyWork  map[string][]*admissionWaiter
	inProgress sets.String

	workers int
	stopCh  <-chan struct{}
	init    sync.Once
}

type admissionWaiter struct {
	attributes *Attributes
	finished   chan struct{}
	result     error
}

type defaultDeny struct{}

func (defaultDeny) Error() string { _ = "STUB: not implemented"; return "" }

// IsDefaultDeny returns true if the error is defaultDeny
func IsDefaultDeny(err error) bool { _ = "STUB: not implemented"; return false }

func newAdmissionWaiter(a *Attributes) *admissionWaiter { _ = "STUB: not implemented"; return nil }

func NewQuotaEvaluator(quotaAccessor QuotaAccessor, workers int, stopCh <-chan struct{}) Evaluator {
	_ = "STUB: not implemented"
	return *new(Evaluator)
}

// start begins watching and syncing.
func (e *quotaEvaluator) start() { _ = "STUB: not implemented"; return }

func (e *quotaEvaluator) shutdownOnStop() { _ = "STUB: not implemented"; return }

func (e *quotaEvaluator) doWork() { _ = "STUB: not implemented"; return }

func (e *quotaEvaluator) checkAttributes(key string, admissionAttributes []*admissionWaiter) {
	_ = "STUB: not implemented"
	// notify all on exit
	return
}

func (e *quotaEvaluator) checkQuota(quota *v1alpha1.ElasticQuota, admissionAttributes []*admissionWaiter, remainingRetries int) {
	_ = "STUB: not implemented"
	// yet another copy to compare against originals to see if we actually have deltas
	return
}

// at this point, errors are fatal.  Update all waiters without status to failed and return

// this means that updates failed.  Anything with a default deny error has failed and we need to let them know

func (e *quotaEvaluator) Handles(a *Attributes) bool { _ = "STUB: not implemented"; return false }

func QuotaV1Pod(pod *corev1.Pod, clock clock.Clock) bool { _ = "STUB: not implemented"; return false }

func PodUsageFunc(pod *corev1.Pod, clock clock.Clock) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func (e *quotaEvaluator) checkRequest(quota *v1alpha1.ElasticQuota, a *Attributes) (*v1alpha1.ElasticQuota, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *quotaEvaluator) Evaluate(a *Attributes) error { _ = "STUB: not implemented"; return nil }

// wait for completion or timeout

func (e *quotaEvaluator) addWork(a *admissionWaiter) { _ = "STUB: not implemented"; return }

func (e *quotaEvaluator) completeWork(key string) { _ = "STUB: not implemented"; return }

func (e *quotaEvaluator) getWork() (string, []*admissionWaiter, bool) {
	_ = "STUB: not implemented"
	return "", nil, false
}

// prettyPrint formats a resource list for usage in errors
// it outputs resources sorted in increasing order
func prettyPrint(item corev1.ResourceList) string { _ = "STUB: not implemented"; return "" }

func prettyPrintResourceNames(a []corev1.ResourceName) string { _ = "STUB: not implemented"; return "" }
