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

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/external-dns/endpoint"
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

// NewIstioVirtualServiceSource creates a new istioVirtualServiceSource with the given config.
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
		return serviceInformer.Informer().HasSynced(), nil
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
// Retrieves all VirtualService resources in the source's namespace(s).
func (sc *istioVirtualServiceSource) Endpoints() ([]*endpoint.Endpoint, error) {
	configs, err := sc.istioClient.List(istiomodel.VirtualService.Type, sc.namespace)
	if err != nil {
		return nil, err
	}

	configs, err = sc.filterByAnnotations(configs)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint

	for _, config := range configs {
		// Check controller annotation to see if we are responsible.
		controller, ok := config.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping VirtualService %s/%s because controller value does not match, found: %s, required: %s",
				config.Namespace, config.Name, controller, controllerAnnotationValue)
			continue
		}

		gwEndpoints, err := sc.endpointsFromVirtualServiceConfig(config)
		if err != nil {
			return nil, err
		}

		// apply template if host is missing on VirtualService
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
			log.Debugf("No endpoints could be generated from VirtualService %s/%s", config.Namespace, config.Name)
			continue
		}

		log.Debugf("Endpoints generated from VirtualService: %s/%s: %v", config.Namespace, config.Name, gwEndpoints)
		sc.setResourceLabel(config, gwEndpoints)
		endpoints = append(endpoints, gwEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (sc *istioVirtualServiceSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
}

func (sc *istioVirtualServiceSource) getGateway(gateway string, vsconfig *istiomodel.Config) *istiomodel.Config {
	if gateway == "" || gateway == "mesh" {
		// This refers to "all sidecars in the mesh"; ignore.
		return nil
	}

	namespace, name, err := parseGateway(gateway)
	if err != nil {
		log.Debugf("Failed parsing gateway %s of VirtualService %s/%s", gateway, vsconfig.Namespace, vsconfig.Name)
		return nil
	}
	if namespace == "" {
		namespace = vsconfig.Namespace
	}

	gwconfig := sc.istioClient.Get(istiomodel.Gateway.Type, name, namespace)
	if gwconfig == nil {
		log.Debugf("Gateway %s referenced by VirtualService %s/%s not found", gateway, vsconfig.Namespace, vsconfig.Name)
		return nil
	}

	return gwconfig
}

func (sc *istioVirtualServiceSource) endpointsFromTemplate(config *istiomodel.Config) ([]*endpoint.Endpoint, error) {
	// Process the whole template string
	var buf bytes.Buffer
	err := sc.fqdnTemplate.Execute(&buf, config)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on istio config %s: %v", config, err)
	}

	hostnamesTemplate := buf.String()

	ttl, err := getTTLFromAnnotations(config.Annotations)
	if err != nil {
		log.Warn(err)
	}

	var endpoints []*endpoint.Endpoint

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(config.Annotations)

	// splits the FQDN template and removes the trailing periods
	hostnames := strings.Split(strings.Replace(hostnamesTemplate, " ", "", -1), ",")
	for _, hostname := range hostnames {
		hostname = strings.TrimSuffix(hostname, ".")
		targets, err := sc.targetsFromVirtualServiceConfig(config, hostname)
		if err != nil {
			return endpoints, err
		}
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

	var filteredList []istiomodel.Config

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

func (sc *istioVirtualServiceSource) targetsFromVirtualServiceConfig(vsconfig *istiomodel.Config, vsHost string) ([]string, error) {
	var targets []string
	// for each host we need to iterate through the gateways because each host might match for only one of the gateways
	for _, gateway := range vsconfig.Spec.(*istionetworking.VirtualService).Gateways {
		gwconfig := sc.getGateway(gateway, vsconfig)
		if gwconfig == nil {
			continue
		}
		if !virtualServiceBindsToGateway(vsconfig, gwconfig, vsHost) {
			continue
		}
		tgs, err := targetsFromGatewayConfig(gwconfig, sc.serviceInformer)
		if err != nil {
			return targets, err
		}
		targets = append(targets, tgs...)
	}

	return targets, nil
}

// endpointsFromVirtualServiceConfig extracts the endpoints from an Istio VirtualService Config object
func (sc *istioVirtualServiceSource) endpointsFromVirtualServiceConfig(vsconfig istiomodel.Config) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(vsconfig.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targetsFromAnnotation := getTargetsFromTargetAnnotation(vsconfig.Annotations)

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(vsconfig.Annotations)

	for _, host := range vsconfig.Spec.(*istionetworking.VirtualService).Hosts {
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
			targets, err = sc.targetsFromVirtualServiceConfig(&vsconfig, host)
			if err != nil {
				return endpoints, err
			}
		}

		endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific, setIdentifier)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(vsconfig.Annotations)
		for _, hostname := range hostnameList {
			targets := targetsFromAnnotation
			if len(targets) == 0 {
				targets, err = sc.targetsFromVirtualServiceConfig(&vsconfig, hostname)
				if err != nil {
					return endpoints, err
				}
			}
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier)...)
		}
	}

	return endpoints, nil
}

// checks if the given VirtualService should actually bind to the given gateway
// see requirements here: https://istio.io/docs/reference/config/networking/gateway/#Server
func virtualServiceBindsToGateway(vsconfig, gwconfig *istiomodel.Config, vsHost string) bool {
	vs := vsconfig.Spec.(*istionetworking.VirtualService)

	isValid := false
	if len(vs.ExportTo) == 0 {
		isValid = true
	} else {
		for _, ns := range vs.ExportTo {
			if ns == "*" || ns == gwconfig.Namespace || (ns == "." && gwconfig.Namespace == vsconfig.Namespace) {
				isValid = true
			}
		}
	}
	if !isValid {
		return false
	}

	gw := gwconfig.Spec.(*istionetworking.Gateway)
	for _, server := range gw.Servers {
		for _, host := range server.Hosts {
			namespace := "*"
			parts := strings.Split(host, "/")
			if len(parts) == 2 {
				namespace = parts[0]
				host = parts[1]
			} else if len(parts) != 1 {
				log.Debugf("Gateway %s/%s has invalid host %s", gwconfig.Namespace, gwconfig.Name, host)
				continue
			}

			if namespace == "*" || namespace == vsconfig.Namespace || (namespace == "." && vsconfig.Namespace == gwconfig.Namespace) {
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

func parseGateway(gateway string) (namespace, name string, err error) {
	parts := strings.Split(gateway, "/")
	if len(parts) == 2 {
		namespace, name = parts[0], parts[1]
	} else if len(parts) == 1 {
		name = parts[0]
	} else {
		err = fmt.Errorf("invalid gateway name (name or namespace/name) found '%v'", gateway)
	}

	return
}
