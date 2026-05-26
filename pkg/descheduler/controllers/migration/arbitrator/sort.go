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

package arbitrator

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	JobKind = "Job"
)

// SortJobsByPod sort a SortFn that sorts PodMigrationJobs by their Pods, including priority, QoS.
func SortJobsByPod(sorter func(pods []*corev1.Pod)) SortFn {
	_ = "STUB: not implemented"
	return *new(SortFn)
}

// using PodSorter to sort

// SortJobsByCreationTime returns a SortFn that stably sorts PodMigrationJobs by create time.
func SortJobsByCreationTime() SortFn { _ = "STUB: not implemented"; return *new(SortFn) }

// SortJobsByMigratingNum returns a SortFn that stably sorts PodMigrationJobs the number of migrating PMJs in the same Job.
func SortJobsByMigratingNum(c client.Client) SortFn { _ = "STUB: not implemented"; return *new(SortFn) }

// get owner of jobs

// SortJobsByController returns a SortFn that places PodMigrationJobs in the same job in adjacent positions.
func SortJobsByController() SortFn { _ = "STUB: not implemented"; return *new(SortFn) }

func getMigratingJobNum(c client.Client, ownerUID types.UID) int {
	_ = "STUB: not implemented"
	return 0
}

func getJobControllerOfPod(pod *corev1.Pod) (*metav1.OwnerReference, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func isJobMigrating(job *v1alpha1.PodMigrationJob) bool { _ = "STUB: not implemented"; return false }
