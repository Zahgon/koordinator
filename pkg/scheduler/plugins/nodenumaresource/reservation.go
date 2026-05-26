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
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

const reservationRestoreStateKey = Name + "/reservationRestoreState"

type reservationRestoreStateData struct {
	lock        sync.RWMutex
	nodeToState frameworkext.NodeReservationRestoreStates
}

type nodeReservationRestoreStateData struct {
	matched            map[types.UID]reusableAlloc   // matched reservation or pre-allocatable pods
	unmatched          map[types.UID]reusableAlloc   // unmatched reservations
	preAllocationRInfo *frameworkext.ReservationInfo // pre-allocating reservation

	mergedMatchedRemainCPUs    cpuset.CPUSet
	mergedMatchedAllocatedCPUs cpuset.CPUSet
	mergedMatchedAllocatable   map[int]corev1.ResourceList
	mergedMatchedAllocated     map[int]corev1.ResourceList
	mergedUnmatchedUsed        map[int]corev1.ResourceList
}

// reusableAlloc represents the allocatable/total reserved CPUs and allocated CPUs of a reservation or pre-allocatable pod.
type reusableAlloc struct {
	rInfo          *frameworkext.ReservationInfo // comparing reservation, i.e. matched reservation to allocate by a normal pod, or pre-allocation reservation to schedule
	preAllocatable *corev1.Pod

	allocatableCPUs cpuset.CPUSet // allocatable/total reserved CPUs
	allocatedCPUs   cpuset.CPUSet // allocated CPUs
	remainedCPUs    cpuset.CPUSet // unallocated reserved CPUs

	allocatable map[int]corev1.ResourceList
	allocated   map[int]corev1.ResourceList
	remained    map[int]corev1.ResourceList
}

func getReservationRestoreState(cycleState fwktype.CycleState) *reservationRestoreStateData {
	_ = "STUB: not implemented"
	return nil
}

func (s *reservationRestoreStateData) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

func (s *reservationRestoreStateData) getNodeState(nodeName string) *nodeReservationRestoreStateData {
	_ = "STUB: not implemented"
	return nil
}

func (s *reservationRestoreStateData) setNodeState(nodeName string, nodeState interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *reservationRestoreStateData) clearData() { _ = "STUB: not implemented"; return }

func (rs *nodeReservationRestoreStateData) mergeReservationAllocations() {
	_ = "STUB: not implemented"
	return
}

// merge matched reservations or pre-allocatable pods

func copyAllocated(m map[int]corev1.ResourceList) map[int]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

func subtractAllocated(m map[int]corev1.ResourceList, allocated map[int]corev1.ResourceList, withNonNegativeResult bool) map[int]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

