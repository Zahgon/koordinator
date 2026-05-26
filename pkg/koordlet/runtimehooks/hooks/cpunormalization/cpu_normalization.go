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

package cpunormalization

import (
	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
)

const (
	name        = "CPUNormalization"
	description = "adjust cpu cgroups value for cpu normalized LS pod"
)

var podQOSConditions = []string{string(extension.QoSLS), string(extension.QoSNone)}

type Plugin struct {
	rule     *Rule
	executor resourceexecutor.ResourceUpdateExecutor
}

var singleton *Plugin

func Object() *Plugin { _ = "STUB: not implemented"; return nil }

func newPlugin() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

func (p *Plugin) AdjustPodCFSQuota(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// currently only reconciler mode provides Resources in ctx

func (p *Plugin) AdjustContainerCFSQuota(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// currently only reconciler mode provides Resources in ctx

func (p *Plugin) adjustPodCFSQuota(podCtx *protocol.PodContext) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to scale when cgroup is unset

func (p *Plugin) adjustContainerCFSQuota(containerCtx *protocol.ContainerContext) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to scale when cgroup is unset

func isPodCPUShare(labels map[string]string, annotations map[string]string) bool {
	_ = "STUB: not implemented"
	// considered None
	return false
}

// consider as LSR if pod is qos=None and has cpuset
