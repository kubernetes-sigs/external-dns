/*
Copyright 2017 The Kubernetes Authors.

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
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
)

// api declares the "API" actions performed against the AdGuard Home server.
type api interface {
	// listRecords returns endpoints for the given record type (A or CNAME).
	listRecords(ctx context.Context) ([]*endpoint.Endpoint, error)
	// createRecord will create a new record for the given endpoint.
	createRecord(ctx context.Context, ep *endpoint.Endpoint) error
	// deleteRecord will delete the given record.
	deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error
}

// client implements the api
type client struct {
	cfg        AdGuardHomeConfig
	token      string
	httpClient *http.Client
	url        *url.URL
}

// newClient creates a new AdGuard Home API client.
func newClient(cfg AdGuardHomeConfig) (api, error) {
	var err error

	if cfg.Server == "" {
		return nil, ErrNoServer
	}

	c := &client{cfg: cfg}

	if c.url, err = url.Parse(cfg.Server); err != nil {
		return nil, err
	}
	c.url = c.url.JoinPath("control")

	if cfg.Username != "" && cfg.Password != "" {
		c.token = base64.StdEncoding.EncodeToString([]byte(cfg.Username + ":" + cfg.Password))
	}

	c.httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: cfg.TLSInsecureSkipVerify,
			},
		},
	}

	if err := c.checkStatus(context.Background()); err != nil {
		return nil, err
	}

	return c, nil
}

func (c client) checkStatus(ctx context.Context) error {
	url := c.url.JoinPath("status").String()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	_ = body.Close()
	return nil
}

func (c client) listRecords(ctx context.Context) ([]*endpoint.Endpoint, error) {
	url := c.url.JoinPath("rewrite", "list").String()

	log.Debugf("Listing records from %s", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var entries []rewriteModel
	if err := json.NewDecoder(body).Decode(&entries); err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(entries))
	for _, entry := range entries {
		recordType := endpoint.RecordTypeCNAME
		if addr := net.ParseIP(entry.Answer); addr != nil {
			recordType = endpoint.RecordTypeA
		}

		if c.cfg.DomainFilter.IsConfigured() && !c.cfg.DomainFilter.Match(entry.Domain) {
			log.Debugf("Skipping %s that does not match domain filter", entry.Domain)
			continue
		}

		endpoints = append(endpoints, &endpoint.Endpoint{
			DNSName:    entry.Domain,
			Targets:    endpoint.Targets{entry.Answer},
			RecordType: recordType,
		})
	}

	return endpoints, nil
}

func (c client) createRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return c.apply(ctx, "add", ep)
}

func (c client) deleteRecord(ctx context.Context, ep *endpoint.Endpoint) error {
	return c.apply(ctx, "delete", ep)
}

func (c client) apply(ctx context.Context, action string, ep *endpoint.Endpoint) error {
	switch action {
	case "add", "delete":
		//
	default:
		return fmt.Errorf("invalid apply action: %s", action)
	}

	if c.cfg.DomainFilter.IsConfigured() && !c.cfg.DomainFilter.Match(ep.DNSName) {
		log.Debugf("Skipping %s %s that does not match domain filter", action, ep.DNSName)
		return nil
	}

	switch ep.RecordType {
	case endpoint.RecordTypeA, endpoint.RecordTypeCNAME:
		//
	default:
		return fmt.Errorf("unsupported record type: %s", ep.RecordType)
	}

	if c.cfg.DryRun {
		log.Infof("DRY RUN: %s %s IN %s -> %s", action, ep.DNSName, ep.RecordType, ep.Targets[0])
		return nil
	}

	log.Infof("%s %s IN %s -> %s", action, ep.DNSName, ep.RecordType, ep.Targets[0])

	reqBody, err := json.Marshal(rewriteModel{
		Domain: ep.DNSName,
		Answer: ep.Targets[0],
	})
	if err != nil {
		return err
	}

	url := c.url.JoinPath("rewrite", action).String()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	respBody, err := c.doRequest(req)
	if err != nil {
		return err
	}
	defer respBody.Close()

	return nil
}

func (c client) doRequest(req *http.Request) (io.ReadCloser, error) {
	if c.token != "" {
		req.Header.Set("Authorization", "Basic "+c.token)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		return nil, fmt.Errorf("%w: %s", ErrInvalidResponse, res.Status)
	}
	return res.Body, nil
}
