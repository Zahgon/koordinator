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

package sloconfig

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	PluginName = "SLOConfig"
)

// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch

type SLOControllerPlugin struct {
	client  ctrlclient.Client
	decoder ctrladmission.Decoder
}

func NewPlugin(decoder ctrladmission.Decoder, client ctrlclient.Client) *SLOControllerPlugin {
	_ = "STUB: not implemented"
	return nil
}

func (cl *SLOControllerPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (cl *SLOControllerPlugin) Admit(ctx context.Context, req ctrladmission.Request, config, oldConfig *corev1.ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *SLOControllerPlugin) Validate(ctx context.Context, req ctrladmission.Request, config, oldConfig *corev1.ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *SLOControllerPlugin) checkConfig(ctx context.Context, oldConfig *corev1.ConfigMap, config *corev1.ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

func parallelizeCheckNode(ctx context.Context, nodeList []corev1.Node, checkers checkers) error {
	_ = "STUB: not implemented"
	return nil
}
