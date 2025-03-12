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
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestGetTTLFromAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title       string
		annotations map[string]string
		expectedTTL endpoint.TTL
	}{
		{
			title:       "TTL annotation not present",
			annotations: map[string]string{"foo": "bar"},
			expectedTTL: endpoint.TTL(0),
		},
		{
			title:       "TTL annotation value is not a number",
			annotations: map[string]string{ttlAnnotationKey: "foo"},
			expectedTTL: endpoint.TTL(0),
		},
		{
			title:       "TTL annotation value is empty",
			annotations: map[string]string{ttlAnnotationKey: ""},
			expectedTTL: endpoint.TTL(0),
		},
		{
			title:       "TTL annotation value is negative number",
			annotations: map[string]string{ttlAnnotationKey: "-1"},
			expectedTTL: endpoint.TTL(0),
		},
		{
			title:       "TTL annotation value is too high",
			annotations: map[string]string{ttlAnnotationKey: fmt.Sprintf("%d", 1<<32)},
			expectedTTL: endpoint.TTL(0),
		},
		{
			title:       "TTL annotation value is set correctly using integer",
			annotations: map[string]string{ttlAnnotationKey: "60"},
			expectedTTL: endpoint.TTL(60),
		},
		{
			title:       "TTL annotation value is set correctly using duration (whole)",
			annotations: map[string]string{ttlAnnotationKey: "10m"},
			expectedTTL: endpoint.TTL(600),
		},
		{
			title:       "TTL annotation value is set correctly using duration (fractional)",
			annotations: map[string]string{ttlAnnotationKey: "20.5s"},
			expectedTTL: endpoint.TTL(20),
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			ttl := getTTLFromAnnotations(tc.annotations, "resource/test")
			assert.Equal(t, tc.expectedTTL, ttl)
		})
	}
}

func TestSuitableType(t *testing.T) {
	for _, tc := range []struct {
		target, recordType, expected string
	}{
		{"8.8.8.8", "", "A"},
		{"2001:db8::1", "", "AAAA"},
		{"::ffff:c0a8:101", "", "AAAA"},
		{"foo.example.org", "", "CNAME"},
		{"bar.eu-central-1.elb.amazonaws.com", "", "CNAME"},
	} {

		recordType := suitableType(tc.target)

		if recordType != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, recordType)
		}
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
			providerSpecificAnnotations, _ := getProviderSpecificAnnotations(tc.annotations)
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
				}
			}
		})
	}
}
