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

package v1

import (
	"github.com/urfave/cli/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// prt returns a reference to whatever type is passed into it
func ptr[T any](x T) *T {
	_ = "STUB: not implemented"

	// updateFromCLIFlag conditionally updates the config flag at 'pflag' to the value of the CLI flag with name 'flagName'
	return nil
}

func updateFromCLIFlag[T any](pflag **T, c *cli.Context, flagName string) {
	_ = "STUB: not implemented"
	return
}

// Flags holds the full list of flags used to configure the device plugin and KDD.
type Flags struct {
	CommandLineFlags
}

// CommandLineFlags holds the list of command line flags used to configure the device plugin and KDD.
type CommandLineFlags struct {
	KDD *KDDCommandLineFlags `json:"kdd,omitempty" yaml:"kdd,omitempty"`
}

// KDDCommandLineFlags holds the list of command line flags specific to KDD.
type KDDCommandLineFlags struct {
	Oneshot          *bool            `json:"oneshot"         yaml:"oneshot"`
	SleepInterval    *metav1.Duration `json:"sleepInterval"   yaml:"sleepInterval"`
	PrintsOutputFile *string          `json:"printsOutputFile" yaml:"printsOutputFile"`
}

// UpdateFromCLIFlags updates Flags from settings in the cli Flags if they are set.
func (f *Flags) UpdateFromCLIFlags(c *cli.Context, flags []cli.Flag) {
	_ = "STUB: not implemented"
	return
}

// KDD specific flags
