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

package options

import (
	"net"

	"github.com/spf13/pflag"
	apiserver "k8s.io/apiserver/pkg/server"

	schedulerappconfig "github.com/koordinator-sh/koordinator/cmd/koord-scheduler/app/config"
)

// DeprecatedInsecureServingOptions provides options for insecure serving (no TLS).
// DEPRECATED: no longer supported.
type DeprecatedInsecureServingOptions struct {
	BindAddress net.IP
	BindPort    int
	BindNetwork string
	Listener    net.Listener
}

// ApplyTo populates the DeprecatedInsecureServingInfo from the options.
func (s *DeprecatedInsecureServingOptions) ApplyTo(c **apiserver.DeprecatedInsecureServingInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// CombinedInsecureServingOptions sets up to two insecure listeners for healthz and metrics. The flags
// override the ComponentConfig and DeprecatedInsecureServingOptions values for both.
type CombinedInsecureServingOptions struct {
	Healthz *DeprecatedInsecureServingOptions

	BindPort    int
	BindAddress string
}

// AddFlags adds flags for the insecure serving options.
func (o *CombinedInsecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

// ApplyTo applies the insecure serving options to the given scheduler app configuration, and updates the componentConfig.
func (o *CombinedInsecureServingOptions) ApplyTo(c *schedulerappconfig.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate validates the insecure serving options.
func (o *CombinedInsecureServingOptions) Validate() []error { _ = "STUB: not implemented"; return nil }
