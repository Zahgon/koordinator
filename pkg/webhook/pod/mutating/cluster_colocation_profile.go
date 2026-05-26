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
	"math/rand"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	configv1alpha1 "github.com/koordinator-sh/koordinator/apis/config/v1alpha1"
	"github.com/koordinator-sh/koordinator/apis/extension"
)

var (
	randIntnFn = rand.Intn
)

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=config.koordinator.sh,resources=clustercolocationprofiles,verbs=get;list;watch

func (h *PodMutatingHandler) clusterColocationProfileMutatingPod(ctx context.Context, req admission.Request, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// sort the profile in lexicographic order

func (h *PodMutatingHandler) matchNamespaceSelector(ctx context.Context, namespaceName string, namespaceSelector *metav1.LabelSelector) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (h *PodMutatingHandler) matchObjectSelector(pod, oldPod *corev1.Pod, objectSelector *metav1.LabelSelector) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func shouldSkipProfile(profile *configv1alpha1.ClusterColocationProfile) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (h *PodMutatingHandler) doMutateByColocationProfile(ctx context.Context, pod *corev1.Pod, profile *configv1alpha1.ClusterColocationProfile) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *PodMutatingHandler) mutatePodResourceSpec(pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func replaceAndEraseResource(priorityClass extension.PriorityClass, resourceList corev1.ResourceList, resourceName corev1.ResourceName) {
	_ = "STUB: not implemented"
	return
}

// TODO move the hook to pod mutating for all Pods
func restrictResourceRequestAndLimit(priorityClass extension.PriorityClass, requirements *corev1.ResourceRequirements, resourceName corev1.ResourceName) {
	_ = "STUB: not implemented"
	return
}
