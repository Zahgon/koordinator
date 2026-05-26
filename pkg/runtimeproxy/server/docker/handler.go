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
)

func (d *RuntimeManagerDockerServer) HandleCreateContainer(ctx context.Context, wr http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	// get create container config
	return
}

// pre check

// refuse the req

// send req to docker

func (d *RuntimeManagerDockerServer) HandleStartContainer(ctx context.Context, wr http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	// we need to get the container id, because we need it to get info from checkpoint
	return
}

// no need to care about the resp

func (d *RuntimeManagerDockerServer) HandleStopContainer(ctx context.Context, wr http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// sandbox container

// kubelet will not treat not found error as error, so we need to return err msg as same with docker server to avoid pod terminating

func (d *RuntimeManagerDockerServer) HandleUpdateContainer(ctx context.Context, wr http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// update resources in cache with UpdateConfig

// send req to docker
