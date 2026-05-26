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

package mutating

import (
	"context"
	"math/rand"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configv1alpha1 "github.com/koordinator-sh/koordinator/apis/config/v1alpha1"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

var (
	randIntnFn = rand.Intn
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=config.koordinator.sh,resources=clustercolocationprofiles,verbs=get;list;watch

func (h *ReservationMutatingHandler) clusterColocationProfileMutatingReservation(ctx context.Context, req admission.Request, reservation *schedulingv1alpha1.Reservation) error {
	_ = "STUB: not implemented"
	return nil
}

// Reservation does not support the namespaceSelector.

// sort the profile in lexicographic order

// TODO: support mutate the resource spec for reservation

func (h *ReservationMutatingHandler) matchObjectSelector(reservation, oldReservation *schedulingv1alpha1.Reservation, objectSelector *metav1.LabelSelector) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func shouldSkipProfile(profile *configv1alpha1.ClusterColocationProfile) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (h *ReservationMutatingHandler) doMutateByColocationProfile(ctx context.Context, reservation *schedulingv1alpha1.Reservation, profile *configv1alpha1.ClusterColocationProfile) error {
	_ = "STUB: not implemented"
	return nil
}
