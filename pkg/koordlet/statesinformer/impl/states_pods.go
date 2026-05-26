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

package impl

import (
	"sync"
	"time"

	"go.uber.org/atomic"
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/pleg"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	podsInformerName PluginName = "podsInformer"
)

type podsInformer struct {
	config *Config

	podRWMutex     sync.RWMutex
	podMap         map[string]*statesinformer.PodMeta
	podUpdatedTime time.Time
	podHasSynced   *atomic.Bool

	// use pleg to accelerate the efficiency of Pod meta update
	pleg       pleg.Pleg
	podCreated chan string

	kubelet      KubeletStub
	nodeInformer *nodeInformer

	callbackRunner *callbackRunner
	cgroupReader   resourceexecutor.CgroupReader
}

func NewPodsInformer() *podsInformer { _ = "STUB: not implemented"; return nil }

func (s *podsInformer) Setup(ctx *PluginOption, states *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (s *podsInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// There is no need to notify to update the data when the channel is not empty

func (s *podsInformer) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (s *podsInformer) GetAllPods() []*statesinformer.PodMeta {
	_ = "STUB: not implemented"
	return nil
}

func (s *podsInformer) getTaskIds(podMeta *statesinformer.PodMeta) {
	_ = "STUB: not implemented"
	return
}

func (s *podsInformer) syncPods() error { _ = "STUB: not implemented"; return nil }

// when kubelet recovers from crash, podList may be empty.

// reset pod container metrics

// no need to deep-copy from unmarshalled

// record pod's containers taskids

// record pod container metrics

func (s *podsInformer) syncKubeletLoop(duration time.Duration, stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// TODO add a config to setup the values

// sync kubelet triggered immediately when the Pod is created

// reset timer to

func newKubeletStubFromConfig(node *corev1.Node, cfg *Config) (KubeletStub, error) {
	_ = "STUB: not implemented"
	return *new(KubeletStub), nil
}

// if the address of the specified type has not been set or error type, InternalIP will be used.

func genPodCgroupParentDir(pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	// todo use cri interface to get pod cgroup dir
	// e.g. kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod9dba1d9e_67ba_4db6_8a73_fb3ea297c363.slice/
	return ""
}

func resetPodMetrics() { _ = "STUB: not implemented"; return }

func recordPodResourceMetrics(podMeta *statesinformer.PodMeta) { _ = "STUB: not implemented"; return }

// record (regular) container metrics

func recordContainerResourceMetrics(container *corev1.Container, containerStatus *corev1.ContainerStatus, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	// record pod requests/limits of BatchCPU & BatchMemory
	return
}

// record pod requests/limits of MidCPU & MidMemory
