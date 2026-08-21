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
	"strings"
	"unicode"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
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
// +externaldns:source:provider-specific=false
// +externaldns:source:events=false
type unstructuredSource struct {
	templateEngine template.Engine
	informers      []kubeinformers.GenericInformer
}

// NewUnstructuredFQDNSource creates a new unstructuredSource.
func NewUnstructuredFQDNSource(
	ctx context.Context,
	dynamicClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	cfg *Config,
) (Source, error) {
	gvrs, err := discoverResources(kubeClient, cfg.UnstructuredResources)
	if err != nil {
		return nil, err
	}

	// Create a single informer factory for all resources
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(
		dynamicClient,
		0,
		cfg.Namespace,
		nil,
	)

	// Create informers for each resource
	resourceInformers := make([]kubeinformers.GenericInformer, 0, len(gvrs))
	for _, gvr := range gvrs {
		informer := informerFactory.ForResource(gvr)

		// Add indexers for efficient lookups by namespace and labels (must be before AddEventHandler)
		informers.MustAddIndexers(informer.Informer(), informers.IndexerWithOptions[*unstructured.Unstructured](
			informers.IndexSelectorWithAnnotationFilter(cfg.AnnotationFilter),
			informers.IndexSelectorWithLabelSelector(cfg.LabelFilter),
			informers.IndexSelectorWithConditions(annotations.IsControllerMatch[*unstructured.Unstructured]),
		))
		informers.MustSetTransform(informer.Informer(), informers.TransformerWithOptions[*unstructured.Unstructured](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
		))

		informers.MustAddEventHandler(informer.Informer(), informers.DefaultEventHandler())
		resourceInformers = append(resourceInformers, informer)
	}

	informerFactory.Start(ctx.Done())
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &unstructuredSource{
		templateEngine: cfg.TemplateEngine,
		informers:      resourceInformers,
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

		hosts := annotations.HostnamesFromAnnotations(el.GetAnnotations())
		addrs := annotations.TargetsFromTargetAnnotation(el.GetAnnotations())
		annotationEdps := endpoint.EndpointsForHostsAndTargets(hosts, addrs)

		edps, err := us.templateEngine.ApplyTemplates(annotationEdps, el)
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

	return endpoint.MergeEndpoints(endpoints), nil
}

// AddEventHandler adds an event handler that is called when resources change.
func (us *unstructuredSource) AddEventHandler(_ context.Context, handler func()) {
	for _, informer := range us.informers {
		informers.MustAddEventHandler(informer.Informer(), eventHandlerFunc(handler))
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
}

func (u *unstructuredWrapper) GetObjectMeta() metav1.Object {
	return u.Unstructured
}

// newUnstructuredWrapper creates a wrapper around an *unstructured.Unstructured,
// exposing typed convenience fields for templateEngine alongside raw map sections.
func newUnstructuredWrapper(u *unstructured.Unstructured) *unstructuredWrapper {
	w := &unstructuredWrapper{
		Unstructured: u,
		Name:         u.GetName(),
		Namespace:    u.GetNamespace(),
		Kind:         u.GetKind(),
		APIVersion:   u.GetAPIVersion(),
		Labels:       u.GetLabels(),
		Annotations:  u.GetAnnotations(),
	}

	// Extract common sections
	if metadata, ok := u.Object["metadata"].(map[string]any); ok {
		w.Metadata = metadata
	}
	if spec, ok := u.Object["spec"].(map[string]any); ok {
		w.Spec = withTitleCaseAliases(spec)
	}
	if status, ok := u.Object["status"].(map[string]any); ok {
		w.Status = withTitleCaseAliases(status)
	}

	return w
}

// withTitleCaseAliases lets a template address a JSON-keyed map field
// (spec.hostnames) by its Go field name too (Spec.Hostnames), since a map
// key miss in text/template renders empty instead of erroring.
//
// Deliberately shallow: an earlier version recursed into nested map values
// to alias their keys too, but that duplicated every key one level down, so
// a template ranging or taking len() over nested data (spec.records) saw
// twice as many entries as were actually declared. Nested JSON keys stay
// reachable through their literal path (spec.endpoint.hostname) instead.
func withTitleCaseAliases(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	maps.Copy(out, m)
	for k := range m {
		title := titleCaseKey(k)
		if title == k {
			continue
		}
		if _, collision := m[title]; collision {
			continue
		}
		out[title] = out[k]
	}
	return out
}

// commonInitialisms are the field-name segments Kubernetes' code generator
// spells fully upper-case in Go structs (DNSNames, not DnsNames; URL, not
// Url), taken from the CRD/API-machinery-relevant subset of the initialisms
// staticcheck/golint treat as standard Go convention.
var commonInitialisms = map[string]bool{
	"acl": true, "api": true, "arn": true, "ca": true, "cidr": true,
	"cpu": true, "crd": true, "db": true, "dns": true, "http": true,
	"https": true, "id": true, "ip": true, "json": true, "jwk": true,
	"jwt": true, "os": true, "pem": true, "rbac": true, "sql": true,
	"ssl": true, "tcp": true, "tls": true, "ttl": true, "udp": true,
	"uid": true, "uuid": true, "uri": true, "url": true, "vm": true,
	"xml": true, "yaml": true,
}

// titleCaseKey converts a lowerCamelCase JSON key into its Go-convention
// exported field name: dnsNames -> DNSNames, url -> URL, hostname ->
// Hostname. It upper-cases the leading word in full when that word is a
// known initialism, matching how Kubernetes generates Go types from CRD
// schemas; otherwise it capitalizes just the first letter, as before.
func titleCaseKey(k string) string {
	if k == "" {
		return k
	}
	r := []rune(k)
	if unicode.IsUpper(r[0]) {
		return k
	}
	end := 0
	for end < len(r) && !unicode.IsUpper(r[end]) {
		end++
	}
	word := string(r[:end])
	if commonInitialisms[strings.ToLower(word)] {
		return strings.ToUpper(word) + string(r[end:])
	}
	return strings.ToUpper(word[:1]) + word[1:] + string(r[end:])
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
