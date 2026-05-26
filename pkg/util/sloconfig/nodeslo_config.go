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

package sloconfig

import (
	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

var defautExtensions = &slov1alpha1.ExtensionsMap{}

func getDefaultExtensionsMap() *slov1alpha1.ExtensionsMap { _ = "STUB: not implemented"; return nil }

func RegisterDefaultExtensionsMap(extKey string, extCfg interface{}) {
	_ = "STUB: not implemented"
	return
}

// DefaultNodeSLOSpecConfig defines the default config of the nodeSLOSpec, which would be used by the resmgr
func DefaultNodeSLOSpecConfig() slov1alpha1.NodeSLOSpec {
	_ = "STUB: not implemented"
	return *new(slov1alpha1.NodeSLOSpec)
}

func DefaultResourceThresholdStrategy() *slov1alpha1.ResourceThresholdStrategy {
	_ = "STUB: not implemented"
	return nil
}

func DefaultCPUQOS(qos apiext.QoSClass) *slov1alpha1.CPUQOS { _ = "STUB: not implemented"; return nil }

// NOTE: Be careful to enable CPU Idle since it overrides and lock the cpu.shares/cpu.weight of the same
// cgroup to a minimal value. This can affect other components like Kubelet which wants to write
// cpu.shares/cpu.weight to other values.
// https://git.kernel.org/pub/scm/linux/kernel/git/tip/tip.git/commit/?id=304000390f88d049c85e9a0958ac5567f38816ee

// TODO https://github.com/koordinator-sh/koordinator/pull/94#discussion_r858786733
func DefaultResctrlQOS(qos apiext.QoSClass) *slov1alpha1.ResctrlQOS {
	_ = "STUB: not implemented"
	return nil
}

// DefaultMemoryQOS returns the recommended configuration for memory qos strategy.
// Please refer to `apis/slo/v1alpha1` for the definition of each field.
// In the recommended configuration, all abilities of memcg qos are disable, including `MinLimitPercent`,
// `LowLimitPercent`, `ThrottlingPercent` since they are not fully beneficial to all scenarios. Whereas, they are still
// useful when the use case is determined. e.g. lock some memory to improve file read performance.
// Asynchronous memory reclaim is enabled by default to alleviate the direct reclaim pressure, including `WmarkRatio`
// and `WmarkScalePermill`. The watermark of async reclaim is not recommended to set too low, since lower the watermark
// the more excess reclamations.
// Memory min watermark grading corresponding to `WmarkMinAdj` is enabled. It benefits high-priority pods by postponing
// global reclaim when machine's free memory is below than `/proc/sys/vm/min_free_kbytes`.
func DefaultMemoryQOS(qos apiext.QoSClass) *slov1alpha1.MemoryQOS {
	_ = "STUB: not implemented"
	return nil
}

func DefaultResourceQOSPolicies() *slov1alpha1.ResourceQOSPolicies {
	_ = "STUB: not implemented"
	return nil
}

func DefaultResourceQOSStrategy() *slov1alpha1.ResourceQOSStrategy {
	_ = "STUB: not implemented"
	return nil
}

func NoneResourceQOS(qos apiext.QoSClass) *slov1alpha1.ResourceQOS {
	_ = "STUB: not implemented"
	// cgroup root case: only used by blkio qos
	return nil
}

func NoneCPUQOS() *slov1alpha1.CPUQOS { _ = "STUB: not implemented"; return nil }

func NoneResctrlQOS() *slov1alpha1.ResctrlQOS { _ = "STUB: not implemented"; return nil }

// NoneMemoryQOS returns the all-disabled configuration for memory qos strategy.
func NoneMemoryQOS() *slov1alpha1.MemoryQOS { _ = "STUB: not implemented"; return nil }

func NoneBlkIOQOS() *slov1alpha1.BlkIOQOS { _ = "STUB: not implemented"; return nil }

func NoneNetworkQOS() *slov1alpha1.NetworkQOS { _ = "STUB: not implemented"; return nil }

func NoneResourceQOSPolicies() *slov1alpha1.ResourceQOSPolicies {
	_ = "STUB: not implemented"
	return nil
}

// NoneResourceQOSStrategy indicates the qos strategy with all qos
func NoneResourceQOSStrategy() *slov1alpha1.ResourceQOSStrategy {
	_ = "STUB: not implemented"
	return nil
}

func DefaultCPUBurstStrategy() *slov1alpha1.CPUBurstStrategy { _ = "STUB: not implemented"; return nil }

func DefaultCPUBurstConfig() slov1alpha1.CPUBurstConfig {
	_ = "STUB: not implemented"
	return *new(slov1alpha1.CPUBurstConfig)
}

func DefaultSystemStrategy() *slov1alpha1.SystemStrategy { _ = "STUB: not implemented"; return nil }

func DefaultExtensions() *slov1alpha1.ExtensionsMap { _ = "STUB: not implemented"; return nil }
