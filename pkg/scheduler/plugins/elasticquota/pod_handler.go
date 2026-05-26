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
)

// todo the eventHandler's operation should be a complete transaction in the future work.

func (g *Plugin) OnPodAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (g *Plugin) OnPodUpdate(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

func (g *Plugin) OnPodDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (g *Plugin) handlePodDelete(pod *corev1.Pod) { _ = "STUB: not implemented"; return }
