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
	"errors"
	"testing"

	egoscale "github.com/exoscale/egoscale/v2"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"

	"github.com/google/uuid"
)

type createRecordExoscale struct {
	domainID string
	record   *egoscale.DNSDomainRecord
}

type deleteRecordExoscale struct {
	domainID string
	recordID string
}

type updateRecordExoscale struct {
	domainID string
	record   *egoscale.DNSDomainRecord
}

var (
	createExoscale []createRecordExoscale
	deleteExoscale []deleteRecordExoscale
	updateExoscale []updateRecordExoscale
)

var defaultTTL int64 = 3600
var domainIDs = []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
var groups = map[string][]egoscale.DNSDomainRecord{
	domainIDs[0]: {
		{ID: strPtr(uuid.New().String()), Name: strPtr("v1"), Type: strPtr("TXT"), Content: strPtr("test"), TTL: &defaultTTL},
		{ID: strPtr(uuid.New().String()), Name: strPtr("v2"), Type: strPtr("CNAME"), Content: strPtr("test"), TTL: &defaultTTL},
	},
	domainIDs[1]: {
		{ID: strPtr(uuid.New().String()), Name: strPtr("v2"), Type: strPtr("A"), Content: strPtr("test"), TTL: &defaultTTL},
		{ID: strPtr(uuid.New().String()), Name: strPtr("v3"), Type: strPtr("ALIAS"), Content: strPtr("test"), TTL: &defaultTTL},
	},
	domainIDs[2]: {
		{ID: strPtr(uuid.New().String()), Name: strPtr("v1"), Type: strPtr("TXT"), Content: strPtr("test"), TTL: &defaultTTL},
	},
	domainIDs[3]: {
		{ID: strPtr(uuid.New().String()), Name: strPtr("v4"), Type: strPtr("ALIAS"), Content: strPtr("test"), TTL: &defaultTTL},
	},
}

func strPtr(s string) *string {
	return &s
}

type ExoscaleClientStub struct{}

func NewExoscaleClientStub() EgoscaleClientI {
	ep := &ExoscaleClientStub{}
	return ep
}

func (ep *ExoscaleClientStub) ListDNSDomains(ctx context.Context, _ string) ([]egoscale.DNSDomain, error) {
	domains := []egoscale.DNSDomain{
		{ID: &domainIDs[0], UnicodeName: strPtr("foo.com")},
		{ID: &domainIDs[1], UnicodeName: strPtr("bar.com")},
	}
	return domains, nil
}

func (ep *ExoscaleClientStub) ListDNSDomainRecords(ctx context.Context, _, domainID string) ([]egoscale.DNSDomainRecord, error) {
	return groups[domainID], nil
}

func (ep *ExoscaleClientStub) GetDNSDomainRecord(ctx context.Context, _, domainID, recordID string) (*egoscale.DNSDomainRecord, error) {
	group := groups[domainID]
	for _, record := range group {
		if *record.ID == recordID {
			return &record, nil
		}
	}

	return nil, errors.New("not found")
}

func (ep *ExoscaleClientStub) CreateDNSDomainRecord(ctx context.Context, _, domainID string, record *egoscale.DNSDomainRecord) (*egoscale.DNSDomainRecord, error) {
	createExoscale = append(createExoscale, createRecordExoscale{domainID: domainID, record: record})
	return record, nil
}

func (ep *ExoscaleClientStub) DeleteDNSDomainRecord(ctx context.Context, _, domainID string, record *egoscale.DNSDomainRecord) error {
	deleteExoscale = append(deleteExoscale, deleteRecordExoscale{domainID: domainID, recordID: *record.ID})
	return nil
}

func (ep *ExoscaleClientStub) UpdateDNSDomainRecord(ctx context.Context, _, domainID string, record *egoscale.DNSDomainRecord) error {
	updateExoscale = append(updateExoscale, updateRecordExoscale{domainID: domainID, record: record})
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
	provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), "", "", false)

	recs, err := provider.Records(context.Background())
	if err == nil {
		assert.Equal(t, 3, len(recs))
		assert.True(t, contains(recs, "v1.foo.com"))
		assert.True(t, contains(recs, "v2.bar.com"))
		assert.True(t, contains(recs, "v2.foo.com"))
		assert.False(t, contains(recs, "v3.bar.com"))
		assert.False(t, contains(recs, "v1.foobar.com"))
	} else {
		assert.Error(t, err)
	}
}

func TestExoscaleApplyChanges(t *testing.T) {
	provider := NewExoscaleProviderWithClient(NewExoscaleClientStub(), "", "", false)

	plan := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
	}
	createExoscale = make([]createRecordExoscale, 0)
	deleteExoscale = make([]deleteRecordExoscale, 0)

	provider.ApplyChanges(context.Background(), plan)

	assert.Equal(t, 1, len(createExoscale))
	assert.Equal(t, domainIDs[0], createExoscale[0].domainID)
	assert.Equal(t, "v1", *createExoscale[0].record.Name)

	assert.Equal(t, 1, len(deleteExoscale))
	assert.Equal(t, domainIDs[0], deleteExoscale[0].domainID)
	assert.Equal(t, *groups[domainIDs[0]][0].ID, deleteExoscale[0].recordID)

	assert.Equal(t, 1, len(updateExoscale))
	assert.Equal(t, domainIDs[0], updateExoscale[0].domainID)
	assert.Equal(t, *groups[domainIDs[0]][0].ID, *updateExoscale[0].record.ID)
}

func TestExoscaleMerge_NoUpdateOnTTL0Changes(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeA,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeA,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	assert.Equal(t, 0, len(merge(updateOld, updateNew)))
}

func TestExoscaleMerge_UpdateOnTTLChanges(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(77),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(10),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	merged := merge(updateOld, updateNew)
	assert.Equal(t, 2, len(merged))
	assert.Equal(t, "name1", merged[0].DNSName)
}

func TestExoscaleMerge_AlwaysUpdateTarget(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(1),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1-changed"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(0),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	merged := merge(updateOld, updateNew)
	assert.Equal(t, 1, len(merged))
	assert.Equal(t, "target1-changed", merged[0].Targets[0])
}

func TestExoscaleMerge_NoUpdateIfTTLUnchanged(t *testing.T) {
	updateOld := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "name1",
			Targets:    endpoint.Targets{"target1"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
		{
			DNSName:    "name2",
			Targets:    endpoint.Targets{"target2"},
			RecordTTL:  endpoint.TTL(55),
			RecordType: endpoint.RecordTypeCNAME,
		},
	}

	merged := merge(updateOld, updateNew)
	assert.Equal(t, 0, len(merged))
}
