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

package validating

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func (h *PodValidatingHandler) deviceResourceValidatingPod(ctx context.Context, req admission.Request) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

func validateDeviceResource(pod *corev1.Pod) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// use gpu api first

// gpu resource will no longer exist

func validatePercentageResource(q resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func validateZeroResource(q resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func validateMultiple(a, b resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func validateLess100(q resource.Quantity) bool { _ = "STUB: not implemented"; return false }

func validateGPU(c *corev1.Container) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateGPUShare(c *corev1.Container) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}
