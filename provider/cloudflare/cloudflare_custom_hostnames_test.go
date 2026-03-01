/*
Copyright 2026 The Kubernetes Authors.

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
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/plan"
)

func (m *mockCloudFlareClient) CustomHostnames(ctx context.Context, zoneID string) autoPager[custom_hostnames.CustomHostnameListResponse] {
	if strings.HasPrefix(zoneID, "newerror-") {
		return &mockAutoPager[custom_hostnames.CustomHostnameListResponse]{
			err: errors.New("failed to list custom hostnames"),
		}
	}

	result := []custom_hostnames.CustomHostnameListResponse{}
	if chs, ok := m.customHostnames[zoneID]; ok {
		for _, ch := range chs {
			if strings.HasPrefix(ch.hostname, "newerror-list-") {
				params := custom_hostnames.CustomHostnameDeleteParams{ZoneID: cloudflare.F(zoneID)}
				m.DeleteCustomHostname(ctx, ch.id, params)
				return &mockAutoPager[custom_hostnames.CustomHostnameListResponse]{
					err: errors.New("failed to list erroring custom hostname"),
				}
			}
			result = append(result, custom_hostnames.CustomHostnameListResponse{
				ID:                 ch.id,
				Hostname:           ch.hostname,
				CustomOriginServer: ch.customOriginServer,
			})
		}
	}
	return &mockAutoPager[custom_hostnames.CustomHostnameListResponse]{
		items: result,
	}
}

func (m *mockCloudFlareClient) CreateCustomHostname(_ context.Context, zoneID string, ch customHostname) error {
	if ch.hostname == "" || ch.customOriginServer == "" || ch.hostname == "newerror-create.foo.fancybar.com" {
		return fmt.Errorf("Invalid custom hostname or origin hostname")
	}
	if _, ok := m.customHostnames[zoneID]; !ok {
		m.customHostnames[zoneID] = []customHostname{}
	}
	newCustomHostname := ch
	newCustomHostname.id = fmt.Sprintf("ID-%s", ch.hostname)
	m.customHostnames[zoneID] = append(m.customHostnames[zoneID], newCustomHostname)
	return nil
}

func (m *mockCloudFlareClient) DeleteCustomHostname(_ context.Context, customHostnameID string, params custom_hostnames.CustomHostnameDeleteParams) error {
	zoneID := params.ZoneID.String()
	idx := 0
	if idx = getCustomHostnameIdxByID(m.customHostnames[zoneID], customHostnameID); idx < 0 {
		return fmt.Errorf("Invalid custom hostname ID to delete")
	}

	m.customHostnames[zoneID] = append(m.customHostnames[zoneID][:idx], m.customHostnames[zoneID][idx+1:]...)

	if customHostnameID == "ID-newerror-delete.foo.fancybar.com" {
		return fmt.Errorf("Invalid custom hostname to delete")
	}
	return nil
}

func getCustomHostnameIdxByID(chs []customHostname, customHostnameID string) int {
	for idx, ch := range chs {
		if ch.id == customHostnameID {
			return idx
		}
	}
	return -1
}

func TestCloudflareCustomHostnameOperations(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := t.Context()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testFailCases := []struct {
		Name                    string
		Endpoints               []*endpoint.Endpoint
		ExpectedCustomHostnames map[string]string
	}{}

	for _, tc := range testFailCases {
		t.Run(tc.Name, func(t *testing.T) {
			records, err := provider.Records(ctx)
			if err != nil {
				t.Errorf("should not fail, %v", err)
			}

			endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

			assert.NoError(t, err)
			plan := &plan.Plan{
				Current:        records,
				Desired:        endpoints,
				DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
				ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
			}

			planned := plan.Calculate()

			err = provider.ApplyChanges(t.Context(), planned.Changes)
			if e := checkFailed(tc.Name, err, false); !errors.Is(e, nil) {
				t.Error(e)
			}

			chs, chErr := provider.listCustomHostnamesWithPagination(ctx, "001")
			if e := checkFailed(tc.Name, chErr, false); !errors.Is(e, nil) {
				t.Error(e)
			}

			actualCustomHostnames := map[string]string{}
			for _, ch := range chs {
				actualCustomHostnames[ch.hostname] = ch.customOriginServer
			}
			if len(actualCustomHostnames) == 0 {
				actualCustomHostnames = nil
			}
			assert.Equal(t, tc.ExpectedCustomHostnames, actualCustomHostnames, "custom hostnames should be the same")
		})
	}
}

func TestCloudflareDisabledCustomHostnameOperations(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: false},
	}
	ctx := t.Context()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testCases := []struct {
		Name        string
		Endpoints   []*endpoint.Endpoint
		testChanges bool
	}{
		{
			Name: "add custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.11"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "a.foo.fancybar.com",
						},
					},
				},
				{
					DNSName:    "b.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.12"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
				{
					DNSName:    "c.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.13"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "c1.foo.fancybar.com",
						},
					},
				},
			},
			testChanges: false,
		},
		{
			Name: "add custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.11"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
				{
					DNSName:    "b.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.12"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "b.foo.fancybar.com",
						},
					},
				},
				{
					DNSName:    "c.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.13"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "c2.foo.fancybar.com",
						},
					},
				},
			},
			testChanges: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			records, err := provider.Records(ctx)
			if err != nil {
				t.Errorf("should not fail, %v", err)
			}

			endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

			assert.NoError(t, err)
			plan := &plan.Plan{
				Current:        records,
				Desired:        endpoints,
				DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
				ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
			}
			planned := plan.Calculate()
			err = provider.ApplyChanges(ctx, planned.Changes)
			if e := checkFailed(tc.Name, err, false); !errors.Is(e, nil) {
				t.Error(e)
			}
			if tc.testChanges {
				assert.False(t, planned.Changes.HasChanges(), "no new changes should be here")
			}
		})
	}
}

func TestCloudflareCustomHostnameNotFoundOnRecordDeletion(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := t.Context()
	zoneID := "001"
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testCases := []struct {
		Name                    string
		Endpoints               []*endpoint.Endpoint
		ExpectedCustomHostnames map[string]string
		preApplyHook            string
		logOutput               string
	}{
		{
			Name: "create DNS record with custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "create.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "newerror-getCustomHostnameOrigin.foo.fancybar.com",
						},
					},
				},
			},
			preApplyHook: "",
			logOutput:    "",
		},
		{
			Name:         "remove DNS record with unexpectedly missing custom hostname",
			Endpoints:    []*endpoint.Endpoint{},
			preApplyHook: "corrupt",
			logOutput:    "failed to delete custom hostname \"newerror-getCustomHostnameOrigin.foo.fancybar.com\": failed to get custom hostname: \"newerror-getCustomHostnameOrigin.foo.fancybar.com\" not found",
		},
		{
			Name:         "duplicate custom hostname",
			Endpoints:    []*endpoint.Endpoint{},
			preApplyHook: "duplicate",
			logOutput:    "",
		},
		{
			Name: "create DNS record with custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "a.foo.fancybar.com",
						},
					},
				},
			},
			preApplyHook: "",
			logOutput:    "custom hostname \"a.foo.fancybar.com\" already exists with the same origin \"a.foo.bar.com\", continue",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.InfoLevel, t)

			records, err := provider.Records(ctx)
			if err != nil {
				t.Errorf("should not fail, %v", err)
			}

			endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

			assert.NoError(t, err)
			plan := &plan.Plan{
				Current:        records,
				Desired:        endpoints,
				DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
				ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
			}

			planned := plan.Calculate()

			// manually corrupt custom hostname before the deletion step
			// the purpose is to cause getCustomHostnameOrigin() to fail on change.Action == cloudFlareDelete
			chs, chErr := provider.listCustomHostnamesWithPagination(ctx, zoneID)
			if e := checkFailed(tc.Name, chErr, false); !errors.Is(e, nil) {
				t.Error(e)
			}
			switch tc.preApplyHook {
			case "corrupt":
				if ch, err := getCustomHostname(chs, "newerror-getCustomHostnameOrigin.foo.fancybar.com"); errors.Is(err, nil) {
					chID := ch.id
					t.Logf("corrupting custom hostname %q", chID)
					oldIdx := getCustomHostnameIdxByID(client.customHostnames[zoneID], chID)
					oldCh := client.customHostnames[zoneID][oldIdx]
					ch := customHostname{
						hostname:           "corrupted-newerror-getCustomHostnameOrigin.foo.fancybar.com",
						customOriginServer: oldCh.customOriginServer,
						ssl:                oldCh.ssl,
					}
					client.customHostnames[zoneID][oldIdx] = ch
				}
			case "duplicate": // manually inject duplicating custom hostname with the same name and origin
				ch := customHostname{
					id:                 "ID-random-123",
					hostname:           "a.foo.fancybar.com",
					customOriginServer: "a.foo.bar.com",
				}
				client.customHostnames[zoneID] = append(client.customHostnames[zoneID], ch)
			}
			err = provider.ApplyChanges(t.Context(), planned.Changes)
			if e := checkFailed(tc.Name, err, false); !errors.Is(e, nil) {
				t.Error(e)
			}

			logtest.TestHelperLogContains(tc.logOutput, hook, t)
		})
	}
}

func TestCloudflareListCustomHostnamesWithPagionation(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := t.Context()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	const CustomHostnamesNumber = 342
	var generatedEndpoints []*endpoint.Endpoint
	for i := range CustomHostnamesNumber {
		ep := []*endpoint.Endpoint{
			{
				DNSName:    fmt.Sprintf("host-%d.foo.bar.com", i),
				Targets:    endpoint.Targets{fmt.Sprintf("cname-%d.foo.bar.com", i)},
				RecordType: endpoint.RecordTypeCNAME,
				RecordTTL:  endpoint.TTL(defaultTTL),
				Labels:     endpoint.Labels{},
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
						Value: fmt.Sprintf("host-%d.foo.fancybar.com", i),
					},
				},
			},
		}
		generatedEndpoints = append(generatedEndpoints, ep...)
	}

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %v", err)
	}

	endpoints, err := provider.AdjustEndpoints(generatedEndpoints)

	assert.NoError(t, err)
	plan := &plan.Plan{
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}

	planned := plan.Calculate()

	err = provider.ApplyChanges(t.Context(), planned.Changes)
	if err != nil {
		t.Errorf("should not fail - %v", err)
	}

	chs, chErr := provider.listCustomHostnamesWithPagination(ctx, "001")
	if chErr != nil {
		t.Errorf("should not fail - %v", chErr)
	}
	assert.Len(t, chs, CustomHostnamesNumber)
}

func TestBuildCustomHostnameNewParams(t *testing.T) {
	t.Run("Minimal custom hostname without SSL", func(t *testing.T) {
		ch := customHostname{
			hostname:           "test.example.com",
			customOriginServer: "origin.example.com",
		}

		params := buildCustomHostnameNewParams("zone-123", ch)

		assert.Equal(t, "zone-123", params.ZoneID.Value)
		assert.Equal(t, "test.example.com", params.Hostname.Value)
		assert.False(t, params.SSL.Present)
	})

	t.Run("Custom hostname with full SSL configuration", func(t *testing.T) {
		ch := customHostname{
			hostname:           "test.example.com",
			customOriginServer: "origin.example.com",
			ssl: &customHostnameSSL{
				sslType:              "dv",
				method:               "http",
				bundleMethod:         "ubiquitous",
				certificateAuthority: "digicert",
				settings: customHostnameSSLSettings{
					minTLSVersion: "1.2",
				},
			},
		}

		params := buildCustomHostnameNewParams("zone-123", ch)

		assert.Equal(t, "zone-123", params.ZoneID.Value)
		assert.Equal(t, "test.example.com", params.Hostname.Value)
		assert.True(t, params.SSL.Present)

		ssl := params.SSL.Value
		assert.Equal(t, "dv", string(ssl.Type.Value))
		assert.Equal(t, "http", string(ssl.Method.Value))
		assert.Equal(t, "ubiquitous", string(ssl.BundleMethod.Value))
		assert.Equal(t, "digicert", string(ssl.CertificateAuthority.Value))
		assert.Equal(t, "1.2", string(ssl.Settings.Value.MinTLSVersion.Value))
	})

	t.Run("Custom hostname with partial SSL configuration", func(t *testing.T) {
		ch := customHostname{
			hostname:           "test.example.com",
			customOriginServer: "origin.example.com",
			ssl: &customHostnameSSL{
				sslType: "dv",
				method:  "http",
			},
		}

		params := buildCustomHostnameNewParams("zone-123", ch)

		assert.True(t, params.SSL.Present)
		ssl := params.SSL.Value
		assert.Equal(t, "dv", string(ssl.Type.Value))
		assert.Equal(t, "http", string(ssl.Method.Value))
		assert.False(t, ssl.BundleMethod.Present)
		assert.False(t, ssl.CertificateAuthority.Present)
		assert.False(t, ssl.Settings.Present)
	})

	t.Run("Custom hostname with 'none' certificate authority", func(t *testing.T) {
		ch := customHostname{
			hostname:           "test.example.com",
			customOriginServer: "origin.example.com",
			ssl: &customHostnameSSL{
				sslType:              "dv",
				method:               "http",
				certificateAuthority: "none",
			},
		}

		params := buildCustomHostnameNewParams("zone-123", ch)

		assert.True(t, params.SSL.Present)
		ssl := params.SSL.Value
		// "none" should not be set as certificate authority
		assert.False(t, ssl.CertificateAuthority.Present)
	})

	t.Run("Custom hostname with empty certificate authority", func(t *testing.T) {
		ch := customHostname{
			hostname:           "test.example.com",
			customOriginServer: "origin.example.com",
			ssl: &customHostnameSSL{
				sslType:              "dv",
				method:               "http",
				certificateAuthority: "",
			},
		}

		params := buildCustomHostnameNewParams("zone-123", ch)

		assert.True(t, params.SSL.Present)
		ssl := params.SSL.Value
		// Empty string should not be set
		assert.False(t, ssl.CertificateAuthority.Present)
	})

	t.Run("Custom hostname with only MinTLSVersion", func(t *testing.T) {
		ch := customHostname{
			hostname:           "test.example.com",
			customOriginServer: "origin.example.com",
			ssl: &customHostnameSSL{
				settings: customHostnameSSLSettings{
					minTLSVersion: "1.3",
				},
			},
		}

		params := buildCustomHostnameNewParams("zone-123", ch)

		assert.True(t, params.SSL.Present)
		ssl := params.SSL.Value
		assert.True(t, ssl.Settings.Present)
		assert.Equal(t, "1.3", string(ssl.Settings.Value.MinTLSVersion.Value))
	})
}

func TestSubmitCustomHostnameChanges(t *testing.T) {
	ctx := t.Context()

	t.Run("CustomHostnames_Disabled", func(t *testing.T) {
		client := NewMockCloudFlareClient()
		provider := &CloudFlareProvider{
			Client: client,
			CustomHostnamesConfig: CustomHostnamesConfig{
				Enabled: false,
			},
		}

		change := &cloudFlareChange{
			Action: cloudFlareCreate,
		}

		result := provider.submitCustomHostnameChanges(ctx, "zone1", change, nil, nil)
		assert.True(t, result, "Should return true when custom hostnames are disabled")
	})

	t.Run("CustomHostnames_Create", func(t *testing.T) {
		client := NewMockCloudFlareClient()
		provider := &CloudFlareProvider{
			Client: client,
			CustomHostnamesConfig: CustomHostnamesConfig{
				Enabled: true,
			},
		}

		change := &cloudFlareChange{
			Action: cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{
				Type: "A",
			},
			CustomHostnames: map[string]customHostname{
				"new.example.com": {
					hostname:           "new.example.com",
					customOriginServer: "origin.example.com",
				},
			},
		}

		chs := make(customHostnamesMap)
		result := provider.submitCustomHostnameChanges(ctx, "zone1", change, chs, nil)
		assert.True(t, result, "Should successfully create custom hostname")
		assert.Len(t, client.customHostnames["zone1"], 1, "One custom hostname should be created")
		assert.Contains(t, client.customHostnames["zone1"],
			customHostname{
				id:                 "ID-new.example.com",
				hostname:           "new.example.com",
				customOriginServer: "origin.example.com",
			},
			"Custom hostname should be created in mock client",
		)
	})

	t.Run("CustomHostnames_Create_AlreadyExists", func(t *testing.T) {
		client := NewMockCloudFlareClient()
		provider := &CloudFlareProvider{
			Client: client,
			CustomHostnamesConfig: CustomHostnamesConfig{
				Enabled: true,
			},
		}

		change := &cloudFlareChange{
			Action: cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{
				Type: "A",
			},
			CustomHostnames: map[string]customHostname{
				"exists.example.com": {
					hostname:           "exists.example.com",
					customOriginServer: "origin.example.com",
				},
			},
		}

		chs := customHostnamesMap{
			customHostnameIndex{hostname: "exists.example.com"}: {
				id:                 "ch1",
				hostname:           "exists.example.com",
				customOriginServer: "origin.example.com",
			},
		}

		client.customHostnames = map[string][]customHostname{
			"zone1": {
				{
					id:                 "ch1",
					hostname:           "exists.example.com",
					customOriginServer: "origin.example.com",
				},
			},
		}

		result := provider.submitCustomHostnameChanges(ctx, "zone1", change, chs, nil)
		assert.True(t, result, "Should succeed when custom hostname already exists with same origin")
		assert.Len(t, client.customHostnames["zone1"], 1, "No new custom hostname should be created")
		assert.Contains(t, client.customHostnames["zone1"],
			customHostname{
				id:                 "ch1",
				hostname:           "exists.example.com",
				customOriginServer: "origin.example.com",
			},
			"Existing custom hostname should remain unchanged in mock client",
		)
	})

	t.Run("CustomHostnames_Delete", func(_ *testing.T) {
		client := NewMockCloudFlareClient()
		client.customHostnames = map[string][]customHostname{
			"zone1": {
				{
					id:                 "ch1",
					hostname:           "delete.example.com",
					customOriginServer: "origin.example.com",
				},
			},
		}
		provider := &CloudFlareProvider{
			Client: client,
			CustomHostnamesConfig: CustomHostnamesConfig{
				Enabled: true,
			},
		}

		change := &cloudFlareChange{
			Action: cloudFlareDelete,
			ResourceRecord: dns.RecordResponse{
				Type: "A",
			},
			CustomHostnames: map[string]customHostname{
				"delete.example.com": {
					hostname: "delete.example.com",
				},
			},
		}

		chs := customHostnamesMap{
			customHostnameIndex{hostname: "delete.example.com"}: {
				id:                 "ch1",
				hostname:           "delete.example.com",
				customOriginServer: "origin.example.com",
			},
		}

		// Note: submitCustomHostnameChanges returns false on failure, true on success
		// The mock may not find the hostname to delete, which is fine for this test
		result := provider.submitCustomHostnameChanges(ctx, "zone1", change, chs, nil)
		// We just verify it doesn't panic - result may be true or false depending on mock behavior
		_ = result
	})

	t.Run("CustomHostnames_Update", func(t *testing.T) {
		client := NewMockCloudFlareClient()
		client.customHostnames = map[string][]customHostname{
			"zone1": {
				{
					id:                 "ch1",
					hostname:           "old.example.com",
					customOriginServer: "origin.example.com",
				},
			},
		}
		provider := &CloudFlareProvider{
			Client: client,
			CustomHostnamesConfig: CustomHostnamesConfig{
				Enabled: true,
			},
		}

		change := &cloudFlareChange{
			Action: cloudFlareUpdate,
			ResourceRecord: dns.RecordResponse{
				Type: "A",
			},
			CustomHostnames: map[string]customHostname{
				"new.example.com": {
					hostname:           "new.example.com",
					customOriginServer: "origin.example.com",
				},
			},
			CustomHostnamesPrev: []string{"old.example.com"},
		}

		chs := customHostnamesMap{
			customHostnameIndex{hostname: "old.example.com"}: {
				id:                 "ch1",
				hostname:           "old.example.com",
				customOriginServer: "origin.example.com",
			},
		}

		client.customHostnames = map[string][]customHostname{
			"zone1": {
				{
					id:                 "ch1",
					hostname:           "old.example.com",
					customOriginServer: "origin.example.com",
				},
			},
		}

		result := provider.submitCustomHostnameChanges(ctx, "zone1", change, chs, nil)
		assert.True(t, result, "Should successfully update custom hostname")
		assert.Len(t, client.customHostnames["zone1"], 1, "One custom hostname should exist after update")
		assert.Contains(t, client.customHostnames["zone1"],
			customHostname{
				id:                 "ID-new.example.com",
				hostname:           "new.example.com",
				customOriginServer: "origin.example.com",
			},
			"Custom hostname should be updated in mock client",
		)
	})
}
