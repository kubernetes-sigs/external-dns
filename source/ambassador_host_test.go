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
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
)

// This is a compile-time validation that ambassadorHostSource is a Source.
var _ Source = &ambassadorHostSource{}

type AmbassadorSuite struct {
	suite.Suite
}

func TestAmbassadorSource(t *testing.T) {
	suite.Run(t, new(AmbassadorSuite))
	t.Run("Interface", testAmbassadorSourceImplementsSource)
	t.Run("Endpoints", testAmbassadorSourceEndpoints)
}

// testAmbassadorSourceImplementsSource tests that ambassadorHostSource is a valid Source.
func testAmbassadorSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(ambassadorHostSource))
}

func TestAmbassadorHostSource(t *testing.T) {
	fakeKubernetesClient := fake.NewSimpleClientset()

	ambassadorScheme := runtime.NewScheme()

	ambassador.AddToScheme(ambassadorScheme)

	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClient(ambassadorScheme)

	ctx := context.Background()

	namespace := "test"

	annotationFilter := ""

	labelSelector :=   labels.Everything()

	host, err := createAmbassadorHost("test-host", "test-service")
	if err != nil {
		t.Fatalf("could not create host resource: %v", err)
	}

	{
		_, err := fakeDynamicClient.Resource(ambHostGVR).Namespace(namespace).Create(ctx, host, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("could not create host: %v", err)
		}
	}

	ambassadorSource, err := NewAmbassadorHostSource(ctx, fakeDynamicClient, fakeKubernetesClient, namespace, annotationFilter,labelSelector)
	if err != nil {
		t.Fatalf("could not create ambassador source: %v", err)
	}

	{
		_, err := ambassadorSource.Endpoints(ctx)
		if err != nil {
			t.Fatalf("could not collect ambassador source endpoints: %v", err)
		}
	}
}

func createAmbassadorHost(name, ambassadorService string) (*unstructured.Unstructured, error) {
	host := &ambassador.Host{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				ambHostAnnotation: ambassadorService,
			},
		},
	}
	obj := &unstructured.Unstructured{}
	uc, _ := newUnstructuredConverter()
	err := uc.scheme.Convert(host, obj, nil)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// testAmbassadorSourceEndpoints tests that various Ambassador Hosts generate the correct endpoints.
func testAmbassadorSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title               string
		targetNamespace     string
		annotationFilter    string
		loadBalancer        fakeAmbassadorLoadBalancerService
		ambassadorHostItems []fakeAmbassadorHost
		expected            []*endpoint.Endpoint
		expectError         bool
		labelSelector       labels.Selector
	}{
		{
			title: "no host",
			labelSelector:   labels.Everything(),
		},
		{
			title:         "two simple hosts",
			labelSelector: labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
				},
				{
					name:      "fake2",
					namespace: "",
					hostname:  "fake2.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
				{
					DNSName: "fake2.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake2.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:         "two simple hosts on different namespaces",
			labelSelector: labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					hostname:  "fake2.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
				{
					DNSName: "fake2.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake2.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:           "two simple hosts on different namespaces and a target namespace",
			targetNamespace: "testing1",
			labelSelector:   labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					hostname:  "fake2.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:         "invalid non matching host ambassador service annotation",
			labelSelector: labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/invalid",
					},
				},
			},
			expected:    []*endpoint.Endpoint{},
		},
		{
			title:            "valid matching annotation filter expression",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "external-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:            "valid non-matching annotation filter expression",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "internal-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "invalid annotation filter expression",
			annotationFilter: "kubernetes.io/ingress.class in (external ingress)",
			labelSelector:    labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "internal-ingress",
					},
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: true,
		},
		{
			title:            "valid matching annotation filter label",
			annotationFilter: "kubernetes.io/ingress.class=external-ingress",
			labelSelector:    labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "external-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:            "valid non-matching annotation filter label",
			annotationFilter: "kubernetes.io/ingress.class=external-ingress",
			labelSelector:    labels.Everything(),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "internal-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "valid non-matching label filter expression",
			labelSelector: labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
					labels: map[string]string{
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "valid matching label filter expression for single host",
			labelSelector: labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
					},
					labels: map[string]string{
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					hostname:  "fake2.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "internal-ingress",
					},
					labels: map[string]string{
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:            "valid matching label filter expression and matching annotation filter",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "external-ingress",
					},
					labels: map[string]string{
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:            "valid non matching label filter expression and valid matching annotation filter",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "external-ingress",
					},
					labels: map[string]string{
						"kubernetes.io/ingress.class": "internal-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "valid matching label filter expression and non matching annotation filter",
			annotationFilter: "kubernetes.io/ingress.class in (external-ingress)",
			labelSelector:    labels.SelectorFromSet(labels.Set{"kubernetes.io/ingress.class": "external-ingress"}),
			loadBalancer: fakeAmbassadorLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
				name:      "emissary",
				namespace: "emissary-ingress",
			},
			ambassadorHostItems: []fakeAmbassadorHost{
				{
					name:      "fake1",
					namespace: "testing1",
					hostname:  "fake1.org",
					annotations: map[string]string{
						"external-dns.ambassador-service": "emissary-ingress/emissary",
						"kubernetes.io/ingress.class":     "internal-ingress",
					},
					labels: map[string]string{
						"kubernetes.io/ingress.class": "external-ingress",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			// Create a slice of Ambassador Hosts
			ambassadorHosts := []*ambassador.Host{}

			// Convert our test data into Ambassador Hosts
			for _, host := range ti.ambassadorHostItems {
				ambassadorHosts = append(ambassadorHosts, host.Host())
			}

			fakeKubernetesClient := fake.NewSimpleClientset()
			service := ti.loadBalancer.Service()
			_, err := fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			fakeDynamicClient, scheme := newAmbassadorDynamicKubernetesClient()

			for _, host := range ambassadorHosts {
				converted, err := convertAmbassadorHostToUnstructured(host, scheme)
				require.NoError(t, err)
				_, err = fakeDynamicClient.Resource(ambHostGVR).Namespace(host.Namespace).Create(context.Background(), converted, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			ambassadorHostSource, err := NewAmbassadorHostSource(
				context.TODO(),
				fakeDynamicClient,
				fakeKubernetesClient,
				ti.targetNamespace,
				ti.annotationFilter,
				ti.labelSelector,
			)
			require.NoError(t, err)

			res, err := ambassadorHostSource.Endpoints(context.Background())
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			validateEndpoints(t, res, ti.expected)
		})
	}
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

func convertAmbassadorHostToUnstructured(hp *ambassador.Host, s *runtime.Scheme) (*unstructured.Unstructured, error) {
	unstructuredAmbassadorHost := &unstructured.Unstructured{}
	if err := s.Convert(hp, unstructuredAmbassadorHost, context.Background()); err != nil {
		return nil, err
	}
	return unstructuredAmbassadorHost, nil
}

func newAmbassadorDynamicKubernetesClient() (*fakeDynamic.FakeDynamicClient, *runtime.Scheme) {
	s := runtime.NewScheme()
	_ = ambassador.AddToScheme(s)
	return fakeDynamic.NewSimpleDynamicClient(s), s
}

type fakeAmbassadorHost struct {
	namespace   string
	name        string
	annotations map[string]string
	hostname    string
	labels      map[string]string
}

func (ir fakeAmbassadorHost) Host() *ambassador.Host {
	spec := ambassador.HostSpec{
		Hostname: ir.hostname,
	}

	host := &ambassador.Host{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   ir.namespace,
			Name:        ir.name,
			Annotations: ir.annotations,
			Labels:      ir.labels,
		},
		Spec: &spec,
	}

	return host
}

type fakeAmbassadorLoadBalancerService struct {
	ips       []string
	hostnames []string
	namespace string
	name      string
}

func (ig fakeAmbassadorLoadBalancerService) Service() *v1.Service {
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ig.namespace,
			Name:      ig.name,
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{
				Ingress: []v1.LoadBalancerIngress{},
			},
		},
	}

	for _, ip := range ig.ips {
		svc.Status.LoadBalancer.Ingress = append(svc.Status.LoadBalancer.Ingress, v1.LoadBalancerIngress{
			IP: ip,
		})
	}
	for _, hostname := range ig.hostnames {
		svc.Status.LoadBalancer.Ingress = append(svc.Status.LoadBalancer.Ingress, v1.LoadBalancerIngress{
			Hostname: hostname,
		})
	}

	return svc
}
