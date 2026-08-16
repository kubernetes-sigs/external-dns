/*
Copyright 2022 The Kubernetes Authors.

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
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/external-dns/source/template"
	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
)

var (
	ingressRouteGVR = schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}
	ingressRouteTCPGVR = schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressroutetcps",
	}
	ingressRouteUDPGVR = schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressrouteudps",
	}
	// TODO: traefik.containo.us CRDs were removed in Traefik v3 (released 2024).
	// Traefik v2 active support ended 2025-04-29; security support ends 2026-02-01.
	// Remove these GVRs and the --traefik-enable-legacy flag after 2026-02-01.
	// See https://doc.traefik.io/traefik/deprecation/releases/
	oldIngressRouteGVR = schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}
	oldIngressRouteTCPGVR = schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressroutetcps",
	}
	oldIngressRouteUDPGVR = schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressrouteudps",
	}
)

var (
	traefikHostExtractor  = regexp.MustCompile(`(?:HostSNI|HostHeader|Host)\s*\(\s*(\x60.*?\x60)\s*\)`)
	traefikValueProcessor = regexp.MustCompile(`\x60([^,\x60]+)\x60`)
)

// +externaldns:source:name=traefik-proxy
// +externaldns:source:category=Ingress Controllers
// +externaldns:source:description=Creates DNS entries from Traefik IngressRoute, IngressRouteTCP, and IngressRouteUDP resources
// +externaldns:source:resources=IngressRoute.traefik.io,IngressRouteTCP.traefik.io,IngressRouteUDP.traefik.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single,multiple
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
type traefikSource struct {
	dynamicKubeClient           dynamic.Interface
	kubeClient                  kubernetes.Interface
	ignoreHostnameAnnotation    bool
	templateEngine              template.Engine
	ingressRouteInformers       *informers.Informers[kubeinformers.GenericInformer]
	ingressRouteTcpInformers    *informers.Informers[kubeinformers.GenericInformer]
	ingressRouteUdpInformers    *informers.Informers[kubeinformers.GenericInformer]
	oldIngressRouteInformers    *informers.Informers[kubeinformers.GenericInformer]
	oldIngressRouteTcpInformers *informers.Informers[kubeinformers.GenericInformer]
	oldIngressRouteUdpInformers *informers.Informers[kubeinformers.GenericInformer]
	unstructuredConverter       *unstructuredConverter
}

func NewTraefikSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	cfg *Config,
) (Source, error) {
	// Use shared informer to listen for add/update/delete of Host in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	factories := informers.NewFactories(cfg.Namespaces, func(namespace string) dynamicinformer.DynamicSharedInformerFactory {
		return dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	})
	forResource := func(gvr schema.GroupVersionResource) *informers.Informers[kubeinformers.GenericInformer] {
		return informers.Map(factories, func(_ string, factory dynamicinformer.DynamicSharedInformerFactory) kubeinformers.GenericInformer {
			return factory.ForResource(gvr)
		})
	}
	var ingressRouteInformers, ingressRouteTcpInformers, ingressRouteUdpInformers *informers.Informers[kubeinformers.GenericInformer]
	var oldIngressRouteInformers, oldIngressRouteTcpInformers, oldIngressRouteUdpInformers *informers.Informers[kubeinformers.GenericInformer]

	// Add default resource event handlers to properly initialize informers.
	indexerOpts := informers.IndexerWithOptions[*unstructured.Unstructured](
		informers.IndexSelectorWithAnnotationFilter(cfg.AnnotationFilter),
		informers.IndexSelectorWithLabelSelector(cfg.LabelFilter),
		informers.IndexSelectorWithConditions(annotations.IsControllerMatch[*unstructured.Unstructured]),
	)
	if !cfg.TraefikDisableNew {
		ingressRouteInformers = forResource(ingressRouteGVR)
		ingressRouteTcpInformers = forResource(ingressRouteTCPGVR)
		ingressRouteUdpInformers = forResource(ingressRouteUDPGVR)
		ingressRouteInformers.MustSetTransform(informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))
		ingressRouteTcpInformers.MustSetTransform(informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))
		ingressRouteUdpInformers.MustSetTransform(informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))
		ingressRouteInformers.MustAddIndexers(indexerOpts)
		ingressRouteTcpInformers.MustAddIndexers(indexerOpts)
		ingressRouteUdpInformers.MustAddIndexers(indexerOpts)
		ingressRouteInformers.MustAddEventHandler(informers.DefaultEventHandler())
		ingressRouteTcpInformers.MustAddEventHandler(informers.DefaultEventHandler())
		ingressRouteUdpInformers.MustAddEventHandler(informers.DefaultEventHandler())
	}
	if cfg.TraefikEnableLegacy {
		oldIngressRouteInformers = forResource(oldIngressRouteGVR)
		oldIngressRouteTcpInformers = forResource(oldIngressRouteTCPGVR)
		oldIngressRouteUdpInformers = forResource(oldIngressRouteUDPGVR)
		oldIngressRouteInformers.MustSetTransform(informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))
		oldIngressRouteTcpInformers.MustSetTransform(informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))
		oldIngressRouteUdpInformers.MustSetTransform(informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))
		oldIngressRouteInformers.MustAddIndexers(indexerOpts)
		oldIngressRouteTcpInformers.MustAddIndexers(indexerOpts)
		oldIngressRouteUdpInformers.MustAddIndexers(indexerOpts)
		oldIngressRouteInformers.MustAddEventHandler(informers.DefaultEventHandler())
		oldIngressRouteTcpInformers.MustAddEventHandler(informers.DefaultEventHandler())
		oldIngressRouteUdpInformers.MustAddEventHandler(informers.DefaultEventHandler())
	}

	factories.Start(ctx.Done())

	// wait for the local caches to be populated.
	if err := informers.WaitForDynamicCacheSyncAll(ctx, factories); err != nil {
		return nil, err
	}

	uc, err := newTraefikUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Unstructured Converter: %w", err)
	}

	return &traefikSource{
		ignoreHostnameAnnotation:    cfg.IgnoreHostnameAnnotation,
		templateEngine:              cfg.TemplateEngine,
		dynamicKubeClient:           dynamicKubeClient,
		ingressRouteInformers:       ingressRouteInformers,
		ingressRouteTcpInformers:    ingressRouteTcpInformers,
		ingressRouteUdpInformers:    ingressRouteUdpInformers,
		oldIngressRouteInformers:    oldIngressRouteInformers,
		oldIngressRouteTcpInformers: oldIngressRouteTcpInformers,
		oldIngressRouteUdpInformers: oldIngressRouteUdpInformers,
		kubeClient:                  kubeClient,
		unstructuredConverter:       uc,
	}, nil
}

func (ts *traefikSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	if ts.ingressRouteInformers != nil {
		ingressRouteEndpoints, err := ts.ingressRouteEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteEndpoints...)
	}
	if ts.oldIngressRouteInformers != nil {
		oldIngressRouteEndpoints, err := ts.oldIngressRouteEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteEndpoints...)
	}
	if ts.ingressRouteTcpInformers != nil {
		ingressRouteTcpEndpoints, err := ts.ingressRouteTCPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteTcpEndpoints...)
	}
	if ts.oldIngressRouteTcpInformers != nil {
		oldIngressRouteTcpEndpoints, err := ts.oldIngressRouteTCPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteTcpEndpoints...)
	}
	if ts.ingressRouteUdpInformers != nil {
		ingressRouteUdpEndpoints, err := ts.ingressRouteUDPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteUdpEndpoints...)
	}
	if ts.oldIngressRouteUdpInformers != nil {
		oldIngressRouteUdpEndpoints, err := ts.oldIngressRouteUDPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteUdpEndpoints...)
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

// ingressRouteEndpoints extracts endpoints from all IngressRoute objects
func (ts *traefikSource) ingressRouteEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints(
		ts.ingressRouteInformers.Indexers(),
		func(u *unstructured.Unstructured) (*IngressRoute, error) {
			typed := &IngressRoute{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.endpointsFromIngressRoute,
	)
}

// ingressRouteTCPEndpoints extracts endpoints from all IngressRouteTCP objects
func (ts *traefikSource) ingressRouteTCPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints(
		ts.ingressRouteTcpInformers.Indexers(),
		func(u *unstructured.Unstructured) (*IngressRouteTCP, error) {
			typed := &IngressRouteTCP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.endpointsFromIngressRouteTCP,
	)
}

// ingressRouteUDPEndpoints extracts endpoints from all IngressRouteUDP objects
func (ts *traefikSource) ingressRouteUDPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints(
		ts.ingressRouteUdpInformers.Indexers(),
		func(u *unstructured.Unstructured) (*IngressRouteUDP, error) {
			typed := &IngressRouteUDP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.endpointsFromIngressRouteUDP,
	)
}

// oldIngressRouteEndpoints extracts endpoints from all IngressRoute objects
func (ts *traefikSource) oldIngressRouteEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints(
		ts.oldIngressRouteInformers.Indexers(),
		func(u *unstructured.Unstructured) (*IngressRoute, error) {
			typed := &IngressRoute{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.endpointsFromIngressRoute,
	)
}

// oldIngressRouteTCPEndpoints extracts endpoints from all IngressRouteTCP objects
func (ts *traefikSource) oldIngressRouteTCPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints(
		ts.oldIngressRouteTcpInformers.Indexers(),
		func(u *unstructured.Unstructured) (*IngressRouteTCP, error) {
			typed := &IngressRouteTCP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.endpointsFromIngressRouteTCP,
	)
}

// oldIngressRouteUDPEndpoints extracts endpoints from all IngressRouteUDP objects
func (ts *traefikSource) oldIngressRouteUDPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints(
		ts.oldIngressRouteUdpInformers.Indexers(),
		func(u *unstructured.Unstructured) (*IngressRouteUDP, error) {
			typed := &IngressRouteUDP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.endpointsFromIngressRouteUDP,
	)
}

// endpointsFromIngressRoute extracts the endpoints from a IngressRoute object
func (ts *traefikSource) endpointsFromIngressRoute(ingressRoute *IngressRoute, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressroute/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := annotations.TTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpoint.EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	for _, route := range ingressRoute.Spec.Routes {
		for _, hostEntry := range traefikHostExtractor.FindAllString(route.Match, -1) {
			for _, host := range traefikValueProcessor.FindAllString(hostEntry, -1) {
				host = strings.Trim(host, "`")

				// Checking for host = * is required, as Host(`*`) can be set
				if host != "*" && host != "" {
					endpoints = append(endpoints, endpoint.EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
				}
			}
		}
	}

	return ts.templateEngine.ApplyTemplates(endpoints, ingressRoute)
}

// endpointsFromIngressRouteTCP extracts the endpoints from a IngressRouteTCP object
func (ts *traefikSource) endpointsFromIngressRouteTCP(ingressRoute *IngressRouteTCP, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressroutetcp/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := annotations.TTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpoint.EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	for _, route := range ingressRoute.Spec.Routes {
		for _, hostEntry := range traefikHostExtractor.FindAllString(route.Match, -1) {
			for _, host := range traefikValueProcessor.FindAllString(hostEntry, -1) {
				host = strings.Trim(host, "`")
				// Checking for host = * is required, as HostSNI(`*`) can be set
				// in the case of TLS passthrough
				if host != "*" && host != "" {
					endpoints = append(endpoints, endpoint.EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
				}
			}
		}
	}

	return ts.templateEngine.ApplyTemplates(endpoints, ingressRoute)
}

// endpointsFromIngressRouteUDP extracts the endpoints from a IngressRouteUDP object
func (ts *traefikSource) endpointsFromIngressRouteUDP(ingressRoute *IngressRouteUDP, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressrouteudp/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := annotations.TTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpoint.EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return ts.templateEngine.ApplyTemplates(endpoints, ingressRoute)
}

func (ts *traefikSource) AddEventHandler(_ context.Context, handler func()) {
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	log.Debug("Adding event handler for IngressRoute")
	if ts.ingressRouteInformers != nil {
		ts.ingressRouteInformers.MustAddEventHandler(eventHandlerFunc(handler))
	}
	if ts.oldIngressRouteInformers != nil {
		ts.oldIngressRouteInformers.MustAddEventHandler(eventHandlerFunc(handler))
	}
	log.Debug("Adding event handler for IngressRouteTCP")
	if ts.ingressRouteTcpInformers != nil {
		ts.ingressRouteTcpInformers.MustAddEventHandler(eventHandlerFunc(handler))
	}
	if ts.oldIngressRouteTcpInformers != nil {
		ts.oldIngressRouteTcpInformers.MustAddEventHandler(eventHandlerFunc(handler))
	}
	log.Debug("Adding event handler for IngressRouteUDP")
	if ts.ingressRouteUdpInformers != nil {
		ts.ingressRouteUdpInformers.MustAddEventHandler(eventHandlerFunc(handler))
	}
	if ts.oldIngressRouteUdpInformers != nil {
		ts.oldIngressRouteUdpInformers.MustAddEventHandler(eventHandlerFunc(handler))
	}
}

// newTraefikUnstructuredConverter returns a new unstructuredConverter initialized
func newTraefikUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Add the core types we need
	uc.scheme.AddKnownTypes(ingressRouteGVR.GroupVersion(), &IngressRoute{}, &IngressRouteList{})
	uc.scheme.AddKnownTypes(oldIngressRouteGVR.GroupVersion(), &IngressRoute{}, &IngressRouteList{})
	uc.scheme.AddKnownTypes(ingressRouteTCPGVR.GroupVersion(), &IngressRouteTCP{}, &IngressRouteTCPList{})
	uc.scheme.AddKnownTypes(oldIngressRouteTCPGVR.GroupVersion(), &IngressRouteTCP{}, &IngressRouteTCPList{})
	uc.scheme.AddKnownTypes(ingressRouteUDPGVR.GroupVersion(), &IngressRouteUDP{}, &IngressRouteUDPList{})
	uc.scheme.AddKnownTypes(oldIngressRouteUDPGVR.GroupVersion(), &IngressRouteUDP{}, &IngressRouteUDPList{})
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}

// Basic redefinition of Traefik 2's CRD: https://github.com/traefik/traefik/tree/v2.8.7/pkg/provider/kubernetes/crd/traefik/v1alpha1

// traefikIngressRouteSpec defines the desired state of IngressRoute.
type traefikIngressRouteSpec struct {
	// Routes defines the list of routes.
	Routes []traefikRoute `json:"routes"`
}

// traefikRoute holds the HTTP route configuration.
type traefikRoute struct {
	// Match defines the router's rule.
	// More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#rule
	Match string `json:"match"`
}

// IngressRoute is the CRD implementation of a Traefik HTTP Router.
type IngressRoute struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	Spec traefikIngressRouteSpec `json:"spec"`
}

// IngressRouteList is a collection of IngressRoute.
type IngressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of IngressRoute.
	Items []IngressRoute `json:"items"`
}

// traefikIngressRouteTCPSpec defines the desired state of IngressRouteTCP.
type traefikIngressRouteTCPSpec struct {
	Routes []traefikRouteTCP `json:"routes"`
}

// traefikRouteTCP holds the TCP route configuration.
type traefikRouteTCP struct {
	// Match defines the router's rule.
	// More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#rule_1
	Match string `json:"match"`
}

// IngressRouteTCP is the CRD implementation of a Traefik TCP Router.
type IngressRouteTCP struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	Spec traefikIngressRouteTCPSpec `json:"spec"`
}

// IngressRouteTCPList is a collection of IngressRouteTCP.
type IngressRouteTCPList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of IngressRouteTCP.
	Items []IngressRouteTCP `json:"items"`
}

// IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
type IngressRouteUDP struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`
}

// IngressRouteUDPList is a collection of IngressRouteUDP.
type IngressRouteUDPList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of IngressRouteUDP.
	Items []IngressRouteUDP `json:"items"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRoute) DeepCopyInto(out *IngressRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRoute.
func (in *IngressRoute) DeepCopy() *IngressRoute {
	if in == nil {
		return nil
	}
	out := new(IngressRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteList) DeepCopyInto(out *IngressRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteList.
func (in *IngressRouteList) DeepCopy() *IngressRouteList {
	if in == nil {
		return nil
	}
	out := new(IngressRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikIngressRouteSpec) DeepCopyInto(out *traefikIngressRouteSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]traefikRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteSpec.
func (in *traefikIngressRouteSpec) DeepCopy() *traefikIngressRouteSpec {
	if in == nil {
		return nil
	}
	out := new(traefikIngressRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikRoute) DeepCopyInto(out *traefikRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *traefikRoute) DeepCopy() *traefikRoute {
	if in == nil {
		return nil
	}
	out := new(traefikRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteTCP) DeepCopyInto(out *IngressRouteTCP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteTCP.
func (in *IngressRouteTCP) DeepCopy() *IngressRouteTCP {
	if in == nil {
		return nil
	}
	out := new(IngressRouteTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteTCP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteTCPList) DeepCopyInto(out *IngressRouteTCPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressRouteTCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteTCPList.
func (in *IngressRouteTCPList) DeepCopy() *IngressRouteTCPList {
	if in == nil {
		return nil
	}
	out := new(IngressRouteTCPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteTCPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikIngressRouteTCPSpec) DeepCopyInto(out *traefikIngressRouteTCPSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]traefikRouteTCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteTCPSpec.
func (in *traefikIngressRouteTCPSpec) DeepCopy() *traefikIngressRouteTCPSpec {
	if in == nil {
		return nil
	}
	out := new(traefikIngressRouteTCPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikRouteTCP) DeepCopyInto(out *traefikRouteTCP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTCP.
func (in *traefikRouteTCP) DeepCopy() *traefikRouteTCP {
	if in == nil {
		return nil
	}
	out := new(traefikRouteTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteUDP) DeepCopyInto(out *IngressRouteUDP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteUDP.
func (in *IngressRouteUDP) DeepCopy() *IngressRouteUDP {
	if in == nil {
		return nil
	}
	out := new(IngressRouteUDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteUDP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteUDPList) DeepCopyInto(out *IngressRouteUDPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressRouteUDP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteUDPList.
func (in *IngressRouteUDPList) DeepCopy() *IngressRouteUDPList {
	if in == nil {
		return nil
	}
	out := new(IngressRouteUDPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteUDPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// GetAnnotations returns the annotations of the IngressRoute.
func (in *IngressRoute) GetAnnotations() map[string]string {
	return in.Annotations
}

// GetAnnotations returns the annotations of the IngressRouteTCP.
func (in *IngressRouteTCP) GetAnnotations() map[string]string {
	return in.Annotations
}

// GetAnnotations returns the annotations of the IngressRouteUDP.
func (in *IngressRouteUDP) GetAnnotations() map[string]string {
	return in.Annotations
}

// extractEndpoints is a generic function that extracts endpoints from Kubernetes resources.
// It performs the following steps:
// 1. Lists all objects admitted by the indexer (annotation + label filters applied at index time).
// 2. Converts the unstructured objects to the desired type using the convertFunc.
// 3. Generates endpoints for each object using the generateEndpoints function.
// Returns a list of generated endpoints or an error if any step fails.
func extractEndpoints[T interface {
	annotations.AnnotatedObject
	runtime.Object
}](
	indexers []cache.Indexer,
	convertFunc func(*unstructured.Unstructured) (T, error),
	generateEndpoints func(T, endpoint.Targets) ([]*endpoint.Endpoint, error),
) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	for _, obj := range informers.ListIndexedAll[*unstructured.Unstructured](indexers...) {
		typed, err := convertFunc(obj)
		if err != nil {
			return nil, err
		}
		typed.GetObjectKind().SetGroupVersionKind(obj.GetObjectKind().GroupVersionKind())

		targets := annotations.TargetsFromTargetAnnotation(typed.GetAnnotations())
		name := getObjectFullName(typed)

		ingressEndpoints, err := generateEndpoints(typed, targets)
		if err != nil {
			return nil, err
		}

		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", name)
			continue
		}

		// All traefik route kinds map to the traefik-proxy source. The concrete
		// CRD types satisfy client.Object; the assertion guards the generic T.
		if cObj, ok := any(typed).(client.Object); ok {
			endpoint.AttachRefObject(ingressEndpoints, events.NewObjectReference(cObj, types.TraefikProxy))
		}

		log.Debugf("Endpoints generated from %s: %v", name, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

func getObjectFullName(obj any) string {
	if m, ok := obj.(metav1.Object); ok {
		return m.GetNamespace() + "/" + m.GetName()
	}
	return ""
}
