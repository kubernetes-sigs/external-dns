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

package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type FakeWebhookProvider struct{}

func (p FakeWebhookProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return []*endpoint.Endpoint{}, nil
}

func (p FakeWebhookProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return nil
}

func (p FakeWebhookProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	return endpoints
}

func (p FakeWebhookProvider) GetDomainFilter() endpoint.DomainFilter {
	return endpoint.DomainFilter{}
}

func TestRecordsHandlerRecords(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/records", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		provider: &FakeWebhookProvider{},
	}
	providerAPIServer.recordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestRecordsHandlerApplyChangesWithBadRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/applychanges", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		provider: &FakeWebhookProvider{},
	}
	providerAPIServer.recordsHandler(w, req)
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
	require.Nil(t, err)

	reader := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodPost, "/applychanges", reader)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		provider: &FakeWebhookProvider{},
	}
	providerAPIServer.recordsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestAdjustEndpointsHandlerWithInvalidRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/adjustendpoints", nil)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		provider: &FakeWebhookProvider{},
	}
	providerAPIServer.adjustEndpointsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)

	req = httptest.NewRequest(http.MethodGet, "/adjustendpoints", nil)

	providerAPIServer.adjustEndpointsHandler(w, req)
	res = w.Result()
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestAdjustEndpointsWithValidRequest(t *testing.T) {
	pve := []*endpoint.Endpoint{
		{
			DNSName:    "foo.bar.com",
			RecordType: "A",
			Targets:    endpoint.Targets{},
			RecordTTL:  0,
		},
	}

	j, err := json.Marshal(pve)
	require.Nil(t, err)

	reader := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/adjustendpoints", reader)
	w := httptest.NewRecorder()

	providerAPIServer := &WebhookServer{
		provider: &FakeWebhookProvider{},
	}
	providerAPIServer.adjustEndpointsHandler(w, req)
	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
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
