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

package nodeslo

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/config"
)

var _ handler.EventHandler = &SLOCfgHandlerForConfigMapEvent{}

type SLOCfgCache interface {
	GetCfgCopy() *SLOCfg
	IsCfgAvailable() bool
}

type SLOCfg struct {
	ThresholdCfgMerged   configuration.ResourceThresholdCfg `json:"thresholdCfgMerged,omitempty"`
	ResourceQOSCfgMerged configuration.ResourceQOSCfg       `json:"resourceQOSCfgMerged,omitempty"`
	CPUBurstCfgMerged    configuration.CPUBurstCfg          `json:"cpuBurstCfgMerged,omitempty"`
	SystemCfgMerged      configuration.SystemCfg            `json:"systemCfgMerged,omitempty"`
	HostAppCfgMerged     configuration.HostApplicationCfg   `json:"hostAppCfgMerged,omitempty"`
	ExtensionCfgMerged   configuration.ExtensionCfgMap      `json:"extensionCfgMerged,omitempty"` // for third-party extension
}

func (in *SLOCfg) DeepCopy() *SLOCfg { _ = "STUB: not implemented"; return nil }

type sLOCfgCache struct {
	lock sync.RWMutex
	// Config could be concurrently used by the Reconciliation and EventHandler
	sloCfg    SLOCfg
	available bool
}

func DefaultSLOCfg() SLOCfg { _ = "STUB: not implemented"; return *new(SLOCfg) }

type SLOCfgHandlerForConfigMapEvent struct {
	config.EnqueueRequestForConfigMap

	Client   client.Client
	cfgCache sLOCfgCache
	recorder record.EventRecorder
}

func NewSLOCfgHandlerForConfigMapEvent(client client.Client, initCfg SLOCfg, recorder record.EventRecorder) *SLOCfgHandlerForConfigMapEvent {
	_ = "STUB: not implemented"
	return nil
}

func (p *SLOCfgHandlerForConfigMapEvent) triggerAllNodeEnqueue(q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (p *SLOCfgHandlerForConfigMapEvent) syncNodeSLOSpecIfChanged(configMap *corev1.ConfigMap) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *SLOCfgHandlerForConfigMapEvent) syncConfig(configMap *corev1.ConfigMap) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *SLOCfgHandlerForConfigMapEvent) updateCacheIfChanged(newSLOCfg SLOCfg) bool {
	_ = "STUB: not implemented"
	return false
}

// set the available flag and never change it

func (p *SLOCfgHandlerForConfigMapEvent) GetCfgCopy() *SLOCfg {
	_ = "STUB: not implemented"
	return nil
}

func (p *SLOCfgHandlerForConfigMapEvent) IsCfgAvailable() bool {
	_ = "STUB: not implemented"
	return false
}

// if config is available, just return

// if config is not available, try to get the configmap from informer cache;
// set available if configmap is found or get not found error