func appendAllocated(m map[int]corev1.ResourceList, allocatedList ...map[int]corev1.ResourceList) map[int]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) PreRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// RestoreReservation restores the fine-grained resources (CPUSet, NUMA) held by matched reservations and unmatched allocated reservations.
func (p *Plugin) RestoreReservation(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, matched []*frameworkext.ReservationInfo, unmatched []*frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo) (interface{}, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NUMA resources: allocatable, allocated, remained

// cpuset resources: allocatable, allocated, remained

// also complete the nodeRestoreState in cycleState

// PreRestoreReservationPreAllocation is called before RestoreReservationPreAllocation to prepare state
func (p *Plugin) PreRestoreReservationPreAllocation(ctx context.Context, cycleState fwktype.CycleState, r *frameworkext.ReservationInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	// PreAllocation restore uses the same state as normal reservation restore
	return nil
}

// RestoreReservationPreAllocation restores the fine-grained resources (CPUSet, NUMA) held by pre-allocatable pods.
func (p *Plugin) RestoreReservationPreAllocation(ctx context.Context, cycleState fwktype.CycleState, r *frameworkext.ReservationInfo, preAllocatable []*corev1.Pod, nodeInfo fwktype.NodeInfo) (interface{}, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// pre-allocatable NUMA resources

// pre-allocatable CPUSet

// All CPUs are considered remained for restoration

// All NUMA resources are considered remained for restoration

// Store pre-allocatable pods' resources directly in the node state

// Merge with existing matched if present

// Trigger merge to compute merged state

// DEPRECATED
func (p *Plugin) FinalRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeToStates frameworkext.NodeReservationRestoreStates) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func tryAllocateFromReusable(
	manager ResourceManager,
	restoreState *nodeReservationRestoreStateData,
	resourceOptions *ResourceOptions,
	matchedReusableAllocs map[types.UID]reusableAlloc, // reusable allocations from matched reservations or pre-allocatable pods
	pod *corev1.Pod, // normal pod or pre-allocatabing reserve pod
	node *corev1.Node,
) (*PodAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// a normal pod with reservation-ignored specified

// update with node preemption state

// use the cycle state to avoid misunderstanding

// update with reservation preemption state

// TODO: Currently the ReservationAllocatePolicyDefault is actually implemented as
//       ReservationAllocatePolicyAligned. Need to re-visit the policies.

// For a Restricted reservation, pod should allocate CPUSet cpus only from it.

// Considering the reservation-ignored pod does not count in allocatedCPUs, we try preempting node CPUs
// which are overlapped with the reservation, while the preferred CPUs are already ready.

// For a Restricted reservation, we should restore the preemptible allocated CPUs twice, where one
// reference is by the reservation itself and the another is by the preempted owner pods.

// pre-allocating reserve pod can allocate more cpus than the pre-allocatable pod remained

// normal pod should allocate no more cpus than the reservation remained

// reserve pod should pre-allocate a superset of the pod requested cpus

func makeReasonsByReservation(reservationReasons []*fwktype.Status) []string {
	_ = "STUB: not implemented"
	return nil
}

// tryAllocateIgnoreReservation will try to allocate where the reserved resources of the node ignored.
func tryAllocateIgnoreReservation(manager ResourceManager,
	restoreState *nodeReservationRestoreStateData,
	resourceOptions *ResourceOptions,
	ignoredReservations map[types.UID]reusableAlloc,
	pod *corev1.Pod,
	node *corev1.Node,
) (*PodAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// update with node preemption state

// accumulate all ignored reserved resources which are not allocated to any owner pods

// update with reservation preemption state

// For a reservation-ignored pod, the pod can allocate resources from:
// (1) the node unallocated and unreserved;
// (2) the unallocated resources from the matched reservations.
// Since the nodeInfo snapshot double-calculate the allocated resources of the scheduled reservations,
// we should add this part to calculate the (1).

// allocate resources with nominated reservation (for a normal pod) or pre-allocatable pods (for a pre-allocating reservation)
func (p *Plugin) allocateWithNominated(
	restoreState *nodeReservationRestoreStateData,
	resourceOptions *ResourceOptions,
	pod *corev1.Pod,
	node *corev1.Node,
) (*PodAllocation, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the pod is reservation-ignored, it should allocate the node unallocated resources and all the reserved
// unallocated resources.

// Use nominated reservation for normal pod.
// Use nominated pre-allocatable pod for pre-allocating reservation.

// When the pod has a concrete nominated reservation but tryAllocateFromReusable
// did not produce an allocation (e.g. Restricted alloc failed in both probe and
// formal Allocate), refuse to fall back to tryAllocateFromNode. Falling back
// here would let the pod silently bypass the reservation and steal node
// capacity that is still reserved for the nominated R, and would desynchronize
// with the reservation plugin which will still write annotation
// scheduling.koordinator.sh/reservation-allocated = <nominated R> in PreBind,
// resulting in an annotation/cpuset mismatch.

func (p *Plugin) getPodNominatedReservationInfo(pod *corev1.Pod, nodeName string) *frameworkext.ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) getNominatedReusableAlloc(restoreState *nodeReservationRestoreStateData, resourceOptions *ResourceOptions, pod *corev1.Pod, node *corev1.Node) (map[types.UID]reusableAlloc, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for a normal pod, return the reusable alloc of the nominated reservation

// for a non-pre-allocating reservation, there will be no reusable alloc
