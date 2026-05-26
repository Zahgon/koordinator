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

import (
	"go.uber.org/atomic"
)

var (
	isSupportColdMemory              *atomic.Bool = atomic.NewBool(false)
	isStartColdMemory                *atomic.Bool = atomic.NewBool(false)
	kidledColdBoundary                            = defaultKidledColdBoundary
	defaultKidledScanPeriodInseconds uint32       = 5
	defaultKidledUseHierarchy        uint8        = 1
	defaultKidledColdBoundary        int          = 5
)

// the unit of Csei, Dsei, Cfei ... is byte
// the detailed description can be seen in https://github.com/alibaba/cloud-kernel/blob/linux-next/Documentation/vm/kidled.rst
type ColdPageInfoByKidled struct {
	Version             string   `json:"version"`
	PageScans           uint64   `json:"page_scans"`
	SlabScans           uint64   `json:"slab_scans"`
	ScanPeriodInSeconds uint64   `json:"scan_period_in_seconds"`
	UseHierarchy        uint64   `json:"use_hierarchy"`
	Buckets             []uint64 `json:"buckets"`
	Csei                []uint64 `json:"csei"`
	Dsei                []uint64 `json:"dsei"`
	Cfei                []uint64 `json:"cfei"`
	Dfei                []uint64 `json:"dfei"`
	Csui                []uint64 `json:"csui"`
	Dsui                []uint64 `json:"dsui"`
	Cfui                []uint64 `json:"cfui"`
	Dfui                []uint64 `json:"dfui"`
	Csea                []uint64 `json:"csea"`
	Dsea                []uint64 `json:"dsea"`
	Cfea                []uint64 `json:"cfea"`
	Dfea                []uint64 `json:"dfea"`
	Csua                []uint64 `json:"csua"`
	Dsua                []uint64 `json:"dsua"`
	Cfua                []uint64 `json:"cfua"`
	Dfua                []uint64 `json:"dfua"`
	Slab                []uint64 `json:"slab"`
}

type KidledConfig struct {
	ScanPeriodInseconds uint32
	UseHierarchy        uint8
	KidledColdBoundary  int
}

func ParseMemoryIdlePageStats(content string) (*ColdPageInfoByKidled, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// boundary is the index of [1,2)  [2,5)  [5,15)  [15,30)  [30,60)  [60,120)  [120,240)  [240,+inf).
// if boundary is equal to 3, it will compute sum([5*scan_period_scands,+inf)) of cold page cache
func (i *ColdPageInfoByKidled) GetColdPageTotalBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// check kidled and set var isSupportColdSupport
func IsKidledSupport() bool { _ = "STUB: not implemented"; return false }

func GetIsSupportColdMemory() bool { _ = "STUB: not implemented"; return false }

func SetIsSupportColdMemory(flag bool) { _ = "STUB: not implemented"; return }

func GetIsStartColdMemory() bool { _ = "STUB: not implemented"; return false }

func SetIsStartColdMemory(flag bool) { _ = "STUB: not implemented"; return }

func SetKidledScanPeriodInSeconds(period uint32) error { _ = "STUB: not implemented"; return nil }

func SetKidledUseHierarchy(useHierarchy uint8) error { _ = "STUB: not implemented"; return nil }

func GetKidledColdBoundary() int { _ = "STUB: not implemented"; return 0 }

func SetKidledColdBoundary(boudary int) { _ = "STUB: not implemented"; return }

func NewDefaultKidledConfig() *KidledConfig { _ = "STUB: not implemented"; return nil }

func sumUint64Slice(nums ...[]uint64) uint64 { _ = "STUB: not implemented"; return 0 }
