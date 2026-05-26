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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/component-helpers/scheduling/corev1/nodeaffinity"
	"k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	// AnnotationReservePod indicates whether the pod is a reserved pod.
	AnnotationReservePod = extension.SchedulingDomainPrefix + "/reserve-pod"
	// AnnotationReservationName indicates the name of the reservation.
	AnnotationReservationName = extension.SchedulingDomainPrefix + "/reservation-name"
	// AnnotationReservationNode indicates the node name if the reservation specifies a node.
	AnnotationReservationNode = extension.SchedulingDomainPrefix + "/reservation-node"
	// AnnotationReservationResizeAllocatable indicates the desired allocatable are to be updated.
	AnnotationReservationResizeAllocatable = extension.SchedulingDomainPrefix + "/reservation-resize-allocatable"
	// AnnotationIsPreAllocation indicates whether the pod is a reserve pod and enables the pre-allocation.
	AnnotationIsPreAllocation = extension.SchedulingDomainPrefix + "/is-pre-allocation"
)

// ErrReasonPrefix is the prefix of the reservation-level scheduling errors.
const ErrReasonPrefix = "Reservation(s) "

// NewReservePod returns a fake pod set as the reservation's specifications.
// The reserve pod is only visible for the scheduler and does not make actual creation on nodes.
func NewReservePod(r *schedulingv1alpha1.Reservation) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// name, uid: reservation uid

// labels, annotations: `objectMeta` overwrites `template.objectMeta`

// annotate the reservePod

// for search inversely

// annotate the pre-allocation info

// annotate node name specified

// if the reservation specifies a nodeName, annotate it and cleanup spec.nodeName for other plugins not
// processing the nodeName before binding

// use reservation status.nodeName as the real scheduled result

// Forces priority to be set to maximum to prevent preemption.

//
// PodRequests is different from r.Status.Allocatable,
// which means that the scheduler allocates additional resources during scheduling
// or Reservation has changed the resource specifications through VPA.
//

func UpdateReservePodWithAllocatable(reservePod *corev1.Pod, podRequests, allocatable corev1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func cleanContainerResources(containers []corev1.Container) { _ = "STUB: not implemented"; return }

func ValidateReservation(r *schedulingv1alpha1.Reservation) error {
	_ = "STUB: not implemented"
	return nil
}

func PodPriority(r *schedulingv1alpha1.Reservation) int32 { _ = "STUB: not implemented"; return 0 }

func IsReservePod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func GetReservePodNamespacedName(r *schedulingv1alpha1.Reservation) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}

func GetReservationKey(r *schedulingv1alpha1.Reservation) string {
	_ = "STUB: not implemented"
	return ""
}

