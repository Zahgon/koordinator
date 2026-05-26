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

package resourceexecutor

import (
	"sync"

	"github.com/koordinator-sh/koordinator/pkg/util/cache"
)

var _ ResourceUpdateExecutor = &ResourceUpdateExecutorImpl{}

type ResourceUpdateExecutor interface {
	Update(cacheable bool, updater ResourceUpdater) (updated bool, err error)
	UpdateBatch(cacheable bool, updaters ...ResourceUpdater)
	// LeveledUpdateBatch is to cacheable update resources by the order of resources' level.
	// For cgroup interfaces like `cpuset.cpus` and `memory.min`, reconciliation from top to bottom should keep the
	// upper value larger/broader than the lower. Thus a Leveled updater is implemented as follows:
	// 1. update batch of cgroup resources group by cgroup interface, i.e. cgroup filename.
	// 2. update each cgroup resource by the order of layers: firstly update resources from upper to lower by merging
	//    the new value with old value; then update resources from lower to upper with the new value.
	LeveledUpdateBatch(updaters [][]ResourceUpdater)
	Run(stopCh <-chan struct{})
}

type ResourceUpdateExecutorImpl struct {
	LeveledUpdateLock sync.Mutex
	ResourceCache     *cache.Cache
	Config            *Config

	onceRun   sync.Once
	gcStarted bool
}

var singleton = &ResourceUpdateExecutorImpl{
	ResourceCache: cache.NewCacheDefault(),
	Config:        Conf,
}

func NewResourceUpdateExecutor() ResourceUpdateExecutor {
	_ = "STUB: not implemented"

	// Update updates the resources with the given cacheable attribute with the cacheable attribute directly.
	return *new(ResourceUpdateExecutor)
}

func (e *ResourceUpdateExecutorImpl) Update(cacheable bool, resource ResourceUpdater) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// UpdateBatch updates a batch of resources with the given cacheable attribute.
// TODO: merge and resolve conflicts of batch updates from multiple callers.
func (e *ResourceUpdateExecutorImpl) UpdateBatch(cacheable bool, updaters ...ResourceUpdater) {
	_ = "STUB: not implemented"
	return
}

func (e *ResourceUpdateExecutorImpl) LeveledUpdateBatch(updaters [][]ResourceUpdater) {
	_ = "STUB: not implemented"
	return
}

// skip update twice for resources specified no merge

// Run runs the ResourceUpdateExecutor.
func (e *ResourceUpdateExecutorImpl) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (e *ResourceUpdateExecutorImpl) run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (e *ResourceUpdateExecutorImpl) needUpdate(updater ResourceUpdater) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *ResourceUpdateExecutorImpl) update(updater ResourceUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

// error can be ignored

func (e *ResourceUpdateExecutorImpl) updateByCache(updater ResourceUpdater) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ResourceUpdateExecutorImpl) isUpdateErrIgnored(err error) bool {
	_ = "STUB: not implemented"
	return false
}

// NewTestResourceExecutor returns a new ResourceUpdateExecutorImpl for testing usage.
// NOTE: Please DO NOT use it except unittests.
func NewTestResourceExecutor() ResourceUpdateExecutor {
	_ = "STUB: not implemented"
	return *new(ResourceUpdateExecutor)
}
