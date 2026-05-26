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

package reconciler

import (
	"sync"
	"time"

	"k8s.io/client-go/tools/record"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

type ReconcilerLevel string

const (
	KubeQOSLevel   ReconcilerLevel = "kubeqos"
	PodLevel       ReconcilerLevel = "pod"
	ContainerLevel ReconcilerLevel = "container"
	SandboxLevel   ReconcilerLevel = "sandbox"
	AllPodsLevel   ReconcilerLevel = "allpods"
)

var globalCgroupReconcilers = struct {
	all []*cgroupReconciler

	kubeQOSLevel   map[string]*cgroupReconciler
	podLevel       map[string]*cgroupReconciler
	containerLevel map[string]*cgroupReconciler

	sandboxContainerLevel map[string]*cgroupReconciler
	allPodsLevel          map[string]*cgroupReconciler
}{
	kubeQOSLevel:   map[string]*cgroupReconciler{},
	podLevel:       map[string]*cgroupReconciler{},
	containerLevel: map[string]*cgroupReconciler{},

	sandboxContainerLevel: map[string]*cgroupReconciler{},
	allPodsLevel:          map[string]*cgroupReconciler{},
}

type cgroupReconciler struct {
	cgroupFile  system.Resource
	description map[string]string
	level       ReconcilerLevel
	filter      Filter
	fn          map[string]reconcileFunc
	fn4AllPods  map[string]reconcileFunc4AllPods
}

// Filter & Conditions:
// 1. a condition for one cgroup file should have no more than one filter/index func
// 2. different indexes of one cgroup file can have different reconcile functions
// 3. indexes for one cgroup should be enumerable
type Filter interface {
	Name() string
	Filter(podMeta *statesinformer.PodMeta) string
}

type noneFilter struct{}

const (
	NoneFilterCondition = ""
	NoneFilterName      = "none"
)

func (d *noneFilter) Name() string { _ = "STUB: not implemented"; return "" }

func (d *noneFilter) Filter(podMeta *statesinformer.PodMeta) string {
	_ = "STUB: not implemented"
	return ""
}

var singletonNoneFilter *noneFilter

// NoneFilter returns a Filter which skip filtering anything (into the same condition)
func NoneFilter() Filter { _ = "STUB: not implemented"; return *new(Filter) }

type podQOSFilter struct{}

const (
	PodQOSFilterName = "podQOS"
	HostNetWork      = "hostNetwork"
)

func (p *podQOSFilter) Name() string { _ = "STUB: not implemented"; return "" }

func (p *podQOSFilter) Filter(podMeta *statesinformer.PodMeta) string {
	_ = "STUB: not implemented"
	return ""
}

// consider as LSR if pod is qos=None and has cpuset

type podHostNetworkFilter struct{}

func (p *podHostNetworkFilter) Name() string { _ = "STUB: not implemented"; return "" }

func (p *podHostNetworkFilter) Filter(podMeta *statesinformer.PodMeta) string {
	_ = "STUB: not implemented"
	return ""
}

var singletonPodQOSFilter *podQOSFilter

// PodQOSFilter returns a Filter which filters pod qos class
func PodQOSFilter() Filter { _ = "STUB: not implemented"; return *new(Filter) }

var singletonPodHostNetworkFilter *podHostNetworkFilter

// PodHostNetworkFilter returns a Filter which filters pod hostnetwork is true
func PodHostNetworkFilter() *podHostNetworkFilter { _ = "STUB: not implemented"; return nil }

type podAnnotationResctrlFilter struct{}

const (
	podAnnotationResctrlFilterName = "resctrl"
)

func (p *podAnnotationResctrlFilter) Name() string { _ = "STUB: not implemented"; return "" }

func (p *podAnnotationResctrlFilter) Filter(podMeta *statesinformer.PodMeta) string {
	_ = "STUB: not implemented"
	return ""
}

var singletonPodAnnotationResctrlFilter *podAnnotationResctrlFilter

// PodQOSFilter returns a Filter which filters pod qos class
func PodAnnotationResctrlFilter() *podAnnotationResctrlFilter {
	_ = "STUB: not implemented"
	return nil
}

type reconcileFunc func(protocol.HooksProtocol) error
type reconcileFunc4AllPods func([]protocol.HooksProtocol) error

func RegisterCgroupReconciler4AllPods(level ReconcilerLevel, cgroupFile system.Resource, description string,
	fn reconcileFunc4AllPods, filter Filter, conditions ...string) {
	_ = "STUB: not implemented"
	return
	// default condition
}

// if reconciler exist

// if reconciler not exist

// RegisterCgroupReconciler registers a cgroup reconciler according to the cgroup file, reconcile function and filter
// conditions. A cgroup file of one level can have multiple reconcile functions with different filtered conditions.
//
//	e.g. pod-level cfs_quota can be registered both by cpuset hook and batchresource hook. While cpuset hook reconciles
//	cfs_quota for LSE and LSR pods, batchresource reconciles pods of BE QoS.
//
// TODO: support priority+qos filter.
func RegisterCgroupReconciler(level ReconcilerLevel, cgroupFile system.Resource, description string,
	fn reconcileFunc, filter Filter, conditions ...string) {
	_ = "STUB: not implemented"
	return
	// default condition
}

// if reconciler exist

// if reconciler not exist

type Reconciler interface {
	Run(stopCh <-chan struct{}) error
}

type Context struct {
	StatesInformer    statesinformer.StatesInformer
	Executor          resourceexecutor.ResourceUpdateExecutor
	ReconcileInterval time.Duration
	EventRecorder     record.EventRecorder
}

func NewReconciler(ctx Context) Reconciler { _ = "STUB: not implemented"; return *new(Reconciler) }

// TODO register individual pod event

type reconciler struct {
	podsMutex         sync.RWMutex
	podsMeta          []*statesinformer.PodMeta
	podUpdated        chan struct{}
	executor          resourceexecutor.ResourceUpdateExecutor
	reconcileInterval time.Duration
	eventRecorder     record.EventRecorder
}

func (c *reconciler) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

func (c *reconciler) podRefreshCallback(t statesinformer.RegisterType, o interface{}, target *statesinformer.CallbackTarget) {
	_ = "STUB: not implemented"
	return
}

func (c *reconciler) getPodsMeta() []*statesinformer.PodMeta { _ = "STUB: not implemented"; return nil }

func (c *reconciler) reconcileKubeQOSCgroup(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// TODO refactor kubeqos reconciler, inotify watch corresponding cgroup file and update only when receive modified event
	return
}

func doKubeQOSCgroup(e resourceexecutor.ResourceUpdateExecutor) { _ = "STUB: not implemented"; return }

// all kube qos reconcilers should register in this condition

func (c *reconciler) reconcilePodCgroup(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// TODO refactor pod reconciler, inotify watch corresponding cgroup file and update only when receive modified event
	// new watcher will be added with new pod created, and deleted with pod destroyed
	return
}
