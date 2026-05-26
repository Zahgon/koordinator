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

package metrics

import (
	"errors"
	"sync"

	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	// kubeSchedulerPort is the default port for the scheduler status server.
	kubeSchedulerPort = 10259
	// kubeControllerManagerPort is the default port for the controller manager status server.
	kubeControllerManagerPort = 10257
	// snapshotControllerPort is the port for the snapshot controller
	snapshotControllerPort = 9102
)

// MetricsGrabbingDisabledError is an error that is wrapped by the
// different MetricsGrabber.Wrap functions when metrics grabbing is
// not supported. Tests that check metrics data should then skip
// the check.
var MetricsGrabbingDisabledError = errors.New("metrics grabbing disabled")

// Collection is metrics collection of components
type Collection struct {
	APIServerMetrics          APIServerMetrics
	ControllerManagerMetrics  ControllerManagerMetrics
	SnapshotControllerMetrics SnapshotControllerMetrics
	KubeletMetrics            map[string]KubeletMetrics
	SchedulerMetrics          SchedulerMetrics
	ClusterAutoscalerMetrics  ClusterAutoscalerMetrics
}

// Grabber provides functions which grab metrics from components
type Grabber struct {
	client                             clientset.Interface
	externalClient                     clientset.Interface
	config                             *rest.Config
	grabFromAPIServer                  bool
	grabFromControllerManager          bool
	grabFromKubelets                   bool
	grabFromScheduler                  bool
	grabFromClusterAutoscaler          bool
	grabFromSnapshotController         bool
	kubeScheduler                      string
	waitForSchedulerReadyOnce          sync.Once
	kubeControllerManager              string
	waitForControllerManagerReadyOnce  sync.Once
	snapshotController                 string
	waitForSnapshotControllerReadyOnce sync.Once
}

// NewMetricsGrabber prepares for grabbing metrics data from several different
// components. It should be called when those components are running because
// it needs to communicate with them to determine for which components
// metrics data can be retrieved.
//
// Collecting metrics data is an optional debug feature. Not all clusters will
// support it. If disabled for a component, the corresponding Grab function
// will immediately return an error derived from MetricsGrabbingDisabledError.
func NewMetricsGrabber(c clientset.Interface, ec clientset.Interface, config *rest.Config, kubelets bool, scheduler bool, controllers bool, apiServer bool, clusterAutoscaler bool, snapshotController bool) (*Grabber, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkPodDebugHandlers(c clientset.Interface, requested bool, component, podName string) bool {
	_ = "STUB: not implemented"
	return false
}

// The debug handlers on the host where the pod runs might be disabled.
// We can check that indirectly by trying to retrieve log output.

// Metrics gathering enabled.

// HasControlPlanePods returns true if metrics grabber was able to find control-plane pods
func (g *Grabber) HasControlPlanePods() bool { _ = "STUB: not implemented"; return false }

// GrabFromKubelet returns metrics from kubelet
func (g *Grabber) GrabFromKubelet(nodeName string) (KubeletMetrics, error) {
	_ = "STUB: not implemented"
	return *new(KubeletMetrics), nil
}

func (g *Grabber) grabFromKubeletInternal(nodeName string, kubeletPort int) (KubeletMetrics, error) {
	_ = "STUB: not implemented"
	return *new(KubeletMetrics), nil
}

// GrabFromScheduler returns metrics from scheduler
func (g *Grabber) GrabFromScheduler() (SchedulerMetrics, error) {
	_ = "STUB: not implemented"
	return *new(SchedulerMetrics), nil
}

// GrabFromClusterAutoscaler returns metrics from cluster autoscaler
func (g *Grabber) GrabFromClusterAutoscaler() (ClusterAutoscalerMetrics, error) {
	_ = "STUB: not implemented"
	return *new(ClusterAutoscalerMetrics), nil
}

// GrabFromControllerManager returns metrics from controller manager
func (g *Grabber) GrabFromControllerManager() (ControllerManagerMetrics, error) {
	_ = "STUB: not implemented"
	return *new(ControllerManagerMetrics), nil
}

// GrabFromSnapshotController returns metrics from controller manager
func (g *Grabber) GrabFromSnapshotController(podName string, port int) (SnapshotControllerMetrics, error) {
	_ = "STUB: not implemented"
	return *new(SnapshotControllerMetrics), nil
}

// Use overrides if provided via test config flags.
// Otherwise, use the default volume-snapshot-controller pod name and port.

// GrabFromAPIServer returns metrics from API server
func (g *Grabber) GrabFromAPIServer() (APIServerMetrics, error) {
	_ = "STUB: not implemented"
	return *new(APIServerMetrics), nil
}

// Grab returns metrics from corresponding component
func (g *Grabber) Grab() (Collection, error) {
	_ = "STUB: not implemented"
	return *new(Collection), nil
}

// getMetricsFromPod retrieves metrics data from an insecure port.
func (g *Grabber) getMetricsFromPod(client clientset.Interface, podName string, namespace string, port int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getSecureMetricsFromPod retrieves metrics from a pod that uses TLS
// and checks client credentials. Conceptually this function is
// similar to "kubectl port-forward" + "kubectl get --raw
// https://localhost:<port>/metrics". It uses the same credentials
// as kubelet.
func (g *Grabber) getSecureMetricsFromPod(podName string, namespace string, port int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// This should make it possible verify the server, but while it
// got past the server name check, certificate validation
// still failed.

// Verifying the pod certificate with the same root CA
// as for the API server led to an error about "unknown root
// certificate". Disabling certificate checking on the client
// side gets around that and should be good enough for
// E2E testing.

// clientset.NewForConfig is used because
// metricClient.RESTClient() is directly usable, in contrast
// to the client constructed by rest.RESTClientFor().
