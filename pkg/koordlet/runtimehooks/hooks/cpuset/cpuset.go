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

package cpuset

import (
	"sync"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
)

const (
	name        = "CPUSetAllocator"
	description = "set cpuset value by pod allocation"
)

type cpusetPlugin struct {
	rule                 *cpusetRule
	ruleRWMutex          sync.RWMutex
	executor             resourceexecutor.ResourceUpdateExecutor
	disableUnsetCPUQuota bool
	// record sharePool, beSharePool cpu ID info.
	recordPerSharePoolCPUInfo bool
}

var (
	cpusetPodQOSConditions   = []string{string(apiext.QoSLSE), string(apiext.QoSLSR)}
	cpusharePodQOSConditions = []string{string(apiext.QoSLS), string(apiext.QoSBE), string(apiext.QoSSystem), string(apiext.QoSNone)}
)

func (p *cpusetPlugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

var singleton *cpusetPlugin

func Object() *cpusetPlugin { _ = "STUB: not implemented"; return nil }

func (p *cpusetPlugin) SetContainerCPUSetAndUnsetCFS(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	// set container-level cpuset.cpus
	return nil
}

// unset container-level cpu.cfs_quota_us if needed

func (p *cpusetPlugin) SetContainerCPUSet(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// cpuset from pod annotation (LSE, LSR)

// cpuset from rule according to pod QoS

func (p *cpusetPlugin) SetHostAppCPUSet(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cpusetPlugin) UnsetPodCPUQuota(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// cpuset from pod annotation (LSE, LSR)
// NOTE: unset cfs quota for cpuset pods to avoid unexpected throttles.
// https://github.com/koordinator-sh/koordinator/issues/489

// do nothing for cpushare pod

func (p *cpusetPlugin) UnsetContainerCPUQuota(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// cpuset from pod annotation (LSE, LSR)
// NOTE: unset cfs quota for cpuset pods to avoid unexpected throttles.
// https://github.com/koordinator-sh/koordinator/issues/489

// do nothing for cpushare pod
