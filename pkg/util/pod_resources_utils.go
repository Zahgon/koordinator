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

package util

import (
	corev1 "k8s.io/api/core/v1"
)

// NOTE: functions in this file can be overwritten for extension

// IsSidecarContainer returns true if the container is an init container with
// restartPolicy=Always, which means it is a sidecar container (KEP-753).
func IsSidecarContainer(c corev1.Container) bool { _ = "STUB: not implemented"; return false }

func GetPodMilliCPULimit(pod *corev1.Pod) int64 { _ = "STUB: not implemented"; return 0 }

// Sidecar containers (init containers with restartPolicy=Always) run alongside
// regular containers, so their limits should be summed rather than max-ed.

func GetPodRequest(pod *corev1.Pod, resourceNames ...corev1.ResourceName) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func GetPodBEMilliCPURequest(pod *corev1.Pod) int64 { _ = "STUB: not implemented"; return 0 }

// Sidecar containers run alongside regular containers and should be summed.

func GetPodBEMilliCPULimit(pod *corev1.Pod) int64 { _ = "STUB: not implemented"; return 0 }

// Sidecar containers run alongside regular containers and should be summed.

func GetPodBEMemoryByteRequestIgnoreUnlimited(pod *corev1.Pod) int64 {
	_ = "STUB: not implemented"
	return 0
}

// consider request of unlimited container as 0

// Sidecar containers run alongside regular containers and should be summed.

func GetPodBEMemoryByteLimit(pod *corev1.Pod) int64 { _ = "STUB: not implemented"; return 0 }

// Sidecar containers run alongside regular containers and should be summed.

// AddResourceList adds the resources in newList to list.
func AddResourceList(list, newList corev1.ResourceList) { _ = "STUB: not implemented"; return }
