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

package generator

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"

	"k8s.io/client-go/util/cert"
)

const (
	rsaKeySize = 2048
)

// ServiceToCommonName generates the CommonName for the certificate when using a k8s service.
func ServiceToCommonName(serviceNamespace, serviceName string) string {
	_ = "STUB: not implemented"
	return ""
}

// SelfSignedCertGenerator implements the certGenerator interface.
// It provisions self-signed certificates.
type SelfSignedCertGenerator struct {
	caKey  []byte
	caCert []byte
}

var _ CertGenerator = &SelfSignedCertGenerator{}

// SetCA sets the PEM-encoded CA private key and CA cert for signing the generated serving cert.
func (cp *SelfSignedCertGenerator) SetCA(caKey, caCert []byte) { _ = "STUB: not implemented"; return }

// Generate creates and returns a CA certificate, certificate and
// key for the server. serverKey and serverCert are used by the server
// to establish trust for clients, CA certificate is used by the
// client to verify the server authentication chain.
// The cert will be valid for 365 days.
func (cp *SelfSignedCertGenerator) Generate(commonName string) (*Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cp *SelfSignedCertGenerator) validCACert() (bool, *rsa.PrivateKey, *x509.Certificate) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// NewPrivateKey creates an RSA private key
func NewPrivateKey() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

// NewSignedCert creates a signed certificate using the given CA certificate and key
func NewSignedCert(cfg cert.Config, key crypto.Signer, caCert *x509.Certificate, caKey crypto.Signer) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodePrivateKeyPEM returns PEM-encoded private key data
func EncodePrivateKeyPEM(key *rsa.PrivateKey) []byte { _ = "STUB: not implemented"; return nil }

// EncodeCertPEM returns PEM-encoded certificate data
func EncodeCertPEM(ct *x509.Certificate) []byte { _ = "STUB: not implemented"; return nil }
