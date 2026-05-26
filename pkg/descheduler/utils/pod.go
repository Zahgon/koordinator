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

package utils

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetResourceRequest finds and returns the request value for a specific resource.
func GetResourceRequest(pod *corev1.Pod, resource corev1.ResourceName) int64 {
	_ = "STUB: not implemented"
	return 0
}

// GetResourceRequestQuantity finds and returns the request quantity for a specific resource.
// It follows the KEP-753 sidecar container resource calculation:
// sidecar containers (initContainers with restartPolicy=Always) are summed like regular containers,
// while regular init containers use max-based comparison.
func GetResourceRequestQuantity(pod *corev1.Pod, resourceName corev1.ResourceName) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// Sidecar containers (initContainers with restartPolicy=Always) run alongside
// regular containers, so their requests should be summed.
// Regular init containers use max-based comparison.

// We assume pod overhead feature gate is enabled.
// We can't import the scheduler settings so we will inherit the default.

// IsMirrorPod returns true if the pod is a Mirror Pod.
func IsMirrorPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsPodTerminating returns true if the pod DeletionTimestamp is set.
func IsPodTerminating(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsStaticPod returns true if the pod is a static pod.
func IsStaticPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsCriticalPriorityPod returns true if the pod has critical priority.
func IsCriticalPriorityPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsDaemonsetPod returns true if the pod is a IsDaemonsetPod.
func IsDaemonsetPod(ownerRefList []metav1.OwnerReference) bool {
	_ = "STUB: not implemented"
	return false
}

// IsPodWithLocalStorage returns true if the pod has local storage.
func IsPodWithLocalStorage(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsPodWithPVC returns true if the pod has claimed a persistent volume.
func IsPodWithPVC(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// GetPodSource returns the source of the pod based on the annotation.
func GetPodSource(pod *corev1.Pod) (string, error) { _ = "STUB: not implemented"; return "", nil }
