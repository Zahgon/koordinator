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

package terwayqos

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	rootPath   = "/host-var-lib/terway/qos"
	podConfig  = "pod.json"
	nodeConfig = "global_bps_config"
)

const (
	name        = "TerwayQoS"
	description = "network qos management"

	ruleNameForNodeQoS = name + " (nodeQoS)"
	ruleNameForAllPods = name + " (allPods)"
)

type Plugin struct {
	executor resourceexecutor.ResourceUpdateExecutor

	lock    sync.RWMutex
	enabled *bool
	node    *Node
	pods    map[string]*Pod

	podFilePath, nodeFilePath string

	syncChan chan struct{}
}

func (p *Plugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

func (p *Plugin) parseRuleForNodeSLO(mergedNodeSLOIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Plugin) parseForAllPods(e interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Plugin) run() { _ = "STUB: not implemented"; return }

func (p *Plugin) syncAll() { _ = "STUB: not implemented"; return }

func (p *Plugin) update(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) syncPodConfig() error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) syncNodeConfig() error { _ = "STUB: not implemented"; return nil }

func ensureConfig(file string) error { _ = "STUB: not implemented"; return nil }

// parseNetQoS only LS and BE is taken into account
func parseNetQoS(slo *slov1alpha1.NodeSLOSpec, node *Node) error {
	_ = "STUB: not implemented"
	return nil
}

// to Byte/s

func parseQoS(qos *slov1alpha1.NetworkQOSCfg, total uint64) (QoS, error) {
	_ = "STUB: not implemented"
	return *new(QoS), nil
}

func parseQuantity(v *intstr.IntOrString, total uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getPodQoS(anno map[string]string) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getPodPrio(pod *corev1.Pod) int { _ = "STUB: not implemented"; return 0 }

func newPlugin() *Plugin { _ = "STUB: not implemented"; return nil }

var singleton *Plugin
var once sync.Once

func Object() *Plugin { _ = "STUB: not implemented"; return nil }

func BitsToBytes[T uint64 | float64 | int](bits T) T { _ = "STUB: not implemented"; return *new(T) }
