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

package nodeslo

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

const Name = "nodeslo"

// NodeSLOReconciler reconciles a NodeSLO object
type NodeSLOReconciler struct {
	client.Client
	sloCfgCache SLOCfgCache
	Scheme      *runtime.Scheme
	Recorder    record.EventRecorder
}

func (r *NodeSLOReconciler) initNodeSLO(node *corev1.Node, nodeSLO *slov1alpha1.NodeSLO) error {
	_ = "STUB: not implemented"
	// NOTE: the node and nodeSLO should not be nil
	// get spec from a configmap
	return nil
}

func (r *NodeSLOReconciler) getNodeSLOSpec(node *corev1.Node, oldSpec *slov1alpha1.NodeSLOSpec) (*slov1alpha1.NodeSLOSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resourceQOS spec

// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=slo.koordinator.sh,resources=nodeslos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=slo.koordinator.sh,resources=nodeslos/status,verbs=get;update;patch

func (r *NodeSLOReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	// reconcile for 2 things:
	//   1. ensuring the NodeSLO exists if the Node exists
	//   2. update NodeSLO Spec
	return *new(ctrl.Result), nil
}

// if cache unavailable, requeue the req

// all nodes would be enqueued once the config is available, so here we just drop the req

// get the node

// get the nodeSLO

// NodeSLO lifecycle management

// do nothing if both does not exist

// delete CR if only the nodeSLO exists

// create and initialize CR if only the node exists

// update nodeSLO spec if both exists

func Add(mgr ctrl.Manager) error { _ = "STUB: not implemented"; return nil }

func (r *NodeSLOReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
