/*
Copyright 2022 The Koordinator Authors.
Copyright 2017 The Kubernetes Authors.

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

package evictions

import (
	"context"
	"sync"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/events"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
	podutil "github.com/koordinator-sh/koordinator/pkg/descheduler/pod"
)

const (
	EvictPodAnnotationKey = "descheduler.alpha.kubernetes.io/evict"
)

type nodePodEvictedCount map[string]uint
type namespacePodEvictCount map[string]uint

type PodEvictor struct {
	client                     clientset.Interface
	eventRecorder              events.EventRecorder
	policyGroupVersion         string
	dryRun                     bool
	maxPodsToEvictPerNode      *uint
	maxPodsToEvictPerNamespace *uint
	lock                       sync.RWMutex
	totalCount                 int
	nodepodCount               nodePodEvictedCount
	namespacePodCount          namespacePodEvictCount
}

func NewPodEvictor(
	client clientset.Interface,
	eventRecorder events.EventRecorder,
	policyGroupVersion string,
	dryRun bool,
	maxPodsToEvictPerNode *uint,
	maxPodsToEvictPerNamespace *uint,
) *PodEvictor {
	_ = "STUB: not implemented"
	return nil
}

// NodeEvicted gives a number of pods evicted for node
func (pe *PodEvictor) NodeEvicted(nodeName string) uint { _ = "STUB: not implemented"; return 0 }

func (pe *PodEvictor) NamespaceEvicted(namespace string) uint { _ = "STUB: not implemented"; return 0 }

// TotalEvicted gives a number of pods evicted through all nodes
func (pe *PodEvictor) TotalEvicted() int { _ = "STUB: not implemented"; return 0 }

// NodeLimitExceeded checks if the number of evictions for a node was exceeded
func (pe *PodEvictor) NodeLimitExceeded(nodeName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (pe *PodEvictor) NamespaceLimitExceeded(namespace string) bool {
	_ = "STUB: not implemented"
	return false
}

func (pe *PodEvictor) Evict(ctx context.Context, pod *corev1.Pod, opts framework.EvictOptions) bool {
	_ = "STUB: not implemented"
	return false
}

// err is used only for logging purposes

func EvictPod(ctx context.Context, client clientset.Interface, pod *corev1.Pod, policyGroupVersion string, deleteOptions *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	// In Kubernetes 1.25+, the v1beta1 eviction API was removed.
	// Always use policy/v1 Eviction API for k8s 1.25 and above.
	return nil
}

type Options struct {
	priority      *int32
	nodeFit       bool
	labelSelector *metav1.LabelSelector
}

// WithPriorityThreshold sets a threshold for pod's priority class.
// Any pod whose priority class is lower is evictable.
func WithPriorityThreshold(priority *int32) func(opts *Options) {
	_ = "STUB: not implemented"
	return nil
}

// WithNodeFit sets whether or not to consider taints, node selectors,
// and pod affinity when evicting. A pod whose tolerations, node selectors,
// and affinity match a node other than the one it is currently running on
// is evictable.
func WithNodeFit(nodeFit bool) func(opts *Options) { _ = "STUB: not implemented"; return nil }

// WithLabelSelector sets whether or not to apply label filtering when evicting.
// Any pod matching the label selector is considered evictable.
func WithLabelSelector(labelSelector *metav1.LabelSelector) func(opts *Options) {
	_ = "STUB: not implemented"
	return nil
}

type nodeGetterFn func() ([]*corev1.Node, error)

type constraint func(pod *corev1.Pod) error

type EvictorFilter struct {
	constraints []constraint
}

func NewEvictorFilter(
	nodeGetter nodeGetterFn,
	nodeIndexer podutil.GetPodsAssignedToNodeFunc,
	evictLocalStoragePods bool,
	evictSystemCriticalPods bool,
	ignorePvcPods bool,
	evictFailedBarePods bool,
	evictAllBarePods bool,
	opts ...func(opts *Options),
) (*EvictorFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Enable evictFailedBarePods to evict bare pods in failed phase

// Moved from IsEvictable function for backward compatibility

// Moved from IsEvictable function to allow for disabling

// todo: to align with the k8s descheduling framework, nodeFit should be moved into PreEvictionFilter in the future

// Filter decides when a pod is evictable
func (ef *EvictorFilter) Filter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// HaveEvictAnnotation checks if the pod have evict annotation
func HaveEvictAnnotation(obj metav1.Object) bool { _ = "STUB: not implemented"; return false }

// IsPodEvictableBasedOnPriority checks if the given pod is evictable based on priority resolved from pod Spec.
func IsPodEvictableBasedOnPriority(pod *corev1.Pod, priority int32) bool {
	_ = "STUB: not implemented"
	return false
}
