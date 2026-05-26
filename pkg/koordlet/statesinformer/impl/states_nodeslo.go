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
	"sync"

	"k8s.io/client-go/tools/cache"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
)

const (
	nodeSLOInformerName PluginName = "nodeSLOInformer"
)

type nodeSLOInformer struct {
	nodeSLOInformer cache.SharedIndexInformer
	nodeSLORWMutex  sync.RWMutex
	nodeSLO         *slov1alpha1.NodeSLO

	callbackRunner *callbackRunner
}

func NewNodeSLOInformer() *nodeSLOInformer { _ = "STUB: not implemented"; return nil }

func (s *nodeSLOInformer) GetNodeSLO() *slov1alpha1.NodeSLO { _ = "STUB: not implemented"; return nil }

func (s *nodeSLOInformer) Setup(ctx *PluginOption, state *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (s *nodeSLOInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *nodeSLOInformer) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (s *nodeSLOInformer) updateNodeSLOSpec(nodeSLO *slov1alpha1.NodeSLO) {
	_ = "STUB: not implemented"
	return
}

func (s *nodeSLOInformer) setNodeSLOSpec(nodeSLO *slov1alpha1.NodeSLO) {
	_ = "STUB: not implemented"
	return
}

// merge nodeSLO spec with the default config

func (s *nodeSLOInformer) mergeNodeSLOSpec(nodeSLO *slov1alpha1.NodeSLO) {
	_ = "STUB: not implemented"
	return
}

// merge ResourceUsedThresholdWithBE individually for nil-ResourceUsedThresholdWithBE case

// merge ResourceQOSStrategy

// merge CPUBurstStrategy

// merge SystemStrategy

// merge Extensions

func newNodeSLOInformer(client koordclientset.Interface, nodeName string) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

// mergeSLOSpecResourceUsedThresholdWithBE merges the nodeSLO ResourceUsedThresholdWithBE with default configs
func mergeSLOSpecResourceUsedThresholdWithBE(defaultSpec, newSpec *slov1alpha1.ResourceThresholdStrategy) *slov1alpha1.ResourceThresholdStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ignore err for serializing/deserializing the same struct type

// NOTE: use deepcopy to avoid a overwrite to the global default

func mergeSLOSpecResourceQOSStrategy(defaultSpec,
	newSpec *slov1alpha1.ResourceQOSStrategy) *slov1alpha1.ResourceQOSStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ignore err for serializing/deserializing the same struct type

// NOTE: use deepcopy to avoid a overwrite to the global default

func mergeSLOSpecCPUBurstStrategy(defaultSpec,
	newSpec *slov1alpha1.CPUBurstStrategy) *slov1alpha1.CPUBurstStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ignore err for serializing/deserializing the same struct type

// NOTE: use deepcopy to avoid a overwrite to the global default

func mergeSLOSpecSystemStrategy(defaultSpec,
	newSpec *slov1alpha1.SystemStrategy) *slov1alpha1.SystemStrategy {
	_ = "STUB: not implemented"
	return nil
}

// ignore err for serializing/deserializing the same struct type

// NOTE: use deepcopy to avoid a overwrite to the global default

func mergeSLOSpecExtensions(defaultSpec,
	newSpec *slov1alpha1.ExtensionsMap) *slov1alpha1.ExtensionsMap {
	_ = "STUB: not implemented"
	return nil
}

// merge extension one by one so that all extension types can be protected during unmarshal

// ignore err for serializing/deserializing the same struct type

// NOTE: use deepcopy to avoid a overwrite to the global default

// mergeNoneResourceQOSIfDisabled complete ResourceQOSStrategy according to enable statuses of qos features
func mergeNoneResourceQOSIfDisabled(resourceQOS *slov1alpha1.ResourceQOSStrategy) {
	_ = "STUB: not implemented"
	return
}

// mergeNoneResctrlQOSIfDisabled completes node's resctrl qos config according to Enable options in ResctrlQOS
func mergeNoneResctrlQOSIfDisabled(resourceQOS *slov1alpha1.ResourceQOSStrategy) {
	_ = "STUB: not implemented"
	return
}

// mergeNoneMemoryQOSIfDisabled completes node's memory qos config according to Enable options in MemoryQOS
func mergeNoneMemoryQOSIfDisabled(resourceQOS *slov1alpha1.ResourceQOSStrategy) {
	_ = "STUB: not implemented"
	// if MemoryQOS.Enable=false, merge with NoneMemoryQOS
	return
}

func mergeNoneCPUQOSIfDisabled(resourceQOS *slov1alpha1.ResourceQOSStrategy) {
	_ = "STUB: not implemented"
	// if CPUQOS.Enabled=false, merge with NoneCPUQOS
	return
}
