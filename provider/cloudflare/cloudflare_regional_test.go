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

	cloudflare "github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
)

func (m *mockCloudFlareClient) ListDataLocalizationRegionalHostnames(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.ListDataLocalizationRegionalHostnamesParams) ([]cloudflare.RegionalHostname, error) {
	if strings.Contains(rc.Identifier, "rherror") {
		return nil, fmt.Errorf("failed to list regional hostnames")
	}
	return m.regionalHostnames[rc.Identifier], nil
}

func (m *mockCloudFlareClient) CreateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDataLocalizationRegionalHostnameParams) error {
	if strings.Contains(rp.Hostname, "rherror") {
		return fmt.Errorf("failed to create regional hostname")
	}

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
	if strings.Contains(rp.Hostname, "rherror") {
		return fmt.Errorf("failed to update regional hostname")
	}

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
	if strings.Contains(hostname, "rherror") {
		return fmt.Errorf("failed to delete regional hostname")
	}
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

func TestCloudflareRegionalHostnameActions(t *testing.T) {
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

func TestRecordsWithListRegionalHostnameFaillure(t *testing.T) {
	client := &mockCloudFlareClient{
		Zones: map[string]string{
			"rherror": "error.com",
		},
		Records: map[string]map[string]cloudflare.DNSRecord{
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
		Records           map[string]cloudflare.DNSRecord
		RegionalHostnames []cloudflare.RegionalHostname
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
				Records:           map[string]cloudflare.DNSRecord{},
				RegionalHostnames: []cloudflare.RegionalHostname{},
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
				Records: map[string]cloudflare.DNSRecord{
					"rherror.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "rherror.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []cloudflare.RegionalHostname{
					{Hostname: "rherror.bar.com", RegionKey: "us"},
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
				Records: map[string]cloudflare.DNSRecord{
					"rherror.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "newerror.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []cloudflare.RegionalHostname{
					{Hostname: "rherror.bar.com", RegionKey: "us"},
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
				Records:           map[string]cloudflare.DNSRecord{},
				RegionalHostnames: []cloudflare.RegionalHostname{},
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
				records = map[string]cloudflare.DNSRecord{}
			}
			p := &CloudFlareProvider{
				Client: &mockCloudFlareClient{
					Zones: map[string]string{
						"001":     "bar.com",
						"rherror": "error.com",
					},
					Records: map[string]map[string]cloudflare.DNSRecord{
						"001": records,
					},
					regionalHostnames: map[string][]cloudflare.RegionalHostname{
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
		Records           map[string]cloudflare.DNSRecord
		RegionalHostnames []cloudflare.RegionalHostname
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
				Records:           map[string]cloudflare.DNSRecord{},
				RegionalHostnames: []cloudflare.RegionalHostname{},
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
				Records: map[string]cloudflare.DNSRecord{
					"foo.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "foo.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []cloudflare.RegionalHostname{
					{Hostname: "foo.bar.com", RegionKey: "us"},
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
				Records: map[string]cloudflare.DNSRecord{
					"foo.bar.com": {
						ID:      "123",
						Type:    "A",
						Name:    "foo.bar.com",
						Content: "127.0.0.1",
					},
				},
				RegionalHostnames: []cloudflare.RegionalHostname{
					{Hostname: "foo.bar.com", RegionKey: "us"},
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
				records = map[string]cloudflare.DNSRecord{}
			}
			p := &CloudFlareProvider{
				DryRun: true,
				Client: &mockCloudFlareClient{
					Zones: map[string]string{
						"001": "bar.com",
					},
					Records: map[string]map[string]cloudflare.DNSRecord{
						"001": records,
					},
					regionalHostnames: map[string][]cloudflare.RegionalHostname{
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
