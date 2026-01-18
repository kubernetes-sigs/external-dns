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

package plan

import (
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestOwnerMismatchMetric(t *testing.T) {
	currentA := &endpoint.Endpoint{
		DNSName:    "example.domain.com",
		Targets:    endpoint.Targets{"1.2.3.4"},
		RecordType: endpoint.RecordTypeA,
		Labels: map[string]string{
			endpoint.OwnerLabelKey: "other-owner",
		},
	}
	desiredCname := &endpoint.Endpoint{
		DNSName:    "example.domain.com",
		Targets:    endpoint.Targets{"target.example.com"},
		RecordType: endpoint.RecordTypeCNAME,
		Labels: map[string]string{
			endpoint.OwnerLabelKey: "my-owner",
		},
	}

	p := &Plan{
		Policies:       []Policy{&SyncPolicy{}},
		Current:        []*endpoint.Endpoint{currentA},
		Desired:        []*endpoint.Endpoint{desiredCname},
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
		OwnerID:        "my-owner",
	}

	changes := p.Calculate().Changes
	assert.Empty(t, changes.Create, "expected no creates due to owner mismatch")

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
		t,
		1.0,
		registryOwnerMismatchPerSync.Gauge,
		map[string]string{
			"record_type":   endpoint.RecordTypeA,
			"foreign_owner": "other-owner",
			"domain":        "domain.com",
		},
	)
}

// TestCalculateOwnerMismatchDetection verifies that owner mismatch is detected
// when desired endpoints want to create new record types on DNS names
// that have current records owned by a different owner.
func TestCalculateOwnerMismatchDetection(t *testing.T) {
	current := testutils.GenerateTestEndpointsWithDistribution(
		map[string]int{endpoint.RecordTypeA: 10},
		map[string]int{"example.com": 1},
		map[string]int{"other-owner": 1},
	)

	// Create desired endpoints: same DNS names but with different type records (new type triggers Create)
	var desired []*endpoint.Endpoint
	for _, ep := range current {
		desired = append(desired, &endpoint.Endpoint{
			DNSName:    ep.DNSName,
			Targets:    endpoint.Targets{"abrakadabra"},
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  300,
		})
	}

	p := &Plan{
		Policies:       []Policy{&SyncPolicy{}},
		Current:        current,
		Desired:        desired,
		ManagedRecords: endpoint.KnownRecordTypes,
		OwnerID:        "my-owner",
	}
	hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)
	changes := p.Calculate().Changes

	assert.Empty(t, changes.Create, "expected no creates due to owner mismatch")
	testutils.TestHelperLogContains("owner id does not match for one or more items to create", hook, t)
}

