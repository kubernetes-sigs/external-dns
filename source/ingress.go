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
	"regexp"
	"sort"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// ingressSource is an implementation of Source for Kubernetes ingress objects.
// Ingress implementation will use the spec.rules.host value for the hostname
// Use targetAnnotationKey to explicitly set Endpoint. (useful if the ingress
// controller does not update, or to override with alternative endpoint)
type ingressSource struct {
	client                kubernetes.Interface
	namespace             string
	annotationFilter      string
	fqdnTemplate          *template.Template
	combineFQDNAnnotation bool
}

// NewIngressSource creates a new ingressSource with the given config.
func NewIngressSource(kubeClient kubernetes.Interface, namespace, annotationFilter string, fqdnTemplate string, combineFqdnAnnotation bool) (Source, error) {
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
		client:                kubeClient,
		namespace:             namespace,
		annotationFilter:      annotationFilter,
		fqdnTemplate:          tmpl,
		combineFQDNAnnotation: combineFqdnAnnotation,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all ingress resources on all namespaces
func (sc *ingressSource) Endpoints() ([]*endpoint.Endpoint, error) {
	ingresses, err := sc.client.Extensions().Ingresses(sc.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	ingresses.Items, err = sc.filterByAnnotations(ingresses.Items)
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
		if (sc.combineFQDNAnnotation || len(ingEndpoints) == 0) && sc.fqdnTemplate != nil {
			iEndpoints, err := sc.endpointsFromTemplate(&ing)
			if err != nil {
				return nil, err
			}

			if sc.combineFQDNAnnotation {
				ingEndpoints = append(ingEndpoints, iEndpoints...)
			} else {
				ingEndpoints = iEndpoints
			}
		}

		if len(ingEndpoints) == 0 {
			log.Debugf("No endpoints could be generated from ingress %s/%s", ing.Namespace, ing.Name)
			continue
		}

		log.Debugf("Endpoints generated from ingress: %s/%s: %v", ing.Namespace, ing.Name, ingEndpoints)
		sc.setResourceLabel(ing, ingEndpoints)
		endpoints = append(endpoints, ingEndpoints...)
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

// get endpoints from optional "target" annotation
// Returns empty endpoints array if none are found.
func getTargetsFromTargetAnnotation(ing *v1beta1.Ingress) endpoint.Targets {
	var targets endpoint.Targets

	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, exists := ing.Annotations[targetAnnotationKey]
	if exists {
		// splits the hostname annotation and removes the trailing periods
		re := regexp.MustCompile(`\s`)
		targetsList := strings.Split(re.ReplaceAllString(targetAnnotation, ""), ",")
		for _, targetHostname := range targetsList {
			targetHostname = strings.TrimSuffix(targetHostname, ".")
			targets = append(targets, targetHostname)
		}
	}
	return targets
}

func (sc *ingressSource) endpointsFromTemplate(ing *v1beta1.Ingress) ([]*endpoint.Endpoint, error) {
	// Process the whole template string
	var buf bytes.Buffer
	err := sc.fqdnTemplate.Execute(&buf, ing)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on ingress %s: %v", ing.String(), err)
	}

	hostnames := buf.String()

	ttl, err := getTTLFromAnnotations(ing.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets := getTargetsFromTargetAnnotation(ing)

	if len(targets) == 0 {
		targets = targetsFromIngressStatus(ing.Status)
	}

	var endpoints []*endpoint.Endpoint
	// splits the FQDN template and removes the trailing periods
	re := regexp.MustCompile(`\s`)
	hostnameList := strings.Split(re.ReplaceAllString(hostnames, ""), ",")
	for _, hostname := range hostnameList {
		hostname = strings.TrimSuffix(hostname, ".")
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl)...)
	}
	return endpoints, nil
}

// filterByAnnotations filters a list of ingresses by a given annotation selector.
func (sc *ingressSource) filterByAnnotations(ingresses []v1beta1.Ingress) ([]v1beta1.Ingress, error) {
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
		return ingresses, nil
	}

	filteredList := []v1beta1.Ingress{}

	for _, ingress := range ingresses {
		// convert the ingress' annotations to an equivalent label selector
		annotations := labels.Set(ingress.Annotations)

		// include ingress if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList = append(filteredList, ingress)
		}
	}

	return filteredList, nil
}

func (sc *ingressSource) setResourceLabel(ingress v1beta1.Ingress, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("ingress/%s/%s", ingress.Namespace, ingress.Name)
	}
}

// endpointsFromIngress extracts the endpoints from ingress object
func endpointsFromIngress(ing *v1beta1.Ingress) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	ttl, err := getTTLFromAnnotations(ing.Annotations)
	if err != nil {
		log.Warn(err)
	}

	targets := getTargetsFromTargetAnnotation(ing)

	if len(targets) == 0 {
		targets = targetsFromIngressStatus(ing.Status)
	}

	for _, rule := range ing.Spec.Rules {
		if rule.Host == "" {
			continue
		}
		endpoints = append(endpoints, endpointsForHostname(rule.Host, targets, ttl)...)
	}

	hostnameList := getHostnamesFromAnnotations(ing.Annotations)
	for _, hostname := range hostnameList {
		endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl)...)
	}

	return endpoints
}

func endpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	var aTargets endpoint.Targets
	var cnameTargets endpoint.Targets

	for _, t := range targets {
		switch suitableType(t) {
		case endpoint.RecordTypeA:
			aTargets = append(aTargets, t)
		default:
			cnameTargets = append(cnameTargets, t)
		}
	}

	if len(aTargets) > 0 {
		epA := &endpoint.Endpoint{
			DNSName:    strings.TrimSuffix(hostname, "."),
			Targets:    aTargets,
			RecordTTL:  ttl,
			RecordType: endpoint.RecordTypeA,
			Labels:     endpoint.NewLabels(),
		}
		endpoints = append(endpoints, epA)
	}

	if len(cnameTargets) > 0 {
		epCNAME := &endpoint.Endpoint{
			DNSName:    strings.TrimSuffix(hostname, "."),
			Targets:    cnameTargets,
			RecordTTL:  ttl,
			RecordType: endpoint.RecordTypeCNAME,
			Labels:     endpoint.NewLabels(),
		}
		endpoints = append(endpoints, epCNAME)
	}

	return endpoints
}

func targetsFromIngressStatus(status v1beta1.IngressStatus) endpoint.Targets {
	var targets endpoint.Targets

	for _, lb := range status.LoadBalancer.Ingress {
		if lb.IP != "" {
			targets = append(targets, lb.IP)
		}
		if lb.Hostname != "" {
			targets = append(targets, lb.Hostname)
		}
	}

	return targets
}
