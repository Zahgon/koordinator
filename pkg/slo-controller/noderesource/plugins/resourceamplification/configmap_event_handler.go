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

package resourceamplification

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/config"
)

const (
	ReasonResourceAmplificationConfigUnmarshalFailed = "ResourceAmplificationCfgUnmarshalFailed"
)

type cfgCache struct {
	sync.RWMutex

	config    *configuration.ResourceAmplificationCfg
	available bool
}

func DefaultResourceAmplificationCfg() *configuration.ResourceAmplificationCfg {
	_ = "STUB: not implemented"
	return nil
}

type configHandler struct {
	config.EnqueueRequestForConfigMap

	Client   ctrlclient.Client
	cache    *cfgCache
	recorder record.EventRecorder
}

func newConfigHandler(c ctrlclient.Client, initCfg *configuration.ResourceAmplificationCfg, recorder record.EventRecorder) *configHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *configHandler) IsCfgAvailable() bool { _ = "STUB: not implemented"; return false }

// if config is not available, try to get the configmap from informer cache;
// set available if configmap is found or get not found error

func (h *configHandler) GetCfgCopy() *configuration.ResourceAmplificationCfg {
	_ = "STUB: not implemented"
	return nil
}

func (h *configHandler) GetStrategyCopy(node *corev1.Node) *configuration.ResourceAmplificationStrategy {
	_ = "STUB: not implemented"
	return nil
}

// assert cache is available

// use cluster strategy

func (h *configHandler) enqueueAllNodes(q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *configHandler) syncCacheIfCfgChanged(configMap *corev1.ConfigMap) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *configHandler) syncConfig(configMap *corev1.ConfigMap) bool {
	_ = "STUB: not implemented"
	return false
}

// merge with clusterStrategy

func (h *configHandler) updateCacheIfChanged(newCfg *configuration.ResourceAmplificationCfg) bool {
	_ = "STUB: not implemented"
	return false
}
