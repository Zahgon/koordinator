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

package runtimehooks

import (
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/nri"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/proxyserver"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/reconciler"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type HookPlugin interface {
	Register(op hooks.Options)
}

type RuntimeHook interface {
	Run(stopCh <-chan struct{}) error
}

type runtimeHook struct {
	statesInformer    statesinformer.StatesInformer
	server            proxyserver.Server
	nriServer         nri.Server
	reconciler        reconciler.Reconciler
	hostAppReconciler reconciler.Reconciler
	reader            resourceexecutor.CgroupReader
	executor          resourceexecutor.ResourceUpdateExecutor
}

func (r *runtimeHook) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

// if NRI is not enabled or container runtime not support NRI, we just skip NRI server start

func NewRuntimeHook(si statesinformer.StatesInformer, cfg *Config, schema *apiruntime.Scheme, kubeClient clientset.Interface, nodeName string) (RuntimeHook, error) {
	_ = "STUB: not implemented"
	return *new(RuntimeHook), nil
}

func registerPlugins(op hooks.Options) { _ = "STUB: not implemented"; return }

func getDisableStagesMap(stagesSlice []string) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
