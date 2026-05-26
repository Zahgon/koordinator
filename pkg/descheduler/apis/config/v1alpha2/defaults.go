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

package v1alpha2

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	migrationevictor "github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/evictor"
)

const (
	defaultMigrationControllerMaxConcurrentReconciles       = 1
	defaultNodeMetricExpirationSeconds                int64 = 180

	defaultMaxMigratingPerNode         = 2
	defaultMigrationJobMode            = sev1alpha1.PodMigrationJobModeReservationFirst
	defaultMigrationJobTTL             = 5 * time.Minute
	defaultMigrationJobEvictionPolicy  = migrationevictor.NativeEvictorName
	defaultMigrationEvictQPS           = 10
	defaultMigrationEvictBurst         = 1
	defaultSchedulerSupportReservation = "koord-scheduler"
	defaultArbitrationInterval         = 500 * time.Millisecond
	defaultDetectorCacheTimeout        = 5 * time.Minute
)

var (
	defaultObjectLimiters = map[MigrationLimitObjectType]MigrationObjectLimiter{
		MigrationLimitObjectWorkload: {
			Duration: metav1.Duration{Duration: 5 * time.Minute},
		},
		// namespace object limiter is disabled as default
	}

	defaultLoadAnomalyCondition = &LoadAnomalyCondition{
		Timeout:                  &metav1.Duration{Duration: 1 * time.Minute},
		ConsecutiveAbnormalities: 5,
		ConsecutiveNormalities:   3,
	}
)

func addDefaultingFuncs(scheme *runtime.Scheme) error { _ = "STUB: not implemented"; return nil }

func pluginsNames(p *Plugins) []string { _ = "STUB: not implemented"; return nil }

func setDefaults_Profile(prof *DeschedulerProfile) {
	_ = "STUB: not implemented"
	// Set default plugins.
	return
}

// Set default plugin configs.

// Append default configs for plugins that didn't have one explicitly set.

// This plugin is out-of-tree or doesn't require configuration.

// SetDefaults_DeschedulerConfiguration sets additional defaults
func SetDefaults_DeschedulerConfiguration(obj *DeschedulerConfiguration) {
	_ = "STUB: not implemented"
	return
}

// Only apply a default scheduler name when there is a single profile.
// Validation will ensure that every profile has a non-empty unique name.

// Add the default set of plugins and apply the configuration.

// Use lease-based leader election to reduce cost.
// We migrated for EndpointsLease lock in 1.17 and starting in 1.20 we
// migrated to Lease lock.

// Use the default LeaderElectionConfiguration options

// Scheduler has an opinion about QPS/Burst, setting specific defaults for itself, instead of generic settings.

// Enable profiling by default in the scheduler

// Enable contention profiling by default if profiling is enabled

func SetDefaults_MigrationControllerArgs(obj *MigrationControllerArgs) {
	_ = "STUB: not implemented"
	return
}

func SetDefaults_LowNodeLoadArgs(obj *LowNodeLoadArgs) { _ = "STUB: not implemented"; return }

func SetDefaults_LowNodeLoadNodePools(args *LowNodeLoadArgs) { _ = "STUB: not implemented"; return }

func SetDefaults_CustomPriorityArgs(obj *CustomPriorityArgs) { _ = "STUB: not implemented"; return }
