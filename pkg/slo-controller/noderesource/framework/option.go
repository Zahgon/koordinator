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

package framework

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Option struct {
	Client   client.Client
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
	Builder  *builder.Builder
}

func NewOption() *Option { _ = "STUB: not implemented"; return nil }

func (o *Option) WithManager(mgr ctrl.Manager) *Option { _ = "STUB: not implemented"; return nil }

func (o *Option) WithControllerBuilder(b *builder.Builder) *Option {
	_ = "STUB: not implemented"
	return nil
}

func (o *Option) WithClient(c client.Client) *Option { _ = "STUB: not implemented"; return nil }

func (o *Option) WithRecorder(r record.EventRecorder) *Option {
	_ = "STUB: not implemented"
	return nil
}

func (o *Option) WithScheme(s *runtime.Scheme) *Option { _ = "STUB: not implemented"; return nil }

func (o *Option) CompleteController(r reconcile.Reconciler) error {
	_ = "STUB: not implemented"
	return nil
}

type FilterFn func(string) bool

var AllPass = func(string) bool {
	return true
}
