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

package evictor

import (
	"context"
	"errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/util/flowcontrol"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

const (
	LabelEvictPolicy = "koordinator.sh/evict-policy"

	AnnotationEvictReason  = "koordinator.sh/evict-reason"
	AnnotationEvictTrigger = "koordinator.sh/evict-trigger"
)

var (
	ErrTooManyEvictions = errors.New("TooManyEvictions")
)

type FactoryFn func(client kubernetes.Interface) (Interface, error)

type Interface interface {
	Evict(ctx context.Context, job *sev1alpha1.PodMigrationJob, pod *corev1.Pod) error
}

var registry = map[string]FactoryFn{}

func RegisterEvictor(name string, factoryFn FactoryFn) { _ = "STUB: not implemented"; return }

type Interpreter interface {
	Interface
}

type interpreterImpl struct {
	evictors       map[string]Interface
	defaultEvictor Interface
	rateLimiter    flowcontrol.RateLimiter
	eventRecorder  events.EventRecorder
}

func NewInterpreter(handle framework.Handle, defaultEvictionPolicy string, evictQPS float32, evictBurst int) (Interpreter, error) {
	_ = "STUB: not implemented"
	return *new(Interpreter), nil
}

func (p *interpreterImpl) Evict(ctx context.Context, job *sev1alpha1.PodMigrationJob, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func getCustomEvictionPolicy(labels map[string]string) string { _ = "STUB: not implemented"; return "" }

func GetEvictionTriggerAndReason(annotations map[string]string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}
