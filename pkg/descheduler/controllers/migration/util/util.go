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

package util

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/reservation"
)

func GetCondition(status *sev1alpha1.PodMigrationJobStatus, conditionType sev1alpha1.PodMigrationJobConditionType) (int, *sev1alpha1.PodMigrationJobCondition) {
	_ = "STUB: not implemented"
	return 0, nil
}

func UpdateCondition(status *sev1alpha1.PodMigrationJobStatus, condition *sev1alpha1.PodMigrationJobCondition) bool {
	_ = "STUB: not implemented"
	return false
}

// Try to find this PodMigrationJob condition.

// We are adding new PodMigrationJob condition.

// We are updating an existing condition, so we need to check if it has changed.

// Return true if one of the fields have changed.

func IsMigratePendingPod(reservationObj reservation.Object) bool {
	_ = "STUB: not implemented"
	return false
}

func GetMaxUnavailable(replicas int, intOrPercent *intstr.IntOrString) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// handle the case of float value that less than 1.0

func GetMaxMigrating(replicas int, intOrPercent *intstr.IntOrString) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetLimiterBurst(burst int) int { _ = "STUB: not implemented"; return 0 }

// FilterPodWithMaxEvictionCost rejects if pod's eviction cost is math.MaxInt32
func FilterPodWithMaxEvictionCost(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
