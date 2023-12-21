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
	"regexp"
	"sort"
	"strings"

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
)

var (
	ingressrouteGVR = schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}
	ingressrouteTCPGVR = schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressroutetcps",
	}
	ingressrouteUDPGVR = schema.GroupVersionResource{
		Group:    "traefik.io",
		Version:  "v1alpha1",
		Resource: "ingressrouteudps",
	}
	oldIngressrouteGVR = schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}
	oldIngressrouteTCPGVR = schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressroutetcps",
	}
	oldIngressrouteUDPGVR = schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressrouteudps",
	}
)

var (
	traefikHostExtractor  = regexp.MustCompile(`(?:HostSNI|HostHeader|Host)\s*\(\s*(\x60.*?\x60)\s*\)`)
	traefikValueProcessor = regexp.MustCompile(`\x60([^,\x60]+)\x60`)
)

type traefikSource struct {
	annotationFilter           string
	ignoreHostnameAnnotation   bool
	dynamicKubeClient          dynamic.Interface
	ingressRouteInformer       informers.GenericInformer
	ingressRouteTcpInformer    informers.GenericInformer
	ingressRouteUdpInformer    informers.GenericInformer
	oldIngressRouteInformer    informers.GenericInformer
	oldIngressRouteTcpInformer informers.GenericInformer
	oldIngressRouteUdpInformer informers.GenericInformer
	kubeClient                 kubernetes.Interface
	namespace                  string
	unstructuredConverter      *unstructuredConverter
}

func NewTraefikSource(ctx context.Context, dynamicKubeClient dynamic.Interface, kubeClient kubernetes.Interface, namespace string, annotationFilter string, ignoreHostnameAnnotation bool, disableLegacy bool, disableNew bool) (Source, error) {
	// Use shared informer to listen for add/update/delete of Host in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	var ingressRouteInformer, ingressRouteTcpInformer, ingressRouteUdpInformer informers.GenericInformer
	var oldIngressRouteInformer, oldIngressRouteTcpInformer, oldIngressRouteUdpInformer informers.GenericInformer

	// Add default resource event handlers to properly initialize informers.
	if !disableNew {
		ingressRouteInformer = informerFactory.ForResource(ingressrouteGVR)
		ingressRouteTcpInformer = informerFactory.ForResource(ingressrouteTCPGVR)
		ingressRouteUdpInformer = informerFactory.ForResource(ingressrouteUDPGVR)
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
	}
	if !disableLegacy {
		oldIngressRouteInformer = informerFactory.ForResource(oldIngressrouteGVR)
		oldIngressRouteTcpInformer = informerFactory.ForResource(oldIngressrouteTCPGVR)
		oldIngressRouteUdpInformer = informerFactory.ForResource(oldIngressrouteUDPGVR)
		oldIngressRouteInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
		oldIngressRouteTcpInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
		oldIngressRouteUdpInformer.Informer().AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {},
			},
		)
	}

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
		annotationFilter:           annotationFilter,
		ignoreHostnameAnnotation:   ignoreHostnameAnnotation,
		dynamicKubeClient:          dynamicKubeClient,
		ingressRouteInformer:       ingressRouteInformer,
		ingressRouteTcpInformer:    ingressRouteTcpInformer,
		ingressRouteUdpInformer:    ingressRouteUdpInformer,
		oldIngressRouteInformer:    oldIngressRouteInformer,
		oldIngressRouteTcpInformer: oldIngressRouteTcpInformer,
		oldIngressRouteUdpInformer: oldIngressRouteUdpInformer,
		kubeClient:                 kubeClient,
		namespace:                  namespace,
		unstructuredConverter:      uc,
	}, nil
}

