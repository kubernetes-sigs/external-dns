/*
Copyright 2025 The Kubernetes Authors.

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
	"context"
	"fmt"
	"sort"
	"strings"
	"text/template"
	"unicode"

	nomad "github.com/hashicorp/nomad/api"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/fqdn"
)

const (
	tagPrefix = "external-dns"
)

// nomadServiceSource is an implementation of Source for Nomad services.
type nomadServiceSource struct {
	client    *nomad.Client
	namespace string

	fqdnTemplate             *template.Template
	combineFQDNAnnotation    bool
	ignoreHostnameAnnotation bool
}

// NewNomadServiceSource creates a new nomadSource.
func NewNomadServiceSource(ctx context.Context, nomadClient *nomad.Client, namespace, fqdnTemplate string, combineFqdnAnnotation bool, ignoreHostnameAnnotation bool) (Source, error) {
	tmpl, err := fqdn.ParseTemplate(fqdnTemplate)
	if err != nil {
		return nil, err
	}

	return &nomadServiceSource{
		client:                   nomadClient,
		namespace:                namespace,
		fqdnTemplate:             tmpl,
		combineFQDNAnnotation:    combineFqdnAnnotation,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
	}, nil
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (ns *nomadServiceSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	namespace := ns.namespace
	if namespace == "" {
		namespace = "*"
	}

	opts := &nomad.QueryOptions{Namespace: namespace}
	opts = opts.WithContext(ctx)

	svcLists, _, err := ns.client.Services().List(opts)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, svcList := range svcLists {
		for _, svc := range svcList.Services {
			annotations := ns.tagsToAnnotations(svc.Tags)

			controller, ok := annotations[controllerAnnotationKey]
			if ok && controller != controllerAnnotationValue {
				log.Debugf("Skipping service %s/%s because controller value does not match, found: %s, required: %s",
					svcList.Namespace, svc.ServiceName, controller, controllerAnnotationValue)
				continue
			}

			svcEndpoints, err := ns.endpoints(ctx, svcList.Namespace, svc.ServiceName, annotations)
			if err != nil {
				return nil, err
			}

			// apply template if none of the above is found
			if (ns.combineFQDNAnnotation || len(svcEndpoints) == 0) && ns.fqdnTemplate != nil {
				sEndpoints, err := ns.endpointsFromTemplate(ctx, svcList.Namespace, svc.ServiceName, annotations)

				if err != nil {
					return nil, err
				}

				if ns.combineFQDNAnnotation {
					svcEndpoints = append(svcEndpoints, sEndpoints...)
				} else {
					svcEndpoints = sEndpoints
				}
			}

			if len(svcEndpoints) == 0 {
				log.Debugf("No endpoints could be generated from service %s/%s", svcList.Namespace, svc.ServiceName)
				continue
			}

			log.Debugf("Endpoints generated from service: %s/%s: %v", svcList.Namespace, svc.ServiceName, svcEndpoints)
			ns.setResourceLabel(svcList.Namespace, svc.ServiceName, svcEndpoints)
			endpoints = append(endpoints, svcEndpoints...)
		}
	}

	// this sorting is required to make merging work.
	// after we merge endpoints that have same DNS, we want to ensure that we end up with the same service being an "owner"
	// of all those records, as otherwise each time we update, we will end up with a different service that gets data merged in
	// and that will cause external-dns to recreate dns record due to different service owner in TXT record.
	// if new service is added or old one removed, that might cause existing record to get re-created due to potentially new
	// owner being selected. Which is fine, since it shouldn't happen often and shouldn't cause any disruption.
	if len(endpoints) > 1 {
		sort.Slice(endpoints, func(i, j int) bool {
			return endpoints[i].Labels[endpoint.ResourceLabelKey] < endpoints[j].Labels[endpoint.ResourceLabelKey]
		})
		// Use stable sort to not disrupt the order of services
		sort.SliceStable(endpoints, func(i, j int) bool {
			if endpoints[i].DNSName != endpoints[j].DNSName {
				return endpoints[i].DNSName < endpoints[j].DNSName
			}
			return endpoints[i].RecordType < endpoints[j].RecordType
		})
		mergedEndpoints := []*endpoint.Endpoint{}
		mergedEndpoints = append(mergedEndpoints, endpoints[0])
		for i := 1; i < len(endpoints); i++ {
			lastMergedEndpoint := len(mergedEndpoints) - 1
			if mergedEndpoints[lastMergedEndpoint].DNSName == endpoints[i].DNSName &&
				mergedEndpoints[lastMergedEndpoint].RecordType == endpoints[i].RecordType &&
				mergedEndpoints[lastMergedEndpoint].RecordType != endpoint.RecordTypeCNAME && // It is against RFC-1034 for CNAME records to have multiple targets, so skip merging
				mergedEndpoints[lastMergedEndpoint].SetIdentifier == endpoints[i].SetIdentifier &&
				mergedEndpoints[lastMergedEndpoint].RecordTTL == endpoints[i].RecordTTL {
				mergedEndpoints[lastMergedEndpoint].Targets = append(mergedEndpoints[lastMergedEndpoint].Targets, endpoints[i].Targets[0])
			} else {
				mergedEndpoints = append(mergedEndpoints, endpoints[i])
			}

			if mergedEndpoints[lastMergedEndpoint].DNSName == endpoints[i].DNSName &&
				mergedEndpoints[lastMergedEndpoint].RecordType == endpoints[i].RecordType &&
				mergedEndpoints[lastMergedEndpoint].RecordType == endpoint.RecordTypeCNAME {
				log.Debugf("CNAME %s with multiple targets found", endpoints[i].DNSName)
			}
		}
		endpoints = mergedEndpoints
	}

	for _, ep := range endpoints {
		sort.Sort(ep.Targets)
	}

	return endpoints, nil
}

func (ns *nomadServiceSource) tagsToAnnotations(tags []string) map[string]string {
	annotations := make(map[string]string, len(tags))
	for _, tag := range tags {
		if strings.HasPrefix(tag, tagPrefix) {
			if parts := strings.SplitN(tag, "=", 2); len(parts) == 2 {
				left, right := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
				key := "external-dns.alpha.kubernetes.io/" + strings.TrimPrefix(left, tagPrefix+".")
				annotations[key] = right
			}
		}
	}
	return annotations
}

func (ns *nomadServiceSource) endpoints(ctx context.Context, namespace string, serviceName string, annotations map[string]string) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	if !ns.ignoreHostnameAnnotation {
		providerSpecific, setIdentifier := getProviderSpecificAnnotations(annotations)

		hostnameList := getHostnamesFromAnnotations(annotations)
		for _, hostname := range hostnameList {
			hnEndpoints, err := ns.generateEndpoints(ctx, namespace, serviceName, annotations, hostname, providerSpecific, setIdentifier)
			if err != nil {
				return nil, err
			}

			endpoints = append(endpoints, hnEndpoints...)
		}
	}

	return endpoints, nil
}

func (ns *nomadServiceSource) endpointsFromTemplate(ctx context.Context, namespace string, serviceName string, annotations map[string]string) ([]*endpoint.Endpoint, error) {
	hostnames, err := ns.execTemplate(ns.fqdnTemplate, nomadServiceMetadata{Namespace: namespace, Name: serviceName})
	if err != nil {
		return nil, err
	}

	providerSpecific, setIdentifier := getProviderSpecificAnnotations(annotations)

	var endpoints []*endpoint.Endpoint
	for _, hostname := range hostnames {
		hnEndpoints, err := ns.generateEndpoints(ctx, namespace, serviceName, annotations, hostname, providerSpecific, setIdentifier)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, hnEndpoints...)
	}

	return endpoints, nil
}

type nomadServiceMetadata struct {
	Namespace string
	Name      string
}

func (ns *nomadServiceSource) execTemplate(tmpl *template.Template, obj nomadServiceMetadata) (hostnames []string, err error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		return nil, fmt.Errorf("failed to apply template on service %s/%s: %w", obj.Namespace, obj.Name, err)
	}
	for _, name := range strings.Split(buf.String(), ",") {
		name = strings.TrimFunc(name, unicode.IsSpace)
		name = strings.TrimSuffix(name, ".")
		hostnames = append(hostnames, name)
	}
	return hostnames, nil
}

func (ns *nomadServiceSource) generateEndpoints(ctx context.Context, namespace string, serviceName string, annotations map[string]string, hostname string, providerSpecific endpoint.ProviderSpecific, setIdentifier string) (endpoints []*endpoint.Endpoint, _ error) {
	hostname = strings.TrimSuffix(hostname, ".")

	resource := fmt.Sprintf("service/%s/%s", namespace, serviceName)

	ttl := getTTLFromAnnotations(annotations, resource)

	targets := getTargetsFromTargetAnnotation(annotations)

	if len(targets) == 0 {
		opts := &nomad.QueryOptions{Namespace: namespace}
		opts = opts.WithContext(ctx)

		svcRegs, _, err := ns.client.Services().Get(serviceName, opts)
		if err != nil {
			return nil, err
		}

		// Collect unique service addresses
		svcAddrs := make(map[string]struct{})
		for _, svcReg := range svcRegs {
			svcAddrs[svcReg.Address] = struct{}{}
		}

		for addr := range svcAddrs {
			targets = append(targets, addr)
		}
	}

	endpoints = append(endpoints, endpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)

	return endpoints, nil
}

func (sc *nomadServiceSource) setResourceLabel(namespace string, serviceName string, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("service/%s/%s", namespace, serviceName)
	}
}

func (ns *nomadServiceSource) AddEventHandler(ctx context.Context, handler func()) {
	// TODO: Implement event listener logic
}
