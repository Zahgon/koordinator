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

package pod

import (
	"time"

	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	// PodDeleteTimeout is how long to wait for a pod to be deleted.
	PodDeleteTimeout = 5 * time.Minute
)

// DeletePodOrFail deletes the pod of the specified namespace and name. Resilient to the pod
// not existing.
func DeletePodOrFail(c clientset.Interface, ns, name string) { _ = "STUB: not implemented"; return }

// DeletePodWithWait deletes the passed-in pod and waits for the pod to be terminated. Resilient to the pod
// not existing.
func DeletePodWithWait(c clientset.Interface, pod *v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// DeletePodWithWaitByName deletes the named and namespaced pod and waits for the pod to be terminated. Resilient to the pod
// not existing.
func DeletePodWithWaitByName(c clientset.Interface, podName, podNamespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// assume pod was already deleted

// DeletePodWithGracePeriod deletes the passed-in pod. Resilient to the pod not existing.
func DeletePodWithGracePeriod(c clientset.Interface, pod *v1.Pod, grace int64) error {
	_ = "STUB: not implemented"
	return nil
}

// DeletePodsWithGracePeriod deletes the passed-in pods. Resilient to the pods not existing.
func DeletePodsWithGracePeriod(c clientset.Interface, pods []v1.Pod, grace int64) error {
	_ = "STUB: not implemented"
	return nil
}

// DeletePodWithGracePeriodByName deletes a pod by name and namespace. Resilient to the pod not existing.
func DeletePodWithGracePeriodByName(c clientset.Interface, podName, podNamespace string, grace int64) error {
	_ = "STUB: not implemented"
	return nil
}

// assume pod was already deleted
