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

package elasticquota

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type QuotaMetaChecker struct {
	client.Client
	admission.Decoder
	QuotaTopo     *quotaTopology
	QuotaInformer cache.Informer
}

var (
	quotaMetaCheck = &QuotaMetaChecker{
		QuotaTopo: nil,
	}
)

func (c *QuotaMetaChecker) Name() string { _ = "STUB: not implemented"; return "" }

func NewPlugin(decoder admission.Decoder, client client.Client) *QuotaMetaChecker {
	_ = "STUB: not implemented"
	return nil
}

func (c *QuotaMetaChecker) AdmitQuota(ctx context.Context, req admission.Request, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *QuotaMetaChecker) ValidateQuota(ctx context.Context, req admission.Request, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *QuotaMetaChecker) ValidatePod(ctx context.Context, req admission.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *QuotaMetaChecker) GetQuotaTopologyInfo() *QuotaTopologySummary {
	_ = "STUB: not implemented"
	return nil
}

func (c *QuotaMetaChecker) GetQuotaInfo(name, namespace string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (c *QuotaMetaChecker) InjectInformer(elasticQuotaInformer cache.Informer) {
	_ = "STUB: not implemented"
	return
}

func NewQuotaInformer(cache cache.Cache, qt *quotaTopology) (cache.Informer, error) {
	_ = "STUB: not implemented"
	return *new(cache.Informer), nil
}
