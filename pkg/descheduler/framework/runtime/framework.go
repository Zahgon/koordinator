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

package runtime

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/events"

	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

type frameworkImpl struct {
	dryRun                    bool
	clientSet                 clientset.Interface
	kubeConfig                *restclient.Config
	eventRecorder             events.EventRecorder
	evictionLimiter           EvictionLimiter
	sharedInformerFactory     informers.SharedInformerFactory
	getPodsAssignedToNodeFunc framework.GetPodsAssignedToNodeFunc
	deschedulePlugins         []framework.DeschedulePlugin
	balancePlugins            []framework.BalancePlugin
	evictPlugins              []framework.EvictPlugin
	filterPlugins             []framework.FilterPlugin
	nodeSelector              *metav1.LabelSelector
}

// Option for the frameworkImpl.
type Option func(*frameworkOptions)

type frameworkOptions struct {
	dryRun                    bool
	clientSet                 clientset.Interface
	kubeConfig                *restclient.Config
	eventRecorder             events.EventRecorder
	sharedInformerFactory     informers.SharedInformerFactory
	getPodsAssignedToNodeFunc framework.GetPodsAssignedToNodeFunc
	evictionLimiter           EvictionLimiter
	captureProfile            CaptureProfile
}

func WithDryRun(dryRun bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClientSet sets clientSet for the scheduling Framework.
func WithClientSet(clientSet clientset.Interface) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithKubeConfig sets kubeConfig for the scheduling frameworkImpl.
func WithKubeConfig(kubeConfig *restclient.Config) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSharedInformerFactory(sharedInformerFactory informers.SharedInformerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGetPodsAssignedToNodeFunc(fn framework.GetPodsAssignedToNodeFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// CaptureProfile is a callback to capture a finalized profile.
type CaptureProfile func(profile deschedulerconfig.DeschedulerProfile)

// WithCaptureProfile sets a callback to capture the finalized profile.
func WithCaptureProfile(c CaptureProfile) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEventRecorder sets clientSet for the scheduling frameworkImpl.
func WithEventRecorder(recorder events.EventRecorder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEvictionLimiter(limiter EvictionLimiter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewFramework(ctx context.Context, r Registry, profile *deschedulerconfig.DeschedulerProfile, opts ...Option) (framework.Handle, error) {
	_ = "STUB: not implemented"
	return *new(framework.Handle), nil
}

func (f *frameworkImpl) initPlugins(ctx context.Context, r Registry, pluginConfig map[string]runtime.Object, extensionPoints []extensionPoint, pluginsMap map[string]framework.Plugin) ([]deschedulerconfig.PluginConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize only needed plugins.

// initialize plugins that have not yet been created

// initialize plugins per individual extension points

func updatePluginList(pluginList interface{}, pluginSet deschedulerconfig.PluginSet, pluginsMap map[string]framework.Plugin) error {
	_ = "STUB: not implemented"
	return nil
}

// extensionPoint encapsulates desired and applied set of plugins at a specific extension
// point. This is used to simplify iterating over all extension points supported by the
// frameworkImpl.
type extensionPoint struct {
	// the set of plugins to be configured at this extension point.
	plugins *deschedulerconfig.PluginSet
	// a pointer to the slice storing plugins implementations that will run at this
	// extension point.
	slicePtr interface{}
}

func (f *frameworkImpl) getExtensionPoints(plugins *deschedulerconfig.Plugins) []extensionPoint {
	_ = "STUB: not implemented"
	return nil
}

func pluginsNeeded(pgSet sets.String, points []extensionPoint) { _ = "STUB: not implemented"; return }

func (f *frameworkImpl) ClientSet() clientset.Interface {
	_ = "STUB: not implemented"
	return *new(clientset.Interface)
}

func (f *frameworkImpl) KubeConfig() *restclient.Config { _ = "STUB: not implemented"; return nil }

func (f *frameworkImpl) EventRecorder() events.EventRecorder {
	_ = "STUB: not implemented"
	return *new(events.EventRecorder)
}

func (f *frameworkImpl) Evictor() framework.Evictor {
	_ = "STUB: not implemented"
	return *new(framework.Evictor)
}

func (f *frameworkImpl) GetPodsAssignedToNodeFunc() framework.GetPodsAssignedToNodeFunc {
	_ = "STUB: not implemented"
	return *new(framework.GetPodsAssignedToNodeFunc)
}

func (f *frameworkImpl) SharedInformerFactory() informers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(informers.SharedInformerFactory)
}

func (f *frameworkImpl) NodeSelector() *metav1.LabelSelector { _ = "STUB: not implemented"; return nil }

func (f *frameworkImpl) RunDeschedulePlugins(ctx context.Context, nodes []*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}

func (f *frameworkImpl) RunBalancePlugins(ctx context.Context, nodes []*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}
