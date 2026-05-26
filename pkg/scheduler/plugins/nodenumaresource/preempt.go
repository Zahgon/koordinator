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
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

type preemptibleAlloc struct {
	cpusToAdd     cpuset.CPUSet
	cpusToRemove  cpuset.CPUSet // only one CPUSet cannot store the removing info
	numaResources map[int]corev1.ResourceList
}

func newPreemptibleAlloc() *preemptibleAlloc { _ = "STUB: not implemented"; return nil }

func (a *preemptibleAlloc) Clone() *preemptibleAlloc { _ = "STUB: not implemented"; return nil }

func (a *preemptibleAlloc) AppendCPUSet(cpus cpuset.CPUSet) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

func (a *preemptibleAlloc) AppendNUMAResources(numaResources map[int]corev1.ResourceList) map[int]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

func (a *preemptibleAlloc) Accumulate(cpus cpuset.CPUSet, numaResources map[int]corev1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (a *preemptibleAlloc) Subtract(cpus cpuset.CPUSet, numaResources map[int]corev1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

type preemptibleNodeState struct {
	// always use a clone so that we don't need to lock the whole state
	nodeAlloc         *preemptibleAlloc               // unallocated and unreserved resources of the node
	reservationsAlloc map[types.UID]*preemptibleAlloc // reservation ID to unallocated resources of the reservation
}

func (ns *preemptibleNodeState) Clone() *preemptibleNodeState {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) AddPod(_ context.Context, cycleState fwktype.CycleState, preemptor *corev1.Pod, podInfoToAdd fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// preempt node unallocated resources

// preempt reservation resources

func (p *Plugin) RemovePod(_ context.Context, cycleState fwktype.CycleState, preemptor *corev1.Pod, podInfoToRemove fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// preempt node unallocated resources

// preempt reservation resources

func (p *Plugin) notNUMAAwareReservation(rInfo *frameworkext.ReservationInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Plugin) getPodAllocated(pod *corev1.Pod, nodeName string) (cpus cpuset.CPUSet, numaResources map[int]corev1.ResourceList) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}
