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
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/controllerfinder"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

type filter struct {
	client client.Client
	clock  clock.RealClock

	nonRetryablePodFilter framework.FilterFunc
	retryablePodFilter    framework.FilterFunc
	defaultFilterPlugin   framework.FilterPlugin

	args              *deschedulerconfig.MigrationControllerArgs
	controllerFinder  controllerfinder.Interface
	skipEvictionGates map[deschedulerconfig.EvictionGate]struct{}

	arbitratedPodMigrationJobs map[types.UID]bool
	arbitratedMapLock          sync.Mutex
}

func newEvictionGateSet(gates []deschedulerconfig.EvictionGate) map[deschedulerconfig.EvictionGate]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (f *filter) isEvictionGateSkipped(gate deschedulerconfig.EvictionGate) bool {
	_ = "STUB: not implemented"
	return false
}

func newFilter(args *deschedulerconfig.MigrationControllerArgs, handle framework.Handle) (*filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *filter) initFilters(args *deschedulerconfig.MigrationControllerArgs, handle framework.Handle) error {
	_ = "STUB: not implemented"
	// Derive effective configuration based on SkipEvictionGates (Skip has the highest priority).
	return nil
}

// NOTE: DefaultEvictorArgs (used by PreEvictionFilter) only supports EvictFailedBarePods.
// We still bypass bare-pod ownerRef constraints in our main filter via evictAllBarePods=true.

// any annotated as evictable pod pass non-retryable filter

func (f *filter) reservationFilter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (f *filter) forEachAvailableMigrationJobs(listOpts *client.ListOptions, handler func(job *sev1alpha1.PodMigrationJob) bool, expectedPhaseContexts ...phaseContext) {
	_ = "STUB: not implemented"
	return
}

func (f *filter) filterExistingPodMigrationJob(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) existingPodMigrationJob(pod *corev1.Pod, expectedPhaseContexts ...phaseContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) filterMaxMigratingGlobally(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) filterMaxMigratingPerNode(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) filterMaxMigratingPerNamespace(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) filterMaxMigratingOrUnavailablePerWorkload(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) filterExpectedReplicas(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO(joseph): There are f few special scenarios where should we allow eviction?

func (f *filter) getUnavailablePods(pods []*corev1.Pod) map[types.NamespacedName]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func mergeUnavailableAndMigratingPods(unavailablePods, migratingPods map[types.NamespacedName]struct{}) {
	_ = "STUB: not implemented"
	return
}

func (f *filter) checkJobPassedArbitration(uid types.UID) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *filter) markJobPassedArbitration(uid types.UID) { _ = "STUB: not implemented"; return }

func (f *filter) removeJobPassedArbitration(uid types.UID) { _ = "STUB: not implemented"; return }

type phaseContext struct {
	phase            sev1alpha1.PodMigrationJobPhase
	checkArbitration bool
}
