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

package nodestorageinfo

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
)

const (
	CollectorName = "NodeStorageInfoCollector"
)

type nodeInfoCollector struct {
	collectInterval time.Duration
	storage         metriccache.KVStorage
	started         *atomic.Bool
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

func (n *nodeInfoCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (n *nodeInfoCollector) Setup(s *framework.Context) { _ = "STUB: not implemented"; return }

func (n *nodeInfoCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (n *nodeInfoCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (n *nodeInfoCollector) collectNodeLocalStorageInfo() { _ = "STUB: not implemented"; return }
