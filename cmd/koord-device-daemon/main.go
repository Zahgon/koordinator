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

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"k8s.io/klog/v2"

	resourceconifg "github.com/koordinator-sh/koordinator/cmd/koord-device-daemon/config/v1"
	printmanager "github.com/koordinator-sh/koordinator/pkg/device-daemon/printer"
	"github.com/koordinator-sh/koordinator/pkg/device-daemon/resource"
)

// Config represents a collection of config options for KDD.
type Config struct {
	configFile string
	// flags stores the CLI flags for later processing.
	flags []cli.Flag
}

func main() {
	klog.InitFlags(nil)
	// Opt into the new klog behavior so that -stderrthreshold is honored even
	// when -logtostderr=true (the default).
	// Ref: kubernetes/klog#212, kubernetes/klog#432
	flag.Set("legacy_stderr_threshold_behavior", "false") //nolint:errcheck
	flag.Set("stderrthreshold", "INFO")                   //nolint:errcheck
	config := &Config{}

	c := cli.NewApp()
	c.Name = "koord device daemon"
	c.Usage = "generate device infos  for heterogeneous devices"
	c.Action = func(ctx *cli.Context) error {
		return start(ctx, config)
	}

	c.Before = func(ctx *cli.Context) error {
		v := ctx.String("v")
		if err := flag.Set("v", v); err != nil {
			return fmt.Errorf("failed to set klog verbosity level: %w", err)
		}
		klog.V(2).InfoS("klog verbosity level set", "level", v)
		return nil
	}

	config.flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "v",
			Value: "1",
			Usage: "klog verbosity level (e.g. 2 for Info, 5 for Debug)",
		},
		&cli.BoolFlag{
			Name:    "oneshot",
			Value:   false,
			Usage:   "Label once and exit",
			EnvVars: []string{"KDD_ONESHOT"},
		},
		&cli.BoolFlag{
			Name:    "no-timestamp",
			Value:   false,
			Usage:   "Do not add the timestamp to the labels",
			EnvVars: []string{"KDD_NO_TIMESTAMP"},
		},
		&cli.DurationFlag{
			Name:    "sleep-interval",
			Value:   900 * time.Second,
			Usage:   "Time to sleep between labeling",
			EnvVars: []string{"KDD_SLEEP_INTERVAL"},
		},
		&cli.StringFlag{
			Name:    "prints-output-file",
			Value:   "/var/run/koordlet/xpu-device-infos/ic-device",
			EnvVars: []string{"KDD_PRINTS_OUTPUT_FILE", "PRINTS_OUTPUT_FILE"},
		},
		&cli.StringFlag{
			Name:        "config-file",
			Usage:       "the path to a config file as an alternative to command line options or environment variables",
			Destination: &config.configFile,
			EnvVars:     []string{"KDD_CONFIG_FILE", "CONFIG_FILE"},
		},
	}

	c.Flags = config.flags

	if err := c.Run(os.Args); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}

// loadConfig loads the config from the spec file.
func (cfg *Config) loadConfig(c *cli.Context) (*resourceconifg.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func start(c *cli.Context, cfg *Config) error { _ = "STUB: not implemented"; return nil }

// Load the configuration file

// Print the config to the output.

type resourceFeatureDiscovery struct {
	manager        map[string]resource.Manager
	config         *resourceconifg.Config
	PrintsOutputer printmanager.Writer
}

func (rfd *resourceFeatureDiscovery) run(sigs chan os.Signal) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Watch for any signals from the OS. On SIGHUP trigger a reload of the config.
// On all other signals, exit the loop and exit the program.

func removeOutputFile(path string) error { _ = "STUB: not implemented"; return nil }

// signals creats a channel for the specified signals.
func signals(sigs ...os.Signal) chan os.Signal { _ = "STUB: not implemented"; return nil }
