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
	"context"
	"sync"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const reservationRestoreStateKey = Name + "/reservationRestoreState"

type reservationRestoreStateData struct {
	lock        sync.RWMutex
	skip        bool
	nodeToState frameworkext.NodeReservationRestoreStates
}

type nodeReservationRestoreStateData struct {
	matched            []reusableAlloc               // matched reservations or pre-allocatable podds
	unmatched          []reusableAlloc               // unmatched reservations
	preAllocationRInfo *frameworkext.ReservationInfo // pre-allocating reservation info

	mergedMatchedAllocatable map[schedulingv1alpha1.DeviceType]deviceResources
	mergedMatchedAllocated   map[schedulingv1alpha1.DeviceType]deviceResources
	mergedUnmatchedUsed      map[schedulingv1alpha1.DeviceType]deviceResources
}

// reusableAlloc represents the allocatable/total reserved devices and allocated devices of a reservation or pre-allocatable pod.
type reusableAlloc struct {
	rInfo          *frameworkext.ReservationInfo // comparing reservation, i.e. the available reservation to allocate or the pre-allocating reservation
	preAllocatable *corev1.Pod
	allocatable    map[schedulingv1alpha1.DeviceType]deviceResources
	allocated      map[schedulingv1alpha1.DeviceType]deviceResources
	remained       map[schedulingv1alpha1.DeviceType]deviceResources
}

func getReservationRestoreState(cycleState fwktype.CycleState) *reservationRestoreStateData {
	_ = "STUB: not implemented"
	return nil
}

func cleanReservationRestoreState(cycleState fwktype.CycleState) { _ = "STUB: not implemented"; return }

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

func (rs *nodeReservationRestoreStateData) mergeReservationAllocations() {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) PreRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) RestoreReservation(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, matched []*frameworkext.ReservationInfo, unmatched []*frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo) (interface{}, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// also complete the nodeRestoreState in cycleState

func (p *Plugin) PreRestoreReservationPreAllocation(ctx context.Context, cycleState fwktype.CycleState, r *frameworkext.ReservationInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) RestoreReservationPreAllocation(ctx context.Context, cycleState fwktype.CycleState, r *frameworkext.ReservationInfo, preAllocatable []*corev1.Pod, nodeInfo fwktype.NodeInfo) (interface{}, *fwktype.Status) {
	_ = "STUB: not implemented"
	// retrieve reusable for pre-allocatable pods
	return nil, nil
}

// merge with nodeState

// DEPRECATED
func (p *Plugin) FinalRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeToStates frameworkext.NodeReservationRestoreStates) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) tryAllocateFromReusable(
	allocator *AutopilotAllocator,
	state *preFilterState,
	restoreState *nodeReservationRestoreStateData,
	matchedReusableAllocs []reusableAlloc,
	pod *corev1.Pod,
	node *corev1.Node,
	basicPreemptible map[schedulingv1alpha1.DeviceType]deviceResources,
	requiredFromReservation bool,
) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

// use the cycle state to avoid misunderstanding

//
// The Aligned and Restricted Policies only allow Pods to be allocated from current Reservation
// and the remaining resources of the node.
// And the Restricted policy requires that if the device resources requested by the Pod overlap with
// the device resources reserved by the current Reservation, such devices can only be allocated from the current Reservation.
// The formula for calculating the remaining amount per device instance in this scenario is as follows:
// basicPreemptible = sum(allocated(unmatched reservations)) + preemptible(node)
// free = total - (used - basicPreemptible - sum(allocated(matched reservations)) - sum(remained(currentReservation)) - preemptible(currentReservation))
//

// TODO: Currently the ReservationAllocatePolicyDefault is actually implemented as
//       ReservationAllocatePolicyAligned. Need to re-visit the policies.

//
// It is necessary to check separately whether the remaining resources of the device instance
// reserved by the Restricted Reservation meet the requirements of the Pod, to ensure that
// the intersecting resources do not exceed the reserved range of the Restricted Reservation.
//
// Example: the node has reservation-ignored pod, matched reservations R1, R2, ..., Ri,
// unmatched reservations U1, U2, ..., Uj, and pods P1, P2, ..., Pk.
// The free device resources for the scheduling pod P0 is:
// min(NodeTotal - P1 - P2 - ... - Pk - U1 - U2 - ... - Uj, R1)

// pre-allocating reservation can allocate more than the pre-allocatable pod remained

// For pre-allocation, validate that the allocated devices include all devices from the pre-allocatable pod

// tryAllocateIgnoreReservation will try to allocate where the reserved resources of the node ignored.
func (p *Plugin) tryAllocateIgnoreReservation(
	allocator *AutopilotAllocator,
	state *preFilterState,
	restoreState *nodeReservationRestoreStateData,
	ignoredReservations []reusableAlloc,
	node *corev1.Node,
	basicPreemptible map[schedulingv1alpha1.DeviceType]deviceResources,
) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

// accumulate all ignored reserved resources which are not allocated to any owner pods

func (p *Plugin) makeReasonsByReservation(reservationReasons []*fwktype.Status) []string {
	_ = "STUB: not implemented"
	return nil
}

// scoreWithReservation combine the reservation with the node's resource usage to calculate the reservation score.
func (p *Plugin) scoreWithReservation(
	allocator *AutopilotAllocator,
	state *preFilterState,
	restoreState *nodeReservationRestoreStateData,
	alloc *reusableAlloc,
	nodeName string,
	basicPreemptible map[schedulingv1alpha1.DeviceType]deviceResources,
) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func calcRequiredDeviceResources(alloc *reusableAlloc, preemptibleInRR map[schedulingv1alpha1.DeviceType]deviceResources) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

// required is empty to indicate that there are no resources left.
// A valid object must be constructed with no remaining capacity.

func (p *Plugin) allocateWithNominated(
	allocator *AutopilotAllocator,
	state *preFilterState,
	restoreState *nodeReservationRestoreStateData,
	node *corev1.Node,
	pod *corev1.Pod,
	basicPreemptible map[schedulingv1alpha1.DeviceType]deviceResources,
) (apiext.DeviceAllocations, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(apiext.DeviceAllocations), nil
}

// if the pod is reservation-ignored, it should allocate the node unallocated resources and all the reserved
// unallocated resources.

func (p *Plugin) scoreWithNominatedReservation(
	allocator *AutopilotAllocator,
	state *preFilterState,
	restoreState *nodeReservationRestoreStateData,
	nodeName string,
	pod *corev1.Pod,
	basicPreemptible map[schedulingv1alpha1.DeviceType]deviceResources,
	reservationInfo *frameworkext.ReservationInfo,
) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Plugin) getNominatedReusableAlloc(restoreState *nodeReservationRestoreStateData, pod *corev1.Pod, node *corev1.Node) ([]reusableAlloc, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isDeviceAllocationsInclude checks if allocations include all devices from required
func isDeviceAllocationsInclude(allocations apiext.DeviceAllocations, required map[schedulingv1alpha1.DeviceType]deviceResources) bool {
	_ = "STUB: not implemented"
	return false
}

// Create a map of allocated device minors

// Check if all required device minors are in allocated
