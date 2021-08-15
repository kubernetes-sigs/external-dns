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
	"text/template"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	cache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/gateway/versioned"
	informers "sigs.k8s.io/gateway-api/pkg/client/informers/gateway/externalversions"
	informers_v1a2 "sigs.k8s.io/gateway-api/pkg/client/informers/gateway/externalversions/apis/v1alpha2"

	"sigs.k8s.io/external-dns/endpoint"
)

type gatewayRoute interface {
	// Object returns the underlying Route object to be used by templates.
	Object() kubeObject
	// Metadata returns the Route's metadata.
	Metadata() *metav1.ObjectMeta
	// Hostnames returns the Route's specified hostnames.
	Hostnames() []v1alpha2.Hostname
	// Status returns the Route's status, including associated gateways.
	Status() v1alpha2.RouteStatus
}

type newGatewayRouteInformerFunc func(informers.SharedInformerFactory) gatewayRouteInfomer

type gatewayRouteInfomer interface {
	List(namespace string, selector labels.Selector) ([]gatewayRoute, error)
	Informer() cache.SharedIndexInformer
}

func newGatewayInformerFactory(client gateway.Interface, namespace string, labelSelector labels.Selector) informers.SharedInformerFactory {
	var opts []informers.SharedInformerOption
	if namespace != "" {
		opts = append(opts, informers.WithNamespace(namespace))
	}
	if labelSelector != nil && !labelSelector.Empty() {
		lbls := labelSelector.String()
		opts = append(opts, informers.WithTweakListOptions(func(o *metav1.ListOptions) {
			o.LabelSelector = lbls
		}))
	}
	return informers.NewSharedInformerFactoryWithOptions(client, 0, opts...)
}

type gatewayRouteSource struct {
	gwNamespace string
	gwLabels    labels.Selector
	gwInformer  informers_v1a2.GatewayInformer

	rtKind        string
	rtNamespace   string
	rtLabels      labels.Selector
	rtAnnotations labels.Selector
	rtInformer    gatewayRouteInfomer

	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
}

func newGatewayRouteSource(clients ClientGenerator, config *Config, kind string, newInformerFn newGatewayRouteInformerFunc) (Source, error) {
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
	tmpl, err := parseTemplate(config.FQDNTemplate)
	if err != nil {
		return nil, err
	}

	client, err := clients.GatewayClient()
	if err != nil {
		return nil, err
	}

	informerFactory := newGatewayInformerFactory(client, config.GatewayNamespace, gwLabels)
	gwInformer := informerFactory.Gateway().V1alpha2().Gateways() // TODO: gateway informer should be shared across gateway sources
	gwInformer.Informer()                                         // Register with factory before starting

	rtInformerFactory := informerFactory
	if config.Namespace != config.GatewayNamespace || !selectorsEqual(rtLabels, gwLabels) {
		rtInformerFactory = newGatewayInformerFactory(client, config.Namespace, rtLabels)
	}
	rtInformer := newInformerFn(rtInformerFactory)
	rtInformer.Informer() // Register with factory before starting

	informerFactory.Start(wait.NeverStop)
	if rtInformerFactory != informerFactory {
		rtInformerFactory.Start(wait.NeverStop)

		if err := waitForCacheSync(context.Background(), rtInformerFactory); err != nil {
			return nil, err
		}
	}
	if err := waitForCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	src := &gatewayRouteSource{
		gwNamespace: config.GatewayNamespace,
		gwLabels:    gwLabels,
		gwInformer:  gwInformer,

		rtKind:        kind,
		rtNamespace:   config.Namespace,
		rtLabels:      rtLabels,
		rtAnnotations: rtAnnotations,
		rtInformer:    rtInformer,

		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    config.CombineFQDNAndAnnotation,
		ignoreHostnameAnnotation: config.IgnoreHostnameAnnotation,
	}
	return src, nil
}

func (src *gatewayRouteSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debugf("Adding event handler for %s", src.rtKind)
	src.gwInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	src.rtInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

func (src *gatewayRouteSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	routes, err := src.rtInformer.List(src.rtNamespace, src.rtLabels)
	if err != nil {
		return nil, err
	}
	gwList, err := src.gwInformer.Lister().Gateways(src.gwNamespace).List(src.gwLabels)
	if err != nil {
		return nil, err
	}
	gateways := gatewaysByRef(gwList)
	for _, rt := range routes {
		eps, err := src.endpoints(rt, gateways)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, eps...)
	}
	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}
	return endpoints, nil
}

