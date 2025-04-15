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
	"context"
	"sort"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

func (m *mockCloudFlareClient) ListDataLocalizationRegionalHostnames(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.ListDataLocalizationRegionalHostnamesParams) ([]cloudflare.RegionalHostname, error) {
	return m.regionalHostnames[rc.Identifier], nil
}

func (m *mockCloudFlareClient) CreateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDataLocalizationRegionalHostnameParams) error {
	m.Actions = append(m.Actions, MockAction{
		Name:     "CreateDataLocalizationRegionalHostname",
		ZoneId:   rc.Identifier,
		RecordId: "",
		RegionalHostname: cloudflare.RegionalHostname{
			Hostname:  rp.Hostname,
			RegionKey: rp.RegionKey,
		},
	})
	return nil
}

func (m *mockCloudFlareClient) UpdateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDataLocalizationRegionalHostnameParams) error {
	m.Actions = append(m.Actions, MockAction{
		Name:     "UpdateDataLocalizationRegionalHostname",
		ZoneId:   rc.Identifier,
		RecordId: "",
		RegionalHostname: cloudflare.RegionalHostname{
			Hostname:  rp.Hostname,
			RegionKey: rp.RegionKey,
		},
	})
	return nil
}

func (m *mockCloudFlareClient) DeleteDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, hostname string) error {
	m.Actions = append(m.Actions, MockAction{
		Name:     "DeleteDataLocalizationRegionalHostname",
		ZoneId:   rc.Identifier,
		RecordId: "",
		RegionalHostname: cloudflare.RegionalHostname{
			Hostname: hostname,
		},
	})
	return nil
}

