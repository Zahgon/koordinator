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

package impl

import (
	"sync"

	topov1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	topologyclientset "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/clientset/versioned"
	_ "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/clientset/versioned/scheme"
	"go.uber.org/atomic"
	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	schedv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/prediction"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	HTTPScheme  = "http"
	HTTPSScheme = "https"
)

type PluginName string

type PluginOption struct {
	config      *Config
	KubeClient  clientset.Interface
	KoordClient koordclientset.Interface
	TopoClient  topologyclientset.Interface
	NodeName    string
}

type PluginState struct {
	metricCache      metriccache.MetricCache
	callbackRunner   *callbackRunner
	informerPlugins  map[PluginName]informerPlugin
	predictorFactory prediction.PredictorFactory
}

type GetGPUDriverAndModelFunc func() (string, string)

type unhealthyGPUInfo struct {
	errCode    string
	errMessage string
}

type statesInformer struct {
	// TODO refactor device as plugin
	config       *Config
	metricsCache metriccache.MetricCache
	deviceClient schedv1alpha1.DeviceInterface
	unhealthyGPU map[string]*unhealthyGPUInfo
	gpuMutex     sync.RWMutex

	option  *PluginOption
	states  *PluginState
	started *atomic.Bool

	getGPUDriverAndModelFunc GetGPUDriverAndModelFunc
}

type informerPlugin interface {
	Setup(ctx *PluginOption, state *PluginState)
	Start(stopCh <-chan struct{})
	HasSynced() bool
}

var _ statesinformer.StatesInformer = &statesInformer{}

// TODO merge all clients into one struct
func NewStatesInformer(config *Config, kubeClient clientset.Interface, crdClient koordclientset.Interface, topologyClient topologyclientset.Interface,
	metricsCache metriccache.MetricCache, nodeName string, schedulingClient schedv1alpha1.SchedulingV1alpha1Interface, predictorFactory prediction.PredictorFactory) statesinformer.StatesInformer {
	_ = "STUB: not implemented"
	return *new(statesinformer.StatesInformer)
}

func (s *statesInformer) initInformerPlugins() { _ = "STUB: not implemented"; return }

func (s *statesInformer) setupPlugins() { _ = "STUB: not implemented"; return }

func (s *statesInformer) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

// waiting for node synced.

// check is nvml is available

// start callback runner after informers synced
// since some callbacks needs the integrated input to execute, e.g. valid pods list
// the initial callback events will not be missing since the callback channels are buffered

func (s *statesInformer) waitForSyncFunc() []cache.InformerSynced {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) startPlugins(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *statesInformer) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (s *statesInformer) GetNode() *corev1.Node { _ = "STUB: not implemented"; return nil }

func (s *statesInformer) GetNodeSLO() *slov1alpha1.NodeSLO { _ = "STUB: not implemented"; return nil }

func (s *statesInformer) GetNodeMetricSpec() *slov1alpha1.NodeMetricSpec {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) GetNodeTopo() *topov1alpha1.NodeResourceTopology {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) GetAllPods() []*statesinformer.PodMeta {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) GetVolumeName(pvcNamespace, pvcName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *statesInformer) RegisterCallbacks(rType statesinformer.RegisterType, name, description string, callbackFn statesinformer.UpdateCbFn) {
	_ = "STUB: not implemented"
	return
}
