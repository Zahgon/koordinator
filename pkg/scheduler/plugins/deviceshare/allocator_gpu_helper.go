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

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

var (
	GPUPartitionIndexOfNVIDIAHopper = GPUPartitionIndexer{
		1: []*PartitionsOfAllocationScore{
			{
				Partitions: []*apiext.GPUPartition{
					{
						Minors:          []int{0},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{0}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{1},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{1}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{2},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{2}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{3},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{3}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{4},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{4}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{5},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{5}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{6},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{6}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{7},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{7}),
						AllocationScore: 1,
					},
				},
				AllocationScore: 1,
			},
		},
		2: []*PartitionsOfAllocationScore{
			{
				Partitions: []*apiext.GPUPartition{
					{
						Minors:          []int{0, 1},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{0, 1}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{2, 3},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{2, 3}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{4, 5},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{4, 5}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{6, 7},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{6, 7}),
						AllocationScore: 1,
					},
				},
				AllocationScore: 1,
			},
		},
		4: []*PartitionsOfAllocationScore{
			{
				Partitions: []*apiext.GPUPartition{
					{
						Minors:          []int{0, 1, 2, 3},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{0, 1, 2, 3}),
						AllocationScore: 1,
					},
					{
						Minors:          []int{4, 5, 6, 7},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{4, 5, 6, 7}),
						AllocationScore: 1,
					},
				},
				AllocationScore: 1,
			},
		},
		8: []*PartitionsOfAllocationScore{
			{
				Partitions: []*apiext.GPUPartition{
					{
						Minors:          []int{0, 1, 2, 3, 4, 5, 6, 7},
						GPULinkType:     apiext.GPUNVLink,
						MinorsHash:      hashMinors([]int{0, 1, 2, 3, 4, 5, 6, 7}),
						AllocationScore: 1,
					},
				},
				AllocationScore: 1,
			},
		},
	}

	GetDesignatedGPUPartitionIndexer = func(node *corev1.Node) (GPUPartitionIndexer, bool) {
		var partitionIndexer GPUPartitionIndexer
		partitionPolicy := apiext.GetGPUPartitionPolicy(node)
		vendor := node.Labels[apiext.LabelGPUVendor]
		// treat empty vendor as nvidia for compatibility
		if vendor == "" || vendor == apiext.GPUVendorNVIDIA {
			switch node.Labels[apiext.LabelGPUModel] {
			case "H100", "H800", "H20":
				partitionIndexer = GPUPartitionIndexOfNVIDIAHopper
			}
		}
		return partitionIndexer, partitionPolicy == apiext.GPUPartitionPolicyHonor
	}
)

func GetGPUPartitionIndexer(table apiext.GPUPartitionTable) GPUPartitionIndexer {
	_ = "STUB: not implemented"
	return *new(GPUPartitionIndexer)
}

func GetGPUTopologyScope(deviceInfos []*schedulingv1alpha1.DeviceInfo, nodeDeviceResources deviceResources) *GPUTopologyScope {
	_ = "STUB: not implemented"
	return nil
}

func getMinorsListFromMap(resources deviceResources) []int { _ = "STUB: not implemented"; return nil }
