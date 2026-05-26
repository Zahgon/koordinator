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

package resourceexecutor

import (
	"errors"

	sysutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

var ErrResourceNotRegistered = errors.New("resource not registered")

type CgroupReader interface {
	ReadCPUQuota(parentDir string) (int64, error)
	ReadCPUPeriod(parentDir string) (int64, error)
	ReadCPUShares(parentDir string) (int64, error)
	ReadCPUSet(parentDir string) (*cpuset.CPUSet, error)
	ReadCPUAcctUsage(parentDir string) (uint64, error)
	ReadCPUStat(parentDir string) (*sysutil.CPUStatRaw, error)
	ReadMemoryUsage(parentDir string) (uint64, error)
	ReadMemoryLimit(parentDir string) (int64, error)
	ReadMemoryStat(parentDir string) (*sysutil.MemoryStatRaw, error)
	ReadMemoryNumaStat(parentDir string) ([]sysutil.NumaMemoryPages, error)
	ReadCPUTasks(parentDir string) ([]int32, error)
	ReadCPUProcs(parentDir string) ([]uint32, error)
	ReadPSI(parentDir string) (*sysutil.PSIByResource, error)
	ReadMemoryColdPageUsage(parentDir string) (uint64, error)
	ReadNetClsId(parentDir string) (uint32, error)
}

var _ CgroupReader = &CgroupV1Reader{}

type CgroupV1Reader struct{}

func (r *CgroupV1Reader) ReadCPUQuota(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV1Reader) ReadCPUPeriod(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV1Reader) ReadCPUShares(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV1Reader) ReadCPUSet(parentDir string) (*cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CgroupV1Reader) ReadCPUAcctUsage(parentDir string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV1Reader) ReadCPUStat(parentDir string) (*sysutil.CPUStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: "nr_periods 0\nnr_throttled 0\nthrottled_time 0\n..."

func (r *CgroupV1Reader) ReadMemoryUsage(parentDir string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV1Reader) ReadMemoryLimit(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// `memory.limit_in_bytes=9223372036854771712` means memory is unlimited, consider as value -1

func (r *CgroupV1Reader) ReadMemoryStat(parentDir string) (*sysutil.MemoryStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: `...total_inactive_anon $total_inactive_anon\ntotal_active_anon $total_active_anon\n
//           total_inactive_file $total_inactive_file\ntotal_active_file $total_active_file\n
//           total_unevictable $total_unevictable\n`

func (r *CgroupV1Reader) ReadMemoryNumaStat(parentDir string) ([]sysutil.NumaMemoryPages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// `total=42227 N0=42184 N1=...\nfile=40094 N0=40126 N1=...\nanon=2133 N0=2058 N1=...\nunevictable=0 N0=0\n...`
// the unit is page

func (r *CgroupV1Reader) ReadMemoryColdPageUsage(parentDir string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV1Reader) ReadCPUTasks(parentDir string) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: `7742\n10971\n11049\n11051...`

func (r *CgroupV1Reader) ReadCPUProcs(parentDir string) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: `7742\n10971\n11049\n11051...`

func (r *CgroupV1Reader) ReadPSI(parentDir string) (*sysutil.PSIByResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CgroupV1Reader) ReadNetClsId(parentDir string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ CgroupReader = &CgroupV2Reader{}

type CgroupV2Reader struct{}

func (r *CgroupV2Reader) ReadCPUQuota(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// content: "max 100000", "100000 100000"

func (r *CgroupV2Reader) ReadCPUPeriod(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// content: "max 100000", "100000 100000"

func (r *CgroupV2Reader) ReadCPUShares(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// convert cpu.weight value into cpu.shares value

func (r *CgroupV2Reader) ReadCPUSet(parentDir string) (*cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	// use `cpuset.cpus.effective` for read cpuset on cgroups-v2
	// https://docs.kernel.org/admin-guide/cgroup-v2.html#cpuset-interface-files
	return nil, nil
}

func (r *CgroupV2Reader) ReadCPUAcctUsage(parentDir string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// content: "usage_usec 1000000\nuser_usec 800000\nsystem_usec 200000\n..."

func (r *CgroupV2Reader) ReadCPUStat(parentDir string) (*sysutil.CPUStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: "...\nnr_periods 0\nnr_throttled 0\nthrottled_usec 0\n..."

func (r *CgroupV2Reader) ReadMemoryUsage(parentDir string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV2Reader) ReadMemoryLimit(parentDir string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *CgroupV2Reader) ReadMemoryStat(parentDir string) (*sysutil.MemoryStatRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: `anon 0\nfile 0\nkernel_stack 0\n...inactive_anon 0\nactive_anon 0\n...`

func (r *CgroupV2Reader) ReadMemoryNumaStat(parentDir string) ([]sysutil.NumaMemoryPages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// `anon N0=193236992 N1=...\nfile N0=1367764992 N1=...`
// the unit is byte, 2Kbyte -> a page

func (r *CgroupV2Reader) ReadMemoryColdPageUsage(parentDir string) (uint64, error) {
	_ = "STUB: not implemented"
	// cgroup v2 has not implemented yet
	return 0, nil
}

func (r *CgroupV2Reader) ReadCPUTasks(parentDir string) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: `7742\n10971\n11049\n11051...`

func (r *CgroupV2Reader) ReadCPUProcs(parentDir string) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: `7742\n10971\n11049\n11051...`

func (r *CgroupV2Reader) ReadPSI(parentDir string) (*sysutil.PSIByResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CgroupV2Reader) ReadNetClsId(parentDir string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func NewCgroupReader() CgroupReader { _ = "STUB: not implemented"; return *new(CgroupReader) }
