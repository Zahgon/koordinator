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

package deviceshare

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

// deviceResources is used to present resources per device.
// we use the minor of device as key
// "0": {koordinator.sh/gpu-core:100, koordinator.sh/gpu-memory-ratio:100, koordinator.sh/gpu-memory: 16GB}
// "1": {koordinator.sh/gpu-core:100, koordinator.sh/gpu-memory-ratio:100, koordinator.sh/gpu-memory: 16GB}
type deviceResources map[int]corev1.ResourceList

func (r deviceResources) DeepCopy() deviceResources {
	_ = "STUB: not implemented"
	return *new(deviceResources)
}

func (r deviceResources) append(in deviceResources, hintMinors sets.Int) {
	_ = "STUB: not implemented"
	return
}

func (r deviceResources) subtract(in deviceResources, withNonNegativeResult bool) {
	_ = "STUB: not implemented"
	return
}

func (r deviceResources) isZero() bool { _ = "STUB: not implemented"; return false }

func copyDeviceResources(m map[schedulingv1alpha1.DeviceType]deviceResources) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

func subtractAllocated(m, allocated map[schedulingv1alpha1.DeviceType]deviceResources, withNonNegativeResult bool) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

func appendAllocated(m map[schedulingv1alpha1.DeviceType]deviceResources, allocatedList ...map[schedulingv1alpha1.DeviceType]deviceResources) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

func appendAllocatedByHints(hints map[schedulingv1alpha1.DeviceType]sets.Int, m map[schedulingv1alpha1.DeviceType]deviceResources, allocatedList ...map[schedulingv1alpha1.DeviceType]deviceResources) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

func newDeviceMinorMap(m map[schedulingv1alpha1.DeviceType]deviceResources) map[schedulingv1alpha1.DeviceType]sets.Int {
	_ = "STUB: not implemented"
	return nil
}

type deviceResourceMinorPair struct {
	preferred bool
	minor     int
	resources corev1.ResourceList
	score     int64
}

func scoreDevices(podRequest corev1.ResourceList, totalResources, freeResources deviceResources, allocationScorer *resourceAllocationScorer) []deviceResourceMinorPair {
	_ = "STUB: not implemented"
	return nil
}

// sortDeviceResourcesByMinor sort devices by preferred than by score than by minor.
func sortDeviceResourcesByMinor(r []deviceResourceMinorPair, preferred sets.Int) []deviceResourceMinorPair {
	_ = "STUB: not implemented"
	return nil
}

// sortDeviceResourcesByPreferredPCIe sort devices by preferred PCIes than by score than by minor.
// Note that it uses a round-robin sort for devices of preferred PCIes, in other words,
// it would sort the 1st device of each PCIe than 2nd, and so on, to make sure that every preferred PCIe gets a device with best efforts.
func sortDeviceResourcesByPreferredPCIe(r []deviceResourceMinorPair, preferredPCIe sets.String, deviceInfos map[int]*schedulingv1alpha1.DeviceInfo) []deviceResourceMinorPair {
	_ = "STUB: not implemented"
	// overall sorting
	return nil
}

// group devices by pcie

// round-robin sorting by group
