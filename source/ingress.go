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
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// ingressSource is an implementation of Source for Kubernetes ingress objects.
// Ingress implementation will use the spec.rules.host value for the hostname
// Use targetAnnotationKey to add an additional Endpoint. (useful if the ingress controller does not update)
type ingressSource struct {
	client       kubernetes.Interface
	namespace    string
	fqdnTemplate *template.Template
}

// NewIngressSource creates a new ingressSource with the given config.
func NewIngressSource(kubeClient kubernetes.Interface, namespace, fqdnTemplate string) (Source, error) {
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

	return &ingressSource{
		client:       kubeClient,
		namespace:    namespace,
		fqdnTemplate: tmpl,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all ingress resources on all namespaces
func (sc *ingressSource) Endpoints() ([]*endpoint.Endpoint, error) {
	ingresses, err := sc.client.Extensions().Ingresses(sc.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, ing := range ingresses.Items {
		// Check controller annotation to see if we are responsible.
		controller, ok := ing.Annotations[controllerAnnotationKey]
		if ok && controller != controllerAnnotationValue {
			log.Debugf("Skipping ingress %s/%s because controller value does not match, found: %s, required: %s",
				ing.Namespace, ing.Name, controller, controllerAnnotationValue)
			continue
		}

		ingEndpoints := endpointsFromIngress(&ing)

		// apply template if host is missing on ingress
		if len(ingEndpoints) == 0 && sc.fqdnTemplate != nil {
			ingEndpoints, err = sc.endpointsFromTemplate(&ing)
			if err != nil {
				return nil, err
			}
		}

		if len(ingEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from ingress %s/%s", ing.Namespace, ing.Name)
			continue
		}

		log.Debugf("Endpoints generated from ingress: %s/%s: %v", ing.Namespace, ing.Name, ingEndpoints)
		endpoints = append(endpoints, ingEndpoints...)
	}

	return endpoints, nil
}

func (sc *ingressSource) endpointsFromTemplate(ing *v1beta1.Ingress) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	var buf bytes.Buffer
	err := sc.fqdnTemplate.Execute(&buf, ing)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on ingress %s: %v", ing.String(), err)
	}

	hostname := buf.String()

	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, exists := ing.Annotations[targetAnnotationKey]
	if exists {
		// splits the hostname annotation and removes the trailing periods
		targetsList := strings.Split(strings.Replace(targetAnnotation, " ", "", -1), ",")
		for _, targetHostname := range targetsList {
			targetHostname = strings.TrimSuffix(targetHostname, ".")
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, targetHostname, ""))
		}
	}

	for _, lb := range ing.Status.LoadBalancer.Ingress {
		if lb.IP != "" {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, lb.IP, ""))
		}
		if lb.Hostname != "" {
			endpoints = append(endpoints, endpoint.NewEndpoint(hostname, lb.Hostname, ""))
		}
	}

	return endpoints, nil
}

// endpointsFromIngress extracts the endpoints from ingress object
func endpointsFromIngress(ing *v1beta1.Ingress) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	for _, rule := range ing.Spec.Rules {
		if rule.Host == "" {
			continue
		}

		// Get the desired hostname of the ingress from the annotation.
		targetAnnotation, exists := ing.Annotations[targetAnnotationKey]
		if exists {
			// splits the hostname annotation and removes the trailing periods
			targetsList := strings.Split(strings.Replace(targetAnnotation, " ", "", -1), ",")
			for _, targetHostname := range targetsList {
				targetHostname = strings.TrimSuffix(targetHostname, ".")
				endpoints = append(endpoints, endpoint.NewEndpoint(rule.Host, targetHostname, ""))
			}
		}

		for _, lb := range ing.Status.LoadBalancer.Ingress {
			if lb.IP != "" {
				endpoints = append(endpoints, endpoint.NewEndpoint(rule.Host, lb.IP, ""))
			}
			if lb.Hostname != "" {
				endpoints = append(endpoints, endpoint.NewEndpoint(rule.Host, lb.Hostname, ""))
			}
		}
	}

	return endpoints
}
