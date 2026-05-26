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

package controller

import (
	"context"
	"sync"
	"time"

	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	webhookutil "github.com/koordinator-sh/koordinator/pkg/webhook/util"
)

const defaultResyncPeriod = time.Minute

var (
	mutatingWebhookConfigurationName   = webhookutil.GetMutatingWebhookName()
	validatingWebhookConfigurationName = webhookutil.GetValidatingWebhookName()

	namespace  = webhookutil.GetNamespace()
	secretName = webhookutil.GetSecretName()

	uninit   = make(chan struct{})
	onceInit = sync.Once{}
)

func Inited() chan struct{} { _ = "STUB: not implemented"; return nil }

type Controller struct {
	kubeClient clientset.Interface
	handlers   map[string]admission.Handler

	informerFactory informers.SharedInformerFactory
	synced          []cache.InformerSynced

	queue workqueue.RateLimitingInterface
}

func New(cfg *rest.Config, handlers map[string]admission.Handler) (*Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

func (c *Controller) sync() error { _ = "STUB: not implemented"; return nil }
