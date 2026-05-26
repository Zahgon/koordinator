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

	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	pvcInformerName PluginName = "pvcInformer"
)

type pvcInformer struct {
	pvcInformer cache.SharedIndexInformer
	pvcRWMutex  sync.RWMutex
	// The key of the map is the name of the PVC, in the format of PVC namespace/PVC name.
	// The value of the map is the name of the PV bound to the PVC.
	volumeNameMap map[string]string
}

func NewPVCInformer() *pvcInformer { _ = "STUB: not implemented"; return nil }

func (s *pvcInformer) GetVolumeName(pvcNamespace, pvcName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *pvcInformer) Setup(ctx *PluginOption, state *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (s *pvcInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *pvcInformer) HasSynced() bool {
	_ = "STUB: not implemented"
	// TODO add interface to check whether a plugin is enabled
	return false
}

func (s *pvcInformer) updateVolumeNameMap(pvc *corev1.PersistentVolumeClaim) {
	_ = "STUB: not implemented"
	return
}

func newPVCInformer(client clientset.Interface) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}
