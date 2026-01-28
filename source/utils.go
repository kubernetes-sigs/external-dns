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
	"fmt"
	"net/netip"
	"sort"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
)

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A/AAAA for IPs and type CNAME for everything else.
// TODO: move this to the endpoint package?
func suitableType(target string) string {
	netIP, err := netip.ParseAddr(target)
	if err != nil {
		return endpoint.RecordTypeCNAME
	}
	switch {
	case netIP.Is4():
		return endpoint.RecordTypeA
	case netIP.Is6():
		return endpoint.RecordTypeAAAA
	default:
		return endpoint.RecordTypeCNAME
	}
}

// ParseIngress parses an ingress string in the format "namespace/name" or "name".
// It returns the namespace and name extracted from the string, or an error if the format is invalid.
// If the namespace is not provided, it defaults to an empty string.
func ParseIngress(ingress string) (string, string, error) {
	var namespace, name string
	var err error
	parts := strings.Split(ingress, "/")
	switch len(parts) {
	case 2:
		namespace, name = parts[0], parts[1]
	case 1:
		name = parts[0]
	default:
		err = fmt.Errorf("invalid ingress name (name or namespace/name) found %q", ingress)
	}

	return namespace, name, err
}

// MatchesServiceSelector checks if all key-value pairs in the selector map
// are present and match the corresponding key-value pairs in the svcSelector map.
// It returns true if all pairs match, otherwise it returns false.
func MatchesServiceSelector(selector, svcSelector map[string]string) bool {
	for k, v := range selector {
		if lbl, ok := svcSelector[k]; !ok || lbl != v {
			return false
		}
	}
	return true
}

// MergeEndpoints merges endpoints with the same key (DNSName + RecordType + RecordTTL)
// by combining their targets. This is useful when multiple resources (e.g., pods, nodes)
// contribute targets to the same DNS record.
//
// TODO: move this to endpoint/utils.go
// TODO: apply to all sources that generate endpoints (e.g., service, ingress, etc.)
func MergeEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpointMap := make(map[endpoint.EndpointKey]*endpoint.Endpoint)

	for _, ep := range endpoints {
		key := endpoint.EndpointKey{
			DNSName:    ep.DNSName,
			RecordType: ep.RecordType,
			RecordTTL:  ep.RecordTTL,
		}
		if existing, ok := endpointMap[key]; ok {
			existing.Targets = append(existing.Targets, ep.Targets...)
		} else {
			endpointMap[key] = ep
		}
	}

	result := make([]*endpoint.Endpoint, 0, len(endpointMap))
	for _, ep := range endpointMap {
		sort.Sort(ep.Targets)
		result = append(result, ep)
	}

	return result
}
