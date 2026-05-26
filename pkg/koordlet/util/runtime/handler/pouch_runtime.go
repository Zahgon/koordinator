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

	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

const (
	PouchEndpointSubFilepath = "pouchcri.sock"
)

func GetPouchEndpoint() string { _ = "STUB: not implemented"; return "" }

type PouchRuntimeHandler struct {
	runtimeServiceClient v1.RuntimeServiceClient
	timeout              time.Duration
	endpoint             string
}

func NewPouchRuntimeHandler(endpoint string) (ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(ContainerRuntimeHandler), nil
}

// use v1alpha2 protocol

func (c *PouchRuntimeHandler) StopContainer(ctx context.Context, containerID string, timeout int64) error {
	_ = "STUB: not implemented"
	return nil
}

// pouch cannot handle context with timeout

func (c *PouchRuntimeHandler) UpdateContainerResources(containerID string, opts UpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func getRuntimeV1alpha2Client(endpoint string) (v1.RuntimeServiceClient, error) {
	_ = "STUB: not implemented"
	return *new(v1.RuntimeServiceClient), nil
}
