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

package loadaware

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

func isNodeMetricExpired(nodeMetric *slov1alpha1.NodeMetric, nodeMetricExpirationSeconds int64) bool {
	_ = "STUB: not implemented"
	return false
}

func getNodeMetricReportInterval(nodeMetric *slov1alpha1.NodeMetric) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type usageThresholdsFilterProfile struct {
	UsageThresholds     ResourceVector
	ProdUsageThresholds ResourceVector
	AggregatedUsage     *aggregatedUsageFilterProfile
}

type aggregatedUsageFilterProfile struct {
	UsageThresholds         ResourceVector
	UsageAggregationType    extension.AggregationType
	UsageAggregatedDuration metav1.Duration
}

func NewUsageThresholdsFilterProfile(args *config.LoadAwareSchedulingArgs, vectorizer ResourceVectorizer) *usageThresholdsFilterProfile {
	_ = "STUB: not implemented"
	return nil
}

// NOTICE: unknown resource name in custom usage thresholds for vectorizer will be skipped in calculation.
//
// Currently, we add all supported resources (cpu, memory) collected by koordlet with hard code
// that can be used in load aware plugin for compatibility.
func (tfp *usageThresholdsFilterProfile) generateUsageThresholdsFilterProfile(node *corev1.Node, vectorizer ResourceVectorizer) *usageThresholdsFilterProfile {
	_ = "STUB: not implemented"
	return nil
}

func getResourceValue(resourceName corev1.ResourceName, quantity resource.Quantity) int64 {
	_ = "STUB: not implemented"
	return 0
}

func getResourceQuantity(resourceName corev1.ResourceName, value int64) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// isDaemonSetPod returns true if the pod is a IsDaemonSetPod.
func isDaemonSetPod(ownerRefList []metav1.OwnerReference) bool {
	_ = "STUB: not implemented"
	return false
}

type ResourceVectorizer []corev1.ResourceName
type ResourceVector []int64

func NewResourceVectorizer(names ...corev1.ResourceName) ResourceVectorizer {
	_ = "STUB: not implemented"
	return *new(ResourceVectorizer)
}

// cpu and memory are added by default for custom usage thresholds compatibility.
func NewResourceVectorizerFromArgs(args *config.LoadAwareSchedulingArgs) ResourceVectorizer {
	_ = "STUB: not implemented"
	return *new(ResourceVectorizer)
}

// NOTICE: unknown resource name will be ignored in vectorization
func (rv ResourceVectorizer) ToVec(list corev1.ResourceList) ResourceVector {
	_ = "STUB: not implemented"
	return *new(ResourceVector)
}

// NOTICE: unknown resource name will be ignored in vectorization
func (rv ResourceVectorizer) ToFactorVec(list map[corev1.ResourceName]int64) ResourceVector {
	_ = "STUB: not implemented"
	return *new(ResourceVector)
}

func (rv ResourceVectorizer) ToList(vec ResourceVector) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func (rv ResourceVectorizer) ToFactorList(vec ResourceVector) map[corev1.ResourceName]int64 {
	_ = "STUB: not implemented"
	return nil
}

func (rv ResourceVectorizer) EmptyVec() ResourceVector {
	_ = "STUB: not implemented"
	return *new(ResourceVector)
}

func (v ResourceVector) Empty() bool { _ = "STUB: not implemented"; return false }

func (v ResourceVector) Add(y ResourceVector) { _ = "STUB: not implemented"; return }

// v = v + max(0, x - y)
func (v ResourceVector) AddDelta(x, y ResourceVector) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

func (v ResourceVector) Sub(y ResourceVector) { _ = "STUB: not implemented"; return }

// v = v - max(0, x - y)
func (v ResourceVector) SubDelta(x, y ResourceVector) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

func (v ResourceVector) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}
