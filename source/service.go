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
	"maps"
	"net"
	"slices"
	"sort"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	discoveryinformers "k8s.io/client-go/informers/discovery/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/source/informers"

	"sigs.k8s.io/external-dns/source/annotations"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/fqdn"
)

var (
	knownServiceTypes = map[v1.ServiceType]struct{}{
		v1.ServiceTypeClusterIP:    {}, // Default service type exposes the service on a cluster-internal IP.
		v1.ServiceTypeNodePort:     {}, // Exposes the service on each node's IP at a static port.
		v1.ServiceTypeLoadBalancer: {}, // Exposes the service externally using a cloud provider's load balancer.
		v1.ServiceTypeExternalName: {}, // Maps the service to an external DNS name.
	}
	serviceNameIndexKey = "serviceName"
)

// serviceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' entrypoints it will return a corresponding
// Endpoint object.
type serviceSource struct {
	client                kubernetes.Interface
	namespace             string
	annotationFilter      string
	labelSelector         labels.Selector
	fqdnTemplate          *template.Template
	combineFQDNAnnotation bool

	ignoreHostnameAnnotation       bool
	publishInternal                bool
	publishHostIP                  bool
	alwaysPublishNotReadyAddresses bool
	resolveLoadBalancerHostname    bool
	listenEndpointEvents           bool
	serviceInformer                coreinformers.ServiceInformer
	endpointSlicesInformer         discoveryinformers.EndpointSliceInformer
	podInformer                    coreinformers.PodInformer
	nodeInformer                   coreinformers.NodeInformer
	serviceTypeFilter              *serviceTypes
	exposeInternalIPv6             bool

	// process Services with legacy annotations
	compatibility string
}

// NewServiceSource creates a new serviceSource with the given config.
func NewServiceSource(ctx context.Context, kubeClient kubernetes.Interface, namespace, annotationFilter, fqdnTemplate string, combineFqdnAnnotation bool, compatibility string, publishInternal, publishHostIP, alwaysPublishNotReadyAddresses bool, serviceTypeFilter []string, ignoreHostnameAnnotation bool, labelSelector labels.Selector, resolveLoadBalancerHostname, listenEndpointEvents bool, exposeInternalIPv6 bool) (Source, error) {
	tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	// Use shared informers to listen for add/update/delete of services/pods/nodes in the specified namespace.
	// Set the resync period to 0 to prevent processing when nothing has changed
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0, kubeinformers.WithNamespace(namespace))
	serviceInformer := informerFactory.Core().V1().Services()
	endpointSlicesInformer := informerFactory.Discovery().V1().EndpointSlices()
	podInformer := informerFactory.Core().V1().Pods()
	nodeInformer := informerFactory.Core().V1().Nodes()

	// Add default resource event handlers to properly initialize informer.
	serviceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)
	endpointSlicesInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)
	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)
	nodeInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)

	// Add an indexer to the EndpointSlice informer to index by the service name label
	err = endpointSlicesInformer.Informer().AddIndexers(cache.Indexers{
		serviceNameIndexKey: func(obj any) ([]string, error) {
			endpointSlice, ok := obj.(*discoveryv1.EndpointSlice)
			if !ok {
				// This should never happen because the Informer should only contain EndpointSlice objects
				return nil, fmt.Errorf("expected %T but got %T instead", endpointSlice, obj)
			}
			serviceName := endpointSlice.Labels[discoveryv1.LabelServiceName]
			if serviceName == "" {
				return nil, nil
			}
			key := types.NamespacedName{Namespace: endpointSlice.Namespace, Name: serviceName}.String()
			return []string{key}, nil
		},
	})
	if err != nil {
		return nil, err
	}

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	// Transform the slice into a map so it will be way much easier and fast to filter later
	sTypesFilter, err := newServiceTypesFilter(serviceTypeFilter)
	if err != nil {
		return nil, err
	}

	return &serviceSource{
		client:                         kubeClient,
		namespace:                      namespace,
		annotationFilter:               annotationFilter,
		compatibility:                  compatibility,
		fqdnTemplate:                   tmpl,
		combineFQDNAnnotation:          combineFqdnAnnotation,
		ignoreHostnameAnnotation:       ignoreHostnameAnnotation,
		publishInternal:                publishInternal,
		publishHostIP:                  publishHostIP,
		alwaysPublishNotReadyAddresses: alwaysPublishNotReadyAddresses,
		serviceInformer:                serviceInformer,
		endpointSlicesInformer:         endpointSlicesInformer,
		podInformer:                    podInformer,
		nodeInformer:                   nodeInformer,
		serviceTypeFilter:              sTypesFilter,
		labelSelector:                  labelSelector,
		resolveLoadBalancerHostname:    resolveLoadBalancerHostname,
		listenEndpointEvents:           listenEndpointEvents,
		exposeInternalIPv6:             exposeInternalIPv6,
	}, nil
}

