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
	"k8s.io/apimachinery/pkg/types"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

func GetReservationNamespacedName(ref *corev1.ObjectReference) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}

func CreateOrUpdateReservationOptions(job *sev1alpha1.PodMigrationJob, pod *corev1.Pod) *sev1alpha1.PodMigrateReservationOptions {
	_ = "STUB: not implemented"
	return nil
}

// Reservation used for migration is no longer reused after consumed

// force removed the assigned nodeName of target Pod to request new Node

func appendSkipNodeAffinity(pod *corev1.Pod, reservationOptions *sev1alpha1.PodMigrateReservationOptions) {
	_ = "STUB: not implemented"
	return
}

func GenerateReserveResourceOwners(pod *corev1.Pod) []sev1alpha1.ReservationOwner {
	_ = "STUB: not implemented"
	return nil
}
