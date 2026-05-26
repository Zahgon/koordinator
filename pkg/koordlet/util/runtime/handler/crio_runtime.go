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

package handler

import (
	"context"
	"time"

	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func GetCrioEndpoint() string { _ = "STUB: not implemented"; return "" }

func GetCrioEndpoint2() string { _ = "STUB: not implemented"; return "" }

type CrioRuntimeHandler struct {
	runtimeServiceClient runtimeapi.RuntimeServiceClient
	timeout              time.Duration
	endpoint             string
}

func NewCrioRuntimeHandler(endpoint string) (ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(ContainerRuntimeHandler), nil
}

func (c *CrioRuntimeHandler) StopContainer(ctx context.Context, containerID string, timeout int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CrioRuntimeHandler) UpdateContainerResources(containerID string, opts UpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}
