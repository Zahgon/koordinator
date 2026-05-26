/*
Copyright 2022 The Koordinator Authors.
Copyright 2020 The Kubernetes Authors.

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

package controller

import (
	"time"

	v1 "k8s.io/api/core/v1"
	coreinformer "k8s.io/client-go/informers/core/v1"
	corelister "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	schedv1alpha1 "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
	schedclientset "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/clientset/versioned"
	schedinformer "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/informers/externalversions/scheduling/v1alpha1"
	schedlister "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/coscheduling/core"
)

const (
	maxGCTime              = 48 * time.Hour
	PodGroupControllerName = "PodGroupController"
)

// PodGroupController  is used to control that process pod groups using provided Handler interface
type PodGroupController struct {
	pgQueue         workqueue.RateLimitingInterface
	pgLister        schedlister.PodGroupLister
	podLister       corelister.PodLister
	pgListerSynced  cache.InformerSynced
	podListerSynced cache.InformerSynced
	pgClient        schedclientset.Interface
	pgManager       core.Manager
	workers         int
}

// NewPodGroupController returns a new *PodGroupController
func NewPodGroupController(
	pgInformer schedinformer.PodGroupInformer,
	podInformer coreinformer.PodInformer,
	pgClient schedclientset.Interface,
	podGroupManager *core.PodGroupManager,
	workers int,
) *PodGroupController {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl PodGroupController) Name() string { _ = "STUB: not implemented"; return "" }

func (ctrl *PodGroupController) Start() { _ = "STUB: not implemented"; return }

// Run starts listening on channel events
func (ctrl *PodGroupController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// pgAdded reacts to a PG creation
func (ctrl *PodGroupController) pgAdded(obj interface{}) { _ = "STUB: not implemented"; return }

// If startScheduleTime - createTime > 2days, do not enqueue again because pod may have been GCed

// pgUpdated reacts to a PG update
func (ctrl *PodGroupController) pgUpdated(old, new interface{}) {
	_ = "STUB: not implemented"

	// podAdded reacts to a PG creation
	return
}

func (ctrl *PodGroupController) podAdded(obj interface{}) { _ = "STUB: not implemented"; return }

// pgUpdated reacts to a PG update
func (ctrl *PodGroupController) podUpdated(old, new interface{}) { _ = "STUB: not implemented"; return }

func (ctrl *PodGroupController) worker() { _ = "STUB: not implemented"; return }

// processNextWorkItem deals with one key off the queue.  It returns false when it's time to quit.
func (ctrl *PodGroupController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// syncHandle syncs pod group and convert status
func (ctrl *PodGroupController) syncHandler(key string) error {
	_ = "STUB: not implemented"
	// Convert the namespace/name string into a distinct namespace and name
	return nil
}

// get all pods belong to the PogGroup from gangCache

// when update the pod's Status, gangCache has not changed,
// so we get the pods' status from the informer according to pods' keys in gangCache
// it may happen that when pod is created, we may get the onAdd event here before the gangCache,
// so we lost the information of the pods, but the lost message can be repaired after the update event

// Final state of pod group

func (ctrl *PodGroupController) patchPodGroup(old, new *schedv1alpha1.PodGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func fillOccupiedObj(pg *schedv1alpha1.PodGroup, pod *v1.Pod) { _ = "STUB: not implemented"; return }
