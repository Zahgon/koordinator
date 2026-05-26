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
	"sync"
	"time"

	"go.uber.org/atomic"
	"k8s.io/utils/clock"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	"github.com/koordinator-sh/koordinator/pkg/util/histogram"
)

var (
	// MinSampleWeight is the minimal weight of any sample (prior to including decaying factor)
	MinSampleWeight = 0.1
	// epsilon is the minimal weight kept in histograms, it should be small enough that old samples
	// (just inside MemoryAggregationWindowLength) added with MinSampleWeight are still kept
	epsilon = 0.001 * MinSampleWeight
	// DefaultHistogramBucketSizeGrowth is the default value for histogramBucketSizeGrowth.
	DefaultHistogramBucketSizeGrowth = 0.05
)

/*
PredictServer is responsible for fetching data from MetricCache, training prediction results
according to predefined models, and providing an interface for obtaining prediction results.

It is important to note that the prediction results made by PredictServer based on the captured
data are only related to the data it sees. For example, when we need to deal with cold starts,
this business logic should be processed when using the predicted data instead of being coupled
to the predictive model.

The predictive model currently provides histogram-based statistics with exponentially decaying
weights over time periods. PredictServer is responsible for storing the intermediate results of
the model and recovering when the process restarts.
*/
type PredictServer interface {
	Setup(statesinformer.StatesInformer, metriccache.MetricCache) error
	Run(stopCh <-chan struct{}) error
	HasSynced() bool
	GetPrediction(MetricDesc) (Result, error)
}

type PredictModel struct {
	CPU    histogram.Histogram
	Memory histogram.Histogram

	LastUpdated      time.Time
	LastCheckpointed time.Time
	Lock             sync.Mutex
}

type peakPredictServer struct {
	cfg          *Config
	informer     Informer
	metricServer MetricServer

	uidGenerator UIDGenerator
	models       map[UIDType]*PredictModel
	modelsLock   sync.Mutex

	clock        clock.Clock
	hasSynced    *atomic.Bool
	checkpointer Checkpointer
}

func NewPeakPredictServer(cfg *Config) PredictServer {
	_ = "STUB: not implemented"
	return *new(PredictServer)
}

func (p *peakPredictServer) Setup(statesInformer statesinformer.StatesInformer, metricCache metriccache.MetricCache) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *peakPredictServer) Run(stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// remove unknown checkpoints before starting to work

func (p *peakPredictServer) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (p *peakPredictServer) training() {
	_ = "STUB: not implemented"
	// get pod metrics
	// 1. list pods, update models
	return
}

// count the node-level usages of different priority classes and system

// update the pod model

// update the node priority metric

// count all pods metric

// 2. get node, update models

// 3. update node priority models

// reset the priority usage

// 4. update system model

// From 0.05 to 1024 cores, maintain the bucket of the CPU histogram at a rate of 5%
func (p *peakPredictServer) defaultCPUHistogram() histogram.Histogram {
	_ = "STUB: not implemented"
	return *new(histogram.Histogram)
}

// From 10M to 2T, maintain the bucket of the Memory histogram at a rate of 5%
func (p *peakPredictServer) defaultMemoryHistogram() histogram.Histogram {
	_ = "STUB: not implemented"
	return *new(histogram.Histogram)
}

func (p *peakPredictServer) updateModel(uid UIDType, cpu, memory float64) {
	_ = "STUB: not implemented"
	return
}

// TODO Add adjusted weights

func (p *peakPredictServer) GetPrediction(metric MetricDesc) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

//

func (p *peakPredictServer) gcModels() { _ = "STUB: not implemented"; return }

// do the io operations out of lock

func (p *peakPredictServer) doCheckpoint() { _ = "STUB: not implemented"; return }

// Sort models and keys by LastCheckpointed time

func (p *peakPredictServer) restoreModels() (unknownUIDs []UIDType) {
	_ = "STUB: not implemented"
	return nil
}

// pods checkpoints

// node checkpoint

// node items checkpoints (priority classes)

type PredictMetric struct {
	LastCPUUsage    float64
	LastMemoryUsage float64
}

type NodeItemsUsage struct {
	// MetricMap maps an item to its predict metric.
	// e.g.
	//      PriorityProd -> {6.2 cores, 20 GiB}
	//      sys          -> {0.1 cores, 4 GiB}
	MetricMap map[string]*PredictMetric
}

func NewNodeItemUsage() *NodeItemsUsage { _ = "STUB: not implemented"; return nil }

func (m *NodeItemsUsage) AddMetric(itemID string, cpuUsage, memoryUsage float64) {
	_ = "STUB: not implemented"
	return
}

func (m *NodeItemsUsage) GetMetric(itemID string) (*PredictMetric, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
