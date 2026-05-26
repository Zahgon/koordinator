/*
Copyright 2022 The Koordinator Authors.
Copyright 2019 The Kubernetes Authors.

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

package statefulset

import (
	"regexp"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// NewStatefulSet creates a new Webserver StatefulSet for testing. The StatefulSet is named name, is in namespace ns,
// statefulPodsMounts are the mounts that will be backed by PVs. podsMounts are the mounts that are mounted directly
// to the Pod. labels are the labels that will be usd for the StatefulSet selector.
func NewStatefulSet(name, ns, governingSvcName string, replicas int32, statefulPodMounts []v1.VolumeMount, podMounts []v1.VolumeMount, labels map[string]string) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

// NewStatefulSetPVC returns a PersistentVolumeClaim named name, for testing StatefulSets.
func NewStatefulSetPVC(name string) v1.PersistentVolumeClaim {
	_ = "STUB: not implemented"
	return *new(v1.PersistentVolumeClaim)
}

func hasPauseProbe(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

var pauseProbe = &v1.Probe{
	ProbeHandler: v1.ProbeHandler{
		Exec: &v1.ExecAction{Command: []string{"test", "-f", "/data/statefulset-continue"}},
	},
	PeriodSeconds:    1,
	SuccessThreshold: 1,
	FailureThreshold: 1,
}

type statefulPodsByOrdinal []v1.Pod

func (sp statefulPodsByOrdinal) Len() int { _ = "STUB: not implemented"; return 0 }

func (sp statefulPodsByOrdinal) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (sp statefulPodsByOrdinal) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// PauseNewPods adds an always-failing ReadinessProbe to the StatefulSet PodTemplate.
// This causes all newly-created Pods to stay Unready until they are manually resumed
// with ResumeNextPod().
// Note that this cannot be used together with SetHTTPProbe().
func PauseNewPods(ss *appsv1.StatefulSet) { _ = "STUB: not implemented"; return }

// ResumeNextPod allows the next Pod in the StatefulSet to continue by removing the ReadinessProbe
// added by PauseNewPods(), if it's still there.
// It fails the test if it finds any pods that are not in phase Running,
// or if it finds more than one paused Pod existing at the same time.
// This is a no-op if there are no paused pods.
func ResumeNextPod(c clientset.Interface, ss *appsv1.StatefulSet) {
	_ = "STUB: not implemented"
	return
}

// SortStatefulPods sorts pods by their ordinals
func SortStatefulPods(pods *v1.PodList) { _ = "STUB: not implemented"; return }

var statefulPodRegex = regexp.MustCompile("(.*)-([0-9]+)$")

func getStatefulPodOrdinal(pod *v1.Pod) int { _ = "STUB: not implemented"; return 0 }
