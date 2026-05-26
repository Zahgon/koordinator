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

package sorter

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"

	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

var koordPriorityClassOrder = map[extension.PriorityClass]int{
	extension.PriorityNone:  5,
	extension.PriorityProd:  4,
	extension.PriorityMid:   3,
	extension.PriorityBatch: 2,
	extension.PriorityFree:  1,
}

var koordQoSClassOrder = map[extension.QoSClass]int{
	extension.QoSNone:   5,
	extension.QoSSystem: 4,
	extension.QoSLSE:    4,
	extension.QoSLSR:    3,
	extension.QoSLS:     2,
	extension.QoSBE:     1,
}

var k8sQoSClassOrder = map[corev1.PodQOSClass]int{
	corev1.PodQOSGuaranteed: 3,
	corev1.PodQOSBurstable:  2,
	corev1.PodQOSBestEffort: 1,
}

// Priority compares pods by Priority
func Priority(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

// KubernetesQoSClass compares pods by Kubernetes QosClass
func KubernetesQoSClass(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

// KoordinatorQoSClass compares pods by the Koordinator QoSClass
func KoordinatorQoSClass(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

// KoordinatorPriorityClass compares pods by the Koordinator PriorityClass
func KoordinatorPriorityClass(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

// PodUsage compares pods by the actual usage
func PodUsage(resourcesThatExceedThresholds map[corev1.ResourceName]resource.Quantity, podMetrics map[types.NamespacedName]*slov1alpha1.ResourceMap, resourceToWeightMap map[corev1.ResourceName]int64) CompareFn {
	_ = "STUB: not implemented"
	return *new(CompareFn)
}

// PodCreationTimestamp compares the pods by the creation timestamp
func PodCreationTimestamp(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

func PodDeletionCost(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

func EvictionCost(p1, p2 *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

func PodSorter(cmp ...CompareFn) *MultiSorter { _ = "STUB: not implemented"; return nil }

func SortPodsByUsage(resourcesThatExceedThresholds map[corev1.ResourceName]resource.Quantity, pods []*corev1.Pod, podMetrics map[types.NamespacedName]*slov1alpha1.ResourceMap, nodeAllocatableMap map[string]corev1.ResourceList, resourceToWeightMap map[corev1.ResourceName]int64) {
	_ = "STUB: not implemented"
	return
}
