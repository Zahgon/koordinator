/*
Copyright 2017 The Kubernetes Authors.

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

package replicaset

import (
	"time"

	appsv1 "k8s.io/api/apps/v1"
	clientset "k8s.io/client-go/kubernetes"
)

// WaitForReadyReplicaSet waits until the replicaset has all of its replicas ready.
func WaitForReadyReplicaSet(c clientset.Interface, ns, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForReplicaSetTargetAvailableReplicas waits for .status.availableReplicas of a RS to equal targetReplicaNum
func WaitForReplicaSetTargetAvailableReplicas(c clientset.Interface, replicaSet *appsv1.ReplicaSet, targetReplicaNum int32) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForReplicaSetTargetAvailableReplicasWithTimeout waits for .status.availableReplicas of a RS to equal targetReplicaNum
// with given timeout.
func WaitForReplicaSetTargetAvailableReplicasWithTimeout(c clientset.Interface, replicaSet *appsv1.ReplicaSet, targetReplicaNum int32, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
