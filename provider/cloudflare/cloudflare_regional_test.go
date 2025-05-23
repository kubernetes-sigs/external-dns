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

package cloudflare

import (
	"reflect"
	"slices"
	"testing"

	"github.com/cloudflare/cloudflare-go"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

func Test_regionalHostname(t *testing.T) {
	type args struct {
		endpoint         *endpoint.Endpoint
		defaultRegionKey string
	}
	tests := []struct {
		name string
		args args
		want cloudflare.RegionalHostname
	}{
		{
			name: "no region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
				},
				defaultRegionKey: "",
			},
			want: cloudflare.RegionalHostname{},
		},
		{
			name: "default region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "us",
			},
		},
		{
			name: "endpoint with region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "eu",
			},
		},
		{
			name: "endpoint with empty region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "",
						},
					},
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "",
			},
		},
		{
			name: "unsupported record type",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "TXT",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CloudFlareProvider{RegionKey: tt.args.defaultRegionKey}
			got := p.regionalHostname(tt.args.endpoint)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_dataLocalizationRegionalHostnamesChanges(t *testing.T) {
	cmpDataLocalizationRegionalHostnameChange := func(i, j regionalHostnameChange) int {
		if i.action == j.action {
			return 0
		}
		if i.Hostname < j.Hostname {
			return -1
		}
		return 1
	}
	type args struct {
		changes []*cloudFlareChange
	}
	tests := []struct {
		name    string
		args    args
		want    []regionalHostnameChange
		wantErr bool
	}{
		{
			name: "empty input",
			args: args{
				changes: []*cloudFlareChange{},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "changes without RegionalHostname",
			args: args{
				changes: []*cloudFlareChange{
					{
						Action: cloudFlareCreate,
						ResourceRecord: cloudflare.DNSRecord{
							Name: "example.com",
						},
						RegionalHostname: cloudflare.RegionalHostname{}, // Empty
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "change with empty RegionKey",
			args: args{
				changes: []*cloudFlareChange{
					{
						Action: cloudFlareCreate,
						ResourceRecord: cloudflare.DNSRecord{
							Name: "example.com",
						},
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "", // Empty region key
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "conflicting region keys",
			args: args{
				changes: []*cloudFlareChange{
					{
						Action: cloudFlareCreate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "eu",
						},
					},
					{
						Action: cloudFlareCreate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "us", // Different region key for same hostname
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "update takes precedence over create & delete",
			args: args{
				changes: []*cloudFlareChange{
					{
						Action: cloudFlareCreate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "eu",
						},
					},
					{
						Action: cloudFlareUpdate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "eu",
						},
					},
					{
						Action: cloudFlareDelete,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "eu",
						},
					},
				},
			},
			want: []regionalHostnameChange{
				{
					action: cloudFlareUpdate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example.com",
						RegionKey: "eu",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "create after delete becomes update",
			args: args{
				changes: []*cloudFlareChange{
					{
						Action: cloudFlareDelete,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "eu",
						},
					},
					{
						Action: cloudFlareCreate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example.com",
							RegionKey: "eu",
						},
					},
				},
			},
			want: []regionalHostnameChange{
				{
					action: cloudFlareUpdate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example.com",
						RegionKey: "eu",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "consolidate mixed actions for different hostnames",
			args: args{
				changes: []*cloudFlareChange{
					{
						Action: cloudFlareCreate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example1.com",
							RegionKey: "eu",
						},
					},
					{
						Action: cloudFlareUpdate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example2.com",
							RegionKey: "us",
						},
					},
					{
						Action: cloudFlareDelete,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example3.com",
							RegionKey: "ap",
						},
					},
					// duplicated actions
					{
						Action: cloudFlareCreate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example1.com",
							RegionKey: "eu",
						},
					},
					{
						Action: cloudFlareUpdate,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example2.com",
							RegionKey: "us",
						},
					},
					{
						Action: cloudFlareDelete,
						RegionalHostname: cloudflare.RegionalHostname{
							Hostname:  "example3.com",
							RegionKey: "ap",
						},
					},
				},
			},
			want: []regionalHostnameChange{
				{
					action: cloudFlareCreate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example1.com",
						RegionKey: "eu",
					},
				},
				{
					action: cloudFlareUpdate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example2.com",
						RegionKey: "us",
					},
				},
				{
					action: cloudFlareDelete,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example3.com",
						RegionKey: "ap",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataLocalizationRegionalHostnamesChanges(tt.args.changes)
			if (err != nil) != tt.wantErr {
				t.Errorf("dataLocalizationRegionalHostnamesChanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			slices.SortFunc(got, cmpDataLocalizationRegionalHostnameChange)
			slices.SortFunc(tt.want, cmpDataLocalizationRegionalHostnameChange)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataLocalizationRegionalHostnamesChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
