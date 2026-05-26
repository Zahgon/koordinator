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

package cri

import (
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"

	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/store"
	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/utils"
)

type PodResourceExecutor struct {
	store.PodSandboxInfo
}

func NewPodResourceExecutor() *PodResourceExecutor { _ = "STUB: not implemented"; return nil }

func (p *PodResourceExecutor) String() string { _ = "STUB: not implemented"; return "" }

func (p *PodResourceExecutor) GetMetaInfo() string { _ = "STUB: not implemented"; return "" }

func (p *PodResourceExecutor) GenerateHookRequest() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodResourceExecutor) loadPodSandboxFromStore(podID string) error {
	_ = "STUB: not implemented"
	return nil
}

// ParseRequest would
func (p *PodResourceExecutor) ParseRequest(req interface{}) (utils.CallHookPluginOperation, error) {
	_ = "STUB: not implemented"
	return *new(utils.CallHookPluginOperation), nil
}

func (p *PodResourceExecutor) ParsePod(podsandbox *runtimeapi.PodSandbox) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: how to get cgroup parent when failOver

func (p *PodResourceExecutor) ResourceCheckPoint(response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodResourceExecutor) DeleteCheckpointIfNeed(req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateRequest will update PodResourceExecutor from hook response and then update CRI request.
func (p *PodResourceExecutor) UpdateRequest(rsp interface{}, req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// update PodResourceExecutor

// update CRI request
