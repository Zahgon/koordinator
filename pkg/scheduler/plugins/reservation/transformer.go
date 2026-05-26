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

package reservation

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	reservationutil "github.com/koordinator-sh/koordinator/pkg/util/reservation"
)

func (pl *Plugin) BeforePreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) (*corev1.Pod, bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (pl *Plugin) AfterPreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, preRes *fwktype.PreFilterResult) *fwktype.Status {
	_ = "STUB: not implemented"
	// Since restoring reserved resources is inefficient, the feature handles it with PreFilter result to reduce the
	// unnecessary restorations.
	return nil
}

// use the hint nodes if exists

// use the PreFilter result if exists

// If the PreFilterResult does not restrict the scope, we try all the valid nodeReservationStates to restore.

func (pl *Plugin) prepareMatchReservationStateForNormalPod(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) (*stateData, bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// check if the node-level preRestore is required for all nodes in the BeforePreFilter

// global preRestore

// use the hint nodes if exists

// check if the reservation matches or can be ignored by the pod

// reservation is matched or ignored for the pod

// reservation is unmatched and not ignored

// The Pod declares a ReservationAffinity, which means that the Pod must reuse the Reservation resources,
// but there are no matching Reservations, which means that the node itself does not need to be processed.
// We can end early to avoid meaningless operations.

// LazyReservationRestore indicates whether to restore reserved resources for the scheduling pod lazily.
// If it is disabled, the reserved resources are ensured to restore in BeforePreFilter/PreFilter phase, where all
// nodes related to reservations will restore reserved resources and refresh node snapshots in the next cycle.
// If it is enabled, the reserved resources are delayed to restore in Filter phase when the pod does not specify
// any pod affinity/anti-affinity or topologySpreadConstraints, it can reduce resource restoration overhead
// especially when there are a large scale of reservations. However, it does not ensure the correctness of the
// existing pod affinities, so it is disabled by default.

// the pre restoration is required in the BeforePreFilter

func (pl *Plugin) prepareMatchReservationStateForReservePod(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) (*stateData, bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// check if the node-level preRestore is required for all nodes in the BeforePreFilter

// by default list all nodes existing a reservation
// merge with pre-allocatable pods for the pre-allocation

// only care about allocated reservations to restore

// global preRestore

// NOTE: Currently, the PreAllocation is only available on a Restricted Reservation.

// if no hint, merge nodes with pre-allocatable

// handle pre-allocation

// The reservation specifies required to pre-allocation, but there is no pre-allocatable pods on the node.

// LazyReservationRestore indicates whether to restore reserved resources for the scheduling pod lazily.
// If it is disabled, the reserved resources are ensured to restore in BeforePreFilter/PreFilter phase, where all
// nodes related to reservations will restore reserved resources and refresh node snapshots in the next cycle.
// If it is enabled, the reserved resources are delayed to restore in Filter phase when the pod does not specify
// any pod affinity/anti-affinity or topologySpreadConstraints, it can reduce resource restoration overhead
// especially when there are a large scale of reservations. However, it does not ensure the correctness of the
// existing pod affinities, so it is disabled by default.

// the pre restoration is required in the BeforePreFilter

func (pl *Plugin) listPreAllocatableCandidates(preAllocationMode schedulingv1alpha1.PreAllocationMode,
	rInfo *frameworkext.ReservationInfo) (map[string][]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	// Cluster mode: retrieve from cache which already has sorted candidates
	return nil, nil
}

// In cluster mode, candidates are cached and sorted in reservationCache
// We retrieve them directly from cache instead of listing from podLister
// The cache is maintained by pod event handlers and updated incrementally
// Returns early with cached data (already grouped by node and sorted)

// Default mode: use OwnerMatchers

// TODO: Reduce the overhead of the finding pre-allocatable pods.

// FIXME: This step also list the unassigned pods.

// checkReservationMatchedOrIgnored checks if the reservation is matched or can be ignored by the pod and
// updates the node diagnosis states.
func checkReservationMatchedOrIgnored(pod *corev1.Pod, rInfo *frameworkext.ReservationInfo, diagnosisState *nodeDiagnosisState, node *corev1.Node, podRequests corev1.ResourceList,
	reservationAffinity *reservationutil.RequiredReservationAffinity, exactMatchReservationSpec *extension.ExactMatchReservationSpec, affinityReservationName string, isReservationIgnored bool) bool {
	_ = "STUB: not implemented"
	// pod specifies reservation ignored
	return false
}

// pod matches the reservation owners

// Check the conditions by the following order:
// 1. check the condition if it has a higher priority than others
// 2. check the more common and fast conditions
// 3. check the more complex conditions

// If reservation name is specified, no longer check the unschedulable, affinity and taints.

// Actually, the reservation name should be unique in the cluster. So if the pod specifies the
// name, only the name matched reservation will check the conditions below.

// exactMatchSpec unmatched

// name matched

// isUnschedulable

// taints not tolerated
// TODO: support effect=PreferNoSchedule

// ReservationAffinity unmatched

// exactMatchSpec unmatched

// matched

func checkPreAllocatableMatched(preAllocationMode schedulingv1alpha1.PreAllocationMode,
	rInfo *frameworkext.ReservationInfo, candidatePod *corev1.Pod, diagnosisState *nodeDiagnosisState,
	node *corev1.Node) (bool, error) {
	_ = "STUB: not implemented"
	// check if candidate pod matches the reservation, the matching logic must differ based on the pre-allocation mode:
	//   - Default mode: OwnerMatchers (ObjectRef, Controller, Labels) check is placed here for performance reasons.
	//   - Cluster mode: Pods are retrieved from a global cache that already contains sorted pre-allocatable candidates
	//  				 grouped by node, so that the OwnerMatchers check is NOT needed by design.
	return false, nil
}

// If reservation name is specified, no longer check the unschedulable, affinity and taints.

// taints not tolerated

// ReservationAffinity unmatched

// exactMatchSpec unmatched

// matched

func preRestoreReservationResourcesForNode(logger klog.Logger, extender frameworkext.FrameworkExtender, pod *corev1.Pod,
	rInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo, nodeRState *nodeReservationState, skipRestoreNodeInfo bool) error {
	_ = "STUB: not implemented"
	return nil
}

// When the nodeInfo restore is skipped, we only record for the reservation restore state.

// no more pre-restore in the same cycle

// Save requested state after trimmed by unmatched to support reservation allocate policy.

// no more pre-restore in the same cycle

func restoreReservationResourcesForNode(ctx context.Context, cycleState fwktype.CycleState, extender frameworkext.FrameworkExtender,
	pod *corev1.Pod, rInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo, nodeRState *nodeReservationState, skipRestoreNodeInfo bool) (bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

// Some attributes like podAffinity and topologySpreadConstraints is pre-processed in the PreFilter phase,
// so we cannot delay the restoration to the Filter.

// no more restore in the same cycle

func restoreMatchedReservation(nodeInfo fwktype.NodeInfo, rInfo *frameworkext.ReservationInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Retain ports that are not used by other Pods. These ports need to be erased from NodeInfo.UsedPorts,
// otherwise it may cause Pod port conflicts

// When AllocateOnce is disabled, some resources may have been allocated,
// and an additional resource record will be accumulated at this time.
// Even if the Reservation is not bound by the Pod (e.g. Reservation is enabled with AllocateOnce),
// these resources held by the Reservation need to be returned, to ensure that
// the Pod can pass through each filter plugin during scheduling.
// The returned resources include scalar resources such as CPU/Memory, ports etc..

func restoreUnmatchedReservations(nodeInfo fwktype.NodeInfo, rInfo *frameworkext.ReservationInfo) error {
	_ = "STUB: not implemented"
	// Here len(rInfo.AssignedPods) == 0 is always false because it was checked before.
	return nil
}

// Reservations and Pods that consume the Reservations are cumulative in resource accounting.
// For example, on a 32C machine, ReservationA reserves 8C, and then PodA uses ReservationA to allocate 4C,
// then the record on NodeInfo is that 12C is allocated. But in fact it should be calculated according to 8C,
// so we need to return some resources.

func restorePreAllocatablePods(nodeInfo fwktype.NodeInfo, rInfo *frameworkext.ReservationInfo, preAllocatablePod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func isPodAllNodesPreRestoreRequired(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	// If a pod specifies required topologySpreadConstraints, podAffinities and podAntiAffinities, we should do the
	// node-level preRestore in the BeforePreFilter for each node, even when the LazyReservationRestore is enabled.
	// FIXME: The existing podAffinities/podAntiAffinities on nodes are not considered.
	return false
}

func updateNodeInfoRequestedForUnmatched(n fwktype.NodeInfo, rInfo *frameworkext.ReservationInfo) {
	_ = "STUB: not implemented"
	return
}

func parseSpecificNodesFromAffinity(pod *corev1.Pod) (sets.String, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(sets.String), nil
}

// Check if there is affinity to a specific node and return it.

// The requirements represent ANDed constraints, and so we need to
// find the intersection of nodes.

// If this term has no node.Name field affinity,
// then all nodes are eligible because the terms are ORed.

// If the set is empty, it means the terms had affinity to different
// sets of nodes, and since they are ANDed, then the pod will not match any node.

func getDiagnosisTaintKey(taint *corev1.Taint) string { _ = "STUB: not implemented"; return "" }

func (pl *Plugin) BeforeFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) (*corev1.Pod, fwktype.NodeInfo, bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	// Both the reserve pod or the normal pod should consider the nominated reserve pods.
	return nil, *new(fwktype.NodeInfo), false, nil
}

// This may happen only in tests.

// Ignore the nominated reservations of the same job
