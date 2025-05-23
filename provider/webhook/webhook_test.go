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
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	webhookapi "sigs.k8s.io/external-dns/provider/webhook/api"
)

func TestNewWebhookProvider_InvalidURL(t *testing.T) {
	_, err := NewWebhookProvider("://invalid-url")
	require.Error(t, err)
}

func TestNewWebhookProvider_HTTPRequestFailure(t *testing.T) {
	_, err := NewWebhookProvider("http://nonexistent.url")
	require.Error(t, err)
}

func TestNewWebhookProvider_InvalidResponseBody(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid-json")) // Invalid JSON
	}))
	defer svr.Close()

	_, err := NewWebhookProvider(svr.URL)
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to unmarshal response body of DomainFilter")
}

func TestNewWebhookProvider_Non2XXStatusCode(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer svr.Close()

	_, err := NewWebhookProvider(svr.URL)
	require.Error(t, err)
	require.Contains(t, err.Error(), "status code < 500")
}

func TestNewWebhookProvider_WrongContentTypeHeader(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion+"wrong")
			_, _ = w.Write([]byte(`{}`))
			return
		}
	}))
	defer svr.Close()

	_, err := NewWebhookProvider(svr.URL)
	require.Error(t, err)
	require.Contains(t, err.Error(), "wrong content type returned from server")
}

func TestInvalidDomainFilter(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.WriteHeader(http.StatusOK)
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
		assert.Equal(t, "/records", r.URL.Path)
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
		assert.Equal(t, "/records", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()

	p, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	_, err = p.Records(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)
}

func TestRecords_HTTPRequestErrorMissingHost0(t *testing.T) {
	wpr := WebhookProvider{
		remoteServerURL: &url.URL{Scheme: "http", Host: "example\\x00.com", Path: "\\x00"},
		client:          &http.Client{},
	}

	_, err := wpr.Records(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid URL escape")
}

func TestRecords_HTTPRequestErrorMissingHost(t *testing.T) {
	wpr := WebhookProvider{
		remoteServerURL: &url.URL{Host: "example.com", Path: "\\x00"},
		client:          &http.Client{},
	}

	_, err := wpr.Records(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestRecords_DecodeError(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == webhookapi.UrlRecords {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("invalid-json")) // Simulate invalid JSON response
			return
		}
	}))
	defer svr.Close()

	parsedURL, _ := url.Parse(svr.URL)
	p := WebhookProvider{
		remoteServerURL: parsedURL,
		client:          &http.Client{},
	}

	_, err := p.Records(context.Background())
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid character 'i' looking for beginning of value")
}

func TestRecords_NonOKStatusCode(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}))
	defer svr.Close()

	parsedURL, _ := url.Parse(svr.URL)

	p := WebhookProvider{
		remoteServerURL: &url.URL{Scheme: parsedURL.Scheme, Host: parsedURL.Host},
		client:          &http.Client{},
	}

	_, err := p.Records(nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get records with code 511")
}

func TestApplyChanges(t *testing.T) {
	successfulApplyChanges := true
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		assert.Equal(t, "/records", r.URL.Path)
		if successfulApplyChanges {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer svr.Close()

	p, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	err = p.ApplyChanges(context.TODO(), nil)
	require.NoError(t, err)

	successfulApplyChanges = false

	err = p.ApplyChanges(context.TODO(), nil)
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)
}

func TestApplyChanges_HTTPNewRequestErrorWrongHost(t *testing.T) {
	wpr := WebhookProvider{
		remoteServerURL: &url.URL{Host: "exa\\x00mple.com"},
		client:          &http.Client{},
	}

	err := wpr.ApplyChanges(context.Background(), nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid URL escape")
}

func TestApplyChanges_GetFailed(t *testing.T) {
	p := WebhookProvider{
		remoteServerURL: &url.URL{Host: "localhost"},
		client:          &http.Client{},
	}

	err := p.ApplyChanges(context.TODO(), &plan.Changes{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestApplyChanges_StatusCodeError(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		assert.Equal(t, webhookapi.UrlRecords, r.URL.Path)
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
	}))
	defer svr.Close()

	p, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)

	err = p.ApplyChanges(context.TODO(), nil)
	require.Error(t, err)
	require.NotErrorIs(t, err, provider.SoftError)
	assert.Contains(t, err.Error(), "failed to apply changes with code 511")
}

