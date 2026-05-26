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
	"path/filepath"
	"sync"

	"github.com/koordinator-sh/koordinator/pkg/util"
)

const (
	ResctrlName string = "resctrl"

	ResctrlDir string = "resctrl/"
	RdtInfoDir string = "info"
	L3CatDir   string = "L3"

	ResctrlSchemataName string = "schemata"
	ResctrlCbmMaskName  string = "cbm_mask"
	ResctrlTasksName    string = "tasks"

	// L3SchemataPrefix is the prefix of l3 cat schemata
	L3SchemataPrefix = "L3"
	// MbSchemataPrefix is the prefix of mba schemata
	MbSchemataPrefix = "MB"

	ResctrlMonData          = "mon_data"
	ResctrlLLCOccupancyName = "llc_occupancy"
	ResctrlMBMLocalName     = "mbm_local_bytes"
	ResctrlMBMTotalName     = "mbm_total_bytes"

	// other cpu vendor like "GenuineIntel"
	AMD_VENDOR_ID   = "AuthenticAMD"
	INTEL_VENDOR_ID = "GenuineIntel"
)

var (
	initLock                  sync.Mutex
	isInit                    bool
	isSupportResctrl          bool
	isSupportResctrlCollector bool
	collectorOnceFunc         sync.Once
	CacheIdsCacheFunc         func() ([]int, error)
	ARM_VENDOR_ID_MAP         = map[string]struct{}{} // support MPAM ARM vendor ids
)

func init() {
	CacheIdsCacheFunc = util.OnceValues(GetCacheIds)
}

func isCPUSupportResctrl() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func isKernelSupportResctrl() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// AMD CPU support resctrl by default

func IsVendorSupportResctrl() bool { _ = "STUB: not implemented"; return false }

func IsSupportResctrl() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Kernel cmdline not set resctrl features does not ensure feature must be disabled.

func IsSupportResctrlCollector() (bool, error) { _ = "STUB: not implemented"; return false, nil }

var (
	ResctrlRoot         = NewCommonResctrlResource("", "")
	ResctrlSchemata     = NewCommonResctrlResource(ResctrlSchemataName, "")
	ResctrlTasks        = NewCommonResctrlResource(ResctrlTasksName, "")
	ResctrlL3CbmMask    = NewCommonResctrlResource(ResctrlCbmMaskName, filepath.Join(RdtInfoDir, L3CatDir))
	ResctrlLLCOccupancy = NewCommonResctrlResource(ResctrlLLCOccupancyName, "")
	ResctrlMBLocal      = NewCommonResctrlResource(ResctrlMBMLocalName, "")
	ResctrlMBTotal      = NewCommonResctrlResource(ResctrlMBMTotalName, "")
)

var _ Resource = &ResctrlResource{}

type ResctrlResource struct {
	Type           ResourceType
	FileName       string
	Subdir         string
	CheckSupported func(r Resource, parentDir string) (isSupported bool, msg string)
	Validator      ResourceValidator
}

func (r *ResctrlResource) ResourceType() ResourceType {
	_ = "STUB: not implemented"
	return *new(ResourceType)
}

func (r *ResctrlResource) Path(parentDir string) string {
	_ = "STUB: not implemented"
	// parentDir for resctrl is like: `/`, `LS/`, `BE`
	return ""
}

