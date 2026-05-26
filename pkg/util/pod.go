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

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

func GetEmptyPodExtendedResources() *apiext.ExtendedResourceSpec {
	_ = "STUB: not implemented"
	return nil
}

func GetPodExtendedResources(pod *corev1.Pod) *apiext.ExtendedResourceSpec {
	_ = "STUB: not implemented"
	return nil
}

// GetPodTargetExtendedResources gets the resource requirements of a pod with given extended resources.
// It returns nil if pod specifies no extended resource.
func GetPodTargetExtendedResources(pod *corev1.Pod, resourceNames ...corev1.ResourceName) *apiext.ExtendedResourceSpec {
	_ = "STUB: not implemented"
	return nil
}

// TODO: count init containers and pod overhead

func GetPodKey(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func GetPodMetricKey(podMetric *slov1alpha1.PodMetricInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func IsPodTerminated(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsPodInactive returns true if the pod is not in Pending or Running phase.
// Pods in other states (Succeeded, Failed, Unknown) typically do not have any running containers
// and therefore do not have an associated cgroup directory.
//
// Note: A PodPhase of Pending does NOT mean the pod has not been scheduled.
// It may have already been assigned to this node, but its regular containers
// have not started yet (e.g., still running initContainers).
//
// Also, podMetas are retrieved from Kubelet's /pods endpoint, which only returns
// Pods that have been successfully scheduled to this node.
func IsPodInactive(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func GetCPUSetFromPod(podAnnotations map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
