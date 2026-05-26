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

package xpu

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

const (
	DeviceCollectorName = "XPU"
)

type xpuCollector struct {
	enabled        bool
	deviceInfosDir string
}

func New(opt *framework.Options) framework.DeviceCollector {
	_ = "STUB: not implemented"
	return *new(framework.DeviceCollector)
}

func (x xpuCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (x xpuCollector) Setup(s *framework.Context) { _ = "STUB: not implemented"; return }

func (x xpuCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (x xpuCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (x xpuCollector) Shutdown() { _ = "STUB: not implemented"; return }

func (x xpuCollector) Infos() metriccache.Devices {
	_ = "STUB: not implemented"
	return *new(metriccache.Devices)
}

func (x xpuCollector) GetNodeMetric() ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x xpuCollector) GetPodMetric(uid, podParentDir string, cs []corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (x xpuCollector) GetContainerMetric(containerID, podParentDir string, c *corev1.ContainerStatus) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetXPUDevice(deviceInfosDir string) (metriccache.Devices, error) {
	_ = "STUB: not implemented"
	// default deviceInfosDir: "/var/run/koordlet/device-infos/"
	return *new(metriccache.Devices), nil
}

func readDeviceInfosFromFile(filePath string) (util.XPUDevices, error) {
	_ = "STUB: not implemented"
	// read the device info from the file
	return *new(util.XPUDevices), nil
}
