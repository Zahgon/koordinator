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
	"k8s.io/apimachinery/pkg/runtime"
	listercorev1 "k8s.io/client-go/listers/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	clientschedulingv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	listerschedulingv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	Name     = "Reservation"
	stateKey = Name

	// ErrReasonReservationAffinity is the reason for Pod's reservation affinity/selector not matching.
	ErrReasonReservationAffinity = "node(s) no reservations match reservation affinity"
	// ErrReasonNodeNotMatchReservation is the reason for node not matching which the reserve pod specifies.
	ErrReasonNodeNotMatchReservation = "node(s) didn't match the nodeName specified by reservation"
	// ErrReasonReservationAllocatePolicyConflict is the reason for the AllocatePolicy of the Reservation conflicts with other Reservations on the node
	ErrReasonReservationAllocatePolicyConflict = "node(s) reservation allocate policy conflict"
	// ErrReasonReservationInactive is the reason for the reservation is failed/succeeded and should not be used.
	ErrReasonReservationInactive = "reservation is not active"
	// ErrReasonNoReservationsMeetRequirements is the reason for no reservation(s) to meet the requirements.
	ErrReasonNoReservationsMeetRequirements = "node(s) no reservation(s) to meet the requirements"
	// ErrReasonReservationPreAllocationRequired is the reason for the Reservation PreAllocationRequired not matching.
	ErrReasonReservationPreAllocationRequired = "node(s) no pod match reservation with pre allocation required"
	// ErrReasonReservationPreAllocationUnsupported is the reason for the Reservation PreAllocation does not support some attributes.
	ErrReasonReservationPreAllocationUnsupported = "node(s) reservation pre allocation is unsupported with current attributes"
	// ErrReasonNoPodsMeetPreAllocationRequirements is the reason for no pod(s) to meet the pre-allocation requirements.
	ErrReasonNoPodsMeetPreAllocationRequirements = "node(s) no pod(s) to meet the pre-allocation requirements"
	// ErrReasonPreemptionFailed is the reason for preemption failed
	ErrReasonPreemptionFailed = "node(s) preemption failed due to insufficient resources"
	// ErrReasonReservationClusterModeDisabled is the reason for cluster mode pre-allocation is disabled
	ErrReasonReservationClusterModeDisabled = "reservation requires cluster mode pre-allocation which is disabled"
)

var (
	_ fwktype.EnqueueExtensions = &Plugin{}

	_ fwktype.PreFilterPlugin  = &Plugin{}
	_ fwktype.FilterPlugin     = &Plugin{}
	_ fwktype.PostFilterPlugin = &Plugin{}
	_ fwktype.ScorePlugin      = &Plugin{}
	_ fwktype.ReservePlugin    = &Plugin{}
	_ fwktype.PreBindPlugin    = &Plugin{}
	_ fwktype.BindPlugin       = &Plugin{}

	_ frameworkext.ControllerProvider      = &Plugin{}
	_ frameworkext.PreFilterTransformer    = &Plugin{}
	_ frameworkext.ReservationCache        = &Plugin{}
	_ frameworkext.ReservationNominator    = &Plugin{}
	_ frameworkext.ReservationFilterPlugin = &Plugin{}
	_ frameworkext.ReservationScorePlugin  = &Plugin{}
)

type Plugin struct {
	handle                         frameworkext.ExtendedHandle
	args                           *config.ReservationArgs
	rLister                        listerschedulingv1alpha1.ReservationLister
	podLister                      listercorev1.PodLister
	client                         clientschedulingv1alpha1.SchedulingV1alpha1Interface
	reservationCache               *reservationCache
	nominator                      *nominator
	preemptionMgr                  *PreemptionMgr
	enableLazyReservationRestore   bool
	enableSkipReservationFitsNode  bool
	enablePreAllocationClusterMode bool
}

func New(_ context.Context, args runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (pl *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (pl *Plugin) NewControllers() ([]frameworkext.Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pl *Plugin) EventsToRegister(_ context.Context) ([]fwktype.ClusterEventWithHint, error) {
	_ = "STUB: not implemented"
	// To register a custom event, follow the naming convention at:
	// https://github.com/kubernetes/kubernetes/blob/e1ad9bee5bba8fbe85a6bf6201379ce8b1a611b1/pkg/scheduler/eventhandlers.go#L415-L422
	return nil, nil
}

// PreFilter checks if the pod is a reserve pod. If it is, update cycle state to annotate reservation scheduling.
// Also do validations in this phase.
func (pl *Plugin) PreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate reserve pod and reservation

// check if pre-allocation requirement not meet

// nor available reservation neither a reserve pod

func (pl *Plugin) PreBindPreFlight(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) PreFilterExtensions() fwktype.PreFilterExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.PreFilterExtensions)
}

