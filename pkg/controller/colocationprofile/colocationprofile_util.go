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

package colocationprofile

import (
	"context"
	"math/rand"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configv1alpha1 "github.com/koordinator-sh/koordinator/apis/config/v1alpha1"
)

var (
	randIntnFn = rand.Intn
)

func (r *Reconciler) listPodsForProfile(profile *configv1alpha1.ClusterColocationProfile) (*corev1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
	// match nothing
}

// list pods with label selectors

// NOTE: Only handle pending pods.

func (r *Reconciler) isPodNamespaceMatched(ctx context.Context, pod *corev1.Pod, nsSelector *metav1.LabelSelector) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// no selector means match all

// empty selector means match all

// namespace not found, so not match

func (r *Reconciler) updatePodByClusterColocationProfile(ctx context.Context, profile *configv1alpha1.ClusterColocationProfile, pod *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Reconciler) doMutateByColocationProfile(pod *corev1.Pod, profile *configv1alpha1.ClusterColocationProfile) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: Below fields are not supported by colocation-profile controller:
// - PriorityClassName
// - SchedulerName

type ReconcileSummary struct {
	Time        string
	Profile     string
	Desired     int
	Succeeded   int
	Changed     int
	RateLimited int
	Skipped     int
	Cached      int
}

func newSummary(profileName string) *ReconcileSummary { _ = "STUB: not implemented"; return nil }

func (s *ReconcileSummary) IsAllSucceeded() bool { _ = "STUB: not implemented"; return false }

func shouldSkipProfile(profile *configv1alpha1.ClusterColocationProfile) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getPodUpdateKey(profile *configv1alpha1.ClusterColocationProfile, pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}