func TestOwnerMismatchMetricDistribution(t *testing.T) {
	p := newOwnerMismatchFixture()

	p.Calculate()
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 44, registryOwnerMismatchPerSync.Gauge,
		map[string]string{"record_type": endpoint.RecordTypeSRV})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 41, registryOwnerMismatchPerSync.Gauge,
		map[string]string{"foreign_owner": "owner1"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 125, registryOwnerMismatchPerSync.Gauge,
		map[string]string{"owner": "my-owner"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 21, registryOwnerMismatchPerSync.Gauge,
		map[string]string{"foreign_owner": "owner1", "domain": "open.net"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 2, registryOwnerMismatchPerSync.Gauge,
		map[string]string{"record_type": endpoint.RecordTypeCNAME, "foreign_owner": "owner1", "domain": "open.net"})
}

func BenchmarkOwnerMismatchMetricDistribution(b *testing.B) {
	p := newOwnerMismatchFixture(1000)

	for b.Loop() {
		p.Calculate()
	}
}

func newOwnerMismatchFixture(scale ...int) *Plan {
	factor := 1
	if len(scale) > 0 && scale[0] > 1 {
		factor = scale[0]
	}
	current := testutils.GenerateTestEndpointsWithDistribution(
		map[string]int{
			endpoint.RecordTypeA:     12 * factor,
			endpoint.RecordTypeAAAA:  27 * factor,
			endpoint.RecordTypeCNAME: 42 * factor,
			endpoint.RecordTypeSRV:   44 * factor,
		},
		map[string]int{
			"example.com": 1,
			"tld.org":     2,
			"open.net":    3,
		},
		map[string]int{"owner1": 1, "owner2": 1, "owner3": 1},
	)

	var desired []*endpoint.Endpoint
	for _, ep := range current {
		desired = append(desired, &endpoint.Endpoint{
			DNSName:    ep.DNSName,
			Targets:    endpoint.Targets{"txt-value"},
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  300,
		})
	}

	return &Plan{
		Policies:       []Policy{&SyncPolicy{}},
		Current:        current,
		Desired:        desired,
		ManagedRecords: endpoint.KnownRecordTypes,
		OwnerID:        "my-owner",
	}
}

func TestFlushMetrics(t *testing.T) {
	tests := []struct {
		name       string
		mismatches map[mismatchKey]float64
		expected   map[string]float64
	}{
		{
			name:       "empty mismatches",
			mismatches: map[mismatchKey]float64{},
			expected:   map[string]float64{},
		},
		{
			name: "single mismatch",
			mismatches: map[mismatchKey]float64{
				{recordType: endpoint.RecordTypeA, owner: "owner1", foreignOwner: "foreign1", domain: "example.com"}: 5,
			},
			expected: map[string]float64{
				"A|owner1|foreign1|example.com": 5,
			},
		},
		{
			name: "multiple mismatches",
			mismatches: map[mismatchKey]float64{
				{recordType: endpoint.RecordTypeA, owner: "owner1", foreignOwner: "foreign1", domain: "example.com"}:    3,
				{recordType: endpoint.RecordTypeCNAME, owner: "owner2", foreignOwner: "foreign2", domain: "test.org"}:   7,
				{recordType: endpoint.RecordTypeAAAA, owner: "owner1", foreignOwner: "foreign3", domain: "example.com"}: 2,
			},
			expected: map[string]float64{
				"A|owner1|foreign1|example.com":    3,
				"CNAME|owner2|foreign2|test.org":   7,
				"AAAA|owner1|foreign3|example.com": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset gauge before test
			registryOwnerMismatchPerSync.Gauge.Reset()

			pm := &planMetric{mismatches: tt.mismatches}
			pm.flushMetrics()

			// Verify each expected metric
			for key, expectedCount := range tt.expected {
				parts := strings.Split(key, "|")
				testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
					t,
					expectedCount,
					registryOwnerMismatchPerSync.Gauge,
					map[string]string{
						"record_type":   parts[0],
						"owner":         parts[1],
						"foreign_owner": parts[2],
						"domain":        parts[3],
					},
				)
			}
		})
	}
}

func TestFlushMetricsResetsGauge(t *testing.T) {
	// Pre-populate gauge with old data
	registryOwnerMismatchPerSync.Gauge.Reset()
	registryOwnerMismatchPerSync.AddWithLabels(10, endpoint.RecordTypeA, "old-owner", "old-foreign", "old.com")

	// Flush with new data
	pm := &planMetric{
		mismatches: map[mismatchKey]float64{
			{recordType: endpoint.RecordTypeCNAME, owner: "new-owner", foreignOwner: "new-foreign", domain: "new.com"}: 5,
		},
	}
	pm.flushMetrics()

	// Verify old metric is gone (gauge was reset)
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
		t,
		0,
		registryOwnerMismatchPerSync.Gauge,
		map[string]string{
			"record_type":   endpoint.RecordTypeA,
			"owner":         "old-owner",
			"foreign_owner": "old-foreign",
			"domain":        "old.com",
		},
	)

	// Verify new metric exists
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(
		t,
		5,
		registryOwnerMismatchPerSync.Gauge,
		map[string]string{
			"record_type":   endpoint.RecordTypeCNAME,
			"owner":         "new-owner",
			"foreign_owner": "new-foreign",
			"domain":        "new.com",
		},
	)
}
