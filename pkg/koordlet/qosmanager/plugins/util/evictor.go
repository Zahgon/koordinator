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
	"go.uber.org/atomic"
	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"

	expireCache "github.com/koordinator-sh/koordinator/pkg/util/cache"
)

type Evictor struct {
	eventRecorder record.EventRecorder
	kubeClient    clientset.Interface
	podsEvicted   *expireCache.Cache
	evictVersion  string
	started       atomic.Bool
}

func NewEvictor(kubeClient clientset.Interface, eventRecorder record.EventRecorder, evictVersion string) *Evictor {
	_ = "STUB: not implemented"
	return nil
}

func (r *Evictor) Start(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

func (r *Evictor) EvictPodsIfNotEvicted(evictPods []*corev1.Pod, reason string, message string) {
	_ = "STUB: not implemented"
	return
}

func (r *Evictor) IsPodEvicted(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (r *Evictor) EvictPodIfNotEvicted(evictPod *corev1.Pod, reason string, message string) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Evictor) evictPod(evictPod *corev1.Pod, reason string, message string) bool {
	_ = "STUB: not implemented"
	return false
}
