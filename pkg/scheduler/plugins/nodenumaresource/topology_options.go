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

package nodenumaresource

import (
	"sync"

	nrtv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

// TopologyOptionsManager manages the System Topology and resource assignments options.
type TopologyOptionsManager interface {
	GetTopologyOptions(nodeName string) TopologyOptions
	UpdateTopologyOptions(nodeName string, updateFn func(options *TopologyOptions))
	Delete(nodeName string)
}

type TopologyOptions struct {
	NodeResourceTopology *nrtv1alpha1.NodeResourceTopology       `json:"-"`
	CPUTopology          *CPUTopology                            `json:"cpuTopology"`
	ReservedCPUs         cpuset.CPUSet                           `json:"reservedCPUs"`
	MaxRefCount          int                                     `json:"maxRefCount"`
	Policy               *extension.KubeletCPUManagerPolicy      `json:"policy,omitempty"`
	NUMATopologyPolicy   extension.NUMATopologyPolicy            `json:"numaTopologyPolicy"`
	NUMANodeResources    []NUMANodeResource                      `json:"numaNodeResources"`
	AmplificationRatios  map[corev1.ResourceName]extension.Ratio `json:"amplificationRatios,omitempty"`
}

type NUMANodeResource struct {
	Node      int                 `json:"node"`
	Resources corev1.ResourceList `json:"resources,omitempty"`
}

type topologyManager struct {
	lock            sync.RWMutex
	topologyOptions map[string]TopologyOptions
}

func NewTopologyOptionsManager() TopologyOptionsManager {
	_ = "STUB: not implemented"
	return *new(TopologyOptionsManager)
}

func (m *topologyManager) GetTopologyOptions(nodeName string) TopologyOptions {
	_ = "STUB: not implemented"
	return *new(TopologyOptions)
}

func (m *topologyManager) UpdateTopologyOptions(nodeName string, updateFn func(options *TopologyOptions)) {
	_ = "STUB: not implemented"
	return
}

func (m *topologyManager) Delete(nodeName string) { _ = "STUB: not implemented"; return }

func NewTopologyOptions(nrt *nrtv1alpha1.NodeResourceTopology) TopologyOptions {
	_ = "STUB: not implemented"
	return *new(TopologyOptions)
}

// remove cpus reserved by node.annotation.

// reservedCPUs = cpus(all) - cpus(guaranteed) - cpus(kubeletReserved) - cpus(nodeReservationReserved) - cpus(systemQOSReserved)

func getPodAllocsCPUSet(podCPUAllocs extension.PodCPUAllocs) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

func convertCPUTopology(reportedCPUTopology *extension.CPUTopology) *CPUTopology {
	_ = "STUB: not implemented"
	return nil
}

func extractNUMANodeResources(nrt *nrtv1alpha1.NodeResourceTopology) []NUMANodeResource {
	_ = "STUB: not implemented"
	return nil
}

func convertToNUMATopologyPolicy(nrt *nrtv1alpha1.NodeResourceTopology) extension.NUMATopologyPolicy {
	_ = "STUB: not implemented"
	return *new(extension.NUMATopologyPolicy)
}

func (opts *TopologyOptions) getNUMANodes() []int { _ = "STUB: not implemented"; return nil }
