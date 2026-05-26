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
	"time"

	utilsysctl "k8s.io/component-helpers/node/util/sysctl"
)

const (
	SysctlSubDir          = "sys"
	KernelCmdlineFileName = "cmdline"
	HugepageDir           = "hugepages"
	nrPath                = "nr_hugepages"

	KernelSchedGroupIdentityEnable = SysKernelRelativePath + SchedGroupIdentityEnabledFileName
	KernelSchedCore                = SysKernelRelativePath + SchedCoreFileName

	SysNUMASubDir   = "bus/node/devices"
	SysPCIDeviceDir = "bus/pci/devices"

	SysCPUSMTActiveSubPath       = "devices/system/cpu/smt/active"
	SysIntelPStateNoTurboSubPath = "devices/system/cpu/intel_pstate/no_turbo"
)

var (
	// Jiffies is the duration unit of CPU stats. Normally, it is 10ms.
	Jiffies = float64(10 * time.Millisecond)
)

// initJiffies use command "getconf CLK_TCK" to fetch the clock tick on current host,
// if the command doesn't exist, uses the default value 10ms for jiffies
func initJiffies() error { _ = "STUB: not implemented"; return nil }

func GetPeriodTicks(start, end time.Time) float64 { _ = "STUB: not implemented"; return 0 }

func GetSysRootDir() string { _ = "STUB: not implemented"; return "" }

func GetSysNUMADir() string { _ = "STUB: not implemented"; return "" }

func GetNUMAMemInfoPath(numaNodeSubDir string) string { _ = "STUB: not implemented"; return "" }

func GetNUMAHugepagesDir(numaNodeSubDir string) string { _ = "STUB: not implemented"; return "" }

func GetNUMAHugepagesNrPath(numaNodeSubDir string, page string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetCPUInfoPath() string { _ = "STUB: not implemented"; return "" }

func GetSysCPUSMTActivePath() string { _ = "STUB: not implemented"; return "" }

func GetSysIntelPStateNoTurboPath() string { _ = "STUB: not implemented"; return "" }

func GetProcSysFilePath(file string) string { _ = "STUB: not implemented"; return "" }

func GetPCIDeviceDir() string { _ = "STUB: not implemented"; return "" }

var _ utilsysctl.Interface = &ProcSysctl{}

// ProcSysctl implements Interface by reading and writing files under /proc/sys
type ProcSysctl struct{}

func NewProcSysctl() utilsysctl.Interface {
	_ = "STUB: not implemented"
	return *new(utilsysctl.Interface)
}

func (*ProcSysctl) GetSysctl(sysctl string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// SetSysctl modifies the specified sysctl flag to the new value
func (*ProcSysctl) SetSysctl(sysctl string, newVal int) error {
	_ = "STUB: not implemented"
	return nil
}

func IsGroupIdentitySysctlSupported() bool { _ = "STUB: not implemented"; return false }

func GetSchedGroupIdentity() (bool, error) {
	_ = "STUB: not implemented"

	// 0: disabled; 1: enabled
	return false, nil
}

func SetSchedGroupIdentity(enable bool) error { _ = "STUB: not implemented"; return nil }

// 0: disabled; 1: enabled

func GetSchedCore() (bool, error) {
	_ = "STUB: not implemented"

	// 0: disabled; 1: enabled
	return false, nil
}

func SetSchedCore(enable bool) error { _ = "STUB: not implemented"; return nil }

// 0: disabled; 1: enabled

func GetSchedFeatures() (map[string]bool, error) { _ = "STUB: not implemented"; return nil, nil }

func SetSchedFeatures(featureMap map[string]bool, valueMap map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

// write XXX to sched_features

// write NO_XXX to sched_features