func TestCloudflareRegionalHostname(t *testing.T) {
	tests := []struct {
		name              string
		records           map[string]cloudflare.DNSRecord
		regionalHostnames []cloudflare.RegionalHostname
		endpoints         []*endpoint.Endpoint
		want              []MockAction
	}{
		{
			name:              "create",
			records:           map[string]cloudflare.DNSRecord{},
			regionalHostnames: []cloudflare.RegionalHostname{},
			endpoints: []*endpoint.Endpoint{
				{
					RecordType: "A",
					DNSName:    "create.bar.com",
					Targets:    endpoint.Targets{"127.0.0.1"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
			},
			want: []MockAction{
				{
					Name:     "Create",
					ZoneId:   "001",
					RecordId: generateDNSRecordID("A", "create.bar.com", "127.0.0.1"),
					RecordData: cloudflare.DNSRecord{
						ID:      generateDNSRecordID("A", "create.bar.com", "127.0.0.1"),
						Type:    "A",
						Name:    "create.bar.com",
						Content: "127.0.0.1",
						TTL:     1,
						Proxied: proxyDisabled,
					},
				},
				{
					Name:   "CreateDataLocalizationRegionalHostname",
					ZoneId: "001",
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "create.bar.com",
						RegionKey: "eu",
					},
				},
			},
		},
		{
			name: "Update",
			records: map[string]cloudflare.DNSRecord{
				"update.bar.com": {
					ID:      generateDNSRecordID("A", "update.bar.com", "127.0.0.1"),
					Type:    "A",
					Name:    "update.bar.com",
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: proxyDisabled,
				},
			},
			regionalHostnames: []cloudflare.RegionalHostname{
				{
					Hostname:  "update.bar.com",
					RegionKey: "us",
				},
			},
			endpoints: []*endpoint.Endpoint{
				{
					RecordType: "A",
					DNSName:    "update.bar.com",
					Targets:    endpoint.Targets{"127.0.0.1"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
			},
			want: []MockAction{
				{
					Name:     "Update",
					ZoneId:   "001",
					RecordId: generateDNSRecordID("A", "update.bar.com", "127.0.0.1"),
					RecordData: cloudflare.DNSRecord{
						ID:      generateDNSRecordID("A", "update.bar.com", "127.0.0.1"),
						Type:    "A",
						Name:    "update.bar.com",
						Content: "127.0.0.1",
						TTL:     1,
						Proxied: proxyDisabled,
					},
				},
				{
					Name:   "UpdateDataLocalizationRegionalHostname",
					ZoneId: "001",
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "update.bar.com",
						RegionKey: "eu",
					},
				},
			},
		},
		{
			name: "Delete",
			records: map[string]cloudflare.DNSRecord{
				"update.bar.com": {
					ID:      generateDNSRecordID("A", "delete.bar.com", "127.0.0.1"),
					Type:    "A",
					Name:    "delete.bar.com",
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: proxyDisabled,
				},
			},
			regionalHostnames: []cloudflare.RegionalHostname{
				{
					Hostname:  "delete.bar.com",
					RegionKey: "us",
				},
			},
			endpoints: []*endpoint.Endpoint{},
			want: []MockAction{
				{
					Name:       "Delete",
					ZoneId:     "001",
					RecordId:   generateDNSRecordID("A", "delete.bar.com", "127.0.0.1"),
					RecordData: cloudflare.DNSRecord{},
				},
				{
					Name:   "DeleteDataLocalizationRegionalHostname",
					ZoneId: "001",
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname: "delete.bar.com",
					},
				},
			},
		},
		{
			name: "No change",
			records: map[string]cloudflare.DNSRecord{
				"nochange.bar.com": {
					ID:      generateDNSRecordID("A", "nochange.bar.com", "127.0.0.1"),
					Type:    "A",
					Name:    "nochange.bar.com",
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: proxyDisabled,
				},
			},
			regionalHostnames: []cloudflare.RegionalHostname{
				{
					Hostname:  "nochange.bar.com",
					RegionKey: "eu",
				},
			},
			endpoints: []*endpoint.Endpoint{
				{
					RecordType: "A",
					DNSName:    "nochange.bar.com",
					Targets:    endpoint.Targets{"127.0.0.1"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := &CloudFlareProvider{
				RegionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"},
				Client: &mockCloudFlareClient{
					Zones: map[string]string{
						"001": "bar.com",
					},
					Records: map[string]map[string]cloudflare.DNSRecord{
						"001": tt.records,
					},
					regionalHostnames: map[string][]cloudflare.RegionalHostname{
						"001": tt.regionalHostnames,
					},
				},
			}

			AssertActions(t, provider, tt.endpoints, tt.want, []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME})
		})
	}
}

func TestCloudflareRegionalHostnameDefaults(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1", "127.0.0.2"},
		},
	}

	AssertActions(t, &CloudFlareProvider{RegionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"}}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.2"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.2"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.2",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:   "CreateDataLocalizationRegionalHostname",
			ZoneId: "001",
			RegionalHostname: cloudflare.RegionalHostname{
				Hostname:  "bar.com",
				RegionKey: "us",
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func Test_regionalHostname(t *testing.T) {
	type args struct {
		endpoint *endpoint.Endpoint
		config   RegionalServicesConfig
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
				config: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: "",
				},
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "",
			},
		},
		{
			name: "default region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
				},
				config: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: "us",
				},
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
				config: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: "us",
				},
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
				config: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: "us",
				},
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
				config: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: "us",
				},
			},
			want: cloudflare.RegionalHostname{},
		},
		{
			name: "disabled",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "us",
						},
					},
				},
				config: RegionalServicesConfig{
					Enabled: false,
				},
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "",
				RegionKey: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CloudFlareProvider{RegionalServicesConfig: tt.args.config}
			got := p.regionalHostname(tt.args.endpoint)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_desiredDataLocalizationRegionalHostnames(t *testing.T) {
	tests := []struct {
		name    string
		changes []*cloudFlareChange
		want    []cloudflare.RegionalHostname
		wantErr bool
	}{
		{
			name:    "empty input",
			changes: []*cloudFlareChange{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "change without regional hostname config",
			changes: []*cloudFlareChange{{
				Action: cloudFlareCreate,
			}},
			want:    nil,
			wantErr: false,
		},
		{
			name: "changes with same hostname and region key",
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
			},
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example.com",
					RegionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "changes with same hostname but different region keys",
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
						RegionKey: "us", // Different region key
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "changes with different hostnames",
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
						RegionKey: "us",
					},
				},
			},
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example1.com",
					RegionKey: "eu",
				},
				{
					Hostname:  "example2.com",
					RegionKey: "us",
				},
				{
					Hostname:  "example3.com",
					RegionKey: "",
				},
			},
			wantErr: false,
		},
		{
			name: "change with empty region key",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example.com",
						RegionKey: "", // Empty region key
					},
				},
			},
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example.com",
					RegionKey: "",
				},
			},
			wantErr: false,
		},
		{
			name: "empty region key followed by region key",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example.com",
						RegionKey: "", // Empty region key
					},
				},
				{
					Action: cloudFlareUpdate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "example.com",
						RegionKey: "eu",
					},
				},
			},
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example.com",
					RegionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "region key followed by empty region key",
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
						RegionKey: "eu", // Empty region key
					},
				},
			},
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example.com",
					RegionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "delete followed by create for the same hostname",
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
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example.com",
					RegionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "create followed by delete for the same hostname",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
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
			want: []cloudflare.RegionalHostname{
				{
					Hostname:  "example.com",
					RegionKey: "eu",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := desiredRegionalHostnames(tt.changes)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i].Hostname < got[j].Hostname
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].Hostname < tt.want[j].Hostname
			})
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_dataLocalizationRegionalHostnamesChanges(t *testing.T) {
	tests := []struct {
		name              string
		desired           []cloudflare.RegionalHostname
		regionalHostnames RegionalHostnamesMap
		want              []regionalHostnameChange
	}{
		{
			name:              "empty desired and current lists",
			desired:           []cloudflare.RegionalHostname{},
			regionalHostnames: RegionalHostnamesMap{},
			want:              []regionalHostnameChange{},
		},
		{
			name: "multiple changes",
			desired: []cloudflare.RegionalHostname{
				{
					Hostname:  "create.example.com",
					RegionKey: "eu",
				},
				{
					Hostname:  "update.example.com",
					RegionKey: "eu",
				},
				{
					Hostname:  "delete.example.com",
					RegionKey: "",
				},
				{
					Hostname:  "nochange.example.com",
					RegionKey: "us",
				},
				{
					Hostname:  "absent.example.com",
					RegionKey: "",
				},
			},
			regionalHostnames: RegionalHostnamesMap{
				"update.example.com": cloudflare.RegionalHostname{
					Hostname:  "update.example.com",
					RegionKey: "us",
				},
				"delete.example.com": cloudflare.RegionalHostname{
					Hostname:  "delete.example.com",
					RegionKey: "ap",
				},
				"nochange.example.com": cloudflare.RegionalHostname{
					Hostname:  "nochange.example.com",
					RegionKey: "us",
				},
			},
			want: []regionalHostnameChange{
				{
					action: cloudFlareCreate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "create.example.com",
						RegionKey: "eu",
					},
				},
				{
					action: cloudFlareUpdate,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "update.example.com",
						RegionKey: "eu",
					},
				},
				{
					action: cloudFlareDelete,
					RegionalHostname: cloudflare.RegionalHostname{
						Hostname:  "delete.example.com",
						RegionKey: "",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := regionalHostnamesChanges(tt.desired, tt.regionalHostnames)
			assert.Equal(t, tt.want, got)
		})
	}
}
