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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

const (
	// batch resource can be shared with other allocators such as Hadoop YARN
	// record origin batch allocatable on node for calculating the batch allocatable of K8s and YARN, e.g.
	// k8s_batch_allocatable = origin_batch_allocatable - yarn_batch_requested
	// yarn_allocatable = origin_batch_allocatable - k8s_batch_requested
	NodeOriginExtendedAllocatableAnnotationKey = "node.koordinator.sh/originExtendedAllocatable"

	// record (batch) allocations of other schedulers such as YARN, which should be excluded before updating node extended resource
	NodeThirdPartyAllocationsAnnotationKey = "node.koordinator.sh/thirdPartyAllocations"
)

type OriginAllocatable struct {
	Resources corev1.ResourceList `json:"resources,omitempty"`
}

func GetOriginExtendedAllocatable(annotations map[string]string) (*OriginAllocatable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetOriginExtendedAllocatableRes(annotations map[string]string, extendedAllocatable corev1.ResourceList) error {
	_ = "STUB: not implemented"
	return nil
}

type ThirdPartyAllocations struct {
	Allocations []ThirdPartyAllocation `json:"allocations,omitempty"`
}

type ThirdPartyAllocation struct {
	Name      string               `json:"name"`
	Priority  apiext.PriorityClass `json:"priority"`
	Resources corev1.ResourceList  `json:"resources,omitempty"`
}

func GetThirdPartyAllocations(annotations map[string]string) (*ThirdPartyAllocations, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetThirdPartyAllocatedResByPriority(annotations map[string]string, priority apiext.PriorityClass) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func SetThirdPartyAllocation(annotations map[string]string, name string, priority apiext.PriorityClass,
	resource corev1.ResourceList) error {
	_ = "STUB: not implemented"
	// parse or init old allocations
	return nil
}

// create or update old alloc

// update allocation string
