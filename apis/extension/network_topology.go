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

package extension

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (

	// AnnotationGangPodNetworkTopologyIndex defines the index of a specific pod in the gang group which should typically
	// match the underlying biz semantics (such as PyTorch rank).
	// This would currently be considered in network topology aware scheduling.
	AnnotationGangPodNetworkTopologyIndex = AnnotationGangPrefix + "/network-topology-index"

	// AnnotationGangNetworkTopologySpec defines the network topology aware requirements of the gang group.
	AnnotationGangNetworkTopologySpec = AnnotationGangPrefix + "/network-topology-spec"

	AnnotationPodNetworkTopologySelector = SchedulingDomainPrefix + "/network-topology-selector"
)

type NetworkTopologySpec struct {
	GatherStrategy []NetworkTopologyGatherRule `json:"gatherStrategy,omitempty"`
}

type NetworkTopologyGatherRule struct {
	Layer            schedulingv1alpha1.TopologyLayer `json:"layer"`
	Strategy         NetworkTopologyGatherStrategy    `json:"strategy"`
	PodCountMultiple int                              `json:"podCountMultiple,omitempty"`
}

type NetworkTopologyGatherStrategy string

const (
	NetworkTopologyGatherStrategyMustGather   NetworkTopologyGatherStrategy = "MustGather"
	NetworkTopologyGatherStrategyPreferGather NetworkTopologyGatherStrategy = "PreferGather"
)

func GetNetworkTopologySpec(obj metav1.Object) (*NetworkTopologySpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetPodNetworkTopologyIndex(pod *corev1.Pod) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SortPodsByIndex sort pods by index than by name,
// pods without valid index would be ordered after pods with index.
func SortPodsByIndex(pods []*corev1.Pod) { _ = "STUB: not implemented"; return }

func GetPodNetworkTopologySelector(obj metav1.Object) string { _ = "STUB: not implemented"; return "" }
