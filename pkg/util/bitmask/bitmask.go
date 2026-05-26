/*
Copyright 2019 The Kubernetes Authors.

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

package bitmask

// BitMask interface allows hint providers to create BitMasks for TopologyHints
type BitMask interface {
	Add(bits ...int) error
	Remove(bits ...int) error
	And(masks ...BitMask)
	Or(masks ...BitMask)
	Clear()
	Fill()
	IsEqual(mask BitMask) bool
	IsEmpty() bool
	IsSet(bit int) bool
	AnySet(bits []int) bool
	IsNarrowerThan(mask BitMask) bool
	IsLessThan(mask BitMask) bool
	IsGreaterThan(mask BitMask) bool
	String() string
	Count() int
	GetBits() []int
}

type bitMask uint64

// NewEmptyBitMask creates a new, empty BitMask
func NewEmptyBitMask() BitMask { _ = "STUB: not implemented"; return *new(BitMask) }

// NewBitMask creates a new BitMask
func NewBitMask(bits ...int) (BitMask, error) { _ = "STUB: not implemented"; return *new(BitMask), nil }

// Add adds the bits with topology affinity to the BitMask
func (s *bitMask) Add(bits ...int) error { _ = "STUB: not implemented"; return nil }

// Remove removes specified bits from BitMask
func (s *bitMask) Remove(bits ...int) error { _ = "STUB: not implemented"; return nil }

// And performs and operation on all bits in masks
func (s *bitMask) And(masks ...BitMask) { _ = "STUB: not implemented"; return }

// Or performs or operation on all bits in masks
func (s *bitMask) Or(masks ...BitMask) { _ = "STUB: not implemented"; return }

// Clear resets all bits in mask to zero
func (s *bitMask) Clear() {
	_ = "STUB: not implemented"

	// Fill sets all bits in mask to one
	return
}

func (s *bitMask) Fill() { _ = "STUB: not implemented"; return }

// IsEmpty checks mask to see if all bits are zero
func (s *bitMask) IsEmpty() bool {
	_ = "STUB: not implemented"

	// IsSet checks bit in mask to see if bit is set to one
	return false
}

func (s *bitMask) IsSet(bit int) bool { _ = "STUB: not implemented"; return false }

// AnySet checks bit in mask to see if any provided bit is set to one
func (s *bitMask) AnySet(bits []int) bool { _ = "STUB: not implemented"; return false }

// IsEqual checks if masks are equal
func (s *bitMask) IsEqual(mask BitMask) bool { _ = "STUB: not implemented"; return false }

// IsNarrowerThan checks if one mask is narrower than another.
//
// A mask is said to be "narrower" than another if it has lets bits set. If the
// same number of bits are set in both masks, then the mask with more
// lower-numbered bits set wins out.
func (s *bitMask) IsNarrowerThan(mask BitMask) bool { _ = "STUB: not implemented"; return false }

// IsLessThan checks which bitmask has more lower-numbered bits set.
func (s *bitMask) IsLessThan(mask BitMask) bool { _ = "STUB: not implemented"; return false }

// IsGreaterThan checks which bitmask has more higher-numbered bits set.
func (s *bitMask) IsGreaterThan(mask BitMask) bool { _ = "STUB: not implemented"; return false }

// String converts mask to string
func (s *bitMask) String() string { _ = "STUB: not implemented"; return "" }

// Count counts number of bits in mask set to one
func (s *bitMask) Count() int { _ = "STUB: not implemented"; return 0 }

// GetBits returns each bit number with bits set to one
func (s *bitMask) GetBits() []int { _ = "STUB: not implemented"; return nil }

// And is a package level implementation of 'and' between first and masks
func And(first BitMask, masks ...BitMask) BitMask { _ = "STUB: not implemented"; return *new(BitMask) }

// Or is a package level implementation of 'or' between first and masks
func Or(first BitMask, masks ...BitMask) BitMask { _ = "STUB: not implemented"; return *new(BitMask) }

// IterateBitMasks iterates all possible masks from a list of bits,
// issuing a callback on each mask.
func IterateBitMasks(bits []int, callback func(BitMask)) { _ = "STUB: not implemented"; return }
