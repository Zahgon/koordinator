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

package sysreconcile

import (
	"time"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	SystemConfigReconcileName = "SystemConfigReconcile"
)

type systemConfig struct {
	reconcileInterval time.Duration
	statesInformer    statesinformer.StatesInformer
	executor          resourceexecutor.ResourceUpdateExecutor
}

var _ framework.QOSStrategy = &systemConfig{}

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (r *systemConfig) Enabled() bool { _ = "STUB: not implemented"; return false }

func (r *systemConfig) Setup(context *framework.Context) { _ = "STUB: not implemented"; return }

func (r *systemConfig) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *systemConfig) init(stopCh <-chan struct{}) {
	s.executor.Run(stopCh)
}

func (s *systemConfig) reconcile() { _ = "STUB: not implemented"; return }

func calculateMemoryConfig(strategy *slov1alpha1.SystemStrategy, nodeMemory int64) []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

//to kbytes
