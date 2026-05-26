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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/koordinator-sh/koordinator/pkg/webhook/cm/plugins"
)

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch

type ConfigMapValidatingHandler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder admission.Decoder
}

func NewConfigMapValidatingHandler(c client.Client, d admission.Decoder) *ConfigMapValidatingHandler {
	_ = "STUB: not implemented"
	return nil
}

var _ admission.Handler = &ConfigMapValidatingHandler{}

func ShouldIgnoreIfNotConfigMap(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to sub resources or resources other than configmaps.
	return false
}

// Handle handles admission requests.
func (h *ConfigMapValidatingHandler) Handle(ctx context.Context, req admission.Request) (resp admission.Response) {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

func (h *ConfigMapValidatingHandler) getPlugins() []plugins.ConfigMapPlugin {
	_ = "STUB: not implemented"
	return nil
}

// var _ inject.Client = &ConfigMapValidatingHandler{}

// InjectClient injects the client into the ValidatingHandler
func (h *ConfigMapValidatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &ConfigMapValidatingHandler{}
}

// InjectDecoder injects the decoder into the ValidatingHandler
func (h *ConfigMapValidatingHandler) InjectDecoder(d admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newDecodeObj() (obj, oldObj *corev1.ConfigMap) { _ = "STUB: not implemented"; return nil, nil }

func getCMInfo(obj runtime.Object) string { _ = "STUB: not implemented"; return "" }
