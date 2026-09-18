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

package noop

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider/inmemory"
)

func TestNoopRegistry(t *testing.T) {
	t.Run("NewNoopRegistry", testNoopInit)
	t.Run("Records", testNoopRecords)
	t.Run("ApplyChanges", testNoopApplyChanges)
}

func TestNew(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	r, err := New(&externaldns.Config{}, p)
	require.NoError(t, err)
	assert.NotNil(t, r)
}

func TestNoopRegistry_OwnerID(t *testing.T) {
	r := newRegistry(inmemory.NewInMemoryProvider())
	assert.Empty(t, r.OwnerID())
}

func TestNoopRegistry_GetDomainFilter(t *testing.T) {
	p := inmemory.NewInMemoryProvider(inmemory.InMemoryWithDomain(endpoint.NewDomainFilter([]string{"example.com"})))
	r := newRegistry(p)
	assert.NotNil(t, r.GetDomainFilter())
}

func TestNoopRegistry_AdjustEndpoints(t *testing.T) {
	r := newRegistry(inmemory.NewInMemoryProvider())
	eps := []*endpoint.Endpoint{endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4")}
	got, err := r.AdjustEndpoints(eps)
	require.NoError(t, err)
	assert.Equal(t, eps, got)
}

func testNoopInit(t *testing.T) {
	p := inmemory.NewInMemoryProvider()
	r := newRegistry(p)
	var err error
	require.NoError(t, err)
	assert.Equal(t, p, r.provider)
}

func testNoopRecords(t *testing.T) {
	ctx := t.Context()
	p := inmemory.NewInMemoryProvider()
	p.CreateZone("org")
	inmemoryRecords := []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			Targets:    endpoint.Targets{"example-lb.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}
	p.ApplyChanges(ctx, &plan.Changes{
		Create: inmemoryRecords,
	})

	r := newRegistry(p)

	eps, err := r.Records(ctx)
	require.NoError(t, err)
	assert.True(t, testutils.SameEndpoints(eps, inmemoryRecords))
}

func testNoopApplyChanges(t *testing.T) {
	// do some prep
	p := inmemory.NewInMemoryProvider()
	p.CreateZone("org")
	inmemoryRecords := []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			Targets:    endpoint.Targets{"old-lb.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}
	expectedUpdate := []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			Targets:    endpoint.Targets{"new-example-lb.com"},
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "new-record.org",
			Targets:    endpoint.Targets{"new-lb.org"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	ctx := t.Context()
	p.ApplyChanges(ctx, &plan.Changes{
		Create: inmemoryRecords,
	})

	// wrong changes
	r := newRegistry(p)
	err := r.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "example.org",
				Targets:    endpoint.Targets{"lb.com"},
				RecordType: endpoint.RecordTypeCNAME,
			},
		},
	})
	assert.EqualError(t, err, inmemory.ErrRecordAlreadyExists.Error())

	// correct changes
	require.NoError(t, r.ApplyChanges(ctx, &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "new-record.org",
				Targets:    endpoint.Targets{"new-lb.org"},
				RecordType: endpoint.RecordTypeCNAME,
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "example.org",
				Targets:    endpoint.Targets{"new-example-lb.com"},
				RecordType: endpoint.RecordTypeCNAME,
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "example.org",
				Targets:    endpoint.Targets{"old-lb.com"},
				RecordType: endpoint.RecordTypeCNAME,
			},
		},
	}))
	res, _ := p.Records(ctx)
	assert.True(t, testutils.SameEndpoints(res, expectedUpdate))
}
