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

package writer

import (
	"github.com/koordinator-sh/koordinator/pkg/webhook/util/generator"
	"github.com/koordinator-sh/koordinator/pkg/webhook/util/writer/atomic"
)

const (
	FsCertWriter = "fs"
)

// fsCertWriter provisions the certificate by reading and writing to the filesystem.
type fsCertWriter struct {
	// dnsName is the DNS name that the certificate is for.
	dnsName string

	*FSCertWriterOptions
}

// FSCertWriterOptions are options for constructing a FSCertWriter.
type FSCertWriterOptions struct {
	// certGenerator generates the certificates.
	CertGenerator generator.CertGenerator
	// path is the directory that the certificate and private key and CA certificate will be written.
	Path string
}

var _ CertWriter = &fsCertWriter{}

func (ops *FSCertWriterOptions) setDefaults() { _ = "STUB: not implemented"; return }

func (ops *FSCertWriterOptions) validate() error { _ = "STUB: not implemented"; return nil }

// NewFSCertWriter constructs a CertWriter that persists the certificate on filesystem.
func NewFSCertWriter(ops FSCertWriterOptions) (CertWriter, error) {
	_ = "STUB: not implemented"
	return *new(CertWriter), nil
}

// EnsureCert provisions certificates for a webhookClientConfig by writing the certificates in the filesystem.
func (f *fsCertWriter) EnsureCert(dnsName string) (*generator.Artifacts, bool, error) {
	_ = "STUB: not implemented"
	// create or refresh cert and write it to fs
	return nil, false, nil
}

func (f *fsCertWriter) write() (*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fsCertWriter) overwrite(_ string) (*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fsCertWriter) doWrite() (*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteCertsToDir(path string, certs *generator.Artifacts) error {
	_ = "STUB: not implemented"
	// Writer's algorithm only manages files using symbolic link.
	// If a file is not a symbolic link, will ignore the update for it.
	// We want to cleanup for Writer by removing old files that are not symbolic links.
	return nil
}

// prepareToWrite ensures it directory is compatible with the atomic.Writer library.
func prepareToWrite(dir string) error { _ = "STUB: not implemented"; return nil }

// TODO: figure out if we can reduce the permission. (Now it's 0777)

// if it's not a symbolic link

func (f *fsCertWriter) read() (*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ensureExist(dir string) error { _ = "STUB: not implemented"; return nil }

func certToProjectionMap(cert *generator.Artifacts) map[string]atomic.FileProjection {
	_ = "STUB: not implemented"
	// TODO: figure out if we can reduce the permission. (Now it's 0666)
	return nil
}