func TestAdjustEndpoints(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		assert.Equal(t, webhookapi.UrlAdjustEndpoints, r.URL.Path)

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
		assert.Equal(t, webhookapi.UrlAdjustEndpoints, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer svr.Close()

	p, err := NewWebhookProvider(svr.URL)
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
	_, err = p.AdjustEndpoints(endpoints)
	require.Error(t, err)
	require.ErrorIs(t, err, provider.SoftError)
}

// test apply changes with an endpoint with a provider specific property
func TestApplyChangesWithProviderSpecificProperty(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.Write([]byte(`{}`))
			return
		}
		if r.URL.Path == "/records" {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			// assert that the request contains the provider-specific property
			var changes plan.Changes
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			err = json.Unmarshal(b, &changes)
			assert.NoError(t, err)
			assert.Len(t, changes.Create, 1)
			assert.Len(t, changes.Create[0].ProviderSpecific, 1)
			assert.Equal(t, "prop1", changes.Create[0].ProviderSpecific[0].Name)
			assert.Equal(t, "value1", changes.Create[0].ProviderSpecific[0].Value)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}))
	defer svr.Close()

	p, err := NewWebhookProvider(svr.URL)
	require.NoError(t, err)
	e := &endpoint.Endpoint{
		DNSName:    "test.example.com",
		RecordTTL:  10,
		RecordType: "A",
		Targets: endpoint.Targets{
			"",
		},
		ProviderSpecific: endpoint.ProviderSpecific{
			endpoint.ProviderSpecificProperty{
				Name:  "prop1",
				Value: "value1",
			},
		},
	}
	err = p.ApplyChanges(context.TODO(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			e,
		},
	})
	require.NoError(t, err)
}

func TestAdjustEndpoints_JoinPathError(t *testing.T) {
	wpr := WebhookProvider{
		remoteServerURL: &url.URL{Scheme: "http", Host: "example\\x00.com"},
	}

	_, err := wpr.AdjustEndpoints(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid URL escape")
}

func TestAdjustEndpoints_HTTPRequestErrorMissingHost(t *testing.T) {
	wpr := WebhookProvider{
		remoteServerURL: &url.URL{Host: "example.com", Path: "\\x00"},
		client:          &http.Client{},
	}

	_, err := wpr.AdjustEndpoints(nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported protocol scheme") // Ensure the "BINGO" log is triggered
}

func TestAdjustEndpoints_NonOKStatusCode(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}))
	defer svr.Close()

	parsedURL, _ := url.Parse(svr.URL)

	p := WebhookProvider{
		remoteServerURL: &url.URL{Scheme: parsedURL.Scheme, Host: parsedURL.Host},
		client:          &http.Client{},
	}

	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "test.example.com",
			RecordTTL:  10,
			RecordType: "A",
			Targets:    endpoint.Targets{""},
		},
	}

	_, err := p.AdjustEndpoints(endpoints)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to AdjustEndpoints with code  511")
}

func TestAdjustEndpoints_DecodeError(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == webhookapi.UrlAdjustEndpoints {
			w.Header().Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("invalid-json")) // Simulate invalid JSON response
			return
		}
	}))
	defer svr.Close()

	parsedURL, _ := url.Parse(svr.URL)
	p := WebhookProvider{
		remoteServerURL: parsedURL,
		client:          &http.Client{},
	}

	var endpoints []*endpoint.Endpoint

	_, err := p.AdjustEndpoints(endpoints)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid character 'i' looking for beginning of value")
}

func TestRequestWithRetry_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := requestWithRetry(client, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRequestWithRetry_NonRetriableStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := requestWithRetry(client, req)
	require.Error(t, err)
	require.Nil(t, resp)
}
