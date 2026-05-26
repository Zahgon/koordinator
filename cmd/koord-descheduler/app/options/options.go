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

// Package options provides the descheduler flags
package options

import (
	apiserveroptions "k8s.io/apiserver/pkg/server/options"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/record"
	cliflag "k8s.io/component-base/cli/flag"
	componentbaseconfig "k8s.io/component-base/config"
	"k8s.io/component-base/logs"
	"k8s.io/component-base/metrics"

	deschedulerappconfig "github.com/koordinator-sh/koordinator/cmd/koord-descheduler/app/config"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
)

func newDefaultComponentConfig() (*deschedulerconfig.DeschedulerConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Options has all the params needed to run a Scheduler
type Options struct {
	ComponentConfig *deschedulerconfig.DeschedulerConfiguration

	LeaderElection          *componentbaseconfig.LeaderElectionConfiguration
	SecureServing           *apiserveroptions.SecureServingOptionsWithLoopback
	CombinedInsecureServing *CombinedInsecureServingOptions
	Metrics                 *metrics.Options
	Logs                    *logs.Options

	// ConfigFile is the location of the scheduler server's configuration file.
	ConfigFile string

	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string

	// Flags hold the parsed CLI flags.
	Flags *cliflag.NamedFlagSets

	Client clientset.Interface
}

// NewOptions returns default scheduler app options.
func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

// ApplyLeaderElectionTo obtains the CLI args related with leaderelection, and override the values in `cfg`.
// Then the `cfg` object is injected into the `options` object.
func (o *Options) ApplyLeaderElectionTo(cfg *deschedulerconfig.DeschedulerConfiguration) {
	_ = "STUB: not implemented"
	return
}

// Obtain CLI args related with leaderElection. Set them to `cfg` if specified in command line.

// initFlags initializes flags by section name.
func (o *Options) initFlags() { _ = "STUB: not implemented"; return }

// ApplyTo applies the scheduler options to the given scheduler app configuration.
func (o *Options) ApplyTo(c *deschedulerappconfig.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// use the loaded config file only, with the exception of --address and --port.

// If the --config arg is specified, honor the leader election CLI args only.

// use the loaded config file only, with the exception of --address and --port.

// Validate validates all the required options.
func (o *Options) Validate() []error { _ = "STUB: not implemented"; return nil }

// Config return a scheduler config object
func (o *Options) Config() (*deschedulerappconfig.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prepare kube config.

// Prepare kube clients.

// Set up leader election if enabled.

// Use the scheduler name in the first profile to record leader election.

// makeLeaderElectionConfig builds a leader election configuration. It will
// create a new resource lock associated with the configuration.
func makeLeaderElectionConfig(config componentbaseconfig.LeaderElectionConfiguration, kubeConfig *restclient.Config, recorder record.EventRecorder) (*leaderelection.LeaderElectionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add a uniquifier so that two processes on the same host don't accidentally both become active

// createKubeConfig creates a kubeConfig from the given config and masterOverride.
func createKubeConfig(config componentbaseconfig.ClientConnectionConfiguration, masterOverride string) (*restclient.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createClients creates a kube client and an event client from the given kubeConfig
func createClients(kubeConfig *restclient.Config) (clientset.Interface, clientset.Interface, error) {
	_ = "STUB: not implemented"
	return *new(clientset.Interface), *new(clientset.Interface), nil
}
