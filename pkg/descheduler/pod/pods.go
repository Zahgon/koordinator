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

package pod

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

// FilterFunc is a filter for a pod.
type FilterFunc = framework.FilterFunc

// GetPodsAssignedToNodeFunc is a function which accept a node name and a pod filter function
// as input and returns the pods that assigned to the node.
type GetPodsAssignedToNodeFunc = framework.GetPodsAssignedToNodeFunc

// WrapFilterFuncs wraps a set of FilterFunc in one.
func WrapFilterFuncs(filters ...FilterFunc) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}

type Options struct {
	filter             FilterFunc
	includedNamespaces sets.String
	excludedNamespaces sets.String
	labelSelector      *metav1.LabelSelector
}

// NewOptions returns an empty Options.
func NewOptions() *Options {
	_ = "STUB: not implemented"

	// WithFilter sets a pod filter.
	// The filter function should return true if the pod should be returned from ListPodsOnANode
	return nil
}

func (o *Options) WithFilter(filter FilterFunc) *Options { _ = "STUB: not implemented"; return nil }

// WithNamespaces sets included namespaces
func (o *Options) WithNamespaces(namespaces sets.String) *Options {
	_ = "STUB: not implemented"
	return nil
}

// WithoutNamespaces sets excluded namespaces
func (o *Options) WithoutNamespaces(namespaces sets.String) *Options {
	_ = "STUB: not implemented"
	return nil
}

// WithLabelSelector sets a pod label selector
func (o *Options) WithLabelSelector(labelSelector *metav1.LabelSelector) *Options {
	_ = "STUB: not implemented"
	return nil
}

// BuildFilterFunc builds a final FilterFunc based on Options.
func (o *Options) BuildFilterFunc() (FilterFunc, error) {
	_ = "STUB: not implemented"
	return *new(FilterFunc), nil
}

// ListPodsOnANode lists all pods on a node.
// It also accepts a "filter" function which can be used to further limit the pods that are returned.
// (Usually this is podEvictor.Evictable().IsEvictable, in order to only list the evictable pods on a node, but can
// be used by strategies to extend it if there are further restrictions, such as with NodeAffinity).
func ListPodsOnANode(
	nodeName string,
	getPodsAssignedToNode GetPodsAssignedToNodeFunc,
	filter FilterFunc,
) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	// Succeeded and failed pods are not considered because they don't occupy any resource.
	return nil, nil
}

// ListAllPodsOnANode lists all the pods on a node no matter what the phase of the pod is.
func ListAllPodsOnANode(
	nodeName string,
	getPodsAssignedToNode GetPodsAssignedToNodeFunc,
	filter FilterFunc,
) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OwnerRef returns the ownerRefList for the pod.
func OwnerRef(pod *corev1.Pod) []metav1.OwnerReference { _ = "STUB: not implemented"; return nil }

func IsBestEffortPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsBurstablePod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func IsGuaranteedPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// SortPodsBasedOnPriorityLowToHigh sorts pods based on their priorities from low to high.
// If pods have same priorities, they will be sorted by QoS in the following order:
// BestEffort, Burstable, Guaranteed
func SortPodsBasedOnPriorityLowToHigh(pods []*corev1.Pod) { _ = "STUB: not implemented"; return }

// SortPodsBasedOnAge sorts Pods from oldest to most recent in place
func SortPodsBasedOnAge(pods []*corev1.Pod) { _ = "STUB: not implemented"; return }
