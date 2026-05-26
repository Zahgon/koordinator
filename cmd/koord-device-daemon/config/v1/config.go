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
	"io"

	"github.com/urfave/cli/v2"
)

// Version indicates the version of the 'Config' struct used to hold configuration information.
const Version = "v1"

// Config is a versioned struct used to hold configuration information.
type Config struct {
	Version string `json:"version"             yaml:"version"`
	Flags   Flags  `json:"flags,omitempty"     yaml:"flags,omitempty"`
}

// NewConfig builds out a Config struct from a config file (or command line flags).
// The data stored in the config will be populated in order of precedence from
// (1) command line, (2) environment variable, (3) config file.
func NewConfig(c *cli.Context, flags []cli.Flag) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseConfig parses a config file as either YAML of JSON and unmarshals it into a Config struct.
func parseConfig(configFile string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func parseConfigFrom(reader io.Reader) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
