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
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"

	"time"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

const (
	defaultTargetsCapacity = 10
)

// serviceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' entrypoints it will return a corresponding
// Endpoint object.
type serviceSource struct {
	client           kubernetes.Interface
	namespace        string
	annotationFilter string
	// process Services with legacy annotations
	compatibility            string
	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
	publishInternal          bool
	publishHostIP            bool
	serviceInformer          coreinformers.ServiceInformer
	podInformer              coreinformers.PodInformer
	nodeInformer             coreinformers.NodeInformer
	serviceTypeFilter        map[string]struct{}
}

// NewServiceSource creates a new serviceSource with the given config.
func NewServiceSource(kubeClient kubernetes.Interface, namespace, annotationFilter string, fqdnTemplate string, combineFqdnAnnotation bool, compatibility string, publishInternal bool, publishHostIP bool, serviceTypeFilter []string, ignoreHostnameAnnotation bool) (Source, error) {
	var (
		tmpl *template.Template
		err  error
	)
	if fqdnTemplate != "" {
		tmpl, err = template.New("endpoint").Funcs(template.FuncMap{
			"trimPrefix": strings.TrimPrefix,
		}).Parse(fqdnTemplate)
		if err != nil {
			return nil, err
		}
	}

	// Use shared informers to listen for add/update/delete of services/pods/nodes in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0, kubeinformers.WithNamespace(namespace))
	serviceInformer := informerFactory.Core().V1().Services()
	podInformer := informerFactory.Core().V1().Pods()
	nodeInformer := informerFactory.Core().V1().Nodes()

	// Add default resource event handlers to properly initialize informer.
	serviceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("service added")
			},
		},
	)
	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("pod added")
			},
		},
	)
	nodeInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("node added")
			},
		},
	)

	// TODO informer is not explicitly stopped since controller is not passing in its channel.
	informerFactory.Start(wait.NeverStop)

	// wait for the local cache to be populated.
	err = wait.Poll(time.Second, 60*time.Second, func() (bool, error) {
		return serviceInformer.Informer().HasSynced() == true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sync cache: %v", err)
	}

	// Transform the slice into a map so it will
	// be way much easier and fast to filter later
	serviceTypes := make(map[string]struct{})
	for _, serviceType := range serviceTypeFilter {
		serviceTypes[serviceType] = struct{}{}
	}

	return &serviceSource{
		client:                   kubeClient,
		namespace:                namespace,
		annotationFilter:         annotationFilter,
		compatibility:            compatibility,
		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    combineFqdnAnnotation,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
		publishInternal:          publishInternal,
		publishHostIP:            publishHostIP,
		serviceInformer:          serviceInformer,
		podInformer:              podInformer,
		nodeInformer:             nodeInformer,
		serviceTypeFilter:        serviceTypes,
	}, nil
}

