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

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

var records []*endpoint.Endpoint

type FakeWebhookProvider struct {
	err          error
	domainFilter endpoint.DomainFilter
}

func (p FakeWebhookProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if p.err != nil {
		return nil, p.err
	}
	return records, nil
}

func (p FakeWebhookProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if p.err != nil {
		return p.err
	}
	records = append(records, changes.Create...)
	return nil
}

func (p FakeWebhookProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	// for simplicity, we do not adjust endpoints in this test
	if p.err != nil {
		return nil, p.err
	}
	return endpoints, nil
}

func (p FakeWebhookProvider) GetDomainFilter() endpoint.DomainFilter {
	return p.domainFilter
}

func TestMain(m *testing.M) {
	records = []*endpoint.Endpoint{
		{
			DNSName:    "foo.bar.com",
			RecordType: "A",
		},
	}
	m.Run()
}

func TestRecordsHandlerRecords(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/records", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{
			domainFilter: endpoint.NewDomainFilter([]string{"foo.bar.com"}),
		},
	}
	providerAPIServer.RecordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	// require that the res has the same endpoints as the records slice
	defer res.Body.Close()
	require.NotNil(t, res.Body)
	endpoints := []*endpoint.Endpoint{}
	if err := json.NewDecoder(res.Body).Decode(&endpoints); err != nil {
		t.Errorf("Failed to decode response body: %s", err.Error())
	}
	require.Equal(t, records, endpoints)
}

func TestRecordsHandlerRecordsWithErrors(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/records", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{
			err: fmt.Errorf("error"),
		},
	}
	providerAPIServer.RecordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestRecordsHandlerApplyChangesWithBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/applychanges", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{},
	}
	providerAPIServer.RecordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestRecordsHandlerApplyChangesWithValidRequest(t *testing.T) {
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "foo.bar.com",
				RecordType: "A",
				Targets:    endpoint.Targets{},
			},
		},
	}
	j, err := json.Marshal(changes)
	require.NoError(t, err)

	reader := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodPost, "/applychanges", reader)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{},
	}
	providerAPIServer.RecordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestRecordsHandlerApplyChangesWithErrors(t *testing.T) {
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "foo.bar.com",
				RecordType: "A",
				Targets:    endpoint.Targets{},
			},
		},
	}
	j, err := json.Marshal(changes)
	require.NoError(t, err)

	reader := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodPost, "/applychanges", reader)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{
			err: fmt.Errorf("error"),
		},
	}
	providerAPIServer.RecordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestRecordsHandlerWithWrongHTTPMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/records", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{},
	}
	providerAPIServer.RecordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAdjustEndpointsHandlerWithInvalidRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/adjustendpoints", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{},
	}
	providerAPIServer.AdjustEndpointsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/adjustendpoints", nil)

	providerAPIServer.AdjustEndpointsHandler(w, req)
	res = w.Result()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAdjustEndpointsHandlerWithValidRequest(t *testing.T) {
	pve := []*endpoint.Endpoint{
		{
			DNSName:    "foo.bar.com",
			RecordType: "A",
			Targets:    endpoint.Targets{},
			RecordTTL:  0,
		},
	}

	j, err := json.Marshal(pve)
	require.NoError(t, err)

	reader := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/adjustendpoints", reader)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{},
	}
	providerAPIServer.AdjustEndpointsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.NotNil(t, res.Body)
}

func TestAdjustEndpointsHandlerWithError(t *testing.T) {
	pve := []*endpoint.Endpoint{
		{
			DNSName:    "foo.bar.com",
			RecordType: "A",
			Targets:    endpoint.Targets{},
			RecordTTL:  0,
		},
	}

	j, err := json.Marshal(pve)
	require.NoError(t, err)

	reader := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/adjustendpoints", reader)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		Provider: &FakeWebhookProvider{
			err: fmt.Errorf("error"),
		},
	}
	providerAPIServer.AdjustEndpointsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
	require.NotNil(t, res.Body)
}

func TestStartHTTPApi(t *testing.T) {
	startedChan := make(chan struct{})
	go StartHTTPApi(FakeWebhookProvider{}, startedChan, 5*time.Second, 10*time.Second, "127.0.0.1:8887")
	<-startedChan
	resp, err := http.Get("http://127.0.0.1:8887")
	require.NoError(t, err)
	// check that resp has a valid domain filter
	defer resp.Body.Close()

	df := endpoint.DomainFilter{}
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, df.UnmarshalJSON(b))
}
