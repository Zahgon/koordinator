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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

type NodeAllocation struct {
	lock               sync.RWMutex
	nodeName           string
	allocatedPods      map[types.UID]PodAllocation
	allocatedCPUs      CPUDetails
	allocatedResources map[int]*NUMANodeResource
	sharedNode         map[int]sets.String
	singleNUMANode     map[int]sets.String
}

type PodAllocation struct {
	UID                types.UID                           `json:"uid,omitempty"`
	Namespace          string                              `json:"namespace,omitempty"`
	Name               string                              `json:"name,omitempty"`
	CPUSet             cpuset.CPUSet                       `json:"cpuset,omitempty"`
	CPUExclusivePolicy schedulingconfig.CPUExclusivePolicy `json:"cpuExclusivePolicy,omitempty"`
	NUMANodeResources  []NUMANodeResource                  `json:"numaNodeResources,omitempty"`
}

func (n *NodeAllocation) GetAllNUMANodeStatus(numaNodes int) []extension.NumaNodeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (n *NodeAllocation) NUMANodeSharedStatus(nodeid int) extension.NumaNodeStatus {
	_ = "STUB: not implemented"
	return *new(extension.NumaNodeStatus)
}

func NewNodeAllocation(nodeName string) *NodeAllocation { _ = "STUB: not implemented"; return nil }

func (n *NodeAllocation) update(allocation *PodAllocation, cpuTopology *CPUTopology) {
	_ = "STUB: not implemented"
	return
}

func (n *NodeAllocation) getCPUs(podUID types.UID) (cpuset.CPUSet, bool) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), false
}

func (n *NodeAllocation) getNUMAResource(podUID types.UID) (map[int]corev1.ResourceList, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *NodeAllocation) addCPUs(cpuTopology *CPUTopology, podUID types.UID, cpuset cpuset.CPUSet, exclusivePolicy schedulingconfig.CPUExclusivePolicy) {
	_ = "STUB: not implemented"
	return
}

func (n *NodeAllocation) addPodAllocation(request *PodAllocation, cpuTopology *CPUTopology) {
	_ = "STUB: not implemented"
	return
}

func (n *NodeAllocation) release(podUID types.UID) { _ = "STUB: not implemented"; return }

func (n *NodeAllocation) getAvailableCPUs(cpuTopology *CPUTopology, maxRefCount int, reservedCPUs cpuset.CPUSet, preferredCPUs ...cpuset.CPUSet) (availableCPUs cpuset.CPUSet, allocateInfo CPUDetails) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), *new(CPUDetails)
}

// NOTE: preferredCPUs is a slice since we may restore a cpu multiple times when it is referenced more than once.
// e.g. For a pod A tries to preempt a pod B in the reservation R, the CPU C is allocated twice by the B and R.
// So the RefCount of CPU C is 2, and the pod A should allocate it by returning the C twice, where the one is
// by the reservation restoring and the another is by the preemption restoring.

func (n *NodeAllocation) getAvailableNUMANodeResources(topologyOptions TopologyOptions, reusableResources map[int]corev1.ResourceList) (totalAvailable, totalAllocated map[int]corev1.ResourceList) {
	_ = "STUB: not implemented"
	return nil, nil
}
