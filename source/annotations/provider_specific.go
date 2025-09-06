/*
Copyright 2025 The Kubernetes Authors.
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

package annotations

import (
	"fmt"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
)

func ProviderSpecificAnnotations(annotations map[string]string) (endpoint.ProviderSpecific, string) {
	providerSpecificAnnotations := make(endpoint.ProviderSpecific)

	if hasAliasFromAnnotations(annotations) {
		providerSpecificAnnotations.Set("alias", "true")
	}
	setIdentifier := ""
	for k, v := range annotations {
		if k == SetIdentifierKey {
			setIdentifier = v
		} else if attr, ok := strings.CutPrefix(k, AWSPrefix); ok {
			providerSpecificAnnotations.Set(fmt.Sprintf("aws/%s", attr), v)
		} else if attr, ok := strings.CutPrefix(k, SCWPrefix); ok {
			providerSpecificAnnotations.Set(fmt.Sprintf("scw/%s", attr), v)
		} else if attr, ok := strings.CutPrefix(k, WebhookPrefix); ok {
			// Support for wildcard annotations for webhook providers
			providerSpecificAnnotations.Set(fmt.Sprintf("webhook/%s", attr), v)
		} else if attr, ok := strings.CutPrefix(k, CoreDNSPrefix); ok {
			providerSpecificAnnotations.Set(fmt.Sprintf("coredns/%s", attr), v)
		} else if strings.HasPrefix(k, CloudflarePrefix) {
			if strings.Contains(k, CloudflareCustomHostnameKey) {
				providerSpecificAnnotations.Set(CloudflareCustomHostnameKey, v)
			} else if strings.Contains(k, CloudflareProxiedKey) {
				providerSpecificAnnotations.Set(CloudflareProxiedKey, v)
			} else if strings.Contains(k, CloudflareRegionKey) {
				providerSpecificAnnotations.Set(CloudflareRegionKey, v)
			} else if strings.Contains(k, CloudflareRecordCommentKey) {
				providerSpecificAnnotations.Set(CloudflareRecordCommentKey, v)
			}
		}
	}
	return providerSpecificAnnotations, setIdentifier
}
