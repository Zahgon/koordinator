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

package frameworkext

import (
	prettytable "github.com/jedib0t/go-pretty/v6/table"
	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
)

var (
	debugTopNScores    = 0
	debugFilterFailure = false
)

// DebugScoresSetter updates debugTopNScores to specified value
func DebugScoresSetter(val string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// DebugFiltersSetter updates debugFilterFailure to specified value
func DebugFiltersSetter(val string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func debugScores(topN int, pod *corev1.Pod, allNodePluginScores []fwktype.NodePluginScores, nodeInfos []fwktype.NodeInfo) prettytable.Writer {
	_ = "STUB: not implemented"
	return *new(prettytable.Writer)
}

// Summarize all scores.

// return writer for UT
