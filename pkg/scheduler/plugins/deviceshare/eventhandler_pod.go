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

package deviceshare

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"

	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
)

func registerPodEventHandler(deviceCache *nodeDeviceCache, sharedInformerFactory informers.SharedInformerFactory, koordSharedInformerFactory koordinatorinformers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

// make sure Pods are loaded before scheduler starts working

func (n *nodeDeviceCache) onPodAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (n *nodeDeviceCache) onPodUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDeviceCache) onPodDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (n *nodeDeviceCache) updatePod(oldPod *corev1.Pod, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	// For multi-scheduler scenarios: clean up old pod when new pod becomes unassigned
	return
}

// avoid leaking for an unassigned pod

func (n *nodeDeviceCache) deletePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }
