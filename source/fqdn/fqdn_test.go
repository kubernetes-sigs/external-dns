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

package fqdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestParseTemplate(t *testing.T) {
	for _, tt := range []struct {
		name                     string
		annotationFilter         string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		expectError              bool
	}{
		{
			name:         "invalid template",
			expectError:  true,
			fqdnTemplate: "{{.Name",
		},
		{
			name:        "valid empty template",
			expectError: false,
		},
		{
			name:         "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			name:         "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
		{
			name:                     "valid template",
			expectError:              false,
			fqdnTemplate:             "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			name:             "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "kubernetes.io/ingress.class=nginx",
		},
		{
			name:         "replace template function",
			expectError:  false,
			fqdnTemplate: "{{\"hello.world\" | replace \".\" \"-\"}}.ext-dns.test.com",
		},
		{
			name:         "isIPv4 template function with valid IPv4",
			expectError:  false,
			fqdnTemplate: "{{if isIPv4 \"192.168.1.1\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name:         "isIPv4 template function with invalid IPv4",
			expectError:  false,
			fqdnTemplate: "{{if isIPv4 \"not.an.ip.addr\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name:         "isIPv6 template function with valid IPv6",
			expectError:  false,
			fqdnTemplate: "{{if isIPv6 \"2001:db8::1\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name:         "isIPv6 template function with invalid IPv6",
			expectError:  false,
			fqdnTemplate: "{{if isIPv6 \"not:ipv6:addr\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTemplate(tt.fqdnTemplate)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExecTemplate(t *testing.T) {
	tests := []struct {
		name    string
		tmpl    string
		obj     kubeObject
		want    []string
		wantErr bool
	}{
		{
			name: "simple template",
			tmpl: "{{ .Name }}.example.com, {{ .Namespace }}.example.org",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
				},
			},
			want: []string{"test.example.com", "default.example.org"},
		},
		{
			name: "multiple hostnames",
			tmpl: "{{.Name}}.example.com, {{.Name}}.example.org",
			obj: &testObject{

				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
				},
			},
			want: []string{"test.example.com", "test.example.org"},
		},
		{
			name: "trim spaces",
			tmpl: "  {{ trim .Name}}.example.com. ",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: " test ",
				},
			},
			want: []string{"test.example.com"},
		},
		{
			name: "annotations and labels",
			tmpl: "{{.Labels.environment }}.example.com, {{ index .ObjectMeta.Annotations \"alb.ingress.kubernetes.io/scheme\" }}.{{ .Labels.environment }}.{{ index .ObjectMeta.Annotations \"dns.company.com/zone\" }}",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/hostname": "test.example.com, test.example.org",
						"kubernetes.io/role/internal-elb":           "true",
						"alb.ingress.kubernetes.io/scheme":          "internal",
						"dns.company.com/zone":                      "company.org",
					},
					Labels: map[string]string{
						"environment": "production",
						"app":         "myapp",
						"tier":        "backend",
						"role":        "worker",
						"version":     "1",
					},
				},
			},
			want: []string{"production.example.com", "internal.production.company.org"},
		},
		{
			name: "labels to lowercase",
			tmpl: "{{ toLower .Labels.department }}.example.org",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
					Labels: map[string]string{
						"department": "FINANCE",
						"app":        "myapp",
					},
				},
			},
			want: []string{"finance.example.org"},
		},
		{
			name: "generate multiple hostnames with if condition",
			tmpl: "{{ if contains (index .ObjectMeta.Annotations \"external-dns.alpha.kubernetes.io/hostname\") \"example.com\" }}{{ toLower .Labels.hostoverride }}{{end}}",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "default",
					Labels: map[string]string{
						"hostoverride": "abrakadabra.google.com",
						"app":          "myapp",
					},
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/hostname": "test.example.com",
					},
				},
			},
			want: []string{"abrakadabra.google.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := ParseTemplate(tt.tmpl)
			require.NoError(t, err)

			got, err := ExecTemplate(tmpl, tt.obj)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExecTemplateEmptyObject(t *testing.T) {
	tmpl, err := ParseTemplate("{{ toLower .Labels.department }}.example.org")
	require.NoError(t, err)
	_, err = ExecTemplate(tmpl, nil)
	assert.Error(t, err)
}

func TestFqdnTemplate(t *testing.T) {
	tests := []struct {
		name          string
		fqdnTemplate  string
		expectedError bool
	}{
		{
			name:          "empty template",
			fqdnTemplate:  "",
			expectedError: false,
		},
		{
			name:          "valid template",
			fqdnTemplate:  "{{ .Name }}.example.com",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := ParseTemplate(tt.fqdnTemplate)
			if tt.expectedError {
				require.Error(t, err)
				assert.Nil(t, tmpl)
			} else {
				require.NoError(t, err)
				if tt.fqdnTemplate == "" {
					assert.Nil(t, tmpl)
				} else {
					assert.NotNil(t, tmpl)
				}
			}
		})
	}
}

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
			result := isIPv6String(tt.input)
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
			result := isIPv4String(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

type testObject struct {
	metav1.ObjectMeta
	runtime.Object
}

func (t *testObject) DeepCopyObject() runtime.Object {
	return &testObject{
		ObjectMeta: *t.ObjectMeta.DeepCopy(),
	}
}
