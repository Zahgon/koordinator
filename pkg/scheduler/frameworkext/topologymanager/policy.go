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
	"github.com/koordinator-sh/koordinator/pkg/util/bitmask"
)

type Policy interface {
	// Name returns Policy Name
	Name() string
	// Merge returns a merged NUMATopologyHint based on input from hint providers
	Merge(providersHints []map[string][]NUMATopologyHint, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) (NUMATopologyHint, bool, []string)
}

// NUMATopologyHint is a struct containing the NUMANodeAffinity for a Container
type NUMATopologyHint struct {
	NUMANodeAffinity bitmask.BitMask
	// Unsatisfied is set to true when the NUMANodeAffinity is insufficient to satisfy the resource request.
	Unsatisfied bool
	// Preferred is set to true when the NUMANodeAffinity encodes a preferred
	// allocation for the Pod. It is set to false otherwise.
	Preferred bool
	// Score is the weight of this hint. For the same Affinity,
	// the one with higher weight will be used first.
	Score int64
}

// IsEqual checks if NUMATopologyHint are equal
func (th *NUMATopologyHint) IsEqual(topologyHint NUMATopologyHint) bool {
	_ = "STUB: not implemented"
	return false
}

// LessThan checks if NUMATopologyHint `a` is less than NUMATopologyHint `b`
// this means that either `a` is a preferred hint and `b` is not
// or `a` NUMANodeAffinity attribute is narrower than `b` NUMANodeAffinity attribute.
func (th *NUMATopologyHint) LessThan(other NUMATopologyHint) bool {
	_ = "STUB: not implemented"
	return false
}

// Check if the affinity match the exclusive policy, return true if match or false otherwise.
func checkExclusivePolicy(affinity NUMATopologyHint, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) bool {
	_ = "STUB: not implemented"
	// check bestHint again if default hint is the best
	return false
}

// we should make sure no numa is in single state

// Merge a TopologyHints permutation to a single hint by performing a bitwise-AND
// of their affinity masks. The hint shall be preferred if all hits in the permutation
// are preferred.
func mergePermutation(numaNodes []int, permutation []NUMATopologyHint) NUMATopologyHint {
	_ = "STUB: not implemented"
	// Get the NUMANodeAffinity from each hint in the permutation and see if any
	// of them encode unpreferred allocations.
	return *new(NUMATopologyHint)
}

// Only consider hints that have an actual NUMANodeAffinity set.

// Only mark preferred if all affinities are equal.

// Only mark preferred if all affinities are preferred.

// Merge the affinities using a bitwise-and operation.

// Build a mergedHint from the merged affinity mask, indicating if an
// preferred allocation was used to generate the affinity mask or not.

func filterProvidersHints(providersHints []map[string][]NUMATopologyHint) ([][]NUMATopologyHint, []string, []string) {
	_ = "STUB: not implemented"
	// Loop through all hint providers and save an accumulated list of the
	// hints returned by each hint provider. If no hints are provided, assume
	// that provider has no preference for topology-aware allocation.
	return nil, nil, nil
}

// If hints is nil, insert a single, preferred any-numa hint into allProviderHints.

// Otherwise, accumulate the hints for each resource type into allProviderHints.

func getSummaryForResource(resource string, hints []NUMATopologyHint) string {
	_ = "STUB: not implemented"
	return ""
}

func mergeFilteredHints(numaNodes []int, filteredHints [][]NUMATopologyHint, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) NUMATopologyHint {
	_ = "STUB: not implemented"
	// Set the default affinity as an any-numa affinity containing the list
	// of NUMA Nodes available on this machine.
	return *new(NUMATopologyHint)
}

// Set the bestHint to return from this function as {nil false}.
// This will only be returned if no better hint can be found when
// merging hints from each hint provider.

// Get the NUMANodeAffinity from each hint in the permutation and see if any
// of them encode unpreferred allocations.

// Only consider mergedHints that result in a NUMANodeAffinity > 0 to
// replace the current bestHint.

// If the current bestHint is non-preferred and the new mergedHint is
// preferred, always choose the preferred hint over the non-preferred one.

// If the current bestHint is preferred and the new mergedHint is
// non-preferred, never update bestHint, regardless of mergedHint's
// narowness.

// If mergedHint and bestHint has the same preference, only consider
// mergedHints that have a narrower NUMANodeAffinity than the
// NUMANodeAffinity in the current bestHint.

// In all other cases, update bestHint to the current mergedHint

// Iterate over all permutations of hints in 'allProviderHints [][]TopologyHint'.
//
// This procedure is implemented as a recursive function over the set of hints
// in 'allproviderHints[i]'. It applies the function 'callback' to each
// permutation as it is found. It is the equivalent of:
//
// for i := 0; i < len(providerHints[0]); i++
//
//	for j := 0; j < len(providerHints[1]); j++
//	    for k := 0; k < len(providerHints[2]); k++
//	        ...
//	        for z := 0; z < len(providerHints[-1]); z++
//	            permutation := []TopologyHint{
//	                providerHints[0][i],
//	                providerHints[1][j],
//	                providerHints[2][k],
//	                ...
//	                providerHints[-1][z]
//	            }
//	            callback(permutation)
func iterateAllProviderTopologyHints(allProviderHints [][]NUMATopologyHint, callback func([]NUMATopologyHint)) {
	_ = "STUB: not implemented"
	// Internal helper function to accumulate the permutation before calling the callback.
	return
}

// Base case: we have looped through all providers and have a full permutation.

// Loop through all hints for provider 'i', and recurse to build the
// the permutation of this hint with all hints from providers 'i++'.
