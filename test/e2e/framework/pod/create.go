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

package pod

import (
	"time"

	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"

	imageutils "github.com/koordinator-sh/koordinator/test/utils/image"
)

var (
	// BusyBoxImage is the image URI of BusyBox.
	BusyBoxImage = imageutils.GetE2EImage(imageutils.BusyBox)
)

// Config is a struct containing all arguments for creating a pod.
// SELinux testing requires to pass HostIPC and HostPID as boolean arguments.
type Config struct {
	NS                     string
	PVCs                   []*v1.PersistentVolumeClaim
	PVCsReadOnly           bool
	InlineVolumeSources    []*v1.VolumeSource
	IsPrivileged           bool
	Command                string
	HostIPC                bool
	HostPID                bool
	SeLinuxLabel           *v1.SELinuxOptions
	FsGroup                *int64
	NodeSelection          NodeSelection
	ImageID                int
	PodFSGroupChangePolicy *v1.PodFSGroupChangePolicy
}

// CreateUnschedulablePod with given claims based on node selector
func CreateUnschedulablePod(client clientset.Interface, namespace string, nodeSelector map[string]string, pvclaims []*v1.PersistentVolumeClaim, isPrivileged bool, command string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Waiting for pod to become Unschedulable

// get fresh pod info

// CreateClientPod defines and creates a pod with a mounted PV. Pod runs infinite loop until killed.
func CreateClientPod(c clientset.Interface, ns string, pvc *v1.PersistentVolumeClaim) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePod with given claims based on node selector
func CreatePod(client clientset.Interface, namespace string, nodeSelector map[string]string, pvclaims []*v1.PersistentVolumeClaim, isPrivileged bool, command string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Waiting for pod to be running

// get fresh pod info

// CreateSecPod creates security pod with given claims
func CreateSecPod(client clientset.Interface, podConfig *Config, timeout time.Duration) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecPodWithNodeSelection creates security pod with given claims
func CreateSecPodWithNodeSelection(client clientset.Interface, podConfig *Config, timeout time.Duration) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Waiting for pod to be running

// get fresh pod info

// MakePod returns a pod definition based on the namespace. The pod references the PVC's
// name.  A slice of BASH commands can be supplied as args to be run by the pod
func MakePod(ns string, nodeSelector map[string]string, pvclaims []*v1.PersistentVolumeClaim, isPrivileged bool, command string) *v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

/*inline volume sources*/ /*PVCs readonly*/

// MakeSecPod returns a pod definition based on the namespace. The pod references the PVC's
// name.  A slice of BASH commands can be supplied as args to be run by the pod.
func MakeSecPod(podConfig *Config) (*v1.Pod, error) { _ = "STUB: not implemented"; return nil, nil }

// MakePodSpec returns a PodSpec definition
func MakePodSpec(podConfig *Config) *v1.PodSpec { _ = "STUB: not implemented"; return nil }

func setVolumes(podSpec *v1.PodSpec, pvcs []*v1.PersistentVolumeClaim, inlineVolumeSources []*v1.VolumeSource, pvcsReadOnly bool) {
	_ = "STUB: not implemented"
	return
}

// In-line volumes can be only filesystem, not block.
