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
	// AnnotationCPUNormalizationRatio denotes the cpu normalization ratio of the node.
	AnnotationCPUNormalizationRatio = NodeDomainPrefix + "/cpu-normalization-ratio"

	// LabelCPUNormalizationEnabled indicates whether the cpu normalization is enabled on the node.
	// If both the label and node-level CPUNormalizationStrategy is set, the label overrides the strategy.
	LabelCPUNormalizationEnabled = NodeDomainPrefix + "/cpu-normalization-enabled"

	// AnnotationCPUBasicInfo denotes the basic CPU info of the node.
	AnnotationCPUBasicInfo = NodeDomainPrefix + "/cpu-basic-info"

	// NormalizationRatioDiffEpsilon is the min difference between two cpu normalization ratios.
	NormalizationRatioDiffEpsilon = 0.01
)

// GetCPUNormalizationRatio gets the cpu normalization ratio from the node.
// It returns -1 without an error when the cpu normalization annotation is missing.
func GetCPUNormalizationRatio(node *corev1.Node) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SetCPUNormalizationRatio sets the node annotation according to the cpu-normalization-ratio.
// It returns true if the label value changes.
// NOTE: The ratio will be converted to string with the precision 2. e.g. 3.1415926 -> 3.14.
func SetCPUNormalizationRatio(node *corev1.Node, ratio float64) bool {
	_ = "STUB: not implemented"
	return false
}

func GetCPUNormalizationEnabled(node *corev1.Node) (*bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IsCPUNormalizationRatioDifferent(old, new float64) bool {
	_ = "STUB: not implemented"
	return false
}

// CPUBasicInfo describes the cpu basic features and status.
type CPUBasicInfo struct {
	CPUModel           string `json:"cpuModel,omitempty"`
	HyperThreadEnabled bool   `json:"hyperThreadEnabled,omitempty"`
	TurboEnabled       bool   `json:"turboEnabled,omitempty"`
	CatL3CbmMask       string `json:"catL3CbmMask,omitempty"`
	VendorID           string `json:"vendorID,omitempty"`
}

func (c *CPUBasicInfo) Key() string { _ = "STUB: not implemented"; return "" }

// GetCPUBasicInfo gets the cpu basic info from the node-level annotations.
// It returns nil info without an error when the cpu basic info annotation is missing.
func GetCPUBasicInfo(annotations map[string]string) (*CPUBasicInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCPUBasicInfo sets the cpu basic info at the node-level annotations.
// It returns true if the annotations changes.
func SetCPUBasicInfo(annotations map[string]string, info *CPUBasicInfo) bool {
	_ = "STUB: not implemented"
	return false
}