// Endpoints return endpoint objects for each service that should be processed.
func (sc *serviceSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	services, err := sc.serviceInformer.Lister().Services(sc.namespace).List(sc.labelSelector)
	if err != nil {
		return nil, err
	}

	// filter on service types if at least one has been provided
	services = sc.filterByServiceType(services)

	services, err = sc.filterByAnnotations(services)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, svc := range services {
		// Check controller annotation to see if we are responsible.
		controller, ok := svc.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping service %s/%s because controller value does not match, found: %s, required: %s",
				svc.Namespace, svc.Name, controller, controllerAnnotationValue)
			continue
		}

		svcEndpoints := sc.endpoints(svc)

		// process legacy annotations if no endpoints were returned and compatibility mode is enabled.
		if len(svcEndpoints) == 0 && sc.compatibility != "" {
			svcEndpoints, err = legacyEndpointsFromService(svc, sc)
			if err != nil {
				return nil, err
			}
		}

		// apply template if none of the above is found
		if (sc.combineFQDNAnnotation || len(svcEndpoints) == 0) && sc.fqdnTemplate != nil {
			sEndpoints, err := sc.endpointsFromTemplate(svc)
			if err != nil {
				return nil, err
			}

			if sc.combineFQDNAnnotation {
				svcEndpoints = append(svcEndpoints, sEndpoints...)
			} else {
				svcEndpoints = sEndpoints
			}
		}

		if len(svcEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from service %s/%s", svc.Namespace, svc.Name)
			continue
		}

		log.Debugf("Endpoints generated from service: %s/%s: %v", svc.Namespace, svc.Name, svcEndpoints)
		endpoints = append(endpoints, svcEndpoints...)
	}

	// this sorting is required to make merging work.
	// after we merge endpoints that have same DNS, we want to ensure that we end up with the same service being an "owner"
	// of all those records, as otherwise each time we update, we will end up with a different service that gets data merged in
	// and that will cause external-dns to recreate dns record due to different service owner in TXT record.
	// if new service is added or old one removed, that might cause existing record to get re-created due to potentially new
	// owner being selected. Which is fine, since it shouldn't happen often and shouldn't cause any disruption.
	if len(endpoints) > 1 {
		sort.Slice(endpoints, func(i, j int) bool {
			return endpoints[i].Labels[endpoint.ResourceLabelKey] < endpoints[j].Labels[endpoint.ResourceLabelKey]
		})
		// Use stable sort to not disrupt the order of services
		sort.SliceStable(endpoints, func(i, j int) bool {
			if endpoints[i].DNSName != endpoints[j].DNSName {
				return endpoints[i].DNSName < endpoints[j].DNSName
			}
			return endpoints[i].RecordType < endpoints[j].RecordType
		})
		mergedEndpoints := []*endpoint.Endpoint{}
		mergedEndpoints = append(mergedEndpoints, endpoints[0])
		for i := 1; i < len(endpoints); i++ {
			lastMergedEndpoint := len(mergedEndpoints) - 1
			if mergedEndpoints[lastMergedEndpoint].DNSName == endpoints[i].DNSName &&
				mergedEndpoints[lastMergedEndpoint].RecordType == endpoints[i].RecordType &&
				mergedEndpoints[lastMergedEndpoint].RecordType != endpoint.RecordTypeCNAME && // It is against RFC-1034 for CNAME records to have multiple targets, so skip merging
				mergedEndpoints[lastMergedEndpoint].SetIdentifier == endpoints[i].SetIdentifier &&
				mergedEndpoints[lastMergedEndpoint].RecordTTL == endpoints[i].RecordTTL {
				mergedEndpoints[lastMergedEndpoint].Targets = append(mergedEndpoints[lastMergedEndpoint].Targets, endpoints[i].Targets[0])
			} else {
				mergedEndpoints = append(mergedEndpoints, endpoints[i])
			}

			if mergedEndpoints[lastMergedEndpoint].DNSName == endpoints[i].DNSName &&
				mergedEndpoints[lastMergedEndpoint].RecordType == endpoints[i].RecordType &&
				mergedEndpoints[lastMergedEndpoint].RecordType == endpoint.RecordTypeCNAME {
				log.Debugf("CNAME %s with multiple targets found", endpoints[i].DNSName)
			}
		}
		endpoints = mergedEndpoints
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

// extractHeadlessEndpoints extracts endpoints from a headless service using the "Endpoints" Kubernetes API resource
func (sc *serviceSource) extractHeadlessEndpoints(svc *v1.Service, hostname string, ttl endpoint.TTL) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	labelSelector, err := metav1.ParseToLabelSelector(labels.Set(svc.Spec.Selector).AsSelectorPreValidated().String())
	if err != nil {
		return nil
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil
	}

	serviceKey := cache.ObjectName{Namespace: svc.Namespace, Name: svc.Name}.String()
	rawEndpointSlices, err := sc.endpointSlicesInformer.Informer().GetIndexer().ByIndex(serviceNameIndexKey, serviceKey)
	if err != nil {
		// Should never happen as long as the index exists
		log.Errorf("Get EndpointSlices of service[%s] error:%v", svc.GetName(), err)
		return nil
	}

	endpointSlices := make([]*discoveryv1.EndpointSlice, 0, len(rawEndpointSlices))
	for _, obj := range rawEndpointSlices {
		endpointSlice, ok := obj.(*discoveryv1.EndpointSlice)
		if !ok {
			// Should never happen as the indexer can only contain EndpointSlice objects
			log.Errorf("Expected %T but got %T instead, skipping", endpointSlice, obj)
			continue
		}
		endpointSlices = append(endpointSlices, endpointSlice)
	}

	pods, err := sc.podInformer.Lister().Pods(svc.Namespace).List(selector)
	if err != nil {
		log.Errorf("List Pods of service[%s] error:%v", svc.GetName(), err)
		return endpoints
	}

	endpointsType := getEndpointsTypeFromAnnotations(svc.Annotations)
	publishPodIPs := endpointsType != EndpointsTypeNodeExternalIP && endpointsType != EndpointsTypeHostIP && !sc.publishHostIP
	publishNotReadyAddresses := svc.Spec.PublishNotReadyAddresses || sc.alwaysPublishNotReadyAddresses

	targetsByHeadlessDomainAndType := make(map[endpoint.EndpointKey]endpoint.Targets)
	for _, endpointSlice := range endpointSlices {
		for _, ep := range endpointSlice.Endpoints {
			if !conditionToBool(ep.Conditions.Ready) && !publishNotReadyAddresses {
				continue
			}

			if publishPodIPs &&
				endpointSlice.AddressType != discoveryv1.AddressTypeIPv4 &&
				endpointSlice.AddressType != discoveryv1.AddressTypeIPv6 {
				log.Debugf("Skipping EndpointSlice %s/%s because its address type is unsupported: %s", endpointSlice.Namespace, endpointSlice.Name, endpointSlice.AddressType)
				continue
			}

			// find pod for this address
			if ep.TargetRef == nil || ep.TargetRef.APIVersion != "" || ep.TargetRef.Kind != "Pod" {
				log.Debugf("Skipping address because its target is not a pod: %v", ep)
				continue
			}
			var pod *v1.Pod
			for _, v := range pods {
				if v.Name == ep.TargetRef.Name {
					pod = v
					break
				}
			}
			if pod == nil {
				log.Errorf("Pod %s not found for address %v", ep.TargetRef.Name, ep)
				continue
			}

			headlessDomains := []string{hostname}
			if pod.Spec.Hostname != "" {
				headlessDomains = append(headlessDomains, fmt.Sprintf("%s.%s", pod.Spec.Hostname, hostname))
			}

			for _, headlessDomain := range headlessDomains {
				targets := annotations.TargetsFromTargetAnnotation(pod.Annotations)
				if len(targets) == 0 {
					if endpointsType == EndpointsTypeNodeExternalIP {
						node, err := sc.nodeInformer.Lister().Get(pod.Spec.NodeName)
						if err != nil {
							log.Errorf("Get node[%s] of pod[%s] error: %v; not adding any NodeExternalIP endpoints", pod.Spec.NodeName, pod.GetName(), err)
							return endpoints
						}
						for _, address := range node.Status.Addresses {
							if address.Type == v1.NodeExternalIP || (sc.exposeInternalIPv6 && address.Type == v1.NodeInternalIP && suitableType(address.Address) == endpoint.RecordTypeAAAA) {
								targets = append(targets, address.Address)
								log.Debugf("Generating matching endpoint %s with NodeExternalIP %s", headlessDomain, address.Address)
							}
						}
					} else if endpointsType == EndpointsTypeHostIP || sc.publishHostIP {
						targets = endpoint.Targets{pod.Status.HostIP}
						log.Debugf("Generating matching endpoint %s with HostIP %s", headlessDomain, pod.Status.HostIP)
					} else {
						if len(ep.Addresses) == 0 {
							log.Warnf("EndpointSlice %s/%s has no addresses for endpoint %v", endpointSlice.Namespace, endpointSlice.Name, ep)
							continue
						}
						address := ep.Addresses[0] // Only use the first address, as additional addresses have no semantic defined
						targets = endpoint.Targets{address}
						log.Debugf("Generating matching endpoint %s with EndpointSliceAddress IP %s", headlessDomain, address)
					}
				}
				for _, target := range targets {
					key := endpoint.EndpointKey{
						DNSName:    headlessDomain,
						RecordType: suitableType(target),
					}
					targetsByHeadlessDomainAndType[key] = append(targetsByHeadlessDomainAndType[key], target)
				}
			}
		}
	}

	headlessKeys := []endpoint.EndpointKey{}
	for headlessKey := range targetsByHeadlessDomainAndType {
		headlessKeys = append(headlessKeys, headlessKey)
	}
	sort.Slice(headlessKeys, func(i, j int) bool {
		if headlessKeys[i].DNSName != headlessKeys[j].DNSName {
			return headlessKeys[i].DNSName < headlessKeys[j].DNSName
		}
		return headlessKeys[i].RecordType < headlessKeys[j].RecordType
	})
	for _, headlessKey := range headlessKeys {
		allTargets := targetsByHeadlessDomainAndType[headlessKey]
		targets := []string{}

		deduppedTargets := map[string]struct{}{}
		for _, target := range allTargets {
			if _, ok := deduppedTargets[target]; ok {
				log.Debugf("Removing duplicate target %s", target)
				continue
			}

			deduppedTargets[target] = struct{}{}
			targets = append(targets, target)
		}

		var ep *endpoint.Endpoint
		if ttl.IsConfigured() {
			ep = endpoint.NewEndpointWithTTL(headlessKey.DNSName, headlessKey.RecordType, ttl, targets...)
		} else {
			ep = endpoint.NewEndpoint(headlessKey.DNSName, headlessKey.RecordType, targets...)
		}

		if ep != nil {
			ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("service/%s/%s", svc.Namespace, svc.Name))
			endpoints = append(endpoints, ep)
		}
	}

	return endpoints
}

