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
package efficientip

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"log"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"strings"
	"testing"
)

type Record struct {
	Name  string
	Type  string
	TTL   int
	Value string
}

type EfficientIPStub struct {
	Zones   []*ZoneAuth
	Records map[string][]*Record
	m       dynamicMock
	t       *testing.T
}

// suitableZones returns all suitable private zones and the most suitable public zone
//   for a given hostname and a set of zones.
func suitableZones(hostname string, zones []*ZoneAuth) *ZoneAuth {
	var matchingZone *ZoneAuth

	for _, z := range zones {
		if z.Name == hostname || strings.HasSuffix(hostname, "."+z.Name) {
			// Only select the best matching public zone
			if matchingZone == nil || len(z.Name) > len(matchingZone.Name) {
				matchingZone = z
			}
		}
	}
	return matchingZone
}

func (e *EfficientIPStub) AddZone(Name string) {
	e.Zones = append(e.Zones, &ZoneAuth{Name: Name, Type: "master", ID: Name})
}

func (e *EfficientIPStub) ZonesList() ([]*ZoneAuth, error) {
	return e.Zones, nil
}

func (e *EfficientIPStub) RecordAdd(rr *endpoint.Endpoint) error {
	zone := suitableZones(rr.DNSName, e.Zones)
	if strings.HasSuffix(rr.DNSName, zone.Name) {
		for _, value := range rr.Targets {
			e.Records[zone.Name] = append(e.Records[zone.Name],
				&Record{Name: rr.DNSName, Type: rr.RecordType, TTL: int(rr.RecordTTL), Value: value})
		}
		return nil
	}
	return errors.New("zone not found")
}

func (e *EfficientIPStub) RecordDelete(rr *endpoint.Endpoint) error {
	zone := suitableZones(rr.DNSName, e.Zones)
	if strings.HasSuffix(rr.DNSName, zone.Name) {
		for _, value := range rr.Targets {
			for i, r := range e.Records[zone.Name] {
				if r.Name == rr.DNSName && r.Type == rr.RecordType && r.Value == value {
					e.Records[zone.Name] = append(e.Records[zone.Name][:i], e.Records[zone.Name][i+1:]...)
				}
			}
		}
		return nil
	}
	return errors.New("record not found")
}

func (e *EfficientIPStub) RecordList(Zone ZoneAuth) (endpoints []*endpoint.Endpoint, _ error) {
	Host := make(map[string]*endpoint.Endpoint)

	for _, rr := range e.Records[Zone.Name] {
		if rr.Type == "A" {
			if h, found := Host[rr.Name+":"+rr.Type]; found {
				h.Targets = append(h.Targets, rr.Value)
			} else {
				Host[rr.Name+":"+rr.Type] = endpoint.NewEndpointWithTTL(rr.Name, endpoint.RecordTypeA, endpoint.TTL(rr.TTL), rr.Value)
			}
		} else {
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(rr.Name, rr.Type, endpoint.TTL(rr.TTL), rr.Value))
		}
	}

	for _, rr := range Host {
		endpoints = append(endpoints, rr)
	}
	return endpoints, nil
}

// Compile time check for interface conformance
var _ EfficientipClient = &EfficientIPStub{}

// MockMethod starts a description of an expectation of the specified method
// being called.
//
//     mockAPIClient.MockMethod("MyMethod", arg1, arg2)
func (e *EfficientIPStub) MockMethod(method string, args ...interface{}) *mock.Call {
	return e.m.On(method, args...)
}

type dynamicMock struct {
	mock.Mock
}

func (d *dynamicMock) isMocked(method string, arguments ...interface{}) bool {
	for _, call := range d.ExpectedCalls {
		if call.Method == method && call.Repeatability > -1 {
			_, diffCount := call.Arguments.Diff(arguments)
			if diffCount == 0 {
				return true
			}
		}
	}
	return false
}

func newEfficientIPStub(t *testing.T) *EfficientIPStub {

	return &EfficientIPStub{
		t: t,
	}
}

func newEfficientipProvider(t *testing.T, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, _ provider.ZoneTypeFilter, dryRun bool, records []*endpoint.Endpoint) (*EfficientIPProvider, *EfficientIPStub) {
	client := newEfficientIPStub(t)

	for _, zone := range domainFilter.Filters {
		client.AddZone(zone)
	}
	client.Records = make(map[string][]*Record)
	for _, record := range records {
		zone := suitableZones(record.DNSName, client.Zones)
		for _, value := range record.Targets {
			rr := &Record{
				Name:  record.DNSName,
				Type:  record.RecordType,
				TTL:   int(record.RecordTTL),
				Value: value,
			}
			client.Records[zone.Name] = append(client.Records[zone.Name], rr)
		}
	}

	p := &EfficientIPProvider{
		client:       client,
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		dryRun:       dryRun,
	}

	return p, client
}

func validateEndpoints(t *testing.T, endpoints []*endpoint.Endpoint, expected []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(endpoints, expected), "actual and expected endpoints don't match. %+v:%+v", endpoints, expected)
}

