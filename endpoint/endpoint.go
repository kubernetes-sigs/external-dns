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

package endpoint

import (
	"fmt"
	"strings"
)

const (
	// RecordTypeA is a RecordType enum value
	RecordTypeA = "A"
	// RecordTypeCNAME is a RecordType enum value
	RecordTypeCNAME = "CNAME"
	// RecordTypeTXT is a RecordType enum value
	RecordTypeTXT = "TXT"
)

// TTL is a structure defining the TTL of a DNS record
type TTL int64

// IsConfigured returns true if TTL is configured, false otherwise
func (ttl TTL) IsConfigured() bool {
	return ttl > 0
}

// Endpoint is a high-level way of a connection between a service and an IP
type Endpoint struct {
	// The hostname of the DNS record
	DNSName string
	// The target the DNS record points to
	Target string
	// RecordType type of record, e.g. CNAME, A, TXT etc
	RecordType string
	// TTL for the record
	RecordTTL TTL
	// Labels stores labels defined for the Endpoint
	Labels Labels
}

// NewEndpoint initialization method to be used to create an endpoint
func NewEndpoint(dnsName, target, recordType string) *Endpoint {
	return NewEndpointWithTTL(dnsName, target, recordType, TTL(0))
}

// NewEndpointWithTTL initialization method to be used to create an endpoint with a TTL struct
func NewEndpointWithTTL(dnsName, target, recordType string, ttl TTL) *Endpoint {
	return &Endpoint{
		DNSName:    strings.TrimSuffix(dnsName, "."),
		Target:     strings.TrimSuffix(target, "."),
		RecordType: recordType,
		Labels:     NewLabels(),
		RecordTTL:  ttl,
	}
}

func (e *Endpoint) String() string {
	return fmt.Sprintf("%s %d IN %s %s", e.DNSName, e.RecordTTL, e.RecordType, e.Target)
}
