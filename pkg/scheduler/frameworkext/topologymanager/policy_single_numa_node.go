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

type singleNumaNodePolicy struct {
	//List of NUMA Nodes available on the underlying machine
	numaNodes []int
}

var _ Policy = &singleNumaNodePolicy{}

// PolicySingleNumaNode policy name.
const PolicySingleNumaNode string = "single-numa-node"

// NewSingleNumaNodePolicy returns single-numa-node policy.
func NewSingleNumaNodePolicy(numaNodes []int) Policy {
	_ = "STUB: not implemented"
	return *new(Policy)
}

func (p *singleNumaNodePolicy) Name() string { _ = "STUB: not implemented"; return "" }

func (p *singleNumaNodePolicy) canAdmitPodResult(hint *NUMATopologyHint) bool {
	_ = "STUB: not implemented"
	return false

	// Return hints that have valid bitmasks with exactly one bit set.
}

func filterSingleNumaHints(allResourcesHints [][]NUMATopologyHint) [][]NUMATopologyHint {
	_ = "STUB: not implemented"
	return nil
}

func (p *singleNumaNodePolicy) Merge(providersHints []map[string][]NUMATopologyHint, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) (NUMATopologyHint, bool, []string) {
	_ = "STUB: not implemented"
	return *new(NUMATopologyHint), false, nil
}

// Filter to only include don't care and hints with a single NUMA node.
