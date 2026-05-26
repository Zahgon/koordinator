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

package extension

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	BatchCPU    corev1.ResourceName = ResourceDomainPrefix + "batch-cpu"
	BatchMemory corev1.ResourceName = ResourceDomainPrefix + "batch-memory"
	MidCPU      corev1.ResourceName = ResourceDomainPrefix + "mid-cpu"
	MidMemory   corev1.ResourceName = ResourceDomainPrefix + "mid-memory"
)

const (
	// AnnotationExtendedResourceSpec specifies the resource requirements of extended resources for internal usage.
	// It annotates the requests/limits of extended resources and can be used by runtime proxy and koordlet that
	// cannot get the original pod spec in CRI requests.
	AnnotationExtendedResourceSpec = NodeDomainPrefix + "/extended-resource-spec"

	// AnnotationPodReplaceResources is used to replace or erase the resource name in pod spec.
	// This annotation allows the scheduler and webhook processing flow to update specified resources before evaluation.
	// NOTE: The usage scenarios must ensure that the resources on the node are not oversold by itself.
	AnnotationPodReplaceResources = PodDomainPrefix + "/replace-resources"
)

var (
	ResourceNameMap = map[PriorityClass]map[corev1.ResourceName]corev1.ResourceName{
		PriorityBatch: {
			corev1.ResourceCPU:    BatchCPU,
			corev1.ResourceMemory: BatchMemory,
		},
		PriorityMid: {
			corev1.ResourceCPU:    MidCPU,
			corev1.ResourceMemory: MidMemory,
		},
	}
	ReverseResourceNameMap = map[corev1.ResourceName]PriorityClass{
		BatchCPU:    PriorityBatch,
		BatchMemory: PriorityBatch,
		MidCPU:      PriorityMid,
		MidMemory:   PriorityMid,
	}
)

// TranslateResourceNameByPriorityClass translates defaultResourceName to extend resourceName by PriorityClass
func TranslateResourceNameByPriorityClass(priorityClass PriorityClass, defaultResourceName corev1.ResourceName) corev1.ResourceName {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceName)
}

type ExtendedResourceSpec struct {
	Containers map[string]ExtendedResourceContainerSpec `json:"containers,omitempty"`
}

type ExtendedResourceContainerSpec struct {
	Limits   corev1.ResourceList `json:"limits,omitempty"`
	Requests corev1.ResourceList `json:"requests,omitempty"`
}

// GetExtendedResourceSpec parses ExtendedResourceSpec from annotations
func GetExtendedResourceSpec(annotations map[string]string) (*ExtendedResourceSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetExtendedResourceSpec(pod *corev1.Pod, spec *ExtendedResourceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPodReplaceResourcesConfig parses replace and erase resource names from pod annotations.
// Annotation format: "resourceA:resourceB,resourceC:resourceD,resourceE:"
// means replace resourceA with resourceB, replace resourceC with resourceD, erase resourceE
func GetPodReplaceResourcesConfig(pod *corev1.Pod) (eraseNames []corev1.ResourceName,
	replaceMappings map[corev1.ResourceName]corev1.ResourceName) {
	_ = "STUB: not implemented"
	return nil, nil
}
