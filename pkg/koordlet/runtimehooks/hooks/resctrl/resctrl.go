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

package resctrl

import (
	"k8s.io/client-go/tools/record"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	util "github.com/koordinator-sh/koordinator/pkg/koordlet/util/resctrl"
)

const (
	name               = "Resctrl"
	description        = "set resctrl for pod"
	ruleNameForAllPods = name + " (AllPods)"
)

type plugin struct {
	rule           *Rule
	engine         util.ResctrlEngine
	executor       resourceexecutor.ResourceUpdateExecutor
	statesInformer statesinformer.StatesInformer
	EventRecorder  record.EventRecorder
}

var singleton *plugin

func Object() *plugin { _ = "STUB: not implemented"; return nil }

func newPlugin() *plugin { _ = "STUB: not implemented"; return nil }

func (p *plugin) Register(op hooks.Options) {
	_ = "STUB: not implemented"
	// skip if host not support resctrl
	return
}

// check if the resctrl root and l3_cat feature are enabled correctly

func (p *plugin) SetPodResctrlResourcesForHooks(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) SetPodResctrlResourcesForReconciler(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) setPodResctrlResources(proto protocol.HooksProtocol, fromNRI bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) RemoveUnusedResctrlPath(protos []protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) UpdatePodTaskIds(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) SetContainerResctrlResources(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) RemovePodResctrlResources(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}
