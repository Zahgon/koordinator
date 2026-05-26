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

// Package app implements a Server object for running the scheduler.
package app

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/apiserver/pkg/server/mux"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	logsapi "k8s.io/component-base/logs/api/v1"
	"k8s.io/component-base/metrics/features"
	"sigs.k8s.io/controller-runtime/pkg/client"

	deschedulerappconfig "github.com/koordinator-sh/koordinator/cmd/koord-descheduler/app/config"
	"github.com/koordinator-sh/koordinator/cmd/koord-descheduler/app/options"
	"github.com/koordinator-sh/koordinator/pkg/descheduler"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	frameworkruntime "github.com/koordinator-sh/koordinator/pkg/descheduler/framework/runtime"
)

func init() {
	utilruntime.Must(logsapi.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
	utilruntime.Must(features.AddFeatureGates(utilfeature.DefaultMutableFeatureGate))
}

// Option configures a framework.Registry.
type Option func(frameworkruntime.Registry) error

// NewDeschedulerCommand creates a *cobra.Command object with default parameters and registryOptions
func NewDeschedulerCommand(registryOptions ...Option) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// runCommand runs the scheduler.
func runCommand(cmd *cobra.Command, opts *options.Options, registryOptions ...Option) error {
	_ = "STUB: not implemented"
	// Activate logging as soon as possible, after that
	// show flags with the final logging configuration.
	return nil
}

// Run executes the scheduler based on the given configuration. It only returns on error or when context is done.
func Run(ctx context.Context, cc *deschedulerappconfig.CompletedConfig, desched *descheduler.Descheduler) error {
	_ = "STUB: not implemented"
	// To help debugging, immediately log version
	return nil
}

// Configz registration.

// Prepare the event broadcaster.

// Setup healthz checks.

// Start up the healthz server.

// Start up the healthz server.

// fail early for secure handlers, removing the old error loop from above

// If leader election is enabled, runCommand via LeaderElector until done and exit.

// We were asked to terminate. Exit 0.

// We lost the lock.

func StartDescheduler(ctx context.Context, cc *deschedulerappconfig.CompletedConfig) {
	_ = "STUB: not implemented"
	// We don't need to start cc.InformerFactory because we adapt the InformerFactory to controller.Manager
	// if cc.InformerFactory != nil {
	// 	cc.InformerFactory.Start(ctx.Done())
	// }
	return
}

// buildHandlerChain wraps the given handler with the standard filters.
func buildHandlerChain(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func installMetricHandler(pathRecorderMux *mux.PathRecorderMux) { _ = "STUB: not implemented"; return }

// newHealthzAndMetricsHandler creates a healthz server from the config, and will also
// embed the metrics handler.
func newHealthzAndMetricsHandler(config *deschedulerconfig.DeschedulerConfiguration, checks ...healthz.HealthChecker) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Setup creates a completed config and a scheduler based on the command args and options
func Setup(ctx context.Context, opts *options.Options, outOfTreeRegistryOptions ...Option) (*deschedulerappconfig.CompletedConfig, *descheduler.Descheduler, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get the completed config

// Ensure DetectorCacheTimeout >= DeschedulingInterval for LowNodeLoad plugins.
// If DetectorCacheTimeout is smaller than DeschedulingInterval, the anomaly
// detector cache entries expire between descheduling cycles, causing evictions
// to silently stop working because the detector never accumulates enough
// consecutive abnormalities.

func podAssignedToNode(clt client.Client) descheduler.PodAssignedToNodeFn {
	_ = "STUB: not implemented"
	return *new(descheduler.PodAssignedToNodeFn)
}
