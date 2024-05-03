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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/provider/arvancloud/dto"
)

func TestArvanCloud_ClientApi_GetDomains(t *testing.T) {
	type given struct {
		ctx     context.Context
		perPage []int
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		init      func(t *testing.T, wants []expectedInit) *clientApi
		wantsErr  expectedErr
		name      string
		given     given
		wants     []dto.Zone
		wantsInit []expectedInit
	}{
		{
			name: "should error get domains when make request",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)
					cntr++

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				ctx: context.Background(),
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should successfully get domains and return empty data",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)
					cntr++

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(`{"data":[]}`)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx: context.Background(),
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
				},
			},
		},
		{
			name: "should successfully get domains data (without pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)
					cntr++

					zone := []dto.Zone{
						{
							ID:          "1",
							Name:        "example1.com",
							Type:        "full",
							Status:      "active",
							OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
							NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
							Plan:        1,
						},
						{
							ID:          "2",
							Name:        "example2.com",
							Type:        "full",
							Status:      "active",
							OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
							NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
							Plan:        1,
						},
					}
					zoneBody, err := json.Marshal(zone)
					assert.NoError(t, err)
					respBody := fmt.Sprintf("{%q:%s}", "data", string(zoneBody))

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx: context.Background(),
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
				},
			},
			wants: []dto.Zone{
				{
					ID:          "1",
					Name:        "example1.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
				{
					ID:          "2",
					Name:        "example2.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
			},
		},
		{
			name: "should error get domains when make request in second try (with pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						zone := []dto.Zone{
							{
								ID:          "1",
								Name:        "example1.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
							{
								ID:          "2",
								Name:        "example2.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 2, "current_page", 1, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						return nil, errors.New("fail to fetch url")
					}
				})

				return c
			},
			given: given{
				ctx: context.Background(),
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?page=2",
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should successfully get domains data and in return empty records in second try (with pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						zone := []dto.Zone{
							{
								ID:          "1",
								Name:        "example1.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
							{
								ID:          "2",
								Name:        "example2.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 2, "current_page", 1, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						respBody := fmt.Sprintf("{%q:[]}", "data")

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					}
				})

				return c
			},
			given: given{
				ctx: context.Background(),
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?page=2",
				},
			},
			wants: []dto.Zone{
				{
					ID:          "1",
					Name:        "example1.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
				{
					ID:          "2",
					Name:        "example2.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
			},
		},
		{
			name: "should successfully get dns records data (with pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						zone := []dto.Zone{
							{
								ID:          "1",
								Name:        "example1.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
							{
								ID:          "2",
								Name:        "example2.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 2, "current_page", 1, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						zone := []dto.Zone{
							{
								ID:          "3",
								Name:        "example3.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 2, "current_page", 2, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					}
				})

				return c
			},
			given: given{
				ctx: context.Background(),
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?page=2",
				},
			},
			wants: []dto.Zone{
				{
					ID:          "1",
					Name:        "example1.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
				{
					ID:          "2",
					Name:        "example2.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
				{
					ID:          "3",
					Name:        "example3.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
			},
		},
		{
			name: "should successfully get dns records data (with pagination and limit item)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						zone := []dto.Zone{
							{
								ID:          "1",
								Name:        "example1.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 3, "current_page", 1, "last_page", 3)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else if cntr == 2 {
						zone := []dto.Zone{
							{
								ID:          "2",
								Name:        "example2.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 3, "current_page", 2, "last_page", 3)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						zone := []dto.Zone{
							{
								ID:          "3",
								Name:        "example3.com",
								Type:        "full",
								Status:      "active",
								OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
								NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
								Plan:        1,
							},
						}
						zoneBody, err := json.Marshal(zone)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(zoneBody), "meta", "total", 3, "current_page", 3, "last_page", 3)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					}
				})

				return c
			},
			given: given{
				ctx:     context.Background(),
				perPage: []int{1},
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?per_page=1",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?page=2&per_page=1",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?page=3&per_page=1",
				},
			},
			wants: []dto.Zone{
				{
					ID:          "1",
					Name:        "example1.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
				{
					ID:          "2",
					Name:        "example2.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
				{
					ID:          "3",
					Name:        "example3.com",
					Type:        "full",
					Status:      "active",
					OriginalNS:  []string{"ns1.cloud.com", "ns2.cloud.com"},
					NameServers: []string{"ns1.cloud.com", "ns2.cloud.com"},
					Plan:        1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.wantsInit)

			out, err := c.GetDomains(tt.given.ctx, tt.given.perPage...)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_ClientApi_GetDnsRecords(t *testing.T) {
	type given struct {
		ctx     context.Context
		zone    string
		perPage []int
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		init      func(t *testing.T, wants []expectedInit) *clientApi
		wantsErr  expectedErr
		name      string
		given     given
		wants     []dto.DnsRecord
		wantsInit []expectedInit
	}{
		{
			name: "should error get dns records when make request",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)
					cntr++

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should successfully get dns records and return empty data",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)
					cntr++

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(`{"data":[]}`)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
				},
			},
		},
		{
			name: "should successfully get dns records data (without pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)
					cntr++

					dns := []dto.DnsRecord{
						{
							ID:       "1",
							Type:     dto.AType,
							Name:     "sub1",
							TTL:      110,
							Cloud:    true,
							Contents: []string{"145.100.10.21"},
						},
						{
							ID:       "2",
							Type:     dto.AType,
							Name:     "sub2",
							TTL:      120,
							Cloud:    true,
							Contents: []string{"145.100.10.22"},
						},
					}
					dnsBody, err := json.Marshal(dns)
					assert.NoError(t, err)
					respBody := fmt.Sprintf("{%q:%s}", "data", string(dnsBody))

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
				},
			},
			wants: []dto.DnsRecord{
				{
					ID:       "1",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub1",
					TTL:      110,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.21"}},
					Contents: []string{"145.100.10.21"},
				},
				{
					ID:       "2",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub2",
					TTL:      120,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.22"}},
					Contents: []string{"145.100.10.22"},
				},
			},
		},
		{
			name: "should error get dns records when make request in second try (with pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						dns := []dto.DnsRecord{
							{
								ID:       "1",
								Type:     dto.AType,
								Name:     "sub1",
								TTL:      110,
								Cloud:    true,
								Contents: []string{"145.100.10.21"},
							},
							{
								ID:       "2",
								Type:     dto.AType,
								Name:     "sub2",
								TTL:      120,
								Cloud:    true,
								Contents: []string{"145.100.10.22"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 2, "current_page", 1, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						return nil, errors.New("fail to fetch url")
					}
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records?page=2",
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should successfully get dns records data and in return empty records in second try (with pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						dns := []dto.DnsRecord{
							{
								ID:       "1",
								Type:     dto.AType,
								Name:     "sub1",
								TTL:      110,
								Cloud:    true,
								Contents: []string{"145.100.10.21"},
							},
							{
								ID:       "2",
								Type:     dto.AType,
								Name:     "sub2",
								TTL:      120,
								Cloud:    true,
								Contents: []string{"145.100.10.22"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 2, "current_page", 1, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						respBody := fmt.Sprintf("{%q:[]}", "data")

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					}
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records?page=2",
				},
			},
			wants: []dto.DnsRecord{
				{
					ID:       "1",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub1",
					TTL:      110,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.21"}},
					Contents: []string{"145.100.10.21"},
				},
				{
					ID:       "2",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub2",
					TTL:      120,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.22"}},
					Contents: []string{"145.100.10.22"},
				},
			},
		},
		{
			name: "should successfully get dns records data (with pagination)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						dns := []dto.DnsRecord{
							{
								ID:       "1",
								Type:     dto.AType,
								Name:     "sub1",
								TTL:      110,
								Cloud:    true,
								Contents: []string{"145.100.10.21"},
							},
							{
								ID:       "2",
								Type:     dto.AType,
								Name:     "sub2",
								TTL:      120,
								Cloud:    true,
								Contents: []string{"145.100.10.22"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 2, "current_page", 1, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						dns := []dto.DnsRecord{
							{
								ID:       "3",
								Type:     dto.AType,
								Name:     "sub3",
								TTL:      130,
								Cloud:    true,
								Contents: []string{"145.100.10.23"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 2, "current_page", 2, "last_page", 2)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					}
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records?page=2",
				},
			},
			wants: []dto.DnsRecord{
				{
					ID:       "1",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub1",
					TTL:      110,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.21"}},
					Contents: []string{"145.100.10.21"},
				},
				{
					ID:       "2",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub2",
					TTL:      120,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.22"}},
					Contents: []string{"145.100.10.22"},
				},
				{
					ID:       "3",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub3",
					TTL:      130,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.23"}},
					Contents: []string{"145.100.10.23"},
				},
			},
		},
		{
			name: "should successfully get dns records data (with pagination and limit item)",
			init: func(t *testing.T, wants []expectedInit) *clientApi {
				t.Helper()
				cntr := 1

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					defer func() { cntr++ }()
					assert.GreaterOrEqual(t, len(wants), cntr)
					assert.Equal(t, req.Method, wants[cntr-1].Method)
					assert.Equal(t, req.URL.String(), wants[cntr-1].URL)

					if cntr == 1 {
						dns := []dto.DnsRecord{
							{
								ID:       "1",
								Type:     dto.AType,
								Name:     "sub1",
								TTL:      110,
								Cloud:    true,
								Contents: []string{"145.100.10.21"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 3, "current_page", 1, "last_page", 3)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else if cntr == 2 {
						dns := []dto.DnsRecord{
							{
								ID:       "2",
								Type:     dto.AType,
								Name:     "sub2",
								TTL:      120,
								Cloud:    true,
								Contents: []string{"145.100.10.22"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 3, "current_page", 2, "last_page", 3)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					} else {
						dns := []dto.DnsRecord{
							{
								ID:       "3",
								Type:     dto.AType,
								Name:     "sub3",
								TTL:      130,
								Cloud:    true,
								Contents: []string{"145.100.10.23"},
							},
						}
						dnsBody, err := json.Marshal(dns)
						assert.NoError(t, err)
						respBody := fmt.Sprintf("{%q:%s,%q:{%q:%d,%q:%d,%q:%d}}", "data", string(dnsBody), "meta", "total", 3, "current_page", 3, "last_page", 3)

						resp := http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(bytes.NewBufferString(respBody)),
						}

						return &resp, nil
					}
				})

				return c
			},
			given: given{
				ctx:     context.Background(),
				zone:    "example.com",
				perPage: []int{1},
			},
			wantsInit: []expectedInit{
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records?per_page=1",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records?page=2&per_page=1",
				},
				{
					Method: "GET",
					URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records?page=3&per_page=1",
				},
			},
			wants: []dto.DnsRecord{
				{
					ID:       "1",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub1",
					TTL:      110,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.21"}},
					Contents: []string{"145.100.10.21"},
				},
				{
					ID:       "2",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub2",
					TTL:      120,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.22"}},
					Contents: []string{"145.100.10.22"},
				},
				{
					ID:       "3",
					Type:     dto.AType,
					Zone:     "example.com",
					Name:     "sub3",
					TTL:      130,
					Cloud:    true,
					Value:    []dto.ARecord{{IP: "145.100.10.23"}},
					Contents: []string{"145.100.10.23"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.wantsInit)

			out, err := c.GetDnsRecords(tt.given.ctx, tt.given.zone, tt.given.perPage...)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_ClientApi_CreateDnsRecord(t *testing.T) {
	type given struct {
		ctx    context.Context
		zone   string
		record dto.DnsRecord
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		init      func(t *testing.T, wants expectedInit) *clientApi
		wantsErr  expectedErr
		wantsInit expectedInit
		name      string
		given     given
		wants     dto.DnsRecord
	}{
		{
			name: "should error create dns when trying make request",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)
				c.apiEndpoint = "invalid&://host"

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error create dns when trying do",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wantsInit: expectedInit{
				Method: "POST",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error create dns when trying unmarshal body",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(&outputReaderUnmarshalResponseInvalid{}),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wantsInit: expectedInit{
				Method: "POST",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully create dns",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					dns := dto.DnsRecord{
						ID:       "1",
						Type:     dto.AType,
						Name:     "sub",
						TTL:      120,
						Cloud:    true,
						Contents: []string{"145.100.10.20"},
					}
					dnsBody, err := json.Marshal(dns)
					assert.NoError(t, err)
					respBody := fmt.Sprintf("{%q:%s}", "data", string(dnsBody))

					resp := http.Response{
						StatusCode: 201,
						Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wants: dto.DnsRecord{
				ID:       "1",
				Type:     dto.AType,
				Name:     "sub",
				TTL:      120,
				Cloud:    true,
				Value:    []dto.ARecord{{IP: "145.100.10.20"}},
				Contents: []string{"145.100.10.20"},
			},
			wantsInit: expectedInit{
				Method: "POST",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.wantsInit)

			out, err := c.CreateDnsRecord(tt.given.ctx, tt.given.zone, tt.given.record)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_ClientApi_UpdateDnsRecord(t *testing.T) {
	type given struct {
		ctx    context.Context
		zone   string
		record dto.DnsRecord
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		init      func(t *testing.T, wants expectedInit) *clientApi
		wantsErr  expectedErr
		wantsInit expectedInit
		name      string
		given     given
		wants     dto.DnsRecord
	}{
		{
			name: "should error update dns when trying make request",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)
				c.apiEndpoint = "invalid&://host"

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					ID:       "1",
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error update dns when trying do",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					ID:       "1",
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wantsInit: expectedInit{
				Method: "PUT",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records/1",
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error update dns when trying unmarshal body",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(&outputReaderUnmarshalResponseInvalid{}),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					ID:       "1",
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wantsInit: expectedInit{
				Method: "PUT",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records/1",
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully dns domain",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					dns := dto.DnsRecord{
						ID:       "1",
						Type:     dto.AType,
						Name:     "sub",
						TTL:      120,
						Cloud:    true,
						Contents: []string{"145.100.10.20"},
					}
					dnsBody, err := json.Marshal(dns)
					assert.NoError(t, err)
					respBody := fmt.Sprintf("{%q:%s}", "data", string(dnsBody))

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(respBody)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				zone: "example.com",
				record: dto.DnsRecord{
					ID:       "1",
					Type:     dto.AType,
					Name:     "sub",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"145.100.10.20"},
				},
			},
			wants: dto.DnsRecord{
				ID:       "1",
				Type:     dto.AType,
				Name:     "sub",
				TTL:      120,
				Cloud:    true,
				Value:    []dto.ARecord{{IP: "145.100.10.20"}},
				Contents: []string{"145.100.10.20"},
			},
			wantsInit: expectedInit{
				Method: "PUT",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records/1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.wantsInit)

			out, err := c.UpdateDnsRecord(tt.given.ctx, tt.given.zone, tt.given.record)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_ClientApi_DeleteDnsRecord(t *testing.T) {
	type given struct {
		ctx      context.Context
		zone     string
		recordId string
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name      string
		init      func(t *testing.T, wants expectedInit) *clientApi
		given     given
		wantsInit expectedInit
		wantsErr  expectedErr
	}{
		{
			name: "should error delete dns when trying make request",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)
				c.apiEndpoint = "invalid&://host"

				return c
			},
			given: given{
				ctx:      context.Background(),
				zone:     "example.com",
				recordId: "1",
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error delete dns when trying do",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				ctx:      context.Background(),
				zone:     "example.com",
				recordId: "1",
			},
			wantsInit: expectedInit{
				Method: "DELETE",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records/1",
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error delete dns when trying unmarshal body",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(&outputReaderUnmarshalResponseInvalid{}),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:      context.Background(),
				zone:     "example.com",
				recordId: "1",
			},
			wantsInit: expectedInit{
				Method: "DELETE",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records/1",
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully delete dns",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString("")),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:      context.Background(),
				zone:     "example.com",
				recordId: "1",
			},
			wantsInit: expectedInit{
				Method: "DELETE",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains/example.com/dns-records/1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.wantsInit)

			err := c.DeleteDnsRecord(tt.given.ctx, tt.given.zone, tt.given.recordId)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestArvanCloud_ClientApi_makeGetWithPagination(t *testing.T) {
	type given struct {
		ctx     context.Context
		resType interface{}
		path    url.URL
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  error
		action   string
		happened bool
	}
	tests := []struct {
		name       string
		init       func(t *testing.T, wants expectedInit) *clientApi
		given      given
		wants      interface{}
		waintsInit expectedInit
		wantsErr   expectedErr
	}{
		{
			name: "should error make get with pagination when trying make new request",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)
				c.apiEndpoint = "invalid&://host"

				return c
			},
			given: given{
				ctx:  context.Background(),
				path: url.URL{Path: "domains"},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
				action:   dto.UnknownActErr,
			},
		},
		{
			name: "should error make get with pagination when trying do",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				path: url.URL{Path: "/domains"},
			},
			waintsInit: expectedInit{
				Method: "GET",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &url.Error{},
			},
		},
		{
			name: "should error make get with pagination when trying unmarshal body",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(&outputReaderUnmarshalResponseInvalid{}),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:  context.Background(),
				path: url.URL{Path: "/domains"},
			},
			waintsInit: expectedInit{
				Method: "GET",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains",
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully make get with pagination",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(`{"id":1,"name":"a"}`)),
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				ctx:     context.Background(),
				path:    url.URL{Path: "/domains", RawQuery: "per_page=10"},
				resType: &outputDataUnmarshalResponseValid{},
			},
			waintsInit: expectedInit{
				Method: "GET",
				URL:    "https://napi.arvancloud.ir/cdn/4.0/domains?per_page=10",
			},
			wants: &outputDataUnmarshalResponseValid{ID: 1, Name: "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.waintsInit)

			err := c.makeGetWithPagination(tt.given.ctx, tt.given.path, tt.given.resType)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants, tt.given.resType)
		})
	}
}

func TestArvanCloud_ClientApi_newRequest(t *testing.T) {
	type given struct {
		reqBody interface{}
		method  string
		path    string
	}
	type expected struct {
		URL     *url.URL
		Header  http.Header
		Method  string
		Body    []byte
		Timeout time.Duration
	}
	type expectedErr struct {
		errType  error
		action   string
		happened bool
	}
	tests := []struct {
		given    given
		init     func(t *testing.T) *clientApi
		wantsErr expectedErr
		name     string
		wants    expected
	}{
		{
			name: "should error make new request when parse request body",
			init: func(t *testing.T) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given:    given{method: "POST", path: "/", reqBody: inputDataNewRequestInvalid{ID: 1}},
			wantsErr: expectedErr{happened: true},
		},
		{
			name: "should error make new request when method invalid (the body is valid)",
			init: func(t *testing.T) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given:    given{method: "@", path: "/", reqBody: inputDataNewRequestValid{ID: 1}},
			wantsErr: expectedErr{happened: true},
		},
		{
			name: "should successfully make new request (without body)",
			init: func(t *testing.T) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{method: "GET", path: "/", reqBody: nil},
			wants: expected{
				Timeout: _defaultTimeout,
				Method:  "GET",
				URL: func() *url.URL {
					u, _ := url.Parse("https://napi.arvancloud.ir/cdn/4.0/")
					return u
				}(),
				Header: map[string][]string{
					"Accept":        {"application/json"},
					"Authorization": {"token"},
					"User-Agent":    {"ExternalDNS/unknown"},
				},
			},
		},
		{
			name: "should successfully make new request (with body)",
			init: func(t *testing.T) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{method: "GET", path: "/", reqBody: inputDataNewRequestValid{ID: 1}},
			wants: expected{
				Timeout: _defaultTimeout,
				Method:  "GET",
				URL: func() *url.URL {
					u, _ := url.Parse("https://napi.arvancloud.ir/cdn/4.0/")
					return u
				}(),
				Header: map[string][]string{
					"Accept":        {"application/json"},
					"Authorization": {"token"},
					"User-Agent":    {"ExternalDNS/unknown"},
					"Content-Type":  {"application/json;charset=utf-8"},
				},
				Body: []byte(`{"ID":1}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t)

			req, err := c.newRequest(tt.given.method, tt.given.path, tt.given.reqBody)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants.Timeout, c.client.Timeout)
			assert.Equal(t, tt.wants.Method, req.Method)
			assert.Equal(t, tt.wants.URL, req.URL)
			assert.Equal(t, tt.wants.Header, req.Header)
			if len(tt.wants.Body) > 0 {
				body, err := io.ReadAll(req.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.wants.Body, body)
			}
		})
	}
}

func TestArvanCloud_ClientApi_do(t *testing.T) {
	type given struct {
		req *http.Request
	}
	type expectedInit struct {
		Method string
		URL    string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		init      func(t *testing.T, wants expectedInit) *clientApi
		given     given
		wantsInit expectedInit
		name      string
		wantsErr  expectedErr
		wants     int
	}{
		{
			name: "should error do request",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					return nil, errors.New("fail to fetch url")
				})

				return c
			},
			given: given{
				req: func() *http.Request {
					req, _ := http.NewRequest("GET", "https://napi.arvancloud.ir/cdn/4.0", nil)
					return req
				}(),
			},
			wantsInit: expectedInit{Method: "GET", URL: "https://napi.arvancloud.ir/cdn/4.0"},
			wantsErr:  expectedErr{happened: true},
		},
		{
			name: "should error do request with action unknown when status found",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 302,
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				req: func() *http.Request {
					req, _ := http.NewRequest("GET", "https://napi.arvancloud.ir/cdn/4.0", nil)
					return req
				}(),
			},
			wantsInit: expectedInit{Method: "GET", URL: "https://napi.arvancloud.ir/cdn/4.0"},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.UnknownActErr,
			},
		},
		{
			name: "should error do request with action unauthorized when credential is empty or is invalid",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 401,
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				req: func() *http.Request {
					req, _ := http.NewRequest("GET", "https://napi.arvancloud.ir/cdn/4.0", nil)
					return req
				}(),
			},
			wantsInit: expectedInit{Method: "GET", URL: "https://napi.arvancloud.ir/cdn/4.0"},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.UnauthorizedActErr,
			},
		},
		{
			name: "should successfully do request",
			init: func(t *testing.T, wants expectedInit) *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				c.client = newHttpClientMock(func(req *http.Request) (*http.Response, error) {
					t.Helper()
					assert.Equal(t, req.Method, wants.Method)
					assert.Equal(t, req.URL.String(), wants.URL)

					resp := http.Response{
						StatusCode: 200,
					}

					return &resp, nil
				})

				return c
			},
			given: given{
				req: func() *http.Request {
					req, _ := http.NewRequest("GET", "https://napi.arvancloud.ir/cdn/4.0", nil)
					return req
				}(),
			},
			wantsInit: expectedInit{Method: "GET", URL: "https://napi.arvancloud.ir/cdn/4.0"},
			wants:     200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init(t, tt.wantsInit)

			resp, err := c.do(tt.given.req)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants, resp.StatusCode)
		})
	}
}

func TestArvanCloud_ClientApi_unmarshalResponse(t *testing.T) {
	type given struct {
		response *http.Response
		resType  interface{}
	}
	type expectedErr struct {
		errType  interface{}
		errStr   string
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		init     func() *clientApi
		given    given
		wants    interface{}
		wantsErr expectedErr
	}{
		{
			name: "should error unmarshal response when read body",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(&outputReaderUnmarshalResponseInvalid{}),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wantsErr: expectedErr{happened: true},
		},
		{
			name: "should error unmarshal response when read body (also error on close response body)",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 200,
					Body:       &mockReadCloserErr{},
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wantsErr: expectedErr{happened: true},
		},
		{
			name: "should error unmarshal response when status code is less than 200 or great and equal than 300 - error is happened when unmarshal error body",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewBufferString(`{"status":error}`)),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &json.SyntaxError{},
			},
		},
		{
			name: "should error unmarshal response when status code is less than 200 or great and equal than 300 (with default error)",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewBufferString(`{"status":"error"}`)),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				errStr:   "error is happened (status: 400)",
				action:   dto.UnknownActErr,
			},
		},
		{
			name: "should error unmarshal response when status code is less than 200 or great and equal than 300 (with message error)",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 400,
					Body:       io.NopCloser(bytes.NewBufferString(`{"status":"error","message": "custom error"}`)),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				errStr:   "custom error (status: 400)",
				action:   dto.UnknownActErr,
			},
		},
		{
			name: "should error unmarshal response when status code is 200 (invalid body)",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(`{"id":"1","name":"a"}`)),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &json.UnmarshalTypeError{},
			},
		},
		{
			name: "should successfully unmarshal response return default struct with empty body",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString("")),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wants: &outputDataUnmarshalResponseValid{},
		},
		{
			name: "should successfully unmarshal response with is filled body",
			init: func() *clientApi {
				t.Helper()

				c, err := NewClientApi("token", "4.0")
				assert.NoError(t, err)

				return c
			},
			given: given{
				response: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(`{"id":1,"name":"a"}`)),
				},
				resType: &outputDataUnmarshalResponseValid{},
			},
			wants: &outputDataUnmarshalResponseValid{ID: 1, Name: "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.init()

			err := c.unmarshalResponse(tt.given.response, tt.given.resType)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&dto.ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*dto.ProviderError).GetAction())
					}
				}
				if tt.wantsErr.errStr != "" {
					assert.ErrorContains(t, err, tt.wantsErr.errStr)
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants, tt.given.resType)
		})
	}
}

func TestArvanCloud_ClientApi_getPage(t *testing.T) {
	type given struct {
		perPage []int
	}
	tests := []struct {
		name  string
		given given
		wants pageIn
	}{
		{
			name:  "should get default if per page is empty",
			given: given{perPage: nil},
			wants: pageIn{current: 1, limit: 0},
		},
		{
			name:  "should set to 100",
			given: given{perPage: []int{100}},
			wants: pageIn{current: 1, limit: 100},
		},
		{
			name:  "should set to 100 from first index of per page",
			given: given{perPage: []int{100, 200}},
			wants: pageIn{current: 1, limit: 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getPage(tt.given.perPage)

			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_ClientApi_setPageQueryString(t *testing.T) {
	type given struct {
		urlPath url.URL
		page    pageIn
	}
	tests := []struct {
		wants url.URL
		name  string
		given given
	}{
		{
			name:  "should skipping add query string when current page less than 1 and limit is equal 0",
			given: given{urlPath: url.URL{Path: "/"}, page: pageIn{}},
			wants: url.URL{Path: "/"},
		},
		{
			name:  "should skipping add query string when limit is equal 0",
			given: given{urlPath: url.URL{Path: "/"}, page: pageIn{current: 2}},
			wants: url.URL{Path: "/", RawQuery: "page=2"},
		},
		{
			name:  "should skipping add query string when current page less than 1",
			given: given{urlPath: url.URL{Path: "/"}, page: pageIn{limit: 10}},
			wants: url.URL{Path: "/", RawQuery: "per_page=10"},
		},
		{
			name:  "should skipping add query string when both of current page and limit is set",
			given: given{urlPath: url.URL{Path: "/"}, page: pageIn{current: 2, limit: 10}},
			wants: url.URL{Path: "/", RawQuery: "page=2&per_page=10"},
		},
		{
			name:  "should skipping add query string when both of current page and limit is set (with exist query string)",
			given: given{urlPath: url.URL{Path: "/", RawQuery: "search=this-is-search"}, page: pageIn{current: 2, limit: 10}},
			wants: url.URL{Path: "/", RawQuery: "page=2&per_page=10&search=this-is-search"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := setPageQueryString(tt.given.urlPath, tt.given.page)

			assert.Equal(t, tt.wants, out)
		})
	}
}
