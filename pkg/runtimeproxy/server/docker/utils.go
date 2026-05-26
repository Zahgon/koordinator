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
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"

	"github.com/docker/docker/api/types/container"

	"github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	resource_executor "github.com/koordinator-sh/koordinator/pkg/runtimeproxy/resexecutor"
)

type mockRespWriter struct {
	http.ResponseWriter
	w    io.Writer
	code int
}

func (m *mockRespWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *mockRespWriter) WriteHeader(c int) { _ = "STUB: not implemented"; return }

func (m mockRespWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func calculateContentLength(body io.Reader) (l int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func encodeBody(obj interface{}) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HostConfigToResource(config *container.HostConfig) *v1alpha1.LinuxContainerResources {
	_ = "STUB: not implemented"
	return nil
}

func getContainerID(urlPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func splitLabelsAndAnnotations(configs map[string]string) (labels map[string]string, annos map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToCriCgroupPath(cgroupDriver, cgroupParent string) string {
	_ = "STUB: not implemented"
	return ""
}

// for docker + systemd combination, the cgroup parent is pod dir, for example:
// kubepods-besteffort-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice,
// so the full path for it is :
// /kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice

func GetRuntimeResourceType(labels map[string]string) resource_executor.RuntimeResourceType {
	_ = "STUB: not implemented"
	return *new(resource_executor.RuntimeResourceType)
}

func UpdateHostConfigByResource(config *container.HostConfig, resources *v1alpha1.LinuxContainerResources) *container.HostConfig {
	_ = "STUB: not implemented"
	return nil
}

func UpdateUpdateConfigByResource(containerConfig *container.UpdateConfig, resources *v1alpha1.LinuxContainerResources) *container.UpdateConfig {
	_ = "STUB: not implemented"
	return nil
}

func MergeResourceByUpdateConfig(resources *v1alpha1.LinuxContainerResources, containerConfig *container.UpdateConfig) *v1alpha1.LinuxContainerResources {
	_ = "STUB: not implemented"
	return nil
}

// generateExpectedCgroupParent is adapted from Dockershim
func generateExpectedCgroupParent(cgroupDriver, cgroupParent string) string {
	_ = "STUB: not implemented"
	return ""

	// if docker uses the systemd cgroup driver, it expects *.slice style names for cgroup parent.
	// if we configured kubelet to use --cgroup-driver=cgroupfs, and docker is configured to use systemd driver
	// docker will fail to launch the container because the name we provide will not be a valid slice.
	// this is a very good thing.
}

// Pass only the last component of the cgroup path to systemd.

func splitDockerEnv(dockerEnvs []string) map[string]string { _ = "STUB: not implemented"; return nil }

func generateEnvList(envs map[string]string) (result []string) {
	_ = "STUB: not implemented"
	return nil
}
