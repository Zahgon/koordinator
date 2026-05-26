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

package noderesourcesfitplus

import (
	v1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

type ResourceAllocationPriority struct {
	scorer func(nodeName string, args *config.NodeResourcesFitPlusArgs, requestedMap, allocatableMap map[v1.ResourceName]int64) int64
}

func mostRequestedScore(requested, capacity int64) int64 { _ = "STUB: not implemented"; return 0 }

func leastRequestedScore(requested, capacity int64) int64 { _ = "STUB: not implemented"; return 0 }

func resourceScorer(nodeName string, args *config.NodeResourcesFitPlusArgs, requestedMap, allocatableMap map[v1.ResourceName]int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (r *ResourceAllocationPriority) getResourceScore(args *config.NodeResourcesFitPlusArgs, podRequestNames []v1.ResourceName, pod *v1.Pod, nodeInfo fwktype.NodeInfo, nodeName string) int64 {
	_ = "STUB: not implemented"
	return 0
}

func computePodResourceRequest(pod *v1.Pod) *preScoreState {
	_ = "STUB: not implemented"
	// pod hasn't scheduled yet so we don't need to worry about InPlacePodVerticalScalingEnabled
	return nil
}

// resourceToValueMap contains resource name and score.
type resourceToValueMap map[v1.ResourceName]int64

// calculateResourceAllocatableRequest returns resources Allocatable and Requested values
func calculateResourceAllocatableRequest(nodeInfo fwktype.NodeInfo, pod *v1.Pod, resource v1.ResourceName) (int64, int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// calculatePodResourceRequest returns the total non-zero requests. If Overhead is defined for the pod and the
// PodOverhead feature is enabled, the Overhead is added to the result.
// It follows the KEP-753 sidecar container resource calculation:
// podResourceRequest = max(sum(Containers) + sum(SidecarInitContainers), max(Regular_InitContainer + preceding_sidecars)) + overHead
func calculatePodResourceRequest(pod *v1.Pod, resource v1.ResourceName) int64 {
	_ = "STUB: not implemented"
	return 0
}

// Sidecar containers (initContainers with restartPolicy=Always) run alongside
// regular containers, so their requests should be summed.
// Regular init containers use max-based comparison.

// If Overhead is being utilized, add to the total requests for the pod

// GetNonzeroRequestForResource returns the default resource request if none is found or
// what is provided on the request.
func GetNonzeroRequestForResource(resource v1.ResourceName, requests *v1.ResourceList) int64 {
	_ = "STUB: not implemented"
	return 0
}

// Override if un-set, but not if explicitly set to zero

// Override if un-set, but not if explicitly set to zero
