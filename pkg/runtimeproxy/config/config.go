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

package config

type FailurePolicyType string

const (
	// PolicyFail returns error to caller when got an error cri hook server
	PolicyFail FailurePolicyType = "Fail"
	// PolicyIgnore transfer cri request to containerd/dockerd when got an error to cri serer
	PolicyIgnore FailurePolicyType = "Ignore"
	// PolicyNone when no Policy configured. Proxy would ignore errors for PolicyNone like PolicyIgnore.
	PolicyNone = ""
)

func GetFailurePolicyType(typeString string) (FailurePolicyType, error) {
	_ = "STUB: not implemented"
	return *new(FailurePolicyType), nil
}

type RuntimeHookType string

const (
	defaultRuntimeHookConfigPath string = "/etc/runtime/hookserver.d"
)

const (
	PreRunPodSandbox            RuntimeHookType = "PreRunPodSandbox"
	PostStopPodSandbox          RuntimeHookType = "PostStopPodSandbox"
	PreCreateContainer          RuntimeHookType = "PreCreateContainer"
	PreStartContainer           RuntimeHookType = "PreStartContainer"
	PostStartContainer          RuntimeHookType = "PostStartContainer"
	PreUpdateContainerResources RuntimeHookType = "PreUpdateContainerResources"
	PostStopContainer           RuntimeHookType = "PostStopContainer"
	PreRemoveRunPodSandbox      RuntimeHookType = "PreRemoveRunPodSandbox"
	NoneRuntimeHookType         RuntimeHookType = "NoneRuntimeHookType"
)

type RuntimeHookConfig struct {
	RemoteEndpoint string            `json:"remote-endpoint,omitempty"`
	FailurePolicy  FailurePolicyType `json:"failure-policy,omitempty"`
	RuntimeHooks   []RuntimeHookType `json:"runtime-hooks,omitempty"`
}

type RuntimeRequestPath string

const (
	RunPodSandbox            RuntimeRequestPath = "RunPodSandbox"
	StopPodSandbox           RuntimeRequestPath = "StopPodSandbox"
	CreateContainer          RuntimeRequestPath = "CreateContainer"
	StartContainer           RuntimeRequestPath = "StartContainer"
	UpdateContainerResources RuntimeRequestPath = "UpdateContainerResources"
	StopContainer            RuntimeRequestPath = "StopContainer"
	NoneRuntimeHookPath      RuntimeRequestPath = "NoneRuntimeHookPath"
)

func (ht RuntimeHookType) OccursOn(path RuntimeRequestPath) bool {
	_ = "STUB: not implemented"
	return false
}

func (hp RuntimeRequestPath) PreHookType() RuntimeHookType {
	_ = "STUB: not implemented"
	return *new(RuntimeHookType)
}

func (hp RuntimeRequestPath) PostHookType() RuntimeHookType {
	_ = "STUB: not implemented"
	return *new(RuntimeHookType)
}

type RuntimeHookStage string

const (
	PreHook     RuntimeHookStage = "PreHook"
	PostHook    RuntimeHookStage = "PostHook"
	UnknownHook RuntimeHookStage = "UnknownHook"
)

func (ht RuntimeHookType) HookStage() RuntimeHookStage {
	_ = "STUB: not implemented"
	return *new(RuntimeHookStage)
}
