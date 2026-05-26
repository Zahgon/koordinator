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

func (h *PodMutatingHandler) deviceResourceSpecMutatingPod(ctx context.Context, req admission.Request, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *PodMutatingHandler) mutateByDeviceResources(pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	// device resource request euqal limit, not overcommit
	return nil
}

// use gpu api first

func injectResourceContainerSpec(c *corev1.Container, s *extension.ExtendedResourceContainerSpec) {
	_ = "STUB: not implemented"
	return
}

func injectGPUShare(c *corev1.Container) { _ = "STUB: not implemented"; return }

func injectGPU(c *corev1.Container) { _ = "STUB: not implemented"; return }
