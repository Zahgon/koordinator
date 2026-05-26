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

package util

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type ReleaseTargetType string
type ReleaseList map[ReleaseTargetType]corev1.ResourceList

const (
	ReleaseTargetTypeBatchResourceRequest ReleaseTargetType = "podBatchResourceRequest"
	ReleaseTargetTypeResourceUsed         ReleaseTargetType = "podUsed"
	ReleaseTargetTypeResourceRequest      ReleaseTargetType = "podResourceRequest"
	EvictedStr                                              = "evicted"
	EvictReasonPrefix                                       = "trigger by koordlet feature "
)

type PodEvictInfo struct {
	Pod *corev1.Pod

	// cpu details
	CpuUsage        float64
	MilliCPURequest int64 // cpu/mid-cpu/batch-cpu
	MilliCPUUsed    int64

	// mem details
	MemoryUsage   float64
	MemoryRequest int64 // memory/mid-memory/batch-memory
	MemoryUsed    int64

	// sort helper
	Priority      int32
	LabelPriority int64
}

type EvictTaskInfo struct {
	Reason            string
	SortedEvictPods   []*PodEvictInfo
	ReleaseTarget     ReleaseTargetType
	ToReleaseResource corev1.ResourceList
	// get relative resource list from pod for task
	GetPodResourceFunc func(*PodEvictInfo) corev1.ResourceList
}

type EvictionExecutor interface {
	Evict(pod *corev1.Pod, node *corev1.Node, releaseReason string, message string) bool
	IsPodEvicted(*corev1.Pod) bool
}

var customExecutorInitializer func(evictor *Evictor, onlyEvictByAPI bool) EvictionExecutor

type DefaultEvictionExecutor struct {
	OnlyEvictByAPI bool
	Evictor        *Evictor
}

func (d *DefaultEvictionExecutor) Evict(pod *corev1.Pod, node *corev1.Node, releaseReason string, message string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DefaultEvictionExecutor) IsPodEvicted(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func SetCustomEvictionExecutorInitializer(initializer func(*Evictor, bool) EvictionExecutor) {
	_ = "STUB: not implemented"
	return
}

func InitializeEvictionExecutor(evictor *Evictor, onlyEvictByAPI bool) EvictionExecutor {
	_ = "STUB: not implemented"
	return *new(EvictionExecutor)
}

func KillAndEvictPods(evictionExecutor EvictionExecutor, node *corev1.Node, tasks []*EvictTaskInfo) (ReleaseList, bool) {
	_ = "STUB: not implemented"
	return *new(ReleaseList), false
}

// note: same content only fetched from pod once only, used max instead of added

func IsEvictionPolicyAllowed(policy string, pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// EvictTaskCheck check if evict tasks finished,return not finished
func EvictTaskCheck(task *EvictTaskInfo, released ReleaseList) (bool, corev1.ResourceList) {
	_ = "STUB: not implemented"
	return false, *new(corev1.ResourceList)
}

// GetRequestTypeAndValueFromPod cpu return millvalue, memory return value
func GetRequestTypeAndValueFromPod(pod *corev1.Pod, name corev1.ResourceName) (corev1.ResourceName, int64) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceName), 0
}

// Sidecar containers run alongside regular containers and should be summed.

func GetRequestFromPod(pod *corev1.Pod, name corev1.ResourceName) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func isZeroResourceList(a corev1.ResourceList) bool { _ = "STUB: not implemented"; return false }

func subReleaseListNoNegative(a, b corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func GetPodPriorityLabel(pod *corev1.Pod, defaultPriority int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// merge b to a
func mergeResourceListByMax(a, b corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func addResource(a, b map[ReleaseTargetType]corev1.ResourceList) { _ = "STUB: not implemented"; return }

func ConvertQuantityToInt64(resourceName corev1.ResourceName, quantity resource.Quantity) int64 {
	_ = "STUB: not implemented"
	return 0
}

// include corev1.ResourceMemory, apiext.BatchCPU, apiext.BatchMemory, apiext.MidCPU, apiext.MidMemory

func ConvertInt64ToQuantity(resourceName corev1.ResourceName, value int64) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// include corev1.ResourceMemory, apiext.BatchMemory, apiext.MidMemory
