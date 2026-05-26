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

type ContainerResourceExecutor struct {
	store.ContainerInfo
}

func NewContainerResourceExecutor() *ContainerResourceExecutor {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContainerResourceExecutor) String() string { _ = "STUB: not implemented"; return "" }

func (c *ContainerResourceExecutor) GetMetaInfo() string { _ = "STUB: not implemented"; return "" }

func (c *ContainerResourceExecutor) GenerateHookRequest() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContainerResourceExecutor) loadContainerInfoFromStore(containerID, stage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContainerResourceExecutor) ParseRequest(req interface{}) (utils.CallHookPluginOperation, error) {
	_ = "STUB: not implemented"
	return *new(utils.CallHookPluginOperation), nil
}

// get the pod info from local store

func (c *ContainerResourceExecutor) ParseContainer(container *runtimeapi.Container) error {
	_ = "STUB: not implemented"
	return nil
}

// envs info would be loss during failover.
// TODO: How to get resource and envs when failOver

func (c *ContainerResourceExecutor) ResourceCheckPoint(rsp interface{}) error {
	_ = "STUB: not implemented"
	// container level resource checkpoint would be triggered during post container create only
	return nil
}

func (c *ContainerResourceExecutor) DeleteCheckpointIfNeed(req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateRequest will update ContainerResourceExecutor from hook response and then update CRI request.
func (c *ContainerResourceExecutor) UpdateRequest(rsp interface{}, req interface{}) error {
	_ = "STUB: not implemented"
	// update ContainerResourceExecutor
	return nil
}

// update PodResourceExecutor

// update CRI request
