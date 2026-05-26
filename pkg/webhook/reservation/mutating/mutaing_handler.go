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

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	ClusterColocationProfile = "ClusterColocationProfile"
)

// ReservationMutatingHandler handles Reservation.
type ReservationMutatingHandler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder admission.Decoder
}

var _ admission.Handler = &ReservationMutatingHandler{}

func shouldIgnoreIfNotReservation(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to sub resources or resources other than reservations.
	return false
}

// Handle handles admission requests.
func (h *ReservationMutatingHandler) Handle(ctx context.Context, req admission.Request) (resp admission.Response) {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

func (h *ReservationMutatingHandler) handleCreate(ctx context.Context, req admission.Request, obj *schedulingv1alpha1.Reservation) error {
	_ = "STUB: not implemented"
	return nil
}

// var _ inject.Client = &ReservationMutatingHandler{}

// InjectClient injects the client into the ReservationMutatingHandler
func (h *ReservationMutatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &ReservationMutatingHandler{}
}

// InjectDecoder injects the decoder into the ReservationMutatingHandler
func (h *ReservationMutatingHandler) InjectDecoder(d admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}
