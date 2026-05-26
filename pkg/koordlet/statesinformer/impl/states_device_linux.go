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

package impl

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

func (s *statesInformer) reportDevice() { _ = "STUB: not implemented"; return }

// if xpu devices exist, report xpu devices; others, fallback to report gpu devices from nvml

func (s *statesInformer) buildBasicDevice(node *corev1.Node) *schedulingv1alpha1.Device {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) fillGPUDevice(device *schedulingv1alpha1.Device,
	gpuDevices []schedulingv1alpha1.DeviceInfo, gpuModel string, gpuDriverVer string) {
	_ = "STUB: not implemented"
	return
}

// currently the built-in informer only supports NVIDIA GPUs

func (s *statesInformer) createDevice(device *schedulingv1alpha1.Device) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) updateDevice(device *schedulingv1alpha1.Device) error {
	_ = "STUB: not implemented"
	return nil
}

// update the device info conditions

func (s *statesInformer) buildGPUDevice() []schedulingv1alpha1.DeviceInfo {
	_ = "STUB: not implemented"
	//queryParam := generateQueryParam()
	return nil
}

func (s *statesInformer) buildRDMADevice() []schedulingv1alpha1.DeviceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *statesInformer) buildXPUDevice(xpuDevices koordletuti.XPUDevices) []schedulingv1alpha1.DeviceInfo {
	_ = "STUB: not implemented"
	return nil
}

// check resource name equal "gpu-core" or "gpu-memory" but not starts with "koordinator.sh/"

func (s *statesInformer) initGPU() bool { _ = "STUB: not implemented"; return false }

func (s *statesInformer) getGPUDriverAndModel() (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// NVIDIA Driver 470 report GPU Model with "NVIDIA " Prefix

// A100 SXM4 80GB -> A100-SXM4-80GB
// Tesla P100-PCIE-16GB -> Tesla-P100-PCIE-16GB
// Tesla V100-SXM2-16GB -> Tesla-V100-SXM2-16GB
// Tesla T4 -> Tesla-T4
// Tesla P40 -> Tesla-P40
// Tesla M40 -> Tesla-M40
// GeForce RTX 2080 Ti -> GeForce-RTX-2080-Ti
// GeForce GTX 1080 Ti -> GeForce-GTX-1080-Ti

func (s *statesInformer) gpuHealCheck(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// FIXME: there is no way to recover from the Unhealthy state.

// check status of gpus, and send unhealthy devices to the unhealthyDeviceChan channel
type gpuHealthEvent struct {
	pciBusID   string
	xidError   uint64
	errMessage string
}

func checkHealth(stopCh <-chan struct{}, devs []string, xids chan<- gpuHealthEvent) {
	_ = "STUB: not implemented"
	return
}

// http://docs.nvidia.com/deploy/xid-errors/index.html#topic_4
// Application errors: the GPU should still be healthy

// All devices are unhealthy

// check if device still exists

func (s *statesInformer) buildXPUDeviceLabels(xpuDevices koordletuti.XPUDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// use the first device's vendor and model as the label value

// If the device has P2P links, set the partition policy

func (s *statesInformer) buildXPUDeviceAnnotations(xpuDevices koordletuti.XPUDevices) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func getPartitionTableFromXPUDevices(xpus koordletuti.XPUDevices) extension.GPUPartitionTable {
	_ = "STUB: not implemented"
	return *new(extension.GPUPartitionTable)
}

func getXPUDevices(metricsCache metriccache.MetricCache) koordletuti.XPUDevices {
	_ = "STUB: not implemented"
	return *new(koordletuti.XPUDevices)
}

func getXPUDeviceTopology(xpu *koordletuti.XPUDeviceInfo) *schedulingv1alpha1.DeviceTopology {
	_ = "STUB: not implemented"
	return nil
}

func getGPUDeviceConditions(gpu *koordletuti.GPUDeviceInfo) []metav1.Condition {
	_ = "STUB: not implemented"
	return nil
}

func getXPUDeviceConditions(xpu *koordletuti.XPUDeviceInfo) []metav1.Condition {
	_ = "STUB: not implemented"
	return nil
}
