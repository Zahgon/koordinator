/*
Copyright 2022 The Koordinator Authors.
Copyright 2019 The Kubernetes Authors.

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

package kubectl

import (
	"os/exec"

	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	maxKubectlExecRetries = 5
)

// TestKubeconfig is a struct containing the needed attributes from TestContext and Framework(Namespace).
type TestKubeconfig struct {
	CertDir     string
	Host        string
	KubeConfig  string
	KubeContext string
	KubectlPath string
	Namespace   string // Every test has at least one namespace unless creation is skipped
}

// NewTestKubeconfig returns a new Kubeconfig struct instance.
func NewTestKubeconfig(certdir, host, kubeconfig, kubecontext, kubectlpath, namespace string) *TestKubeconfig {
	_ = "STUB: not implemented"
	return nil
}

// KubectlCmd runs the kubectl executable through the wrapper script.
func (tk *TestKubeconfig) KubectlCmd(args ...string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil

	// Reference a --server option so tests can run anywhere.
}

// Reference the KubeContext

// We allow users to specify path to kubectl, so you can test either "kubectl" or "cluster/kubectl.sh"
// and so on.

// caller will invoke this and wait on it.

// LogFailedContainers runs `kubectl logs` on a failed containers.
func LogFailedContainers(c clientset.Interface, ns string, logFunc func(ftm string, args ...interface{})) {
	_ = "STUB: not implemented"
	return
}

func kubectlLogPod(c clientset.Interface, pod v1.Pod, containerNameSubstr string, logFunc func(ftm string, args ...interface{})) {
	_ = "STUB: not implemented"
	return
}

// Contains() matches all strings if substr is empty

// WriteFileViaContainer writes a file using kubectl exec echo <contents> > <path> via specified container
// because of the primitive technique we're using here, we only allow ASCII alphanumeric characters
func (tk *TestKubeconfig) WriteFileViaContainer(podName, containerName string, path string, contents string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(mauriciopoppe): remove this statement once we add `sync` to the test image, ref #101172

// ReadFileViaContainer reads a file using kubectl exec cat <path>.
func (tk *TestKubeconfig) ReadFileViaContainer(podName, containerName string, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (tk *TestKubeconfig) kubectlExecWithRetry(namespace string, podName, containerName string, args ...string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Retry on "i/o timeout" errors

// Retry on "container not found" errors

func (tk *TestKubeconfig) kubectlExec(namespace string, podName, containerName string, args ...string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
