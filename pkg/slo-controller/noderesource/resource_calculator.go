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

package noderesource

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

func (r *NodeResourceReconciler) isColocationCfgDisabled(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *NodeResourceReconciler) resetNodeResource(node *corev1.Node, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeResourceReconciler) calculateNodeResource(node *corev1.Node,
	nodeMetric *slov1alpha1.NodeMetric, podList *corev1.PodList) *framework.NodeResource {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeResourceReconciler) updateNodeResource(node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
	// avoid overwriting the cache
}

// pre-update once

func (r *NodeResourceReconciler) updateNodeStatus(node *corev1.Node, strategy *configuration.ColocationStrategy, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
}

// avoid overwriting the cache

func (r *NodeResourceReconciler) updateNodeMeta(node *corev1.Node, strategy *configuration.ColocationStrategy, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
}

// avoid overwriting the cache

// updateNodeExtensions is an extension point for updating node other than node metric resources.
func (r *NodeResourceReconciler) updateNodeExtensions(node *corev1.Node, nodeMetric *slov1alpha1.NodeMetric, podList *corev1.PodList) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeResourceReconciler) isNodeResourceSyncNeeded(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (r *NodeResourceReconciler) isCommonNodeNeedSync(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) bool {
	_ = "STUB: not implemented"
	// update time gap is bigger than UpdateTimeThresholdSeconds
	return false
}

func (r *NodeResourceReconciler) prepareNodeResource(strategy *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) {
	_ = "STUB: not implemented"
	return
}
