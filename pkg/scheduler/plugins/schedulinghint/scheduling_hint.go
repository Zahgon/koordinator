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

package schedulinghint

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	Name = "SchedulingHint"
)

var (
	_ fwktype.PreFilterPlugin                = &Plugin{}
	_ frameworkext.PreFilterTransformer      = &Plugin{}
	_ frameworkext.PreferNodesPluginProvider = &Plugin{}
	_ frameworkext.PreferNodesPlugin         = &Plugin{}
)

type Plugin struct {
	handle       frameworkext.ExtendedHandle
	maxHintNodes int32
}

func New(_ context.Context, args runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (p *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) PreFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) PreFilterExtensions() fwktype.PreFilterExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.PreFilterExtensions)
}

func (p *Plugin) BeforePreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) (*corev1.Pod, bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (p *Plugin) AfterPreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, preRes *fwktype.PreFilterResult) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) PreferNodesPlugin() frameworkext.PreferNodesPlugin {
	_ = "STUB: not implemented"
	return *new(frameworkext.PreferNodesPlugin)
}

func (p *Plugin) PreferNodes(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, result *fwktype.PreFilterResult) ([]string, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// truncate nodes exceeding the max length

// Check if the preferred nodes are in the prefilter result.
// We suppose the hint nodes should be no larger than the prefilter result.
