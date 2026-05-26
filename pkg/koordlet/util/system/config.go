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

package system

import (
	"flag"
	"os"

	"go.uber.org/atomic"
)

const (
	DS_MODE   = "dsMode"
	HOST_MODE = "hostMode"
)

var Conf = NewDsModeConfig()
var AgentMode = DS_MODE

var UseCgroupsV2 = atomic.NewBool(false)

type Config struct {
	CgroupRootDir         string
	SysRootDir            string
	SysFSRootDir          string
	ProcRootDir           string
	VarRunRootDir         string
	VarLibKubeletRootDir  string
	RunRootDir            string
	RuntimeHooksConfigDir string

	ContainerdEndPoint           string
	PouchEndpoint                string
	DockerEndPoint               string
	CrioEndPoint                 string
	DefaultRuntimeType           string
	HAMICoreLibraryDirectoryPath string
	PodResourcesProxyPath        string
	XPUDeviceInfosDir            string
}

func init() {
	agentMode := os.Getenv("agent_mode")
	if agentMode == HOST_MODE {
		Conf = NewHostModeConfig()
		AgentMode = agentMode
	}
}

// InitSupportConfigs initializes the system support status.
// e.g. the cgroup version, resctrl capability
func InitSupportConfigs() {
	_ = "STUB: not implemented"
	// $ getconf CLK_TCK > jiffies
	return
}

func NewHostModeConfig() *Config { _ = "STUB: not implemented"; return nil }

func NewDsModeConfig() *Config { _ = "STUB: not implemented"; return nil }

// some dirs are not covered by ns, or unused with `hostPID` is on

func SetConf(config Config) { _ = "STUB: not implemented"; return }

func (c *Config) InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }
