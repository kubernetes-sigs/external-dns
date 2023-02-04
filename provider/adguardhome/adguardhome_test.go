/*
Copyright 2020 The Kubernetes Authors.

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
package adguardhome

import (
	"context"
	"os"
	"reflect"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type mockAdguardClient struct {
	rules []string
}

func (m *mockAdguardClient) GetFilteringRules(_ context.Context) ([]string, error) {
	return m.rules, nil
}

func (m *mockAdguardClient) SaveFilteringRules(_ context.Context, rules []string) error {
	m.rules = rules
	return nil
}

func newMockClient() *mockAdguardClient {
	return &mockAdguardClient{
		rules: []string{
			"# I am not for external-dns",
			"1.1.1.1 example.com #$managed by external-dns",
			"# myresponse notexample.com $managed by external-dns",
		},
	}
}

func TestAdguardHomeProvider_ApplyChanges(t *testing.T) {
	p := &AdguardHomeProvider{
		client: newMockClient(),
	}

	create := []*endpoint.Endpoint{
		{
			DNSName:    "example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"2.2.2.2"},
		},
	}

	updateNew := []*endpoint.Endpoint{
		{
			DNSName:    "example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"3.3.3.3"},
		},
	}

	changes := &plan.Changes{
		Create:    create,
		UpdateNew: updateNew,
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "notexample.com",
				RecordType: endpoint.RecordTypeTXT,
				Targets:    endpoint.Targets{"myresponse"},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: endpoint.RecordTypeA,
				Targets:    endpoint.Targets{"1.1.1.1"},
			},
		},
	}

	err := p.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("failed to apply changes: %v", err)
	}

	r, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("failed to fetch records: %v", err)
	}

	expected := []*endpoint.Endpoint{
		{
			DNSName:    "example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"2.2.2.2", "3.3.3.3"},
		},
	}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("records do not match: got: %v, expected: %v", r, expected)
	}

	// Ensure original rules were kept in place
	expectedRules := []string{
		"# I am not for external-dns",
		"2.2.2.2 example.com #$managed by external-dns",
		"3.3.3.3 example.com #$managed by external-dns",
	}

	if !reflect.DeepEqual(p.client.(*mockAdguardClient).rules, expectedRules) {
		t.Errorf("rules do not match: got: %v, expected: %v", p.client.(*mockAdguardClient).rules, expectedRules)
	}
}

func TestAdguardHomeProvider_Records(t *testing.T) {
	type fields struct {
		BaseProvider provider.BaseProvider
		client       Client
		domainFilter endpoint.DomainFilter
		DryRun       bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*endpoint.Endpoint
		wantErr bool
	}{
		{
			name: "test fetches records",
			fields: fields{
				client: newMockClient(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					RecordType: "A",
					Targets:    endpoint.Targets{"1.1.1.1"},
				},
				{
					DNSName:    "notexample.com",
					RecordType: "TXT",
					Targets:    endpoint.Targets{"myresponse"},
				},
			},
		},
		{
			name: "applies domain filter",
			fields: fields{
				client: newMockClient(),
				domainFilter: endpoint.NewDomainFilter(
					[]string{"example.com"},
				),
			},
			args: args{
				ctx: context.Background(),
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					RecordType: "A",
					Targets:    endpoint.Targets{"1.1.1.1"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AdguardHomeProvider{
				BaseProvider: tt.fields.BaseProvider,
				client:       tt.fields.client,
				domainFilter: tt.fields.domainFilter,
			}
			got, err := p.Records(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdguardHomeProvider.Records() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdguardHomeProvider.Records() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAdguardHomeProvider(t *testing.T) {
	got, err := NewAdguardHomeProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("NewAdguardHomeProvider() error = %v", err)
		return
	}

	if got != nil {
		t.Errorf("NewAdguardHomeProvider() = %v, want %v", got, "not nil")
	}

	_ = os.Setenv(envURL, "http://localhost:3000")
	_ = os.Setenv(envUser, "pw")
	_ = os.Setenv(envPassword, "user")

	got, err = NewAdguardHomeProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("NewAdguardHomeProvider() error = %v", err)
	}
}

func TestAdguardHomeProvider_MergeRecords(t *testing.T) {
	c := newMockClient()
	p := &AdguardHomeProvider{
		client: c,
	}

	c.rules = []string{
		"1.1.1.1 example.com #$managed by external-dns",
		"1.2.1.1 example.com #$managed by external-dns",
		"1.3.1.1 example.com #$managed by external-dns",
		"1.4.1.1 example.com #$managed by external-dns",
	}

	got, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("failed to fetch records: %v", err)
	}

	expected := []*endpoint.Endpoint{
		{
			DNSName:    "example.com",
			RecordType: "A",
			Targets: endpoint.Targets{
				"1.1.1.1",
				"1.2.1.1",
				"1.3.1.1",
				"1.4.1.1",
			},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("records do not match: got: %v, expected: %v", got, expected)
	}
}
