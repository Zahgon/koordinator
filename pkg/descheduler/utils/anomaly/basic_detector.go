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

package anomaly

import (
	"sync"
	"time"
)

const (
	defaultTimeout = 60 * time.Second
)

func defaultAnomalyCondition(counter Counter) bool { _ = "STUB: not implemented"; return false }

func defaultNormalCondition(counter Counter) bool { _ = "STUB: not implemented"; return false }

// Options configures BasicDetector
type Options struct {
	// Timeout is the period of the open state,
	// after which the state of the BasicDetector becomes half-open.
	// If Timeout is less than or equal to 0, the timeout value of the BasicDetector is set to 60 seconds.
	Timeout time.Duration
	// NormalConditionFn is called with a copy of Counter whenever a request success in the anomaly state.
	// If NormalConditionFn returns true, the BasicDetector will be placed into the ok state.
	// If NormalConditionFn is nil, default NormalConditionFn is used.
	// Default NormalConditionFn returns true when the number of consecutive normalities is more than 3.
	NormalConditionFn func(counter Counter) bool
	// AnomalyConditionFn is called with a copy of Counter whenever a request fails in the ok state.
	// If AnomalyConditionFn returns true, the BasicDetector will be placed into the anomaly state.
	// If AnomalyConditionFn is nil, default AnomalyConditionFn is used.
	// Default AnomalyConditionFn returns true when the number of consecutive abnormalities is more than 5.
	AnomalyConditionFn func(counter Counter) bool
	// OnStateChange is called whenever the state of the BasicDetector changes.
	OnStateChange func(name string, from State, to State)
}

var _ Detector = &BasicDetector{}

// BasicDetector is a state machine to prevent sending requests that are likely to fail.
type BasicDetector struct {
	name               string
	timeout            time.Duration
	anomalyConditionFn func(counts Counter) bool
	normalConditionFn  func(counts Counter) bool
	onStateChange      func(name string, from State, to State)

	mutex      sync.Mutex
	state      State
	generation uint64
	counter    Counter
	expiration time.Time
}

// NewBasicDetector returns a new BasicDetector configured with the given Options.
func NewBasicDetector(name string, opts Options) *BasicDetector {
	_ = "STUB: not implemented"
	return nil
}

// Name returns the name of the BasicDetector.
func (d *BasicDetector) Name() string {
	_ = "STUB: not implemented"

	// State returns the current state of the BasicDetector.
	return ""
}

func (d *BasicDetector) State() State { _ = "STUB: not implemented"; return *new(State) }

// Counter returns internal counters
func (d *BasicDetector) Counter() Counter { _ = "STUB: not implemented"; return *new(Counter) }

func (d *BasicDetector) Mark(normality bool) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (d *BasicDetector) Reset() { _ = "STUB: not implemented"; return }

func (d *BasicDetector) onNormality(state State, now time.Time) { _ = "STUB: not implemented"; return }

func (d *BasicDetector) onAbnormalities(state State, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (d *BasicDetector) currentState(now time.Time) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (d *BasicDetector) setState(state State, now time.Time) { _ = "STUB: not implemented"; return }

func (d *BasicDetector) toNewGeneration(now time.Time) { _ = "STUB: not implemented"; return }
