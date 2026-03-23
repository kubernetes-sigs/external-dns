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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestNewEngine(t *testing.T) {
	for _, tt := range []struct {
		name        string
		fqdn        string
		target      string
		fqdnTarget  string
		errContains string
	}{
		{
			name:        "invalid fqdn template",
			fqdn:        "{{.Name",
			errContains: `parse --fqdn-template: "{{.Name"`,
		},
		{
			name: "empty fqdn template",
		},
		{
			name: "valid fqdn template",
			fqdn: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			name: "valid fqdn template with multiple hosts",
			fqdn: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
		{
			name: "replace template function",
			fqdn: "{{\"hello.world\" | replace \".\" \"-\"}}.ext-dns.test.com",
		},
		{
			name: "isIPv4 template function with valid IPv4",
			fqdn: "{{if isIPv4 \"192.168.1.1\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name: "isIPv4 template function with invalid IPv4",
			fqdn: "{{if isIPv4 \"not.an.ip.addr\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name: "isIPv6 template function with valid IPv6",
			fqdn: "{{if isIPv6 \"2001:db8::1\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name: "isIPv6 template function with invalid IPv6",
			fqdn: "{{if isIPv6 \"not:ipv6:addr\"}}valid{{else}}invalid{{end}}.ext-dns.test.com",
		},
		{
			name:        "invalid target template",
			target:      "{{.Status.LoadBalancer.Ingress",
			errContains: `parse --target-template: "{{.Status.LoadBalancer.Ingress"`,
		},
		{
			name:   "valid target template",
			target: "{{.Name}}.targets.example.com",
		},
		{
			name:        "invalid fqdn-target template",
			fqdnTarget:  "{{.Name}}.example.com:{{.Status",
			errContains: `parse --fqdn-target-template: "{{.Name}}.example.com:{{.Status"`,
		},
		{
			name:       "valid fqdn-target template",
			fqdnTarget: "{{.Name}}.example.com:{{.Name}}.targets.example.com",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEngine(tt.fqdn, tt.target, tt.fqdnTarget, false)
			if tt.errContains != "" {
				assert.ErrorContains(t, err, tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTemplateEngineIsConfigured(t *testing.T) {
	empty, err := NewEngine("", "", "", false)
	require.NoError(t, err)
	assert.False(t, empty.IsConfigured())

	configured, err := NewEngine("{{ .Name }}.example.com", "", "", false)
	require.NoError(t, err)
	assert.True(t, configured.IsConfigured())
}

func TestExecFQDN(t *testing.T) {
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
			engine, err := NewEngine(tt.tmpl, "", "", false)
			require.NoError(t, err)

			got, err := engine.ExecFQDN(tt.obj)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExecFQDNNilObject(t *testing.T) {
	engine, err := NewEngine("{{ toLower .Labels.department }}.example.org", "", "", false)
	require.NoError(t, err)
	_, err = engine.ExecFQDN(nil)
	assert.Error(t, err)
}

func TestExecFQDNPopulatesEmptyKind(t *testing.T) {
	// Test that Kind is populated when initially empty (simulates informer behavior)
	engine, err := NewEngine("{{ .Kind }}.{{ .Name }}.example.com", "", "", false)
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

	got, err := engine.ExecFQDN(obj)
	require.NoError(t, err)

	// Kind should now be populated via reflection
	assert.Equal(t, "testObject", obj.GetObjectKind().GroupVersionKind().Kind)
	assert.Equal(t, []string{"testObject.test.example.com"}, got)
}

func TestExecFQDNPreservesExistingKind(t *testing.T) {
	// Test that existing Kind is not overwritten
	engine, err := NewEngine("{{ .Kind }}.{{ .Name }}.example.com", "", "", false)
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

	got, err := engine.ExecFQDN(obj)
	require.NoError(t, err)

	// Kind should remain unchanged
	assert.Equal(t, "CustomKind", obj.GetObjectKind().GroupVersionKind().Kind)
	assert.Equal(t, []string{"CustomKind.test.example.com"}, got)
}

func TestExecFQDNExecutionError(t *testing.T) {
	engine, err := NewEngine("{{ call .Name }}", "", "", false)
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

	_, err = engine.ExecFQDN(obj)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to apply template on TestKind default/test-name")
}

func TestCombineWithEndpoints(t *testing.T) {
	configured, err := NewEngine("{{.Name}}", "", "", false)
	require.NoError(t, err)
	configuredCombine, err := NewEngine("{{.Name}}", "", "", true)
	require.NoError(t, err)
	unconfigured, err := NewEngine("", "", "", false)
	require.NoError(t, err)

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
		name         string
		endpoints    []*endpoint.Endpoint
		engine       Engine
		templateFunc func() ([]*endpoint.Endpoint, error)
		want         []*endpoint.Endpoint
		wantErr      bool
	}{
		{
			name:         "unconfigured engine returns original endpoints",
			endpoints:    annotationEndpoints,
			engine:       unconfigured,
			templateFunc: successTemplateFunc,
			want:         annotationEndpoints,
		},
		{
			name:         "combine=false with existing endpoints returns original",
			endpoints:    annotationEndpoints,
			engine:       configured,
			templateFunc: successTemplateFunc,
			want:         annotationEndpoints,
		},
		{
			name:         "combine=false with empty endpoints returns templated",
			endpoints:    []*endpoint.Endpoint{},
			engine:       configured,
			templateFunc: successTemplateFunc,
			want:         templatedEndpoints,
		},
		{
			name:         "combine=true appends templated to existing",
			endpoints:    annotationEndpoints,
			engine:       configuredCombine,
			templateFunc: successTemplateFunc,
			want:         append(annotationEndpoints, templatedEndpoints...),
		},
		{
			name:         "combine=true with empty endpoints returns templated",
			endpoints:    []*endpoint.Endpoint{},
			engine:       configuredCombine,
			templateFunc: successTemplateFunc,
			want:         templatedEndpoints,
		},
		{
			name:         "template error is propagated",
			endpoints:    []*endpoint.Endpoint{},
			engine:       configured,
			templateFunc: errorTemplateFunc,
			want:         nil,
			wantErr:      true,
		},
		{
			name:         "nil endpoints with combine=false returns templated",
			endpoints:    nil,
			engine:       configured,
			templateFunc: successTemplateFunc,
			want:         templatedEndpoints,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.engine.CombineWithEndpoints(
				tt.endpoints,
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