func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
	hostnames, err := fqdn.ExecTemplate(sc.fqdnTemplate, svc)
	if err != nil {
		return nil, err
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(svc.Annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		endpoints = append(endpoints, sc.generateEndpoints(svc, hostname, providerSpecific, setIdentifier, false)...)
	}

	return endpoints, nil
}

// endpointsFromService extracts the endpoints from a service object
func (sc *serviceSource) endpoints(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint
	// Skip endpoints if we do not want entries from annotations
	if sc.ignoreHostnameAnnotation {
		return endpoints
	}

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(svc.Annotations)
	var hostnameList []string
	var internalHostnameList []string

	hostnameList = annotations.HostnamesFromAnnotations(svc.Annotations)
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, sc.generateEndpoints(svc, hostname, providerSpecific, setIdentifier, false)...)
	}

	internalHostnameList = annotations.InternalHostnamesFromAnnotations(svc.Annotations)
	for _, hostname := range internalHostnameList {
		endpoints = append(endpoints, sc.generateEndpoints(svc, hostname, providerSpecific, setIdentifier, true)...)
	}

	return endpoints
}

// filterByAnnotations filters a list of services by a given annotation selector.
func (sc *serviceSource) filterByAnnotations(services []*v1.Service) ([]*v1.Service, error) {
	selector, err := annotations.ParseFilter(sc.annotationFilter)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return services, nil
	}

	var filteredList []*v1.Service

	for _, service := range services {
		// include service if its annotations match the selector
		if selector.Matches(labels.Set(service.Annotations)) {
			filteredList = append(filteredList, service)
		}
	}
	log.Debugf("filtered %d services out of %d with annotation filter", len(filteredList), len(services))
	return filteredList, nil
}

