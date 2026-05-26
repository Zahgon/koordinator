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
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"

	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/config"
	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/dispatcher"
	resource_executor "github.com/koordinator-sh/koordinator/pkg/runtimeproxy/resexecutor"
)

const (
	defaultTimeout = 5 * time.Second
)

type RuntimeRequestInterceptor interface {
	InterceptRuntimeRequest(serviceType RuntimeServiceType, ctx context.Context, request interface{}, handler grpc.UnaryHandler, alphaRuntime bool) (interface{}, error)
}

var _ runtimeapi.RuntimeServiceServer = &criServer{}

type criServer struct {
	runtimeapi.UnimplementedRuntimeServiceServer
	RuntimeRequestInterceptor
	backendRuntimeServiceClient runtimeapi.RuntimeServiceClient
}

type RuntimeManagerCriServer struct {
	hookDispatcher *dispatcher.RuntimeHookDispatcher
	criServer      *criServer
}

func NewRuntimeManagerCriServer() *RuntimeManagerCriServer { _ = "STUB: not implemented"; return nil }

func (c *RuntimeManagerCriServer) Name() string { _ = "STUB: not implemented"; return "" }

func (c *RuntimeManagerCriServer) Run() error { _ = "STUB: not implemented"; return nil }

// For unsupported requests, pass through directly to the backend

func (c *RuntimeManagerCriServer) getRuntimeHookInfo(serviceType RuntimeServiceType) (config.RuntimeRequestPath,
	resource_executor.RuntimeResourceType) {
	_ = "STUB: not implemented"
	return *new(config.RuntimeRequestPath), *new(resource_executor.RuntimeResourceType)
}

func (c *RuntimeManagerCriServer) InterceptRuntimeRequest(serviceType RuntimeServiceType,
	ctx context.Context, request interface{}, handler grpc.UnaryHandler, alphaRuntime bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//if alphaRuntime {
//	request, err = alphaObjectToV1Object(request)
//	if err != nil {
//		return nil, err
//	}
//}

// TODO deal with the Dispatch response

// call the backend runtime engine
//if alphaRuntime {
//	request, err = v1ObjectToAlphaObject(request)
//	if err != nil {
//		return nil, err
//	}
//}

// responseConverted := false

//if alphaRuntime {
//	responseConverted = true
//	res, err = alphaObjectToV1Object(res)
//	if err != nil {
//		return nil, err
//	}
//}

// store checkpoint info basing request only when response success

// post call hook server
// TODO the response

// if responseConverted {
//res, err = v1ObjectToAlphaObject(res)
//if err != nil {
//	return nil, err
//}
// }

func dialer(ctx context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// UpdatePodSandboxResources implements runtimeapi.RuntimeServiceServer.
func (c *criServer) UpdatePodSandboxResources(ctx context.Context, req *runtimeapi.UpdatePodSandboxResourcesRequest) (*runtimeapi.UpdatePodSandboxResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *RuntimeManagerCriServer) initCriServer(runtimeSockPath string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// According to the version of cri api supported by backend runtime, create the corresponding cri server.

func (c *RuntimeManagerCriServer) failOver() error {
	_ = "STUB: not implemented"
	// Only CRI v1 API is supported. CRI v1alpha2 has been removed since K8s 1.26.
	return nil
}
