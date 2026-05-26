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

package services

import (
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler"
)

const (
	servicesBaseRelativePath       = "/apis/v1/"
	pluginServicesBaseRelativePath = servicesBaseRelativePath + "plugins"
)

var once sync.Once
var engine atomic.Value

type mux interface {
	Handle(string, http.Handler)
	HandlePrefix(string, http.Handler)
}

func InstallAPIHandler(mux mux, e *Engine, sched *scheduler.Scheduler, isLeader func() bool) {
	_ = "STUB: not implemented"
	return
}

func handle(isLeader func() bool) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

type Engine struct {
	*gin.Engine

	registeredProviders map[string]struct{}
	mu                  sync.Mutex
}

func NewEngine(e *gin.Engine) *Engine { _ = "STUB: not implemented"; return nil }

func (e *Engine) RegisterPluginService(plugin fwktype.Plugin, profileName string) {
	_ = "STUB: not implemented"
	return
}

func listRegisteredServices(e *gin.Engine) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

func queryNodeInfo(sched *scheduler.Scheduler) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}
