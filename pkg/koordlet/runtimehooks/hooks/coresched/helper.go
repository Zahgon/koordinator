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
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type containerPID struct {
	ContainerName string
	ContainerID   string
	PID           []uint32
}

// getCookie retrieves the last core sched cookies applied to the PIDs.
// If multiple cookies are set for the PIDs, only the first non-default cookie is picked.
// It returns the last cookie ID, PIDs synced and the error.
func (p *Plugin) getCookie(pids []uint32, groupID string) (uint64, []uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// no cookie to sync, all PIDs are default or unknown

// only one cookie to sync

// else newCookieIDMap > 1

// When got more than one non-default cookie for given group, use the first synced new cookie ID.
// Let the PIDs of different cookies fixed by the next container-level reconciliation.

// addCookie creates a new cookie for the given PIDs[0], and assign the cookie to PIDs[1:].
// It returns the new cookie ID, the assigned PIDs, and the error.
// TODO: refactor to resource updater.
func (p *Plugin) addCookie(pids []uint32, groupID string) (uint64, []uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// perhaps the group is changed

// assignCookie assigns the target cookieID to the given PIDs.
// It returns the PIDs assigned, PIDs to delete, and the error (when exists, fallback adding new cookie).
// TODO: refactor to resource updater.
func (p *Plugin) assignCookie(pids, siblingPIDs []uint32, groupID string, targetCookieID uint64) ([]uint32, []uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// in case the given pids terminate, e.g. the container is restarting, aborted

// all PIDs are assigned, just refresh reference

// find one valid sibling PID to share from

// get the first valid sibling PID

// assign to valid sibling PID

// clearCookie clears the cookie for the given PIDs to the default cookie 0.
// It returns the PIDs cleared.
func (p *Plugin) clearCookie(pids []uint32, groupID string, lastCookieID uint64) []uint32 {
	_ = "STUB: not implemented"
	return nil
}

// getPodEnabledAndGroup gets whether the pod enables the core scheduling and the group ID if it does.
func (p *Plugin) getPodEnabledAndGroup(podAnnotations, podLabels map[string]string, podKubeQOS corev1.PodQOSClass, podUID string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (p *Plugin) getContainerUID(podUID string, containerID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Plugin) getContainerPIDs(containerCgroupParent string) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getSandboxContainerPIDs(podMeta *statesinformer.PodMeta) ([]uint32, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (p *Plugin) getNormalContainerPIDs(podMeta *statesinformer.PodMeta, containerStatus *corev1.ContainerStatus) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) getAllContainerPIDs(podMeta *statesinformer.PodMeta) []*containerPID {
	_ = "STUB: not implemented"
	return nil
}

// for sandbox container

// for containers

func recordContainerCookieMetrics(containerCtx *protocol.ContainerContext, groupID string, cookieID uint64) {
	_ = "STUB: not implemented"
	return
}

func resetContainerCookieMetrics(containerCtx *protocol.ContainerContext, groupID string, lastCookieID uint64) {
	_ = "STUB: not implemented"
	return
}
