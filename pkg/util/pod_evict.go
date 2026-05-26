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
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	EvictionKind            = "Eviction"
	EvictionGroupName       = "policy"
	EvictionSubResourceName = "pods/eviction"
)

// EvictPodByVersion evicts Pods using the policy/v1 Eviction API (k8s >= 1.22).
// The v1beta1 eviction API was removed in k8s 1.25 and is no longer supported.
func EvictPodByVersion(ctx context.Context, kubernetes kubernetes.Interface, namespace, name string, opts metav1.DeleteOptions, evictVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

func FindSupportedEvictVersion(client kubernetes.Interface) (version string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func SupportEviction(client kubernetes.Interface) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
