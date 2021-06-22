// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	defaultBaseURL    = "https://my.rcodezero.at/api/"
	defaultAPIVersion = "v1"
	userAgent         = "rc0go"

	defaultPageSize	  = 100

	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
)

type ClientInterface interface {
	NewRequest() *resty.Request
	ResponseToRC0StatusResponse(response *resty.Response) (*StatusResponse, error)
}

type Client struct {

	// Base URL for API requests. Defaults to the rcode0 dev API, but can be
	// set to a production or test domain. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Version of Rcode0 API
	APIVersion string

	// API Token
	Token string

	// User agent used when communicating with the rcode0 API.
	UserAgent string

	// HTTP client used to communicate with the API.
	client *http.Client

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	Zones  	  ZoneManagementServiceInterface
	RRSet  	  RRSetServiceInterface
	DNSSEC    *DNSSECService
	ZoneStats *ZoneStatsService
	AccStats  *AccountStatsService
	Reports   *ReportService
	Messages  *MessageService
	Settings  *AccSettingsService
}

type service struct {
	client *Client
}

type StatusResponse struct {
	Status  string `json:"status, omitempty"`
	Message string `json:"message, omitempty"`
}

func (sr *StatusResponse) HasError() bool {
	return !strings.EqualFold(sr.Status, "ok")
}

type Page struct {
	Data        []interface{} `json:"data"`
	CurrentPage int           `json:"current_page, omitempty"`
	From        int           `json:"from, omitempty"`
	LastPage    int           `json:"last_page, omitempty"`
	NextPageURL string        `json:"next_page_url, omitempty"`
	Path        string        `json:"path, omitempty"`
	PerPage     int           `json:"per_page, omitempty"`
	PrevPageURL string        `json:"prev_page_url, omitempty"`
	To          int           `json:"to, omitempty"`
	Total       int           `json:"total, omitempty"`
}

type ListOptions struct {
	pageSize 	int
	pageNumber	int
}

func NewListOptions() *ListOptions {
	return &ListOptions{pageNumber: 1, pageSize: defaultPageSize}
}

func (o *ListOptions) SetPageSize(size int) {
	if size < 1 {
		o.pageSize = 1
	}

	o.pageSize = size
}

func (o *ListOptions) PageSize() int {
	if o.pageSize < 1 {
		return 1
	}

	return o.pageSize
}

func (o *ListOptions) PageSizeAsString() string {
	return strconv.Itoa(o.pageSize)
}

func (o *ListOptions) SetPageNumber(number int) {
	if o.pageNumber < 1 {
		o.pageNumber = 1
	}

	o.pageNumber = number
}

func (o *ListOptions) PageNumber() int {
	if o.pageNumber < 1 {
		return 1
	}

	return o.pageNumber
}

func (o *ListOptions) PageNumberAsString() string {
	return strconv.Itoa(o.pageNumber)
}

func (p* Page) IsLastPage() bool {
	return p.CurrentPage == p.LastPage || p.CurrentPage > p.LastPage
}

// NewClient returns a new rcode0 API client.
func NewClient(token string) (*Client, error) {

	if strings.Compare(token, "") == 0 {
		return nil, fmt.Errorf("rcodezero API token is not provided")
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		BaseURL:    baseURL,
		APIVersion: defaultAPIVersion,
		Token:      token,
		UserAgent:  userAgent,
		client:     http.DefaultClient,
	}

	c.common.client = c
	c.Zones   		= (*ZoneManagementService)(&c.common)
	c.RRSet   		= (*RRSetService)(&c.common)
	c.DNSSEC    	= (*DNSSECService)(&c.common)
	c.ZoneStats 	= (*ZoneStatsService)(&c.common)
	c.AccStats 		= (*AccountStatsService)(&c.common)
	c.Settings		= (*AccSettingsService)(&c.common)
	c.Reports 		= (*ReportService)(&c.common)
	c.Messages 		= (*MessageService)(&c.common)

	return c, nil
}

// @todo
func (c *Client) NewRequest() *resty.Request {

	return resty.R().
		SetAuthToken(c.Token)
}

// @todo
func (c *Client) ResponseToRC0StatusResponse(response *resty.Response) (*StatusResponse, error) {
	var statusResponse *StatusResponse

	err := json.Unmarshal(response.Body(), &statusResponse)

	if err != nil {
		return nil, err
	}

	return statusResponse, nil
}