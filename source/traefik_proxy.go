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
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
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

type traefikSource struct {
	dynamicKubeClient          dynamic.Interface
	kubeClient                 kubernetes.Interface
	annotationFilter           string
	namespace                  string
	ignoreHostnameAnnotation   bool
	ingressRouteInformer       kubeinformers.GenericInformer
	ingressRouteTcpInformer    kubeinformers.GenericInformer
	ingressRouteUdpInformer    kubeinformers.GenericInformer
	oldIngressRouteInformer    kubeinformers.GenericInformer
	oldIngressRouteTcpInformer kubeinformers.GenericInformer
	oldIngressRouteUdpInformer kubeinformers.GenericInformer
	unstructuredConverter      *unstructuredConverter
}

func NewTraefikSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	namespace, annotationFilter string,
	ignoreHostnameAnnotation, disableLegacy, disableNew bool) (Source, error) {
	// Use shared informer to listen for add/update/delete of Host in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	var ingressRouteInformer, ingressRouteTcpInformer, ingressRouteUdpInformer kubeinformers.GenericInformer
	var oldIngressRouteInformer, oldIngressRouteTcpInformer, oldIngressRouteUdpInformer kubeinformers.GenericInformer

	// Add default resource event handlers to properly initialize informers.
	if !disableNew {
		ingressRouteInformer = informerFactory.ForResource(ingressRouteGVR)
		ingressRouteTcpInformer = informerFactory.ForResource(ingressRouteTCPGVR)
		ingressRouteUdpInformer = informerFactory.ForResource(ingressRouteUDPGVR)
		_, _ = ingressRouteInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
		_, _ = ingressRouteTcpInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
		_, _ = ingressRouteUdpInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
	}
	if !disableLegacy {
		oldIngressRouteInformer = informerFactory.ForResource(oldIngressRouteGVR)
		oldIngressRouteTcpInformer = informerFactory.ForResource(oldIngressRouteTCPGVR)
		oldIngressRouteUdpInformer = informerFactory.ForResource(oldIngressRouteUDPGVR)
		_, _ = oldIngressRouteInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
		_, _ = oldIngressRouteTcpInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
		_, _ = oldIngressRouteUdpInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
	}

	informerFactory.Start((ctx.Done()))

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	uc, err := newTraefikUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Unstructured Converter: %w", err)
	}

	return &traefikSource{
		annotationFilter:           annotationFilter,
		ignoreHostnameAnnotation:   ignoreHostnameAnnotation,
		dynamicKubeClient:          dynamicKubeClient,
		ingressRouteInformer:       ingressRouteInformer,
		ingressRouteTcpInformer:    ingressRouteTcpInformer,
		ingressRouteUdpInformer:    ingressRouteUdpInformer,
		oldIngressRouteInformer:    oldIngressRouteInformer,
		oldIngressRouteTcpInformer: oldIngressRouteTcpInformer,
		oldIngressRouteUdpInformer: oldIngressRouteUdpInformer,
		kubeClient:                 kubeClient,
		namespace:                  namespace,
		unstructuredConverter:      uc,
	}, nil
}

