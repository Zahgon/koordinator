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

	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type Param struct {
	IsPodEnabled bool
	IsExpeller   bool
	IsCPUIdle    bool
}

func newParam(qosCfg *slov1alpha1.CPUQOSCfg, policy slov1alpha1.CPUQOSPolicy) Param {
	_ = "STUB: not implemented"
	return *new(Param)
}

type Rule struct {
	lock             sync.RWMutex
	enable           bool // node-level switch
	podQOSParams     map[extension.QoSClass]Param
	kubeQOSPodParams map[corev1.PodQOSClass]Param
}

func newRule() *Rule { _ = "STUB: not implemented"; return nil }

func (r *Rule) IsInited() bool { _ = "STUB: not implemented"; return false }

func (r *Rule) IsEnabled() bool { _ = "STUB: not implemented"; return false }

// IsPodEnabled returns if the pod's core sched is enabled by the rule, and if the QoS-level core expeller is enabled.
func (r *Rule) IsPodEnabled(podQoSClass extension.QoSClass, podKubeQOS corev1.PodQOSClass) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// core sched is not needed for all types of pods, so it should be disabled by default

func (r *Rule) IsKubeQOSCPUIdle(KubeQOS corev1.PodQOSClass) bool {
	_ = "STUB: not implemented"
	return false
}

// cpu idle disabled by default

func (r *Rule) Update(ruleNew *Rule) bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) parseRuleForNodeSLO(mergedNodeSLOIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// default policy disables

// setting pod rule by qos config

// setting guaranteed pod enabled if LS or LSR enabled

func (p *Plugin) parseForAllPods(e interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Plugin) ruleUpdateCb(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// check the kernel feature and enable if needed

func (p *Plugin) refreshForAllPods(podMetas []*statesinformer.PodMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// sandbox-container-level

// container-level
