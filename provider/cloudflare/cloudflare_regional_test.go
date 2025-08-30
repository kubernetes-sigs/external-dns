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
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5/addressing"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source/annotations"
)

func (m *mockCloudFlareClient) ListDataLocalizationRegionalHostnames(ctx context.Context, params addressing.RegionalHostnameListParams) autoPager[addressing.RegionalHostnameListResponse] {
	zoneID := params.ZoneID.Value
	if strings.Contains(zoneID, "rherror") {
		return &mockAutoPager[addressing.RegionalHostnameListResponse]{err: fmt.Errorf("failed to list regional hostnames")}
	}
	results := make([]addressing.RegionalHostnameListResponse, 0, len(m.regionalHostnames[zoneID]))
	for _, rh := range m.regionalHostnames[zoneID] {
		results = append(results, addressing.RegionalHostnameListResponse{
			Hostname:  rh.hostname,
			RegionKey: rh.regionKey,
		})
	}
	return &mockAutoPager[addressing.RegionalHostnameListResponse]{
		items: results,
	}
}

func (m *mockCloudFlareClient) CreateDataLocalizationRegionalHostname(ctx context.Context, params addressing.RegionalHostnameNewParams) error {
	if strings.Contains(params.Hostname.Value, "rherror") {
		return fmt.Errorf("failed to create regional hostname")
	}

	m.Actions = append(m.Actions, MockAction{
		Name:     "CreateDataLocalizationRegionalHostname",
		ZoneId:   params.ZoneID.Value,
		RecordId: "",
		RegionalHostname: regionalHostname{
			hostname:  params.Hostname.Value,
			regionKey: params.RegionKey.Value,
		},
	})
	return nil
}

func (m *mockCloudFlareClient) UpdateDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameEditParams) error {
	if strings.Contains(hostname, "rherror") {
		return fmt.Errorf("failed to update regional hostname")
	}

	m.Actions = append(m.Actions, MockAction{
		Name:     "UpdateDataLocalizationRegionalHostname",
		ZoneId:   params.ZoneID.Value,
		RecordId: "",
		RegionalHostname: regionalHostname{
			hostname:  hostname,
			regionKey: params.RegionKey.Value,
		},
	})
	return nil
}

func (m *mockCloudFlareClient) DeleteDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameDeleteParams) error {
	if strings.Contains(hostname, "rherror") {
		return fmt.Errorf("failed to delete regional hostname")
	}
	m.Actions = append(m.Actions, MockAction{
		Name:     "DeleteDataLocalizationRegionalHostname",
		ZoneId:   params.ZoneID.Value,
		RecordId: "",
		RegionalHostname: regionalHostname{
			hostname: hostname,
		},
	})
	return nil
}

