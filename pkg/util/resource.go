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

package util

import (
	"github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	NodeZoneType = "Node"
)

func NewZeroResourceList() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func MinQuant(quant1, quant2 resource.Quantity) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// MultiplyMilliQuant scales quantity by factor
func MultiplyMilliQuant(quant resource.Quantity, factor float64) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// MultiplyQuant scales quantity by factor
func MultiplyQuant(quant resource.Quantity, factor float64) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// MinResourceList returns the result of Min(a, b) for each named resource, corresponding to the quotav1.Max().
// It should be semantically equivalent to the result of `quotav1.Subtract(quotav1.Add(a, b), quotav1.Max(a, b))`.
//
// e.g.
//
//	a = {"cpu": "10", "memory": "20Gi"}, b = {"cpu": "6", "memory": "24Gi", "nvidia.com/gpu": "2"}
//	=> {"cpu": "6", "memory": "20Gi"}
func MinResourceList(a corev1.ResourceList, b corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// IsResourceListEqual checks if the two resource lists are numerically equivalent.
func IsResourceListEqual(a corev1.ResourceList, b corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

// IsResourceListEqualIgnoreZeroValues checks if the two resource lists are numerically equivalent.
// NOTE: Resource name with a zero value will be ignored in comparison.
// e.g. a = {"cpu": "10", "memory": "0"}, b = {"cpu": "10"} => true
func IsResourceListEqualIgnoreZeroValues(a, b corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

// IsResourceDiff returns whether the new resource has big enough difference with the old one or not
func IsResourceDiff(old, new corev1.ResourceList, resourceName corev1.ResourceName, diffThreshold float64) bool {
	_ = "STUB: not implemented"
	return false
}

// not equal for both are zero

func IsQuantityDiff(old, new resource.Quantity, diffThreshold float64) bool {
	_ = "STUB: not implemented"
	return false
}

// use multiplication and larger than in case the oldMilli is zero

func QuantityPtr(q resource.Quantity) *resource.Quantity {
	_ = "STUB: not implemented"

	// GenNodeZoneName generates the zone name according to the NUMA node ID.
	return nil
}

func GenNodeZoneName(nodeID int) string { _ = "STUB: not implemented"; return "" }

// ZoneListToZoneResourceList transforms the zone list into a map from zone name to its resource list.
// It supposes the capacity, allocatable, available of a ResourceInfo is the same.
func ZoneListToZoneResourceList(zoneList v1alpha1.ZoneList) map[string]corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

// ZoneResources holds both capacity and allocatable resources for a zone.
type ZoneResources struct {
	Capacity    corev1.ResourceList
	Allocatable corev1.ResourceList
}

// ZoneResourceListToZoneList converts zone resource list to zone list.
// When resourceList contains only capacity, allocatable and available will be set to capacity.
// When resourceList contains both capacity and allocatable, they will be set accordingly.
func ZoneResourceListToZoneList(zoneResourceList map[string]corev1.ResourceList) v1alpha1.ZoneList {
	_ = "STUB: not implemented"
	return *new(v1alpha1.ZoneList)
}

// ZoneResourcesToZoneList converts zone resources (with capacity and allocatable) to zone list.
// This function properly handles the case where capacity >= allocatable.
func ZoneResourcesToZoneList(zoneResources map[string]ZoneResources) v1alpha1.ZoneList {
	_ = "STUB: not implemented"
	return *new(v1alpha1.ZoneList)
}

// Merge capacity and allocatable resource names

// If capacity is not set, use allocatable; if allocatable is not set, use capacity

// Available should match allocatable

func TrimDifferentZone(a, b v1alpha1.ZoneList) v1alpha1.ZoneList {
	_ = "STUB: not implemented"
	return *new(v1alpha1.ZoneList)
}

// MergeZoneList merges ZoneList b override ZoneList a.
func MergeZoneList(a, b v1alpha1.ZoneList) v1alpha1.ZoneList {
	_ = "STUB: not implemented"
	return *new(v1alpha1.ZoneList)
}

func IsZoneListResourceEqual(a, b v1alpha1.ZoneList, resourceNames ...corev1.ResourceName) bool {
	_ = "STUB: not implemented"
	return false
}

// keep B as the larger one

// different resource name

// both have no resource

// keep B as the larger one

// if resourceNames not specified, compare each resources in zone B

// current resource has different quantity

// LessThanOrEqualCompletely is different quotav1.LessThanOrEqual. It will compare non-exist value in b
func LessThanOrEqualCompletely(a corev1.ResourceList, b corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}
