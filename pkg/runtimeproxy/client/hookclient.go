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

package client

import (
	"sync"

	"github.com/golang/groupcache/lru"

	"github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
)

type HookServerClientManagerInterface interface {
	RuntimeHookServerClient(serverPath HookServerPath) (*RuntimeHookClient, error)
}

type HookServerClientManager struct {
	sync.RWMutex
	cache *lru.Cache
}

const (
	defaultCacheSize = 10
)

// NewClientManager
// TODO: garbage client gc
func NewClientManager() *HookServerClientManager { _ = "STUB: not implemented"; return nil }

type HookServerPath struct {
	Path string
	Port int64
}

type RuntimeHookClient struct {
	SockPath string
	v1alpha1.RuntimeHookServiceClient
}

func newRuntimeHookClient(sockPath string) (*RuntimeHookClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cm *HookServerClientManager) RuntimeHookServerClient(serverPath HookServerPath) (*RuntimeHookClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
