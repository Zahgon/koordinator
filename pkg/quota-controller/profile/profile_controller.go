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

package profile

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/quota/v1alpha1"
)

const Name = "quotaprofile"

const (
	ReasonCreateQuotaFailed = "CreateQuotaFailed"
	ReasonUpdateQuotaFailed = "UpdateQuotaFailed"
)

var resourceDecorators = []func(profile *v1alpha1.ElasticQuotaProfile, total corev1.ResourceList){
	DecorateResourceByResourceRatio,
}

func decorateTotalResource(profile *v1alpha1.ElasticQuotaProfile, total corev1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

// QuotaProfileReconciler reconciles a QuotaProfile object
type QuotaProfileReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
// +kubebuilder:rbac:groups=scheduling.sigs.k8s.io,resources=elasticquotas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=quota.koordinator.sh,resources=elasticquotaprofiles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=quota.koordinator.sh,resources=elasticquotaprofiles/status,verbs=get;update;patch

func (r *QuotaProfileReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	// reconcile for 2 things:
	//   1. ensuring the Quota exists if the QuotaProfile exists
	//   2. update Quota Spec
	return *new(ctrl.Result), nil
}

// get the profile

// generate quota tree id

// TODO: consider node status.

// update min and max

// update quota label

// update quota tree id

// update quota root label

// update total resource

// update unschedulable resource

func Add(mgr ctrl.Manager) error { _ = "STUB: not implemented"; return nil }

func (r *QuotaProfileReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func MultiplyQuantity(value resource.Quantity, resName corev1.ResourceName, ratio float64) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

func hash(s string) string { _ = "STUB: not implemented"; return "" }

func DecorateResourceByResourceRatio(profile *v1alpha1.ElasticQuotaProfile, total corev1.ResourceList) {
	_ = "STUB: not implemented"
	return
}
