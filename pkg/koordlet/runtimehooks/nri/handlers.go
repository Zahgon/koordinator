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

package nri

import (
	"context"

	"github.com/containerd/nri/pkg/api"
	"github.com/containerd/nri/pkg/stub"
)

func (p *NriServer) Configure(_ context.Context, config, runtime, version string) (stub.EventMask, error) {
	_ = "STUB: not implemented"
	return *new(stub.EventMask), nil
}

func (p *NriServer) Synchronize(_ context.Context, pods []*api.PodSandbox, containers []*api.Container) ([]*api.ContainerUpdate, error) {
	_ = "STUB: not implemented"
	// todo: update existed containers configure
	return nil, nil
}

func (p *NriServer) RunPodSandbox(_ context.Context, pod *api.PodSandbox) error {
	_ = "STUB: not implemented"
	return nil
}

// todo: return error or bypass error based on PluginFailurePolicy

func (p *NriServer) CreateContainer(_ context.Context, pod *api.PodSandbox, container *api.Container) (*api.ContainerAdjustment, []*api.ContainerUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// todo: return error or bypass error based on PluginFailurePolicy

func (p *NriServer) UpdateContainer(_ context.Context, pod *api.PodSandbox, container *api.Container, r *api.LinuxResources) ([]*api.ContainerUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// todo: return error or bypass error based on PluginFailurePolicy

func (p *NriServer) RemovePodSandbox(_ context.Context, pod *api.PodSandbox) error {
	_ = "STUB: not implemented"
	return nil
}

// todo: return error or bypass error based on PluginFailurePolicy

func (p *NriServer) RemoveContainer(context.Context, *api.PodSandbox, *api.Container) error {
	_ = "STUB: not implemented"
	// TODO
	return nil
}

func (p *NriServer) StopContainer(context.Context, *api.PodSandbox, *api.Container) ([]*api.ContainerUpdate, error) {
	_ = "STUB: not implemented"
	// TODO
	return nil, nil
}

func (p *NriServer) StartContainer(context.Context, *api.PodSandbox, *api.Container) error {
	_ = "STUB: not implemented"
	// TODO
	return nil
}

func (p *NriServer) StopPodSandbox(context.Context, *api.PodSandbox) error {
	_ = "STUB: not implemented"
	// TODO
	return nil
}