func (src *gatewayRouteSource) endpoints(rt gatewayRoute, gateways map[types.NamespacedName]*v1alpha2.Gateway) ([]*endpoint.Endpoint, error) {
	// Filter by annotations.
	meta := rt.Metadata()
	annotations := meta.Annotations
	if !src.rtAnnotations.Matches(labels.Set(meta.Annotations)) {
		return nil, nil
	}

	// Check controller annotation to see if we are responsible.
	if v, ok := meta.Annotations[controllerAnnotationKey]; ok && v != controllerAnnotationValue {
		log.Debugf("Skipping %s %s/%s because controller value does not match, found: %s, required: %s",
			src.rtKind, meta.Namespace, meta.Name, v, controllerAnnotationValue)
		return nil, nil
	}

	// Get hostnames.
	hostnames, err := src.hostnames(rt)
	if err != nil {
		return nil, err
	}
	if len(hostnames) == 0 {
		log.Debugf("No hostnames could be generated from %s %s/%s", src.rtKind, meta.Namespace, meta.Name)
		return nil, nil
	}

	// Get targets.
	targets := src.targets(rt, gateways)
	if len(targets) == 0 {
		log.Debugf("No targets could be generated from %s %s/%s", src.rtKind, meta.Namespace, meta.Name)
		return nil, nil
	}

	// Create endpoints.
	ttl, err := getTTLFromAnnotations(annotations)
	if err != nil {
		log.Warn(err)
	}
	providerSpecific, setIdentifier := getProviderSpecificAnnotations(annotations)
	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
	}
	log.Debugf("Endpoints generated from %s %s/%s: %v", src.rtKind, meta.Namespace, meta.Name, endpoints)

	kind := strings.ToLower(src.rtKind)
	resourceKey := fmt.Sprintf("%s/%s/%s", kind, meta.Namespace, meta.Name)
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = resourceKey
	}
	return endpoints, nil
}

func (src *gatewayRouteSource) hostnames(rt gatewayRoute) ([]string, error) {
	var hostnames []string
	for _, name := range rt.Hostnames() {
		hostnames = append(hostnames, string(name))
	}
	meta := rt.Metadata()
	// TODO: The ignore-hostname-annotation flag help says "valid only when using fqdn-template"
	// but other sources don't check if fqdn-template is set. Which should it be?
	if !src.ignoreHostnameAnnotation {
		hostnames = append(hostnames, getHostnamesFromAnnotations(meta.Annotations)...)
	}
	// TODO: The combine-fqdn-annotation flag is similarly vague.
	if src.fqdnTemplate != nil && (len(hostnames) == 0 || src.combineFQDNAnnotation) {
		hosts, err := execTemplate(src.fqdnTemplate, rt.Object())
		if err != nil {
			return nil, err
		}
		hostnames = append(hostnames, hosts...)
	}
	return hostnames, nil
}

func (src *gatewayRouteSource) targets(rt gatewayRoute, gateways map[types.NamespacedName]*v1alpha2.Gateway) endpoint.Targets {
	var targets endpoint.Targets
	meta := rt.Metadata()
	for _, rps := range rt.Status().Parents {
		ref := rps.ParentRef
		if (ref.Group != nil && *ref.Group != "gateway.networking.k8s.io") || (ref.Kind != nil && *ref.Kind != "Gateway") {
			log.Debugf("Unsupported parent %v/%v for %s %s/%s", ref.Group, ref.Kind, src.rtKind, meta.Namespace, meta.Name)
			continue
		}
		namespace := meta.Namespace
		if ref.Namespace != nil {
			namespace = string(*ref.Namespace)
		}
		gw, ok := gateways[types.NamespacedName{
			Namespace: namespace,
			Name:      string(ref.Name),
		}]
		if !ok {
			log.Debugf("Gateway %s/%s not found for %s %s/%s", namespace, ref.Name, src.rtKind, meta.Namespace, meta.Name)
			continue
		}
		if !gwRouteIsAdmitted(rps.Conditions) {
			log.Debugf("Gateway %s/%s has not admitted %s %s/%s", namespace, ref.Name, src.rtKind, meta.Namespace, meta.Name)
			continue
		}
		for _, addr := range gw.Status.Addresses {
			// TODO: Should we validate address type?
			// The spec says it should always be an IP.
			targets = append(targets, addr.Value)
		}
	}
	return targets
}

func gwRouteIsAdmitted(conds []metav1.Condition) bool {
	for _, c := range conds {
		if v1alpha2.RouteConditionType(c.Type) == v1alpha2.ConditionRouteAccepted {
			return c.Status == metav1.ConditionTrue
		}
	}
	return false
}

func gatewaysByRef(list []*v1alpha2.Gateway) map[types.NamespacedName]*v1alpha2.Gateway {
	if len(list) == 0 {
		return nil
	}
	set := make(map[types.NamespacedName]*v1alpha2.Gateway, len(list))
	for _, gw := range list {
		set[types.NamespacedName{Namespace: gw.Namespace, Name: gw.Name}] = gw
	}
	return set
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
