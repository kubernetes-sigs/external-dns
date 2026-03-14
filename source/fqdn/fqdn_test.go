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
	"errors"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/external-dns/endpoint"
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
			want: []string{"default.example.org", "test.example.com"},
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
			name: "trim prefix",
			tmpl: `{{ trimPrefix .Name "the-" }}.example.com`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "the-test",
					Namespace: "default",
				},
			},
			want: []string{"test.example.com"},
		},
		{
			name: "trim suffix",
			tmpl: `{{ trimSuffix .Name "-v2" }}.example.com`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-v2",
					Namespace: "default",
				},
			},
			want: []string{"test.example.com"},
		},
		{
			name: "replace dash",
			tmpl: `{{ replace "-" "." .Name }}.example.com`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-v2",
					Namespace: "default",
				},
			},
			want: []string{"test.v2.example.com"},
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
			want: []string{"internal.production.company.org", "production.example.com"},
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
		{
			name: "ignore empty template output",
			tmpl: "{{ if eq .Name \"other\" }}{{ .Name }}.example.com{{ end }}",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
			},
			want: nil,
		},
		{
			name: "ignore trailing comma output",
			tmpl: "{{ .Name }}.example.com,",
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
			},
			want: []string{"test.example.com"},
		},
		{
			name: "contains label with empty value",
			tmpl: `{{if hasKey .Labels "service.kubernetes.io/headless"}}{{ .Name }}.example.com,{{end}}`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
					Labels: map[string]string{
						"service.kubernetes.io/headless": "",
					},
				},
			},
			want: []string{"test.example.com"},
		},
		{
			name: "result only contains unique values",
			tmpl: `{{ .Name }}.example.com,{{ .Name }}.example.com,{{ .Name }}.example.com`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
					Labels: map[string]string{
						"service.kubernetes.io/headless": "",
					},
				},
			},
			want: []string{"test.example.com"},
		},
		{
			name: "dns entries in labels",
			tmpl: `
{{ if hasKey .Labels "records" }}{{ range $entry := (index .Labels "records" | fromJson) }}{{ index $entry "dns" }},{{ end }}{{ end }}`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
					Labels: map[string]string{
						"records": `
[{"dns":"entry1.internal.tld","target":"10.10.10.10"},{"dns":"entry2.example.tld","target":"my.cluster.local"}]`,
					},
				},
			},
			want: []string{"entry1.internal.tld", "entry2.example.tld"},
		},
		{
			name: "configmap with multiple entries",
			tmpl: `{{ range $entry := (index .Data "entries" | fromJson) }}{{ index $entry "dns" }},{{ end }}`,
			obj: &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-configmap",
				},
				Data: map[string]string{
					"entries": `
[{"dns":"entry1.internal.tld","target":"10.10.10.10"},{"dns":"entry2.example.tld","target":"my.cluster.local"}]`,
				},
			},
			want: []string{"entry1.internal.tld", "entry2.example.tld"},
		},
		{
			name: "rancher publicEndpoints annotation",
			tmpl: `
{{ range $entry := (index .Annotations "field.cattle.io/publicEndpoints" | fromJson) }}{{ index $entry "hostname" }},{{ end }}`,
			obj: &testObject{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
					Annotations: map[string]string{
						"field.cattle.io/publicEndpoints": `
							[{"addresses":[""],"port":80,"protocol":"HTTP",
								"serviceName":"development:keycloak-ha-service",
								"ingressName":"development:keycloak-ha-ingress",
								"hostname":"keycloak.snip.com","allNodes":false
							}]`,
					},
				},
			},
			want: []string{"keycloak.snip.com"},
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

func TestExecTemplatePopulatesEmptyKind(t *testing.T) {
	// Test that Kind is populated when initially empty (simulates informer behavior)
	tmpl, err := ParseTemplate("{{ .Kind }}.{{ .Name }}.example.com")
	require.NoError(t, err)

	// Create object with empty TypeMeta (Kind == "")
	obj := &testObject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
	}

	// Kind should be empty initially
	assert.Empty(t, obj.GetObjectKind().GroupVersionKind().Kind)

	got, err := ExecTemplate(tmpl, obj)
	require.NoError(t, err)

	// Kind should now be populated via reflection
	assert.Equal(t, "testObject", obj.GetObjectKind().GroupVersionKind().Kind)
	assert.Equal(t, []string{"testObject.test.example.com"}, got)
}

