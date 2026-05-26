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

package colocationprofile

import (
	"context"
	"flag"
	"time"

	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const Name = "colocationprofile"

var (
	ReconcileByDefault     = false
	ReconcileInterval      = 30 * time.Second
	ForceUpdatePodDuration = 5 * time.Minute
	ExpirePodCacheDuration = 10 * time.Minute
	MaxUpdatePodQPS        = 10.0
	MaxUpdatePodQPSBurst   = 100
)

type Reconciler struct {
	client.Client
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme

	rateLimiter    *rate.Limiter
	podUpdateCache gocache.Cache
}

func newReconciler(mgr ctrl.Manager) *Reconciler { _ = "STUB: not implemented"; return nil }

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// not found

// skip for a terminating profile

// should not handle the profile

// NOTE: Only handle pending and unscheduled pods.

// rate limit the pod updates

// TODO: handle reservations

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func Add(mgr ctrl.Manager) error { _ = "STUB: not implemented"; return nil }
