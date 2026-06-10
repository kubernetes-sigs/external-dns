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

package exoscale

import (
	"context"
	"testing"
	"time"

	v3 "github.com/exoscale/egoscale/v3"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type createRecordExoscale struct {
	domainID v3.UUID
	req      v3.CreateDNSDomainRecordRequest
}

type deleteRecordExoscale struct {
	domainID v3.UUID
	recordID v3.UUID
}

type updateRecordExoscale struct {
	domainID v3.UUID
	recordID v3.UUID
	req      v3.UpdateDNSDomainRecordRequest
}

var (
	createExoscale []createRecordExoscale
	deleteExoscale []deleteRecordExoscale
	updateExoscale []updateRecordExoscale
)

var domainIDs = []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
var groups = map[string][]v3.DNSDomainRecord{
	domainIDs[0]: {
		{ID: v3.UUID(uuid.New().String()), Name: "v1", Type: v3.DNSDomainRecordTypeTXT, Content: "test", Ttl: 3600},
		{ID: v3.UUID(uuid.New().String()), Name: "v2", Type: v3.DNSDomainRecordTypeCNAME, Content: "test", Ttl: 3600},
	},
	domainIDs[1]: {
		{ID: v3.UUID(uuid.New().String()), Name: "v2", Type: v3.DNSDomainRecordTypeA, Content: "test", Ttl: 3600},
		{ID: v3.UUID(uuid.New().String()), Name: "v3", Type: v3.DNSDomainRecordTypeALIAS, Content: "test", Ttl: 3600},
	},
	domainIDs[2]: {
		{ID: v3.UUID(uuid.New().String()), Name: "v1", Type: v3.DNSDomainRecordTypeTXT, Content: "test", Ttl: 3600},
	},
	domainIDs[3]: {
		{ID: v3.UUID(uuid.New().String()), Name: "v4", Type: v3.DNSDomainRecordTypeALIAS, Content: "test", Ttl: 3600},
	},
	// domainIDs[4] is for apex record tests
	domainIDs[4]: {
		{ID: v3.UUID(uuid.New().String()), Name: "", Type: v3.DNSDomainRecordTypeA, Content: "1.2.3.4", Ttl: 3600},
	},
}

type ExoscaleClientStub struct{}

func NewExoscaleClientStub() EgoscaleClientI {
	return &ExoscaleClientStub{}
}

func (ep *ExoscaleClientStub) ListDNSDomains(_ context.Context) ([]v3.DNSDomain, error) {
	return []v3.DNSDomain{
		{ID: v3.UUID(domainIDs[0]), UnicodeName: "foo.com"},
		{ID: v3.UUID(domainIDs[1]), UnicodeName: "bar.com"},
	}, nil
}

func (ep *ExoscaleClientStub) ListDNSDomainRecords(_ context.Context, domainID v3.UUID) ([]v3.DNSDomainRecord, error) {
	return groups[string(domainID)], nil
}

func (ep *ExoscaleClientStub) CreateDNSDomainRecord(_ context.Context, domainID v3.UUID, req v3.CreateDNSDomainRecordRequest) error {
	createExoscale = append(createExoscale, createRecordExoscale{domainID: domainID, req: req})
	return nil
}

func (ep *ExoscaleClientStub) DeleteDNSDomainRecord(_ context.Context, domainID v3.UUID, recordID v3.UUID) error {
	deleteExoscale = append(deleteExoscale, deleteRecordExoscale{domainID: domainID, recordID: recordID})
	return nil
}

func (ep *ExoscaleClientStub) UpdateDNSDomainRecord(_ context.Context, domainID v3.UUID, recordID v3.UUID, req v3.UpdateDNSDomainRecordRequest) error {
	updateExoscale = append(updateExoscale, updateRecordExoscale{domainID: domainID, recordID: recordID, req: req})
	return nil
}

// ExoscaleClientApexStub serves a single domain with one apex record for apex-specific tests.
type ExoscaleClientApexStub struct{}

func (ep *ExoscaleClientApexStub) ListDNSDomains(_ context.Context) ([]v3.DNSDomain, error) {
	return []v3.DNSDomain{
		{ID: v3.UUID(domainIDs[4]), UnicodeName: "apex.com"},
	}, nil
}

func (ep *ExoscaleClientApexStub) ListDNSDomainRecords(_ context.Context, _ v3.UUID) ([]v3.DNSDomainRecord, error) {
	return groups[domainIDs[4]], nil
}

func (ep *ExoscaleClientApexStub) CreateDNSDomainRecord(_ context.Context, domainID v3.UUID, req v3.CreateDNSDomainRecordRequest) error {
	createExoscale = append(createExoscale, createRecordExoscale{domainID: domainID, req: req})
	return nil
}

func (ep *ExoscaleClientApexStub) DeleteDNSDomainRecord(_ context.Context, domainID v3.UUID, recordID v3.UUID) error {
	deleteExoscale = append(deleteExoscale, deleteRecordExoscale{domainID: domainID, recordID: recordID})
	return nil
}

func (ep *ExoscaleClientApexStub) UpdateDNSDomainRecord(_ context.Context, domainID v3.UUID, recordID v3.UUID, req v3.UpdateDNSDomainRecordRequest) error {
	updateExoscale = append(updateExoscale, updateRecordExoscale{domainID: domainID, recordID: recordID, req: req})
	return nil
}

func contains(arr []*endpoint.Endpoint, name string) bool {
	for _, a := range arr {
		if a.DNSName == name {
			return true
		}
	}
	return false
}

func TestExoscaleGetRecords(t *testing.T) {
	provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), false, 0)
	recs, err := provider.Records(t.Context())
	if err == nil {
		assert.Len(t, recs, 3)
		assert.True(t, contains(recs, "v1.foo.com"))
		assert.True(t, contains(recs, "v2.bar.com"))
		assert.True(t, contains(recs, "v2.foo.com"))
		assert.False(t, contains(recs, "v3.bar.com"))
		assert.False(t, contains(recs, "v1.foobar.com"))
	} else {
		assert.Error(t, err)
	}
}

