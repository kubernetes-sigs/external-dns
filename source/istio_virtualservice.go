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
	"cmp"
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	istioinformers "istio.io/client-go/pkg/informers/externalversions"
	networkingv1beta1informer "istio.io/client-go/pkg/informers/externalversions/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/fqdn"
	"sigs.k8s.io/external-dns/source/informers"
)

// IstioMeshGateway is the built in gateway for all sidecars
const IstioMeshGateway = "mesh"

// virtualServiceSource is an implementation of Source for Istio VirtualService objects.
// The implementation uses the spec.hosts values for the hostnames.
// Use targetAnnotationKey to explicitly set Endpoint.
type virtualServiceSource struct {
	kubeClient               kubernetes.Interface
	istioClient              istioclient.Interface
	namespace                string
	annotationFilter         string
	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
	serviceInformer          coreinformers.ServiceInformer
	vServiceInformer         networkingv1beta1informer.VirtualServiceInformer
	gatewayInformer          networkingv1beta1informer.GatewayInformer
}

// NewIstioVirtualServiceSource creates a new virtualServiceSource with the given config.
func NewIstioVirtualServiceSource(
	ctx context.Context,
	kubeClient kubernetes.Interface,
	istioClient istioclient.Interface,
	namespace string,
	annotationFilter string,
	fqdnTemplate string,
	combineFQDNAnnotation bool,
	ignoreHostnameAnnotation bool,
) (Source, error) {
	tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	// Use shared informers to listen for add/update/delete of services/pods/nodes in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0, kubeinformers.WithNamespace(namespace))
	serviceInformer := informerFactory.Core().V1().Services()
	istioInformerFactory := istioinformers.NewSharedInformerFactoryWithOptions(istioClient, 0, istioinformers.WithNamespace(namespace))
	virtualServiceInformer := istioInformerFactory.Networking().V1beta1().VirtualServices()
	gatewayInformer := istioInformerFactory.Networking().V1beta1().Gateways()

	// Add default resource event handlers to properly initialize informer.
	_, _ = serviceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("service added")
			},
		},
	)

	_, _ = virtualServiceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("virtual service added")
			},
		},
	)

	_, _ = gatewayInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("gateway added")
			},
		},
	)

	informerFactory.Start(ctx.Done())
	istioInformerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}
	if err := informers.WaitForCacheSync(context.Background(), istioInformerFactory); err != nil {
		return nil, err
	}

	return &virtualServiceSource{
		kubeClient:               kubeClient,
		istioClient:              istioClient,
		namespace:                namespace,
		annotationFilter:         annotationFilter,
		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    combineFQDNAnnotation,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
		serviceInformer:          serviceInformer,
		vServiceInformer:         virtualServiceInformer,
		gatewayInformer:          gatewayInformer,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all VirtualService resources in the source's namespace(s).
func (sc *virtualServiceSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	virtualServices, err := sc.vServiceInformer.Lister().VirtualServices(sc.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}
	virtualServices, err = sc.filterByAnnotations(virtualServices)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint

	log.Debugf("Found %d virtualservice in namespace %s", len(virtualServices), sc.namespace)

	for _, vService := range virtualServices {
		// Check controller annotation to see if we are responsible.
		controller, ok := vService.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping VirtualService %s/%s.%s because controller value does not match, found: %s, required: %s",
				vService.Namespace, vService.APIVersion, vService.Name, controller, controllerAnnotationValue)
			continue
		}

		gwEndpoints, err := sc.endpointsFromVirtualService(ctx, vService)
		if err != nil {
			return nil, err
		}

		// apply template if host is missing on VirtualService
		if (sc.combineFQDNAnnotation || len(gwEndpoints) == 0) && sc.fqdnTemplate != nil {
			iEndpoints, err := sc.endpointsFromTemplate(ctx, vService)
			if err != nil {
				return nil, err
			}

			if sc.combineFQDNAnnotation {
				gwEndpoints = append(gwEndpoints, iEndpoints...)
			} else {
				gwEndpoints = iEndpoints
			}
		}

		if len(gwEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from VirtualService %s/%s", vService.Namespace, vService.Name)
			continue
		}

		log.Debugf("Endpoints generated from %q '%s/%s.%s': %q", vService.Kind, vService.Namespace, vService.APIVersion, vService.Name, gwEndpoints)
		endpoints = append(endpoints, gwEndpoints...)
	}

	// TODO: sort on endpoint creation
	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

