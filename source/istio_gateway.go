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
	"strings"

	log "github.com/sirupsen/logrus"
	networkingv1 "istio.io/client-go/pkg/apis/networking/v1"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	istioinformers "istio.io/client-go/pkg/informers/externalversions"
	networkingv1informer "istio.io/client-go/pkg/informers/externalversions/networking/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	netinformers "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
)

// IstioGatewayIngressSource is the annotation used to determine if the gateway is implemented by an Ingress object
// instead of a standard LoadBalancer service type
// Using var instead of const because annotation keys can be customized
var IstioGatewayIngressSource = annotations.Ingress

// gatewaySource is an implementation of Source for Istio Gateway objects.
// The gateway implementation uses the spec.servers.hosts values for the hostnames.
// Use annotations.TargetKey to explicitly set Endpoint.
//
// +externaldns:source:name=istio-gateway
// +externaldns:source:category=Service Mesh
// +externaldns:source:description=Creates DNS entries from Istio Gateway resources
// +externaldns:source:resources=Gateway.networking.istio.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
type gatewaySource struct {
	namespace                string
	annotationFilter         string
	templateEngine           template.Engine
	ignoreHostnameAnnotation bool
	serviceInformer          coreinformers.ServiceInformer
	gatewayInformer          networkingv1informer.GatewayInformer
	ingressInformer          netinformers.IngressInformer
}

