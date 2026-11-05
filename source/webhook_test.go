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

package source

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/wait"

	"sigs.k8s.io/external-dns/endpoint"
	extdnshttp "sigs.k8s.io/external-dns/pkg/http"
	"sigs.k8s.io/external-dns/provider"
)

const testWebhookMaxBodySize = 32 << 20

// fastWebhookBackoff keeps negotiation failures from taking seconds per test.
func fastWebhookBackoff(t *testing.T) {
	t.Helper()
	orig := webhookSourceBackoff
	webhookSourceBackoff = wait.Backoff{Duration: time.Millisecond, Factor: 1.5, Steps: 3}
	t.Cleanup(func() { webhookSourceBackoff = orig })
}

// newWebhookServer answers negotiation on "/" and delegates "/endpoints" to handler.
func newWebhookServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", WebhookSourceAccept)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		handler(w, r)
	}))
}

func serveEndpoints(body string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", WebhookSourceAccept)
		_, _ = w.Write([]byte(body))
	}
}

func newTestWebhookSource(t *testing.T, srv *httptest.Server) Source {
	t.Helper()
	fastWebhookBackoff(t)
	src, err := NewWebhookSource(t.Context(), []string{srv.URL}, 5*time.Second, testWebhookMaxBodySize)
	require.NoError(t, err)
	return src
}

func TestParseWebhookSourceURL(t *testing.T) {
	for _, raw := range []string{
		"http://localhost:8889",
		"http://127.0.0.1:8889",
		"http://127.1.2.3:8889/prefix",
		"http://[::1]:8889",
	} {
		_, err := parseWebhookSourceURL(raw)
		require.NoError(t, err, raw)
	}

	for raw, want := range map[string]string{
		"://bad":                   "invalid --webhook-source-url",
		"https://localhost:8889":   "must use the http scheme",
		"http://10.0.0.1:8889":     "must point at a loopback host",
		"http://webhook.svc:8889":  "must point at a loopback host",
		"http://[fe80::1]:8889":    "must point at a loopback host",
		"http://localhost.io:8889": "must point at a loopback host",
	} {
		_, err := parseWebhookSourceURL(raw)
		require.ErrorContains(t, err, want, raw)
	}
}

func TestNewWebhookSource_ValidatesBeforeConnecting(t *testing.T) {
	fastWebhookBackoff(t)

	_, err := NewWebhookSource(t.Context(), nil, time.Second, testWebhookMaxBodySize)
	require.ErrorContains(t, err, "at least one --webhook-source-url")

	_, err = NewWebhookSource(t.Context(), []string{"http://localhost:8889", "http://localhost:8889"}, time.Second, testWebhookMaxBodySize)
	require.ErrorContains(t, err, "duplicate --webhook-source-url")

	_, err = NewWebhookSource(t.Context(), []string{"http://localhost:8889", "https://localhost:8890"}, time.Second, testWebhookMaxBodySize)
	require.ErrorContains(t, err, "must use the http scheme")
}

func TestNewWebhookSource_Unreachable(t *testing.T) {
	fastWebhookBackoff(t)
	srv := httptest.NewServer(http.NotFoundHandler())
	srv.Close()

	_, err := NewWebhookSource(t.Context(), []string{srv.URL}, time.Second, testWebhookMaxBodySize)
	require.ErrorContains(t, err, "negotiation failed")
}

