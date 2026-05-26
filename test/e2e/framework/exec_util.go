/*
Copyright 2016 The Kubernetes Authors.

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

package framework

import (
	"io"
	"net/url"

	restclient "k8s.io/client-go/rest"
)

// ExecOptions passed to ExecWithOptions
type ExecOptions struct {
	Command       []string
	Namespace     string
	PodName       string
	ContainerName string
	Stdin         io.Reader
	CaptureStdout bool
	CaptureStderr bool
	// If false, whitespace in std{err,out} will be removed.
	PreserveWhitespace bool
	Quiet              bool
}

// ExecWithOptions executes a command in the specified container,
// returning stdout, stderr and error. `options` allowed for
// additional parameters to be passed.
func (f *Framework) ExecWithOptions(options ExecOptions) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ExecCommandInContainerWithFullOutput executes a command in the
// specified container and return stdout, stderr and error
func (f *Framework) ExecCommandInContainerWithFullOutput(podName, containerName string, cmd ...string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ExecCommandInContainer executes a command in the specified container.
func (f *Framework) ExecCommandInContainer(podName, containerName string, cmd ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// ExecShellInContainer executes the specified command on the pod's container.
func (f *Framework) ExecShellInContainer(podName, containerName string, cmd string) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Framework) execCommandInPod(podName string, cmd ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *Framework) execCommandInPodWithFullOutput(podName string, cmd ...string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ExecShellInPod executes the specified command on the pod.
func (f *Framework) ExecShellInPod(podName string, cmd string) string {
	_ = "STUB: not implemented"
	return ""
}

// ExecShellInPodWithFullOutput executes the specified command on the Pod and returns stdout, stderr and error.
func (f *Framework) ExecShellInPodWithFullOutput(podName string, cmd string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool) error {
	_ = "STUB: not implemented"
	return nil
}
