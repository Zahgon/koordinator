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

package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/nodeutilization"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/podlifetime"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removeduplicates"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removefailedpods"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removepodshavingtoomanyrestarts"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removepodsviolatinginterpodantiaffinity"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removepodsviolatingnodeaffinity"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removepodsviolatingnodetaints"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/removepodsviolatingtopologyspreadconstraint"
	k8sdeschedulerframework "sigs.k8s.io/descheduler/pkg/framework/types"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
	frameworkruntime "github.com/koordinator-sh/koordinator/pkg/descheduler/framework/runtime"
)

type (
	PluginFactory       = func(ctx context.Context, args runtime.Object, handle k8sdeschedulerframework.Handle) (k8sdeschedulerframework.Plugin, error)
	PluginArgsDefaulter = func(obj runtime.Object)
	PluginArgsValidator = func(obj runtime.Object) error
)

type PluginDescriptor struct {
	Name          string
	Factory       PluginFactory
	ArgsPrototype runtime.Object
	ArgsDefaulter PluginArgsDefaulter
	ArgsValidator PluginArgsValidator
}

var Plugins = []PluginDescriptor{
	{
		Name:          nodeutilization.HighNodeUtilizationPluginName,
		Factory:       nodeutilization.NewHighNodeUtilization,
		ArgsPrototype: &nodeutilization.HighNodeUtilizationArgs{},
		ArgsDefaulter: nodeutilization.SetDefaults_HighNodeUtilizationArgs,
		ArgsValidator: nodeutilization.ValidateHighNodeUtilizationArgs,
	},
	{
		Name:          nodeutilization.LowNodeUtilizationPluginName,
		Factory:       nodeutilization.NewLowNodeUtilization,
		ArgsPrototype: &nodeutilization.LowNodeUtilizationArgs{},
		ArgsDefaulter: nodeutilization.SetDefaults_LowNodeUtilizationArgs,
		ArgsValidator: nodeutilization.ValidateLowNodeUtilizationArgs,
	},
	{
		Name:          podlifetime.PluginName,
		Factory:       podlifetime.New,
		ArgsPrototype: &podlifetime.PodLifeTimeArgs{},
		ArgsDefaulter: podlifetime.SetDefaults_PodLifeTimeArgs,
		ArgsValidator: podlifetime.ValidatePodLifeTimeArgs,
	},
	{
		Name:          removefailedpods.PluginName,
		Factory:       removefailedpods.New,
		ArgsPrototype: &removefailedpods.RemoveFailedPodsArgs{},
		ArgsDefaulter: removefailedpods.SetDefaults_RemoveFailedPodsArgs,
		ArgsValidator: removefailedpods.ValidateRemoveFailedPodsArgs,
	},
	{
		Name:          removeduplicates.PluginName,
		Factory:       removeduplicates.New,
		ArgsPrototype: &removeduplicates.RemoveDuplicatesArgs{},
		ArgsDefaulter: removeduplicates.SetDefaults_RemoveDuplicatesArgs,
		ArgsValidator: removeduplicates.ValidateRemoveDuplicatesArgs,
	},
	{
		Name:          removepodshavingtoomanyrestarts.PluginName,
		Factory:       removepodshavingtoomanyrestarts.New,
		ArgsPrototype: &removepodshavingtoomanyrestarts.RemovePodsHavingTooManyRestartsArgs{},
		ArgsDefaulter: removepodshavingtoomanyrestarts.SetDefaults_RemovePodsHavingTooManyRestartsArgs,
		ArgsValidator: removepodshavingtoomanyrestarts.ValidateRemovePodsHavingTooManyRestartsArgs,
	},
	{
		Name:          removepodsviolatinginterpodantiaffinity.PluginName,
		Factory:       removepodsviolatinginterpodantiaffinity.New,
		ArgsPrototype: &removepodsviolatinginterpodantiaffinity.RemovePodsViolatingInterPodAntiAffinityArgs{},
		ArgsDefaulter: removepodsviolatinginterpodantiaffinity.SetDefaults_RemovePodsViolatingInterPodAntiAffinityArgs,
		ArgsValidator: removepodsviolatinginterpodantiaffinity.ValidateRemovePodsViolatingInterPodAntiAffinityArgs,
	},
	{
		Name:          removepodsviolatingnodeaffinity.PluginName,
		Factory:       removepodsviolatingnodeaffinity.New,
		ArgsPrototype: &removepodsviolatingnodeaffinity.RemovePodsViolatingNodeAffinityArgs{},
		ArgsDefaulter: removepodsviolatingnodeaffinity.SetDefaults_RemovePodsViolatingNodeAffinityArgs,
		ArgsValidator: removepodsviolatingnodeaffinity.ValidateRemovePodsViolatingNodeAffinityArgs,
	},

	{
		Name:          removepodsviolatingnodetaints.PluginName,
		Factory:       removepodsviolatingnodetaints.New,
		ArgsPrototype: &removepodsviolatingnodetaints.RemovePodsViolatingNodeTaintsArgs{},
		ArgsDefaulter: removepodsviolatingnodetaints.SetDefaults_RemovePodsViolatingNodeTaintsArgs,
		ArgsValidator: removepodsviolatingnodetaints.ValidateRemovePodsViolatingNodeTaintsArgs,
	},
	{
		Name:          removepodsviolatingtopologyspreadconstraint.PluginName,
		Factory:       removepodsviolatingtopologyspreadconstraint.New,
		ArgsPrototype: &removepodsviolatingtopologyspreadconstraint.RemovePodsViolatingTopologySpreadConstraintArgs{},
		ArgsDefaulter: removepodsviolatingtopologyspreadconstraint.SetDefaults_RemovePodsViolatingTopologySpreadConstraintArgs,
		ArgsValidator: removepodsviolatingtopologyspreadconstraint.ValidateRemovePodsViolatingTopologySpreadConstraintArgs,
	},
}

func SetupK8sDeschedulerPlugins(registry frameworkruntime.Registry) {
	_ = "STUB: not implemented"
	return
}

func (d *PluginDescriptor) New(ctx context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

var _ framework.DeschedulePlugin = &PluginAdaptor{}
var _ framework.BalancePlugin = &PluginAdaptor{}

type PluginAdaptor struct {
	descriptor *PluginDescriptor
	pl         k8sdeschedulerframework.Plugin
}

func (a *PluginAdaptor) Name() string { _ = "STUB: not implemented"; return "" }

func (a *PluginAdaptor) Deschedule(ctx context.Context, nodes []*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}

func (a *PluginAdaptor) Balance(ctx context.Context, nodes []*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}
