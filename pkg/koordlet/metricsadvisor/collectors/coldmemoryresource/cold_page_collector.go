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

package coldmemoryresource

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
)

const (
	CollectorName = "ColdPageCollector"
)

type nonColdPageCollector struct {
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	// check whether support kidled cold page info collector
	return *new(framework.Collector)
}

// TODO(BUPT-wxq): implement podFilter for the VM-based pods and containers

// TODO(BUPT-wxq): check kstaled cold page collector
// nonCollector does nothing

func (n *nonColdPageCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (n *nonColdPageCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (n *nonColdPageCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (n *nonColdPageCollector) Setup(c1 *framework.Context) { _ = "STUB: not implemented"; return }
