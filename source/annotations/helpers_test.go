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
				TargetAnnotationKey: "example.com",
			},
			expected: endpoint.Targets{"example.com"},
		},
		{
			name: "multiple target annotations",
			annotations: map[string]string{
				TargetAnnotationKey: "example.com,example.org",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
		{
			name: "target annotation with trailing periods",
			annotations: map[string]string{
				TargetAnnotationKey: "example.com.,example.org.",
			},
			expected: endpoint.Targets{"example.com", "example.org"},
		},
		{
			name: "target annotation with spaces",
			annotations: map[string]string{
				TargetAnnotationKey: " example.com , example.org ",
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
				TtlAnnotationKey: "600",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(600),
		},
		{
			name: "invalid TTL annotation",
			annotations: map[string]string{
				TtlAnnotationKey: "invalid",
			},
			resource:    "test-resource",
			expectedTTL: endpoint.TTL(0),
		},
		{
			name: "TTL annotation out of range",
			annotations: map[string]string{
				TtlAnnotationKey: "999999",
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
			annotations: map[string]string{AliasAnnotationKey: "true"},
			expected:    true,
		},
		{
			name:        "alias annotation exists and is false",
			annotations: map[string]string{AliasAnnotationKey: "false"},
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
