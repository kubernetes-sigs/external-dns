/*
Copyright 2024 The Kubernetes Authors.

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

package wrappers

import (
	"context"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source"
)

// Validates that dedupSource is a Source
var _ source.Source = &nat64Source{}

func TestNAT64Source(t *testing.T) {
	t.Run("Endpoints", testNat64Source)
}

// testDedupEndpoints tests that duplicates from the wrapped source are removed.
func testNat64Source(t *testing.T) {
	for _, tc := range []struct {
		title     string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			"single non-nat64 ipv6 endpoint returns one ipv6 endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8:1::1"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8:1::1"}},
			},
		},
		{
			"single nat64 ipv6 endpoint returns one ipv4 endpoint and one ipv6 endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::192.0.2.42"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::192.0.2.42"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.0.2.42"}},
			},
		},
		{
			"single nat64 ipv6 endpoint returns one ipv4 endpoint and one ipv6 endpoint",
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::c000:22a"}},
			},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::c000:22a"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.0.2.42"}},
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.endpoints, nil)

			// Create our object under test and get the endpoints.
			source, err := NewNAT64Source(mockSource, []string{"2001:DB8::/96"})
			require.NoError(t, err)

			endpoints, err := source.Endpoints(context.Background())
			require.NoError(t, err)

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)

			// Validate that the mock source was called.
			mockSource.AssertExpectations(t)
		})
	}
}

func TestNat64Source_AddEventHandler(t *testing.T) {
	tests := []struct {
		title string
		input []string
		times int
	}{
		{
			title: "should add event handler when prefixes are provided",
			input: []string{"2001:DB8::/96"},
			times: 1,
		},
		{
			title: "should add event handler when prefixes not provided",
			input: []string{},
			times: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			mockSource := testutils.NewMockSource()

			src, err := NewNAT64Source(mockSource, tt.input)
			require.NoError(t, err)

			src.AddEventHandler(t.Context(), func() {})

			mockSource.AssertNumberOfCalls(t, "AddEventHandler", tt.times)
		})
	}
}

func TestNewNAT64Source(t *testing.T) {
	type args struct {
		source        source.Source
		nat64Prefixes []string
	}
	tests := []struct {
		name    string
		args    args
		want    source.Source
		wantErr bool
	}{
		{
			name: "empty NAT64 prefixes should succeed",
			args: args{
				source:        &testutils.MockSource{},
				nat64Prefixes: []string{},
			},
			want: &nat64Source{source: &testutils.MockSource{}, nat64Prefixes: []netip.Prefix{}},
		},
		{
			name: "multiple valid NAT64 prefixes should succeed",
			args: args{
				source:        &testutils.MockSource{},
				nat64Prefixes: []string{"2001:db8::/96", "64:ff9b::/96"},
			},
			want: &nat64Source{source: &testutils.MockSource{}, nat64Prefixes: []netip.Prefix{netip.MustParsePrefix("2001:db8::/96"), netip.MustParsePrefix("64:ff9b::/96")}},
		},
		{
			name: "invalid NAT64 prefix should fail",
			args: args{
				source:        &testutils.MockSource{},
				nat64Prefixes: []string{"invalid-prefix"},
			},
			wantErr: true,
		},
		{
			name: "NAT64 prefix with wrong mask length should fail",
			args: args{
				source:        &testutils.MockSource{},
				nat64Prefixes: []string{"2001:db8::/64"},
			},
			wantErr: true,
		},
		{
			name: "IPv4 address as NAT64 prefix should fail",
			args: args{
				source:        &testutils.MockSource{},
				nat64Prefixes: []string{"192.0.2.0/24"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, err := NewNAT64Source(tt.args.source, tt.args.nat64Prefixes)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, src)
		})
	}
}

func TestNat64SourceEndpoints_VariousCases(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn []*endpoint.Endpoint
		mockError  error
		setup      func()
		asserts    func([]*endpoint.Endpoint, error)
	}{
		{
			name:      "expect source error propagation",
			mockError: assert.AnError,
			asserts: func(eps []*endpoint.Endpoint, err error) {
				assert.Nil(t, eps)
				require.Error(t, err)
				require.ErrorIs(t, err, assert.AnError)
			},
		},
		{
			name: "skip nat64 processing for non-AAAA records",
			mockReturn: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.10.10.11"}},
			},
			asserts: func(eps []*endpoint.Endpoint, err error) {
				assert.NotNil(t, eps)
				assert.Len(t, eps, 1)
				require.NoError(t, err)
			},
		},
		{
			name: "target is not a valid IPv6 address",
			mockReturn: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"not-an-ip"}},
			},
			asserts: func(eps []*endpoint.Endpoint, err error) {
				assert.Nil(t, eps)
				require.Error(t, err)
				assert.EqualError(t, err, "ParseAddr(\"not-an-ip\"): unable to parse IP")
			},
		},
		{
			name: "addr from slice fails",
			mockReturn: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::192.0.2.42"}},
			},
			setup: func() {
				originalAddrFromSlice := addrFromSlice
				addrFromSlice = func([]byte) (netip.Addr, bool) {
					return netip.Addr{}, false
				}
				t.Cleanup(func() {
					addrFromSlice = originalAddrFromSlice
				})
			},
			asserts: func(eps []*endpoint.Endpoint, err error) {
				assert.Nil(t, eps)
				require.Error(t, err)
				assert.EqualError(t, err, "could not parse [192 0 2 42] to IPv4 address")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup()
			}
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.mockReturn, tc.mockError)

			src, err := NewNAT64Source(mockSource, []string{"2001:db8::/96"})
			require.NoError(t, err)

			eps, err := src.Endpoints(context.Background())
			tc.asserts(eps, err)

			mockSource.AssertExpectations(t)
		})
	}
}
