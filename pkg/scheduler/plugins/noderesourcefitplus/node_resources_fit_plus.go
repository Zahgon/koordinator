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

package noderesourcesfitplus

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fwk "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

const (
	// Name is plugin name
	Name = "NodeResourcesFitPlus"

	preScoreStateKey = "PreScore" + Name
)

var (
	_ fwk.ScorePlugin = &Plugin{}
)

type Plugin struct {
	handle fwk.Handle
	args   *config.NodeResourcesFitPlusArgs
}

func New(_ context.Context, args runtime.Object, handle fwk.Handle) (fwk.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwk.Plugin), nil
}

func (s *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

type preScoreState struct {
	framework.Resource
	ResourceName []v1.ResourceName
}

// Clone the prefilter state.
func (s *preScoreState) Clone() fwk.StateData {
	_ = "STUB: not implemented"
	return *new(fwk.StateData)
}

func (s *Plugin) PreScore(ctx context.Context, cycleState fwk.CycleState, pod *v1.Pod, nodes []fwk.NodeInfo) *fwk.Status {
	_ = "STUB: not implemented"
	return nil
}

func (s *Plugin) Score(ctx context.Context, state fwk.CycleState, p *v1.Pod, nodeInfo fwk.NodeInfo) (int64, *fwk.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Plugin) ScoreExtensions() fwk.ScoreExtensions {
	_ = "STUB: not implemented"
	return *new(fwk.ScoreExtensions)
}

func fitsPodRequestName(podRequest framework.Resource) []v1.ResourceName {
	_ = "STUB: not implemented"
	return nil
}

func getPreScoreState(cycleState fwk.CycleState) (*preScoreState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// preFilterState doesn't exist, likely PreFilter wasn't invoked.