func TestNewWebhookSource_WrongMediaType(t *testing.T) {
	fastWebhookBackoff(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// A provider webhook answers with its own media type.
		w.Header().Set("Content-Type", "application/external.dns.webhook+json;version=1")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	_, err := NewWebhookSource(t.Context(), []string{srv.URL}, time.Second, testWebhookMaxBodySize)
	require.ErrorContains(t, err, "wrong content type")
}

func TestNewWebhookSource_RetriesNegotiation(t *testing.T) {
	fastWebhookBackoff(t)
	attempts := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		assert.Equal(t, WebhookSourceAccept, r.Header.Get("Accept"))
		if attempts < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", WebhookSourceAccept+"; charset=utf-8")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	_, err := NewWebhookSource(t.Context(), []string{srv.URL}, time.Second, testWebhookMaxBodySize)
	require.NoError(t, err)
	assert.Equal(t, 2, attempts)
}

func TestWebhookSource_Endpoints(t *testing.T) {
	srv := newWebhookServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/endpoints", r.URL.Path)
		assert.Equal(t, WebhookSourceAccept, r.Header.Get("Accept"))
		w.Header().Set("Content-Type", WebhookSourceAccept+"; charset=utf-8")
		_, _ = w.Write([]byte(`[
			{"dnsName": "app.example.com", "recordType": "A", "targets": ["192.0.2.10"], "recordTTL": 300,
			 "labels": {"resource": "ingress/default/forged", "owner": "someone-else"}},
			{"dnsName": "_acme-challenge.app.example.com", "recordType": "CNAME", "targets": ["app.example.com.acme.example.net"],
			 "setIdentifier": "id-1", "providerSpecific": [{"name": "alias", "value": "false"}]}
		]`))
	})
	defer srv.Close()
	src := newTestWebhookSource(t, srv)

	eps, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	host := srv.Listener.Addr().String()
	require.Equal(t, []*endpoint.Endpoint{
		{
			DNSName:    "app.example.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"192.0.2.10"},
			RecordTTL:  300,
			Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "webhook/" + host},
		},
		{
			DNSName:          "_acme-challenge.app.example.com",
			RecordType:       endpoint.RecordTypeCNAME,
			Targets:          endpoint.Targets{"app.example.com.acme.example.net"},
			SetIdentifier:    "id-1",
			ProviderSpecific: endpoint.ProviderSpecific{{Name: "alias", Value: "false"}},
			Labels:           endpoint.Labels{endpoint.ResourceLabelKey: "webhook/" + host},
		},
	}, eps)
}

func TestWebhookSource_EmptyListIsValid(t *testing.T) {
	srv := newWebhookServer(t, serveEndpoints(`[]`))
	defer srv.Close()
	src := newTestWebhookSource(t, srv)

	eps, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	assert.Empty(t, eps)
}

func TestWebhookSource_SoftErrors(t *testing.T) {
	tests := map[string]struct {
		handler http.HandlerFunc
		want    string
	}{
		"not ready": {
			handler: func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusServiceUnavailable) },
			want:    "not ready",
		},
		"server error": {
			handler: func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusInternalServerError) },
			want:    "unexpected status code 500",
		},
		"not found": {
			handler: func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNotFound) },
			want:    "unexpected status code 404",
		},
		"redirect is not followed": {
			handler: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "http://10.0.0.1/endpoints", http.StatusFound)
			},
			want: "unexpected status code 302",
		},
		"wrong media type": {
			handler: func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`[]`))
			},
			want: "wrong content type",
		},
		"wrong version": {
			handler: func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", WebhookSourceMediaType+";version=2")
				_, _ = w.Write([]byte(`[]`))
			},
			want: "wrong content type",
		},
		"invalid json": {
			handler: serveEndpoints(`not json`),
			want:    "decoding response",
		},
		"null body": {
			handler: serveEndpoints(`null`),
			want:    "response body is null",
		},
		"null element": {
			handler: serveEndpoints(`[null]`),
			want:    "null endpoint",
		},
		"unsupported record type": {
			handler: serveEndpoints(`[{"dnsName": "a.example.com", "recordType": "SPF", "targets": ["v=spf1"]}]`),
			want:    `unsupported record type "SPF"`,
		},
		"invalid endpoint rejects the whole response": {
			handler: serveEndpoints(`[
				{"dnsName": "ok.example.com", "recordType": "A", "targets": ["192.0.2.1"]},
				{"dnsName": "bad.example.com", "recordType": "A", "targets": ["not-an-ip"]}
			]`),
			want: `endpoint "bad.example.com"`,
		},
		"missing targets": {
			handler: serveEndpoints(`[{"dnsName": "a.example.com", "recordType": "TXT"}]`),
			want:    "needs at least one target",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			srv := newWebhookServer(t, tc.handler)
			defer srv.Close()
			src := newTestWebhookSource(t, srv)

			eps, err := src.Endpoints(t.Context())
			require.ErrorIs(t, err, provider.SoftError)
			require.ErrorContains(t, err, tc.want)
			assert.Nil(t, eps)
		})
	}
}

