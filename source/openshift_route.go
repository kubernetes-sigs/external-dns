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
	"text/template"
	"time"

	routev1 "github.com/openshift/api/route/v1"
	"github.com/openshift/client-go/route/clientset/versioned"
	extInformers "github.com/openshift/client-go/route/informers/externalversions"
	routeInformer "github.com/openshift/client-go/route/informers/externalversions/route/v1"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/fqdn"
	"sigs.k8s.io/external-dns/source/informers"
)

// ocpRouteSource is an implementation of Source for OpenShift Route objects.
// The Route implementation will use the Route spec.host field for the hostname,
// and the Route status' canonicalHostname field as the target.
// The annotations.TargetKey can be used to explicitly set an alternative
// endpoint, if desired.
//
// +externaldns:source:name=openshift-route
// +externaldns:source:category=OpenShift
// +externaldns:source:description=Creates DNS entries from OpenShift Route resources
// +externaldns:source:resources=Route.route.openshift.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
type ocpRouteSource struct {
	client                   versioned.Interface
	namespace                string
	annotationFilter         string
	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
	routeInformer            routeInformer.RouteInformer
	labelSelector            labels.Selector
	ocpRouterName            string
}

// NewOcpRouteSource creates a new ocpRouteSource with the given config.
func NewOcpRouteSource(
	ctx context.Context,
	ocpClient versioned.Interface,
	namespace string,
	annotationFilter string,
	fqdnTemplate string,
	combineFQDNAnnotation bool,
	ignoreHostnameAnnotation bool,
	labelSelector labels.Selector,
	ocpRouterName string,
) (Source, error) {
	tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	// Use a shared informer to listen for add/update/delete of Routes in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := extInformers.NewSharedInformerFactoryWithOptions(ocpClient, 0*time.Second, extInformers.WithNamespace(namespace))
	informer := informerFactory.Route().V1().Routes()

	// Add default resource event handlers to properly initialize informer.
	_, _ = informer.Informer().AddEventHandler(informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &ocpRouteSource{
		client:                   ocpClient,
		namespace:                namespace,
		annotationFilter:         annotationFilter,
		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    combineFQDNAnnotation,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
		routeInformer:            informer,
		labelSelector:            labelSelector,
		ocpRouterName:            ocpRouterName,
	}, nil
}

func (ors *ocpRouteSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for openshift route")

	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	_, _ = ors.routeInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all OpenShift Route resources on all namespaces, unless an explicit namespace
// is specified in ocpRouteSource.
func (ors *ocpRouteSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	ocpRoutes, err := ors.routeInformer.Lister().Routes(ors.namespace).List(ors.labelSelector)
	if err != nil {
		return nil, err
	}

	ocpRoutes, err = annotations.Filter(ocpRoutes, ors.annotationFilter)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, ocpRoute := range ocpRoutes {
		if annotations.IsControllerMismatch(ocpRoute, types.OpenShiftRoute) {
			continue
		}

		orEndpoints := ors.endpointsFromOcpRoute(ocpRoute, ors.ignoreHostnameAnnotation)

		// apply template if host is missing on OpenShift Route
		orEndpoints, err = fqdn.CombineWithTemplatedEndpoints(
			orEndpoints,
			ors.fqdnTemplate,
			ors.combineFQDNAnnotation,
			func() ([]*endpoint.Endpoint, error) { return ors.endpointsFromTemplate(ocpRoute) },
		)
		if err != nil {
			return nil, err
		}

		if endpoint.HasNoEmptyEndpoints(orEndpoints, types.OpenShiftRoute, ocpRoute) {
			continue
		}

		log.Debugf("Endpoints generated from OpenShift Route: %s/%s: %v", ocpRoute.Namespace, ocpRoute.Name, orEndpoints)
		endpoints = append(endpoints, orEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
}

func (ors *ocpRouteSource) endpointsFromTemplate(ocpRoute *routev1.Route) ([]*endpoint.Endpoint, error) {
	hostnames, err := fqdn.ExecTemplate(ors.fqdnTemplate, ocpRoute)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf("route/%s/%s", ocpRoute.Namespace, ocpRoute.Name)

	ttl := annotations.TTLFromAnnotations(ocpRoute.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(ocpRoute.Annotations)
	if len(targets) == 0 {
		targetsFromRoute, _ := ors.getTargetsFromRouteStatus(ocpRoute.Status)
		targets = targetsFromRoute
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ocpRoute.Annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}
	return endpoints, nil
}

// endpointsFromOcpRoute extracts the endpoints from a OpenShift Route object
func (ors *ocpRouteSource) endpointsFromOcpRoute(ocpRoute *routev1.Route, ignoreHostnameAnnotation bool) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("route/%s/%s", ocpRoute.Namespace, ocpRoute.Name)

	ttl := annotations.TTLFromAnnotations(ocpRoute.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(ocpRoute.Annotations)
	targetsFromRoute, host := ors.getTargetsFromRouteStatus(ocpRoute.Status)

	if len(targets) == 0 {
		targets = targetsFromRoute
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(ocpRoute.Annotations)

	if host != "" {
		endpoints = append(endpoints, EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(ocpRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}
	return endpoints
}

// getTargetsFromRouteStatus returns the router's canonical hostname and host
// either for the given router if it admitted the route
// or for the first (in the status list) router that admitted the route.
func (ors *ocpRouteSource) getTargetsFromRouteStatus(status routev1.RouteStatus) (endpoint.Targets, string) {
	for _, ing := range status.Ingress {
		// if this Ingress didn't admit the route or it doesn't have the canonical hostname, then ignore it
		if ingressConditionStatus(&ing, routev1.RouteAdmitted) != corev1.ConditionTrue || ing.RouterCanonicalHostname == "" {
			continue
		}

		// if the router name is specified for the Route source and it matches the route's ingress name, then return it
		if ors.ocpRouterName != "" && ors.ocpRouterName == ing.RouterName {
			return endpoint.Targets{ing.RouterCanonicalHostname}, ing.Host
		}

		// if the router name is not specified in the Route source then return the first ingress
		if ors.ocpRouterName == "" {
			return endpoint.Targets{ing.RouterCanonicalHostname}, ing.Host
		}
	}
	return endpoint.Targets{}, ""
}

func ingressConditionStatus(ingress *routev1.RouteIngress, t routev1.RouteIngressConditionType) corev1.ConditionStatus {
	for _, condition := range ingress.Conditions {
		if t != condition.Type {
			continue
		}
		return condition.Status
	}
	return corev1.ConditionUnknown
}
