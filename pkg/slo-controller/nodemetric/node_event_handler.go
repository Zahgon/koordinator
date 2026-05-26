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

package nodemetric

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ handler.TypedEventHandler[client.Object, reconcile.Request] = &EnqueueRequestForNode{}

type EnqueueRequestForNode struct {
	client.Client
}

func (n *EnqueueRequestForNode) Create(ctx context.Context, e event.TypedCreateEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (n *EnqueueRequestForNode) Update(ctx context.Context, e event.TypedUpdateEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

// TODO, only use for noderesource

func (n *EnqueueRequestForNode) Delete(ctx context.Context, e event.TypedDeleteEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"
	return
}

func (n *EnqueueRequestForNode) Generic(ctx context.Context, e event.TypedGenericEvent[client.Object], q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	_ = "STUB: not implemented"

	// isNodeUpdated returns whether the new node's allocatable or labels is different from the old one's
	return
}

func isNodeUpdated(newNode *corev1.Node, oldNode *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}
