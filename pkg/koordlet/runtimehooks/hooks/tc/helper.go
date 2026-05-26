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

package tc

import (
	"k8s.io/apimachinery/pkg/util/intstr"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

func loadConfigFromNodeSlo(nodesloSpec *slov1alpha1.NodeSLOSpec) *NetQosGlobalConfig {
	_ = "STUB: not implemented"
	return nil
}

func getBandwidthVal(total uint64, intOrPercent *intstr.IntOrString) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func getBandwidthByQuantityFormat(quanityStr string) uint64 { _ = "STUB: not implemented"; return 0 }

func getBandwidthByPercentageFormat(total uint64, percentage int) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func convertToClassId(major, minor int) string { _ = "STUB: not implemented"; return "" }

// convertToHexClassId get class id in hex.
func convertToHexClassId(major, minor int) uint32 { _ = "STUB: not implemented"; return 0 }

// convertIpToHex convert ip to it's hex format
// 10.211.248.149 => 0ad3f895
func convertIpToHex(ip string) string { _ = "STUB: not implemented"; return "" }

// each ip segment takes up two hexadecimal digits, and when it does not take up all the bits,
// it needs to be filled with 0.
