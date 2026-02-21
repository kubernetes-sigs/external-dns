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
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/fqdn"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/types"
)

// unstructuredSource is a Source that creates DNS records from unstructured resources.
//
// +externaldns:source:name=unstructured
// +externaldns:source:category=Custom Resources
// +externaldns:source:description=Creates DNS entries from unstructured Kubernetes resources
// +externaldns:source:resources=Unstructured
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
type unstructuredSource struct {
	combineFqdnAnnotation bool
	fqdnTemplate          *template.Template
	targetTemplate        *template.Template
	hostTargetTemplate    *template.Template
	informers             []kubeinformers.GenericInformer
}

// NewUnstructuredFQDNSource creates a new unstructuredSource.
func NewUnstructuredFQDNSource(
	ctx context.Context,
	dynamicClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	namespace, annotationFilter string,
	labelSelector labels.Selector,
	resources []string,
	fqdnTemplate, targetTemplate, hostTargetTemplate string,
	combineFqdnAnnotation bool,
) (Source, error) {
	fqdnTmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	targetTmpl, err := fqdn.ParseTemplate(targetTemplate)
	if err != nil {
		return nil, err
	}

	hostTargetTmpl, err := fqdn.ParseTemplate(hostTargetTemplate)
	if err != nil {
		return nil, err
	}

	gvrs, err := discoverResources(kubeClient, resources)
	if err != nil {
		return nil, err
	}

	// Create a single informer factory for all resources
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(
		dynamicClient,
		0,
		namespace,
		nil,
	)

	// Create informers for each resource
	resourceInformers := make([]kubeinformers.GenericInformer, 0, len(gvrs))
	for _, gvr := range gvrs {
		informer := informerFactory.ForResource(gvr)

		// Add indexers for efficient lookups by namespace and labels (must be before AddEventHandler)
		err := informer.Informer().AddIndexers(
			informers.IndexerWithOptions[*unstructured.Unstructured](
				informers.IndexSelectorWithAnnotationFilter(annotationFilter),
				informers.IndexSelectorWithLabelSelector(labelSelector),
			),
		)
		if err != nil {
			return nil, err
		}

		_, _ = informer.Informer().AddEventHandler(informers.DefaultEventHandler())
		resourceInformers = append(resourceInformers, informer)
	}

	informerFactory.Start(ctx.Done())
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &unstructuredSource{
		fqdnTemplate:          fqdnTmpl,
		targetTemplate:        targetTmpl,
		hostTargetTemplate:    hostTargetTmpl,
		informers:             resourceInformers,
		combineFqdnAnnotation: combineFqdnAnnotation,
	}, nil
}

// Endpoints returns the list of endpoints from unstructured resources.
func (us *unstructuredSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	for _, informer := range us.informers {
		resourceEndpoints, err := us.endpointsFromInformer(informer)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, resourceEndpoints...)
	}

	return endpoints, nil
}

// endpointsFromInformer returns endpoints for a single resource type.
func (us *unstructuredSource) endpointsFromInformer(informer kubeinformers.GenericInformer) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	// Get objects that match the indexer filter (annotation and label selectors)
	indexKeys := informer.Informer().GetIndexer().ListIndexFuncValues(informers.IndexWithSelectors)
	if len(indexKeys) == 0 {
		return nil, nil
	}
	for _, key := range indexKeys {
		obj, err := informers.GetByKey[*unstructured.Unstructured](informer.Informer().GetIndexer(), key)
		if err != nil {
			continue
		}

		el := newUnstructuredWrapper(obj)
		if el == nil {
			continue
		}

		if annotations.IsControllerMismatch(el, types.Unstructured) {
			continue
		}

		var edps []*endpoint.Endpoint
		// Get endpoints from annotations if no template or combining both
		if us.fqdnTemplate == nil || us.combineFqdnAnnotation {
			hosts := annotations.HostnamesFromAnnotations(el.GetAnnotations())
			addrs := annotations.TargetsFromTargetAnnotation(el.GetAnnotations())

			edps = EndpointsForHostsAndTargets(hosts, addrs)
		}

		if us.hostTargetTemplate != nil {
			edps, err = fqdn.CombineWithTemplatedEndpoints(
				edps, us.hostTargetTemplate, us.combineFqdnAnnotation,
				func() ([]*endpoint.Endpoint, error) {
					return us.endpointsFromHostTargetTemplate(el)
				},
			)
		} else if us.fqdnTemplate != nil {
			edps, err = fqdn.CombineWithTemplatedEndpoints(
				edps, us.fqdnTemplate, us.combineFqdnAnnotation,
				func() ([]*endpoint.Endpoint, error) {
					return us.endpointsFromTemplate(el)
				},
			)
		}
		if err != nil {
			return nil, err
		}

		ttl := annotations.TTLFromAnnotations(el.GetAnnotations(),
			fmt.Sprintf("%s/%s", strings.ToLower(el.GetKind()), el.GetName()))

		for _, ep := range edps {
			ep.
				WithRefObject(events.NewObjectReference(el, types.Unstructured)).
				WithLabel(endpoint.ResourceLabelKey,
					fmt.Sprintf("%s/%s/%s", strings.ToLower(el.GetKind()), el.GetNamespace(), el.GetName())).
				WithMinTTL(int64(ttl))
			endpoints = append(endpoints, ep)
		}
	}

	return MergeEndpoints(endpoints), nil
}