func (ts *traefikSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	if ts.ingressRouteInformer != nil {
		ingressRouteEndpoints, err := ts.ingressRouteEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteEndpoints...)
	}
	if ts.oldIngressRouteInformer != nil {
		oldIngressRouteEndpoints, err := ts.oldIngressRouteEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteEndpoints...)
	}
	if ts.ingressRouteTcpInformer != nil {
		ingressRouteTcpEndpoints, err := ts.ingressRouteTCPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteTcpEndpoints...)
	}
	if ts.oldIngressRouteTcpInformer != nil {
		oldIngressRouteTcpEndpoints, err := ts.oldIngressRouteTCPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteTcpEndpoints...)
	}
	if ts.ingressRouteUdpInformer != nil {
		ingressRouteUdpEndpoints, err := ts.ingressRouteUDPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, ingressRouteUdpEndpoints...)
	}
	if ts.oldIngressRouteUdpInformer != nil {
		oldIngressRouteUdpEndpoints, err := ts.oldIngressRouteUDPEndpoints()
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, oldIngressRouteUdpEndpoints...)
	}

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

	var ingressRoutes []*IngressRoute
	for _, ingressRouteObj := range irs {
		unstructuredHost, ok := ingressRouteObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRoute object to unstructured")
		}

		ingressRoute := &IngressRoute{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRoutes = append(ingressRoutes, ingressRoute)
	}

	ingressRoutes, err = ts.filterIngressRouteByAnnotation(ingressRoutes)
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

	var ingressRouteTCPs []*IngressRouteTCP
	for _, ingressRouteTCPObj := range irs {
		unstructuredHost, ok := ingressRouteTCPObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRouteTCP object to unstructured")
		}

		ingressRouteTCP := &IngressRouteTCP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRouteTCP, nil)
		if err != nil {
			return nil, err
		}
		ingressRouteTCPs = append(ingressRouteTCPs, ingressRouteTCP)
	}

	ingressRouteTCPs, err = ts.filterIngressRouteTcpByAnnotations(ingressRouteTCPs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRouteTCP")
	}

	for _, ingressRouteTCP := range ingressRouteTCPs {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRouteTCP.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRouteTCP.Namespace, ingressRouteTCP.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRouteTCP(ingressRouteTCP, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRouteTCP: %s: %v", fullname, ingressEndpoints)
		ts.setDualstackLabelIngressRouteTCP(ingressRouteTCP, ingressEndpoints)
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

	var ingressRouteUDPs []*IngressRouteUDP
	for _, ingressRouteUDPObj := range irs {
		unstructuredHost, ok := ingressRouteUDPObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRouteUDP object to unstructured")
		}

		ingressRoute := &IngressRouteUDP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRouteUDPs = append(ingressRouteUDPs, ingressRoute)
	}

	ingressRouteUDPs, err = ts.filterIngressRouteUdpByAnnotations(ingressRouteUDPs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRouteUDP")
	}

	for _, ingressRouteUDP := range ingressRouteUDPs {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRouteUDP.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRouteUDP.Namespace, ingressRouteUDP.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRouteUDP(ingressRouteUDP, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRouteUDP: %s: %v", fullname, ingressEndpoints)
		ts.setDualstackLabelIngressRouteUDP(ingressRouteUDP, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// oldIngressRouteEndpoints extracts endpoints from all IngressRoute objects
func (ts *traefikSource) oldIngressRouteEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.oldIngressRouteInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRoutes []*IngressRoute
	for _, ingressRouteObj := range irs {
		unstructuredHost, ok := ingressRouteObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRoute object to unstructured")
		}

		ingressRoute := &IngressRoute{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRoutes = append(ingressRoutes, ingressRoute)
	}

	ingressRoutes, err = ts.filterIngressRouteByAnnotation(ingressRoutes)
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
		ts.setDualstackLabelIngressRoute(ingressRoute, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// oldIngressRouteTCPEndpoints extracts endpoints from all IngressRouteTCP objects
func (ts *traefikSource) oldIngressRouteTCPEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.oldIngressRouteTcpInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRouteTCPs []*IngressRouteTCP
	for _, ingressRouteTCPObj := range irs {
		unstructuredHost, ok := ingressRouteTCPObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRouteTCP object to unstructured")
		}

		ingressRouteTCP := &IngressRouteTCP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRouteTCP, nil)
		if err != nil {
			return nil, err
		}
		ingressRouteTCPs = append(ingressRouteTCPs, ingressRouteTCP)
	}

	ingressRouteTCPs, err = ts.filterIngressRouteTcpByAnnotations(ingressRouteTCPs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRouteTCP")
	}

	for _, ingressRouteTCP := range ingressRouteTCPs {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRouteTCP.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRouteTCP.Namespace, ingressRouteTCP.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRouteTCP(ingressRouteTCP, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRouteTCP: %s: %v", fullname, ingressEndpoints)
		ts.setDualstackLabelIngressRouteTCP(ingressRouteTCP, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// oldIngressRouteUDPEndpoints extracts endpoints from all IngressRouteUDP objects
func (ts *traefikSource) oldIngressRouteUDPEndpoints() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	irs, err := ts.oldIngressRouteUdpInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var ingressRouteUDPs []*IngressRouteUDP
	for _, ingressRouteUDPObj := range irs {
		unstructuredHost, ok := ingressRouteUDPObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert IngressRouteUDP object to unstructured")
		}

		ingressRoute := &IngressRouteUDP{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, ingressRoute, nil)
		if err != nil {
			return nil, err
		}
		ingressRouteUDPs = append(ingressRouteUDPs, ingressRoute)
	}

	ingressRouteUDPs, err = ts.filterIngressRouteUdpByAnnotations(ingressRouteUDPs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter IngressRouteUDP")
	}

	for _, ingressRouteUDP := range ingressRouteUDPs {
		var targets endpoint.Targets

		targets = append(targets, getTargetsFromTargetAnnotation(ingressRouteUDP.Annotations)...)

		fullname := fmt.Sprintf("%s/%s", ingressRouteUDP.Namespace, ingressRouteUDP.Name)

		ingressEndpoints, err := ts.endpointsFromIngressRouteUDP(ingressRouteUDP, targets)
		if err != nil {
			return nil, err
		}
		if len(ingressEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from Host %s", fullname)
			continue
		}

		log.Debugf("Endpoints generated from IngressRouteUDP: %s: %v", fullname, ingressEndpoints)
		ts.setDualstackLabelIngressRouteUDP(ingressRouteUDP, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return endpoints, nil
}

// filterIngressRouteByAnnotation filters a list of IngressRoute by a given annotation selector.
func (ts *traefikSource) filterIngressRouteByAnnotation(ingressRoutes []*IngressRoute) ([]*IngressRoute, error) {
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

	filteredList := []*IngressRoute{}

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

// filterIngressRouteTcpByAnnotations filters a list of IngressRouteTCP by a given annotation selector.
func (ts *traefikSource) filterIngressRouteTcpByAnnotations(ingressRoutes []*IngressRouteTCP) ([]*IngressRouteTCP, error) {
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

	filteredList := []*IngressRouteTCP{}

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

// filterIngressRouteUdpByAnnotations filters a list of IngressRoute by a given annotation selector.
func (ts *traefikSource) filterIngressRouteUdpByAnnotations(ingressRoutes []*IngressRouteUDP) ([]*IngressRouteUDP, error) {
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

	filteredList := []*IngressRouteUDP{}

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

func (ts *traefikSource) setDualstackLabelIngressRoute(ingressRoute *IngressRoute, endpoints []*endpoint.Endpoint) {
	val, ok := ingressRoute.Annotations[ALBDualstackAnnotationKey]
	if ok && val == ALBDualstackAnnotationValue {
		log.Debugf("Adding dualstack label to IngressRoute %s/%s.", ingressRoute.Namespace, ingressRoute.Name)
		for _, ep := range endpoints {
			ep.Labels[endpoint.DualstackLabelKey] = "true"
		}
	}
}
func (ts *traefikSource) setDualstackLabelIngressRouteTCP(ingressRoute *IngressRouteTCP, endpoints []*endpoint.Endpoint) {
	val, ok := ingressRoute.Annotations[ALBDualstackAnnotationKey]
	if ok && val == ALBDualstackAnnotationValue {
		log.Debugf("Adding dualstack label to IngressRouteTCP %s/%s.", ingressRoute.Namespace, ingressRoute.Name)
		for _, ep := range endpoints {
			ep.Labels[endpoint.DualstackLabelKey] = "true"
		}
	}
}
func (ts *traefikSource) setDualstackLabelIngressRouteUDP(ingressRoute *IngressRouteUDP, endpoints []*endpoint.Endpoint) {
	val, ok := ingressRoute.Annotations[ALBDualstackAnnotationKey]
	if ok && val == ALBDualstackAnnotationValue {
		log.Debugf("Adding dualstack label to IngressRouteUDP %s/%s.", ingressRoute.Namespace, ingressRoute.Name)
		for _, ep := range endpoints {
			ep.Labels[endpoint.DualstackLabelKey] = "true"
		}
	}
}

// endpointsFromIngressRoute extracts the endpoints from a IngressRoute object
func (ts *traefikSource) endpointsFromIngressRoute(ingressRoute *IngressRoute, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressroute/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := getTTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	for _, route := range ingressRoute.Spec.Routes {
		match := route.Match

		for _, hostEntry := range traefikHostExtractor.FindAllString(match, -1) {
			for _, host := range traefikValueProcessor.FindAllString(hostEntry, -1) {
				host = strings.TrimPrefix(host, "`")
				host = strings.TrimSuffix(host, "`")

				// Checking for host = * is required, as Host(`*`) can be set
				if host != "*" && host != "" {
					endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
				}
			}
		}
	}

	return endpoints, nil
}

// endpointsFromIngressRouteTCP extracts the endpoints from a IngressRouteTCP object
func (ts *traefikSource) endpointsFromIngressRouteTCP(ingressRoute *IngressRouteTCP, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressroutetcp/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := getTTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	for _, route := range ingressRoute.Spec.Routes {
		match := route.Match

		for _, hostEntry := range traefikHostExtractor.FindAllString(match, -1) {
			for _, host := range traefikValueProcessor.FindAllString(hostEntry, -1) {
				host = strings.TrimPrefix(host, "`")
				host = strings.TrimSuffix(host, "`")

				// Checking for host = * is required, as HostSNI(`*`) can be set
				// in the case of TLS passthrough
				if host != "*" && host != "" {
					endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
				}
			}
		}
	}

	return endpoints, nil
}

// endpointsFromIngressRouteUDP extracts the endpoints from a IngressRouteUDP object
func (ts *traefikSource) endpointsFromIngressRouteUDP(ingressRoute *IngressRouteUDP, targets endpoint.Targets) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("ingressrouteudp/%s/%s", ingressRoute.Namespace, ingressRoute.Name)

	ttl := getTTLFromAnnotations(ingressRoute.Annotations, resource)

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(ingressRoute.Annotations)

	if !ts.ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(ingressRoute.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return endpoints, nil
}

func (ts *traefikSource) AddEventHandler(ctx context.Context, handler func()) {
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	log.Debug("Adding event handler for IngressRoute")
	ts.ingressRouteInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	ts.oldIngressRouteInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	log.Debug("Adding event handler for IngressRouteTCP")
	ts.ingressRouteTcpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	ts.oldIngressRouteTcpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	log.Debug("Adding event handler for IngressRouteUDP")
	ts.ingressRouteUdpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	ts.oldIngressRouteUdpInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

// newTraefikUnstructuredConverter returns a new unstructuredConverter initialized
func newTraefikUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Add the core types we need
	uc.scheme.AddKnownTypes(ingressrouteGVR.GroupVersion(), &IngressRoute{}, &IngressRouteList{})
	uc.scheme.AddKnownTypes(oldIngressrouteGVR.GroupVersion(), &IngressRoute{}, &IngressRouteList{})
	uc.scheme.AddKnownTypes(ingressrouteTCPGVR.GroupVersion(), &IngressRouteTCP{}, &IngressRouteTCPList{})
	uc.scheme.AddKnownTypes(oldIngressrouteTCPGVR.GroupVersion(), &IngressRouteTCP{}, &IngressRouteTCPList{})
	uc.scheme.AddKnownTypes(ingressrouteUDPGVR.GroupVersion(), &IngressRouteUDP{}, &IngressRouteUDPList{})
	uc.scheme.AddKnownTypes(oldIngressrouteUDPGVR.GroupVersion(), &IngressRouteUDP{}, &IngressRouteUDPList{})
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}

// Basic redefinition of Traefik 2's CRD: https://github.com/traefik/traefik/tree/v2.8.7/pkg/provider/kubernetes/crd/traefik/v1alpha1

// traefikIngressRouteSpec defines the desired state of IngressRoute.
type traefikIngressRouteSpec struct {
	// Routes defines the list of routes.
	Routes []traefikRoute `json:"routes"`
}

// traefikRoute holds the HTTP route configuration.
type traefikRoute struct {
	// Match defines the router's rule.
	// More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#rule
	Match string `json:"match"`
}

// IngressRoute is the CRD implementation of a Traefik HTTP Router.
type IngressRoute struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	Spec traefikIngressRouteSpec `json:"spec"`
}

// IngressRouteList is a collection of IngressRoute.
type IngressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of IngressRoute.
	Items []IngressRoute `json:"items"`
}

// traefikIngressRouteTCPSpec defines the desired state of IngressRouteTCP.
type traefikIngressRouteTCPSpec struct {
	Routes []traefikRouteTCP `json:"routes"`
}

// traefikRouteTCP holds the TCP route configuration.
type traefikRouteTCP struct {
	// Match defines the router's rule.
	// More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#rule_1
	Match string `json:"match"`
}

// IngressRouteTCP is the CRD implementation of a Traefik TCP Router.
type IngressRouteTCP struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`

	Spec traefikIngressRouteTCPSpec `json:"spec"`
}

// IngressRouteTCPList is a collection of IngressRouteTCP.
type IngressRouteTCPList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of IngressRouteTCP.
	Items []IngressRouteTCP `json:"items"`
}

// IngressRouteUDP is a CRD implementation of a Traefik UDP Router.
type IngressRouteUDP struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata"`
}

// IngressRouteUDPList is a collection of IngressRouteUDP.
type IngressRouteUDPList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata"`

	// Items is the list of IngressRouteUDP.
	Items []IngressRouteUDP `json:"items"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRoute) DeepCopyInto(out *IngressRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRoute.
func (in *IngressRoute) DeepCopy() *IngressRoute {
	if in == nil {
		return nil
	}
	out := new(IngressRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteList) DeepCopyInto(out *IngressRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteList.
func (in *IngressRouteList) DeepCopy() *IngressRouteList {
	if in == nil {
		return nil
	}
	out := new(IngressRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikIngressRouteSpec) DeepCopyInto(out *traefikIngressRouteSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]traefikRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteSpec.
func (in *traefikIngressRouteSpec) DeepCopy() *traefikIngressRouteSpec {
	if in == nil {
		return nil
	}
	out := new(traefikIngressRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikRoute) DeepCopyInto(out *traefikRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *traefikRoute) DeepCopy() *traefikRoute {
	if in == nil {
		return nil
	}
	out := new(traefikRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteTCP) DeepCopyInto(out *IngressRouteTCP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteTCP.
func (in *IngressRouteTCP) DeepCopy() *IngressRouteTCP {
	if in == nil {
		return nil
	}
	out := new(IngressRouteTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteTCP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteTCPList) DeepCopyInto(out *IngressRouteTCPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressRouteTCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteTCPList.
func (in *IngressRouteTCPList) DeepCopy() *IngressRouteTCPList {
	if in == nil {
		return nil
	}
	out := new(IngressRouteTCPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteTCPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikIngressRouteTCPSpec) DeepCopyInto(out *traefikIngressRouteTCPSpec) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]traefikRouteTCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteTCPSpec.
func (in *traefikIngressRouteTCPSpec) DeepCopy() *traefikIngressRouteTCPSpec {
	if in == nil {
		return nil
	}
	out := new(traefikIngressRouteTCPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *traefikRouteTCP) DeepCopyInto(out *traefikRouteTCP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTCP.
func (in *traefikRouteTCP) DeepCopy() *traefikRouteTCP {
	if in == nil {
		return nil
	}
	out := new(traefikRouteTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteUDP) DeepCopyInto(out *IngressRouteUDP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteUDP.
func (in *IngressRouteUDP) DeepCopy() *IngressRouteUDP {
	if in == nil {
		return nil
	}
	out := new(IngressRouteUDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteUDP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRouteUDPList) DeepCopyInto(out *IngressRouteUDPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressRouteUDP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRouteUDPList.
func (in *IngressRouteUDPList) DeepCopy() *IngressRouteUDPList {
	if in == nil {
		return nil
	}
	out := new(IngressRouteUDPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressRouteUDPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