func (ts *traefikSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	if ts.ingressRouteInformer != nil {
		ingressRouteEndpoints, err := ts.ingressRouteEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteEndpoints...)
	}
	if ts.oldIngressRouteInformer != nil {
		oldIngressRouteEndpoints, err := ts.oldIngressRouteEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteEndpoints...)
	}
	if ts.ingressRouteTcpInformer != nil {
		ingressRouteTcpEndpoints, err := ts.ingressRouteTCPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteTcpEndpoints...)
	}
	if ts.oldIngressRouteTcpInformer != nil {
		oldIngressRouteTcpEndpoints, err := ts.oldIngressRouteTCPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteTcpEndpoints...)
	}
	if ts.ingressRouteUdpInformer != nil {
		ingressRouteUdpEndpoints, err := ts.ingressRouteUDPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteUdpEndpoints...)
	}
	if ts.oldIngressRouteUdpInformer != nil {
		oldIngressRouteUdpEndpoints, err := ts.oldIngressRouteUDPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteUdpEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

// ingressRouteEndpoints extracts endpoints from all IngressRoute objects
func (ts *traefikSource) ingressRouteEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints[IngressRoute](
		ts.ingressRouteInformer.Lister(),
		ts.namespace,
		func(u *unstructured.Unstructured) (*IngressRoute, error) {
			typed := &IngressRoute{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.filterIngressRouteByAnnotation,
		func(r *IngressRoute, targets endpoint.Targets) []*endpoint.Endpoint {
			return ts.endpointsFromIngressRoute(r, targets)
		},
	)
}

// ingressRouteTCPEndpoints extracts endpoints from all IngressRouteTCP objects
func (ts *traefikSource) ingressRouteTCPEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.ingressRouteTcpInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRouteTCPs []*IngressRouteTCP
	for _, ingressRouteTCPObj := range irs {
		unstructuredHost, ok := ingressRouteTCPObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRouteTCP object to unstructured")
		}

		ingressRouteTCP := &IngressRouteTCP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRouteTCP, nil)
		if err != nil {
			return nil, err
		}
		ingressRouteTCPs = append(ingressRouteTCPs, ingressRouteTCP)
	}

	ingressRouteTCPs, err = ts.filterIngressRouteTcpByAnnotations(ingressRouteTCPs)
	if err != nil {
		return nil, fmt.Errorf("failed to filter IngressRouteTCP: %w", err)
	}

	for _, ingressRouteTCP := range ingressRouteTCPs {
		var targets endpoint.Targets

		targets = append(targets, annotations.TargetsFromTargetAnnotation(ingressRouteTCP.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRouteTCP.Namespace, ingressRouteTCP.Name)

		ingressEndpoints := ts.endpointsFromIngressRouteTCP(ingressRouteTCP, targets)
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRouteTCP: %s: %v", fullname, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// ingressRouteUDPEndpoints extracts endpoints from all IngressRouteUDP objects
func (ts *traefikSource) ingressRouteUDPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints[IngressRouteUDP](
		ts.ingressRouteUdpInformer.Lister(),
		ts.namespace,
		func(u *unstructured.Unstructured) (*IngressRouteUDP, error) {
			typed := &IngressRouteUDP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.filterIngressRouteUdpByAnnotations,
		ts.endpointsFromIngressRouteUDP,
	)
}

// oldIngressRouteEndpoints extracts endpoints from all IngressRoute objects
func (ts *traefikSource) oldIngressRouteEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints[IngressRoute](
		ts.oldIngressRouteInformer.Lister(),
		ts.namespace,
		func(u *unstructured.Unstructured) (*IngressRoute, error) {
			typed := &IngressRoute{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.filterIngressRouteByAnnotation,
		func(r *IngressRoute, targets endpoint.Targets) []*endpoint.Endpoint {
			return ts.endpointsFromIngressRoute(r, targets)
		},
	)
}

// oldIngressRouteTCPEndpoints extracts endpoints from all IngressRouteTCP objects
func (ts *traefikSource) oldIngressRouteTCPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints[IngressRouteTCP](
		ts.oldIngressRouteTcpInformer.Lister(),
		ts.namespace,
		func(u *unstructured.Unstructured) (*IngressRouteTCP, error) {
			typed := &IngressRouteTCP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.filterIngressRouteTcpByAnnotations,
		ts.endpointsFromIngressRouteTCP,
	)
}

// oldIngressRouteUDPEndpoints extracts endpoints from all IngressRouteUDP objects
func (ts *traefikSource) oldIngressRouteUDPEndpoints() ([]*endpoint.Endpoint, error) {
	return extractEndpoints[IngressRouteUDP](
		ts.oldIngressRouteUdpInformer.Lister(),
		ts.namespace,
		func(u *unstructured.Unstructured) (*IngressRouteUDP, error) {
			typed := &IngressRouteUDP{}
			return typed, ts.unstructuredConverter.scheme.Convert(u, typed, nil)
		},
		ts.filterIngressRouteUdpByAnnotations,
		ts.endpointsFromIngressRouteUDP,
	)
}

// filterIngressRouteByAnnotation filters a list of IngressRoute by a given annotation selector.
func (ts *traefikSource) filterIngressRouteByAnnotation(input []*IngressRoute) ([]*IngressRoute, error) {
	return filterResourcesByAnnotations(input, ts.annotationFilter, func(ir *IngressRoute) map[string]string {
		return ir.Annotations
	})
}

// filterIngressRouteTcpByAnnotations filters a list of IngressRouteTCP by a given annotation selector.
func (ts *traefikSource) filterIngressRouteTcpByAnnotations(input []*IngressRouteTCP) ([]*IngressRouteTCP, error) {
	return filterResourcesByAnnotations(input, ts.annotationFilter, func(ir *IngressRouteTCP) map[string]string {
		return ir.Annotations
	})
}

// filterIngressRouteUdpByAnnotations filters a list of IngressRoute by a given annotation selector.
func (ts *traefikSource) filterIngressRouteUdpByAnnotations(input []*IngressRouteUDP) ([]*IngressRouteUDP, error) {
	return filterResourcesByAnnotations(input, ts.annotationFilter, func(ir *IngressRouteUDP) map[string]string {
		return ir.Annotations
	})
}

// endpointsFromIngressRoute extracts the endpoints from a IngressRoute object
func (ts *traefikSource) endpointsFromIngressRoute(ingressRoute *IngressRoute, targets endpoint.Targets) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressroute/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := annotations.TTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	for _, route := range ingressRoute.Spec.Routes {
		for _, hostEntry := range traefikHostExtractor.FindAllString(route.Match, -1) {
			for _, host := range traefikValueProcessor.FindAllString(hostEntry, -1) {
				host = strings.Trim(host, "`")

				// Checking for host = * is required, as Host(`*`) can be set
				if host != "*" && host != "" {
					endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
				}
			}
		}
	}

	return endpoints
}

// endpointsFromIngressRouteTCP extracts the endpoints from a IngressRouteTCP object
func (ts *traefikSource) endpointsFromIngressRouteTCP(ingressRoute *IngressRouteTCP, targets endpoint.Targets) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressroutetcp/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := annotations.TTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	for _, route := range ingressRoute.Spec.Routes {
		for _, hostEntry := range traefikHostExtractor.FindAllString(route.Match, -1) {
			for _, host := range traefikValueProcessor.FindAllString(hostEntry, -1) {
				host = strings.Trim(host, "`")
				// Checking for host = * is required, as HostSNI(`*`) can be set
				// in the case of TLS passthrough
				if host != "*" && host != "" {
					endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
				}
			}
		}
	}

	return endpoints
}

// endpointsFromIngressRouteUDP extracts the endpoints from a IngressRouteUDP object
func (ts *traefikSource) endpointsFromIngressRouteUDP(ingressRoute *IngressRouteUDP, targets endpoint.Targets) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressrouteudp/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := annotations.TTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return endpoints
}