// NewIstioGatewaySource creates a new gatewaySource with the given config.
func NewIstioGatewaySource(
	ctx context.Context,
	kubeClient kubernetes.Interface,
	istioClient istioclient.Interface,
	cfg *Config,
) (Source, error) {
	// Use shared informers to listen for add/update/delete of services/pods/nodes in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0, kubeinformers.WithNamespace(cfg.Namespace))
	serviceInformer := informerFactory.Core().V1().Services()
	istioInformerFactory := istioinformers.NewSharedInformerFactoryWithOptions(istioClient, 0, istioinformers.WithNamespace(cfg.Namespace))
	gatewayInformer := istioInformerFactory.Networking().V1().Gateways()
	ingressInformer := informerFactory.Networking().V1().Ingresses()

	informers.MustSetTransform(serviceInformer.Informer(), informers.TransformerWithOptions[*corev1.Service](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
		informers.TransformRemoveStatusConditions(),
	))
	informers.MustSetTransform(ingressInformer.Informer(), informers.TransformerWithOptions[*networkv1.Ingress](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
	))
	informers.MustSetTransform(gatewayInformer.Informer(), informers.TransformerWithOptions[*networkingv1.Gateway](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
	))

	informers.MustAddIndexers(gatewayInformer.Informer(), informers.IndexerWithOptions[*networkingv1.Gateway](
		informers.IndexSelectorWithAnnotationFilter(cfg.AnnotationFilter),
		informers.IndexSelectorWithLabelSelector(cfg.LabelFilter),
	))

	// Add default resource event handlers to properly initialize informer.
	informers.MustAddEventHandler(serviceInformer.Informer(), informers.DefaultEventHandler())
	informers.MustAddEventHandler(ingressInformer.Informer(), informers.DefaultEventHandler())
	informers.MustAddEventHandler(gatewayInformer.Informer(), informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())
	istioInformerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}
	if err := informers.WaitForCacheSync(ctx, istioInformerFactory); err != nil {
		return nil, err
	}

	return &gatewaySource{
		namespace:                cfg.Namespace,
		annotationFilter:         cfg.AnnotationFilter,
		templateEngine:           cfg.TemplateEngine,
		ignoreHostnameAnnotation: cfg.IgnoreHostnameAnnotation,
		serviceInformer:          serviceInformer,
		gatewayInformer:          gatewayInformer,
		ingressInformer:          ingressInformer,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all gateway resources in the source's namespace(s).
func (sc *gatewaySource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	indexer := sc.gatewayInformer.Informer().GetIndexer()
	indexKeys := indexer.ListIndexFuncValues(informers.IndexWithSelectors)

	endpoints := make([]*endpoint.Endpoint, 0, len(indexKeys))

	log.Debugf("Found %d gateways in namespace %s", len(indexKeys), sc.namespace)

	for _, key := range indexKeys {
		gateway, err := informers.GetByKey[*networkingv1.Gateway](indexer, key)
		if err != nil || gateway == nil {
			continue
		}

		if annotations.IsControllerMismatch(gateway, types.IstioGateway) {
			continue
		}

		gwHostnames := sc.hostNamesFromGateway(gateway)

		log.Debugf("Processing gateway '%s/%s.%s' and hosts %q", gateway.Namespace, gateway.APIVersion, gateway.Name, strings.Join(gwHostnames, ","))

		gwEndpoints, err := sc.endpointsFromGateway(gwHostnames, gateway)
		if err != nil {
			return nil, err
		}

		// apply template if host is missing on gateway
		gwEndpoints, err = sc.templateEngine.CombineWithEndpoints(
			gwEndpoints,
			func() ([]*endpoint.Endpoint, error) {
				hostnames, err := sc.templateEngine.ExecFQDN(gateway)
				if err != nil {
					return nil, err
				}
				return sc.endpointsFromGateway(hostnames, gateway)
			},
		)
		if err != nil {
			return nil, err
		}

		if endpoint.HasNoEmptyEndpoints(gwEndpoints, types.IstioGateway, gateway) {
			continue
		}

		log.Debugf("Endpoints generated from '%s/%s/%s': %q", strings.ToLower(gateway.Kind), gateway.Namespace, gateway.Name, gwEndpoints)
		endpoints = append(endpoints, gwEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
}

// AddEventHandler adds an event handler that should be triggered if the watched Istio Gateway changes.
func (sc *gatewaySource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for Istio Gateway")

	informers.MustAddEventHandler(sc.gatewayInformer.Informer(), eventHandlerFunc(handler))
}

func (sc *gatewaySource) targetsFromIngress(ingressStr string, gateway *networkingv1.Gateway) (endpoint.Targets, error) {
	namespace, name, err := ParseIngress(ingressStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Ingress annotation on Gateway (%s/%s): %w", gateway.Namespace, gateway.Name, err)
	}
	if namespace == "" {
		namespace = gateway.Namespace
	}

	targets := make(endpoint.Targets, 0)

	ingress, err := sc.ingressInformer.Lister().Ingresses(namespace).Get(name)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, lb := range ingress.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			targets = append(targets, lb.IP)
		} else if lb.Hostname != "" {
			targets = append(targets, lb.Hostname)
		}
	}
	return targets, nil
}

func (sc *gatewaySource) targetsFromGateway(gateway *networkingv1.Gateway) (endpoint.Targets, error) {
	targets := annotations.TargetsFromTargetAnnotation(gateway.Annotations)
	if len(targets) > 0 {
		return targets, nil
	}

	ingressStr, ok := gateway.Annotations[IstioGatewayIngressSource]
	if ok && ingressStr != "" {
		return sc.targetsFromIngress(ingressStr, gateway)
	}

	return EndpointTargetsFromServices(sc.serviceInformer, sc.namespace, gateway.Spec.Selector)
}

// endpointsFromGatewayConfig extracts the endpoints from an Istio Gateway Config object
func (sc *gatewaySource) endpointsFromGateway(hostnames []string, gateway *networkingv1.Gateway) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	var err error

	targets, err := sc.targetsFromGateway(gateway)
	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return endpoints, nil
	}

	resource := fmt.Sprintf("gateway/%s/%s", gateway.Namespace, gateway.Name)
	ttl := annotations.TTLFromAnnotations(gateway.Annotations, resource)
	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(gateway.Annotations)

	for _, host := range hostnames {
		endpoints = append(endpoints, endpoint.EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}

	return endpoints, nil
}

func (sc *gatewaySource) hostNamesFromGateway(gateway *networkingv1.Gateway) []string {
	var hostnames []string
	for _, server := range gateway.Spec.Servers {
		for _, host := range server.Hosts {
			if host == "" {
				continue
			}

			parts := strings.Split(host, "/")

			// If the input hostname is of the form my-namespace/foo.bar.com, remove the namespace
			// before appending it to the list of endpoints to create
			if len(parts) == 2 {
				host = parts[1]
			}

			if host != "*" {
				hostnames = append(hostnames, host)
			}
		}
	}

	if !sc.ignoreHostnameAnnotation {
		hostnames = append(hostnames, annotations.HostnamesFromAnnotations(gateway.Annotations)...)
	}

	return hostnames
}
