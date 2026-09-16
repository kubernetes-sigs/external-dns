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
	"sync"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/sets"
)

// Canonical Cloudflare property names. Deliberately independent of
// --annotation-prefix: they are part of the DNSEndpoint API, so a prefixed name
// would tie a manifest to the controller's flags.
const (
	CloudflareProxiedProperty        = "cloudflare/proxied"
	CloudflareCustomHostnameProperty = "cloudflare/custom-hostname"
	CloudflareRegionProperty         = "cloudflare/region-key"
	CloudflareRecordCommentProperty  = "cloudflare/record-comment"
	CloudflareTagsProperty           = "cloudflare/tags"
)

// Keyed by the attribute part of "<prefix>cloudflare-<attribute>".
var cloudflareProperties = map[string]string{
	"proxied":         CloudflareProxiedProperty,
	"custom-hostname": CloudflareCustomHostnameProperty,
	"region-key":      CloudflareRegionProperty,
	"record-comment":  CloudflareRecordCommentProperty,
	"tags":            CloudflareTagsProperty,
}

// Warn once per distinct name: NormalizeProviderSpecific runs on every endpoint
// of every sync.
var warnedLegacyNames sync.Map

func ProviderSpecificAnnotations(annotations map[string]string) (endpoint.ProviderSpecific, string) {
	providerSpecificAnnotations := endpoint.ProviderSpecific{}

	if hasAliasFromAnnotations(annotations) {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  endpoint.ProviderSpecificAlias,
			Value: "true",
		})
	}
	if v, ok := annotations[RecordTypeKey]; ok {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  endpoint.ProviderSpecificRecordType,
			Value: v,
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
		} else if k == AzureTagsKey {
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  "azure/tags",
				Value: v,
			})
		} else if attr, ok := strings.CutPrefix(k, CloudflarePrefix); ok {
			if name, known := cloudflareProperties[attr]; known {
				providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
					Name:  name,
					Value: v,
				})
			}
		}
	}
	return providerSpecificAnnotations, setIdentifier
}

// LegacyProviderSpecificName maps a property name written in the old Cloudflare
// annotation form onto its canonical name, reporting whether it rewrote
// anything.
// Any prefix is accepted, not just the current --annotation-prefix, for easier migration
func LegacyProviderSpecificName(name string) (string, bool) {
	attr := name
	if i := strings.LastIndex(name, "/"); i >= 0 {
		attr = name[i+1:]
	}
	attr, ok := strings.CutPrefix(attr, "cloudflare-")
	if !ok {
		return name, false
	}
	canonical, known := cloudflareProperties[attr]
	if !known || canonical == name {
		return name, false
	}
	return canonical, true
}

// NormalizeProviderSpecific rewrites ep's legacy property names to their
// canonical form, so sources, the plan and providers only ever compare
// canonical names.
// A canonical property already present wins over the legacy one it collides with.
func NormalizeProviderSpecific(ep *endpoint.Endpoint) {
	if len(ep.ProviderSpecific) == 0 {
		return
	}

	canonical := sets.New[string]()
	for _, prop := range ep.ProviderSpecific {
		if _, legacy := LegacyProviderSpecificName(prop.Name); !legacy {
			canonical.Insert(prop.Name)
		}
	}

	normalized := make(endpoint.ProviderSpecific, 0, len(ep.ProviderSpecific))
	for _, prop := range ep.ProviderSpecific {
		name, legacy := LegacyProviderSpecificName(prop.Name)
		if !legacy {
			normalized = append(normalized, prop)
			continue
		}
		if canonical.Has(name) {
			log.Debugf("%s: ignoring provider-specific property %q because %q is already set", ep.DNSName, prop.Name, name)
			continue
		}
		warnLegacyProviderSpecificName(prop.Name, name)
		prop.Name = name
		canonical.Insert(name)
		normalized = append(normalized, prop)
	}
	ep.ProviderSpecific = normalized
}

func warnLegacyProviderSpecificName(legacy, canonical string) {
	if _, seen := warnedLegacyNames.LoadOrStore(legacy, struct{}{}); seen {
		return
	}
	log.Warnf("Provider-specific property %q is deprecated and will be removed in a future release; use %q instead", legacy, canonical)
}