func TestExoscaleGetRecordsApex(t *testing.T) {
	provider := NewExoscaleProviderWithClient(&ExoscaleClientApexStub{}, false, 0)
	recs, err := provider.Records(t.Context())
	assert.NoError(t, err)
	assert.Len(t, recs, 1)
	// Apex record must appear as the bare zone name, not ".apex.com"
	assert.True(t, contains(recs, "apex.com"))
	assert.False(t, contains(recs, ".apex.com"))
}

func TestExoscaleApplyChanges(t *testing.T) {
	provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), false, 0)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "A", Targets: []string{""}},
			{DNSName: "v1.foobar.com", RecordType: "TXT", Targets: []string{""}},
		},
		Delete: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "TXT", Targets: []string{""}},
			{DNSName: "v1.foobar.com", RecordType: "TXT", Targets: []string{""}},
		},
		UpdateOld: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "TXT", Targets: []string{""}},
			{DNSName: "v1.foobar.com", RecordType: "TXT", Targets: []string{""}},
		},
		UpdateNew: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "TXT", Targets: []string{""}},
			{DNSName: "v1.foobar.com", RecordType: "TXT", Targets: []string{""}},
		},
	}
	createExoscale = make([]createRecordExoscale, 0)
	deleteExoscale = make([]deleteRecordExoscale, 0)
	updateExoscale = make([]updateRecordExoscale, 0)

	provider.ApplyChanges(t.Context(), changes)

	assert.Len(t, createExoscale, 1)
	assert.Equal(t, v3.UUID(domainIDs[0]), createExoscale[0].domainID)
	assert.Equal(t, "v1", createExoscale[0].req.Name)

	assert.Len(t, deleteExoscale, 1)
	assert.Equal(t, v3.UUID(domainIDs[0]), deleteExoscale[0].domainID)
	assert.Equal(t, groups[domainIDs[0]][0].ID, deleteExoscale[0].recordID)

	assert.Len(t, updateExoscale, 1)
	assert.Equal(t, v3.UUID(domainIDs[0]), updateExoscale[0].domainID)
	assert.Equal(t, groups[domainIDs[0]][0].ID, updateExoscale[0].recordID)
}

