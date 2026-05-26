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

package mutating

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// ElasticQuotaMutatingHandler handles ElasticQuota
type ElasticQuotaMutatingHandler struct {
	Client client.Client

	// Decoder decodes the objects
	Decoder admission.Decoder
}

var _ admission.Handler = &ElasticQuotaMutatingHandler{}

func shouldIgnoreIfNotElasticQuotas(req admission.Request) bool {
	_ = "STUB: not implemented"
	// Ignore all calls to sub resources or resources other than pods.
	return false
}

func (h *ElasticQuotaMutatingHandler) Handle(ctx context.Context, request admission.Request) (resp admission.Response) {
	_ = "STUB: not implemented"
	return *new(admission.Response)
}

// var _ inject.Client = &ElasticQuotaMutatingHandler{}

// InjectClient injects the client into the ElasticQuotaMutatingHandler
func (h *ElasticQuotaMutatingHandler) InjectClient(c client.Client) error {
	_ = "STUB: not implemented"
	return nil

	// var _ admission.DecoderInjector = &ElasticQuotaMutatingHandler{}
}

// InjectDecoder injects the decoder into the ElasticQuotaMutatingHandler
func (h *ElasticQuotaMutatingHandler) InjectDecoder(decoder admission.Decoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *ElasticQuotaMutatingHandler) InjectCache(cache cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}
