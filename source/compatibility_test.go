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

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

func lbService(annotations map[string]string, svcType v1.ServiceType) *v1.Service {
	return &v1.Service{
		Namespace:   "default",
		Name:        "legacy",
		Annotations: annotations,
		Spec:        v1.ServiceSpec{Type: svcType},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{
				Ingress: []v1.LoadBalancerIngress{
					{IP: "54.10.11.1"},
					{Hostname: "lb.example.org"},
				},
			},
		},
	}
}

func TestLegacyEndpointsFromService(t *testing.T) {
	t.Parallel()

	svc := lbService(map[string]string{mateAnnotationKey: "mate.example.org"}, v1.ServiceTypeLoadBalancer)

	for _, tc := range []struct {
		title         string
		compatibility string
		svc           *v1.Service
		expectEmpty   bool
		expectCount   int
	}{
		{title: "no compatibility returns nothing", compatibility: "", svc: svc, expectEmpty: true},
		{title: "unknown compatibility returns nothing", compatibility: "bogus", svc: svc, expectEmpty: true},
		{title: "mate delegates to mate semantics", compatibility: "mate", svc: svc, expectCount: 2},
		{
			title:         "molecule delegates to molecule semantics",
			compatibility: "molecule",
			svc: func() *v1.Service {
				s := lbService(map[string]string{moleculeAnnotationKey: "mol.example.org"}, v1.ServiceTypeLoadBalancer)
				s.Labels = map[string]string{"dns": "route53"}
				return s
			}(),
			expectCount: 2,
		},
		{
			title:         "kops delegates to dns-controller semantics",
			compatibility: "kops-dns-controller",
			svc:           lbService(map[string]string{kopsDNSControllerHostnameAnnotationKey: "kops.example.org"}, v1.ServiceTypeLoadBalancer),
			expectCount:   2,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			sc := &serviceSource{compatibility: tc.compatibility}
			got, err := legacyEndpointsFromService(tc.svc, sc)
			require.NoError(t, err)
			if tc.expectEmpty {
				require.Empty(t, got)
			} else {
				require.Len(t, got, tc.expectCount)
			}
		})
	}
}

func TestLegacyEndpointsFromMateService(t *testing.T) {
	t.Parallel()

	t.Run("missing annotation returns nil", func(t *testing.T) {
		t.Parallel()
		require.Nil(t, legacyEndpointsFromMateService(lbService(nil, v1.ServiceTypeLoadBalancer)))
	})

	t.Run("annotation without ingress returns nothing", func(t *testing.T) {
		t.Parallel()
		svc := lbService(map[string]string{mateAnnotationKey: "mate.example.org"}, v1.ServiceTypeLoadBalancer)
		svc.Status.LoadBalancer.Ingress = nil
		require.Empty(t, legacyEndpointsFromMateService(svc))
	})

	t.Run("IP and hostname ingress produce A and CNAME", func(t *testing.T) {
		t.Parallel()
		require.Equal(t,
			[]*endpoint.Endpoint{
				endpoint.NewEndpoint("mate.example.org", endpoint.RecordTypeA, "54.10.11.1"),
				endpoint.NewEndpoint("mate.example.org", endpoint.RecordTypeCNAME, "lb.example.org"),
			},
			legacyEndpointsFromMateService(lbService(map[string]string{mateAnnotationKey: "mate.example.org"}, v1.ServiceTypeLoadBalancer)),
		)
	})
}

func TestLegacyEndpointsFromMoleculeService(t *testing.T) {
	t.Parallel()

	t.Run("non-route53 label opts out", func(t *testing.T) {
		t.Parallel()
		svc := lbService(map[string]string{moleculeAnnotationKey: "mol.example.org"}, v1.ServiceTypeLoadBalancer)
		svc.Labels = map[string]string{"dns": "other"}
		require.Nil(t, legacyEndpointsFromMoleculeService(svc))
	})

	t.Run("missing annotation returns nil", func(t *testing.T) {
		t.Parallel()
		svc := lbService(nil, v1.ServiceTypeLoadBalancer)
		svc.Labels = map[string]string{"dns": "route53"}
		require.Nil(t, legacyEndpointsFromMoleculeService(svc))
	})

	t.Run("comma-separated hostnames with spaces each get endpoints", func(t *testing.T) {
		t.Parallel()
		svc := lbService(map[string]string{moleculeAnnotationKey: "one.example.org, two.example.org"}, v1.ServiceTypeLoadBalancer)
		svc.Labels = map[string]string{"dns": "route53"}
		require.Equal(t,
			[]*endpoint.Endpoint{
				endpoint.NewEndpoint("one.example.org", endpoint.RecordTypeA, "54.10.11.1"),
				endpoint.NewEndpoint("one.example.org", endpoint.RecordTypeCNAME, "lb.example.org"),
				endpoint.NewEndpoint("two.example.org", endpoint.RecordTypeA, "54.10.11.1"),
				endpoint.NewEndpoint("two.example.org", endpoint.RecordTypeCNAME, "lb.example.org"),
			},
			legacyEndpointsFromMoleculeService(svc),
		)
	})
}