// filterByServiceType filters services according to their types
func (sc *serviceSource) filterByServiceType(services []*v1.Service) []*v1.Service {
	if !sc.serviceTypeFilter.enabled || len(services) == 0 {
		return services
	}
	var result []*v1.Service
	for _, service := range services {
		if _, ok := sc.serviceTypeFilter.types[service.Spec.Type]; ok {
			result = append(result, service)
		}
	}
	log.Debugf("filtered %d services out of %d with service types filter %q", len(result), len(services), slices.Collect(maps.Keys(sc.serviceTypeFilter.types)))
	return result
}

func (sc *serviceSource) generateEndpoints(svc *v1.Service, hostname string, providerSpecific endpoint.ProviderSpecific, setIdentifier string, useClusterIP bool) (endpoints []*endpoint.Endpoint) {
	hostname = strings.TrimSuffix(hostname, ".")

	resource := fmt.Sprintf("service/%s/%s", svc.Namespace, svc.Name)

	ttl := annotations.TTLFromAnnotations(svc.Annotations, resource)

	targets := annotations.TargetsFromTargetAnnotation(svc.Annotations)

	if len(targets) == 0 {
		switch svc.Spec.Type {
		case v1.ServiceTypeLoadBalancer:
			if useClusterIP {
				targets = extractServiceIps(svc)
			} else {
				targets = extractLoadBalancerTargets(svc, sc.resolveLoadBalancerHostname)
			}
		case v1.ServiceTypeClusterIP:
			if svc.Spec.ClusterIP == v1.ClusterIPNone {
				endpoints = append(endpoints, sc.extractHeadlessEndpoints(svc, hostname, ttl)...)
			} else if useClusterIP || sc.publishInternal {
				targets = extractServiceIps(svc)
			}
		case v1.ServiceTypeNodePort:
			// add the nodeTargets and extract an SRV endpoint
			var err error
			targets, err = sc.extractNodePortTargets(svc)
			if err != nil {
				log.Errorf("Unable to extract targets from service %s/%s error: %v", svc.Namespace, svc.Name, err)
				return endpoints
			}
			endpoints = append(endpoints, sc.extractNodePortEndpoints(svc, hostname, ttl)...)
		case v1.ServiceTypeExternalName:
			targets = extractServiceExternalName(svc)
		}

		for _, en := range endpoints {
			en.ProviderSpecific = providerSpecific
			en.SetIdentifier = setIdentifier
		}
	}

	endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)

	return endpoints
}