// Endpoints returns endpoint objects for each service that should be processed.
func (sc *serviceSource) Endpoints() ([]*endpoint.Endpoint, error) {
	services, err := sc.serviceInformer.Lister().Services(sc.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}
	services, err = sc.filterByAnnotations(services)
	if err != nil {
		return nil, err
	}

	// filter on service types if at least one has been provided
	if len(sc.serviceTypeFilter) > 0 {
		services = sc.filterByServiceType(services)
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
			svcEndpoints = legacyEndpointsFromService(svc, sc.compatibility)
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
		sc.setResourceLabel(svc, svcEndpoints)
		endpoints = append(endpoints, svcEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

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

	pods, err := sc.podInformer.Lister().Pods(svc.Namespace).List(selector)
	if err != nil {
		log.Errorf("List Pods of service[%s] error:%v", svc.GetName(), err)
		return endpoints
	}

	targetsByHeadlessDomain := make(map[string][]string)
	for _, v := range pods {
		headlessDomains := []string{hostname}

		if v.Spec.Hostname != "" {
			headlessDomains = append(headlessDomains, fmt.Sprintf("%s.%s", v.Spec.Hostname, hostname))
		}
		for _, headlessDomain := range headlessDomains {
			if sc.publishHostIP == true {
				log.Debugf("Generating matching endpoint %s with HostIP %s", headlessDomain, v.Status.HostIP)
				// To reduce traffice on the DNS API only add record for running Pods. Good Idea?
				if v.Status.Phase == v1.PodRunning {
					targetsByHeadlessDomain[headlessDomain] = append(targetsByHeadlessDomain[headlessDomain], v.Status.HostIP)
				} else {
					log.Debugf("Pod %s is not in running phase", v.Spec.Hostname)
				}
			} else {
				log.Debugf("Generating matching endpoint %s with PodIP %s", headlessDomain, v.Status.PodIP)
				// To reduce traffice on the DNS API only add record for running Pods. Good Idea?
				if v.Status.Phase == v1.PodRunning {
					targetsByHeadlessDomain[headlessDomain] = append(targetsByHeadlessDomain[headlessDomain], v.Status.PodIP)
				} else {
					log.Debugf("Pod %s is not in running phase", v.Spec.Hostname)
				}
			}
		}

	}

	headlessDomains := []string{}
	for headlessDomain := range targetsByHeadlessDomain {
		headlessDomains = append(headlessDomains, headlessDomain)
	}
	sort.Strings(headlessDomains)
	for _, headlessDomain := range headlessDomains {
		targets := targetsByHeadlessDomain[headlessDomain]
		if ttl.IsConfigured() {
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(headlessDomain, endpoint.RecordTypeA, ttl, targets...))
		} else {
			endpoints = append(endpoints, endpoint.NewEndpoint(headlessDomain, endpoint.RecordTypeA, targets...))
		}
	}

	return endpoints
}

func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	// Process the whole template string
	var buf bytes.Buffer
	err := sc.fqdnTemplate.Execute(&buf, svc)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on service %s: %v", svc.String(), err)
	}

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(svc.Annotations)
	hostnameList := strings.Split(strings.Replace(buf.String(), " ", "", -1), ",")
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, sc.generateEndpoints(svc, hostname, providerSpecific, setIdentifier)...)
	}

	return endpoints, nil
}

// endpointsFromService extracts the endpoints from a service object
func (sc *serviceSource) endpoints(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint
	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		providerSpecific, setIdentifier := getProviderSpecificAnnotations(svc.Annotations)
		hostnameList := getHostnamesFromAnnotations(svc.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, sc.generateEndpoints(svc, hostname, providerSpecific, setIdentifier)...)
		}
	}
	return endpoints
}

// filterByAnnotations filters a list of services by a given annotation selector.
func (sc *serviceSource) filterByAnnotations(services []*v1.Service) ([]*v1.Service, error) {
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
		return services, nil
	}

	filteredList := []*v1.Service{}

	for _, service := range services {
		// convert the service's annotations to an equivalent label selector
		annotations := labels.Set(service.Annotations)

		// include service if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, service)
		}
	}

	return filteredList, nil
}

// filterByServiceType filters services according their types
func (sc *serviceSource) filterByServiceType(services []*v1.Service) []*v1.Service {
	filteredList := []*v1.Service{}
	for _, service := range services {
		// Check if the service is of the given type or not
		if _, ok := sc.serviceTypeFilter[string(service.Spec.Type)]; ok {
			filteredList = append(filteredList, service)
		}
	}

	return filteredList
}

func (sc *serviceSource) setResourceLabel(service *v1.Service, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("service/%s/%s", service.Namespace, service.Name)
	}
}

