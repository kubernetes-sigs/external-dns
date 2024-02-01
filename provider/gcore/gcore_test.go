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

package gcore

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	gdns "github.com/G-Core/gcore-dns-sdk-go"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type dnsManagerMock struct {
	addZoneRRSet      func(ctx context.Context, zone, recordName, recordType string, values []gdns.ResourceRecord, ttl int) error
	zonesWithRecords  func(ctx context.Context, filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error)
	zones             func(ctx context.Context, filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error)
	deleteRRSetRecord func(ctx context.Context, zone, name, recordType string, contents ...string) error
}

func (d dnsManagerMock) AddZoneRRSet(ctx context.Context,
	zone, recordName, recordType string,
	values []gdns.ResourceRecord, ttl int, _ ...gdns.AddZoneOpt) error {
	return d.addZoneRRSet(ctx, zone, recordName, recordType, values, ttl)
}
func (d dnsManagerMock) ZonesWithRecords(ctx context.Context, filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
	return d.zonesWithRecords(ctx, filters...)
}
func (d dnsManagerMock) Zones(ctx context.Context, filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
	return d.zones(ctx, filters...)
}
func (d dnsManagerMock) DeleteRRSetRecord(ctx context.Context, zone, name, recordType string, contents ...string) error {
	return d.deleteRRSetRecord(ctx, zone, name, recordType, contents...)
}

