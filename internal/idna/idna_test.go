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

package idna

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileWithDefault(t *testing.T) {
	tets := []struct {
		input    string
		expected string
	}{
		{
			input:    "*.GÖPHER.com",
			expected: "*.göpher.com",
		},
		{
			input:    "*._abrakadabra.com",
			expected: "*._abrakadabra.com",
		},
		{
			input:    "_abrakadabra.com",
			expected: "_abrakadabra.com",
		},
		{
			input:    "*.foo.kube.example.com",
			expected: "*.foo.kube.example.com",
		},
		{
			input:    "xn--bcher-kva.example.com",
			expected: "bücher.example.com",
		},
	}
	for _, tt := range tets {
		t.Run(strings.ToLower(tt.input), func(t *testing.T) {
			result, err := Profile.ToUnicode(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