func TestCloudflareRegionalHostnameActions(t *testing.T) {
	tests := []struct {
		name              string
		records           map[string]dns.RecordResponse
		regionalHostnames []regionalHostname
		endpoints         []*endpoint.Endpoint
		want              []MockAction
	}{
		{
			name:              "create",
			records:           map[string]dns.RecordResponse{},
			regionalHostnames: []regionalHostname{},
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
					RecordData: dns.RecordResponse{
						ID:      generateDNSRecordID("A", "create.bar.com", "127.0.0.1"),
						Type:    "A",
						Name:    "create.bar.com",
						Content: "127.0.0.1",
						TTL:     1,
						Proxied: false,
					},
				},
				{
					Name:   "CreateDataLocalizationRegionalHostname",
					ZoneId: "001",
					RegionalHostname: regionalHostname{
						hostname:  "create.bar.com",
						regionKey: "eu",
					},
				},
			},
		},
		{
			name: "Update",
			records: map[string]dns.RecordResponse{
				"update.bar.com": {
					ID:      generateDNSRecordID("A", "update.bar.com", "127.0.0.1"),
					Type:    "A",
					Name:    "update.bar.com",
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: false,
				},
			},
			regionalHostnames: []regionalHostname{
				{
					hostname:  "update.bar.com",
					regionKey: "us",
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
					RecordData: dns.RecordResponse{
						ID:      generateDNSRecordID("A", "update.bar.com", "127.0.0.1"),
						Type:    "A",
						Name:    "update.bar.com",
						Content: "127.0.0.1",
						TTL:     1,
						Proxied: false,
					},
				},
				{
					Name:   "UpdateDataLocalizationRegionalHostname",
					ZoneId: "001",
					RegionalHostname: regionalHostname{
						hostname:  "update.bar.com",
						regionKey: "eu",
					},
				},
			},
		},
		{
			name: "Delete",
			records: map[string]dns.RecordResponse{
				"update.bar.com": {
					ID:      generateDNSRecordID("A", "delete.bar.com", "127.0.0.1"),
					Type:    "A",
					Name:    "delete.bar.com",
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: false,
				},
			},
			regionalHostnames: []regionalHostname{
				{
					hostname:  "delete.bar.com",
					regionKey: "us",
				},
			},
			endpoints: []*endpoint.Endpoint{},
			want: []MockAction{
				{
					Name:       "Delete",
					ZoneId:     "001",
					RecordId:   generateDNSRecordID("A", "delete.bar.com", "127.0.0.1"),
					RecordData: dns.RecordResponse{},
				},
				{
					Name:   "DeleteDataLocalizationRegionalHostname",
					ZoneId: "001",
					RegionalHostname: regionalHostname{
						hostname: "delete.bar.com",
					},
				},
			},
		},
		{
			name: "No change",
			records: map[string]dns.RecordResponse{
				"nochange.bar.com": {
					ID:      generateDNSRecordID("A", "nochange.bar.com", "127.0.0.1"),
					Type:    "A",
					Name:    "nochange.bar.com",
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: false,
				},
			},
			regionalHostnames: []regionalHostname{
				{
					hostname:  "nochange.bar.com",
					regionKey: "eu",
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
					Records: map[string]map[string]dns.RecordResponse{
						"001": tt.records,
					},
					regionalHostnames: map[string][]regionalHostname{
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: false,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.2"),
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.2"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.2",
				TTL:     1,
				Proxied: false,
			},
		},
		{
			Name:   "CreateDataLocalizationRegionalHostname",
			ZoneId: "001",
			RegionalHostname: regionalHostname{
				hostname:  "bar.com",
				regionKey: "us",
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
		want regionalHostname
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
			want: regionalHostname{
				hostname:  "example.com",
				regionKey: "",
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
			want: regionalHostname{
				hostname:  "example.com",
				regionKey: "us",
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
			want: regionalHostname{
				hostname:  "example.com",
				regionKey: "eu",
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
			want: regionalHostname{
				hostname:  "example.com",
				regionKey: "",
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
			want: regionalHostname{},
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
			want: regionalHostname{
				hostname:  "",
				regionKey: "",
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
		want    []regionalHostname
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
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
				{
					Action: cloudFlareUpdate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example.com",
					regionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "changes with same hostname but different region keys",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
				{
					Action: cloudFlareUpdate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "us", // Different region key
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
					RegionalHostname: regionalHostname{
						hostname:  "example1.com",
						regionKey: "eu",
					},
				},
				{
					Action: cloudFlareUpdate,
					RegionalHostname: regionalHostname{
						hostname:  "example2.com",
						regionKey: "us",
					},
				},
				{
					Action: cloudFlareDelete,
					RegionalHostname: regionalHostname{
						hostname:  "example3.com",
						regionKey: "us",
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example1.com",
					regionKey: "eu",
				},
				{
					hostname:  "example2.com",
					regionKey: "us",
				},
				{
					hostname:  "example3.com",
					regionKey: "",
				},
			},
			wantErr: false,
		},
		{
			name: "change with empty region key",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "", // Empty region key
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example.com",
					regionKey: "",
				},
			},
			wantErr: false,
		},
		{
			name: "empty region key followed by region key",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "", // Empty region key
					},
				},
				{
					Action: cloudFlareUpdate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example.com",
					regionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "region key followed by empty region key",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
				{
					Action: cloudFlareUpdate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu", // Empty region key
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example.com",
					regionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "delete followed by create for the same hostname",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareDelete,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
				{
					Action: cloudFlareCreate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example.com",
					regionKey: "eu",
				},
			},
			wantErr: false,
		},
		{
			name: "create followed by delete for the same hostname",
			changes: []*cloudFlareChange{
				{
					Action: cloudFlareCreate,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
				{
					Action: cloudFlareDelete,
					RegionalHostname: regionalHostname{
						hostname:  "example.com",
						regionKey: "eu",
					},
				},
			},
			want: []regionalHostname{
				{
					hostname:  "example.com",
					regionKey: "eu",
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
				return got[i].hostname < got[j].hostname
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].hostname < tt.want[j].hostname
			})
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_dataLocalizationRegionalHostnamesChanges(t *testing.T) {
	tests := []struct {
		name              string
		desired           []regionalHostname
		regionalHostnames regionalHostnamesMap
		want              []regionalHostnameChange
	}{
		{
			name:              "empty desired and current lists",
			desired:           []regionalHostname{},
			regionalHostnames: regionalHostnamesMap{},
			want:              []regionalHostnameChange{},
		},
		{
			name: "multiple changes",
			desired: []regionalHostname{
				{
					hostname:  "create.example.com",
					regionKey: "eu",
				},
				{
					hostname:  "update.example.com",
					regionKey: "eu",
				},
				{
					hostname:  "delete.example.com",
					regionKey: "",
				},
				{
					hostname:  "nochange.example.com",
					regionKey: "us",
				},
				{
					hostname:  "absent.example.com",
					regionKey: "",
				},
			},
			regionalHostnames: regionalHostnamesMap{
				"update.example.com": regionalHostname{
					hostname:  "update.example.com",
					regionKey: "us",
				},
				"delete.example.com": regionalHostname{
					hostname:  "delete.example.com",
					regionKey: "ap",
				},
				"nochange.example.com": regionalHostname{
					hostname:  "nochange.example.com",
					regionKey: "us",
				},
			},
			want: []regionalHostnameChange{
				{
					action: cloudFlareCreate,
					regionalHostname: regionalHostname{
						hostname:  "create.example.com",
						regionKey: "eu",
					},
				},
				{
					action: cloudFlareUpdate,
					regionalHostname: regionalHostname{
						hostname:  "update.example.com",
						regionKey: "eu",
					},
				},
				{
					action: cloudFlareDelete,
					regionalHostname: regionalHostname{
						hostname:  "delete.example.com",
						regionKey: "",
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

func TestRecordsWithListRegionalHostnameFaillure(t *testing.T) {
	client := &mockCloudFlareClient{
		Zones: map[string]string{
			"rherror": "error.com",
		},
		Records: map[string]map[string]dns.RecordResponse{
			"rherror": {"foo.error.com": {Type: "A"}},
		},
	}
	failingProvider := &CloudFlareProvider{
		Client:                 client,
		RegionalServicesConfig: RegionalServicesConfig{Enabled: true},
	}
	_, err := failingProvider.Records(t.Context())
	assert.Error(t, err, "listing regional hostnames should fail")
}

func TestApplyChangesWithRegionalHostnamesFaillures(t *testing.T) {
	t.Parallel()
	type fields struct {
		Records           map[string]dns.RecordResponse
		RegionalHostnames []regionalHostname
		RegionKey         string
	}
	type args struct {
		changes *plan.Changes
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		errMsg      string
		expectDebug string
	}{
		{
			name: "list zone fails",
			args: args{
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "foo.error.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
				},
			},
			errMsg: "failed to list regional hostnames",
		},
		{
			name: "create fails",
			fields: fields{
				Records:           map[string]dns.RecordResponse{},
				RegionalHostnames: []regionalHostname{},
				RegionKey:         "us",
			},
			args: args{
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "rherror.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
				},
			},
			expectDebug: "failed to create regional hostname",
		},
		{
			name: "update fails",
			fields: fields{
				Records: map[string]dns.RecordResponse{
					"rherror.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "rherror.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []regionalHostname{
					{hostname: "rherror.bar.com", regionKey: "us"},
				},
				RegionKey: "us",
			},
			args: args{
				changes: &plan.Changes{
					UpdateOld: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "rherror.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
					UpdateNew: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "rherror.bar.com",
							Targets:    endpoint.Targets{"127.0.0.2"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
				},
			},
			expectDebug: "failed to update regional hostname",
		},
		{
			name: "delete fails",
			fields: fields{
				Records: map[string]dns.RecordResponse{
					"rherror.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "newerror.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []regionalHostname{
					{hostname: "rherror.bar.com", regionKey: "us"},
				},
				RegionKey: "us",
			},
			args: args{
				changes: &plan.Changes{
					Delete: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "rherror.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
						},
					},
				},
			},
			expectDebug: "failed to delete regional hostname",
		},
		{
			// This should not happen in practice, but we test it to ensure we return an error.
			name: "conflicting regional keys",
			fields: fields{
				Records:           map[string]dns.RecordResponse{},
				RegionalHostnames: []regionalHostname{},
				RegionKey:         "us",
			},
			args: args{
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "foo.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
						{
							RecordType: "A",
							DNSName:    "foo.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "us"},
							},
						},
					},
				},
			},
			errMsg: "conflicting region keys",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			records := tt.fields.Records
			if records == nil {
				records = map[string]dns.RecordResponse{}
			}
			p := &CloudFlareProvider{
				Client: &mockCloudFlareClient{
					Zones: map[string]string{
						"001":     "bar.com",
						"rherror": "error.com",
					},
					Records: map[string]map[string]dns.RecordResponse{
						"001": records,
					},
					regionalHostnames: map[string][]regionalHostname{
						"001": tt.fields.RegionalHostnames,
					},
				},
				RegionalServicesConfig: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: tt.fields.RegionKey,
				},
			}
			hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)
			err := p.ApplyChanges(t.Context(), tt.args.changes)
			assert.Error(t, err, "ApplyChanges should return an error")
			if tt.errMsg != "" && err != nil {
				assert.Contains(t, err.Error(), tt.errMsg, "Expected error message to contain: %s", tt.errMsg)
			}
			if tt.expectDebug != "" {
				testutils.TestHelperLogContains(tt.expectDebug, hook, t)
			}
		})
	}
}

func TestApplyChangesWithRegionalHostnamesDryRun(t *testing.T) {
	t.Parallel()
	type fields struct {
		Records           map[string]dns.RecordResponse
		RegionalHostnames []regionalHostname
		RegionKey         string
	}
	type args struct {
		changes *plan.Changes
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectDebug string
	}{
		{
			name: "create dry run",
			fields: fields{
				Records:           map[string]dns.RecordResponse{},
				RegionalHostnames: []regionalHostname{},
				RegionKey:         "us",
			},
			args: args{
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "foo.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
				},
			},
			expectDebug: "Dry run: skipping regional hostname change",
		},
		{
			name: "update fails",
			fields: fields{
				Records: map[string]dns.RecordResponse{
					"foo.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "foo.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []regionalHostname{
					{hostname: "foo.bar.com", regionKey: "us"},
				},
				RegionKey: "us",
			},
			args: args{
				changes: &plan.Changes{
					UpdateOld: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "foo.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
					UpdateNew: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "foo.bar.com",
							Targets:    endpoint.Targets{"127.0.0.2"},
							ProviderSpecific: endpoint.ProviderSpecific{
								{Name: "external-dns.alpha.kubernetes.io/cloudflare-region-key", Value: "eu"},
							},
						},
					},
				},
			},
			expectDebug: "Dry run: skipping regional hostname change",
		},
		{
			name: "delete fails",
			fields: fields{
				Records: map[string]dns.RecordResponse{
					"foo.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "foo.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []regionalHostname{
					{hostname: "foo.bar.com", regionKey: "us"},
				},
				RegionKey: "us",
			},
			args: args{
				changes: &plan.Changes{
					Delete: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "foo.bar.com",
							Targets:    endpoint.Targets{"127.0.0.1"},
						},
					},
				},
			},
			expectDebug: "Dry run: skipping regional hostname change",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			records := tt.fields.Records
			if records == nil {
				records = map[string]dns.RecordResponse{}
			}
			p := &CloudFlareProvider{
				DryRun: true,
				Client: &mockCloudFlareClient{
					Zones: map[string]string{
						"001": "bar.com",
					},
					Records: map[string]map[string]dns.RecordResponse{
						"001": records,
					},
					regionalHostnames: map[string][]regionalHostname{
						"001": tt.fields.RegionalHostnames,
					},
				},
				RegionalServicesConfig: RegionalServicesConfig{
					Enabled:   true,
					RegionKey: tt.fields.RegionKey,
				},
			}
			hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)
			err := p.ApplyChanges(t.Context(), tt.args.changes)
			assert.NoError(t, err, "ApplyChanges should not fail")
			if tt.expectDebug != "" {
				testutils.TestHelperLogContains(tt.expectDebug, hook, t)
			}
		})
	}
}

func TestCloudflareAdjustEndpointsRegionalServices(t *testing.T) {
	testCases := []struct {
		name                   string
		recordType             string
		regionalServicesConfig RegionalServicesConfig
		initialRegionKey       string  // existing region key on endpoint
		expectedRegionKey      *string // expected region key after AdjustEndpoints (nil = should not be present)
	}{
		// Supported types should get region key when enabled
		{
			name:                   "A record with regional services enabled",
			recordType:             "A",
			regionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"},
			initialRegionKey:       "",
			expectedRegionKey:      testutils.ToPtr("us"),
		},
		{
			name:                   "AAAA record with regional services enabled",
			recordType:             "AAAA",
			regionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"},
			initialRegionKey:       "",
			expectedRegionKey:      testutils.ToPtr("us"),
		},
		{
			name:                   "CNAME record with regional services enabled",
			recordType:             "CNAME",
			regionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"},
			initialRegionKey:       "",
			expectedRegionKey:      testutils.ToPtr("us"),
		},

		// Unsupported types should NOT get region key even when enabled
		{
			name:                   "TXT record with regional services enabled",
			recordType:             "TXT",
			regionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"},
			initialRegionKey:       "",
			expectedRegionKey:      nil,
		},

		// Disabled regional services should remove region key for all types
		{
			name:                   "A record with regional services disabled",
			recordType:             "A",
			regionalServicesConfig: RegionalServicesConfig{Enabled: false},
			initialRegionKey:       "existing-region",
			expectedRegionKey:      nil,
		},
		{
			name:                   "TXT record with regional services disabled",
			recordType:             "TXT",
			regionalServicesConfig: RegionalServicesConfig{Enabled: false},
			initialRegionKey:       "existing-region",
			expectedRegionKey:      nil,
		},

		// Existing region key should be preserved when already set
		{
			name:                   "A record with existing custom region key",
			recordType:             "A",
			regionalServicesConfig: RegionalServicesConfig{Enabled: true, RegionKey: "us"},
			initialRegionKey:       "eu",
			expectedRegionKey:      testutils.ToPtr("eu"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create endpoint with initial region key if specified
			testEndpoint := &endpoint.Endpoint{
				RecordType: tc.recordType,
				DNSName:    "test.bar.com",
				Targets:    endpoint.Targets{"127.0.0.1"},
			}

			if tc.initialRegionKey != "" {
				testEndpoint.ProviderSpecific = endpoint.ProviderSpecific{
					endpoint.ProviderSpecificProperty{
						Name:  annotations.CloudflareRegionKey,
						Value: tc.initialRegionKey,
					},
				}
			}

			provider := &CloudFlareProvider{
				RegionalServicesConfig: tc.regionalServicesConfig,
			}

			adjustedEndpoints, err := provider.AdjustEndpoints([]*endpoint.Endpoint{testEndpoint})
			assert.NoError(t, err)
			assert.Len(t, adjustedEndpoints, 1)

			regionKey, exists := adjustedEndpoints[0].GetProviderSpecificProperty(annotations.CloudflareRegionKey)

			if tc.expectedRegionKey != nil {
				// Region key should be present with expected value
				assert.True(t, exists, "Region key should be present")
				assert.Equal(t, *tc.expectedRegionKey, regionKey, "Region key value should match expected")
			} else {
				// Region key should not be present
				assert.False(t, exists, "Region key should not be present")
			}
		})
	}
}
