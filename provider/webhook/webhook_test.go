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
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	webhookapi "sigs.k8s.io/external-dns/provider/webhook/api"
)

func TestInvalidDomainFilter(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.WriteHeader(200)
			return
		}
		w.Write([]byte(`[{
			"dnsName" : "test.example.com"
		}]`))
	}))
	defer svr.Close()

	_, err := NewWebhookProvider(svr.URL)
	require.Error(t, err)
}

func TestValidDomainfilter(t *testing.T) {
	// initialize domain filter
	domainFilter := endpoint.NewDomainFilter([]string{"example.com"})
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			json.NewEncoder(w).Encode(domainFilter)
			return
		}
	}))
	defer svr.Close()

	p, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	require.Equal(t, p.GetDomainFilter(), endpoint.NewDomainFilter([]string{"example.com"}))
}

func TestRecords(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		require.Equal(t, "/records", r.URL.Path)
		w.Write([]byte(`[{
			"dnsName" : "test.example.com"
		}]`))
	}))
	defer svr.Close()

	provider, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	endpoints, err := provider.Records(context.TODO())
	require.NoError(t, err)
	require.NotNil(t, endpoints)
	require.Equal(t, []*endpoint.Endpoint{{
		DNSName: "test.example.com",
	}}, endpoints)
}

func TestRecordsWithErrors(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		require.Equal(t, "/records", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()

	provider, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	_, err = provider.Records(context.Background())
	require.NotNil(t, err)
}

func TestApplyChanges(t *testing.T) {
	successfulApplyChanges := true
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		require.Equal(t, "/records", r.URL.Path)
		if successfulApplyChanges {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer svr.Close()

	provider, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	err = provider.ApplyChanges(context.TODO(), nil)
	require.NoError(t, err)

	successfulApplyChanges = false

	err = provider.ApplyChanges(context.TODO(), nil)
	require.NotNil(t, err)
}

func TestAdjustEndpoints(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		require.Equal(t, "/adjustendpoints", r.URL.Path)

		var endpoints []*endpoint.Endpoint
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(b, &endpoints)
		if err != nil {
			t.Fatal(err)
		}

		for _, e := range endpoints {
			e.RecordTTL = 0
		}
		j, _ := json.Marshal(endpoints)
		w.Write(j)

	}))
	defer svr.Close()

	provider, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "test.example.com",
			RecordTTL:  10,
			RecordType: "A",
			Targets: endpoint.Targets{
				"",
			},
		},
	}
	adjustedEndpoints, err := provider.AdjustEndpoints(endpoints)
	require.NoError(t, err)
	require.Equal(t, []*endpoint.Endpoint{{
		DNSName:    "test.example.com",
		RecordTTL:  0,
		RecordType: "A",
		Targets: endpoint.Targets{
			"",
		},
	}}, adjustedEndpoints)
}

func TestAdjustendpointsWithError(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		require.Equal(t, "/adjustendpoints", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()

	provider, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "test.example.com",
			RecordTTL:  10,
			RecordType: "A",
			Targets: endpoint.Targets{
				"",
			},
		},
	}
	_, err = provider.AdjustEndpoints(endpoints)
	require.Error(t, err)
}
