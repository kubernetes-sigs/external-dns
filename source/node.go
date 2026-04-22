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
	"fmt"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
)

// nodeSource is an implementation of Source for Kubernetes Node objects.
//
// +externaldns:source:name=node
// +externaldns:source:category=Kubernetes Core
// +externaldns:source:description=Creates DNS entries based on Kubernetes Node resources
// +externaldns:source:resources=Node
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=false
// +externaldns:source:events=true
type nodeSource struct {
	client           kubernetes.Interface
	annotationFilter string
	templateEngine   template.Engine

	nodeInformer         coreinformers.NodeInformer
	labelSelector        labels.Selector
	excludeUnschedulable bool
	exposeInternalIPv6   bool
}

// NewNodeSource creates a new nodeSource with the given config.
func NewNodeSource(
	ctx context.Context,
	kubeClient kubernetes.Interface,
	cfg *Config) (Source, error) {
	// Use shared informers to listen for add/update/delete of nodes.
	// Set resync period to 0, to prevent processing when nothing has changed
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0)
	nodeInformer := informerFactory.Core().V1().Nodes()

	informers.MustSetTransform(nodeInformer.Informer(), informers.TransformerWithOptions[*v1.Node](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
		informers.TransformRemoveStatusConditions(),
	))

	// Add default resource event handler to properly initialize informer.
	informers.MustAddEventHandler(nodeInformer.Informer(), informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &nodeSource{
		client:               kubeClient,
		annotationFilter:     cfg.AnnotationFilter,
		templateEngine:       cfg.TemplateEngine,
		nodeInformer:         nodeInformer,
		labelSelector:        cfg.LabelFilter,
		excludeUnschedulable: cfg.ExcludeUnschedulable,
		exposeInternalIPv6:   cfg.ExposeInternalIPv6,
	}, nil
}

// Endpoints returns endpoint objects for each service that should be processed.
func (ns *nodeSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	nodes, err := ns.nodeInformer.Lister().List(ns.labelSelector)
	if err != nil {
		return nil, err
	}

	nodes, err = annotations.Filter(nodes, ns.annotationFilter)
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)

	// create endpoints for all nodes
	for _, node := range nodes {
		if annotations.IsControllerMismatch(node, types.Node) {
			continue
		}

		if node.Spec.Unschedulable && ns.excludeUnschedulable {
			log.Debugf("Skipping node %s because it is unschedulable", node.Name)
			continue
		}

		log.Debugf("creating endpoint for node %s", node.Name)

		// Only generate node name endpoints when there's no template or when combining
		var nodeEndpoints []*endpoint.Endpoint
		if !ns.templateEngine.IsConfigured() || ns.templateEngine.Combining() {
			nodeEndpoints, err = ns.endpointsForDNSNames(node, []string{node.Name})
			if err != nil {
				return nil, err
			}
		}

		nodeEndpoints, err = ns.templateEngine.CombineWithEndpoints(
			nodeEndpoints,
			func() ([]*endpoint.Endpoint, error) { return ns.endpointsFromNodeTemplate(node) },
		)
		if err != nil {
			return nil, err
		}

		if len(nodeEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from node %s", node.Name)
			continue
		}

		endpoint.AttachRefObject(nodeEndpoints, events.NewObjectReference(node, types.Node))

		endpoints = append(endpoints, nodeEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
}

func (ns *nodeSource) AddEventHandler(_ context.Context, handler func()) {
	informers.MustAddEventHandler(ns.nodeInformer.Informer(), eventHandlerFunc(handler))
}

// endpointsFromNodeTemplate creates endpoints using DNS names from the FQDN template.
func (ns *nodeSource) endpointsFromNodeTemplate(node *v1.Node) ([]*endpoint.Endpoint, error) {
	names, err := ns.templateEngine.ExecFQDN(node)
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		log.Debugf("applied template for %s, converting to %s", node.Name, name)
	}

	return ns.endpointsForDNSNames(node, names)
}

// endpointsForDNSNames creates endpoints for the given DNS names using the node's addresses.
func (ns *nodeSource) endpointsForDNSNames(node *v1.Node, dnsNames []string) ([]*endpoint.Endpoint, error) {
	ttl := annotations.TTLFromAnnotations(node.Annotations, fmt.Sprintf("node/%s", node.Name))

	addrs := annotations.TargetsFromTargetAnnotation(node.Annotations)
	if len(addrs) == 0 {
		var err error
		addrs, err = ns.nodeAddresses(node)
		if err != nil {
			return nil, fmt.Errorf("failed to get node address from %s: %w", node.Name, err)
		}
	}

	var endpoints []*endpoint.Endpoint
	for _, dns := range dnsNames {
		log.Debugf("adding endpoint with %d targets", len(addrs))

		for _, addr := range addrs {
			ep := endpoint.NewEndpointWithTTL(dns, endpoint.SuitableType(addr), ttl, addr)
			ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("node/%s", node.Name))
			log.Debugf("adding endpoint %s target %s", ep, addr)
			endpoints = append(endpoints, ep)
		}
	}

	return MergeEndpoints(endpoints), nil
}

// nodeAddress returns the node's externalIP and if that's not found, the node's internalIP
// basically what k8s.io/kubernetes/pkg/util/node.GetPreferredNodeAddress does
func (ns *nodeSource) nodeAddresses(node *v1.Node) ([]string, error) {
	addresses := map[v1.NodeAddressType][]string{
		v1.NodeExternalIP: {},
		v1.NodeInternalIP: {},
	}
	var internalIpv6Addresses []string

	for _, addr := range node.Status.Addresses {
		// IPv6 InternalIP addresses have special handling.
		// Refer to https://github.com/kubernetes-sigs/external-dns/pull/5192 for more details.
		if addr.Type == v1.NodeInternalIP && endpoint.SuitableType(addr.Address) == endpoint.RecordTypeAAAA {
			internalIpv6Addresses = append(internalIpv6Addresses, addr.Address)
		}
		addresses[addr.Type] = append(addresses[addr.Type], addr.Address)
	}

	if len(addresses[v1.NodeExternalIP]) > 0 {
		if ns.exposeInternalIPv6 {
			return append(addresses[v1.NodeExternalIP], internalIpv6Addresses...), nil
		}
		return addresses[v1.NodeExternalIP], nil
	}

	if len(addresses[v1.NodeInternalIP]) > 0 {
		return addresses[v1.NodeInternalIP], nil
	}

	return nil, fmt.Errorf("could not find node address for %s", node.Name)
}
