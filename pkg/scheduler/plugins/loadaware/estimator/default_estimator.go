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

package estimator

import (
	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

const (
	defaultEstimatorName = "defaultEstimator"

	// DefaultMilliCPURequest defines default milli cpu request number.
	DefaultMilliCPURequest int64 = 250 // 0.25 core
	// DefaultMemoryRequest defines default memory request size.
	DefaultMemoryRequest int64 = 200 * 1024 * 1024 // 200 MB
)

type DefaultEstimator struct {
	scalingFactors map[corev1.ResourceName]int64
	allowCustomize bool
}

func NewDefaultEstimator(args *config.LoadAwareSchedulingArgs, handle fwktype.Handle) (Estimator, error) {
	_ = "STUB: not implemented"
	return *new(Estimator), nil
}

func (e *DefaultEstimator) Name() string { _ = "STUB: not implemented"; return "" }

func (e *DefaultEstimator) EstimatePod(pod *corev1.Pod) (map[corev1.ResourceName]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func estimatedPodUsed(pod *corev1.Pod, scalingFactors map[corev1.ResourceName]int64) map[corev1.ResourceName]int64 {
	_ = "STUB: not implemented"
	return nil
}

// TODO(joseph): Do we need to differentiate scalingFactor according to Koordinator Priority type?
func estimatedUsedByResource(requests, limits corev1.ResourceList, resourceName corev1.ResourceName, scalingFactor int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *DefaultEstimator) EstimateNode(node *corev1.Node) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}
