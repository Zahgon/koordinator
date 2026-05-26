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

package groupidentity

import (
	"sync"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
)

const (
	name        = "GroupIdentity"
	description = "set bvt value by priority and qos class"
)

type bvtPlugin struct {
	rule        *bvtRule
	ruleRWMutex sync.RWMutex

	sysSupported             *bool
	hasKernelEnabled         *bool // whether kernel is configurable for enabling bvt (via `kernel.sched_group_identity_enabled`)
	coreSchedSysctlSupported *bool // whether core sched is supported by the sysctl

	executor resourceexecutor.ResourceUpdateExecutor
}

func (b *bvtPlugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

func (b *bvtPlugin) SystemSupported() bool { _ = "STUB: not implemented"; return false }

func (b *bvtPlugin) hasKernelEnable() bool { _ = "STUB: not implemented"; return false }

// tryDisableCoreSched tries disabling the core scheduling via sysctl to safely enable the group identity.
func (b *bvtPlugin) tryDisableCoreSched() error { _ = "STUB: not implemented"; return nil }

// initSysctl initializes the sysctl for the group identity.
// It returns whether cgroups need to set, e.g. if the sysctl config is not disabled.
func (b *bvtPlugin) initSysctl() error { _ = "STUB: not implemented"; return nil }

// kernel does not support bvt sysctl

// NOTE: Currently the kernel feature core scheduling is strictly excluded with the group identity's
//       bvt=-1. So we have to check if the CoreSched can be disabled before enabling group identity.

// try to set bvt kernel enabled via sysctl when the sysctl config is disabled or unknown
// https://github.com/koordinator-sh/koordinator/pull/1172

// isSysctlEnabled checks if the sysctl configuration for the bvt (group identity) is enabled.
// It returns whether cgroups need to set, e.g. if the sysctl config is not disabled.
func (b *bvtPlugin) isSysctlEnabled() (bool, error) {
	_ = "STUB: not implemented"
	// NOTE: bvt (group identity) is supported and can be initialized in the system if:
	//  1. anolis os kernel (<26.4): cgroup cpu.bvt_warp_ns exists but sysctl kernel.sched_group_identity_enabled no exist,
	//     the bvt feature is enabled by default, no need to set sysctl.
	//  2. anolis os kernel (>=26.4): both cgroup cpu.bvt_warp_ns and sysctl kernel.sched_group_identity_enabled exist,
	//     the bvt feature is enabled when kernel.sched_group_identity_enabled is set as `1`.
	return false, nil
}

// consider enabled when kernel does not support bvt sysctl

var singleton *bvtPlugin

func Object() *bvtPlugin { _ = "STUB: not implemented"; return nil }