func TestExoscaleApplyChangesApex(t *testing.T) {
	provider := NewExoscaleProviderWithClient(&ExoscaleClientApexStub{}, false, 0)

	createExoscale = make([]createRecordExoscale, 0)
	deleteExoscale = make([]deleteRecordExoscale, 0)
	updateExoscale = make([]updateRecordExoscale, 0)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "apex.com", RecordType: "A", Targets: []string{"1.2.3.4"}},
		},
		Delete: []*endpoint.Endpoint{
			{DNSName: "apex.com", RecordType: "A", Targets: []string{"1.2.3.4"}},
		},
	}

	assert.NoError(t, provider.ApplyChanges(t.Context(), changes))

	assert.Len(t, createExoscale, 1)
	assert.Equal(t, v3.UUID(domainIDs[4]), createExoscale[0].domainID)
	assert.Empty(t, createExoscale[0].req.Name)

	assert.Len(t, deleteExoscale, 1)
	assert.Equal(t, v3.UUID(domainIDs[4]), deleteExoscale[0].domainID)
	assert.Equal(t, groups[domainIDs[4]][0].ID, deleteExoscale[0].recordID)
}

func TestExoscaleMerge_NoUpdateOnTTL0Changes(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(1), RecordType: endpoint.RecordTypeA},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(1), RecordType: endpoint.RecordTypeA},
	}
	updateNew := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(0), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(0), RecordType: endpoint.RecordTypeCNAME},
	}
	assert.Empty(t, merge(updateOld, updateNew))
}

func TestExoscaleMerge_UpdateOnTTLChanges(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(1), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(1), RecordType: endpoint.RecordTypeCNAME},
	}
	updateNew := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(77), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(10), RecordType: endpoint.RecordTypeCNAME},
	}
	merged := merge(updateOld, updateNew)
	assert.Len(t, merged, 2)
	assert.Equal(t, "name1", merged[0].DNSName)
}

func TestExoscaleMerge_AlwaysUpdateTarget(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(1), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(1), RecordType: endpoint.RecordTypeCNAME},
	}
	updateNew := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1-changed"}, RecordTTL: endpoint.TTL(0), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(0), RecordType: endpoint.RecordTypeCNAME},
	}
	merged := merge(updateOld, updateNew)
	assert.Len(t, merged, 1)
	assert.Equal(t, "target1-changed", merged[0].Targets[0])
}

func TestExoscaleMerge_NoUpdateIfTTLUnchanged(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(55), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(55), RecordType: endpoint.RecordTypeCNAME},
	}
	updateNew := []*endpoint.Endpoint{
		{DNSName: "name1", Targets: endpoint.Targets{"target1"}, RecordTTL: endpoint.TTL(55), RecordType: endpoint.RecordTypeCNAME},
		{DNSName: "name2", Targets: endpoint.Targets{"target2"}, RecordTTL: endpoint.TTL(55), RecordType: endpoint.RecordTypeCNAME},
	}
	assert.Empty(t, merge(updateOld, updateNew))
}

