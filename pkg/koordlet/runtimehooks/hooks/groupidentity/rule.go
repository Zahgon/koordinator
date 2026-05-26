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

package groupidentity

import (
	corev1 "k8s.io/api/core/v1"

	ext "github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type bvtRule struct {
	enable           bool
	podQOSParams     map[ext.QoSClass]int64
	kubeQOSDirParams map[corev1.PodQOSClass]int64
	kubeQOSPodParams map[corev1.PodQOSClass]int64
}

func (r *bvtRule) getEnable() bool { _ = "STUB: not implemented"; return false }

func (r *bvtRule) getPodBvtValue(podQoSClass ext.QoSClass, podKubeQOS corev1.PodQOSClass) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (r *bvtRule) getKubeQOSDirBvtValue(kubeQOS corev1.PodQOSClass) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (r *bvtRule) getHostQOSBvtValue(qosClass ext.QoSClass) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (b *bvtPlugin) parseRule(mergedNodeSLOIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// default policy enables

// check if bvt (group identity) is enabled

// setting pod rule by qos config
// Group Identity should be reset if the CPU QOS disables (already merged in states informer) or the CPU QoS policy
// is not "groupIdentity".

// setting besteffort according to BE

// setting burstable according to LS

// NOTE: guaranteed root dir must set as 0 until kernel supported

// setting guaranteed pod enabled if LS or LSR enabled

func (b *bvtPlugin) ruleUpdateCb(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// check sysctl
// Currently, the kernel feature core scheduling is conflict to the group identity. So before we enable the
// group identity, we should check if the GroupIdentity can be enabled via sysctl and the CoreSched can be
// disabled via sysctl. And when we disable the group identity, we can check if the GroupIdentity is already
// disabled which means we do not need to update the cgroups.

// no need to update cgroups if both rule and sysctl disabled

// FIXME(saintube): Currently the kernel feature core scheduling is strictly excluded with the group identity's
//   bvt=-1. So we have to check and disable all the BE cgroups' bvt for the GroupIdentity before creating the
//   core sched cookies. To keep the consistency of the cgroup tree's configuration, we list and update the cgroups
//   by the level order when the group identity is globally disabled.
//   This check should be removed after the kernel provides a more stable interface.

// pod-level

// exclude qos cgroup

// container-level
// NOTE: Although we do not set the container's cpu.bvt_warp_ns directly, it is inheritable from the pod-level,
//       we have to handle the container-level only when we want to disable the group identity.

// handle the remaining pod cgroups, which can belong to the dangling pods

// container-level

func (b *bvtPlugin) getRule() *bvtRule { _ = "STUB: not implemented"; return nil }

func (b *bvtPlugin) updateRule(newRule *bvtRule) bool { _ = "STUB: not implemented"; return false }
