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

package statefulset

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// WaitForRunning waits for numPodsRunning in ss to be Running and for the first
// numPodsReady ordinals to be Ready.
func WaitForRunning(c clientset.Interface, numPodsRunning, numPodsReady int32, ss *appsv1.StatefulSet) {
	_ = "STUB: not implemented"
	return
}

// WaitForState periodically polls for the ss and its pods until the until function returns either true or an error
func WaitForState(c clientset.Interface, ss *appsv1.StatefulSet, until func(*appsv1.StatefulSet, *v1.PodList) (bool, error)) {
	_ = "STUB: not implemented"
	return
}

// WaitForRunningAndReady waits for numStatefulPods in ss to be Running and Ready.
func WaitForRunningAndReady(c clientset.Interface, numStatefulPods int32, ss *appsv1.StatefulSet) {
	_ = "STUB: not implemented"
	return
}

// WaitForPodReady waits for the Pod named podName in set to exist and have a Ready condition.
func WaitForPodReady(c clientset.Interface, set *appsv1.StatefulSet, podName string) (*appsv1.StatefulSet, *v1.PodList) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WaitForStatusReadyReplicas waits for the ss.Status.ReadyReplicas to be equal to expectedReplicas
func WaitForStatusReadyReplicas(c clientset.Interface, ss *appsv1.StatefulSet, expectedReplicas int32) {
	_ = "STUB: not implemented"
	return
}

// WaitForStatusAvailableReplicas waits for the ss.Status.Available to be equal to expectedReplicas
func WaitForStatusAvailableReplicas(c clientset.Interface, ss *appsv1.StatefulSet, expectedReplicas int32) {
	_ = "STUB: not implemented"
	return
}

// WaitForStatusReplicas waits for the ss.Status.Replicas to be equal to expectedReplicas
func WaitForStatusReplicas(c clientset.Interface, ss *appsv1.StatefulSet, expectedReplicas int32) {
	_ = "STUB: not implemented"
	return
}

// Saturate waits for all Pods in ss to become Running and Ready.
func Saturate(c clientset.Interface, ss *appsv1.StatefulSet) { _ = "STUB: not implemented"; return }
