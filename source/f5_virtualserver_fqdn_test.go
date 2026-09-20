/*
Copyright 2026 The Kubernetes Authors.

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
	"encoding/json"
	"testing"

	f5 "github.com/F5Networks/k8s-bigip-ctlr/v2/config/apis/cis/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestF5VirtualServerFQDNTemplate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title              string
		virtualServer      f5.VirtualServer
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		expected           []*endpoint.Endpoint
	}{
		{
			title: "fqdn-template and target-template generate endpoint when no valid IP",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host: "www.example.com",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "",
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vs.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template generates endpoint when no valid IP",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host: "www.example.com",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "",
				},
			},
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vs.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "fqdn-template with combine adds template endpoint alongside IP-based endpoint",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "my-vs.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "www.example.com",
					Targets:          endpoint.Targets{"192.168.1.100"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        0,
					Labels:           endpoint.Labels{"resource": "f5-virtualserver/virtualserver/my-vs"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-target-template with combine adds template endpoint alongside IP-based endpoint",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			combine:            true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "my-vs.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "www.example.com",
					Targets:          endpoint.Targets{"192.168.1.100"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        0,
					Labels:           endpoint.Labels{"resource": "f5-virtualserver/virtualserver/my-vs"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-target-template without combine is ignored when IP-based endpoints exist",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			fqdnTargetTemplate: "{{.Name}}.example.com:lb.example.com",
			combine:            false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "www.example.com",
					Targets:          endpoint.Targets{"192.168.1.100"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        0,
					Labels:           endpoint.Labels{"resource": "f5-virtualserver/virtualserver/my-vs"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template with combine alongside HostAliases-based endpoints",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
					HostAliases:          []string{"alias.example.com"},
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "alias.example.com",
					Targets:          endpoint.Targets{"192.168.1.100"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        0,
					Labels:           endpoint.Labels{"resource": "f5-virtualserver/virtualserver/my-vs"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "my-vs.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "www.example.com",
					Targets:          endpoint.Targets{"192.168.1.100"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        0,
					Labels:           endpoint.Labels{"resource": "f5-virtualserver/virtualserver/my-vs"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template without combine is ignored when IP-based endpoints exist",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec: f5.VirtualServerSpec{
					Host:                 "www.example.com",
					VirtualServerAddress: "192.168.1.100",
				},
				Status: f5.CustomResourceStatus{
					VSAddress: "192.168.1.100",
					Status:    "OK",
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			combine:        false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "www.example.com",
					Targets:          endpoint.Targets{"192.168.1.100"},
					RecordType:       endpoint.RecordTypeA,
					RecordTTL:        0,
					Labels:           endpoint.Labels{"resource": "f5-virtualserver/virtualserver/my-vs"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template can reference .Kind",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec:   f5.VirtualServerSpec{Host: "www.example.com"},
				Status: f5.CustomResourceStatus{VSAddress: ""},
			},
			fqdnTemplate:   "{{.Kind | toLower}}.{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("virtualserver.my-vs.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template can reference .APIVersion",
			virtualServer: f5.VirtualServer{
				TypeMeta: metav1.TypeMeta{
					APIVersion: f5VirtualServerGVR.GroupVersion().String(),
					Kind:       "VirtualServer",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-vs",
					Namespace: defaultF5VirtualServerNamespace,
				},
				Spec:   f5.VirtualServerSpec{Host: "www.example.com"},
				Status: f5.CustomResourceStatus{VSAddress: ""},
			},
			fqdnTargetTemplate: `{{.Name}}.{{replace "/" "." .APIVersion}}.example.com:1.2.3.4`,
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-vs.cis.f5.com.v1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(f5VirtualServerGVR.GroupVersion(), &f5.VirtualServer{}, &f5.VirtualServerList{})
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)
			fakeKubernetesClient := fakeKube.NewClientset()

			vsJSON, err := json.Marshal(tt.virtualServer)
			require.NoError(t, err)
			vsObj := unstructured.Unstructured{}
			require.NoError(t, vsObj.UnmarshalJSON(vsJSON))

			_, err = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).
				Create(t.Context(), &vsObj, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewF5VirtualServerSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{
				Namespace:      defaultF5VirtualServerNamespace,
				LabelFilter:    labels.Everything(),
				TemplateEngine: templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
			})
			require.NoError(t, err)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(f5VirtualServerGVR).Namespace(defaultF5VirtualServerNamespace).
					List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := src.Endpoints(t.Context())
			assert.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
