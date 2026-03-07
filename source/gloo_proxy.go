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
	"fmt"
	"maps"
	"strings"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"

	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	netinformers "k8s.io/client-go/informers/networking/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
)

var (
	proxyGVR = schema.GroupVersionResource{
		Group:    "gloo.solo.io",
		Version:  "v1",
		Resource: "proxies",
	}
	virtualServiceGVR = schema.GroupVersionResource{
		Group:    "gateway.solo.io",
		Version:  "v1",
		Resource: "virtualservices",
	}
	gatewayGVR = schema.GroupVersionResource{
		Group:    "gateway.solo.io",
		Version:  "v1",
		Resource: "gateways",
	}
)

// Basic redefinition of "Proxy" CRD : https://github.com/solo-io/gloo/blob/v1.4.6/projects/gloo/pkg/api/v1/proxy.pb.go
type proxy struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata"`
	Spec            proxySpec         `json:"spec"`
}

type proxySpec struct {
	Listeners []proxySpecListener `json:"listeners,omitempty"`
}

type proxySpecListener struct {
	HTTPListener   proxySpecHTTPListener `json:"httpListener"`
	MetadataStatic proxyMetadataStatic   `json:"metadataStatic"`
}

type proxyMetadataStatic struct {
	Source []proxyMetadataStaticSource `json:"sources,omitempty"`
}

type proxyMetadataStaticSource struct {
	ResourceKind string                               `json:"resourceKind,omitempty"`
	ResourceRef  proxyMetadataStaticSourceResourceRef `json:"resourceRef"`
}

type proxyMetadataStaticSourceResourceRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type proxySpecHTTPListener struct {
	VirtualHosts []proxyVirtualHost `json:"virtualHosts,omitempty"`
}

type proxyVirtualHost struct {
	Domains        []string                       `json:"domains,omitempty"`
	Metadata       proxyVirtualHostMetadata       `json:"metadata"`
	MetadataStatic proxyVirtualHostMetadataStatic `json:"metadataStatic"`
}

type proxyVirtualHostMetadata struct {
	Source []proxyVirtualHostMetadataSource `json:"sources,omitempty"`
}

type proxyVirtualHostMetadataStatic struct {
	Source []proxyVirtualHostMetadataStaticSource `json:"sources"`
}

type proxyVirtualHostMetadataSource struct {
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type proxyVirtualHostMetadataStaticSource struct {
	ResourceKind string                                    `json:"resourceKind"`
	ResourceRef  proxyVirtualHostMetadataSourceResourceRef `json:"resourceRef"`
}

type proxyVirtualHostMetadataSourceResourceRef struct {
	proxyVirtualHost
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// glooSource is an implementation of Source for Gloo Proxy objects.
//
// +externaldns:source:name=gloo-proxy
// +externaldns:source:category=Service Mesh
// +externaldns:source:description=Creates DNS entries from Gloo Proxy resources
// +externaldns:source:resources=Proxy.gloo.solo.io
// +externaldns:source:filters=
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=false
// +externaldns:source:provider-specific=true
type glooSource struct {
	serviceInformer        coreinformers.ServiceInformer
	ingressInformer        netinformers.IngressInformer
	proxyInformer          kubeinformers.GenericInformer
	virtualServiceInformer kubeinformers.GenericInformer
	gatewayInformer        kubeinformers.GenericInformer
	glooNamespaces         []string
}

// NewGlooSource creates a new glooSource with the given config
func NewGlooSource(ctx context.Context, dynamicKubeClient dynamic.Interface, kubeClient kubernetes.Interface,
	glooNamespaces []string) (Source, error) {
	informerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, 0)
	serviceInformer := informerFactory.Core().V1().Services()
	ingressInformer := informerFactory.Networking().V1().Ingresses()

	_, _ = serviceInformer.Informer().AddEventHandler(informers.DefaultEventHandler())
	_, _ = ingressInformer.Informer().AddEventHandler(informers.DefaultEventHandler())

	dynamicInformerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynamicKubeClient, 0)

	proxyInformer := dynamicInformerFactory.ForResource(proxyGVR)
	virtualServiceInformer := dynamicInformerFactory.ForResource(virtualServiceGVR)
	gatewayInformer := dynamicInformerFactory.ForResource(gatewayGVR)

	_, _ = proxyInformer.Informer().AddEventHandler(informers.DefaultEventHandler())
	_, _ = virtualServiceInformer.Informer().AddEventHandler(informers.DefaultEventHandler())
	_, _ = gatewayInformer.Informer().AddEventHandler(informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())
	dynamicInformerFactory.Start(ctx.Done())
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}
	if err := informers.WaitForDynamicCacheSync(ctx, dynamicInformerFactory); err != nil {
		return nil, err
	}

	return &glooSource{
		serviceInformer,
		ingressInformer,
		proxyInformer,
		virtualServiceInformer,
		gatewayInformer,
		glooNamespaces,
	}, nil
}

