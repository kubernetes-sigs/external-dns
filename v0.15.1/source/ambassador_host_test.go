/*
Copyright 2019 The Kubernetes Authors.

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
	"fmt"
	"testing"

	ambassador "github.com/datawire/ambassador/pkg/api/getambassador.io/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
)

const defaultAmbassadorNamespace = "ambassador"
const defaultAmbassadorServiceName = "ambassador"

// This is a compile-time validation that ambassadorHostSource is a Source.
var _ Source = &ambassadorHostSource{}

type AmbassadorSuite struct {
	suite.Suite
}

func TestAmbassadorSource(t *testing.T) {
	suite.Run(t, new(AmbassadorSuite))
	t.Run("Interface", testAmbassadorSourceImplementsSource)
}

// testAmbassadorSourceImplementsSource tests that ambassadorHostSource is a valid Source.
func testAmbassadorSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(ambassadorHostSource))
}

func TestAmbassadorHostSource(t *testing.T) {
	t.Parallel()

	hostAnnotation := fmt.Sprintf("%s/%s", defaultAmbassadorNamespace, defaultAmbassadorServiceName)

	for _, ti := range []struct {
		title            string
		annotationFilter string
		labelSelector    labels.Selector
		host             ambassador.Host
		service          v1.Service
		expected         []*endpoint.Endpoint
	}{
		{
			title:         "Simple host",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation: hostAnnotation,
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1"},
				},
			},
		}, {
			title:         "Service with load balancer hostname",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation: hostAnnotation,
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							Hostname: "dns.google",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"dns.google"},
				},
			},
		}, {
			title:         "Service with external IP",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "service-external-ip",
					Annotations: map[string]string{
						ambHostAnnotation: hostAnnotation,
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Spec: v1.ServiceSpec{
					ExternalIPs: []string{"2.2.2.2"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"2.2.2.2"},
				},
			},
		}, {
			title:         "Host with target annotation",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:   hostAnnotation,
						targetAnnotationKey: "3.3.3.3",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"3.3.3.3"},
				},
			},
		}, {
			title:         "Host with TTL annotation",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation: hostAnnotation,
						ttlAnnotationKey:  "180",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1"},
					RecordTTL:  180,
				},
			},
		}, {
			title:         "Host with provider specific annotation",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:    hostAnnotation,
						CloudflareProxiedKey: "true",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1"},
					ProviderSpecific: endpoint.ProviderSpecific{{
						Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
						Value: "true",
					}},
				},
			},
		}, {
			title:         "Host with missing Ambassador annotation",
			labelSelector: labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		}, {
			title:            "valid matching annotation filter expression",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1"},
				},
			},
		}, {
			title:            "valid non-matching annotation filter expression",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		}, {
			title:            "invalid annotation filter expression",
			annotationFilter: "kubernetes.io/ingress.class in (invalid-ingress)",
			labelSelector:    labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		}, {
			title:            "valid non-matching annotation filter label",
			annotationFilter: "kubernetes.io/ingress.class=external-ingress",
			labelSelector:    labels.Everything(),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:         "valid non-matching label filter expression",
			labelSelector: labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "internal-ingress",
					},
					Labels: map[string]string{
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP: "1.1.1.1",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:         "valid matching label filter expression for single host",
			labelSelector: labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation: hostAnnotation,
					},
					Labels: map[string]string{
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP:       "1.1.1.1",
							Hostname: "dns.google",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1"},
				},
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"dns.google"},
				},
			},
		},
		{
			title:            "valid matching label filter expression and matching annotation filter",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "external-ingress",
					},
					Labels: map[string]string{
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP:       "1.1.1.1",
							Hostname: "dns.google",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1"},
				},
				{
					DNSName:    "www.example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"dns.google"},
				},
			},
		},
		{
			title:            "valid non matching label filter expression and valid matching annotation filter",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "external-ingress",
					},
					Labels: map[string]string{
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP:       "1.1.1.1",
							Hostname: "dns.google",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "valid matching label filter expression and non matching annotation filter",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			host: ambassador.Host{
				ObjectMeta: metav1.ObjectMeta{
					Name: "basic-host",
					Annotations: map[string]string{
						ambHostAnnotation:             hostAnnotation,
						"kubernetes.io/ingress.class": "internal-ingress",
					},
					Labels: map[string]string{
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
				Spec: &ambassador.HostSpec{
					Hostname: "www.example.org",
				},
			},
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaultAmbassadorServiceName,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{{
							IP:       "1.1.1.1",
							Hostname: "dns.google",
						}},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			fakeKubernetesClient := fakeKube.NewSimpleClientset()
			ambassadorScheme := runtime.NewScheme()
			ambassador.AddToScheme(ambassadorScheme)
			fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(ambassadorScheme)

			namespace := "default"

			// Create Ambassador service
			_, err := fakeKubernetesClient.CoreV1().Services(defaultAmbassadorNamespace).Create(context.Background(), &ti.service, metav1.CreateOptions{})
			assert.NoError(t, err)

			// Create host resource
			host, err := createAmbassadorHost(&ti.host)
			assert.NoError(t, err)

			_, err = fakeDynamicClient.Resource(ambHostGVR).Namespace(namespace).Create(context.Background(), host, metav1.CreateOptions{})
			assert.NoError(t, err)

			source, err := NewAmbassadorHostSource(context.TODO(), fakeDynamicClient, fakeKubernetesClient, namespace, ti.annotationFilter, ti.labelSelector)
			assert.NoError(t, err)
			assert.NotNil(t, source)

			endpoints, err := source.Endpoints(context.Background())
			assert.NoError(t, err)
			// Validate returned endpoints against expected endpoints.
			validateEndpoints(t, endpoints, ti.expected)
		})
	}
}

func createAmbassadorHost(host *ambassador.Host) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	uc, _ := newUnstructuredConverter()
	err := uc.scheme.Convert(host, obj, nil)

	return obj, err
}

// TestParseAmbLoadBalancerService tests our parsing of Ambassador service info.
func TestParseAmbLoadBalancerService(t *testing.T) {
	vectors := []struct {
		input  string
		ns     string
		svc    string
		errstr string
	}{
		{"svc", "default", "svc", ""},
		{"ns/svc", "ns", "svc", ""},
		{"svc.ns", "ns", "svc", ""},
		{"svc.ns.foo.bar", "ns.foo.bar", "svc", ""},
		{"ns/svc/foo/bar", "", "", "invalid external-dns service: ns/svc/foo/bar"},
		{"ns/svc/foo.bar", "", "", "invalid external-dns service: ns/svc/foo.bar"},
		{"ns.foo/svc/bar", "", "", "invalid external-dns service: ns.foo/svc/bar"},
	}

	for _, v := range vectors {
		ns, svc, err := parseAmbLoadBalancerService(v.input)

		errstr := ""

		if err != nil {
			errstr = err.Error()
		}
		if v.ns != ns {
			t.Errorf("%s: got ns \"%s\", wanted \"%s\"", v.input, ns, v.ns)
		}

		if v.svc != svc {
			t.Errorf("%s: got svc \"%s\", wanted \"%s\"", v.input, svc, v.svc)
		}

		if v.errstr != errstr {
			t.Errorf("%s: got err \"%s\", wanted \"%s\"", v.input, errstr, v.errstr)
		}
	}
}
