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

package reservation

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

var _ Interpreter = &interpreterImpl{}

type interpreterImpl struct {
	mgr ctrl.Manager
	client.Client
}

func newInterpreter(mgr ctrl.Manager) Interpreter {
	_ = "STUB: not implemented"
	return *new(Interpreter)
}

func (p *interpreterImpl) GetReservationType() client.Object {
	_ = "STUB: not implemented"
	return *new(client.Object)
}

func (p *interpreterImpl) Preemption() Preemption {
	_ = "STUB: not implemented"
	return *new(Preemption)
}

func (p *interpreterImpl) GetReservation(ctx context.Context, ref *corev1.ObjectReference) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}

func (p *interpreterImpl) DeleteReservation(ctx context.Context, ref *corev1.ObjectReference) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *interpreterImpl) CreateReservation(ctx context.Context, job *sev1alpha1.PodMigrationJob) (Object, error) {
	_ = "STUB: not implemented"
	return *new(Object), nil
}
