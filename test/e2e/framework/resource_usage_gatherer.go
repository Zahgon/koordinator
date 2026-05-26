/*
Copyright 2022 The Koordinator Authors.
Copyright 2015 The Kubernetes Authors.

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

package framework

import (
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	kubeletstatsv1alpha1 "k8s.io/kubelet/pkg/apis/stats/v1alpha1"
	// TODO: Remove the following imports (ref: https://github.com/kubernetes/kubernetes/issues/81245)
)

// ResourceConstraint is a struct to hold constraints.
type ResourceConstraint struct {
	CPUConstraint    float64
	MemoryConstraint uint64
}

// SingleContainerSummary is a struct to hold single container summary.
type SingleContainerSummary struct {
	Name string
	CPU  float64
	Mem  uint64
}

// ContainerResourceUsage is a structure for gathering container resource usage.
type ContainerResourceUsage struct {
	Name                    string
	Timestamp               time.Time
	CPUUsageInCores         float64
	MemoryUsageInBytes      uint64
	MemoryWorkingSetInBytes uint64
	MemoryRSSInBytes        uint64
	// The interval used to calculate CPUUsageInCores.
	CPUInterval time.Duration
}

// ResourceUsagePerContainer is map of ContainerResourceUsage
type ResourceUsagePerContainer map[string]*ContainerResourceUsage

// ResourceUsageSummary is a struct to hold resource usage summary.
// we can't have int here, as JSON does not accept integer keys.
type ResourceUsageSummary map[string][]SingleContainerSummary

// PrintHumanReadable prints resource usage summary in human readable.
func (s *ResourceUsageSummary) PrintHumanReadable() string { _ = "STUB: not implemented"; return "" }

// PrintJSON prints resource usage summary in JSON.
func (s *ResourceUsageSummary) PrintJSON() string { _ = "STUB: not implemented"; return "" }

// SummaryKind returns string of ResourceUsageSummary
func (s *ResourceUsageSummary) SummaryKind() string { _ = "STUB: not implemented"; return "" }

type uint64arr []uint64

func (a uint64arr) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a uint64arr) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a uint64arr) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type usageDataPerContainer struct {
	cpuData        []float64
	memUseData     []uint64
	memWorkSetData []uint64
}

func computePercentiles(timeSeries []ResourceUsagePerContainer, percentilesToCompute []int) map[int]ResourceUsagePerContainer {
	_ = "STUB: not implemented"
	return nil
}

func leftMergeData(left, right map[int]ResourceUsagePerContainer) map[int]ResourceUsagePerContainer {
	_ = "STUB: not implemented"
	return nil
}

type resourceGatherWorker struct {
	c                           clientset.Interface
	nodeName                    string
	wg                          *sync.WaitGroup
	containerIDs                []string
	stopCh                      chan struct{}
	dataSeries                  []ResourceUsagePerContainer
	finished                    bool
	inKubemark                  bool
	resourceDataGatheringPeriod time.Duration
	probeDuration               time.Duration
	printVerboseLogs            bool
}

func (w *resourceGatherWorker) singleProbe() { _ = "STUB: not implemented"; return }

// getOneTimeResourceUsageOnNode queries the node's /stats/summary endpoint
// and returns the resource usage of all containerNames for the past
// cpuInterval.
// The acceptable range of the interval is 2s~120s. Be warned that as the
// interval (and #containers) increases, the size of kubelet's response
// could be significant. E.g., the 60s interval stats for ~20 containers is
// ~1.5MB. Don't hammer the node with frequent, heavy requests.
//
// cadvisor records cumulative cpu usage in nanoseconds, so we need to have two
// stats points to compute the cpu usage over the interval. Assuming cadvisor
// polls every second, we'd need to get N stats points for N-second interval.
// Note that this is an approximation and may not be accurate, hence we also
// write the actual interval used for calculation (based on the timestamps of
// the stats points in ContainerResourceUsage.CPUInterval.
//
// containerNames is a function returning a collection of container names in which
// user is interested in.
func getOneTimeResourceUsageOnNode(
	c clientset.Interface,
	nodeName string,
	cpuInterval time.Duration,
	containerNames func() []string,
) (ResourceUsagePerContainer, error) {
	_ = "STUB: not implemented"

	// cadvisor records stats about every second.
	return *new(ResourceUsagePerContainer), nil
}

// cadvisor caches up to 2 minutes of stats (configured by kubelet).

// Get information of all containers on the node.

// Process container infos that are relevant to us.

// getStatsSummary contacts kubelet for the container information.
func getStatsSummary(c clientset.Interface, nodeName string) (*kubeletstatsv1alpha1.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeUint64Ptr(ptr *uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func (w *resourceGatherWorker) gather(initialSleep time.Duration) {
	_ = "STUB: not implemented"
	return
}

// ContainerResourceGatherer is a struct for gathering container resource.
type ContainerResourceGatherer struct {
	client       clientset.Interface
	stopCh       chan struct{}
	workers      []resourceGatherWorker
	workerWg     sync.WaitGroup
	containerIDs []string
	options      ResourceGathererOptions
}

// ResourceGathererOptions is a struct to hold options for resource.
type ResourceGathererOptions struct {
	InKubemark                  bool
	Nodes                       NodesSet
	ResourceDataGatheringPeriod time.Duration
	ProbeDuration               time.Duration
	PrintVerboseLogs            bool
}

// NodesSet is a value of nodes set.
type NodesSet int

const (
	// AllNodes means all containers on all nodes.
	AllNodes NodesSet = 0
	// MasterNodes means all containers on Master nodes only.
	MasterNodes NodesSet = 1
	// MasterAndDNSNodes means all containers on Master nodes and DNS containers on other nodes.
	MasterAndDNSNodes NodesSet = 2
)

// nodeHasControlPlanePods returns true if specified node has control plane pods
// (kube-scheduler and/or kube-controller-manager).
func nodeHasControlPlanePods(c clientset.Interface, nodeName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NewResourceUsageGatherer returns a new ContainerResourceGatherer.
func NewResourceUsageGatherer(c clientset.Interface, options ResourceGathererOptions, pods *v1.PodList) (*ContainerResourceGatherer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tracks kube-system pods if no valid PodList is passed in.

// StartGatheringData starts a stat gathering worker blocks for each node to track,
// and blocks until StopAndSummarize is called.
func (g *ContainerResourceGatherer) StartGatheringData() { _ = "STUB: not implemented"; return }

// StopAndSummarize stops stat gathering workers, processes the collected stats,
// generates resource summary for the passed-in percentiles, and returns the summary.
// It returns an error if the resource usage at any percentile is beyond the
// specified resource constraints.
func (g *ContainerResourceGatherer) StopAndSummarize(percentiles []int, constraints map[string]ResourceConstraint) (*ResourceUsageSummary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Workers has been stopped. We need to gather data stored in them.

// Verifying 99th percentile of resource usage

// Name has a form: <pod_name>/<container_name>

// kubemarkResourceUsage is a struct for tracking the resource usage of kubemark.
type kubemarkResourceUsage struct {
	Name                    string
	MemoryWorkingSetInBytes uint64
	CPUUsageInCores         float64
}

func getMasterUsageByPrefix(prefix string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getKubemarkMasterComponentsResourceUsage returns the resource usage of kubemark which contains multiple combinations of cpu and memory usage for each pod name.
func getKubemarkMasterComponentsResourceUsage() map[string]*kubemarkResourceUsage {
	_ = "STUB: not implemented"
	return nil
}

// Get kubernetes component resource usage

// Gatherer expects pod_name/container_name format

// Get etcd resource usage

// Gatherer expects pod_name/container_name format
