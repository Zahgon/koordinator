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

package frameworkext

import (
	"context"
	"time"

	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var (
	// JobPreemptionSuccessPlugin If the waitingPod is rejected by the Job preemption plugin, the assumedNodeName is the nominatedNodeName.
	JobPreemptionSuccessPlugin = "JobPreemptionSuccessPlugin"
	// JobPreemptionFailurePlugin If Job preemption fails, the nominatedNodeName of the WaitingPod should be erased.
	JobPreemptionFailurePlugin = "JobPreemptionFailurePlugin"
	// JobRejectPlugin If the waitingPod is only recycled by the JobRejectPlugin, the NominatedNodeName should not be modified.
	JobRejectPlugin = "JobRejectPlugin"
)

func TakeoverNominatingInfo(ctx context.Context, fwk framework.Framework, podInfo *framework.QueuedPodInfo, status *fwktype.Status, nominatingInfo *fwktype.NominatingInfo, start time.Time) *fwktype.NominatingInfo {
	_ = "STUB: not implemented"
	return nil
}

func getRejecterPlugin(status *fwktype.Status) string { _ = "STUB: not implemented"; return "" }

// 1.33 fwktype.NewStatus(fwktype.Unschedulable, msg).WithPlugin(pluginName):
//		fitErr := &framework.FitError{
//			NumAllNodes: 1,
//			Pod:         assumedPodInfo.Pod,
//			Diagnosis: framework.Diagnosis{
//				NodeToStatus:         framework.NewDefaultNodeToStatus(),
//				UnschedulablePlugins: sets.New(status.Plugin()),
//			},
//		}
//		fwktype.NewStatus(status.Code()).WithError(fitErr)

// 1.28 fwktype.NewStatus(fwktype.Unschedulable, msg).WithFailedPlugin(pluginName):
