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

package nodenumaresource

import (
	"k8s.io/apimachinery/pkg/util/sets"

	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

func takePreferredCPUs(
	topology *CPUTopology,
	maxRefCount int,
	availableCPUs cpuset.CPUSet,
	preferredCPUs cpuset.CPUSet,
	allocatedCPUs CPUDetails,
	numCPUsNeeded int,
	cpuBindPolicy schedulingconfig.CPUBindPolicy,
	cpuExclusivePolicy schedulingconfig.CPUExclusivePolicy,
	numaAllocatedStrategy schedulingconfig.NUMAAllocateStrategy,
) (cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}

func takeCPUs(
	topology *CPUTopology,
	maxRefCount int,
	availableCPUs cpuset.CPUSet,
	allocatedCPUs CPUDetails,
	numCPUsNeeded int,
	cpuBindPolicy schedulingconfig.CPUBindPolicy,
	cpuExclusivePolicy schedulingconfig.CPUExclusivePolicy,
	numaAllocatedStrategy schedulingconfig.NUMAAllocateStrategy,
) (cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}

// According to the NUMA allocation strategy,
// select the NUMA Node with the most remaining amount or the least amount remaining
// and the total amount of available CPUs in the NUMA Node is greater than or equal to the number of CPUs needed

// According to the NUMA allocation strategy,
// select the NUMA Socket with the most remaining amount or the least amount remaining
// and the total amount of available CPUs in the NUMA Socket is greater than or equal to the number of CPUs needed

// There are some scenarios at this time,
// - the amount of CPUs needed exceeds the total amount of a NUMA Socket,
// - or each NUMA Node/Socket is allocated a part and the remaining number of CPUs does not meet the needed.
// For better performance, allocate from the NUMA Socket with the most remaining physical cores as much as possible
// and no longer follow the NUMA Allocate Strategy.

// There are still some unsatisfied physical core requests,
// which need to be allocated from the NUMA Socket with the fewest remaining physical cores as much as possible

// The following scenarios need to be allocated through the SpreadByPCPUs policy
// - The application requires CPU allocation according to the SpreadByPCPUs policy
// - Not enough physical cores to satisfy FullPCPUs policy
// - It is required to be allocated in the FullPCPUs policy, but the number of CPUs requested is
//   not enough to monopolize a complete physical core

// Try to allocate CPUs in the same NUMA Node according NUMA Allocate Strategy,
// even if the SpreadByPCPUs policy cannot be satisfied

// Try to allocate CPUs in the same NUMA Socket according NUMA Allocate Strategy,
// even if the SpreadByPCPUs policy cannot be satisfied

// Try to allocate CPUs on the NUMA Node/Socket which the allocated CPU belongs

type cpuAccumulator struct {
	topology             *CPUTopology
	maxRefCount          int
	allocatableCPUs      CPUDetails
	numCPUsNeeded        int
	exclusive            bool
	exclusiveInCores     sets.Int
	exclusiveInNUMANodes sets.Int
	exclusivePolicy      schedulingconfig.CPUExclusivePolicy
	numaAllocateStrategy schedulingconfig.NUMAAllocateStrategy
	result               cpuset.CPUSet
}

func newCPUAccumulator(
	topology *CPUTopology,
	maxRefCount int,
	availableCPUs cpuset.CPUSet,
	allocatedCPUs CPUDetails,
	numCPUsNeeded int,
	exclusivePolicy schedulingconfig.CPUExclusivePolicy,
	numaAllocateStrategy schedulingconfig.NUMAAllocateStrategy,
) *cpuAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func (a *cpuAccumulator) take(cpus ...int) { _ = "STUB: not implemented"; return }

func (a *cpuAccumulator) needs(n int) bool { _ = "STUB: not implemented"; return false }

func (a *cpuAccumulator) isSatisfied() bool { _ = "STUB: not implemented"; return false }

func (a *cpuAccumulator) isFailed() bool { _ = "STUB: not implemented"; return false }

func (a *cpuAccumulator) isCPUExclusivePCPULevel(cpuInfo *CPUInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *cpuAccumulator) isCPUExclusiveNUMANodeLevel(cpuInfo *CPUInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *cpuAccumulator) extractCPU(cpus []int) []int { _ = "STUB: not implemented"; return nil }

func (a *cpuAccumulator) sortCores(details CPUDetails, cores []int, cpusInCores map[int][]int) {
	_ = "STUB: not implemented"
	return
}

// freeCoresInNode returns the logical cpus of the free cores in nodes that sorted
func (a *cpuAccumulator) freeCoresInNode(filterFullFreeCore bool, filterExclusive bool) [][]int {
	_ = "STUB: not implemented"
	return nil
}

// each cpu's socketId and nodeId in same node are same

// Compute the number of available CPUs available on the same node as each core.

// Compute the number of available CPUs available on the same socket as each core.

// freeCoresInSocket returns the logical cpus of the free cores in sockets that sorted
func (a *cpuAccumulator) freeCoresInSocket(filterFullFreeCore bool) [][]int {
	_ = "STUB: not implemented"
	return nil
}

// freeCPUsInNode returns free logical cpus in nodes that sorted in ascending order.
func (a *cpuAccumulator) freeCPUsInNode(filterExclusive bool) [][]int {
	_ = "STUB: not implemented"
	return nil
}

// Compute the number of available CPUs available on the same node as each core.

// Compute the number of available CPUs available on the same socket as each core.

// freeCPUsInSocket returns free logical cpus in sockets that sorted in ascending order.
func (a *cpuAccumulator) freeCPUsInSocket(filterExclusive bool) [][]int {
	_ = "STUB: not implemented"
	return nil
}

// Returns CPU IDs as a slice sorted in ascending order by:
// - socket affinity with result
// - number of CPUs available on the same socket
// - number of CPUs available on the same core
// - socket ID
// - node ID
// - reference count if maxRefCount > 1
// - core ID
func (a *cpuAccumulator) freeCPUs(filterExclusive bool) []int {
	_ = "STUB: not implemented"
	return nil
}

// Compute the number of CPUs in the result reside on the same socket as each core.

// Compute the number of available CPUs available on the same socket as each core.

// Compute the number of available CPUs available on the same node as each core.

// Compute the number of available CPUs on each core.

// For each core, append sorted CPU IDs to result.

func getCoreRefCount(details CPUDetails, core int) int { _ = "STUB: not implemented"; return 0 }

func (a *cpuAccumulator) sortCPUsByRefCount(cpus []int) { _ = "STUB: not implemented"; return }

func (a *cpuAccumulator) spreadCPUs(cpus []int) []int { _ = "STUB: not implemented"; return nil }
