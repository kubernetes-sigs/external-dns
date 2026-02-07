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
	namespace             string
	annotationFilter      string
	labelSelector         labels.Selector
	combineFqdnAnnotation bool
	fqdnTemplate          *template.Template
	targetFqdnTemplate    *template.Template
	resources             []resourceConfig
}

// NewUnstructuredFQDNSource creates a new unstructuredSource.
func NewUnstructuredFQDNSource(
	ctx context.Context,
	dynamicClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	namespace, annotationFilter string,
	labelSelector labels.Selector,
	resources []string,
	fqdnTemplate, targetFqdnTemplate string,
	combineFqdnAnnotation bool,
) (Source, error) {
	fqdnTmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	targetTmpl, err := fqdn.ParseTemplate(targetFqdnTemplate)
	if err != nil {
		return nil, err
	}

	// Discover and validate all resources using cached discovery client
	cachedDiscovery := memory.NewMemCacheClient(kubeClient.Discovery())
	resourceConfigs := make([]resourceConfig, 0, len(resources))
	for _, r := range resources {
		gvr, err := parseResourceIdentifier(r)
		if err != nil {
			return nil, err
		}

		rc, err := discoverResource(cachedDiscovery, gvr)
		if err != nil {
			return nil, err
		}

		resourceConfigs = append(resourceConfigs, *rc)
	}

	// Create a single informer factory for all resources
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(
		dynamicClient,
		0,
		namespace,
		nil,
	)

	// Create informers for each resource
	for i := range resourceConfigs {
		resourceConfigs[i].informer = informerFactory.ForResource(resourceConfigs[i].gvr)

		// Add indexers for efficient lookups by namespace and labels (must be before AddEventHandler)
		err := resourceConfigs[i].informer.Informer().AddIndexers(
			informers.IndexerWithOptions[*unstructured.Unstructured](
				informers.IndexSelectorWithAnnotationFilter(annotationFilter),
				informers.IndexSelectorWithLabelSelector(labelSelector),
			),
		)
		if err != nil {
			return nil, err
		}

		_, _ = resourceConfigs[i].informer.Informer().AddEventHandler(informers.DefaultEventHandler())
	}

	informerFactory.Start(ctx.Done())
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &unstructuredSource{
		namespace:             namespace,
		annotationFilter:      annotationFilter,
		labelSelector:         labelSelector,
		fqdnTemplate:          fqdnTmpl,
		targetFqdnTemplate:    targetTmpl,
		resources:             resourceConfigs,
		combineFqdnAnnotation: combineFqdnAnnotation,
	}, nil
}

// Endpoints returns the list of endpoints from unstructured resources.
func (us *unstructuredSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	for _, rc := range us.resources {
		resourceEndpoints, err := us.endpointsForResource(ctx, rc)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, resourceEndpoints...)
	}

	return endpoints, nil
}

// endpointsForResource returns endpoints for a single resource type.
func (us *unstructuredSource) endpointsForResource(_ context.Context, rc resourceConfig) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	// Get objects that match the indexer filter (annotation and label selectors)
	indexKeys := rc.informer.Informer().GetIndexer().ListIndexFuncValues(informers.IndexWithSelectors)

	for _, key := range indexKeys {
		obj, err := informers.GetByKey[*unstructured.Unstructured](rc.informer.Informer().GetIndexer(), key)
		if err != nil {
			log.Debugf("failed to get object by key %q: %v", key, err)
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

			edps = endpoint.EndpointsForHostsAndTargets(hosts, addrs)
		}

		if us.targetFqdnTemplate != nil {
			edps, err = fqdn.CombineWithTemplatedEndpoints(
				edps,
				us.fqdnTemplate,
				us.combineFqdnAnnotation,
				func() ([]*endpoint.Endpoint, error) {
					return us.endpointsFromTemplate(el)
				},
			)
			if err != nil {
				return nil, err
			}
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

	targets, err := fqdn.ExecTemplate(us.targetFqdnTemplate, el)
	if err != nil {
		return nil, err
	}

	return endpoint.EndpointsForHostsAndTargets(hostnames, targets), nil
}

// AddEventHandler adds an event handler that is called when resources change.
func (us *unstructuredSource) AddEventHandler(_ context.Context, handler func()) {
	for _, rc := range us.resources {
		_, _ = rc.informer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
}

// resourceConfig holds the parsed configuration for a single resource type.
type resourceConfig struct {
	gvr      schema.GroupVersionResource
	informer kubeinformers.GenericInformer
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

// parseResourceIdentifier parses a resource identifier in the format "resource.version.group"
// (e.g., "virtualmachineinstances.v1.kubevirt.io") and returns a GroupVersionResource.
//
// Format: resource.version.group
// - resource: plural resource name (e.g., "vmachines")
// - version: API version (e.g., "v1", "v1beta1")
// - group: API group (e.g., "kubevirt.io", "apps")
//
// For core API resources (e.g., pods.v1), the group is empty.
func parseResourceIdentifier(identifier string) (schema.GroupVersionResource, error) {
	parts := strings.SplitN(identifier, ".", 3)
	if len(parts) < 2 {
		return schema.GroupVersionResource{}, fmt.Errorf("invalid resource identifier %q: expected format resource.version.group (e.g., virtualmachineinstances.v1.kubevirt.io)", identifier)
	}

	resource := parts[0]
	version := parts[1]
	group := ""
	if len(parts) == 3 {
		group = parts[2]
	}

	if resource == "" {
		return schema.GroupVersionResource{}, fmt.Errorf("invalid resource identifier %q: resource name cannot be empty", identifier)
	}
	if version == "" {
		return schema.GroupVersionResource{}, fmt.Errorf("invalid resource identifier %q: version cannot be empty", identifier)
	}

	return schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}, nil
}

// discoverResource validates that a resource exists in the cluster and returns its configuration.
// It uses the Discovery API to verify the resource and determine if it's namespaced.
func discoverResource(discoveryClient discovery.DiscoveryInterface, gvr schema.GroupVersionResource) (*resourceConfig, error) {
	gv := gvr.GroupVersion().String()

	apiResourceList, err := discoveryClient.ServerResourcesForGroupVersion(gv)
	if err != nil {
		return nil, fmt.Errorf("failed to discover resources for %q: %w", gv, err)
	}

	var apiResource *metav1.APIResource
	for i := range apiResourceList.APIResources {
		ar := &apiResourceList.APIResources[i]
		if ar.Name == gvr.Resource {
			apiResource = ar
			break
		}
	}

	if apiResource == nil {
		return nil, fmt.Errorf("resource %q not found in %q", gvr.Resource, gv)
	}

	return &resourceConfig{
		gvr: gvr,
	}, nil
}
