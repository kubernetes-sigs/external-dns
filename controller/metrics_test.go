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

package controller

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func TestRecordKnownEndpointType(t *testing.T) {
	mr := newMetricsRecorder()

	// Recording a built-in type should start at 1 and increment
	mr.recordEndpointType(endpoint.RecordTypeA)
	assert.Equal(t, 1, mr.getEndpointTypeCount(endpoint.RecordTypeA))

	mr.recordEndpointType(endpoint.RecordTypeA)
	assert.Equal(t, 2, mr.getEndpointTypeCount(endpoint.RecordTypeA))
}

func TestRecordUnknownEndpointType(t *testing.T) {
	mr := newMetricsRecorder()
	const customType = "CUSTOM"

	// Unknown types start at zero
	assert.Equal(t, 0, mr.getEndpointTypeCount(customType))

	// First record sets to 1
	mr.recordEndpointType(customType)
	assert.Equal(t, 1, mr.getEndpointTypeCount(customType))

	// Subsequent records increment
	mr.recordEndpointType(customType)
	assert.Equal(t, 2, mr.getEndpointTypeCount(customType))
}

func TestLoadFloat64(t *testing.T) {
	mr := newMetricsRecorder()

	// loadFloat64 should return the float64 representation of the count
	mr.recordEndpointType(endpoint.RecordTypeAAAA)
	assert.InDelta(t, float64(1), mr.loadFloat64(endpoint.RecordTypeAAAA), 0.0001)
}

func TestVerifyARecords(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "create-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
			{
				DNSName:    "create-record.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
		},
		[]*plan.Changes{},
	)
	assert.Equal(t, math.Float64bits(2), valueFromMetric(verifiedARecords.Gauge))

	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.1.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "some-record.2.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
			{
				DNSName:    "some-record.3.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"24.24.24.24"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.1.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "some-record.2.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
		},
		[]*plan.Changes{{
			Create: []*endpoint.Endpoint{
				{
					DNSName:    "some-record.3.used.tld",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"24.24.24.24"},
				},
			},
		}},
	)
	assert.Equal(t, math.Float64bits(2), valueFromMetric(verifiedARecords.Gauge))
	assert.Equal(t, math.Float64bits(0), valueFromMetric(verifiedAAAARecords.Gauge))
}

func TestVerifyAAAARecords(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "create-record.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::2"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::2"},
			},
			{
				DNSName:    "create-record.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
		},
		[]*plan.Changes{},
	)
	assert.Equal(t, math.Float64bits(2), valueFromMetric(verifiedAAAARecords.Gauge))

	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.1.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "some-record.2.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::2"},
			},
			{
				DNSName:    "some-record.3.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::3"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "some-record.1.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "some-record.2.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::2"},
			},
		},
		[]*plan.Changes{{
			Create: []*endpoint.Endpoint{
				{
					DNSName:    "some-record.3.used.tld",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets{"2001:DB8::3"},
				},
			},
		}},
	)
	assert.Equal(t, math.Float64bits(0), valueFromMetric(verifiedARecords.Gauge))
	assert.Equal(t, math.Float64bits(2), valueFromMetric(verifiedAAAARecords.Gauge))
}

func TestARecords(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "record1.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "record2.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"8.8.8.8"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "record1.used.tld",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
		},
		[]*plan.Changes{{
			Create: []*endpoint.Endpoint{
				{
					DNSName:    "record2.used.tld",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
			},
		}},
	)
	assert.Equal(t, math.Float64bits(2), valueFromMetric(sourceARecords.Gauge))
	assert.Equal(t, math.Float64bits(1), valueFromMetric(registryARecords.Gauge))
}

func TestAAAARecords(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "record1.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "record2.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::2"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "record1.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
		},
		[]*plan.Changes{{
			Create: []*endpoint.Endpoint{
				{
					DNSName:    "record2.used.tld",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets{"2001:DB8::2"},
				},
			},
		}},
	)
	assert.Equal(t, math.Float64bits(2), valueFromMetric(sourceAAAARecords.Gauge))
	assert.Equal(t, math.Float64bits(1), valueFromMetric(registryAAAARecords.Gauge))
}

func TestMixedRecords(t *testing.T) {
	testControllerFiltersDomains(
		t,
		[]*endpoint.Endpoint{
			{
				DNSName:    "record1.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "record2.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::2"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
			{
				DNSName:    "example.com",
				RecordType: endpoint.RecordTypeMX,
				Targets:    endpoint.Targets{"10 example.com"},
			},
			{
				DNSName:    "cloud-ttl",
				RecordType: endpoint.RecordTypeNS,
				Targets:    endpoint.Targets{"ns1-ttl.example.com."},
			},
			{
				DNSName:    "cloud.example.com",
				RecordType: endpoint.RecordTypeNS,
				Targets:    endpoint.Targets{"ns1.example.com."},
			},
		},
		endpoint.NewDomainFilter([]string{"used.tld"}),
		[]*endpoint.Endpoint{
			{
				DNSName:    "record1.used.tld",
				RecordType: endpoint.RecordTypeAAAA,
				Targets:    endpoint.Targets{"2001:DB8::1"},
			},
			{
				DNSName:    "_mysql-svc._tcp.mysql.used.tld",
				RecordType: endpoint.RecordTypeSRV,
				Targets:    endpoint.Targets{"0 50 30007 mysql.used.tld"},
			},
			{
				DNSName:    "100.2.0.192.in-addr.arpa",
				RecordType: endpoint.RecordTypePTR,
				Targets:    endpoint.Targets{"mail.example.com"},
			},
		},
		[]*plan.Changes{{
			Create: []*endpoint.Endpoint{
				{
					DNSName:    "record2.used.tld",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets{"2001:DB8::2"},
				},
			},
		}},
	)

	assert.Equal(t, math.Float64bits(2), valueFromMetric(sourceSRVRecords.Gauge))
	assert.Equal(t, math.Float64bits(1), valueFromMetric(registrySRVRecords.Gauge))

	assert.Equal(t, math.Float64bits(1), valueFromMetric(sourceMXRecords.Gauge))
	assert.Equal(t, math.Float64bits(0), valueFromMetric(registryMXRecords.Gauge))

	assert.Equal(t, math.Float64bits(0), valueFromMetric(sourcePTRRecords.Gauge))
	assert.Equal(t, math.Float64bits(1), valueFromMetric(registryPTRRecords.Gauge))

	assert.Equal(t, math.Float64bits(2), valueFromMetric(sourceNSRecords.Gauge))
	assert.Equal(t, math.Float64bits(0), valueFromMetric(registryNSRecords.Gauge))
}