func extractServiceIps(svc *v1.Service) endpoint.Targets {
	if svc.Spec.ClusterIP == v1.ClusterIPNone {
		log.Debugf("Unable to associate %s headless service with a Cluster IP", svc.Name)
		return endpoint.Targets{}
	}
	return endpoint.Targets{svc.Spec.ClusterIP}
}

func extractServiceExternalName(svc *v1.Service) endpoint.Targets {
	if len(svc.Spec.ExternalIPs) > 0 {
		return svc.Spec.ExternalIPs
	}
	return endpoint.Targets{svc.Spec.ExternalName}
}

func extractLoadBalancerTargets(svc *v1.Service, resolveLoadBalancerHostname bool) endpoint.Targets {
	if len(svc.Spec.ExternalIPs) > 0 {
		return svc.Spec.ExternalIPs
	}

	// Create a corresponding endpoint for each configured external entrypoint.
	var targets endpoint.Targets
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			targets = append(targets, lb.IP)
		}
		if lb.Hostname != "" {
			if resolveLoadBalancerHostname {
				ips, err := net.LookupIP(lb.Hostname)
				if err != nil {
					log.Errorf("Unable to resolve %q: %v", lb.Hostname, err)
					continue
				}
				for _, ip := range ips {
					targets = append(targets, ip.String())
				}
			} else {
				targets = append(targets, lb.Hostname)
			}
		}
	}

	return targets
}

