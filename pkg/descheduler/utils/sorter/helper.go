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

package sorter

import (
	corev1 "k8s.io/api/core/v1"
)

// CompareFn compares p1 and p2 and returns:
//
//	-1 if p1 <  p2
//	 0 if p1 == p2
//	+1 if p1 >  p2
type CompareFn func(p1, p2 *corev1.Pod) int

// MultiSorter implements the Sort interface
type MultiSorter struct {
	ascending bool
	pods      []*corev1.Pod
	cmp       []CompareFn
}

// Sort sorts the pods according to the cmp functions passed to OrderedBy.
func (ms *MultiSorter) Sort(pods []*corev1.Pod) { _ = "STUB: not implemented"; return }

// OrderedBy returns a Sorter sorted using the cmp functions, sorts in ascending order by default
func OrderedBy(cmp ...CompareFn) *MultiSorter { _ = "STUB: not implemented"; return nil }

func (ms *MultiSorter) Ascending() *MultiSorter { _ = "STUB: not implemented"; return nil }

func (ms *MultiSorter) Descending() *MultiSorter { _ = "STUB: not implemented"; return nil }

// Len is part of sort.Interface.
func (ms *MultiSorter) Len() int { _ = "STUB: not implemented"; return 0 }

// Swap is part of sort.Interface.
func (ms *MultiSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Less is part of sort.Interface.
func (ms *MultiSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// p1 is less than p2

// p1 is greater than p2

// cmpBool compares booleans, placing true before false
func cmpBool(a, b bool) int { _ = "STUB: not implemented"; return 0 }

func Reverse(cmp CompareFn) CompareFn { _ = "STUB: not implemented"; return *new(CompareFn) }
