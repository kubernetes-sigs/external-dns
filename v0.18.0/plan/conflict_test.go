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
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"sigs.k8s.io/external-dns/endpoint"
)

var _ ConflictResolver = PerResource{}

type ResolverSuite struct {
	// resolvers
	perResource PerResource
	// endpoints
	fooV1Cname          *endpoint.Endpoint
	fooV2Cname          *endpoint.Endpoint
	fooV2CnameDuplicate *endpoint.Endpoint
	fooA5               *endpoint.Endpoint
	fooAAAA5            *endpoint.Endpoint
	bar127A             *endpoint.Endpoint
	bar192A             *endpoint.Endpoint
	bar127AAnother      *endpoint.Endpoint
	legacyBar192A       *endpoint.Endpoint // record created in AWS now without resource label
	suite.Suite
}

func (suite *ResolverSuite) SetupTest() {
	suite.perResource = PerResource{}
	// initialize endpoints used in tests
	suite.fooV1Cname = &endpoint.Endpoint{
		DNSName:    "foo",
		Targets:    endpoint.Targets{"v1"},
		RecordType: "CNAME",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-v1",
		},
	}
	suite.fooV2Cname = &endpoint.Endpoint{
		DNSName:    "foo",
		Targets:    endpoint.Targets{"v2"},
		RecordType: "CNAME",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-v2",
		},
	}
	suite.fooV2CnameDuplicate = &endpoint.Endpoint{
		DNSName:    "foo",
		Targets:    endpoint.Targets{"v2"},
		RecordType: "CNAME",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-v2-duplicate",
		},
	}
	suite.fooA5 = &endpoint.Endpoint{
		DNSName:    "foo",
		Targets:    endpoint.Targets{"5.5.5.5"},
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-5",
		},
	}
	suite.fooAAAA5 = &endpoint.Endpoint{
		DNSName:    "foo",
		Targets:    endpoint.Targets{"2001:DB8::1"},
		RecordType: "AAAA",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/foo-5",
		},
	}
	suite.bar127A = &endpoint.Endpoint{
		DNSName:    "bar",
		Targets:    endpoint.Targets{"127.0.0.1"},
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/bar-127",
		},
	}
	suite.bar127AAnother = &endpoint.Endpoint{ // TODO: remove this once we move to multiple targets under same endpoint
		DNSName:    "bar",
		Targets:    endpoint.Targets{"8.8.8.8"},
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/bar-127",
		},
	}
	suite.bar192A = &endpoint.Endpoint{
		DNSName:    "bar",
		Targets:    endpoint.Targets{"192.168.0.1"},
		RecordType: "A",
		Labels: map[string]string{
			endpoint.ResourceLabelKey: "ingress/default/bar-192",
		},
	}
	suite.legacyBar192A = &endpoint.Endpoint{
		DNSName:    "bar",
		Targets:    endpoint.Targets{"192.168.0.1"},
		RecordType: "A",
	}
}

func (suite *ResolverSuite) TestStrictResolver() {
	// test that perResource resolver picks min for create list
	suite.Equal(suite.bar127A, suite.perResource.ResolveCreate([]*endpoint.Endpoint{suite.bar127A, suite.bar192A}), "should pick min one")
	suite.Equal(suite.fooA5, suite.perResource.ResolveCreate([]*endpoint.Endpoint{suite.fooA5, suite.fooV1Cname}), "should pick min one")
	suite.Equal(suite.fooV1Cname, suite.perResource.ResolveCreate([]*endpoint.Endpoint{suite.fooV2Cname, suite.fooV1Cname}), "should pick min one")

	// test that perResource resolver preserves resource if it still exists
	suite.Equal(suite.bar127AAnother, suite.perResource.ResolveUpdate(suite.bar127A, []*endpoint.Endpoint{suite.bar127AAnother, suite.bar127A}), "should pick min for update when same resource endpoint occurs multiple times (remove after multiple-target support") // TODO:remove this test
	suite.Equal(suite.bar127A, suite.perResource.ResolveUpdate(suite.bar127A, []*endpoint.Endpoint{suite.bar192A, suite.bar127A}), "should pick existing resource")
	suite.Equal(suite.fooV2Cname, suite.perResource.ResolveUpdate(suite.fooV2Cname, []*endpoint.Endpoint{suite.fooV2Cname, suite.fooV2CnameDuplicate}), "should pick existing resource even if targets are same")
	suite.Equal(suite.fooA5, suite.perResource.ResolveUpdate(suite.fooV1Cname, []*endpoint.Endpoint{suite.fooA5, suite.fooV2Cname}), "should pick new if resource was deleted")
	// should actually get the updated record (note ttl is different)
	newFooV1Cname := &endpoint.Endpoint{
		DNSName:    suite.fooV1Cname.DNSName,
		Targets:    suite.fooV1Cname.Targets,
		Labels:     suite.fooV1Cname.Labels,
		RecordType: suite.fooV1Cname.RecordType,
		RecordTTL:  suite.fooV1Cname.RecordTTL + 1, // ttl is different
	}
	suite.Equal(newFooV1Cname, suite.perResource.ResolveUpdate(suite.fooV1Cname, []*endpoint.Endpoint{suite.fooA5, suite.fooV2Cname, newFooV1Cname}), "should actually pick same resource with updates")

	// legacy record's resource value will not match any candidates resource label
	// therefore pick minimum again
	suite.Equal(suite.bar127A, suite.perResource.ResolveUpdate(suite.legacyBar192A, []*endpoint.Endpoint{suite.bar127A, suite.bar192A}), " legacy record's resource value will not match, should pick minimum")
}

