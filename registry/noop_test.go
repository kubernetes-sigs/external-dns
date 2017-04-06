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
	if err != nil {
		t.Error(err)
	}
	if r.provider != p {
		t.Error("noop registry incorrectly initialized")
	}
}

func testNoopRecords(t *testing.T) {
	p := provider.NewInMemoryProvider()
	p.CreateZone("zone")
	providerRecords := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Target:  "example-lb.com",
		},
	}
	p.ApplyChanges("zone", &plan.Changes{
		Create: providerRecords,
	})

	r, err := NewNoopRegistry(p)
	_, err = r.Records("wrong-zone")
	if err == nil {
		t.Error("Should fail for wrong zone: wrong-zone")
	}

	r, err = NewNoopRegistry(p)
	eps, err := r.Records("zone")
	if !testutils.SameEndpoints(eps, providerRecords) {
		t.Error("incorrect result is returned")
	}
}

func testNoopApplyChanges(t *testing.T) {
	// do some prep
	p := provider.NewInMemoryProvider()
	p.CreateZone("zone")
	providerRecords := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Target:  "example-lb.com",
		},
	}
	expectedUpdate := []*endpoint.Endpoint{
		{
			DNSName: "example.org",
			Target:  "new-example-lb.com",
		},
		{
			DNSName: "new-record.org",
			Target:  "new-lb.org",
		},
	}

	p.ApplyChanges("zone", &plan.Changes{
		Create: providerRecords,
	})

	// wrong zone
	r, _ := NewNoopRegistry(p)
	err := r.ApplyChanges("wrong-zone", &plan.Changes{})
	if err != provider.ErrZoneNotFound {
		t.Error("should return zone not found for apply changes on wrong zone")
	}

	// wrong changes
	r, _ = NewNoopRegistry(p)
	err = r.ApplyChanges("zone", &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName: "example.org",
				Target:  "lb.com",
			},
		},
	})
	if err != provider.ErrRecordAlreadyExists {
		t.Error("should return record already exists")
	}

	//correct changes
	err = r.ApplyChanges("zone", &plan.Changes{
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
				Target:  "example-lb.com",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	res, _ := p.Records("zone")
	if !testutils.SameEndpoints(res, expectedUpdate) {
		t.Error("incorrectly updated dns provider")
	}
}
