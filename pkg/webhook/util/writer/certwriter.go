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
)

const (
	// CAKeyName is the name of the CA private key
	CAKeyName = "ca-key.pem"
	// CACertName is the name of the CA certificate
	CACertName = "ca-cert.pem"
	// ServerKeyName is the name of the server private key
	ServerKeyName  = "key.pem"
	ServerKeyName2 = "tls.key"
	// ServerCertName is the name of the serving certificate
	ServerCertName  = "cert.pem"
	ServerCertName2 = "tls.crt"
)

// CertWriter provides method to handle webhooks.
type CertWriter interface {
	// EnsureCert provisions the cert for the webhookClientConfig.
	EnsureCert(dnsName string) (*generator.Artifacts, bool, error)
}

// handleCommon ensures the given webhook has a proper certificate.
// It uses the given certReadWriter to read and (or) write the certificate.
func handleCommon(dnsName string, ch certReadWriter) (*generator.Artifacts, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Recreate the cert if it's invalid.

func createIfNotExists(ch certReadWriter) (*generator.Artifacts, bool, error) {
	_ = "STUB: not implemented"
	// Try to read first
	return nil, false, nil
}

// Create if not exists

// This may happen if there is another racer.

// certReadWriter provides methods for reading and writing certificates.
type certReadWriter interface {
	// read reads a webhook name and returns the certs for it.
	read() (*generator.Artifacts, error)
	// write writes the certs and return the certs it wrote.
	write() (*generator.Artifacts, error)
	// overwrite overwrites the existing certs and return the certs it wrote.
	overwrite(resourceVersion string) (*generator.Artifacts, error)
}

func validCert(certs *generator.Artifacts, dnsName string) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify key and cert are valid pair

// Verify cert is good for desired DNS name and signed by CA and will be valid for desired period of time.
