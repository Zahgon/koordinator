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

package store

import (
	"sync"

	"github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
)

const (
	defaultPoolSize = 10
)

// PodSandboxInfo is almost the same with v1alpha.RunPodSandboxHookRequest
type PodSandboxInfo struct {
	*v1alpha1.PodSandboxHookRequest
}

// ContainerInfo is almost the same with v1alpha.ContainerResourceHookRequest
type ContainerInfo struct {
	*v1alpha1.ContainerResourceHookRequest
}

func (p *PodSandboxInfo) GetPodSandboxHookRequest() *v1alpha1.PodSandboxHookRequest {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContainerInfo) GetContainerResourceHookRequest() *v1alpha1.ContainerResourceHookRequest {
	_ = "STUB: not implemented"
	return nil
}

type metaManager struct {
	sync.RWMutex
	podInfos       map[string]*PodSandboxInfo
	containerInfos map[string]*ContainerInfo
}

// reset. currently only used by test case
func (mm *metaManager) reset() { _ = "STUB: not implemented"; return }

var m = &metaManager{
	podInfos:       make(map[string]*PodSandboxInfo, defaultPoolSize),
	containerInfos: make(map[string]*ContainerInfo, defaultPoolSize),
}

// WritePodSandboxInfo checkpoints the pod level info
func WritePodSandboxInfo(podUID string, pod *PodSandboxInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteContainerInfo returns
func WriteContainerInfo(containerUID string, container *ContainerInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPodSandboxInfo returns sandbox info
func GetPodSandboxInfo(podUID string) *PodSandboxInfo { _ = "STUB: not implemented"; return nil }

func GetContainerInfo(containerUID string) *ContainerInfo { _ = "STUB: not implemented"; return nil }

// DeletePodSandboxInfo delete pod checkpoint indexed by podUID
func DeletePodSandboxInfo(podUID string) { _ = "STUB: not implemented"; return }

// DeleteContainerInfo delete container checkpoint indexed by containerUID
func DeleteContainerInfo(containerUID string) { _ = "STUB: not implemented"; return }
