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

package system

type CPUStatV2Raw struct {
	UsageUsec  int64
	UserUsec   int64
	SystemUSec int64

	NrPeriods     int64
	NrThrottled   int64
	ThrottledUSec int64
}

func initCgroupsVersion() { _ = "STUB: not implemented"; return }

func ParseCPUCFSQuotaV2(content string) (int64, error) {
	_ = "STUB: not implemented"
	// content: "max 100000", "100000 100000"; the first field indicates cfs quota
	return 0, nil
}

// "max" means unlimited

func ParseCPUCFSPeriodV2(content string) (int64, error) {
	_ = "STUB: not implemented"
	// content: "max 100000", "100000 100000"; the second field indicates cfs period
	return 0, nil
}

func ParseCPUAcctStatRawV2(content string) (*CPUStatV2Raw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseCPUAcctUsageV2(content string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// return usage in nanosecond (compatible to v1)
// assert no overflow

func ParseCPUStatRawV2(content string) (*CPUStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assert no overflow

func ParseMemoryStatRawV2(content string) (*MemoryStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseMemoryNumaStatV2(content string) ([]NumaMemoryPages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertCPUWeightToShares converts the value of `cpu.weight` (cgroups-v2) into the value of `cpu.shares` (cgroups-v1)
func ConvertCPUWeightToShares(v int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Use the inverse conversion of the kubelet.
// https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2254-cgroup-v2
// Map weights [1, 10000] to shares [2, 262144]:
// shares = (weights - 1) * 262142 / 9999 + 2

func ConvertCPUSharesToWeight(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// the valid value must be an integer
// Use the same conversion of the kubelet.
// https://github.com/kubernetes/enhancements/tree/master/keps/sig-node/2254-cgroup-v2
// Map shares [2, 262144] to weights [1, 10000]:
// weights = 1 + ((shares - 2) * 9999) / 262142
