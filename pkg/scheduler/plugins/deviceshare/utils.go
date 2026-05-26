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

package deviceshare

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	NvidiaGPU = 1 << iota
	AMDGPU
	HygonDCU
	KoordGPU
	GPUShared
	GPUCore
	GPUMemory
	GPUMemoryRatio
	HuaweiNPUCore
	HuaweiNPUCPU
	HuaweiNPUDVPP
	FPGA
	RDMA
)

var DeviceResourceNames = map[schedulingv1alpha1.DeviceType][]corev1.ResourceName{
	schedulingv1alpha1.GPU: {
		apiext.ResourceNvidiaGPU,
		apiext.ResourceAMDGPU,
		apiext.ResourceHygonDCU,
		apiext.ResourceGPU,
		apiext.ResourceGPUShared,
		apiext.ResourceGPUCore,
		apiext.ResourceGPUMemory,
		apiext.ResourceGPUMemoryRatio,
		apiext.ResourceHuaweiNPUCore,
		apiext.ResourceHuaweiNPUCPU,
		apiext.ResourceHuaweiNPUDVPP,
	},
	schedulingv1alpha1.RDMA: {apiext.ResourceRDMA},
	schedulingv1alpha1.FPGA: {apiext.ResourceFPGA},
}

var DeviceResourceFlags = map[corev1.ResourceName]uint{
	apiext.ResourceNvidiaGPU:      NvidiaGPU,
	apiext.ResourceAMDGPU:         AMDGPU,
	apiext.ResourceHygonDCU:       HygonDCU,
	apiext.ResourceGPU:            KoordGPU,
	apiext.ResourceGPUCore:        GPUCore,
	apiext.ResourceGPUMemory:      GPUMemory,
	apiext.ResourceGPUMemoryRatio: GPUMemoryRatio,
	apiext.ResourceGPUShared:      GPUShared,
	apiext.ResourceHuaweiNPUCore:  HuaweiNPUCore,
	apiext.ResourceHuaweiNPUCPU:   HuaweiNPUCPU,
	apiext.ResourceHuaweiNPUDVPP:  HuaweiNPUDVPP,
	apiext.ResourceFPGA:           FPGA,
	apiext.ResourceRDMA:           RDMA,
}

var ValidDeviceResourceCombinations = map[uint]func(resources corev1.ResourceList) bool{
	NvidiaGPU:                            ValidDeviceResourceCombinationsDefaultTrue,
	AMDGPU:                               ValidDeviceResourceCombinationsDefaultTrue,
	HygonDCU:                             ValidDeviceResourceCombinationsDefaultTrue,
	KoordGPU:                             ValidDeviceResourceCombinationsDefaultTrue,
	GPUMemory:                            ValidDeviceResourceCombinationsGPUPercentage,
	GPUMemoryRatio:                       ValidDeviceResourceCombinationsGPUPercentage,
	GPUCore | GPUMemory:                  ValidDeviceResourceCombinationsGPUPercentage,
	GPUCore | GPUMemoryRatio:             ValidDeviceResourceCombinationsGPUPercentage,
	HuaweiNPUCore | GPUMemoryRatio:       ValidDeviceResourceCombinationsGPUPercentage,
	GPUShared | GPUMemory:                ValidDeviceResourceCombinationsGPUShared,
	GPUShared | GPUMemoryRatio:           ValidDeviceResourceCombinationsGPUShared,
	GPUShared | GPUCore | GPUMemory:      ValidDeviceResourceCombinationsGPUShared,
	GPUShared | GPUCore | GPUMemoryRatio: ValidDeviceResourceCombinationsGPUShared,
	GPUShared | HuaweiNPUCore | HuaweiNPUCPU | GPUMemory:                 ValidDeviceResourceCombinationsHuaweiNPUShared,
	GPUShared | HuaweiNPUCore | HuaweiNPUCPU | HuaweiNPUDVPP | GPUMemory: ValidDeviceResourceCombinationsHuaweiNPUShared,
	FPGA: ValidDeviceResourceCombinationsDefaultTrue,
	RDMA: ValidDeviceResourceCombinationsDefaultTrue,
}