func (suite *ResolverSuite) TestPerResource_ResolveRecordTypes() {
	type args struct {
		key planKey
		row *planTableRow
	}
	tests := []struct {
		name string
		args args
		want map[string]*domainEndpoints
	}{
		{
			name: "no conflict: cname record",
			args: args{
				key: planKey{dnsName: "foo"},
				row: &planTableRow{
					candidates: []*endpoint.Endpoint{suite.fooV1Cname},
					records: map[string]*domainEndpoints{
						endpoint.RecordTypeCNAME: {
							candidates: []*endpoint.Endpoint{suite.fooV1Cname},
						},
					},
				},
			},
			want: map[string]*domainEndpoints{
				endpoint.RecordTypeCNAME: {
					candidates: []*endpoint.Endpoint{suite.fooV1Cname},
				},
			},
		},
		{
			name: "no conflict: a record",
			args: args{
				key: planKey{dnsName: "foo"},
				row: &planTableRow{
					current:    []*endpoint.Endpoint{suite.fooA5},
					candidates: []*endpoint.Endpoint{suite.fooA5},
					records: map[string]*domainEndpoints{
						endpoint.RecordTypeA: {
							current:    suite.fooA5,
							candidates: []*endpoint.Endpoint{suite.fooA5},
						},
					},
				},
			},
			want: map[string]*domainEndpoints{
				endpoint.RecordTypeA: {
					current:    suite.fooA5,
					candidates: []*endpoint.Endpoint{suite.fooA5},
				},
			},
		},
		{
			name: "no conflict: a and aaaa records",
			args: args{
				key: planKey{dnsName: "foo"},
				row: &planTableRow{
					candidates: []*endpoint.Endpoint{suite.fooA5, suite.fooAAAA5},
					records: map[string]*domainEndpoints{
						endpoint.RecordTypeA: {
							candidates: []*endpoint.Endpoint{suite.fooA5},
						},
						endpoint.RecordTypeAAAA: {
							candidates: []*endpoint.Endpoint{suite.fooAAAA5},
						},
					},
				},
			},
			want: map[string]*domainEndpoints{
				endpoint.RecordTypeA: {
					candidates: []*endpoint.Endpoint{suite.fooA5},
				},
				endpoint.RecordTypeAAAA: {
					candidates: []*endpoint.Endpoint{suite.fooAAAA5},
				},
			},
		},
		{
			name: "conflict: cname and a records",
			args: args{
				key: planKey{dnsName: "foo"},
				row: &planTableRow{
					current:    []*endpoint.Endpoint{suite.fooV1Cname},
					candidates: []*endpoint.Endpoint{suite.fooV1Cname, suite.fooA5},
					records: map[string]*domainEndpoints{
						endpoint.RecordTypeCNAME: {
							current:    suite.fooV1Cname,
							candidates: []*endpoint.Endpoint{suite.fooV1Cname},
						},
						endpoint.RecordTypeA: {
							candidates: []*endpoint.Endpoint{suite.fooA5},
						},
					},
				},
			},
			want: map[string]*domainEndpoints{
				endpoint.RecordTypeCNAME: {
					current:    suite.fooV1Cname,
					candidates: []*endpoint.Endpoint{},
				},
				endpoint.RecordTypeA: {
					candidates: []*endpoint.Endpoint{suite.fooA5},
				},
			},
		},
		{
			name: "conflict: cname, a, and aaaa records",
			args: args{
				key: planKey{dnsName: "foo"},
				row: &planTableRow{
					current:    []*endpoint.Endpoint{suite.fooA5, suite.fooAAAA5},
					candidates: []*endpoint.Endpoint{suite.fooV1Cname, suite.fooA5, suite.fooAAAA5},
					records: map[string]*domainEndpoints{
						endpoint.RecordTypeCNAME: {
							candidates: []*endpoint.Endpoint{suite.fooV1Cname},
						},
						endpoint.RecordTypeA: {
							current:    suite.fooA5,
							candidates: []*endpoint.Endpoint{suite.fooA5},
						},
						endpoint.RecordTypeAAAA: {
							current:    suite.fooAAAA5,
							candidates: []*endpoint.Endpoint{suite.fooAAAA5},
						},
					},
				},
			},
			want: map[string]*domainEndpoints{
				endpoint.RecordTypeCNAME: {
					candidates: []*endpoint.Endpoint{},
				},
				endpoint.RecordTypeA: {
					current:    suite.fooA5,
					candidates: []*endpoint.Endpoint{suite.fooA5},
				},
				endpoint.RecordTypeAAAA: {
					current:    suite.fooAAAA5,
					candidates: []*endpoint.Endpoint{suite.fooAAAA5},
				},
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if got := suite.perResource.ResolveRecordTypes(tt.args.key, tt.args.row); !reflect.DeepEqual(got, tt.want) {
				suite.T().Errorf("PerResource.ResolveRecordTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConflictResolver(t *testing.T) {
	suite.Run(t, new(ResolverSuite))
}
