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

package qosmanager

import (
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"

	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	_ "github.com/koordinator-sh/koordinator/pkg/koordlet/metrics"
	ma "github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type QOSManager interface {
	Run(stopCh <-chan struct{}) error
}

type qosManager struct {
	options *framework.Options
	context *framework.Context
}

func NewQOSManager(cfg *framework.Config, schema *apiruntime.Scheme, kubeClient clientset.Interface, crdClient *koordclientset.Clientset, nodeName string,
	statesInformer statesinformer.StatesInformer, metricCache metriccache.MetricCache, metricAdvisorConfig *ma.Config, evictVersion string) QOSManager {
	_ = "STUB: not implemented"
	return *new(QOSManager)
}

func (r *qosManager) setup() { _ = "STUB: not implemented"; return }

func (r *qosManager) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

// minimum interval is one second.
