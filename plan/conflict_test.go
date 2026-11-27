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
	"maps"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	logtest "sigs.k8s.io/external-dns/internal/testutils/log"

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

func TestPerResource_ResolveRecordTypes_LogsWarning(t *testing.T) {
	const warnMsg = "contains conflicting record type candidates; discarding CNAME record"

	cname := &endpoint.Endpoint{DNSName: "foo", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"v1"}}
	a := &endpoint.Endpoint{DNSName: "foo", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.5.5.5"}}
	aaaa := &endpoint.Endpoint{DNSName: "foo", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1"}}

	tests := []struct {
		name        string
		row         *planTableRow
		expectsWarn bool
	}{
		{
			name: "no warning: cname only",
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{cname},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeCNAME: {candidates: []*endpoint.Endpoint{cname}},
				},
			},
		},
		{
			name: "no warning: a and aaaa only",
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{a, aaaa},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeA:    {candidates: []*endpoint.Endpoint{a}},
					endpoint.RecordTypeAAAA: {candidates: []*endpoint.Endpoint{aaaa}},
				},
			},
		},
		{
			name: "warning: cname conflicts with a record",
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{cname, a},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeCNAME: {candidates: []*endpoint.Endpoint{cname}},
					endpoint.RecordTypeA:     {candidates: []*endpoint.Endpoint{a}},
				},
			},
			expectsWarn: true,
		},
		{
			name: "warning: cname conflicts with a and aaaa records",
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{cname, a, aaaa},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeCNAME: {candidates: []*endpoint.Endpoint{cname}},
					endpoint.RecordTypeA:     {candidates: []*endpoint.Endpoint{a}},
					endpoint.RecordTypeAAAA:  {candidates: []*endpoint.Endpoint{aaaa}},
				},
			},
			expectsWarn: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)
			PerResource{}.ResolveRecordTypes(planKey{dnsName: "foo"}, tt.row)
			if tt.expectsWarn {
				logtest.TestHelperLogContainsWithLogLevel(warnMsg, log.WarnLevel, hook, t)
			} else {
				logtest.TestHelperLogNotContains(warnMsg, hook, t)
			}
		})
	}
}

func TestConflictResolver(t *testing.T) {
	suite.Run(t, new(ResolverSuite))
}

