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

package impl

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	podresourcesapi "k8s.io/kubelet/pkg/apis/podresources/v1"
)

const (
	podResourcesInformerName PluginName = "podResourcesInformer"
)

var (
	_ podresourcesapi.PodResourcesListerServer = &podResourcesServer{}
)

type podResourcesInformer struct {
	config         *Config
	nodeInformer   *nodeInformer
	resourceServer podresourcesapi.PodResourcesListerServer
}

func newPodResourcesInformer() *podResourcesInformer { _ = "STUB: not implemented"; return nil }

func (s *podResourcesInformer) Setup(ctx *PluginOption, states *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (s *podResourcesInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *podResourcesInformer) startServer(stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanup(grpcServerSocketFullPath string) error { _ = "STUB: not implemented"; return nil }

func (s *podResourcesInformer) HasSynced() bool { _ = "STUB: not implemented"; return false }

type podResourcesServer struct {
	podresourcesapi.UnimplementedPodResourcesListerServer
	podResourceClient podresourcesapi.PodResourcesListerClient
	kubeletStub       KubeletStub
}

func (p *podResourcesServer) List(ctx context.Context, request *podresourcesapi.ListPodResourcesRequest) (*podresourcesapi.ListPodResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fillPodDevicesAllocatedByKoord(response *podresourcesapi.ListPodResourcesResponse, allPods *corev1.PodList) {
	_ = "STUB: not implemented"
	return
}

func (p *podResourcesServer) GetAllocatableResources(ctx context.Context, request *podresourcesapi.AllocatableResourcesRequest) (*podresourcesapi.AllocatableResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *podResourcesServer) Get(ctx context.Context, request *podresourcesapi.GetPodResourcesRequest) (*podresourcesapi.GetPodResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
