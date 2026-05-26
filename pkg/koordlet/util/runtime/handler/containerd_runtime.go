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
	"net"
	"time"

	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var (
	GrpcDial = grpc.DialContext // for test
)

func GetContainerdEndpoint() string { _ = "STUB: not implemented"; return "" }

func GetContainerdEndpoint2() string { _ = "STUB: not implemented"; return "" }

type ContainerdRuntimeHandler struct {
	runtimeServiceClient runtimeapi.RuntimeServiceClient
	timeout              time.Duration
	endpoint             string
}

func NewContainerdRuntimeHandler(endpoint string) (ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(ContainerRuntimeHandler), nil
}

func (c *ContainerdRuntimeHandler) StopContainer(con context.Context, containerID string, timeout int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContainerdRuntimeHandler) UpdateContainerResources(containerID string, opts UpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func getRuntimeClient(endpoint string) (runtimeapi.RuntimeServiceClient, error) {
	_ = "STUB: not implemented"
	return *new(runtimeapi.RuntimeServiceClient), nil
}

func getClientConnection(endpoint string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAddressAndDialer(endpoint string) (string, func(context context.Context, addr string) (net.Conn, error), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func parseEndpoint(endpoint string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func dial(context context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
