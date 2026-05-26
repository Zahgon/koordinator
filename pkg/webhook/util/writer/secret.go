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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/koordinator-sh/koordinator/pkg/webhook/util/generator"
)

const (
	SecretCertWriter = "secret"
)

// secretCertWriter provisions the certificate by reading and writing to the k8s secrets.
type secretCertWriter struct {
	*SecretCertWriterOptions

	// dnsName is the DNS name that the certificate is for.
	dnsName string
}

// SecretCertWriterOptions is options for constructing a secretCertWriter.
type SecretCertWriterOptions struct {
	// client talks to a kubernetes cluster for creating the secret.
	Clientset clientset.Interface
	// certGenerator generates the certificates.
	CertGenerator generator.CertGenerator
	// secret points the secret that contains certificates that written by the CertWriter.
	Secret *types.NamespacedName
}

var _ CertWriter = &secretCertWriter{}

func (ops *SecretCertWriterOptions) setDefaults() { _ = "STUB: not implemented"; return }

func (ops *SecretCertWriterOptions) validate() error { _ = "STUB: not implemented"; return nil }

// NewSecretCertWriter constructs a CertWriter that persists the certificate in a k8s secret.
func NewSecretCertWriter(ops SecretCertWriterOptions) (CertWriter, error) {
	_ = "STUB: not implemented"
	return *new(CertWriter), nil
}

// EnsureCert provisions certificates for a webhookClientConfig by writing the certificates to a k8s secret.
func (s *secretCertWriter) EnsureCert(dnsName string) (*generator.Artifacts, bool, error) {
	_ = "STUB: not implemented"
	// Create or refresh the certs based on clientConfig
	return nil, false, nil
}

var _ certReadWriter = &secretCertWriter{}

func (s *secretCertWriter) buildSecret() (*corev1.Secret, *generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *secretCertWriter) write() (*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *secretCertWriter) overwrite(resourceVersion string) (
	*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *secretCertWriter) read() (*generator.Artifacts, error) {
	_ = "STUB: not implemented"
	//secret := &corev1.Secret{
	//	TypeMeta: metav1.TypeMeta{
	//		APIVersion: "v1",
	//		Kind:       "Secret",
	//	},
	//}
	return nil, nil
}

// Store the CA for next usage.

func secretToCerts(secret *corev1.Secret) *generator.Artifacts {
	_ = "STUB: not implemented"
	return nil
}

func certsToSecret(certs *generator.Artifacts, sec types.NamespacedName) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}
