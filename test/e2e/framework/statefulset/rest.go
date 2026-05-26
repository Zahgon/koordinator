/*
Copyright 2022 The Koordinator Authors.
Copyright 2019 The Kubernetes Authors.

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

package statefulset

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// CreateStatefulSet creates a StatefulSet from the manifest at manifestPath in the Namespace ns using kubectl create.
func CreateStatefulSet(c clientset.Interface, manifestPath, ns string) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

// GetPodList gets the current Pods in ss.
func GetPodList(c clientset.Interface, ss *appsv1.StatefulSet) *v1.PodList {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAllStatefulSets deletes all StatefulSet API Objects in Namespace ns.
func DeleteAllStatefulSets(c clientset.Interface, ns string) { _ = "STUB: not implemented"; return }

// Scale down each statefulset, then delete it completely.
// Deleting a pvc without doing this will leak volumes, #25101.

// Use OrphanDependents=false so it's deleted synchronously.
// We already made sure the Pods are gone inside Scale().

// pvs are global, so we need to wait for the exact ones bound to the statefulset pvcs.

// TODO: Don't assume all pvcs in the ns belong to a statefulset

// TODO: Double check that there are no pods referencing the pvc

// Scale scales ss to count replicas.
func Scale(c clientset.Interface, ss *appsv1.StatefulSet, count int32) (*appsv1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateReplicas updates the replicas of ss to count.
func UpdateReplicas(c clientset.Interface, ss *appsv1.StatefulSet, count int32) {
	_ = "STUB: not implemented"
	return
}

// Restart scales ss to 0 and then back to its previous number of replicas.
func Restart(c clientset.Interface, ss *appsv1.StatefulSet) { _ = "STUB: not implemented"; return }

// Wait for controller to report the desired number of Pods.
// This way we know the controller has observed all Pod deletions
// before we scale it back up.

// CheckHostname verifies that all Pods in ss have the correct Hostname. If the returned error is not nil than verification failed.
func CheckHostname(c clientset.Interface, ss *appsv1.StatefulSet) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckMount checks that the mount at mountPath is valid for all Pods in ss.
func CheckMount(c clientset.Interface, ss *appsv1.StatefulSet, mountPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Print inode, size etc

// Print subdirs

// Try writing

// CheckServiceName asserts that the ServiceName for ss is equivalent to expectedServiceName.
func CheckServiceName(ss *appsv1.StatefulSet, expectedServiceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// ExecInStatefulPods executes cmd in all Pods in ss. If a error occurs it is returned and cmd is not execute in any subsequent Pods.
func ExecInStatefulPods(c clientset.Interface, ss *appsv1.StatefulSet, cmd string) error {
	_ = "STUB: not implemented"
	return nil
}

// udpate updates a statefulset, and it is only used within rest.go
func update(c clientset.Interface, ns, name string, update func(ss *appsv1.StatefulSet)) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}
