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
	"net/http"

	dclient "github.com/docker/docker/client"
)

var GetDockerClient = createDockerClient // for test

func GetDockerEndpoint() string { _ = "STUB: not implemented"; return "" }

type DockerRuntimeHandler struct {
	dockerClient *dclient.Client
	endpoint     string
}

func NewDockerRuntimeHandler(endpoint string) (ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(ContainerRuntimeHandler), nil
}

func createDockerClient(httpClient *http.Client, endPoint string) (*dclient.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DockerRuntimeHandler) StopContainer(c context.Context, containerID string, timeout int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DockerRuntimeHandler) UpdateContainerResources(containerID string, opts UpdateOptions) error {
	_ = "STUB: not implemented"
	return nil
}