// endpointsFromTemplate creates endpoints using DNS names from the FQDN template.
func (us *unstructuredSource) endpointsFromTemplate(el *unstructuredWrapper) ([]*endpoint.Endpoint, error) {
	hostnames, err := fqdn.ExecTemplate(us.fqdnTemplate, el)
	if err != nil {
		return nil, err
	}

	if len(hostnames) == 0 {
		return nil, nil
	}

	targets, err := fqdn.ExecTemplate(us.targetTemplate, el)
	if err != nil {
		return nil, err
	}

	return EndpointsForHostsAndTargets(hostnames, targets), nil
}

// endpointsFromHostTargetTemplate creates endpoints from a template that returns host:target pairs.
// Each pair creates a single endpoint with 1:1 mapping between host and target.
func (us *unstructuredSource) endpointsFromHostTargetTemplate(el *unstructuredWrapper) ([]*endpoint.Endpoint, error) {
	pairs, err := fqdn.ExecTemplate(us.hostTargetTemplate, el)
	if err != nil {
		return nil, err
	}

	if len(pairs) == 0 {
		return nil, nil
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(pairs))
	for _, pair := range pairs {
		// Split at first colon (hostnames can't contain colons, IPv6 targets can)
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) != 2 {
			log.Debugf("Skipping invalid host:target pair %q from %s %s/%s: missing ':' separator",
				pair, strings.ToLower(el.GetKind()), el.GetNamespace(), el.GetName())
			continue
		}

		host := strings.TrimSpace(parts[0])
		target := strings.TrimSpace(parts[1])
		if host == "" || target == "" {
			log.Debugf("Skipping incomplete host:target pair %q from %s %s/%s: field may not yet be populated",
				pair, strings.ToLower(el.GetKind()), el.GetNamespace(), el.GetName())
			continue
		}

		endpoints = append(endpoints, endpoint.NewEndpoint(host, suitableType(target), target))
	}

	return MergeEndpoints(endpoints), nil
}

// AddEventHandler adds an event handler that is called when resources change.
func (us *unstructuredSource) AddEventHandler(_ context.Context, handler func()) {
	for _, informer := range us.informers {
		_, _ = informer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
}

// unstructuredWrapper wraps an unstructured.Unstructured to provide both
// typed-style template access ({{ .Name }}, {{ .Namespace }}) and raw map access
// ({{ .Spec.field }}, {{ index .Status.interfaces 0 "ipAddress" }}).
// By embedding *unstructured.Unstructured, it implements kubeObject (runtime.Object + metav1.Object).
type unstructuredWrapper struct {
	*unstructured.Unstructured

	// Typed-style convenience fields (like typed Kubernetes objects)
	Name        string
	Namespace   string
	Kind        string
	APIVersion  string
	Labels      map[string]string
	Annotations map[string]string

	// Raw map sections for custom field access
	Metadata map[string]any
	Spec     map[string]any
	Status   map[string]any

	// Full object for arbitrary access
	Object map[string]any
}

func (u *unstructuredWrapper) GetObjectMeta() metav1.Object {
	return u.Unstructured
}

// newUnstructuredWrapper creates a wrapper from a runtime.Object.
// Returns nil if the object is not an *unstructured.Unstructured.
func newUnstructuredWrapper(obj runtime.Object) *unstructuredWrapper {
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil
	}

	w := &unstructuredWrapper{
		Unstructured: u,
		Name:         u.GetName(),
		Namespace:    u.GetNamespace(),
		Kind:         u.GetKind(),
		APIVersion:   u.GetAPIVersion(),
		Labels:       u.GetLabels(),
		Annotations:  u.GetAnnotations(),
		Object:       u.Object,
	}

	// Extract common sections
	if metadata, ok := u.Object["metadata"].(map[string]any); ok {
		w.Metadata = metadata
	}
	if spec, ok := u.Object["spec"].(map[string]any); ok {
		w.Spec = spec
	}
	if status, ok := u.Object["status"].(map[string]any); ok {
		w.Status = status
	}

	return w
}

