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
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/external-dns/endpoint"
)

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
			selector, err := ParseFilter(tt.annotationFilter)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedSelector, selector)
			}
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
				TargetKey: "example.com",
			},
			expected: endpoint.Targets{"example.com"},
		},
		{
			name: "multiple target annotations",
			annotations: map[string]string{
				TargetKey: "example.com,example.org",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
		{
			name: "target annotation with trailing periods",
			annotations: map[string]string{
				TargetKey: "example.com.,example.org.",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
		{
			name: "target annotation with spaces",
			annotations: map[string]string{
				TargetKey: " example.com , example.org ",
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
				TtlKey: "600",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(600),
		},
		{
			name: "invalid TTL annotation",
			annotations: map[string]string{
				TtlKey: "invalid",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(0),
		},
		{
			name: "TTL annotation out of range",
			annotations: map[string]string{
				TtlKey: "999999",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(999999),
		},
		{
			name:        "TTL annotation not present",
			annotations: map[string]string{"foo": "bar"},
			expectedTTL: endpoint.TTL(0),
		},
		{
			name:        "TTL annotation value is not a number",
			annotations: map[string]string{TtlKey: "foo"},
			expectedTTL: endpoint.TTL(0),
		},
		{
			name:        "TTL annotation value is empty",
			annotations: map[string]string{TtlKey: ""},
			expectedTTL: endpoint.TTL(0),
		},
		{
			name:        "TTL annotation value is negative number",
			annotations: map[string]string{TtlKey: "-1"},
			expectedTTL: endpoint.TTL(0),
		},
		{
			name:        "TTL annotation value is too high",
			annotations: map[string]string{TtlKey: fmt.Sprintf("%d", 1<<32)},
			expectedTTL: endpoint.TTL(0),
		},
		{
			name:        "TTL annotation value is set correctly using integer",
			annotations: map[string]string{TtlKey: "60"},
			expectedTTL: endpoint.TTL(60),
		},
		{
			name:        "TTL annotation value is set correctly using duration (whole)",
			annotations: map[string]string{TtlKey: "10m"},
			expectedTTL: endpoint.TTL(600),
		},
		{
			name:        "TTL annotation value is set correctly using duration (fractional)",
			annotations: map[string]string{TtlKey: "20.5s"},
			expectedTTL: endpoint.TTL(20),
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
			annotations: map[string]string{AliasKey: "true"},
			expected:    true,
		},
		{
			name:        "alias annotation exists and is false",
			annotations: map[string]string{AliasKey: "false"},
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
			result := hasAliasFromAnnotations(tt.annotations)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHostnamesFromAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		expected    []string
	}{
		{
			name:        "no hostname annotation",
			annotations: map[string]string{},
			expected:    nil,
		},
		{
			name: "single hostname annotation",
			annotations: map[string]string{
				HostnameKey: "example.com",
			},
			expected: []string{"example.com"},
		},
		{
			name: "multiple hostname annotations",
			annotations: map[string]string{
				HostnameKey: "example.com,example.org",
			},
			expected: []string{"example.com", "example.org"},
		},
		{
			name: "hostname annotation with spaces",
			annotations: map[string]string{
				HostnameKey: " example.com , example.org ",
			},
			expected: []string{"example.com", "example.org"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HostnamesFromAnnotations(tt.annotations)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSplitHostnameAnnotation(t *testing.T) {
	tests := []struct {
		name       string
		annotation string
		expected   []string
	}{
		{
			name:       "empty annotation",
			annotation: "",
			expected:   []string{""},
		},
		{
			name:       "single hostname",
			annotation: "example.com",
			expected:   []string{"example.com"},
		},
		{
			name:       "multiple hostnames",
			annotation: "example.com,example.org",
			expected:   []string{"example.com", "example.org"},
		},
		{
			name:       "hostnames with spaces",
			annotation: " example.com , example.org ",
			expected:   []string{"example.com", "example.org"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitHostnameAnnotation(tt.annotation)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInternalHostnamesFromAnnotations(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		expected    []string
	}{
		{
			name:        "no internal hostname annotation",
			annotations: map[string]string{},
			expected:    nil,
		},
		{
			name: "single internal hostname annotation",
			annotations: map[string]string{
				InternalHostnameKey: "internal.example.com",
			},
			expected: []string{"internal.example.com"},
		},
		{
			name: "multiple internal hostname annotations",
			annotations: map[string]string{
				InternalHostnameKey: "internal.example.com,internal.example.org",
			},
			expected: []string{"internal.example.com", "internal.example.org"},
		},
		{
			name: "internal hostname annotation with spaces",
			annotations: map[string]string{
				InternalHostnameKey: " internal.example.com , internal.example.org ",
			},
			expected: []string{"internal.example.com", "internal.example.org"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InternalHostnamesFromAnnotations(tt.annotations)
			assert.Equal(t, tt.expected, result)
		})
	}
}
