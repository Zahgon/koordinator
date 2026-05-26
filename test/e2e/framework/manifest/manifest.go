/*
Copyright 2022 The Koordinator Authors.
Copyright 2017 The Kubernetes Authors.

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

package manifest

import (
	appsv1 "k8s.io/api/apps/v1"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

// PodFromManifest reads a .json/yaml file and returns the pod in it.
func PodFromManifest(filename string) (*v1.Pod, error) { _ = "STUB: not implemented"; return nil, nil }

// SvcFromManifest reads a .json/yaml file and returns the service in it.
func SvcFromManifest(fileName string) (*v1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StatefulSetFromManifest returns a StatefulSet from a manifest stored in fileName in the Namespace indicated by ns.
func StatefulSetFromManifest(fileName, ns string) (*appsv1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DaemonSetFromURL reads from a url and returns the daemonset in it.
func DaemonSetFromURL(url string) (*appsv1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DaemonSetFromData reads a byte slice and returns the daemonset in it.
func DaemonSetFromData(data []byte) (*appsv1.DaemonSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigMapFromManifest reads a .json/yaml file and returns the pod in it.
func ConfigMapFromManifest(filename string) (*v1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReservationFromManifest(filename string) (*schedulingv1alpha1.Reservation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
