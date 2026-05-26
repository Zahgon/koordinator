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

package coresched

import (
	"sync"
)

// CookieCacheEntry is an entry which stores the cookie ID and its belonging PIDs.
type CookieCacheEntry struct {
	rwMutex  sync.RWMutex
	cookieID uint64
	pidCache *PIDCache
}

func newCookieCacheEntry(cookieID uint64, pids ...uint32) *CookieCacheEntry {
	_ = "STUB: not implemented"
	return nil
}

func (c *CookieCacheEntry) DeepCopy() *CookieCacheEntry { _ = "STUB: not implemented"; return nil }

func (c *CookieCacheEntry) GetCookieID() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *CookieCacheEntry) SetCookieID(cookieID uint64) { _ = "STUB: not implemented"; return }

func (c *CookieCacheEntry) IsEntryInvalid() bool { _ = "STUB: not implemented"; return false }

func (c *CookieCacheEntry) HasPID(pid uint32) bool { _ = "STUB: not implemented"; return false }

func (c *CookieCacheEntry) ContainsPIDs(pids ...uint32) []uint32 {
	_ = "STUB: not implemented"
	return nil
}

// GetAllPIDs gets all PIDs sorted in ascending order.
func (c *CookieCacheEntry) GetAllPIDs() []uint32 { _ = "STUB: not implemented"; return nil }

func (c *CookieCacheEntry) AddPIDs(pids ...uint32) { _ = "STUB: not implemented"; return }

func (c *CookieCacheEntry) DeletePIDs(pids ...uint32) { _ = "STUB: not implemented"; return }

type PIDCache map[uint32]struct{}

func NewPIDCache(pids ...uint32) *PIDCache { _ = "STUB: not implemented"; return nil }

func (p PIDCache) DeepCopy() *PIDCache { _ = "STUB: not implemented"; return nil }

func (p PIDCache) Len() int { _ = "STUB: not implemented"; return 0 }

func (p PIDCache) Has(pid uint32) bool { _ = "STUB: not implemented"; return false }

func (p PIDCache) GetAllSorted() []uint32 { _ = "STUB: not implemented"; return nil }

func (p PIDCache) AddAny(pids ...uint32) { _ = "STUB: not implemented"; return }

func (p PIDCache) DeleteAny(pids ...uint32) { _ = "STUB: not implemented"; return }
