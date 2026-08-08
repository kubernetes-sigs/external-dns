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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestTraefikFQDNTemplateIngressRoute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title              string
		ingressRoute       IngressRoute
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		expected           []*endpoint.Endpoint
	}{
		{
			title: "fqdn-template and target-template generate endpoint when no route spec",
			ingressRoute: IngressRoute{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteGVR.GroupVersion().String(),
					Kind:       "IngressRoute",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-app.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template generates endpoint when no route spec",
			ingressRoute: IngressRoute{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteGVR.GroupVersion().String(),
					Kind:       "IngressRoute",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-app.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "fqdn-template with combine adds template endpoint alongside route-based endpoint",
			ingressRoute: IngressRoute{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteGVR.GroupVersion().String(),
					Kind:       "IngressRoute",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/target": "lb.example.com",
						"kubernetes.io/ingress.class":       "traefik",
					},
				},
				Spec: traefikIngressRouteSpec{
					Routes: []traefikRoute{
						{Match: "Host(`route.example.com`)"},
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "template-lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "my-app.example.com",
					Targets:          endpoint.Targets{"template-lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "route.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{"resource": "ingressroute/traefik/my-app"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template without combine is ignored when route-based endpoints exist",
			ingressRoute: IngressRoute{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteGVR.GroupVersion().String(),
					Kind:       "IngressRoute",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/target": "lb.example.com",
						"kubernetes.io/ingress.class":       "traefik",
					},
				},
				Spec: traefikIngressRouteSpec{
					Routes: []traefikRoute{
						{Match: "Host(`route.example.com`)"},
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "template-lb.example.com",
			combine:        false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "route.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{"resource": "ingressroute/traefik/my-app"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template can reference .Kind",
			ingressRoute: IngressRoute{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteGVR.GroupVersion().String(),
					Kind:       "IngressRoute",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Kind | toLower}}.{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("ingressroute.my-app.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template can reference .APIVersion",
			ingressRoute: IngressRoute{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteGVR.GroupVersion().String(),
					Kind:       "IngressRoute",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTargetTemplate: `{{.Name}}.{{replace "/" "." .APIVersion}}.example.com:1.2.3.4`,
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-app.traefik.io.v1alpha1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			uc, err := newTraefikUnstructuredConverter()
			require.NoError(t, err)

			fakeKubeClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

			obj := &unstructured.Unstructured{}
			require.NoError(t, uc.scheme.Convert(&tt.ingressRoute, obj, t.Context()))

			_, err = fakeDynamicClient.Resource(ingressRouteGVR).Namespace(defaultTraefikNamespace).
				Create(t.Context(), obj, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewTraefikSource(t.Context(), fakeDynamicClient, fakeKubeClient, &Config{
				Namespace:        defaultTraefikNamespace,
				AnnotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class=traefik"),
				LabelFilter:      labels.Everything(),
				TemplateEngine:   templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
			})
			require.NoError(t, err)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(ingressRouteGVR).Namespace(defaultTraefikNamespace).
					List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := src.Endpoints(t.Context())
			assert.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestTraefikFQDNTemplateIngressRouteTCP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title              string
		ingressRouteTCP    IngressRouteTCP
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		expected           []*endpoint.Endpoint
	}{
		{
			title: "fqdn-template and target-template generate endpoint when no HostSNI rule",
			ingressRouteTCP: IngressRouteTCP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteTCPGVR.GroupVersion().String(),
					Kind:       "IngressRouteTCP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-tcp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-tcp-app.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template generates endpoint when no HostSNI rule",
			ingressRouteTCP: IngressRouteTCP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteTCPGVR.GroupVersion().String(),
					Kind:       "IngressRouteTCP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-tcp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-tcp-app.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "fqdn-template with combine adds template endpoint alongside HostSNI-based endpoint",
			ingressRouteTCP: IngressRouteTCP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteTCPGVR.GroupVersion().String(),
					Kind:       "IngressRouteTCP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-tcp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/target": "lb.example.com",
						"kubernetes.io/ingress.class":       "traefik",
					},
				},
				Spec: traefikIngressRouteTCPSpec{
					Routes: []traefikRouteTCP{
						{Match: "HostSNI(`sni.example.com`)"},
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "template-lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "my-tcp-app.example.com",
					Targets:          endpoint.Targets{"template-lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "sni.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{"resource": "ingressroutetcp/traefik/my-tcp-app"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template can reference .Kind",
			ingressRouteTCP: IngressRouteTCP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteTCPGVR.GroupVersion().String(),
					Kind:       "IngressRouteTCP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-tcp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Kind | toLower}}.{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("ingressroutetcp.my-tcp-app.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template can reference .APIVersion",
			ingressRouteTCP: IngressRouteTCP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteTCPGVR.GroupVersion().String(),
					Kind:       "IngressRouteTCP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-tcp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTargetTemplate: `{{.Name}}.{{replace "/" "." .APIVersion}}.example.com:1.2.3.4`,
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-tcp-app.traefik.io.v1alpha1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			uc, err := newTraefikUnstructuredConverter()
			require.NoError(t, err)

			fakeKubeClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

			obj := &unstructured.Unstructured{}
			require.NoError(t, uc.scheme.Convert(&tt.ingressRouteTCP, obj, t.Context()))

			_, err = fakeDynamicClient.Resource(ingressRouteTCPGVR).Namespace(defaultTraefikNamespace).
				Create(t.Context(), obj, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewTraefikSource(t.Context(), fakeDynamicClient, fakeKubeClient, &Config{
				Namespace:        defaultTraefikNamespace,
				AnnotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class=traefik"),
				LabelFilter:      labels.Everything(),
				TemplateEngine:   templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
			})
			require.NoError(t, err)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(ingressRouteTCPGVR).Namespace(defaultTraefikNamespace).
					List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := src.Endpoints(t.Context())
			assert.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestTraefikFQDNTemplateIngressRouteUDP(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title              string
		ingressRouteUDP    IngressRouteUDP
		fqdnTemplate       string
		targetTemplate     string
		fqdnTargetTemplate string
		combine            bool
		expected           []*endpoint.Endpoint
	}{
		{
			title: "fqdn-template and target-template generate endpoint",
			ingressRouteUDP: IngressRouteUDP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteUDPGVR.GroupVersion().String(),
					Kind:       "IngressRouteUDP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-udp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-udp-app.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template generates endpoint",
			ingressRouteUDP: IngressRouteUDP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteUDPGVR.GroupVersion().String(),
					Kind:       "IngressRouteUDP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-udp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTargetTemplate: "{{.Name}}.example.com:1.2.3.4",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-udp-app.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title: "fqdn-template with combine adds template endpoint alongside annotation-based endpoint",
			ingressRouteUDP: IngressRouteUDP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteUDPGVR.GroupVersion().String(),
					Kind:       "IngressRouteUDP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-udp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "annotated.example.com",
						"external-dns.kubernetes.io/target":   "lb.example.com",
						"kubernetes.io/ingress.class":         "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Name}}.example.com",
			targetTemplate: "template-lb.example.com",
			combine:        true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "annotated.example.com",
					Targets:          endpoint.Targets{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{"resource": "ingressrouteudp/traefik/my-udp-app"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:          "my-udp-app.example.com",
					Targets:          endpoint.Targets{"template-lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "fqdn-template can reference .Kind",
			ingressRouteUDP: IngressRouteUDP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteUDPGVR.GroupVersion().String(),
					Kind:       "IngressRouteUDP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-udp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTemplate:   "{{.Kind | toLower}}.{{.Name}}.example.com",
			targetTemplate: "lb.example.com",
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("ingressrouteudp.my-udp-app.example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			title: "fqdn-target-template can reference .APIVersion",
			ingressRouteUDP: IngressRouteUDP{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ingressRouteUDPGVR.GroupVersion().String(),
					Kind:       "IngressRouteUDP",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-udp-app",
					Namespace: defaultTraefikNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "traefik",
					},
				},
			},
			fqdnTargetTemplate: `{{.Name}}.{{replace "/" "." .APIVersion}}.example.com:1.2.3.4`,
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("my-udp-app.traefik.io.v1alpha1.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			uc, err := newTraefikUnstructuredConverter()
			require.NoError(t, err)

			fakeKubeClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

			obj := &unstructured.Unstructured{}
			require.NoError(t, uc.scheme.Convert(&tt.ingressRouteUDP, obj, t.Context()))

			_, err = fakeDynamicClient.Resource(ingressRouteUDPGVR).Namespace(defaultTraefikNamespace).
				Create(t.Context(), obj, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewTraefikSource(t.Context(), fakeDynamicClient, fakeKubeClient, &Config{
				Namespace:        defaultTraefikNamespace,
				AnnotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class=traefik"),
				LabelFilter:      labels.Everything(),
				TemplateEngine:   templatetest.MustEngine(t, tt.fqdnTemplate, tt.targetTemplate, tt.fqdnTargetTemplate, tt.combine),
			})
			require.NoError(t, err)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(ingressRouteUDPGVR).Namespace(defaultTraefikNamespace).
					List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := src.Endpoints(t.Context())
			assert.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, tt.expected)
		})
	}
}