func TestPerResourceResolveRecordTypesApexCNAMEFlattening(t *testing.T) {
	cname := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"origin.example.net"}}
	txt := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"google-site-verification=abc"}}
	mx := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.com"}}
	a := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.0.2.1"}}
	aaaa := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}}
	ns := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeNS, Targets: endpoint.Targets{"ns.example.net"}}
	srv := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 0 443 origin.example.net"}}
	wwwCNAME := &endpoint.Endpoint{DNSName: "www.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"origin.example.net"}}
	wwwTXT := &endpoint.Endpoint{DNSName: "www.example.com", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"v=spf1"}}

	row := func(candidates []*endpoint.Endpoint, current map[string]*endpoint.Endpoint) *planTableRow {
		records := map[string]*domainEndpoints{}
		for _, ep := range candidates {
			recs := records[ep.RecordType]
			if recs == nil {
				recs = &domainEndpoints{}
				records[ep.RecordType] = recs
			}
			recs.candidates = append(recs.candidates, ep)
		}
		for recordType, ep := range current {
			recs := records[recordType]
			if recs == nil {
				recs = &domainEndpoints{}
				records[recordType] = recs
			}
			recs.current = ep
		}
		return &planTableRow{candidates: candidates, records: records}
	}

	discardedCNAME := func(current *endpoint.Endpoint, kept map[string]*domainEndpoints) map[string]*domainEndpoints {
		out := map[string]*domainEndpoints{
			endpoint.RecordTypeCNAME: {current: current, candidates: []*endpoint.Endpoint{}},
		}
		maps.Copy(out, kept)
		return out
	}

	tests := []struct {
		name     string
		resolver PerResource
		key      planKey
		row      *planTableRow
		want     map[string]*domainEndpoints
	}{
		{
			name:     "default drops CNAME beside TXT",
			resolver: PerResource{},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, txt}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: cname}),
			want: discardedCNAME(cname, map[string]*domainEndpoints{
				endpoint.RecordTypeTXT: {candidates: []*endpoint.Endpoint{txt}},
			}),
		},
		{
			name:     "default drops CNAME beside MX",
			resolver: PerResource{},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, mx}, nil),
			want: discardedCNAME(nil, map[string]*domainEndpoints{
				endpoint.RecordTypeMX: {candidates: []*endpoint.Endpoint{mx}},
			}),
		},
		{
			name:     "flattening keeps CNAME beside TXT and MX at the apex",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, txt, mx}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: cname}),
			want:     row([]*endpoint.Endpoint{cname, txt, mx}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: cname}).records,
		},
		{
			name:     "flattening matches apex case and trailing dot",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"Example.COM."}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, txt}, nil),
			want:     row([]*endpoint.Endpoint{cname, txt}, nil).records,
		},
		{
			name:     "flattening keeps SRV beside the apex CNAME",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, srv}, nil),
			want:     row([]*endpoint.Endpoint{cname, srv}, nil).records,
		},
		{
			name:     "flattening still drops CNAME beside A",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, a}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: cname}),
			want: discardedCNAME(cname, map[string]*domainEndpoints{
				endpoint.RecordTypeA: {candidates: []*endpoint.Endpoint{a}},
			}),
		},
		{
			name:     "flattening still drops CNAME beside AAAA",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, aaaa}, nil),
			want: discardedCNAME(nil, map[string]*domainEndpoints{
				endpoint.RecordTypeAAAA: {candidates: []*endpoint.Endpoint{aaaa}},
			}),
		},
		{
			name:     "flattening drops CNAME when TXT shares the name with A",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, txt, a}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: cname}),
			want: discardedCNAME(cname, map[string]*domainEndpoints{
				endpoint.RecordTypeTXT: {candidates: []*endpoint.Endpoint{txt}},
				endpoint.RecordTypeA:   {candidates: []*endpoint.Endpoint{a}},
			}),
		},
		{
			name:     "flattening still drops CNAME beside NS",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, ns}, nil),
			want: discardedCNAME(nil, map[string]*domainEndpoints{
				endpoint.RecordTypeNS: {candidates: []*endpoint.Endpoint{ns}},
			}),
		},
		{
			name:     "flattening does not apply off the apex",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			key:      planKey{dnsName: "www.example.com"},
			row:      row([]*endpoint.Endpoint{wwwCNAME, wwwTXT}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: wwwCNAME}),
			want: discardedCNAME(wwwCNAME, map[string]*domainEndpoints{
				endpoint.RecordTypeTXT: {candidates: []*endpoint.Endpoint{wwwTXT}},
			}),
		},
		{
			name:     "flattening without apex names keeps RFC 1034",
			resolver: PerResource{ApexCNAMEFlattening: true},
			key:      planKey{dnsName: "example.com"},
			row:      row([]*endpoint.Endpoint{cname, txt}, map[string]*endpoint.Endpoint{endpoint.RecordTypeCNAME: cname}),
			want: discardedCNAME(cname, map[string]*domainEndpoints{
				endpoint.RecordTypeTXT: {candidates: []*endpoint.Endpoint{txt}},
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.resolver.ResolveRecordTypes(tt.key, tt.row)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveRecordTypes() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestPerResourceResolveRecordTypesApexCNAMEFlatteningLogs(t *testing.T) {
	const warnMsg = "contains conflicting record type candidates; discarding CNAME record"
	const keepMsg = "keeps CNAME alongside other record types at a flattened zone apex"

	cname := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"origin.example.net"}}
	txt := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"v=spf1"}}
	a := &endpoint.Endpoint{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.0.2.1"}}

	tests := []struct {
		name         string
		resolver     PerResource
		row          *planTableRow
		expectsWarn  bool
		expectsDebug bool
	}{
		{
			name:     "default warns for CNAME and TXT",
			resolver: PerResource{},
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{cname, txt},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeCNAME: {candidates: []*endpoint.Endpoint{cname}},
					endpoint.RecordTypeTXT:   {candidates: []*endpoint.Endpoint{txt}},
				},
			},
			expectsWarn: true,
		},
		{
			name:     "flattening does not warn when the apex CNAME is kept",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{cname, txt},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeCNAME: {candidates: []*endpoint.Endpoint{cname}},
					endpoint.RecordTypeTXT:   {candidates: []*endpoint.Endpoint{txt}},
				},
			},
			expectsDebug: true,
		},
		{
			name:     "flattening still warns for CNAME and A",
			resolver: PerResource{ApexCNAMEFlattening: true, ZoneApexes: []string{"example.com"}},
			row: &planTableRow{
				candidates: []*endpoint.Endpoint{cname, a},
				records: map[string]*domainEndpoints{
					endpoint.RecordTypeCNAME: {candidates: []*endpoint.Endpoint{cname}},
					endpoint.RecordTypeA:     {candidates: []*endpoint.Endpoint{a}},
				},
			},
			expectsWarn: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)
			tt.resolver.ResolveRecordTypes(planKey{dnsName: "example.com"}, tt.row)
			if tt.expectsWarn {
				logtest.TestHelperLogContainsWithLogLevel(warnMsg, log.WarnLevel, hook, t)
			} else {
				logtest.TestHelperLogNotContains(warnMsg, hook, t)
			}
			if tt.expectsDebug {
				logtest.TestHelperLogContainsWithLogLevel(keepMsg, log.DebugLevel, hook, t)
			}
		})
	}
}
