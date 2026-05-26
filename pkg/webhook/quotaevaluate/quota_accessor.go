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

package quotaevaluate

import (
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/utils/lru"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

// QuotaAccessor abstracts the get/set logic from the rest of the Evaluator.  This could be a test stub, a straight passthrough,
// or most commonly a series of deconflicting caches.
type QuotaAccessor interface {
	// UpdateQuotaStatus is called to persist final status.  This method should write to persistent storage.
	// An error indicates that write didn't complete successfully.
	UpdateQuotaStatus(newQuota *v1alpha1.ElasticQuota) error

	// GetQuota gets the specificated elastic quota.
	GetQuota(key string) (*v1alpha1.ElasticQuota, error)
}

type quotaAccessor struct {
	client client.Client
	// updatedQuotas holds a cache of quotas that we've updated.  This is used to pull the "really latest" during back to
	// back quota evaluations that touch the same quota doc.  This only works because we can compare etcd resourceVersions
	// for the same resource as integers.  Before this change: 22 updates with 12 conflicts.  after this change: 15 updates with 0 conflicts
	updatedQuotas *lru.Cache
}

// NewQuotaAccessor creates an object that conforms to the QuotaAccessor interface to be used to retrieve quota objects.
func NewQuotaAccessor(c client.Client) *quotaAccessor { _ = "STUB: not implemented"; return nil }

func (q *quotaAccessor) UpdateQuotaStatus(newQuota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

var etcdVersioner = storage.APIObjectVersioner{}

func (q *quotaAccessor) checkCache(quota *v1alpha1.ElasticQuota) *v1alpha1.ElasticQuota {
	_ = "STUB: not implemented"
	return nil
}

func (q *quotaAccessor) GetQuota(key string) (*v1alpha1.ElasticQuota, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
