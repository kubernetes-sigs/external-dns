/*
Copyright 2020 The Kubernetes Authors.

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
	"context"
	"reflect"
	"testing"

	adguard "github.com/markussiebert/go-adguardhome-client/client"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

func Test_basicAuth(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basicauth works",
			args: args{
				username: "user",
				password: "password",
			},
			want: "Basic dXNlcjpwYXNzd29yZA==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicAuth(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("basicAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAdguardHomeProvider(t *testing.T) {
	type args struct {
		domainFilter endpoint.DomainFilter
		dryRun       bool
	}
	tests := []struct {
		name    string
		args    args
		want    *AdguardHomeProvider
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAdguardHomeProvider(tt.args.domainFilter, tt.args.dryRun)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAdguardHomeProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdguardHomeProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdguardHomeProvider_ApplyChanges(t *testing.T) {
	type fields struct {
		BaseProvider provider.BaseProvider
		client       adguard.ClientWithResponses
		domainFilter endpoint.DomainFilter
		DryRun       bool
	}
	type args struct {
		ctx     context.Context
		changes *plan.Changes
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AdguardHomeProvider{
				BaseProvider: tt.fields.BaseProvider,
				client:       tt.fields.client,
				domainFilter: tt.fields.domainFilter,
				DryRun:       tt.fields.DryRun,
			}
			if err := p.ApplyChanges(tt.args.ctx, tt.args.changes); (err != nil) != tt.wantErr {
				t.Errorf("AdguardHomeProvider.ApplyChanges() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdguardHomeProvider_Records(t *testing.T) {
	type fields struct {
		BaseProvider provider.BaseProvider
		client       adguard.ClientWithResponses
		domainFilter endpoint.DomainFilter
		DryRun       bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*endpoint.Endpoint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AdguardHomeProvider{
				BaseProvider: tt.fields.BaseProvider,
				client:       tt.fields.client,
				domainFilter: tt.fields.domainFilter,
				DryRun:       tt.fields.DryRun,
			}
			got, err := p.Records(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AdguardHomeProvider.Records() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdguardHomeProvider.Records() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genAdguardCustomRuleRecord(t *testing.T) {
	type args struct {
		e *endpoint.Endpoint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic record",
			args: args{
				e: &endpoint.Endpoint{
					DNSName:    "some.basic.record",
					RecordType: "A",
					Targets: endpoint.Targets{
						"1.2.3.4",
					},
				},
			},
			want: "some.basic.record$dnstype=A,dnsrewrite=NOERROR;A;1.2.3.4",
		},
		{
			name: "wildcard record",
			args: args{
				e: &endpoint.Endpoint{
					DNSName:    "*.some.basic.record",
					RecordType: "A",
					Targets: endpoint.Targets{
						"1.2.3.4",
					},
				},
			},
			want: "*.some.basic.record$dnstype=A,dnsrewrite=NOERROR;A;1.2.3.4",
		},
		{
			name: "prefixed wildcard record",
			args: args{
				e: &endpoint.Endpoint{
					DNSName:    "a-*.some.basic.record",
					RecordType: "A",
					Targets: endpoint.Targets{
						"1.2.3.4",
					},
				},
			},
			want: "a-*.some.basic.record$dnstype=A,dnsrewrite=NOERROR;A;1.2.3.4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genAdguardCustomRuleRecord(tt.args.e); got != tt.want {
				t.Errorf("genAdguardCustomRuleRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseAdguardCustomRuleRecord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *endpoint.Endpoint
	}{
		{
			name: "won't parse",
			args: args{
				s: "hopefully.this.wont?parse",
			},
			want: nil,
		},
		{
			name: "basic record parsing",
			args: args{
				s: "some.basic.record$dnstype=A,dnsrewrite=NOERROR;A;1.2.3.4",
			},
			want: &endpoint.Endpoint{
				DNSName:    "some.basic.record",
				RecordType: "A",
				Targets: endpoint.Targets{
					"1.2.3.4",
				},
			},
		},
		{
			name: "wildcard record parsing",
			args: args{
				s: "*.some.basic.record$dnstype=A,dnsrewrite=NOERROR;A;1.2.3.4",
			},
			want: &endpoint.Endpoint{
				DNSName:    "*.some.basic.record",
				RecordType: "A",
				Targets: endpoint.Targets{
					"1.2.3.4",
				},
			},
		},
		{
			name: "prefixed wildcard record parsing",
			args: args{
				s: "a-*.some.basic.record$dnstype=A,dnsrewrite=NOERROR;A;1.2.3.4",
			},
			want: &endpoint.Endpoint{
				DNSName:    "a-*.some.basic.record",
				RecordType: "A",
				Targets: endpoint.Targets{
					"1.2.3.4",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAdguardCustomRuleRecord(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAdguardCustomRuleRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeRule(t *testing.T) {
	type args struct {
		rules            []string
		ruleToRemove     string
		finishAfterFirst bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple remove rule",
			args: args{
				rules:        []string{"one", "twoe", "three", "four", "five"},
				ruleToRemove: "three",
			},
			want: []string{"one", "twoe", "four", "five"},
		},
		{
			name: "simple remove rule, finish after first match",
			args: args{
				rules:            []string{"one", "twoe", "three", "four", "three", "five"},
				ruleToRemove:     "three",
				finishAfterFirst: true,
			},
			want: []string{"one", "twoe", "four", "three", "five"},
		},
		{
			name: "simple remove rule, don't finish after first match",
			args: args{
				rules:            []string{"one", "twoe", "three", "four", "three", "five"},
				ruleToRemove:     "three",
				finishAfterFirst: false,
			},
			want: []string{"one", "twoe", "four", "five"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeRule(tt.args.rules, tt.args.ruleToRemove, tt.args.finishAfterFirst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
