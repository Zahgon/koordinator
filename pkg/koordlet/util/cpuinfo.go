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

package util

import (
	"time"

	"github.com/koordinator-sh/koordinator/apis/extension"
)

const cpuCmdTimeout = 5 * time.Second // maybe run slowly on some platforms

// ProcessorInfo describes the processor topology information of a single logic cpu, including the core, socket and numa
// node it belongs to
type ProcessorInfo struct {
	// logic CPU/ processor ID
	CPUID int32 `json:"cpu"`
	// physical CPU core ID
	CoreID int32 `json:"core"`
	// cpu socket ID
	SocketID int32 `json:"socket"`
	// numa node ID
	NodeID int32 `json:"node"`
	// L1 L2 cache ID
	L1dl1il2 string `json:"l1dl1il2"`
	// L3 cache ID
	L3 int32 `json:"l3"`
	// online
	Online string `json:"online"`
}

// CPUTotalInfo describes the total number infos of the local cpu, e.g. the number of cores, the number of numa nodes
type CPUTotalInfo struct {
	NumberCPUs  int32                     `json:"numberCPUs"`
	CoreToCPU   map[int32][]ProcessorInfo `json:"coreToCPU"`
	NodeToCPU   map[int32][]ProcessorInfo `json:"nodeToCPU"`
	SocketToCPU map[int32][]ProcessorInfo `json:"socketToCPU"`
	L3ToCPU     map[int32][]ProcessorInfo `json:"l3ToCPU"`
}

// LocalCPUInfo contains the cpu information collected from the node
type LocalCPUInfo struct {
	// BasicInfo describe the cpu features and their status
	BasicInfo extension.CPUBasicInfo `json:"basicInfo,omitempty"`
	// ProcessorInfos contains topology information of all available CPUs
	ProcessorInfos []ProcessorInfo `json:"processorInfos,omitempty"`
	// TotalInfo stores the numbers of cpu processors, cores, sockets and nodes
	TotalInfo CPUTotalInfo `json:"totalInfo,omitempty"`
}

// getCPUModel gets the Model name of the CPU.
func getCPUModel() (string, error) { _ = "STUB: not implemented"; return "", nil }

// getHyperThreadEnabled returns whether the cpu is HT-enabled or not
// NOTE: currently only support intel cpu, otherwise it can always return false
func getHyperThreadEnabled() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func getCPUTurboEnabled() (bool, error) {
	_ = "STUB: not implemented"
	// TODO: In the current version, only intel cpu is collected turbo status. The other vendors' interfaces are not
	//
	//	supported yet. We may check the frequency in the future.
	return false, nil
}

func getCPUBasicInfo() (*extension.CPUBasicInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func lsCPU(option string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getProcessorInfos(lsCPUStr string) ([]ProcessorInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sorted by cpu topology
// NOTE: in some cases, max(cpuId[...]) can be not equal to len(processors)

func calculateCPUTotalInfo(processorInfos []ProcessorInfo) *CPUTotalInfo {
	_ = "STUB: not implemented"
	return nil
}

// GetLocalCPUInfo returns the local cpu info for cpuset allocation, NUMA-aware scheduling
func GetLocalCPUInfo() (*LocalCPUInfo, error) { _ = "STUB: not implemented"; return nil, nil }
