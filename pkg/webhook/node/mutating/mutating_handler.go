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

	"github.com/koordinator-sh/koordinator/pkg/webhook/node/plugins"
	"github.com/koordinator-sh/koordinator/pkg/webhook/node/plugins/resourceamplification"
)

var (
	nodeMutatingPlugins = []plugins.NodePlugin{
		resourceamplification.NewPlugin(),
	}
)

type IgnoreFilter func(req admission.Request) bool

// NodeMutatingHandler handles Node
type NodeMutatingHandler struct {
	Client client.Client

	// Decoder decodes objects
	Decoder admission.Decoder

	ignoreFilter IgnoreFilter
}

/* Uncomment the following lines if you want to enable mutating for node
// NewNodeMutatingHandler creates a new handler for node/status.
func NewNodeMutatingHandler() *NodeMutatingHandler {
	handler := &NodeMutatingHandler{
		ignoreFilter: shouldIgnoreIfNotNode,
	}
	return handler
}

func shouldIgnoreIfNotNode(req admission.Request) bool {
	// Ignore all calls to nodes status or resources other than node.
	if len(req.AdmissionRequest.SubResource) != 0 || req.AdmissionRequest.Resource.Resource != "nodes"{
		return true
	}
	return false
}
*/

// NewNodeStatusMutatingHandler creates a new handler for node/status.
func NewNodeStatusMutatingHandler(c client.Client, d admission.Decoder) *NodeMutatingHandler {
	_ = "STUB: not implemented"
	return nil
}

var _ admission.Handler = &NodeMutatingHandler{}

func shouldIgnoreIfNotNodeStatus(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to nodes or resources other than node status.
	return false
}

// Handle handles admission requests.
func (h *NodeMutatingHandler) Handle(ctx context.Context, req admission.Request) (resp admission.Response) {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

// var _ inject.Client = &NodeMutatingHandler{}

// InjectClient injects the client into the PodMutatingHandler
func (n *NodeMutatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &NodeMutatingHandler{}
}

// InjectDecoder injects the decoder into the PodMutatingHandler
func (n *NodeMutatingHandler) InjectDecoder(d admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}
