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
	// The targets the DNS record points to
	Targets []string
	// RecordType type of record, e.g. CNAME, A, TXT etc
	RecordType string
	// TTL for the record
	RecordTTL TTL
	// Labels stores labels defined for the Endpoint
	Labels map[string]string
}

// NewEndpoint initialization method to be used to create an endpoint
func NewEndpoint(dnsName string, targets []string, recordType string) *Endpoint {
	return NewEndpointWithTTL(dnsName, targets, recordType, TTL(0))
}

// NewEndpointWithTTL initialization method to be used to create an endpoint with a TTL struct
func NewEndpointWithTTL(dnsName string, targets []string, recordType string, ttl TTL) *Endpoint {
	for i := range targets {
		targets[i] = strings.TrimSuffix(targets[i], ".")
	}
	return &Endpoint{
		DNSName:    strings.TrimSuffix(dnsName, "."),
		Targets:    targets,
		RecordType: recordType,
		Labels:     map[string]string{},
		RecordTTL:  ttl,
	}
}

func NewEndpointWithLabels(dnsName string, targets []string, recordType string, labels map[string]string) *Endpoint {
	for i := range targets {
		targets[i] = strings.TrimSuffix(targets[i], ".")
	}
	return &Endpoint{
		DNSName:    strings.TrimSuffix(dnsName, "."),
		Targets:    targets,
		RecordType: recordType,
		Labels:     labels,
		RecordTTL:  TTL(0),
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
	return fmt.Sprintf("%s %d IN %s %s %s", e.DNSName, e.RecordTTL, e.RecordType, e.Targets, e.Labels)
}
