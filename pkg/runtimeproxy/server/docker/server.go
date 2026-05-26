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

package docker

import (
	"context"
	"net/http"
	"regexp"

	dockertypes "github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
	dockersystem "github.com/docker/docker/api/types/system"

	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/dispatcher"
	"github.com/koordinator-sh/koordinator/pkg/util/httputil"
)

type RuntimeManagerDockerServer struct {
	dispatcher   *dispatcher.RuntimeHookDispatcher
	reverseProxy *httputil.ReverseProxy
	router       map[*regexp.Regexp]func(context.Context, http.ResponseWriter, *http.Request)
	cgroupDriver string
}

type proxyDockerClient interface {
	Info(ctx context.Context) (dockersystem.Info, error)
	ContainerList(ctx context.Context, options dockercontainer.ListOptions) ([]dockercontainer.Summary, error)
	ContainerInspect(ctx context.Context, containerID string) (dockertypes.ContainerJSON, error)
}

func (d *RuntimeManagerDockerServer) Name() string { _ = "STUB: not implemented"; return "" }

func NewRuntimeManagerDockerServer() *RuntimeManagerDockerServer {
	_ = "STUB: not implemented"
	return nil
}

func (d *RuntimeManagerDockerServer) Direct(wr http.ResponseWriter, req *http.Request) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *RuntimeManagerDockerServer) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// fall back to reverse proxy

func (d *RuntimeManagerDockerServer) failOver(dockerClient proxyDockerClient) error {
	_ = "STUB: not implemented"
	return nil
}

// need to backup pod meta first

func (d *RuntimeManagerDockerServer) Run() error { _ = "STUB: not implemented"; return nil }
