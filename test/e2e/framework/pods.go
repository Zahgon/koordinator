/*
Copyright 2022 The Koordinator Authors.
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
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	// TODO: Remove the following imports (ref: https://github.com/kubernetes/kubernetes/issues/81245)
)

const (
	// DefaultPodDeletionTimeout is the default timeout for deleting pod
	DefaultPodDeletionTimeout = 3 * time.Minute

	// the status of container event, copied from k8s.io/kubernetes/pkg/kubelet/events
	killingContainer = "Killing"

	// the status of container event, copied from k8s.io/kubernetes/pkg/kubelet/events
	failedToCreateContainer = "Failed"

	// the status of container event, copied from k8s.io/kubernetes/pkg/kubelet/events
	startedContainer = "Started"

	// it is copied from k8s.io/kubernetes/pkg/kubelet/sysctl
	forbiddenReason = "SysctlForbidden"
)

// ImagePrePullList is the images used in the current test suite. It should be initialized in test suite and
// the images in the list should be pre-pulled in the test suite.  Currently, this is only used by
// node e2e test.
var ImagePrePullList sets.String

// PodClient is a convenience method for getting a pod client interface in the framework's namespace,
// possibly applying test-suite specific transformations to the pod spec, e.g. for
// node e2e pod scheduling.
func (f *Framework) PodClient() *PodClient { _ = "STUB: not implemented"; return nil }

// PodClientNS is a convenience method for getting a pod client interface in an alternative namespace,
// possibly applying test-suite specific transformations to the pod spec, e.g. for
// node e2e pod scheduling.
func (f *Framework) PodClientNS(namespace string) *PodClient { _ = "STUB: not implemented"; return nil }

// PodClient is a struct for pod client.
type PodClient struct {
	f *Framework
	v1core.PodInterface
}

// Create creates a new pod according to the framework specifications (don't wait for it to start).
func (c *PodClient) Create(pod *v1.Pod) *v1.Pod { _ = "STUB: not implemented"; return nil }

// CreateSync creates a new pod according to the framework specifications, and wait for it to start and be running and ready.
func (c *PodClient) CreateSync(pod *v1.Pod) *v1.Pod { _ = "STUB: not implemented"; return nil }

// Get the newest pod after it becomes running and ready, some status may change after pod created, such as pod ip.

// CreateBatch create a batch of pods. All pods are created before waiting.
func (c *PodClient) CreateBatch(pods []*v1.Pod) []*v1.Pod { _ = "STUB: not implemented"; return nil }

// Update updates the pod object. It retries if there is a conflict, throw out error if
// there is any other apierrors. name is the pod name, updateFn is the function updating the
// pod object.
func (c *PodClient) Update(name string, updateFn func(pod *v1.Pod)) {
	_ = "STUB: not implemented"
	return
}

// DeleteSync deletes the pod and wait for the pod to disappear for `timeout`. If the pod doesn't
// disappear before the timeout, it will fail the test.
func (c *PodClient) DeleteSync(name string, options metav1.DeleteOptions, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// mungeSpec apply test-suite specific transformations to the pod spec.
func (c *PodClient) mungeSpec(pod *v1.Pod) { _ = "STUB: not implemented"; return }

// Node e2e does not support the default DNSClusterFirst policy. Set
// the policy to DNSDefault, which is configured per node.

// PrepullImages only works for node e2e now. For cluster e2e, image prepull is not enforced,
// we should not munge ImagePullPolicy for cluster e2e pods.

// If prepull is enabled, munge the container spec to make sure the images are not pulled
// during the test.

// If the image pull policy is PullAlways, the image doesn't need to be in
// the allow list or pre-pulled, because the image is expected to be pulled
// in the test anyway.

// If the image policy is not PullAlways, the image must be in the pre-pull list and
// pre-pulled.

// Do not pull images during the tests because the images in pre-pull list should have
// been prepulled.

// WaitForSuccess waits for pod to succeed.
// TODO(random-liu): Move pod wait function into this file
func (c *PodClient) WaitForSuccess(name string, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// WaitForFinish waits for pod to finish running, regardless of success or failure.
func (c *PodClient) WaitForFinish(name string, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// WaitForErrorEventOrSuccess waits for pod to succeed or an error event for that pod.
func (c *PodClient) WaitForErrorEventOrSuccess(pod *v1.Pod) (*v1.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ignore all other errors

// MatchContainerOutput gets output of a container and match expected regexp in the output.
func (c *PodClient) MatchContainerOutput(name string, containerName string, expectedRegexp string) error {
	_ = "STUB: not implemented"
	return nil
}

// PodIsReady returns true if the specified pod is ready. Otherwise false.
func (c *PodClient) PodIsReady(name string) bool { _ = "STUB: not implemented"; return false }
