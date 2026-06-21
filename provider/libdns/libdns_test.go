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

package libdns

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/libdns/libdns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// fakeClient is an in-memory libdns module for tests. It implements the optional
// RecordSetter and ZoneLister interfaces too.
type fakeClient struct {
	records []libdns.Record
	zones   []string
	setN    int
	delN    int
	appendN int
}

func (f *fakeClient) GetRecords(_ context.Context, _ string) ([]libdns.Record, error) {
	return f.records, nil
}

func (f *fakeClient) AppendRecords(_ context.Context, _ string, recs []libdns.Record) ([]libdns.Record, error) {
	f.appendN++
	f.records = append(f.records, recs...)
	return recs, nil
}

func (f *fakeClient) DeleteRecords(_ context.Context, _ string, recs []libdns.Record) ([]libdns.Record, error) {
	f.delN++
	return recs, nil
}

func (f *fakeClient) SetRecords(_ context.Context, _ string, recs []libdns.Record) ([]libdns.Record, error) {
	f.setN++
	return recs, nil
}

func (f *fakeClient) ListZones(_ context.Context) ([]libdns.Zone, error) {
	zones := make([]libdns.Zone, 0, len(f.zones))
	for _, z := range f.zones {
		zones = append(zones, libdns.Zone{Name: z})
	}
	return zones, nil
}

func TestRecordsGroupsRRsets(t *testing.T) {
	client := &fakeClient{records: []libdns.Record{
		libdns.RR{Name: "www", Type: "A", TTL: 300 * time.Second, Data: "1.2.3.4"},
		libdns.RR{Name: "www", Type: "A", TTL: 300 * time.Second, Data: "5.6.7.8"},
		libdns.RR{Name: "@", Type: "TXT", TTL: 60 * time.Second, Data: "owner=me"},
	}}
	p := &Provider{client: client, zones: []string{"example.com"}}

	endpoints, err := p.Records(t.Context())
	require.NoError(t, err)
	require.Len(t, endpoints, 2)

	byName := map[string]*endpoint.Endpoint{}
	for _, ep := range endpoints {
		byName[ep.DNSName] = ep
	}

	a := byName["www.example.com"]
	require.NotNil(t, a)
	assert.Equal(t, endpoint.RecordTypeA, a.RecordType)
	assert.Equal(t, endpoint.TTL(300), a.RecordTTL)
	assert.ElementsMatch(t, []string{"1.2.3.4", "5.6.7.8"}, []string(a.Targets))

	txt := byName["example.com"]
	require.NotNil(t, txt)
	assert.Equal(t, endpoint.RecordTypeTXT, txt.RecordType)
}

func TestApplyChangesUsesSetAndDelete(t *testing.T) {
	client := &fakeClient{}
	p := &Provider{client: client, zones: []string{"example.com"}}

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("new.example.com", endpoint.RecordTypeA, 300, "1.1.1.1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, 300, "2.2.2.2"),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, "3.3.3.3"),
		},
	}

	require.NoError(t, p.ApplyChanges(t.Context(), changes))
	assert.Equal(t, 1, client.setN, "Create and UpdateNew batched into one SetRecords per zone")
	assert.Equal(t, 1, client.delN, "Delete should be deleted")
}

func TestApplyChangesDryRun(t *testing.T) {
	client := &fakeClient{}
	p := &Provider{client: client, zones: []string{"example.com"}, dryRun: true}

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{endpoint.NewEndpoint("new.example.com", endpoint.RecordTypeA, "1.1.1.1")},
		Delete: []*endpoint.Endpoint{endpoint.NewEndpoint("old.example.com", endpoint.RecordTypeA, "3.3.3.3")},
	}
	require.NoError(t, p.ApplyChanges(t.Context(), changes))
	assert.Zero(t, client.setN)
	assert.Zero(t, client.delN)
}

func TestZoneForLongestSuffix(t *testing.T) {
	p := &Provider{zones: []string{"example.com", "sub.example.com"}}

	zone, ok := p.zoneFor("a.sub.example.com")
	require.True(t, ok)
	assert.Equal(t, "sub.example.com", zone)

	zone, ok = p.zoneFor("a.example.com")
	require.True(t, ok)
	assert.Equal(t, "example.com", zone)

	_, ok = p.zoneFor("a.other.org")
	assert.False(t, ok)
}

func TestToRecordsRelativeNameAndTTL(t *testing.T) {
	ep := endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, 120, "1.2.3.4", "5.6.7.8")
	recs := toRecords("example.com", ep)
	require.Len(t, recs, 2)

	rr := recs[0].RR()
	assert.Equal(t, "www", rr.Name)
	assert.Equal(t, "A", rr.Type)
	assert.Equal(t, 120*time.Second, rr.TTL)

	// Unconfigured TTL maps to zero duration (provider default).
	noTTL := toRecords("example.com", endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "1.2.3.4"))
	assert.Equal(t, time.Duration(0), noTTL[0].RR().TTL)
}

func TestAdjustEndpointsStripsSetIdentifier(t *testing.T) {
	ep := endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "1.2.3.4")
	ep.SetIdentifier = "blue"
	p := &Provider{}

	adjusted, err := p.AdjustEndpoints([]*endpoint.Endpoint{ep})
	require.NoError(t, err)
	assert.Empty(t, adjusted[0].SetIdentifier)
}

func TestResolveZones(t *testing.T) {
	// From --domain-filter.
	zones, err := resolveZones(t.Context(), &fakeClient{}, []string{"example.com.", " other.org "})
	require.NoError(t, err)
	sort.Strings(zones)
	assert.Equal(t, []string{"example.com", "other.org"}, zones)

	// Auto-discovered via ZoneLister when no domain filter is set.
	zones, err = resolveZones(t.Context(), &fakeClient{zones: []string{"discovered.com."}}, nil)
	require.NoError(t, err)
	assert.Equal(t, []string{"discovered.com"}, zones)
}
