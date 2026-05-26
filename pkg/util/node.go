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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

// GenerateNodeKey returns a generated key with given meta
func GenerateNodeKey(node *metav1.ObjectMeta) string { _ = "STUB: not implemented"; return "" }

// GetNodeAddress get node specified type address.
func GetNodeAddress(node *corev1.Node, addrType corev1.NodeAddressType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsNodeAddressTypeSupported determine whether addrType is a supported type.
func IsNodeAddressTypeSupported(addrType corev1.NodeAddressType) bool {
	_ = "STUB: not implemented"
	return false
}

func GetNodeReservationFromKubelet(node *corev1.Node) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func GetNodeReservationFromAnnotation(anno map[string]string) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func GetNodeReservationResources(reservation *apiext.NodeReservation) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func TrimNodeAllocatableByNodeReservation(node *corev1.Node) (trimmedAllocatable corev1.ResourceList, trimmed bool) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), false
}

// node.alloc(batch-memory) and node.alloc(batch-memory) have subtracted the reserved resources from the koord-manager,
// so we should keep the original data here.

func GetNodeAnnoReservedJson(reserved apiext.NodeReservation) string {
	_ = "STUB: not implemented"
	return ""
}

func GetNodeAllocatableBatchMilliCPU(node *corev1.Node) int64 {
	_ = "STUB: not implemented"
	// assert node != nil
	return 0
}

func GetNodeAllocatableBatchMemory(node *corev1.Node) int64 {
	_ = "STUB: not implemented"
	// assert node != nil
	return 0
}
