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

package informer

import (
	"time"

	policyv1 "k8s.io/api/policy/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	storagev1 "k8s.io/api/storage/v1"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type SetupInformerFn func(informerFactory informers.SharedInformerFactory)

var setupInformers = []SetupInformerFn{
	setupCompatibleInformers,
}

func SetupCustomInformers(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

func setupCompatibleInformers(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

// Clusters below k8s v1.32 do not have resource.k8s.io/v1 APIs (ResourceClaims,
// ResourceSlices, DeviceClasses, DeviceTaintRules). When DynamicResourceAllocation
// is locked to true in k8s 1.35, we must stub these informers with a fake client
// to prevent watch errors like "resourceclaims.resource.k8s.io is forbidden".

// Versions below k8s v1.22 need to disable CSIStorageCapacity

// The k8s v1.22 version needs to enable the FeatureGate to convert v1beta1.CSIStorageCapacity to v1.CSIStorageCapacity

// Versions below k8s v1.22 need to enable the FeatureGate to convert v1beta1.PodDisruptionBudget to v1.PodDisruptionBudget

// disableDynamicResourceAllocationInformer stubs out all DRA informers with a fake client
// so they never connect to the real API server. This is needed for clusters running
// Kubernetes < 1.32 that do not have the resource.k8s.io/v1 API group.
func disableDynamicResourceAllocationInformer(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

func disableCSIStorageCapacityInformer(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

func setupCompatibleCSICapacityInformer(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

// set transform funcs after the InformerFor

func newCSIStorageCapacityInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

func storagev1beta1CSIStorageCapacityTransformer(obj interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertV1Beta1CSIStorageCapacityToV1CSIStorageCapacity(in *storagev1beta1.CSIStorageCapacity) (*storagev1.CSIStorageCapacity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func disablePodDisruptionBudgetInformer(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

func setupCompatiblePodDisruptionBudgetInformer(informerFactory informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

// set transform funcs after the InformerFor

func newPodDisruptionBudgetInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

func policyv1beta1PodDisruptionBudgetTransformer(obj interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertV1Beta1PodDisruptionBudgetToV1PodDisruptionBudget(in *policyv1beta1.PodDisruptionBudget) (*policyv1.PodDisruptionBudget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
