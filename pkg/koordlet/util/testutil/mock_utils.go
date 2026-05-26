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

package testutil

import (
	"go.uber.org/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	mock_metriccache "github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache/mockmetriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type FakeRecorder struct {
	EventReason string
}

func (f *FakeRecorder) Event(object runtime.Object, eventType, reason, message string) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeRecorder) Eventf(object runtime.Object, eventType, reason, messageFmt string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeRecorder) AnnotatedEventf(object runtime.Object, annotations map[string]string, eventType, reason, messageFmt string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func MockTestNode(cpu, memory string) *corev1.Node { _ = "STUB: not implemented"; return nil }

func MockTestNodeWithExtendResource(cpu, memory string, allocatable, capacity corev1.ResourceList) *corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

func MockTestPod(qosClass apiext.QoSClass, name string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

var PodsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

var EvictionResource = schema.GroupVersionResource{Group: "policy", Version: "v1beta1", Resource: "evictions"}

var EvictionKind = schema.GroupVersionKind{Group: "policy", Version: "v1beta1", Kind: "Eviction"}

func BuildMockQueryResult(ctrl *gomock.Controller, querier *mock_metriccache.MockQuerier, factory *mock_metriccache.MockAggregateResultFactory,
	queryMeta metriccache.MetricMeta, value float64) {
	_ = "STUB: not implemented"
	return
}

func BuildMockQueryResultAndCount(ctrl *gomock.Controller, querier *mock_metriccache.MockQuerier, factory *mock_metriccache.MockAggregateResultFactory,
	queryMeta metriccache.MetricMeta) *mock_metriccache.MockAggregateResult {
	_ = "STUB: not implemented"
	return nil
}

func GetPodMetas(pods []*corev1.Pod) []*statesinformer.PodMeta {
	_ = "STUB: not implemented"
	return nil
}

func GetNodeSLOByThreshold(thresholdConfig *slov1alpha1.ResourceThresholdStrategy) *slov1alpha1.NodeSLO {
	_ = "STUB: not implemented"
	return nil
}

func MockTestPodWithQOS(kubeQosClass corev1.PodQOSClass, qosClass apiext.QoSClass) *statesinformer.PodMeta {
	_ = "STUB: not implemented"
	return nil
}

func DefaultQOSStrategy() *slov1alpha1.ResourceQOSStrategy { _ = "STUB: not implemented"; return nil }
