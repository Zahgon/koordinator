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
	"sync"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const ratioDiffEpsilon = 0.01

type Rule struct {
	lock sync.RWMutex

	enableCFSQuota        *bool    // default = true
	cpuNormalizationRatio *float64 // default = -1, -1 means disabled
}

func newRule() *Rule { _ = "STUB: not implemented"; return nil }

// GetCFSQuotaScaleRatio returns
// 1. if the cfs quota should be unlimited (-1 / max)
// 2. the scale ratio for the cfs quota when the quota is limited (-1 means do not scale)
func (r *Rule) GetCFSQuotaScaleRatio() (bool, float64) { _ = "STUB: not implemented"; return false, 0 }

func (r *Rule) UpdateCFSQuotaEnabled(enabled bool) bool { _ = "STUB: not implemented"; return false }

func (r *Rule) UpdateCPUNormalizationRatio(ratio float64) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *plugin) parseRuleForNodeSLO(mergedNodeSLOIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NOTE: If CPU Suppress Policy `CPUCfsQuotaPolicy` is enabled for batch pods, batch pods' cfs_quota should be unset
// since the cfs quota of `kubepods-besteffort` is required to be no less than the children's. Then the cpu usage
// of Batch is limited by pod-level cpu.shares and qos-level cfs_quota.

func (p *plugin) parseRuleForNodeMeta(nodeIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *plugin) ruleUpdateCbForNodeSLO(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// pod-level

// only need to change cfs quota

// container-level

// TODO set for sandbox container

func (p *plugin) ruleUpdateCbForNodeMeta(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: if the ratio becomes bigger, scale top down, otherwise, scale bottom up

// pod-level

// container-level

// ignore sandbox containers

// NOTE: Update cgroups by the level since some resources like cfs quota requires the upper level value is no less
//       than the lower.

func getCPUSuppressPolicy(nodeSLOSpec *slov1alpha1.NodeSLOSpec) (bool, slov1alpha1.CPUSuppressPolicy) {
	_ = "STUB: not implemented"
	return false, *new(slov1alpha1.CPUSuppressPolicy)
}
