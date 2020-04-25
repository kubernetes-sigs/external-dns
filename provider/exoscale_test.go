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

package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/exoscale/egoscale"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type createRecordExoscale struct {
	name string
	rec  egoscale.DNSRecord
}

type deleteRecordExoscale struct {
	name     string
	recordID int64
}

type updateRecordExoscale struct {
	name            string
	updateDNSRecord egoscale.UpdateDNSRecord
}

var createExoscale []createRecordExoscale
var deleteExoscale []deleteRecordExoscale
var updateExoscale []updateRecordExoscale

type ExoscaleClientStub struct {
}

func NewExoscaleClientStub() EgoscaleClientI {
	ep := &ExoscaleClientStub{}
	return ep
}

func (ep *ExoscaleClientStub) DeleteRecord(ctx context.Context, name string, recordID int64) error {
	deleteExoscale = append(deleteExoscale, deleteRecordExoscale{name: name, recordID: recordID})
	return nil
}
func (ep *ExoscaleClientStub) GetRecords(ctx context.Context, name string) ([]egoscale.DNSRecord, error) {
	init := []egoscale.DNSRecord{
		{ID: 0, Name: "v4.barfoo.com", RecordType: "ALIAS"},
		{ID: 1, Name: "v1.foo.com", RecordType: "TXT"},
		{ID: 2, Name: "v2.bar.com", RecordType: "A"},
		{ID: 3, Name: "v3.bar.com", RecordType: "ALIAS"},
		{ID: 4, Name: "v2.foo.com", RecordType: "CNAME"},
		{ID: 5, Name: "v1.foobar.com", RecordType: "TXT"},
	}

	rec := make([]egoscale.DNSRecord, 0)
	for _, r := range init {
		if strings.HasSuffix(r.Name, "."+name) {
			r.Name = strings.TrimSuffix(r.Name, "."+name)
			rec = append(rec, r)
		}
	}

	return rec, nil
}
func (ep *ExoscaleClientStub) UpdateRecord(ctx context.Context, name string, rec egoscale.UpdateDNSRecord) (*egoscale.DNSRecord, error) {
	updateExoscale = append(updateExoscale, updateRecordExoscale{name: name, updateDNSRecord: rec})
	return nil, nil
}
func (ep *ExoscaleClientStub) CreateRecord(ctx context.Context, name string, rec egoscale.DNSRecord) (*egoscale.DNSRecord, error) {
	createExoscale = append(createExoscale, createRecordExoscale{name: name, rec: rec})
	return nil, nil
}
func (ep *ExoscaleClientStub) GetDomains(ctx context.Context) ([]egoscale.DNSDomain, error) {
	dom := []egoscale.DNSDomain{
		{ID: 1, Name: "foo.com"},
		{ID: 2, Name: "bar.com"},
	}
	return dom, nil
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
	provider := NewExoscaleProviderWithClient("", "", "", NewExoscaleClientStub(), false)

	if recs, err := provider.Records(context.Background()); err == nil {
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
	provider := NewExoscaleProviderWithClient("", "", "", NewExoscaleClientStub(), false)

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
	assert.Equal(t, "foo.com", createExoscale[0].name)
	assert.Equal(t, "v1", createExoscale[0].rec.Name)

	assert.Equal(t, 1, len(deleteExoscale))
	assert.Equal(t, "foo.com", deleteExoscale[0].name)
	assert.Equal(t, int64(1), deleteExoscale[0].recordID)

	assert.Equal(t, 1, len(updateExoscale))
	assert.Equal(t, "foo.com", updateExoscale[0].name)
	assert.Equal(t, int64(1), updateExoscale[0].updateDNSRecord.ID)
}
