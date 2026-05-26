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

package impl

import (
	"net/http"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	kubeletconfigv1beta1 "k8s.io/kubelet/config/v1beta1"
	kubeletconfiginternal "k8s.io/kubernetes/pkg/kubelet/apis/config"
)

type KubeletStub interface {
	GetAllPods() (corev1.PodList, error)
	GetKubeletConfiguration() (*kubeletconfiginternal.KubeletConfiguration, error)
}

type kubeletStub struct {
	addr       string
	port       int
	scheme     string
	httpClient *http.Client
}

func NewKubeletStub(addr string, port int, scheme string, timeout time.Duration, cfg *rest.Config) (KubeletStub, error) {
	_ = "STUB: not implemented"
	return *new(KubeletStub), nil
}

func (k *kubeletStub) GetAllPods() (corev1.PodList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.PodList), nil
}

// parse json data

type kubeletConfigz struct {
	ComponentConfig kubeletconfigv1beta1.KubeletConfiguration `json:"kubeletconfig"`
}

// GetKubeletConfiguration removes the logging field from the configz during unmarshall to make sure the configz is compatible
func (k *kubeletStub) GetKubeletConfiguration() (*kubeletconfiginternal.KubeletConfiguration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO remove unmarshalFromNewVersionConfig after upgrade k8s dependency to 1.28

// In the Kubernetes 1.28, the Kubelet configuration introduces an API change that is not forward-compatible:
// v1.24: https://github.com/kubernetes/component-base/blob/release-1.24/config/types.go#L99
// v1.26: https://github.com/kubernetes/component-base/blob/release-1.26/logs/api/v1/types.go#L45
// v1.28: https://github.com/kubernetes/component-base/blob/release-1.28/logs/api/v1/types.go#L48
// unmarshalFromNewVersionConfig removes the logging field from the configz
func unmarshalFromNewVersionConfig(body []byte, configz *kubeletConfigz) error {
	_ = "STUB: not implemented"
	return nil
}
