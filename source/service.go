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
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// serviceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' entrypoints it will return a corresponding
// Endpoint object.
type serviceSource struct {
	client    kubernetes.Interface
	namespace string
	// process Services with legacy annotations
	compatibility   string
	fqdnTemplate    *template.Template
	publishInternal bool
}

// NewServiceSource creates a new serviceSource with the given config.
func NewServiceSource(kubeClient kubernetes.Interface, namespace, fqdnTemplate, compatibility string, publishInternal bool) (Source, error) {
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

	return &serviceSource{
		client:          kubeClient,
		namespace:       namespace,
		compatibility:   compatibility,
		fqdnTemplate:    tmpl,
		publishInternal: publishInternal,
	}, nil
}

// Endpoints returns endpoint objects for each service that should be processed.
func (sc *serviceSource) Endpoints() ([]*endpoint.Endpoint, error) {
	services, err := sc.client.CoreV1().Services(sc.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, svc := range services.Items {
		// Check controller annotation to see if we are responsible.
		controller, ok := svc.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping service %s/%s because controller value does not match, found: %s, required: %s",
				svc.Namespace, svc.Name, controller, controllerAnnotationValue)
			continue
		}

		svcEndpoints := sc.endpoints(&svc)

		// process legacy annotations if no endpoints were returned and compatibility mode is enabled.
		if len(svcEndpoints) == 0 && sc.compatibility != "" {
			svcEndpoints = legacyEndpointsFromService(&svc, sc.compatibility)
		}

		// apply template if none of the above is found
		if len(svcEndpoints) == 0 && sc.fqdnTemplate != nil {
			svcEndpoints, err = sc.endpointsFromTemplate(&svc)
			if err != nil {
				return nil, err
			}
		}

		if len(svcEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from service %s/%s", svc.Namespace, svc.Name)
			continue
		}

		log.Debugf("Endpoints generated from service: %s/%s: %v", svc.Namespace, svc.Name, svcEndpoints)
		endpoints = append(endpoints, svcEndpoints...)
	}

	return endpoints, nil
}

func (sc *serviceSource) endpointsFromTemplate(svc *v1.Service) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	var buf bytes.Buffer
	err := sc.fqdnTemplate.Execute(&buf, svc)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on service %s: %v", svc.String(), err)
	}

	hostname := buf.String()

	endpoints = sc.generateEndpoints(svc, hostname)

	return endpoints, nil
}

// endpointsFromService extracts the endpoints from a service object
func (sc *serviceSource) endpoints(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// Get the desired hostname of the service from the annotation.
	hostnameAnnotation, exists := svc.Annotations[hostnameAnnotationKey]
	if !exists {
		return nil
	}

	hostnameList := strings.Split(strings.Replace(hostnameAnnotation, " ", "", -1), ",")
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, sc.generateEndpoints(svc, hostname)...)
	}

	return endpoints
}

func (sc *serviceSource) generateEndpoints(svc *v1.Service, hostname string) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	hostname = strings.TrimSuffix(hostname, ".")
	switch svc.Spec.Type {
	case v1.ServiceTypeLoadBalancer:
		endpoints = append(endpoints, extractLoadBalancerEndpoints(svc, hostname)...)
	case v1.ServiceTypeClusterIP:
		if sc.publishInternal {
			endpoints = append(endpoints, extractServiceIps(svc, hostname)...)
		}
	}
	return endpoints
}

func extractServiceIps(svc *v1.Service, hostname string) []*endpoint.Endpoint {
	ttl, err := getTTLFromAnnotations(svc.Annotations)
	if err != nil {
		log.Warn(err)
	}
	if svc.Spec.ClusterIP == v1.ClusterIPNone {
		log.Debugf("Unable to associate %s headless service with a Cluster IP", svc.Name)
		return []*endpoint.Endpoint{}
	}

	return []*endpoint.Endpoint{endpoint.NewEndpointWithTTL(hostname, svc.Spec.ClusterIP, endpoint.RecordTypeA, ttl)}
}

func extractLoadBalancerEndpoints(svc *v1.Service, hostname string) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(svc.Annotations)
	if err != nil {
		log.Warn(err)
	}
	// Create a corresponding endpoint for each configured external entrypoint.
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			//TODO(ideahitme): consider retrieving record type from resource annotation instead of empty
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(hostname, lb.IP, endpoint.RecordTypeA, ttl))
		}
		if lb.Hostname != "" {
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(hostname, lb.Hostname, endpoint.RecordTypeCNAME, ttl))
		}
	}

	return endpoints
}
