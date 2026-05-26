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

package validating

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/koordinator-sh/koordinator/pkg/webhook/quotaevaluate"
)

const (
	ClusterReservation       = "ClusterReservation"
	ClusterColocationProfile = "ClusterColocationProfile"
	EvaluateQuota            = "EvaluateQuota"
	DeviceResource           = "DeviceResource"
	EnhancedValidation       = "EnhancedValidation"
)

// PodValidatingHandler handles Pod
type PodValidatingHandler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder admission.Decoder

	// QuotaEvaluator evaluate pod quota usage
	QuotaEvaluator quotaevaluate.Evaluator

	// PodEnhancedValidator manages pod enhanced validation configuration
	PodEnhancedValidator *PodEnhancedValidator
}

var _ admission.Handler = &PodValidatingHandler{}

func shouldIgnoreIfNotPod(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to sub resources or resources other than pods.
	return false
}

func (h *PodValidatingHandler) validatingPodFn(ctx context.Context, req admission.Request) (allowed bool, reason string, err error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

var _ admission.Handler = &PodValidatingHandler{}

// Handle handles admission requests.
func (h *PodValidatingHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

// var _ inject.Client = &PodValidatingHandler{}

// InjectClient injects the client into the PodValidatingHandler
func (h *PodValidatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &PodValidatingHandler{}
}

// InjectDecoder injects the decoder into the PodValidatingHandler
func (h *PodValidatingHandler) InjectDecoder(d admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}
