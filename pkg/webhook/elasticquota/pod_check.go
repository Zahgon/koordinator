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

package elasticquota

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO If the parentQuotaGroup submits pods, the runtime will be calculated incorrectly.
// Supporting the parentQuotaGroup to submit pods is the future work.

func (qt *quotaTopology) ValidateAddPod(pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (qt *quotaTopology) ValidateUpdatePod(oldPod, newPod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (qt *quotaTopology) getQuotaNameFromPodNoLock(pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}

func GetQuotaName(pod *corev1.Pod, kubeClient client.Client) string {
	_ = "STUB: not implemented"
	return ""
}

// hasQuotaBoundedPods returns true if the quota has bounded pods.
func hasQuotaBoundedPods(kubeClient client.Client, quotaName string, namespaces []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
