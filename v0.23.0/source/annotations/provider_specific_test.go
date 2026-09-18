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
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
)

// Spelled out rather than derived from CloudflarePrefix so a change to the
// prefix handling shows up as a test failure.
const (
	cloudflareProxiedAnnotation        = DefaultAnnotationPrefix + "cloudflare-proxied"
	cloudflareCustomHostnameAnnotation = DefaultAnnotationPrefix + "cloudflare-custom-hostname"
	cloudflareRegionAnnotation         = DefaultAnnotationPrefix + "cloudflare-region-key"
	cloudflareRecordCommentAnnotation  = DefaultAnnotationPrefix + "cloudflare-record-comment"
	cloudflareTagsAnnotation           = DefaultAnnotationPrefix + "cloudflare-tags"
)

func TestMain(m *testing.M) {
	// Initialize annotation prefixes before running tests
	SetAnnotationPrefix(DefaultAnnotationPrefix)
	os.Exit(m.Run())
}

func TestProviderSpecificAnnotations(t *testing.T) {
	tests := []struct {
		name          string
		annotations   map[string]string
		expected      endpoint.ProviderSpecific
		setIdentifier string
	}{
		{
			name:          "no annotations",
			annotations:   map[string]string{},
			expected:      endpoint.ProviderSpecific{},
			setIdentifier: "",
		},
		{
			name: "Cloudflare proxied annotation",
			annotations: map[string]string{
				cloudflareProxiedAnnotation: "true",
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareProxiedProperty, Value: "true"},
			},
			setIdentifier: "",
		},
		{
			name: "Cloudflare custom hostname annotation",
			annotations: map[string]string{
				cloudflareCustomHostnameAnnotation: "custom.example.com",
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareCustomHostnameProperty, Value: "custom.example.com"},
			},
			setIdentifier: "",
		},
		{
			name: "AWS annotation",
			annotations: map[string]string{
				"external-dns.kubernetes.io/aws-weight": "100",
			},
			expected: endpoint.ProviderSpecific{
				{Name: "aws/weight", Value: "100"},
			},
			setIdentifier: "",
		},
		{
			name: "CoreDNS annotation",
			annotations: map[string]string{
				"external-dns.kubernetes.io/coredns-group": "g1",
			},
			expected: endpoint.ProviderSpecific{
				{Name: "coredns/group", Value: "g1"},
			},
			setIdentifier: "",
		},
		{
			name: "Azure tags annotation",
			annotations: map[string]string{
				AzureTagsKey: "cost-center=12345,owner=backend-team",
			},
			expected: endpoint.ProviderSpecific{
				{Name: "azure/tags", Value: "cost-center=12345,owner=backend-team"},
			},
			setIdentifier: "",
		},
		{
			name: "Azure tags annotation with spaces",
			annotations: map[string]string{
				AzureTagsKey: "environment=production, app=myapp ",
			},
			expected: endpoint.ProviderSpecific{
				{Name: "azure/tags", Value: "environment=production, app=myapp "},
			},
			setIdentifier: "",
		},
		{
			name: "Set identifier annotation",
			annotations: map[string]string{
				SetIdentifierKey: "identifier",
			},
			expected:      endpoint.ProviderSpecific{},
			setIdentifier: "identifier",
		},
		{
			name: "Record type annotation",
			annotations: map[string]string{
				RecordTypeKey: "ptr",
			},
			expected: endpoint.ProviderSpecific{
				{Name: endpoint.ProviderSpecificRecordType, Value: "ptr"},
			},
			setIdentifier: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, setIdentifier := ProviderSpecificAnnotations(tt.annotations)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.setIdentifier, setIdentifier)

			for _, prop := range result {
				slashIdx := strings.Index(prop.Name, "/")
				if slashIdx == -1 {
					continue
				}
				assert.NotContains(t, prop.Name[:slashIdx], ".",
					"property %q uses a full annotation name; use the short \"provider/attr\" form instead", prop.Name)
			}
		})
	}
}

