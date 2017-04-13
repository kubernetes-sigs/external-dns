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
	"html/template"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"

	"bytes"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"strings"
)

// serviceSource is an implementation of Source for Kubernetes service objects.
// It will find all services that are under our jurisdiction, i.e. annotated
// desired hostname and matching or no controller annotation. For each of the
// matched services' external entrypoints it will return a corresponding
// Endpoint object.
type serviceSource struct {
	client    kubernetes.Interface
	namespace string
	// set to true to process Services with legacy annotations
	compatibility bool
	fqdntemplate  string
}

// NewServiceSource creates a new serviceSource with the given client and namespace scope.
func NewServiceSource(client kubernetes.Interface, namespace string, compatibility bool, fqdntemplate string) Source {
	return &serviceSource{
		client:        client,
		namespace:     namespace,
		compatibility: compatibility,
		fqdntemplate:  fqdntemplate,
	}
}

// Endpoints returns endpoint objects for each service that should be processed.
func (sc *serviceSource) Endpoints() ([]*endpoint.Endpoint, error) {
	services, err := sc.client.CoreV1().Services(sc.namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, svc := range services.Items {
		svcEndpoints := endpointsFromService(&svc, sc.fqdntemplate)

		// process legacy annotations if no endpoints were returned and compatibility mode is enabled.
		if len(svcEndpoints) == 0 && sc.compatibility {
			svcEndpoints = legacyEndpointsFromService(&svc)
		}

		if len(svcEndpoints) != 0 {
			endpoints = append(endpoints, svcEndpoints...)
		}
	}

	return endpoints, nil
}

// endpointsFromService extracts the endpoints from a service object
func endpointsFromService(svc *v1.Service, fqdntemplate string) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// Check controller annotation to see if we are responsible.
	controller, exists := svc.Annotations[controllerAnnotationKey]
	if exists && controller != controllerAnnotationValue {
		return nil
	}

	// Get the desired hostname of the service from the annotation.
	hostname, exists := svc.Annotations[hostnameAnnotationKey]
	if !exists {
		tmpl, err := template.New("endpoint").Funcs(template.FuncMap{
			"trimPrefix": strings.TrimPrefix,
		}).Parse(fqdntemplate)
		if err != nil {
			return nil
		}

		var buf bytes.Buffer

		tmpl.Execute(&buf, svc)
		hostname = buf.String()
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
