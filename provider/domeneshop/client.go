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

package domeneshop

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

const apiURL string = "https://api.domeneshop.no/v0"
const defaultTTL int = 300

type DomeneshopClient interface {
	ListDomains(ctx context.Context) ([]*Domain, error)
	ListDNSRecords(ctx context.Context, domain *Domain, host, recordType string) ([]*DNSRecord, error)
	AddDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error
	UpdateDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error
	DeleteDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error
}

type domeneshopClient struct {
	apiToken   string
	apiSecret  string
	appVersion string
	http       http.Client
}

var _ DomeneshopClient = &domeneshopClient{}

// NewClient returns an instance of the Domeneshop API wrapper
func NewClient(apiToken, apiSecret string, appVersion string) DomeneshopClient {
	client := domeneshopClient{
		apiToken:   apiToken,
		apiSecret:  apiSecret,
		appVersion: appVersion,
		http:       http.Client{},
	}

	return &client
}

// Request makes a request against the API with an optional body, and makes sure
// that the required Authorization header is set using `setBasicAuth`
func (c *domeneshopClient) request(ctx context.Context, method string, endpoint string, reqBody []byte, v interface{}) error {
	var buf = bytes.NewBuffer(reqBody)

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", apiURL, endpoint), buf)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.apiToken, c.apiSecret)

	req.Header.Set("User-Agent", fmt.Sprintf("ExternalDNS/%s", c.appVersion))

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 399 {
		var httpErrorBody HttpErrorBody
		if err := json.Unmarshal(respBody, &httpErrorBody); err != nil {
			log.WithFields(log.Fields{
				"status_code": resp.StatusCode,
				"body":        string(respBody),
			}).Warn("Error parsing body of error response.")
		}
		return &HttpError{
			Message:   fmt.Errorf("API returned %s: %s", resp.Status, respBody),
			Response:  resp,
			ErrorBody: httpErrorBody,
		}
	}

	if v != nil {
		return json.Unmarshal(respBody, &v)
	}
	return nil
}

func (c *domeneshopClient) ListDomains(ctx context.Context) ([]*Domain, error) {
	var domains []*Domain
	err := c.request(ctx, "GET", "domains", nil, &domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

func (c *domeneshopClient) ListDNSRecords(ctx context.Context, domain *Domain, host, recordType string) ([]*DNSRecord, error) {
	var values = url.Values{}
	if host != "" {
		values.Add("host", host)
	}
	if recordType != "" {
		values.Add("type", recordType)
	}

	var records []*DNSRecord
	err := c.request(ctx, "GET", fmt.Sprintf("domains/%d/dns?%s", domain.ID, values.Encode()), nil, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (c *domeneshopClient) AddDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error {
	jsonRecord, err := json.Marshal(record)

	if err != nil {
		return err
	}

	return c.request(ctx, "POST", fmt.Sprintf("domains/%d/dns", domain.ID), jsonRecord, nil)
}

func (c *domeneshopClient) UpdateDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error {
	jsonRecord, err := json.Marshal(record)

	if err != nil {
		return err
	}

	return c.request(ctx, "PUT", fmt.Sprintf("domains/%d/dns/%d", domain.ID, record.ID), jsonRecord, nil)
}

func (c *domeneshopClient) DeleteDNSRecord(ctx context.Context, domain *Domain, record *DNSRecord) error {
	return c.request(ctx, "DELETE", fmt.Sprintf("domains/%d/dns/%d", domain.ID, record.ID), nil, nil)
}
