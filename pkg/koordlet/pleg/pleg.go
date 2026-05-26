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

package pleg

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
)

const (
	eventsChanCapacity = 128
)

type PodLifeCycleHandler interface {
	OnPodAdded(podID string)
	OnPodDeleted(podID string)
	OnContainerAdded(podID, containerID string)
	OnContainerDeleted(podID, containerID string)
}

type PodLifeCycleHandlerFuncs struct {
	PodAddedFunc         func(podID string)
	PodDeletedFunc       func(podID string)
	ContainerAddedFunc   func(podID, containerID string)
	ContainerDeletedFunc func(podID, containerID string)
}

func (r PodLifeCycleHandlerFuncs) OnPodAdded(podID string) { _ = "STUB: not implemented"; return }

func (r PodLifeCycleHandlerFuncs) OnPodDeleted(podID string) { _ = "STUB: not implemented"; return }

func (r PodLifeCycleHandlerFuncs) OnContainerAdded(podID, containerID string) {
	_ = "STUB: not implemented"
	return
}

func (r PodLifeCycleHandlerFuncs) OnContainerDeleted(podID, containerID string) {
	_ = "STUB: not implemented"
	return
}

type HandlerID uint32

type Pleg interface {
	Run(<-chan struct{}) error
	AddHandler(PodLifeCycleHandler) HandlerID
	RemoverHandler(id HandlerID) PodLifeCycleHandler
}

func NewPLEG(cgroupRootPath string) (Pleg, error) {
	_ = "STUB: not implemented"
	return *new(Pleg), nil
}

type pleg struct {
	cgroupRootPath string

	// internal status of pleg
	idGenerator HandlerID

	handlerMutex sync.Mutex
	handlers     map[HandlerID]PodLifeCycleHandler

	podWatcher       Watcher
	containerWatcher Watcher

	events chan *event
}

func (p *pleg) AddHandler(handler PodLifeCycleHandler) HandlerID {
	_ = "STUB: not implemented"
	return *new(HandlerID)
}

func (p *pleg) RemoverHandler(id HandlerID) PodLifeCycleHandler {
	_ = "STUB: not implemented"
	return *new(PodLifeCycleHandler)
}

func (p *pleg) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

// here we choose cpu subsystem as ground truth,
// since we only need to watch one of all subsystems, and cpu subsystem always and must exist

// handle Pod event

// register watcher for containers

// handle Pod event

// remove watcher for containers

// handle Container event

// handle Container event

func (p *pleg) runEventHandler(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (p *pleg) handleEvent(event *event) { _ = "STUB: not implemented"; return }

type internalEventType int

const (
	podAdded internalEventType = iota
	podDeleted
	containerAdded
	containerDeleted
)

type event struct {
	eventType   internalEventType
	podID       string
	containerID string
}

func newPodEvent(podID string, eventType internalEventType) *event {
	_ = "STUB: not implemented"
	return nil
}

func newContainerEvent(podID, containerID string, eventType internalEventType) *event {
	_ = "STUB: not implemented"
	return nil
}

// getWatchCgroupPath gets the directory path to watch for the given qos class, compatible to cgroups-v2.
func getWatchCgroupPath(cgroupRootDir string, qosClass corev1.PodQOSClass) string {
	_ = "STUB: not implemented"
	return ""
}

// cgroups-v2

// cgroups-v1
// here we choose cpu subsystem as ground truth,
// since we only need to watch one of all subsystems, and cpu subsystem always and must exist
