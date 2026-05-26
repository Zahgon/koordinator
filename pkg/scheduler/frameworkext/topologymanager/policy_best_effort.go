/*
Copyright 2022 The Koordinator Authors.
Copyright 2019 The Kubernetes Authors.

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

package topologymanager

import (
	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

type bestEffortPolicy struct {
	//List of NUMA Nodes available on the underlying machine
	numaNodes []int
}

var _ Policy = &bestEffortPolicy{}

// PolicyBestEffort policy name.
const PolicyBestEffort string = "best-effort"

// NewBestEffortPolicy returns best-effort policy.
func NewBestEffortPolicy(numaNodes []int) Policy { _ = "STUB: not implemented"; return *new(Policy) }

func (p *bestEffortPolicy) Name() string { _ = "STUB: not implemented"; return "" }

func (p *bestEffortPolicy) canAdmitPodResult(hint *NUMATopologyHint) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *bestEffortPolicy) Merge(providersHints []map[string][]NUMATopologyHint, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) (NUMATopologyHint, bool, []string) {
	_ = "STUB: not implemented"
	return *new(NUMATopologyHint), false, nil
}

// 如果 bestHint 不是一个所有资源都可分的 numa affinity，则应该返回 bestHint 为亲和所有 NUMANode，即放弃任何 NUMA 倾向
