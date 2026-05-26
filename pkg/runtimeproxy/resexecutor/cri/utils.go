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

package cri

import (
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"

	"github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
)

func transferToKoordResources(r *runtimeapi.LinuxContainerResources) *v1alpha1.LinuxContainerResources {
	_ = "STUB: not implemented"
	return nil
}

func transferToCRIResources(r *v1alpha1.LinuxContainerResources) *runtimeapi.LinuxContainerResources {
	_ = "STUB: not implemented"
	return nil
}

func updateResource(a, b *v1alpha1.LinuxContainerResources) *v1alpha1.LinuxContainerResources {
	_ = "STUB: not implemented"
	return nil
}

// -1 is valid

// updateResourceByUpdateContainerResourceRequest updates resources in cache by UpdateContainerResource request.
// updateResourceByUpdateContainerResourceRequest will omit OomScoreAdj.
//
// Normally kubelet won't send UpdateContainerResource request, so if some components want to send it and want to update OomScoreAdj,
// please use hook to achieve it.
func updateResourceByUpdateContainerResourceRequest(a, b *v1alpha1.LinuxContainerResources) *v1alpha1.LinuxContainerResources {
	_ = "STUB: not implemented"
	return nil
}

// -1 is valid

func transferToKoordContainerEnvs(envs []*runtimeapi.KeyValue) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func transferToCRIContainerEnvs(envs map[string]string) []*runtimeapi.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func IsKeyValExistInLabels(labels map[string]string, key, val string) bool {
	_ = "STUB: not implemented"
	return false
}