var DeviceResourceValidators = map[corev1.ResourceName]func(q resource.Quantity) bool{
	apiext.ResourceGPU:  ValidatePercentageResource,
	apiext.ResourceFPGA: ValidatePercentageResource,
	apiext.ResourceRDMA: ValidatePercentageResource,
}

var ResourceCombinationsMapper = map[uint]func(podRequest corev1.ResourceList) corev1.ResourceList{
	GPUMemory: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUMemory: podRequest[apiext.ResourceGPUMemory],
		}
	},
	GPUMemoryRatio: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUMemoryRatio: podRequest[apiext.ResourceGPUMemoryRatio],
		}
	},
	GPUCore | GPUMemory: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUCore:   podRequest[apiext.ResourceGPUCore],
			apiext.ResourceGPUMemory: podRequest[apiext.ResourceGPUMemory],
		}
	},
	GPUCore | GPUMemoryRatio: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUCore:        podRequest[apiext.ResourceGPUCore],
			apiext.ResourceGPUMemoryRatio: podRequest[apiext.ResourceGPUMemoryRatio],
		}
	},
	HuaweiNPUCore | GPUMemoryRatio: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceHuaweiNPUCore:  podRequest[apiext.ResourceHuaweiNPUCore],
			apiext.ResourceGPUMemoryRatio: podRequest[apiext.ResourceGPUMemoryRatio],
		}
	},
	KoordGPU: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUCore:        podRequest[apiext.ResourceGPU],
			apiext.ResourceGPUMemoryRatio: podRequest[apiext.ResourceGPU],
		}
	},
	GPUShared | GPUMemory: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUShared: podRequest[apiext.ResourceGPUShared],
			apiext.ResourceGPUMemory: podRequest[apiext.ResourceGPUMemory],
		}
	},
	GPUShared | GPUMemoryRatio: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUShared:      podRequest[apiext.ResourceGPUShared],
			apiext.ResourceGPUMemoryRatio: podRequest[apiext.ResourceGPUMemoryRatio],
		}
	},
	GPUShared | GPUCore | GPUMemory: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUShared: podRequest[apiext.ResourceGPUShared],
			apiext.ResourceGPUCore:   podRequest[apiext.ResourceGPUCore],
			apiext.ResourceGPUMemory: podRequest[apiext.ResourceGPUMemory],
		}
	},
	GPUShared | GPUCore | GPUMemoryRatio: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUShared:      podRequest[apiext.ResourceGPUShared],
			apiext.ResourceGPUCore:        podRequest[apiext.ResourceGPUCore],
			apiext.ResourceGPUMemoryRatio: podRequest[apiext.ResourceGPUMemoryRatio],
		}
	},
	GPUShared | HuaweiNPUCore | HuaweiNPUCPU | GPUMemory: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUShared:     podRequest[apiext.ResourceGPUShared],
			apiext.ResourceHuaweiNPUCore: podRequest[apiext.ResourceHuaweiNPUCore],
			apiext.ResourceHuaweiNPUCPU:  podRequest[apiext.ResourceHuaweiNPUCPU],
			apiext.ResourceGPUMemory:     podRequest[apiext.ResourceGPUMemory],
		}
	},
	GPUShared | HuaweiNPUCore | HuaweiNPUCPU | HuaweiNPUDVPP | GPUMemory: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceGPUShared:     podRequest[apiext.ResourceGPUShared],
			apiext.ResourceHuaweiNPUCore: podRequest[apiext.ResourceHuaweiNPUCore],
			apiext.ResourceHuaweiNPUCPU:  podRequest[apiext.ResourceHuaweiNPUCPU],
			apiext.ResourceHuaweiNPUDVPP: podRequest[apiext.ResourceHuaweiNPUDVPP],
			apiext.ResourceGPUMemory:     podRequest[apiext.ResourceGPUMemory],
		}
	},
	NvidiaGPU: func(podRequest corev1.ResourceList) corev1.ResourceList {
		nvidiaGPU := podRequest[apiext.ResourceNvidiaGPU]
		return corev1.ResourceList{
			apiext.ResourceGPUCore:        *resource.NewQuantity(nvidiaGPU.Value()*100, resource.DecimalSI),
			apiext.ResourceGPUMemoryRatio: *resource.NewQuantity(nvidiaGPU.Value()*100, resource.DecimalSI),
		}
	},
	AMDGPU: func(podRequest corev1.ResourceList) corev1.ResourceList {
		amdGPU := podRequest[apiext.ResourceAMDGPU]
		return corev1.ResourceList{
			apiext.ResourceGPUCore:        *resource.NewQuantity(amdGPU.Value()*100, resource.DecimalSI),
			apiext.ResourceGPUMemoryRatio: *resource.NewQuantity(amdGPU.Value()*100, resource.DecimalSI),
		}
	},
	HygonDCU: func(podRequest corev1.ResourceList) corev1.ResourceList {
		hygonDCU := podRequest[apiext.ResourceHygonDCU]
		return corev1.ResourceList{
			apiext.ResourceGPUCore:        *resource.NewQuantity(hygonDCU.Value()*100, resource.DecimalSI),
			apiext.ResourceGPUMemoryRatio: *resource.NewQuantity(hygonDCU.Value()*100, resource.DecimalSI),
		}
	},
	FPGA: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceFPGA: podRequest[apiext.ResourceFPGA],
		}
	},
	RDMA: func(podRequest corev1.ResourceList) corev1.ResourceList {
		return corev1.ResourceList{
			apiext.ResourceRDMA: podRequest[apiext.ResourceRDMA],
		}
	},
}