func (r *ResctrlResource) IsSupported(parentDir string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (r *ResctrlResource) IsValid(v string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (r *ResctrlResource) WithValidator(validator ResourceValidator) Resource {
	_ = "STUB: not implemented"
	return *new(Resource)
}

func (r *ResctrlResource) WithSupported(isSupported bool, msg string) Resource {
	_ = "STUB: not implemented"
	return *new(Resource)
}

func (r *ResctrlResource) WithCheckSupported(checkSupportedFn func(r Resource, parentDir string) (isSupported bool, msg string)) Resource {
	_ = "STUB: not implemented"
	return *new(Resource)
}

func (r *ResctrlResource) WithCheckOnce(isCheckOnce bool) Resource {
	_ = "STUB: not implemented"
	return *new(Resource)
}

func NewCommonResctrlResource(filename string, subdir string) Resource {
	_ = "STUB: not implemented"
	return *new(Resource)
}

type ResctrlSchemataRaw struct {
	L3    map[int]int64
	MB    map[int]int64
	L3Num int
}

func NewResctrlSchemataRaw(cacheids []int) *ResctrlSchemataRaw {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResctrlSchemataRaw) WithL3Num(l3Num int) *ResctrlSchemataRaw {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResctrlSchemataRaw) WithL3Mask(mask string) *ResctrlSchemataRaw {
	_ = "STUB: not implemented"
	// l3 mask MUST be a valid hex
	return nil
}

func (r *ResctrlSchemataRaw) WithMB(valueOrPercent string) *ResctrlSchemataRaw {
	_ = "STUB: not implemented"
	// mba valueOrPercent MUST be a valid integer
	// for intel: "MB:0=100;1=100"
	// for amd format: "MB:0=2048;1=2048;2=2048;3=2048"
	return nil
}

func (r *ResctrlSchemataRaw) DeepCopy() *ResctrlSchemataRaw { _ = "STUB: not implemented"; return nil }

func (r *ResctrlSchemataRaw) Prefix() string { _ = "STUB: not implemented"; return "" }

func (r *ResctrlSchemataRaw) L3Number() int { _ = "STUB: not implemented"; return 0 }

func (r *ResctrlSchemataRaw) CacheIds() []int {
	_ = "STUB: not implemented"
	// TODO: consider situation that L3 number and the MB number are the same.
	return nil
}

func (r *ResctrlSchemataRaw) L3String() string { _ = "STUB: not implemented"; return "" }

// the last ';' will be auto ignored

// the trailing '\n' is necessary to append

func (r *ResctrlSchemataRaw) MBString() string { _ = "STUB: not implemented"; return "" }

// the last ';' will be auto ignored

// the trailing '\n' is necessary to append

func (r *ResctrlSchemataRaw) Equal(a *ResctrlSchemataRaw) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (r *ResctrlSchemataRaw) Validate() (bool, string) { _ = "STUB: not implemented"; return false, "" }

func (r *ResctrlSchemataRaw) ValidateL3() (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (r *ResctrlSchemataRaw) ValidateMB() (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// ParseResctrlSchemata parses the resctrl schemata of given cgroup, and returns the l3_cat masks and mba masks.
// Set l3Num=-1 to use the read L3 number from the schemata.
// @content `L3:0=fff;1=fff\nMB:0=100;1=100\n` (may have additional lines (e.g. ARM MPAM))
// @l3Num 2
// @return {L3: {0: "fff", 1: "fff"}, MB: {0: "100", 1: "100"}}, nil
func (r *ResctrlSchemataRaw) ParseResctrlSchemata(content string, l3Num int) error {
	_ = "STUB: not implemented"
	return nil
}

// @return /sys/fs/resctrl
func GetResctrlSubsystemDirPath() string { _ = "STUB: not implemented"; return "" }

// @groupPath BE
// @return /sys/fs/resctrl/BE
func GetResctrlGroupRootDirPath(groupPath string) string { _ = "STUB: not implemented"; return "" }

// @return /sys/fs/resctrl/info/L3/cbm_mask
func GetResctrlL3CbmFilePath() string { _ = "STUB: not implemented"; return "" }

// @groupPath BE
// @return /sys/fs/resctrl/BE/schemata
func GetResctrlSchemataFilePath(groupPath string) string { _ = "STUB: not implemented"; return "" }

// @groupPath BE
// @return /sys/fs/resctrl/BE/tasks
func GetResctrlTasksFilePath(groupPath string) string { _ = "STUB: not implemented"; return "" }

// @parentDir BE
// @return /sys/fs/resctrl/BE/mon_data
func GetResctrlMonDataPath(parentDir string) string { _ = "STUB: not implemented"; return "" }

func ReadResctrlSchemataRaw(schemataFile string, l3Num int) (*ResctrlSchemataRaw, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCacheIds get cache ids from schemata file.
// e.g. schemata=`L3:0=fff;1=fff\nMB:0=100;1=100\n` -> [0, 1]
func GetCacheIds() ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseResctrlSchemataMap parses the content of resctrl schemata.
// e.g. schemata=`L3:0=fff;1=fff\nMB:0=100;1=100\n` -> `{"L3": {0: "fff", 1: "fff"}, "MB": {0: "100", 1: "100"}}`
func ParseResctrlSchemataMap(content string) map[string]map[int]string {
	_ = "STUB: not implemented"
	return nil
}

// `L3:0=fff;1=fff`, `MB:0=100;1=100`

// {`L3`, `0=fff;1=fff`}, {`MB`, `0=100;1=100`}

// {`0=fff`, `1=fff`}, {`0=100`, `1=100`}

// {`0`, `fff`}, {`1`, `100`}

// {0: `fff`}

// {`L3`: {0: `fff`, 1: `fff`} }

// ReadCatL3CbmString reads and returns the value of cat l3 cbm_mask
func ReadCatL3CbmString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadResctrlTasksMap reads and returns the map of given resctrl group's task ids
func ReadResctrlTasksMap(groupPath string) (map[int32]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckAndTryEnableResctrlCat checks if resctrl and l3_cat are enabled; if not, try to enable the features by mount
// resctrl subsystem; See MountResctrlSubsystem() for the detail.
// It returns whether the resctrl cat is enabled, and the error if failed to enable or to check resctrl interfaces
func CheckAndTryEnableResctrlCat() error {
	_ = "STUB: not implemented"
	// resctrl cat is correctly enabled: l3_cbm path exists
	return nil
}

// double check l3_cbm path to ensure both resctrl and cat are correctly enabled

func InitCatGroupIfNotExist(group string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func CheckResctrlSchemataValid() error { _ = "STUB: not implemented"; return nil }

func CalculateCatL3MaskValue(cbm uint, startPercent, endPercent int64) (string, error) {
	_ = "STUB: not implemented"
	// check if the parsed cbm value is valid, eg. 0xff, 0x1, 0x7ff, ...
	// NOTE: (Cache Bit Masks) X86 hardware requires that these masks have all the '1' bits in a contiguous block.
	//       ref: https://www.kernel.org/doc/Documentation/x86/intel_rdt_ui.txt
	// since the input cbm here is the cbm value of the resctrl root, every lower bit is required to be `1` additionally
	return "", nil
}

// check if the startPercent and endPercent are valid

// calculate a bit mask belonging to interval [startPercent% * ways, endPercent% * ways)
// eg.
// cbm 0x3ff ('b1111111111), start 10%, end 80%
// ways 10, l3Mask 0xfe ('b11111110)
// cbm 0x7ff ('b11111111111), start 10%, end 50%
// ways 11, l3Mask 0x3c ('b111100)
// cbm 0x7ff ('b11111111111), start 0%, end 30%
// ways 11, l3Mask 0xf ('b1111)

// GetVendorIDByCPUInfo returns vendor_id like AuthenticAMD from cpu info, e.g.
// vendor_id       : AuthenticAMD
// vendor_id       : GenuineIntel
func GetVendorIDByCPUInfo(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// get "vendor_id" from first line

func isResctrlAvailableByCpuInfo(path string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func isResctrlMBMAvailableByCpuInfo(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isResctrlCQMAvailableByCpuInfo(path string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getCPUFlags(path string) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// file content example:
// BOOT_IMAGE=/boot/vmlinuz-4.19.91-24.1.al7.x86_64 root=UUID=231efa3b-302b-4e82-9445-0f7d5d353dda \
// crashkernel=0M-2G:0M,2G-8G:192M,8G-:256M cryptomgr.notests cgroup.memory=nokmem rcupdate.rcu_cpu_stall_timeout=300 \
// vring_force_dma_api biosdevname=0 net.ifnames=0 console=tty0 console=ttyS0,115200n8 noibrs \
// nvme_core.io_timeout=4294967295 nomodeset intel_idle.max_cstate=1 rdt=cmt,l3cat,l3cdp,mba
func isResctrlAvailableByKernelCmd(path string) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

type MBStatData map[string]uint64
