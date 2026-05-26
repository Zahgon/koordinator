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
	corev1 "k8s.io/api/core/v1"
)

type NetQosGlobalConfig struct {
	HwTxBpsMax uint64 `json:"hw_tx_bps_max"`
	HwRxBpsMax uint64 `json:"hw_rx_bps_max"`
	L1TxBpsMin uint64 `json:"l1_tx_bps_min"`
	L1TxBpsMax uint64 `json:"l1_tx_bps_max"`
	L2TxBpsMin uint64 `json:"l2_tx_bps_min"`
	L2TxBpsMax uint64 `json:"l2_tx_bps_max"`
	L1RxBpsMin uint64 `json:"l1_rx_bps_min"`
	L1RxBpsMax uint64 `json:"l1_rx_bps_max"`
	L2RxBpsMin uint64 `json:"l2_rx_bps_min"`
	L2RxBpsMax uint64 `json:"l2_rx_bps_max"`
}

type NetQoSClass string

const (
	NETQoSSystem NetQoSClass = "system_class"
	NETQoSLS     NetQoSClass = "ls_class"
	NETQoSBE     NetQoSClass = "be_class"
	NETQoSNone   NetQoSClass = ""
)

func GetPodNetQoSClassByName(qos string) NetQoSClass {
	_ = "STUB: not implemented"
	return *new(NetQoSClass)
}

func GetPodNetQoSClass(pod *corev1.Pod) NetQoSClass {
	_ = "STUB: not implemented"
	return *new(NetQoSClass)
}

func GetNetQoSClassByAttrs(labels, annotations map[string]string) NetQoSClass {
	_ = "STUB: not implemented"
	return *new(NetQoSClass)
}

// annotations are for old format adaption reason

type NetworkQoS struct {
	// IngressLimit and EgressLimit is the bandwidth in bps
	// are used to set bandwidth for Pod. The unit is bps.
	// For example, 10M means 10 megabits per second.
	IngressLimit string `json:"ingressLimit"`
	EgressLimit  string `json:"egressLimit"`
}

func getIngressAndEgress(anno map[string]string) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func BitsToBytes[T uint64 | float64 | int](bits T) T { _ = "STUB: not implemented"; return *new(T) }