func (gs *glooSource) AddEventHandler(_ context.Context, _ func()) {
}

// Endpoints returns endpoint objects
func (gs *glooSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	for _, ns := range gs.glooNamespaces {
		proxyObjects, err := gs.proxyInformer.Lister().ByNamespace(ns).List(labels.Everything())
		if err != nil {
			return nil, err
		}

		for _, obj := range proxyObjects {
			unstructuredObj, ok := obj.(*unstructured.Unstructured)
			if !ok {
				return nil, err
			}

			jsonData, err := json.Marshal(unstructuredObj.Object)
			if err != nil {
				return nil, err
			}

			var proxy proxy
			if err = json.Unmarshal(jsonData, &proxy); err != nil {
				return nil, err
			}
			log.Debugf("Gloo: Find %s proxy", proxy.Metadata.Name)

			proxyTargets := annotations.TargetsFromTargetAnnotation(proxy.Metadata.Annotations)
			if len(proxyTargets) == 0 {
				proxyTargets, err = gs.targetsFromGatewayIngress(&proxy)
				if err != nil {
					return nil, err
				}
			}

			if len(proxyTargets) == 0 {
				proxyTargets, err = gs.proxyTargets(proxy.Metadata.Name, ns)
				if err != nil {
					return nil, err
				}
			}
			log.Debugf("Gloo[%s]: Find %d target(s) (%+v)", proxy.Metadata.Name, len(proxyTargets), proxyTargets)

			proxyEndpoints, err := gs.generateEndpointsFromProxy(&proxy, proxyTargets)
			if err != nil {
				return nil, err
			}
			log.Debugf("Gloo[%s]: Generate %d endpoint(s)", proxy.Metadata.Name, len(proxyEndpoints))
			endpoints = append(endpoints, proxyEndpoints...)
		}
	}
	return MergeEndpoints(endpoints), nil
}

func (gs *glooSource) generateEndpointsFromProxy(proxy *proxy, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	resource := fmt.Sprintf("proxy/%s/%s", proxy.Metadata.Namespace, proxy.Metadata.Name)

	for _, listener := range proxy.Spec.Listeners {
		for _, virtualHost := range listener.HTTPListener.VirtualHosts {
			ants, err := gs.annotationsFromProxySource(virtualHost)
			if err != nil {
				return nil, err
			}
			ttl := annotations.TTLFromAnnotations(ants, resource)
			providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ants)
			for _, domain := range virtualHost.Domains {
				endpoints = append(endpoints, EndpointsForHostname(strings.TrimSuffix(domain, "."), targets, ttl, providerSpecific, setIdentifier, "")...)
			}
		}
	}
	return endpoints, nil
}

