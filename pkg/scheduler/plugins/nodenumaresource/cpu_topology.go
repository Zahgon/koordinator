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
	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

// TODO: migrate to latest cpu-manager code

// CPUTopology contains details of node cpu
type CPUTopology struct {
	NumCPUs    int        `json:"numCPUs"`
	NumCores   int        `json:"numCores"`
	NumNodes   int        `json:"numNodes"`
	NumSockets int        `json:"numSockets"`
	CPUDetails CPUDetails `json:"cpuDetails"`
}

type CPUTopologyBuilder struct {
	topologyTracker map[int] /*socket*/ map[int] /*node*/ map[int] /*core*/ struct{}
	topology        CPUTopology
}

func NewCPUTopologyBuilder() *CPUTopologyBuilder { _ = "STUB: not implemented"; return nil }

func (b *CPUTopologyBuilder) AddCPUInfo(socketID, nodeID, coreID, cpuID int) *CPUTopologyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *CPUTopologyBuilder) Result() *CPUTopology {
	_ = "STUB: not implemented"

	// IsValid checks if the topology is valid
	return nil
}

func (topo *CPUTopology) IsValid() bool { _ = "STUB: not implemented"; return false }

// CPUsPerCore returns the number of logical CPUs are associated with each core.
func (topo *CPUTopology) CPUsPerCore() int { _ = "STUB: not implemented"; return 0 }

// CPUsPerSocket returns the number of logical CPUs are associated with each socket.
func (topo *CPUTopology) CPUsPerSocket() int { _ = "STUB: not implemented"; return 0 }

// CPUsPerNode returns the number of logical CPUs are associated with each node.
func (topo *CPUTopology) CPUsPerNode() int { _ = "STUB: not implemented"; return 0 }

// CPUDetails is a map from logical CPU ID to CPUInfo.
type CPUDetails map[int]CPUInfo

// NewCPUDetails returns CPUDetails instance
func NewCPUDetails() CPUDetails {
	_ = "STUB: not implemented"
	return *

	// CPUInfo contains the NUMA, socket, and core IDs associated with a CPU.
	new(CPUDetails)
}

type CPUInfo struct {
	CPUID           int                                 `json:"cpuID"`
	CoreID          int                                 `json:"coreID"`
	NodeID          int                                 `json:"nodeID"`
	SocketID        int                                 `json:"socketID"`
	RefCount        int                                 `json:"refCount"`
	ExclusivePolicy schedulingconfig.CPUExclusivePolicy `json:"exclusivePolicy"`
}

// Clone clones the CPUDetails
func (d CPUDetails) Clone() CPUDetails { _ = "STUB: not implemented"; return *new(CPUDetails) }

// KeepOnly returns a new CPUDetails object with only the supplied cpus.
func (d CPUDetails) KeepOnly(cpus cpuset.CPUSet) CPUDetails {
	_ = "STUB: not implemented"
	return *new(CPUDetails)
}

// NUMANodes returns the NUMANode IDs associated with the CPUs in this CPUDetails.
func (d CPUDetails) NUMANodes() cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// NUMANodesInSockets returns the logical NUMANode IDs associated with the given socket IDs in this CPUDetails.
func (d CPUDetails) NUMANodesInSockets(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// Sockets returns the socket IDs associated with the CPUs in this CPUDetails.
func (d CPUDetails) Sockets() cpuset.CPUSet { _ = "STUB: not implemented"; return *new(cpuset.CPUSet) }

// CPUsInSockets returns logical CPU IDs associated with the given socket IDs in this CPUDetails.
func (d CPUDetails) CPUsInSockets(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// SocketsInNUMANodes returns the socket IDs associated with the given NUMANode IDs in this CPUDetails.
func (d CPUDetails) SocketsInNUMANodes(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// Cores returns the core IDs associated with the CPUs in this CPUDetails.
func (d CPUDetails) Cores() cpuset.CPUSet { _ = "STUB: not implemented"; return *new(cpuset.CPUSet) }

// CoresInNUMANodes returns the core IDs associated with the given NUMANode IDs in this CPUDetails.
func (d CPUDetails) CoresInNUMANodes(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// CoresInSockets returns the core IDs associated with the given socket IDs in this CPUDetails.
func (d CPUDetails) CoresInSockets(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// CPUs returns the logical CPU IDs in this CPUDetails.
func (d CPUDetails) CPUs() cpuset.CPUSet { _ = "STUB: not implemented"; return *new(cpuset.CPUSet) }

// CPUsInNUMANodes returns the logical CPU IDs associated with the given NUMANode IDs in this CPUDetails.
func (d CPUDetails) CPUsInNUMANodes(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}

// CPUsInCores returns the logical CPU IDs associated with the given core IDs in this CPUDetails.
func (d CPUDetails) CPUsInCores(ids ...int) cpuset.CPUSet {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet)
}
