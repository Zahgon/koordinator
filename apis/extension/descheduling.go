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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// AnnotationEvictionCost indicates the eviction cost. It can be used to set to an int32.
	// Although the K8s community has [Pod Deletion Cost #2255](https://github.com/kubernetes/enhancements/issues/2255),
	// it is not a general mechanism. To avoid conflicts with components that use `Pod Deletion Cost`,
	// users can individually mark the eviction cost for Pods.
	// The implicit eviction cost for pods that don't set the annotation is 0, negative values are permitted.
	// If set the cost with `math.MaxInt32`, it means the Pod will not be evicted.
	// Pods with lower eviction cost are preferred to be evicted before pods with higher eviction cost.
	// If a batch of Pods to be evicted have the same priority, they will be sorted by cost,
	// and the Pod with the smallest cost will be evicted.
	AnnotationEvictionCost = SchedulingDomainPrefix + "/eviction-cost"
)

const (
	// AnnotationSoftEviction indicates custom eviction. It can be used to set to an "true".
	AnnotationSoftEviction = SchedulingDomainPrefix + "/soft-eviction"
)

type SoftEvictionSpec struct {
	// Timestamp indicates time when custom eviction occurs . It can be used to set a second timestamp.
	Timestamp *metav1.Time `json:"timestamp,omitempty"`
	// DeleteOptions indicates the options to delete the pod.
	DeleteOptions *metav1.DeleteOptions `json:"deleteOptions,omitempty"`
	// Initiator indicates the initiator of the eviction.
	Initiator string `json:"initiator,omitempty"`
	// Reason indicates reason for eviction.
	Reason string `json:"reason,omitempty"`
}

func GetSoftEvictionSpec(annotations map[string]string) (*SoftEvictionSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetEvictionCost(annotations map[string]string) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// values that start with plus sign (e.g, "+10") or leading zeros (e.g., "008") are not valid.

// make sure we default to 0 on error.

func validFirstDigit(str string) bool { _ = "STUB: not implemented"; return false }
