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

package plan

import (
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/stretchr/testify/suite"
)

type PlanTestSuite struct {
	suite.Suite
	fooV1Cname        *endpoint.Endpoint
	fooV2Cname        *endpoint.Endpoint
	fooV2CnameNoLabel *endpoint.Endpoint
	fooA5             *endpoint.Endpoint
	bar127A           *endpoint.Endpoint
	bar127AWithTTL    *endpoint.Endpoint
	bar192A           *endpoint.Endpoint
}

func (suite *PlanTestSuite) SetupTest() {
	suite.fooV1Cname = &endpoint.Endpoint{
		DNSName:    "foo",
		Target:     "v1",
		RecordType: "CNAME",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-v1",
			endpoint.OwnerLabelKey:    "pwner",
		},
	}
	suite.fooV2Cname = &endpoint.Endpoint{
		DNSName:    "foo",
		Target:     "v2",
		RecordType: "CNAME",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-v2",
		},
	}
	suite.fooV2CnameNoLabel = &endpoint.Endpoint{
		DNSName:    "foo",
		Target:     "v2",
		RecordType: "CNAME",
	}
	suite.fooA5 = &endpoint.Endpoint{
		DNSName:    "foo",
		Target:     "5.5.5.5",
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-5",
		},
	}
	suite.bar127A = &endpoint.Endpoint{
		DNSName:    "bar",
		Target:     "127.0.0.1",
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/bar-127",
		},
	}
	suite.bar127AWithTTL = &endpoint.Endpoint{
		DNSName:    "bar",
		Target:     "127.0.0.1",
		RecordType: "A",
		RecordTTL:  300,
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/bar-127",
		},
	}
	suite.bar192A = &endpoint.Endpoint{
		DNSName:    "bar",
		Target:     "192.168.0.1",
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/bar-192",
		},
	}
}

