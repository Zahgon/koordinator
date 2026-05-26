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

package descheduler

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	corev1informers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
	frameworkruntime "github.com/koordinator-sh/koordinator/pkg/descheduler/framework/runtime"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/profile"
)

type Descheduler struct {
	// Profiles are the descheduling profiles.
	Profiles profile.Map

	// Close this to shut down the scheduler.
	StopEverything <-chan struct{}

	clientSet    clientset.Interface
	nodeInformer corev1informers.NodeInformer

	deschedulingInterval time.Duration
	nodeSelector         string
	evictionLimiter      frameworkruntime.EvictionLimiter
}

type deschedulerOptions struct {
	componentConfigVersion string
	kubeConfig             *restclient.Config
	frameworkCapturer      FrameworkCapturer
	podAssignedToNodeFn    PodAssignedToNodeFn
	outOfTreeRegistry      frameworkruntime.Registry
	profiles               []deschedulerconfig.DeschedulerProfile
	applyDefaultProfile    bool
	dryRun                 bool
	deschedulingInterval   time.Duration
	nodeSelector           *metav1.LabelSelector
	evictionLimiter        frameworkruntime.EvictionLimiter
}

// Option configures a Scheduler
type Option func(*deschedulerOptions)

// WithComponentConfigVersion sets the component config version to the
// DeschedulerConfiguration version used. The string should be the full
// scheme group/version of the external type we converted from (for example
// "descheduler/v1alpha2")
func WithComponentConfigVersion(apiVersion string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithKubeConfig(cfg *restclient.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProfiles(p ...deschedulerconfig.DeschedulerProfile) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDryRun(dryRun bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNodeSelector(nodeSelector *metav1.LabelSelector) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDeschedulingInterval(interval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFrameworkOutOfTreeRegistry sets the registry for out-of-tree plugins. Those plugins
// will be appended to the default registry.
func WithFrameworkOutOfTreeRegistry(registry frameworkruntime.Registry) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// FrameworkCapturer is used for registering a notify function in building framework.
type FrameworkCapturer func(deschedulerconfig.DeschedulerProfile)

// WithBuildFrameworkCapturer sets a notify function for getting buildFramework details.
func WithBuildFrameworkCapturer(fc FrameworkCapturer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type PodAssignedToNodeFn func(nodeName string) ([]*corev1.Pod, error)

func WithPodAssignedToNodeFn(fn PodAssignedToNodeFn) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEvictionLimiter(limiter frameworkruntime.EvictionLimiter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

var defaultDeschedulerOptions = deschedulerOptions{
	applyDefaultProfile: true,
}

func New(client clientset.Interface,
	informerFactory informers.SharedInformerFactory,
	dynInformerFactory dynamicinformer.DynamicSharedInformerFactory,
	recorderFactory profile.RecorderFactory,
	stopCh <-chan struct{},
	opts ...Option,
) (*Descheduler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create the informers before starting the informer factory

func (d *Descheduler) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// If there was no interval specified, send a signal to the stopChannel to end the wait.Until loop after 1 iteration

func (d *Descheduler) deschedulerOnce(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func podAssignedToNodeAdaptor(fn PodAssignedToNodeFn) framework.GetPodsAssignedToNodeFunc {
	_ = "STUB: not implemented"
	return *new(framework.GetPodsAssignedToNodeFunc)
}

func filterNodes(nodeSelector *metav1.LabelSelector, nodes []*corev1.Node, processedNodes sets.String) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
