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
		{
			name: "Multiple annotations",
			annotations: map[string]string{
				CloudflareProxiedKey:                          "true",
				"external-dns.alpha.kubernetes.io/aws-weight": "100",
				SetIdentifierKey:                              "identifier",
			},
			expected: endpoint.ProviderSpecific{
				{Name: CloudflareProxiedKey, Value: "true"},
				{Name: "aws/weight", Value: "100"},
			},
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
