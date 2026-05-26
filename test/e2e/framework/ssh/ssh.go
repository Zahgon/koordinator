/*
Copyright 2022 The Koordinator Authors.
Copyright 2018 The Kubernetes Authors.

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

package ssh

import (
	"time"

	"golang.org/x/crypto/ssh"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	// SSHPort is tcp port number of SSH
	SSHPort = "22"

	// pollNodeInterval is how often to Poll pods.
	pollNodeInterval = 2 * time.Second

	// singleCallTimeout is how long to try single API calls (like 'get' or 'list'). Used to prevent
	// transient failures from failing tests.
	singleCallTimeout = 5 * time.Minute

	// sshBastionEnvKey is the environment variable key for running SSH commands via bastion.
	sshBastionEnvKey = "KUBE_SSH_BASTION"
)

// GetSigner returns an ssh.Signer for the provider ("gce", etc.) that can be
// used to SSH to their nodes.
func GetSigner(provider string) (ssh.Signer, error) {
	_ = "STUB: not implemented"
	// honor a consistent SSH key across all providers
	return *new(ssh.Signer), nil
}

// Select the key itself to use. When implementing more providers here,
// please also add them to any SSH tests that are disabled because of signer
// support.

// Respect absolute paths for keys given by user, fallback to assuming
// relative paths are in ~/.ssh

func makePrivateKeySignerFromFile(key string) (ssh.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ssh.Signer), nil
}

// NodeSSHHosts returns SSH-able host names for all schedulable nodes.
// If it can't find any external IPs, it falls back to
// looking for internal IPs. If it can't find an internal IP for every node it
// returns an error, though it still returns all hosts that it found in that
// case.
func NodeSSHHosts(c clientset.Interface) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If  ExternalIPs aren't available for all nodes, try falling back to the InternalIPs.

// Error if neither External nor Internal IPs weren't available for all nodes.

// canConnect returns true if a network connection is possible to the SSHPort.
func canConnect(host string) bool { _ = "STUB: not implemented"; return false }

// Result holds the execution result of SSH command
type Result struct {
	User   string
	Host   string
	Cmd    string
	Stdout string
	Stderr string
	Code   int
}

// NodeExec execs the given cmd on node via SSH. Note that the nodeName is an sshable name,
// eg: the name returned by framework.GetMasterHost(). This is also not guaranteed to work across
// cloud providers since it involves ssh.
func NodeExec(nodeName, cmd, provider string) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

// SSH synchronously SSHs to a node running on provider and runs cmd. If there
// is no error performing the SSH, the stdout, stderr, and exit code are
// returned.
func SSH(cmd, host, provider string) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

// Get a signer for the provider.

// RunSSHCommand will default to Getenv("USER") if user == "", but we're
// defaulting here as well for logging clarity.

// runSSHCommandViaBastion returns the stdout, stderr, and exit code from running cmd on
// host as specific user, along with any SSH-level error.
func runSSHCommand(cmd, user, host string, signer ssh.Signer) (string, string, int, error) {
	_ = "STUB: not implemented"
	return "", "", 0, nil
}

// Setup the config, dial the server, and open a session.

// Run the command.

// Check whether the command failed to run or didn't complete.

// If we got an ExitError and the exit code is nonzero, we'll
// consider the SSH itself successful (just that the command run
// errored on the host).

// Some other kind of error happened (e.g. an IOError); consider the
// SSH unsuccessful.

// runSSHCommandViaBastion returns the stdout, stderr, and exit code from running cmd on
// host as specific user, along with any SSH-level error. It uses an SSH proxy to connect
// to bastion, then via that tunnel connects to the remote host. Similar to
// sshutil.RunSSHCommand but scoped to the needs of the test infrastructure.
func runSSHCommandViaBastion(cmd, user, bastion, host string, signer ssh.Signer) (string, string, int, error) {
	_ = "STUB: not implemented"
	// Setup the config, dial the server, and open a session.
	return "", "", 0, nil
}

// Run the command.

// Check whether the command failed to run or didn't complete.

// If we got an ExitError and the exit code is nonzero, we'll
// consider the SSH itself successful (just that the command run
// errored on the host).

// Some other kind of error happened (e.g. an IOError); consider the
// SSH unsuccessful.

// LogResult records result log
func LogResult(result Result) { _ = "STUB: not implemented"; return }

// IssueSSHCommandWithResult tries to execute a SSH command and returns the execution result
func IssueSSHCommandWithResult(cmd, provider string, node *v1.Node) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No external IPs were found, let's try to use internal as plan B

// IssueSSHCommand tries to execute a SSH command
func IssueSSHCommand(cmd, provider string, node *v1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// nodeAddresses returns the first address of the given type of each node.
func nodeAddresses(nodelist *v1.NodeList, addrType v1.NodeAddressType) []string {
	_ = "STUB: not implemented"
	return nil
}

// waitListSchedulableNodes is a wrapper around listing nodes supporting retries.
func waitListSchedulableNodes(c clientset.Interface) (*v1.NodeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// waitListSchedulableNodesOrDie is a wrapper around listing nodes supporting retries.
func waitListSchedulableNodesOrDie(c clientset.Interface) *v1.NodeList {
	_ = "STUB: not implemented"
	return nil
}

// expectNoError checks if "err" is set, and if so, fails assertion while logging the error.
func expectNoError(err error, explain ...interface{}) { _ = "STUB: not implemented"; return }

// expectNoErrorWithOffset checks if "err" is set, and if so, fails assertion while logging the error at "offset" levels above its caller
// (for example, for call chain f -> g -> ExpectNoErrorWithOffset(1, ...) error would be logged for "f").
func expectNoErrorWithOffset(offset int, err error, explain ...interface{}) {
	_ = "STUB: not implemented"
	return
}
