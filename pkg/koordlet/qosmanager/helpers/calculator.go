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

package helpers

import (
	corev1 "k8s.io/api/core/v1"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

// CalculateFilterPodsUsed calculates the sum used of filtered pods and hostApps.
// It returns the sum used of the filtered pods, the sum used of the filtered hostApps and the system used.
// If hostApps not passed, the hostApps part will be counted into the system used.
//
// e.g. To calculate non-BE used,
//   - filterPodsUsed := sum(podUsed[i] if f_pod(pod[i]))
//   - filterHostAppUsed := sum(hostAppUsed[i] if f_hostApp(hostApp[i]))
//   - systemdUsed := max(nodeReserved, nodeUsed - sum(podUsed[i]) - sum(hostAppUsed[i]))
func CalculateFilterPodsUsed(nodeUsed float64, nodeReserved float64,
	podMetas []*statesinformer.PodMeta, podUsedMap map[string]float64,
	hostApps []slov1alpha1.HostApplicationSpec, hostAppMetrics map[string]float64,
	podFilterFn func(*corev1.Pod) bool,
	hostAppFilterFn func(*slov1alpha1.HostApplicationSpec) bool) (float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// NOTE: consider podMeta-missing pods as filtered

// systemUsed means the remain used excluding the filtered pods and hostApps

// set systemUsed always no less than 0

// systemUsed = max(nodeUsed - podsAllUsed - hostAppsAllUsed, nodeAnnoReserved, nodeKubeletReserved)

func NotBatchOrFreePodFilter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func NonBEPodFilter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func NonBEHostAppFilter(hostAppSpec *slov1alpha1.HostApplicationSpec) bool {
	_ = "STUB: not implemented"
	// host app qos must be BE and also run under best-effort dir
	return false
}

func NonePodHighPriority(_ *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func NoneHostAppHighPriority(_ *slov1alpha1.HostApplicationSpec) bool {
	_ = "STUB: not implemented"
	return false
}

func GetNodeResourceReserved(node *corev1.Node) corev1.ResourceList {
	_ = "STUB: not implemented"
	// nodeReserved := max(nodeKubeletReserved, nodeAnnoReserved)
	return *new(corev1.ResourceList)
}
