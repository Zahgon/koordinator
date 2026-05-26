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

package batchresource

import (
	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
)

const (
	name        = "BatchResource"
	description = "set fundamental cgroups value for batch pod"

	ruleNameForNodeSLO  = name + " (nodeSLO)"
	ruleNameForNodeMeta = name + " (nodeMeta)"
)

type plugin struct {
	rule     *Rule
	executor resourceexecutor.ResourceUpdateExecutor
}

var podQOSConditions = []string{string(apiext.QoSBE)}

func (p *plugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

var singleton *plugin

func Object() *plugin { _ = "STUB: not implemented"; return nil }

func newPlugin() *plugin { _ = "STUB: not implemented"; return nil }

func (p *plugin) SetPodResources(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) SetPodCPUShares(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// if the extendedResourceSpec is empty, do nothing and keep the original cgroup configs

// TODO: count init container and pod overhead

func (p *plugin) SetPodCFSQuota(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// if the extendedResourceSpec is empty, do nothing and keep the original cgroup configs

// if cfs quota is disabled, set as -1

// TODO: count init container and pod overhead

// pod unlimited once a container is unlimited

// no support ratio in (0, 1) yet

func (p *plugin) SetPodMemoryLimit(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// if the extendedResourceSpec is empty, do nothing and keep the original cgroup configs

// TODO: count init container and pod overhead

// pod unlimited once a container is unlimited

func (p *plugin) SetContainerResources(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *plugin) SetContainerCPUShares(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// if the extendedResourceSpec is empty, do nothing and keep the original cgroup configs

func (p *plugin) SetContainerCFSQuota(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// if the extendedResourceSpec is empty, do nothing and keep the original cgroup configs

// if cfs quota is disabled, set as -1

// no support ratio in (0, 1) yet

func (p *plugin) SetContainerMemoryLimit(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// if the extendedResourceSpec is empty, do nothing and keep the original cgroup configs

func isPodQoSBEByAttr(labels map[string]string, annotations map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}
