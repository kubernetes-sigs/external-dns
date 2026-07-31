/*
Copyright 2021 The Kubernetes Authors.

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

	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/types"
)

// This is a compile-time validation that kongTCPIngressSource is a Source.
var _ Source = &kongTCPIngressSource{}

const defaultKongNamespace = "kong"

func TestKongTCPIngressEndpoints(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title                    string
		tcpProxy                 TCPIngress
		labelFilter              labels.Selector
		ignoreHostnameAnnotation bool
		expected                 []*endpoint.Endpoint
	}{
		{
			title: "TCPIngress with hostname annotation",
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-annotation",
					Namespace: defaultKongNamespace,
					UID:       "kong-tcpingress-uid",
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "a.example.com",
						"kubernetes.io/ingress.class":         "kong",
					},
				},
				Spec: tcpIngressSpec{
					Rules: []tcpIngressRule{
						{
							Port: 30000,
						},
						{
							Port: 30001,
						},
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								Hostname: "a691234567a314e71861a4303f06a3bd-1291189659.us-east-1.elb.amazonaws.com",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				(&endpoint.Endpoint{
					DNSName:    "a.example.com",
					Targets:    []string{"a691234567a314e71861a4303f06a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-annotation",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				}).WithRefObject(testutils.RefSource(types.KongTCPIngress)),
			},
		},
		{
			title: "TCPIngress using SNI",
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-sni",
					Namespace: defaultKongNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class": "kong",
					},
				},
				Spec: tcpIngressSpec{
					Rules: []tcpIngressRule{
						{
							Port: 30002,
							Host: "b.example.com",
						},
						{
							Port: 30003,
							Host: "c.example.com",
						},
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								Hostname: "a123456769a314e71861a4303f06a3bd-1291189659.us-east-1.elb.amazonaws.com",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "b.example.com",
					Targets:    []string{"a123456769a314e71861a4303f06a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-sni",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:    "c.example.com",
					Targets:    []string{"a123456769a314e71861a4303f06a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-sni",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "TCPIngress with hostname annotation and using SNI",
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-both",
					Namespace: defaultKongNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "d.example.com",
						"kubernetes.io/ingress.class":         "kong",
					},
				},
				Spec: tcpIngressSpec{
					Rules: []tcpIngressRule{
						{
							Port: 30004,
							Host: "e.example.com",
						},
						{
							Port: 30005,
							Host: "f.example.com",
						},
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								Hostname: "a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "d.example.com",
					Targets:    []string{"a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-both",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:    "e.example.com",
					Targets:    []string{"a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-both",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:    "f.example.com",
					Targets:    []string{"a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-both",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "TCPIngress ignoring hostname annotation",
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-both",
					Namespace: defaultKongNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "d.example.com",
						"kubernetes.io/ingress.class":         "kong",
					},
				},
				Spec: tcpIngressSpec{
					Rules: []tcpIngressRule{
						{
							Port: 30004,
							Host: "e.example.com",
						},
						{
							Port: 30005,
							Host: "f.example.com",
						},
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								Hostname: "a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com",
							},
						},
					},
				},
			},
			ignoreHostnameAnnotation: true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "e.example.com",
					Targets:    []string{"a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-both",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:    "f.example.com",
					Targets:    []string{"a12e71861a4303f063456769a314a3bd-1291189659.us-east-1.elb.amazonaws.com"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-both",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title: "TCPIngress with target annotation",
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-sni",
					Namespace: defaultKongNamespace,
					Annotations: map[string]string{
						"kubernetes.io/ingress.class":       "kong",
						"external-dns.kubernetes.io/target": "203.2.45.7",
					},
				},
				Spec: tcpIngressSpec{
					Rules: []tcpIngressRule{
						{
							Port: 30002,
							Host: "b.example.com",
						},
						{
							Port: 30003,
							Host: "c.example.com",
						},
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								Hostname: "a123456769a314e71861a4303f06a3bd-1291189659.us-east-1.elb.amazonaws.com",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "b.example.com",
					Targets:    []string{"203.2.45.7"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-sni",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
				{
					DNSName:    "c.example.com",
					Targets:    []string{"203.2.45.7"},
					RecordType: endpoint.RecordTypeA,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-sni",
					},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title:       "TCPIngress with matching label filter",
			labelFilter: labels.SelectorFromSet(labels.Set{"app": "test"}),
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-label-match",
					Namespace: defaultKongNamespace,
					Labels:    map[string]string{"app": "test"},
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "label.example.com",
						"kubernetes.io/ingress.class":         "kong",
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{Hostname: "lb.example.com"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:          "label.example.com",
					Targets:          []string{"lb.example.com"},
					RecordType:       endpoint.RecordTypeCNAME,
					Labels:           endpoint.Labels{"resource": "tcpingress/kong/tcp-ingress-label-match"},
					ProviderSpecific: endpoint.ProviderSpecific{},
				},
			},
		},
		{
			title:       "TCPIngress with non-matching label filter",
			labelFilter: labels.SelectorFromSet(labels.Set{"app": "test"}),
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-label-no-match",
					Namespace: defaultKongNamespace,
					Labels:    map[string]string{"app": "other"},
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "label.example.com",
						"kubernetes.io/ingress.class":         "kong",
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{Hostname: "lb.example.com"},
						},
					},
				},
			},
			expected: nil,
		},
		{
			title: "TCPIngress with provider-specific annotation",
			tcpProxy: TCPIngress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: kongGroupdVersionResource.GroupVersion().String(),
					Kind:       "TCPIngress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tcp-ingress-provider-specific",
					Namespace: defaultKongNamespace,
					Annotations: map[string]string{
						"external-dns.kubernetes.io/hostname": "a.example.com",
						"kubernetes.io/ingress.class":         "kong",
						annotations.AWSPrefix + "weight":      "10",
					},
				},
				Status: tcpIngressStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{
							{
								IP: "1.2.3.4",
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "a.example.com",
					Targets:    []string{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  0,
					Labels: endpoint.Labels{
						"resource": "tcpingress/kong/tcp-ingress-provider-specific",
					},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/weight", Value: "10"},
					},
				},
			},
		},
	} {

		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			scheme := runtime.NewScheme()
			scheme.AddKnownTypes(kongGroupdVersionResource.GroupVersion(), &TCPIngress{}, &TCPIngressList{})
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(scheme)

			tcpi := unstructured.Unstructured{}

			tcpIngressAsJSON, err := json.Marshal(ti.tcpProxy)
			assert.NoError(t, err)

			assert.NoError(t, tcpi.UnmarshalJSON(tcpIngressAsJSON))

			// Create proxy resources
			_, err = fakeDynamicClient.Resource(kongGroupdVersionResource).Namespace(defaultKongNamespace).Create(t.Context(), &tcpi, metav1.CreateOptions{})
			assert.NoError(t, err)

			labelFilter := ti.labelFilter
			if labelFilter == nil {
				labelFilter = labels.Everything()
			}
			source, err := NewKongTCPIngressSource(t.Context(), fakeDynamicClient, fakeKubernetesClient,
				&Config{
					Namespaces:               []string{defaultKongNamespace},
					AnnotationFilter:         parseAnnotationFilterOrNil("kubernetes.io/ingress.class=kong"),
					IgnoreHostnameAnnotation: ti.ignoreHostnameAnnotation,
					LabelFilter:              labelFilter,
				})
			assert.NoError(t, err)
			assert.NotNil(t, source)

			count := &unstructured.UnstructuredList{}
			for len(count.Items) < 1 {
				count, _ = fakeDynamicClient.Resource(kongGroupdVersionResource).Namespace(defaultKongNamespace).List(t.Context(), metav1.ListOptions{})
			}

			endpoints, err := source.Endpoints(t.Context())
			assert.NoError(t, err)
			testutils.ValidateEndpoints(t, endpoints, ti.expected)
		})
	}
}

func TestKongTCPIngressSource_InformerTransform(t *testing.T) {
	t.Parallel()

	uc, err := newKongUnstructuredConverter()
	require.NoError(t, err)

	fakeClient := fakeKube.NewSimpleClientset()
	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

	source, err := NewKongTCPIngressSource(t.Context(), fakeDynamicClient, fakeClient, &Config{LabelFilter: labels.Everything()})
	require.NoError(t, err)
	require.IsType(t, &kongTCPIngressSource{}, source)

	testDynamicInformerTransformHelper(t,
		kongGroupdVersionResource,
		fakeDynamicClient,
		source.(*kongTCPIngressSource).kongTCPIngressInformer,
		withRemovedLastAppliedConfigAnnotation(),
		withRemovedManagedFields(),
	)
}

// TestKongTCPIngressIndexer verifies that the TCPIngress indexer correctly filters resources
// by annotation filter and label selector at index time, so that only matching resources are
// returned by Endpoints().
func TestKongTCPIngressIndexer(t *testing.T) {
	t.Parallel()

	makeEntity := func(name string, ann, lbls map[string]string) *TCPIngress {
		if ann == nil {
			ann = map[string]string{}
		}
		if lbls == nil {
			lbls = map[string]string{}
		}
		return &TCPIngress{
			TypeMeta: metav1.TypeMeta{
				APIVersion: kongGroupdVersionResource.GroupVersion().String(),
				Kind:       "TCPIngress",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        name,
				Namespace:   defaultKongNamespace,
				Annotations: ann,
				Labels:      lbls,
			},
			Spec: tcpIngressSpec{
				Rules: []tcpIngressRule{{Host: name + ".example.org", Port: 80}},
			},
			Status: tcpIngressStatus{
				LoadBalancer: corev1.LoadBalancerStatus{
					Ingress: []corev1.LoadBalancerIngress{{IP: "1.2.3.4"}},
				},
			},
		}
	}

	tests := []struct {
		name             string
		annotationFilter string
		labelFilter      string
		ingresses        []*TCPIngress
		expectedCount    int
	}{
		{
			name:          "no filters returns all ingresses",
			expectedCount: 3,
			ingresses: []*TCPIngress{
				makeEntity("ti1", nil, nil),
				makeEntity("ti2", nil, nil),
				makeEntity("ti3", nil, nil),
			},
		},
		{
			name:             "annotation filter includes matching ingresses",
			annotationFilter: "tier=frontend",
			expectedCount:    2,
			ingresses: []*TCPIngress{
				makeEntity("ti1", map[string]string{"tier": "frontend"}, nil),
				makeEntity("ti2", map[string]string{"tier": "frontend"}, nil),
				makeEntity("ti3", map[string]string{"tier": "backend"}, nil),
			},
		},
		{
			name:          "label filter includes matching ingresses",
			labelFilter:   "env=prod",
			expectedCount: 1,
			ingresses: []*TCPIngress{
				makeEntity("ti1", nil, map[string]string{"env": "prod"}),
				makeEntity("ti2", nil, map[string]string{"env": "staging"}),
				makeEntity("ti3", nil, nil),
			},
		},
		{
			name:             "annotation and label filter combined",
			annotationFilter: "tier=frontend",
			labelFilter:      "env=prod",
			expectedCount:    1,
			ingresses: []*TCPIngress{
				makeEntity("ti1", map[string]string{"tier": "frontend"}, map[string]string{"env": "prod"}),
				makeEntity("ti2", map[string]string{"tier": "frontend"}, map[string]string{"env": "staging"}),
				makeEntity("ti3", map[string]string{"tier": "backend"}, map[string]string{"env": "prod"}),
			},
		},
		{
			name:             "no matches returns empty",
			annotationFilter: "tier=missing",
			expectedCount:    0,
			ingresses: []*TCPIngress{
				makeEntity("ti1", map[string]string{"tier": "frontend"}, nil),
			},
		},
		{
			name:          "controller mismatch is excluded",
			expectedCount: 0,
			ingresses: []*TCPIngress{
				makeEntity("ti1", map[string]string{annotations.ControllerKey: "other-controller"}, nil),
			},
		},
	}

	uc, err := newKongUnstructuredConverter()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(uc.scheme)

			for _, ing := range tt.ingresses {
				data, err := json.Marshal(ing)
				require.NoError(t, err)
				obj := unstructured.Unstructured{}
				require.NoError(t, obj.UnmarshalJSON(data))
				_, err = fakeDynamicClient.Resource(kongGroupdVersionResource).Namespace(defaultKongNamespace).Create(t.Context(), &obj, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			src, err := NewKongTCPIngressSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{
				Namespaces:       []string{defaultKongNamespace},
				AnnotationFilter: parseAnnotationFilterOrNil(tt.annotationFilter),
				LabelFilter:      parseLabelSelectorOrEverything(t, tt.labelFilter),
			})
			require.NoError(t, err)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Len(t, endpoints, tt.expectedCount)
		})
	}
}
