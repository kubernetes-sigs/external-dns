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
	"sigs.k8s.io/external-dns/plan"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	mediaTypeFormatAndVersion = "application/external.dns.webhook+json;version=1"
	contentTypeHeader         = "Content-Type"
	acceptHeader              = "Accept"
	varyHeader                = "Vary"
	maxRetries                = 5
)

var (
	recordsErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "records_errors",
			Help:      "Errors with Records method",
		},
	)
	applyChangesErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "applychanges_errors",
			Help:      "Errors with ApplyChanges method",
		},
	)
	propertyValuesEqualErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "propertyvaluesequal_errors",
			Help:      "Errors with PropertyValuesEqual method",
		},
	)
	adjustEndpointsErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "webhook_provider",
			Name:      "adjustendpointsgauge_errors",
			Help:      "Errors with AdjustEndpoints method",
		},
	)
)

type WebhookProvider struct {
	client          *http.Client
	remoteServerURL *url.URL
}

type PropertyValuesEqualRequest struct {
	Name     string `json:"name"`
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

type PropertyValuesEqualResponse struct {
	Equals bool `json:"equals"`
}

func init() {
	prometheus.MustRegister(recordsErrorsGauge)
	prometheus.MustRegister(applyChangesErrorsGauge)
	prometheus.MustRegister(propertyValuesEqualErrorsGauge)
	prometheus.MustRegister(adjustEndpointsErrorsGauge)
}

func NewWebhookProvider(u string) (*WebhookProvider, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	return &WebhookProvider{
		client:          client,
		remoteServerURL: parsedURL,
	}, nil
}

// Records will make a GET call to remoteServerURL/records and return the results
func (p WebhookProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	u := p.remoteServerURL.JoinPath("records").String()
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to create request: %s", err.Error())
		return nil, err
	}
	req.Header.Set(acceptHeader, mediaTypeFormatAndVersion)
	resp, err := p.client.Do(req)
	if err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to perform request: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to get records with code %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to get records with code %d", resp.StatusCode)
	}

	endpoints := []*endpoint.Endpoint{}
	if err := json.NewDecoder(resp.Body).Decode(&endpoints); err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to decode response body: %s", err.Error())
		return nil, err
	}
	return endpoints, nil
}

// ApplyChanges will make a POST to remoteServerURL/records with the changes
func (p WebhookProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	u := p.remoteServerURL.JoinPath("records").String()

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(changes); err != nil {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to encode changes: %s", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", u, b)
	if err != nil {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to create request: %s", err.Error())
		return err
	}

	req.Header.Set(contentTypeHeader, mediaTypeFormatAndVersion)

	resp, err := p.client.Do(req)
	if err != nil {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to perform request: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to apply changes with code %d", resp.StatusCode)
		return fmt.Errorf("failed to apply changes with code %d", resp.StatusCode)
	}
	return nil
}

// PropertyValuesEqual will call the provider doing a POST on `/propertyvaluesequal` which will return a boolean in the format
// `{propertyvaluesequal: true}`
// Errors in anything technically happening from the provider will return true so that no update is performed.
// Errors will also be logged and exposed as metrics so that it is possible to alert on them if needed.
func (p WebhookProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	u := p.remoteServerURL.JoinPath("records").String()

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(&PropertyValuesEqualRequest{
		Name:     name,
		Previous: previous,
		Current:  current,
	}); err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to encode, %s", err)
		return true
	}

	req, err := http.NewRequest("POST", u, b)
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Debugf("Failed to create request: %s", err)
		return true
	}

	req.Header.Set(contentTypeHeader, mediaTypeFormatAndVersion)
	req.Header.Set(acceptHeader, mediaTypeFormatAndVersion)

	resp, err := p.client.Do(req)
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Debugf("Failed to perform request: %s", err)
		return true
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		propertyValuesEqualErrorsGauge.Inc()
		log.Debugf("Failed to run PropertyValuesEqual with code %d", resp.StatusCode)
		return true
	}

	r := PropertyValuesEqualResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to decode response body: %s", err.Error())
		return true
	}

	return r.Equals
}

// AdjustEndpoints will call the provider doing a POST on `/adjustendpoints` which will return a list of modified endpoints
// based on a provider specific requirement.
// This method returns an empty slice in case there is a technical error on the provider's side so that no endpoints will be considered.
func (p WebhookProvider) AdjustEndpoints(e []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}
	u, err := url.JoinPath(p.remoteServerURL.String(), "adjustendpoints")
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to join path, %s", err)
		return endpoints
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(e); err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to encode endpoints, %s", err)
		return endpoints
	}

	req, err := http.NewRequest("POST", u, b)
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to create new HTTP request, %s", err)
		return endpoints
	}

	req.Header.Set(contentTypeHeader, mediaTypeFormatAndVersion)
	req.Header.Set(acceptHeader, mediaTypeFormatAndVersion)

	resp, err := p.client.Do(req)
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed executing http request, %s", err)
		return endpoints
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to AdjustEndpoints with code %d", resp.StatusCode)
		return endpoints
	}

	if err := json.NewDecoder(resp.Body).Decode(&endpoints); err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to decode response body: %s", err.Error())
		return endpoints
	}

	return endpoints
}

// GetDomainFilter is the default implementation of GetDomainFilter.
func (p WebhookProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return endpoint.DomainFilter{}
}
