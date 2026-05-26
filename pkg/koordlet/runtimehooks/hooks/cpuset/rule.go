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
	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type cpusetRule struct {
	kubeletPolicy   extension.KubeletCPUManagerPolicy
	sharePools      []extension.CPUSharedPool
	beSharePools    []extension.CPUSharedPool
	systemQOSCPUSet string
	// TODO: support per-node disable
}

func (r *cpusetRule) getContainerCPUSet(containerReq *protocol.ContainerRequest) (*string, error) {
	_ = "STUB: not implemented"
	// pod specifies QoS=BE and share pool id in annotations, use part be cpu share pool if BECPUManager enabled
	// pod specifies share pool id in annotations, use part cpu share pool
	// pod specifies QoS=SYSTEM in labels, use system qos resource if rule exist
	// pod specifies QoS=LS in labels, use all share pool
	// besteffort pod(including QoS=BE) will be managed by cpu suppress policy, inject empty string
	// guaranteed/bustable pod without QoS label, if kubelet use none policy, use all share pool, and if kubelet use
	// static policy, do nothing
	return nil, nil
}

// check if numa-aware

// check if cpu resource is allocated in numa-level since there can be numa allocation without cpu

// BE pods which have specified cpu share pool

// LS pods which have specified cpu share pool

// SYSTEM QoS cpuset
// TBD: support numa-aware

// LS pods use all share pool

// besteffort pods including QoS=BE, clear cpuset of BE container to avoid conflict with kubelet static policy,
// which will pass cpuset in StartContainerRequest of CRI
// TODO remove this in the future since cpu suppress will keep besteffort dir as all cpuset

// none policy

func (r *cpusetRule) getHostAppCpuset(hostAppReq *protocol.HostAppRequest) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *cpusetPlugin) parseRule(nodeTopoIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check cpuset format

// sharePool CPUSet ID info may be expired and needs to be reset before record new metrics.

func (p *cpusetPlugin) ruleUpdateCb(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cpusetPlugin) getRule() *cpusetRule { _ = "STUB: not implemented"; return nil }

func (p *cpusetPlugin) updateRule(newRule *cpusetRule) bool {
	_ = "STUB: not implemented"
	return false
}
