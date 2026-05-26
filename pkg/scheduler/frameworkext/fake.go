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

package frameworkext

import (
	"context"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// nominatedPodMap is a structure that stores pods nominated to run on nodes.
// It exists because nominatedNodeName of pod objects stored in the structure
// may be different than what scheduler has here. We should be able to find pods
// by their UID and update/delete them.
type nominatedPodMap struct {
	// nominatedPods is a map keyed by a node name and the value is a list of
	// pods which are nominated to run on the node. These are pods which can be in
	// the activeQ or unschedulableQ.
	nominatedPods map[string][]fwktype.PodInfo
	// nominatedPodToNode is map keyed by a Pod UID to the node name where it is
	// nominated.
	nominatedPodToNode map[types.UID]string

	sync.RWMutex
}

// NewFakePodNominator creates a nominatedPodMap as a backing of fwktype.PodNominator.
func NewFakePodNominator() fwktype.PodNominator {
	_ = "STUB: not implemented"
	return *new(fwktype.PodNominator)
}

func (npm *nominatedPodMap) add(pi fwktype.PodInfo, nodeName string) {
	_ = "STUB: not implemented"
	// always delete the pod if it already exist, to ensure we never store more than
	// one instance of the pod.
	return
}

func (npm *nominatedPodMap) delete(p *corev1.Pod) { _ = "STUB: not implemented"; return }

// UpdateNominatedPod updates the <oldPod> with <newPod>.
func (npm *nominatedPodMap) UpdateNominatedPod(logr klog.Logger, oldPod *corev1.Pod, newPodInfo fwktype.PodInfo) {
	_ = "STUB: not implemented"
	return
}

// In some cases, an Update event with no "NominatedNode" present is received right
// after a node("NominatedNode") is reserved for this pod in memory.
// In this case, we need to keep reserving the NominatedNode when updating the pod pointer.

// We won't fall into below `if` block if the Update event represents:
// (1) NominatedNode info is added
// (2) NominatedNode info is updated
// (3) NominatedNode info is removed

// This is the only case we should continue reserving the NominatedNode

// We update irrespective of the nominatedNodeName changed or not, to ensure
// that pod pointer is updated.

// DeleteNominatedPodIfExists deletes <pod> from nominatedPods.
func (npm *nominatedPodMap) DeleteNominatedPodIfExists(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// AddNominatedPod adds a pod to the nominated pods of the given node.
// This is called during the preemption process after a node is nominated to run
// the pod. We update the structure before sending a request to update the pod
// object to avoid races with the following scheduling cycles.
func (npm *nominatedPodMap) AddNominatedPod(logger klog.Logger, pi fwktype.PodInfo, nominatingInfo *fwktype.NominatingInfo) {
	_ = "STUB: not implemented"
	return
}

// NominatedPodsForNode returns pods that are nominated to run on the given node,
// but they are waiting for other pods to be removed from the node.
func (npm *nominatedPodMap) NominatedPodsForNode(nodeName string) []fwktype.PodInfo {
	_ = "STUB: not implemented"
	return nil
}

// TODO: we may need to return a copy of []PodInfo to avoid modification
// on the caller side.

var _ ReservationNominator = &FakeNominator{}

type FakeNominator struct {
	lock sync.RWMutex
	// nominatedPodToNode is map keyed by a Pod UID to the node name where it is nominated.
	nominatedPodToNode map[types.UID]map[string]types.UID
	reservations       map[types.UID]*ReservationInfo
	preAllocatable     map[types.UID]map[string][]*corev1.Pod
	// nominatedReservePod is map keyed by nodeName, value is the nominated reservations
	nominatedReservePod       map[string][]*framework.PodInfo
	nominatedReservePodToNode map[types.UID]string
}

func NewFakeReservationNominator() *FakeNominator { _ = "STUB: not implemented"; return nil }

func (nm *FakeNominator) Name() string { _ = "STUB: not implemented"; return "" }

func (nm *FakeNominator) AddNominatedReservation(pod *corev1.Pod, nodeName string, rInfo *ReservationInfo) {
	_ = "STUB: not implemented"
	return
}

func (nm *FakeNominator) RemoveNominatedReservations(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (nm *FakeNominator) NominatedReservePodForNode(nodeName string) []*framework.PodInfo {
	_ = "STUB: not implemented"
	return nil
}

func (nm *FakeNominator) GetNominatedReservation(pod *corev1.Pod, nodeName string) *ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (nm *FakeNominator) NominateReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) (*ReservationInfo, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nm *FakeNominator) AddNominatedReservePod(pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// Always delete the reservation if it already exists, to ensure we never store more than
// one instance of the reservation.

func (nm *FakeNominator) DeleteNominatedReservePod(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (nm *FakeNominator) deleteReservePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (nm *FakeNominator) NominatePreAllocation(ctx context.Context, cycleState fwktype.CycleState, rInfo *ReservationInfo, nodeName string) (*corev1.Pod, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nm *FakeNominator) AddNominatedPreAllocation(rInfo *ReservationInfo, nodeName string, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (nm *FakeNominator) GetNominatedPreAllocation(rInfo *ReservationInfo, nodeName string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (nm *FakeNominator) deletePreAllocation(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (nm *FakeNominator) RemoveNominatedPreAllocation(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (nm *FakeNominator) AddNominatedPreAllocations(rInfo *ReservationInfo, nodeName string, pods []*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (nm *FakeNominator) GetNominatedPreAllocations(rInfo *ReservationInfo, nodeName string) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// GetNominatedNodeForReservePod returns the node name that the reserve pod is nominated to.
// This is only for testing purposes.
func (nm *FakeNominator) GetNominatedNodeForReservePod(pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}

func (nm *FakeNominator) ReservationNominate(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (nm *FakeNominator) DeleteNominatedReservePodOrReservation(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}
