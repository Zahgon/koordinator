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

type restrictedPolicy struct {
	bestEffortPolicy
}

var _ Policy = &restrictedPolicy{}

// PolicyRestricted policy name.
const PolicyRestricted string = "restricted"

// NewRestrictedPolicy returns restricted policy.
func NewRestrictedPolicy(numaNodes []int) Policy { _ = "STUB: not implemented"; return *new(Policy) }

func (p *restrictedPolicy) Name() string { _ = "STUB: not implemented"; return "" }

func (p *restrictedPolicy) canAdmitPodResult(hint *NUMATopologyHint) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *restrictedPolicy) Merge(providersHints []map[string][]NUMATopologyHint, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) (NUMATopologyHint, bool, []string) {
	_ = "STUB: not implemented"
	return *new(NUMATopologyHint), false, nil
}