func (pl *Plugin) AddPod(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, podInfoToAdd fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// count pods resources

func (pl *Plugin) RemovePod(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, podInfoToRemove fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// count pods resources

// Filter only processes pods either the pod is a reserve pod or a pod can allocate reserved resources on the node.
func (pl *Plugin) Filter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// if the reservation specifies a nodeName initially, check if the nodeName matches

// ReservationAllocatePolicyDefault cannot coexist with other allocate policies

// handle pre-allocation cases

// Return Unschedulable (not UnschedulableAndUnresolvable) so that the preemption evaluator
// can consider this node as a potential preemption candidate via NodesForStatusCode(Unschedulable).

func (pl *Plugin) filterWithReservations(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo, matchedReservations []*frameworkext.ReservationInfo, requiredFromReservation bool) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// For reservation-ignored pods, the transformer has already restored all reservation resources back to the node.
// NodeResourceFit plugin will validate if the node has sufficient resources for the pod.
// So we can skip the resource validation here and return success directly.

// Making resource list and framework.Resource is heavy, skip it when there is no preemptible resources.

// Contextualization: If a pod only have one reservation matched, the fitsNode should be equivalent to
// the NodeResourceFit's Filter, so we can skip the fitsNode to reduce overhead.

// NOTE: The reservation may not consider the irrelevant pods that have no matched resource names since it makes
// no sense in most cases but introduces a performance overhead. However, we allow pods to allocate reserved
// resources to accomplish their reservation affinities.

// When the pod specifies a reservation name, we record the admission reasons.

// Making resource list and framework.Resource is heavy, skip it when there is no preemptible resources.

// Before nominating a reservation in PreScore or Reserve, check the reservation by multiple plugins to make
// the Filter phase give a more accurate result. It is extensible to support more policies.

// The Pod requirement must be allocated from Reservation, but currently no Reservation meets the requirement.
// We will keep all failure reasons.

// If the combination of reservation and node cannot satisfy the pod, then the node alone cannot satisfy it either.

// try to allocate from node alone

func (pl *Plugin) filterWithPreAllocatablePods(ctx context.Context, cycleState fwktype.CycleState, rInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo, preAllocatablePods []*corev1.Pod, isPreAllocationRequired bool) (selectedPreAllocatablePods []*corev1.Pod, result *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filterWithPreAllocatablePod checks if the reservation can fit the node with or without at most one pre-allocated pod.
func (pl *Plugin) filterWithPreAllocatablePod(ctx context.Context, cycleState fwktype.CycleState, rInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo, preAllocatablePods []*corev1.Pod, isPreAllocationRequired bool) (selectedPreAllocatablePods []*corev1.Pod, result *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check node-unallocated resources only once

// use nodeInfo.GetRequested() if nodeRState.podRequested is nil (when there is no pre-allocatable pods).

// Check if the reserve pod can be placed with node-unallocated resource when pre-allocation is not required.
// For reserve pod, matchedOrIgnored should be 0, both rAllocated and rRemained should be nil.

// Fast-path: check node-unallocated resources first when pre-allocation is not required

// Directly return if the reservation can place without any pre-allocatable pod

// To filter a reservation pre-allocate with the pod:
// (0) No need to check if the pod can place to the node without reservation.
// (1) Reservation Restricted policy: Check if the pod can place into the reservation:
//     podRequest <= rRequest - rAllocated
// (2) Reservation PreAllocation: Check if the reservation can place without the preAllocatable pod:
//     rRequest <= nodeAllocatable - (allPodsRequested - podRequest - allRAllocated - preemptible)
// (3) Pod allocate Reservation: Check if the pod can place with reserved free resources:
//     podRequest <= nodeAllocatable - (allPodsRequested - allRAllocated - rRemained - preemptible)
// Where preemptible > 0 when there is victim pods can be preempted.

// 1. Check if the reservation can place into the node if pod uses the reserved resource.
// 2. Check if the pod can place into the reservation.

// Before nominating a reservation in PreScore or Reserve, check the reservation by multiple plugins to make
// the Filter phase give a more accurate result. It is extensible to support more policies.

// Build failure reasons

// If the combination of pre-allocatable and node cannot satisfy the reservation, then the node alone cannot satisfy it either.

// If node-unallocated resources have already been checked, reuse the reasons directly.

// If no insufficient resources by node with pre-allocatable pods, use the node-unallocated reasons,
// avoid passing in Filter phase but later fail in Reserve phase.
// When no pre-allocatable pods exist, NodeResourcesFit performs the node-unallocated check.

// filterWithMultiplePreAllocatablePods checks if the reservation can fit the node with or without multiple pre-allocated pods.
func (pl *Plugin) filterWithMultiplePreAllocatablePods(ctx context.Context, cycleState fwktype.CycleState,
	rInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo, preAllocatablePods []*corev1.Pod,
	isPreAllocationRequired bool) (selectedPreAllocatablePods []*corev1.Pod, result *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For default mode, pre-allocatable pods should be sorted before selecting.
// Not for cluster mode since pre-allocatable pods has already been sorted in cache.

// Check node-unallocated resources only once

// use nodeInfo.Requested if nodeRState.podRequested is nil (when there is no pre-allocatable pods).

// Check if the reserve pod can be placed with node-unallocated resource when pre-allocation is not required.
// For reserve pod, matchedOrIgnored should be 0, both rAllocated and rRemained should be nil.

// Fast-path: check node-unallocated resources first when pre-allocation is not required

// Directly return if the reservation can place without any pre-allocatable pod

// Accumulate resources from pods until all dimensions are satisfied:
// 	 skip any pod that causes any dimension to exceed, only count pods where all dimensions fit.

// Try accumulating this pod's resources

// Calculate pod requested without the trial accumulated requests

// Check if adding this pod still fits

// If any dimension exceeds, skip this pod

// Record reasons for later diagnosis

// Skip this pod, do not add to accumulated resources

// Run reservation filter plugins for this pod

// Skip this pod if it doesn't pass filter plugins

// This pod fits, add it to selected pods and accumulate resources

// No dimension exceeds, update insufficient resources by node
// All dimensions are satisfied, update selected pre-allocatable pods

// Build failure reasons

// If the combination of pre-allocatable and node cannot satisfy the reservation, then the node alone cannot satisfy it either.

// If node-unallocated resources have already been checked, reuse the reasons directly.

// If no insufficient resources by node with pre-allocatable pods, use the node-unallocated reasons,
// avoid passing in Filter phase but later fail in Reserve phase.
// When no pre-allocatable pods exist, NodeResourcesFit performs the node-unallocated check.

// sortPreAllocatablePodsForDefaultMode sorts pre-allocatable pods by priority if needed.
func sortPreAllocatablePodsForDefaultMode(ctx context.Context, extender frameworkext.FrameworkExtender,
	cycleState fwktype.CycleState, rInfo *frameworkext.ReservationInfo,
	preAllocatablePods []*corev1.Pod, nodeName string) ([]*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sort by priority descending

var dummyResource = framework.NewResource(nil)

func fitsNodeAndReservation(podRequestsResources, allPodsRequested, allRAllocated, preemptible, rRemained fwktype.Resource,
	podRequests, preemptibleInRR corev1.ResourceList, pod *corev1.Pod, rInfo *frameworkext.ReservationInfo,
	nodeInfo fwktype.NodeInfo, matchedCount int, requireDetailReasons, isFitsNodeSkipped bool) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fit the reservation

// fitsNode checks if node have enough resources to host the pod.
func fitsNode(podRequest, nodeAllocatable, allPodsRequested, allRAllocated, rRemained fwktype.Resource, matchedOrIgnored, allocatedPods int, preemptible fwktype.Resource) []string {
	_ = "STUB: not implemented"
	return nil
}

func fitsReservation(podRequest corev1.ResourceList, rInfo *frameworkext.ReservationInfo, preemptibleInRR corev1.ResourceList, isDetailed bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// check "pods" resource in the reservation when reserved explicitly

// assert no overflow

// print a reason with resource amounts if needed

// keep allocated >= 0

// NOTE: capacity excludes the reserved resource

// just give the resource name

// print a reason with resource amounts if needed

// print a reason with resource amounts if needed

func (pl *Plugin) PostFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, filteredNodeStatusMap fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If Reservation Preemption is enabled, try preemption before aggregating failure reasons.

func (pl *Plugin) makePostFilterReasons(state *stateData, filteredNodeStatusMap fwktype.NodeToStatusReader) []string {
	_ = "STUB: not implemented"
	return nil
}

// failure reasons and counts for the nodes which have not been handled by the Reservation's Filter

// summarize node diagnosis states

// calculate the remaining unmatched which is owner-matched and Reservation BeforePreFilter matched

// no need to check other reasons

// count the failure reasons which is neither counted by the PreFilterTransformer nor by the Reservation Filter.

// reservation-level reasons are not counted

// capped with zero

// if the pod specifies an affinity, we always show the owner matched reason

// Make the error messages: The framework does not aggregate the same reasons for the PostFilter, so we need
// to prepare the exact messages by ourselves.

// node reason Filter failed

func (pl *Plugin) FilterReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) FilterNominateReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, rInfo *frameworkext.ReservationInfo, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	// TODO(joseph): We can consider optimizing these codes. It seems that there is no need to exist at present.
	return nil
}

// For a PreAllocation, needs to filter the pre-allocatable pod for the reservation.

func (pl *Plugin) ReservationNominate(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// reservation
// If the pod represents to a pre-allocation reservation, we need to nominate a pre-allocatable pod for it.
// Otherwise, assuming the reservation is enough for a non-pre-allocation reservation.

// Check if multiple pre-allocation is enabled

// Single pre-allocation mode (original logic)

// normal pod

// The scheduleOne skip scores and reservation nomination if there is only one node available.

func (pl *Plugin) Reserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// clean scheduling cycle to avoid unnecessary memory cost before entering the binding

// reservation

// If the pod represents to a pre-allocation reservation, we need to nominate a pre-allocatable pod for it.
// Otherwise, assuming the reservation is enough for a non-pre-allocation reservation.

// Check if multiple pre-allocation is enabled

// Single pre-allocation mode (original logic)

// Store all nominated pods in state for PreBind phase

// normal pod

// The scheduleOne skip scores and reservation nomination if there is only one node available.

func (pl *Plugin) Unreserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// If the reservation has been deleted or the lister fails, construct a temporary reservation to clean up the cache.

// reserve pod without pre-allocation

// If the reserve pod is in pre-allocation, need to clean the nominator and the pod annotation.

// no reservation is assumed

// clean the assumed in cache

// clean the reservation-allocated annotation of the allocated pod
// no reservation-allocated has set

func (pl *Plugin) unreservePod(ctx context.Context, pod, allocatedPod *corev1.Pod,
	rInfo *frameworkext.ReservationInfo, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// avoid modifying a homonymous pod

// In some corner cases, the pod could fail the Bind request but finally assigned to the node.
// To protect the ownership between the assigned pod and the reservation and avoid cache leak,
// we keep the reservation-allocated annotation.

// no need to fix annotation

func (pl *Plugin) PreBind(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) PreBindReservation(ctx context.Context, cycleState fwktype.CycleState, reservation *schedulingv1alpha1.Reservation, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// PreAllocation disable but has a preAllocated result

// expect a pre-allocation but no reservation is assumed

// Apply PreAllocation result to the pod(s)

// Apply PreAllocation annotation to all nominated pods

// avoid modify a homonymous pod

// Bind fake binds reserve pod and mark corresponding reservation as Available.
// NOTE: This Bind plugin should get called before DefaultBinder; plugin order should be configured.
func (pl *Plugin) Bind(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// check if the reservation has been inactive

// mark reservation as available

func (pl *Plugin) DeleteReservation(r *schedulingv1alpha1.Reservation) *frameworkext.ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) GetReservationInfoByPod(pod *corev1.Pod, nodeName string) *frameworkext.ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

// IsPreferNoPreAllocatedPods returns whether to prefer placing reservation
// without pre-allocated pods when pre-allocation is not required and node unallocated resources are sufficient.
// Defaults to false.
func (pl *Plugin) IsPreferNoPreAllocatedPods() bool { _ = "STUB: not implemented"; return false }
