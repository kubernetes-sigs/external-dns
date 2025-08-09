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

package informers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/labels"
)

func TestToSHA(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "test",
			expected: "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
		},
		{
			input:    "",
			expected: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		{
			input: labels.Set(map[string]string{
				"app": "test",
				"env": "production",
			}).String(),
			expected: "29eda95ee609e3186afe17e3bf988a654bc5b739",
		},
		{
			input: labels.Set(map[string]string{
				"app":       "test",
				"env":       "production",
				"version":   "v1",
				"component": "frontend",
			}).String(),
			expected: "446f9bdf6ba5c7edf324a07482bcd5c3b6c6ce38",
		},
	}

	for _, tt := range tests {
		got := ToSHA(tt.input)
		assert.Equal(t, tt.expected, got)
	}
}
