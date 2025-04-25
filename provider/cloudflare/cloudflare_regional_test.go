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

package cloudflare

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

func Test_regionalHostname(t *testing.T) {
	type args struct {
		endpoint         *endpoint.Endpoint
		defaultRegionKey string
	}
	tests := []struct {
		name string
		args args
		want cloudflare.RegionalHostname
	}{
		{
			name: "no region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
				},
				defaultRegionKey: "",
			},
			want: cloudflare.RegionalHostname{},
		},
		{
			name: "default region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "us",
			},
		},
		{
			name: "endpoint with region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "eu",
			},
		},
		{
			name: "endpoint with empty region key",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "A",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "",
						},
					},
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{
				Hostname:  "example.com",
				RegionKey: "",
			},
		},
		{
			name: "unsupported record type",
			args: args{
				endpoint: &endpoint.Endpoint{
					RecordType: "TXT",
					DNSName:    "example.com",
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-region-key",
							Value: "eu",
						},
					},
				},
				defaultRegionKey: "us",
			},
			want: cloudflare.RegionalHostname{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CloudFlareProvider{RegionKey: tt.args.defaultRegionKey}
			got := p.regionalHostname(tt.args.endpoint)
			assert.Equal(t, got, tt.want)
		})
	}
}