func GetReservePodNodeName(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func GetReservationNameFromReservePod(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func GetReservationSchedulerName(r *schedulingv1alpha1.Reservation) string {
	_ = "STUB: not implemented"
	return ""
}

// IsReservationActive checks if the reservation is scheduled and its status is Available/Waiting (active to use).
func IsReservationActive(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}

// IsReservationAvailable checks if the reservation is scheduled on a node and its status is Available.
func IsReservationAvailable(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}

func IsReservationSucceeded(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}

func IsReservationFailed(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}

func IsReservationExpired(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}

func GetReservationNodeName(r *schedulingv1alpha1.Reservation) string {
	_ = "STUB: not implemented"
	return ""
}

func SetReservationUnschedulable(r *schedulingv1alpha1.Reservation, msg string) {
	_ = "STUB: not implemented"
	// unschedule reservations can try scheduling in next cycles, so we does not update its phase
	// not duplicate condition info
	return
}

// if not set condition

// if is scheduled, keep the condition status

// if already unschedulable, update the message

func SetReservationExpired(r *schedulingv1alpha1.Reservation) { _ = "STUB: not implemented"; return }

// not duplicate expired info

// if not set condition

// if was ready

// if already not ready

func SetReservationSucceeded(r *schedulingv1alpha1.Reservation) { _ = "STUB: not implemented"; return }

// if not set condition

func SetReservationAvailable(r *schedulingv1alpha1.Reservation, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// initialize the conditions

func IsReservePodPreAllocation(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func ReservationRequests(r *schedulingv1alpha1.Reservation) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func ReservePorts(r *schedulingv1alpha1.Reservation) framework.HostPortInfo {
	_ = "STUB: not implemented"
	return *new(framework.HostPortInfo)
}

type ReservationOwnerMatcher struct {
	schedulingv1alpha1.ReservationOwner
	Selector labels.Selector
}

func ParseReservationOwnerMatchers(owners []schedulingv1alpha1.ReservationOwner) ([]ReservationOwnerMatcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ReservationOwnerMatcher) Match(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// MatchReservationOwners checks if the scheduling pod matches the reservation's owner spec.
// `reservation.spec.owners` defines the DNF (disjunctive normal form) of ObjectReference, ControllerReference
// (extended), LabelSelector, which means multiple selectors are firstly ANDed and secondly ORed.
func MatchReservationOwners(pod *corev1.Pod, matchers []ReservationOwnerMatcher) bool {
	_ = "STUB: not implemented"
	// assert pod != nil && r != nil
	// Owners == nil matches nothing, while Owners = [{}] matches everything
	return false
}

func MatchObjectRef(pod *corev1.Pod, objRef *corev1.ObjectReference) bool {
	_ = "STUB: not implemented"
	// `ResourceVersion`, `FieldPath` are ignored.
	// since only pod type are compared, `Kind` field is also ignored.
	return false
}

func MatchReservationControllerReference(pod *corev1.Pod, controllerRef *schedulingv1alpha1.ReservationControllerReference) bool {
	_ = "STUB: not implemented"
	// controllerRef matched if any of pod owner references matches the controllerRef;
	// typically a pod has only one controllerRef
	return false
}

// namespace field is extended

// currently `BlockOwnerDeletion` is ignored

func MatchLabels(podLabels map[string]string, selector labels.Selector) bool {
	_ = "STUB: not implemented"
	return false
}

type RequiredReservationAffinity struct {
	labelSelector         labels.Selector
	nodeSelector          *nodeaffinity.NodeSelector
	tolerations           []corev1.Toleration
	tolerateUnschedulable bool
	name                  string
}

// GetRequiredReservationAffinity returns the parsing result of pod's nodeSelector and nodeAffinity.
func GetRequiredReservationAffinity(pod *corev1.Pod) (*RequiredReservationAffinity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetName returns the reservation name if it is specified in the reservation affinity.
func (s *RequiredReservationAffinity) GetName() string { _ = "STUB: not implemented"; return "" }

// MatchName checks if the reservation affinity specifies a reservation name, and it matches the given reservation's.
func (s *RequiredReservationAffinity) MatchName(reservationName string) bool {
	_ = "STUB: not implemented"
	return false
}

// Match checks whether the pod is schedulable onto nodes according to
// the requirements in both nodeSelector and nodeAffinity.
// DEPRECATED: use MatchAffinity instead.
func (s *RequiredReservationAffinity) Match(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false

	// MatchAffinity checks whether the pod is schedulable onto nodes according to
	// the requirements in both nodeSelector and nodeAffinity.
}

func (s *RequiredReservationAffinity) MatchAffinity(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// FindMatchingUntoleratedTaint checks if the reservation tolerations tolerates all the filtered reservation taints.
// It returns the first taint without a toleration and the status if any taint is NOT tolerated.
func (s *RequiredReservationAffinity) FindMatchingUntoleratedTaint(taints []corev1.Taint, inclusionFilter func(*corev1.Taint) bool) (corev1.Taint, bool) {
	_ = "STUB: not implemented"
	return *new(corev1.Taint), false
}

// TolerateUnschedulable returns true if the pod tolerates the reservation unschedulable.
func (s *RequiredReservationAffinity) TolerateUnschedulable() bool {
	_ = "STUB: not implemented"
	return false
}

type ReservationResizeAllocatable struct {
	Resources corev1.ResourceList `json:"resources,omitempty"`
}

func GetReservationResizeAllocatable(annotations map[string]string) (*ReservationResizeAllocatable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetReservationResizeAllocatable(obj metav1.Object, resizeAllocatable *ReservationResizeAllocatable) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateReservationResizeAllocatable replaces or add the resources to the resizeAllocatable.
func UpdateReservationResizeAllocatable(obj metav1.Object, resources corev1.ResourceList) error {
	_ = "STUB: not implemented"
	return nil
}

func GetReservationRestrictedResources(allocatableResources []corev1.ResourceName, options *extension.ReservationRestrictedOptions) []corev1.ResourceName {
	_ = "STUB: not implemented"
	return nil
}

// NewReservationReason creates a reservation-level error reason with the given message.
func NewReservationReason(fmtMsg string, args ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// IsReservationReason checks if the error reason is at the reservation-level.
func IsReservationReason(reason string) bool { _ = "STUB: not implemented"; return false }

// DoNotScheduleTaintsFilter can filter out the node taints that reject scheduling Pod on a Node.
func DoNotScheduleTaintsFilter(t *corev1.Taint) bool {
	_ = "STUB: not implemented"
	// PodToleratesNodeTaints is only interested in NoSchedule and NoExecute taints.
	return false
}

// GetPreAllocationMode returns the pre-allocation mode of the reservation.
func GetPreAllocationMode(reservation *schedulingv1alpha1.Reservation) schedulingv1alpha1.PreAllocationMode {
	_ = "STUB: not implemented"
	return *new(schedulingv1alpha1.PreAllocationMode)
}

// IsMultiplePAPodsEnabled returns whether multiple pre-allocated pods are enabled for the reservation.
func IsMultiplePAPodsEnabled(reservation *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}
