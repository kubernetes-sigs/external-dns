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

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/pkg/events"
)

const (
	msg = "No endpoints could be generated from '%s/%s/%s'"
)

// SuitableType returns the DNS record type for the given target:
// A for IPv4, AAAA for IPv6, CNAME for everything else.
func SuitableType(target string) string {
	ip, err := netip.ParseAddr(target)
	if err != nil {
		return RecordTypeCNAME
	}
	switch {
	case ip.Is4():
		return RecordTypeA
	case ip.Is6():
		return RecordTypeAAAA
	default:
		return RecordTypeCNAME
	}
}

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

// AttachRefObject sets the same ObjectReference on every endpoint in eps.
// The reference is shared across all endpoints, so callers should create it once
// per source object rather than once per endpoint.
func AttachRefObject(eps []*Endpoint, ref *events.ObjectReference) {
	for _, ep := range eps {
		ep.WithRefObject(ref)
	}
}
