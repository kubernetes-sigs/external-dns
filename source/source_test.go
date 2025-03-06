/*
Copyright 2017 The Kubernetes Authors.

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

package source

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	for _, tt := range []struct {
		name                     string
		annotationFilter         string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		expectError              bool
	}{
		{
			name:         "invalid template",
			expectError:  true,
			fqdnTemplate: "{{.Name",
		},
		{
			name:        "valid empty template",
			expectError: false,
		},
		{
			name:         "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			name:         "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
		{
			name:                     "valid template",
			expectError:              false,
			fqdnTemplate:             "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			name:             "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "kubernetes.io/ingress.class=nginx",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseTemplate(tt.fqdnTemplate)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFqdnTemplate(t *testing.T) {
	tests := []struct {
		name          string
		fqdnTemplate  string
		expectedError bool
	}{
		{
			name:          "empty template",
			fqdnTemplate:  "",
			expectedError: false,
		},
		{
			name:          "valid template",
			fqdnTemplate:  "{{ .Name }}.example.com",
			expectedError: false,
		},
	}

<<<<<<< HEAD
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := parseTemplate(tt.fqdnTemplate)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, tmpl)
			} else {
				assert.NoError(t, err)
				if tt.fqdnTemplate == "" {
					assert.Nil(t, tmpl)
				} else {
					assert.NotNil(t, tmpl)
=======
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
			providerSpecificAnnotations, _ := getProviderSpecificAnnotations(tc.annotations)
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
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := getProviderSpecificAnnotations(tc.annotations)
			for _, providerSpecificAnnotation := range providerSpecificAnnotations {
				if providerSpecificAnnotation.Name == tc.expectedKey {
					assert.Equal(t, tc.expectedValue, providerSpecificAnnotation.Value)
					return
				}
			}
			t.Errorf("Cloudflare provider specific annotation %s is not set correctly to %v", tc.expectedKey, tc.expectedValue)
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
			annotations:   map[string]string{aliasAnnotationKey: "true"},
			expectedKey:   aliasAnnotationKey,
			expectedValue: true,
		},
		{
			title: "alias annotation among another annotations is set correctly to true",
			annotations: map[string]string{
				"random annotation 1": "random value 1",
				aliasAnnotationKey:    "true",
				"random annotation 2": "random value 2",
			},
			expectedKey:   aliasAnnotationKey,
			expectedValue: true,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			providerSpecificAnnotations, _ := getProviderSpecificAnnotations(tc.annotations)
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
			annotations: map[string]string{aliasAnnotationKey: "false"},
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
			providerSpecificAnnotations, _ := getProviderSpecificAnnotations(tc.annotations)
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
			title: "ibmcloud- provider specific annotations are set correctly",
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/ibmcloud-annotation-1": "value 1",
				SetIdentifierKey: "id1",
				"external-dns.alpha.kubernetes.io/ibmcloud-annotation-2": "value 2",
			},
			expectedResult: map[string]string{
				"ibmcloud-annotation-1": "value 1",
				"ibmcloud-annotation-2": "value 2",
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
			providerSpecificAnnotations, identifier := getProviderSpecificAnnotations(tc.annotations)
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
>>>>>>> c3f0cd66 (fix cloudflare regional hostnames)
				}
			}
		})
	}
}

type mockInformerFactory struct {
	syncResults map[reflect.Type]bool
}

func (m *mockInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	return m.syncResults
}

func TestWaitForCacheSync(t *testing.T) {
	tests := []struct {
		name        string
		syncResults map[reflect.Type]bool
		expectError bool
	}{
		{
			name:        "all caches synced",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): true},
			expectError: false,
		},
		{
			name:        "some caches not synced",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): false},
			expectError: true,
		},
		{
			name:        "context timeout",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): false},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()

			factory := &mockInformerFactory{syncResults: tt.syncResults}
			err := waitForCacheSync(ctx, factory)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
