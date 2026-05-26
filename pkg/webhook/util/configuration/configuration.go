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

package configuration

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	clientset "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	webhookutil "github.com/koordinator-sh/koordinator/pkg/webhook/util"
)

var (
	mutatingWebhookConfigurationName   = webhookutil.GetMutatingWebhookName()
	validatingWebhookConfigurationName = webhookutil.GetValidatingWebhookName()
)

func Ensure(kubeClient clientset.Interface, handlers map[string]admission.Handler, caBundle []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func getPath(clientConfig *admissionregistrationv1.WebhookClientConfig) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func convertClientConfig(clientConfig *admissionregistrationv1.WebhookClientConfig, host string, port int) {
	_ = "STUB: not implemented"
	return
}

func parseMutatingTemplate(mutatingConfig *admissionregistrationv1.MutatingWebhookConfiguration) ([]admissionregistrationv1.MutatingWebhook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseValidatingTemplate(validatingConfig *admissionregistrationv1.ValidatingWebhookConfiguration) ([]admissionregistrationv1.ValidatingWebhook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
