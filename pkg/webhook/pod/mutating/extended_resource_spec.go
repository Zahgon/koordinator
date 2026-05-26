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

package mutating

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/koordinator-sh/koordinator/apis/extension"
)

func (h *PodMutatingHandler) extendedResourceSpecMutatingPod(ctx context.Context, req admission.Request, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *PodMutatingHandler) mutateByExtendedResources(pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	// dump batch-resource of pod.spec.containers[*].resources.requests/limits into ExtendedResourceSpec{}
	return nil
}

// TODO: count init containers and pod overhead

// no requirement of specified extended resources

// compare annotation values

// if resource requirements not changed, just return

// mutate pod annotation

func getContainerExtendedResourcesRequirement(container *corev1.Container, resourceNames []corev1.ResourceName) *extension.ExtendedResourceContainerSpec {
	_ = "STUB: not implemented"
	return nil
}

// container has no requirement of specified extended resources
