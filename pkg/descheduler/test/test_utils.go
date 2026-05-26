/*
Copyright 2022 The Koordinator Authors.
Copyright 2017 The Kubernetes Authors.

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

package test

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/informers/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

// BuildTestPod creates a test pod with given parameters.
func BuildTestPod(name string, cpu int64, memory int64, nodeName string, apply func(*corev1.Pod)) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// GetMirrorPodAnnotation returns the annotation needed for mirror pod.
func GetMirrorPodAnnotation() map[string]string { _ = "STUB: not implemented"; return nil }

// GetNormalPodOwnerRefList returns the ownerRef needed for a pod.
func GetNormalPodOwnerRefList() []metav1.OwnerReference { _ = "STUB: not implemented"; return nil }

// GetReplicaSetOwnerRefList returns the ownerRef needed for replicaset pod.
func GetReplicaSetOwnerRefList() []metav1.OwnerReference { _ = "STUB: not implemented"; return nil }

// GetStatefulSetOwnerRefList returns the ownerRef needed for statefulset pod.
func GetStatefulSetOwnerRefList() []metav1.OwnerReference { _ = "STUB: not implemented"; return nil }

// GetDaemonSetOwnerRefList returns the ownerRef needed for daemonset pod.
func GetDaemonSetOwnerRefList() []metav1.OwnerReference { _ = "STUB: not implemented"; return nil }

// BuildTestNode creates a node with specified capacity.
func BuildTestNode(name string, millicpu int64, mem int64, pods int64, apply func(*corev1.Node)) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// MakeBestEffortPod makes the given pod a BestEffort pod
func MakeBestEffortPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// MakeBurstablePod makes the given pod a Burstable pod
func MakeBurstablePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// MakeGuaranteedPod makes the given pod an Guaranteed pod
func MakeGuaranteedPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// SetRSOwnerRef sets the given pod's owner to ReplicaSet
func SetRSOwnerRef(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// SetSSOwnerRef sets the given pod's owner to StatefulSet
func SetSSOwnerRef(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// SetDSOwnerRef sets the given pod's owner to DaemonSet
func SetDSOwnerRef(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// SetNormalOwnerRef sets the given pod's owner to Pod
func SetNormalOwnerRef(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// SetPodPriority sets the given pod's priority
func SetPodPriority(pod *corev1.Pod, priority int32) { _ = "STUB: not implemented"; return }

// SetNodeUnschedulable sets the given node unschedulable
func SetNodeUnschedulable(node *corev1.Node) { _ = "STUB: not implemented"; return }

// SetPodExtendedResourceRequest sets the given pod's extended resources
func SetPodExtendedResourceRequest(pod *corev1.Pod, resourceName corev1.ResourceName, requestQuantity int64) {
	_ = "STUB: not implemented"
	return
}

// SetNodeExtendedResouces sets the given node's extended resources
func SetNodeExtendedResource(node *corev1.Node, resourceName corev1.ResourceName, requestQuantity int64) {
	_ = "STUB: not implemented"
	return
}

const (
	nodeNameKeyIndex = "spec.nodeName"
)

// BuildGetPodsAssignedToNodeFunc establishes an indexer to map the pods and their assigned nodes.
// It returns a function to help us get all the pods that assigned to a node based on the indexer.
func BuildGetPodsAssignedToNodeFunc(podInformer v1.PodInformer) (framework.GetPodsAssignedToNodeFunc, error) {
	_ = "STUB: not implemented"
	// Establish an indexer to map the pods and their assigned nodes.
	return *new(framework.GetPodsAssignedToNodeFunc), nil
}

// The indexer helps us get all the pods that assigned to a node.
