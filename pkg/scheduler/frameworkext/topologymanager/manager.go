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

package topologymanager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/kube-scheduler/framework"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

const (
	ErrUnsatisfiedNUMAResource = "Unsatisfied NUMA %s"
	ErrNUMAHintCannotAligned   = "Unaligned NUMA Hint cause %v"
)

type Interface interface {
	Admit(ctx context.Context, cycleState framework.CycleState, pod *corev1.Pod, node *corev1.Node, numaNodes []int, policyType apiext.NUMATopologyPolicy, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) *fwktype.Status
}

type NUMATopologyHintProvider interface {
	// GetPodTopologyHints returns a map of resource names to a list of possible
	// concrete resource allocations per Pod in terms of NUMA locality hints.
	GetPodTopologyHints(ctx context.Context, cycleState framework.CycleState, pod *corev1.Pod, node *corev1.Node) (map[string][]NUMATopologyHint, *fwktype.Status)
	// Allocate triggers resource allocation to occur on the HintProvider after
	// all hints have been gathered and the aggregated Hint
	Allocate(ctx context.Context, cycleState framework.CycleState, affinity NUMATopologyHint, pod *corev1.Pod, node *corev1.Node) *fwktype.Status
}

var _ Interface = &topologyManager{}

type topologyManager struct {
	hintProviderFactory NUMATopologyHintProviderFactory
}

type NUMATopologyHintProviderFactory interface {
	GetNUMATopologyHintProvider() []NUMATopologyHintProvider
}

func New(hintProviderFactory NUMATopologyHintProviderFactory) Interface {
	_ = "STUB: not implemented"
	return *new(Interface)
}

func (m *topologyManager) Admit(ctx context.Context, cycleState framework.CycleState, pod *corev1.Pod, node *corev1.Node, numaNodes []int, policyType apiext.NUMATopologyPolicy, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// TODO If the Affinity above is confirmed to be allocatable, it seems unnecessary to call this again here.

func (m *topologyManager) calculateAffinity(ctx context.Context, cycleState framework.CycleState, policy Policy, pod *corev1.Pod, node *corev1.Node, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) (NUMATopologyHint, bool, []string) {
	_ = "STUB: not implemented"
	return *new(NUMATopologyHint), false, nil
}

func (m *topologyManager) accumulateProvidersHints(ctx context.Context, cycleState framework.CycleState, pod *corev1.Pod, node *corev1.Node) ([]map[string][]NUMATopologyHint, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the TopologyHints for a Pod from a provider.

func (m *topologyManager) allocateResources(ctx context.Context, cycleState framework.CycleState, affinity NUMATopologyHint, pod *corev1.Pod, node *corev1.Node) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func createNUMATopologyPolicy(policyType apiext.NUMATopologyPolicy, numaNodes []int) Policy {
	_ = "STUB: not implemented"
	return *new(Policy)
}