func Test_dnsProvider_Records(t *testing.T) {
	type fields struct {
		domainFilter endpoint.DomainFilter
		client       dnsManager
		dryRun       bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []endpoint.Endpoint
		wantErr bool
	}{
		{
			name: "no_filter",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zonesWithRecords: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{
							{
								Name: "example.com",
								Records: []gdns.ZoneRecord{
									{
										Name:         "test.example.com",
										Type:         "A",
										TTL:          10,
										ShortAnswers: []string{"1.1.1.1"},
									},
								},
							},
						}, nil
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []endpoint.Endpoint{
				*endpoint.NewEndpointWithTTL(
					"test.example.com", "A", endpoint.TTL(10), []string{"1.1.1.1"}...),
			},
			wantErr: false,
		},
		{
			name: "filtered",
			fields: fields{
				domainFilter: endpoint.DomainFilter{Filters: []string{"example.com"}},
				client: dnsManagerMock{
					zonesWithRecords: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						zoneFilter := &gdns.ZonesFilter{}
						for _, op := range filters {
							op(zoneFilter)
						}
						val := []gdns.Zone{
							{
								Name: "example.com",
								Records: []gdns.ZoneRecord{
									{
										Name:         "test.example.com",
										Type:         "A",
										TTL:          10,
										ShortAnswers: []string{"1.1.1.1"},
									},
								},
							},
						}
						res := make([]gdns.Zone, 0)
						for _, v := range val {
							for _, name := range zoneFilter.Names {
								if name == v.Name {
									res = append(res, v)
								}
							}
						}
						return res, nil
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []endpoint.Endpoint{
				*endpoint.NewEndpointWithTTL(
					"test.example.com", "A", endpoint.TTL(10), []string{"1.1.1.1"}...),
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				domainFilter: endpoint.DomainFilter{Filters: []string{"example.com"}},
				client: dnsManagerMock{
					zonesWithRecords: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return nil, fmt.Errorf("test")
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dnsProvider{
				domainFilter: tt.fields.domainFilter,
				client:       tt.fields.client,
				dryRun:       tt.fields.dryRun,
			}
			got, err := p.Records(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Records() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			toCompare := make([]endpoint.Endpoint, len(got))
			for i, e := range got {
				toCompare[i] = *e
			}
			if !reflect.DeepEqual(toCompare, tt.want) {
				t.Errorf("Records() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnsProvider_GetDomainFilter(t *testing.T) {
	type fields struct {
		domainFilter endpoint.DomainFilter
		client       dnsManager
		dryRun       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   endpoint.DomainFilter
	}{
		{
			name: "not_empty",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "example.com"}}, nil
					},
				},
				dryRun: false,
			},
			want: endpoint.NewDomainFilter([]string{"example.com", ".example.com"}),
		},
		{
			name: "empty",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{}, nil
					},
				},
				dryRun: false,
			},
			want: endpoint.NewDomainFilter([]string{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dnsProvider{
				domainFilter: tt.fields.domainFilter,
				client:       tt.fields.client,
				dryRun:       tt.fields.dryRun,
			}
			if got := p.GetDomainFilter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDomainFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dnsProvider_ApplyChanges(t *testing.T) {
	type fields struct {
		domainFilter endpoint.DomainFilter
		client       dnsManager
		dryRun       bool
	}
	type args struct {
		ctx     context.Context
		changes *plan.Changes
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete exist in filter",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "test.com"}}, nil
					},
					deleteRRSetRecord: func(ctx context.Context, zone, name, recordType string, contents ...string) error {
						if zone == "test.com" && name == "my.test.com" && recordType == "A" && contents[0] == "1.1.1.1" {
							return nil
						}
						return fmt.Errorf("deleteRRSetRecord wrong params")
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
				changes: &plan.Changes{
					Delete: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL("my.test.com", "A", 10, "1.1.1.1"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "delete exist in filter with left dot",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "test.com"}}, nil
					},
					deleteRRSetRecord: func(ctx context.Context, zone, name, recordType string, contents ...string) error {
						if zone == "test.com" && name == ".my.test.com" && recordType == "A" && contents[0] == "1.1.1.1" {
							return nil
						}
						return fmt.Errorf("deleteRRSetRecord wrong params")
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
				changes: &plan.Changes{
					Delete: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL(".my.test.com", "A", 10, "1.1.1.1"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "delete not exist in filter",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "test.com"}}, nil
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
				changes: &plan.Changes{
					Delete: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL("my.t.com", "A", 10, "1.1.1.1"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "delete error",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "test.com"}}, nil
					},
					deleteRRSetRecord: func(ctx context.Context, zone, name, recordType string, contents ...string) error {
						if zone == "test.com" && name == "my.test.com" && recordType == "A" && contents[0] == "1.1.1.1" {
							return nil
						}
						return fmt.Errorf("deleteRRSetRecord wrong params")
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
				changes: &plan.Changes{
					Delete: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL("my1.test.com", "A", 10, "1.1.1.1"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "create ok",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "test.com"}}, nil
					},
					addZoneRRSet: func(ctx context.Context, zone, recordName, recordType string, values []gdns.ResourceRecord, ttl int) error {
						if zone == "test.com" &&
							ttl == 10 &&
							recordName == "my.test.com" &&
							recordType == "A" &&
							values[0].Content[0] == "1.1.1.1" {
							return nil
						}
						return fmt.Errorf("addZoneRRSet wrong params")
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL("my.test.com", "A", 10, "1.1.1.1"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "update ok",
			fields: fields{
				domainFilter: endpoint.DomainFilter{},
				client: dnsManagerMock{
					zones: func(ctx context.Context,
						filters ...func(zone *gdns.ZonesFilter)) ([]gdns.Zone, error) {
						return []gdns.Zone{{Name: "test.com"}}, nil
					},
					deleteRRSetRecord: func(ctx context.Context, zone, name, recordType string, contents ...string) error {
						if zone == "test.com" && name == "my.test.com" && recordType == "A" && contents[0] == "1.1.1.1" {
							return nil
						}
						return fmt.Errorf("deleteRRSetRecord wrong params: %s %s %s %+v",
							zone, name, recordType, contents)
					},
					addZoneRRSet: func(ctx context.Context, zone, recordName, recordType string, values []gdns.ResourceRecord, ttl int) error {
						if zone == "test.com" &&
							ttl == 10 &&
							recordName == "my.test.com" &&
							recordType == "A" &&
							values[0].Content[0] == "1.2.3.4" {
							return nil
						}
						return fmt.Errorf("addZoneRRSet wrong params")
					},
				},
				dryRun: false,
			},
			args: args{
				ctx: context.Background(),
				changes: &plan.Changes{
					UpdateOld: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL("my.test.com", "A", 10, "1.1.1.1"),
						endpoint.NewEndpointWithTTL("my1.test.com", "A", 10, "1.1.1.2"),
					},
					UpdateNew: []*endpoint.Endpoint{
						endpoint.NewEndpointWithTTL("my.test.com", "A", 10, "1.2.3.4"),
						endpoint.NewEndpointWithTTL("my1.test.com", "A", 10, "1.1.1.2"),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dnsProvider{
				domainFilter: tt.fields.domainFilter,
				client:       tt.fields.client,
				dryRun:       tt.fields.dryRun,
			}
			if err := p.ApplyChanges(tt.args.ctx, tt.args.changes); (err != nil) != tt.wantErr {
				t.Errorf("ApplyChanges() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_extractAllZones(t *testing.T) {
	type args struct {
		dnsName string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "3th level",
			args: args{
				dnsName: "my.test.com",
			},
			want: []string{"my.test.com", "test.com"},
		},
		{
			name: "with dots",
			args: args{
				dnsName: ".my.test.com.",
			},
			want: []string{"my.test.com", "test.com"},
		},
		{
			name: "2d level",
			args: args{
				dnsName: "test.com",
			},
			want: []string{"test.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractAllZones(tt.args.dnsName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractAllZones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNonExistingTargets(t *testing.T) {
	type args struct {
		existing         *endpoint.Endpoint
		toCompare        []*endpoint.Endpoint
		diffFromExisting bool
	}
	tests := []struct {
		name string
		args args
		want endpoint.Targets
	}{
		{
			name: "not from existing",
			args: args{
				existing: endpoint.NewEndpointWithTTL(
					"my.test.com", "A", 10, "1.1.1.1", "1.2.2.2"),
				toCompare: []*endpoint.Endpoint{
					endpoint.NewEndpointWithTTL(
						"my.test.com", "A", 10, "1.1.1.1", "1.2.3.4"),
					endpoint.NewEndpointWithTTL(
						"my.test.com", "AAAA", 10, "1.1.1.1", "1.3.3.4"),
					endpoint.NewEndpointWithTTL(
						"no.test.com", "A", 10, "1.1.1.1", "1.3.3.4"),
				},
				diffFromExisting: false,
			},
			want: endpoint.Targets{"1.2.3.4"},
		},
		{
			name: "from existing",
			args: args{
				existing: endpoint.NewEndpointWithTTL(
					"my.test.com", "A", 10, "1.1.1.1", "1.2.2.2"),
				toCompare: []*endpoint.Endpoint{
					endpoint.NewEndpointWithTTL(
						"my.test.com", "A", 10, "1.1.1.1", "1.2.3.4"),
					endpoint.NewEndpointWithTTL(
						"my.test.com", "AAAA", 10, "1.1.1.1", "1.3.3.4"),
					endpoint.NewEndpointWithTTL(
						"no.test.com", "A", 10, "1.1.1.1", "1.3.3.4"),
				},
				diffFromExisting: true,
			},
			want: endpoint.Targets{"1.2.2.2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unexistingTargets(tt.args.existing, tt.args.toCompare, tt.args.diffFromExisting); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unexistingTargets() = %v, want %v", got, tt.want)
			}
		})
	}
}
