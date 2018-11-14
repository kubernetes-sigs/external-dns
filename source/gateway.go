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

	log "github.com/sirupsen/logrus"

	istionetworking "istio.io/api/networking/v1alpha3"
	istiomodel "istio.io/istio/pilot/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"k8s.io/client-go/kubernetes"
)

// gatewaySource is an implementation of Source for Istio Gateway objects.
// The gateway implementation uses the spec.servers.hosts values for the hostnames.
// Use targetAnnotationKey to explicitly set Endpoint.
type gatewaySource struct {
	kubeClient              kubernetes.Interface
	istioClient             istiomodel.ConfigStore
	istioNamespace          string
	istioIngressGatewayName string
	namespace               string
	annotationFilter        string
	fqdnTemplate            *template.Template
	combineFQDNAnnotation   bool
}

// NewIstioGatewaySource creates a new gatewaySource with the given config.
func NewIstioGatewaySource(
	kubeClient kubernetes.Interface,
	istioClient istiomodel.ConfigStore,
	istioIngressGateway string,
	namespace string,
	annotationFilter string,
	fqdnTemplate string,
	combineFqdnAnnotation bool,
) (Source, error) {
	var (
		tmpl *template.Template
		err  error
	)
	istioNamespace, istioIngressGatewayName, err := parseIngressGateway(istioIngressGateway)
	if err != nil {
		return nil, err
	}

	if fqdnTemplate != "" {
		tmpl, err = template.New("endpoint").Funcs(template.FuncMap{
			"trimPrefix": strings.TrimPrefix,
		}).Parse(fqdnTemplate)
		if err != nil {
			return nil, err
		}
	}

	return &gatewaySource{
		kubeClient:              kubeClient,
		istioClient:             istioClient,
		istioNamespace:          istioNamespace,
		istioIngressGatewayName: istioIngressGatewayName,
		namespace:               namespace,
		annotationFilter:        annotationFilter,
		fqdnTemplate:            tmpl,
		combineFQDNAnnotation:   combineFqdnAnnotation,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all gateway resources in the source's namespace(s).
func (sc *gatewaySource) Endpoints() ([]*endpoint.Endpoint, error) {
	configs, err := sc.istioClient.List(istiomodel.Gateway.Type, sc.namespace)
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
			log.Debugf("Skipping gateway %s/%s because controller value does not match, found: %s, required: %s",
				config.Namespace, config.Name, controller, controllerAnnotationValue)
			continue
		}

		gwEndpoints, err := sc.endpointsFromGatewayConfig(config)
		if err != nil {
			return nil, err
		}

		// apply template if host is missing on gateway
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
			log.Debugf("No endpoints could be generated from gateway %s/%s", config.Namespace, config.Name)
			continue
		}

		log.Debugf("Endpoints generated from gateway: %s/%s: %v", config.Namespace, config.Name, gwEndpoints)
		sc.setResourceLabel(config, gwEndpoints)
		endpoints = append(endpoints, gwEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (sc *gatewaySource) endpointsFromTemplate(config *istiomodel.Config) ([]*endpoint.Endpoint, error) {
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

	targets := getTargetsFromTargetAnnotation(config.Annotations)

	if len(targets) == 0 {
		targets, err = sc.targetsFromIstioIngressStatus()
		if err != nil {
			return nil, err
		}
	}

	providerSpecific := getProviderSpecificAnnotations(config.Annotations)

	var endpoints []*endpoint.Endpoint
	// splits the FQDN template and removes the trailing periods
	hostnameList := strings.Split(strings.Replace(hostnames, " ", "", -1), ",")
	for _, hostname := range hostnameList {
		hostname = strings.TrimSuffix(hostname, ".")
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific)...)
	}
	return endpoints, nil
}

// filterByAnnotations filters a list of configs by a given annotation selector.
func (sc *gatewaySource) filterByAnnotations(configs []istiomodel.Config) ([]istiomodel.Config, error) {
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

func (sc *gatewaySource) setResourceLabel(config istiomodel.Config, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("gateway/%s/%s", config.Namespace, config.Name)
	}
}

func (sc *gatewaySource) targetsFromIstioIngressStatus() (targets endpoint.Targets, err error) {
	if svc, e := sc.kubeClient.CoreV1().Services(sc.istioNamespace).Get(sc.istioIngressGatewayName, metav1.GetOptions{}); e != nil {
		err = e
	} else {
		for _, lb := range svc.Status.LoadBalancer.Ingress {
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

// endpointsFromGatewayConfig extracts the endpoints from an Istio Gateway Config object
func (sc *gatewaySource) endpointsFromGatewayConfig(config istiomodel.Config) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(config.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets := getTargetsFromTargetAnnotation(config.Annotations)

	if len(targets) == 0 {
		targets, err = sc.targetsFromIstioIngressStatus()
		if err != nil {
			return nil, err
		}
	}

	gateway := config.Spec.(*istionetworking.Gateway)

	providerSpecific := getProviderSpecificAnnotations(config.Annotations)

	for _, server := range gateway.Servers {
		for _, host := range server.Hosts {
			if host == "" {
				continue
			}
			endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific)...)
		}
	}

	hostnameList := getHostnamesFromAnnotations(config.Annotations)
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific)...)
	}

	return endpoints, nil
}

func parseIngressGateway(ingressGateway string) (namespace, name string, err error) {
	parts := strings.Split(ingressGateway, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid ingress gateway service (namespace/name) found '%v'", ingressGateway)
	} else {
		namespace, name = parts[0], parts[1]
	}

	return
}