func (ts *traefikSource) AddEventHandler(ctx context.Context, handler func()) {
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	log.Debug("Adding event handler for IngressRoute")
	if ts.ingressRouteInformer != nil {
		_, _ = ts.ingressRouteInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
	if ts.oldIngressRouteInformer != nil {
		_, _ = ts.oldIngressRouteInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
	log.Debug("Adding event handler for IngressRouteTCP")
	if ts.ingressRouteTcpInformer != nil {
		_, _ = ts.ingressRouteTcpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
	if ts.oldIngressRouteTcpInformer != nil {
		_, _ = ts.oldIngressRouteTcpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
	log.Debug("Adding event handler for IngressRouteUDP")
	if ts.ingressRouteUdpInformer != nil {
		_, _ = ts.ingressRouteUdpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
	if ts.oldIngressRouteUdpInformer != nil {
		_, _ = ts.oldIngressRouteUdpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
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

// extractEndpoints is a generic function that extracts endpoints from Kubernetes resources.
// It performs the following steps:
// 1. Lists all objects in the specified namespace using the provided informer.
// 2. Converts the unstructured objects to the desired type using the convertFunc.
// 3. Filters the converted objects based on the provided filterFunc.
// 4. Generates endpoints for each filtered object using the generateEndpoints function.
// Returns a list of generated endpoints or an error if any step fails.
func extractEndpoints[T any](
	informer cache.GenericLister,
	namespace string,
	convertFunc func(*unstructured.Unstructured) (*T, error),
	filterFunc func([]*T) ([]*T, error),
	generateEndpoints func(*T, endpoint.Targets) []*endpoint.Endpoint,
) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	objs, err := informer.ByNamespace(namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var typedObjs []*T
	for _, obj := range objs {
		unstructuredObj, ok := obj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("failed to cast to unstructured.Unstructured")
		}

		typed, err := convertFunc(unstructuredObj)
		if err != nil {
			return nil, err
		}
		typedObjs = append(typedObjs, typed)
	}

	typedObjs, err = filterFunc(typedObjs)
	if err != nil {
		return nil, err
	}

	for _, item := range typedObjs {
		targets := annotations.TargetsFromTargetAnnotation(getAnnotations(item))

		name := getObjectFullName(item)
		ingressEndpoints := generateEndpoints(item, targets)

		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", name)
			continue
		}

		log.Debugf("Endpoints generated from %s: %v", name, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// filterResourcesByAnnotations filters a list of resources based on a given annotation selector.
// It performs the following steps:
// 1. Parses the annotation filter into a label selector.
// 2. Converts the label selector into a Kubernetes selector.
// 3. If the selector is empty, returns the original list of resources.
// 4. Iterates through the resources and matches their annotations against the selector.
// 5. Returns the filtered list of resources or an error if any step fails.
func filterResourcesByAnnotations[T any](resources []*T, annotationFilter string, getAnnotations func(*T) map[string]string) ([]*T, error) {
	labelSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	if selector.Empty() {
		return resources, nil
	}

	var filteredList []*T
	for _, resource := range resources {
		if selector.Matches(labels.Set(getAnnotations(resource))) {
			filteredList = append(filteredList, resource)
		}
	}

	return filteredList, nil
}

func getAnnotations(obj interface{}) map[string]string {
	switch o := obj.(type) {
	case *IngressRouteUDP:
		return o.Annotations
	case *IngressRoute:
		return o.Annotations
	case *IngressRouteTCP:
		return o.Annotations
	default:
		return nil
	}
}

func getObjectFullName(obj interface{}) string {
	switch o := obj.(type) {
	case *IngressRouteUDP:
		return fmt.Sprintf("%s/%s", o.Namespace, o.Name)
	case *IngressRoute:
		return fmt.Sprintf("%s/%s", o.Namespace, o.Name)
	case *IngressRouteTCP:
		return fmt.Sprintf("%s/%s", o.Namespace, o.Name)
	default:
		return ""
	}
}
