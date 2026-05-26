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

package nodenumaresource

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/topologymanager"
	"github.com/koordinator-sh/koordinator/pkg/util/bitmask"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

type ResourceManager interface {
	GetTopologyHints(node *corev1.Node, pod *corev1.Pod, options *ResourceOptions, policy apiext.NUMATopologyPolicy, restoreState *nodeReservationRestoreStateData) (map[string][]topologymanager.NUMATopologyHint, error)
	Allocate(node *corev1.Node, pod *corev1.Pod, options *ResourceOptions) (*PodAllocation, *fwktype.Status)

	Update(nodeName string, allocation *PodAllocation)
	Release(nodeName string, podUID types.UID)

	GetNodeAllocation(nodeName string) *NodeAllocation
	GetAllocatedCPUSet(nodeName string, podUID types.UID) (cpuset.CPUSet, bool)
	GetAllocatedNUMAResource(nodeName string, podUID types.UID) (map[int]corev1.ResourceList, bool)
	GetAvailableCPUs(nodeName string, preferredCPUs ...cpuset.CPUSet) (availableCPUs cpuset.CPUSet, allocated CPUDetails, err error)
}

type ResourceOptions struct {
	numCPUsNeeded           int
	requestCPUBind          bool
	requests                corev1.ResourceList
	originalRequests        corev1.ResourceList
	requiredCPUBindPolicy   bool
	cpuBindPolicy           schedulingconfig.CPUBindPolicy
	cpuExclusivePolicy      schedulingconfig.CPUExclusivePolicy
	preferredCPUs           cpuset.CPUSet
	preemptibleCPUs         cpuset.CPUSet // cpus could be allocated by preemption
	reusableResources       map[int]corev1.ResourceList
	requiredResources       map[int]corev1.ResourceList
	requiredFromReservation bool
	requiredPreAllocation   bool
	hint                    topologymanager.NUMATopologyHint
	topologyOptions         TopologyOptions
	numaScorer              *resourceAllocationScorer
	nodePreemptionState     *preemptibleNodeState
}

type resourceManager struct {
	numaAllocateStrategy   schedulingconfig.NUMAAllocateStrategy
	topologyOptionsManager TopologyOptionsManager
	lock                   sync.Mutex
	nodeAllocations        map[string]*NodeAllocation
}

func NewResourceManager(
	handle fwktype.Handle,
	defaultNUMAAllocateStrategy schedulingconfig.NUMAAllocateStrategy,
	topologyOptionsManager TopologyOptionsManager,
) ResourceManager {
	_ = "STUB: not implemented"
	return *new(ResourceManager)
}

func (c *resourceManager) onNodeDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *resourceManager) getOrCreateNodeAllocation(nodeName string) *NodeAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (c *resourceManager) GetTopologyHints(node *corev1.Node, pod *corev1.Pod, options *ResourceOptions, policy apiext.NUMATopologyPolicy, restoreStateData *nodeReservationRestoreStateData) (map[string][]topologymanager.NUMATopologyHint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: restore reserved cpus according to reservations' allocatePolicy

// update with preemptible resources

func (c *resourceManager) trimNUMANodeResources(nodeName string, totalAvailable map[int]corev1.ResourceList, options *ResourceOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Allocate NUMA resources and CPUSet for the normal pod or reserve pod.
func (c *resourceManager) Allocate(node *corev1.Node, pod *corev1.Pod, options *ResourceOptions) (*PodAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *resourceManager) allocateResourcesByHint(node *corev1.Node, pod *corev1.Pod, options *ResourceOptions) ([]NUMANodeResource, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryBestToDistributeEvenly(requests corev1.ResourceList, totalAvailable map[int]corev1.ResourceList, options *ResourceOptions) ([]NUMANodeResource, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitQuantity(resourceName corev1.ResourceName, quantity resource.Quantity, numaNodeCount int, options *ResourceOptions) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

func allocateRes(available, request resource.Quantity) (resource.Quantity, resource.Quantity, resource.Quantity) {
	_ = "STUB: not implemented"
	return *new(resource.Quantity), *new(resource.Quantity), *new(resource.Quantity)
}

func (c *resourceManager) allocateCPUSet(node *corev1.Node, pod *corev1.Pod, allocatedNUMANodes []NUMANodeResource, options *ResourceOptions) (cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}

func (c *resourceManager) Update(nodeName string, allocation *PodAllocation) {
	_ = "STUB: not implemented"
	return
}

func (c *resourceManager) Release(nodeName string, podUID types.UID) {
	_ = "STUB: not implemented"
	return
}

func (c *resourceManager) GetAllocatedCPUSet(nodeName string, podUID types.UID) (cpuset.CPUSet, bool) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), false
}

func (c *resourceManager) GetAllocatedNUMAResource(nodeName string, podUID types.UID) (map[int]corev1.ResourceList, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *resourceManager) GetAvailableCPUs(nodeName string, preferredCPUs ...cpuset.CPUSet) (availableCPUs cpuset.CPUSet, allocated CPUDetails, err error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), *new(CPUDetails), nil
}

func (c *resourceManager) GetNodeAllocation(nodeName string) *NodeAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (c *resourceManager) getAvailableNUMANodeResources(nodeName string, topologyOptions TopologyOptions, reusableResources map[int]corev1.ResourceList) (totalAvailable, totalAllocated map[int]corev1.ResourceList, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *resourceManager) generateResourceHints(node *corev1.Node, pod *corev1.Pod, numaNodeResources []NUMANodeResource, options *ResourceOptions, totalAvailable map[int]corev1.ResourceList, policy apiext.NUMATopologyPolicy, restoreState *nodeReservationRestoreStateData) map[string][]topologymanager.NUMATopologyHint {
	_ = "STUB: not implemented"
	return nil
}

// update hints preferred according to multiNUMAGroups, in case when it wasn't provided, the default
// behavior to prefer the minimal amount of NUMA nodes will be used

// no possible NUMA affinities for resource

type hintsGenerator struct {
	numaNodesLackResource map[corev1.ResourceName][]int
	minAffinitySize       map[corev1.ResourceName]int
	hints                 map[string][]topologymanager.NUMATopologyHint
}

func (g *hintsGenerator) generateHints(mask bitmask.BitMask, score int64, resourceNames ...corev1.ResourceName) {
	_ = "STUB: not implemented"
	return
}

func filterCPUsByRequiredCPUBindPolicy(policy schedulingconfig.CPUBindPolicy, availableCPUs cpuset.CPUSet, cpuDetails CPUDetails, cpusPerCore int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// TODO(joseph): Maybe we should support required exclusive policy as following
//	allocated := allocatedCPUs.CPUsInCores(core)
//	if allocated.Size() > 0 {
//		cpuInfo := allocatedCPUs[allocated.ToSliceNoSort()[0]]
//		if cpuInfo.ExclusivePolicy != "" &&
//			cpuInfo.ExclusivePolicy != schedulingconfig.CPUExclusivePolicyNone &&
//			cpuInfo.ExclusivePolicy == exclusivePolicy {
//			continue
//		}
//	}
// Using only one CPU per core ensures correct hints are generated

func satisfiedRequiredCPUBindPolicy(policy schedulingconfig.CPUBindPolicy, cpus cpuset.CPUSet, topology *CPUTopology) error {
	_ = "STUB: not implemented"
	return nil
}

func determineFullPCPUs(cpus cpuset.CPUSet, details CPUDetails, cpusPerCore int) bool {
	_ = "STUB: not implemented"
	return false
}

func determineSpreadByPCPUs(cpus cpuset.CPUSet, details CPUDetails) bool {
	_ = "STUB: not implemented"
	return false
}
