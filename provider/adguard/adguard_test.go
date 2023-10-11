/*
Copyright 2023 The Kubernetes Authors.

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

package adguard

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// mockClient mocks the http client and response
//
// maintains a map of records based on operations
type mockClient struct {
	resp *http.Response
	err  error

	t *testing.T

	// records are potentially modified when Do is called
	records map[string]RewriteEntry
}

func (mc mockClient) Do(req *http.Request) (*http.Response, error) {
	mc.t.Helper()
	defer func() {
		if req != nil && req.Body != nil {
			req.Body.Close()
		}
	}()

	if mc.resp != nil || mc.err != nil {
		return mc.resp, mc.err
	}

	switch req.URL.Path {
	case ListRewrites:
		entries := make([]RewriteEntry, 0, len(mc.records))
		for _, entry := range mc.records {
			entries = append(entries, entry)
		}
		jsonBytes, err := json.Marshal(entries)
		if err != nil {
			mc.t.Errorf("failed to marshal json: %v", err)
		}
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(jsonBytes)),
		}
		return resp, nil
	case DeleteRewrite:
		body, err := io.ReadAll(req.Body)
		if err != nil {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request")),
			}
			return resp, err
		}
		entry := &RewriteEntry{}
		if err := json.Unmarshal(body, entry); err != nil {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request")),
			}
			return resp, err
		}
		delete(mc.records, entry.Domain)
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("OK"))}, nil
	case CreateRwrite:
		body, err := io.ReadAll(req.Body)
		if err != nil {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request")),
			}
			return resp, err
		}
		entry := &RewriteEntry{}
		if err := json.Unmarshal(body, entry); err != nil {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request")),
			}
			return resp, err
		}
		mc.records[entry.Domain] = *entry
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("OK"))}, nil
	case UpdateRewrite:
		body, err := io.ReadAll(req.Body)
		if err != nil {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request")),
			}
			return resp, err
		}
		entry := &UpdateRewriteEntry{}
		if err := json.Unmarshal(body, entry); err != nil {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request")),
			}
			return resp, err
		}
		existing, ok := mc.records[entry.Target.Domain]
		if !ok || existing.Answer != entry.Target.Answer || entry.Update.Domain != entry.Target.Domain {
			resp := &http.Response{
				StatusCode: 400,
				Body:       io.NopCloser(strings.NewReader("bad request - existing not found")),
			}
			return resp, nil
		}
		mc.records[entry.Update.Domain] = entry.Update
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("OK"))}, nil
	default:
		mc.t.Errorf("Received invalid path: %s", req.URL.Path)
		return nil, nil
	}
}

func TestNewAdguardProvider(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  Config
		want    *Provider
		wantErr error
	}{
		{
			name: "Invalid Username",
			config: Config{
				Password: "superDuperSecret",
				Server:   "https://where.you.run.adguard.fqdn",
			},
			wantErr: ErrInvalidUsername,
		},
		{
			name: "Invalid Password",
			config: Config{
				Username: "greatestUsernameEver",
				Server:   "https://where.you.run.adguard.fqdn",
			},
			wantErr: ErrInvalidPassword,
		},
		{
			name: "Invalid Endpoint",
			config: Config{
				Username: "greatestUsernameEver",
				Password: "superDuperSecret",
			},
			wantErr: ErrInvalidEndpoint,
		},
		{
			name:    "Multiple Issues",
			config:  Config{},
			wantErr: ErrInvalidConfig,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewProvider(tt.config)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewAdguardProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdguardProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_Records(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  Config
		client  httpClient
		want    []*endpoint.Endpoint
		wantErr error
	}{
		{
			name:   "Happy Path - has rewrites",
			config: Config{},
			client: mockClient{
				t: t,
				records: map[string]RewriteEntry{
					"foo.bar.baz":      {Domain: "foo.bar.baz", Answer: "10.11.12.13"},
					"stuff.and.things": {Domain: "stuff.and.things", Answer: "foo.bar.baz"},
				},
			},
			want: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.bar.baz", endpoint.RecordTypeA, "10.11.12.13"),
				endpoint.NewEndpoint("stuff.and.things", endpoint.RecordTypeCNAME, "foo.bar.baz"),
			},
		},
		{
			name:   "Happy Path - no rewrites",
			config: Config{},
			client: mockClient{
				t:       t,
				records: map[string]RewriteEntry{},
			},
			want: []*endpoint.Endpoint{},
		},
		{
			name:   "Throws an error",
			config: Config{},
			client: mockClient{
				t:       t,
				records: map[string]RewriteEntry{},
				err:     assert.AnError,
			},
			want:    []*endpoint.Endpoint{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ap := Provider{
				config: tt.config,
				client: tt.client,
			}

			got, err := ap.Records(ctx)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_recordsToEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		records []RewriteEntry
		want    []*endpoint.Endpoint
	}{
		{
			name: "Mixed CNAME and IP answer",
			records: []RewriteEntry{
				{
					Domain: "best.domain.ever",
					Answer: "10.11.12.13",
				},
				{
					Domain: "worst.domain.ever",
					Answer: "best.domain.ever",
				},
			},
			want: []*endpoint.Endpoint{
				endpoint.NewEndpoint("best.domain.ever", endpoint.RecordTypeA, "10.11.12.13"),
				endpoint.NewEndpoint("worst.domain.ever", endpoint.RecordTypeCNAME, "best.domain.ever"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := recordsToEndpoints(tt.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recordsToEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_ApplyChanges(t *testing.T) {
	t.Parallel()

	type fields struct {
		config Config
		client *mockClient
	}
	tests := []struct {
		name        string
		fields      fields
		changes     *plan.Changes
		wantErr     error
		wantRecords map[string]RewriteEntry
	}{
		{
			name: "Delete happy path",
			fields: fields{
				config: Config{
					Username: "userHere",
					Password: "superSecret",
					Server:   "https://where.adguard.is.hosted",
				},
				client: &mockClient{
					t: t,
					records: map[string]RewriteEntry{
						"some.cname.record": {Domain: "some.cname.record", Answer: "some.other.record"},
						"some.a.record":     {Domain: "some.a.record", Answer: "10.0.0.10"},
						"existing.entry":    {Domain: "existing.entry", Answer: "1.2.3.4"},
					},
				},
			},
			changes: &plan.Changes{
				Delete: []*endpoint.Endpoint{
					endpoint.NewEndpoint("some.cname.record", endpoint.RecordTypeCNAME, "some.other.record"),
					endpoint.NewEndpoint("some.a.record", endpoint.RecordTypeA, "10.0.0.10"),
				},
			},
			wantRecords: map[string]RewriteEntry{
				"existing.entry": {Domain: "existing.entry", Answer: "1.2.3.4"},
			},
		},
		{
			name: "Create happy path",
			fields: fields{
				config: Config{
					Username: "userHere",
					Password: "superSecret",
					Server:   "https://where.adguard.is.hosted",
				},
				client: &mockClient{
					t: t,
					records: map[string]RewriteEntry{
						"existing.entry": {Domain: "existing.entry", Answer: "1.2.3.4"},
					},
				},
			},
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("create.domain.here", endpoint.RecordTypeA, "10.10.10.10"),
				},
			},
			wantRecords: map[string]RewriteEntry{
				"create.domain.here": {Domain: "create.domain.here", Answer: "10.10.10.10"},
				"existing.entry":     {Domain: "existing.entry", Answer: "1.2.3.4"},
			},
		},
		{
			name: "Update happy path",
			fields: fields{
				config: Config{
					Username: "userHere",
					Password: "superSecret",
					Server:   "https://where.adguard.is.hosted",
				},
				client: &mockClient{
					t: t,
					records: map[string]RewriteEntry{
						"update.domain.here": {Domain: "update.domain.here", Answer: "10.0.0.1"},
					},
				},
			},
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{
					endpoint.NewEndpoint("update.domain.here", endpoint.RecordTypeA, "10.10.10.10"),
				},
				UpdateOld: []*endpoint.Endpoint{
					endpoint.NewEndpoint("update.domain.here", endpoint.RecordTypeA, "10.0.0.1"),
				},
			},
			wantRecords: map[string]RewriteEntry{
				"update.domain.here": {Domain: "update.domain.here", Answer: "10.10.10.10"},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			ap := Provider{
				BaseProvider: provider.BaseProvider{},
				config:       tt.fields.config,
				client:       tt.fields.client,
			}
			assert.ErrorIs(t, ap.ApplyChanges(ctx, tt.changes), tt.wantErr)
			assert.Equal(t, tt.wantRecords, tt.fields.client.records)
		})
	}
}
