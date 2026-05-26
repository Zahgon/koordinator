//go:build linux
// +build linux

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

package gpu

import (
	"sync"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"go.uber.org/atomic"
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
)

type gpuDeviceManager struct {
	sync.RWMutex
	deviceCount      int
	devices          []*device
	collectTime      time.Time
	start            *atomic.Bool
	processesMetrics map[uint32][]*rawGPUMetric
}

type rawGPUMetric struct {
	SMUtil     uint32 // current utilization rate for the device
	MemoryUsed uint64
}

type device struct {
	Minor       int32 // index starting from 0
	DeviceUUID  string
	MemoryTotal uint64
	NodeID      int32
	PCIE        string
	BusID       string
	Device      nvml.Device
}

// initGPUDeviceManager will not retry if init fails,
func initGPUDeviceManager() GPUDeviceManager {
	_ = "STUB: not implemented"
	return *new(GPUDeviceManager)
}

func (g *gpuDeviceManager) shutdown() error { _ = "STUB: not implemented"; return nil }

func (g *gpuDeviceManager) initGPUData() error { _ = "STUB: not implemented"; return nil }

func (g *gpuDeviceManager) deviceInfos() metriccache.Devices {
	_ = "STUB: not implemented"
	return *new(metriccache.Devices)
}

func (g *gpuDeviceManager) getNodeGPUUsage() []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

func (g *gpuDeviceManager) getPodOrContainerTotalGPUUsageOfPIDs(id string, isPodID bool, pids []uint32) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

func (g *gpuDeviceManager) getPodGPUUsage(uid, podParentDir string, cs []corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *gpuDeviceManager) getContainerGPUUsage(containerID, podParentDir string, c *corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *gpuDeviceManager) collectGPUUsage() { _ = "STUB: not implemented"; return }

// Sort by pid.

// pid not exist.
// init processes gpu metric array.

func (g *gpuDeviceManager) started() bool { _ = "STUB: not implemented"; return false }

func buildMetricSample(mr metriccache.MetricResource, properties map[metriccache.MetricProperty]string, t time.Time, val float64) metriccache.MetricSample {
	_ = "STUB: not implemented"
	return *new(metriccache.MetricSample)
}
