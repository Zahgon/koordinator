/*
Copyright 2022 The Koordinator Authors.
Copyright 2016 The Kubernetes Authors.

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

package auth

import (
	"sync"
	"time"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1authorization "k8s.io/client-go/kubernetes/typed/authorization/v1"
	v1rbac "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

const (
	policyCachePollInterval = 100 * time.Millisecond
	policyCachePollTimeout  = 5 * time.Second
)

type bindingsGetter interface {
	v1rbac.RoleBindingsGetter
	v1rbac.ClusterRoleBindingsGetter
	v1rbac.ClusterRolesGetter
}

// WaitForAuthorizationUpdate checks if the given user can perform the named verb and action.
// If policyCachePollTimeout is reached without the expected condition matching, an error is returned
func WaitForAuthorizationUpdate(c v1authorization.SubjectAccessReviewsGetter, user, namespace, verb string, resource schema.GroupResource, allowed bool) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForNamedAuthorizationUpdate checks if the given user can perform the named verb and action on the named resource.
// If policyCachePollTimeout is reached without the expected condition matching, an error is returned
func WaitForNamedAuthorizationUpdate(c v1authorization.SubjectAccessReviewsGetter, user, namespace, verb, resourceName string, resource schema.GroupResource, allowed bool) error {
	_ = "STUB: not implemented"
	return nil
}

// BindClusterRole binds the cluster role at the cluster scope. If RBAC is not enabled, nil
// is returned with no action.
func BindClusterRole(c bindingsGetter, clusterRole, ns string, subjects ...rbacv1.Subject) error {
	_ = "STUB: not implemented"
	return nil
}

// Since the namespace names are unique, we can leave this lying around so we don't have to race any caches

// BindClusterRoleInNamespace binds the cluster role at the namespace scope. If RBAC is not enabled, nil
// is returned with no action.
func BindClusterRoleInNamespace(c bindingsGetter, clusterRole, ns string, subjects ...rbacv1.Subject) error {
	_ = "STUB: not implemented"
	return nil
}

// BindRoleInNamespace binds the role at the namespace scope. If RBAC is not enabled, nil
// is returned with no action.
func BindRoleInNamespace(c bindingsGetter, role, ns string, subjects ...rbacv1.Subject) error {
	_ = "STUB: not implemented"
	return nil
}

func bindInNamespace(c bindingsGetter, roleType, role, ns string, subjects ...rbacv1.Subject) error {
	_ = "STUB: not implemented"
	return nil
}

// Since the namespace names are unique, we can leave this lying around so we don't have to race any caches

var (
	isRBACEnabledOnce sync.Once
	isRBACEnabled     bool
)

// IsRBACEnabled returns true if RBAC is enabled. Otherwise false.
func IsRBACEnabled(crGetter v1rbac.ClusterRolesGetter) bool {
	_ = "STUB: not implemented"
	return false
}
