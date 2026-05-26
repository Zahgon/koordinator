/*
Copyright 2018 The Kubernetes Authors.

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
	"sync"

	clientset "k8s.io/client-go/kubernetes"
)

// Factory is a func which operates provider specific behavior.
type Factory func() (ProviderInterface, error)

var (
	providers = make(map[string]Factory)
	mutex     sync.Mutex
)

// RegisterProvider is expected to be called during application init,
// typically by an init function in a provider package.
func RegisterProvider(name string, factory Factory) { _ = "STUB: not implemented"; return }

// GetProviders returns the names of all currently registered providers.
func GetProviders() []string { _ = "STUB: not implemented"; return nil }

func init() {
	// "local" or "skeleton" can always be used.
	RegisterProvider("local", func() (ProviderInterface, error) {
		return NullProvider{}, nil
	})
	RegisterProvider("skeleton", func() (ProviderInterface, error) {
		return NullProvider{}, nil
	})
	// The empty string used to be accepted in the past, but is not
	// a valid value anymore.
}

// SetupProviderConfig validates the chosen provider and creates
// an interface instance for it.
func SetupProviderConfig(providerName string) (ProviderInterface, error) {
	_ = "STUB: not implemented"
	return *new(ProviderInterface), nil
}

// ProviderInterface contains the implementation for certain
// provider-specific functionality.
type ProviderInterface interface {
	FrameworkBeforeEach(f *Framework)
	FrameworkAfterEach(f *Framework)

	ResizeGroup(group string, size int32) error
	GetGroupNodes(group string) ([]string, error)
	GroupSize(group string) (int, error)

	DeleteNode(node *v1.Node) error

	CreatePD(zone string) (string, error)
	DeletePD(pdName string) error
	CreatePVSource(zone, diskName string) (*v1.PersistentVolumeSource, error)
	DeletePVSource(pvSource *v1.PersistentVolumeSource) error

	CleanupServiceResources(c clientset.Interface, loadBalancerName, region, zone string)

	EnsureLoadBalancerResourcesDeleted(ip, portRange string) error
	LoadBalancerSrcRanges() []string
	EnableAndDisableInternalLB() (enable, disable func(svc *v1.Service))
}

// NullProvider is the default implementation of the ProviderInterface
// which doesn't do anything.
type NullProvider struct{}

// FrameworkBeforeEach is a base implementation which does BeforeEach.
func (n NullProvider) FrameworkBeforeEach(f *Framework) {
	_ = "STUB: not implemented"

	// FrameworkAfterEach is a base implementation which does AfterEach.
	return
}

func (n NullProvider) FrameworkAfterEach(f *Framework) {
	_ = "STUB: not implemented"

	// ResizeGroup is a base implementation which resizes group.
	return
}

func (n NullProvider) ResizeGroup(string, int32) error { _ = "STUB: not implemented"; return nil }

// GetGroupNodes is a base implementation which returns group nodes.
func (n NullProvider) GetGroupNodes(group string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GroupSize returns the size of an instance group
func (n NullProvider) GroupSize(group string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DeleteNode is a base implementation which deletes a node.
func (n NullProvider) DeleteNode(node *v1.Node) error { _ = "STUB: not implemented"; return nil }

// CreatePD is a base implementation which creates PD.
func (n NullProvider) CreatePD(zone string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DeletePD is a base implementation which deletes PD.
func (n NullProvider) DeletePD(pdName string) error { _ = "STUB: not implemented"; return nil }

// CreatePVSource is a base implementation which creates PV source.
func (n NullProvider) CreatePVSource(zone, diskName string) (*v1.PersistentVolumeSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePVSource is a base implementation which deletes PV source.
func (n NullProvider) DeletePVSource(pvSource *v1.PersistentVolumeSource) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanupServiceResources is a base implementation which cleans up service resources.
func (n NullProvider) CleanupServiceResources(c clientset.Interface, loadBalancerName, region, zone string) {
	_ = "STUB: not implemented"

	// EnsureLoadBalancerResourcesDeleted is a base implementation which ensures load balancer is deleted.
	return
}

func (n NullProvider) EnsureLoadBalancerResourcesDeleted(ip, portRange string) error {
	_ = "STUB: not implemented"

	// LoadBalancerSrcRanges is a base implementation which returns the ranges of ips used by load balancers.
	return nil
}

func (n NullProvider) LoadBalancerSrcRanges() []string {
	_ = "STUB: not implemented"

	// EnableAndDisableInternalLB is a base implementation which returns functions for enabling/disabling an internal LB.
	return nil
}

func (n NullProvider) EnableAndDisableInternalLB() (enable, disable func(svc *v1.Service)) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ ProviderInterface = NullProvider{}
