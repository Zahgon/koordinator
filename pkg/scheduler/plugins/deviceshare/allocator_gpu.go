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
	"k8s.io/apimachinery/pkg/util/sets"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	ErrNoGPURequirements                    = "No GPU Requirements"
	ErrInsufficientPartitionedDevice        = "Insufficient Partitioned GPU Devices"
	ErrInsufficientTopologyScopedGPUDevices = "Insufficient Topology Scoped GPU Devices"
	ErrInsufficientGPUDevices               = "Insufficient GPU Devices"
	ErrNodeMissingGPUPartitionTable         = "node(s) missing GPU Partition Table"
	ErrUnsupportedGPURequests               = "node(s) Unsupported number of GPU requests"
	ErrUnsupportedMultiSharedGPU            = "node(s) Unsupported Multi-Shared GPU"
	ErrNodeMissingGPUDeviceTopologyTree     = "node(s) missing GPU Device Topology Tree"
	ErrNoMatchedGPUSharedResourceTemplate   = "no matched GPU shared resource template"
)

func init() {
	deviceAllocators[schedulingv1alpha1.GPU] = &GPUAllocator{}
}

var _ DeviceAllocator = &GPUAllocator{}

type GPUAllocator struct {
}

type AllocateContext struct {
	deviceUsedMinorsHash int
	deviceFree           deviceResources
	deviceTotal          deviceResources
	allocationScorer     *resourceAllocationScorer
}

func getRealUsed(originalUsed, refinedTotal, refinedUsed deviceResources) deviceResources {
	_ = "STUB: not implemented"
	return *new(deviceResources)
}

func (a *GPUAllocator) Allocate(requestCtx *requestContext, nodeDevice *nodeDevice, desiredCount int, maxDesiredCount int, preferredPCIEs sets.String) ([]*apiext.DeviceAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if pod does not enforce GPU shared resource template, allocateByTemplate may return (nil, nil), so we should check if allocations is nil

// if honorGPUPartition is false, allocateByPartition may return (nil, nil), so we should check if allocations is nil

func removeZeroDevice(originalResources deviceResources) deviceResources {
	_ = "STUB: not implemented"
	return *new(deviceResources)
}

func generalAllocate(requestCtx *requestContext, nodeDevice *nodeDevice, desiredCount int, maxDesiredCount int, allocateContext *AllocateContext) ([]*apiext.DeviceAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if same NUMANode or PCIE are not required, allocateByDeviceTopology may return (nil, nil), so we should check if allocations is nil

func allocateByTemplate(requestCtx *requestContext, nodeDevice *nodeDevice, desiredCount int, maxDesiredCount int, allocateContext *AllocateContext) ([]*apiext.DeviceAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// koord style: matching only one template for accurate resources specified by user

// TODO(zqzten): volcano style: automatically choose a template which meets user's requirement

func appendTemplateInfoToAllocations(allocations []*apiext.DeviceAllocation, templateName string) {
	_ = "STUB: not implemented"
	return
}

type GPUPartitionIndexer map[int][]*PartitionsOfAllocationScore

type PartitionsOfAllocationScore struct {
	Partitions      []*apiext.GPUPartition
	AllocationScore int
}

func allocateByPartition(honorGPUPartition bool, gpuRequirements *GPURequirements, gpuPartitionIndexer GPUPartitionIndexer, allocateContext *AllocateContext) (allocations []*apiext.DeviceAllocation, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if honorGPUPartition is false, allocateByPartition should return (allocation, nil)

// TODO when allocate shared gpu, partition binPack logic is equivalent with topology binPack in most machine models. Bus there may be still some unexpected machine model need to be considered

// we have to calculate this hash during scheduling cycle because reservation restore and preemption may happen

// we definitely prefer the partition with the higher allocation score

func hashDevices(resources deviceResources) int { _ = "STUB: not implemented"; return 0 }

func hashMinors(minors []int) int { _ = "STUB: not implemented"; return 0 }

type partitionOfBinPackScore struct {
	Partition    *apiext.GPUPartition
	BinPackScore int
}

func selectPartitionByBinPack(deviceUsedMinorsHash int, feasiblePartitions []*apiext.GPUPartition, partitionIndexer GPUPartitionIndexer, desiredNumberOfGPU int) *apiext.GPUPartition {
	_ = "STUB: not implemented"
	return nil
}

type GPUTopologyScope struct {
	scopeName       apiext.DeviceTopologyScope
	scopeLevel      int
	minorsResources deviceResources
	minors          []int
	minorsHash      int
	childScopes     []*GPUTopologyScope

	// only used for sort the scope
	pcieID     string
	numaNodeID int32
	minor      int32
}

func allocateByDeviceTopology(gpuRequirements *GPURequirements, gpuTopologyScope *GPUTopologyScope, allocateContext *AllocateContext) (allocations []*apiext.DeviceAllocation, status *fwktype.Status) {
	_ = "STUB: not implemented"

	// if allocateByDeviceTopology is not required and unsupported in some cases, then we can return nil instead of fwktype.UnschedulableAndUnresolvable to give the change of success
	return nil, nil
}

type ScopeLevelContext struct {
	cumulativeNotEmpties int
	depth                int
	contextOfDevices     map[int]*DeviceLevelContext
}

type DeviceLevelContext struct {
	satisfied bool
	score     int64
}

type ScopeLevelAllocateResult struct {
	allocations          []*apiext.DeviceAllocation
	cumulativeNotEmpties int
	depth                int
	score                int64
}

func allocateFromScope(requirements *GPURequirements, scope *GPUTopologyScope, allocateContext *AllocateContext, scopeLevelContext ScopeLevelContext) *ScopeLevelAllocateResult {
	_ = "STUB: not implemented"
	return nil
}
