/*
Copyright 2022 The Koordinator Authors.
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

package deployment

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	testutils "k8s.io/kubernetes/test/utils"
)

// UpdateDeploymentWithRetries updates the specified deployment with retries.
func UpdateDeploymentWithRetries(c clientset.Interface, namespace, name string, applyUpdate testutils.UpdateDeploymentFunc) (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDeployment returns a deployment spec with the specified argument.
func NewDeployment(deploymentName string, replicas int32, podLabels map[string]string, imageName, image string, strategyType appsv1.DeploymentStrategyType) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

// CreateDeployment creates a deployment.
func CreateDeployment(client clientset.Interface, replicas int32, podLabels map[string]string, nodeSelector map[string]string, namespace string, pvclaims []*v1.PersistentVolumeClaim, command string) (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodsForDeployment gets pods for the given deployment
func GetPodsForDeployment(client clientset.Interface, deployment *appsv1.Deployment) (*v1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We ignore pod-template-hash because:
// 1. The hash result would be different upon podTemplateSpec API changes
//    (e.g. the addition of a new field will cause the hash code to change)
// 2. The deployment template won't have hash labels

// Remove hash labels from template.Labels before comparing

// In rare cases, such as after cluster upgrades, Deployment may end up with
// having more than one new ReplicaSets that have the same template as its template,
// see https://github.com/kubernetes/kubernetes/issues/40415
// We deterministically choose the oldest new ReplicaSet.

// replicaSetsByCreationTimestamp sorts a list of ReplicaSet by creation timestamp, using their names as a tie breaker.
type replicaSetsByCreationTimestamp []*appsv1.ReplicaSet

func (o replicaSetsByCreationTimestamp) Len() int      { _ = "STUB: not implemented"; return 0 }
func (o replicaSetsByCreationTimestamp) Swap(i, j int) { _ = "STUB: not implemented"; return }
func (o replicaSetsByCreationTimestamp) Less(i, j int) bool {
	_ = "STUB: not implemented"
	return false
}

// testDeployment creates a deployment definition based on the namespace. The deployment references the PVC's
// name.  A slice of BASH commands can be supplied as args to be run by the pod
func testDeployment(replicas int32, podLabels map[string]string, nodeSelector map[string]string, namespace string, pvclaims []*v1.PersistentVolumeClaim, isPrivileged bool, command string) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}
