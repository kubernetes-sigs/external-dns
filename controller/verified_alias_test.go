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

package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/registry/noop"
)

// aliasingProvider rewrites a CNAME into an A record in AdjustEndpoints, the
// way the AWS provider turns a CNAME to an ELB into an alias A record. Records
// returns the alias as an A record, as Route53 does.
type aliasingProvider struct {
	mockProvider
}

func (p *aliasingProvider) AdjustEndpoints(eps []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	out := make([]*endpoint.Endpoint, 0, len(eps))
	for _, ep := range eps {
		adjusted := ep.DeepCopy()
		if adjusted.RecordType == endpoint.RecordTypeCNAME {
			adjusted.RecordType = endpoint.RecordTypeA
		}
		out = append(out, adjusted)
	}
	return out, nil
}

func (p *aliasingProvider) ApplyChanges(_ context.Context, _ *plan.Changes) error { return nil }

func TestVerifiedRecordsCountsProviderAdjustedRecordType(t *testing.T) {
	source := new(testutils.MockSource)
	source.On("Endpoints").Return([]*endpoint.Endpoint{
		{DNSName: "*.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb-123.elb.us-east-1.amazonaws.com"}},
	}, nil)

	p := &aliasingProvider{mockProvider{RecordsStore: []*endpoint.Endpoint{
		{DNSName: "*.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"lb-123.elb.us-east-1.amazonaws.com"}},
	}}}
	r, err := noop.New(nil, p)
	require.NoError(t, err)

	ctrl := &Controller{
		Source:             source,
		Registry:           r,
		Policy:             &plan.SyncPolicy{},
		ManagedRecordTypes: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}
	require.NoError(t, ctrl.RunOnce(t.Context()))

	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, registryRecords.Gauge, map[string]string{"record_type": "a"})
	testutils.TestHelperVerifyMetricsGaugeVectorWithLabels(t, 1, verifiedRecords.Gauge, map[string]string{"record_type": "a"})
}
