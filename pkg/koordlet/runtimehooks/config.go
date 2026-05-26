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

package runtimehooks

import (
	"flag"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"

	"github.com/koordinator-sh/koordinator/pkg/features"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/batchresource"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/coresched"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/cpunormalization"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/cpuset"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/gpu"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/groupidentity"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/rdma"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/resctrl"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/tc"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks/terwayqos"
)

const (
	// GroupIdentity sets pod cpu group identity(bvt) according to QoS.
	//
	// owner: @zwzhang0107 @saintube
	// alpha: v0.3
	// beta: v1.1
	GroupIdentity featuregate.Feature = "GroupIdentity"

	// CPUSetAllocator sets container cpuset according to allocate result from koord-scheduler for LSR/LS pods.
	//
	// owner: @saintube @zwzhang0107
	// alpha: v0.3
	// beta: v1.1
	CPUSetAllocator featuregate.Feature = "CPUSetAllocator"

	// GPUEnvInject injects gpu allocated env info according to allocate result from koord-scheduler.
	//
	// owner: @ZYecho @jasonliu747
	// alpha: v0.3
	// beta: v1.1
	GPUEnvInject featuregate.Feature = "GPUEnvInject"

	// RDMADeviceInject injects rdma device info according to allocate result from koord-scheduler.
	//
	// owner: @ZiMengSheng
	// alpha: v1.6
	RDMADeviceInject featuregate.Feature = "RDMADeviceInject"

	// BatchResource sets request and limits of cpu and memory on cgroup file according batch resources.
	//
	// owner: @saintube @zwzhang0107
	// alpha: v1.1
	BatchResource featuregate.Feature = "BatchResource"

	// CPUNormalization adjusts cpu cgroups value for cpu normalized LS pod.
	//
	// owner: @saintube @zwzhang0107
	// alpha: v1.4
	CPUNormalization featuregate.Feature = "CPUNormalization"

	// CoreSched manages Linux Core Scheduling cookies for containers who enable the core sched.
	// NOTE: CoreSched is an alternative policy of the CPU QoS, and it is exclusive to the Group Identity feature.
	//
	// owner: @saintube @zwzhang0107
	// alpha: v1.4
	CoreSched featuregate.Feature = "CoreSched"

	// TerwayQoS enables net QoS feature of koordlet.
	// owner: @l1b0k
	// alpha: v1.5
	TerwayQoS featuregate.Feature = "TerwayQoS"

	// TCNetworkQoS indicates a network qos implementation based on tc.
	// owner: @lucming
	// alpha: v1.5
	TCNetworkQoS featuregate.Feature = "TCNetworkQoS"

	// Resctrl adjusts LLC/MB value for pod.
	//
	// owner: @kangclzjc @saintube @zwzhang0107
	// alpha: v1.5
	Resctrl featuregate.Feature = "Resctrl"
)

var (
	defaultRuntimeHooksFG = map[featuregate.Feature]featuregate.FeatureSpec{
		GroupIdentity:    {Default: true, PreRelease: featuregate.Beta},
		CPUSetAllocator:  {Default: true, PreRelease: featuregate.Beta},
		GPUEnvInject:     {Default: false, PreRelease: featuregate.Alpha},
		RDMADeviceInject: {Default: false, PreRelease: featuregate.Alpha},
		BatchResource:    {Default: true, PreRelease: featuregate.Beta},
		CPUNormalization: {Default: false, PreRelease: featuregate.Alpha},
		CoreSched:        {Default: false, PreRelease: featuregate.Alpha},
		TerwayQoS:        {Default: false, PreRelease: featuregate.Alpha},
		TCNetworkQoS:     {Default: false, PreRelease: featuregate.Alpha},
		Resctrl:          {Default: false, PreRelease: featuregate.Alpha},
	}

	runtimeHookPlugins = map[featuregate.Feature]HookPlugin{
		GroupIdentity:    groupidentity.Object(),
		CPUSetAllocator:  cpuset.Object(),
		GPUEnvInject:     gpu.Object(),
		RDMADeviceInject: rdma.Object(),
		BatchResource:    batchresource.Object(),
		CPUNormalization: cpunormalization.Object(),
		CoreSched:        coresched.Object(),
		TerwayQoS:        terwayqos.Object(),
		TCNetworkQoS:     tc.Object(),
		Resctrl:          resctrl.Object(),
	}
)

type Config struct {
	RuntimeHooksNetwork             string
	RuntimeHooksAddr                string
	RuntimeHooksFailurePolicy       string
	RuntimeHooksPluginFailurePolicy string
	RuntimeHookConfigFilePath       string
	RuntimeHookHostEndpoint         string
	RuntimeHookDisableStages        []string
	RuntimeHooksNRI                 bool
	RuntimeHooksNRIConnectTimeout   time.Duration
	RuntimeHooksNRIBackOffDuration  time.Duration
	RuntimeHooksNRIBackOffCap       time.Duration
	RuntimeHooksNRIBackOffFactor    float64
	RuntimeHooksNRIBackOffSteps     int
	RuntimeHooksNRISocketPath       string
	RuntimeHooksNRIPluginName       string
	RuntimeHooksNRIPluginIndex      string
	RuntimeHookReconcileInterval    time.Duration
	RuntimeHookDisableUnsetCPUQuota bool
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

func (c *Config) InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func init() {
	runtime.Must(features.DefaultMutableKoordletFeatureGate.Add(defaultRuntimeHooksFG))
}
