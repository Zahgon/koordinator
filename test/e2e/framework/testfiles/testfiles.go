/*
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

// Package testfiles provides a wrapper around various optional ways
// of retrieving additional files needed during a test run:
// - builtin bindata
// - filesystem access
//
// Because it is a is self-contained package, it can be used by
// test/e2e/framework and test/e2e/manifest without creating
// a circular dependency.
package testfiles

import (
	"embed"
)

var filesources []FileSource

// AddFileSource registers another provider for files that may be
// needed at runtime. Should be called during initialization of a test
// binary.
func AddFileSource(filesource FileSource) { _ = "STUB: not implemented"; return }

// FileSource implements one way of retrieving test file content.  For
// example, one file source could read from the original source code
// file tree, another from bindata compiled into a test executable.
type FileSource interface {
	// ReadTestFile retrieves the content of a file that gets maintained
	// alongside a test's source code. Files are identified by the
	// relative path inside the repository containing the tests, for
	// example "cluster/gce/upgrade.sh" inside kubernetes/kubernetes.
	//
	// When the file is not found, a nil slice is returned. An error is
	// returned for all fatal errors.
	ReadTestFile(filePath string) ([]byte, error)

	// DescribeFiles returns a multi-line description of which
	// files are available via this source. It is meant to be
	// used as part of the error message when a file cannot be
	// found.
	DescribeFiles() string
}

// Read tries to retrieve the desired file content from
// one of the registered file sources.
func Read(filePath string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Here we try to generate an error that points test authors
// or users in the right direction for resolving the problem.

// Exists checks whether a file could be read. Unexpected errors
// are handled by calling the fail function, which then should
// abort the current test.
func Exists(filePath string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// RootFileSource looks for files relative to a root directory.
type RootFileSource struct {
	Root string
}

// ReadTestFile looks for the file relative to the configured
// root directory. If the path is already absolute, for example
// in a test that has its own method of determining where
// files are, then the path will be used directly.
func (r RootFileSource) ReadTestFile(filePath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Not an error (yet), some other provider may have the file.

// DescribeFiles explains that it looks for files inside a certain
// root directory.
func (r RootFileSource) DescribeFiles() string { _ = "STUB: not implemented"; return "" }

// The default in test_context.go is the relative path
// ../../, which doesn't really help locating the
// actual location. Therefore we add also the absolute
// path if necessary.

// EmbeddedFileSource handles files stored in a package generated with bindata.
type EmbeddedFileSource struct {
	EmbeddedFS embed.FS
	Root       string
	fileList   []string
}

// ReadTestFile looks for an embedded file with the given path.
func (e EmbeddedFileSource) ReadTestFile(filepath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DescribeFiles explains that it is looking inside an embedded filesystem
func (e EmbeddedFileSource) DescribeFiles() string { _ = "STUB: not implemented"; return "" }

func (e *EmbeddedFileSource) populateFileList() { _ = "STUB: not implemented"; return }
