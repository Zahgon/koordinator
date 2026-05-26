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

package v1

import (
	"math"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/ptr"

	"github.com/koordinator-sh/koordinator/apis/extension"
)

var (
	defaultNodeMetricExpirationSeconds int64 = 180

	defaultResourceWeights = map[corev1.ResourceName]int64{
		corev1.ResourceCPU:    1,
		corev1.ResourceMemory: 1,
	}

	defaultUsageThresholds = map[corev1.ResourceName]int64{
		corev1.ResourceCPU:    65, // 65%
		corev1.ResourceMemory: 95, // 95%
	}

	defaultEstimatedScalingFactors = map[corev1.ResourceName]int64{
		corev1.ResourceCPU:    85, // 85%
		corev1.ResourceMemory: 70, // 70%
	}

	defaultPreferredCPUBindPolicy = CPUBindPolicyFullPCPUs

	defaultEnablePreemption     = ptr.To[bool](false)
	defaultAwareNetworkTopology = ptr.To[bool](false)
	defaultGangMatchPolicy      = ptr.To[string](extension.GangMatchPolicyOnceSatisfied)

	defaultMinCandidateNodesPercentage  = ptr.To[int32](10)
	defaultMinCandidateNodesAbsolute    = ptr.To[int32](100)
	defaultReservationControllerWorkers = ptr.To[int32](1)
	defaultGCDurationSeconds            = ptr.To[int64](86400)
	defaultGCIntervalSeconds            = ptr.To[int64](60)
	defaultResyncIntervalSeconds        = ptr.To[int64](60)

	defaultDelayEvictTime       = 120 * time.Second
	defaultRevokePodInterval    = 1 * time.Second
	defaultDefaultQuotaGroupMax = corev1.ResourceList{
		// pkg/scheduler/plugins/elasticquota/controller.go syncHandler patch will overflow when the spec Max/Min is too high.
		corev1.ResourceCPU:    *resource.NewQuantity(math.MaxInt64/5, resource.DecimalSI),
		corev1.ResourceMemory: *resource.NewQuantity(math.MaxInt64/5, resource.DecimalSI),
	}
	defaultSystemQuotaGroupMax = corev1.ResourceList{
		corev1.ResourceCPU:    *resource.NewQuantity(math.MaxInt64/5, resource.DecimalSI),
		corev1.ResourceMemory: *resource.NewQuantity(math.MaxInt64/5, resource.BinarySI),
	}

	defaultQuotaGroupNamespace = "koordinator-system"

	defaultMonitorAllQuotas              = ptr.To[bool](false)
	defaultEnableCheckParentQuota        = ptr.To[bool](false)
	defaultEnableRuntimeQuota            = ptr.To[bool](true)
	defaultEnableMinQuotaScale           = ptr.To[bool](true)
	defaultDisableDefaultQuotaPreemption = ptr.To[bool](true)
	defaultEnableQueueHint               = ptr.To[bool](false)

	defaultTimeout                     = 600 * time.Second
	defaultControllerWorkers           = 1
	defaultQuotaSnapshotUpdateInterval = 120 * time.Second

	defaultGPUSharedResourceTemplatesConfig = &GPUSharedResourceTemplatesConfig{
		ConfigMapNamespace: "koordinator-system",
		ConfigMapName:      "gpu-shared-resource-templates",
		MatchedResources: []corev1.ResourceName{
			extension.ResourceHuaweiNPUCore,
		},
	}

	defaultMaxHintNodes = ptr.To[int32](100)
)

// SetDefaults_LoadAwareSchedulingArgs sets the default parameters for LoadAwareScheduling plugin.
func SetDefaults_LoadAwareSchedulingArgs(obj *LoadAwareSchedulingArgs) {
	_ = "STUB: not implemented"
	return
}

// SetDefaults_NodeNUMAResourceArgs sets the default parameters for NodeNUMANodeResource plugin.
func SetDefaults_NodeNUMAResourceArgs(obj *NodeNUMAResourceArgs) { _ = "STUB: not implemented"; return }

func SetDefaults_ReservationArgs(obj *ReservationArgs) { _ = "STUB: not implemented"; return }

func SetDefaults_ElasticQuotaArgs(obj *ElasticQuotaArgs) { _ = "STUB: not implemented"; return }

func SetDefaults_CoschedulingArgs(obj *CoschedulingArgs) { _ = "STUB: not implemented"; return }

func SetDefaults_DeviceShareArgs(obj *DeviceShareArgs) { _ = "STUB: not implemented"; return }

// By default, LeastAllocate is used to ensure high availability of applications

// SetDefaults_SchedulingHintArgs sets the default parameters for SchedulingHint plugin.
func SetDefaults_SchedulingHintArgs(obj *SchedulingHintArgs) { _ = "STUB: not implemented"; return }
