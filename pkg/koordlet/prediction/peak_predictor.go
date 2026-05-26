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

package prediction

import (
	"time"

	v1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/extension"
)

// PredictorType defines constants for different types of predictors.
type PredictorType int

type PredictorContext struct {
	Node *v1.Node
}

const (
	// ProdReclaimablePredictor represents the type of a reclaimable production predictor.
	ProdReclaimablePredictor PredictorType = iota
)

// PredictorFactory is an interface for creating predictors of different types.
type PredictorFactory interface {
	New(predictorType PredictorType, context PredictorContext) Predictor
}

type Predictor interface {
	GetPredictorName() string
	AddPod(pod *v1.Pod) error
	GetResult() (v1.ResourceList, error)
}

type predictorFactory struct {
	predictServer       PredictServer
	coldStartDuration   time.Duration
	safetyMarginPercent int
}

// NewPredictorFactory creates a new instance of PredictorFactory.
func NewPredictorFactory(predictServer PredictServer, coldStartDuration time.Duration, safetyMarginPercent int) PredictorFactory {
	_ = "STUB: not implemented"
	return *new(PredictorFactory)
}

// New creates a new instance of a predictor based on the given type.
func (f *predictorFactory) New(t PredictorType, context PredictorContext) Predictor {
	_ = "STUB: not implemented"
	return *new(Predictor)
}

var _ Predictor = (*emptyPredictor)(nil)

type emptyPredictor struct {
}

func (p *emptyPredictor) GetPredictorName() string { _ = "STUB: not implemented"; return "" }

// AddPod adds a pod to the predictor for resource prediction.
func (p *emptyPredictor) AddPod(pod *v1.Pod) error {
	_ = "STUB: not implemented"

	// GetResult returns an error indicating that the predictor is empty.
	return nil
}

func (p *emptyPredictor) GetResult() (v1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList), nil
}

func NewEmptyPredictorFactory() PredictorFactory {
	_ = "STUB: not implemented"
	return *new(PredictorFactory)
}

type emptyPredictorFactory struct {
}

func (f *emptyPredictorFactory) New(t PredictorType, context PredictorContext) Predictor {
	_ = "STUB: not implemented"
	return *new(Predictor)
}

var _ Predictor = (*podReclaimablePredictor)(nil)

// podReclaimablePredictor predicts the peak according to historical metrics of the pods.
// e.g. A podReclaimablePredictor for Prod pods calculates the result based on the sum of the percentile of Prod pods.
type podReclaimablePredictor struct {
	predictServer       PredictServer
	node                *v1.Node
	coldStartDuration   time.Duration
	safetyMarginPercent int
	podFilterFn         func(pod *v1.Pod) bool // return true if the pod is reclaimable
	reclaimable         v1.ResourceList
	unReclaimable       v1.ResourceList
	pods                map[string]bool
}

// GetPredictorName is used to obtain the predictor name.
func (p *podReclaimablePredictor) GetPredictorName() string { _ = "STUB: not implemented"; return "" }

// AddPod adds a pod to the predictor for resource prediction.
func (p *podReclaimablePredictor) AddPod(pod *v1.Pod) error {
	_ = "STUB: not implemented"
	// podReclaimablePredictor process only specified PriorityClass pods.
	return nil
}

// Pods in cold start have 0 reclaimable resources

// Pods in terminating stage have 0 reclaimable resources
// Terminated pods are not running and do not need to predict.

// TODO: customize the percentile

// calculate the reclaimable resources: reclaimable = podRequest - peak
// calculate the unReclaimable resources: unReclaimable = peak

// update the unReclaimable resources

// update the reclaimableCPUMilli resources

// GetResult returns the predicted resource list for the added pods.
// The result is the sum of the reclaimable resources of the added pods.
func (p *podReclaimablePredictor) GetResult() (v1.ResourceList, error) {
	_ = "STUB: not implemented"
	// if failed to get node info, stop the reclaimPredictor
	return *new(v1.ResourceList), nil
}

var _ Predictor = (*priorityReclaimablePredictor)(nil)

// priorityReclaimablePredictor predicts the peak according to historical metrics of the node priority resources.
// e.g. A priorityReclaimablePredictor for Prod calculates the result based on the sum of the percentile of the
// Prod-tier and the system components parts.
type priorityReclaimablePredictor struct {
	predictServer         PredictServer
	node                  *v1.Node
	safetyMarginPercent   int
	priorityClassFilterFn func(p extension.PriorityClass) bool // return true if the priority class is reclaimable

	reclaimRequest v1.ResourceList
}

// GetPredictorName is used to obtain the predictor name.
func (p *priorityReclaimablePredictor) GetPredictorName() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *priorityReclaimablePredictor) AddPod(pod *v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// TBD: handle the cold start pods if necessary.

// Pods in terminating stage have 0 reclaimable resources
// Terminated pods are not running and do not need to predict.

func (p *priorityReclaimablePredictor) GetResult() (v1.ResourceList, error) {
	_ = "STUB: not implemented"
	// if failed to get node info, stop the reclaimPredictor
	return *new(v1.ResourceList), nil
}

// get sys prediction

// get reclaimable priority class prediction

// scale with the safety margin

// reclaimable[P] := max(request[P] - peak[P], 0)

// fixReclaimable[P] := min(nodeAllocatable[P]-unReclaimable[P],reclaimable[P])

var _ Predictor = (*minPredictor)(nil)

// minPredictor predicts the peak according to the minimal of the results of the sub-predictors.
type minPredictor struct {
	predictors []Predictor
}

// GetPredictorName is used to obtain the predictor name.
func (m *minPredictor) GetPredictorName() string { _ = "STUB: not implemented"; return "" }

func (m *minPredictor) AddPod(pod *v1.Pod) error { _ = "STUB: not implemented"; return nil }

func (m *minPredictor) GetResult() (v1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList), nil
}

func isPodReclaimableForProd(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func isPriorityClassReclaimableForProd(priorityClass extension.PriorityClass) bool {
	_ = "STUB: not implemented"
	return false
}

func getNodeAllocatable(node *v1.Node) (v1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList), nil
}
