/*
Copyright 2022 The Koordinator Authors.
Copyright 2014 The Kubernetes Authors.

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
	"context"
	"io"
	"os/exec"
	"time"

	gomegatypes "github.com/onsi/gomega/types"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	watchtools "k8s.io/client-go/tools/watch"
	"k8s.io/component-base/featuregate"

	// TODO: Remove the following imports (ref: https://github.com/kubernetes/kubernetes/issues/81245)

	imageutils "github.com/koordinator-sh/koordinator/test/utils/image"
)

const (
	// Minimal number of nodes for the cluster to be considered large.
	largeClusterThreshold = 100

	// TODO(justinsb): Avoid hardcoding this.
	awsMasterIP = "172.20.0.9"

	// AllContainers specifies that all containers be visited
	// Copied from pkg/api/v1/pod to avoid pulling extra dependencies
	AllContainers = InitContainers | Containers | EphemeralContainers
)

// DEPRECATED constants. Use the timeouts in framework.Framework instead.
const (
	// PodListTimeout is how long to wait for the pod to be listable.
	PodListTimeout = time.Minute

	// PodStartTimeout is how long to wait for the pod to be started.
	PodStartTimeout = 5 * time.Minute

	// PodStartShortTimeout is same as `PodStartTimeout` to wait for the pod to be started, but shorter.
	// Use it case by case when we are sure pod start will not be delayed.
	// minutes by slow docker pulls or something else.
	PodStartShortTimeout = 2 * time.Minute

	// PodDeleteTimeout is how long to wait for a pod to be deleted.
	PodDeleteTimeout = 5 * time.Minute

	// PodGetTimeout is how long to wait for a pod to be got.
	PodGetTimeout = 2 * time.Minute

	// PodEventTimeout is how much we wait for a pod event to occur.
	PodEventTimeout = 2 * time.Minute

	// ServiceStartTimeout is how long to wait for a service endpoint to be resolvable.
	ServiceStartTimeout = 3 * time.Minute

	// Poll is how often to Poll pods, nodes and claims.
	Poll = 2 * time.Second

	// PollShortTimeout is the short timeout value in polling.
	PollShortTimeout = 1 * time.Minute

	// ServiceAccountProvisionTimeout is how long to wait for a service account to be provisioned.
	// service accounts are provisioned after namespace creation
	// a service account is required to support pod creation in a namespace as part of admission control
	ServiceAccountProvisionTimeout = 2 * time.Minute

	// SingleCallTimeout is how long to try single API calls (like 'get' or 'list'). Used to prevent
	// transient failures from failing tests.
	SingleCallTimeout = 5 * time.Minute

	// NodeReadyInitialTimeout is how long nodes have to be "ready" when a test begins. They should already
	// be "ready" before the test starts, so this is small.
	NodeReadyInitialTimeout = 20 * time.Second

	// PodReadyBeforeTimeout is how long pods have to be "ready" when a test begins.
	PodReadyBeforeTimeout = 5 * time.Minute

	// ClaimProvisionShortTimeout is same as `ClaimProvisionTimeout` to wait for claim to be dynamically provisioned, but shorter.
	// Use it case by case when we are sure this timeout is enough.
	ClaimProvisionShortTimeout = 1 * time.Minute

	// ClaimProvisionTimeout is how long claims have to become dynamically provisioned.
	ClaimProvisionTimeout = 5 * time.Minute

	// RestartNodeReadyAgainTimeout is how long a node is allowed to become "Ready" after it is restarted before
	// the test is considered failed.
	RestartNodeReadyAgainTimeout = 5 * time.Minute

	// RestartPodReadyAgainTimeout is how long a pod is allowed to become "running" and "ready" after a node
	// restart before test is considered failed.
	RestartPodReadyAgainTimeout = 5 * time.Minute

	// SnapshotCreateTimeout is how long for snapshot to create snapshotContent.
	SnapshotCreateTimeout = 5 * time.Minute

	// SnapshotDeleteTimeout is how long for snapshot to delete snapshotContent.
	SnapshotDeleteTimeout = 5 * time.Minute
)

var (
	// BusyBoxImage is the image URI of BusyBox.
	BusyBoxImage = imageutils.GetE2EImage(imageutils.BusyBox)

	// ProvidersWithSSH are those providers where each node is accessible with SSH
	ProvidersWithSSH = []string{"gce", "gke", "aws", "local"}

	// ServeHostnameImage is a serve hostname image name.
	ServeHostnameImage = imageutils.GetE2EImage(imageutils.Agnhost)
)

// RunID is a unique identifier of the e2e run.
// Beware that this ID is not the same for all tests in the e2e run, because each Ginkgo node creates it separately.
var RunID = uuid.NewUUID()

// CreateTestingNSFn is a func that is responsible for creating namespace used for executing e2e tests.
type CreateTestingNSFn func(baseName string, c clientset.Interface, labels map[string]string) (*v1.Namespace, error)

// APIAddress returns a address of an instance.
func APIAddress() string { _ = "STUB: not implemented"; return "" }

// ProviderIs returns true if the provider is included is the providers. Otherwise false.
func ProviderIs(providers ...string) bool { _ = "STUB: not implemented"; return false }

// MasterOSDistroIs returns true if the master OS distro is included in the supportedMasterOsDistros. Otherwise false.
func MasterOSDistroIs(supportedMasterOsDistros ...string) bool {
	_ = "STUB: not implemented"
	return false
}

// NodeOSDistroIs returns true if the node OS distro is included in the supportedNodeOsDistros. Otherwise false.
func NodeOSDistroIs(supportedNodeOsDistros ...string) bool { _ = "STUB: not implemented"; return false }

// NodeOSArchIs returns true if the node OS arch is included in the supportedNodeOsArchs. Otherwise false.
func NodeOSArchIs(supportedNodeOsArchs ...string) bool { _ = "STUB: not implemented"; return false }

// DeleteNamespaces deletes all namespaces that match the given delete and skip filters.
// Filter is by simple strings.Contains; first skip filter, then delete filter.
// Returns the list of deleted namespaces or an error.
func DeleteNamespaces(c clientset.Interface, deleteFilter, skipFilter []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WaitForNamespacesDeleted waits for the namespaces to be deleted.
func WaitForNamespacesDeleted(c clientset.Interface, namespaces []string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Now POLL until all namespaces have been eradicated.

func waitForServiceAccountInNamespace(c clientset.Interface, ns, serviceAccountName string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// serviceAccountHasSecrets returns true if the service account has at least one secret,
// false if it does not, or an error.
func serviceAccountHasSecrets(event watch.Event) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// WaitForDefaultServiceAccountInNamespace waits for the default service account to be provisioned
// the default service account is what is associated with pods when they do not specify a service account
// as a result, pods are not able to be provisioned in a namespace until the service account is provisioned
func WaitForDefaultServiceAccountInNamespace(c clientset.Interface, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTestingNS should be used by every test, note that we append a common prefix to the provided test name.
// Please see NewFramework instead of using this directly.
func CreateTestingNS(baseName string, c clientset.Interface, labels map[string]string) (*v1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We don't use ObjectMeta.GenerateName feature, as in case of API call
// failure we don't know whether the namespace was created and what is its
// name.

// Be robust about making the namespace creation call.

// regenerate on conflict

// Even if we fail to create serviceAccount in the namespace,
// we have successfully create a namespace.
// So, return the created namespace.

// CheckTestingNSDeletedExcept checks whether all e2e based existing namespaces are in the Terminating state
// and waits until they are finally deleted. It ignores namespace skip.
func CheckTestingNSDeletedExcept(c clientset.Interface, skip string) error {
	_ = "STUB: not implemented"
	// TODO: Since we don't have support for bulk resource deletion in the API,
	// while deleting a namespace we are deleting all objects from that namespace
	// one by one (one deletion == one API call). This basically exposes us to
	// throttling - currently controller-manager has a limit of max 20 QPS.
	// Once #10217 is implemented and used in namespace-controller, deleting all
	// object from a given namespace should be much faster and we will be able
	// to lower this timeout.
	// However, now Density test is producing ~26000 events and Load capacity test
	// is producing ~35000 events, thus assuming there are no other requests it will
	// take ~30 minutes to fully delete the namespace. Thus I'm setting it to 60
	// minutes to avoid any timeouts here.
	return nil
}

// WaitForServiceEndpointsNum waits until the amount of endpoints that implement service to expectNum.
func WaitForServiceEndpointsNum(c clientset.Interface, namespace, serviceName string, expectNum int, interval, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func countEndpointsNum(e *v1.Endpoints) int { _ = "STUB: not implemented"; return 0 }

// restclientConfig returns a config holds the information needed to build connection to kubernetes clusters.
func restclientConfig(kubeContext string) (*clientcmdapi.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientConfigGetter is a func that returns getter to return a config.
type ClientConfigGetter func() (*restclient.Config, error)

// LoadConfig returns a config for a rest client with the UserAgent set to include the current test name.
func LoadConfig() (config *restclient.Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is a node e2e test, apply the node e2e configuration

// In case Host is not set in TestContext, sets it as
// CurrentContext Server for k8s API client to connect to.

// LoadClientset returns clientset for connecting to kubernetes clusters.
func LoadClientset() (*clientset.Clientset, error) { _ = "STUB: not implemented"; return nil, nil }

// RandomSuffix provides a random sequence to append to pods,services,rcs.
func RandomSuffix() string { _ = "STUB: not implemented"; return "" }

// LookForStringInPodExec looks for the given string in the output of a command
// executed in the first container of specified pod.
// TODO(alejandrox1): move to pod/ subpkg once kubectl methods are refactored.
func LookForStringInPodExec(ns, podName string, command []string, expectedString string, timeout time.Duration) (result string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LookForStringInPodExecToContainer looks for the given string in the output of a
// command executed in specified pod container, or first container if not specified.
func LookForStringInPodExecToContainer(ns, podName, containerName string, command []string, expectedString string, timeout time.Duration) (result string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// lookForString looks for the given string in the output of fn, repeatedly calling fn until
// the timeout is reached or the string is found. Returns last log and possibly
// error if the string was not found.
// TODO(alejandrox1): move to pod/ subpkg once kubectl methods are refactored.
func lookForString(expectedString string, timeout time.Duration, fn func() string) (result string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// KubectlBuilder is used to build, customize and execute a kubectl Command.
// Add more functions to customize the builder as needed.
type KubectlBuilder struct {
	cmd     *exec.Cmd
	timeout <-chan time.Time
}

// NewKubectlCommand returns a KubectlBuilder for running kubectl.
func NewKubectlCommand(namespace string, args ...string) *KubectlBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithEnv sets the given environment and returns itself.
func (b *KubectlBuilder) WithEnv(env []string) *KubectlBuilder {
	_ = "STUB: not implemented"
	return nil

	// WithTimeout sets the given timeout and returns itself.
}

func (b *KubectlBuilder) WithTimeout(t <-chan time.Time) *KubectlBuilder {
	_ = "STUB: not implemented"
	return nil

	// WithStdinData sets the given data to stdin and returns itself.
}

func (b KubectlBuilder) WithStdinData(data string) *KubectlBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithStdinReader sets the given reader and returns itself.
func (b KubectlBuilder) WithStdinReader(reader io.Reader) *KubectlBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ExecOrDie runs the kubectl executable or dies if error occurs.
func (b KubectlBuilder) ExecOrDie(namespace string) string {
	_ = "STUB: not implemented"

	// In case of i/o timeout error, try talking to the apiserver again after 2s before dying.
	// Note that we're still dying after retrying so that we can get visibility to triage it further.
	return ""
}

func isTimeout(err error) bool { _ = "STUB: not implemented"; return false }

// Exec runs the kubectl executable.
func (b KubectlBuilder) Exec() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ExecWithFullOutput runs the kubectl executable, and returns the stdout and stderr.
func (b KubectlBuilder) ExecWithFullOutput() (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// skip arg[0] as it is printed separately

// RunKubectlOrDie is a convenience wrapper over kubectlBuilder
func RunKubectlOrDie(namespace string, args ...string) string { _ = "STUB: not implemented"; return "" }

// RunKubectl is a convenience wrapper over kubectlBuilder
func RunKubectl(namespace string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RunKubectlWithFullOutput is a convenience wrapper over kubectlBuilder
// It will also return the command's stderr.
func RunKubectlWithFullOutput(namespace string, args ...string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// RunKubectlOrDieInput is a convenience wrapper over kubectlBuilder that takes input to stdin
func RunKubectlOrDieInput(namespace string, data string, args ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// RunKubectlInput is a convenience wrapper over kubectlBuilder that takes input to stdin
func RunKubectlInput(namespace string, data string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RunKubemciWithKubeconfig is a convenience wrapper over RunKubemciCmd
func RunKubemciWithKubeconfig(args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RunKubemciCmd is a convenience wrapper over kubectlBuilder to run kubemci.
// It assumes that kubemci exists in PATH.
func RunKubemciCmd(args ...string) (string, error) {
	_ = "STUB: not implemented"
	// kubemci is assumed to be in PATH.
	return "", nil
}

// StartCmdAndStreamOutput returns stdout and stderr after starting the given cmd.
func StartCmdAndStreamOutput(cmd *exec.Cmd) (stdout, stderr io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), *new(io.ReadCloser), nil
}

// TryKill is rough equivalent of ctrl+c for cleaning up processes. Intended to be run in defer.
func TryKill(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }

// testContainerOutputMatcher runs the given pod in the given namespace and waits
// for all of the containers in the podSpec to move into the 'Success' status, and tests
// the specified container log against the given expected output using the given matcher.
func (f *Framework) testContainerOutputMatcher(scenarioName string,
	pod *v1.Pod,
	containerIndex int,
	expectedOutput []string,
	matcher func(string, ...interface{}) gomegatypes.GomegaMatcher) {
	_ = "STUB: not implemented"
	return
}

// ContainerType signifies container type
type ContainerType int

const (
	// FeatureEphemeralContainers allows running an ephemeral container in pod namespaces to troubleshoot a running pod
	FeatureEphemeralContainers featuregate.Feature = "EphemeralContainers"
	// Containers is for normal containers
	Containers ContainerType = 1 << iota
	// InitContainers is for init containers
	InitContainers
	// EphemeralContainers is for ephemeral containers
	EphemeralContainers
)

// allFeatureEnabledContainers returns a ContainerType mask which includes all container
// types except for the ones guarded by feature gate.
// Copied from pkg/api/v1/pod to avoid pulling extra dependencies
func allFeatureEnabledContainers() ContainerType {
	_ = "STUB: not implemented"
	return *new(ContainerType)
}

// ContainerVisitor is called with each container spec, and returns true
// if visiting should continue.
// Copied from pkg/api/v1/pod to avoid pulling extra dependencies
type ContainerVisitor func(container *v1.Container, containerType ContainerType) (shouldContinue bool)

// visitContainers invokes the visitor function with a pointer to every container
// spec in the given pod spec with type set in mask. If visitor returns false,
// visiting is short-circuited. visitContainers returns true if visiting completes,
// false if visiting was short-circuited.
// Copied from pkg/api/v1/pod to avoid pulling extra dependencies
func visitContainers(podSpec *v1.PodSpec, mask ContainerType, visitor ContainerVisitor) bool {
	_ = "STUB: not implemented"
	return false
}

// MatchContainerOutput creates a pod and waits for all it's containers to exit with success.
// It then tests that the matcher with each expectedOutput matches the output of the specified container.
func (f *Framework) MatchContainerOutput(
	pod *v1.Pod,
	containerName string,
	expectedOutput []string,
	matcher func(string, ...interface{}) gomegatypes.GomegaMatcher) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for client pod to complete.

// Grab its logs.  Get host first.

// Pod failed. Dump all logs from all containers to see what's wrong

// Sometimes the actual containers take a second to get started, try to get logs for 60s

// EventsLister is a func that lists events.
type EventsLister func(opts metav1.ListOptions, ns string) (*v1.EventList, error)

// dumpEventsInNamespace dumps events in the given namespace.
func dumpEventsInNamespace(eventsLister EventsLister, namespace string) {
	_ = "STUB: not implemented"
	return
}

// Sort events by their first timestamp

// Note that we don't wait for any Cleanup to propagate, which means
// that if you delete a bunch of pods right before ending your test,
// you may or may not see the killing/deletion/Cleanup events.

// DumpAllNamespaceInfo dumps events, pods and nodes information in the given namespace.
func DumpAllNamespaceInfo(c clientset.Interface, namespace string) {
	_ = "STUB: not implemented"
	return
}

// If cluster is large, then the following logs are basically useless, because:
// 1. it takes tens of minutes or hours to grab all of them
// 2. there are so many of them that working with them are mostly impossible
// So we dump them only if the cluster is relatively small.

// byFirstTimestamp sorts a slice of events by first timestamp, using their involvedObject's name as a tie breaker.
type byFirstTimestamp []v1.Event

func (o byFirstTimestamp) Len() int      { _ = "STUB: not implemented"; return 0 }
func (o byFirstTimestamp) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (o byFirstTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func dumpAllNodeInfo(c clientset.Interface, nodes *v1.NodeList) { _ = "STUB: not implemented"; return }

// DumpNodeDebugInfo dumps debug information of the given nodes.
func DumpNodeDebugInfo(c clientset.Interface, nodeNames []string, logFunc func(fmt string, args ...interface{})) {
	_ = "STUB: not implemented"
	return
}

// TODO: Log node resource info

// getKubeletPods retrieves the list of pods on the kubelet.
func getKubeletPods(c clientset.Interface, node string) (*v1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// call chain tends to hang in some cases when Node is not ready. Add an artificial timeout for this call. #22165

// logNodeEvents logs kubelet events from the given node. This includes kubelet
// restart and node unhealthy events. Note that listing events like this will mess
// with latency metrics, beware of calling it during a test.
func getNodeEvents(c clientset.Interface, nodeName string) []v1.Event {
	_ = "STUB: not implemented"
	return nil
}

// WaitForAllNodesSchedulable waits up to timeout for all
// (but TestContext.AllowedNotReadyNodes) to become schedulable.
func WaitForAllNodesSchedulable(c clientset.Interface, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// AddOrUpdateLabelOnNode adds the given label key and value to the given node or updates value.
func AddOrUpdateLabelOnNode(c clientset.Interface, nodeName string, labelKey, labelValue string) {
	_ = "STUB: not implemented"
	return
}

// ExpectNodeHasLabel expects that the given node has the given label pair.
func ExpectNodeHasLabel(c clientset.Interface, nodeName string, labelKey string, labelValue string) {
	_ = "STUB: not implemented"
	return
}

// RemoveLabelOffNode is for cleaning up labels temporarily added to node,
// won't fail if target label doesn't exist or has been removed.
func RemoveLabelOffNode(c clientset.Interface, nodeName string, labelKey string) {
	_ = "STUB: not implemented"
	return
}

// ExpectNodeHasTaint expects that the node has the given taint.
func ExpectNodeHasTaint(c clientset.Interface, nodeName string, taint *v1.Taint) {
	_ = "STUB: not implemented"
	return
}

// NodeHasTaint returns true if the node has the given taint, else returns false.
func NodeHasTaint(c clientset.Interface, nodeName string, taint *v1.Taint) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RunHostCmd runs the given cmd in the context of the given pod using `kubectl exec`
// inside of a shell.
func RunHostCmd(ns, name, cmd string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// RunHostCmdWithFullOutput runs the given cmd in the context of the given pod using `kubectl exec`
// inside of a shell. It will also return the command's stderr.
func RunHostCmdWithFullOutput(ns, name, cmd string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// RunHostCmdOrDie calls RunHostCmd and dies on error.
func RunHostCmdOrDie(ns, name, cmd string) string { _ = "STUB: not implemented"; return "" }

// RunHostCmdWithRetries calls RunHostCmd and retries all errors
// until it succeeds or the specified timeout expires.
// This can be used with idempotent commands to deflake transient Node issues.
func RunHostCmdWithRetries(ns, name, cmd string, interval, timeout time.Duration) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AllNodesReady checks whether all registered nodes are ready. Setting -1 on
// TestContext.AllowedNotReadyNodes will bypass the post test node readiness check.
// TODO: we should change the AllNodesReady call in AfterEach to WaitForAllNodesHealthy,
// and figure out how to do it in a configurable way, as we can't expect all setups to run
// default test add-ons.
func AllNodesReady(c clientset.Interface, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// It should be OK to list unschedulable Nodes here.

// Framework allows for <TestContext.AllowedNotReadyNodes> nodes to be non-ready,
// to make it possible e.g. for incorrect deployment of some small percentage
// of nodes (which we allow in cluster validation). Some nodes that are not
// provisioned correctly at startup will never become ready (e.g. when something
// won't install correctly), so we can't expect them to be ready at any point.

// LookForStringInLog looks for the given string in the log of a specific pod container
func LookForStringInLog(ns, podName, container, expectedString string, timeout time.Duration) (result string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// EnsureLoadBalancerResourcesDeleted ensures that cloud load balancer resources that were created
// are actually cleaned up.  Currently only implemented for GCE/GKE.
func EnsureLoadBalancerResourcesDeleted(ip, portRange string) error {
	_ = "STUB: not implemented"
	return nil
}

// CoreDump SSHs to the master and all nodes and dumps their logs into dir.
// It shells out to cluster/log-dump/log-dump.sh to accomplish this.
func CoreDump(dir string) { _ = "STUB: not implemented"; return }

// parseSystemdServices converts services separator from comma to space.
func parseSystemdServices(services string) string { _ = "STUB: not implemented"; return "" }

// RunCmd runs cmd using args and returns its stdout and stderr. It also outputs
// cmd's stdout and stderr to their respective OS streams.
func RunCmd(command string, args ...string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// RunCmdEnv runs cmd with the provided environment and args and
// returns its stdout and stderr. It also outputs cmd's stdout and
// stderr to their respective OS streams.
func RunCmdEnv(env []string, command string, args ...string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// We also output to the OS stdout/stderr to aid in debugging in case cmd
// hangs and never returns before the test gets killed.
//
// This creates some ugly output because gcloud doesn't always provide
// newlines.

// getControlPlaneAddresses returns the externalIP, internalIP and hostname fields of control plane nodes.
// If any of these is unavailable, empty slices are returned.
func getControlPlaneAddresses(c clientset.Interface) ([]string, []string, []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Populate the internal IPs.

// Populate the external IP/hostname.

// GetControlPlaneAddresses returns all IP addresses on which the kubelet can reach the control plane.
// It may return internal and external IPs, even if we expect for
// e.g. internal IPs to be used (issue #56787), so that we can be
// sure to block the control plane fully during tests.
func GetControlPlaneAddresses(c clientset.Interface) []string {
	_ = "STUB: not implemented"
	return nil
}

// CreateEmptyFileOnPod creates empty file at given path on the pod.
// TODO(alejandrox1): move to subpkg pod once kubectl methods have been refactored.
func CreateEmptyFileOnPod(namespace string, podName string, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// DumpDebugInfo dumps debug info of tests.
func DumpDebugInfo(c clientset.Interface, ns string) { _ = "STUB: not implemented"; return }

// PrettyPrintJSON converts metrics to JSON format.
func PrettyPrintJSON(metrics interface{}) string { _ = "STUB: not implemented"; return "" }

// taintExists checks if the given taint exists in list of taints. Returns true if exists false otherwise.
func taintExists(taints []v1.Taint, taintToFind *v1.Taint) bool {
	_ = "STUB: not implemented"
	return false
}

// WatchEventSequenceVerifier ...
// manages a watch for a given resource, ensures that events take place in a given order, retries the test on failure
//
//	testContext         cancelation signal across API boundries, e.g: context.TODO()
//	dc                  sets up a client to the API
//	resourceType        specify the type of resource
//	namespace           select a namespace
//	resourceName        the name of the given resource
//	listOptions         options used to find the resource, recommended to use listOptions.labelSelector
//	expectedWatchEvents array of events which are expected to occur
//	scenario            the test itself
//	retryCleanup        a function to run which ensures that there are no dangling resources upon test failure
//
// this tooling relies on the test to return the events as they occur
// the entire scenario must be run to ensure that the desired watch events arrive in order (allowing for interweaving of watch events)
//
//	if an expected watch event is missing we elect to clean up and run the entire scenario again
//
// we try the scenario three times to allow the sequencing to fail a couple of times
func WatchEventSequenceVerifier(ctx context.Context, dc dynamic.Interface, resourceType schema.GroupVersionResource, namespace string, resourceName string, listOptions metav1.ListOptions, expectedWatchEvents []watch.Event, scenario func(*watchtools.RetryWatcher) []watch.Event, retryCleanup func() error) {
	_ = "STUB: not implemented"
	return
}

// NOTE the test may need access to the events to see what's going on, such as a change in status
