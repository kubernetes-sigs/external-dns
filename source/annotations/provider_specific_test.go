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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

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
				CloudflareProxiedKey: "true",
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareProxiedKey, Value: "true"},
			},
			setIdentifier: "",
		},
		{
			name: "Cloudflare custom hostname annotation",
			annotations: map[string]string{
				CloudflareCustomHostnameKey: "custom.example.com",
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareCustomHostnameKey, Value: "custom.example.com"},
			},
			setIdentifier: "",
		},
		{
			name: "AWS annotation",
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/aws-weight": "100",
			},
			expected: endpoint.ProviderSpecific{
				{Name: "aws/weight", Value: "100"},
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, setIdentifier := ProviderSpecificAnnotations(tt.annotations)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.setIdentifier, setIdentifier)
		})
	}
}

func TestGetProviderSpecificCloudflareAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title         string
		annotations   map[string]string
		expectedKey   string
		expectedValue bool
	}{
		{
			title:         "Cloudflare proxied annotation is set correctly to true",
			annotations:   map[string]string{CloudflareProxiedKey: "true"},
			expectedKey:   CloudflareProxiedKey,
			expectedValue: true,
		},
		{
			title:         "Cloudflare proxied annotation is set correctly to false",
			annotations:   map[string]string{CloudflareProxiedKey: "false"},
			expectedKey:   CloudflareProxiedKey,
			expectedValue: false,
		},
		{
			title: "Cloudflare proxied annotation among another annotations is set correctly to true",
			annotations: map[string]string{
				"random annotation 1": "random value 1",
				CloudflareProxiedKey:  "false",
				"random annotation 2": "random value 2",
			},
			expectedKey:   CloudflareProxiedKey,
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
			annotations:   map[string]string{CloudflareRegionKey: "us"},
			expectedKey:   CloudflareRegionKey,
			expectedValue: "us",
		},
		{
			title: "Cloudflare region key annotation among another annotations is set correctly",
			annotations: map[string]string{
				"random annotation 1": "random value 1",
				CloudflareRegionKey:   "us",
				"random annotation 2": "random value 2",
			},
			expectedKey:   CloudflareRegionKey,
			expectedValue: "us",
		},
		{
			title: "Cloudflare DNS record comment annotation is set correctly",
			annotations: map[string]string{
				CloudflareRecordCommentKey: "comment",
			},
			expectedKey:   CloudflareRecordCommentKey,
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
			annotations:   map[string]string{CloudflareCustomHostnameKey: "a.foo.fancybar.com"},
			expectedKey:   CloudflareCustomHostnameKey,
			expectedValue: "a.foo.fancybar.com",
		},
		{
			title: "Cloudflare custom hostname annotation among another annotations is set correctly",
			annotations: map[string]string{
				"random annotation 1":       "random value 1",
				CloudflareCustomHostnameKey: "a.foo.fancybar.com",
				"random annotation 2":       "random value 2"},
			expectedKey:   CloudflareCustomHostnameKey,
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
				if providerSpecificAnnotation.Name == "alias" {
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
				if providerSpecificAnnotation.Name == "alias" {
					t.Error("provider specific annotation alias is not expected to be set")
				}
			}

		})
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
				"external-dns.alpha.kubernetes.io/aws-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.alpha.kubernetes.io/aws-annotation-2": "value 2",
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
				"external-dns.alpha.kubernetes.io/scw-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.alpha.kubernetes.io/scw-annotation-2": "value 2",
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
				"external-dns.alpha.kubernetes.io/webhook-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.alpha.kubernetes.io/webhook-annotation-2": "value 2",
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
