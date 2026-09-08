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
	rgv1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"
	rgversioned "github.com/szuecs/routegroup-client/client/clientset/versioned"
	rginformers "github.com/szuecs/routegroup-client/client/informers/externalversions"
	rginformersv1 "github.com/szuecs/routegroup-client/client/informers/externalversions/zalando.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
	"sigs.k8s.io/external-dns/source/types"
)

const (
	// DefaultRoutegroupVersion is the default version for route groups.
	//
	// Deprecated: the informer-based implementation is hardwired to zalando.org/v1.
	DefaultRoutegroupVersion = "zalando.org/v1"
)

// +externaldns:source:name=skipper-routegroup
// +externaldns:source:category=Ingress Controllers
// +externaldns:source:description=Creates DNS entries from Skipper RouteGroup resources
// +externaldns:source:resources=RouteGroup.zalando.org
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
// +externaldns:source:events=true
type routeGroupSource struct {
	annotationFilter         labels.Selector
	labelSelector            labels.Selector
	templateEngine           template.Engine
	ignoreHostnameAnnotation bool
	rgInformer               rginformersv1.RouteGroupInformer
}

// routeGroupWrapper adds a Metadata() accessor to *rgv1.RouteGroup so that
// fqdn templates using {{.Metadata.Name}} continue to work alongside the
// canonical {{.Name}} form.
//
// Deprecated: use top-level fields directly (e.g. {{.Name}} instead of {{.Metadata.Name}}).
type routeGroupWrapper struct {
	*rgv1.RouteGroup
}

// Metadata returns the ObjectMeta for backward-compatible template access.
//
// Deprecated: use top-level fields directly (e.g. {{.Name}} instead of {{.Metadata.Name}}).
func (w *routeGroupWrapper) Metadata() *metav1.ObjectMeta {
	return &w.ObjectMeta
}

// NewRouteGroupSource creates a new routeGroupSource with the given config.
func NewRouteGroupSource(ctx context.Context, client rgversioned.Interface, cfg *Config) (Source, error) {
	informerFactory := rginformers.NewSharedInformerFactoryWithOptions(
		client, 0,
		rginformers.WithNamespace(cfg.Namespace),
	)
	rgInformer := informerFactory.Zalando().V1().RouteGroups()

	informers.MustAddIndexers(rgInformer.Informer(), informers.IndexerWithOptions[*rgv1.RouteGroup](
		informers.IndexSelectorWithAnnotationFilter(cfg.AnnotationFilter),
		informers.IndexSelectorWithLabelSelector(cfg.LabelFilter),
		informers.IndexSelectorWithConditions(annotations.IsControllerMatch[*rgv1.RouteGroup]),
	))

	informers.MustSetTransform(rgInformer.Informer(), informers.TransformerWithOptions[*rgv1.RouteGroup](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
	))

	// Add default resource event handlers to properly initialize informer.
	informers.MustAddEventHandler(rgInformer.Informer(), informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &routeGroupSource{
		annotationFilter:         cfg.AnnotationFilter,
		labelSelector:            cfg.LabelFilter,
		templateEngine:           cfg.TemplateEngine,
		ignoreHostnameAnnotation: cfg.IgnoreHostnameAnnotation,
		rgInformer:               rgInformer,
	}, nil
}

// AddEventHandler adds an event handler that can be triggered on RouteGroup changes.
func (sc *routeGroupSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for routegroup")
	informers.MustAddEventHandler(sc.rgInformer.Informer(), eventHandlerFunc(handler))
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all routeGroup resources on all namespaces.
// Logic is ported from ingress without fqdnTemplate
func (sc *routeGroupSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	routeGroups := informers.ListIndexed[*rgv1.RouteGroup](sc.rgInformer.Informer().GetIndexer())

	var endpoints []*endpoint.Endpoint
	for _, rg := range routeGroups {
		eps := sc.endpointsFromRouteGroup(rg)

		var err error
		eps, err = sc.templateEngine.CombineWithEndpoints(
			eps,
			func() ([]*endpoint.Endpoint, error) { return sc.endpointsFromTemplate(rg) },
		)
		if err != nil {
			return nil, err
		}

		if endpoint.HasNoEmptyEndpoints(eps, types.SkipperRouteGroup, rg) {
			continue
		}

		endpoint.AttachRefObject(eps, events.NewObjectReference(rg, types.SkipperRouteGroup))

		log.Debugf("Endpoints generated from routegroup: %s/%s: %v", rg.Namespace, rg.Name, eps)
		endpoints = append(endpoints, eps...)
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

func (sc *routeGroupSource) endpointsFromTemplate(rg *rgv1.RouteGroup) ([]*endpoint.Endpoint, error) {
	hostnames, err := sc.templateEngine.ExecFQDN(&routeGroupWrapper{rg})
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf("routegroup/%s/%s", rg.Namespace, rg.Name)

	// error handled in endpointsFromRouteGroup(), otherwise duplicate log
	ttl := annotations.TTLFromAnnotations(rg.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(rg.Annotations)

	if len(targets) == 0 {
		targets = targetsFromRouteGroupStatus(rg.Status)
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(rg.Annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		endpoints = append(endpoints, endpoint.EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}
	return endpoints, nil
}

func (sc *routeGroupSource) endpointsFromRouteGroup(rg *rgv1.RouteGroup) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	resource := fmt.Sprintf("routegroup/%s/%s", rg.Namespace, rg.Name)

	ttl := annotations.TTLFromAnnotations(rg.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(rg.Annotations)
	if len(targets) == 0 {
		for _, lb := range rg.Status.LoadBalancer.RouteGroup {
			if lb.IP != "" {
				targets = append(targets, lb.IP)
			}
			if lb.Hostname != "" {
				targets = append(targets, lb.Hostname)
			}
		}
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(rg.Annotations)

	for _, src := range rg.Spec.Hosts {
		if src == "" {
			continue
		}
		endpoints = append(endpoints, endpoint.EndpointsForHostname(src, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(rg.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpoint.EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}
	return endpoints
}

func targetsFromRouteGroupStatus(status rgv1.RouteGroupStatus) endpoint.Targets {
	var targets endpoint.Targets

	for _, lb := range status.LoadBalancer.RouteGroup {
		if lb.IP != "" {
			targets = append(targets, lb.IP)
		}
		if lb.Hostname != "" {
			targets = append(targets, lb.Hostname)
		}
	}

	return targets
}
