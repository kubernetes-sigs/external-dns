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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	netinformers "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayclient "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	gwinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	gwinformers_v1 "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/apis/v1"

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
)

// IstioGatewayIngressSource returns the annotation key used to determine if the gateway
// is implemented by an Ingress object instead of a standard LoadBalancer service type.
// This must be a function (not a package-level var) because the annotation prefix can
// be customized at runtime via --annotation-prefix / SetAnnotationPrefix.
func IstioGatewayIngressSource() string { return annotations.Ingress }

// K8sGatewaySource returns the annotation key used to reference a Kubernetes Gateway API
// Gateway object for endpoint target resolution.
func K8sGatewaySource() string { return annotations.GatewayKey }

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
	gwAPIInformer            gwinformers_v1.GatewayInformer
}

// NewIstioGatewaySource creates a new gatewaySource with the given config.
func NewIstioGatewaySource(
	ctx context.Context,
	kubeClient kubernetes.Interface,
	istioClient istioclient.Interface,
	cfg *Config,
	gwAPIClient gatewayclient.Interface,
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
		informers.IndexSelectorWithConditions(annotations.IsControllerMatch[*networkingv1.Gateway]),
	))

	// Add default resource event handlers to properly initialize informer.
	informers.MustAddEventHandler(serviceInformer.Informer(), informers.DefaultEventHandler())
	informers.MustAddEventHandler(ingressInformer.Informer(), informers.DefaultEventHandler())
	informers.MustAddEventHandler(gatewayInformer.Informer(), informers.DefaultEventHandler())

	// Optionally set up a Gateway API Gateway informer if a client is provided
	// and the Gateway API CRDs are installed in the cluster.
	var gwAPIInformer gwinformers_v1.GatewayInformer
	var gwAPIInformerFactory gwinformers.SharedInformerFactory
	if gwAPIClient != nil {
		if _, err := gwAPIClient.GatewayV1().Gateways("").List(ctx, metav1.ListOptions{Limit: 1}); err != nil {
			log.Debugf("Gateway API not available, %q annotation will not be supported: %v", K8sGatewaySource(), err)
		} else {
			gwAPIInformerFactory = gwinformers.NewSharedInformerFactory(gwAPIClient, 0)
			gwAPIInformer = gwAPIInformerFactory.Gateway().V1().Gateways()
			informers.MustSetTransform(gwAPIInformer.Informer(), informers.TransformerWithOptions[*gatewayv1.Gateway](
				informers.TransformRemoveManagedFields(),
				informers.TransformRemoveLastAppliedConfig(),
			))
			informers.MustAddEventHandler(gwAPIInformer.Informer(), informers.DefaultEventHandler())
		}
	}

	informerFactory.Start(ctx.Done())
	istioInformerFactory.Start(ctx.Done())
	if gwAPIInformerFactory != nil {
		gwAPIInformerFactory.Start(ctx.Done())
	}

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}
	if err := informers.WaitForCacheSync(ctx, istioInformerFactory); err != nil {
		return nil, err
	}
	if gwAPIInformerFactory != nil {
		if err := informers.WaitForCacheSync(ctx, gwAPIInformerFactory); err != nil {
			return nil, err
		}
	}

	return &gatewaySource{
		namespace:                cfg.Namespace,
		annotationFilter:         cfg.AnnotationFilter,
		templateEngine:           cfg.TemplateEngine,
		ignoreHostnameAnnotation: cfg.IgnoreHostnameAnnotation,
		serviceInformer:          serviceInformer,
		gatewayInformer:          gatewayInformer,
		ingressInformer:          ingressInformer,
		gwAPIInformer:            gwAPIInformer,
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

		gwHostnames := sc.hostNamesFromGateway(gateway)

		log.Debugf("Processing gateway '%s/%s.%s' and hosts %q", gateway.Namespace, gateway.APIVersion, gateway.Name, strings.Join(gwHostnames, ","))

		gwEndpoints, err := sc.endpointsFromGateway(gwHostnames, gateway)
		if err != nil {
			log.Warnf("Could not generate endpoints for gateway '%s/%s': %v", gateway.Namespace, gateway.Name, err)
			continue
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
			log.Warnf("Could not apply template for gateway '%s/%s': %v", gateway.Namespace, gateway.Name, err)
			continue
		}

		if endpoint.HasNoEmptyEndpoints(gwEndpoints, types.IstioGateway, gateway) {
			continue
		}

		log.Debugf("Endpoints generated from '%s/%s/%s': %q", strings.ToLower(gateway.Kind), gateway.Namespace, gateway.Name, gwEndpoints)
		endpoints = append(endpoints, gwEndpoints...)
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

// AddEventHandler adds an event handler that should be triggered if the watched Istio Gateway changes.
func (sc *gatewaySource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for Istio Gateway")

	informers.MustAddEventHandler(sc.gatewayInformer.Informer(), eventHandlerFunc(handler))
}

func (sc *gatewaySource) targetsFromIngress(ingressStr string, gateway *networkingv1.Gateway) (endpoint.Targets, error) {
	namespace, name, err := ParseNamespacedName(ingressStr)
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

func (sc *gatewaySource) targetsFromGatewayAPIGateway(gatewayStr string, gateway *networkingv1.Gateway) (endpoint.Targets, error) {
	return EndpointTargetsFromK8sGateway(sc.gwAPIInformer, gatewayStr, gateway.Namespace)
}

func (sc *gatewaySource) targetsFromGateway(gateway *networkingv1.Gateway) (endpoint.Targets, error) {
	targets := annotations.TargetsFromTargetAnnotation(gateway.Annotations)
	if len(targets) > 0 {
		return targets, nil
	}

	ingressStr, ok := gateway.Annotations[IstioGatewayIngressSource()]
	if ok && ingressStr != "" {
		return sc.targetsFromIngress(ingressStr, gateway)
	}

	gatewayStr, ok := gateway.Annotations[K8sGatewaySource()]
	if ok && gatewayStr != "" {
		return sc.targetsFromGatewayAPIGateway(gatewayStr, gateway)
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
