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

package migration

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

const (
	AnnotationJobCreatedBy = "koordinator.sh/job-created-by"
)

var applyJobContextFn = func(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	jobCtx := FromContext(ctx)
	return jobCtx.ApplyTo(job)
}

// Evict evicts a pod
func (r *Reconciler) Evict(ctx context.Context, pod *corev1.Pod, evictOptions framework.EvictOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func CreatePodMigrationJob(ctx context.Context, pod *corev1.Pod, evictOptions framework.EvictOptions, client client.Client, args *deschedulerconfig.MigrationControllerArgs, reconcilerUID types.UID) error {
	_ = "STUB: not implemented"
	return nil
}
