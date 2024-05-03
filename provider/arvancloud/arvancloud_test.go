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

package arvancloud

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sigs.k8s.io/external-dns/provider/arvancloud/api"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/arvancloud/dto"
)

func TestArvanCloud_NewArvanCloudProvider(t *testing.T) {
	type given struct {
		domainFilter endpoint.DomainFilter
		zoneIDFilter provider.ZoneIDFilter
		dryRun       bool
		opts         []Option
	}
	type expectedErr struct {
		errType  interface{}
		errStr   string
		action   string
		happened bool
	}
	tests := []struct {
		wantsErr expectedErr
		init     func(t *testing.T)
		name     string
		given    given
		wants    *Provider
	}{
		{
			name: "should error create new instance of ArvanCloud provider when validate options",
			init: func(t *testing.T) {
				t.Helper()
			},
			given: given{
				domainFilter: endpoint.DomainFilter{},
				zoneIDFilter: provider.ZoneIDFilter{},
				dryRun:       false,
				opts:         []Option{WithApiVersion("invalid-version")},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.ApiVersionActErr,
			},
		},
		{
			name: "should error create new instance of ArvanCloud provider when error on get token",
			init: func(t *testing.T) {
				t.Helper()
			},
			given: given{
				domainFilter: endpoint.DomainFilter{},
				zoneIDFilter: provider.ZoneIDFilter{},
				dryRun:       false,
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.ApiTokenRequireActErr,
			},
		},
		{
			name: "should successfully create new instance of ArvanCloud provider",
			init: func(t *testing.T) {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})
			},
			given: given{
				domainFilter: endpoint.DomainFilter{},
				zoneIDFilter: provider.ZoneIDFilter{},
				dryRun:       false,
			},
			wants: &Provider{
				domainFilter: endpoint.DomainFilter{},
				zoneFilter:   provider.ZoneIDFilter{},
				client: func() ArvanAdapter {
					t.Helper()
					out, err := api.NewClientApi("this-is-a-fake-value", "4.0")
					assert.NoError(t, err)
					return out
				}(),
				options: providerOptions{
					apiVersion:       "4.0",
					dnsPerPage:       300,
					domainPerPage:    15,
					enableCloudProxy: true,
					dryRun:           false,
				},
			},
		},
		{
			name: "should successfully create new instance of ArvanCloud provider (with change default options)",
			init: func(t *testing.T) {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})
			},
			given: given{
				domainFilter: endpoint.DomainFilter{},
				zoneIDFilter: provider.ZoneIDFilter{},
				dryRun:       true,
				opts: []Option{
					WithApiVersion("3.0"),
					WithDisableCloudProxy(),
					WithDnsPerPage(20),
					WithDomainPerPage(10),
				},
			},
			wants: &Provider{
				domainFilter: endpoint.DomainFilter{},
				zoneFilter:   provider.ZoneIDFilter{},
				client: func() ArvanAdapter {
					t.Helper()
					out, err := api.NewClientApi("this-is-a-fake-value", "3.0")
					assert.NoError(t, err)
					return out
				}(),
				options: providerOptions{
					apiVersion:       "3.0",
					dnsPerPage:       20,
					domainPerPage:    10,
					enableCloudProxy: false,
					dryRun:           true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init(t)

			out, err := NewArvanCloudProvider(tt.given.domainFilter, tt.given.zoneIDFilter, tt.given.dryRun, tt.given.opts...)

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
					fmt.Println(err, tt.wantsErr.errStr)
					assert.ErrorContains(t, err, tt.wantsErr.errStr)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_Records(t *testing.T) {
	type given struct {
		ctx context.Context
	}
	type expectedErr struct {
		errType  interface{}
		errStr   string
		action   string
		happened bool
	}
	tests := []struct {
		init     func(t *testing.T) *Provider
		wantsErr expectedErr
		name     string
		given    given
		wants    []*endpoint.Endpoint
	}{
		{
			name: "should error get records when get domains",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						return nil, errors.New("fail to fetch domain")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			wantsErr: expectedErr{
				happened: true,
				errStr:   "fail to fetch domain",
			},
		},
		{
			name: "should error get records when get dns records",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						return nil, errors.New("fail to fetch record")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			wantsErr: expectedErr{
				happened: true,
				errStr:   "fail to fetch record",
			},
		},
		{
			name: "should successfully get records",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				cntr := 0
				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
							{
								ID:   "2",
								Name: "example2.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						defer func() { cntr++ }()
						var out []dto.DnsRecord

						if cntr == 0 {
							data := []dto.DnsRecord{
								{
									ID:       "11",
									Name:     "sub11",
									Zone:     "example1.com",
									Type:     dto.AType,
									TTL:      111,
									Contents: []string{"192.168.1.1"},
								},
								{
									ID:       "12",
									Name:     "sub12",
									Zone:     "example1.com",
									Type:     dto.AType,
									TTL:      121,
									Contents: []string{"192.168.1.2"},
								},
							}
							out = append(out, data...)
						} else {
							data := []dto.DnsRecord{
								{
									ID:       "21",
									Name:     "sub21",
									Zone:     "example2.com",
									Type:     dto.AType,
									TTL:      100,
									Contents: []string{"192.168.2.1"},
								},
								{
									ID:       "22",
									Name:     "sub22",
									Zone:     "example2.com",
									Type:     dto.CAAType,
									TTL:      110,
									Contents: []string{`issue "letsencrypt.org"`},
								},
							}
							out = append(out, data...)
						}

						return out, nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			wants: []*endpoint.Endpoint{
				{
					RecordType: "A",
					RecordTTL:  111,
					DNSName:    "sub11.example1.com",
					Targets:    []string{"192.168.1.1"},
					Labels:     endpoint.Labels{},
				},
				{
					RecordType: "A",
					RecordTTL:  121,
					DNSName:    "sub12.example1.com",
					Targets:    []string{"192.168.1.2"},
					Labels:     endpoint.Labels{},
				},
				{
					RecordType: "A",
					RecordTTL:  100,
					DNSName:    "sub21.example2.com",
					Targets:    []string{"192.168.2.1"},
					Labels:     endpoint.Labels{},
				},
			},
		},
		{
			name: "should successfully get records (with match zone)",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{Filters: []string{"example2.com"}}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
							{
								ID:   "2",
								Name: "example2.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "11",
								Name:     "sub11",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      111,
								Contents: []string{"192.168.1.1"},
							},
							{
								ID:       "12",
								Name:     "sub12",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.1.2"},
							},
						}

						return out, nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			wants: []*endpoint.Endpoint{
				{
					RecordType: "A",
					RecordTTL:  111,
					DNSName:    "sub11.example1.com",
					Targets:    []string{"192.168.1.1"},
					Labels:     endpoint.Labels{},
				},
				{
					RecordType: "A",
					RecordTTL:  121,
					DNSName:    "sub12.example1.com",
					Targets:    []string{"192.168.1.2"},
					Labels:     endpoint.Labels{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.init(t)

			out, err := p.Records(tt.given.ctx)

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
					fmt.Println(err, tt.wantsErr.errStr)
					assert.ErrorContains(t, err, tt.wantsErr.errStr)
				}
				return
			}
			assert.NoError(t, err)

			sort.SliceStable(tt.wants, func(i, j int) bool {
				return tt.wants[i].DNSName < tt.wants[j].DNSName
			})
			sort.SliceStable(out, func(i, j int) bool {
				return out[i].DNSName < out[j].DNSName
			})
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_AdjustEndpoints(t *testing.T) {
	type given struct {
		endpoints []*endpoint.Endpoint
	}
	type expectedErr struct {
		errType  interface{}
		errStr   string
		action   string
		happened bool
	}
	tests := []struct {
		init     func(t *testing.T) *Provider
		wantsErr expectedErr
		name     string
		given    given
		wants    []*endpoint.Endpoint
	}{
		{
			name: "should successfully skip adjust endpoints if no endpoint records exist",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
		},
		{
			name: "should successfully adjust endpoints (the default proxy mode is false)",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false, WithDisableCloudProxy())
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				endpoints: []*endpoint.Endpoint{
					{
						RecordType: "A",
						DNSName:    "sub1.example.com",
						Targets:    []string{"192.168.1.1", "192.168.1.2"},
					},
					{
						RecordType:       "A",
						DNSName:          "sub2.example.com",
						Targets:          []string{"192.168.2.1", "192.168.2.2"},
						ProviderSpecific: endpoint.ProviderSpecific{{Name: "arvan/cloud-proxy", Value: "true"}},
					},
					{
						RecordType: "A",
						RecordTTL:  200,
						DNSName:    "sub3.example.com",
						Targets:    []string{"192.168.3.1", "192.168.3.2"},
					},
				},
			},
			wants: []*endpoint.Endpoint{
				{
					RecordType:       "A",
					DNSName:          "sub1.example.com",
					Targets:          []string{"192.168.1.1", "192.168.1.2"},
					ProviderSpecific: endpoint.ProviderSpecific{{Name: "arvan/cloud-proxy", Value: "false"}},
				},
				{
					RecordType:       "A",
					DNSName:          "sub2.example.com",
					Targets:          []string{"192.168.2.1", "192.168.2.2"},
					ProviderSpecific: endpoint.ProviderSpecific{{Name: "arvan/cloud-proxy", Value: "true"}},
				},
				{
					RecordType:       "A",
					RecordTTL:        200,
					DNSName:          "sub3.example.com",
					Targets:          []string{"192.168.3.1", "192.168.3.2"},
					ProviderSpecific: endpoint.ProviderSpecific{{Name: "arvan/cloud-proxy", Value: "false"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.init(t)

			out, err := p.AdjustEndpoints(tt.given.endpoints)

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
					fmt.Println(err, tt.wantsErr.errStr)
					assert.ErrorContains(t, err, tt.wantsErr.errStr)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_ApplyChanges(t *testing.T) {
	type given struct {
		ctx     context.Context
		changes *plan.Changes
	}
	type expectedErr struct {
		errType  interface{}
		errStr   string
		action   string
		happened bool
	}
	tests := []struct {
		init     func(t *testing.T) *Provider
		wantsErr expectedErr
		name     string
		given    given
	}{
		{
			name: "should successfully skip apply changes if no change exist",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, true)
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
			},
		},
		{
			name: "should error apply changes when submit changes",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, true)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						return nil, errors.New("fail to fetch domain")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "sub1.example.com",
							Targets:    []string{"192.168.1.1", "192.168.1.2"},
						},
					},
					UpdateNew: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "sub2.example.com",
							RecordTTL:  200,
							Targets:    []string{"192.168.2.1", "192.168.2.2"},
						},
					},
					Delete: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "sub3.example.com",
							Targets:    []string{"192.168.3.1", "192.168.3.2"},
						},
					},
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errStr:   "fail to fetch domain",
			},
		},
		{
			name: "should successfully apply changes",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, true)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						return nil, nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changes: &plan.Changes{
					Create: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "sub1.example.com",
							Targets:    []string{"192.168.1.1", "192.168.1.2"},
						},
					},
					UpdateNew: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "sub2.example.com",
							RecordTTL:  200,
							Targets:    []string{"192.168.2.1", "192.168.2.2"},
						},
					},
					Delete: []*endpoint.Endpoint{
						{
							RecordType: "A",
							DNSName:    "sub3.example.com",
							Targets:    []string{"192.168.3.1", "192.168.3.2"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.init(t)

			err := p.ApplyChanges(tt.given.ctx, tt.given.changes)

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
					fmt.Println(err, tt.wantsErr.errStr)
					assert.ErrorContains(t, err, tt.wantsErr.errStr)
				}
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestArvanCloud_applyDnsChange(t *testing.T) {
	type given struct {
		endpointData *endpoint.Endpoint
		action       dto.DnsChangeAction
	}
	tests := []struct {
		name  string
		init  func(t *testing.T) *Provider
		given given
		wants dto.DnsChange
	}{
		{
			name: "should successfully apply dns change with default ttl and enable proxy",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false, WithEnableCloudProxy())
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				action: dto.CreateDns,
				endpointData: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "sub1.example.com",
					Targets:    []string{"192.168.1.1", "192.168.1.2"},
				},
			},
			wants: dto.DnsChange{
				Action: dto.CreateDns,
				Record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub1.example.com",
					TTL:      120,
					Cloud:    true,
					Contents: []string{"192.168.1.1", "192.168.1.2"},
				},
			},
		},
		{
			name: "should successfully apply dns change with default ttl and disable proxy",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false, WithDisableCloudProxy())
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				action: dto.CreateDns,
				endpointData: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "sub1.example.com",
					Targets:    []string{"192.168.1.1", "192.168.1.2"},
				},
			},
			wants: dto.DnsChange{
				Action: dto.CreateDns,
				Record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub1.example.com",
					TTL:      120,
					Cloud:    false,
					Contents: []string{"192.168.1.1", "192.168.1.2"},
				},
			},
		},
		{
			name: "should successfully apply dns change with specific ttl and disable proxy",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false, WithDisableCloudProxy())
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				action: dto.CreateDns,
				endpointData: &endpoint.Endpoint{
					RecordType: "A",
					RecordTTL:  300,
					DNSName:    "sub1.example.com",
					Targets:    []string{"192.168.1.1", "192.168.1.2"},
				},
			},
			wants: dto.DnsChange{
				Action: dto.CreateDns,
				Record: dto.DnsRecord{
					Type:     dto.AType,
					Name:     "sub1.example.com",
					TTL:      300,
					Cloud:    false,
					Contents: []string{"192.168.1.1", "192.168.1.2"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.init(t)

			out := p.applyDnsChange(tt.given.action, tt.given.endpointData)

			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_submitChanges(t *testing.T) {
	type given struct {
		ctx       context.Context
		changeSet []dto.DnsChange
	}
	type expectedErr struct {
		errType  interface{}
		errStr   string
		action   string
		happened bool
	}
	tests := []struct {
		init     func(t *testing.T) *Provider
		wantsErr expectedErr
		name     string
		given    given
	}{
		{
			name: "should successfully submit change and return if change is empty",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
		},
		{
			name: "should error submit change when get domains",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						return nil, errors.New("fail to fetch domain")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "11",
							Name: "sub1.example1.com",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errStr:   "fail to fetch domain",
			},
		},
		{
			name: "should successfully submit change and doesn't do anythings because get domains return empty records",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						return nil, nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "11",
							Name: "sub1.example1.com",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
				},
			},
		},
		{
			name: "should error submit change when get dns records",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						return nil, errors.New("fail to fetch record")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "11",
							Name: "sub1.example1.com",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errStr:   "fail to fetch record",
			},
		},
		{
			name: "should successfully submit change (with dry-run mode)",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, true)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "12",
								Name:     "sub2",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.2.2"},
							},
							{
								ID:       "13",
								Name:     "sub3",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      131,
								Contents: []string{"192.168.3.3"},
							},
							{
								ID:       "14",
								Name:     "sub4",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      140,
								Contents: []string{"192.168.1.4"},
							},
						}

						return out, nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							Name:     "sub2.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      120,
							Contents: []string{"192.168.1.2"},
						},
					},
					{
						Action: dto.UpdateDns,
						Record: dto.DnsRecord{
							Name:     "sub3.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      130,
							Contents: []string{"192.168.1.3"},
						},
					},
					{
						Action: dto.DeleteDns,
						Record: dto.DnsRecord{
							Name:     "sub4.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      140,
							Contents: []string{"192.168.1.4"},
						},
					},
				},
			},
		},
		{
			name: "should skip submit change when create dns record",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						return nil, nil
					},
					createDnsRecordOut: func() (dto.DnsRecord, error) {
						return dto.DnsRecord{}, errors.New("fail to do create dns")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
				},
			},
		},
		{
			name: "should skip submit change when update dns record (action is create)",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "11",
								Name:     "sub1",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.1.1"},
							},
						}

						return out, nil
					},
					updateDnsRecordOut: func() (dto.DnsRecord, error) {
						return dto.DnsRecord{}, errors.New("fail to do create dns")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
				},
			},
		},
		{
			name: "should skip submit change when update dns record (record id is not found)",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "11",
								Name:     "sub2",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.1.1"},
							},
						}

						return out, nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.UpdateDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
				},
			},
		},
		{
			name: "should skip submit change when update dns record",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "11",
								Name:     "sub1",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.1.1"},
							},
						}

						return out, nil
					},
					updateDnsRecordOut: func() (dto.DnsRecord, error) {
						return dto.DnsRecord{}, errors.New("fail to do create dns")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.UpdateDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
				},
			},
		},
		{
			name: "should skip submit change when delete dns record",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "11",
								Name:     "sub1",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.1.1"},
							},
						}

						return out, nil
					},
					deleteDnsRecordOut: func() error {
						return errors.New("fail to do create dns")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.DeleteDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
				},
			},
		},
		{
			name: "should skip submit change when delete dns record (record id is not found)",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "11",
								Name:     "sub2",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.1.1"},
							},
						}

						return out, nil
					},
					deleteDnsRecordOut: func() error {
						return errors.New("fail to do create dns")
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.DeleteDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
				},
			},
		},
		{
			name: "should successfully submit change",
			init: func(t *testing.T) *Provider {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
				p, err := NewArvanCloudProvider(endpoint.DomainFilter{}, provider.ZoneIDFilter{}, false)
				assert.NoError(t, err)

				p.client = &mockClient{
					t: t,
					getDomainsOut: func() ([]dto.Zone, error) {
						out := []dto.Zone{
							{
								ID:   "1",
								Name: "example1.com",
							},
						}

						return out, nil
					},
					getDnsRecordsOut: func() ([]dto.DnsRecord, error) {
						out := []dto.DnsRecord{
							{
								ID:       "12",
								Name:     "sub2",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      121,
								Contents: []string{"192.168.2.2"},
							},
							{
								ID:       "13",
								Name:     "sub3",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      131,
								Contents: []string{"192.168.3.3"},
							},
							{
								ID:       "14",
								Name:     "sub4",
								Zone:     "example1.com",
								Type:     dto.AType,
								TTL:      140,
								Contents: []string{"192.168.1.4"},
							},
						}

						return out, nil
					},
					createDnsRecordOut: func() (dto.DnsRecord, error) {
						return dto.DnsRecord{}, nil
					},
					updateDnsRecordOut: func() (dto.DnsRecord, error) {
						return dto.DnsRecord{}, nil
					},
					deleteDnsRecordOut: func() error {
						return nil
					},
				}

				t.Cleanup(func() {
					_ = os.Unsetenv("AC_API_TOKEN")
				})

				return p
			},
			given: given{
				ctx: context.Background(),
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							Name:     "sub1.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      110,
							Contents: []string{"192.168.1.1"},
						},
					},
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							Name:     "sub2.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      120,
							Contents: []string{"192.168.1.2"},
						},
					},
					{
						Action: dto.UpdateDns,
						Record: dto.DnsRecord{
							Name:     "sub3.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      130,
							Contents: []string{"192.168.1.3"},
						},
					},
					{
						Action: dto.DeleteDns,
						Record: dto.DnsRecord{
							Name:     "sub4.example1.com",
							Zone:     "example1.com",
							Type:     dto.AType,
							TTL:      140,
							Contents: []string{"192.168.1.4"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.init(t)

			err := p.submitChanges(tt.given.ctx, tt.given.changeSet)

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
		})
	}
}

