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

package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	mediaTypeFormatAndVersion = "application/external.dns.plugin+json;version=1"
	contentTypeHeader         = "Content-Type"
	acceptHeader              = "Accept"
	varyHeader                = "Vary"
	maxRetries                = 5
)

var (
	recordsErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "plugin_provider",
			Name:      "records_errors",
			Help:      "Errors with Records method",
		},
	)
	applyChangesErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "plugin_provider",
			Name:      "applychanges_errors",
			Help:      "Errors with ApplyChanges method",
		},
	)
	propertyValuesEqualErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "plugin_provider",
			Name:      "propertyvaluesequal_errors",
			Help:      "Errors with PropertyValuesEqual method",
		},
	)
	adjustEndpointsErrorsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "external_dns",
			Subsystem: "plugin_provider",
			Name:      "adjustendpointsgauge_errors",
			Help:      "Errors with AdjustEndpoints method",
		},
	)
)

type PluginProvider struct {
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

func NewPluginProvider(u string) (*PluginProvider, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	// negotiate API information
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(acceptHeader, mediaTypeFormatAndVersion)

	client := &http.Client{}
	var resp *http.Response
	err = backoff.Retry(func() error {
		resp, err = client.Do(req)
		if err != nil {
			log.Debugf("Failed to connect to plugin api: %v", err)
			return err
		}
		// we currently only use 200 as success, but considering okay all 2XX for future usage
		if resp.StatusCode >= 300 && resp.StatusCode < 500 {
			return backoff.Permanent(fmt.Errorf("status code < 500"))
		}
		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetries))

	if err != nil {
		return nil, fmt.Errorf("failed to connect to plugin api: %v", err)
	}

	vary := resp.Header.Get(varyHeader)
	contentType := resp.Header.Get(contentTypeHeader)

	if vary != contentTypeHeader {
		return nil, fmt.Errorf("wrong vary value returned from server: %s", vary)
	}

	if contentType != mediaTypeFormatAndVersion {
		return nil, fmt.Errorf("wrong content type returned from server: %s", contentType)
	}

	return &PluginProvider{
		client:          client,
		remoteServerURL: parsedURL,
	}, nil
}

// Records will make a GET call to remoteServerURL/records and return the results
func (p PluginProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	u, err := url.JoinPath(p.remoteServerURL.String(), "records")
	if err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to join path: %s", err.Error())
		return nil, err
	}
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to read response body: %s", err.Error())
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	err = json.Unmarshal(b, &endpoints)
	if err != nil {
		recordsErrorsGauge.Inc()
		log.Debugf("Failed to unmarshal response body: %s", err.Error())
		return nil, err
	}
	return endpoints, nil
}

// ApplyChanges will make a POST to remoteServerURL/records with the changes
func (p PluginProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	u, err := url.JoinPath(p.remoteServerURL.String(), "records")
	if err != nil {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to join path: %s", err.Error())
		return err
	}
	b, err := json.Marshal(changes)
	if err != nil {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to marshal changes: %s", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(b))
	if err != nil {
		applyChangesErrorsGauge.Inc()
		log.Debugf("Failed to create request: %s", err.Error())
		return err
	}
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
func (p PluginProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	u, err := url.JoinPath(p.remoteServerURL.String(), "propertyvaluesequal")
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Debugf("Failed to join path: %s", err)
		return true
	}
	b, err := json.Marshal(&PropertyValuesEqualRequest{
		Name:     name,
		Previous: previous,
		Current:  current,
	})
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Debugf("Failed to marshal request: %s", err)
		return true
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(b))
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Debugf("Failed to create request: %s", err)
		return true
	}
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

	respoBody, err := io.ReadAll(resp.Body)
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Errorf("Failed to read body: %v", err)
		return true
	}
	r := PropertyValuesEqualResponse{}
	err = json.Unmarshal(respoBody, &r)
	if err != nil {
		propertyValuesEqualErrorsGauge.Inc()
		log.Errorf("Failed to unmarshal body: %v", err)
		return true
	}
	return r.Equals
}

// AdjustEndpoints will call the provider doing a POST on `/adjustendpoints` which will return a list of modified endpoints
// based on a provider specific requirement.
// This method returns an empty slice in case there is a technical error on the provider's side so that no endpoints will be considered.
func (p PluginProvider) AdjustEndpoints(e []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}
	u, err := url.JoinPath(p.remoteServerURL.String(), "adjustendpoints")
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to join path, %s", err)
		return endpoints
	}
	b, err := json.Marshal(e)
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to marshal endpoints, %s", err)
		return endpoints
	}
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(b))
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to create new HTTP request, %s", err)
		return endpoints
	}
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

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Failed to read response body, %s", err)
		return endpoints
	}

	err = json.Unmarshal(b, &endpoints)
	if err != nil {
		adjustEndpointsErrorsGauge.Inc()
		log.Debugf("Faile to unmarshal response body, %s", err)
		return endpoints
	}
	return endpoints
}

// GetDomainFilter is the default implementation of GetDomainFilter.
func (p PluginProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return endpoint.DomainFilter{}
}
