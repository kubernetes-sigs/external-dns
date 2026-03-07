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
	"strings"

	ambassador "github.com/datawire/ambassador/pkg/api/getambassador.io/v2"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
)

const (
	// ambHostAnnotation is the annotation in the Host that maps to a Service
	ambHostAnnotation = "external-dns.ambassador-service"
	// groupName is the group name for the Ambassador API
	groupName = "getambassador.io"
)

var (
	schemeGroupVersion = schema.GroupVersion{Group: groupName, Version: "v2"}
	ambHostGVR         = schemeGroupVersion.WithResource("hosts")
)

// ambassadorHostSource is an implementation of Source for Ambassador Host objects.
// The IngressRoute implementation uses the spec.virtualHost.fqdn value for the hostname.
// Use annotations.TargetKey to explicitly set Endpoint.
//
// +externaldns:source:name=ambassador-host
// +externaldns:source:category=Ingress Controllers
// +externaldns:source:description=Creates DNS entries from Ambassador Host resources
// +externaldns:source:resources=Host.getambassador.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=false
// +externaldns:source:events=false
// +externaldns:source:provider-specific=true
type ambassadorHostSource struct {
	dynamicKubeClient      dynamic.Interface
	kubeClient             kubernetes.Interface
	namespace              string
	annotationFilter       string
	ambassadorHostInformer kubeinformers.GenericInformer
	unstructuredConverter  *unstructuredConverter
	labelSelector          labels.Selector
}

// NewAmbassadorHostSource creates a new ambassadorHostSource with the given config.
func NewAmbassadorHostSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	namespace string,
	annotationFilter string,
	labelSelector labels.Selector,
) (Source, error) {
	// Use shared informer to listen for add/update/delete of Host in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	ambassadorHostInformer := informerFactory.ForResource(ambHostGVR)

	// Add default resource event handlers to properly initialize informer.
	_, _ = ambassadorHostInformer.Informer().AddEventHandler(informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	uc, err := newUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Unstructured Converter: %w", err)
	}

	return &ambassadorHostSource{
		dynamicKubeClient:      dynamicKubeClient,
		kubeClient:             kubeClient,
		namespace:              namespace,
		annotationFilter:       annotationFilter,
		ambassadorHostInformer: ambassadorHostInformer,
		unstructuredConverter:  uc,
		labelSelector:          labelSelector,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all Hosts in the source's namespace(s).
func (sc *ambassadorHostSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	hosts, err := sc.ambassadorHostInformer.Lister().ByNamespace(sc.namespace).List(sc.labelSelector)
	if err != nil {
		return nil, err
	}

	// Get a list of Ambassador Host resources
	var ambassadorHosts []*ambassador.Host
	for _, hostObj := range hosts {
		unstructuredHost, ok := hostObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		host := &ambassador.Host{}
		err := sc.unstructuredConverter.scheme.Convert(unstructuredHost, host, nil)
		if err != nil {
			return nil, err
		}
		ambassadorHosts = append(ambassadorHosts, host)
	}

	// Filter Ambassador Hosts
	ambassadorHosts, err = annotations.Filter(ambassadorHosts, sc.annotationFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to filter Ambassador Hosts by annotation: %w", err)
	}

	var endpoints []*endpoint.Endpoint

	for _, host := range ambassadorHosts {
		fullname := fmt.Sprintf("%s/%s", host.Namespace, host.Name)

		// look for the "external-dns.ambassador-service" annotation. If it is not there then just ignore this `Host`
		service, found := host.Annotations[ambHostAnnotation]
		if !found {
			log.Debugf("Host %s ignored: no annotation %q found", fullname, ambHostAnnotation)
			continue
		}

		targets := annotations.TargetsFromTargetAnnotation(host.Annotations)
		if len(targets) == 0 {
			targets, err = sc.targetsFromAmbassadorLoadBalancer(ctx, service)
			if err != nil {
				log.Warningf("Could not find targets for service %s for Host %s: %v", service, fullname, err)
				continue
			}
		}

		hostEndpoints := sc.endpointsFromHost(host, targets)
		if endpoint.HasNoEmptyEndpoints(hostEndpoints, types.AmbassadorHost, host) {
			continue
		}

		log.Debugf("Endpoints generated from Host: %s: %v", fullname, hostEndpoints)
		endpoints = append(endpoints, hostEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
}

// endpointsFromHost extracts the endpoints from a Host object
func (sc *ambassadorHostSource) endpointsFromHost(host *ambassador.Host, targets endpoint.Targets) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("host/%s/%s", host.Namespace, host.Name)
	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(host.Annotations)
	ttl := annotations.TTLFromAnnotations(host.Annotations, resource)

	if host.Spec != nil {
		hostname := host.Spec.Hostname
		if hostname != "" {
			endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return endpoints
}

func (sc *ambassadorHostSource) targetsFromAmbassadorLoadBalancer(ctx context.Context, service string) (endpoint.Targets, error) {
	lbNamespace, lbName, err := parseAmbLoadBalancerService(service)
	if err != nil {
		return nil, err
	}

	svc, err := sc.kubeClient.CoreV1().Services(lbNamespace).Get(ctx, lbName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	targets := extractLoadBalancerTargets(svc, false)

	return targets, nil
}

// parseAmbLoadBalancerService returns a name/namespace tuple from the annotation in
// an Ambassador Host CRD
//
// This is a thing because Ambassador has historically supported cross-namespace
// references using a name.namespace syntax, but here we want to also support
// namespace/name.
//
// Returns namespace, name, error.

func parseAmbLoadBalancerService(service string) (string, string, error) {
	// Start by assuming that we have namespace/name.
	parts := strings.Split(service, "/")

	if len(parts) == 1 {
		// No "/" at all, so let's try for name.namespace. To be consistent with the
		// rest of Ambassador, use SplitN to limit this to one split, so that e.g.
		// svc.foo.bar uses service "svc" in namespace "foo.bar".
		parts = strings.SplitN(service, ".", 2)

		if len(parts) == 2 {
			// We got a namespace, great.
			name := parts[0]
			namespace := parts[1]

			return namespace, name, nil
		}

		// If here, we have no separator, so the whole string is the service, and
		// we can assume the default namespace.
		name := service
		namespace := "default"

		return namespace, name, nil
	} else if len(parts) == 2 {
		// This is "namespace/name". Note that the name could be qualified,
		// which is fine.
		namespace := parts[0]
		name := parts[1]

		return namespace, name, nil
	}

	// If we got here, this string is simply ill-formatted. Return an error.
	return "", "", fmt.Errorf("invalid external-dns service: %s", service)
}

func (sc *ambassadorHostSource) AddEventHandler(_ context.Context, _ func()) {
}

// unstructuredConverter handles conversions between unstructured.Unstructured and Ambassador types
type unstructuredConverter struct {
	// scheme holds an initializer for converting Unstructured to a type
	scheme *runtime.Scheme
}

// newUnstructuredConverter returns a new unstructuredConverter initialized
func newUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Setup converter to understand custom CRD types
	ambassador.AddToScheme(uc.scheme)

	// Add the core types we need
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}
