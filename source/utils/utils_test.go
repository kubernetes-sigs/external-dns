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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestTTLFromAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		resource    string
		expectedTTL endpoint.TTL
	}{
		{
			name:        "no TTL annotation",
			annotations: map[string]string{},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(0),
		},
		{
			name: "valid TTL annotation",
			annotations: map[string]string{
				ttlAnnotationKey: "600",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(600),
		},
		{
			name: "invalid TTL annotation",
			annotations: map[string]string{
				ttlAnnotationKey: "invalid",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(0),
		},
		{
			name: "TTL annotation out of range",
			annotations: map[string]string{
				ttlAnnotationKey: "999999",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(999999),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttl := TTLFromAnnotations(tt.annotations, tt.resource)
			assert.Equal(t, tt.expectedTTL, ttl)
		})
	}
}

func TestGetAliasFromAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		expected    bool
	}{
		{
			name:        "alias annotation exists and is true",
			annotations: map[string]string{aliasAnnotationKey: "true"},
			expected:    true,
		},
		{
			name:        "alias annotation exists and is false",
			annotations: map[string]string{aliasAnnotationKey: "false"},
			expected:    false,
		},
		{
			name:        "alias annotation does not exist",
			annotations: map[string]string{},
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getAliasFromAnnotations(tt.annotations)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSuitableType(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		expected string
	}{
		{
			name:     "valid IPv4 address",
			target:   "192.168.1.1",
			expected: endpoint.RecordTypeA,
		},
		{
			name:     "valid IPv6 address",
			target:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: endpoint.RecordTypeAAAA,
		},
		{
			name:     "invalid IP address, should return CNAME",
			target:   "example.com",
			expected: endpoint.RecordTypeCNAME,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := suitableType(tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseIngress(t *testing.T) {
	tests := []struct {
		name      string
		ingress   string
		wantNS    string
		wantName  string
		wantError bool
	}{
		{
			name:      "valid namespace and name",
			ingress:   "default/test-ingress",
			wantNS:    "default",
			wantName:  "test-ingress",
			wantError: false,
		},
		{
			name:      "only name provided",
			ingress:   "test-ingress",
			wantNS:    "",
			wantName:  "test-ingress",
			wantError: false,
		},
		{
			name:      "invalid format",
			ingress:   "default/test/ingress",
			wantNS:    "",
			wantName:  "",
			wantError: true,
		},
		{
			name:      "empty string",
			ingress:   "",
			wantNS:    "",
			wantName:  "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNS, gotName, err := ParseIngress(tt.ingress)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantNS, gotNS)
			assert.Equal(t, tt.wantName, gotName)
		})
	}
}

func TestSelectorMatchesServiceSelector(t *testing.T) {
	tests := []struct {
		name        string
		selector    map[string]string
		svcSelector map[string]string
		expected    bool
	}{
		{
			name:        "all key-value pairs match",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{"app": "nginx", "env": "prod"},
			expected:    true,
		},
		{
			name:        "one key-value pair does not match",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{"app": "nginx", "env": "dev"},
			expected:    false,
		},
		{
			name:        "key not present in svcSelector",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{"app": "nginx"},
			expected:    false,
		},
		{
			name:        "empty selector",
			selector:    map[string]string{},
			svcSelector: map[string]string{"app": "nginx", "env": "prod"},
			expected:    true,
		},
		{
			name:        "empty svcSelector",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{},
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectorMatchesServiceSelector(tt.selector, tt.svcSelector)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTargetsFromTargetAnnotation(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		expected    endpoint.Targets
	}{
		{
			name:        "no target annotation",
			annotations: map[string]string{},
			expected:    endpoint.Targets(nil),
		},
		{
			name: "single target annotation",
			annotations: map[string]string{
				targetAnnotationKey: "example.com",
			},
			expected: endpoint.Targets{"example.com"},
		},
		{
			name: "multiple target annotations",
			annotations: map[string]string{
				targetAnnotationKey: "example.com,example.org",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
		{
			name: "target annotation with trailing periods",
			annotations: map[string]string{
				targetAnnotationKey: "example.com.,example.org.",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
		{
			name: "target annotation with spaces",
			annotations: map[string]string{
				targetAnnotationKey: " example.com , example.org ",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TargetsFromTargetAnnotation(tt.annotations)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseAnnotationFilter(t *testing.T) {
	tests := []struct {
		name             string
		annotationFilter string
		expectedSelector labels.Selector
		expectError      bool
	}{
		{
			name:             "valid annotation filter",
			annotationFilter: "key1=value1,key2=value2",
			expectedSelector: labels.Set{"key1": "value1", "key2": "value2"}.AsSelector(),
			expectError:      false,
		},
		{
			name:             "invalid annotation filter",
			annotationFilter: "key1==value1",
			expectedSelector: labels.Set{"key1": "value1"}.AsSelector(),
			expectError:      false,
		},
		{
			name:             "empty annotation filter",
			annotationFilter: "",
			expectedSelector: labels.Set{}.AsSelector(),
			expectError:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector, err := ParseAnnotationFilter(tt.annotationFilter)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSelector, selector)
			}
		})
	}
}
