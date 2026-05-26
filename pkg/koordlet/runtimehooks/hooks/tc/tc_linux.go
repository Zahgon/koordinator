//go:build linux
// +build linux

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

package tc

import (
	"sync"

	"github.com/coreos/go-iptables/iptables"
	"github.com/vishvananda/netlink"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/exec"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	name        = "tcPlugin"
	description = "setup tc rules for node"

	ruleNameForNodeSLO = name + " (nodeSLO)"
	ruleNameForAllPods = name + " (allPods)"
)

const (
	MAJOR_ID              = 1
	QDISC_MINOR_ID        = 0
	ROOT_CLASS_MINOR_ID   = 1
	SYSTEM_CLASS_MINOR_ID = 2
	LS_CLASS_MINOR_ID     = 3
	BE_CLASS_MINOR_ID     = 4

	// 0-7, In the round-robin process, classes with the lowest priority field are tried for packets first.
	SYSTEM_CLASS_PRIO = 1
	LS_CLASS_PRIO     = 2
	BE_CLASS_PRIO     = 3

	POD_FILTER_PRIO = 4

	// Maximum rate this class and all its children are guaranteed. Mandatory.
	// attention: the values below only represent the percentage of bandwidth can be used by different tc classes on the host,
	// the real values need to be calculated based on the physical network bandwidth.
	// eg: eth0: speed:200Mbit => high_clss.rate = 200Mbit * 40 / 100 = 80Mbit
	BE_CLASS_RATE_PERCENTAGE     = 30
	LS_CLASS_RATE_PERCENTAGE     = 30
	SYSTEM_CLASS_RATE_PERCENTAGE = 40

	// Maximum rate at which a class can send, if its parent has bandwidth to spare.  Defaults to the configured rate,
	// which implies no borrowing
	CEIL_PERCENTAGE = 100

	DEFAULT_INTERFACE_NAME = "eth0"
)

var (
	rootClass   = netlink.MakeHandle(MAJOR_ID, ROOT_CLASS_MINOR_ID)
	systemClass = netlink.MakeHandle(MAJOR_ID, SYSTEM_CLASS_MINOR_ID)
	lsClass     = netlink.MakeHandle(MAJOR_ID, LS_CLASS_MINOR_ID)
	beClass     = netlink.MakeHandle(MAJOR_ID, BE_CLASS_MINOR_ID)

	ipsets = []string{string(NETQoSSystem), string(NETQoSLS), string(NETQoSBE)}
)

type tcPlugin struct {
	ruleRWMutex sync.RWMutex
	rule        *tcRule

	// this is the physical NIC on host, default eth0
	interfLink netlink.Link

	// for executing the iptables command.
	iptablesHandler *iptables.IPTables
	// for executing the tc and ipset command.
	netLinkHandler netlink.Handle

	allPodsSyncOnce sync.Once

	executor resourceexecutor.ResourceUpdateExecutor
}

var singleton *tcPlugin

func Object() *tcPlugin { _ = "STUB: not implemented"; return nil }

func newPlugin() *tcPlugin { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

// TODO register NRI after there is pod ip in NRI request

func (p *tcPlugin) SetPodNetCls(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) createTcRulesForHostPod(rule *tcRule, pod *v1.Pod, egress uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// tc filter add dev eth0 parent 1: protocol ip prio 2 handle 5: cgroup
// means to 1:5 class

func (p *tcPlugin) delTcRules(handle uint32) error { _ = "STUB: not implemented"; return nil }

func getFilterPrio(nic string, handle uint32) string { _ = "STUB: not implemented"; return "" }

// find filter by a key value, just as "handle 0x**"
func getPrioForFilter(nic string, key string) string {
	_ = "STUB: not implemented"
	// output just like this:
	// filter parent 1: protocol ip pref 1 cgroup chain 0
	// filter parent 1: protocol ip pref 1 cgroup chain 0 handle 0x3
	return ""
}

func (p *tcPlugin) prepare() error { _ = "STUB: not implemented"; return nil }

func getMinorId(num uint32) string { _ = "STUB: not implemented"; return "" }

func (p *tcPlugin) refreshForAllPods(pods []*statesinformer.PodMeta, rule *tcRule) error {
	_ = "STUB: not implemented"
	return nil
}

// handle active pod

// create tc rules

// pod in host network namespace, network bandwidth can be limited by net_cls cgroup.

// finally, handled by network rules at the node level

// delete netqos rules for the deleted pods.

// handle netqos rules at the node level.

func (p *tcPlugin) createRulesForPod(rule *tcRule, pod *v1.Pod, netqos NetQoSClass, egress uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// tc filter add dev eth0 parent 1:0 protocol ip prio 2 u32 match ip dst 0.0.0.0/0 flowid 1:5

// tc filter add dev br0 parent 1:0 protocol ip prio 2 u32 match ip src 1.2.0.0 classid 1:5

// getPrio get next available priority for tc filter.
func getPrio(name string) int {
	_ = "STUB: not implemented"
	// output just like this:
	// filter parent 1: protocol ip pref 1 cgroup chain 0
	// filter parent 1: protocol ip pref 1 cgroup chain 0 handle 0x3
	return 0
}

// initHandleId get class minor id from pod.uid(last 4 digits).
func initHandleId(uid types.UID, handleToUid map[uint32]types.UID, uidToHandle map[types.UID]uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) InitRelatedRules() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) CleanUp() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) EnsureQdisc() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) DelQdisc() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) EnsureClasses() error { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) EnsureCgroupFilters() error { _ = "STUB: not implemented"; return nil }

func KeyByHandle(clsMinorId string) string { _ = "STUB: not implemented"; return "" }

func KeyByFlowId(clsMinorId string) string { _ = "STUB: not implemented"; return "" }

func GetPrio(qos NetQoSClass) uint32 { _ = "STUB: not implemented"; return 0 }

func newClass(index int, parent, handle uint32, rate, ceil uint64, prio uint32) *netlink.HtbClass {
	_ = "STUB: not implemented"
	return nil
}

// NewHtbClass NOTE: function is in here because it uses other linux functions
func NewHtbClass(attrs netlink.ClassAttrs, cattrs netlink.HtbClassAttrs) *netlink.HtbClass {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) ensureClass(nic netlink.Link, expect *netlink.HtbClass) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) deleteClass(nic netlink.Link, expect *netlink.HtbClass) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) deleteFilter(key string, delFunc exec.Cmd) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) ensureFilter(find string, createCmd, updateCmd exec.Cmd) error {
	_ = "STUB: not implemented"
	return nil
}

// handled by command because netlink does not support creating filters for cgroup types.
// creating a tc filter for be netqos pod, just as follows:
// tc filter add dev eth0 parent 1:0 protocol ip prio 2 match ip src 1.2.0.0 classid 1:5

func (p *tcPlugin) checkAllRulesExisted() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *tcPlugin) QdiscExisted() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (p *tcPlugin) classesExisted() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// other leaf class

func (p *tcPlugin) classExisted(nic netlink.Link, expect *netlink.HtbClass) error {
	_ = "STUB: not implemented"
	return nil
}
