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

package resourceamplification

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/koordinator-sh/koordinator/pkg/webhook/node/plugins"
)

var (
	// Only supports amplification of cpu and memory resources
	supportedResources = []corev1.ResourceName{
		corev1.ResourceCPU,
		corev1.ResourceMemory,
	}
)

func NewPlugin() *NodeResourceAmplificationPlugin { _ = "STUB: not implemented"; return nil }

var _ plugins.NodePlugin = &NodeResourceAmplificationPlugin{}

type NodeResourceAmplificationPlugin struct {
}

func (n *NodeResourceAmplificationPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (n *NodeResourceAmplificationPlugin) Validate(ctx context.Context, req admission.Request, node, oldNode *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Admit makes an admission decision based on the request attributes
func (n *NodeResourceAmplificationPlugin) Admit(ctx context.Context, req admission.Request, node, oldNode *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// isSupportedResourceChanged checks if the supported resources are changed
func (n *NodeResourceAmplificationPlugin) isSupportedResourceChanged(oldNode, node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *NodeResourceAmplificationPlugin) handleUpdate(oldNode, node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// When the feature is turned off, the saved raw allocatable needs to be cleaned up

// do nothing

// Set original resource under two conditions
// 1. if the original resource is not set
// 2. Node allocatable are updated, such as the administrator adjusting the reserved resource
//    https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources
//    Here, we rely on an assumption that only kubelet will update allocatable, and other roles will only expand new fields and not overwrite native fields.
// FIXME
// 1. If the updated resource of kubelet happens to be the same as the amplificated allocatable, the update will not take effect.
// 2. If the webhook policy is configured to ignore, allocatable changes in the kubelet while the webhook is unavailable will not be observed.

// resources in other dimensions will be dynamically updated, so they are not saved here. The caller needs to be aware of this feature.

// parse resource amplification ratios from annotations

// update allocatable by originalResources * amplificationRatios:
//     originalResources: {cpu: 1000, memory: 2000} * amplificationRatios: {cpu: 2, memory: 3} = {cpu: 2000, memory: 6000}
// If the resources saved in original are missing, the original will still prevail, and the missing dimensions will not be scaled:
//     originalResources: {cpu: 1000} * amplificationRatios: {cpu: 2, memory: 3} = {cpu: 2000}
// If the resource dimensions supported by the amplification change, the latest one shall prevail, and the missing dimensions will not be amplified:
//     originalResources: {cpu: 1000, memory: 2000} * amplificationRatios: {cpu: 2} = {cpu: 2000}