func (sc *serviceSource) generateEndpoints(svc *v1.Service, hostname string, providerSpecific endpoint.ProviderSpecific, setIdentifier string) []*endpoint.Endpoint {
	hostname = strings.TrimSuffix(hostname, ".")
	ttl, err := getTTLFromAnnotations(svc.Annotations)
	if err != nil {
		log.Warn(err)
	}

	epA := &endpoint.Endpoint{
		RecordTTL:  ttl,
		RecordType: endpoint.RecordTypeA,
		Labels:     endpoint.NewLabels(),
		Targets:    make(endpoint.Targets, 0, defaultTargetsCapacity),
		DNSName:    hostname,
	}

	epCNAME := &endpoint.Endpoint{
		RecordTTL:  ttl,
		RecordType: endpoint.RecordTypeCNAME,
		Labels:     endpoint.NewLabels(),
		Targets:    make(endpoint.Targets, 0, defaultTargetsCapacity),
		DNSName:    hostname,
	}

	var endpoints []*endpoint.Endpoint
	var targets endpoint.Targets

	switch svc.Spec.Type {
	case v1.ServiceTypeLoadBalancer:
		targets = append(targets, extractLoadBalancerTargets(svc)...)
	case v1.ServiceTypeClusterIP:
		if sc.publishInternal {
			targets = append(targets, extractServiceIps(svc)...)
		}
		if svc.Spec.ClusterIP == v1.ClusterIPNone {
			endpoints = append(endpoints, sc.extractHeadlessEndpoints(svc, hostname, ttl)...)
		}
	case v1.ServiceTypeNodePort:
		// add the nodeTargets and extract an SRV endpoint
		targets, err = sc.extractNodePortTargets(svc)
		if err != nil {
			log.Errorf("Unable to extract targets from service %s/%s error: %v", svc.Namespace, svc.Name, err)
			return endpoints
		}
		endpoints = append(endpoints, sc.extractNodePortEndpoints(svc, targets, hostname, ttl)...)
	case v1.ServiceTypeExternalName:
		targets = append(targets, extractServiceExternalName(svc)...)
	}

	for _, t := range targets {
		if suitableType(t) == endpoint.RecordTypeA {
			epA.Targets = append(epA.Targets, t)
		}
		if suitableType(t) == endpoint.RecordTypeCNAME {
			epCNAME.Targets = append(epCNAME.Targets, t)
		}
	}

	if len(epA.Targets) > 0 {
		endpoints = append(endpoints, epA)
	}
	if len(epCNAME.Targets) > 0 {
		endpoints = append(endpoints, epCNAME)
	}
	for _, endpoint := range endpoints {
		endpoint.ProviderSpecific = providerSpecific
		endpoint.SetIdentifier = setIdentifier
	}
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
	return endpoint.Targets{svc.Spec.ExternalName}
}

func extractLoadBalancerTargets(svc *v1.Service) endpoint.Targets {
	var targets endpoint.Targets

	// Create a corresponding endpoint for each configured external entrypoint.
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			targets = append(targets, lb.IP)
		}
		if lb.Hostname != "" {
			targets = append(targets, lb.Hostname)
		}
	}

	return targets
}

func (sc *serviceSource) extractNodePortTargets(svc *v1.Service) (endpoint.Targets, error) {
	var (
		internalIPs endpoint.Targets
		externalIPs endpoint.Targets
		nodes       []*v1.Node
		err         error
	)

	switch svc.Spec.ExternalTrafficPolicy {
	case v1.ServiceExternalTrafficPolicyTypeLocal:
		labelSelector, err := metav1.ParseToLabelSelector(labels.Set(svc.Spec.Selector).AsSelectorPreValidated().String())
		if err != nil {
			return nil, err
		}
		selector, err := metav1.LabelSelectorAsSelector(labelSelector)
		if err != nil {
			return nil, err
		}
		pods, err := sc.podInformer.Lister().Pods(svc.Namespace).List(selector)
		if err != nil {
			return nil, err
		}

		for _, v := range pods {
			if v.Status.Phase == v1.PodRunning {
				node, err := sc.nodeInformer.Lister().Get(v.Spec.NodeName)
				if err != nil {
					log.Debugf("Unable to find node where Pod %s is running", v.Spec.Hostname)
					continue
				}
				nodes = append(nodes, node)
			}
		}
	default:
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
			}
		}
	}

	if len(externalIPs) > 0 {
		return externalIPs, nil
	}

	return internalIPs, nil
}

func (sc *serviceSource) extractNodePortEndpoints(svc *v1.Service, nodeTargets endpoint.Targets, hostname string, ttl endpoint.TTL) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	for _, port := range svc.Spec.Ports {
		if port.NodePort > 0 {
			// build a target with a priority of 0, weight of 0, and pointing the given port on the given host
			target := fmt.Sprintf("0 50 %d %s", port.NodePort, hostname)

			// figure out the portname
			portName := port.Name
			if portName == "" {
				portName = fmt.Sprintf("%d", port.NodePort)
			}

			// figure out the protocol
			protocol := strings.ToLower(string(port.Protocol))
			if protocol == "" {
				protocol = "tcp"
			}

			recordName := fmt.Sprintf("_%s._%s.%s", portName, protocol, hostname)

			var ep *endpoint.Endpoint
			if ttl.IsConfigured() {
				ep = endpoint.NewEndpointWithTTL(recordName, endpoint.RecordTypeSRV, ttl, target)
			} else {
				ep = endpoint.NewEndpoint(recordName, endpoint.RecordTypeSRV, target)
			}

			endpoints = append(endpoints, ep)
		}
	}

	return endpoints
}
