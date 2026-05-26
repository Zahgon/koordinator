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

const (
	DefaultCPUCFSPeriod int64 = 100000
	CPUShareKubeBEValue int64 = 2
	CPUShareUnitValue   int64 = 1024
	// MemoryLimitUnlimitedValue denotes the unlimited value of cgroups-v1 memory.limit_in_bytes.
	// It derives from linux PAGE_COUNTER_MAX and may be different according to the PAGE_SIZE (here we suppose `4k`).
	// https://github.com/torvalds/linux/blob/ea4424be16887a37735d6550cfd0611528dbe5d9/mm/memcontrol.c#L5337
	MemoryLimitUnlimitedValue int64 = 0x7FFFFFFFFFFFF000 // 9223372036854771712 < math.MaxInt64

	// CgroupMaxSymbolStr only appears in cgroups-v2 files, we consider the value as MaxInt64
	CgroupMaxSymbolStr string = "max"
	// CgroupUnlimitedSymbolStr indicates value is unlimited. It appears in cfs_quota which regarded as max in comparison.
	CgroupUnlimitedSymbolStr string = "-1"
	// CgroupMaxValueStr math.MaxInt64; writing `memory.high` with this do the same as set as "max"
	CgroupMaxValueStr string = "9223372036854775807"
)

const ErrCgroupDir = "cgroup path or file not exist"

type CPUStatRaw struct {
	NrPeriods            int64
	NrThrottled          int64
	ThrottledNanoSeconds int64
}

type MemoryStatRaw struct {
	Cache        int64
	RSS          int64
	InactiveFile int64
	ActiveFile   int64
	InactiveAnon int64
	ActiveAnon   int64
	Unevictable  int64
	// add more fields
}

type NumaMemoryPages struct {
	NumaId   int
	PagesNum uint64
}

func (m *MemoryStatRaw) Usage() int64 {
	_ = "STUB: not implemented"
	// memory.stat usage: total_inactive_anon + total_active_anon + total_unevictable
	return 0
}

func (m *MemoryStatRaw) UsageWithPageCache() int64 {
	_ = "STUB: not implemented"
	// memory.stat usage: total_inactive_anon + total_active_anon + total_unevictable + total_activefile + total_inactivefile
	return 0
}

// GetCgroupFilePath gets the full path of the given cgroup dir and resource.
// @cgroupTaskDir kubepods.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// @return /sys/fs/cgroup/cpu/kubepods.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/cpu.shares
func GetCgroupFilePath(cgroupTaskDir string, r Resource) string {
	_ = "STUB: not implemented"
	return ""
}

func ParseCPUStatRaw(content string) (*CPUStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseMemoryStatRaw(content string) (*MemoryStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseMemoryNumaStat(content string) ([]NumaMemoryPages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseCgroupProcs parses the content in cgroup.procs.
// pattern: `7742\n10971\n11049\n11051...`
// TODO: refactor with readCgroupAndParseInt32Slice via Generics.
func ParseCgroupProcs(content string) ([]uint32, error) { _ = "STUB: not implemented"; return nil, nil }

func CalcCPUThrottledRatio(curPoint, prePoint *CPUStatRaw) float64 {
	_ = "STUB: not implemented"
	return 0
}

func GetRootCgroupSubfsDir(subfs string) string { _ = "STUB: not implemented"; return "" }

func MilliCPUToShares(milliCPURequest int64) int64 { _ = "STUB: not implemented"; return 0 }

func MilliCPUToQuota(milliCPULimit int64) int64 { _ = "STUB: not implemented"; return 0 }

// TBD: assert base cfs period not changed
// unlimited

// cfs_quota_us should be no less than 1000
