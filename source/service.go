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

	log "github.com/Sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// serviceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' external entrypoints it will return a corresponding
// Endpoint object.
type serviceSource struct {
	client    kubernetes.Interface
	namespace string
	// process Services with legacy annotations
	compatibility string
	fqdntemplate  *template.Template
}

// NewServiceSource creates a new serviceSource with the given client and namespace scope.
func NewServiceSource(client kubernetes.Interface, namespace, fqdntemplate string, compatibility string) (Source, error) {
	var tmpl *template.Template
	var err error
	if fqdntemplate != "" {
		tmpl, err = template.New("endpoint").Funcs(template.FuncMap{
			"trimPrefix": strings.TrimPrefix,
		}).Parse(fqdntemplate)
		if err != nil {
			return nil, err
		}
	}

	return &serviceSource{
		client:        client,
		namespace:     namespace,
		compatibility: compatibility,
		fqdntemplate:  tmpl,
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

		svcEndpoints := endpointsFromService(&svc)

		// process legacy annotations if no endpoints were returned and compatibility mode is enabled.
		if len(svcEndpoints) == 0 && sc.compatibility != "" {
			svcEndpoints = legacyEndpointsFromService(&svc, sc.compatibility)
		}

		// apply template if none of the above is found
		if len(svcEndpoints) == 0 && sc.fqdntemplate != nil {
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
	err := sc.fqdntemplate.Execute(&buf, svc)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on service %s: %v", svc.String(), err)
	}

	hostname := buf.String()
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			//TODO(ideahitme): consider retrieving record type from resource annotation instead of empty
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, lb.IP, ""))
		}
		if lb.Hostname != "" {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, lb.Hostname, ""))
		}
	}

	return endpoints, nil
}

// endpointsFromService extracts the endpoints from a service object
func endpointsFromService(svc *v1.Service) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// Get the desired hostname of the service from the annotation.
	hostname, exists := svc.Annotations[hostnameAnnotationKey]
	if !exists {
		return nil
	}

	// Create a corresponding endpoint for each configured external entrypoint.
	for _, lb := range svc.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			//TODO(ideahitme): consider retrieving record type from resource annotation instead of empty
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, lb.IP, ""))
		}
		if lb.Hostname != "" {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, lb.Hostname, ""))
		}
	}

	return endpoints
}
