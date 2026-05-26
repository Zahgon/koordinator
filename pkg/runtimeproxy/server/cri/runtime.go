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

	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func (c *criServer) Version(ctx context.Context, req *runtimeapi.VersionRequest) (*runtimeapi.VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) RunPodSandbox(ctx context.Context, req *runtimeapi.RunPodSandboxRequest) (*runtimeapi.RunPodSandboxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) StopPodSandbox(ctx context.Context, req *runtimeapi.StopPodSandboxRequest) (*runtimeapi.StopPodSandboxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) RemovePodSandbox(ctx context.Context, req *runtimeapi.RemovePodSandboxRequest) (*runtimeapi.RemovePodSandboxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) PodSandboxStatus(ctx context.Context, req *runtimeapi.PodSandboxStatusRequest) (*runtimeapi.PodSandboxStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ListPodSandbox(ctx context.Context, req *runtimeapi.ListPodSandboxRequest) (*runtimeapi.ListPodSandboxResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ListPodSandboxStats(ctx context.Context, req *runtimeapi.ListPodSandboxStatsRequest) (*runtimeapi.ListPodSandboxStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) CreateContainer(ctx context.Context, req *runtimeapi.CreateContainerRequest) (*runtimeapi.CreateContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) StartContainer(ctx context.Context, req *runtimeapi.StartContainerRequest) (*runtimeapi.StartContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) StopContainer(ctx context.Context, req *runtimeapi.StopContainerRequest) (*runtimeapi.StopContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) RemoveContainer(ctx context.Context, req *runtimeapi.RemoveContainerRequest) (*runtimeapi.RemoveContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ContainerStatus(ctx context.Context, req *runtimeapi.ContainerStatusRequest) (*runtimeapi.ContainerStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ListContainers(ctx context.Context, req *runtimeapi.ListContainersRequest) (*runtimeapi.ListContainersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) UpdateContainerResources(ctx context.Context, req *runtimeapi.UpdateContainerResourcesRequest) (*runtimeapi.UpdateContainerResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ContainerStats(ctx context.Context, req *runtimeapi.ContainerStatsRequest) (*runtimeapi.ContainerStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ListContainerStats(ctx context.Context, req *runtimeapi.ListContainerStatsRequest) (*runtimeapi.ListContainerStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) Status(ctx context.Context, req *runtimeapi.StatusRequest) (*runtimeapi.StatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ReopenContainerLog(ctx context.Context, in *runtimeapi.ReopenContainerLogRequest) (*runtimeapi.ReopenContainerLogResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ExecSync(ctx context.Context, in *runtimeapi.ExecSyncRequest) (*runtimeapi.ExecSyncResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) Exec(ctx context.Context, in *runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) Attach(ctx context.Context, in *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) PortForward(ctx context.Context, in *runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) UpdateRuntimeConfig(ctx context.Context, in *runtimeapi.UpdateRuntimeConfigRequest) (*runtimeapi.UpdateRuntimeConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) PodSandboxStats(ctx context.Context, in *runtimeapi.PodSandboxStatsRequest) (*runtimeapi.PodSandboxStatsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) CheckpointContainer(ctx context.Context, req *runtimeapi.CheckpointContainerRequest) (*runtimeapi.CheckpointContainerResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) GetContainerEvents(req *runtimeapi.GetEventsRequest, server runtimeapi.RuntimeService_GetContainerEventsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *criServer) ListMetricDescriptors(ctx context.Context, req *runtimeapi.ListMetricDescriptorsRequest) (*runtimeapi.ListMetricDescriptorsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) ListPodSandboxMetrics(ctx context.Context, req *runtimeapi.ListPodSandboxMetricsRequest) (*runtimeapi.ListPodSandboxMetricsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *criServer) RuntimeConfig(ctx context.Context, req *runtimeapi.RuntimeConfigRequest) (*runtimeapi.RuntimeConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