func TestWebhookSource_ConnectionErrorIsSoft(t *testing.T) {
	srv := newWebhookServer(t, serveEndpoints(`[]`))
	src := newTestWebhookSource(t, srv)
	srv.Close()

	_, err := src.Endpoints(t.Context())
	require.ErrorIs(t, err, provider.SoftError)
}

func TestWebhookSource_BodyOverLimit(t *testing.T) {
	const limit = 1 << 20
	srv := newWebhookServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", WebhookSourceAccept)
		_, _ = w.Write([]byte(`[{"dnsName":"`))
		_, _ = w.Write(bytes.Repeat([]byte("a"), 2*limit))
		_, _ = w.Write([]byte(`","recordType":"A","targets":["192.0.2.1"]}]`))
	})
	defer srv.Close()
	fastWebhookBackoff(t)
	src, err := NewWebhookSource(t.Context(), []string{srv.URL}, 5*time.Second, limit)
	require.NoError(t, err)

	_, err = src.Endpoints(t.Context())
	require.ErrorIs(t, err, provider.SoftError)
	require.ErrorIs(t, err, extdnshttp.ErrResponseTooLarge)
}

func TestWebhookSource_MultipleWebhooks(t *testing.T) {
	first := newWebhookServer(t, serveEndpoints(`[{"dnsName": "one.example.com", "recordType": "A", "targets": ["192.0.2.1"]}]`))
	defer first.Close()
	second := newWebhookServer(t, serveEndpoints(`[{"dnsName": "two.example.com", "recordType": "A", "targets": ["192.0.2.2"]}]`))
	defer second.Close()
	fastWebhookBackoff(t)
	src, err := NewWebhookSource(t.Context(), []string{first.URL, second.URL}, 5*time.Second, testWebhookMaxBodySize)
	require.NoError(t, err)

	eps, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	require.Len(t, eps, 2)
	assert.Equal(t, "webhook/"+first.Listener.Addr().String(), eps[0].Labels[endpoint.ResourceLabelKey])
	assert.Equal(t, "webhook/"+second.Listener.Addr().String(), eps[1].Labels[endpoint.ResourceLabelKey])

	// One failing webhook fails the whole iteration; a partial result would
	// delete the records of the webhook that is down.
	second.Close()
	eps, err = src.Endpoints(t.Context())
	require.ErrorIs(t, err, provider.SoftError)
	assert.Nil(t, eps)
}

func TestWebhookSource_AddEventHandlerIsNoop(t *testing.T) {
	src := &webhookSource{}
	src.AddEventHandler(t.Context(), func() { t.Fatal("handler must never be called") })
}

func TestIsWebhookSourceMediaType(t *testing.T) {
	assert.True(t, isWebhookSourceMediaType(WebhookSourceAccept))
	assert.True(t, isWebhookSourceMediaType(WebhookSourceMediaType+"; version=1; charset=utf-8"))
	assert.True(t, isWebhookSourceMediaType("Application/External.DNS.Webhook.Source+JSON;version=1"))
	assert.False(t, isWebhookSourceMediaType(WebhookSourceMediaType))
	assert.False(t, isWebhookSourceMediaType(WebhookSourceMediaType+";version=2"))
	assert.False(t, isWebhookSourceMediaType("application/external.dns.webhook+json;version=1"))
	assert.False(t, isWebhookSourceMediaType(""))
}
