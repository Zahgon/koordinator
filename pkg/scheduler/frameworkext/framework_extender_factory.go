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

package frameworkext

import (
	"context"

	nrtinformers "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions"
	"github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	frameworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
	"k8s.io/kubernetes/pkg/scheduler/metrics"

	koordinatorclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/networktopology"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/services"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/workloadauditor"
)

var (
	EnableNetworkTopologyManager = false
)

func AddFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type extendedHandleOptions struct {
	servicesEngine                      *services.Engine
	koordinatorClientSet                koordinatorclientset.Interface
	koordinatorSharedInformerFactory    koordinatorinformers.SharedInformerFactory
	nodeResourceTopologyInformerFactory nrtinformers.SharedInformerFactory
	reservationCache                    ReservationCache
	reservationNominator                ReservationNominator
	networkTopologyManager              networktopology.TreeManager
	crossSchedulerNominator             *CrossSchedulerPodNominator
	workloadAuditor                     workloadauditor.WorkloadAuditor
}

type Option func(*extendedHandleOptions)

func WithServicesEngine(engine *services.Engine) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithKoordinatorClientSet(koordinatorClientSet koordinatorclientset.Interface) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithKoordinatorSharedInformerFactory(informerFactory koordinatorinformers.SharedInformerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithNodeResourceTopologySharedInformerFactory(informerFactory nrtinformers.SharedInformerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithReservationCache(cache ReservationCache) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithReservationNominator(nominator ReservationNominator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithNetworkTopologyManager(manager networktopology.TreeManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCrossSchedulerPodNominator(nominator *CrossSchedulerPodNominator) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithWorkloadAuditor(auditor workloadauditor.WorkloadAuditor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// FrameworkExtenderFactory is a factory for creating a FrameworkExtender.
// NOTE: DO NOT put framework-level data here.
type FrameworkExtenderFactory struct {
	controllerMaps                      *ControllersMap
	servicesEngine                      *services.Engine
	koordinatorClientSet                koordinatorclientset.Interface
	koordinatorSharedInformerFactory    koordinatorinformers.SharedInformerFactory
	nodeResourceTopologyInformerFactory nrtinformers.SharedInformerFactory
	reservationCache                    ReservationCache     // for testing
	reservationNominator                ReservationNominator // for testing
	nextPodPlugin                       NextPodPlugin
	profiles                            map[string]FrameworkExtender
	monitor                             *SchedulerMonitor
	scheduler                           Scheduler
	schedulePod                         func(ctx context.Context, fwk framework.Framework, state fwktype.CycleState, pod *corev1.Pod) (scheduler.ScheduleResult, error)
	*errorHandlerDispatcher

	networkTopologyTreeManager networktopology.TreeManager
	crossSchedulerNominator    *CrossSchedulerPodNominator

	workloadAuditor workloadauditor.WorkloadAuditor

	pluginInformerFactories []SharedInformerFactory

	metricsRecorder *metrics.MetricAsyncRecorder
}

func NewFrameworkExtenderFactory(options ...Option) (*FrameworkExtenderFactory, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FrameworkExtenderFactory) NewFrameworkExtender(fw framework.Framework) FrameworkExtender {
	_ = "STUB: not implemented"
	return *new(FrameworkExtender)
}

func (f *FrameworkExtenderFactory) GetExtender(profileName string) FrameworkExtender {
	_ = "STUB: not implemented"
	return *new(FrameworkExtender)
}

func (f *FrameworkExtenderFactory) KoordinatorClientSet() koordinatorclientset.Interface {
	_ = "STUB: not implemented"
	return *new(koordinatorclientset.Interface)
}

func (f *FrameworkExtenderFactory) KoordinatorSharedInformerFactory() koordinatorinformers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(koordinatorinformers.SharedInformerFactory)
}

// Scheduler return the scheduler adapter to support operating with cache and schedulingQueue.
// NOTE: Plugins do not acquire a dispatcher instance during plugin initialization,
// nor are they allowed to hold the object within the plugin object.
func (f *FrameworkExtenderFactory) Scheduler() Scheduler {
	_ = "STUB: not implemented"
	return *new(Scheduler)
}

func (f *FrameworkExtenderFactory) InitScheduler(sched Scheduler) {
	_ = "STUB: not implemented"
	return
}

// NextPodPlugin but has no suggestion for which Pod to dequeue next and falls back to the original nextPod logic

// just for plugins to get Pod queue information

// The podInfo can be nil when the queue is closing.
// Deep copy podInfo to allow pod modification during scheduling

func (f *FrameworkExtenderFactory) runNextPodPlugin() (*framework.QueuedPodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// should not delete it from nominator, so use a fake UID

const (
	initialTimestampManager = "scheduler.scheduling.koordinator.sh/initialTimestamp"
	attemptsManager         = "scheduler.scheduling.koordinator.sh/attempts"
)

func CopyQueueInfoToPod(podHasQueueInfo, podNeedQueueInfo *corev1.Pod) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// avoid directly modifying the original Pod and only copy the fields that are needed

func RecordPodQueueInfoToPod(podInfo *framework.QueuedPodInfo) { _ = "STUB: not implemented"; return }

// the podInfo can be nil when the queue is closing

// avoid directly modifying the original Pod and only copy the fields that are needed

func makePodInfoFromPod(pod *corev1.Pod) (*framework.QueuedPodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FrameworkExtenderFactory) scheduleOne(ctx context.Context, fwk framework.Framework, cycleState fwktype.CycleState, pod *corev1.Pod) (scheduler.ScheduleResult, error) {
	_ = "STUB: not implemented"
	return *new(scheduler.ScheduleResult), nil
}

// Due to some ResizePod plugins need the reservation nomination before the real Reserve phase,
// and the PreScore phase might be skipped, we force to nominate reservation here.
// https://github.com/koordinator-sh/koordinator/issues/2753

// NOTE(joseph): We can modify the Pod because we have cloned the Pod in the NextPod function.
// Make sure to modify the Pod only related to AssumePod, and do not modify the plugins' cache, since
// when the assume failure would not do Unreserve the plugins' cache.
// Other resizing logic about the plugins' cache can be convergent to its Reserve phase.

func recordScheduleDiagnosis(cycleState fwktype.CycleState, err error) {
	_ = "STUB: not implemented"
	return
}

func (f *FrameworkExtenderFactory) CollectSchedulePodResult(sched *scheduler.Scheduler) {
	_ = "STUB: not implemented"
	return
}

// avoid recording metrics when there is no feasible node or internal error in scheduling

func (f *FrameworkExtenderFactory) InterceptSchedulerError(sched *scheduler.Scheduler) {
	_ = "STUB: not implemented"
	return
}

func (f *FrameworkExtenderFactory) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (f *FrameworkExtenderFactory) updatePlugins(pl fwktype.Plugin, profileName string) {
	_ = "STUB: not implemented"
	return
}

// GetPluginInformerFactories returns all informer factories registered by plugins
// via the InformerFactoryProvider interface.
func (f *FrameworkExtenderFactory) GetPluginInformerFactories() []SharedInformerFactory {
	_ = "STUB: not implemented"
	return nil
}

// PluginFactoryProxy is used to proxy the call to the PluginFactory function and pass in the ExtendedHandle for the custom plugin
func PluginFactoryProxy(extenderFactory *FrameworkExtenderFactory, factoryFn frameworkruntime.PluginFactory) frameworkruntime.PluginFactory {
	_ = "STUB: not implemented"
	return *new(frameworkruntime.PluginFactory)
}
