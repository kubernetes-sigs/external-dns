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
	"testing"

	"sigs.k8s.io/external-dns/pkg/apis"

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

func TestProviderSpecificAnnotations(t *testing.T) {
	for _, tc := range []struct {
		title                  string
		annotations            map[string]string
		providerSpecificConfig apis.ProviderSpecificConfig
		expectedPS             endpoint.ProviderSpecific
	}{
		{
			title:       "PS annotations are not present",
			annotations: map[string]string{"foo": "bar"},
			expectedPS:  endpoint.ProviderSpecific{},
		},
		{
			title:       "Cloudflare annotation is present",
			annotations: map[string]string{"external-dns.alpha.kubernetes.io/cloudflare-proxied": "true"},
			providerSpecificConfig: apis.ProviderSpecificConfig{
				Translation: map[string]string{
					"external-dns.alpha.kubernetes.io/cloudflare-proxied": "external-dns.alpha.kubernetes.io/cloudflare-proxied",
				},
			},
			expectedPS: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
					Value: "true",
				},
			},
		},
		{
			title:       "AWS annotation is present",
			annotations: map[string]string{"external-dns.alpha.kubernetes.io/aws-weight": "10"},
			providerSpecificConfig: apis.ProviderSpecificConfig{
				PrefixTranslation: map[string]string{
					"external-dns.alpha.kubernetes.io/aws-": "aws/",
				},
			},
			expectedPS: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "aws/weight",
					Value: "10",
				},
			},
		},
		{
			title:       "AWS annotation is present with generic",
			annotations: map[string]string{"aws.provider.external-dns.alpha.kubernetes.io/weight": "10"},
			providerSpecificConfig: apis.ProviderSpecificConfig{
				PrefixTranslation: map[string]string{
					"aws.provider.external-dns.alpha.kubernetes.io/": "aws/",
				},
			},
			expectedPS: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "aws/weight",
					Value: "10",
				},
			},
		},
		{
			title:       "Generic annotation is present",
			annotations: map[string]string{"generic.provider.external-dns.alpha.kubernetes.io/weight": "10"},
			providerSpecificConfig: apis.ProviderSpecificConfig{
				PrefixTranslation: map[string]string{
					"generic.provider.external-dns.alpha.kubernetes.io/": "generic/",
				},
			},
			expectedPS: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "generic/weight",
					Value: "10",
				},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			source := BaseSource{}
			source.ProviderSpecificConfig = tc.providerSpecificConfig
			providerSpecificAnnotations, _ := source.GetProviderSpecificAnnotations(tc.annotations)
			assert.Equal(t, tc.expectedPS, providerSpecificAnnotations)
		})
	}
}

func TestSuitableType(t *testing.T) {
	for _, tc := range []struct {
		target, recordType, expected string
	}{
		{"8.8.8.8", "", "A"},
		{"2001:db8::1", "", "AAAA"},
		{"foo.example.org", "", "CNAME"},
		{"bar.eu-central-1.elb.amazonaws.com", "", "CNAME"},
	} {

		recordType := suitableType(tc.target)

		if recordType != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, recordType)
		}
	}
}
