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

// virtualServiceSource is an implementation of Source for Istio VirtualService objects.
// The virtualService implementation uses the spec.hosts values for the hostnames.
// Use targetAnnotationKey to explicitly set Endpoint.
type virtualServiceSource struct {
	kubeClient                  kubernetes.Interface
	istioClient                 istiomodel.ConfigStore
	IstioIngressGatewayServices []string
	namespace                   string
	annotationFilter            string
	fqdnTemplate                *template.Template
	combineFQDNAnnotation       bool
	ignoreHostnameAnnotation    bool
}

// NewIstioVirtualServiceSource creates a new virtualServiceSource with the given config.
func NewIstioVirtualServiceSource(
	kubeClient kubernetes.Interface,
	istioClient istiomodel.ConfigStore,
	IstioIngressGatewayServices []string,
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
	for _, lbService := range IstioIngressGatewayServices {
		if _, _, err = parseIngressGateway(lbService); err != nil {
			return nil, err
		}
	}

	if fqdnTemplate != "" {
		tmpl, err = template.New("endpoint").Funcs(template.FuncMap{
			"trimPrefix": strings.TrimPrefix,
		}).Parse(fqdnTemplate)
		if err != nil {
			return nil, err
		}
	}

	return &virtualServiceSource{
		kubeClient:                  kubeClient,
		istioClient:                 istioClient,
		IstioIngressGatewayServices: IstioIngressGatewayServices,
		namespace:                   namespace,
		annotationFilter:            annotationFilter,
		fqdnTemplate:                tmpl,
		combineFQDNAnnotation:       combineFqdnAnnotation,
		ignoreHostnameAnnotation:    ignoreHostnameAnnotation,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all virtualService resources in the source's namespace(s).
func (sc *virtualServiceSource) Endpoints() ([]*endpoint.Endpoint, error) {
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
			log.Debugf("Skipping virtualService %s/%s because controller value does not match, found: %s, required: %s",
				config.Namespace, config.Name, controller, controllerAnnotationValue)
			continue
		}

		vsEndpoints, err := sc.endpointsFromVirtualServiceConfig(config)
		if err != nil {
			return nil, err
		}

		// apply template if host is missing on virtualService
		if (sc.combineFQDNAnnotation || len(vsEndpoints) == 0) && sc.fqdnTemplate != nil {
			iEndpoints, err := sc.endpointsFromTemplate(&config)
			if err != nil {
				return nil, err
			}

			if sc.combineFQDNAnnotation {
				vsEndpoints = append(vsEndpoints, iEndpoints...)
			} else {
				vsEndpoints = iEndpoints
			}
		}

		if len(vsEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from virtualService %s/%s", config.Namespace, config.Name)
			continue
		}

		log.Debugf("Endpoints generated from virtualService: %s/%s: %v", config.Namespace, config.Name, vsEndpoints)
		sc.setResourceLabel(config, vsEndpoints)
		endpoints = append(endpoints, vsEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (sc *virtualServiceSource) endpointsFromTemplate(config *istiomodel.Config) ([]*endpoint.Endpoint, error) {
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
		targets, err = sc.targetsFromIstioIngressGatewayServices()
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
func (sc *virtualServiceSource) filterByAnnotations(configs []istiomodel.Config) ([]istiomodel.Config, error) {
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

func (sc *virtualServiceSource) setResourceLabel(config istiomodel.Config, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("virtualService/%s/%s", config.Namespace, config.Name)
	}
}

func (sc *virtualServiceSource) targetsFromIstioIngressGatewayServices() (targets endpoint.Targets, err error) {
	for _, lbService := range sc.IstioIngressGatewayServices {
		lbNamespace, lbName, err := parseIngressGateway(lbService)
		if err != nil {
			return nil, err
		}
		if svc, err := sc.kubeClient.CoreV1().Services(lbNamespace).Get(lbName, metav1.GetOptions{}); err != nil {
			log.Warn(err)
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
	}

	return
}

// endpointsFromVirtualServiceConfig extracts the endpoints from an Istio VirtualService Config object
func (sc *virtualServiceSource) endpointsFromVirtualServiceConfig(config istiomodel.Config) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(config.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets := getTargetsFromTargetAnnotation(config.Annotations)

	if len(targets) == 0 {
		targets, err = sc.targetsFromIstioIngressGatewayServices()
		if err != nil {
			return nil, err
		}
	}

	virtualService := config.Spec.(*istionetworking.VirtualService)

	providerSpecific := getProviderSpecificAnnotations(config.Annotations)

	for _, host := range virtualService.Hosts {
		if host == "" {
			continue
		}
		endpoints = append(endpoints, endpointsForHostname(host, targets, ttl, providerSpecific)...)
	}

	// Skip endpoints if we do not want entries from annotations
	if !sc.ignoreHostnameAnnotation {
		hostnameList := getHostnamesFromAnnotations(config.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific)...)
		}
	}

	return endpoints, nil
}
