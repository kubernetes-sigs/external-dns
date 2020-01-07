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
	"time"

	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	log "github.com/sirupsen/logrus"

	istionetworking "istio.io/api/networking/v1alpha3"
	istiomodel "istio.io/istio/pilot/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/kubernetes-sigs/external-dns/endpoint"
	"k8s.io/client-go/kubernetes"
)

// istioVirtualServiceSource is an implementation of Source for Istio VirtualService objects.
// The implementation uses the spec.hosts values for the hostnames.
// Use targetAnnotationKey to explicitly set Endpoint.
type istioVirtualServiceSource struct {
	kubeClient               kubernetes.Interface
	istioClient              istiomodel.ConfigStore
	namespace                string
	annotationFilter         string
	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
	serviceInformer          coreinformers.ServiceInformer
}

// NewIstioGatewaySource creates a new istioVirtualServiceSource with the given config.
func NewIstioVirtualServiceSource(
	kubeClient kubernetes.Interface,
	istioClient istiomodel.ConfigStore,
	namespace string,
	annotationFilter string,
	fqdnTemplate string,
	combineFqdnAnnotation bool,
	ignoreHostnameAnnotation bool,
) (Source, error) {
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

	// Add default resource event handlers to properly initialize informer.
	serviceInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				log.Debug("service added")
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

	return &istioVirtualServiceSource{
		kubeClient:               kubeClient,
		istioClient:              istioClient,
		namespace:                namespace,
		annotationFilter:         annotationFilter,
		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    combineFqdnAnnotation,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
		serviceInformer:          serviceInformer,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all virtualservice resources in the source's namespace(s).
func (sc *istioVirtualServiceSource) Endpoints() ([]*endpoint.Endpoint, error) {
	configs, err := sc.istioClient.List(istiomodel.VirtualService.Type, sc.namespace)
	if err != nil {
		return nil, err
	}

	configs, err = sc.filterByAnnotations(configs)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, config := range configs {
		// Check controller annotation to see if we are responsible.
		controller, ok := config.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping virtualservice %s/%s because controller value does not match, found: %s, required: %s",
				config.Namespace, config.Name, controller, controllerAnnotationValue)
			continue
		}

		gwEndpoints, err := sc.endpointsFromVirtualServiceConfig(config)
		if err != nil {
			return nil, err
		}

		// apply template if host is missing on virtualservice
		if (sc.combineFQDNAnnotation || len(gwEndpoints) == 0) && sc.fqdnTemplate != nil {
			iEndpoints, err := sc.endpointsFromTemplate(&config)
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
			log.Debugf("No endpoints could be generated from virtualservice %s/%s", config.Namespace, config.Name)
			continue
		}

		log.Debugf("Endpoints generated from virtualservice: %s/%s: %v", config.Namespace, config.Name, gwEndpoints)
		sc.setResourceLabel(config, gwEndpoints)
		endpoints = append(endpoints, gwEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (sc *istioVirtualServiceSource) endpointsFromTemplate(config *istiomodel.Config) ([]*endpoint.Endpoint, error) {
	// Process the whole template string
	var buf bytes.Buffer
	err := sc.fqdnTemplate.Execute(&buf, config)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on istio config %s: %v", config, err)
	}

	hostnames := buf.String()

	ttl, err := getTTLFromAnnotations(config.Annotations)
	if err != nil {
		log.Warn(err)
	}

	var endpoints []*endpoint.Endpoint

	targets, err := sc.targetsFromVirtualServiceConfig(config)
	if err != nil {
		return endpoints, err
	}

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(config.Annotations)

	// splits the FQDN template and removes the trailing periods
	hostnameList := strings.Split(strings.Replace(hostnames, " ", "", -1), ",")
	for _, hostname := range hostnameList {
		hostname = strings.TrimSuffix(hostname, ".")
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
	}
	return endpoints, nil
}

// filterByAnnotations filters a list of configs by a given annotation selector.
func (sc *istioVirtualServiceSource) filterByAnnotations(configs []istiomodel.Config) ([]istiomodel.Config, error) {
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
		return configs, nil
	}

	filteredList := []istiomodel.Config{}

	for _, config := range configs {
		// convert the annotations to an equivalent label selector
		annotations := labels.Set(config.Annotations)

		// include if the annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, config)
		}
	}

	return filteredList, nil
}

func (sc *istioVirtualServiceSource) setResourceLabel(config istiomodel.Config, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("virtualservice/%s/%s", config.Namespace, config.Name)
	}
}

func (sc *istioVirtualServiceSource) targetsFromGatewayConfig(config *istiomodel.Config) (targets endpoint.Targets, err error) {
	gateway := config.Spec.(*istionetworking.Gateway)

	services, err := sc.serviceInformer.Lister().Services(config.Namespace).List(labels.Everything())
	if err != nil {
		log.Error(err)
		return
	}

	for _, service := range services {
		if !labels.Equals(service.Spec.Selector, gateway.Selector) {
			// Only consider services which have the same selector as the Gateway CR.
			// Use the target annotation to override if this picks the wrong service.
			continue
		}
		for _, lb := range service.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				targets = append(targets, lb.IP)
			}
			if lb.Hostname != "" {
				targets = append(targets, lb.Hostname)
			}
		}
	}

	return
}

func (sc *istioVirtualServiceSource) targetsFromVirtualServiceConfig(config *istiomodel.Config) (targets endpoint.Targets, err error) {
	targets = getTargetsFromTargetAnnotation(config.Annotations)
	if len(targets) > 0 {
		return targets, nil
	}

	virtualservice := config.Spec.(*istionetworking.VirtualService)

	for _, gateway := range virtualservice.Gateways {
		if gateway == "" || gateway == "mesh" {
			// This refers to "all sidecars in the mesh"; ignore.
			continue
		}
		namespace, name, err := parseIngressGateway(gateway)
		if err != nil {
			log.Debugf("Failed parsing gateway %s of virtualservice %s/%s", gateway, config.Namespace, config.Name)
			continue
		}

		gwconfig := sc.istioClient.Get(istiomodel.Gateway.Type, name, namespace)
		if err != nil {
			log.Debugf("Failed reading gateway %s/%s of virtualservice %s/%s", namespace, name, config.Namespace, config.Name)
			continue
		}

		// TODO check if this virtualservice can be bound to this gateway
		// (i.e. if gateway doesn't have a namespace restriction in its hosts list)

		targets = getTargetsFromTargetAnnotation(gwconfig.Annotations)
		if len(targets) > 0 {
			break
		}

		targets, err = sc.targetsFromGatewayConfig(gwconfig)
		if err != nil {
			return targets, err
		}

		if len(targets) > 0 {
			break
		}
	}

	return targets, err
}

// endpointsFromGatewayConfig extracts the endpoints from an Istio Gateway Config object
func (sc *istioVirtualServiceSource) endpointsFromVirtualServiceConfig(config istiomodel.Config) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(config.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets, err := sc.targetsFromVirtualServiceConfig(&config)
	if err != nil {
		return endpoints, err
	}

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(config.Annotations)

	virtualservice := config.Spec.(*istionetworking.VirtualService)

	for _, host := range virtualservice.Hosts {
		if host == "" || host == "*" {
			continue
		}

		parts := strings.Split(host, "/")

		// If the input hostname is of the form my-namespace/foo.bar.com, remove the namespace
		// before appending it to the list of endpoints to create
		if len(parts) == 2 {
			host = parts[1]
		}

		endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(config.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
		}
	}

	return endpoints, nil
}
