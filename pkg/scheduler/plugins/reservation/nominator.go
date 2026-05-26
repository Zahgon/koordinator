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
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	corelisters "k8s.io/client-go/listers/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	listerschedulingv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

type nominator struct {
	podLister corelisters.PodLister
	rLister   listerschedulingv1alpha1.ReservationLister
	// nominatedPodToNode is map keyed by a Pod UID to the node name where it is nominated to reserve.
	nominatedPodToNode map[types.UID]map[string]types.UID
	// nominatedPreAllocatable is map keyed by a Reservation UID to the node name where there are
	// pre-allocatable pods being nominated. Supports both single and multiple pods.
	nominatedPreAllocatable map[types.UID]map[string][]*corev1.Pod
	// nominatedReservePod is map keyed by nodeName to the nominated reservation's PodInfo for preemption.
	nominatedReservePod       map[string][]*framework.PodInfo
	nominatedReservePodToNode map[types.UID]string
	lock                      sync.RWMutex
}

func newNominator(podLister corelisters.PodLister, rLister listerschedulingv1alpha1.ReservationLister) *nominator {
	_ = "STUB: not implemented"
	return nil
}

func (nm *nominator) AddNominatedReservation(pod *corev1.Pod, nodeName string, rInfo *frameworkext.ReservationInfo) {
	_ = "STUB: not implemented"
	return
}

// nominate it only if the pod is unscheduled and reservation is active

// cannot nominate to an inactive reservation

func (nm *nominator) AddNominatedReservePod(pi *framework.PodInfo, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// Always delete the reservation if it already exists, to ensure we never store more than
// one instance of the reservation.

// not a reserve pod

// If a reservation is just deleted or scheduled, don't nominate it.

func (nm *nominator) AddNominatedPreAllocation(rInfo *frameworkext.ReservationInfo, nodeName string, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// nominate it only if the pod is unscheduled and reservation is active

func (nm *nominator) NominatedReservePodForNode(nodeName string) []*framework.PodInfo {
	_ = "STUB: not implemented"
	return nil
}

// Make a copy of the nominated Pods so the caller can mutate safely.

func (nm *nominator) DeleteReservePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (nm *nominator) deleteReservePod(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	// delete pre-allocation for the reservation if exists
	return
}

// RemoveNominatedReservation removes the nominated reservation of a pod from the nominator.
func (nm *nominator) RemoveNominatedReservation(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (nm *nominator) RemoveNominatedPreAllocation(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// DeleteNominatedReservePodOrReservation is used to delete the nominated reserve pod or
// the nominated reservation for the pod.
func (nm *nominator) DeleteNominatedReservePodOrReservation(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (nm *nominator) GetNominatedReservation(pod *corev1.Pod, nodeName string) types.UID {
	_ = "STUB: not implemented"
	return *new(types.UID)
}

func (nm *nominator) GetNominatedPreAllocation(rInfo *frameworkext.ReservationInfo, nodeName string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (nm *nominator) AddNominatedPreAllocations(rInfo *frameworkext.ReservationInfo, nodeName string, pods []*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// Validate all pods are assigned to the node

func (nm *nominator) GetNominatedPreAllocations(rInfo *frameworkext.ReservationInfo, nodeName string) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// TODO(joseph): Should move the function into frameworkext package as default nominator

func (pl *Plugin) NominateReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) (*frameworkext.ReservationInfo, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pl *Plugin) AddNominatedReservation(pod *corev1.Pod, nodeName string, rInfo *frameworkext.ReservationInfo) {
	_ = "STUB: not implemented"
	return
}

// RemoveNominatedReservations is used to delete the nominated reserve pod.
// DEPRECATED: use DeleteNominatedReservePodOrReservation instead.
func (pl *Plugin) RemoveNominatedReservations(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (pl *Plugin) AddNominatedReservePod(pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// DeleteNominatedReservePodOrReservation is used to delete the nominated reserve pod or
// the nominated reservation for the pod.
func (pl *Plugin) DeleteNominatedReservePodOrReservation(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// DeleteNominatedReservePod is used to delete the nominated reserve pod.
// DEPRECATED: use DeleteNominatedReservePodOrReservation instead.
func (pl *Plugin) DeleteNominatedReservePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (pl *Plugin) NominatedReservePodForNode(nodeName string) []*framework.PodInfo {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) GetNominatedReservation(pod *corev1.Pod, nodeName string) *frameworkext.ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) NominatePreAllocation(ctx context.Context, cycleState fwktype.CycleState, rInfo *frameworkext.ReservationInfo, nodeName string) (*corev1.Pod, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pl *Plugin) AddNominatedPreAllocation(rInfo *frameworkext.ReservationInfo, nodeName string, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (pl *Plugin) GetNominatedPreAllocation(rInfo *frameworkext.ReservationInfo, nodeName string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) AddNominatedPreAllocations(rInfo *frameworkext.ReservationInfo, nodeName string, pods []*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (pl *Plugin) GetNominatedPreAllocations(rInfo *frameworkext.ReservationInfo, nodeName string) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// NominatePreAllocations nominates multiple pre-allocatable pods for a reservation.
// It accumulates resources from pods until all dimensions are satisfied.
func (pl *Plugin) NominatePreAllocations(cycleState fwktype.CycleState, rInfo *frameworkext.ReservationInfo, nodeName string) ([]*corev1.Pod, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prioritizeReservations(
	ctx context.Context,
	fwk frameworkext.FrameworkExtender,
	state fwktype.CycleState,
	pod *corev1.Pod,
	reservations []*frameworkext.ReservationInfo,
	nodeName string,
) (frameworkext.ReservationScoreList, error) {
	_ = "STUB: not implemented"
	return *new(frameworkext.ReservationScoreList), nil
}

// Summarize all scores.

func prioritizePreAllocatablePods(
	ctx context.Context,
	fwk frameworkext.FrameworkExtender,
	state fwktype.CycleState,
	rInfo *frameworkext.ReservationInfo,
	pods []*corev1.Pod,
	nodeName string,
) (frameworkext.ReservationScoreList, error) {
	_ = "STUB: not implemented"
	return *new(frameworkext.ReservationScoreList), nil
}

// Summarize all scores.
