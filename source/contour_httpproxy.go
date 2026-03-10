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
	"text/template"

	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/fqdn"
	"sigs.k8s.io/external-dns/source/informers"
)

// HTTPProxySource is an implementation of Source for ProjectContour HTTPProxy objects.
// The HTTPProxy implementation uses the spec.virtualHost.fqdn value for the hostname.
// Use annotations.TargetKey to explicitly set Endpoint.
//
// +externaldns:source:name=contour-httpproxy
// +externaldns:source:category=Ingress Controllers
// +externaldns:source:description=Creates DNS entries from Contour HTTPProxy resources
// +externaldns:source:resources=HTTPProxy.projectcontour.io
// +externaldns:source:filters=annotation
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
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
	_, _ = httpProxyInformer.Informer().AddEventHandler(informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
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
func (sc *httpProxySource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
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

	httpProxies, err = annotations.Filter(httpProxies, sc.annotationFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to filter HTTPProxies: %w", err)
	}

	endpoints := []*endpoint.Endpoint{}

	for _, hp := range httpProxies {
		if annotations.IsControllerMismatch(hp, types.ContourHTTPProxy) {
			continue
		}

		hpEndpoints := sc.endpointsFromHTTPProxy(hp)

		// apply template if fqdn is missing on HTTPProxy
		hpEndpoints, err = fqdn.CombineWithTemplatedEndpoints(
			hpEndpoints,
			sc.fqdnTemplate,
			sc.combineFQDNAnnotation,
			func() ([]*endpoint.Endpoint, error) { return sc.endpointsFromTemplate(hp) },
		)
		if err != nil {
			return nil, err
		}

		if endpoint.HasNoEmptyEndpoints(hpEndpoints, types.ContourHTTPProxy, hp) {
			continue
		}

		log.Debugf("Endpoints generated from HTTPProxy: %s/%s: %v", hp.Namespace, hp.Name, hpEndpoints)
		endpoints = append(endpoints, hpEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
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
		endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}
	return endpoints, nil
}

// endpointsFromHTTPProxyConfig extracts the endpoints from a Contour HTTPProxy object
func (sc *httpProxySource) endpointsFromHTTPProxy(httpProxy *projectcontour.HTTPProxy) []*endpoint.Endpoint {
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
			endpoints = append(endpoints, EndpointsForHostname(fqdn, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(httpProxy.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return endpoints
}

func (sc *httpProxySource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for httpproxy")

	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	_, _ = sc.httpProxyInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}
