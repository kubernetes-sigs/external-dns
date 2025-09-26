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
package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestToInternalEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		input    *apiv1alpha1.Endpoint
		expected *endpoint.Endpoint
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name: "basic endpoint conversion",
			input: &apiv1alpha1.Endpoint{
				DNSName:    "test.example.com",
				Targets:    []string{"1.2.3.4"},
				RecordType: "A",
				RecordTTL:  300,
			},
			expected: &endpoint.Endpoint{
				DNSName:          "test.example.com",
				Targets:          []string{"1.2.3.4"},
				RecordType:       "A",
				RecordTTL:        300,
				ProviderSpecific: endpoint.ProviderSpecific{},
			},
		},
		{
			name: "endpoint with provider specific properties",
			input: &apiv1alpha1.Endpoint{
				DNSName:    "test.example.com",
				Targets:    []string{"1.2.3.4"},
				RecordType: "A",
				RecordTTL:  300,
				ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{
					{Name: "aws-alias", Value: "true"},
					{Name: "aws-zone-id", Value: "Z12345"},
				},
			},
			expected: &endpoint.Endpoint{
				DNSName:    "test.example.com",
				Targets:    []string{"1.2.3.4"},
				RecordType: "A",
				RecordTTL:  300,
				ProviderSpecific: endpoint.ProviderSpecific{
					"aws-alias":   "true",
					"aws-zone-id": "Z12345",
				},
			},
		},
		{
			name: "endpoint with labels and set identifier",
			input: &apiv1alpha1.Endpoint{
				DNSName:       "test.example.com",
				Targets:       []string{"1.2.3.4"},
				RecordType:    "A",
				RecordTTL:     300,
				SetIdentifier: "us-east-1",
				Labels: map[string]string{
					"owner":     "external-dns",
					"resource":  "ingress/default/test",
					"namespace": "default",
				},
			},
			expected: &endpoint.Endpoint{
				DNSName:       "test.example.com",
				Targets:       []string{"1.2.3.4"},
				RecordType:    "A",
				RecordTTL:     300,
				SetIdentifier: "us-east-1",
				Labels: map[string]string{
					"owner":     "external-dns",
					"resource":  "ingress/default/test",
					"namespace": "default",
				},
				ProviderSpecific: endpoint.ProviderSpecific{},
			},
		},
		{
			name: "complete endpoint with all fields",
			input: &apiv1alpha1.Endpoint{
				DNSName:       "api.example.com",
				Targets:       []string{"10.0.0.1", "10.0.0.2"},
				RecordType:    "A",
				RecordTTL:     600,
				SetIdentifier: "primary",
				Labels: map[string]string{
					"owner": "external-dns",
					"type":  "public",
				},
				ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{
					{Name: "weight", Value: "100"},
					{Name: "policy", Value: "weighted"},
				},
			},
			expected: &endpoint.Endpoint{
				DNSName:       "api.example.com",
				Targets:       []string{"10.0.0.1", "10.0.0.2"},
				RecordType:    "A",
				RecordTTL:     600,
				SetIdentifier: "primary",
				Labels: map[string]string{
					"owner": "external-dns",
					"type":  "public",
				},
				ProviderSpecific: endpoint.ProviderSpecific{
					"weight": "100",
					"policy": "weighted",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromAPIEndpoint(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.True(t, testutils.SameEndpoint(tt.expected, result))
			}
		})
	}
}

func TestToAPIEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		input    *endpoint.Endpoint
		expected *apiv1alpha1.Endpoint
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name: "basic endpoint conversion",
			input: &endpoint.Endpoint{
				DNSName:          "test.example.com",
				Targets:          []string{"1.2.3.4"},
				RecordType:       "A",
				RecordTTL:        300,
				ProviderSpecific: endpoint.ProviderSpecific{},
			},
			expected: &apiv1alpha1.Endpoint{
				DNSName:          "test.example.com",
				Targets:          []string{"1.2.3.4"},
				RecordType:       "A",
				RecordTTL:        300,
				ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{},
			},
		},
		{
			name: "endpoint with provider specific properties",
			input: &endpoint.Endpoint{
				DNSName:    "test.example.com",
				Targets:    []string{"1.2.3.4"},
				RecordType: "A",
				RecordTTL:  300,
				ProviderSpecific: endpoint.ProviderSpecific{
					"aws-alias":   "true",
					"aws-zone-id": "Z12345",
				},
			},
			expected: &apiv1alpha1.Endpoint{
				DNSName:    "test.example.com",
				Targets:    []string{"1.2.3.4"},
				RecordType: "A",
				RecordTTL:  300,
				ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{
					{Name: "aws-alias", Value: "true"},
					{Name: "aws-zone-id", Value: "Z12345"},
				},
			},
		},
		{
			name: "endpoint with labels and set identifier",
			input: &endpoint.Endpoint{
				DNSName:       "test.example.com",
				Targets:       []string{"1.2.3.4"},
				RecordType:    "A",
				RecordTTL:     300,
				SetIdentifier: "us-east-1",
				Labels: map[string]string{
					"owner":     "external-dns",
					"resource":  "ingress/default/test",
					"namespace": "default",
				},
				ProviderSpecific: endpoint.ProviderSpecific{},
			},
			expected: &apiv1alpha1.Endpoint{
				DNSName:       "test.example.com",
				Targets:       []string{"1.2.3.4"},
				RecordType:    "A",
				RecordTTL:     300,
				SetIdentifier: "us-east-1",
				Labels: map[string]string{
					"owner":     "external-dns",
					"resource":  "ingress/default/test",
					"namespace": "default",
				},
				ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToAPIEndpoint(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tt.expected.DNSName, result.DNSName)
				assert.Equal(t, tt.expected.Targets, result.Targets)
				assert.Equal(t, tt.expected.RecordType, result.RecordType)
				assert.Equal(t, tt.expected.RecordTTL, result.RecordTTL)
				assert.Equal(t, tt.expected.SetIdentifier, result.SetIdentifier)
				assert.Equal(t, tt.expected.Labels, result.Labels)

				// Convert ProviderSpecific slice to map for comparison since order doesn't matter
				expectedPS := make(map[string]string)
				for _, ps := range tt.expected.ProviderSpecific {
					expectedPS[ps.Name] = ps.Value
				}
				resultPS := make(map[string]string)
				for _, ps := range result.ProviderSpecific {
					resultPS[ps.Name] = ps.Value
				}
				assert.Equal(t, expectedPS, resultPS)
			}
		})
	}
}

func TestToInternalEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		input    []*apiv1alpha1.Endpoint
		expected []*endpoint.Endpoint
	}{
		{
			name:     "empty slice",
			input:    []*apiv1alpha1.Endpoint{},
			expected: nil,
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name: "single endpoint",
			input: []*apiv1alpha1.Endpoint{
				{
					DNSName:    "test.example.com",
					Targets:    []string{"1.2.3.4"},
					RecordType: "A",
					RecordTTL:  300,
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "test.example.com",
					Targets:          []string{"1.2.3.4"},
					RecordType:       "A",
					RecordTTL:        300,
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			name: "multiple endpoints",
			input: []*apiv1alpha1.Endpoint{
				{
					DNSName:    "api.example.com",
					Targets:    []string{"1.2.3.4"},
					RecordType: "A",
					RecordTTL:  300,
					ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{
						{Name: "type", Value: "api"},
					},
				},
				{
					DNSName:    "www.example.com",
					Targets:    []string{"example.com"},
					RecordType: "CNAME",
					RecordTTL:  600,
					Labels: map[string]string{
						"owner": "external-dns",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "api.example.com",
					Targets:    []string{"1.2.3.4"},
					RecordType: "A",
					RecordTTL:  300,
					ProviderSpecific: endpoint.ProviderSpecific{
						"type": "api",
					},
				},
				{
					DNSName:    "www.example.com",
					Targets:    []string{"example.com"},
					RecordType: "CNAME",
					RecordTTL:  600,
					Labels: map[string]string{
						"owner": "external-dns",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromAPIEndpoints(tt.input)
			assert.True(t, testutils.SameEndpoints(tt.expected, result))
		})
	}
}

func TestToAPIEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		input    []*endpoint.Endpoint
		expected []*apiv1alpha1.Endpoint
	}{
		{
			name:     "empty slice",
			input:    []*endpoint.Endpoint{},
			expected: nil,
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name: "single endpoint",
			input: []*endpoint.Endpoint{
				{
					DNSName:          "test.example.com",
					Targets:          []string{"1.2.3.4"},
					RecordType:       "A",
					RecordTTL:        300,
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
			expected: []*apiv1alpha1.Endpoint{
				{
					DNSName:          "test.example.com",
					Targets:          []string{"1.2.3.4"},
					RecordType:       "A",
					RecordTTL:        300,
					ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{},
				},
			},
		},
		{
			name: "multiple endpoints with provider specific",
			input: []*endpoint.Endpoint{
				{
					DNSName:    "api.example.com",
					Targets:    []string{"1.2.3.4"},
					RecordType: "A",
					RecordTTL:  300,
					ProviderSpecific: endpoint.ProviderSpecific{
						"weight": "100",
					},
				},
				{
					DNSName:    "www.example.com",
					Targets:    []string{"example.com"},
					RecordType: "CNAME",
					RecordTTL:  600,
					Labels: map[string]string{
						"owner": "external-dns",
					},
					ProviderSpecific: endpoint.ProviderSpecific{
						"alias": "true",
					},
				},
			},
			expected: []*apiv1alpha1.Endpoint{
				{
					DNSName:    "api.example.com",
					Targets:    []string{"1.2.3.4"},
					RecordType: "A",
					RecordTTL:  300,
					ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{
						{Name: "weight", Value: "100"},
					},
				},
				{
					DNSName:    "www.example.com",
					Targets:    []string{"example.com"},
					RecordType: "CNAME",
					RecordTTL:  600,
					Labels: map[string]string{
						"owner": "external-dns",
					},
					ProviderSpecific: []apiv1alpha1.ProviderSpecificProperty{
						{Name: "alias", Value: "true"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToAPIEndpoints(tt.input)
			assert.Len(t, result, len(tt.expected))

			for i, expected := range tt.expected {
				if expected == nil {
					assert.Nil(t, result[i])
					continue
				}

				assert.Equal(t, expected.DNSName, result[i].DNSName)
				assert.Equal(t, expected.Targets, result[i].Targets)
				assert.Equal(t, expected.RecordType, result[i].RecordType)
				assert.Equal(t, expected.RecordTTL, result[i].RecordTTL)
				assert.Equal(t, expected.SetIdentifier, result[i].SetIdentifier)
				assert.Equal(t, expected.Labels, result[i].Labels)

				// Convert ProviderSpecific slice to map for comparison since order doesn't matter
				expectedPS := make(map[string]string)
				for _, ps := range expected.ProviderSpecific {
					expectedPS[ps.Name] = ps.Value
				}
				resultPS := make(map[string]string)
				for _, ps := range result[i].ProviderSpecific {
					resultPS[ps.Name] = ps.Value
				}
				assert.Equal(t, expectedPS, resultPS)
			}
		})
	}
}

func TestRoundTripConversion(t *testing.T) {
	tests := []struct {
		name     string
		original *endpoint.Endpoint
	}{
		{
			name: "basic endpoint round trip",
			original: &endpoint.Endpoint{
				DNSName:          "test.example.com",
				Targets:          []string{"1.2.3.4"},
				RecordType:       "A",
				RecordTTL:        300,
				ProviderSpecific: endpoint.ProviderSpecific{},
			},
		},
		{
			name: "complex endpoint round trip",
			original: &endpoint.Endpoint{
				DNSName:       "api.example.com",
				Targets:       []string{"10.0.0.1", "10.0.0.2"},
				RecordType:    "A",
				RecordTTL:     600,
				SetIdentifier: "primary",
				Labels: map[string]string{
					"owner": "external-dns",
					"type":  "public",
				},
				ProviderSpecific: endpoint.ProviderSpecific{
					"weight":       "100",
					"policy":       "weighted",
					"health-check": "enabled",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert internal -> API -> internal
			apiEndpoint := ToAPIEndpoint(tt.original)
			result := FromAPIEndpoint(apiEndpoint)

			assert.True(t, testutils.SameEndpoint(tt.original, result))
		})
	}
}