func (suite *PlanTestSuite) TestSyncFirstRound() {
	current := []*endpoint.Endpoint{}
	desired := []*endpoint.Endpoint{suite.fooV1Cname, suite.fooV2Cname, suite.bar127A}
	expectedCreate := []*endpoint.Endpoint{suite.fooV1Cname, suite.bar127A} //v1 is chosen because of resolver taking "min"
	expectedUpdateOld := []*endpoint.Endpoint{}
	expectedUpdateNew := []*endpoint.Endpoint{}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestSyncSecondRound() {
	current := []*endpoint.Endpoint{suite.fooV1Cname}
	desired := []*endpoint.Endpoint{suite.fooV2Cname, suite.fooV1Cname, suite.bar127A}
	expectedCreate := []*endpoint.Endpoint{suite.bar127A}
	expectedUpdateOld := []*endpoint.Endpoint{}
	expectedUpdateNew := []*endpoint.Endpoint{}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestSyncSecondRoundMigration() {
	current := []*endpoint.Endpoint{suite.fooV2CnameNoLabel}
	desired := []*endpoint.Endpoint{suite.fooV2Cname, suite.fooV1Cname, suite.bar127A}
	expectedCreate := []*endpoint.Endpoint{suite.bar127A}
	expectedUpdateOld := []*endpoint.Endpoint{suite.fooV2CnameNoLabel}
	expectedUpdateNew := []*endpoint.Endpoint{suite.fooV1Cname}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestSyncSecondRoundWithTTLChange() {
	current := []*endpoint.Endpoint{suite.bar127A}
	desired := []*endpoint.Endpoint{suite.bar127AWithTTL}
	expectedCreate := []*endpoint.Endpoint{}
	expectedUpdateOld := []*endpoint.Endpoint{suite.bar127A}
	expectedUpdateNew := []*endpoint.Endpoint{suite.bar127AWithTTL}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestSyncSecondRoundWithOwnerInherited() {
	current := []*endpoint.Endpoint{suite.fooV1Cname}
	desired := []*endpoint.Endpoint{suite.fooV2Cname}

	expectedCreate := []*endpoint.Endpoint{}
	expectedUpdateOld := []*endpoint.Endpoint{suite.fooV1Cname}
	expectedUpdateNew := []*endpoint.Endpoint{{
		DNSName:    suite.fooV2Cname.DNSName,
		Target:     suite.fooV2Cname.Target,
		RecordType: suite.fooV2Cname.RecordType,
		RecordTTL:  suite.fooV2Cname.RecordTTL,
		Labels: map[string]string{
			endpoint.ResourceLabelKey: suite.fooV2Cname.Labels[endpoint.ResourceLabelKey],
			endpoint.OwnerLabelKey:    suite.fooV1Cname.Labels[endpoint.OwnerLabelKey],
		},
	}}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestIdempotency() {
	current := []*endpoint.Endpoint{suite.fooV1Cname, suite.fooV2Cname}
	desired := []*endpoint.Endpoint{suite.fooV1Cname, suite.fooV2Cname}
	expectedCreate := []*endpoint.Endpoint{}
	expectedUpdateOld := []*endpoint.Endpoint{}
	expectedUpdateNew := []*endpoint.Endpoint{}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestDifferentTypes() {
	current := []*endpoint.Endpoint{suite.fooV1Cname}
	desired := []*endpoint.Endpoint{suite.fooV2Cname, suite.fooA5}
	expectedCreate := []*endpoint.Endpoint{}
	expectedUpdateOld := []*endpoint.Endpoint{suite.fooV1Cname}
	expectedUpdateNew := []*endpoint.Endpoint{suite.fooA5}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestRemoveEndpoint() {
	current := []*endpoint.Endpoint{suite.fooV1Cname, suite.bar192A}
	desired := []*endpoint.Endpoint{suite.fooV1Cname}
	expectedCreate := []*endpoint.Endpoint{}
	expectedUpdateOld := []*endpoint.Endpoint{}
	expectedUpdateNew := []*endpoint.Endpoint{}
	expectedDelete := []*endpoint.Endpoint{suite.bar192A}

	p := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func (suite *PlanTestSuite) TestRemoveEndpointWithUpsert() {
	current := []*endpoint.Endpoint{suite.fooV1Cname, suite.bar192A}
	desired := []*endpoint.Endpoint{suite.fooV1Cname}
	expectedCreate := []*endpoint.Endpoint{}
	expectedUpdateOld := []*endpoint.Endpoint{}
	expectedUpdateNew := []*endpoint.Endpoint{}
	expectedDelete := []*endpoint.Endpoint{}

	p := &Plan{
		Policies: []Policy{&UpsertOnlyPolicy{}},
		Current:  current,
		Desired:  desired,
	}

	changes := p.Calculate().Changes
	validateEntries(suite.T(), changes.Create, expectedCreate)
	validateEntries(suite.T(), changes.UpdateNew, expectedUpdateNew)
	validateEntries(suite.T(), changes.UpdateOld, expectedUpdateOld)
	validateEntries(suite.T(), changes.Delete, expectedDelete)
}

func TestPlan(t *testing.T) {
	suite.Run(t, new(PlanTestSuite))
}

//
//// TestCalculate tests that a plan can calculate actions to move a list of
//// current records to a list of desired records.
//func TestCalculate(t *testing.T) {
//	// empty list of records
//	empty := []*endpoint.Endpoint{}
//	// a simple entry
//	fooV1 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v1", endpoint.RecordTypeCNAME)}
//	// the same entry but with different target
//	fooV2 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", endpoint.RecordTypeCNAME)}
//	// another simple entry
//	bar := []*endpoint.Endpoint{endpoint.NewEndpoint("bar", "v1", endpoint.RecordTypeCNAME)}
//
//	// test case with labels
//	noLabels := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", endpoint.RecordTypeCNAME)}
//	labeledV2 := []*endpoint.Endpoint{newEndpointWithOwner("foo", "v2", "123")}
//	labeledV1 := []*endpoint.Endpoint{newEndpointWithOwner("foo", "v1", "123")}
//
//	// test case with type inheritance
//	noType := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", "")}
//	typedV2 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", endpoint.RecordTypeA)}
//	typedV1 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v1", endpoint.RecordTypeA)}
//
//	// test case with TTL
//	ttl := endpoint.TTL(300)
//	ttl2 := endpoint.TTL(50)
//	ttlV1 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v1", endpoint.RecordTypeCNAME, ttl)}
//	ttlV2 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v1", endpoint.RecordTypeCNAME)}
//	ttlV3 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v1", endpoint.RecordTypeCNAME, ttl)}
//	ttlV4 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v1", endpoint.RecordTypeCNAME, ttl2)}
//	ttlV5 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", endpoint.RecordTypeCNAME)}
//	ttlV6 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", endpoint.RecordTypeCNAME, ttl)}
//
//	for _, tc := range []struct {
//		policies                             []Policy
//		current, desired                     []*endpoint.Endpoint
//		create, updateOld, updateNew, delete []*endpoint.Endpoint
//	}{
//		// Nothing exists and nothing desired doesn't change anything.
//		{[]Policy{&SyncPolicy{}}, empty, empty, empty, empty, empty, empty},
//		// More desired than current creates the desired.
//		{[]Policy{&SyncPolicy{}}, empty, fooV1, fooV1, empty, empty, empty},
//		// Desired equals current doesn't change anything.
//		{[]Policy{&SyncPolicy{}}, fooV1, fooV1, empty, empty, empty, empty},
//		// Nothing is desired deletes the current.
//		{[]Policy{&SyncPolicy{}}, fooV1, empty, empty, empty, empty, fooV1},
//		// Current and desired match but Target is different triggers an update.
//		{[]Policy{&SyncPolicy{}}, fooV1, fooV2, empty, fooV1, fooV2, empty},
//		// Both exist but are different creates desired and deletes current.
//		{[]Policy{&SyncPolicy{}}, fooV1, bar, bar, empty, empty, fooV1},
//		// Nothing is desired but policy doesn't allow deletions.
//		{[]Policy{&UpsertOnlyPolicy{}}, fooV1, empty, empty, empty, empty, empty},
//		// Labels should be inherited
//		{[]Policy{&SyncPolicy{}}, labeledV1, noLabels, empty, labeledV1, labeledV2, empty},
//		// RecordType should be inherited
//		{[]Policy{&SyncPolicy{}}, typedV1, noType, empty, typedV1, typedV2, empty},
//		// If desired TTL is not configured, do not update
//		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV2, empty, empty, empty, empty},
//		// If desired TTL is configured but is the same as current TTL, do not update
//		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV3, empty, empty, empty, empty},
//		// If desired TTL is configured and is not the same as current TTL, need to update
//		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV4, empty, ttlV1, ttlV4, empty},
//		// If target changed and desired TTL is not configured, do not update TTL
//		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV5, empty, ttlV1, ttlV6, empty},
//	} {
//		// setup plan
//		plan := &Plan{
//			Policies: tc.policies,
//			Current:  tc.current,
//			Desired:  tc.desired,
//		}
//		// calculate actions
//		plan = plan.Calculate()
//
//		// validate actions
//		validateEntries(t, plan.Changes.Create, tc.create)
//		validateEntries(t, plan.Changes.UpdateOld, tc.updateOld)
//		validateEntries(t, plan.Changes.UpdateNew, tc.updateNew)
//		validateEntries(t, plan.Changes.Delete, tc.delete)
//	}
//}

// validateEntries validates that the list of entries matches expected.
func validateEntries(t *testing.T, entries, expected []*endpoint.Endpoint) {
	if !testutils.SameEndpoints(entries, expected) {
		t.Fatalf("expected %q to match %q", entries, expected)
	}
}
