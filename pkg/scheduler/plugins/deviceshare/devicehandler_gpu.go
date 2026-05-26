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
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

func init() {
	deviceHandlers[schedulingv1alpha1.GPU] = &GPUHandler{}
}

var _ DeviceHandler = &GPUHandler{}

type GPUHandler struct {
}

func (h *GPUHandler) CalcDesiredRequestsAndCount(node *corev1.Node, pod *corev1.Pod, podRequests corev1.ResourceList, nodeDevice *nodeDevice, hint *apiext.DeviceHint, state *preFilterState) (corev1.ResourceList, int, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), 0, nil
}

// just for test

func calcDesiredRequestsAndCountForGPU(podRequests corev1.ResourceList) (corev1.ResourceList, int, bool) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), 0, false
}

// gpu share mode

func fillGPUTotalMem(allocations apiext.DeviceAllocations, nodeDeviceInfo *nodeDevice) error {
	_ = "STUB: not implemented"
	return nil
}

func memoryRatioToBytes(ratio, totalMemory resource.Quantity) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

func memoryBytesToRatio(bytes, totalMemory resource.Quantity) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}
