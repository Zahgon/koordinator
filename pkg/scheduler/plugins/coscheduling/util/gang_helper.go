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
	"time"

	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

func GetGangGroupId(s []string) string { _ = "STUB: not implemented"; return "" }

func GetId(namespace, name string) string { _ = "STUB: not implemented"; return "" }

func GetGangNameByPod(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func GetGangMinNumFromPod(pod *corev1.Pod) (minNum int, err error) {
	_ = "STUB: not implemented"
	// nolint:staticcheck // SA1019: extension.LabelLightweightCoschedulingPodGroupMinAvailable is deprecated
	return 0, nil
}

func IsPodNeedGang(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// GetWaitTimeDuration returns a wait timeout based on the following precedences:
// 1. spec.scheduleTimeoutSeconds of the given pg, if specified
// 2. fall back to defaultTimeout
func GetWaitTimeDuration(pg *v1alpha1.PodGroup, defaultTimeout time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// StringToGangGroupSlice
// Parse gang group's annotation like :"["nsA/gangA","nsB/gangB"]"  => goLang slice : []string{"nsA/gangA"."nsB/gangB"}
func StringToGangGroupSlice(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateMergePatch return patch generated from original and new interfaces
func CreateMergePatch(original, new interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
