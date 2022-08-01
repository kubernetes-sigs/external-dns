/*
Copyright 2022 The Kubernetes Authors.

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

package plural

import (
	"context"
	"testing"

	"sigs.k8s.io/external-dns/plan"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/provider"
)

type ClientStub struct {
	mockDnsRecords []*DnsRecord
}

// CreateRecord provides a mock function with given fields: record
func (c *ClientStub) CreateRecord(record *DnsRecord) (*DnsRecord, error) {
	c.mockDnsRecords = append(c.mockDnsRecords, record)
	return record, nil
}

// DeleteRecord provides a mock function with given fields: name, ttype
func (c *ClientStub) DeleteRecord(name string, ttype string) error {
	newRecords := make([]*DnsRecord, 0)
	for _, record := range c.mockDnsRecords {
		if record.Name == name && record.Type == ttype {
			continue
		}
		newRecords = append(newRecords, record)
	}
	c.mockDnsRecords = newRecords
	return nil
}

// DnsRecords provides a mock function with given fields:
func (c *ClientStub) DnsRecords() ([]*DnsRecord, error) {
	return c.mockDnsRecords, nil
}

func newPluralProvider(pluralDNSRecord []*DnsRecord) *PluralProvider {
	if pluralDNSRecord == nil {
		pluralDNSRecord = make([]*DnsRecord, 0)
	}
	return &PluralProvider{
		BaseProvider: provider.BaseProvider{},
		Client: &ClientStub{
			mockDnsRecords: pluralDNSRecord,
		},
	}
}

func TestPluralRecords(t *testing.T) {

	tests := []struct {
		name              string
		expectedEndpoints []*endpoint.Endpoint
		records           []*DnsRecord
	}{
		{
			name: "check records",
			records: []*DnsRecord{
				{
					Type:    endpoint.RecordTypeA,
					Name:    "example.com",
					Records: []string{"123.123.123.122"},
				},
				{
					Type:    endpoint.RecordTypeA,
					Name:    "nginx.example.com",
					Records: []string{"123.123.123.123"},
				},
				{
					Type:    endpoint.RecordTypeCNAME,
					Name:    "hack.example.com",
					Records: []string{"bluecatnetworks.com"},
				},
				{
					Type:    endpoint.RecordTypeTXT,
					Name:    "kdb.example.com",
					Records: []string{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
			expectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.122"},
				},
				{
					DNSName:    "nginx.example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.123"},
				},
				{
					DNSName:    "hack.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"bluecatnetworks.com"},
				},
				{
					DNSName:    "kdb.example.com",
					RecordType: endpoint.RecordTypeTXT,
					Targets:    endpoint.Targets{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := newPluralProvider(test.records)

			actual, err := provider.Records(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			validateEndpoints(t, actual, test.expectedEndpoints)
		})
	}

}

func TestPluralApplyChangesCreate(t *testing.T) {

	tests := []struct {
		name              string
		expectedEndpoints []*endpoint.Endpoint
	}{
		{
			name: "create new endpoints",
			expectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.122"},
				},
				{
					DNSName:    "nginx.example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.123"},
				},
				{
					DNSName:    "hack.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"bluecatnetworks.com"},
				},
				{
					DNSName:    "kdb.example.com",
					RecordType: endpoint.RecordTypeTXT,
					Targets:    endpoint.Targets{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := newPluralProvider(nil)

			// no records
			actual, err := provider.Records(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, len(actual), 0, "expected no entries")

			err = provider.ApplyChanges(context.Background(), &plan.Changes{Create: test.expectedEndpoints})
			if err != nil {
				t.Fatal(err)
			}

			actual, err = provider.Records(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			validateEndpoints(t, actual, test.expectedEndpoints)
		})
	}
}

func TestPluralApplyChangesDelete(t *testing.T) {

	tests := []struct {
		name              string
		records           []*DnsRecord
		deleteEndpoints   []*endpoint.Endpoint
		expectedEndpoints []*endpoint.Endpoint
	}{
		{
			name: "delete not existing record",
			records: []*DnsRecord{
				{
					Type:    endpoint.RecordTypeA,
					Name:    "example.com",
					Records: []string{"123.123.123.122"},
				},
				{
					Type:    endpoint.RecordTypeA,
					Name:    "nginx.example.com",
					Records: []string{"123.123.123.123"},
				},
				{
					Type:    endpoint.RecordTypeCNAME,
					Name:    "hack.example.com",
					Records: []string{"bluecatnetworks.com"},
				},
				{
					Type:    endpoint.RecordTypeTXT,
					Name:    "kdb.example.com",
					Records: []string{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
			deleteEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "fake.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.2.3.4"},
				},
			},
			expectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.122"},
				},
				{
					DNSName:    "nginx.example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.123"},
				},
				{
					DNSName:    "hack.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"bluecatnetworks.com"},
				},
				{
					DNSName:    "kdb.example.com",
					RecordType: endpoint.RecordTypeTXT,
					Targets:    endpoint.Targets{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
		},
		{
			name: "delete one record",
			records: []*DnsRecord{
				{
					Type:    endpoint.RecordTypeA,
					Name:    "example.com",
					Records: []string{"123.123.123.122"},
				},
				{
					Type:    endpoint.RecordTypeA,
					Name:    "nginx.example.com",
					Records: []string{"123.123.123.123"},
				},
				{
					Type:    endpoint.RecordTypeCNAME,
					Name:    "hack.example.com",
					Records: []string{"bluecatnetworks.com"},
				},
				{
					Type:    endpoint.RecordTypeTXT,
					Name:    "kdb.example.com",
					Records: []string{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
			deleteEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "kdb.example.com",
					RecordType: endpoint.RecordTypeTXT,
					Targets:    endpoint.Targets{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
			expectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.122"},
				},
				{
					DNSName:    "nginx.example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.123"},
				},
				{
					DNSName:    "hack.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"bluecatnetworks.com"},
				},
			},
		},
		{
			name: "delete all records",
			records: []*DnsRecord{
				{
					Type:    endpoint.RecordTypeA,
					Name:    "example.com",
					Records: []string{"123.123.123.122"},
				},
				{
					Type:    endpoint.RecordTypeA,
					Name:    "nginx.example.com",
					Records: []string{"123.123.123.123"},
				},
				{
					Type:    endpoint.RecordTypeCNAME,
					Name:    "hack.example.com",
					Records: []string{"bluecatnetworks.com"},
				},
				{
					Type:    endpoint.RecordTypeTXT,
					Name:    "kdb.example.com",
					Records: []string{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
			},
			deleteEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "kdb.example.com",
					RecordType: endpoint.RecordTypeTXT,
					Targets:    endpoint.Targets{"heritage=external-dns,external-dns/owner=default,external-dns/resource=service/openshift-ingress/router-default"},
				},
				{
					DNSName:    "example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.122"},
				},
				{
					DNSName:    "nginx.example.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"123.123.123.123"},
				},
				{
					DNSName:    "hack.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"bluecatnetworks.com"},
				},
			},
			expectedEndpoints: []*endpoint.Endpoint{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := newPluralProvider(test.records)

			err := provider.ApplyChanges(context.Background(), &plan.Changes{Delete: test.deleteEndpoints})
			if err != nil {
				t.Fatal(err)
			}

			actual, err := provider.Records(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			validateEndpoints(t, actual, test.expectedEndpoints)
		})
	}
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "expected and actual endpoints don't match. %s:%s", endpoints, expected)
}
