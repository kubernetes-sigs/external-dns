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
	providerSpecificAnnotations := endpoint.ProviderSpecific{}

	if hasAliasFromAnnotations(annotations) {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  "alias",
			Value: "true",
		})
	}
	setIdentifier := ""
	for k, v := range annotations {
		if k == SetIdentifierKey {
			setIdentifier = v
		} else if attr, ok := strings.CutPrefix(k, AWSPrefix); ok {
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("aws/%s", attr),
				Value: v,
			})
		} else if attr, ok := strings.CutPrefix(k, SCWPrefix); ok {
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("scw/%s", attr),
				Value: v,
			})
		} else if attr, ok := strings.CutPrefix(k, WebhookPrefix); ok {
			// Support for wildcard annotations for webhook providers
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("webhook/%s", attr),
				Value: v,
			})
		} else if attr, ok := strings.CutPrefix(k, CoreDNSPrefix); ok {
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("coredns/%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, CloudflarePrefix) {
			// TODO: unlike other providers which normalise to "provider/attr",
			// Cloudflare retains the full annotation key as the property name
			// (e.g. "external-dns.alpha.kubernetes.io/cloudflare-proxied").
			// This is why RetainProviderProperties has a special case for cloudflare.
			// Should be aligned with the standard convention in a future change.
			switch {
			case strings.Contains(k, CloudflareCustomHostnameKey):
				providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
					Name:  CloudflareCustomHostnameKey,
					Value: v,
				})
			case strings.Contains(k, CloudflareProxiedKey):
				providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
					Name:  CloudflareProxiedKey,
					Value: v,
				})
			case strings.Contains(k, CloudflareRegionKey):
				providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
					Name:  CloudflareRegionKey,
					Value: v,
				})
			case strings.Contains(k, CloudflareRecordCommentKey):
				providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
					Name:  CloudflareRecordCommentKey,
					Value: v,
				})
			case strings.Contains(k, CloudflareTagsKey):
				providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
					Name:  CloudflareTagsKey,
					Value: v,
				})
			}
		}
	}
	return providerSpecificAnnotations, setIdentifier
}
