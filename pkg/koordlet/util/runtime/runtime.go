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

package runtime

import (
	"sync"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/runtime/handler"
)

var (
	DockerHandler     handler.ContainerRuntimeHandler
	ContainerdHandler handler.ContainerRuntimeHandler
	PouchHandler      handler.ContainerRuntimeHandler
	CrioHandler       handler.ContainerRuntimeHandler
	mutex             = &sync.Mutex{}
)

func GetRuntimeHandler(runtimeType string) (handler.ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(handler.ContainerRuntimeHandler), nil
}

func getDockerHandler() (handler.ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(handler.ContainerRuntimeHandler), nil
}

func getDockerEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getContainerdHandler() (handler.ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(handler.ContainerRuntimeHandler), nil
}

func getContainerdEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getPouchHandler() (handler.ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(handler.ContainerRuntimeHandler), nil
}

func getPouchEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getCrioHandler() (handler.ContainerRuntimeHandler, error) {
	_ = "STUB: not implemented"
	return *new(handler.ContainerRuntimeHandler), nil
}

func getCrioEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }

func isFile(path string) bool { _ = "STUB: not implemented"; return false }
