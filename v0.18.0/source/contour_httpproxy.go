/*
Copyright 2020 The Kubernetes Authors.

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
	"errors"
	"fmt"
	"sort"
	"text/template"

	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/fqdn"
	"sigs.k8s.io/external-dns/source/informers"
)

// HTTPProxySource is an implementation of Source for ProjectContour HTTPProxy objects.
// The HTTPProxy implementation uses the spec.virtualHost.fqdn value for the hostname.
// Use targetAnnotationKey to explicitly set Endpoint.
type httpProxySource struct {
	dynamicKubeClient        dynamic.Interface
	namespace                string
	annotationFilter         string
	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
	httpProxyInformer        kubeinformers.GenericInformer
	unstructuredConverter    *UnstructuredConverter
}

// NewContourHTTPProxySource creates a new contourHTTPProxySource with the given config.
func NewContourHTTPProxySource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	namespace string,
	annotationFilter string,
	fqdnTemplate string,
	combineFqdnAnnotation bool,
	ignoreHostnameAnnotation bool,
) (Source, error) {
	tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	// Use shared informer to listen for add/update/delete of HTTPProxys in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	httpProxyInformer := informerFactory.ForResource(projectcontour.HTTPProxyGVR)

	// Add default resource event handlers to properly initialize informer.
	httpProxyInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	uc, err := NewUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Unstructured Converter: %w", err)
	}

	return &httpProxySource{
		dynamicKubeClient:        dynamicKubeClient,
		namespace:                namespace,
		annotationFilter:         annotationFilter,
		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    combineFqdnAnnotation,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
		httpProxyInformer:        httpProxyInformer,
		unstructuredConverter:    uc,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all HTTPProxy resources in the source's namespace(s).
func (sc *httpProxySource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	hps, err := sc.httpProxyInformer.Lister().ByNamespace(sc.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var httpProxies []*projectcontour.HTTPProxy
	for _, hp := range hps {
		unstructuredHP, ok := hp.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		hpConverted := &projectcontour.HTTPProxy{}
		err := sc.unstructuredConverter.scheme.Convert(unstructuredHP, hpConverted, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to HTTPProxy: %w", err)
		}
		httpProxies = append(httpProxies, hpConverted)
	}

	httpProxies, err = sc.filterByAnnotations(httpProxies)
	if err != nil {
		return nil, fmt.Errorf("failed to filter HTTPProxies: %w", err)
	}

	endpoints := []*endpoint.Endpoint{}

	for _, hp := range httpProxies {
		// Check controller annotation to see if we are responsible.
		controller, ok := hp.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping HTTPProxy %s/%s because controller value does not match, found: %s, required: %s",
				hp.Namespace, hp.Name, controller, controllerAnnotationValue)
			continue
		}

		hpEndpoints, err := sc.endpointsFromHTTPProxy(hp)
		if err != nil {
			return nil, fmt.Errorf("failed to get endpoints from HTTPProxy: %w", err)
		}

		// apply template if fqdn is missing on HTTPProxy
		if (sc.combineFQDNAnnotation || len(hpEndpoints) == 0) && sc.fqdnTemplate != nil {
			tmplEndpoints, err := sc.endpointsFromTemplate(hp)
			if err != nil {
				return nil, fmt.Errorf("failed to get endpoints from template: %w", err)
			}

			if sc.combineFQDNAnnotation {
				hpEndpoints = append(hpEndpoints, tmplEndpoints...)
			} else {
				hpEndpoints = tmplEndpoints
			}
		}

		if len(hpEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from HTTPProxy %s/%s", hp.Namespace, hp.Name)
			continue
		}

		log.Debugf("Endpoints generated from HTTPProxy: %s/%s: %v", hp.Namespace, hp.Name, hpEndpoints)
		endpoints = append(endpoints, hpEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (sc *httpProxySource) endpointsFromTemplate(httpProxy *projectcontour.HTTPProxy) ([]*endpoint.Endpoint, error) {
	hostnames, err := fqdn.ExecTemplate(sc.fqdnTemplate, httpProxy)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf("HTTPProxy/%s/%s", httpProxy.Namespace, httpProxy.Name)

	ttl := annotations.TTLFromAnnotations(httpProxy.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(httpProxy.Annotations)
	if len(targets) == 0 {
		for _, lb := range httpProxy.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				targets = append(targets, lb.IP)
			}
			if lb.Hostname != "" {
				targets = append(targets, lb.Hostname)
			}
		}
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(httpProxy.Annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}
	return endpoints, nil
}

// filterByAnnotations filters a list of configs by a given annotation selector.
func (sc *httpProxySource) filterByAnnotations(httpProxies []*projectcontour.HTTPProxy) ([]*projectcontour.HTTPProxy, error) {
	labelSelector, err := metav1.ParseToLabelSelector(sc.annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return httpProxies, nil
	}

	var filteredList []*projectcontour.HTTPProxy

	for _, httpProxy := range httpProxies {
		// include HTTPProxy if its annotations match the selector
		if selector.Matches(labels.Set(httpProxy.Annotations)) {
			filteredList = append(filteredList, httpProxy)
		}
	}

	return filteredList, nil
}

// endpointsFromHTTPProxyConfig extracts the endpoints from a Contour HTTPProxy object
func (sc *httpProxySource) endpointsFromHTTPProxy(httpProxy *projectcontour.HTTPProxy) ([]*endpoint.Endpoint, error) {
	resource := fmt.Sprintf("HTTPProxy/%s/%s", httpProxy.Namespace, httpProxy.Name)

	ttl := annotations.TTLFromAnnotations(httpProxy.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(httpProxy.Annotations)

	if len(targets) == 0 {
		for _, lb := range httpProxy.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				targets = append(targets, lb.IP)
			}
			if lb.Hostname != "" {
				targets = append(targets, lb.Hostname)
			}
		}
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(httpProxy.Annotations)

	var endpoints []*endpoint.Endpoint

	if virtualHost := httpProxy.Spec.VirtualHost; virtualHost != nil {
		if fqdn := virtualHost.Fqdn; fqdn != "" {
			endpoints = append(endpoints, endpointsForHostname(fqdn, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(httpProxy.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return endpoints, nil
}

func (sc *httpProxySource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("Adding event handler for httpproxy")

	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	sc.httpProxyInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}
