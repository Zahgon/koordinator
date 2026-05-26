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

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	ClusterColocationProfile = "ClusterColocationProfile"
	ExtendedResourceSpec     = "ExtendedResourceSpec"
	MultiQuotaTree           = "MultiQuotaTree"
	DeviceResourceSpec       = "DeviceResourceSpec"
)

// PodMutatingHandler handles Pod
type PodMutatingHandler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder admission.Decoder
}

var _ admission.Handler = &PodMutatingHandler{}

func shouldIgnoreIfNotPod(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to sub resources or resources other than pods.
	return false
}

// Handle handles admission requests.
func (h *PodMutatingHandler) Handle(ctx context.Context, req admission.Request) (resp admission.Response) {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

// when pod.namespace is empty, using req.namespace

// Do not modify namespace in webhook

func (h *PodMutatingHandler) handleCreate(ctx context.Context, req admission.Request, obj *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *PodMutatingHandler) handleUpdate(ctx context.Context, req admission.Request, obj *corev1.Pod) error {
	_ = "STUB: not implemented"
	// TODO: add mutating logic for pod update here
	return nil
}

// var _ inject.Client = &PodMutatingHandler{}

// InjectClient injects the client into the PodMutatingHandler
func (h *PodMutatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &PodMutatingHandler{}
}

// InjectDecoder injects the decoder into the PodMutatingHandler
func (h *PodMutatingHandler) InjectDecoder(d admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}
