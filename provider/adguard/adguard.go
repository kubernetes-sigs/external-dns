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

package adguard

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	ListRewrites  = "/control/rewrite/list"
	DeleteRewrite = "/control/rewrite/delete"
	CreateRwrite  = "/control/rewrite/add"
	UpdateRewrite = "/control/rewrite/update"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

// RewriteEntry represents the RewriteEntry in Adguard's API
type RewriteEntry struct {
	// Domain is the domain Adguard will rewrite
	Domain string `json:"domain"`
	// Answer is the A, AAA, or CNAME Adguard will rewrite to
	Answer string `json:"answer"`
}

type UpdateRewriteEntry struct {
	Target RewriteEntry `json:"target"`
	Update RewriteEntry `json:"update"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func defaultHttpClient() httpClient {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("redirect attmempted")
		},
	}
}

// Provider is the provider implementation for AdGuard Home
//
// AdGuard Home API documentation is located at https://github.com/AdguardTeam/AdGuardHome/tree/master/openapi
type Provider struct {
	provider.BaseProvider
	config Config

	client httpClient
}

// Explicitly ensure interface is implemented
var _ provider.Provider = &Provider{}

// NewProvider creates a provder for AdGuard Home
func NewProvider(config Config) (*Provider, error) {
	if err := config.validate(); err != nil {
		return nil, errors.Join(ErrInvalidConfig, err)
	}
	return &Provider{
		config: config,
		client: defaultHttpClient(),
	}, nil
}

func (ap Provider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	resp, err := ap.makeRequest(ctx, http.MethodGet, ListRewrites, nil)
	if err != nil {
		return nil, fmt.Errorf("adguard records: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	var records []RewriteEntry
	if err := json.Unmarshal(body, &records); err != nil {
		return nil, fmt.Errorf("unmarshal body: %w", err)
	}

	return recordsToEndpoints(records), nil
}

func (ap Provider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if changes == nil || !changes.HasChanges() {
		return nil
	}

	if err := ap.createEndpoints(ctx, changes.Create); err != nil {
		return fmt.Errorf("adguard create: %w", err)
	}

	if err := ap.deleteEndpoints(ctx, changes.Delete); err != nil {
		return fmt.Errorf("adguard delete: %w", err)
	}

	if err := ap.updateEndpoints(ctx, changes.UpdateOld, changes.UpdateNew); err != nil {
		return fmt.Errorf("adguard update: %w", err)
	}

	return nil
}

func (ap Provider) createEndpoints(ctx context.Context, endpoints []*endpoint.Endpoint) error {
	for _, endpoint := range endpoints {
		if err := ap.createEndpoint(ctx, endpoint); err != nil {
			return fmt.Errorf("create %s: %w", endpoint.DNSName, err)
		}
	}

	return nil
}

func (ap Provider) createEndpoint(ctx context.Context, endpoint *endpoint.Endpoint) error {
	resp, err := ap.postRequest(ctx, endpoint, CreateRwrite)
	if err != nil {
		return err
	}

	return validateResponse(resp)
}

func (ap Provider) deleteEndpoints(ctx context.Context, endpoints []*endpoint.Endpoint) error {
	for _, endpoint := range endpoints {
		if err := ap.deleteEndpoint(ctx, endpoint); err != nil {
			return fmt.Errorf("delete %s: %w", endpoint.DNSName, err)
		}
	}
	return nil
}

func (ap Provider) deleteEndpoint(ctx context.Context, endpoint *endpoint.Endpoint) error {
	resp, err := ap.postRequest(ctx, endpoint, DeleteRewrite)
	if err != nil {
		return err
	}

	return validateResponse(resp)
}

func (ap Provider) updateEndpoints(ctx context.Context, oldEndpoints, newEndpoints []*endpoint.Endpoint) error {
	oldMapping := map[string]*endpoint.Endpoint{}
	for _, old := range oldEndpoints {
		oldMapping[old.DNSName] = old
	}

	for _, endpoint := range newEndpoints {
		if err := ap.updateEndpoint(ctx, oldMapping[endpoint.DNSName], endpoint); err != nil {
			return fmt.Errorf("update %s: %w", endpoint.DNSName, err)
		}
	}

	return nil
}

func (ap Provider) updateEndpoint(ctx context.Context, old, new *endpoint.Endpoint) error {
	updateData := UpdateRewriteEntry{
		Target: RewriteEntry{
			Domain: old.DNSName,
			Answer: old.Targets.String(),
		},
		Update: RewriteEntry{
			Domain: new.DNSName,
			Answer: new.Targets.String(),
		},
	}
	json, err := json.Marshal(updateData)
	if err != nil {
		return err
	}

	resp, err := ap.makeRequest(ctx, http.MethodPut, UpdateRewrite, bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	return validateResponse(resp)
}

func (ap Provider) postRequest(ctx context.Context, endpoint *endpoint.Endpoint, reqPath string) (*http.Response, error) {
	entryBytes, err := getRewriteEntryJson(endpoint)
	if err != nil {
		return nil, err
	}

	resp, err := ap.makeRequest(ctx, http.MethodPost, reqPath, bytes.NewBuffer(entryBytes))
	if err != nil {
		return nil, fmt.Errorf("make request: %w", err)
	}

	return resp, nil
}

func (ap Provider) makeRequest(ctx context.Context, method string, reqPath string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, ap.config.Server+reqPath, data)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.SetBasicAuth(ap.config.Username, ap.config.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := ap.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return resp, err
}

func validateResponse(resp *http.Response) error {
	if resp != nil {
		defer func() { resp.Body.Close() }()
		io.ReadAll(resp.Body)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response: %v", resp.StatusCode)
	}

	return nil
}

func getRewriteEntryJson(endpoint *endpoint.Endpoint) ([]byte, error) {
	entry := RewriteEntry{
		Domain: endpoint.DNSName,
		Answer: endpoint.Targets.String(),
	}

	entryBytes, err := json.Marshal(entry)
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}
	return entryBytes, nil
}

func recordsToEndpoints(records []RewriteEntry) []*endpoint.Endpoint {
	endpoints := make([]*endpoint.Endpoint, len(records))
	for i, record := range records {
		if isAnswerCNAME(record.Answer) {
			endpoints[i] = endpoint.NewEndpoint(record.Domain, endpoint.RecordTypeCNAME, record.Answer)
		} else {
			endpoints[i] = endpoint.NewEndpoint(record.Domain, endpoint.RecordTypeA, record.Answer)
		}
	}

	return endpoints
}

func isAnswerCNAME(answer string) bool {
	ip := net.ParseIP(answer)
	return ip == nil
}
