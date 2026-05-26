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

package resource

import (
	"regexp"

	"github.com/jaypipes/ghw/pkg/pci"

	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

var (
	devRegexp                   = regexp.MustCompile(`^/dev/davinci(\d+)$`)
	HUAWEIProductExcludeNPUList = []string{"1710", "1711"}
)

// IsNPUDevice only consider 910B/910B3/310P3
func IsNPUDevice(device *pci.Device) bool { _ = "STUB: not implemented"; return false }

// The HUAWEI bridge device will also be detected as HUAWEI vendor

func IsXPUDevice(device *pci.Device) bool { _ = "STUB: not implemented"; return false }

func IsMLUDevice(device *pci.Device) bool { _ = "STUB: not implemented"; return false }

func IsMXDevice(device *pci.Device) bool { _ = "STUB: not implemented"; return false }

// GetXPUName exec cmd in chroot to get xpu name
func GetXPUName(deviceType string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// npu-smi info -m

// keep the format consistent with NVIDIA

// ixsmi --query-gpu=gpu_name --format=csv,noheader --id=0 | sed -e 's/ /-/g'

// mx-smi --show-hwinfo | grep -i "Model Name" | awk '{print $4}' | head -1

// GetXPUCount exec cmd in chroot to get xpu count
func GetXPUCount(deviceType string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// nvidia-smi --list-gpus

// npu-smi info -l

// ixsmi -L | wc -l

// cnmon -l | grep UUID | wc -l

// mx-smi --show-hwinfo | grep -i "Model Name" | wc -l

// GetXPUMemory exec cmd in chroot to get xpu memory, return Mi(K8s resource unit)
func GetXPUMemory(deviceType string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// nvidia-smi --query-gpu=memory.total --format=csv,noheader --id=0

// npu-smi -t memory -i 1

// for 310P3

// ixsmi --id=0 --query-gpu=memory.total --format=csv,noheader | sed -e 's/ //g'

// mx-smi --show-hwinfo | grep -i "Memory Capacity" | awk '{print $4}' | head -1

func GetXPUCards(deviceType string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// npu-smi info -l

func GetXPUMinors(deviceType string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// ls /dev/davinci*

// GetAiCore exec cmd in chroot to get ai core count from npu
func GetAiCore() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// npu-smi info -t common -i 0

// GetAiCpu exec cmd in chroot to get ai cpu count from npu
func GetAiCpu(npuID string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// npu-smi info -t cpu-num-cfg -i 1

func GetDeviceInfo(deviceType, index string) (koordletuti.DeviceTopology, koordletuti.DeviceStatus, string, error) {
	_ = "STUB: not implemented"
	return *new(koordletuti.DeviceTopology), *new(koordletuti.DeviceStatus), "", nil
}

//TODO jess add PCIE switch

// 使用正则表达式精确匹配 "Device : ..."
// 支持 Device 后面有任意空格，冒号前后有任意空格

// 如果是数字

// TODO jess add PCIE switch

// npu-smi info -t board -i 1

// Placeholder UUID, replace with actual if available

// npu-smi info -t topo -i 1

// TODO for 910B need parse info

// xpu_smi -m -d /dev/xpu*

// mx-smi --list | awk '{print $6}'

// mx-smi --show-hwinfo -i 0 | grep -i "GPU#0" | awk '{print $3}'

// mx-smi  --count-ecc -i 0

func isNumericRegex(s string) bool { _ = "STUB: not implemented"; return false }