func TestGetProviderSpecificCloudflareAnnotations(t *testing.T) {

	for _, tc := range []struct {
		title         string
		annotations   map[string]string
		expectedKey   string
		expectedValue string
	}{
		{
			title:         "Cloudflare tags annotation is set correctly",
			annotations:   map[string]string{cloudflareTagsAnnotation: "env:test,owner:team-a"},
			expectedKey:   CloudflareTagsProperty,
			expectedValue: "env:test,owner:team-a",
		},
		{
			title: "Cloudflare tags annotation among another annotations is set correctly",
			annotations: map[string]string{
				"random annotation 1":    "random value 1",
				cloudflareTagsAnnotation: "env:test,owner:team-b",
				"random annotation 2":    "random value 2"},
			expectedKey:   CloudflareTagsProperty,
			expectedValue: "env:test,owner:team-b",
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := ProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == tc.expectedKey {
					assert.Equal(t, tc.expectedValue, providerSpecificAnnotation.Value)
					return
				}
			}
			t.Errorf("Cloudflare provider specific annotation %s is not set correctly to %s", tc.expectedKey, tc.expectedValue)
		})
	}

	for _, tc := range []struct {
		title         string
		annotations   map[string]string
		expectedKey   string
		expectedValue bool
	}{
		{
			title:         "Cloudflare proxied annotation is set correctly to true",
			annotations:   map[string]string{cloudflareProxiedAnnotation: "true"},
			expectedKey:   CloudflareProxiedProperty,
			expectedValue: true,
		},
		{
			title:         "Cloudflare proxied annotation is set correctly to false",
			annotations:   map[string]string{cloudflareProxiedAnnotation: "false"},
			expectedKey:   CloudflareProxiedProperty,
			expectedValue: false,
		},
		{
			title: "Cloudflare proxied annotation among another annotations is set correctly to true",
			annotations: map[string]string{
				"random annotation 1":       "random value 1",
				cloudflareProxiedAnnotation: "false",
				"random annotation 2":       "random value 2",
			},
			expectedKey:   CloudflareProxiedProperty,
			expectedValue: false,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := ProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == tc.expectedKey {
					assert.Equal(t, strconv.FormatBool(tc.expectedValue), providerSpecificAnnotation.Value)
					return
				}
			}
			t.Errorf("Cloudflare provider specific annotation %s is not set correctly to %v", tc.expectedKey, tc.expectedValue)
		})
	}

	for _, tc := range []struct {
		title         string
		annotations   map[string]string
		expectedKey   string
		expectedValue string
	}{
		{
			title:         "Cloudflare region key annotation is set correctly",
			annotations:   map[string]string{cloudflareRegionAnnotation: "us"},
			expectedKey:   CloudflareRegionProperty,
			expectedValue: "us",
		},
		{
			title: "Cloudflare region key annotation among another annotations is set correctly",
			annotations: map[string]string{
				"random annotation 1":      "random value 1",
				cloudflareRegionAnnotation: "us",
				"random annotation 2":      "random value 2",
			},
			expectedKey:   CloudflareRegionProperty,
			expectedValue: "us",
		},
		{
			title: "Cloudflare DNS record comment annotation is set correctly",
			annotations: map[string]string{
				cloudflareRecordCommentAnnotation: "comment",
			},
			expectedKey:   CloudflareRecordCommentProperty,
			expectedValue: "comment",
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := ProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == tc.expectedKey {
					assert.Equal(t, tc.expectedValue, providerSpecificAnnotation.Value)
					return
				}
			}
			t.Errorf("Cloudflare provider specific annotation %s is not set correctly to %v", tc.expectedKey, tc.expectedValue)
		})
	}

	for _, tc := range []struct {
		title         string
		annotations   map[string]string
		expectedKey   string
		expectedValue string
	}{
		{
			title:         "Cloudflare custom hostname annotation is set correctly",
			annotations:   map[string]string{cloudflareCustomHostnameAnnotation: "a.foo.fancybar.com"},
			expectedKey:   CloudflareCustomHostnameProperty,
			expectedValue: "a.foo.fancybar.com",
		},
		{
			title: "Cloudflare custom hostname annotation among another annotations is set correctly",
			annotations: map[string]string{
				"random annotation 1":              "random value 1",
				cloudflareCustomHostnameAnnotation: "a.foo.fancybar.com",
				"random annotation 2":              "random value 2"},
			expectedKey:   CloudflareCustomHostnameProperty,
			expectedValue: "a.foo.fancybar.com",
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := ProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == tc.expectedKey {
					assert.Equal(t, tc.expectedValue, providerSpecificAnnotation.Value)
					return
				}
			}
			t.Errorf("Cloudflare provider specific annotation %s is not set correctly to %s", tc.expectedKey, tc.expectedValue)
		})
	}
}

func TestGetProviderSpecificAliasAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title         string
		annotations   map[string]string
		expectedKey   string
		expectedValue bool
	}{
		{
			title:         "alias annotation is set correctly to true",
			annotations:   map[string]string{AliasKey: "true"},
			expectedKey:   AliasKey,
			expectedValue: true,
		},
		{
			title: "alias annotation among another annotations is set correctly to true",
			annotations: map[string]string{
				"random annotation 1": "random value 1",
				AliasKey:              "true",
				"random annotation 2": "random value 2",
			},
			expectedKey:   AliasKey,
			expectedValue: true,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := ProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == endpoint.ProviderSpecificAlias {
					assert.Equal(t, strconv.FormatBool(tc.expectedValue), providerSpecificAnnotation.Value)
					return
				}
			}
			t.Errorf("provider specific annotation alias is not set correctly to %v", tc.expectedValue)
		})
	}

	for _, tc := range []struct {
		title       string
		annotations map[string]string
	}{
		{
			title:       "alias annotation is set to false",
			annotations: map[string]string{AliasKey: "false"},
		},
		{
			title: "alias annotation is not set",
			annotations: map[string]string{
				"random annotation 1": "random value 1",
				"random annotation 2": "random value 2",
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := ProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == endpoint.ProviderSpecificAlias {
					t.Error("provider specific annotation alias is not expected to be set")
				}
			}

		})
	}
}

// Property names must be the short "provider/attr" form (e.g. "aws/weight",
// "cloudflare/proxied") so they do not depend on --annotation-prefix. Catches a
// new provider that emits the full annotation name instead.
func TestProviderSpecificPropertyNameConvention(t *testing.T) {
	annotations := map[string]string{
		AnnotationKeyPrefix + "aws-weight":        "10",
		AnnotationKeyPrefix + "scw-something":     "val",
		AnnotationKeyPrefix + "webhook-something": "val",
		AnnotationKeyPrefix + "coredns-group":     "g1",
		cloudflareProxiedAnnotation:               "true",
		cloudflareTagsAnnotation:                  "tag1",
		cloudflareRegionAnnotation:                "us",
		cloudflareRecordCommentAnnotation:         "comment",
		cloudflareCustomHostnameAnnotation:        "host.example.com",
		AliasKey:                                  "true",
	}

	props, _ := ProviderSpecificAnnotations(annotations)
	for _, prop := range props {
		name := prop.Name
		providerSegment, _, ok := strings.Cut(name, "/")
		if !ok {
			// No slash: provider-agnostic property (e.g. "alias") — always OK.
			continue
		}
		// Every provider must use the short "provider/attr" form.
		// The segment before "/" must be a plain word with no dots.
		assert.NotContains(t, providerSegment, ".",
			"property %q uses a full annotation name; use the short \"provider/attr\" form instead", name)
	}
}

func TestGetProviderSpecificIdentifierAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title              string
		annotations        map[string]string
		expectedResult     map[string]string
		expectedIdentifier string
	}{
		{
			title: "aws- provider specific annotations are set correctly",
			annotations: map[string]string{
				"external-dns.kubernetes.io/aws-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.kubernetes.io/aws-annotation-2": "value 2",
			},
			expectedResult: map[string]string{
				"aws/annotation-1": "value 1",
				"aws/annotation-2": "value 2",
			},
			expectedIdentifier: "id1",
		},
		{
			title: "scw- provider specific annotations are set correctly",
			annotations: map[string]string{
				"external-dns.kubernetes.io/scw-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.kubernetes.io/scw-annotation-2": "value 2",
			},
			expectedResult: map[string]string{
				"scw/annotation-1": "value 1",
				"scw/annotation-2": "value 2",
			},
			expectedIdentifier: "id1",
		},
		{
			title: "webhook- provider specific annotations are set correctly",
			annotations: map[string]string{
				"external-dns.kubernetes.io/webhook-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.kubernetes.io/webhook-annotation-2": "value 2",
			},
			expectedResult: map[string]string{
				"webhook/annotation-1": "value 1",
				"webhook/annotation-2": "value 2",
			},
			expectedIdentifier: "id1",
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, identifier := ProviderSpecificAnnotations(tc.annotations)
			assert.Equal(t, tc.expectedIdentifier, identifier)
			for expectedAnnotationKey, expectedAnnotationValue := range tc.expectedResult {
				expectedResultFound := false
				for _, providerSpecificAnnotation := range providerSpecificAnnotations {
					if providerSpecificAnnotation.Name == expectedAnnotationKey {
						assert.Equal(t, expectedAnnotationValue, providerSpecificAnnotation.Value)
						expectedResultFound = true
						break
					}
				}
				if !expectedResultFound {
					t.Errorf("provider specific annotation %s has not been set", expectedAnnotationKey)
				}
			}
		})
	}
}

func TestLegacyProviderSpecificName(t *testing.T) {
	for _, tc := range []struct {
		name     string
		expected string
		legacy   bool
	}{
		{name: CloudflareProxiedProperty, expected: CloudflareProxiedProperty},
		{name: "external-dns.kubernetes.io/cloudflare-proxied", expected: CloudflareProxiedProperty, legacy: true},
		{name: "external-dns.alpha.kubernetes.io/cloudflare-proxied", expected: CloudflareProxiedProperty, legacy: true},
		{name: "custom.io/cloudflare-proxied", expected: CloudflareProxiedProperty, legacy: true},
		{name: "cloudflare-proxied", expected: CloudflareProxiedProperty, legacy: true},
		{name: "external-dns.kubernetes.io/cloudflare-region-key", expected: CloudflareRegionProperty, legacy: true},
		{name: "external-dns.kubernetes.io/cloudflare-unknown", expected: "external-dns.kubernetes.io/cloudflare-unknown"},
		{name: "aws/evaluate-target-health", expected: "aws/evaluate-target-health"},
		{name: endpoint.ProviderSpecificAlias, expected: endpoint.ProviderSpecificAlias},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, legacy := LegacyProviderSpecificName(tc.name)
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.legacy, legacy)
		})
	}
}

func TestNormalizeProviderSpecific(t *testing.T) {
	for _, tc := range []struct {
		title    string
		given    endpoint.ProviderSpecific
		expected endpoint.ProviderSpecific
	}{
		{
			title:    "no properties",
			given:    nil,
			expected: nil,
		},
		{
			title: "legacy names are rewritten, others untouched",
			given: endpoint.ProviderSpecific{
				{Name: "external-dns.alpha.kubernetes.io/cloudflare-proxied", Value: "true"},
				{Name: "aws/evaluate-target-health", Value: "true"},
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareProxiedProperty, Value: "true"},
				{Name: "aws/evaluate-target-health", Value: "true"},
			},
		},
		{
			title: "canonical wins over a colliding legacy name",
			given: endpoint.ProviderSpecific{
				{Name: "external-dns.kubernetes.io/cloudflare-proxied", Value: "false"},
				{Name: CloudflareProxiedProperty, Value: "true"},
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareProxiedProperty, Value: "true"},
			},
		},
		{
			title: "duplicate legacy names collapse to one",
			given: endpoint.ProviderSpecific{
				{Name: "external-dns.kubernetes.io/cloudflare-tags", Value: "tag1"},
				{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag2"},
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareTagsProperty, Value: "tag1"},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			ep := &endpoint.Endpoint{DNSName: "foo.example.org", ProviderSpecific: tc.given}
			NormalizeProviderSpecific(ep)
			assert.Equal(t, tc.expected, ep.ProviderSpecific)
		})
	}
}
