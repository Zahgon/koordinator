/*
Copyright 2022 The Koordinator Authors.
Copyright 2021 The Kruise Authors.

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

package controllerfinder

import (
	appsv1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	appsv1beta1 "github.com/openkruise/kruise-api/apps/v1beta1"
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	scaleclient "k8s.io/client-go/scale"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// ScaleAndSelector is used to return (controller, scale, selector) fields from the
// controller finder functions.
type ScaleAndSelector struct {
	ControllerReference
	// controller.spec.Replicas
	Scale int32
	// kruise statefulSet.spec.ReserveOrdinals
	ReserveOrdinals []int
	// controller.spec.Selector
	Selector *metav1.LabelSelector
	// metadata
	Metadata metav1.ObjectMeta
}

type ControllerReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion"`
	// Kind of the referent.
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	// Name of the referent.
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
	// UID of the referent.
	UID types.UID `json:"uid" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`
}

// PodControllerFinder is a function type that maps a pod to a list of
// controllers and their scale.
type PodControllerFinder func(ref ControllerReference, namespace string) (*ScaleAndSelector, error)

type Interface interface {
	GetPodsForRef(ownerReference *metav1.OwnerReference, ns string, labelSelector *metav1.LabelSelector, active bool) ([]*corev1.Pod, int32, error)
	GetExpectedScaleForPod(pods *corev1.Pod) (int32, error)
	ListPodsByWorkloads(workloadUIDs []types.UID, ns string, labelSelector *metav1.LabelSelector, active bool) ([]*corev1.Pod, error)
}

type ControllerFinder struct {
	client.Client

	mapper          meta.RESTMapper
	scaleNamespacer scaleclient.ScalesGetter
	discoveryClient discovery.DiscoveryInterface
}

var New = func(manager manager.Manager) (Interface, error) {
	finder := &ControllerFinder{
		Client: manager.GetClient(),
		mapper: manager.GetRESTMapper(),
	}
	cfg := manager.GetConfig()
	if cfg.GroupVersion == nil {
		cfg.GroupVersion = &schema.GroupVersion{}
	}
	codecs := serializer.NewCodecFactory(manager.GetScheme())
	cfg.NegotiatedSerializer = codecs.WithoutConversion()
	restClient, err := rest.RESTClientFor(cfg)
	if err != nil {
		return nil, err
	}
	k8sClient, err := clientset.NewForConfig(manager.GetConfig())
	if err != nil {
		return nil, err
	}
	finder.discoveryClient = k8sClient.Discovery()
	scaleKindResolver := scaleclient.NewDiscoveryScaleKindResolver(finder.discoveryClient)
	finder.scaleNamespacer = scaleclient.New(restClient, finder.mapper, dynamic.LegacyAPIPathResolverFunc, scaleKindResolver)
	return finder, nil
}

func (r *ControllerFinder) GetExpectedScaleForPod(pod *corev1.Pod) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *ControllerFinder) GetScaleAndSelectorForRef(apiVersion, kind, ns, name string, uid types.UID) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ControllerFinder) Finders() []PodControllerFinder { _ = "STUB: not implemented"; return nil }

var (
	ControllerKindRS       = apps.SchemeGroupVersion.WithKind("ReplicaSet")
	ControllerKindSS       = apps.SchemeGroupVersion.WithKind("StatefulSet")
	ControllerKindRC       = corev1.SchemeGroupVersion.WithKind("ReplicationController")
	ControllerKindDep      = apps.SchemeGroupVersion.WithKind("Deployment")
	ControllerKruiseKindCS = appsv1alpha1.SchemeGroupVersion.WithKind("CloneSet")
	ControllerKruiseKindSS = appsv1beta1.SchemeGroupVersion.WithKind("StatefulSet")

	validWorkloadList = []schema.GroupVersionKind{ControllerKindRS, ControllerKindSS, ControllerKindRC, ControllerKindDep, ControllerKruiseKindCS, ControllerKruiseKindSS}
)

// getPodReplicaSet finds a replicaset which has no matching deployments.
func (r *ControllerFinder) getPodReplicaSet(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// getPodReplicaSet finds a replicaset which has no matching deployments.
func (r *ControllerFinder) getReplicaSet(ref ControllerReference, namespace string) (*apps.ReplicaSet, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// when error is NotFound, it is ok here.

// getPodStatefulSet returns the statefulset referenced by the provided controllerRef.
func (r *ControllerFinder) getPodStatefulSet(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// when error is NotFound, it is ok here.

// getPodDeployments finds deployments for any replicasets which are being managed by deployments.
func (r *ControllerFinder) getPodDeployment(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// when error is NotFound, it is ok here.

func (r *ControllerFinder) getPodReplicationController(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// when error is NotFound, it is ok here.

// getPodStatefulSet returns the kruise cloneSet referenced by the provided controllerRef.
func (r *ControllerFinder) getPodKruiseCloneSet(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// when error is NotFound, it is ok here.

// getPodStatefulSet returns the kruise statefulset referenced by the provided controllerRef.
func (r *ControllerFinder) getPodKruiseStatefulSet(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	// This error is irreversible, so there is no need to return error
	return nil, nil
}

// when error is NotFound, it is ok here.

func (r *ControllerFinder) getScaleController(ref ControllerReference, namespace string) (*ScaleAndSelector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO, implementsScale

func verifyGroupKind(apiVersion, kind string, gvk schema.GroupVersionKind) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isValidGroupVersionKind(apiVersion, kind string) bool { _ = "STUB: not implemented"; return false }
