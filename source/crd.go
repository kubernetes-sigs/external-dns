/*
Copyright 2018 The Kubernetes Authors.

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

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	toolscache "k8s.io/client-go/tools/cache"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/types"
)

// crdSource is an implementation of Source that provides endpoints by listing
// specified CRD and fetching Endpoints embedded in Spec.
//
// +externaldns:source:name=crd
// +externaldns:source:category=ExternalDNS
// +externaldns:source:description=Creates DNS entries from DNSEndpoint CRD resources
// +externaldns:source:resources=DNSEndpoint.externaldns.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=false
// +externaldns:source:events=true
// +externaldns:source:provider-specific=true
type crdSource struct {
	crReader client.Reader
	crWriter client.Client // status writes
	informer crcache.Informer
	listOpts []client.ListOption
}

type legacyCRDSource struct {
	crdClient        rest.Interface
	namespace        string
	crdResource      string
	codec            runtime.ParameterCodec
	annotationFilter labels.Selector
	labelSelector    labels.Selector
	informer         toolscache.SharedInformer
}

// NewCRDSource creates a new crdSource backed by a controller-runtime cache.
// It builds the scheme, cache, and status-write client from restConfig and cfg.
func NewCRDSource(ctx context.Context, restConfig *rest.Config, cfg *Config) (Source, error) {
	if !usesBuiltInDNSEndpointCRD(cfg) {
		return newLegacyCRDSource(restConfig, cfg)
	}

	opts, err := buildCacheOptions(cfg.Namespace, cfg.LabelFilter, cfg.AnnotationFilter)
	if err != nil {
		return nil, err
	}

	c, err := crcache.New(restConfig, opts)
	if err != nil {
		return nil, err
	}

	// crWriter is used exclusively for status writes; reads come from the cache.
	crWriter, err := client.New(restConfig, client.Options{Scheme: opts.Scheme})
	if err != nil {
		return nil, err
	}

	return newCrdSource(ctx, c, crWriter, cfg.Namespace, cfg.LabelFilter)
}

func usesBuiltInDNSEndpointCRD(cfg *Config) bool {
	apiVersion := cfg.CRDSourceAPIVersion
	if apiVersion == "" {
		apiVersion = apiv1alpha1.GroupVersion.String()
	}

	kind := cfg.CRDSourceKind
	if kind == "" {
		kind = "DNSEndpoint"
	}

	return apiVersion == apiv1alpha1.GroupVersion.String() && kind == "DNSEndpoint"
}

func (cs *crdSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("crd: adding event handler")
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	_, _ = cs.informer.AddEventHandler(eventHandlerFunc(handler))
}

func (cs *legacyCRDSource) AddEventHandler(_ context.Context, handler func()) {
	if cs.informer == nil {
		return
	}

	log.Debug("crd: adding event handler for custom CRD")
	informers.MustAddEventHandler(cs.informer, toolscache.ResourceEventHandlerFuncs{
		AddFunc:    func(_ any) { handler() },
		UpdateFunc: func(_, _ any) { handler() },
		DeleteFunc: func(_ any) { handler() },
	})
}

// Endpoints returns endpoint objects for all DNSEndpoint resources visible to
// this source. Namespace, label, and annotation filtering are handled at the
// cache level via buildCacheOptions; target-format validation is applied here.
func (cs *crdSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	list := &apiv1alpha1.DNSEndpointList{}
	if err := cs.crReader.List(ctx, list, cs.listOpts...); err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(list.Items))
	for i := range list.Items {
		dnsEndpoint := &list.Items[i]
		crdEndpoints := endpointsFromDNSEndpoint(dnsEndpoint)
		endpoints = append(endpoints, crdEndpoints...)

		if dnsEndpoint.Status.ObservedGeneration == dnsEndpoint.Generation {
			continue
		}

		dnsEndpoint.Status.ObservedGeneration = dnsEndpoint.Generation
		if err := cs.crWriter.Status().Update(ctx, dnsEndpoint); err != nil {
			log.Warnf("Could not update ObservedGeneration of [%s/%s/%s]: %v",
				"dnsendpoint", dnsEndpoint.Namespace, dnsEndpoint.Name, err)
		}
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

func (cs *legacyCRDSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	result, err := cs.List(ctx, &metav1.ListOptions{LabelSelector: cs.labelSelector.String()})
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(result.Items))
	for i := range result.Items {
		dnsEndpoint := &result.Items[i]
		if !matchLabelSelector(cs.annotationFilter, dnsEndpoint.Annotations) {
			continue
		}

		crdEndpoints := endpointsFromDNSEndpoint(dnsEndpoint)
		endpoints = append(endpoints, crdEndpoints...)

		if dnsEndpoint.Status.ObservedGeneration == dnsEndpoint.Generation {
			continue
		}

		dnsEndpoint.Status.ObservedGeneration = dnsEndpoint.Generation
		if _, err := cs.UpdateStatus(ctx, dnsEndpoint); err != nil {
			log.Warnf("Could not update ObservedGeneration of [%s/%s/%s]: %v",
				cs.crdResource, dnsEndpoint.Namespace, dnsEndpoint.Name, err)
		}
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

// newCrdSource wires a cache and writer into a running crdSource.
func newCrdSource(
	ctx context.Context,
	c crcache.Cache,
	crWriter client.Client,
	namespace string,
	labelSelector labels.Selector) (*crdSource, error) {
	inf, err := c.GetInformer(ctx, &apiv1alpha1.DNSEndpoint{})
	if err != nil {
		return nil, err
	}

	_, _ = inf.AddEventHandler(informers.DefaultEventHandler())

	listOpts := []client.ListOption{client.InNamespace(namespace)}
	if labelSelector != nil && !labelSelector.Empty() {
		listOpts = append(listOpts, client.MatchingLabelsSelector{Selector: labelSelector})
	}

	cs := &crdSource{
		crReader: c,
		crWriter: crWriter,
		informer: inf,
		listOpts: listOpts,
	}

	if err := startAndSync(ctx, c); err != nil {
		return nil, err
	}

	return cs, nil
}

// startAndSync starts the cache in a goroutine and waits for it to sync.
// Returns an error if the cache fails to start or sync.
func startAndSync(ctx context.Context, c crcache.Cache) error {
	errCh := make(chan error, 1)
	go func() { errCh <- c.Start(ctx) }()
	if !c.WaitForCacheSync(ctx) {
		select {
		case err := <-errCh:
			if err != nil {
				return fmt.Errorf("cache failed to sync: %w", err)
			}
			return fmt.Errorf("cache failed to sync")
		case <-ctx.Done():
			return fmt.Errorf("cache failed to sync: %w", ctx.Err())
		}
	}
	return nil
}

// buildCacheOptions constructs the controller-runtime cache options for the
// given namespace and label selector. Extracted so the namespace/label scoping
// logic can be unit-tested without a running API server.
func buildCacheOptions(namespace string, labelFilter, annotationSelector labels.Selector) (crcache.Options, error) {
	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return crcache.Options{}, err
	}
	// metav1.AddToGroupVersion registers ListOptions (and other meta types) under
	// apiv1alpha1.GroupVersion so that runtime.NewParameterCodec can encode them
	// as URL parameters when building watch requests for this group.
	metav1.AddToGroupVersion(scheme, apiv1alpha1.GroupVersion)

	nsMap := map[string]crcache.Config{
		namespace: {}, // "" == NamespaceAll
	}
	byObj := crcache.ByObject{
		Namespaces: nsMap,
		Transform: informers.TransformerWithOptions[*apiv1alpha1.DNSEndpoint](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
			informers.TransformRequireAnnotation(annotationSelector),
		),
	}
	if labelFilter != nil && !labelFilter.Empty() {
		byObj.Label = labelFilter
	}
	return crcache.Options{
		Scheme: scheme,
		ByObject: map[client.Object]crcache.ByObject{
			&apiv1alpha1.DNSEndpoint{}: byObj,
		},
	}, nil
}

func endpointsFromDNSEndpoint(dnsEndpoint *apiv1alpha1.DNSEndpoint) []*endpoint.Endpoint {
	var crdEndpoints []*endpoint.Endpoint

	for _, ep := range dnsEndpoint.Spec.Endpoints {
		if ep == nil {
			log.Debugf(
				"Skipping nil endpoint in DNSEndpoint %s/%s at spec.endpoints",
				dnsEndpoint.Namespace,
				dnsEndpoint.Name,
			)
			continue
		}

		if endpointHasIllegalCRDTarget(dnsEndpoint, ep) {
			continue
		}

		ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("crd/%s/%s", dnsEndpoint.Namespace, dnsEndpoint.Name))
		crdEndpoints = append(crdEndpoints, ep)
	}

	endpoint.AttachRefObject(crdEndpoints, events.NewObjectReference(dnsEndpoint, types.CRD))
	return crdEndpoints
}

func endpointHasIllegalCRDTarget(dnsEndpoint *apiv1alpha1.DNSEndpoint, ep *endpoint.Endpoint) bool {
	if (ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) && len(ep.Targets) < 1 {
		log.Debugf("Endpoint %s with DNSName %s has an empty list of targets, allowing it to pass through for default-targets processing", dnsEndpoint.Name, ep.DNSName)
	}

	for _, target := range ep.Targets {
		switch ep.RecordType {
		case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
			continue // no format constraint on targets
		case endpoint.RecordTypeCNAME:
			continue // RFC 1035 §5.1: trailing dot denotes an absolute FQDN in zone file notation; both forms are valid
		case endpoint.RecordTypeSRV:
			// SRV targets are "<prio> <weight> <port> <host>"; RFC 2782
			// requires the host to be an absolute FQDN and
			// Endpoint.ValidateSRVRecord enforces the trailing dot.
			// Reject-on-trailing-dot (the default branch below) would
			// loop users between this warning and ValidateSRVRecord's
			// "does not end with a dot" error (#6357).
			continue
		}

		hasDot := strings.HasSuffix(target, ".")
		illegalTarget := hasDot
		if ep.RecordType == endpoint.RecordTypeNAPTR {
			illegalTarget = !hasDot
		}

		if illegalTarget {
			fixed := target + "."
			if ep.RecordType != endpoint.RecordTypeNAPTR {
				fixed = strings.TrimSuffix(target, ".")
			}
			log.Warnf("Endpoint %s/%s with DNSName %s has an illegal target %q for %s record — use %q not %q.",
				dnsEndpoint.Namespace, dnsEndpoint.Name, ep.DNSName, target, ep.RecordType, fixed, target)
			return true
		}
	}

	return false
}

func newLegacyCRDSource(restConfig *rest.Config, cfg *Config) (Source, error) {
	crdClient, scheme, resource, err := newLegacyCRDClient(restConfig, cfg)
	if err != nil {
		return nil, err
	}

	sourceCrd := &legacyCRDSource{
		crdResource:      resource,
		namespace:        cfg.Namespace,
		annotationFilter: cfg.AnnotationFilter,
		labelSelector:    cfg.LabelFilter,
		crdClient:        crdClient,
		codec:            runtime.NewParameterCodec(scheme),
	}
	if cfg.UpdateEvents {
		// external-dns already runs its sync-handler periodically to cover missed events.
		sourceCrd.informer = toolscache.NewSharedInformer(
			&toolscache.ListWatch{
				ListWithContextFunc: func(ctx context.Context, lo metav1.ListOptions) (runtime.Object, error) {
					return sourceCrd.List(ctx, &lo)
				},
				WatchFuncWithContext: func(ctx context.Context, lo metav1.ListOptions) (watch.Interface, error) {
					return sourceCrd.watch(ctx, &lo)
				},
			},
			&apiv1alpha1.DNSEndpoint{},
			0,
		)
		go sourceCrd.informer.Run(wait.NeverStop)
	}

	return sourceCrd, nil
}

func newLegacyCRDClient(restConfig *rest.Config, cfg *Config) (rest.Interface, *runtime.Scheme, string, error) {
	groupVersion, err := schema.ParseGroupVersion(cfg.CRDSourceAPIVersion)
	if err != nil {
		return nil, nil, "", err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(restConfig)
	if err != nil {
		return nil, nil, "", err
	}

	apiResourceList, err := discoveryClient.ServerResourcesForGroupVersion(groupVersion.String())
	if err != nil {
		return nil, nil, "", fmt.Errorf("error listing resources in GroupVersion %q: %w", groupVersion.String(), err)
	}

	var crdResource string
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Kind == cfg.CRDSourceKind {
			crdResource = apiResource.Name
			break
		}
	}
	if crdResource == "" {
		return nil, nil, "", fmt.Errorf("unable to find Resource Kind %q in GroupVersion %q", cfg.CRDSourceKind, cfg.CRDSourceAPIVersion)
	}

	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return nil, nil, "", err
	}
	scheme.AddKnownTypes(groupVersion, &apiv1alpha1.DNSEndpoint{}, &apiv1alpha1.DNSEndpointList{})
	metav1.AddToGroupVersion(scheme, groupVersion)

	config := rest.CopyConfig(restConfig)
	config.GroupVersion = &groupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{
		CodecFactory: serializer.NewCodecFactory(scheme),
	}

	crdClient, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, nil, "", err
	}

	return crdClient, scheme, crdResource, nil
}

func (cs *legacyCRDSource) watch(ctx context.Context, opts *metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return cs.crdClient.Get().
		Namespace(cs.namespace).
		Resource(cs.crdResource).
		VersionedParams(opts, cs.codec).
		Watch(ctx)
}

func (cs *legacyCRDSource) List(ctx context.Context, opts *metav1.ListOptions) (*apiv1alpha1.DNSEndpointList, error) {
	result := &apiv1alpha1.DNSEndpointList{}
	return result, cs.crdClient.Get().
		Namespace(cs.namespace).
		Resource(cs.crdResource).
		VersionedParams(opts, cs.codec).
		Do(ctx).
		Into(result)
}

func (cs *legacyCRDSource) UpdateStatus(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint) (*apiv1alpha1.DNSEndpoint, error) {
	result := &apiv1alpha1.DNSEndpoint{}
	return result, cs.crdClient.Put().
		Namespace(dnsEndpoint.Namespace).
		Resource(cs.crdResource).
		Name(dnsEndpoint.Name).
		SubResource("status").
		Body(dnsEndpoint).
		Do(ctx).
		Into(result)
}