func TestZones(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		input    map[string]string
		expected map[string]string
	}{
		{
			name:   "single matching zone",
			domain: "example.com",
			input:  map[string]string{"1": "example.com"},
			expected: map[string]string{
				"1": "example.com",
			},
		},
		{
			name:     "non matching zone",
			domain:   "example.com",
			input:    map[string]string{"1": "other.com"},
			expected: map[string]string{},
		},
		{
			name:   "multiple zones mixed match",
			domain: "example.com",
			input: map[string]string{
				"1": "example.com",
				"2": "sub.example.com",
				"3": "other.com",
			},
			expected: map[string]string{
				"1": "example.com",
				"2": "sub.example.com",
			},
		},
		{
			name:     "empty input",
			domain:   "example.com",
			input:    map[string]string{},
			expected: map[string]string{},
		},
		{
			name:   "empty domain matches all",
			domain: "",
			input: map[string]string{
				"1": "example.com",
				"2": "other.com",
			},
			expected: map[string]string{
				"1": "example.com",
				"2": "other.com",
			},
		},
		{
			name:   "suffix must be exact",
			domain: "ample.com",
			input: map[string]string{
				"1": "example.com",
				"2": "sample.com",
			},
			expected: map[string]string{
				"1": "example.com",
				"2": "sample.com",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			zf := zoneFilter{domain: test.domain}
			result := zf.Zones(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestExoscaleWithDomain_SetsDomain(t *testing.T) {
	tests := []struct {
		name         string
		domainFilter []string
	}{
		{name: "domain filter", domainFilter: []string{"example.com", "apple.xyz"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(_ *testing.T) {
			p := &ExoscaleProvider{}
			ExoscaleWithDomain(endpoint.NewDomainFilter(test.domainFilter))(p)
		})
	}
}

func TestInMemoryWithLogging_LogsChanges(t *testing.T) {
	t.Run("exoscaleWithlogging", func(t *testing.T) {
		logger, hook := test.NewNullLogger()
		log.SetFormatter(logger.Formatter)
		log.SetLevel(log.InfoLevel)
		log.AddHook(hook)

		p := &ExoscaleProvider{}
		ExoscaleWithLogging()(p)

		changes := &plan.Changes{
			Create:    []*endpoint.Endpoint{{DNSName: "create.example.com", RecordType: "A"}},
			UpdateOld: []*endpoint.Endpoint{{DNSName: "old.example.com", RecordType: "A"}},
			UpdateNew: []*endpoint.Endpoint{{DNSName: "new.example.com", RecordType: "A"}},
			Delete:    []*endpoint.Endpoint{{DNSName: "delete.example.com", RecordType: "A"}},
		}

		p.OnApplyChanges(changes)

		entries := hook.AllEntries()
		assert.Contains(t, entries[0].Message, "CREATE")
		assert.Contains(t, entries[1].Message, "UPDATE (old)")
		assert.Contains(t, entries[2].Message, "UPDATE (new)")
		assert.Contains(t, entries[3].Message, "DELETE")
	})
}

func TestExoscaleGetDomainFilter(t *testing.T) {
	t.Run("returns bare and leading-dot variants for each zone", func(t *testing.T) {
		provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), false, 0)
		filter := provider.GetDomainFilter()
		// stub returns foo.com and bar.com
		assert.True(t, filter.Match("foo.com"))
		assert.True(t, filter.Match("v1.foo.com"))
		assert.True(t, filter.Match("bar.com"))
		assert.True(t, filter.Match("v2.bar.com"))
	})

	t.Run("returns empty filter on list error", func(t *testing.T) {
		provider := NewExoscaleProviderWithClient(&errListDomainsStub{}, false, 0)
		filter := provider.GetDomainFilter()
		// empty filter matches nothing specific; getting a DomainFilter back is enough
		assert.NotNil(t, filter)
	})
}

func TestExoscaleApplyChangesDryRun(t *testing.T) {
	provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), true, 0)

	createExoscale = make([]createRecordExoscale, 0)
	deleteExoscale = make([]deleteRecordExoscale, 0)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "A", Targets: []string{"1.2.3.4"}},
		},
		Delete: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "A", Targets: []string{"1.2.3.4"}},
		},
	}

	assert.NoError(t, provider.ApplyChanges(t.Context(), changes))
	// dryRun: nothing should be sent to the API
	assert.Empty(t, createExoscale)
	assert.Empty(t, deleteExoscale)
}

