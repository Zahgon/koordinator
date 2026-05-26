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

package cpunormalization

import (
	"context"

	topologyv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	"k8s.io/client-go/util/workqueue"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ handler.TypedEventHandler[ctrlclient.Object, reconcile.Request] = &nrtHandler{}

type nrtHandler struct{}

func (h *nrtHandler) Create(ctx context.Context, evt event.TypedCreateEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *nrtHandler) Update(ctx context.Context, evt event.TypedUpdateEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *nrtHandler) Delete(ctx context.Context, evt event.TypedDeleteEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (h *nrtHandler) Generic(ctx context.Context, evt event.TypedGenericEvent[ctrlclient.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func isNRTCPUBasicInfoCreated(nrt *topologyv1alpha1.NodeResourceTopology) bool {
	_ = "STUB: not implemented"
	return false
}

func isNRTCPUBasicInfoChanged(nrtOld, nrtNew *topologyv1alpha1.NodeResourceTopology) bool {
	_ = "STUB: not implemented"
	return false
}

// ignore old error
