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

package webhook

import (
	"context"
	"net/http"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/koordinator-sh/koordinator/pkg/webhook/util/framework"
)

type GateFunc func() (enabled bool)

var (
	// handlerMap contains all admission webhook handlers.
	handlerMap        = map[string]admission.Handler{}
	handlerGates      = map[string]GateFunc{}
	HandlerBuilderMap = map[string]framework.HandlerBuilder{}
)

func addHandlersWithGate(m map[string]framework.HandlerBuilder, fn GateFunc) {
	_ = "STUB: not implemented"
	return
}

func filterActiveHandlers() { _ = "STUB: not implemented"; return }

func SetupWithWebhookOpt(opt *manager.Options) { _ = "STUB: not implemented"; return }

func SetupWithManager(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

// register admission handlers

// register conversion webhook

// register health handler

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingwebhookconfigurations,verbs=get;list;watch;update;patch

func Initialize(ctx context.Context, cfg *rest.Config) error { _ = "STUB: not implemented"; return nil }

func Checker(req *http.Request) error {
	_ = "STUB: not implemented"
	// Firstly wait webhook controller initialized
	return nil
}

func WaitReady() error { _ = "STUB: not implemented"; return nil }
