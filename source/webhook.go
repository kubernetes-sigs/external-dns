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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"net/netip"
	"net/url"
	"slices"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/sets"
	extdnshttp "sigs.k8s.io/external-dns/pkg/http"
	"sigs.k8s.io/external-dns/pkg/metrics"
	"sigs.k8s.io/external-dns/provider"
)

const (
	// WebhookSourceMediaType is the media type a source webhook answers with. It
	// differs from the provider webhook's so that a URL pointing at the wrong
	// kind of webhook fails negotiation instead of misbehaving later.
	WebhookSourceMediaType = "application/external.dns.webhook.source+json"
	// WebhookSourceMediaTypeVersion is the contract version this client speaks.
	WebhookSourceMediaTypeVersion = "1"
	// WebhookSourceAccept is the exact Accept header value ExternalDNS sends.
	WebhookSourceAccept = WebhookSourceMediaType + ";version=" + WebhookSourceMediaTypeVersion

	webhookSourceEndpointsPath = "/endpoints"
	webhookSourceSubsystem     = "webhook_source"
)

var (
	webhookSourceRequests = metrics.NewCounterVecWithOpts(
		prometheus.CounterOpts{
			Subsystem: webhookSourceSubsystem,
			Name:      "endpoints_requests_total",
			Help:      "Requests to the /endpoints route of a source webhook",
		},
		[]string{"webhook"},
	)
	webhookSourceErrors = metrics.NewCounterVecWithOpts(
		prometheus.CounterOpts{
			Subsystem: webhookSourceSubsystem,
			Name:      "endpoints_errors_total",
			Help:      "Failed requests to the /endpoints route of a source webhook",
		},
		[]string{"webhook"},
	)

	// webhookSourceBackoff paces the start-up negotiation; a sidecar usually
	// listens within a few seconds of the ExternalDNS container.
	webhookSourceBackoff = wait.Backoff{Duration: 250 * time.Millisecond, Factor: 1.5, Jitter: 2, Steps: 5}
)

func init() {
	metrics.RegisterMetric.MustRegister(webhookSourceRequests)
	metrics.RegisterMetric.MustRegister(webhookSourceErrors)
}

// webhookSource fetches the desired endpoints from one or more out-of-process
// webhooks over HTTP/JSON, the source-side counterpart of the provider webhook.
//
// +externaldns:source:name=webhook
// +externaldns:source:category=Special
// +externaldns:source:description=Fetches DNS endpoints over HTTP/JSON from a webhook sidecar on loopback
// +externaldns:source:resources=Remote HTTP server
// +externaldns:source:filters=
// +externaldns:source:namespace=
// +externaldns:source:fqdn-template=false
// +externaldns:source:events=false
// +externaldns:source:provider-specific=false
type webhookSource struct {
	webhooks []*webhookClient
}

// webhookClient talks to a single source webhook.
type webhookClient struct {
	client  *http.Client
	baseURL *url.URL
	// name is the host:port of the webhook; it labels the metrics and the
	// resource label of every endpoint the webhook returns.
	name        string
	maxBodySize int64
}

// NewWebhookSource validates the URLs, negotiates the media type with every
// webhook and returns a Source that polls all of them on each reconcile.
func NewWebhookSource(ctx context.Context, urls []string, timeout time.Duration, maxBodySize int64) (Source, error) {
	if len(urls) == 0 {
		return nil, errors.New("webhook source: at least one --webhook-source-url is required")
	}
	seen := sets.New[string]()
	parsed := make([]*url.URL, 0, len(urls))
	for _, raw := range urls {
		if seen.Has(raw) {
			return nil, fmt.Errorf("webhook source: duplicate --webhook-source-url %q", raw)
		}
		seen.Insert(raw)
		u, err := parseWebhookSourceURL(raw)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, u)
	}

	webhooks := make([]*webhookClient, 0, len(parsed))
	for _, u := range parsed {
		wc := &webhookClient{
			client: extdnshttp.NewInstrumentedClient(&http.Client{
				// Timeout covers the whole round trip, including reading the body.
				Timeout: timeout,
				// A redirect could point off loopback, so it is reported instead of followed.
				CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
			}),
			baseURL:     u,
			name:        u.Host,
			maxBodySize: maxBodySize,
		}
		if err := wc.negotiate(ctx); err != nil {
			return nil, err
		}
		webhooks = append(webhooks, wc)
	}
	return &webhookSource{webhooks: webhooks}, nil
}

// parseWebhookSourceURL enforces the v1 contract: plain HTTP on a loopback
// host. Nothing authenticates the caller, so a webhook reachable from the
// network would let any client decide which records ExternalDNS manages.
func parseWebhookSourceURL(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("webhook source: invalid --webhook-source-url %q: %w", raw, err)
	}
	if u.Scheme != "http" {
		return nil, fmt.Errorf("webhook source: --webhook-source-url %q must use the http scheme", raw)
	}
	if host := u.Hostname(); host != "localhost" {
		addr, err := netip.ParseAddr(host)
		if err != nil || !addr.IsLoopback() {
			return nil, fmt.Errorf("webhook source: --webhook-source-url %q must point at a loopback host (localhost, 127.0.0.0/8 or ::1)", raw)
		}
	}
	return u, nil
}

// negotiate checks the media type once at start-up, so a URL that points at
// nothing or at a provider webhook fails fast instead of on the first reconcile.
func (wc *webhookClient) negotiate(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, wc.baseURL.String(), nil)
	if err != nil {
		return fmt.Errorf("webhook source %s: %w", wc.name, err)
	}
	req.Header.Set("Accept", WebhookSourceAccept)
	resp, err := extdnshttp.RequestWithRetry(wc.client, req, webhookSourceBackoff)
	if err != nil {
		return fmt.Errorf("webhook source %s: negotiation failed: %w", wc.name, err)
	}
	defer extdnshttp.DrainAndClose(resp.Body)
	if ct := resp.Header.Get("Content-Type"); !isWebhookSourceMediaType(ct) {
		return fmt.Errorf("webhook source %s: wrong content type %q, want %s", wc.name, ct, WebhookSourceAccept)
	}
	return nil
}

// isWebhookSourceMediaType accepts the media type with any extra parameters,
// such as the charset many HTTP frameworks append, as long as the version matches.
func isWebhookSourceMediaType(contentType string) bool {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	return mediaType == WebhookSourceMediaType && params["version"] == WebhookSourceMediaTypeVersion
}

// Endpoints returns the endpoints of every webhook. It fails as a whole when
// one webhook fails: applying the others' endpoints would delete the records
// owned by the one that is down.
func (ws *webhookSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	for _, wc := range ws.webhooks {
		eps, err := wc.endpoints(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, eps...)
	}
	return result, nil
}

// AddEventHandler is a no-op: the contract has no change notifications, so a
// change shows up on the next --interval tick.
func (ws *webhookSource) AddEventHandler(_ context.Context, _ func()) {}

func (wc *webhookClient) endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	webhookSourceRequests.CounterVec.WithLabelValues(wc.name).Inc()
	eps, err := wc.fetch(ctx)
	if err != nil {
		webhookSourceErrors.CounterVec.WithLabelValues(wc.name).Inc()
		// Every runtime failure is soft: a plain error would exit the process
		// for all sources, and an empty list would delete this webhook's records.
		return nil, provider.NewSoftErrorf("webhook source %s: %w", wc.name, err)
	}
	return eps, nil
}

func (wc *webhookClient) fetch(ctx context.Context) ([]*endpoint.Endpoint, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, wc.baseURL.JoinPath(webhookSourceEndpointsPath).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", WebhookSourceAccept)
	resp, err := wc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer extdnshttp.DrainAndClose(resp.Body)

	if resp.StatusCode == http.StatusServiceUnavailable {
		return nil, errors.New("webhook is not ready yet (503)")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	if ct := resp.Header.Get("Content-Type"); !isWebhookSourceMediaType(ct) {
		return nil, fmt.Errorf("wrong content type %q", ct)
	}

	var eps []*endpoint.Endpoint
	if err := json.NewDecoder(extdnshttp.LimitBody(resp.Body, wc.maxBodySize)).Decode(&eps); err != nil {
		if errors.Is(err, extdnshttp.ErrResponseTooLarge) {
			return nil, fmt.Errorf("%w (--webhook-source-max-body-size)", err)
		}
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	// A JSON null decodes to a nil slice without error; it is not an empty list.
	if eps == nil {
		return nil, errors.New("response body is null, expected a JSON array")
	}
	// One bad entry rejects the whole response. Dropping only that entry would
	// turn a malformed record into the deletion of its existing one at the next plan.
	for _, ep := range eps {
		if ep == nil {
			return nil, errors.New("response contains a null endpoint")
		}
		// The crd source gets these from the API server's schema; a webhook has no schema in front of it.
		if ep.DNSName == "" {
			return nil, errors.New("endpoint without a dnsName")
		}
		if label := endpoint.OverlongDNSLabel(ep.DNSName); label != "" {
			return nil, fmt.Errorf("endpoint %q: label %q is longer than 63 characters", ep.DNSName, label)
		}
		if !slices.Contains(endpoint.KnownRecordTypes, ep.RecordType) {
			return nil, fmt.Errorf("endpoint %q: unknown record type %q", ep.DNSName, ep.RecordType)
		}
		if ep.RecordTTL < 0 {
			return nil, fmt.Errorf("endpoint %q: negative TTL %d", ep.DNSName, ep.RecordTTL)
		}
		if err := ep.Validate(); err != nil {
			return nil, fmt.Errorf("endpoint %q: %w", ep.DNSName, err)
		}
		// Labels from the webhook are discarded. The resource label decides
		// conflicts and ownership, and a webhook must not impersonate another source.
		ep.Labels = endpoint.Labels{endpoint.ResourceLabelKey: "webhook/" + wc.name}
	}
	log.Debugf("webhook source %s: received %d endpoints", wc.name, len(eps))
	return eps, nil
}
