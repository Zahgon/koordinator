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

package skipper

import (
	"regexp"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilversion "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/component-base/featuregate"
)

// New local storage types to support local storage capacity isolation
var localStorageCapacityIsolation featuregate.Feature = "LocalStorageCapacityIsolation"

var (
	downwardAPIHugePages featuregate.Feature = "DownwardAPIHugePages"
	execProbeTimeout     featuregate.Feature = "ExecProbeTimeout"
	csiMigration         featuregate.Feature = "CSIMigration"
)

func skipInternalf(caller int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// SkipPanic is the value that will be panicked from Skip.
type SkipPanic struct {
	Message        string // The failure message passed to Fail
	Filename       string // The filename that is the source of the failure
	Line           int    // The line number of the filename that is the source of the failure
	FullStackTrace string // A full stack trace starting at the source of the failure
}

// String makes SkipPanic look like the old Ginkgo panic when printed.
func (SkipPanic) String() string { _ = "STUB: not implemented"; return "" }

// Skip wraps ginkgo.Skip so that it panics with more useful
// information about why the test is being skipped. This function will
// panic with a SkipPanic.
func skip(message string, callerSkip ...int) { _ = "STUB: not implemented"; return }

// ginkgo adds a lot of test running infrastructure to the stack, so
// we filter those out
var stackSkipPattern = regexp.MustCompile(`onsi/ginkgo`)

func pruneStack(skip int) string {
	_ = "STUB: not implemented"
	// one for pruneStack and one for debug.Stack
	return ""
}

// skip the top of the stack

// these come in pairs

// these come in pairs

// Skipf skips with information about why the test is being skipped.
func Skipf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// SkipUnlessAtLeast skips if the value is less than the minValue.
func SkipUnlessAtLeast(value int, minValue int, message string) { _ = "STUB: not implemented"; return }

// SkipUnlessLocalEphemeralStorageEnabled skips if the LocalStorageCapacityIsolation is not enabled.
func SkipUnlessLocalEphemeralStorageEnabled() { _ = "STUB: not implemented"; return }

func SkipUnlessDownwardAPIHugePagesEnabled() { _ = "STUB: not implemented"; return }

func SkipUnlessExecProbeTimeoutEnabled() { _ = "STUB: not implemented"; return }

func SkipIfCSIMigrationEnabled() { _ = "STUB: not implemented"; return }

// SkipIfMissingResource skips if the gvr resource is missing.
func SkipIfMissingResource(dynamicClient dynamic.Interface, gvr schema.GroupVersionResource, namespace string) {
	_ = "STUB: not implemented"
	return
}

// not all resources support list, so we ignore those

// SkipUnlessNodeCountIsAtLeast skips if the number of nodes is less than the minNodeCount.
func SkipUnlessNodeCountIsAtLeast(minNodeCount int) { _ = "STUB: not implemented"; return }

// SkipUnlessNodeCountIsAtMost skips if the number of nodes is greater than the maxNodeCount.
func SkipUnlessNodeCountIsAtMost(maxNodeCount int) { _ = "STUB: not implemented"; return }

// SkipIfProviderIs skips if the provider is included in the unsupportedProviders.
func SkipIfProviderIs(unsupportedProviders ...string) { _ = "STUB: not implemented"; return }

// SkipUnlessProviderIs skips if the provider is not included in the supportedProviders.
func SkipUnlessProviderIs(supportedProviders ...string) { _ = "STUB: not implemented"; return }

// SkipUnlessMultizone skips if the cluster does not have multizone.
func SkipUnlessMultizone(c clientset.Interface) { _ = "STUB: not implemented"; return }

// SkipIfMultizone skips if the cluster has multizone.
func SkipIfMultizone(c clientset.Interface) { _ = "STUB: not implemented"; return }

// SkipUnlessMasterOSDistroIs skips if the master OS distro is not included in the supportedMasterOsDistros.
func SkipUnlessMasterOSDistroIs(supportedMasterOsDistros ...string) {
	_ = "STUB: not implemented"
	return
}

// SkipUnlessNodeOSDistroIs skips if the node OS distro is not included in the supportedNodeOsDistros.
func SkipUnlessNodeOSDistroIs(supportedNodeOsDistros ...string) { _ = "STUB: not implemented"; return }

// SkipUnlessNodeOSArchIs skips if the node OS distro is not included in the supportedNodeOsArchs.
func SkipUnlessNodeOSArchIs(supportedNodeOsArchs ...string) { _ = "STUB: not implemented"; return }

// SkipIfNodeOSDistroIs skips if the node OS distro is included in the unsupportedNodeOsDistros.
func SkipIfNodeOSDistroIs(unsupportedNodeOsDistros ...string) { _ = "STUB: not implemented"; return }

// SkipUnlessServerVersionGTE skips if the server version is less than v.
func SkipUnlessServerVersionGTE(v *utilversion.Version, c discovery.ServerVersionInterface) {
	_ = "STUB: not implemented"
	return
}

// SkipUnlessSSHKeyPresent skips if no SSH key is found.
func SkipUnlessSSHKeyPresent() { _ = "STUB: not implemented"; return }

// serverVersionGTE returns true if v is greater than or equal to the server version.
func serverVersionGTE(v *utilversion.Version, c discovery.ServerVersionInterface) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AppArmorDistros are distros with AppArmor support
var AppArmorDistros = []string{"gci", "ubuntu"}

// SkipIfAppArmorNotSupported skips if the AppArmor is not supported by the node OS distro.
func SkipIfAppArmorNotSupported() { _ = "STUB: not implemented"; return }

// RunIfContainerRuntimeIs runs if the container runtime is included in the runtimes.
func RunIfContainerRuntimeIs(runtimes ...string) { _ = "STUB: not implemented"; return }

// RunIfSystemSpecNameIs runs if the system spec name is included in the names.
func RunIfSystemSpecNameIs(names ...string) { _ = "STUB: not implemented"; return }

// SkipUnlessComponentRunsAsPodsAndClientCanDeleteThem run if the component run as pods and client can delete them
func SkipUnlessComponentRunsAsPodsAndClientCanDeleteThem(componentName string, c clientset.Interface, ns string, labelSet labels.Set) {
	_ = "STUB: not implemented"
	// verify if component run as pod
	return
}

// verify if client can delete pod