func ValidatePercentageResource(q resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func ValidateMultiple(a, b resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func ValidateLessThan100Times(a, b resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func ValidDeviceResourceCombinationsGPUShared(podRequest corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

func ValidDeviceResourceCombinationsHuaweiNPUShared(podRequest corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

// multiple npu share is not supported on device side

func ValidDeviceResourceCombinationsGPUPercentage(podRequest corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

func ValidDeviceResourceCombinationsDefaultTrue(podRequest corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

func ValidateDeviceRequest(podRequest corev1.ResourceList) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ConvertDeviceRequest(podRequest corev1.ResourceList, combination uint) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func hasVirtualFunctions(nodeDevice *nodeDevice, deviceType schedulingv1alpha1.DeviceType) bool {
	_ = "STUB: not implemented"
	// TODO 这里可以异步掉，虽然计算量也不多
	return false
}

func mustAllocateVF(hint *apiext.DeviceHint) bool { _ = "STUB: not implemented"; return false }

func preparePod(pod *corev1.Pod, gpuSharedResourceTemplatesCache *gpuSharedResourceTemplatesCache, templateMatchedResources []corev1.ResourceName) (state *preFilterState, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func constructDesignatedVF(designatedAllocation apiext.DeviceAllocations) map[schedulingv1alpha1.DeviceType]map[int32]sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func GetPodDeviceRequests(pod *corev1.Pod) (map[schedulingv1alpha1.DeviceType]corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parsePodDeviceShareExtensions(pod *corev1.Pod, podRequests map[schedulingv1alpha1.DeviceType]corev1.ResourceList, state *preFilterState) error {
	_ = "STUB: not implemented"
	return nil
}

// selectors is type of [2]labels.Selector, so we don't need to worry it is nil or len != 2

func newHintSelectors(hints apiext.DeviceAllocateHints) (map[schedulingv1alpha1.DeviceType][2]labels.Selector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseGPURequirements(pod *corev1.Pod, podRequests map[schedulingv1alpha1.DeviceType]corev1.ResourceList, gpuHints *apiext.DeviceHint, gpuSharedResourceTemplatesCache *gpuSharedResourceTemplatesCache, templateMatchedResources []corev1.ResourceName) (*GPURequirements, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(zqzten): use non-strict finding for volcano style usage of huawei npu
