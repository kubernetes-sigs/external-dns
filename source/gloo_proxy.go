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
	"strings"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/external-dns/endpoint"
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
)

// Basic redefinition of "Proxy" CRD : https://github.com/solo-io/gloo/blob/v1.4.6/projects/gloo/pkg/api/v1/proxy.pb.go
type proxy struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec            proxySpec         `json:"spec,omitempty"`
}

type proxySpec struct {
	Listeners []proxySpecListener `json:"listeners,omitempty"`
}

type proxySpecListener struct {
	HTTPListener   proxySpecHTTPListener   `json:"httpListener,omitempty"`
	HybridListener proxySpecHybridListener `json:"hybridListener,omitempty"`
}

type proxySpecHybridListener struct {
	MatchedListeners []proxyMatchedListeners `json:"matchedListeners,omitempty"`
}

type proxyMatchedListeners struct {
	HTTPListener proxySpecHTTPListener `json:"httpListener,omitempty"`
}

type proxySpecHTTPListener struct {
	VirtualHosts []proxyVirtualHost `json:"virtualHosts,omitempty"`
}

type proxyVirtualHost struct {
	Domains        []string                       `json:"domains,omitempty"`
	Metadata       proxyVirtualHostMetadata       `json:"metadata,omitempty"`
	MetadataStatic proxyVirtualHostMetadataStatic `json:"metadataStatic,omitempty"`
}

type proxyVirtualHostMetadata struct {
	Source []proxyVirtualHostMetadataSource `json:"sources,omitempty"`
}

type proxyVirtualHostMetadataStatic struct {
	Source []proxyVirtualHostMetadataSourceStatic `json:"sources,omitempty"`
}

type proxyVirtualHostMetadataSource struct {
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type proxyVirtualHostMetadataSourceStatic struct {
	Kind        string      `json:"resourceKind,omitempty"`
	ResourceRef resourceRef `json:"resourceRef,omitempty"`
}

type resourceRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type glooSource struct {
	dynamicKubeClient dynamic.Interface
	kubeClient        kubernetes.Interface
	glooNamespace     string
}

// NewGlooSource creates a new glooSource with the given config
func NewGlooSource(dynamicKubeClient dynamic.Interface, kubeClient kubernetes.Interface, glooNamespace string) (Source, error) {
	return &glooSource{
		dynamicKubeClient,
		kubeClient,
		glooNamespace,
	}, nil
}

// virtualHosts aggregates all virtual hosts from the available listener types.
// Only 1 of HybridListener or HTTPListener will be set
func (l proxySpecListener) virtualHosts() []proxyVirtualHost {
	virtualHosts := l.HTTPListener.VirtualHosts
	for _, hybridMatchedListener := range l.HybridListener.MatchedListeners {
		virtualHosts = append(virtualHosts, hybridMatchedListener.HTTPListener.VirtualHosts...)
	}
	return virtualHosts
}

// sources aggregates all sources within a proxyVirtualHost.
// Only 1 of Metadata for MetadataStatic will be set
func (vh proxyVirtualHost) sources() []proxyVirtualHostMetadataSource {
	sources := vh.Metadata.Source
	for _, srcStatic := range vh.MetadataStatic.Source {
		// Use the same struct type as proxyVirtualHostMetadataSource to enable reuse
		// and remove additional logic branches due to object property layout differences only
		sources = append(sources, proxyVirtualHostMetadataSource{
			Kind:      srcStatic.Kind,
			Name:      srcStatic.ResourceRef.Name,
			Namespace: srcStatic.ResourceRef.Namespace,
		})
	}
	return sources
}

func (gs *glooSource) AddEventHandler(ctx context.Context, handler func()) {
}

// Endpoints returns endpoint objects
func (gs *glooSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	proxies, err := gs.dynamicKubeClient.Resource(proxyGVR).Namespace(gs.glooNamespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, obj := range proxies.Items {
		proxy := proxy{}
		jsonString, err := obj.MarshalJSON()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonString, &proxy)
		if err != nil {
			return nil, err
		}
		log.Debugf("Gloo: Find %s proxy", proxy.Metadata.Name)
		proxyTargets, err := gs.proxyTargets(ctx, proxy.Metadata.Name)
		if err != nil {
			return nil, err
		}
		log.Debugf("Gloo[%s]: Find %d target(s) (%+v)", proxy.Metadata.Name, len(proxyTargets), proxyTargets)
		proxyEndpoints, err := gs.generateEndpointsFromProxy(ctx, &proxy, proxyTargets)
		if err != nil {
			return nil, err
		}
		log.Debugf("Gloo[%s]: Generate %d endpoint(s)", proxy.Metadata.Name, len(proxyEndpoints))
		endpoints = append(endpoints, proxyEndpoints...)
	}
	return endpoints, nil
}

func (gs *glooSource) generateEndpointsFromProxy(ctx context.Context, proxy *proxy, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}
	for _, listener := range proxy.Spec.Listeners {
		for _, virtualHost := range listener.virtualHosts() {
			annotations, err := gs.annotationsFromProxySource(ctx, virtualHost)
			if err != nil {
				return nil, err
			}
			ttl, err := getTTLFromAnnotations(annotations)
			if err != nil {
				return nil, err
			}
			providerSpecific, setIdentifier := getProviderSpecificAnnotations(annotations)
			for _, domain := range virtualHost.Domains {
				endpoints = append(endpoints, endpointsForHostname(strings.TrimSuffix(domain, "."), targets, ttl, providerSpecific, setIdentifier)...)
			}
		}
	}
	return endpoints, nil
}

func (gs *glooSource) annotationsFromProxySource(ctx context.Context, virtualHost proxyVirtualHost) (map[string]string, error) {
	annotations := map[string]string{}
	for _, src := range virtualHost.sources() {
		kind := sourceKind(src.Kind)
		if kind != nil {
			source, err := gs.dynamicKubeClient.Resource(*kind).Namespace(src.Namespace).Get(ctx, src.Name, metav1.GetOptions{})
			if err != nil {
				return nil, err
			}
			for key, value := range source.GetAnnotations() {
				annotations[key] = value
			}
		}
	}
	return annotations, nil
}

func (gs *glooSource) proxyTargets(ctx context.Context, name string) (endpoint.Targets, error) {
	svc, err := gs.kubeClient.CoreV1().Services(gs.glooNamespace).Get(ctx, name, metav1.GetOptions{})
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

func sourceKind(kind string) *schema.GroupVersionResource {
	switch kind {
	case "*v1.VirtualService":
		return &virtualServiceGVR
	}
	return nil
}
