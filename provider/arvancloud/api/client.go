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

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider/arvancloud/dto"
)

const (
	_defaultTimeout           = 180 * time.Second
	_defaultPrefixApiEndpoint = "https://napi.arvancloud.ir/cdn/%s"
)

type clientApi struct {
	client      *http.Client
	apiToken    string
	apiEndpoint string
	timeout     time.Duration
}

type OutResp[T interface{}] struct {
	Data  T        `json:"data"`
	Links LinkResp `json:"links"`
	Meta  MetaResp `json:"meta"`
}

type LinkResp struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type MetaResp struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

type pageIn struct {
	current int
	limit   int
}

type ClientError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (err *ClientError) Error() string {
	msg := "error is happened"
	if err.Message != "" {
		msg = err.Message
	}

	return fmt.Sprintf("%s (status: %s)", msg, err.Code)
}

func NewClientApi(apiToken, apiVersion string) (*clientApi, error) {
	return &clientApi{
		client:      &http.Client{},
		apiToken:    apiToken,
		apiEndpoint: fmt.Sprintf(_defaultPrefixApiEndpoint, apiVersion),
		timeout:     _defaultTimeout,
	}, nil
}

// GetDomains get a list of domains to exist in the ArvanCloud
func (c *clientApi) GetDomains(ctx context.Context, perPage ...int) ([]dto.Zone, error) {
	page := getPage(perPage)
	domainUrl := setPageQueryString(url.URL{Path: "/domains"}, page)

	resp := OutResp[[]dto.Zone]{}
	if err := c.makeGetWithPagination(ctx, domainUrl, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, nil
	}

	zl := make([]dto.Zone, 0, len(resp.Data))
	zl = append(zl, resp.Data...)

	for {
		if resp.Meta.CurrentPage == resp.Meta.LastPage {
			break
		}

		page.current = resp.Meta.CurrentPage + 1
		domainUrl = setPageQueryString(domainUrl, page)

		resp = OutResp[[]dto.Zone]{}
		if err := c.makeGetWithPagination(ctx, domainUrl, &resp); err != nil {
			return nil, err
		}

		if len(resp.Data) == 0 {
			break
		}

		zl = append(zl, resp.Data...)
	}

	return zl, nil
}

// GetDnsRecords get a list of dns records to exist on zone
func (c *clientApi) GetDnsRecords(ctx context.Context, zone string, perPage ...int) ([]dto.DnsRecord, error) {
	page := getPage(perPage)
	dnsUrl := setPageQueryString(url.URL{Path: fmt.Sprintf("/domains/%s/dns-records", zone)}, page)

	resp := OutResp[[]dto.DnsRecord]{}
	if err := c.makeGetWithPagination(ctx, dnsUrl, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, nil
	}

	dl := make([]dto.DnsRecord, 0, len(resp.Data))
	dl = append(dl, resp.Data...)

	for {
		if resp.Meta.CurrentPage == resp.Meta.LastPage {
			break
		}

		page.current = resp.Meta.CurrentPage + 1
		dnsUrl = setPageQueryString(dnsUrl, page)

		resp = OutResp[[]dto.DnsRecord]{}
		if err := c.makeGetWithPagination(ctx, dnsUrl, &resp); err != nil {
			return nil, err
		}

		if len(resp.Data) == 0 {
			break
		}

		dl = append(dl, resp.Data...)
	}

	for i := 0; i < len(dl); i++ {
		dl[i].Zone = zone
	}

	return dl, nil
}

// CreateDnsRecord create new dns record for domain
func (c *clientApi) CreateDnsRecord(ctx context.Context, zone string, record dto.DnsRecord) (out dto.DnsRecord, err error) {
	dnsUrl := url.URL{Path: fmt.Sprintf("/domains/%s/dns-records", zone)}

	req, err := c.newRequest("POST", dnsUrl.String(), record)
	if err != nil {
		return out, err
	}
	req = req.WithContext(ctx)
	response, err := c.do(req)
	if err != nil {
		return out, err
	}

	resp := OutResp[dto.DnsRecord]{}
	if err = c.unmarshalResponse(response, &resp); err != nil {
		return out, err
	}

	return resp.Data, nil
}

// UpdateDnsRecord update exist dns record for domain
func (c *clientApi) UpdateDnsRecord(ctx context.Context, zone string, record dto.DnsRecord) (out dto.DnsRecord, err error) {
	dnsUrl := url.URL{Path: fmt.Sprintf("/domains/%s/dns-records/%s", zone, record.ID)}

	req, err := c.newRequest("PUT", dnsUrl.String(), record)
	if err != nil {
		return out, err
	}
	req = req.WithContext(ctx)
	response, err := c.do(req)
	if err != nil {
		return out, err
	}

	resp := OutResp[dto.DnsRecord]{}
	if err = c.unmarshalResponse(response, &resp); err != nil {
		return out, err
	}

	return resp.Data, nil
}

// DeleteDnsRecord delete exist dns record from domain
func (c *clientApi) DeleteDnsRecord(ctx context.Context, zone, recordId string) error {
	dnsUrl := url.URL{Path: fmt.Sprintf("/domains/%s/dns-records/%s", zone, recordId)}

	req, err := c.newRequest("DELETE", dnsUrl.String(), nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	response, err := c.do(req)
	if err != nil {
		return err
	}

	resp := OutResp[dto.DnsRecord]{}
	if err = c.unmarshalResponse(response, &resp); err != nil {
		return err
	}

	return nil
}

// makeGetWithPagination make new GET request with pagination
func (c *clientApi) makeGetWithPagination(ctx context.Context, path url.URL, resType interface{}) error {
	req, err := c.newRequest("GET", path.String(), nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	response, err := c.do(req)
	if err != nil {
		return err
	}

	if err := c.unmarshalResponse(response, resType); err != nil {
		return err
	}

	return nil
}

// newRequest returns a new HTTP request
func (c *clientApi) newRequest(method, path string, reqBody interface{}) (*http.Request, error) {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
	}

	target := fmt.Sprintf("%s%s", c.apiEndpoint, path)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Set("Authorization", c.apiToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ExternalDNS/"+externaldns.Version)

	c.client.Timeout = c.timeout

	return req, nil
}

// do send an HTTP request and returns an HTTP response
func (c *clientApi) do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusFound {
		return nil, dto.NewUnknownError(nil)
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, dto.NewUnauthorizedError()
	}

	return resp, nil
}

// unmarshalResponse checks the response and unmarshal it into the response
// type if needed Helper function, called from CallAPI
func (c *clientApi) unmarshalResponse(response *http.Response, resType interface{}) error {
	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Errorf("failed to close response: %v", err)
		}
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		log.Tracef("error on make request: %s", string(body))
		apiError := &ClientError{
			Code: strconv.Itoa(response.StatusCode),
		}

		if err = json.Unmarshal(body, apiError); err != nil {
			return err
		}

		return dto.NewUnknownError(apiError)
	}

	if len(body) == 0 || resType == nil {
		return nil
	}

	return json.Unmarshal(body, &resType)
}

func getPage(perPage []int) pageIn {
	var limit int
	if len(perPage) > 0 && perPage[0] > 0 {
		limit = perPage[0]
	}

	return pageIn{current: 1, limit: limit}
}

func setPageQueryString(urlPath url.URL, page pageIn) url.URL {
	q := urlPath.Query()
	if page.current > 1 {
		q.Set("page", strconv.Itoa(page.current))
	}
	if page.limit != 0 {
		q.Set("per_page", strconv.Itoa(page.limit))
	}
	urlPath.RawQuery = q.Encode()

	return urlPath
}