func TestLegacyEndpointsFromDNSControllerLoadBalancerService(t *testing.T) {
	t.Parallel()

	t.Run("missing annotations return nil", func(t *testing.T) {
		t.Parallel()
		require.Nil(t, legacyEndpointsFromDNSControllerLoadBalancerService(lbService(nil, v1.ServiceTypeLoadBalancer)))
	})

	t.Run("external and internal annotations combine", func(t *testing.T) {
		t.Parallel()
		svc := lbService(map[string]string{
			kopsDNSControllerHostnameAnnotationKey:         "ext.example.org",
			kopsDNSControllerInternalHostnameAnnotationKey: "int.example.org",
		}, v1.ServiceTypeLoadBalancer)
		require.Equal(t,
			[]*endpoint.Endpoint{
				endpoint.NewEndpoint("ext.example.org", endpoint.RecordTypeA, "54.10.11.1"),
				endpoint.NewEndpoint("ext.example.org", endpoint.RecordTypeCNAME, "lb.example.org"),
				endpoint.NewEndpoint("int.example.org", endpoint.RecordTypeA, "54.10.11.1"),
				endpoint.NewEndpoint("int.example.org", endpoint.RecordTypeCNAME, "lb.example.org"),
			},
			legacyEndpointsFromDNSControllerLoadBalancerService(svc),
		)
	})
}

func nodePortServiceSource(t *testing.T, exposeInternalIPv6 bool) *serviceSource {
	t.Helper()

	nodes := []*v1.Node{
		{
			Name:   "worker",
			Labels: map[string]string{"node-role.kubernetes.io/node": ""},
			Status: v1.NodeStatus{
				Addresses: []v1.NodeAddress{
					{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
					{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
					{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
				},
			},
		},
		{
			// No node role label: skipped as a target.
			Name: "control-plane",
			Status: v1.NodeStatus{
				Addresses: []v1.NodeAddress{
					{Type: v1.NodeExternalIP, Address: "54.10.11.99"},
				},
			},
		},
	}

	client := fake.NewClientset()
	for _, node := range nodes {
		_, err := client.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	sc, err := NewServiceSource(t.Context(), client, &Config{
		Compatibility:      "kops-dns-controller",
		ExposeInternalIPv6: exposeInternalIPv6,
		LabelFilter:        labels.Everything(),
	})
	require.NoError(t, err)
	require.IsType(t, &serviceSource{}, sc)
	return sc.(*serviceSource)
}

func TestLegacyEndpointsFromDNSControllerNodePortService(t *testing.T) {
	t.Parallel()

	newSvc := func(annotations map[string]string) *v1.Service {
		svc := lbService(annotations, v1.ServiceTypeNodePort)
		return svc
	}

	t.Run("missing annotations return nil", func(t *testing.T) {
		t.Parallel()
		got, err := legacyEndpointsFromService(newSvc(nil), nodePortServiceSource(t, false))
		require.NoError(t, err)
		require.Nil(t, got)
	})

	t.Run("both annotations return nil like dns-controller", func(t *testing.T) {
		t.Parallel()
		got, err := legacyEndpointsFromService(newSvc(map[string]string{
			kopsDNSControllerHostnameAnnotationKey:         "ext.example.org",
			kopsDNSControllerInternalHostnameAnnotationKey: "int.example.org",
		}), nodePortServiceSource(t, false))
		require.NoError(t, err)
		require.Nil(t, got)
	})

	t.Run("external uses external IPs of nodes with the node role", func(t *testing.T) {
		t.Parallel()
		got, err := legacyEndpointsFromService(newSvc(map[string]string{
			kopsDNSControllerHostnameAnnotationKey: "ext.example.org",
		}), nodePortServiceSource(t, false))
		require.NoError(t, err)
		require.Equal(t,
			[]*endpoint.Endpoint{
				endpoint.NewEndpoint("ext.example.org", endpoint.RecordTypeA, "54.10.11.1"),
			},
			got,
		)
	})

	t.Run("internal uses internal IPs", func(t *testing.T) {
		t.Parallel()
		got, err := legacyEndpointsFromService(newSvc(map[string]string{
			kopsDNSControllerInternalHostnameAnnotationKey: "int.example.org",
		}), nodePortServiceSource(t, false))
		require.NoError(t, err)
		require.Equal(t,
			[]*endpoint.Endpoint{
				endpoint.NewEndpoint("int.example.org", endpoint.RecordTypeA, "10.0.1.1"),
				endpoint.NewEndpoint("int.example.org", endpoint.RecordTypeAAAA, "2001:DB8::2"),
			},
			got,
		)
	})

	t.Run("exposed internal IPv6 yields AAAA", func(t *testing.T) {
		t.Parallel()
		got, err := legacyEndpointsFromService(newSvc(map[string]string{
			kopsDNSControllerHostnameAnnotationKey: "ext.example.org",
		}), nodePortServiceSource(t, true))
		require.NoError(t, err)
		require.Equal(t,
			[]*endpoint.Endpoint{
				endpoint.NewEndpoint("ext.example.org", endpoint.RecordTypeA, "54.10.11.1"),
				endpoint.NewEndpoint("ext.example.org", endpoint.RecordTypeAAAA, "2001:DB8::2"),
			},
			got,
		)
	})

	t.Run("non NodePort or LB type returns empty", func(t *testing.T) {
		t.Parallel()
		svc := lbService(map[string]string{
			kopsDNSControllerHostnameAnnotationKey: "ext.example.org",
		}, v1.ServiceTypeClusterIP)
		got, err := legacyEndpointsFromService(svc, &serviceSource{compatibility: "kops-dns-controller"})
		require.NoError(t, err)
		require.Empty(t, got)
	})
}