func TestExecTemplatePreservesExistingKind(t *testing.T) {
	// Test that existing Kind is not overwritten
	tmpl, err := ParseTemplate("{{ .Kind }}.{{ .Name }}.example.com")
	require.NoError(t, err)

	obj := &testObject{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CustomKind",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
	}

	got, err := ExecTemplate(tmpl, obj)
	require.NoError(t, err)

	// Kind should remain unchanged
	assert.Equal(t, "CustomKind", obj.GetObjectKind().GroupVersionKind().Kind)
	assert.Equal(t, []string{"CustomKind.test.example.com"}, got)
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

type testObject struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

func (t *testObject) DeepCopyObject() runtime.Object {
	return &testObject{
		TypeMeta:   t.TypeMeta,
		ObjectMeta: *t.ObjectMeta.DeepCopy(),
	}
}

func TestExecTemplateExecutionError(t *testing.T) {
	tmpl, err := ParseTemplate("{{ call .Name }}")
	require.NoError(t, err)

	obj := &metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			Kind: "TestKind",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-name",
			Namespace: "default",
		},
	}

	_, err = ExecTemplate(tmpl, obj)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to apply template on TestKind default/test-name")
}

func TestCombineWithTemplatedEndpoints(t *testing.T) {
	// Create a dummy template for tests that need one
	dummyTemplate := template.Must(template.New("test").Parse("{{.Name}}"))

	annotationEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpoint("annotation.example.com", endpoint.RecordTypeA, "1.2.3.4"),
	}
	templatedEndpoints := []*endpoint.Endpoint{
		endpoint.NewEndpoint("template.example.com", endpoint.RecordTypeA, "5.6.7.8"),
	}

	successTemplateFunc := func() ([]*endpoint.Endpoint, error) {
		return templatedEndpoints, nil
	}
	errorTemplateFunc := func() ([]*endpoint.Endpoint, error) {
		return nil, errors.New("template error")
	}

	tests := []struct {
		name                  string
		endpoints             []*endpoint.Endpoint
		fqdnTemplate          *template.Template
		combineFQDNAnnotation bool
		templateFunc          func() ([]*endpoint.Endpoint, error)
		want                  []*endpoint.Endpoint
		wantErr               bool
	}{
		{
			name:         "nil template returns original endpoints",
			endpoints:    annotationEndpoints,
			fqdnTemplate: nil,
			templateFunc: successTemplateFunc,
			want:         annotationEndpoints,
		},
		{
			name:         "combine=false with existing endpoints returns original",
			endpoints:    annotationEndpoints,
			fqdnTemplate: dummyTemplate,
			templateFunc: successTemplateFunc,
			want:         annotationEndpoints,
		},
		{
			name:         "combine=false with empty endpoints returns templated",
			endpoints:    []*endpoint.Endpoint{},
			fqdnTemplate: dummyTemplate,
			templateFunc: successTemplateFunc,
			want:         templatedEndpoints,
		},
		{
			name:                  "combine=true appends templated to existing",
			endpoints:             annotationEndpoints,
			fqdnTemplate:          dummyTemplate,
			combineFQDNAnnotation: true,
			templateFunc:          successTemplateFunc,
			want:                  append(annotationEndpoints, templatedEndpoints...),
		},
		{
			name:                  "combine=true with empty endpoints returns templated",
			endpoints:             []*endpoint.Endpoint{},
			fqdnTemplate:          dummyTemplate,
			combineFQDNAnnotation: true,
			templateFunc:          successTemplateFunc,
			want:                  templatedEndpoints,
		},
		{
			name:         "template error is propagated",
			endpoints:    []*endpoint.Endpoint{},
			fqdnTemplate: dummyTemplate,
			templateFunc: errorTemplateFunc,
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "nil endpoints with combine=false returns templated",
			endpoints:    nil,
			fqdnTemplate: dummyTemplate,
			templateFunc: successTemplateFunc,
			want:         templatedEndpoints,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CombineWithTemplatedEndpoints(
				tt.endpoints,
				tt.fqdnTemplate,
				tt.combineFQDNAnnotation,
				tt.templateFunc,
			)
			if tt.wantErr {
				require.Error(t, err)
				require.ErrorContains(t, err, "failed to get endpoints from template")
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
