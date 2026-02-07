/*
Copyright 2026 The Kubernetes Authors.

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

package endpoint

import (
	"net/netip"
	"slices"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	msg = "No endpoints could be generated from '%s/%s/%s'"
)

// HasNoEmptyEndpoints checks if the endpoint list is empty and logs
// a debug message if so. Returns true if empty, false otherwise.
func HasNoEmptyEndpoints(
	endpoints []*Endpoint,
	rType string, entity metav1.ObjectMetaAccessor,
) bool {
	if len(endpoints) == 0 {
		log.Debugf(msg, rType, entity.GetObjectMeta().GetNamespace(), entity.GetObjectMeta().GetName())
		return true
	}
	return false
}

// SuitableType returns the DNS resource record type suitable for the target.
// Returns type A for IPv4 addresses, AAAA for IPv6 addresses, and CNAME for everything else.
func SuitableType(target string) string {
	netIP, err := netip.ParseAddr(target)
	if err != nil {
		return RecordTypeCNAME
	}
	switch {
	case netIP.Is4():
		return RecordTypeA
	case netIP.Is6():
		return RecordTypeAAAA
	default:
		return RecordTypeCNAME
	}
}

// EndpointsForHostsAndTargets creates endpoints by grouping targets by record type
// and creating an endpoint for each hostname/record-type combination.
// The function returns endpoints in deterministic order (sorted by record type).
func EndpointsForHostsAndTargets(hostnames, targets []string) []*Endpoint {
	if len(hostnames) == 0 || len(targets) == 0 {
		return nil
	}

	// Group targets by record type
	targetsByType := make(map[string][]string)
	for _, target := range targets {
		recordType := SuitableType(target)
		targetsByType[recordType] = append(targetsByType[recordType], target)
	}

	// Pre-allocate with estimated capacity
	endpoints := make([]*Endpoint, 0, len(hostnames)*len(targetsByType))

	// Sort record types for deterministic ordering
	recordTypes := make([]string, 0, len(targetsByType))
	for recordType := range targetsByType {
		recordTypes = append(recordTypes, recordType)
	}
	slices.Sort(recordTypes)

	for _, hostname := range hostnames {
		for _, recordType := range recordTypes {
			endpoints = append(endpoints, NewEndpoint(hostname, recordType, targetsByType[recordType]...))
		}
	}

	return endpoints
}
