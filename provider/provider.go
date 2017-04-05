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

package provider

import (
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// Provider defines the interface DNS providers should implement.
type Provider interface {
	Records(zone string) ([]*endpoint.Endpoint, error)
	ApplyChanges(zone string, changes *plan.Changes) error
}

// DNSRecord is dns provider agnostic concept
type DNSRecord struct {
	DNSName    string
	Target     string
	RecordType string
	Value      string
}

// NewDNSRecord creates new dns record
func NewDNSRecord(dnsName, target, recordType, value string) *DNSRecord {
	return &DNSRecord{
		DNSName:    dnsName,
		Target:     target,
		RecordType: recordType,
		Value:      value,
	}
}
