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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/registry"
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

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 2, verifiedRecords.Gauge, map[string]string{"record_type": "a"})

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

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 2, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})
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

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 2, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})

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

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 2, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})
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
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})
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

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, sourceRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 2, sourceRecords.Gauge, map[string]string{"record_type": "aaaa"})

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "aaaa"})
}

func TestGaugeMetricsWithMixedRecords(t *testing.T) {
	configuredEndpoints := testutils.GenerateTestEndpointsByType(map[string]int{
		endpoint.RecordTypeA:     534,
		endpoint.RecordTypeAAAA:  324,
		endpoint.RecordTypeCNAME: 2,
		endpoint.RecordTypeTXT:   56,
		endpoint.RecordTypeSRV:   11,
		endpoint.RecordTypeNS:    3,
	})

	providerEndpoints := testutils.GenerateTestEndpointsByType(map[string]int{
		endpoint.RecordTypeA:     5334,
		endpoint.RecordTypeAAAA:  324,
		endpoint.RecordTypeCNAME: 23,
		endpoint.RecordTypeTXT:   6,
		endpoint.RecordTypeSRV:   25,
		endpoint.RecordTypeNS:    1,
		endpoint.RecordTypePTR:   43,
	})

	cfg := externaldns.NewConfig()
	cfg.ManagedDNSRecordTypes = endpoint.KnownRecordTypes

	source := new(testutils.MockSource)
	source.On("Endpoints").Return(configuredEndpoints, nil)

	provider := &filteredMockProvider{
		RecordsStore: providerEndpoints,
	}
	r, err := registry.NewNoopRegistry(provider)

	require.NoError(t, err)

	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		DomainFilter:       endpoint.NewDomainFilter([]string{}),
		ManagedRecordTypes: cfg.ManagedDNSRecordTypes,
	}

	assert.NoError(t, ctrl.RunOnce(t.Context()))

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 534, sourceRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 324, sourceRecords.Gauge, map[string]string{"record_type": "aaaa"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, sourceRecords.Gauge, map[string]string{"record_type": "cname"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 11, sourceRecords.Gauge, map[string]string{"record_type": "srv"})

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 5334, registryRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 324, registryRecords.Gauge, map[string]string{"record_type": "aaaa"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 0, registryRecords.Gauge, map[string]string{"record_type": "mx"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 43, registryRecords.Gauge, map[string]string{"record_type": "ptr"})
}