func isPodStatusReady(status v1.PodStatus) bool {
	_, condition := getPodCondition(&status, v1.PodReady)
	return condition != nil && condition.Status == v1.ConditionTrue
}

func getPodCondition(status *v1.PodStatus, conditionType v1.PodConditionType) (int, *v1.PodCondition) {
	if status == nil {
		return -1, nil
	}
	return getPodConditionFromList(status.Conditions, conditionType)
}

func getPodConditionFromList(conditions []v1.PodCondition, conditionType v1.PodConditionType) (int, *v1.PodCondition) {
	if conditions == nil {
		return -1, nil
	}
	for i := range conditions {
		if conditions[i].Type == conditionType {
			return i, &conditions[i]
		}
	}
	return -1, nil
}

// nodesExternalTrafficPolicyTypeLocal filters nodes that have running pods belonging to the given NodePort service
// with externalTrafficPolicy=Local. Returns a prioritized slice of nodes, favoring those with ready, non-terminating pods.
func (sc *serviceSource) nodesExternalTrafficPolicyTypeLocal(svc *v1.Service) []*v1.Node {
	var nodesReady []*v1.Node
	var nodesRunning []*v1.Node
	var nodes []*v1.Node
	nodesMap := map[*v1.Node]struct{}{}

	pods := sc.pods(svc)

	for _, v := range pods {
		if v.Status.Phase == v1.PodRunning {
			node, err := sc.nodeInformer.Lister().Get(v.Spec.NodeName)
			if err != nil {
				log.Debugf("Unable to find node where Pod %s is running", v.Spec.Hostname)
				continue
			}

			if _, ok := nodesMap[node]; !ok {
				nodesMap[node] = *new(struct{})
				nodesRunning = append(nodesRunning, node)

				if isPodStatusReady(v.Status) {
					nodesReady = append(nodesReady, node)
					// Check pod not terminating
					if v.GetDeletionTimestamp() == nil {
						nodes = append(nodes, node)
					}
				}
			}
		}
	}

	// Prioritize nodes with non-terminating ready pods
	// If none available, fall back to nodes with ready pods
	// If still none, use nodes with any running pods
	if len(nodes) > 0 {
		// Works the same as service endpoints
	} else if len(nodesReady) > 0 {
		// 2 level of panic modes as safeguard, because old wrong behavior can be used by someone
		// Publish all endpoints not always a bad thing
		log.Debugf("All pods in terminating state, use ready")
		nodes = nodesReady
	} else {
		log.Debugf("All pods not ready, use all running")
		nodes = nodesRunning
	}

	return nodes
}

// pods retrieve a slice of pods associated with the given Service
func (sc *serviceSource) pods(svc *v1.Service) []*v1.Pod {
	labelSelector, err := metav1.ParseToLabelSelector(labels.Set(svc.Spec.Selector).AsSelectorPreValidated().String())
	if err != nil {
		return nil
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil
	}
	pods, err := sc.podInformer.Lister().Pods(svc.Namespace).List(selector)
	if err != nil {
		return nil
	}

	return pods
}

