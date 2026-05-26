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

package cache

import (
	"sync"
	"time"
)

const (
	defaultExpiration = 2 * time.Minute
	defaultGCInterval = time.Minute
)

type item struct {
	object         interface{}
	expirationTime time.Time
}

type Cache struct {
	items             map[string]item
	defaultExpiration time.Duration
	gcInterval        time.Duration
	gcStarted         bool
	mu                sync.Mutex
}

func NewCacheDefault() *Cache { _ = "STUB: not implemented"; return nil }

func NewCache(expiration time.Duration, gcInterval time.Duration) *Cache {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

func (c *Cache) gcExpiredCache() { _ = "STUB: not implemented"; return }

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) SetDefault(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) set(key string, value interface{}, expiration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }
