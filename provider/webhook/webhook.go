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
	"fmt"
	"net/http"
	"net/url"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/metrics"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	webhookapi "sigs.k8s.io/external-dns/provider/webhook/api"

	"github.com/cenkalti/backoff/v5"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	acceptHeader = "Accept"
	maxRetries   = 5
)

var (
	recordsErrorsGauge = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "records_errors_total",
			Help:      "Errors with Records method",
		},
	)
	recordsRequestsGauge = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "records_requests_total",
			Help:      "Requests with Records method",
		},
	)
	applyChangesErrorsGauge = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "applychanges_errors_total",
			Help:      "Errors with ApplyChanges method",
		},
	)
	applyChangesRequestsGauge = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "applychanges_requests_total",
			Help:      "Requests with ApplyChanges method",
		},
	)
	adjustEndpointsErrorsGauge = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "adjustendpoints_errors_total",
			Help:      "Errors with AdjustEndpoints method",
		},
	)
	adjustEndpointsRequestsGauge = metrics.NewGaugeWithOpts(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "adjustendpoints_requests_total",
			Help:      "Requests with AdjustEndpoints method",
		},
	)
)

type WebhookProvider struct {
	client          *http.Client
	remoteServerURL *url.URL
	DomainFilter    *endpoint.DomainFilter
}

func init() {
	metrics.RegisterMetric.MustRegister(recordsErrorsGauge)
	metrics.RegisterMetric.MustRegister(recordsRequestsGauge)
	metrics.RegisterMetric.MustRegister(applyChangesErrorsGauge)
	metrics.RegisterMetric.MustRegister(applyChangesRequestsGauge)
	metrics.RegisterMetric.MustRegister(adjustEndpointsErrorsGauge)
	metrics.RegisterMetric.MustRegister(adjustEndpointsRequestsGauge)
}

func NewWebhookProvider(u string) (*WebhookProvider, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	// negotiate API information
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(acceptHeader, webhookapi.MediaTypeFormatAndVersion)

	client := &http.Client{}

	resp, err := requestWithRetry(client, req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to webhook: %w", err)
	}
	// read the serialized DomainFilter from the response body and set it in the webhook provider struct
	defer resp.Body.Close()

	if ct := resp.Header.Get(webhookapi.ContentTypeHeader); ct != webhookapi.MediaTypeFormatAndVersion {
		return nil, fmt.Errorf("wrong content type returned from server: %s", ct)
	}

	df := &endpoint.DomainFilter{}
	if err := json.NewDecoder(resp.Body).Decode(df); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body of DomainFilter: %w", err)
	}

	return &WebhookProvider{
		client:          client,
		remoteServerURL: parsedURL,
		DomainFilter:    df,
	}, nil
}

func requestWithRetry(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := backoff.Retry(context.Background(), func() (*http.Response, error) {
		resp, err := client.Do(req)
		if err != nil {
			log.Debugf("Failed to connect to webhook: %v", err)
			return nil, err
		}
		// we currently only use 200 as success, but considering okay all 2XX for future usage
		if resp.StatusCode >= http.StatusMultipleChoices && resp.StatusCode < http.StatusInternalServerError {
			return nil, backoff.Permanent(fmt.Errorf("status code < %d", http.StatusInternalServerError))
		}
		return resp, nil
	}, backoff.WithMaxTries(maxRetries))
	return resp, err
}

// Records will make a GET call to remoteServerURL/records and return the results
func (p WebhookProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	recordsRequestsGauge.Gauge.Inc()
	u := p.remoteServerURL.JoinPath("records").String()

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		recordsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to create request: %s", err.Error())
		return nil, err
	}
	req.Header.Set(acceptHeader, webhookapi.MediaTypeFormatAndVersion)
	resp, err := p.client.Do(req)
	if err != nil {
		recordsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to perform request: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		recordsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to get records with code %d", resp.StatusCode)
		err := fmt.Errorf("failed to get records with code %d", resp.StatusCode)
		if isRetryableError(resp.StatusCode) {
			return nil, provider.NewSoftError(err)
		}
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	if err := json.NewDecoder(resp.Body).Decode(&endpoints); err != nil {
		recordsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to decode response body: %s", err.Error())
		return nil, err
	}
	return endpoints, nil
}

// ApplyChanges will make a POST to remoteServerURL/records with the changes
func (p WebhookProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	applyChangesRequestsGauge.Gauge.Inc()
	u := p.remoteServerURL.JoinPath(webhookapi.UrlRecords).String()

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(changes); err != nil {
		applyChangesErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to encode changes: %s", err.Error())
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u, b)
	if err != nil {
		applyChangesErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to create request: %s", err.Error())
		return err
	}

	req.Header.Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)

	resp, err := p.client.Do(req)
	if err != nil {
		applyChangesErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to perform request: %s", err.Error())
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		applyChangesErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to apply changes with code %d", resp.StatusCode)
		err := fmt.Errorf("failed to apply changes with code %d", resp.StatusCode)
		if isRetryableError(resp.StatusCode) {
			return provider.NewSoftError(err)
		}
		return err
	}
	return nil
}

// AdjustEndpoints will call the provider doing a POST on `/adjustendpoints` which will return a list of modified endpoints
// based on a provider-specific requirement.
// This method returns an empty slice in case there is a technical error on the provider's side so that no endpoints will be considered.
func (p WebhookProvider) AdjustEndpoints(e []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	adjustEndpointsRequestsGauge.Gauge.Inc()
	var endpoints []*endpoint.Endpoint
	u, err := url.JoinPath(p.remoteServerURL.String(), webhookapi.UrlAdjustEndpoints)
	if err != nil {
		adjustEndpointsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to join path, %s", err)
		return nil, err
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(e); err != nil {
		adjustEndpointsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to encode endpoints, %s", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u, b)
	if err != nil {
		adjustEndpointsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to create new HTTP request, %s", err)
		return nil, err
	}

	req.Header.Set(webhookapi.ContentTypeHeader, webhookapi.MediaTypeFormatAndVersion)
	req.Header.Set(acceptHeader, webhookapi.MediaTypeFormatAndVersion)

	resp, err := p.client.Do(req)
	if err != nil {
		adjustEndpointsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed executing http request, %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		adjustEndpointsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to AdjustEndpoints with code %d", resp.StatusCode)
		err := fmt.Errorf("failed to AdjustEndpoints with code  %d", resp.StatusCode)
		if isRetryableError(resp.StatusCode) {
			return nil, provider.NewSoftError(err)
		}
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&endpoints); err != nil {
		adjustEndpointsErrorsGauge.Gauge.Inc()
		log.Debugf("Failed to decode response body: %s", err.Error())
		return nil, err
	}

	return endpoints, nil
}

// GetDomainFilter make calls to get the serialized version of the domain filter
func (p WebhookProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return p.DomainFilter
}

// isRetryableError returns true for HTTP status codes between 500 and 510 (inclusive)
func isRetryableError(statusCode int) bool {
	return statusCode >= http.StatusInternalServerError && statusCode <= http.StatusNotExtended
}
