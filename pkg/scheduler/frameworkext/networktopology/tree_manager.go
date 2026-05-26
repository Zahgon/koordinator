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
	"context"
	"sync"
	"time"

	"k8s.io/client-go/informers"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	clientv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
	"github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
)

const (
	Name = "NetworkTopology"

	defaultClusterNetworkTopologyName = "default"
	defaultRebuildTreeDuration        = time.Second * 10
)

var _ TreeManager = &treeManager{}

type TreeManager interface {
	Name() string
	Run(ctx context.Context)
	GetSnapshot() *TreeSnapshot
}

type treeManager struct {
	informerFactory      informers.SharedInformerFactory
	koordInformerFactory koordinatorinformers.SharedInformerFactory
	topologyLister       v1alpha1.ClusterNetworkTopologyLister
	topologyClient       clientv1alpha1.ClusterNetworkTopologyInterface

	sync.RWMutex
	tree Tree
}

func (tm *treeManager) Name() string { _ = "STUB: not implemented"; return "" }

func NewTreeManager(
	koordInformerFactory koordinatorinformers.SharedInformerFactory,
	informerFactory informers.SharedInformerFactory,
	client koordclientset.Interface,
) TreeManager {
	_ = "STUB: not implemented"
	return *new(TreeManager)
}

func (tm *treeManager) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (tm *treeManager) run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (tm *treeManager) buildTree() Tree { _ = "STUB: not implemented"; return *new(Tree) }

func (tm *treeManager) updateStatus(ctx context.Context, t Tree) { _ = "STUB: not implemented"; return }

func makeStatus(treeNode *TreeNode, clusterNetworkTopology *schedulingv1alpha1.ClusterNetworkTopology) (nodeNum int32) {
	_ = "STUB: not implemented"
	return 0
}

func (tm *treeManager) cacheTree(t Tree) { _ = "STUB: not implemented"; return }

func (tm *treeManager) GetSnapshot() *TreeSnapshot { _ = "STUB: not implemented"; return nil }
