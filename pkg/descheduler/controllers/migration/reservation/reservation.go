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
	"sigs.k8s.io/controller-runtime/pkg/client"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

var _ Object = &Reservation{}

type Reservation struct {
	*sev1alpha1.Reservation
}

func NewReservation(reservation *sev1alpha1.Reservation) Object {
	_ = "STUB: not implemented"
	return *new(Object)
}

func (r *Reservation) String() string { _ = "STUB: not implemented"; return "" }

func (r *Reservation) OriginObject() client.Object {
	_ = "STUB: not implemented"
	return *new(client.Object)
}

func (r *Reservation) GetReservationConditions() []sev1alpha1.ReservationCondition {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reservation) QueryPreemptedPodsRefs() []corev1.ObjectReference {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reservation) GetBoundPod() *corev1.ObjectReference { _ = "STUB: not implemented"; return nil }

func (r *Reservation) GetReservationOwners() []sev1alpha1.ReservationOwner {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reservation) GetScheduledNodeName() string { _ = "STUB: not implemented"; return "" }

func (r *Reservation) GetPhase() sev1alpha1.ReservationPhase {
	_ = "STUB: not implemented"
	return *new(sev1alpha1.ReservationPhase)
}

func (r *Reservation) NeedPreemption() bool { _ = "STUB: not implemented"; return false }

func GetReservationCondition(r Object, conditionType sev1alpha1.ReservationConditionType, reason string) *sev1alpha1.ReservationCondition {
	_ = "STUB: not implemented"
	return nil
}

func GetUnschedulableCondition(r Object) *sev1alpha1.ReservationCondition {
	_ = "STUB: not implemented"
	return nil
}

func IsReservationScheduled(r Object) bool { _ = "STUB: not implemented"; return false }

func IsReservationPending(r Object) bool { _ = "STUB: not implemented"; return false }

// IsReservationAvailable checks if the reservation is scheduled on a node and its status is Available.
func IsReservationAvailable(r Object) bool { _ = "STUB: not implemented"; return false }

func IsReservationSucceeded(r Object) bool { _ = "STUB: not implemented"; return false }

func IsReservationFailed(r Object) bool { _ = "STUB: not implemented"; return false }

func IsReservationExpired(r Object) bool { _ = "STUB: not implemented"; return false }
