/*
Copyright 2017 The Kubernetes Authors.

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
	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"

	routev1 "github.com/openshift/api/route/v1"
	ingressoperatorclient "github.com/openshift/client-go/operator/clientset/versioned/fake"
	fake "github.com/openshift/client-go/route/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetesFakeClient "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

type OCPRouteSuite struct {
	suite.Suite
	sc               Source
	routeWithTargets *routev1.Route
	//loadbalancerService *v1.Service
}

func (suite *OCPRouteSuite) SetupTest() {
	fakeClient := fake.NewSimpleClientset()
	kubeFakeClient := kubernetesFakeClient.NewSimpleClientset()
	var err error

	suite.sc, err = NewOcpRouteSource(
		fakeClient,
		kubeFakeClient,
		"",
		"",
		"",
		false,
		true,
	)

	suite.routeWithTargets = &routev1.Route{
		Spec: routev1.RouteSpec{
			Host: "my-domain.com",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   "default",
			Name:        "route-with-targets",
			Annotations: map[string]string{},
		},
		Status: routev1.RouteStatus{
			Ingress: []routev1.RouteIngress{
				{
					RouterCanonicalHostname: "apps.my-domain.com",
				},
			},
		},
	}

	suite.NoError(err, "should initialize route source")

	_, err = fakeClient.RouteV1().Routes(suite.routeWithTargets.Namespace).Create(context.Background(), suite.routeWithTargets, metav1.CreateOptions{})
	suite.NoError(err, "should successfully create route")
}

func (suite *OCPRouteSuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.sc.Endpoints(context.Background())
	for _, ep := range endpoints {
		suite.Equal("route/default/route-with-targets", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestOcpRouteSource(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(OCPRouteSuite))
	//t.Run("Interface", testOcpRouteSourceImplementsSource)
	//t.Run("NewOcpRouteSource", testOcpRouteSourceNewOcpRouteSource)
	//t.Run("Endpoints", testOcpRouteSourceEndpoints)
	t.Run("NewOcpRouteSourceForOCP4", testOcpRouteSourceEndpointsForOCP4)
}

// testOcpRouteSourceImplementsSource tests that ocpRouteSource is a valid Source.
func testOcpRouteSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(ocpRouteSource))
}

// testOcpRouteSourceNewOcpRouteSource tests that NewOcpRouteSource doesn't return an error.
func testOcpRouteSourceNewOcpRouteSource(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title            string
		annotationFilter string
		fqdnTemplate     string
		expectError      bool
	}{
		{
			title:        "invalid template",
			expectError:  true,
			fqdnTemplate: "{{.Name",
		},
		{
			title:       "valid empty template",
			expectError: false,
		},
		{
			title:        "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			title:            "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "kubernetes.io/ingress.class=nginx",
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewOcpRouteSource(
				fake.NewSimpleClientset(),
				kubernetesFakeClient.NewSimpleClientset(),
				"",
				ti.annotationFilter,
				ti.fqdnTemplate,
				false,
				true,
			)

			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// testOcpRouteSourceEndpoints tests that various OCP routes generate the correct endpoints.
func testOcpRouteSourceEndpoints(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		fqdnTemplate             string
		ignoreHostnameAnnotation bool
		ocpRoute                 *routev1.Route
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			title:                    "route with basic hostname and route status target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "my-domain.com",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							RouterCanonicalHostname: "apps.my-domain.com",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "my-domain.com",
					Targets: []string{
						"apps.my-domain.com",
					},
				},
			},
			expectError: false,
		},
		{
			title:                    "route with incorrect externalDNS controller annotation",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-ignore-annotation",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/controller": "foo",
					},
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: false,
		},
		{
			title:                    "route with basic hostname and annotation target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "my-annotation-domain.com",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-annotation-target",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/target": "my.site.foo.com",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "my-annotation-domain.com",
					Targets: []string{
						"my.site.foo.com",
					},
				},
			},
			expectError: false,
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			fakeClient := fake.NewSimpleClientset()

			_, err := fakeClient.RouteV1().Routes(tc.ocpRoute.Namespace).Create(context.Background(), tc.ocpRoute, metav1.CreateOptions{})
			require.NoError(t, err)

			kubeFakeClient := kubernetesFakeClient.NewSimpleClientset()

			source, err := NewOcpRouteSource(
				fakeClient,
				kubeFakeClient,
				"",
				"",
				"",
				false,
				true,
			)
			require.NoError(t, err)

			res, err := source.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, res, tc.expected)
		})
	}
}

// testOcpRouteSourceEndpoints tests that various OCP routes generate the correct endpoints.
func testOcpRouteSourceEndpointsForOCP4(t *testing.T) {
	t.Parallel()
	ingressOperatorClient := ingressoperatorclient.NewSimpleClientset()
	kubeFakeClient := kubernetesFakeClient.NewSimpleClientset()

	for _, tc := range []struct {
		title                     string
		targetNamespace           string
		annotationFilter          string
		fqdnTemplate              string
		ignoreHostnameAnnotation  bool
		ocpRoute                  *routev1.Route
		loadbalancerService       *v1.Service
		secondloadbalancerService *v1.Service
		expected                  []*endpoint.Endpoint
		expectError               bool
	}{
		{
			title:                    "route with basic hostname and route status and service LB with single IP target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-target-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{IP: "8.8.8.8"},
							//{Hostname: "foo"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"8.8.8.8",
					},
				},
			},
			expectError: false,
		},
		{
			title:                    "route with basic hostname and route status and service LB with multiple IP targets",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-targets-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{IP: "1.1.1.1"},
							{IP: "8.8.8.8"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"1.1.1.1",
						"8.8.8.8",
					},
				},
			},
			expectError: false,
		},
		{
			title:                    "route with basic hostname and route status and service LB with single hostname target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-target-for-LB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{Hostname: "foo"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"foo",
					},
				},
			},
			expectError: false,
		},
		{
			title:                    "route with basic hostname and route status and service LB with multiple hostname targets",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-targets-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{Hostname: "abc.com"},
							{Hostname: "xyz.com"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"abc.com",
						//"xyz.com",
					},
				},
			},
			expectError: false,
		},

		{
			title:                    "route with basic hostname and route status and service LB with one hostname and one IP target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-targets-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{Hostname: "abc.com"},
							{IP: "1.1.1.1"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"abc.com",
					},
				},
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"1.1.1.1",
					},
				},
			},
			expectError: false,
		},

		{
			title:                    "route with basic hostname and route status and service LB with multiple hostnames and multiple IP targets",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-targets-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{Hostname: "abc.com"},
							{IP: "1.1.1.1"},
							{Hostname: "xyz.com"},
							{IP: "8.8.8.8"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"abc.com",
						//	"xyz.com",
					},
				},
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"1.1.1.1",
						"8.8.8.8",
					},
				},
			},
			expectError: false,
		},
		{
			title:                    "route with basic hostname, multiple router names and route status and service LB with single IP target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
						{Host: "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "test",
							RouterCanonicalHostname: "apps.example.com",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-target-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{IP: "8.8.8.8"},
						},
					},
				},
			},
			secondloadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "test"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-target-forLB1",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "test"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{IP: "1.1.1.1"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"8.8.8.8",
						//"1.1.1.1",
					},
				},
			},
			expectError: false,
		},
		{
			title:                    "route with basic hostname, multiple router names and route status and service LB with single hostname target",
			targetNamespace:          "",
			annotationFilter:         "",
			fqdnTemplate:             "",
			ignoreHostnameAnnotation: false,
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "hello-openshift.apps.misalunk.externaldns",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "default",
					Name:        "route-with-target",
					Annotations: map[string]string{},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.apps.misalunk.externaldns",
						},
						{
							Host:                    "hello-openshift.apps.misalunk.externaldns",
							RouterName:              "test",
							RouterCanonicalHostname: "apps.example.com",
						},
					},
				},
			},
			loadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "default"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-target-forLB",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "default"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{Hostname: "foo"},
						},
					},
				},
			},
			secondloadbalancerService: &v1.Service{
				Spec: v1.ServiceSpec{
					Type:     v1.ServiceTypeLoadBalancer,
					Selector: map[string]string{"ingresscontroller.operator.openshift.io/deployment-ingresscontroller": "test"},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   "openshift-ingress",
					Name:        "foo-with-target-forLB1",
					Annotations: map[string]string{},
					Labels:      map[string]string{"ingresscontroller.operator.openshift.io/owning-ingresscontroller": "test"},
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: []v1.LoadBalancerIngress{
							{Hostname: "bar"},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "hello-openshift.apps.misalunk.externaldns",
					Targets: []string{
						"foo",
						//"bar",
					},
				},
			},
			expectError: false,
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {

			// Create a Kubernetes testing client
			fakeClient := fake.NewSimpleClientset()

			_, err := fakeClient.RouteV1().Routes(tc.ocpRoute.Namespace).Create(context.Background(), tc.ocpRoute, metav1.CreateOptions{})
			require.NoError(t, err)

			kubeFakeClient.Fake.Resources = getfakeResource(kubeFakeClient)

			createDefaultIngressController(ingressOperatorClient)

			createTestIngressController(ingressOperatorClient)

			require.NoError(t, err)

			_, err = kubeFakeClient.CoreV1().Services("openshift-ingress").Create(context.Background(), tc.loadbalancerService, metav1.CreateOptions{})
			//require.NoError(t, err)

			_, err = kubeFakeClient.CoreV1().Services("openshift-ingress").Create(context.Background(), tc.secondloadbalancerService, metav1.CreateOptions{})
			//require.NoError(t, err)

			source, err := NewOcpRouteSource(
				fakeClient,
				kubeFakeClient,
				"",
				"",
				"",
				false,
				true,
			)
			require.NoError(t, err)

			res, err := source.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			fmt.Printf("tc.expected %v\n", tc.expected)
			fmt.Printf("res %v\n", res)
			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, res, tc.expected)
			kubeFakeClient.CoreV1().Services("openshift-ingress").Delete(context.Background(), tc.loadbalancerService.Name, metav1.DeleteOptions{})
			if tc.secondloadbalancerService != nil {
				kubeFakeClient.CoreV1().Services("openshift-ingress").Delete(context.Background(), tc.secondloadbalancerService.Name, metav1.DeleteOptions{})
			}

		})
	}
}

func getfakeResource(kubeFakeClient *kubernetesFakeClient.Clientset) []*metav1.APIResourceList {
	return []*metav1.APIResourceList{
		{
			GroupVersion: "operator.openshift.io/v1",
			APIResources: []metav1.APIResource{
				{
					Kind: "IngressController",
				},
			},
		},
	}
}

func createDefaultIngressController(ingressOperatorClient *ingressoperatorclient.Clientset) *operatorv1.IngressController {

	ingresscontroller := &operatorv1.IngressController{
		TypeMeta: metav1.TypeMeta{
			Kind:       "IngressController",
			APIVersion: "operator.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "default",
		},
		Status: operatorv1.IngressControllerStatus{
			Domain: "abc.com",
			EndpointPublishingStrategy: &operatorv1.EndpointPublishingStrategy{
				LoadBalancer: &operatorv1.LoadBalancerStrategy{Scope: operatorv1.ExternalLoadBalancer},
				Type:         "LoadBalancerService",
			},
		},
	}

	if ingresscontroller, err := ingressOperatorClient.OperatorV1().IngressControllers("openshift-ingress-operator").Create(context.Background(), ingresscontroller, metav1.CreateOptions{}); err != nil {
		return ingresscontroller
	}

	return nil

}

func createTestIngressController(ingressOperatorClient *ingressoperatorclient.Clientset) *operatorv1.IngressController {

	ingresscontroller := &operatorv1.IngressController{
		TypeMeta: metav1.TypeMeta{
			Kind:       "IngressController",
			APIVersion: "operator.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Status: operatorv1.IngressControllerStatus{
			Domain: "abc.com",
			EndpointPublishingStrategy: &operatorv1.EndpointPublishingStrategy{
				LoadBalancer: &operatorv1.LoadBalancerStrategy{Scope: operatorv1.ExternalLoadBalancer},
				Type:         "LoadBalancerService",
			},
		},
	}

	if ingresscontroller, err := ingressOperatorClient.OperatorV1().IngressControllers("openshift-ingress-operator").Create(context.Background(), ingresscontroller, metav1.CreateOptions{}); err != nil {
		return ingresscontroller
	}

	return nil

}
