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
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

const (
	// The annotation used for figuring out which controller is responsible
	controllerAnnotationKey = "external-dns.alpha.kubernetes.io/controller"
	// The annotation used for defining the desired hostname
	hostnameAnnotationKey = "external-dns.alpha.kubernetes.io/hostname"
	// The annotation used for defining the desired ingress target
	targetAnnotationKey = "external-dns.alpha.kubernetes.io/target"
	// The annotation used for defining the desired DNS record TTL
	ttlAnnotationKey = "external-dns.alpha.kubernetes.io/ttl"
	// The annotation used for switching to the alias record types e. g. AWS Alias records instead of a normal CNAME
	aliasAnnotationKey = "external-dns.alpha.kubernetes.io/alias"
	// The value of the controller annotation so that we feel responsible
	controllerAnnotationValue = "dns-controller"
	// The annotation used to configure which nodes should be returned for nodeport services
	nodeLabelFilterKey = "external-dns.alpha.kubernetes.io/node-label-filter"
)

const (
	ttlMinimum = 1
	ttlMaximum = math.MaxUint32
)

// Source defines the interface Endpoint sources should implement.
type Source interface {
	Endpoints() ([]*endpoint.Endpoint, error)
}

func getTTLFromAnnotations(annotations map[string]string) (endpoint.TTL, error) {
	ttlNotConfigured := endpoint.TTL(0)
	ttlAnnotation, exists := annotations[ttlAnnotationKey]
	if !exists {
		return ttlNotConfigured, nil
	}
	ttlValue, err := strconv.ParseInt(ttlAnnotation, 10, 64)
	if err != nil {
		return ttlNotConfigured, fmt.Errorf("\"%v\" is not a valid TTL value", ttlAnnotation)
	}
	if ttlValue < ttlMinimum || ttlValue > ttlMaximum {
		return ttlNotConfigured, fmt.Errorf("TTL value must be between [%d, %d]", ttlMinimum, ttlMaximum)
	}
	return endpoint.TTL(ttlValue), nil
}

func getHostnamesFromAnnotations(annotations map[string]string) []string {
	hostnameAnnotation, exists := annotations[hostnameAnnotationKey]
	if !exists {
		return nil
	}

	return strings.Split(strings.Replace(hostnameAnnotation, " ", "", -1), ",")
}

func getAliasFromAnnotations(annotations map[string]string) bool {
	aliasAnnotation, exists := annotations[aliasAnnotationKey]
	return exists && aliasAnnotation == "true"
}

func getProviderSpecificAnnotations(annotations map[string]string) endpoint.ProviderSpecific {
	if getAliasFromAnnotations(annotations) {
		return map[string]string{"alias": "true"}
	}
	return map[string]string{}
}

// getTargetsFromTargetAnnotation gets endpoints from optional "target" annotation.
// Returns empty endpoints array if none are found.
func getTargetsFromTargetAnnotation(annotations map[string]string) endpoint.Targets {
	var targets endpoint.Targets

	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, exists := annotations[targetAnnotationKey]
	if exists && targetAnnotation != "" {
		// splits the hostname annotation and removes the trailing periods
		targetsList := strings.Split(strings.Replace(targetAnnotation, " ", "", -1), ",")
		for _, targetHostname := range targetsList {
			targetHostname = strings.TrimSuffix(targetHostname, ".")
			targets = append(targets, targetHostname)
		}
	}
	return targets
}

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A for IPs and type CNAME for everything else.
func suitableType(target string) string {
	if net.ParseIP(target) != nil {
		return endpoint.RecordTypeA
	}
	return endpoint.RecordTypeCNAME
}

// endpointsForHostname returns the endpoint objects for each host-target combination.
func endpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL, providerSpecific endpoint.ProviderSpecific) []*endpoint.Endpoint {
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
			DNSName:          strings.TrimSuffix(hostname, "."),
			Targets:          aTargets,
			RecordTTL:        ttl,
			RecordType:       endpoint.RecordTypeA,
			Labels:           endpoint.NewLabels(),
			ProviderSpecific: providerSpecific,
		}
		endpoints = append(endpoints, epA)
	}

	if len(cnameTargets) > 0 {
		epCNAME := &endpoint.Endpoint{
			DNSName:          strings.TrimSuffix(hostname, "."),
			Targets:          cnameTargets,
			RecordTTL:        ttl,
			RecordType:       endpoint.RecordTypeCNAME,
			Labels:           endpoint.NewLabels(),
			ProviderSpecific: providerSpecific,
		}
		endpoints = append(endpoints, epCNAME)
	}

	return endpoints
}
