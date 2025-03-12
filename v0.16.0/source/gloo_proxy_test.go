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
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
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

// Internal proxy test
var internalProxy = proxy{
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

var internalProxySvc = corev1.Service{
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
				{
					IP: "203.0.113.1",
				},
				{
					IP: "203.0.113.2",
				},
				{
					IP: "203.0.113.3",
				},
			},
		},
	},
}

var internalProxySource = metav1.PartialObjectMetadata{
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
var externalProxy = proxy{
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

var externalProxySvc = corev1.Service{
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
				{
					Hostname: "a.example.org",
				},
				{
					Hostname: "b.example.org",
				},
				{
					Hostname: "c.example.org",
				},
			},
		},
	},
}

var externalProxySource = metav1.PartialObjectMetadata{
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
var proxyMetadataStatic = proxy{
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

var proxyMetadataStaticSvc = corev1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      proxyMetadataStatic.Metadata.Name,
		Namespace: proxyMetadataStatic.Metadata.Namespace,
	},
	Spec: corev1.ServiceSpec{
		Type: corev1.ServiceTypeLoadBalancer,
	},
	Status: corev1.ServiceStatus{
		LoadBalancer: corev1.LoadBalancerStatus{
			Ingress: []corev1.LoadBalancerIngress{
				{
					IP: "203.0.115.1",
				},
				{
					IP: "203.0.115.2",
				},
				{
					IP: "203.0.115.3",
				},
			},
		},
	},
}

var proxyMetadataStaticSource = metav1.PartialObjectMetadata{
	TypeMeta: metav1.TypeMeta{
		APIVersion: virtualServiceGVR.GroupVersion().String(),
		Kind:       "VirtualService",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      proxyMetadataStatic.Spec.Listeners[0].HTTPListener.VirtualHosts[1].MetadataStatic.Source[0].ResourceRef.Name,
		Namespace: proxyMetadataStatic.Spec.Listeners[0].HTTPListener.VirtualHosts[1].MetadataStatic.Source[0].ResourceRef.Namespace,
		Annotations: map[string]string{
			"external-dns.alpha.kubernetes.io/ttl":                          "420",
			"external-dns.alpha.kubernetes.io/aws-geolocation-country-code": "ES",
			"external-dns.alpha.kubernetes.io/set-identifier":               "identifier",
		},
	},
}

// Proxy with target annotation test
var targetAnnotatedProxy = proxy{
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

var targetAnnotatedProxySvc = corev1.Service{
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
				{
					IP: "203.1.115.1",
				},
				{
					IP: "203.1.115.2",
				},
				{
					IP: "203.1.115.3",
				},
			},
		},
	},
}

var targetAnnotatedProxySource = metav1.PartialObjectMetadata{
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

func TestGlooSource(t *testing.T) {
	t.Parallel()

	fakeKubernetesClient := fakeKube.NewSimpleClientset()
	fakeDynamicClient := fakeDynamic.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(),
		map[schema.GroupVersionResource]string{
			proxyGVR: "ProxyList",
		})

	source, err := NewGlooSource(fakeDynamicClient, fakeKubernetesClient, []string{defaultGlooNamespace})
	assert.NoError(t, err)
	assert.NotNil(t, source)

	internalProxyUnstructured := unstructured.Unstructured{}
	externalProxyUnstructured := unstructured.Unstructured{}
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

	proxyMetadataStaticAsJSON, err := json.Marshal(proxyMetadataStatic)
	assert.NoError(t, err)

	targetAnnotatedProxyAsJSON, err := json.Marshal(targetAnnotatedProxy)
	assert.NoError(t, err)

	internalProxySvcAsJSON, err := json.Marshal(internalProxySource)
	assert.NoError(t, err)

	externalProxySvcAsJSON, err := json.Marshal(externalProxySource)
	assert.NoError(t, err)

	proxyMetadataStaticSvcAsJSON, err := json.Marshal(proxyMetadataStaticSource)
	assert.NoError(t, err)

	targetAnnotatedProxySvcAsJSON, err := json.Marshal(targetAnnotatedProxySource)
	assert.NoError(t, err)

	assert.NoError(t, internalProxyUnstructured.UnmarshalJSON(internalProxyAsJSON))
	assert.NoError(t, externalProxyUnstructured.UnmarshalJSON(externalProxyAsJSON))
	assert.NoError(t, proxyMetadataStaticUnstructured.UnmarshalJSON(proxyMetadataStaticAsJSON))
	assert.NoError(t, targetAnnotatedProxyUnstructured.UnmarshalJSON(targetAnnotatedProxyAsJSON))

	assert.NoError(t, internalProxySourceUnstructured.UnmarshalJSON(internalProxySvcAsJSON))
	assert.NoError(t, externalProxySourceUnstructured.UnmarshalJSON(externalProxySvcAsJSON))
	assert.NoError(t, proxyMetadataStaticSourceUnstructured.UnmarshalJSON(proxyMetadataStaticSvcAsJSON))
	assert.NoError(t, targetAnnotatedProxySourceUnstructured.UnmarshalJSON(targetAnnotatedProxySvcAsJSON))

	// Create proxy resources
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(context.Background(), &internalProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(context.Background(), &externalProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(context.Background(), &proxyMetadataStaticUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(proxyGVR).Namespace(defaultGlooNamespace).Create(context.Background(), &targetAnnotatedProxyUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)

	// Create proxy source
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(internalProxySource.Namespace).Create(context.Background(), &internalProxySourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(externalProxySource.Namespace).Create(context.Background(), &externalProxySourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(proxyMetadataStaticSource.Namespace).Create(context.Background(), &proxyMetadataStaticSourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeDynamicClient.Resource(virtualServiceGVR).Namespace(targetAnnotatedProxySource.Namespace).Create(context.Background(), &targetAnnotatedProxySourceUnstructured, metav1.CreateOptions{})
	assert.NoError(t, err)

	// Create proxy service resources
	_, err = fakeKubernetesClient.CoreV1().Services(internalProxySvc.GetNamespace()).Create(context.Background(), &internalProxySvc, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeKubernetesClient.CoreV1().Services(externalProxySvc.GetNamespace()).Create(context.Background(), &externalProxySvc, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeKubernetesClient.CoreV1().Services(proxyMetadataStaticSvc.GetNamespace()).Create(context.Background(), &proxyMetadataStaticSvc, metav1.CreateOptions{})
	assert.NoError(t, err)
	_, err = fakeKubernetesClient.CoreV1().Services(targetAnnotatedProxySvc.GetNamespace()).Create(context.Background(), &targetAnnotatedProxySvc, metav1.CreateOptions{})
	assert.NoError(t, err)

	endpoints, err := source.Endpoints(context.Background())
	assert.NoError(t, err)
	assert.Len(t, endpoints, 10)
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
			Targets:          []string{proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[0].IP, proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[1].IP, proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:       endpoint.RecordTypeA,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:          "g.test",
			Targets:          []string{proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[0].IP, proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[1].IP, proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[2].IP},
			RecordType:       endpoint.RecordTypeA,
			RecordTTL:        0,
			Labels:           endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
		{
			DNSName:       "h.test",
			Targets:       []string{proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[0].IP, proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[1].IP, proxyMetadataStaticSvc.Status.LoadBalancer.Ingress[2].IP},
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
	})
}
