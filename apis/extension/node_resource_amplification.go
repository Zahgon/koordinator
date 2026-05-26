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

package extension

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	// AnnotationNodeResourceAmplificationRatio denotes the resource amplification ratio of the node.
	AnnotationNodeResourceAmplificationRatio = NodeDomainPrefix + "/resource-amplification-ratio"

	// AnnotationNodeRawAllocatable denotes the un-amplified raw allocatable of the node.
	AnnotationNodeRawAllocatable = NodeDomainPrefix + "/raw-allocatable"
)

// Ratio is a float64 wrapper which will always be json marshalled with precision 2.
type Ratio float64

func (f Ratio) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetNodeResourceAmplificationRatios gets the resource amplification ratios of node from annotations.
func GetNodeResourceAmplificationRatios(annotations map[string]string) (map[corev1.ResourceName]Ratio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeResourceAmplificationRatio gets the amplification ratio of a specific resource of node from annotations.
// It returns -1 without an error when the amplification ratio is not set for this resource.
func GetNodeResourceAmplificationRatio(annotations map[string]string, resource corev1.ResourceName) (Ratio, error) {
	_ = "STUB: not implemented"
	return *new(Ratio), nil
}

// SetNodeResourceAmplificationRatios sets the node annotation according to the resource amplification ratios.
// NOTE: The ratio will be converted to string with the precision 2. e.g. 3.1415926 -> 3.14.
func SetNodeResourceAmplificationRatios(node *corev1.Node, ratios map[corev1.ResourceName]Ratio) {
	_ = "STUB: not implemented"
	return
}

// SetNodeResourceAmplificationRatio sets the amplification ratio of a specific resource of the node.
// It returns true if the ratio changes.
// NOTE: The ratio will be converted to string with the precision 2. e.g. 3.1415926 -> 3.14.
func SetNodeResourceAmplificationRatio(node *corev1.Node, resource corev1.ResourceName, ratio Ratio) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// HasNodeRawAllocatable checks if the node has raw allocatable annotation.
func HasNodeRawAllocatable(annotations map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetNodeRawAllocatable gets the raw allocatable of node from annotations.
func GetNodeRawAllocatable(annotations map[string]string) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

// SetNodeRawAllocatable sets the node annotation according to the raw allocatable.
func SetNodeRawAllocatable(node *corev1.Node, allocatable corev1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func AmplifyResourceList(requests corev1.ResourceList, amplificationRatios map[corev1.ResourceName]Ratio, resourceNames ...corev1.ResourceName) {
	_ = "STUB: not implemented"
	return
}

func Amplify(origin int64, ratio Ratio) int64 { _ = "STUB: not implemented"; return 0 }
