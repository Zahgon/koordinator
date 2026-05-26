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

package gpudeviceresource

import (
	corev1 "k8s.io/api/core/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

const PluginName = "GPUDeviceResource"

const (
	ResetResourcesMsg  = "reset node gpu resources"
	UpdateResourcesMsg = "node gpu resources from device"
	UpdateLabelsMsg    = "node gpu labels from device"

	NeedSyncForResourceDiffMsg = "gpu resource diff is big than threshold"
	NeedSyncForGPUModelMsgFmt  = "gpu device label %s changed"
)

var client ctrlclient.Client

type Plugin struct{}

func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
	// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;patch
	// +kubebuilder:rbac:groups=core,resources=nodes/status,verbs=get;update;patch
	// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
	// +kubebuilder:rbac:groups=scheduling.koordinator.sh,resources=devices,verbs=get;list;watch
	// +kubebuilder:rbac:groups=topology.node.k8s.io,resources=noderesourcetopologies,verbs=get;list;watch;create;update
	return ""
}

func (p *Plugin) Setup(opt *framework.Option) error {
	_ = "STUB: not implemented"

	// schedulingv1alpha1.AddToScheme(opt.Scheme)
	return nil
}

func (p *Plugin) NeedSync(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (p *Plugin) NeedSyncMeta(_ *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (p *Plugin) Prepare(_ *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	// prepare node resources
	return nil
}

// ignore missing resources
// TBD: shall we remove the resource when some resource types are missing

// prepare node labels
// TBD: shall we reset labels if not exist in the NR

func (p *Plugin) Reset(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Calculate(_ *configuration.ColocationStrategy, node *corev1.Node, _ *corev1.PodList, _ *framework.ResourceMetrics) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// calculate device resources

// device not found, reset gpu resources on node

// TODO: calculate NUMA-level resources against NRT

func (p *Plugin) calculate(node *corev1.Node, device *schedulingv1alpha1.Device) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// calculate gpu resources

// Currently only nvidia gpu supports the koordinator.sh/gpu wrapper resource,
// also keep empty vendor behavior as is for compatibility.

// FIXME: shall we add node resources in devices but not in ResourceNames?

// calculate labels about gpu driver and model

func (p *Plugin) resetGPUNodeResource() ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: shall we reset node resources in devices but not in ResourceNames?
