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

package config

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/koordinator-sh/koordinator/apis/configuration"
)

const (
	ReasonColocationConfigUnmarshalFailed = "ColocationCfgUnmarshalFailed"
	ReasonSLOConfigUnmarshalFailed        = "SLOCfgUnmarshalFailed"
)

var _ handler.EventHandler = &ColocationHandlerForConfigMapEvent{}

type ColocationCfgCache interface {
	GetCfgCopy() *configuration.ColocationCfg
	IsCfgAvailable() bool
	IsErrorStatus() bool
}

type colocationCfgCache struct {
	lock          sync.RWMutex
	colocationCfg configuration.ColocationCfg
	available     bool
	errorStatus   bool
}

type ColocationHandlerForConfigMapEvent struct {
	EnqueueRequestForConfigMap

	Client   client.Client
	cfgCache colocationCfgCache
	recorder record.EventRecorder
}

func NewColocationHandlerForConfigMapEvent(client client.Client, initCfg configuration.ColocationCfg, recorder record.EventRecorder) *ColocationHandlerForConfigMapEvent {
	_ = "STUB: not implemented"
	return nil
}

// syncColocationCfgIfChanged is a locked version of syncConfig
func (p *ColocationHandlerForConfigMapEvent) syncColocationCfgIfChanged(configMap *corev1.ConfigMap) bool {
	_ = "STUB: not implemented"
	// get co-location config from the configmap
	// if the configmap does not exist, use the default
	return false
}

// syncConfig syncs valid colocation config from the configmap request
func (p *ColocationHandlerForConfigMapEvent) syncConfig(configMap *corev1.ConfigMap) bool {
	_ = "STUB: not implemented"
	// get co-location config from the configmap
	// if the configmap does not exist, use the default
	return false
}

//if controller restart ,cache will unavailable, else use old cfg

// merge default cluster strategy

//if controller restart ,cache will unavailable, else use old cfg

// merge with clusterStrategy

func (p *ColocationHandlerForConfigMapEvent) updateCacheIfChanged(newCfg *configuration.ColocationCfg, errorStatus bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ColocationHandlerForConfigMapEvent) triggerAllNodeEnqueue(q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (p *ColocationHandlerForConfigMapEvent) GetCfgCopy() *configuration.ColocationCfg {
	_ = "STUB: not implemented"
	return nil
}

func (p *ColocationHandlerForConfigMapEvent) IsErrorStatus() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ColocationHandlerForConfigMapEvent) IsCfgAvailable() bool {
	_ = "STUB: not implemented"
	return false
}

// if config is available, just return

// if config is not available, try to get the configmap from informer cache;
// set available if configmap is found or get not found error

func GetConfigMapForCache(client client.Client) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	// try to get the configmap from informer cache;
	// if not found, set configmap to nil and ignore error
	return nil, nil
}
