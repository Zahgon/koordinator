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

package prediction

import (
	"time"

	v1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	DefaultNodeID        = "__node__"
	DefaultNodeItemIDFmt = "__node-%s__"
	SystemItemID         = "sys"      // node item ID for the system overhead which is not counted in any pod
	AllPodsItemID        = "all-pods" // not stored for now, just used for calculating the sys
)

type UIDType string

type UIDGenerator interface {
	Pod(pod *v1.Pod) UIDType
	Node() UIDType
	NodeItem(itemID string) UIDType // generate a UID for the item supposed to unique on the node
}

type Options struct {
	Filepath string
	// TODO add configs here
}

// The kubernetes UID is unique within the Pod lifecycle, so use this first. If there
// are some special scenarios in the future, such as deleting a Pod and creating a Pod
// with the same name, consider using NamespacedName as the UID.
type generator struct {
}

func (gen *generator) Pod(pod *v1.Pod) UIDType { _ = "STUB: not implemented"; return *new(UIDType) }

func (gen *generator) Node() UIDType { _ = "STUB: not implemented"; return *new(UIDType) }

func (gen *generator) NodeItem(itemID string) UIDType {
	_ = "STUB: not implemented"
	return *new(UIDType)
}

func getNodeItemUID(itemID string) UIDType { _ = "STUB: not implemented"; return *new(UIDType) }

type Result struct {
	// Use different quantile type as key, currently support "p60", "p90", "p95" "p98", "max".
	Data map[string]v1.ResourceList
}

// FIXME
// This is used for the agent's dependence on the basic data structure, and the
// basic data structure will be reconstructed later to better support testing.
type Informer interface {
	HasSynced() bool
	ListPods() []*v1.Pod
	GetNode() *v1.Node
}

func NewInformer(statesInformer statesinformer.StatesInformer) Informer {
	_ = "STUB: not implemented"
	return *new(Informer)
}

type informer struct {
	statesInformer statesinformer.StatesInformer
}

func (i *informer) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (i *informer) ListPods() []*v1.Pod { _ = "STUB: not implemented"; return nil }

func (i *informer) GetNode() *v1.Node { _ = "STUB: not implemented"; return nil }

type MetricDesc struct {
	UID UIDType
}

type MetricKey int

const (
	CPUUsage MetricKey = iota
	MemoryUsage
)

type MetricServer interface {
	GetPodMetric(desc MetricDesc, m MetricKey) (float64, error)
	GetNodeMetric(desc MetricDesc, m MetricKey) (float64, error)
}

func NewMetricServer(metricCache metriccache.MetricCache, dataInterval time.Duration) MetricServer {
	_ = "STUB: not implemented"
	return *new(MetricServer)
}

type metricServer struct {
	metricCache  metriccache.MetricCache
	dataInterval time.Duration
}

func (ms *metricServer) GetPodMetric(desc MetricDesc, m MetricKey) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ms *metricServer) GetNodeMetric(desc MetricDesc, m MetricKey) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
