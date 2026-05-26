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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/util/bitmask"
)

var deviceHandlers = map[schedulingv1alpha1.DeviceType]DeviceHandler{}
var deviceAllocators = map[schedulingv1alpha1.DeviceType]DeviceAllocator{}

type DeviceHandler interface {
	CalcDesiredRequestsAndCount(node *corev1.Node, pod *corev1.Pod, podRequests corev1.ResourceList, nodeDevice *nodeDevice, hint *apiext.DeviceHint, state *preFilterState) (corev1.ResourceList, int, *fwktype.Status)
}

type DeviceAllocator interface {
	Allocate(requestCtx *requestContext, nodeDevice *nodeDevice, desiredCount int, maxDesiredCount int, preferredPCIEs sets.String) ([]*apiext.DeviceAllocation, *fwktype.Status)
}

type requestContext struct {
	pod                       *corev1.Pod
	node                      *corev1.Node
	requestsPerInstance       map[schedulingv1alpha1.DeviceType]corev1.ResourceList
	desiredCountPerDeviceType map[schedulingv1alpha1.DeviceType]int
	gpuRequirements           *GPURequirements
	hints                     apiext.DeviceAllocateHints
	hintSelectors             map[schedulingv1alpha1.DeviceType][2]labels.Selector
	required                  map[schedulingv1alpha1.DeviceType]sets.Int
	preferred                 map[schedulingv1alpha1.DeviceType]sets.Int
	allocationScorer          *resourceAllocationScorer
	nodeDevice                *nodeDevice

	designatedVF map[schedulingv1alpha1.DeviceType]map[int32]sets.Set[string]
}

type AutopilotAllocator struct {
	state                     *preFilterState
	phaseBeingExecuted        string
	nodeDevice                *nodeDevice
	node                      *corev1.Node
	pod                       *corev1.Pod
	scorer                    *resourceAllocationScorer
	numaNodes                 bitmask.BitMask
	requestsPerInstance       map[schedulingv1alpha1.DeviceType]corev1.ResourceList
	desiredCountPerDeviceType map[schedulingv1alpha1.DeviceType]int
}

func (a *AutopilotAllocator) Prepare() *fwktype.Status { _ = "STUB: not implemented"; return nil }

func (a *AutopilotAllocator) Allocate(
	required, preferred map[schedulingv1alpha1.DeviceType]sets.Int,
	requiredDeviceResources, preemptibleDeviceResources map[schedulingv1alpha1.DeviceType]deviceResources,
) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

func (a *AutopilotAllocator) filterNodeDevice(
	requiredDeviceResources, preemptibleDeviceResources map[schedulingv1alpha1.DeviceType]deviceResources,
) *nodeDevice {
	_ = "STUB: not implemented"
	return nil
}

// TODO if a.numaNodes == nil && selector == nil return all device of this deviceType

// TODO Device allocation logic hotspots discovered through flame graphs

func (a *AutopilotAllocator) calcRequestsAndCountByDeviceType(
	podRequests map[schedulingv1alpha1.DeviceType]corev1.ResourceList, nodeDevice *nodeDevice,
	hints apiext.DeviceAllocateHints, primaryDeviceType schedulingv1alpha1.DeviceType,
	podFitsSecondaryDeviceWellPlanned bool,
) (map[schedulingv1alpha1.DeviceType]corev1.ResourceList, map[schedulingv1alpha1.DeviceType]int, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (a *AutopilotAllocator) tryJointAllocate(requestCtx *requestContext, jointAllocate *apiext.DeviceJointAllocate, nodeDevice *nodeDevice) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

func (a *AutopilotAllocator) validateJointAllocation(jointAllocate *apiext.DeviceJointAllocate, nodeDevice *nodeDevice, deviceAllocations apiext.DeviceAllocations) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (a *AutopilotAllocator) jointAllocate(nodeDevice *nodeDevice, requestCtx *requestContext, jointAllocate *apiext.DeviceJointAllocate, primaryDeviceType schedulingv1alpha1.DeviceType, secondaryDeviceTypes []schedulingv1alpha1.DeviceType) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

func (a *AutopilotAllocator) allocateDevices(requestCtx *requestContext, nodeDevice *nodeDevice, deviceAllocations apiext.DeviceAllocations) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

func allocateDevices(requestCtx *requestContext, nodeDevice *nodeDevice, deviceType schedulingv1alpha1.DeviceType, requestPerInstance corev1.ResourceList, desiredCount int, preferredPCIEs sets.String) (allocations []*apiext.DeviceAllocation, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaultAllocateDevices(
	nodeDevice *nodeDevice,
	requestCtx *requestContext,
	podRequestPerInstance corev1.ResourceList,
	desiredCount int,
	maxDesiredCount int,
	deviceType schedulingv1alpha1.DeviceType,
	preferredPCIEs sets.String,
) ([]*apiext.DeviceAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO Device allocation logic hotspots discovered through flame graphs

// Skip unhealthy Device instances with zero resources

// TODO Device allocation logic hotspots discovered through flame graphs

func allocateVF(vfAllocation *VFAllocation, deviceInfos map[int]*schedulingv1alpha1.DeviceInfo, minor int, vfSelector labels.Selector, designatedVFOfMinor sets.Set[string]) *schedulingv1alpha1.VirtualFunction {
	_ = "STUB: not implemented"
	return nil
}

func newPreferredPCIes(nodeDevice *nodeDevice, deviceType schedulingv1alpha1.DeviceType, allocations []*apiext.DeviceAllocation) sets.String {
	_ = "STUB: not implemented"
	return *new(sets.String)
}

func (a *AutopilotAllocator) score(
	requiredDeviceResources, preemptibleDeviceResources map[schedulingv1alpha1.DeviceType]deviceResources,
) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO(joseph): Maybe different device types have different weights, but that's not currently supported.
