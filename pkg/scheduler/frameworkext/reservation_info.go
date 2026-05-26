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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	reservationutil "github.com/koordinator-sh/koordinator/pkg/util/reservation"
)

type ReservationInfo struct {
	Reservation           *schedulingv1alpha1.Reservation
	Pod                   *corev1.Pod
	ResourceNames         []corev1.ResourceName
	Allocatable           corev1.ResourceList // RO
	Allocated             corev1.ResourceList // RO
	Reserved              corev1.ResourceList // reserved inside the reservation
	Available             *framework.Resource // pre-calculated info: Allocatable - Reserved - Allocated
	AllocatedResource     *framework.Resource // pre-calculated info: Allocated
	Non0AllocatedMilliCPU int64               // pre-calculated info: non-zero milli-CPU of Allocated
	Non0AllocatedMem      int64               // pre-calculated info: non-zero Memory of Allocated
	AllocatablePorts      fwktype.HostPortInfo
	AllocatedPorts        fwktype.HostPortInfo
	AssignedPods          map[types.UID]*PodRequirement
	OwnerMatchers         []reservationutil.ReservationOwnerMatcher
	ParseError            error
}

type PodRequirement struct {
	Namespace string
	Name      string
	UID       types.UID
	Requests  corev1.ResourceList
	Ports     fwktype.HostPortInfo
}

func NewPodRequirement(pod *corev1.Pod) *PodRequirement { _ = "STUB: not implemented"; return nil }

func (p *PodRequirement) Clone() *PodRequirement { _ = "STUB: not implemented"; return nil }

func NewReservationInfo(r *schedulingv1alpha1.Reservation) *ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func NewReservationInfoFromPod(pod *corev1.Pod) *ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (ri *ReservationInfo) GetName() string { _ = "STUB: not implemented"; return "" }

func (ri *ReservationInfo) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (ri *ReservationInfo) UID() types.UID { _ = "STUB: not implemented"; return *new(types.UID) }

func (ri *ReservationInfo) GetObject() metav1.Object {
	_ = "STUB: not implemented"
	return *new(metav1.Object)
}

func (ri *ReservationInfo) GetReservePod() *corev1.Pod { _ = "STUB: not implemented"; return nil }

func (ri *ReservationInfo) GetNodeName() string { _ = "STUB: not implemented"; return "" }

func (ri *ReservationInfo) IsAllocateOnce() bool { _ = "STUB: not implemented"; return false }

// Reservation Operating Mode Pod MUST BE AllocateOnce

func (ri *ReservationInfo) GetAllocatePolicy() schedulingv1alpha1.ReservationAllocatePolicy {
	_ = "STUB: not implemented"
	return *new(schedulingv1alpha1.ReservationAllocatePolicy)
}

func (ri *ReservationInfo) GetPriority() int32 { _ = "STUB: not implemented"; return 0 }

func (ri *ReservationInfo) GetAllocatedPods() int { _ = "STUB: not implemented"; return 0 }

func (ri *ReservationInfo) GetPodOwners() []schedulingv1alpha1.ReservationOwner {
	_ = "STUB: not implemented"
	return nil
}

func (ri *ReservationInfo) MatchOwners(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (ri *ReservationInfo) IsAvailable() bool { _ = "STUB: not implemented"; return false }

func (ri *ReservationInfo) IsUnschedulable() bool { _ = "STUB: not implemented"; return false }

func (ri *ReservationInfo) IsTerminating() bool { _ = "STUB: not implemented"; return false }

func (ri *ReservationInfo) IsPreAllocation() bool { _ = "STUB: not implemented"; return false }

func (ri *ReservationInfo) GetTaints() []corev1.Taint { _ = "STUB: not implemented"; return nil }

// MatchReservationAffinity returns the statuses of whether the reservation affinity matches, whether the reservation
// taints are tolerated, and whether the reservation name matches.
func (ri *ReservationInfo) MatchReservationAffinity(reservationAffinity *reservationutil.RequiredReservationAffinity, node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// In this case, we suppose the necessary node labels have been patched to the reservation.

// NOTE: There are some special scenarios.
// For example, the AZ where the Pod wants to select the Reservation is cn-hangzhou, but the Reservation itself
// does not have this information, so it needs to perceive the label of the Node when Matching Affinity.

func (ri *ReservationInfo) MatchExactMatchSpec(podRequests corev1.ResourceList, spec *apiext.ExactMatchReservationSpec) bool {
	_ = "STUB: not implemented"
	return false
}

func (ri *ReservationInfo) FindMatchingUntoleratedTaint(reservationAffinity *reservationutil.RequiredReservationAffinity) (corev1.Taint, bool) {
	_ = "STUB: not implemented"
	return *new(corev1.Taint), false
}

func (ri *ReservationInfo) Clone() *ReservationInfo { _ = "STUB: not implemented"; return nil }

// use a shallow copy to reduce overhead

func (ri *ReservationInfo) UpdateReservation(r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

func (ri *ReservationInfo) UpdatePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (ri *ReservationInfo) AddAssignedPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (ri *ReservationInfo) RemoveAssignedPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (ri *ReservationInfo) RefreshPreCalculated() {
	_ = "STUB: not implemented"
	// Reservation available = Allocatable - Allocated - InnerReserved
	return
}

func (ri *ReservationInfo) GetAvailable() *framework.Resource {
	_ = "STUB: not implemented"
	return nil
}

func (ri *ReservationInfo) GetAllocatedResource() (*framework.Resource, int64, int64) {
	_ = "STUB: not implemented"
	return nil, 0, 0
}

// IsMatchable checks if the reservation is available to match any pods.
func (ri *ReservationInfo) IsMatchable() bool {
	_ = "STUB: not implemented"
	// if phase is available
	return false
}

// reservation/reserve pod

// operating pod

// In this case, the Controller has not yet updated the status of the Reservation to Succeeded,
// but in fact it can no longer be used for allocation. So it's better to skip first.

// IsMultiplePAPodsEnabled checks if multiple pre-allocated pods are enabled for the reservation.
func (ri *ReservationInfo) IsMultiplePAPodsEnabled() bool { _ = "STUB: not implemented"; return false }

// GetNonZeroRequestForResource returns the requested values,
// if the resource has undefined request for CPU or memory, it returns a default value.
func GetNonZeroRequestForResource(resourceName corev1.ResourceName, requests *corev1.ResourceList) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// Override if un-set, but not if explicitly set to zero

// Override if un-set, but not if explicitly set to zero