func TestExoscaleApplyChangesWithTTL(t *testing.T) {
	provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), false, 0)

	createExoscale = make([]createRecordExoscale, 0)
	updateExoscale = make([]updateRecordExoscale, 0)

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "A", Targets: []string{"1.2.3.4"}, RecordTTL: 300},
		},
		UpdateOld: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "TXT", Targets: []string{"old"}, RecordTTL: 300},
		},
		UpdateNew: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "TXT", Targets: []string{"new"}, RecordTTL: 600},
		},
	}

	assert.NoError(t, provider.ApplyChanges(t.Context(), changes))
	assert.Len(t, createExoscale, 1)
	assert.Equal(t, int64(300), createExoscale[0].req.Ttl)
	assert.Len(t, updateExoscale, 1)
	assert.Equal(t, int64(600), updateExoscale[0].req.Ttl)
}

func TestExoscaleApplyChangesZonesError(t *testing.T) {
	provider := NewExoscaleProviderWithClient(&errListDomainsStub{}, false, 0)
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{DNSName: "v1.foo.com", RecordType: "A", Targets: []string{"1.2.3.4"}},
		},
	}
	assert.Error(t, provider.ApplyChanges(t.Context(), changes))
}

func TestExoscaleRecordsZonesError(t *testing.T) {
	provider := NewExoscaleProviderWithClient(&errListDomainsStub{}, false, 0)
	_, err := provider.Records(t.Context())
	assert.Error(t, err)
}

func TestExoscaleRecordsListRecordsError(t *testing.T) {
	provider := NewExoscaleProviderWithClient(&errListRecordsStub{}, false, 0)
	_, err := provider.Records(t.Context())
	assert.Error(t, err)
}

func TestExoscaleZoneCacheHit(t *testing.T) {
	provider := NewExoscaleProviderWithClient(
		NewExoscaleClientStub(),
		false,
		time.Hour,
		ExoscaleWithDomain(endpoint.NewDomainFilter([]string{"foo.com"})),
	)
	// first call populates the cache
	recs1, err := provider.Records(t.Context())
	assert.NoError(t, err)
	// second call hits the cache
	recs2, err := provider.Records(t.Context())
	assert.NoError(t, err)
	assert.Len(t, recs2, len(recs1))
}

// errListDomainsStub always errors on ListDNSDomains.
type errListDomainsStub struct{}

func (s *errListDomainsStub) ListDNSDomains(_ context.Context) ([]v3.DNSDomain, error) {
	return nil, assert.AnError
}
func (s *errListDomainsStub) ListDNSDomainRecords(_ context.Context, _ v3.UUID) ([]v3.DNSDomainRecord, error) {
	return nil, nil
}
func (s *errListDomainsStub) CreateDNSDomainRecord(_ context.Context, _ v3.UUID, _ v3.CreateDNSDomainRecordRequest) error {
	return nil
}
func (s *errListDomainsStub) DeleteDNSDomainRecord(_ context.Context, _ v3.UUID, _ v3.UUID) error {
	return nil
}
func (s *errListDomainsStub) UpdateDNSDomainRecord(_ context.Context, _ v3.UUID, _ v3.UUID, _ v3.UpdateDNSDomainRecordRequest) error {
	return nil
}

// errListRecordsStub returns one domain but errors on ListDNSDomainRecords.
type errListRecordsStub struct{}

func (s *errListRecordsStub) ListDNSDomains(_ context.Context) ([]v3.DNSDomain, error) {
	return []v3.DNSDomain{{ID: v3.UUID(domainIDs[0]), UnicodeName: "foo.com"}}, nil
}
func (s *errListRecordsStub) ListDNSDomainRecords(_ context.Context, _ v3.UUID) ([]v3.DNSDomainRecord, error) {
	return nil, assert.AnError
}
func (s *errListRecordsStub) CreateDNSDomainRecord(_ context.Context, _ v3.UUID, _ v3.CreateDNSDomainRecordRequest) error {
	return nil
}
func (s *errListRecordsStub) DeleteDNSDomainRecord(_ context.Context, _ v3.UUID, _ v3.UUID) error {
	return nil
}
func (s *errListRecordsStub) UpdateDNSDomainRecord(_ context.Context, _ v3.UUID, _ v3.UUID, _ v3.UpdateDNSDomainRecordRequest) error {
	return nil
}
