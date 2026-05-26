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

package batchresource

import (
	"context"

	topologyv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/util/workqueue"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

var checkedNRTResourceSet = sets.New[corev1.ResourceName](corev1.ResourceCPU, corev1.ResourceMemory)

var _ handler.EventHandler = &NRTHandler{}

type NRTHandler struct {
	syncContext *framework.SyncContext
}

func (h *NRTHandler) Create(ctx context.Context, evt event.TypedCreateEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *NRTHandler) Update(ctx context.Context, evt event.TypedUpdateEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *NRTHandler) Delete(ctx context.Context, evt event.TypedDeleteEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *NRTHandler) Generic(ctx context.Context, evt event.TypedGenericEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func isNRTResourcesCreated(nrt *topologyv1alpha1.NodeResourceTopology) bool {
	_ = "STUB: not implemented"
	return false
}

// check if any zone has the target resource allocatable

func isNRTResourcesChanged(nrtOld, nrtNew *topologyv1alpha1.NodeResourceTopology) bool {
	_ = "STUB: not implemented"
	// check if target resources not equal
	return false
}

func cleanupContextForNRT(syncContext *framework.SyncContext, nrt *topologyv1alpha1.NodeResourceTopology) error {
	_ = "STUB: not implemented"
	return nil
}

// NRT name = node name
