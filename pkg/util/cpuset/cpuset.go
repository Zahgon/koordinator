/*
Copyright 2022 The Koordinator Authors.
Copyright 2017 The Kubernetes Authors.

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

package cpuset

const (
	maxAvailableCPUCount = 4096
)

// CPUSetBuilder is a mutable builder for CPUSet.
// Functions that mutate instances of this type are not thread-safe.
type CPUSetBuilder struct {
	result CPUSet
	done   bool
}

// NewCPUSetBuilder returns a mutable CPUSet builder.
func NewCPUSetBuilder() *CPUSetBuilder { _ = "STUB: not implemented"; return nil }

// Add adds the supplied elements to the result.
// Calling Add after calling Result has no effect.
func (b *CPUSetBuilder) Add(elems ...int) { _ = "STUB: not implemented"; return }

// Result returns the result CPUSet containing all elements that were
// previously added to this builder. Subsequent calls to Add have no effect.
func (b *CPUSetBuilder) Result() CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// CPUSet is a thread-safe, immutable set-like data structure for CPU IDs.
type CPUSet struct {
	elems map[int]struct{}
}

// NewCPUSet returns a new CPUSet containing the supplied elements.
func NewCPUSet(cpus ...int) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// Clone returns a copy of this CPUSet.
func (s CPUSet) Clone() CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// Size returns the number of elements in this CPUSet.
func (s CPUSet) Size() int { _ = "STUB: not implemented"; return 0 }

// IsEmpty returns true if there are zero elements in this CPUSet.
func (s CPUSet) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Contains returns true if the supplied element is present in this CPUSet.
func (s CPUSet) Contains(cpu int) bool { _ = "STUB: not implemented"; return false }

// Equals returns true if the supplied CPUSet contains exactly the same elements
// as this CPUSet (s IsSubsetOf s2 and s2 IsSubsetOf s).
func (s CPUSet) Equals(s2 CPUSet) bool { _ = "STUB: not implemented"; return false }

// Filter returns a new CPUSet that contains the elements from this
// CPUSet that match the supplied predicate, without mutating the source CPUSet.
func (s CPUSet) Filter(predicate func(int) bool) CPUSet {
	_ = "STUB: not implemented"
	return *new(CPUSet)
}

// FilterNot returns a new CPUSet that contains the elements from this
// CPUSet that do not match the supplied predicate, without mutating the source CPUSet.
func (s CPUSet) FilterNot(predicate func(int) bool) CPUSet {
	_ = "STUB: not implemented"
	return *new(CPUSet)
}

// IsSubsetOf returns true if the supplied CPUSet contains all the elements
func (s CPUSet) IsSubsetOf(s2 CPUSet) bool { _ = "STUB: not implemented"; return false }

// Union returns a new CPUSet that contains the elements from this CPUSet
// and the elements from the supplied CPUSet, without mutating either source CPUSet.
func (s CPUSet) Union(s2 CPUSet) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// UnionSlice returns a new CPUSet that contains the elements from this CPUSet
// and the elements from the supplied CPUSet, without mutating either source CPUSet.
func (s CPUSet) UnionSlice(s2 ...int) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// UnionAll returns a new CPUSet that contains the elements from this
// CPUSet and the elements from the supplied sets, without mutating either source CPUSet.
func (s CPUSet) UnionAll(s2 []CPUSet) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// Intersection returns a new CPUSet that contains the elements
// that are present in both this CPUSet and the supplied CPUSet, without mutating either source CPUSet.
func (s CPUSet) Intersection(s2 CPUSet) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// Difference returns a new CPUSet that contains the elements that
// are present in this CPUSet and not the supplied CPUSet, without mutating either source CPUSet.
func (s CPUSet) Difference(s2 CPUSet) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// ToSlice returns a slice of integers that contains all elements from this CPUSet.
func (s CPUSet) ToSlice() []int { _ = "STUB: not implemented"; return nil }

// ToSliceNoSort returns a slice of integers that contains all elements from this CPUSet.
func (s CPUSet) ToSliceNoSort() []int { _ = "STUB: not implemented"; return nil }

// ToInt32Slice returns a slice of int32 values that contains all elements from this CPUSet.
func (s CPUSet) ToInt32Slice() []int32 { _ = "STUB: not implemented"; return nil }

// assert cpu id is in int32 range

func (s CPUSet) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *CPUSet) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// String returns a new string representation of the elements in this CPUSet
// in canonical linux CPU list format.
//
// See: http://man7.org/linux/man-pages/man7/cpuset.7.html#FORMATS
func (s CPUSet) String() string { _ = "STUB: not implemented"; return "" }

// if this element is adjacent to the high end of the last range

// then extend the last range to include this element

// otherwise, start a new range beginning with this element

// construct string from ranges

// MustParse CPUSet constructs a new CPUSet from a Linux CPU list formatted
// string. Unlike Parse, it does not return an error but rather panics if the
// input cannot be used to construct a CPUSet.
func MustParse(s string) CPUSet { _ = "STUB: not implemented"; return *new(CPUSet) }

// Parse CPUSet constructs a new CPUSet from a Linux CPU list formatted string.
//
// See: http://man7.org/linux/man-pages/man7/cpuset.7.html#FORMATS
func Parse(s string) (CPUSet, error) {
	_ = "STUB: not implemented"
	return *

	// Handle empty string.
	new(CPUSet), nil
}

// Split CPU list string:
// "0-5,34,46-48 => ["0-5", "34", "46-48"]

// Handle ranges that consist of only one element like "34".
// assert cpu id is in range of int32

// Handle multi-element ranges like "0-5".
// assert cpu id is in range of int32

// Add all elements to the result.
// e.g. "0-5", "46-48" => [0, 1, 2, 3, 4, 5, 46, 47, 48].

func IsEqualStrCpus(a, b string) bool { _ = "STUB: not implemented"; return false }
