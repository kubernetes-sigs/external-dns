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

package adguardhome

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

// adguardAPI declares the "API" actions performed against the AdGuardHome server.
type adGuardHomeAPI interface {
	// listRecords returns endpoints currently configured in AdGuard
	listRecords(ctx context.Context) ([]*endpoint.Endpoint, error)
	// createRecord will create a new record for the given endpoint.
	createRecord(ctx context.Context, ep *endpoint.Endpoint) error
	// deleteRecord will delete the given record.
	deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error
	//updateRecord will update the given record.
	updateRecord(ctx context.Context, oldEP, newEP *endpoint.Endpoint) error
}

// AdGuardHomeClient implements the AdGuardHomeAPI.
type adGuardHomeClient struct {
	BaseURL    *url.URL
	cfg        AdGuardHomeConfig
	httpClient *http.Client
}

type adGuardHomeEntry struct {
	Domain string `json:"domain"`
	Answer string `json:"answer"`
}

type adGuardHomeUpdateEntry struct {
	Target adGuardHomeEntry `json:"target"`
	Update adGuardHomeEntry `json:"update"`
}

// newAdGuardHomeClient creates a new AdGuardHome API client.
func newAdGuardHomeClient(cfg AdGuardHomeConfig) (adGuardHomeAPI, error) {
	if cfg.Username == "" {
		err := errors.New("no username supplied, this is required")
		return nil, err
	}

	if cfg.Password == "" {
		err := errors.New("no password supplied, this is required")
		return nil, err
	}

	parsedURL, err := url.Parse(cfg.Server)
	if err != nil {
		return nil, err
	}

	httpClient := instrumented_http.NewClient(
		&http.Client{},
		&instrumented_http.Callbacks{})

	client := &adGuardHomeClient{
		BaseURL:    parsedURL,
		cfg:        cfg,
		httpClient: httpClient,
	}

	return client, nil
}

func (c *adGuardHomeClient) do(req *http.Request) (io.ReadCloser, error) {
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(c.cfg.Username, c.cfg.Password)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("received non-200 status code from %s request: %s", res.Request.URL, res.Status)
	}
	return res.Body, nil
}

func (c *adGuardHomeClient) getType(entry *adGuardHomeEntry) string {
	ip := net.ParseIP(entry.Answer)
	if ip == nil {
		return endpoint.RecordTypeCNAME
	}
	ip4 := ip.To4()
	if ip4 != nil {
		return endpoint.RecordTypeA
	}

	return endpoint.RecordTypeAAAA
}

func (c *adGuardHomeClient) listRecords(ctx context.Context) ([]*endpoint.Endpoint, error) {
	url := c.BaseURL.ResolveReference(&url.URL{Path: "/control/rewrite/list"})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	raw, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var entries []adGuardHomeEntry
	if err := json.Unmarshal(raw, &entries); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	out := make([]*endpoint.Endpoint, 0)
	for _, rec := range entries {
		if c.cfg.DomainFilter.IsConfigured() && !c.cfg.DomainFilter.Match(rec.Domain) {
			log.Debugf("Skipping %s that does not match domain filter", rec.Domain)
			continue
		}
		out = append(out, &endpoint.Endpoint{
			DNSName:    rec.Domain,
			Targets:    []string{rec.Answer},
			RecordType: c.getType(&rec),
		})
	}
	return out, nil
}

func (c *adGuardHomeClient) handleRecord(ctx context.Context, ep *endpoint.Endpoint, action string) error {
	if c.cfg.DryRun {
		return nil
	}

	record := adGuardHomeEntry{
		Domain: ep.DNSName,
		Answer: ep.Targets[0],
	}

	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	url := c.BaseURL.ResolveReference(&url.URL{Path: "/control/rewrite/" + action})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if _, err := c.do(req); err != nil {
		return err
	}

	return nil
}

func (c *adGuardHomeClient) createRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	if ep.RecordType != "A" && ep.RecordType != "AAAA" && ep.RecordType != "CNAME" {
		return fmt.Errorf("unsupported record type: %s for %s", ep.RecordType, ep.DNSName)
	}
	return c.handleRecord(ctx, ep, "add")
}

func (c *adGuardHomeClient) deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return c.handleRecord(ctx, ep, "delete")
}

func (c *adGuardHomeClient) updateRecord(ctx context.Context, oldEp, newEp *endpoint.Endpoint) error {
	if c.cfg.DryRun {
		return nil
	}

	record := adGuardHomeUpdateEntry{
		Target: adGuardHomeEntry{
			Domain: oldEp.DNSName,
			Answer: oldEp.Targets[0],
		},
		Update: adGuardHomeEntry{
			Domain: newEp.DNSName,
			Answer: newEp.Targets[0],
		},
	}
	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	url := c.BaseURL.ResolveReference(&url.URL{Path: "/control/rewrite/update"})
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if _, err := c.do(req); err != nil {
		return err
	}

	return nil
}
