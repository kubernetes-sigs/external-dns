/*
Copyright 2020 The Kubernetes Authors.
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

package infomaniak

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	// "sigs.k8s.io/external-dns/provider"
)

var domainsFiltered = []string{"infomaniak.dev"}
var fakeAPIKey = "abcdef12345"

var domain1 = InfomaniakDNSDomain{
	ID:           123,
	CustomerName: "infomaniak.dev",
}
var domain2 = InfomaniakDNSDomain{
	ID:           124,
	CustomerName: "longer.infomaniak.dev",
}
var domain3 = InfomaniakDNSDomain{
	ID:           125,
	CustomerName: "infomaniak.xxx",
}
var domains = []InfomaniakDNSDomain{domain1, domain2, domain3}

var records = map[string][]InfomaniakDNSRecord{
	"infomaniak.dev": {
		{
			Source: "record1",
			Type:   "A",
			TTL:    300,
			Target: "192.168.1.1",
		},
		{
			Source: "record2",
			Type:   "A",
			TTL:    300,
			Target: "192.168.1.2",
		},
		{
			Source: ".",
			Type:   "A",
			TTL:    300,
			Target: "192.168.1.3",
		},
	},
	"infomaniak.xxx": {
		{
			Source: "ignore1",
			Type:   "A",
			TTL:    300,
			Target: "192.168.2.1",
		},
		{
			Source: "ignore2",
			Type:   "A",
			TTL:    300,
			Target: "192.168.2.2",
		},
	},
}

func TestNewInfomaniakProvider(t *testing.T) {
	domainFilter := endpoint.NewDomainFilterWithExclusions(domainsFiltered, nil)

	os.Unsetenv(APITokenVariable)

	p, err := NewInfomaniakProvider(context.TODO(), domainFilter, false)
	// assert.Nil(err, fmt.Sprintf("should not fail, %s", err))
	assert.NotEqual(t, err, nil, "should fail before env is set")

	os.Setenv(APITokenVariable, fakeAPIKey)

	p, err = NewInfomaniakProvider(context.TODO(), domainFilter, false)
	// assert.Nil(err, fmt.Sprintf("should not fail, %s", err))
	assert.Equal(t, err, nil, "should not fail")

	// compare the domainFilter property
	assert.Equal(t, len(p.domainFilter.Filters), 1, "domain filter has wrong size")
	assert.Equal(t, p.domainFilter.Filters[0], domainsFiltered[0], "domain filter not set")
	// compare dryRun property
	assert.Equal(t, p.DryRun, false, "DryRun property should be false")
}

func TestFindMatchingZone(t *testing.T) {
	assert := assert.New(t)

	var testData = []struct {
		in     string
		domain string
		source string
		err    error
	}{
		{in: "abc.longer.infomaniak.dev", domain: "longer.infomaniak.dev", source: "abc", err: nil},
		{in: "abc.infomaniak.dev", domain: "infomaniak.dev", source: "abc", err: nil},
		{in: "abc.def.infomaniak.dev", domain: "infomaniak.dev", source: "abc.def", err: nil},
		{in: "fake.infomaniak.404", domain: "", source: "", err: errorNotFound},
	}

	for _, tt := range testData {
		t.Run(tt.in, func(t *testing.T) {
			domain, source, err := findMatchingZone(&domains, tt.in)
			assert.Equal(err, tt.err)
			if tt.err == nil {
				assert.Equal(domain.CustomerName, tt.domain)
				assert.Equal(source, tt.source)
			}
		})
	}
}

type MockInfomaniakAPI struct{}

type stubRecord struct {
	domain string
	source string
	target string
	ttl uint64
	updated bool
}

var recordsCreated []stubRecord
var recordsDeleted []stubRecord
var recordsUpdated []stubRecord

func NewMockInfomaniakAPI() *MockInfomaniakAPI {
	return &MockInfomaniakAPI{}
}

func (ik *MockInfomaniakAPI) ListDomains() (*[]InfomaniakDNSDomain, error) {
	return &domains, nil
}

func (ik *MockInfomaniakAPI) GetRecords(domain *InfomaniakDNSDomain) (*[]InfomaniakDNSRecord, error) {
	var domainRecords = records[domain.CustomerName]
	return &domainRecords, nil
}

func (ik *MockInfomaniakAPI) EnsureDNSRecord(domain *InfomaniakDNSDomain, source, target, rtype string, ttl uint64) error {
	recordsCreated = append(recordsCreated, stubRecord{domain: domain.CustomerName, source: source, target: target})
	return nil
}

func (ik *MockInfomaniakAPI) RemoveDNSRecord(domain *InfomaniakDNSDomain, source, target, rtype string) error {
	recordsDeleted = append(recordsDeleted, stubRecord{domain: domain.CustomerName, source: source, target: target})
	return nil
}

func (ik *MockInfomaniakAPI) ModifyDNSRecord(domain *InfomaniakDNSDomain, source, oldTarget, newTarget, rtype string, ttl uint64) error {
	for i := range recordsUpdated {
		if recordsUpdated[i].domain == domain.CustomerName &&
			recordsUpdated[i].source == source &&
			recordsUpdated[i].target == oldTarget {
			recordsUpdated[i].target = newTarget
			recordsUpdated[i].updated = true
		}
	}
	return nil
}

func contains(arr []*endpoint.Endpoint, name string) bool {
	for _, a := range arr {
		if a.DNSName == name {
			return true
		}
	}
	return false
}

// func (ik *InfomaniakAPI) ListDomains() (*[]InfomaniakDNSDomain, error) {
func TestRecords(t *testing.T) {
	assert := assert.New(t)
	domainFilter := endpoint.NewDomainFilterWithExclusions(domainsFiltered, nil)

	// var api InfomaniakAPIAdapter = mockapi
	p, err := NewInfomaniakProvider(context.TODO(), domainFilter, false)
	assert.Equal(err, nil)

	mockapi := NewMockInfomaniakAPI()
	p.API = mockapi

	endpoints, err := p.Records(context.TODO())

	assert.Equal(err, nil)

	assert.Equal(len(endpoints), 3, "should contain only the two records of infomaniak.dev")
	assert.True(contains(endpoints, "record1.infomaniak.dev"))
	assert.True(contains(endpoints, "record2.infomaniak.dev"))
	assert.True(contains(endpoints, "infomaniak.dev"))
	assert.False(contains(endpoints, "ignore1.infomaniak.xxx"))
	assert.False(contains(endpoints, "ignore2.infomaniak.xxx"))
}

func TestApplyChanges(t *testing.T) {
	assert := assert.New(t)
	domainFilter := endpoint.NewDomainFilterWithExclusions(domainsFiltered, nil)

	// var api InfomaniakAPIAdapter = mockapi
	p, err := NewInfomaniakProvider(context.TODO(), domainFilter, false)
	assert.Equal(err, nil)

	mockapi := NewMockInfomaniakAPI()
	p.API = mockapi

	emptyPlan := &plan.Changes{
		Create: []*endpoint.Endpoint{},
		Delete: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}

	err = p.ApplyChanges(context.TODO(), emptyPlan)
	assert.Equal(err, nil)

	fullPlan := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "serviceadd.infomaniak.dev",
				RecordType: "A",
				Targets:    []string{"10.233.12.123"},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "servicedelete.infomaniak.dev",
				RecordType: "A",
				Targets:    []string{"10.233.99.123"},
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "updated.infomaniak.dev",
				RecordType: "A",
				Targets:    []string{"10.233.10.10"},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "updated.infomaniak.dev",
				RecordType: "A",
				Targets:    []string{"10.233.20.20"},
			},
		},
	}

	recordsCreated = make([]stubRecord, 0)
	recordsDeleted = make([]stubRecord, 0)
	recordsUpdated = []stubRecord{
		{
			domain: "infomaniak.dev",
			source: "updated",
			target: "10.233.10.10",
			ttl: 300,
		},
		{
			domain: "infomaniak.dev",
			source: "notupdated",
			target: "10.233.30.30",
			ttl: 300,
		},
	}

	err = p.ApplyChanges(context.TODO(), fullPlan)
	assert.Equal(err, nil)

	// test Create
	assert.Equal(recordsCreated[0].domain, "infomaniak.dev")
	assert.Equal(recordsCreated[0].source, "serviceadd")
	assert.Equal(recordsCreated[0].target, "10.233.12.123")

	// test Delete
	assert.Equal(recordsDeleted[0].domain, "infomaniak.dev")
	assert.Equal(recordsDeleted[0].source, "servicedelete")
	assert.Equal(recordsDeleted[0].target, "10.233.99.123")

	// test Update: record to update
	assert.Equal(recordsUpdated[0].target, "10.233.20.20")
	// test Update: record to leave untouched
	assert.Equal(recordsUpdated[1].target, "10.233.30.30")

	// Error when creating record out of known zones
	fullPlan = &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "serviceadd.infomaniak.404",
				RecordType: "A",
				Targets:    []string{"127.0.0.1"},
			},
		},
		Delete: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}
	err = p.ApplyChanges(context.TODO(), fullPlan)
	assert.NotEqual(err, nil)

	// Error when deleting record out of known zones
	fullPlan = &plan.Changes{
		Create: []*endpoint.Endpoint{},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "servicedel.infomaniak.404",
				RecordType: "A",
				Targets:    []string{"127.0.0.1"},
			},
		},
		UpdateOld: []*endpoint.Endpoint{},
		UpdateNew: []*endpoint.Endpoint{},
	}
	err = p.ApplyChanges(context.TODO(), fullPlan)
	assert.NotEqual(err, nil)

	// No update when ttl is equal
	fullPlan = &plan.Changes{
		Create: []*endpoint.Endpoint{},
		Delete: []*endpoint.Endpoint{},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "updated.infomaniak.dev",
				RecordType: "A",
				Targets:    []string{"10.233.10.10"},
				RecordTTL:  300,
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "updated.infomaniak.dev",
				RecordType: "A",
				Targets:    []string{"10.233.10.10"},
				RecordTTL:  300,
			},
		},
	}
	recordsUpdated = []stubRecord{
		{
			domain: "infomaniak.dev",
			source: "updated",
			target: "10.233.10.10",
			ttl: 300,
			updated: false,
		},
	}

	err = p.ApplyChanges(context.TODO(), fullPlan)
	assert.Equal(err, nil)
	assert.False(recordsUpdated[0].updated)

}
