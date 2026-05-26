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
	"context"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/topologymanager"
)

func (p *Plugin) FilterByNUMANode(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, node *corev1.Node, policyType apiext.NUMATopologyPolicy, exclusivePolicy apiext.NumaTopologyExclusive, topologyOptions TopologyOptions) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) GetPodTopologyHints(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, node *corev1.Node) (map[string][]topologymanager.NUMATopologyHint, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we have check in filter, so we will not get error in reserve

func (p *Plugin) Allocate(ctx context.Context, cycleState fwktype.CycleState, affinity topologymanager.NUMATopologyHint, pod *corev1.Pod, node *corev1.Node) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}
