/*
Copyright 2020n The Kubernetes Authors.

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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	fakeDynamic "k8s.io/client-go/dynamic/fake"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

// This is a compile-time validation that glooSource is a Source.
var _ Source = &glooSource{}

const defaultGlooNamespace = "gloo-system"

var (
	// Internal proxy test
	internalProxy = proxy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: proxyGVR.GroupVersion().String(),
			Kind:       "Proxy",
		},
		Metadata: metav1.ObjectMeta{
			Name:      "internal",
			Namespace: defaultGlooNamespace,
		},
		Spec: proxySpec{
			Listeners: []proxySpecListener{
				{
					HTTPListener: proxySpecHTTPListener{
						VirtualHosts: []proxyVirtualHost{
							{
								Domains: []string{"a.test", "b.test"},
								Metadata: proxyVirtualHostMetadata{
									Source: []proxyVirtualHostMetadataSource{
										{
											Kind:      "*v1.Unknown",
											Name:      "my-unknown-svc",
											Namespace: "unknown",
										},
									},
								},
							},
							{
								Domains: []string{"c.test"},
								Metadata: proxyVirtualHostMetadata{
									Source: []proxyVirtualHostMetadataSource{
										{
											Kind:      "*v1.VirtualService",
											Name:      "my-internal-svc",
											Namespace: "internal",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	internalProxySvc = corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      internalProxy.Metadata.Name,
			Namespace: internalProxy.Metadata.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeLoadBalancer,
		},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{IP: "203.0.113.1"},
					{IP: "203.0.113.2"},
					{IP: "203.0.113.3"},
				},
			},
		},
	}
	internalProxySource = metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			APIVersion: virtualServiceGVR.GroupVersion().String(),
			Kind:       "VirtualService",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      internalProxy.Spec.Listeners[0].HTTPListener.VirtualHosts[1].Metadata.Source[0].Name,
			Namespace: internalProxy.Spec.Listeners[0].HTTPListener.VirtualHosts[1].Metadata.Source[0].Namespace,
			Annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/ttl":                          "42",
				"external-dns.alpha.kubernetes.io/aws-geolocation-country-code": "LU",
				"external-dns.alpha.kubernetes.io/set-identifier":               "identifier",
			},
		},
	}

	// External proxy test
	externalProxy = proxy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: proxyGVR.GroupVersion().String(),
			Kind:       "Proxy",
		},
		Metadata: metav1.ObjectMeta{
			Name:      "external",
			Namespace: defaultGlooNamespace,
		},
		Spec: proxySpec{
			Listeners: []proxySpecListener{
				{
					HTTPListener: proxySpecHTTPListener{
						VirtualHosts: []proxyVirtualHost{
							{
								Domains: []string{"d.test"},
								Metadata: proxyVirtualHostMetadata{
									Source: []proxyVirtualHostMetadataSource{
										{
											Kind:      "*v1.Unknown",
											Name:      "my-unknown-svc",
											Namespace: "unknown",
										},
									},
								},
							},
							{
								Domains: []string{"e.test"},
								Metadata: proxyVirtualHostMetadata{
									Source: []proxyVirtualHostMetadataSource{
										{
											Kind:      "*v1.VirtualService",
											Name:      "my-external-svc",
											Namespace: "external",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	externalProxySvc = corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      externalProxy.Metadata.Name,
			Namespace: externalProxy.Metadata.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeLoadBalancer,
		},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{Hostname: "a.example.org"},
					{Hostname: "b.example.org"},
					{Hostname: "c.example.org"},
				},
			},
		},
	}
	externalProxySource = metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			APIVersion: virtualServiceGVR.GroupVersion().String(),
			Kind:       "VirtualService",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      externalProxy.Spec.Listeners[0].HTTPListener.VirtualHosts[1].Metadata.Source[0].Name,
			Namespace: externalProxy.Spec.Listeners[0].HTTPListener.VirtualHosts[1].Metadata.Source[0].Namespace,
			Annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/ttl":                          "24",
				"external-dns.alpha.kubernetes.io/aws-geolocation-country-code": "JP",
				"external-dns.alpha.kubernetes.io/set-identifier":               "identifier-external",
			},
		},
	}

	// Proxy with metadata static test
	proxyWithMetadataStatic = proxy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: proxyGVR.GroupVersion().String(),
			Kind:       "Proxy",
		},
		Metadata: metav1.ObjectMeta{
			Name:      "internal-static",
			Namespace: defaultGlooNamespace,
		},
		Spec: proxySpec{
			Listeners: []proxySpecListener{
				{
					HTTPListener: proxySpecHTTPListener{
						VirtualHosts: []proxyVirtualHost{
							{
								Domains: []string{"f.test", "g.test"},
								MetadataStatic: proxyVirtualHostMetadataStatic{
									Source: []proxyVirtualHostMetadataStaticSource{
										{
											ResourceKind: "*v1.Unknown",
											ResourceRef: proxyVirtualHostMetadataSourceResourceRef{
												Name:      "my-unknown-svc",
												Namespace: "unknown",
											},
										},
									},
								},
							},
							{
								Domains: []string{"h.test"},
								MetadataStatic: proxyVirtualHostMetadataStatic{
									Source: []proxyVirtualHostMetadataStaticSource{
										{
											ResourceKind: "*v1.VirtualService",
											ResourceRef: proxyVirtualHostMetadataSourceResourceRef{
												Name:      "my-internal-static-svc",
												Namespace: "internal-static",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	proxyWithMetadataStaticSvc = corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxyWithMetadataStatic.Metadata.Name,
			Namespace: proxyWithMetadataStatic.Metadata.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeLoadBalancer,
		},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{IP: "203.0.115.1"},
					{IP: "203.0.115.2"},
					{IP: "203.0.115.3"},
				},
			},
		},
	}
	proxyWithMetadataStaticSource = metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			APIVersion: virtualServiceGVR.GroupVersion().String(),
			Kind:       "VirtualService",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      proxyWithMetadataStatic.Spec.Listeners[0].HTTPListener.VirtualHosts[1].MetadataStatic.Source[0].ResourceRef.Name,
			Namespace: proxyWithMetadataStatic.Spec.Listeners[0].HTTPListener.VirtualHosts[1].MetadataStatic.Source[0].ResourceRef.Namespace,
			Annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/ttl":                          "420",
				"external-dns.alpha.kubernetes.io/aws-geolocation-country-code": "ES",
				"external-dns.alpha.kubernetes.io/set-identifier":               "identifier",
			},
		},
	}

	// Proxy with target annotation test
	targetAnnotatedProxy = proxy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: proxyGVR.GroupVersion().String(),
			Kind:       "Proxy",
		},
		Metadata: metav1.ObjectMeta{
			Name:      "target-ann",
			Namespace: defaultGlooNamespace,
			Annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/target": "203.2.45.7",
			},
		},
		Spec: proxySpec{
			Listeners: []proxySpecListener{
				{
					HTTPListener: proxySpecHTTPListener{
						VirtualHosts: []proxyVirtualHost{
							{
								Domains: []string{"i.test"},
								Metadata: proxyVirtualHostMetadata{
									Source: []proxyVirtualHostMetadataSource{
										{
											Kind:      "*v1.Unknown",
											Name:      "my-unknown-svc",
											Namespace: "unknown",
										},
									},
								},
							},
							{
								Domains: []string{"j.test"},
								Metadata: proxyVirtualHostMetadata{
									Source: []proxyVirtualHostMetadataSource{
										{
											Kind:      "*v1.VirtualService",
											Name:      "my-annotated-svc",
											Namespace: "internal",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	targetAnnotatedProxySvc = corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      targetAnnotatedProxy.Metadata.Name,
			Namespace: targetAnnotatedProxy.Metadata.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeLoadBalancer,
		},
		Status: corev1.ServiceStatus{
			LoadBalancer: corev1.LoadBalancerStatus{
				Ingress: []corev1.LoadBalancerIngress{
					{IP: "203.1.115.1"},
					{IP: "203.1.115.2"},
					{IP: "203.1.115.3"},
				},
			},
		},
	}
	targetAnnotatedProxySource = metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			APIVersion: virtualServiceGVR.GroupVersion().String(),
			Kind:       "VirtualService",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      targetAnnotatedProxy.Spec.Listeners[0].HTTPListener.VirtualHosts[1].Metadata.Source[0].Name,
			Namespace: targetAnnotatedProxy.Spec.Listeners[0].HTTPListener.VirtualHosts[1].Metadata.Source[0].Namespace,
			Annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/ttl":                          "460",
				"external-dns.alpha.kubernetes.io/aws-geolocation-country-code": "IT",
				"external-dns.alpha.kubernetes.io/set-identifier":               "identifier-annotated",
			},
		},
	}

	// Proxy backed by Ingress
	gatewayIngressAnnotatedProxy = proxy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: proxyGVR.GroupVersion().String(),
			Kind:       "Proxy",
		},
		Metadata: metav1.ObjectMeta{
			Name:      "gateway-ingress-annotated",
			Namespace: defaultGlooNamespace,
		},
		Spec: proxySpec{
			Listeners: []proxySpecListener{
				{
					HTTPListener: proxySpecHTTPListener{
						VirtualHosts: []proxyVirtualHost{
							{
								Domains: []string{"k.test"},
								MetadataStatic: proxyVirtualHostMetadataStatic{
									Source: []proxyVirtualHostMetadataStaticSource{
										{
											ResourceKind: "*v1.Unknown",
											ResourceRef: proxyVirtualHostMetadataSourceResourceRef{
												Name:      "my-unknown-svc",
												Namespace: "unknown",
											},
										},
									},
								},
							},
						},
					},
					MetadataStatic: proxyMetadataStatic{
						Source: []proxyMetadataStaticSource{
							{
								ResourceKind: "*v1.Gateway",
								ResourceRef: proxyMetadataStaticSourceResourceRef{
									Name:      "gateway-ingress-annotated",
									Namespace: defaultGlooNamespace,
								},
							},
						},
					},
				},
			},
		},
	}
	gatewayIngressAnnotatedProxyGateway = metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			APIVersion: gatewayGVR.GroupVersion().String(),
			Kind:       "Gateway",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      gatewayIngressAnnotatedProxy.Spec.Listeners[0].MetadataStatic.Source[0].ResourceRef.Name,
			Namespace: gatewayIngressAnnotatedProxy.Spec.Listeners[0].MetadataStatic.Source[0].ResourceRef.Namespace,
			Annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/ingress": fmt.Sprintf("%s/%s", gatewayIngressAnnotatedProxy.Spec.Listeners[0].MetadataStatic.Source[0].ResourceRef.Namespace, gatewayIngressAnnotatedProxy.Spec.Listeners[0].MetadataStatic.Source[0].ResourceRef.Name),
			},
		},
	}
	gatewayIngressAnnotatedProxyIngress = networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gatewayIngressAnnotatedProxy.Spec.Listeners[0].MetadataStatic.Source[0].ResourceRef.Name,
			Namespace: gatewayIngressAnnotatedProxy.Spec.Listeners[0].MetadataStatic.Source[0].ResourceRef.Namespace,
		},
		Status: networkingv1.IngressStatus{
			LoadBalancer: networkingv1.IngressLoadBalancerStatus{
				Ingress: []networkingv1.IngressLoadBalancerIngress{
					{Hostname: "example.com"},
				},
			},
		},
	}
)

func TestGlooSource(t *testing.T) {
	t.Parallel()

	fakeKubernetesClient := fakeKube.NewSimpleClientset()
	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(),
		map[schema.GroupVersionResource]string{
			proxyGVR:          "ProxyList",
			virtualServiceGVR: "VirtualServiceList",
			gatewayGVR:        "GatewayList",
		})

	internalProxyUnstructured := unstructured.Unstructured{}
	externalProxyUnstructured := unstructured.Unstructured{}
	gatewayIngressAnnotatedProxyUnstructured := unstructured.Unstructured{}
	gatewayIngressAnnotatedProxyGatewayUnstructured := unstructured.Unstructured{}
	proxyMetadataStaticUnstructured := unstructured.Unstructured{}
	targetAnnotatedProxyUnstructured := unstructured.Unstructured{}

	internalProxySourceUnstructured := unstructured.Unstructured{}
	externalProxySourceUnstructured := unstructured.Unstructured{}
	proxyMetadataStaticSourceUnstructured := unstructured.Unstructured{}
	targetAnnotatedProxySourceUnstructured := unstructured.Unstructured{}

	internalProxyAsJSON, err := json.Marshal(internalProxy)
	assert.NoError(t, err)

	externalProxyAsJSON, err := json.Marshal(externalProxy)
	assert.NoError(t, err)

	gatewayIngressAnnotatedProxyAsJSON, err := json.Marshal(gatewayIngressAnnotatedProxy)
	assert.NoError(t, err)

	gatewayIngressAnnotatedProxyGatewayAsJSON, err := json.Marshal(gatewayIngressAnnotatedProxyGateway)
	assert.NoError(t, err)

	proxyMetadataStaticAsJSON, err := json.Marshal(proxyWithMetadataStatic)
	assert.NoError(t, err)

	targetAnnotatedProxyAsJSON, err := json.Marshal(targetAnnotatedProxy)
	assert.NoError(t, err)

	internalProxySvcAsJSON, err := json.Marshal(internalProxySource)
	assert.NoError(t, err)

	externalProxySvcAsJSON, err := json.Marshal(externalProxySource)
	assert.NoError(t, err)

	proxyMetadataStaticSvcAsJSON, err := json.Marshal(proxyWithMetadataStaticSource)
	assert.NoError(t, err)

	targetAnnotatedProxySvcAsJSON, err := json.Marshal(targetAnnotatedProxySource)
	assert.NoError(t, err)

	assert.NoError(t, internalProxyUnstructured.UnmarshalJSON(internalProxyAsJSON))
	assert.NoError(t, externalProxyUnstructured.UnmarshalJSON(externalProxyAsJSON))
	assert.NoError(t, gatewayIngressAnnotatedProxyUnstructured.UnmarshalJSON(gatewayIngressAnnotatedProxyAsJSON))
	assert.NoError(t, gatewayIngressAnnotatedProxyGatewayUnstructured.UnmarshalJSON(gatewayIngressAnnotatedProxyGatewayAsJSON))
	assert.NoError(t, proxyMetadataStaticUnstructured.UnmarshalJSON(proxyMetadataStaticAsJSON))
	assert.NoError(t, targetAnnotatedProxyUnstructured.UnmarshalJSON(targetAnnotatedProxyAsJSON))

	assert.NoError(t, internalProxySourceUnstructured.UnmarshalJSON(internalProxySvcAsJSON))
	assert.NoError(t, externalProxySourceUnstructured.UnmarshalJSON(externalProxySvcAsJSON))
	assert.NoError(t, proxyMetadataStaticSourceUnstructured.UnmarshalJSON(proxyMetadataStaticSvcAsJSON))
	assert.NoError(t, targetAnnotatedProxySourceUnstructured.UnmarshalJSON(targetAnnotatedProxySvcAsJSON))

	_, err = fakeKubernetesClient.CoreV1().Services(internalProxySvc.GetNamespace()).Create(t.Context(), &internalProxySvc, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeKubernetesClient.CoreV1().Services(externalProxySvc.GetNamespace()).Create(t.Context(), &externalProxySvc, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeKubernetesClient.CoreV1().Services(proxyWithMetadataStaticSvc.GetNamespace()).Create(t.Context(), &proxyWithMetadataStaticSvc, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeKubernetesClient.CoreV1().Services(targetAnnotatedProxySvc.GetNamespace()).Create(t.Context(), &targetAnnotatedProxySvc, metav1.CreateOptions{})
	assert.NoError(t, err)

	_, err = fakeKubernetesClient.NetworkingV1().Ingresses(gatewayIngressAnnotatedProxyIngress.GetNamespace()).Create(t.Context(), &gatewayIngressAnnotatedProxyIngress, metav1.CreateOptions{})
	assert.NoError(t, err)

	// Create proxy resources
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(t.Context(), &internalProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(t.Context(), &externalProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(t.Context(), &proxyMetadataStaticUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(t.Context(), &targetAnnotatedProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(t.Context(), &gatewayIngressAnnotatedProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)

	// Create proxy source
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(internalProxySource.Namespace).Create(t.Context(), &internalProxySourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(externalProxySource.Namespace).Create(t.Context(), &externalProxySourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(proxyWithMetadataStaticSource.Namespace).Create(t.Context(), &proxyMetadataStaticSourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(targetAnnotatedProxySource.Namespace).Create(t.Context(), &targetAnnotatedProxySourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)

	// Create gateway resource
	_, err = fakeDynamicClient.Resource(gatewayGVR).Namespace(gatewayIngressAnnotatedProxyGateway.Namespace).Create(t.Context(), &gatewayIngressAnnotatedProxyGatewayUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)

	source, err := NewGlooSource(t.Context(), fakeDynamicClient, fakeKubernetesClient, &Config{
		GlooNamespaces: []string{defaultGlooNamespace},
	})
	assert.NoError(t, err)
	assert.NotNil(t, source)

	endpoints, err := source.Endpoints(t.Context())
	assert.NoError(t, err)
	assert.Len(t, endpoints, 11)

	assert.ElementsMatch(t, endpoints, []*endpoint.Endpoint{
		{
			DNSName:          "a.test",
			Targets:          []string{internalProxySvc.Status.LoadBalancer.Ingress[0].IP, internalProxySvc.Status.LoadBalancer.Ingress[1].IP, internalProxySvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:       endpoint.RecordTypeA,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:          "b.test",
			Targets:          []string{internalProxySvc.Status.LoadBalancer.Ingress[0].IP, internalProxySvc.Status.LoadBalancer.Ingress[1].IP, internalProxySvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:       endpoint.RecordTypeA,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:       "c.test",
			Targets:       []string{internalProxySvc.Status.LoadBalancer.Ingress[0].IP, internalProxySvc.Status.LoadBalancer.Ingress[1].IP, internalProxySvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "identifier",
			RecordTTL:     42,
			Labels:        endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "aws/geolocation-country-code",
					Value: "LU",
				},
			},
		},
		{
			DNSName:          "d.test",
			Targets:          []string{externalProxySvc.Status.LoadBalancer.Ingress[0].Hostname, externalProxySvc.Status.LoadBalancer.Ingress[1].Hostname, externalProxySvc.Status.LoadBalancer.Ingress[2].Hostname},
			RecordType:       endpoint.RecordTypeCNAME,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:       "e.test",
			Targets:       []string{externalProxySvc.Status.LoadBalancer.Ingress[0].Hostname, externalProxySvc.Status.LoadBalancer.Ingress[1].Hostname, externalProxySvc.Status.LoadBalancer.Ingress[2].Hostname},
			RecordType:    endpoint.RecordTypeCNAME,
			SetIdentifier: "identifier-external",
			RecordTTL:     24,
			Labels:        endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "aws/geolocation-country-code",
					Value: "JP",
				},
			},
		},
		{
			DNSName:          "f.test",
			Targets:          []string{proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[0].IP, proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[1].IP, proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:       endpoint.RecordTypeA,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:          "g.test",
			Targets:          []string{proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[0].IP, proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[1].IP, proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:       endpoint.RecordTypeA,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:       "h.test",
			Targets:       []string{proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[0].IP, proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[1].IP, proxyWithMetadataStaticSvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "identifier",
			RecordTTL:     420,
			Labels:        endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "aws/geolocation-country-code",
					Value: "ES",
				},
			},
		},
		{
			DNSName:          "i.test",
			Targets:          []string{"203.2.45.7"},
			RecordType:       endpoint.RecordTypeA,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:       "j.test",
			Targets:       []string{"203.2.45.7"},
			RecordType:    endpoint.RecordTypeA,
			SetIdentifier: "identifier-annotated",
			RecordTTL:     460,
			Labels:        endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "aws/geolocation-country-code",
					Value: "IT",
				},
			},
		},
		{
			DNSName:          "k.test",
			Targets:          []string{gatewayIngressAnnotatedProxyIngress.Status.LoadBalancer.Ingress[0].Hostname},
			RecordType:       endpoint.RecordTypeCNAME,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{}},
	})
}

func TestTransformerInGlooSource(t *testing.T) {
	newSource := func(t *testing.T, dClient *fakeDynamic.FakeDynamicClient, kClient *fakeKube.Clientset, ns ...string) *glooSource {
		t.Helper()
		src, err := NewGlooSource(t.Context(), dClient, kClient, &Config{GlooNamespaces: ns})
		require.NoError(t, err)
		gs, ok := src.(*glooSource)
		require.True(t, ok)
		return gs
	}

	t.Run("service strips managed fields and status conditions", func(t *testing.T) {
		var (
			svc = &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-service",
					Namespace: "default",
					Labels:    map[string]string{"label1": "value1"},
					Annotations: map[string]string{
						"user-annotation":                  "value",
						corev1.LastAppliedConfigAnnotation: `{"apiVersion":"v1"}`,
					},
					UID: "someuid",
					ManagedFields: []metav1.ManagedFieldsEntry{
						{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
					},
				},
				Spec: corev1.ServiceSpec{Type: corev1.ServiceTypeLoadBalancer},
				Status: corev1.ServiceStatus{
					LoadBalancer: corev1.LoadBalancerStatus{
						Ingress: []corev1.LoadBalancerIngress{{IP: "1.2.3.4"}},
					},
					Conditions: []metav1.Condition{
						{Type: "Available", Status: metav1.ConditionTrue, Reason: "Ready"},
					},
				},
			}
			gs = newSource(t, newGlooDynamicClient(), fakeKube.NewSimpleClientset(svc), "default")
		)

		retrieved, err := gs.serviceInformer.Lister().Services(svc.Namespace).Get(svc.Name)
		require.NoError(t, err)

		assert.Equal(t, svc.Name, retrieved.Name)
		assert.Equal(t, svc.Labels, retrieved.Labels)
		assert.Equal(t, svc.UID, retrieved.UID)
		assert.Empty(t, retrieved.ManagedFields)
		assert.NotContains(t, retrieved.Annotations, corev1.LastAppliedConfigAnnotation)
		assert.Contains(t, retrieved.Annotations, "user-annotation")
		// Status.LoadBalancer preserved — used for endpoint generation
		assert.Equal(t, svc.Status.LoadBalancer, retrieved.Status.LoadBalancer)
		// Status.Conditions stripped
		assert.Empty(t, retrieved.Status.Conditions)
	})

	t.Run("ingress strips managed fields", func(t *testing.T) {
		var (
			ingress = &networkingv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ingress",
					Namespace: "default",
					Labels:    map[string]string{"label1": "value1"},
					Annotations: map[string]string{
						"user-annotation":                  "value",
						corev1.LastAppliedConfigAnnotation: `{"apiVersion":"networking.k8s.io/v1"}`,
					},
					UID: "someuid",
					ManagedFields: []metav1.ManagedFieldsEntry{
						{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
					},
				},
				Status: networkingv1.IngressStatus{
					LoadBalancer: networkingv1.IngressLoadBalancerStatus{
						Ingress: []networkingv1.IngressLoadBalancerIngress{{IP: "1.2.3.4"}},
					},
				},
			}
			gs = newSource(t, newGlooDynamicClient(), fakeKube.NewSimpleClientset(ingress), "default", "kube-system")
		)

		retrieved, err := gs.ingressInformer.Lister().Ingresses(ingress.Namespace).Get(ingress.Name)
		require.NoError(t, err)

		assert.Equal(t, ingress.Name, retrieved.Name)
		assert.Equal(t, ingress.Labels, retrieved.Labels)
		assert.Equal(t, ingress.UID, retrieved.UID)
		assert.Empty(t, retrieved.ManagedFields)
		assert.NotContains(t, retrieved.Annotations, corev1.LastAppliedConfigAnnotation)
		assert.Contains(t, retrieved.Annotations, "user-annotation")
		// Status.LoadBalancer preserved — used for endpoint generation
		assert.Equal(t, ingress.Status.LoadBalancer, retrieved.Status.LoadBalancer)
	})

	t.Run("proxy strips managed fields", func(t *testing.T) {
		var proxyObj unstructured.Unstructured
		proxyObj.SetName("test-proxy")
		proxyObj.SetNamespace(defaultGlooNamespace)
		proxyObj.SetUID("someuid")
		proxyObj.SetLabels(map[string]string{"label1": "value1"})
		proxyObj.SetAnnotations(map[string]string{
			"user-annotation":                  "value",
			corev1.LastAppliedConfigAnnotation: `{"apiVersion":"gloo.solo.io/v1"}`,
		})
		proxyObj.SetManagedFields([]metav1.ManagedFieldsEntry{
			{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
		})
		proxyObj.SetGroupVersionKind(proxyGVR.GroupVersion().WithKind("Proxy"))

		gs := newSource(t, newGlooDynamicClient(&proxyObj), fakeKube.NewSimpleClientset(), defaultGlooNamespace, "default", "production")

		retrieved, err := gs.proxyInformer.Lister().ByNamespace(defaultGlooNamespace).Get(proxyObj.GetName())
		require.NoError(t, err)

		obj, ok := retrieved.(*unstructured.Unstructured)
		require.True(t, ok)

		assert.Equal(t, proxyObj.GetName(), obj.GetName())
		assert.Equal(t, proxyObj.GetLabels(), obj.GetLabels())
		assert.Equal(t, proxyObj.GetUID(), obj.GetUID())
		assert.Empty(t, obj.GetManagedFields())
		assert.NotContains(t, obj.GetAnnotations(), corev1.LastAppliedConfigAnnotation)
		assert.Contains(t, obj.GetAnnotations(), "user-annotation")
	})
}

func newGlooDynamicClient(objs ...runtime.Object) *fakeDynamic.FakeDynamicClient {
	return fakeDynamic.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(),
		map[schema.GroupVersionResource]string{
			proxyGVR:          "ProxyList",
			virtualServiceGVR: "VirtualServiceList",
			gatewayGVR:        "GatewayList",
		}, objs...)
}
