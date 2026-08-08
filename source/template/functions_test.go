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

package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	for _, tt := range []struct {
		name     string
		oldValue string
		newValue string
		target   string
		expected string
	}{
		{
			name:     "simple replacement",
			oldValue: "old",
			newValue: "new",
			target:   "old-value",
			expected: "new-value",
		},
		{
			name:     "multiple replacements",
			oldValue: ".",
			newValue: "-",
			target:   "hello.world.com",
			expected: "hello-world-com",
		},
		{
			name:     "no replacement needed",
			oldValue: "x",
			newValue: "y",
			target:   "hello-world",
			expected: "hello-world",
		},
		{
			name:     "empty strings",
			oldValue: "",
			newValue: "",
			target:   "test",
			expected: "test",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := replace(tt.oldValue, tt.newValue, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsIPv6String(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid IPv6",
			input:    "2001:db8::1",
			expected: true,
		},
		{
			name:     "valid IPv6 with multiple segments",
			input:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: true,
		},
		{
			name:     "valid IPv4-mapped IPv6",
			input:    "::ffff:192.168.1.1",
			expected: true,
		},
		{
			name:     "invalid IPv6",
			input:    "not:ipv6:addr",
			expected: false,
		},
		{
			name:     "IPv4 address",
			input:    "192.168.1.1",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := isIPv6(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsIPv4String(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid IPv4",
			input:    "192.168.1.1",
			expected: true,
		},
		{
			name:     "invalid IPv4",
			input:    "256.256.256.256",
			expected: false,
		},
		{
			name:     "IPv6 address",
			input:    "2001:db8::1",
			expected: false,
		},
		{
			name:     "invalid format",
			input:    "not.an.ip",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := isIPv4(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHasKey(t *testing.T) {
	for _, tt := range []struct {
		name     string
		m        map[string]string
		key      string
		expected bool
	}{
		{
			name:     "key exists with non-empty value",
			m:        map[string]string{"foo": "bar"},
			key:      "foo",
			expected: true,
		},
		{
			name:     "key exists with empty value",
			m:        map[string]string{"service.kubernetes.io/headless": ""},
			key:      "service.kubernetes.io/headless",
			expected: true,
		},
		{
			name:     "key does not exist",
			m:        map[string]string{"foo": "bar"},
			key:      "baz",
			expected: false,
		},
		{
			name:     "nil map",
			m:        nil,
			key:      "foo",
			expected: false,
		},
		{
			name:     "empty map",
			m:        map[string]string{},
			key:      "foo",
			expected: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := hasKey(tt.m, tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFromJson(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected any
	}{
		{
			name:     "map of strings",
			input:    `{"dns":"entry1.internal.tld","target":"10.10.10.10"}`,
			expected: map[string]any{"dns": "entry1.internal.tld", "target": "10.10.10.10"},
		},
		{
			name:  "slice of maps",
			input: `[{"dns":"entry1.internal.tld","target":"10.10.10.10"},{"dns":"entry2.example.tld","target":"my.cluster.local"}]`,
			expected: []any{
				map[string]any{"dns": "entry1.internal.tld", "target": "10.10.10.10"},
				map[string]any{"dns": "entry2.example.tld", "target": "my.cluster.local"},
			},
		},
		{
			name:     "null input",
			input:    "null",
			expected: nil,
		},
		{
			name:     "empty object",
			input:    "{}",
			expected: map[string]any{},
		},
		{
			name:     "string value",
			input:    `"hello"`,
			expected: "hello",
		},
		{
			name:     "invalid json",
			input:    "not valid json",
			expected: nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := fromJson(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
