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

package framework

import (
	"sync"
	"time"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
)

type Context struct {
	DeviceCollectors map[string]DeviceCollector
	Collectors       map[string]Collector
	State            *SharedState
}

func DeviceCollectorsStarted(devices map[string]DeviceCollector) bool {
	_ = "STUB: not implemented"
	return false
}

func CollectorsHasStarted(collectors map[string]Collector) bool {
	_ = "STUB: not implemented"
	return false
}

type CPUStat struct {
	// TODO check CPUTick or CPUUsage can be abandoned
	CPUTick   uint64
	CPUUsage  uint64
	Timestamp time.Time
}

// SharedState is for sharing infos across collectors, for example the system resource collector use the result of
// pod and node resource collector for calculating system usage
type SharedState struct {
	LatestMetric
}

func NewSharedState() *SharedState { _ = "STUB: not implemented"; return nil }

type LatestMetric struct {
	nodeMutex  sync.RWMutex
	nodeCPU    *metriccache.Point
	nodeMemory *metriccache.Point

	podMutex              sync.RWMutex
	podsCPUByCollector    map[string]metriccache.Point
	podsMemoryByCollector map[string]metriccache.Point

	hostAppMutex  sync.RWMutex
	hostAppCPU    *metriccache.Point
	hostAppMemory *metriccache.Point
}

func (r *SharedState) UpdateNodeUsage(cpu, memory metriccache.Point) {
	_ = "STUB: not implemented"
	return
}

func (r *SharedState) UpdatePodUsage(collectorName string, cpu, memory metriccache.Point) {
	_ = "STUB: not implemented"
	return
}

func (r *SharedState) UpdateHostAppUsage(cpu, memory metriccache.Point) {
	_ = "STUB: not implemented"
	return
}

func (r *SharedState) GetNodeUsage() (cpu, memory *metriccache.Point) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SharedState) GetHostAppUsage() (cpu, memory *metriccache.Point) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SharedState) GetPodsUsageByCollector() (cpu, memory map[string]metriccache.Point) {
	_ = "STUB: not implemented"
	return nil, nil
}