// discoverResources parses and validates resource identifiers against the cluster.
// It uses a cached discovery client to minimize API calls.
func discoverResources(kubeClient kubernetes.Interface, resources []string) ([]schema.GroupVersionResource, error) {
	cachedDiscovery := memory.NewMemCacheClient(kubeClient.Discovery())
	gvrs := make([]schema.GroupVersionResource, 0, len(resources))

	for _, r := range resources {
		// Handle core API resources (e.g., "configmaps.v1" -> "configmaps.v1.")
		if strings.Count(r, ".") == 1 {
			r += "."
		}

		gvr, _ := schema.ParseResourceArg(r)
		if gvr == nil {
			return nil, fmt.Errorf("invalid resource identifier %q: expected format resource.version.group (e.g., certificates.v1.cert-manager.io)", r)
		}

		if err := validateResource(cachedDiscovery, *gvr); err != nil {
			return nil, err
		}

		gvrs = append(gvrs, *gvr)
	}

	return gvrs, nil
}

// validateResource validates that a resource exists in the cluster.
// It uses the Discovery API to verify the resource is available.
func validateResource(discoveryClient discovery.DiscoveryInterface, gvr schema.GroupVersionResource) error {
	gv := gvr.GroupVersion().String()

	apiResourceList, err := discoveryClient.ServerResourcesForGroupVersion(gv)
	if err != nil {
		return fmt.Errorf("failed to discover resources for %q: %w", gv, err)
	}

	for i := range apiResourceList.APIResources {
		if apiResourceList.APIResources[i].Name == gvr.Resource {
			return nil
		}
	}

	return fmt.Errorf("resource %q not found in %q", gvr.Resource, gv)
}

// EndpointsForHostsAndTargets creates endpoints by grouping targets by record type
// and creating an endpoint for each hostname/record-type combination.
// The function returns endpoints in deterministic order (sorted by record type).
func EndpointsForHostsAndTargets(hostnames, targets []string) []*endpoint.Endpoint {
	if len(hostnames) == 0 || len(targets) == 0 {
		return nil
	}

	// Deduplicate hostnames
	hostSet := make(map[string]struct{}, len(hostnames))
	for _, h := range hostnames {
		hostSet[h] = struct{}{}
	}
	sortedHosts := slices.Sorted(maps.Keys(hostSet))

	// Group and deduplicate targets by record type
	targetsByType := make(map[string]map[string]struct{})
	for _, target := range targets {
		recordType := suitableType(target)
		if targetsByType[recordType] == nil {
			targetsByType[recordType] = make(map[string]struct{})
		}
		targetsByType[recordType][target] = struct{}{}
	}

	// Resolve to sorted slices once
	sortedTypes := slices.Sorted(maps.Keys(targetsByType))
	sortedTargets := make(map[string][]string, len(targetsByType))
	for _, recordType := range sortedTypes {
		sortedTargets[recordType] = slices.Sorted(maps.Keys(targetsByType[recordType]))
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(sortedHosts)*len(sortedTypes))
	for _, hostname := range sortedHosts {
		for _, recordType := range sortedTypes {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, recordType, sortedTargets[recordType]...))
		}
	}

	return endpoints
}
