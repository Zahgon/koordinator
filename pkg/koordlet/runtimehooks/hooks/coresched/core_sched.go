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
	"time"

	gocache "github.com/patrickmn/go-cache"
	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	sysutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

const (
	name        = "CoreSched"
	description = "manage core sched cookies for pod and containers"

	ruleNameForNodeSLO = name + " (nodeSLO)"
	ruleNameForAllPods = name + " (allPods)"

	defaultCacheExpiration     = 300 * time.Second
	defaultCacheDeleteInterval = 600 * time.Second

	// ExpellerGroupSuffix is the default suffix of the expeller core sched group.
	ExpellerGroupSuffix = "-expeller"
	// NoneGroupID is the special ID denoting none core sched group.
	NoneGroupID = "__0__"
)

// SYSTEM QoS is excluded from the cookie mutating.
// All SYSTEM pods use the default cookie so the agent can reset the cookie of a container by ShareTo its cookie to
// the target.
var podQOSConditions = []string{string(extension.QoSBE), string(extension.QoSLS), string(extension.QoSLSR),
	string(extension.QoSLSE), string(extension.QoSNone)}

// Plugin is responsible for managing core sched cookies and cpu.idle for containers.
type Plugin struct {
	rule *Rule

	initialized     *atomic.Bool // whether the cache has been initialized
	allPodsSyncOnce sync.Once    // sync once for AllPods

	sysSupported      *bool
	supportedMsg      string
	giSysctlSupported *bool

	cookieCache        *gocache.Cache // core-sched-group-id -> cookie id, set<pid>; if the group has had cookie
	cookieCacheRWMutex sync.RWMutex
	groupCache         *gocache.Cache // pod-uid+container-id -> core-sched-group-id (note that it caches the last state); if the container has had cookie of the group

	reader   resourceexecutor.CgroupReader
	executor resourceexecutor.ResourceUpdateExecutor
	cse      sysutil.CoreSchedExtendedInterface
}

var singleton *Plugin

func Object() *Plugin { _ = "STUB: not implemented"; return nil }

func newPlugin() *Plugin { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

// TODO: hook NRI events RunPodSandbox, PostStartContainer

// TODO: support host application

func (p *Plugin) Setup(op hooks.Options) { _ = "STUB: not implemented"; return }

func (p *Plugin) SystemSupported() bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) IsCacheInited() bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) SetKubeQOSCPUIdle(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// SetContainerCookie reconciles the core sched cookie for the container.
// There are the following operations about the cookies:
//  1. Get: Get the cookie for a core sched group, firstly try finding in cache and then get from the existing PIDs.
//  2. Add: Add a new cookie for a core sched group for a container, and add a new entry into the cache.
//  3. Assign: Assign a cookie of an existing core sched group for a container and update the cache entry.
//     The cached sibling PIDs (i.e. the PIDs of the same core sched group) will be fetched in the Assign. If all
//     cookies of the sibling PIDs are default or invalid, the Assign should fall back to Add.
//  4. Clear: Clear a cookie of an existing core sched group for a container (reset to default cookie 0), and the
//     containers' PIDs are removed from the cache. The cache entry of the group is removed when the number of the
//     cached PIDs decreases to zero.
//
// If multiple non-default cookies are assigned to existing containers of the same group, the firstly-created and
// available cookie will be retained and the PIDs of others will be moved to the former.
// NOTE: The agent itself should be set the default cookie. It can be excluded by setting QoS to SYSTEM.
func (p *Plugin) SetContainerCookie(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// only process sandbox container or container has valid ID

// expect enabled
// 1. disabled -> enabled: Add or Assign.
// 2. keep enabled: Check the differences of cookie, group ID and the PIDs, and do Assign.

// FIXME(saintube): Currently we need to ensure the group identity is disabled via sysctl before enabling
//   the core sched cookie in the container reconciler, because the disabling during the rule update might
//   fail. This check can be removed once the kernel feature provides a way to disable the group identity.

// else pod disables

func (p *Plugin) initCache(podMetas []*statesinformer.PodMeta) bool {
	_ = "STUB: not implemented"
	return false
}

// loadAllCookies syncs the current core sched cookies of all pods into the cookie cache.
func (p *Plugin) loadAllCookies(podMetas []*statesinformer.PodMeta) bool {
	_ = "STUB: not implemented"
	return false
}

// container synced including using the default cookie

// If multiple cookie exists for a group, aborted to sync cache. Let the reconciliation fix these.

func (p *Plugin) initSystem(isEnabled bool) error {
	_ = "STUB: not implemented"
	// only switch sysctl if enabled
	return nil
}

// NOTE: Currently the kernel feature core scheduling is strictly excluded with the group identity's
//       bvt=-1. So we have to check if the GroupIdentity can be disabled before creating core sched cookies.

// tryDisableGroupIdentity tries disabling the group identity via sysctl to safely enable the core sched.
func (p *Plugin) tryDisableGroupIdentity() error { _ = "STUB: not implemented"; return nil }

// not support either the group identity or the core sched

// enableContainerCookie adds or assigns a core sched cookie for the container.
func (p *Plugin) enableContainerCookie(containerCtx *protocol.ContainerContext, groupID string) error {
	_ = "STUB: not implemented"
	return nil
}

// assert groupID != "0"
// NOTE: if the group ID changed for a enabled pod, the cookie will be updated while the old PIDs should expire
// in the old cookie's cache.

// firstly try Assign, if all cached sibling pids invalid, then try Add
// else cookie exists for group:
// 1. assign cookie if the container has not set cookie
// 2. assign cookie if some process of the container has missing cookie or set incoherent cookie

// do Assign

// no pid is successfully assigned

// no valid sibling PID, fallback to Add

// group has no cookie, do Add

// disableContainerCookie clears a core sched cookie for the container.
func (p *Plugin) disableContainerCookie(containerCtx *protocol.ContainerContext, groupID string) error {
	_ = "STUB: not implemented"
	return nil
}

// invalid lastGroupID means container not in group cache (container should be cleared or not ever added)
// invalid lastCookieEntry means group not in cookie cache (group should be cleared)
// let its cached PIDs expire or removed by siblings' Assign

// In case the pod has group set before while no cookie entry, do Clear to fix it

// do Clear:
// - clear cookie if any process of the container has set cookie

// getCookieCacheForPod gets the last group ID, the last cookie entry and the cookie entry for the current group.
// If a pod has not set cookie before, return lastGroupID=0 and lastCookieEntry=nil.
func (p *Plugin) getCookieCacheForContainer(groupID, containerUID string) (string, *CookieCacheEntry, *CookieCacheEntry) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// no valid cookie ref

// no valid cookie ref

func (p *Plugin) updateCookieCacheForContainer(groupID, containerUID string, cookieEntry *CookieCacheEntry) {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) cleanCookieCacheForContainer(groupID, containerUID string, cookieEntry *CookieCacheEntry) {
	_ = "STUB: not implemented"
	return
}
