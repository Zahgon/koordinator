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

package coscheduling

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fwktype "k8s.io/kube-scheduler/framework"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	pgclientset "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/clientset/versioned"
	pgformers "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/informers/externalversions"
	schedinformers "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/informers/externalversions/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/coscheduling/core"
)

// Coscheduling is a plugin that schedules pods in a group.
type Coscheduling struct {
	args              *config.CoschedulingArgs
	frameworkHandler  fwktype.Handle
	pgClient          pgclientset.Interface
	pgInformerFactory pgformers.SharedInformerFactory
	pgInformer        schedinformers.PodGroupInformer
	pgMgr             core.Manager
}

var _ fwktype.PreEnqueuePlugin = &Coscheduling{}
var _ frameworkext.NextPodPlugin = &Coscheduling{}
var _ frameworkext.PreFilterTransformer = &Coscheduling{}
var _ fwktype.PreFilterPlugin = &Coscheduling{}
var _ frameworkext.FindOneNodePluginProvider = &Coscheduling{}
var _ frameworkext.FindOneNodePlugin = &Coscheduling{}
var _ frameworkext.PostFilterTransformer = &Coscheduling{}
var _ fwktype.PostFilterPlugin = &Coscheduling{}
var _ fwktype.PreScorePlugin = &Coscheduling{}
var _ fwktype.ScorePlugin = &Coscheduling{}
var _ fwktype.PermitPlugin = &Coscheduling{}
var _ fwktype.ReservePlugin = &Coscheduling{}
var _ fwktype.PreBindPlugin = &Coscheduling{}
var _ frameworkext.ReservationPreBindPlugin = &Coscheduling{}
var _ fwktype.PostBindPlugin = &Coscheduling{}
var _ fwktype.EnqueueExtensions = &Coscheduling{}
var _ frameworkext.InformerFactoryProvider = &Coscheduling{}

const (
	// Name is the name of the plugin used in Registry and configurations.
	Name = core.Name
)

// New initializes and returns a new Coscheduling plugin.
func New(_ context.Context, obj runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (cs *Coscheduling) EventsToRegister(_ context.Context) ([]fwktype.ClusterEventWithHint, error) {
	_ = "STUB: not implemented"
	// indicates that we are not interested in any events
	return nil, nil
}

// Name returns name of the plugin. It is used in logs, etc.
func (cs *Coscheduling) Name() string {
	_ = "STUB: not implemented"

	// PreEnqueue
	// i.Check whether childes in Gang has met the requirements of minimum number under each Gang, and reject the pod if negative.
	// ii.Check whether the Gang has been timeout(check the pod's annotation,later introduced at Permit section) or is inited, and reject the pod if positive.
	return ""
}

func (cs *Coscheduling) PreEnqueue(ctx context.Context, pod *v1.Pod) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (cs *Coscheduling) NextPod() *v1.Pod { _ = "STUB: not implemented"; return nil }

func (cs *Coscheduling) FindOneNodePlugin() frameworkext.FindOneNodePlugin {
	_ = "STUB: not implemented"
	return *new(frameworkext.FindOneNodePlugin)
}

func (cs *Coscheduling) FindOneNode(ctx context.Context, cycleState fwktype.CycleState, pod *v1.Pod, result *fwktype.PreFilterResult) (string, *fwktype.Status) {
	_ = "STUB: not implemented"
	return "", nil
}

// BeforePreFilter
// i.Check whether the Gang has met the scheduleCycleValid check, and reject the pod if negative.
// ii.Try update scheduleCycle, scheduleCycleValid, childrenScheduleRoundMap as mentioned above.
func (cs *Coscheduling) BeforePreFilter(ctx context.Context, state fwktype.CycleState, pod *v1.Pod) (*v1.Pod, bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	// If PreFilter fails, return fwktype.UnschedulableAndUnresolvable to avoid any preemption attempts.
	return nil, false, nil
}

func (cs *Coscheduling) PreFilter(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *Coscheduling) AfterPreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *v1.Pod, preFilterResult *fwktype.PreFilterResult) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (cs *Coscheduling) AfterPostFilter(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, filteredNodeStatusMap fwktype.NodeToStatusReader) {
	_ = "STUB: not implemented"
	return
}

// PostFilter
// i. If strict-mode, we will set scheduleCycleValid to false and release all assumed pods.
// ii. If non-strict mode, we will do nothing.
func (cs *Coscheduling) PostFilter(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, filteredNodeStatusMap fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PreFilterExtensions returns a PreFilterExtensions interface if the plugin implements one.
func (cs *Coscheduling) PreFilterExtensions() fwktype.PreFilterExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.PreFilterExtensions)
}

func (cs *Coscheduling) PreScore(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, nodes []fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (cs *Coscheduling) Score(ctx context.Context, state fwktype.CycleState, p *v1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cs *Coscheduling) ScoreExtensions() fwktype.ScoreExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.ScoreExtensions)
}

func (cs *Coscheduling) NormalizeScore(ctx context.Context, state fwktype.CycleState, p *v1.Pod, scores fwktype.NodeScoreList) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// Permit
// we will calculate all Gangs in GangGroup whether the current number of assumed-pods in each Gang meets the Gang's minimum requirement.
// and decide whether we should let the pod wait in Permit stage or let the whole gangGroup go binding
func (cs *Coscheduling) Permit(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, nodeName string) (*fwktype.Status, time.Duration) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration)
}

// Reserve is the functions invoked by the framework at "reserve" extension point.
func (cs *Coscheduling) Reserve(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"

	// Unreserve
	// i. handle the timeout gang
	// ii. do nothing when bound failed
	return nil
}

func (cs *Coscheduling) Unreserve(ctx context.Context, state fwktype.CycleState, pod *v1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (cs *Coscheduling) PreBindPreFlight(ctx context.Context, cycleState fwktype.CycleState, pod *v1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (cs *Coscheduling) PreBind(ctx context.Context, cycleState fwktype.CycleState, pod *v1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// DEPRECATED: This api is marked as internal and will be removed next version.

// DEPRECATED: This api is marked as internal and will be removed next version.

func (cs *Coscheduling) PreBindReservation(ctx context.Context, cycleState fwktype.CycleState, r *schedulingv1alpha1.Reservation, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// DEPRECATED: This api is marked as internal and will be removed next version.

// DEPRECATED: This api is marked as internal and will be removed next version.

// PostBind is called after a pod is successfully bound. These plugins are used update PodGroup when pod is bound.
func (cs *Coscheduling) PostBind(ctx context.Context, _ fwktype.CycleState, pod *v1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// GetInformerFactories returns the PodGroup informer factory for central startup management.
func (cs *Coscheduling) GetInformerFactories() []frameworkext.SharedInformerFactory {
	_ = "STUB: not implemented"
	return nil
}
