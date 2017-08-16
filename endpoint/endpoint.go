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
	// OwnerLabelKey is the name of the label that defines the owner of an Endpoint.
	OwnerLabelKey = "owner"
)

// TTL is a structure defining the TTL of a DNS record
type TTL struct {
	// The value of the TTL
	Value int64
	// Whether or not the TTL value is configured
	IsConfigured bool
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
	Labels map[string]string
}

// NewEndpoint initialization method to be used to create an endpoint
func NewEndpoint(dnsName, target, recordType string) *Endpoint {
	return NewEndpointWithTTL(dnsName, target, recordType, nil)
}

// NewEndpointWithTTL initialization method to be used to create an endpoint with TTL
func NewEndpointWithTTL(dnsName, target, recordType string, ttlValue *int64) *Endpoint {
	var ttl TTL
	if ttlValue != nil {
		ttl = TTL{Value: *ttlValue, IsConfigured: true}
	} else {
		ttl = TTL{IsConfigured: false}
	}
	return &Endpoint{
		DNSName:    strings.TrimSuffix(dnsName, "."),
		Target:     strings.TrimSuffix(target, "."),
		RecordType: recordType,
		Labels:     map[string]string{},
		RecordTTL:  ttl,
	}
}

// MergeLabels adds keys to labels if not defined for the endpoint
func (e *Endpoint) MergeLabels(labels map[string]string) {
	for k, v := range labels {
		if e.Labels[k] == "" {
			e.Labels[k] = v
		}
	}
}

func (e *Endpoint) String() string {
	var ttl string
	if e.RecordTTL.IsConfigured {
		ttl = fmt.Sprintf(" TTL: %v", e.RecordTTL.Value)
	}
	return fmt.Sprintf(`%s -> %s (type "%s") %s`, e.DNSName, e.Target, e.RecordType, ttl)
}
