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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
)

const (
	nominatedPodsOfTheSameJob = extension.SchedulingDomainPrefix + "/nominated-pods-to-ignore"
)

type NominatedPodsOfTheSameJob struct {
	UIDs sets.Set[string]
}

func (s *NominatedPodsOfTheSameJob) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

func MakeNominatedPodsOfTheSameJob(cycleState fwktype.CycleState, uids []string) {
	_ = "STUB: not implemented"
	return
}

func GetNominatedPodsOfTheSameJob(cycleState fwktype.CycleState) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// runFilterPluginsWithNominatedPods is the unified implementation for running filter plugins
// with nominated pods. It handles all feature gate combinations:
//   - CrossSchedulerNomination: controls whether cross-scheduler nominated pods are included.
//     When enabled, uses addMergedNominatedPods (native >= + cross-scheduler >).
//     When disabled, uses addNominatedPods (native >= only).
//   - SkipFilterWithNominatedPods: controls the number of filter passes.
//     When enabled, runs a single pass with nominated pods (skipping the second pass without
//     nominated pods that handles the pod affinity corner case).
//     When disabled, runs two passes (with and without nominated pods) for conservative scheduling.
//
// It also excludes same-job nominated pods and includes NominatedNodeName diagnostic logging,
// compatible with the previous runFilterPluginsWithNominatedPodsIgnoreSameJob behavior.
func (ext *frameworkExtenderImpl) runFilterPluginsWithNominatedPods(
	ctx context.Context,
	state fwktype.CycleState,
	pod *corev1.Pod,
	info fwktype.NodeInfo,
) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// Pass 1: add nominated pods to nodeInfo overlay.

// Pass 2: only needed if pass 1 added pods AND pass 1 succeeded.

// NominatedNodeName diagnostic logging.

// In single-pass mode, skip the second pass.

// addNominatedPods adds pods with equal or greater priority which are nominated
// to run on the node. It returns 1) whether any pod was added, 2) augmented cycleState,
// 3) augmented nodeInfo.
func addNominatedPods(ctx context.Context, fh fwktype.Handle, pod *corev1.Pod, state fwktype.CycleState, nodeInfo fwktype.NodeInfo, podsOfSameJob sets.Set[string]) (bool, fwktype.CycleState, fwktype.NodeInfo, error, []string) {
	_ = "STUB: not implemented"

	// This may happen only in tests.
	return false, *new(fwktype.CycleState), *new(fwktype.NodeInfo), nil, nil
}

// addMergedNominatedPods adds both native and cross-scheduler nominated pods to a cloned nodeInfo.
// For native nominated pods: uses priority >= (consistent with k8s native behavior).
// For cross-scheduler nominated pods: uses priority > (strictly greater than, to avoid same-priority deadlock).
func addMergedNominatedPods(
	ctx context.Context,
	ext *frameworkExtenderImpl,
	pod *corev1.Pod,
	state fwktype.CycleState,
	nodeInfo fwktype.NodeInfo,
	podsOfSameJob sets.Set[string],
) (bool, fwktype.CycleState, fwktype.NodeInfo, error, []string) {
	_ = "STUB: not implemented"
	return false, *new(fwktype.CycleState), *new(fwktype.NodeInfo), nil, nil
}

// Get native nominated pods via the embedded framework handle.

// Get cross-scheduler nominated pods.

// Add responsible nominated pods not in the same job (priority >= current pod, consistent with k8s native behavior).

// Add cross-scheduler nominated pods not in the same job (priority > current pod, strictly greater to avoid same-priority deadlock).
