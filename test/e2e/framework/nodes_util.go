/*
Copyright 2022 The Koordinator Authors.
Copyright 2014 The Kubernetes Authors.

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

package framework

import (
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	// TODO: Remove the following imports (ref: https://github.com/kubernetes/kubernetes/issues/81245)
)

const etcdImage = "3.5.0-0"

// EtcdUpgrade upgrades etcd on GCE.
func EtcdUpgrade(targetStorage, targetVersion string) error { _ = "STUB: not implemented"; return nil }

func etcdUpgradeGCE(targetStorage, targetVersion string) error {
	_ = "STUB: not implemented"
	return nil
}

// LocationParamGKE returns parameter related to location for gcloud command.
func LocationParamGKE() string { _ = "STUB: not implemented"; return "" }

// GKE Regional Clusters are being tested.

// AppendContainerCommandGroupIfNeeded returns container command group parameter if necessary.
func AppendContainerCommandGroupIfNeeded(args []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// TODO(wojtek-t): Get rid of it once Regional Clusters go to GA.

// MasterUpgradeGKE upgrades master node to the specified version on GKE.
func MasterUpgradeGKE(namespace string, v string) error { _ = "STUB: not implemented"; return nil }

// GCEUpgradeScript returns path of script for upgrading on GCE.
func GCEUpgradeScript() string { _ = "STUB: not implemented"; return "" }

// WaitForSSHTunnels waits for establishing SSH tunnel to busybox pod.
func WaitForSSHTunnels(namespace string) { _ = "STUB: not implemented"; return }

// allow up to a minute for new ssh tunnels to establish

// NodeKiller is a utility to simulate node failures.
type NodeKiller struct {
	config   NodeKillerConfig
	client   clientset.Interface
	provider string
}

// NewNodeKiller creates new NodeKiller.
func NewNodeKiller(config NodeKillerConfig, client clientset.Interface, provider string) *NodeKiller {
	_ = "STUB: not implemented"
	return nil
}

// Run starts NodeKiller until stopCh is closed.
func (k *NodeKiller) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// wait.JitterUntil starts work immediately, so wait first.
	return
}

func (k *NodeKiller) pickNodes() []v1.Node { _ = "STUB: not implemented"; return nil }

func (k *NodeKiller) kill(nodes []v1.Node) { _ = "STUB: not implemented"; return }
