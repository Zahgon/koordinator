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

package defaultprebind

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"
	fwktype "k8s.io/kube-scheduler/framework"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	koordinatorclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	Name = "DefaultPreBind"
)

var _ fwktype.PreBindPlugin = &Plugin{}
var _ frameworkext.PreBindExtensions = &Plugin{}

type Plugin struct {
	clientSet      clientset.Interface
	koordClientSet koordinatorclientset.Interface
}

type koordClientSetHandle interface {
	KoordinatorClientSet() koordinatorclientset.Interface
}

func New(_ context.Context, args runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (pl *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (pl *Plugin) PreBindPreFlight(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) PreBind(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) ApplyPatch(ctx context.Context, cycleState fwktype.CycleState, originalObj, modifiedObj metav1.Object) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) applyPodPatch(ctx context.Context, originalPod, modifiedPod *corev1.Pod) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: Patch might succeed for a deleting object without updating the data when the apiserver receive a Delete earlier.
// In this case, we should clean up the reserved resources to avoid cache leak.

func (pl *Plugin) applyReservationPatch(ctx context.Context, originalReservation, modifiedReservation *schedulingv1alpha1.Reservation) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: Patch might succeed for a deleting object without updating the data when the apiserver receive a Delete earlier.
// In this case, we should clean up the reserved resources to avoid cache leak.
