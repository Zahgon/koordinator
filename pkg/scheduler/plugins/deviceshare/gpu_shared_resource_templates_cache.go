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

package deviceshare

import (
	"sync"

	corev1 "k8s.io/api/core/v1"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

type gpuSharedResourceTemplatesCache struct {
	lock sync.RWMutex
	// gpuSharedResourceTemplatesInfos stores GPUSharedResourceTemplates for each model of GPU which has it.
	gpuSharedResourceTemplatesInfos map[string]apiext.GPUSharedResourceTemplates
}

func newGPUSharedResourceTemplatesCache() *gpuSharedResourceTemplatesCache {
	_ = "STUB: not implemented"
	// no need to make infos map because it would be directly initialized from configmap data
	return nil
}

func (c *gpuSharedResourceTemplatesCache) findMatchedTemplates(resources corev1.ResourceList, strict bool) map[string]apiext.GPUSharedResourceTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (c *gpuSharedResourceTemplatesCache) setTemplatesInfos(infos map[string]apiext.GPUSharedResourceTemplates) {
	_ = "STUB: not implemented"
	return
}

func (c *gpuSharedResourceTemplatesCache) setTemplatesInfosFromConfigMap(cm *corev1.ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

func buildGPUSharedResourceTemplatesKey(vendor, model string) string {
	_ = "STUB: not implemented"
	return ""
}
