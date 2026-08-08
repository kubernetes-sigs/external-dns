/*
Copyright 2025 The Kubernetes Authors.

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

package pihole

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	extdnshttp "sigs.k8s.io/external-dns/pkg/http"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

const (
	contentTypeJSON = "application/json"
	apiAuthPath     = "/api/auth"
	apiConfigDNS    = "/api/config/dns"
)

// piholeAPI is the interface for interacting with Pi-hole's DNS records.
type piholeAPI interface {
	listRecords(ctx context.Context, rtype string) ([]*endpoint.Endpoint, error)
	createRecord(ctx context.Context, ep *endpoint.Endpoint) error
	deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error
}

// piholeClient implements the piholeAPI.
type piholeClient struct {
	cfg        PiholeConfig
	httpClient *http.Client
	token      string
}

// newPiholeClient creates a new Pi-hole API client.
func newPiholeClient(cfg PiholeConfig) (piholeAPI, error) {
	if cfg.Server == "" {
		return nil, ErrNoPiholeServer
	}

	// Setup an HTTP client
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.TLSInsecureSkipVerify,
			},
		},
	}

	cl := extdnshttp.NewInstrumentedClient(httpClient)

	p := &piholeClient{
		cfg:        cfg,
		httpClient: cl,
	}

	if cfg.Password != "" {
		if err := p.retrieveNewToken(context.Background()); err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p *piholeClient) getConfigValue(ctx context.Context, rtype string) ([]string, error) {
	apiUrl, err := p.urlForRecordType(rtype)
	if err != nil {
		return nil, err
	}

	log.Debugf("Listing %s records from %s", rtype, apiUrl)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	jRes, err := p.do(req)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var apiResponse ApiRecordsResponse
	if err := json.Unmarshal(jRes, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
	}

	var results []string
	if endpoint.RecordTypeCNAME == rtype {
		results = apiResponse.Config.DNS.CnameRecords
	} else {
		results = apiResponse.Config.DNS.Hosts
	}

	return results, nil
}

func (p *piholeClient) listRecords(ctx context.Context, rtype string) ([]*endpoint.Endpoint, error) {
	results, err := p.getConfigValue(ctx, rtype)
	if err != nil {
		return nil, err
	}

	endpoints := make(map[string]*endpoint.Endpoint)

	for _, rec := range results {
		recs := strings.FieldsFunc(rec, func(r rune) bool {
			return r == ' ' || r == ','
		})
		if len(recs) < 2 {
			log.Warnf("skipping record %s: invalid format received from PiHole", rec)
			continue
		}
		dnsName, target := recs[1], recs[0]
		ttl := endpoint.TTL(0)
		switch rtype {
		case endpoint.RecordTypeA:
			if endpoint.SuitableType(target) != endpoint.RecordTypeA {
				continue
			}
		case endpoint.RecordTypeAAAA:
			if endpoint.SuitableType(target) != endpoint.RecordTypeAAAA {
				continue
			}
		case endpoint.RecordTypeCNAME:
			// CNAME format is dnsName,target[,ttl]
			dnsName, target = recs[0], recs[1]
			if len(recs) == 3 {
				if ttlInt, err := strconv.ParseInt(recs[2], 10, 64); err == nil {
					ttl = endpoint.TTL(ttlInt)
				} else {
					log.Warnf("failed to parse TTL value received from PiHole '%s': %v; using a TTL of %d", recs[2], err, ttl)
				}
			}
		}

		ep := endpoint.NewEndpointWithTTL(dnsName, rtype, ttl, target)

		if oldEp, ok := endpoints[dnsName]; ok {
			ep.Targets = append(oldEp.Targets, target) // nolint: gocritic // appendAssign
		}

		endpoints[dnsName] = ep
	}

	out := make([]*endpoint.Endpoint, 0, len(endpoints))
	for _, ep := range endpoints {
		out = append(out, ep)
	}
	return out, nil
}

func (p *piholeClient) createRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return p.apply(ctx, http.MethodPut, ep)
}

func (p *piholeClient) deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return p.apply(ctx, http.MethodDelete, ep)
}

func (p *piholeClient) urlForRecordType(rtype string) (string, error) {
	switch rtype {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA:
		return p.cfg.Server + apiConfigDNS + "/hosts", nil
	case endpoint.RecordTypeCNAME:
		return p.cfg.Server + apiConfigDNS + "/cnameRecords", nil
	default:
		return "", fmt.Errorf("unsupported record type: %s", rtype)
	}
}

// ApiAuthResponse Define a struct to match the JSON response /auth/app structure
type ApiAuthResponse struct {
	Session struct {
		Valid    bool   `json:"valid"`
		TOTP     bool   `json:"totp"`
		SID      string `json:"sid"`
		CSRF     string `json:"csrf"`
		Validity int    `json:"validity"`
		Message  string `json:"message"`
	} `json:"session"`
	Took float64 `json:"took"`
}

// ApiErrorResponse Define struct to match the JSON structure
type ApiErrorResponse struct {
	Error struct {
		Key     string `json:"key"`
		Message string `json:"message"`
		Hint    string `json:"hint"`
	} `json:"error"`
	Took float64 `json:"took"`
}

// ApiRecordsResponse Define struct to match JSON structure
type ApiRecordsResponse struct {
	Config struct {
		DNS struct {
			Hosts        []string `json:"hosts"`
			CnameRecords []string `json:"cnameRecords"`
		} `json:"dns"`
	} `json:"config"`
	Took float64 `json:"took"`
}

func (p *piholeClient) apply(ctx context.Context, action string, ep *endpoint.Endpoint) error {
	if !p.cfg.DomainFilter.Match(ep.DNSName) {
		log.Debugf("Skipping : %s %s that does not match domain filter", action, ep.DNSName)
		return nil
	}
	apiUrl, err := p.urlForRecordType(ep.RecordType)
	if err != nil {
		log.Warnf("Skipping : unsupported endpoint %s %s %v", ep.DNSName, ep.RecordType, ep.Targets)
		return nil
	}

	if len(ep.Targets) == 0 {
		log.Infof("Skipping : missing targets  %s %s %s", action, ep.DNSName, ep.RecordType)
		return nil
	}

	if strings.Contains(ep.DNSName, "*") {
		return provider.NewSoftError(errors.New("UNSUPPORTED: Pihole DNS names cannot return wildcard"))
	}

	if ep.RecordType == endpoint.RecordTypeCNAME && len(ep.Targets) > 1 {
		return provider.NewSoftError(errors.New("UNSUPPORTED: Pihole CNAME records cannot have multiple targets"))
	}

	for _, target := range ep.Targets {
		if p.cfg.DryRun {
			log.Infof("DRY RUN: %s %s IN %s -> %s", action, ep.DNSName, ep.RecordType, target)
			continue
		}

		log.Infof("%s %s IN %s -> %s", action, ep.DNSName, ep.RecordType, target)

		targetApiUrl := apiUrl

		switch ep.RecordType {
		case endpoint.RecordTypeA, endpoint.RecordTypeAAAA:
			targetApiUrl += "/" + url.PathEscape(fmt.Sprintf("%s %s", target, ep.DNSName))
		case endpoint.RecordTypeCNAME:
			if ep.RecordTTL.IsConfigured() {
				targetApiUrl += "/" + url.PathEscape(fmt.Sprintf("%s,%s,%d", ep.DNSName, target, ep.RecordTTL))
			} else {
				targetApiUrl += "/" + url.PathEscape(fmt.Sprintf("%s,%s", ep.DNSName, target))
			}
		}
		req, err := http.NewRequestWithContext(ctx, action, targetApiUrl, nil)
		if err != nil {
			return err
		}

		_, err = p.do(req)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *piholeClient) retrieveNewToken(ctx context.Context) error {
	if p.cfg.Password == "" {
		return nil
	}

	apiUrl := fmt.Sprintf("%s"+apiAuthPath, p.cfg.Server)
	log.Debugf("Fetching new token from %s", apiUrl)

	jsonData, err := json.Marshal(map[string]string{"password": p.cfg.Password})
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	jRes, err := p.do(req)
	if err != nil {
		return err
	}

	// Parse JSON response
	var apiResponse ApiAuthResponse
	if err := json.Unmarshal(jRes, &apiResponse); err != nil {
		return fmt.Errorf("failed to unmarshal auth response: %w", err)
	}
	if apiResponse.Session.SID != "" {
		p.token = apiResponse.Session.SID
	}
	return nil
}

func (p *piholeClient) checkTokenValidity(ctx context.Context) (bool, error) {
	if p.token == "" {
		return false, nil
	}

	apiUrl := fmt.Sprintf("%s"+apiAuthPath, p.cfg.Server)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return false, nil
	}
	req.Header.Add("content-type", contentTypeJSON)
	req.Header.Add("X-FTL-SID", p.token)
	res, err := p.httpClient.Do(req)
	if err != nil {
		return false, err
	}

	jRes, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return false, err
	}

	// Parse JSON response
	var apiResponse ApiAuthResponse
	if err := json.Unmarshal(jRes, &apiResponse); err != nil {
		return false, fmt.Errorf("failed to unmarshal error response: %w", err)
	}
	return apiResponse.Session.Valid, nil
}

func (p *piholeClient) do(req *http.Request) ([]byte, error) {
	req.Header.Add("content-type", contentTypeJSON)
	if p.token != "" {
		req.Header.Add("X-FTL-SID", p.token)
	}
	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	jRes, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusCreated &&
		res.StatusCode != http.StatusNoContent {
		// Parse JSON response
		var apiError ApiErrorResponse
		if err := json.Unmarshal(jRes, &apiError); err != nil {
			return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}
		// Ignore if the entry already exists when adding a record
		if strings.Contains(apiError.Error.Message, "Item already present") {
			return jRes, nil
		}
		// Ignore if the entry does not exist when deleting a record
		if res.StatusCode == http.StatusNotFound && req.Method == http.MethodDelete {
			return jRes, nil
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("Error on request %s", req.URL)
			if req.Body != nil {
				log.Debugf("Body of the request %s", req.Body)
			}
		}

		if res.StatusCode == http.StatusUnauthorized && p.token != "" {
			tryCount := 1
			maxRetries := 3
			// Try to fetch a new token and redo the request.
			for tryCount <= maxRetries {
				valid, err := p.checkTokenValidity(req.Context())
				if err != nil {
					return nil, err
				}
				if !valid {
					log.Debugf("Pihole token has expired, fetching a new one. Try (%d/%d)", tryCount, maxRetries)
					if err := p.retrieveNewToken(req.Context()); err != nil {
						return nil, err
					}
					tryCount++
					continue
				}
				break
			}
			if tryCount > maxRetries {
				return nil, errors.New("max tries reached for token renewal")
			}
			return p.do(req)
		}
		return nil, fmt.Errorf("received %d status code from request: [%s] %s (%s) - %fs", res.StatusCode, apiError.Error.Key, apiError.Error.Message, apiError.Error.Hint, apiError.Took)
	}
	return jRes, nil
}
