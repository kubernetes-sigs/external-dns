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

func TestNormalizeDNSName(tt *testing.T) {
	records := []struct {
		dnsName string
		expect  string
	}{
		{
			"3AAAA.FOO.BAR.COM    ",
			"3aaaa.foo.bar.com.",
		},
		{
			"   example.foo.com.",
			"example.foo.com.",
		},
		{
			"example123.foo.com ",
			"example123.foo.com.",
		},
		{
			"foo",
			"foo.",
		},
		{
			"123foo.bar",
			"123foo.bar.",
		},
		{
			"foo.com",
			"foo.com.",
		},
		{
			"foo.com.",
			"foo.com.",
		},
		{
			"_foo.com.",
			"_foo.com.",
		},
		{
			"\u005Ffoo.com.",
			"_foo.com.",
		},
		{
			".foo.com.",
			".foo.com.",
		},
		{
			"foo123.COM",
			"foo123.com.",
		},
		{
			"my-exaMple3.FOO.BAR.COM",
			"my-example3.foo.bar.com.",
		},
		{
			"   my-example1214.FOO-1235.BAR-foo.COM   ",
			"my-example1214.foo-1235.bar-foo.com.",
		},
		{
			"my-example-my-example-1214.FOO-1235.BAR-foo.COM",
			"my-example-my-example-1214.foo-1235.bar-foo.com.",
		},
		{
			"點看.org.",
			"xn--c1yn36f.org.",
		},
		{
			"nordic-ø.xn--kitty-點看pd34d.com",
			"xn--nordic--w1a.xn--xn--kitty-pd34d-hn01b3542b.com.",
		},
		{
			"nordic-ø.kitty😸.com.",
			"xn--nordic--w1a.xn--kitty-pd34d.com.",
		},
		{
			"  nordic-ø.kitty😸.COM",
			"xn--nordic--w1a.xn--kitty-pd34d.com.",
		},
		{
			"xn--nordic--w1a.kitty😸.com.",
			"xn--nordic--w1a.xn--kitty-pd34d.com.",
		},
		{
			"*.example.com.",
			"*.example.com.",
		},
		{
			"*.example.com",
			"*.example.com.",
		},
	}
	for _, r := range records {
		tt.Run(r.dnsName, func(t *testing.T) {
			gotName := NormalizeDNSName(r.dnsName)
			assert.Equal(t, r.expect, gotName)
		})
	}
}

func TestToASCII(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "unicode domain",
			input: "例子.com",
			want:  "xn--fsqu00a.com",
		},
		{
			name:  "invalid input returns original",
			input: "a\x00b.com",
			want:  "a\x00b.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ToASCII(tt.input))
		})
	}
}

func TestToUnicode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "punycode domain",
			input: "xn--fsqu00a.com",
			want:  "例子.com",
		},
		{
			name:  "invalid input returns original",
			input: "a\x00b.com",
			want:  "a\x00b.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ToUnicode(tt.input))
		})
	}
}
