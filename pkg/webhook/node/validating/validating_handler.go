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
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/koordinator-sh/koordinator/pkg/webhook/node/plugins"
)

// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch

type NodeValidatingHandler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder admission.Decoder
}

func NewNodeValidatingHandler(c client.Client, d admission.Decoder) *NodeValidatingHandler {
	_ = "STUB: not implemented"
	return nil
}

var _ admission.Handler = &NodeValidatingHandler{}

func ShouldIgnoreIfNotNode(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to sub resources or resources other than nodes.
	return false
}

// Handle handles admission requests.
func (h *NodeValidatingHandler) Handle(ctx context.Context, req admission.Request) (resp admission.Response) {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

func (h *NodeValidatingHandler) getPlugins() []plugins.NodePlugin {
	_ = "STUB: not implemented"
	return nil
}

// var _ inject.Client = &NodeValidatingHandler{}

// InjectClient injects the client into the ValidatingHandler
func (h *NodeValidatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &NodeValidatingHandler{}
}

// InjectDecoder injects the decoder into the ValidatingHandler
func (h *NodeValidatingHandler) InjectDecoder(d admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newDecodeObj() (obj, oldObj *corev1.Node) { _ = "STUB: not implemented"; return nil, nil }
