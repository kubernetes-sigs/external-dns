/*
Copyright 2023 The Kubernetes Authors.

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

package source

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

type UnstructuredSuite struct {
	suite.Suite
}

func (suite *UnstructuredSuite) SetupTest() {
}

func TestUnstructuredSource(t *testing.T) {
	suite.Run(t, new(UnstructuredSuite))
	t.Run("Interface", testUnstructuredSourceImplementsSource)
	t.Run("Endpoints", testUnstructuredSourceEndpoints)
}

// testUnstructuredSourceImplementsSource tests that unstructuredSource
// is a valid Source.
func testUnstructuredSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(unstructuredSource))
}

// testUnstructuredSourceEndpoints tests various scenarios of using
// Unstructured source.
func testUnstructuredSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title                string
		registeredNamespace  string
		namespace            string
		registeredAPIVersion string
		apiVersion           string
		registeredKind       string
		kind                 string
		targetJsonPath       string
		hostNameJsonPath     string
		endpoints            []*endpoint.Endpoint
		expectEndpoints      bool
		expectError          bool
		annotationFilter     string
		labelFilter          string
		annotations          map[string]string
		labels               map[string]string
		setFn                func(*unstructured.Unstructured)
	}{

		{
			title:                "invalid api version",
			registeredAPIVersion: "v1",
			apiVersion:           "v2",
			registeredKind:       "ConfigMap",
			kind:                 "ConfigMap",
			targetJsonPath:       "{.data.ip-addr}",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     true,
		},
		{
			title:                "invalid kind",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "ConfigMap",
			kind:                 "FakeConfigMap",
			targetJsonPath:       "{.data.ip-addr}",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     true,
		},
		{
			title:                "endpoints from ConfigMap within a specific namespace",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "ConfigMap",
			kind:                 "ConfigMap",
			targetJsonPath:       "{.data.ip-addr}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedStringMap(
					obj.Object,
					map[string]string{"ip-addr": "1.2.3.4"},
					"data",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "endpoints from Pod within a specific namespace",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "podIP",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "endpoints from annotation on ConfigMap within a specific namespace",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "ConfigMap",
			kind:                 "ConfigMap",
			targetJsonPath:       "{.metadata.annotations.target}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
				"target":                "1.2.3.4",
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "endpoints from labels on Pod within a specific namespace",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.metadata.labels.target}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			labels: map[string]string{
				"target": "1.2.3.4",
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "no endpoints within a specific namespace",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "ConfigMap",
			kind:                 "ConfigMap",
			targetJsonPath:       "{.data.ip-addr}",
			namespace:            "foo",
			registeredNamespace:  "bar",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectEndpoints: false,
			expectError:     false,
		},

		{
			title:                "invalid api resource with no targets",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			expectEndpoints: false,
			expectError:     false,
		},
		{
			title:                "valid api gvk with single endpoint",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "podIP",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid api gvk with single CNAME endpoint",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "ConfigMap",
			kind:                 "ConfigMap",
			targetJsonPath:       "{.data.hostname}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"def.example.org"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeCNAME,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedStringMap(
					obj.Object,
					map[string]string{"hostname": "def.example.org"},
					"data",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with single endpoint",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with multiple endpoints from hostnames via a string slice derived from a JSONPath expression that evaluates to a comma-delimited string",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			hostNameJsonPath:     `{range .status.hostnames[*]}{@}.example.org,{end}`,
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
				{
					DNSName:    "def.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
				unstructured.SetNestedStringSlice(
					obj.Object,
					[]string{"abc", "def"},
					"status", "hostnames",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with multiple endpoints from hostnames from nested objects in status",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			hostNameJsonPath:     `{.status.nodes[*].name}`,
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
				{
					DNSName:    "def.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
				unstructured.SetNestedSlice(
					obj.Object,
					[]interface{}{
						map[string]interface{}{
							"name": "abc.example.org",
						},
						map[string]interface{}{
							"name": "def.example.org",
						},
					},
					"status", "nodes",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with multiple endpoints with multiple targets from nested objects in status",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.addrs}",
			hostNameJsonPath:     `{.status.nodes[*].name}`,
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4", "5.6.7.8"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
				{
					DNSName:    "abc-alias.example.org",
					Targets:    endpoint.Targets{"1.2.3.4", "5.6.7.8"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedStringSlice(
					obj.Object,
					[]string{"1.2.3.4", "5.6.7.8"},
					"status", "addrs",
				)
				unstructured.SetNestedSlice(
					obj.Object,
					[]interface{}{
						map[string]interface{}{
							"name": "abc.example.org",
						},
						map[string]interface{}{
							"name": "abc-alias.example.org",
						},
					},
					"status", "nodes",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with multiple endpoints from hostnames via string slice in status",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			hostNameJsonPath:     "{.status.hostnames}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
				{
					DNSName:    "def.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
				unstructured.SetNestedStringSlice(
					obj.Object,
					[]string{"abc.example.org", "def.example.org"},
					"status", "hostnames",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with multiple endpoints from host names via comma-delimited string in status",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			hostNameJsonPath:     "{.status.hostnames}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
				{
					DNSName:    "def.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
				unstructured.SetNestedField(
					obj.Object,
					"abc.example.org,def.example.org",
					"status", "hostnames",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with endpoint formatted using JSONPath expression",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			hostNameJsonPath:     "{.status.hostname}.example.org",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
				},
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
				unstructured.SetNestedField(
					obj.Object,
					"abc",
					"status", "hostname",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid cluster-scoped crd gvk with single endpoint",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyClusterScopedCRD",
			kind:                 "MyClusterScopedCRD",
			targetJsonPath:       "{.status.ipAddr}",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "ipAddr",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid namespace-scoped crd gvk with single endpoint from terminating indexed property",
			registeredAPIVersion: "example.com/v1alpha1",
			apiVersion:           "example.com/v1alpha1",
			registeredKind:       "MyNamespaceScopedCRD",
			kind:                 "MyNamespaceScopedCRD",
			targetJsonPath:       "{.status.ipAddrs[0]}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedStringSlice(
					obj.Object,
					[]string{"1.2.3.4", "5.6.7.8"},
					"status", "ipAddrs",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid api gvk with multiple endpoints",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "ConfigMap",
			kind:                 "ConfigMap",
			targetJsonPath:       "{.data.ip-addr}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
				{
					DNSName:    "def.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org,def.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedStringMap(
					obj.Object,
					map[string]string{"ip-addr": "1.2.3.4"},
					"data",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
		{
			title:                "valid api gvk with annotation and non matching annotation filter",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			annotationFilter:     "test=filter_something_else",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "podIP",
				)
			},
			expectEndpoints: false,
			expectError:     false,
		},

		{
			title:                "valid api with annotation and matching annotation filter",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			annotationFilter:     "test=filter_something_else",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			annotations: map[string]string{
				"test":                  "filter_something_else",
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "podIP",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},

		{
			title:                "valid api gvk with label and non matching label filter",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			labelFilter: "test=that",
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "podIP",
				)
			},
			expectEndpoints: false,
			expectError:     false,
		},
		{
			title:                "valid api gvk with label and matching label filter",
			registeredAPIVersion: "v1",
			apiVersion:           "v1",
			registeredKind:       "Pod",
			kind:                 "Pod",
			targetJsonPath:       "{.status.podIP}",
			namespace:            "foo",
			registeredNamespace:  "foo",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			labelFilter: "test=that",
			labels: map[string]string{
				"test": "that",
			},
			annotations: map[string]string{
				hostnameAnnotationKey:   "abc.example.org",
				ttlAnnotationKey:        "180",
				recordTypeAnnotationKey: endpoint.RecordTypeA,
			},
			setFn: func(obj *unstructured.Unstructured) {
				unstructured.SetNestedField(
					obj.Object,
					"1.2.3.4",
					"status", "podIP",
				)
			},
			expectEndpoints: true,
			expectError:     false,
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			groupVersion, err := schema.ParseGroupVersion(ti.apiVersion)
			require.NoError(t, err)

			scheme := runtime.NewScheme()
			scheme.AddKnownTypeWithName(
				groupVersion.WithKind(ti.kind),
				&unstructured.Unstructured{})
			scheme.AddKnownTypeWithName(
				groupVersion.WithKind(ti.kind+"List"),
				&unstructured.UnstructuredList{})

			labelSelector, err := labels.Parse(ti.labelFilter)
			require.NoError(t, err)

			fakeKubeClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)
			ctx := context.Background()

			fakeKubeClient.Resources = []*metav1.APIResourceList{
				{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "v1",
						Kind:       "APIResourceList",
					},
					GroupVersion: "v1",
					APIResources: []metav1.APIResource{
						{
							Kind:       "ConfigMap",
							Namespaced: true,
						},
						{
							Kind:       "Pod",
							Namespaced: true,
						},
					},
				},
				{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "v1",
						Kind:       "APIResourceList",
					},
					GroupVersion: "apps/v1",
					APIResources: []metav1.APIResource{
						{
							Kind:       "Deployment",
							Namespaced: true,
						},
					},
				},
				{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "v1",
						Kind:       "APIResourceList",
					},
					GroupVersion: "example.com/v1alpha1",
					APIResources: []metav1.APIResource{
						{
							Kind:       "MyNamespaceScopedCRD",
							Namespaced: true,
						},
						{
							Kind:       "MyClusterScopedCRD",
							Namespaced: false,
						},
					},
				},
			}

			// Create the object in the fake client.
			obj := &unstructured.Unstructured{Object: map[string]interface{}{}}
			obj.SetAPIVersion(ti.registeredAPIVersion)
			obj.SetKind(ti.registeredKind)
			obj.SetName("test")
			obj.SetNamespace(ti.registeredNamespace)
			obj.SetAnnotations(ti.annotations)
			obj.SetLabels(ti.labels)
			obj.SetGeneration(1)
			if ti.setFn != nil {
				ti.setFn(obj)
			}
			groupVersionResource := groupVersion.WithResource(strings.ToLower(ti.registeredKind) + "s")

			var createErr error
			if ti.namespace == "" {
				_, createErr = fakeDynamicClient.Resource(
					groupVersionResource).Create(ctx, obj, metav1.CreateOptions{})
			} else {
				_, createErr = fakeDynamicClient.Resource(
					groupVersionResource).Namespace(
					ti.registeredNamespace).Create(ctx, obj, metav1.CreateOptions{})
			}
			require.NoError(t, createErr)

			src, err := NewUnstructuredSource(
				ctx,
				fakeDynamicClient,
				fakeKubeClient,
				ti.namespace,
				ti.apiVersion,
				ti.kind,
				ti.targetJsonPath,
				ti.hostNameJsonPath,
				labelSelector,
				ti.annotationFilter)
			if ti.expectError {
				require.Nil(t, src)
				require.Error(t, err)
				return
			}
			require.NotNil(t, src)
			require.NoError(t, err)

			receivedEndpoints, err := src.Endpoints(ctx)
			if ti.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			if len(receivedEndpoints) == 0 && !ti.expectEndpoints {
				return
			}

			// Validate received endpoints against expected endpoints.
			validateEndpoints(t, receivedEndpoints, ti.endpoints)
		})
	}
}
