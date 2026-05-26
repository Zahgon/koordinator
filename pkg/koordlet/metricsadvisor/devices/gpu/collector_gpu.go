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
	"time"

	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
)

const (
	DeviceCollectorName = "GPU"
)

type gpuCollector struct {
	enabled          bool
	collectInterval  time.Duration
	gpuDeviceManager GPUDeviceManager
}

func New(opt *framework.Options) framework.DeviceCollector {
	_ = "STUB: not implemented"
	return *new(framework.DeviceCollector)
}

func (g *gpuCollector) Shutdown() { _ = "STUB: not implemented"; return }

func (g *gpuCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (g *gpuCollector) Setup(fra *framework.Context) { _ = "STUB: not implemented"; return }

func (g *gpuCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (g *gpuCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (g *gpuCollector) Infos() metriccache.Devices {
	_ = "STUB: not implemented"
	return *new(metriccache.Devices)
}

func (g *gpuCollector) GetNodeMetric() ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *gpuCollector) GetPodMetric(uid, podParentDir string, cs []corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *gpuCollector) GetContainerMetric(ContainerID, podParentDir string, c *corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GPUDeviceManager interface {
	started() bool
	deviceInfos() metriccache.Devices
	collectGPUUsage()
	getNodeGPUUsage() []metriccache.MetricSample
	getPodGPUUsage(uid, podParentDir string, cs []corev1.ContainerStatus) ([]metriccache.MetricSample, error)
	getContainerGPUUsage(containerID, podParentDir string, c *corev1.ContainerStatus) ([]metriccache.MetricSample, error)
	shutdown() error
}

type dummyDeviceManager struct{}

func (d *dummyDeviceManager) started() bool { _ = "STUB: not implemented"; return false }

func (d *dummyDeviceManager) deviceInfos() metriccache.Devices {
	_ = "STUB: not implemented"
	return *new(metriccache.Devices)
}

func (d *dummyDeviceManager) collectGPUUsage() { _ = "STUB: not implemented"; return }

func (d *dummyDeviceManager) getNodeGPUUsage() []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

func (d *dummyDeviceManager) getPodGPUUsage(uid, podParentDir string, cs []corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dummyDeviceManager) getContainerGPUUsage(containerID, podParentDir string, c *corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *dummyDeviceManager) shutdown() error { _ = "STUB: not implemented"; return nil }
