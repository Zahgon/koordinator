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

package networktopology

import (
	"sync"

	corev1 "k8s.io/api/core/v1"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

type Tree interface {
	AddNode(*corev1.Node)
	GetSnapshot() *TreeSnapshot
}

type tree struct {
	lock  sync.RWMutex
	root  *TreeNode
	index map[TreeNodeMeta]*TreeNode

	sortedTopologySpec []schedulingv1alpha1.NetworkTopologySpec
	layerToIndex       map[schedulingv1alpha1.TopologyLayer]int
}

func NewTree(clusterNetworkTopology *schedulingv1alpha1.ClusterNetworkTopology) (Tree, error) {
	_ = "STUB: not implemented"
	return *new(Tree), nil
}

func sortNetworkTopologySpecFromParentToChild(networkTopologySpec []schedulingv1alpha1.NetworkTopologySpec) ([]schedulingv1alpha1.NetworkTopologySpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TreeNode struct {
	TreeNodeMeta
	Parent   *TreeNode
	Children map[string]*TreeNode

	OfferSlot      int
	Score          int
	ExistingPodNum int
}

type TreeNodeMeta struct {
	Layer schedulingv1alpha1.TopologyLayer
	Name  string
}

func (t *tree) AddNode(node *corev1.Node) { _ = "STUB: not implemented"; return }

func (t *tree) addNode(treeNodesMeta []TreeNodeMeta, node *corev1.Node) {
	_ = "STUB: not implemented"
	return
}

// getTreeNodesMeta extracts all related TreeNodeMeta from a K8s Node.
// The returned order is the same as the order of []NetworkTopologySpec.
func (t *tree) getTreeNodesMeta(node *corev1.Node) ([]TreeNodeMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// currently we only allow node missing its direct parent layer (typically accelerators)

// make a virtual tree node so that we get a tree with each layer in the same tree level

type IsLayerAncestorFunc func(a, b schedulingv1alpha1.TopologyLayer) bool

type TreeSnapshot struct {
	TreeNode   *TreeNode
	IsAncestor IsLayerAncestorFunc
}

func (t *tree) GetSnapshot() *TreeSnapshot { _ = "STUB: not implemented"; return nil }

func DeepCopyTreeNode(origin, parent *TreeNode) *TreeNode { _ = "STUB: not implemented"; return nil }
