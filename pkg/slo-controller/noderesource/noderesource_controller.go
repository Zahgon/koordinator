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

package noderesource

import (
	"context"
	"flag"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/clock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/pkg/slo-controller/config"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

const (
	Name = "noderesource"

	disableInConfig string = "DisableInConfig"
)

var (
	NodeResourcePlugins []string
	AllPlugins          []string
)

type NodeResourceReconciler struct {
	Client          client.Client
	Recorder        record.EventRecorder
	Scheme          *runtime.Scheme
	Clock           clock.Clock
	NodeSyncContext *framework.SyncContext
	GPUSyncContext  *framework.SyncContext
	cfgCache        config.ColocationCfgCache
}

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;patch
// +kubebuilder:rbac:groups=core,resources=nodes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=scheduling.koordinator.sh,resources=devices,verbs=get;list;watch
// +kubebuilder:rbac:groups=slo.koordinator.sh,resources=nodemetrics,verbs=get;list;watch

func (r *NodeResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// skip non-existing node and return no error to forget the request

// disable all resources

// the node metric might be not exist or abnormal, resource calculation should handle this case

// calculate node resources

// update node status

// do other node updates.

func InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func addPluginOption(plugin framework.Plugin, enabled bool) { _ = "STUB: not implemented"; return }

func isPluginEnabled(pluginName string) bool { _ = "STUB: not implemented"; return false }

func Add(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	// init plugins for NodeResource
	return nil
}

func (r *NodeResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// avoid conflict with others reconciling `Node`

// setup plugins
// allow plugins to mutate controller via the builder
