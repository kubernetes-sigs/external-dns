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

package registry

import (
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ Registry = &NoopRegistry{}

func TestNoopRegistry(t *testing.T) {
	t.Run("NewNoopRegistry", testNoopInit)
	t.Run("Records", testNoopRecords)
	t.Run("ApplyChanges", testNoopApplyChanges)
}

func testNoopInit(t *testing.T) {
	p := provider.NewInMemoryProvider()
	r, err := NewNoopRegistry(p)
	require.NoError(t, err)
	assert.Equal(t, p, r.provider)
}

func testNoopRecords(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone("org")
	providerRecords := []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			Target:     "example-lb.com",
			RecordType: "CNAME",
		},
	}
	p.ApplyChanges(&plan.Changes{
		Create: providerRecords,
	})

	r, _ := NewNoopRegistry(p)
	eps, err := r.Records()
	assert.NoError(t, err)
	assert.True(t, testutils.SameEndpoints(eps, providerRecords))
}

func testNoopApplyChanges(t *testing.T) {
	// do some prep
	p := provider.NewInMemoryProvider()
	p.CreateZone("org")
	providerRecords := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Target:  "old-lb.com",
		},
	}
	expectedUpdate := []*endpoint.Endpoint{
		{
			DNSName:    "example.org",
			Target:     "new-example-lb.com",
			RecordType: "CNAME",
		},
		{
			DNSName:    "new-record.org",
			Target:     "new-lb.org",
			RecordType: "CNAME",
		},
	}

	p.ApplyChanges(&plan.Changes{
		Create: providerRecords,
	})

	// wrong changes
	r, _ := NewNoopRegistry(p)
	err := r.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName: "example.org",
				Target:  "lb.com",
			},
		},
	})
	assert.EqualError(t, err, provider.ErrRecordAlreadyExists.Error())

	//correct changes
	err = r.ApplyChanges(&plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName: "new-record.org",
				Target:  "new-lb.org",
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName: "example.org",
				Target:  "new-example-lb.com",
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName: "example.org",
				Target:  "old-lb.com",
			},
		},
	})
	if err != nil {
		require.NoError(t, err)
	}
	res, _ := p.Records()
	assert.True(t, testutils.SameEndpoints(res, expectedUpdate))
}
