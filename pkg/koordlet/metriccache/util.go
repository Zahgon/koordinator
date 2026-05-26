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

package metriccache

func fieldAvgOfMetricList(metricsList interface{}, aggregateParam AggregateParam) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// convert to struct for list with ptr

func fieldPercentileOfMetricList(metricsList interface{}, aggregateParam AggregateParam, percentile float32) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NOTE: use a general sort for the small case

// convert to struct for list with ptr

func fieldLastOfMetricList(metricsList interface{}, aggregateParam AggregateParam) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// convert to struct for list with ptr

func fieldCountOfMetricList(metricsList interface{}, aggregateParam AggregateParam) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func percentileFuncOfMetricList(percentile float32) AggregationFunc {
	_ = "STUB: not implemented"
	return *new(AggregationFunc)
}

func fieldLastOfMetricListBool(metricsList interface{}, aggregateParam AggregateParam) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// convert to struct for list with ptr
