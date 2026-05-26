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

	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	koordinatorclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
)

// MergeCfg returns a merged interface. Value in new will
// override old's when both fields exist.
// It will throw an error if:
//  1. either of the inputs was nil;
//  2. inputs were not a pointer of the same json struct.
func MergeCfg(old, new interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MinInt64(i, j int64) int64 { _ = "STUB: not implemented"; return 0 }

func MaxInt64(i, j int64) int64 { _ = "STUB: not implemented"; return 0 }

func MinFloat64(i, j float64) float64 { _ = "STUB: not implemented"; return 0 }

func MaxFloat64(i, j float64) float64 { _ = "STUB: not implemented"; return 0 }

func RetryOnConflictOrTooManyRequests(fn func() error) error { _ = "STUB: not implemented"; return nil }

func RetryOnConflictOrTooManyRequestsOrConnectionClose(fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func isErrorConnectionClosed(err error) bool { _ = "STUB: not implemented"; return false }

func GeneratePodPatch(oldPod, newPod *corev1.Pod) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GeneratePodPatchWithUID(oldPod, newPod *corev1.Pod) ([]byte, error) {
	_ = "STUB: not implemented"
	// For safely patch, generate with the object UID.
	// This ensures we will not patch the different object with the same name.
	return nil, nil
}

func PatchPod(ctx context.Context, clientset clientset.Interface, oldPod, newPod *corev1.Pod, subResources ...string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// generate patch bytes for the update

// nothing to patch

// patch with pod client

// PatchPodSafe patches the pod with the object UID for safety.
// This ensures we will not patch the different object with the same name.
func PatchPodSafe(ctx context.Context, clientset clientset.Interface, oldPod, newPod *corev1.Pod, subResources ...string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// generate patch bytes for the update

// patch with pod client

func GenerateReservationPatch(oldReservation, newReservation *schedulingv1alpha1.Reservation) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateReservationPatchWithUID(oldReservation, newReservation *schedulingv1alpha1.Reservation) ([]byte, error) {
	_ = "STUB: not implemented"
	// For safely patch, generate with the object UID.
	// This ensures we will not patch the different object with the same name.
	return nil, nil
}

func PatchReservation(ctx context.Context, clientset koordinatorclientset.Interface, oldReservation, newReservation *schedulingv1alpha1.Reservation) (*schedulingv1alpha1.Reservation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nothing to patch

// NOTE: CRDs do not support strategy merge patch, so here falls back to merge patch.
// link: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#advanced-features-and-flexibility

// PatchReservationSafe patches the reservation with the object UID for safety.
// This ensures we will not patch the different object with the same name.
func PatchReservationSafe(ctx context.Context, clientset koordinatorclientset.Interface, oldReservation, newReservation *schedulingv1alpha1.Reservation) (*schedulingv1alpha1.Reservation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: CRDs do not support strategy merge patch, so here falls back to merge patch.
// link: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#advanced-features-and-flexibility

func GenerateNodePatch(oldNode, newNode *corev1.Node) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PatchNode(ctx context.Context, clientset clientset.Interface, oldNode, newNode *corev1.Node, subResources ...string) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// generate patch bytes for the update

// nothing to patch

// patch with node client

func GetNamespacedName(namespace, name string) string { _ = "STUB: not implemented"; return "" }

func BoolToFloat64(b bool) float64 { _ = "STUB: not implemented"; return 0 }

func IsIn(arr []string, val string) bool { _ = "STUB: not implemented"; return false }

// TODO: Replace this function with the standard library after go1.21+ version
func OnceValues(f func() ([]int, error)) func() ([]int, error) {
	_ = "STUB: not implemented"
	return nil
}
