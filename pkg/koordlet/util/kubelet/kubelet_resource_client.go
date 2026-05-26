/*
Copyright (c) 2019 Intel Corporation
Copyright (c) 2021 Multus Authors
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

package kubelet

import (
	"net"
	"net/url"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	podresourcesapi "k8s.io/kubelet/pkg/apis/podresources/v1"
)

const (
	defaultKubeletSocket       = "kubelet" // which is defined in k8s.io/kubernetes/pkg/kubelet/apis/podresources
	kubeletConnectionTimeout   = 10 * time.Second
	defaultPodResourcesMaxSize = 1024 * 1024 * 16 // 16 Mb
	defaultPodResourcesPath    = "/var/lib/kubelet/pod-resources"
	unixProtocol               = "unix"
)

// LocalEndpoint returns the full path to a unix socket at the given endpoint
// which is in k8s.io/kubernetes/pkg/kubelet/util
func localEndpoint(path string) *url.URL { _ = "STUB: not implemented"; return nil }

// GetResourceClient returns an instance of ResourceClient interface initialized with Pod resource information
func GetResourceClient(kubeletSocket string) (podresourcesapi.PodResourcesListerClient, error) {
	_ = "STUB: not implemented"
	return *new(podresourcesapi.PodResourcesListerClient), nil
}

// If Kubelet resource API endpoint exist use that by default
// Or else fallback with checkpoint file

func hasKubeletAPIEndpoint(url *url.URL) bool {
	_ = "STUB: not implemented"
	// Check for kubelet resource API socket file
	return false
}

func getKubeletResourceClient(kubeletSocketURL *url.URL, timeout time.Duration) (podresourcesapi.PodResourcesListerClient, *grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return *new(podresourcesapi.PodResourcesListerClient), nil, nil
}

func dial(ctx context.Context, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
