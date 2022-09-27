/*
Copyright 2022 The Kubernetes Authors.

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

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/external-dns/endpoint"

	traefikV1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
)

var (
	ingressrouteGVR = schema.GroupVersionResource{
		Group:    traefikV1alpha1.SchemeGroupVersion.Group,
		Version:  traefikV1alpha1.SchemeGroupVersion.Version,
		Resource: "ingressroutes",
	}
	ingressrouteTCPGVR = schema.GroupVersionResource{
		Group:    traefikV1alpha1.SchemeGroupVersion.Group,
		Version:  traefikV1alpha1.SchemeGroupVersion.Version,
		Resource: "ingressroutetcps",
	}
	ingressrouteUDPGVR = schema.GroupVersionResource{
		Group:    traefikV1alpha1.SchemeGroupVersion.Group,
		Version:  traefikV1alpha1.SchemeGroupVersion.Version,
		Resource: "ingressrouteudps",
	}
)

type traefikSource struct {
	annotationFilter        string
	dynamicKubeClient       dynamic.Interface
	ingressRouteInformer    informers.GenericInformer
	ingressRouteTcpInformer informers.GenericInformer
	ingressRouteUdpInformer informers.GenericInformer
	kubeClient              kubernetes.Interface
	namespace               string
	unstructuredConverter   *unstructuredConverter
}

func NewTraefikSource(ctx context.Context, dynamicKubeClient dynamic.Interface, kubeClient kubernetes.Interface, namespace string, annotationFilter string) (Source, error) {
	// Use shared informer to listen for add/update/delete of Host in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	ingressRouteInformer := informerFactory.ForResource(ingressrouteGVR)
	ingressRouteTcpInformer := informerFactory.ForResource(ingressrouteTCPGVR)
	ingressRouteUdpInformer := informerFactory.ForResource(ingressrouteUDPGVR)

	// Add default resource event handlers to properly initialize informers.
	ingressRouteInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {},
		},
	)
	ingressRouteTcpInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {},
		},
	)
	ingressRouteUdpInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {},
		},
	)

	informerFactory.Start((ctx.Done()))

	// wait for the local cache to be populated.
	if err := waitForDynamicCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	uc, err := newTraefikUnstructuredConverter()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to setup Unstructured Converter")
	}

	return &traefikSource{
		annotationFilter:        annotationFilter,
		dynamicKubeClient:       dynamicKubeClient,
		ingressRouteInformer:    ingressRouteInformer,
		ingressRouteTcpInformer: ingressRouteTcpInformer,
		ingressRouteUdpInformer: ingressRouteUdpInformer,
		kubeClient:              kubeClient,
		namespace:               namespace,
		unstructuredConverter:   uc,
	}, nil
}

func (ts *traefikSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	ingressRouteEndpoints, err := ts.ingressRouteEndpoints()
	if err != nil {
		return nil, err
	}
	ingressRouteTCPEndpoints, err := ts.ingressRouteTCPEndpoints()
	if err != nil {
		return nil, err
	}
	ingressRouteUDPEndpoints, err := ts.ingressRouteUDPEndpoints()
	if err != nil {
		return nil, err
	}

	endpoints = append(endpoints, ingressRouteEndpoints...)
	endpoints = append(endpoints, ingressRouteTCPEndpoints...)
	endpoints = append(endpoints, ingressRouteUDPEndpoints...)

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

// ingressRouteEndpoints extracts endpoints from all IngressRoute objects
func (ts *traefikSource) ingressRouteEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.ingressRouteInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRoutes []*traefikV1alpha1.IngressRoute
	for _, ingressRouteObj := range irs {
		unstructuredHost, ok := ingressRouteObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		ingressRoute := &traefikV1alpha1.IngressRoute{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRoutes = append(ingressRoutes, ingressRoute)
	}

	ingressRoutes, err = ts.filterByAnnotationsIngressRoute(ingressRoutes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRoute")
	}

	for _, ingressRoute := range ingressRoutes {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRoute.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRoute.Namespace, ingressRoute.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRoute(ingressRoute, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRoute: %s: %v", fullname, ingressEndpoints)
		ts.setResourceLabelIngressRoute(ingressRoute, ingressEndpoints)
		ts.setDualstackLabelIngressRoute(ingressRoute, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// ingressRouteTCPEndpoints extracts endpoints from all IngressRouteTCP objects
func (ts *traefikSource) ingressRouteTCPEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.ingressRouteTcpInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRoutes []*traefikV1alpha1.IngressRouteTCP
	for _, ingressRouteObj := range irs {
		unstructuredHost, ok := ingressRouteObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		ingressRoute := &traefikV1alpha1.IngressRouteTCP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRoutes = append(ingressRoutes, ingressRoute)
	}

	ingressRoutes, err = ts.filterByAnnotationsIngressRouteTCP(ingressRoutes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRoute")
	}

	for _, ingressRoute := range ingressRoutes {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRoute.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRoute.Namespace, ingressRoute.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRouteTCP(ingressRoute, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRoute: %s: %v", fullname, ingressEndpoints)
		ts.setResourceLabelIngressRouteTCP(ingressRoute, ingressEndpoints)
		ts.setDualstackLabelIngressRouteTCP(ingressRoute, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// ingressRouteUDPEndpoints extracts endpoints from all IngressRouteUDP objects
func (ts *traefikSource) ingressRouteUDPEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.ingressRouteUdpInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRoutes []*traefikV1alpha1.IngressRouteUDP
	for _, ingressRouteObj := range irs {
		unstructuredHost, ok := ingressRouteObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		ingressRoute := &traefikV1alpha1.IngressRouteUDP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRoutes = append(ingressRoutes, ingressRoute)
	}

	ingressRoutes, err = ts.filterByAnnotationsIngressRouteUDP(ingressRoutes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRoute")
	}

	for _, ingressRoute := range ingressRoutes {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRoute.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRoute.Namespace, ingressRoute.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRouteUDP(ingressRoute, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRoute: %s: %v", fullname, ingressEndpoints)
		ts.setResourceLabelIngressRouteUDP(ingressRoute, ingressEndpoints)
		ts.setDualstackLabelIngressRouteUDP(ingressRoute, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// filterByAnnotations filters a list of IngressRoute by a given annotation selector.
func (ts *traefikSource) filterByAnnotationsIngressRoute(ingressRoutes []*traefikV1alpha1.IngressRoute) ([]*traefikV1alpha1.IngressRoute, error) {
	labelSelector, err := metav1.ParseToLabelSelector(ts.annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return ingressRoutes, nil
	}

	filteredList := []*traefikV1alpha1.IngressRoute{}

	for _, ingressRoute := range ingressRoutes {
		// convert the IngressRoute's annotations to an equivalent label selector
		annotations := labels.Set(ingressRoute.Annotations)

		// include IngressRoute if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, ingressRoute)
		}
	}

	return filteredList, nil
}

// filterByAnnotations filters a list of IngressRouteTCP by a given annotation selector.
func (ts *traefikSource) filterByAnnotationsIngressRouteTCP(ingressRoutes []*traefikV1alpha1.IngressRouteTCP) ([]*traefikV1alpha1.IngressRouteTCP, error) {
	labelSelector, err := metav1.ParseToLabelSelector(ts.annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return ingressRoutes, nil
	}

	filteredList := []*traefikV1alpha1.IngressRouteTCP{}

	for _, ingressRoute := range ingressRoutes {
		// convert the IngressRoute's annotations to an equivalent label selector
		annotations := labels.Set(ingressRoute.Annotations)

		// include IngressRoute if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, ingressRoute)
		}
	}

	return filteredList, nil
}

// filterByAnnotations filters a list of IngressRoute by a given annotation selector.
func (ts *traefikSource) filterByAnnotationsIngressRouteUDP(ingressRoutes []*traefikV1alpha1.IngressRouteUDP) ([]*traefikV1alpha1.IngressRouteUDP, error) {
	labelSelector, err := metav1.ParseToLabelSelector(ts.annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return ingressRoutes, nil
	}

	filteredList := []*traefikV1alpha1.IngressRouteUDP{}

	for _, ingressRoute := range ingressRoutes {
		// convert the IngressRoute's annotations to an equivalent label selector
		annotations := labels.Set(ingressRoute.Annotations)

		// include IngressRoute if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, ingressRoute)
		}
	}

	return filteredList, nil
}

func (ts *traefikSource) setResourceLabelIngressRoute(ingressroute *traefikV1alpha1.IngressRoute, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("ingressroute/%s/%s", ingressroute.Namespace, ingressroute.Name)
	}
}
func (ts *traefikSource) setResourceLabelIngressRouteTCP(ingressroute *traefikV1alpha1.IngressRouteTCP, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("ingressroutetcp/%s/%s", ingressroute.Namespace, ingressroute.Name)
	}
}
func (ts *traefikSource) setResourceLabelIngressRouteUDP(ingressroute *traefikV1alpha1.IngressRouteUDP, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("ingressrouteudp/%s/%s", ingressroute.Namespace, ingressroute.Name)
	}
}

func (ts *traefikSource) setDualstackLabelIngressRoute(ingressRoute *traefikV1alpha1.IngressRoute, endpoints []*endpoint.Endpoint) {
	val, ok := ingressRoute.Annotations[ALBDualstackAnnotationKey]
	if ok && val == ALBDualstackAnnotationValue {
		log.Debugf("Adding dualstack label to IngressRoute %s/%s.", ingressRoute.Namespace, ingressRoute.Name)
		for _, ep := range endpoints {
			ep.Labels[endpoint.DualstackLabelKey] = "true"
		}
	}
}
func (ts *traefikSource) setDualstackLabelIngressRouteTCP(ingressRoute *traefikV1alpha1.IngressRouteTCP, endpoints []*endpoint.Endpoint) {
	val, ok := ingressRoute.Annotations[ALBDualstackAnnotationKey]
	if ok && val == ALBDualstackAnnotationValue {
		log.Debugf("Adding dualstack label to IngressRouteTCP %s/%s.", ingressRoute.Namespace, ingressRoute.Name)
		for _, ep := range endpoints {
			ep.Labels[endpoint.DualstackLabelKey] = "true"
		}
	}
}
func (ts *traefikSource) setDualstackLabelIngressRouteUDP(ingressRoute *traefikV1alpha1.IngressRouteUDP, endpoints []*endpoint.Endpoint) {
	val, ok := ingressRoute.Annotations[ALBDualstackAnnotationKey]
	if ok && val == ALBDualstackAnnotationValue {
		log.Debugf("Adding dualstack label to IngressRouteUDP %s/%s.", ingressRoute.Namespace, ingressRoute.Name)
		for _, ep := range endpoints {
			ep.Labels[endpoint.DualstackLabelKey] = "true"
		}
	}
}

// endpointsFromIngressRoute extracts the endpoints from a IngressRoute object
func (ts *traefikSource) endpointsFromIngressRoute(ingressRoute *traefikV1alpha1.IngressRoute, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ingressRoute.Annotations)

	ttl, err := getTTLFromAnnotations(ingressRoute.Annotations)
	if err != nil {
		return nil, err
	}

	hostnameList := getHostnamesFromAnnotations(ingressRoute.Annotations)
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
	}

	// TODO: Implement Traefik router rule logic/regex magic
	// if ingressRoute.Spec.Rules != nil {
	// 	for _, rule := range ingressRoute.Spec.Rules {
	// 		if rule.Host != "" {
	// 			endpoints = append(endpoints, endpointsForHostname(rule.Host, targets, ttl, providerSpecific, setIdentifier)...)
	// 		}
	// 	}
	// }

	return endpoints, nil
}

// endpointsFromIngressRouteTCP extracts the endpoints from a IngressRouteTCP object
func (ts *traefikSource) endpointsFromIngressRouteTCP(ingressRoute *traefikV1alpha1.IngressRouteTCP, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ingressRoute.Annotations)

	ttl, err := getTTLFromAnnotations(ingressRoute.Annotations)
	if err != nil {
		return nil, err
	}

	hostnameList := getHostnamesFromAnnotations(ingressRoute.Annotations)
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
	}

	// TODO: Implement Traefik router rule logic/regex magic
	// if ingressRoute.Spec.Rules != nil {
	// 	for _, rule := range ingressRoute.Spec.Rules {
	// 		if rule.Host != "" {
	// 			endpoints = append(endpoints, endpointsForHostname(rule.Host, targets, ttl, providerSpecific, setIdentifier)...)
	// 		}
	// 	}
	// }

	return endpoints, nil
}

// endpointsFromIngressRouteUDP extracts the endpoints from a IngressRouteUDP object
func (ts *traefikSource) endpointsFromIngressRouteUDP(ingressRoute *traefikV1alpha1.IngressRouteUDP, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ingressRoute.Annotations)

	ttl, err := getTTLFromAnnotations(ingressRoute.Annotations)
	if err != nil {
		return nil, err
	}

	hostnameList := getHostnamesFromAnnotations(ingressRoute.Annotations)
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
	}

	// TODO: Implement Traefik router rule logic/regex magic
	// if ingressRoute.Spec.Rules != nil {
	// 	for _, rule := range ingressRoute.Spec.Rules {
	// 		if rule.Host != "" {
	// 			endpoints = append(endpoints, endpointsForHostname(rule.Host, targets, ttl, providerSpecific, setIdentifier)...)
	// 		}
	// 	}
	// }

	return endpoints, nil
}

func (ts *traefikSource) AddEventHandler(ctx context.Context, handler func()) {
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	log.Debug("Adding event handler for IngressRoute")
	ts.ingressRouteInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	log.Debug("Adding event handler for IngressRouteTCP")
	ts.ingressRouteTcpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	log.Debug("Adding event handler for IngressRouteUDP")
	ts.ingressRouteUdpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

// newTraefikUnstructuredConverter returns a new unstructuredConverter initialized
func newTraefikUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Add the core types we need
	uc.scheme.AddKnownTypes(ingressrouteGVR.GroupVersion(), &traefikV1alpha1.IngressRoute{}, &traefikV1alpha1.IngressRouteList{})
	uc.scheme.AddKnownTypes(ingressrouteTCPGVR.GroupVersion(), &traefikV1alpha1.IngressRouteTCP{}, &traefikV1alpha1.IngressRouteTCPList{})
	uc.scheme.AddKnownTypes(ingressrouteUDPGVR.GroupVersion(), &traefikV1alpha1.IngressRouteUDP{}, &traefikV1alpha1.IngressRouteUDP{})
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}