func (sc *serviceSource) extractNodePortTargets(svc *v1.Service) (endpoint.Targets, error) {
	var (
		internalIPs endpoint.Targets
		externalIPs endpoint.Targets
		ipv6IPs     endpoint.Targets
		nodes       []*v1.Node
	)

	if svc.Spec.ExternalTrafficPolicy == v1.ServiceExternalTrafficPolicyTypeLocal {
		nodes = sc.nodesExternalTrafficPolicyTypeLocal(svc)
	} else {
		var err error
		nodes, err = sc.nodeInformer.Lister().List(labels.Everything())
		if err != nil {
			return nil, err
		}
	}

	for _, node := range nodes {
		for _, address := range node.Status.Addresses {
			switch address.Type {
			case v1.NodeExternalIP:
				externalIPs = append(externalIPs, address.Address)
			case v1.NodeInternalIP:
				internalIPs = append(internalIPs, address.Address)
				if suitableType(address.Address) == endpoint.RecordTypeAAAA {
					ipv6IPs = append(ipv6IPs, address.Address)
				}
			}
		}
	}

	access := getAccessFromAnnotations(svc.Annotations)
	switch access {
	case "public":
		return append(externalIPs, ipv6IPs...), nil
	case "private":
		return internalIPs, nil
	}

	if len(externalIPs) > 0 {
		return append(externalIPs, ipv6IPs...), nil
	}

	return internalIPs, nil
}

func (sc *serviceSource) extractNodePortEndpoints(svc *v1.Service, hostname string, ttl endpoint.TTL) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	for _, port := range svc.Spec.Ports {
		if port.NodePort > 0 {
			// following the RFC 2782, SRV record must have a following format
			// _service._proto.name. TTL class SRV priority weight port
			// see https://en.wikipedia.org/wiki/SRV_record

			// build a target with a priority of 0, weight of 50, and pointing the given port on the given host
			target := fmt.Sprintf("0 50 %d %s", port.NodePort, hostname)

			// take the service name from the K8s Service object
			// it is safe to use since it is DNS compatible
			// see https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-label-names
			serviceName := svc.Name

			// figure out the protocol
			protocol := strings.ToLower(string(port.Protocol))
			if protocol == "" {
				protocol = "tcp"
			}

			recordName := fmt.Sprintf("_%s._%s.%s", serviceName, protocol, hostname)

			var ep *endpoint.Endpoint
			if ttl.IsConfigured() {
				ep = endpoint.NewEndpointWithTTL(recordName, endpoint.RecordTypeSRV, ttl, target)
			} else {
				ep = endpoint.NewEndpoint(recordName, endpoint.RecordTypeSRV, target)
			}

			if ep != nil {
				ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("service/%s/%s", svc.Namespace, svc.Name))
				endpoints = append(endpoints, ep)
			}
		}
	}

	return endpoints
}

func (sc *serviceSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for service")

	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	sc.serviceInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	if sc.listenEndpointEvents {
		sc.endpointSlicesInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
	}
}

type serviceTypes struct {
	enabled bool
	types   map[v1.ServiceType]bool
}

// newServiceTypesFilter processes a slice of service type filter strings and returns a serviceTypes struct.
// It validates the filter against known Kubernetes service types. If the filter is empty or contains an empty string,
// service type filtering is disabled. If an unknown type is found, an error is returned.
func newServiceTypesFilter(filter []string) (*serviceTypes, error) {
	if len(filter) == 0 || slices.Contains(filter, "") {
		return &serviceTypes{
			enabled: false,
		}, nil
	}
	types := make(map[v1.ServiceType]bool)
	for _, serviceType := range filter {
		if _, ok := knownServiceTypes[v1.ServiceType(serviceType)]; !ok {
			return nil, fmt.Errorf("unsupported service type filter: %q. Supported types are: %q", serviceType, slices.Collect(maps.Keys(knownServiceTypes)))
		}
		types[v1.ServiceType(serviceType)] = true
	}

	return &serviceTypes{
		enabled: true,
		types:   types,
	}, nil
}

// conditionToBool converts an EndpointConditions condition to a bool value.
func conditionToBool(v *bool) bool {
	if v == nil {
		return true // nil should be interpreted as "true" as per EndpointConditions spec
	}
	return *v
}
