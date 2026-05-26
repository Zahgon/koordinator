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
	PluginName = "SLOManagerConfigConflictForNode"
)

type SLOControllerConfigConflict struct {
	client  ctrlclient.Client
	decoder ctrladmission.Decoder
}

func NewPlugin(decoder ctrladmission.Decoder, client ctrlclient.Client) *SLOControllerConfigConflict {
	_ = "STUB: not implemented"
	return nil
}

func (cl *SLOControllerConfigConflict) Name() string { _ = "STUB: not implemented"; return "" }

func (cl *SLOControllerConfigConflict) Admit(ctx context.Context, req ctrladmission.Request, node, oldNode *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *SLOControllerConfigConflict) Validate(ctx context.Context, req ctrladmission.Request, node, oldNode *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// create just check and log

func (cl *SLOControllerConfigConflict) checkConflict(node, oldNode *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func needCheck(node, oldNode *corev1.Node) bool { _ = "STUB: not implemented"; return false }
