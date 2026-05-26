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

package scheduling

import (
	"time"

	nrtv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	nrtclientset "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/clientset/versioned"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	"github.com/koordinator-sh/koordinator/test/e2e/framework"
)

type pausePodConfig struct {
	Name                              string
	Namespace                         string
	Affinity                          *corev1.Affinity
	Annotations, Labels, NodeSelector map[string]string
	Resources                         *corev1.ResourceRequirements
	RuntimeClassHandler               *string
	Tolerations                       []corev1.Toleration
	NodeName                          string
	Ports                             []corev1.ContainerPort
	OwnerReferences                   []metav1.OwnerReference
	PriorityClassName                 string
	DeletionGracePeriodSeconds        *int64
	TopologySpreadConstraints         []corev1.TopologySpreadConstraint
	SchedulerName                     string
}

func initPausePod(f *framework.Framework, conf pausePodConfig) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// TODO: setting the Pod's nodeAffinity instead of setting .spec.nodeName works around the
// Preemption e2e flake (#88441), but we should investigate deeper to get to the bottom of it.

func createPausePod(f *framework.Framework, conf pausePodConfig) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func runPausePod(f *framework.Framework, conf pausePodConfig) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func runPausePodWithTimeout(f *framework.Framework, conf pausePodConfig, timeout time.Duration) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func runPodAndGetNodeName(f *framework.Framework, conf pausePodConfig) string {
	_ = "STUB: not implemented"
	// launch a pod to find a node which can launch a pod. We intentionally do
	// not just take the node list and choose the first of them. Depending on the
	// cluster and the scheduler it might be that a "normal" pod cannot be
	// scheduled onto it.
	return ""
}

// GetNodeThatCanRunPod trying to launch a pod without a label to get a node which can launch it
func GetNodeThatCanRunPod(f *framework.Framework) string { _ = "STUB: not implemented"; return "" }

// Get2NodesThatCanRunPod return a 2-node slice where can run pod.
func Get2NodesThatCanRunPod(f *framework.Framework) []string { _ = "STUB: not implemented"; return nil }

type pauseRSConfig struct {
	Replicas  int32
	PodConfig pausePodConfig
}

func initPauseRS(f *framework.Framework, conf pauseRSConfig) *appsv1.ReplicaSet {
	_ = "STUB: not implemented"
	return nil
}

func createPauseRS(f *framework.Framework, conf pauseRSConfig) *appsv1.ReplicaSet {
	_ = "STUB: not implemented"
	return nil
}

func runPauseRS(f *framework.Framework, conf pauseRSConfig) *appsv1.ReplicaSet {
	_ = "STUB: not implemented"
	return nil
}

func waitingForReservationScheduled(koordinatorClientSet koordclientset.Interface, reservation *schedulingv1alpha1.Reservation) *schedulingv1alpha1.Reservation {
	_ = "STUB: not implemented"
	return nil
}

func expectPodBoundReservation(clientSet clientset.Interface, koordinatorClientSet koordclientset.Interface, podNamespace, podName, reservationName string) {
	_ = "STUB: not implemented"
	return
}

func getSuitableNodeResourceTopology(client nrtclientset.Interface, expectNumNodes int) *nrtv1alpha1.NodeResourceTopology {
	_ = "STUB: not implemented"
	return nil
}
