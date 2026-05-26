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

package proxyserver

import (
	"context"

	runtimeapi "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
)

func (s *server) PreRunPodSandboxHook(ctx context.Context,
	req *runtimeapi.PodSandboxHookRequest) (*runtimeapi.PodSandboxHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PostStopPodSandboxHook(ctx context.Context,
	req *runtimeapi.PodSandboxHookRequest) (*runtimeapi.PodSandboxHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PreCreateContainerHook(ctx context.Context,
	req *runtimeapi.ContainerResourceHookRequest) (*runtimeapi.ContainerResourceHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PreStartContainerHook(ctx context.Context,
	req *runtimeapi.ContainerResourceHookRequest) (*runtimeapi.ContainerResourceHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PostStartContainerHook(ctx context.Context,
	req *runtimeapi.ContainerResourceHookRequest) (*runtimeapi.ContainerResourceHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PostStopContainerHook(ctx context.Context,
	req *runtimeapi.ContainerResourceHookRequest) (*runtimeapi.ContainerResourceHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *server) PreUpdateContainerResourcesHook(ctx context.Context,
	req *runtimeapi.ContainerResourceHookRequest) (*runtimeapi.ContainerResourceHookResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
