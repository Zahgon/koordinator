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

package agent

import (
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/config"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/prediction"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

var (
	scheme = apiruntime.NewScheme()

	extensionControllerInitFuncs = map[string]extension.ControllerInitFunc{}
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
}

type Daemon interface {
	Run(stopCh <-chan struct{})
}

type daemon struct {
	metricAdvisor  metricsadvisor.MetricAdvisor
	statesInformer statesinformer.StatesInformer
	metricCache    metriccache.MetricCache
	qosManager     qosmanager.QOSManager
	runtimeHook    runtimehooks.RuntimeHook
	predictServer  prediction.PredictServer
	executor       resourceexecutor.ResourceUpdateExecutor

	extensionControllers []extension.Controller
}

func NewDaemon(config *config.Configuration) (Daemon, error) {
	_ = "STUB: not implemented"
	// get node name
	return *new(Daemon), nil
}

// use json for CRD clients

func (d *daemon) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// start resource executor cache

// start states informer

// wait for metric advisor sync

// start metric advisor

// wait for metric advisor sync

// start predict server

// start qos manager
