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

package hooks

import (
	"k8s.io/client-go/tools/record"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	rmconfig "github.com/koordinator-sh/koordinator/pkg/runtimeproxy/config"
)

type Hook struct {
	name        string
	stage       rmconfig.RuntimeHookType
	description string
	fn          HookFn
}

type Options struct {
	Reader                           resourceexecutor.CgroupReader
	Executor                         resourceexecutor.ResourceUpdateExecutor
	StatesInformer                   statesinformer.StatesInformer
	EventRecorder                    record.EventRecorder
	DisableUnsetCPUQuotaForCPUSetPod bool
}

type HookFn func(protocol.HooksProtocol) error

var globalStageHooks map[rmconfig.RuntimeHookType][]*Hook

func Register(stage rmconfig.RuntimeHookType, name, description string, hookFn HookFn) *Hook {
	_ = "STUB: not implemented"
	return nil
}

func generateNewHook(stage rmconfig.RuntimeHookType, name string) (*Hook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort hooks by name for a stable order

func getHooksByStage(stage rmconfig.RuntimeHookType) []*Hook { _ = "STUB: not implemented"; return nil }

func RunHooks(failPolicy rmconfig.FailurePolicyType, stage rmconfig.RuntimeHookType, protocol protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	globalStageHooks = map[rmconfig.RuntimeHookType][]*Hook{
		rmconfig.PreRunPodSandbox:            make([]*Hook, 0),
		rmconfig.PreCreateContainer:          make([]*Hook, 0),
		rmconfig.PreStartContainer:           make([]*Hook, 0),
		rmconfig.PostStartContainer:          make([]*Hook, 0),
		rmconfig.PostStopContainer:           make([]*Hook, 0),
		rmconfig.PostStopPodSandbox:          make([]*Hook, 0),
		rmconfig.PreUpdateContainerResources: make([]*Hook, 0),
		rmconfig.PreRemoveRunPodSandbox:      make([]*Hook, 0),
	}
}

func GetStages(disable map[string]struct{}) []rmconfig.RuntimeHookType {
	_ = "STUB: not implemented"
	return nil
}
