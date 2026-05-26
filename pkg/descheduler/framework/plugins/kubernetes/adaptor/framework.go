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

package adaptor

import (
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	metricscollector "sigs.k8s.io/descheduler/pkg/descheduler/metricscollector"
	podutil "sigs.k8s.io/descheduler/pkg/descheduler/pod"
	k8sdeschedulerframework "sigs.k8s.io/descheduler/pkg/framework/types"

	promapi "github.com/prometheus/client_golang/api"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

var _ k8sdeschedulerframework.Handle = &frameworkHandleAdaptor{}

type frameworkHandleAdaptor struct {
	handle framework.Handle
}

func NewFrameworkHandleAdaptor(handle framework.Handle) k8sdeschedulerframework.Handle {
	_ = "STUB: not implemented"
	return *new(k8sdeschedulerframework.Handle)
}

// ClientSet returns a kubernetes clientSet.
func (a *frameworkHandleAdaptor) ClientSet() clientset.Interface {
	_ = "STUB: not implemented"
	return *new(clientset.Interface)
}

func (a *frameworkHandleAdaptor) Evictor() k8sdeschedulerframework.Evictor {
	_ = "STUB: not implemented"
	return *new(k8sdeschedulerframework.Evictor)
}

func (a *frameworkHandleAdaptor) GetPodsAssignedToNodeFunc() podutil.GetPodsAssignedToNodeFunc {
	_ = "STUB: not implemented"
	return *new(podutil.GetPodsAssignedToNodeFunc)
}

func (a *frameworkHandleAdaptor) SharedInformerFactory() informers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(informers.SharedInformerFactory)
}

// MetricsCollector returns nil as this adaptor does not support metrics collection.
func (a *frameworkHandleAdaptor) MetricsCollector() *metricscollector.MetricsCollector {
	_ = "STUB: not implemented"

	// PluginInstanceID returns a unique identifier for this plugin instance.
	return nil
}

func (a *frameworkHandleAdaptor) PluginInstanceID() string {
	_ = "STUB: not implemented"

	// PrometheusClient returns nil as this adaptor does not support prometheus.
	return ""
}

func (a *frameworkHandleAdaptor) PrometheusClient() promapi.Client {
	_ = "STUB: not implemented"
	return *new(promapi.Client)
}
