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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/config"
)

const Name = "nodemetric"

// NodeMetricReconciler reconciles a NodeMetric object
type NodeMetricReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	cfgCache config.ColocationCfgCache
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=slo.koordinator.sh,resources=nodemetrics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=slo.koordinator.sh,resources=nodemetrics/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *NodeMetricReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// if cache unavailable, requeue the req

// all nodes would be enqueued once the config is available, so here we just drop the req

// return if neither Node nor NodeMetric exists.

// if !nodeExist && nodeMetricExist, delete NodeMetric.

// if nodeExist && !nodeMetricExist, create an empty NodeMetric CR.

// update nodeMetric spec if both exists

func (r *NodeMetricReconciler) getNodeMetricSpec(node *corev1.Node, oldSpec *slov1alpha1.NodeMetricSpec) (*slov1alpha1.NodeMetricSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if cfg cache error status(like unmarshall error),then use oldSpec

func getDefaultSpec() *slov1alpha1.NodeMetricSpec { _ = "STUB: not implemented"; return nil }

func Add(mgr ctrl.Manager) error { _ = "STUB: not implemented"; return nil }

// SetupWithManager sets up the controller with the Manager.
func (r *NodeMetricReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
