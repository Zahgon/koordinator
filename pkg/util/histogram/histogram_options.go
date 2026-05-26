/*
Copyright 2023 The Koordinator Authors.
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

package histogram

// HistogramOptions define the number and size of buckets of a histogram.
type HistogramOptions interface {
	// Returns the number of buckets in the histogram.
	NumBuckets() int
	// Returns the index of the bucket to which the given value falls.
	// If the value is outside of the range covered by the histogram, it
	// returns the closest bucket (either the first or the last one).
	FindBucket(value float64) int
	// Returns the start of the bucket with a given index. If the index is
	// outside the [0..NumBuckets() - 1] range, the result is undefined.
	GetBucketStart(bucket int) float64
	// Returns the minimum weight for a bucket to be considered non-empty.
	Epsilon() float64
}

// NewLinearHistogramOptions returns HistogramOptions describing a histogram
// with a given number of fixed-size buckets, with the first bucket start at 0.0
// and the last bucket start larger or equal to maxValue.
// Requires maxValue > 0, bucketSize > 0, epsilon > 0.
func NewLinearHistogramOptions(
	maxValue float64, bucketSize float64, epsilon float64) (HistogramOptions, error) {
	_ = "STUB: not implemented"
	return *new(HistogramOptions), nil
}

// NewExponentialHistogramOptions returns HistogramOptions describing a
// histogram with exponentially growing bucket boundaries. The first bucket
// covers the range [0..firstBucketSize). Bucket with index n has size equal to firstBucketSize * ratio^n.
// It follows that the bucket with index n >= 1 starts at:
//
//	firstBucketSize * (1 + ratio + ratio^2 + ... + ratio^(n-1)) =
//	firstBucketSize * (ratio^n - 1) / (ratio - 1).
//
// The last bucket start is larger or equal to maxValue.
// Requires maxValue > 0, firstBucketSize > 0, ratio > 1, epsilon > 0.
func NewExponentialHistogramOptions(
	maxValue float64, firstBucketSize float64, ratio float64, epsilon float64) (HistogramOptions, error) {
	_ = "STUB: not implemented"
	return *new(HistogramOptions), nil
}

type linearHistogramOptions struct {
	numBuckets int
	bucketSize float64
	epsilon    float64
}

type exponentialHistogramOptions struct {
	numBuckets      int
	firstBucketSize float64
	ratio           float64
	epsilon         float64
}

func (o *linearHistogramOptions) NumBuckets() int { _ = "STUB: not implemented"; return 0 }

func (o *linearHistogramOptions) FindBucket(value float64) int { _ = "STUB: not implemented"; return 0 }

func (o *linearHistogramOptions) GetBucketStart(bucket int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (o *linearHistogramOptions) Epsilon() float64 { _ = "STUB: not implemented"; return 0 }

func (o *exponentialHistogramOptions) NumBuckets() int { _ = "STUB: not implemented"; return 0 }

// Returns the index of the bucket for given value. This is the inverse function to
// GetBucketStart(), which yields the following formula for the bucket index:
//
//	bucket(value) = floor(log(value/firstBucketSize*(ratio-1)+1) / log(ratio))
func (o *exponentialHistogramOptions) FindBucket(value float64) int {
	_ = "STUB: not implemented"
	return 0
}

// Returns the start of the bucket with given index, according to the formula:
//
//	bucketStart(bucket) = firstBucketSize * (ratio^bucket - 1) / (ratio - 1).
func (o *exponentialHistogramOptions) GetBucketStart(bucket int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (o *exponentialHistogramOptions) Epsilon() float64 {
	_ = "STUB: not implemented"

	// Returns the logarithm of x to given base, so that: base^log(base, x) == x.
	return 0
}

func log(base, x float64) float64 { _ = "STUB: not implemented"; return 0 }
