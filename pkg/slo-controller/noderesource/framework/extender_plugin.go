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

package framework

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/configuration"
)

// The plugins called in the node resource initialization:
// - Setup
//
// The plugins called order in the node resource reconcile loop:
// - Calculate/Reset -> NodePreUpdate -> NodeUpdate(NodePrepare -> NodeStatusCheck,NodeMetaCheck)
// The NodeUpdate stage can be called for multiple times with retries.
//
// For more info, please see the README.md.
var (
	globalSetupExtender             = NewRegistry("Setup")
	globalNodePreUpdateExtender     = NewRegistry("NodePreUpdate")
	globalNodePrepareExtender       = NewRegistry("NodePrepare")
	globalNodeStatusCheckExtender   = NewRegistry("NodeStatusCheck")
	globalNodeMetaCheckExtender     = NewRegistry("NodeMetaCheck")
	globalResourceCalculateExtender = NewRegistry("ResourceCalculate")
)

// Plugin has its name. Plugins in a registry are executed in order of the registration.
type Plugin interface {
	Name() string
}

// SetupPlugin implements setup for the plugin.
// The framework exposes the kube ClientSet and controller builder to the plugins thus the plugins can set up their
// necessary clients, add new watches and initialize their internal states.
// The Setup of each plugin will be called before other extension stages and invoked only once.
type SetupPlugin interface {
	Plugin
	Setup(opt *Option) error
}

func RegisterSetupExtender(filter FilterFn, plugins ...SetupPlugin) {
	_ = "STUB: not implemented"
	return
}

func RunSetupExtenders(opt *Option) { _ = "STUB: not implemented"; return }

func UnregisterSetupExtender(name string) { _ = "STUB: not implemented"; return }

// NodePreUpdatePlugin implements preprocessing for the calculated results called before updating the Node.
// There are mainly two use cases for this stage:
// 1. A plugin may prepare and update some Objects like CRDs before updating the Node obj (NodePrepare and NodeXXXCheck).
// 2. A plugin may need to mutate the internal NodeResource object before updating the Node object.
// It differs from the NodePreparePlugin in that a NodePreUpdatePlugin will be invoked only once in one loop (so the
// plugin should consider implement a retry login itself if needed), while the NodePreparePlugin is not expected to
// update other objects or mutate the NodeResource.
type NodePreUpdatePlugin interface {
	Plugin
	PreUpdate(strategy *configuration.ColocationStrategy, node *corev1.Node, nr *NodeResource) error
}

func RegisterNodePreUpdateExtender(filter FilterFn, plugins ...NodePreUpdatePlugin) {
	_ = "STUB: not implemented"
	return
}

func RunNodePreUpdateExtenders(strategy *configuration.ColocationStrategy, node *corev1.Node, nr *NodeResource) {
	_ = "STUB: not implemented"
	return
}

func UnregisterNodePreUpdateExtender(name string) { _ = "STUB: not implemented"; return }

// NodePreparePlugin implements node resource preparing for the calculated results.
// For example, assign extended resources in the node allocatable.
// It is invoked each time the controller tries updating the latest NodeResource object with calculated results.
// NOTE: The Prepare should be idempotent since it can be called multiple times in one reconciliation.
type NodePreparePlugin interface {
	Plugin
	Prepare(strategy *configuration.ColocationStrategy, node *corev1.Node, nr *NodeResource) error
}

func RegisterNodePrepareExtender(filter FilterFn, plugins ...NodePreparePlugin) {
	_ = "STUB: not implemented"
	return
}

func RunNodePrepareExtenders(strategy *configuration.ColocationStrategy, node *corev1.Node, nr *NodeResource) {
	_ = "STUB: not implemented"
	return
}

func UnregisterNodePrepareExtender(name string) { _ = "STUB: not implemented"; return }

// NodeStatusCheckPlugin implements the check of resource updating.
// For example, trigger an update if the values of the current is more than 10% different with the former.
type NodeStatusCheckPlugin interface {
	Plugin
	NeedSync(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string)
}

func RegisterNodeStatusCheckExtender(filter FilterFn, plugins ...NodeStatusCheckPlugin) {
	_ = "STUB: not implemented"
	return
}

func UnregisterNodeStatusCheckExtender(name string) { _ = "STUB: not implemented"; return }

func RunNodeStatusCheckExtenders(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

type NodeMetaCheckPlugin interface {
	Plugin
	NeedSyncMeta(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string)
}

func RegisterNodeMetaCheckExtender(filter FilterFn, plugins ...NodeMetaCheckPlugin) {
	_ = "STUB: not implemented"
	return
}

func UnregisterNodeMetaCheckExtender(name string) { _ = "STUB: not implemented"; return }

func RunNodeMetaCheckExtenders(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

type ResourceResetPlugin interface {
	Plugin
	Reset(node *corev1.Node, message string) []ResourceItem
}

func RunResourceResetExtenders(nr *NodeResource, node *corev1.Node, message string) {
	_ = "STUB: not implemented"
	return
}

// ResourceCalculatePlugin implements resource counting and overcommitment algorithms.
// It implements Reset which can be invoked when the calculated resources need a reset.
// A ResourceCalculatePlugin can handle the case when the metrics are abnormal by implementing degraded calculation.
type ResourceCalculatePlugin interface {
	ResourceResetPlugin
	Calculate(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList, metrics *ResourceMetrics) ([]ResourceItem, error)
}

func RegisterResourceCalculateExtender(filter FilterFn, plugins ...ResourceCalculatePlugin) {
	_ = "STUB: not implemented"
	return
}

func UnregisterResourceCalculateExtender(name string) { _ = "STUB: not implemented"; return }

func RunResourceCalculateExtenders(nr *NodeResource, strategy *configuration.ColocationStrategy, node *corev1.Node,
	podList *corev1.PodList, resourceMetrics *ResourceMetrics) {
	_ = "STUB: not implemented"
	return
}