func (gs *glooSource) annotationsFromProxySource(virtualHost proxyVirtualHost) (map[string]string, error) {
	ants := map[string]string{}
	for _, src := range virtualHost.Metadata.Source {
		if src.Kind != "*v1.VirtualService" {
			log.Debugf("Unsupported listener source. Expecting '*v1.VirtualService', got (%s)", src.Kind)
			continue
		}

		virtualServiceObj, err := gs.virtualServiceInformer.Lister().ByNamespace(src.Namespace).Get(src.Name)
		if err != nil {
			return nil, err
		}
		unstructuredVirtualService, ok := virtualServiceObj.(*unstructured.Unstructured)
		if !ok {
			log.Error("unexpected object: it is not *unstructured.Unstructured")
			continue
		}

		maps.Copy(ants, unstructuredVirtualService.GetAnnotations())
	}

	for _, src := range virtualHost.MetadataStatic.Source {
		if src.ResourceKind != "*v1.VirtualService" {
			log.Debugf("Unsupported listener source. Expecting '*v1.VirtualService', got (%s)", src.ResourceKind)
			continue
		}
		virtualServiceObj, err := gs.virtualServiceInformer.Lister().ByNamespace(src.ResourceRef.Namespace).Get(src.ResourceRef.Name)
		if err != nil {
			return nil, err
		}
		unstructuredVirtualService, ok := virtualServiceObj.(*unstructured.Unstructured)
		if !ok {
			log.Error("unexpected object: it is not *unstructured.Unstructured")
			continue
		}

		maps.Copy(ants, unstructuredVirtualService.GetAnnotations())
	}
	return ants, nil
}

func (gs *glooSource) proxyTargets(name string, namespace string) (endpoint.Targets, error) {
	svc, err := gs.serviceInformer.Lister().Services(namespace).Get(name)
	if err != nil {
		return nil, err
	}

	var targets endpoint.Targets
	switch svc.Spec.Type {
	case corev1.ServiceTypeLoadBalancer:
		for _, lb := range svc.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				targets = append(targets, lb.IP)
			}
			if lb.Hostname != "" {
				targets = append(targets, lb.Hostname)
			}
		}
	default:
		log.WithField("gateway", name).WithField("service", svc).Warn("Gloo: Proxy service type not supported")
	}
	return targets, nil
}

func (gs *glooSource) targetsFromGatewayIngress(proxy *proxy) (endpoint.Targets, error) {
	targets := make(endpoint.Targets, 0)

	for _, listener := range proxy.Spec.Listeners {
		for _, source := range listener.MetadataStatic.Source {
			if source.ResourceKind != "*v1.Gateway" {
				log.Debugf("Unsupported listener source. Expecting '*v1.Gateway', got (%s)", source.ResourceKind)
				continue
			}
			gatewayObj, err := gs.gatewayInformer.Lister().ByNamespace(source.ResourceRef.Namespace).Get(source.ResourceRef.Name)
			if err != nil {
				return nil, err
			}
			unstructuredGateway, ok := gatewayObj.(*unstructured.Unstructured)
			if !ok {
				log.Error("unexpected object: it is not *unstructured.Unstructured")
				continue
			}

			if ingressStr, ok := unstructuredGateway.GetAnnotations()[annotations.Ingress]; ok && ingressStr != "" {
				namespace, name, err := ParseIngress(ingressStr)
				if err != nil {
					return nil, fmt.Errorf("failed to parse Ingress annotation on Gateway (%s/%s): %w", unstructuredGateway.GetNamespace(), unstructuredGateway.GetName(), err)
				}
				if namespace == "" {
					namespace = unstructuredGateway.GetNamespace()
				}

				ingress, err := gs.ingressInformer.Lister().Ingresses(namespace).Get(name)
				if err != nil {
					return nil, err
				}

				for _, lb := range ingress.Status.LoadBalancer.Ingress {
					if lb.IP != "" {
						targets = append(targets, lb.IP)
					} else if lb.Hostname != "" {
						targets = append(targets, lb.Hostname)
					}
				}
			}
		}
	}
	return targets, nil
}
