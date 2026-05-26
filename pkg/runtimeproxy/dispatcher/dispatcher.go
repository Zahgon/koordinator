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

package dispatcher

import (
	"context"

	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/client"
	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/config"
)

// RuntimeHookDispatcher dispatches hook request to RuntimeHookServer(e.g. koordlet)
type RuntimeHookDispatcher struct {
	cm          client.HookServerClientManagerInterface
	hookManager config.ManagerInterface
}

func NewRuntimeDispatcher() *RuntimeHookDispatcher { _ = "STUB: not implemented"; return nil }

func (rd *RuntimeHookDispatcher) dispatchInternal(ctx context.Context, hookType config.RuntimeHookType,
	client *client.RuntimeHookClient, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rd *RuntimeHookDispatcher) Dispatch(ctx context.Context, runtimeRequestPath config.RuntimeRequestPath,
	stage config.RuntimeHookStage, request interface{}) (interface{}, error, config.FailurePolicyType) {
	_ = "STUB: not implemented"
	return nil, nil, *new(config.FailurePolicyType)
}

// currently, only one hook be called during one runtime
// TODO: multi hook server to merge response