func TestArvanCloud_changesByZone(t *testing.T) {
	type given struct {
		zones     []dto.Zone
		changeSet []dto.DnsChange
	}
	tests := []struct {
		wants map[string][]dto.DnsChange
		name  string
		given given
	}{
		{
			name:  "should successfully get changes with empty records",
			wants: map[string][]dto.DnsChange{},
		},
		{
			name: "should successfully get changes",
			given: given{
				zones: []dto.Zone{
					{
						ID:   "1",
						Name: "example1.com",
					},
					{
						ID:   "2",
						Name: "example2.com",
					},
				},
				changeSet: []dto.DnsChange{
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "11",
							Name: "sub1.example1.com",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "12",
							Name: "sub2.example1.com",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "13",
							Name: "sub3.test.com",
							Zone: "test.com",
							Type: dto.AType,
						},
					},
				},
			},
			wants: map[string][]dto.DnsChange{
				"example1.com": {
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "11",
							Name: "sub1",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
					{
						Action: dto.CreateDns,
						Record: dto.DnsRecord{
							ID:   "12",
							Name: "sub2",
							Zone: "example1.com",
							Type: dto.AType,
						},
					},
				},
				"example2.com": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := changesByZone(tt.given.zones, tt.given.changeSet)

			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_getRecordID(t *testing.T) {
	type given struct {
		records []dto.DnsRecord
		record  dto.DnsRecord
	}
	tests := []struct {
		name  string
		wants string
		given given
	}{
		{
			name: "should successfully get record ID with empty value if records is empty",
			given: given{
				record: dto.DnsRecord{Name: "sub1", Type: dto.AType},
			},
		},
		{
			name: "should successfully get record ID with empty value if not found ID in records",
			given: given{
				records: []dto.DnsRecord{{ID: "1", Name: "sub", Type: dto.MXType}, {ID: "2", Name: "sub", Type: dto.AType}},
				record:  dto.DnsRecord{Name: "sub", Type: dto.TXTType},
			},
		},
		{
			name: "should successfully get record ID",
			given: given{
				records: []dto.DnsRecord{{ID: "1", Name: "sub", Type: dto.MXType}, {ID: "2", Name: "sub", Type: dto.AType}},
				record:  dto.DnsRecord{Name: "sub", Type: dto.AType},
			},
			wants: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := getRecordID(tt.given.records, tt.given.record)

			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_groupByNameAndType(t *testing.T) {
	type given struct {
		records []dto.DnsRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		wantsErr expectedErr
		name     string
		given    given
		wants    []*endpoint.Endpoint
	}{
		{
			name: "should successfully group by name and type and return empty records",
		},
		{
			name: "should successfully group by name and type",
			given: given{
				records: []dto.DnsRecord{
					{
						Type:     dto.AType,
						Name:     "sub",
						Zone:     "example.com",
						TTL:      110,
						Contents: []string{"192.168.1.1"},
					},
					{
						Type:     dto.AType,
						Name:     "sub",
						Zone:     "example.com",
						TTL:      120,
						Contents: []string{"192.168.1.2"},
					},
					{
						Type:     dto.AType,
						Name:     "@",
						Zone:     "example.com",
						TTL:      130,
						Contents: []string{"192.168.1.3"},
					},
					{
						Type:     dto.SPFType,
						Name:     "sub",
						Zone:     "example.com",
						TTL:      140,
						Contents: []string{"v=spf1 include:_spf.example.com -all"},
					},
				},
			},
			wants: []*endpoint.Endpoint{
				{
					DNSName:    "sub.example.com",
					Targets:    []string{"192.168.1.1", "192.168.1.2"},
					RecordType: "A",
					RecordTTL:  110,
					Labels:     endpoint.Labels{},
				},
				{
					DNSName:    "example.com",
					Targets:    []string{"192.168.1.3"},
					RecordType: "A",
					RecordTTL:  130,
					Labels:     endpoint.Labels{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := groupByNameAndType(tt.given.records)

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

			sort.SliceStable(tt.wants, func(i, j int) bool {
				return tt.wants[i].DNSName < tt.wants[j].DNSName
			})
			sort.SliceStable(out, func(i, j int) bool {
				return out[i].DNSName < out[j].DNSName
			})
			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_shouldBeProxy(t *testing.T) {
	type given struct {
		endpoint       *endpoint.Endpoint
		proxyByDefault bool
	}
	tests := []struct {
		name  string
		given given
		wants bool
	}{
		{
			name: "should successfully proxy is equal false (default value) when endpoint is nil pointer",
			given: given{
				proxyByDefault: false,
			},
			wants: false,
		},
		{
			name: "should successfully proxy is equal true (default value) when endpoint is nil pointer",
			given: given{
				proxyByDefault: true,
			},
			wants: true,
		},
		{
			name: "should successfully proxy is equal true (default value) when doesn't found in provider - skip error when covert string to bool",
			given: given{
				endpoint: &endpoint.Endpoint{
					ProviderSpecific: []endpoint.ProviderSpecificProperty{
						{Name: "arvan/cloud-proxy", Value: "invalid"},
					},
				},
				proxyByDefault: true,
			},
			wants: true,
		},
		{
			name: "should successfully proxy is equal false when found in provider",
			given: given{
				endpoint: &endpoint.Endpoint{
					ProviderSpecific: []endpoint.ProviderSpecificProperty{
						{Name: "arvan/cloud-proxy", Value: "false"},
					},
				},
				proxyByDefault: true,
			},
			wants: false,
		},
		{
			name: "should successfully proxy is equal false when proxy is disable for record type",
			given: given{
				endpoint: &endpoint.Endpoint{
					RecordType: "AAAA",
				},
				proxyByDefault: true,
			},
			wants: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := shouldBeProxy(tt.given.endpoint, tt.given.proxyByDefault)

			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_defaultOptions(t *testing.T) {
	tests := []struct {
		name  string
		wants providerOptions
	}{
		{
			name: "should successfully get default option",
			wants: providerOptions{
				domainPerPage:    15,
				dnsPerPage:       300,
				enableCloudProxy: true,
				apiVersion:       "4.0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := defaultOptions()

			assert.Equal(t, tt.wants, out)
		})
	}
}

func TestArvanCloud_optionValidation(t *testing.T) {
	type given struct {
		option providerOptions
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		wantsErr expectedErr
		name     string
		given    given
	}{
		{
			name: "should error validate option for check api version",
			given: given{
				option: providerOptions{apiVersion: "invalid"},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.ApiVersionActErr,
			},
		},
		{
			name: "should successfully validate option",
			given: given{
				option: providerOptions{apiVersion: "4.0"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := optionValidation(tt.given.option)

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

func TestArvanCloud_getToken(t *testing.T) {
	defer func() {
		_ = os.Unsetenv("AC_API_TOKEN")
	}()

	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		init     func()
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "should error get token when env variable is empty",
			init: func() {},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.ApiTokenRequireActErr,
			},
		},
		{
			name: "should error get token if is registered as file but fail to read file",
			init: func() {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "file:/path/of/file")
				assert.NoError(t, err)
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &dto.ProviderError{},
				action:   dto.ApiTokenFromFileActErr,
			},
		},
		{
			name: "should successfully get token if is registered in env variable",
			init: func() {
				t.Helper()
				file, err := os.CreateTemp(os.TempDir(), "external-dns-arvancloud-token-*")
				assert.NoError(t, err)
				t.Cleanup(func() {
					_ = os.Remove(file.Name())
				})
				err = os.WriteFile(file.Name(), []byte("this-is-a-fake-value-from-file"), 0666)
				assert.NoError(t, err)
				err = os.Setenv("AC_API_TOKEN", fmt.Sprintf("file:%s", file.Name()))
				assert.NoError(t, err)

			},
			wants: "this-is-a-fake-value-from-file",
		},
		{
			name: "should successfully get token if is registered in env variable",
			init: func() {
				t.Helper()
				err := os.Setenv("AC_API_TOKEN", "this-is-a-fake-value")
				assert.NoError(t, err)
			},
			wants: "this-is-a-fake-value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init()

			out, err := getToken()

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
