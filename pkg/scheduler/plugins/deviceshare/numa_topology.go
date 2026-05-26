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
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

type NUMATopology struct {
	numNodePerSocket int
	nodes            map[int][]PCIe
	deviceToNodeID   map[schedulingv1alpha1.DeviceType]map[int32]int
}

type PCIe struct {
	PCIeIndex
	devices map[schedulingv1alpha1.DeviceType][]int
}

type PCIeIndex struct {
	socket int
	node   int
	pcie   string
}

func newNUMATopology(deviceObj *schedulingv1alpha1.Device) *NUMATopology {
	_ = "STUB: not implemented"
	return nil
}

//
// NOTE: By default, it must be assigned according to the topology,
// and the Required/Preferred strategy should be provided later.
//
