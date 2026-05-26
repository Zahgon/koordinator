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
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/loadaware/estimator"
)

const (
	Name = "LoadAwareScheduling"

	// incomingPodEstimatedStateKey is the key in CycleState to LoadAwareScheduling estimated resources for incoming pod.
	// Using the name of the plugin will likely help us avoid collisions with other plugins.
	incomingPodEstimatedStateKey = "IncomingPodEstimated" + Name

	ErrReasonNodeMetricExpired              = "node(s) nodeMetric expired"
	ErrReasonUsageExceedThreshold           = "node(s) %s usage exceed threshold"
	ErrReasonAggregatedUsageExceedThreshold = "node(s) %s aggregated usage exceed threshold"
	ErrReasonFailedEstimatePod
)

const (
	// DefaultMilliCPURequest defines default milli cpu request number.
	DefaultMilliCPURequest int64 = 250 // 0.25 core
	// DefaultMemoryRequest defines default memory request size.
	DefaultMemoryRequest int64 = 200 * 1024 * 1024 // 200 MB
	// DefaultNodeMetricReportInterval defines the default koodlet report NodeMetric interval.
	DefaultNodeMetricReportInterval = 60 * time.Second
)

var (
	_ fwktype.EnqueueExtensions = &Plugin{}

	_ fwktype.PreFilterPlugin = &Plugin{}
	_ fwktype.FilterPlugin    = &Plugin{}
	_ fwktype.ScorePlugin     = &Plugin{}
	_ fwktype.ReservePlugin   = &Plugin{}
)

type Plugin struct {
	handle         fwktype.Handle
	args           *config.LoadAwareSchedulingArgs
	vectorizer     ResourceVectorizer
	filterProfile  *usageThresholdsFilterProfile
	scoreWeights   ResourceVector
	estimator      estimator.Estimator
	podAssignCache *podAssignCache
}

func New(_ context.Context, args runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (p *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) EventsToRegister(_ context.Context) ([]fwktype.ClusterEventWithHint, error) {
	_ = "STUB: not implemented"
	// To register a custom event, follow the naming convention at:
	// https://github.com/kubernetes/kubernetes/blob/e1ad9bee5bba8fbe85a6bf6201379ce8b1a611b1/pkg/scheduler/eventhandlers.go#L415-L422
	return nil, nil
}

func (p *Plugin) PreFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	// add estimated resources for incoming pod to cache
	return nil, nil
}

// PreFilterExtensions returns a PreFilterExtensions interface if the plugin implements one.
func (p *Plugin) PreFilterExtensions() fwktype.PreFilterExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.PreFilterExtensions)
}

func (p *Plugin) Filter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// skip if filter thresholds are disabled

// For nodes that lack load information, fall back to the situation where there is no load-aware scheduling.
// Some nodes in the cluster do not install the koordlet, but users newly created Pod use koord-scheduler to schedule,
// and the load-aware scheduling itself is an optimization, so we should skip these nodes.

func (p *Plugin) ScoreExtensions() fwktype.ScoreExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.ScoreExtensions)
}

func (p *Plugin) Reserve(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Unreserve(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) Score(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// skip if score weights are disabled

// caused by load-aware scheduling itself is an optimization,
// so we should skip the node and score the node 0

// try to use state cache before EstimatePod
func (p *Plugin) addEstimatedOfIncoming(estimated ResourceVector, cycleState fwktype.CycleState, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// use len=0 but not empty vector to indicate error occurred

func (p *Plugin) filterNodeUsage(nodeName string, pod *corev1.Pod, usageThresholds, estimatedUsed, allocatable ResourceVector, isAgg bool) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func loadAwareSchedulingScorer(dominantWeight int64, resToWeightMap, used, allocatable ResourceVector) int64 {
	_ = "STUB: not implemented"
	return 0
}

func leastUsedScore(used, capacity int64) int64 { _ = "STUB: not implemented"; return 0 }
