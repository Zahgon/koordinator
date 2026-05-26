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

package impl

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	nodeInformerName PluginName = "nodeInformer"
)

type nodeInformer struct {
	nodeInformer   cache.SharedIndexInformer
	nodeRWMutex    sync.RWMutex
	node           *corev1.Node
	callbackRunner *callbackRunner
}

func NewNodeInformer() *nodeInformer { _ = "STUB: not implemented"; return nil }

func (s *nodeInformer) GetNode() *corev1.Node { _ = "STUB: not implemented"; return nil }

func (s *nodeInformer) Setup(ctx *PluginOption, state *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (s *nodeInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *nodeInformer) HasSynced() bool { _ = "STUB: not implemented"; return false }

// maybe the cache has synced ,but the event handlers haven't beed called yet

var newNodeInformer = func(client clientset.Interface, nodeName string) cache.SharedIndexInformer {
	tweakListOptionsFunc := func(opt *metav1.ListOptions) {
		opt.FieldSelector = "metadata.name=" + nodeName
	}

	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (apiruntime.Object, error) {
				tweakListOptionsFunc(&options)
				return client.CoreV1().Nodes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				tweakListOptionsFunc(&options)
				return client.CoreV1().Nodes().Watch(context.TODO(), options)
			},
		},
		&corev1.Node{},
		time.Hour*12,
		cache.Indexers{},
	)
}

func (s *nodeInformer) syncNode(newNode *corev1.Node) { _ = "STUB: not implemented"; return }

// also register node for metrics

func isNodeMetadataUpdated(oldNode, newNode *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func recordNodeResourceMetrics(node *corev1.Node) {
	_ = "STUB: not implemented"
	// register node labels
	return
}

// record node resource metrics

func recordNodeResources(node *corev1.Node) { _ = "STUB: not implemented"; return }

// record node allocatable of BatchCPU & BatchMemory

// record node allocatable of MidCPU & MidMemory
