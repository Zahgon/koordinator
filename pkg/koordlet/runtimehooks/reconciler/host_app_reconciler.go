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

package reconciler

import (
	"sync"
	"time"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

var globalHostAppReconcilers = struct {
	hostApps []*hostAppReconciler
}{}

type hostAppReconciler struct {
	resourceFile system.Resource
	description  string
	fn           reconcileFunc
}

func RegisterHostAppReconciler(resource system.Resource, description string, fn reconcileFunc, opt *ReconcilerOption) {
	_ = "STUB: not implemented"
	return
}

type ReconcilerOption struct {
	// TODO mv filter and condition
}

type hostReconciler struct {
	appMutex          sync.RWMutex
	hostAppMap        map[string]*slov1alpha1.HostApplicationSpec
	appUpdated        chan struct{}
	executor          resourceexecutor.ResourceUpdateExecutor
	reconcileInterval time.Duration
}

func NewHostAppReconciler(ctx Context) Reconciler {
	_ = "STUB: not implemented"
	return *new(Reconciler)
}

func (r *hostReconciler) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

func (r *hostReconciler) reconcile(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (r *hostReconciler) appRefreshCallback(t statesinformer.RegisterType, mergedNodeSLOSpecIf interface{},
	target *statesinformer.CallbackTarget) {
	_ = "STUB: not implemented"
	return
}

func (r *hostReconciler) parseHostApp(hostApps []slov1alpha1.HostApplicationSpec) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *hostReconciler) updateHostApp(hostAppMap map[string]*slov1alpha1.HostApplicationSpec) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *hostReconciler) getHostApps() map[string]*slov1alpha1.HostApplicationSpec {
	_ = "STUB: not implemented"
	return nil
}

func (r *hostReconciler) doHostAppCgroup(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }
