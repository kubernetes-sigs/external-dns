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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"k8s.io/apimachinery/pkg/labels"

	routev1 "github.com/openshift/api/route/v1"
	fake "github.com/openshift/client-go/route/clientset/versioned/fake"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
)

type OCPRouteSuite struct {
	suite.Suite
	sc               Source
	routeWithTargets *routev1.Route
}

func (suite *OCPRouteSuite) SetupTest() {
	fakeClient := fake.NewSimpleClientset()
	var err error

	suite.sc, err = NewOcpRouteSource(
		context.TODO(),
		fakeClient,
		"",
		"",
		"{{.Name}}",
		false,
		false,
		labels.Everything(),
		"",
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
	t.Run("Interface", testOcpRouteSourceImplementsSource)
	t.Run("NewOcpRouteSource", testOcpRouteSourceNewOcpRouteSource)
	t.Run("Endpoints", testOcpRouteSourceEndpoints)
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
		labelFilter      string
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
		{
			title:       "valid label selector",
			expectError: false,
			labelFilter: "app=web-external",
		},
	} {
		ti := ti
		labelSelector, err := labels.Parse(ti.labelFilter)
		require.NoError(t, err)
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewOcpRouteSource(
				context.TODO(),
				fake.NewSimpleClientset(),
				"",
				ti.annotationFilter,
				ti.fqdnTemplate,
				false,
				false,
				labelSelector,
				"",
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
	for _, tc := range []struct {
		title         string
		ocpRoute      *routev1.Route
		expected      []*endpoint.Endpoint
		expectError   bool
		labelFilter   string
		ocpRouterName string
	}{
		{
			title: "route with basic hostname and route status target",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-target",
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-domain.com",
							RouterCanonicalHostname: "apps.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
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
		},
		{
			title: "route with basic hostname, route status target and ocpRouterName defined",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-target",
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			ocpRouterName: "default",
			expected: []*endpoint.Endpoint{
				{
					DNSName: "my-domain.com",
					Targets: []string{
						"router-default.my-domain.com",
					},
				},
			},
		},
		{
			title: "route with basic hostname, route status target, one ocpRouterName and two router canonical names",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-target",
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
						{
							Host:                    "my-domain.com",
							RouterName:              "test",
							RouterCanonicalHostname: "router-test.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			ocpRouterName: "default",
			expected: []*endpoint.Endpoint{
				{
					DNSName: "my-domain.com",
					Targets: []string{
						"router-default.my-domain.com",
					},
				},
			},
		},
		{
			title: "route not admitted by the given router",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-target",
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
						{
							Host:                    "my-domain.com",
							RouterName:              "test",
							RouterCanonicalHostname: "router-test.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionFalse,
								},
							},
						},
					},
				},
			},
			ocpRouterName: "test",
			expected:      []*endpoint.Endpoint{},
		},
		{
			title: "route not admitted by any router",
			ocpRoute: &routev1.Route{
				Spec: routev1.RouteSpec{
					Host: "my-domain.com",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-target",
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionFalse,
								},
							},
						},
						{
							Host:                    "my-domain.com",
							RouterName:              "test",
							RouterCanonicalHostname: "router-test.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionFalse,
								},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "route admitted by first appropriate router",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-target",
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionFalse,
								},
							},
						},
						{
							Host:                    "my-domain.com",
							RouterName:              "test",
							RouterCanonicalHostname: "router-test.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "my-domain.com",
					Targets: []string{"router-test.my-domain.com"},
				},
			},
		},
		{
			title: "route with incorrect externalDNS controller annotation",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-ignore-annotation",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/controller": "foo",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "route with basic hostname and annotation target",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-annotation-target",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/target": "my.site.foo.com",
					},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-annotation-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
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
		},
		{
			title:       "route with matching labels",
			labelFilter: "app=web-external",
			ocpRoute: &routev1.Route{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-with-matching-labels",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/target": "my.site.foo.com",
					},
					Labels: map[string]string{
						"app":  "web-external",
						"name": "service-frontend",
					},
				},
				Status: routev1.RouteStatus{
					Ingress: []routev1.RouteIngress{
						{
							Host:                    "my-annotation-domain.com",
							RouterName:              "default",
							RouterCanonicalHostname: "router-default.my-domain.com",
							Conditions: []routev1.RouteIngressCondition{
								{
									Type:   routev1.RouteAdmitted,
									Status: corev1.ConditionTrue,
								},
							},
						},
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
		},
		{
			title:       "route without matching labels",
			labelFilter: "app=web-external",
			ocpRoute: &routev1.Route{

				Spec: routev1.RouteSpec{
					Host: "my-annotation-domain.com",
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "route-without-matching-labels",
					Annotations: map[string]string{
						"external-dns.alpha.kubernetes.io/target": "my.site.foo.com",
					},
					Labels: map[string]string{
						"app":  "web-internal",
						"name": "service-frontend",
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()
			// Create a Kubernetes testing client
			fakeClient := fake.NewSimpleClientset()
			_, err := fakeClient.RouteV1().Routes(tc.ocpRoute.Namespace).Create(context.Background(), tc.ocpRoute, metav1.CreateOptions{})
			require.NoError(t, err)

			labelSelector, err := labels.Parse(tc.labelFilter)
			require.NoError(t, err)

			source, err := NewOcpRouteSource(
				context.TODO(),
				fakeClient,
				"",
				"",
				"{{.Name}}",
				false,
				false,
				labelSelector,
				tc.ocpRouterName,
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
