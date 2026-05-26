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

package transformer

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

var podTransformers = []func(pod *corev1.Pod){
	TransformDeprecatedBatchResources,
	TransformDeprecatedDeviceResources,
	TransformReplaceResources,
}

var podTransformerFactories = []func() func(pod *corev1.Pod){
	TransformKoordPriorityClassFunc,
	TransformKoordPreemptionPolicyFunc,
	TransformSchedulerName,
}

func InstallPodTransformer(informer cache.SharedIndexInformer) { _ = "STUB: not implemented"; return }

func TransformPodFactory() cache.TransformFunc {
	_ = "STUB: not implemented"
	return *new(cache.TransformFunc)
}

func TransformKoordPriorityClassFunc() func(pod *corev1.Pod) { _ = "STUB: not implemented"; return nil }

func TransformKoordPreemptionPolicyFunc() func(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return nil
}

func TransformDeprecatedBatchResources(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func TransformDeprecatedDeviceResources(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func transformDeviceAllocations(deviceAllocations apiext.DeviceAllocations) bool {
	_ = "STUB: not implemented"
	return false
}

func transformDeprecatedResources(pod *corev1.Pod, resourceNames map[corev1.ResourceName]corev1.ResourceName) {
	_ = "STUB: not implemented"
	return
}

func replaceAndEraseResource(resourceList corev1.ResourceList, from, to corev1.ResourceName) bool {
	_ = "STUB: not implemented"
	return false
}

func replaceAndEraseWithResourcesMapper(resList corev1.ResourceList, mapper map[corev1.ResourceName]corev1.ResourceName) bool {
	_ = "STUB: not implemented"
	return false
}

func TransformSchedulerName() func(pod *corev1.Pod) { _ = "STUB: not implemented"; return nil }

// TransformReplaceResources transforms pod resources according to the replace-resources annotation.
func TransformReplaceResources(pod *corev1.Pod) { _ = "STUB: not implemented"; return }