func TestSuitableZones(t *testing.T) {
	zones := map[string]*ZoneAuth{
		// Public domain
		"example.org": {Name: "example.org", ID: "example.org"},
		// Public subdomain
		"bar.example.org": {Name: "bar.example.org", ID: "bar.example.org"},
		// Public subdomain
		"longfoo.bar.example.org": {Name: "longfoo.bar.example.org", ID: "longfoo.bar.example.org"},
	}
	db := []*ZoneAuth{
		zones["example.org"],
		zones["bar.example.org"],
		zones["longfoo.bar.example.org"],
	}

	for _, tc := range []struct {
		hostname string
		expected *ZoneAuth
	}{
		// bar.example.org is NOT suitable
		{"foobar.example.org", zones["example.org"]},

		// all matching private zones are suitable
		// https://github.com/kubernetes-sigs/external-dns/pull/356
		{"bar.example.org", zones["bar.example.org"]},

		{"foo.bar.example.org", zones["bar.example.org"]},
		{"foo.example.org", zones["example.org"]},
		{"foo.kubernetes.io", nil},
	} {
		log.Printf("Current DB : %+v", db)
		suitableZone := suitableZones(tc.hostname, db)
		assert.Equal(t, tc.expected, suitableZone)
	}
}

func TestEfficientIPProvider_ApplyChanges(t *testing.T) {
	recordTTL := 60

	p, _ := newEfficientipProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), false, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.1.1.1"),
		endpoint.NewEndpointWithTTL("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.eu-central-1.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "qux.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "qux.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "4.3.2.1"),
	})

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("create-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("create-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("create-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("create-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("create-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.1.1.1"),
		endpoint.NewEndpoint("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8", "8.8.4.4"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "4.3.2.1"),
		endpoint.NewEndpointWithTTL("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "my-internal-host.example.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "baz.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "baz.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "4.3.2.1"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4", "4.3.2.1"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	ctx := context.Background()

	require.NoError(t, p.ApplyChanges(ctx, changes))

	records, err := p.Records(ctx)
	require.NoError(t, err)

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("create-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("create-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "4.3.2.1"),
		endpoint.NewEndpointWithTTL("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "my-internal-host.example.com"),
		endpoint.NewEndpointWithTTL("create-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "baz.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("create-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "baz.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("create-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "4.3.2.1"),
	})
}

func TestEfficientIPProvider_DryRunApplyChanges(t *testing.T) {
	recordTTL := 60

	p, _ := newEfficientipProvider(t, endpoint.NewDomainFilter([]string{"ext-dns-test-2.teapot.zalan.do."}), provider.NewZoneIDFilter([]string{}), provider.NewZoneTypeFilter(""), true, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.1.1.1"),
		endpoint.NewEndpointWithTTL("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.eu-central-1.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "qux.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "qux.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "4.3.2.1"),
	})

	createRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("create-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("create-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("create-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("create-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("create-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
	}

	currentRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.1.1.1"),
		endpoint.NewEndpoint("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "foo.eu-central-1.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "bar.elb.amazonaws.com"),
		endpoint.NewEndpoint("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8", "8.8.4.4"),
	}
	updatedRecords := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "4.3.2.1"),
		endpoint.NewEndpointWithTTL("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "my-internal-host.example.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "baz.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "baz.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "4.3.2.1"),
	}

	deleteRecords := []*endpoint.Endpoint{
		endpoint.NewEndpoint("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.8.8"),
		endpoint.NewEndpoint("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "8.8.4.4"),
		endpoint.NewEndpoint("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, "qux.elb.amazonaws.com"),
		endpoint.NewEndpoint("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, "1.2.3.4", "4.3.2.1"),
	}

	changes := &plan.Changes{
		Create:    createRecords,
		UpdateNew: updatedRecords,
		UpdateOld: currentRecords,
		Delete:    deleteRecords,
	}

	ctx := context.Background()

	require.NoError(t, p.ApplyChanges(ctx, changes))

	records, err := p.Records(ctx)
	require.NoError(t, err)

	validateEndpoints(t, records, []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("update-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("delete-test.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8"),
		endpoint.NewEndpointWithTTL("update-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.4.4"),
		endpoint.NewEndpointWithTTL("update-test-a-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.1.1.1"),
		endpoint.NewEndpointWithTTL("update-test-alias-to-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "foo.eu-central-1.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "qux.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "bar.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("delete-test-cname-alias.zone-1.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeCNAME, endpoint.TTL(recordTTL), "qux.elb.amazonaws.com"),
		endpoint.NewEndpointWithTTL("update-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "8.8.8.8", "8.8.4.4"),
		endpoint.NewEndpointWithTTL("delete-test-multiple.zone-2.ext-dns-test-2.teapot.zalan.do", endpoint.RecordTypeA, endpoint.TTL(recordTTL), "1.2.3.4", "4.3.2.1"),
	})
}
