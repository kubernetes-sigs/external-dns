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
	"sort"
	"text/template"

	routev1 "github.com/openshift/api/route/v1"
	versioned "github.com/openshift/client-go/route/clientset/versioned"
	extInformers "github.com/openshift/client-go/route/informers/externalversions"
	routeInformer "github.com/openshift/client-go/route/informers/externalversions/route/v1"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
)

// ocpRouteSource is an implementation of Source for OpenShift Route objects.
// The Route implementation will use the Route spec.host field for the hostname,
// and the Route status' canonicalHostname field as the target.
// The targetAnnotationKey can be used to explicitly set an alternative
// endpoint, if desired.
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
	tmpl, err := parseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	// Use a shared informer to listen for add/update/delete of Routes in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := extInformers.NewSharedInformerFactoryWithOptions(ocpClient, 0, extInformers.WithNamespace(namespace))
	informer := informerFactory.Route().V1().Routes()

	// Add default resource event handlers to properly initialize informer.
	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := waitForCacheSync(context.Background(), informerFactory); err != nil {
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

func (ors *ocpRouteSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("Adding event handler for openshift route")

	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	ors.routeInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all OpenShift Route resources on all namespaces, unless an explicit namespace
// is specified in ocpRouteSource.
func (ors *ocpRouteSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	ocpRoutes, err := ors.routeInformer.Lister().Routes(ors.namespace).List(ors.labelSelector)
	if err != nil {
		return nil, err
	}

	ocpRoutes, err = ors.filterByAnnotations(ocpRoutes)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, ocpRoute := range ocpRoutes {
		// Check controller annotation to see if we are responsible.
		controller, ok := ocpRoute.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping OpenShift Route %s/%s because controller value does not match, found: %s, required: %s",
				ocpRoute.Namespace, ocpRoute.Name, controller, controllerAnnotationValue)
			continue
		}

		orEndpoints := ors.endpointsFromOcpRoute(ocpRoute, ors.ignoreHostnameAnnotation)

		// apply template if host is missing on OpenShift Route
		if (ors.combineFQDNAnnotation || len(orEndpoints) == 0) && ors.fqdnTemplate != nil {
			oEndpoints, err := ors.endpointsFromTemplate(ocpRoute)
			if err != nil {
				return nil, err
			}

			if ors.combineFQDNAnnotation {
				orEndpoints = append(orEndpoints, oEndpoints...)
			} else {
				orEndpoints = oEndpoints
			}
		}

		if len(orEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from OpenShift Route %s/%s", ocpRoute.Namespace, ocpRoute.Name)
			continue
		}

		log.Debugf("Endpoints generated from OpenShift Route: %s/%s: %v", ocpRoute.Namespace, ocpRoute.Name, orEndpoints)
		ors.setResourceLabel(ocpRoute, orEndpoints)
		endpoints = append(endpoints, orEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (ors *ocpRouteSource) endpointsFromTemplate(ocpRoute *routev1.Route) ([]*endpoint.Endpoint, error) {
	hostnames, err := execTemplate(ors.fqdnTemplate, ocpRoute)
	if err != nil {
		return nil, err
	}

	ttl, err := getTTLFromAnnotations(ocpRoute.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets := getTargetsFromTargetAnnotation(ocpRoute.Annotations)
	if len(targets) == 0 {
		targetsFromRoute, _ := ors.getTargetsFromRouteStatus(ocpRoute.Status)
		targets = targetsFromRoute
	}

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ocpRoute.Annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
	}
	return endpoints, nil
}

func (ors *ocpRouteSource) filterByAnnotations(ocpRoutes []*routev1.Route) ([]*routev1.Route, error) {
	labelSelector, err := metav1.ParseToLabelSelector(ors.annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return ocpRoutes, nil
	}

	filteredList := []*routev1.Route{}

	for _, ocpRoute := range ocpRoutes {
		// convert the Route's annotations to an equivalent label selector
		annotations := labels.Set(ocpRoute.Annotations)

		// include ocpRoute if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, ocpRoute)
		}
	}

	return filteredList, nil
}

func (ors *ocpRouteSource) setResourceLabel(ocpRoute *routev1.Route, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("route/%s/%s", ocpRoute.Namespace, ocpRoute.Name)
	}
}

// endpointsFromOcpRoute extracts the endpoints from a OpenShift Route object
func (ors *ocpRouteSource) endpointsFromOcpRoute(ocpRoute *routev1.Route, ignoreHostnameAnnotation bool) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(ocpRoute.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets := getTargetsFromTargetAnnotation(ocpRoute.Annotations)
	targetsFromRoute, host := ors.getTargetsFromRouteStatus(ocpRoute.Status)

	if len(targets) == 0 {
		targets = targetsFromRoute
	}

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ocpRoute.Annotations)

	if host != "" {
		endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(ocpRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
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