// AddEventHandler adds an event handler that should be triggered if the watched Istio VirtualService changes.
func (sc *virtualServiceSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for Istio VirtualService")

	_, _ = sc.vServiceInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

func (sc *virtualServiceSource) getGateway(_ context.Context, gatewayStr string, virtualService *v1beta1.VirtualService) (*v1beta1.Gateway, error) {
	if gatewayStr == "" || gatewayStr == IstioMeshGateway {
		// This refers to "all sidecars in the mesh"; ignore.
		return nil, nil
	}

	namespace, name, err := ParseIngress(gatewayStr)
	if err != nil {
		log.Debugf("Failed parsing gatewayStr %s of VirtualService %s/%s", gatewayStr, virtualService.Namespace, virtualService.Name)
		return nil, err
	}
	namespace = cmp.Or(namespace, virtualService.Namespace)

	gateway, err := sc.gatewayInformer.Lister().Gateways(namespace).Get(name)
	if errors.IsNotFound(err) {
		log.Warnf("VirtualService (%s/%s) references non-existent gateway: %s ", virtualService.Namespace, virtualService.Name, gatewayStr)
		return gateway, nil
	} else if err != nil {
		log.Errorf("Failed retrieving gateway %s referenced by VirtualService %s/%s: %v", gatewayStr, virtualService.Namespace, virtualService.Name, err)
		return nil, err
	}
	if gateway == nil {
		log.Debugf("Gateway %s referenced by VirtualService %s/%s not found: %v", gatewayStr, virtualService.Namespace, virtualService.Name, err)
		return gateway, nil
	}
	return gateway, nil
}

func (sc *virtualServiceSource) endpointsFromTemplate(ctx context.Context, virtualService *v1beta1.VirtualService) ([]*endpoint.Endpoint, error) {
	hostnames, err := fqdn.ExecTemplate(sc.fqdnTemplate, virtualService)
	if err != nil {
		return nil, err
	}

	resource := fmt.Sprintf("virtualservice/%s/%s", virtualService.Namespace, virtualService.Name)

	ttl := annotations.TTLFromAnnotations(virtualService.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(virtualService.Annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		targets, err := sc.targetsFromVirtualService(ctx, virtualService, hostname)
		if err != nil {
			return endpoints, err
		}
		endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}
	return endpoints, nil
}

// filterByAnnotations filters a list of configs by a given annotation selector.
func (sc *virtualServiceSource) filterByAnnotations(vServices []*v1beta1.VirtualService) ([]*v1beta1.VirtualService, error) {
	selector, err := annotations.ParseFilter(sc.annotationFilter)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return vServices, nil
	}

	var filteredList []*v1beta1.VirtualService

	for _, vs := range vServices {
		// include if the annotations match the selector
		if selector.Matches(labels.Set(vs.Annotations)) {
			filteredList = append(filteredList, vs)
		}
	}

	return filteredList, nil
}

// append a target to the list of targets unless it's already in the list
func appendUnique(targets []string, target string) []string {
	if slices.Contains(targets, target) {
		return targets
	}
	return append(targets, target)
}

func (sc *virtualServiceSource) targetsFromVirtualService(ctx context.Context, vService *v1beta1.VirtualService, vsHost string) ([]string, error) {
	var targets []string
	// for each host we need to iterate through the gateways because each host might match for only one of the gateways
	for _, gateway := range vService.Spec.Gateways {
		gw, err := sc.getGateway(ctx, gateway, vService)
		if err != nil {
			return nil, err
		}
		if gw == nil {
			continue
		}
		if !virtualServiceBindsToGateway(vService, gw, vsHost) {
			continue
		}
		tgs, err := sc.targetsFromGateway(ctx, gw)
		if err != nil {
			return targets, err
		}
		for _, target := range tgs {
			targets = appendUnique(targets, target)
		}
	}
	return targets, nil
}

// endpointsFromVirtualService extracts the endpoints from an Istio VirtualService Config object
func (sc *virtualServiceSource) endpointsFromVirtualService(ctx context.Context, vService *v1beta1.VirtualService) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	var err error

	resource := fmt.Sprintf("virtualservice/%s/%s", vService.Namespace, vService.Name)

	ttl := annotations.TTLFromAnnotations(vService.Annotations, resource)

	targetsFromAnnotation := annotations.TargetsFromTargetAnnotation(vService.Annotations)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(vService.Annotations)

	for _, host := range vService.Spec.Hosts {
		if host == "" || host == "*" {
			continue
		}

		parts := strings.Split(host, "/")

		// If the input hostname is of the form my-namespace/foo.bar.com, remove the namespace
		// before appending it to the list of endpoints to create
		if len(parts) == 2 {
			host = parts[1]
		}

		targets := targetsFromAnnotation
		if len(targets) == 0 {
			targets, err = sc.targetsFromVirtualService(ctx, vService, host)
			if err != nil {
				return endpoints, err
			}
		}

		endpoints = append(endpoints, EndpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier, resource)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(vService.Annotations)
		for _, hostname := range hostnameList {
			targets := targetsFromAnnotation
			if len(targets) == 0 {
				targets, err = sc.targetsFromVirtualService(ctx, vService, hostname)
				if err != nil {
					return endpoints, err
				}
			}
			endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	return endpoints, nil
}

// checks if the given VirtualService should actually bind to the given gateway
// see requirements here: https://istio.io/docs/reference/config/networking/gateway/#Server
func virtualServiceBindsToGateway(vService *v1beta1.VirtualService, gateway *v1beta1.Gateway, vsHost string) bool {
	isValid := false
	if len(vService.Spec.ExportTo) == 0 {
		isValid = true
	} else {
		for _, ns := range vService.Spec.ExportTo {
			if ns == "*" || ns == gateway.Namespace || (ns == "." && gateway.Namespace == vService.Namespace) {
				isValid = true
			}
		}
	}
	if !isValid {
		return false
	}

	for _, server := range gateway.Spec.Servers {
		for _, host := range server.Hosts {
			namespace := "*"
			parts := strings.Split(host, "/")
			if len(parts) == 2 {
				namespace = parts[0]
				host = parts[1]
			} else if len(parts) != 1 {
				log.Debugf("Gateway %s/%s has invalid host %s", gateway.Namespace, gateway.Name, host)
				continue
			}

			if namespace == "*" || namespace == vService.Namespace || (namespace == "." && vService.Namespace == gateway.Namespace) {
				if host == "*" {
					return true
				}

				suffixMatch := false
				if strings.HasPrefix(host, "*.") {
					suffixMatch = true
				}

				if host == vsHost || (suffixMatch && strings.HasSuffix(vsHost, host[1:])) {
					return true
				}
			}
		}
	}

	return false
}

func (sc *virtualServiceSource) targetsFromIngress(ctx context.Context, ingressStr string, gateway *v1beta1.Gateway) (endpoint.Targets, error) {
	namespace, name, err := ParseIngress(ingressStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Ingress annotation on Gateway (%s/%s): %w", gateway.Namespace, gateway.Name, err)
	}
	if namespace == "" {
		namespace = gateway.Namespace
	}

	ingress, err := sc.kubeClient.NetworkingV1().Ingresses(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	targets := make(endpoint.Targets, 0)

	for _, lb := range ingress.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			targets = append(targets, lb.IP)
		} else if lb.Hostname != "" {
			targets = append(targets, lb.Hostname)
		}
	}
	return targets, nil
}

func (sc *virtualServiceSource) targetsFromGateway(ctx context.Context, gateway *v1beta1.Gateway) (endpoint.Targets, error) {
	targets := annotations.TargetsFromTargetAnnotation(gateway.Annotations)
	if len(targets) > 0 {
		return targets, nil
	}

	ingressStr, ok := gateway.Annotations[IstioGatewayIngressSource]
	if ok && ingressStr != "" {
		return sc.targetsFromIngress(ctx, ingressStr, gateway)
	}

	return EndpointTargetsFromServices(sc.serviceInformer, sc.namespace, gateway.Spec.Selector)
}
