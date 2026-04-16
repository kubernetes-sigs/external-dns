/*
Copyright 2021 The Kubernetes Authors.

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
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	gwinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	informers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
)

const (
	gatewayGroup                               = "gateway.networking.k8s.io"
	gatewayKind                                = "Gateway"
	listenerSetKind                            = "ListenerSet"
	gatewayHostnameSourceAnnotationOnlyValue   = "annotation-only"
	gatewayHostnameSourceDefinedHostsOnlyValue = "defined-hosts-only"
)

type gatewayRoute interface {
	// Object returns the underlying route object to be used by templates.
	Object() kubeObject
	// Metadata returns the route's metadata.
	Metadata() *metav1.ObjectMeta
	// Hostnames returns the route's specified hostnames.
	Hostnames() []v1.Hostname
	// ParentRefs returns the route's parent references as defined in the route spec.
	ParentRefs() []v1.ParentReference
	// Protocol returns the route's protocol type.
	Protocol() v1.ProtocolType
	// RouteStatus returns the route's common status.
	RouteStatus() v1.RouteStatus
}

type newGatewayRouteInformerFunc func(gwinformers.SharedInformerFactory) gatewayRouteInformer

type gatewayRouteInformer interface {
	List(namespace string, selector labels.Selector) ([]gatewayRoute, error)
	Informer() cache.SharedIndexInformer
}

func newGatewayInformerFactory(client gateway.Interface, namespace string, labelSelector labels.Selector) gwinformers.SharedInformerFactory {
	var opts []gwinformers.SharedInformerOption
	if namespace != "" {
		opts = append(opts, gwinformers.WithNamespace(namespace))
	}
	if labelSelector != nil && !labelSelector.Empty() {
		lbls := labelSelector.String()
		opts = append(opts, gwinformers.WithTweakListOptions(func(o *metav1.ListOptions) {
			o.LabelSelector = lbls
		}))
	}
	return gwinformers.NewSharedInformerFactoryWithOptions(client, 0, opts...)
}

// gatewayRouteSource is an implementation of Source for Gateway API Route objects.
//
// +externaldns:source:name=gateway-httproute
// +externaldns:source:category=Gateway API
// +externaldns:source:description=Creates DNS entries from Gateway API HTTPRoute resources
// +externaldns:source:resources=HTTPRoute.gateway.networking.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
//
// +externaldns:source:name=gateway-grpcroute
// +externaldns:source:category=Gateway API
// +externaldns:source:description=Creates DNS entries from Gateway API GRPCRoute resources
// +externaldns:source:resources=GRPCRoute.gateway.networking.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
//
// +externaldns:source:name=gateway-tcproute
// +externaldns:source:category=Gateway API
// +externaldns:source:description=Creates DNS entries from Gateway API TCPRoute resources
// +externaldns:source:resources=TCPRoute.gateway.networking.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
//
// +externaldns:source:name=gateway-tlsroute
// +externaldns:source:category=Gateway API
// +externaldns:source:description=Creates DNS entries from Gateway API TLSRoute resources
// +externaldns:source:resources=TLSRoute.gateway.networking.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
//
// +externaldns:source:name=gateway-udproute
// +externaldns:source:category=Gateway API
// +externaldns:source:description=Creates DNS entries from Gateway API UDPRoute resources
// +externaldns:source:resources=UDPRoute.gateway.networking.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
type gatewayRouteSource struct {
	gwName      string
	gwNamespace string
	gwLabels    labels.Selector
	gwInformer  informers_v1.GatewayInformer
	lsInformer  informers_v1.ListenerSetInformer

	rtKind        string
	rtNamespace   string
	rtLabels      labels.Selector
	rtAnnotations labels.Selector
	rtInformer    gatewayRouteInformer

	nsInformer coreinformers.NamespaceInformer

	templateEngine           template.Engine
	ignoreHostnameAnnotation bool
}

func newGatewayRouteSource(
	ctx context.Context,
	clients ClientGenerator,
	config *Config,
	kind string,
	newInformerFn newGatewayRouteInformerFunc) (Source, error) {
	gwLabels, err := getLabelSelector(config.GatewayLabelFilter)
	if err != nil {
		return nil, err
	}
	rtLabels := config.LabelFilter
	if rtLabels == nil {
		rtLabels = labels.Everything()
	}
	rtAnnotations, err := getLabelSelector(config.AnnotationFilter)
	if err != nil {
		return nil, err
	}

	client, err := clients.GatewayClient()
	if err != nil {
		return nil, err
	}

	gwInformerFactory := newGatewayInformerFactory(client, config.GatewayNamespace, gwLabels)
	gwInformer := gwInformerFactory.Gateway().V1().Gateways() // TODO: Gateway informer should be shared across gateway sources.
	gwInformer.Informer()                                     // Register with factory before starting.

	var lsInformer informers_v1.ListenerSetInformer
	lsInformerFactory := gwInformerFactory
	if config.GatewayListenerSets {
		// Gateway filters should apply only to Gateways, not ListenerSets.
		if config.GatewayNamespace != "" || (gwLabels != nil && !gwLabels.Empty()) {
			lsInformerFactory = newGatewayInformerFactory(client, "", nil)
		}
		lsInformer = lsInformerFactory.Gateway().V1().ListenerSets() // TODO: ListenerSet informer should be shared across gateway sources.
		lsInformer.Informer()                                        // Register with factory before starting.
	}

	rtInformerFactory := gwInformerFactory
	if config.Namespace != config.GatewayNamespace || !selectorsEqual(rtLabels, gwLabels) {
		rtInformerFactory = newGatewayInformerFactory(client, config.Namespace, rtLabels)
	}
	rtInformer := newInformerFn(rtInformerFactory)
	rtInformer.Informer() // Register with factory before starting.

	kubeClient, err := clients.KubeClient()
	if err != nil {
		return nil, err
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, 0)
	nsInformer := kubeInformerFactory.Core().V1().Namespaces() // TODO: Namespace informer should be shared across gateway sources.
	nsInformer.Informer()                                      // Register with factory before starting.

	gwInformerFactory.Start(ctx.Done())
	if lsInformerFactory != gwInformerFactory {
		lsInformerFactory.Start(ctx.Done())
	}
	kubeInformerFactory.Start(ctx.Done())
	if rtInformerFactory != gwInformerFactory {
		rtInformerFactory.Start(ctx.Done())
	}
	if err := informers.WaitForCacheSync(ctx, gwInformerFactory); err != nil {
		return nil, err
	}
	if lsInformer != nil && lsInformerFactory != gwInformerFactory {
		if err := informers.WaitForCacheSync(ctx, lsInformerFactory); err != nil {
			return nil, err
		}
	}
	if err := informers.WaitForCacheSync(ctx, rtInformerFactory); err != nil {
		return nil, err
	}
	if err := informers.WaitForCacheSync(ctx, kubeInformerFactory); err != nil {
		return nil, err
	}

	src := &gatewayRouteSource{
		gwName:      config.GatewayName,
		gwNamespace: config.GatewayNamespace,
		gwLabels:    gwLabels,
		gwInformer:  gwInformer,
		lsInformer:  lsInformer,

		rtKind:        kind,
		rtNamespace:   config.Namespace,
		rtLabels:      rtLabels,
		rtAnnotations: rtAnnotations,
		rtInformer:    rtInformer,

		nsInformer: nsInformer,

		templateEngine:           config.TemplateEngine,
		ignoreHostnameAnnotation: config.IgnoreHostnameAnnotation,
	}
	return src, nil
}

func (src *gatewayRouteSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debugf("Adding event handlers for %s", src.rtKind)
	eventHandler := eventHandlerFunc(handler)
	informers.MustAddEventHandler(src.gwInformer.Informer(), eventHandler)
	if src.lsInformer != nil {
		informers.MustAddEventHandler(src.lsInformer.Informer(), eventHandler)
	}
	informers.MustAddEventHandler(src.rtInformer.Informer(), eventHandler)
	informers.MustAddEventHandler(src.nsInformer.Informer(), eventHandler)
}

func (src *gatewayRouteSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	routes, err := src.rtInformer.List(src.rtNamespace, src.rtLabels)
	if err != nil {
		return nil, err
	}
	gateways, err := src.gwInformer.Lister().Gateways(src.gwNamespace).List(src.gwLabels)
	if err != nil {
		return nil, err
	}
	var listenerSets []*v1.ListenerSet
	if src.lsInformer != nil {
		listenerSets, err = src.lsInformer.Lister().List(labels.Everything())
		if err != nil {
			return nil, err
		}
	}
	namespaces, err := src.nsInformer.Lister().List(labels.Everything())
	if err != nil {
		return nil, err
	}
	kind := strings.ToLower(src.rtKind)
	resolver := newGatewayRouteResolver(src, gateways, listenerSets, namespaces)
	for _, rt := range routes {
		// Filter by annotations.
		meta := rt.Metadata()
		annots := meta.Annotations
		if !src.rtAnnotations.Matches(labels.Set(annots)) {
			continue
		}

		if annotations.IsControllerMismatch(meta, src.rtKind) {
			continue
		}

		// Get Route hostnames and their targets.
		hostTargets, err := resolver.resolve(rt)
		if err != nil {
			return nil, err
		}
		// TODO: does not follow the pattern of other sources to log empty hostTargets
		if len(hostTargets) == 0 {
			log.Debugf("No endpoints could be generated from %s %s/%s", src.rtKind, meta.Namespace, meta.Name)
			continue
		}

		// Create endpoints from hostnames and targets.
		var routeEndpoints []*endpoint.Endpoint
		resource := fmt.Sprintf("%s/%s/%s", kind, meta.Namespace, meta.Name)
		providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(annots)
		ttl := annotations.TTLFromAnnotations(annots, resource)
		for host, targets := range hostTargets {
			routeEndpoints = append(routeEndpoints, endpoint.EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
		log.Debugf("Endpoints generated from %s %s/%s: %v", src.rtKind, meta.Namespace, meta.Name, routeEndpoints)

		endpoints = append(endpoints, routeEndpoints...)
	}
	return MergeEndpoints(endpoints), nil
}

func namespacedName(namespace, name string) types.NamespacedName {
	return types.NamespacedName{Namespace: namespace, Name: name}
}

type gatewayRouteResolver struct {
	src     *gatewayRouteSource
	objects map[objectRef]*listenerObject
	nss     map[string]*corev1.Namespace
}

type objectRef struct {
	schema.GroupKind
	types.NamespacedName
}

type listenerObject struct {
	gateway        *v1.Gateway
	gatewayRef     types.NamespacedName
	listenerSet    *v1.ListenerSet
	ownerNamespace string
	sections       map[v1.SectionName][]listenerSection
	attached       bool
	attachReason   string
	overrides      endpoint.Targets
}

type listenerSection struct {
	listener     v1.Listener
	attached     bool
	attachReason string
}

type resolvedParent struct {
	kind      string
	namespace string
	ref       v1.ParentReference
	obj       *listenerObject
	section   v1.SectionName
	listeners []listenerSection
}

func newGatewayRouteResolver(src *gatewayRouteSource, gateways []*v1.Gateway, listenerSets []*v1.ListenerSet, namespaces []*corev1.Namespace) *gatewayRouteResolver {
	// Create Namespace lookup table.
	nss := make(map[string]*corev1.Namespace, len(namespaces))
	for _, ns := range namespaces {
		nss[ns.Name] = ns
	}

	objects := make(map[objectRef]*listenerObject, len(gateways)+len(listenerSets))
	for _, gw := range gateways {
		objects[newObjectRef(gatewayGroup, gatewayKind, gw.Namespace, gw.Name)] = newGatewayListenerObject(gw)
	}
	for _, ls := range listenerSets {
		objects[newObjectRef(gatewayGroup, listenerSetKind, ls.Namespace, ls.Name)] = newListenerSetObject(ls, objects, nss)
	}
	return &gatewayRouteResolver{
		src:     src,
		objects: objects,
		nss:     nss,
	}
}

func (c *gatewayRouteResolver) resolve(rt gatewayRoute) (map[string]endpoint.Targets, error) {
	rtHosts, err := c.hosts(rt)
	if err != nil {
		return nil, err
	}
	hostTargets := make(map[string]endpoint.Targets)

	routeParentRefs := rt.ParentRefs()

	if len(routeParentRefs) == 0 {
		log.Debugf("No parent references found for %s %s/%s", c.src.rtKind, rt.Metadata().Namespace, rt.Metadata().Name)
		return hostTargets, nil
	}

	for _, rps := range rt.RouteStatus().Parents {
		parent, ok := c.resolveParentRef(rt, routeParentRefs, rps)
		if !ok {
			continue
		}
		c.matchRouteToParent(rt, rtHosts, parent, hostTargets)
	}
	// If a Gateway has multiple matching Listeners for the same host, then we'll
	// add its IPs to the target list multiple times and should dedupe them.
	for host, targets := range hostTargets {
		hostTargets[host] = uniqueTargets(targets)
	}
	return hostTargets, nil
}

func (c *gatewayRouteResolver) resolveParentRef(rt gatewayRoute, routeParentRefs []v1.ParentReference, rps v1.RouteParentStatus) (*resolvedParent, bool) {
	meta := rt.Metadata()
	ref := rps.ParentRef
	namespace := strVal((*string)(ref.Namespace), meta.Namespace)
	if !gwRouteHasParentRef(routeParentRefs, ref, meta) {
		log.Debugf("Parent reference %s/%s not found in routeParentRefs for %s %s/%s", namespace, string(ref.Name), c.src.rtKind, meta.Namespace, meta.Name)
		return nil, false
	}

	group := strVal((*string)(ref.Group), gatewayGroup)
	kind := strVal((*string)(ref.Kind), gatewayKind)
	if group != gatewayGroup || (kind != gatewayKind && kind != listenerSetKind) {
		log.Debugf("Unsupported parent %s/%s for %s %s/%s", group, kind, c.src.rtKind, meta.Namespace, meta.Name)
		return nil, false
	}

	obj, ok := c.objects[newObjectRef(group, kind, namespace, string(ref.Name))]
	if !ok {
		log.Debugf("%s %s/%s not found for %s %s/%s", kind, namespace, ref.Name, c.src.rtKind, meta.Namespace, meta.Name)
		return nil, false
	}
	if c.src.gwName != "" && c.src.gwName != obj.gatewayRef.Name {
		log.Debugf("Gateway %s/%s does not match %s %s/%s", obj.gatewayRef.Namespace, obj.gatewayRef.Name, c.src.gwName, meta.Namespace, meta.Name)
		return nil, false
	}
	if !gwRouteIsAccepted(rps.Conditions) {
		log.Debugf("%s %s/%s has not accepted the current generation %s %s/%s", kind, namespace, ref.Name, c.src.rtKind, meta.Namespace, meta.Name)
		return nil, false
	}
	if !obj.attached {
		log.Debugf("%s %s/%s is not attached for %s %s/%s: %s", kind, namespace, ref.Name, c.src.rtKind, meta.Namespace, meta.Name, obj.attachReason)
		return nil, false
	}

	sectionName := sectionVal(ref.SectionName, "")
	listeners := obj.sections[sectionName]
	if len(listeners) == 0 {
		log.Debugf("%s %s/%s section %q not found for %s %s/%s", kind, namespace, ref.Name, sectionName, c.src.rtKind, meta.Namespace, meta.Name)
		return nil, false
	}

	return &resolvedParent{
		kind:      kind,
		namespace: namespace,
		ref:       ref,
		obj:       obj,
		section:   sectionName,
		listeners: listeners,
	}, true
}

func (c *gatewayRouteResolver) matchRouteToParent(rt gatewayRoute, rtHosts []string, parent *resolvedParent, hostTargets map[string]endpoint.Targets) {
	meta := rt.Metadata()
	match := false
	var unattachedReason string
	for i := range parent.listeners {
		section := &parent.listeners[i]
		if !section.attached {
			if parent.section != "" && unattachedReason == "" {
				unattachedReason = section.attachReason
			}
			continue
		}
		if c.matchRouteToListener(rt, rtHosts, parent, &section.listener, hostTargets) {
			match = true
		}
	}
	if match {
		return
	}
	if parent.section != "" && unattachedReason != "" {
		log.Debugf("%s %s/%s section %q is not attached for %s %s/%s: %s", parent.kind, parent.namespace, parent.ref.Name, parent.section, c.src.rtKind, meta.Namespace, meta.Name, unattachedReason)
		return
	}
	log.Debugf("%s %s/%s section %q does not match %s %s/%s hostnames %q", parent.kind, parent.namespace, parent.ref.Name, parent.section, c.src.rtKind, meta.Namespace, meta.Name, rtHosts)
}

func (c *gatewayRouteResolver) matchRouteToListener(rt gatewayRoute, rtHosts []string, parent *resolvedParent, lis *v1.Listener, hostTargets map[string]endpoint.Targets) bool {
	if !gwProtocolMatches(rt.Protocol(), lis.Protocol) {
		return false
	}
	// EXPERIMENTAL: https://gateway-api.sigs.k8s.io/geps/gep-957/
	if parent.ref.Port != nil && *parent.ref.Port != lis.Port {
		return false
	}
	if !c.routeIsAllowed(parent.obj.ownerNamespace, lis, rt) {
		return false
	}

	match := false
	gwHost := ""
	if lis.Hostname != nil {
		gwHost = string(*lis.Hostname)
	}
	for _, rtHost := range rtHosts {
		if gwHost == "" && rtHost == "" {
			continue
		}
		host, ok := gwMatchingHost(gwHost, rtHost)
		if !ok {
			continue
		}
		override := parent.obj.overrides
		hostTargets[host] = append(hostTargets[host], override...)
		if len(override) == 0 {
			for _, addr := range parent.obj.gateway.Status.Addresses {
				hostTargets[host] = append(hostTargets[host], addr.Value)
			}
		}
		match = true
	}
	return match
}

func (c *gatewayRouteResolver) hosts(rt gatewayRoute) ([]string, error) {
	var hostnames []string
	for _, name := range rt.Hostnames() {
		hostnames = append(hostnames, string(name))
	}
	// TODO: The combine-fqdn-annotation flag is similarly vague.
	if c.src.templateEngine.IsConfigured() && (len(hostnames) == 0 || c.src.templateEngine.Combining()) {
		hosts, err := c.src.templateEngine.ExecFQDN(rt.Object())
		if err != nil {
			return nil, err
		}
		hostnames = append(hostnames, hosts...)
	}

	hostNameAnnotation, hostNameAnnotationExists := rt.Metadata().Annotations[annotations.GatewayHostnameSourceKey]
	if !hostNameAnnotationExists {
		// This means that the route doesn't specify a hostname and should use any provided by
		// attached Gateway Listeners. This is only useful for {HTTP,TLS}Routes, but it doesn't
		// break {TCP,UDP}Routes.
		if len(rt.Hostnames()) == 0 {
			hostnames = append(hostnames, "")
		}
		if !c.src.ignoreHostnameAnnotation {
			hostnames = append(hostnames, annotations.HostnamesFromAnnotations(rt.Metadata().Annotations)...)
		}
		return hostnames, nil
	}

	switch strings.ToLower(hostNameAnnotation) {
	case gatewayHostnameSourceAnnotationOnlyValue:
		if c.src.ignoreHostnameAnnotation {
			return []string{}, nil
		}
		return annotations.HostnamesFromAnnotations(rt.Metadata().Annotations), nil
	case gatewayHostnameSourceDefinedHostsOnlyValue:
		// Explicitly use only defined hostnames (route spec and optional template result)
		return hostnames, nil
	default:
		// Invalid value provided: warn and fall back to default behavior (as if the annotation is absent)
		log.Warnf("Invalid value for %q on %s/%s: %q. Falling back to default behavior.",
			annotations.GatewayHostnameSourceKey, rt.Metadata().Namespace, rt.Metadata().Name, hostNameAnnotation)
		if len(rt.Hostnames()) == 0 {
			hostnames = append(hostnames, "")
		}
		if !c.src.ignoreHostnameAnnotation {
			hostnames = append(hostnames, annotations.HostnamesFromAnnotations(rt.Metadata().Annotations)...)
		}
		return hostnames, nil
	}
}

func (c *gatewayRouteResolver) routeIsAllowed(ownerNamespace string, lis *v1.Listener, rt gatewayRoute) bool {
	meta := rt.Metadata()
	allow := lis.AllowedRoutes

	// Check the route's namespace.
	from := v1.NamespacesFromSame
	if allow != nil && allow.Namespaces != nil && allow.Namespaces.From != nil {
		from = *allow.Namespaces.From
	}
	switch from {
	case v1.NamespacesFromAll:
		// OK
	case v1.NamespacesFromSame:
		if ownerNamespace != meta.Namespace {
			return false
		}
	case v1.NamespacesFromSelector:
		selector, err := metav1.LabelSelectorAsSelector(allow.Namespaces.Selector)
		if err != nil {
			log.Debugf("Listener %q has invalid namespace selector: %v", lis.Name, err)
			return false
		}
		// Get namespace.
		ns, ok := c.nss[meta.Namespace]
		if !ok {
			log.Errorf("Namespace not found for %s %s/%s", c.src.rtKind, meta.Namespace, meta.Name)
			return false
		}
		if !selector.Matches(labels.Set(ns.Labels)) {
			return false
		}
	default:
		log.Debugf("Listener %q has unknown namespace from %q", lis.Name, from)
		return false
	}

	// Check the route's kind, if any are specified by the listener.
	// TODO: Do we need to consider SupportedKinds in the ListenerStatus instead of the Spec?
	// We only support core kinds and already check the protocol... Does this matter at all?
	if allow == nil || len(allow.Kinds) == 0 {
		return true
	}
	gvk := rt.Object().GetObjectKind().GroupVersionKind()
	for _, gk := range allow.Kinds {
		group := strVal((*string)(gk.Group), gatewayGroup)
		if gvk.Group == group && gvk.Kind == string(gk.Kind) {
			return true
		}
	}
	return false
}

func newObjectRef(group, kind, namespace, name string) objectRef {
	return objectRef{
		GroupKind:      schema.GroupKind{Group: group, Kind: kind},
		NamespacedName: namespacedName(namespace, name),
	}
}

func newGatewayListenerObject(gw *v1.Gateway) *listenerObject {
	return &listenerObject{
		gateway:        gw,
		gatewayRef:     namespacedName(gw.Namespace, gw.Name),
		ownerNamespace: gw.Namespace,
		sections:       gatewaySections(gw.Spec.Listeners),
		attached:       true,
		overrides:      annotations.TargetsFromTargetAnnotation(gw.Annotations),
	}
}

func newListenerSetObject(ls *v1.ListenerSet, objects map[objectRef]*listenerObject, namespaces map[string]*corev1.Namespace) *listenerObject {
	gwNamespace := strVal((*string)(ls.Spec.ParentRef.Namespace), ls.Namespace)
	gwRef := namespacedName(gwNamespace, string(ls.Spec.ParentRef.Name))
	obj := &listenerObject{
		gatewayRef:     gwRef,
		listenerSet:    ls,
		ownerNamespace: ls.Namespace,
	}

	gwObj, ok := objects[newObjectRef(gatewayGroup, gatewayKind, gwRef.Namespace, gwRef.Name)]
	switch {
	case !ok:
		obj.attachReason = fmt.Sprintf("parent Gateway %s/%s not found", gwRef.Namespace, gwRef.Name)
	case !gatewayAllowsListenerSet(gwObj.gateway, ls, namespaces):
		obj.gateway = gwObj.gateway
		obj.attachReason = fmt.Sprintf("parent Gateway %s/%s does not allow ListenerSets from namespace %s", gwRef.Namespace, gwRef.Name, ls.Namespace)
	case !listenerSetIsAccepted(ls.Status.Conditions):
		obj.gateway = gwObj.gateway
		obj.attachReason = fmt.Sprintf("ListenerSet %s/%s has not been accepted", ls.Namespace, ls.Name)
	default:
		obj.gateway = gwObj.gateway
		obj.attached = true
	}

	if obj.gateway != nil {
		obj.overrides = annotations.TargetsFromTargetAnnotation(ls.Annotations)
		if len(obj.overrides) == 0 {
			obj.overrides = annotations.TargetsFromTargetAnnotation(obj.gateway.Annotations)
		}
	}

	obj.sections = listenerSetSections(ls, obj.attached, obj.attachReason)
	return obj
}

func gatewaySections(listeners []v1.Listener) map[v1.SectionName][]listenerSection {
	sections := make([]listenerSection, 0, len(listeners))
	bySection := make(map[v1.SectionName][]listenerSection, len(listeners)+1)
	for _, lis := range listeners {
		section := listenerSection{
			listener: lis,
			attached: true,
		}
		sections = append(sections, section)
		bySection[lis.Name] = []listenerSection{section}
	}
	bySection[""] = sections
	return bySection
}

func listenerSetSections(ls *v1.ListenerSet, objectAttached bool, objectReason string) map[v1.SectionName][]listenerSection {
	sections := make([]listenerSection, 0, len(ls.Spec.Listeners))
	bySection := make(map[v1.SectionName][]listenerSection, len(ls.Spec.Listeners)+1)
	for _, entry := range ls.Spec.Listeners {
		section := listenerSection{
			listener: v1.Listener(entry),
		}
		accepted, hasStatus := listenerSetListenerAcceptedCondition(ls.Status.Listeners, entry.Name)
		switch {
		case !objectAttached:
			section.attachReason = objectReason
		case !hasStatus:
			section.attachReason = fmt.Sprintf("ListenerSet %s/%s section %q has no Accepted status", ls.Namespace, ls.Name, entry.Name)
		case accepted:
			section.attached = true
		default:
			section.attachReason = fmt.Sprintf("ListenerSet %s/%s section %q has not been accepted", ls.Namespace, ls.Name, entry.Name)
		}
		sections = append(sections, section)
		bySection[entry.Name] = []listenerSection{section}
	}
	bySection[""] = sections
	return bySection
}

func gatewayAllowsListenerSet(gw *v1.Gateway, ls *v1.ListenerSet, namespaces map[string]*corev1.Namespace) bool {
	from := v1.NamespacesFromNone
	if gw.Spec.AllowedListeners != nil && gw.Spec.AllowedListeners.Namespaces != nil && gw.Spec.AllowedListeners.Namespaces.From != nil {
		from = *gw.Spec.AllowedListeners.Namespaces.From
	}
	switch from {
	case v1.NamespacesFromAll:
		return true
	case v1.NamespacesFromSame:
		return gw.Namespace == ls.Namespace
	case v1.NamespacesFromSelector:
		if gw.Spec.AllowedListeners == nil || gw.Spec.AllowedListeners.Namespaces == nil || gw.Spec.AllowedListeners.Namespaces.Selector == nil {
			log.Debugf("Gateway %s/%s has invalid AllowedListeners selector", gw.Namespace, gw.Name)
			return false
		}
		selector, err := metav1.LabelSelectorAsSelector(gw.Spec.AllowedListeners.Namespaces.Selector)
		if err != nil {
			log.Debugf("Gateway %s/%s has invalid AllowedListeners selector: %v", gw.Namespace, gw.Name, err)
			return false
		}
		ns, ok := namespaces[ls.Namespace]
		if !ok {
			log.Errorf("Namespace %q not found for ListenerSet %s/%s", ls.Namespace, ls.Namespace, ls.Name)
			return false
		}
		return selector.Matches(labels.Set(ns.Labels))
	case v1.NamespacesFromNone:
		return false
	default:
		log.Debugf("Gateway %s/%s has unknown AllowedListeners namespace from %q", gw.Namespace, gw.Name, from)
		return false
	}
}

func gwRouteHasParentRef(routeParentRefs []v1.ParentReference, ref v1.ParentReference, meta *metav1.ObjectMeta) bool {
	// Ensure that the parent reference is in the routeParentRefs list
	namespace := strVal((*string)(ref.Namespace), meta.Namespace)
	group := strVal((*string)(ref.Group), gatewayGroup)
	kind := strVal((*string)(ref.Kind), gatewayKind)
	for _, rpr := range routeParentRefs {
		rprGroup := strVal((*string)(rpr.Group), gatewayGroup)
		rprKind := strVal((*string)(rpr.Kind), gatewayKind)
		if rprGroup != group || rprKind != kind {
			continue
		}
		rprNamespace := strVal((*string)(rpr.Namespace), meta.Namespace)
		if string(rpr.Name) != string(ref.Name) || rprNamespace != namespace {
			continue
		}
		return true
	}
	return false
}

func gwRouteIsAccepted(conds []metav1.Condition) bool {
	return conditionStatusIsTrue(conds, string(v1.RouteConditionAccepted))
}

func listenerSetIsAccepted(conds []metav1.Condition) bool {
	return conditionStatusIsTrue(conds, string(v1.ListenerSetConditionAccepted))
}

func listenerSetListenerAcceptedCondition(statuses []v1.ListenerEntryStatus, name v1.SectionName) (bool, bool) {
	for _, status := range statuses {
		if status.Name != name {
			continue
		}
		return conditionStatusIsTrue(status.Conditions, string(v1.ListenerEntryConditionAccepted)), true
	}
	return false, false
}

func conditionStatusIsTrue(conds []metav1.Condition, conditionType string) bool {
	for _, c := range conds {
		if c.Type == conditionType {
			return c.Status == metav1.ConditionTrue
		}
	}
	return false
}

func uniqueTargets(targets endpoint.Targets) endpoint.Targets {
	if len(targets) < 2 {
		return targets
	}
	sort.Strings([]string(targets))
	prev := targets[0]
	n := 1
	for _, v := range targets[1:] {
		if v == prev {
			continue
		}
		prev = v
		targets[n] = v
		n++
	}
	return targets[:n]
}

// gwProtocolMatches returns whether a and b are the same protocol,
// where HTTP and HTTPS are considered the same.
// and TLS and TCP are considered the same.
func gwProtocolMatches(a, b v1.ProtocolType) bool {
	if a == v1.HTTPSProtocolType {
		a = v1.HTTPProtocolType
	}
	if b == v1.HTTPSProtocolType {
		b = v1.HTTPProtocolType
	}
	// if Listener is TLS and Route is TCP set Listener type to TCP as to pass true and return valid match
	if a == v1.TCPProtocolType && b == v1.TLSProtocolType {
		b = v1.TCPProtocolType
	}
	return a == b
}

// gwMatchingHost returns the most-specific overlapping host and a bool indicating if one was found.
// Hostnames that are prefixed with a wildcard label (`*.`) are interpreted as a suffix match.
// That means that "*.example.com" would match both "test.example.com" and "foo.test.example.com",
// but not "example.com". An empty string matches anything.
func gwMatchingHost(a, b string) (string, bool) {
	var ok bool
	if a, ok = gwHost(a); !ok {
		return "", false
	}
	if b, ok = gwHost(b); !ok {
		return "", false
	}

	if a == "" {
		return b, true
	}
	if b == "" || a == b {
		return a, true
	}
	if na, nb := len(a), len(b); nb < na || (na == nb && strings.HasPrefix(b, "*.")) {
		a, b = b, a
	}
	if strings.HasPrefix(a, "*.") && strings.HasSuffix(b, a[1:]) {
		return b, true
	}
	return "", false
}

// gwHost returns the canonical host and a value indicating if it's valid.
func gwHost(host string) (string, bool) {
	if host == "" {
		return "", true
	}
	if isIPAddr(host) || !isDNS1123Domain(strings.TrimPrefix(host, "*.")) {
		return "", false
	}
	return toLowerCaseASCII(host), true
}

// isIPAddr returns whether s in an IP address.
func isIPAddr(s string) bool {
	return endpoint.SuitableType(s) != endpoint.RecordTypeCNAME
}

// isDNS1123Domain returns whether s is a valid domain name according to RFC 1123.
func isDNS1123Domain(s string) bool {
	if n := len(s); n == 0 || n > 255 {
		return false
	}
	for lbl, rest := "", s; rest != ""; {
		if lbl, rest, _ = strings.Cut(rest, "."); !isDNS1123Label(lbl) {
			return false
		}
	}
	return true
}

// isDNS1123Label returns whether s is a valid domain label according to RFC 1123.
func isDNS1123Label(s string) bool {
	n := len(s)
	if n == 0 || n > 63 {
		return false
	}
	if !isAlphaNum(s[0]) || !isAlphaNum(s[n-1]) {
		return false
	}
	for i, k := 1, n-1; i < k; i++ {
		if b := s[i]; b != '-' && !isAlphaNum(b) {
			return false
		}
	}
	return true
}

func isAlphaNum(b byte) bool {
	switch {
	case 'a' <= b && b <= 'z',
		'A' <= b && b <= 'Z',
		'0' <= b && b <= '9':
		return true
	default:
		return false
	}
}

func strVal(ptr *string, def string) string {
	if ptr == nil || *ptr == "" {
		return def
	}
	return *ptr
}

func sectionVal(ptr *v1.SectionName, def v1.SectionName) v1.SectionName {
	if ptr == nil || *ptr == "" {
		return def
	}
	return *ptr
}

func selectorsEqual(a, b labels.Selector) bool {
	if a == nil || b == nil {
		return a == b
	}
	aReq, aOK := a.DeepCopySelector().Requirements()
	bReq, bOK := b.DeepCopySelector().Requirements()
	if aOK != bOK || len(aReq) != len(bReq) {
		return false
	}
	sort.Stable(labels.ByKey(aReq))
	sort.Stable(labels.ByKey(bReq))
	for i, r := range aReq {
		if !r.Equal(bReq[i]) {
			return false
		}
	}
	return true
}
