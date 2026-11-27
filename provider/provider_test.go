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

package provider

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type apexStub struct {
	BaseProvider
	names []string
	err   error
}

type plainStub struct {
	BaseProvider
}

type loopUnwrapper struct {
	BaseProvider
}

func (s *apexStub) Records(context.Context) ([]*endpoint.Endpoint, error) { return nil, nil }
func (s *apexStub) ApplyChanges(context.Context, *plan.Changes) error     { return nil }
func (s *apexStub) ZoneApexes(context.Context) ([]string, error) {
	return s.names, s.err
}

func (s *plainStub) Records(context.Context) ([]*endpoint.Endpoint, error) { return nil, nil }
func (s *plainStub) ApplyChanges(context.Context, *plan.Changes) error     { return nil }

func (w *loopUnwrapper) Records(context.Context) ([]*endpoint.Endpoint, error) { return nil, nil }
func (w *loopUnwrapper) ApplyChanges(context.Context, *plan.Changes) error     { return nil }
func (w *loopUnwrapper) Unwrap() Provider                                      { return w }

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}

func TestEnsureTrailingDot(t *testing.T) {
	for _, tc := range []struct {
		input, expected string
	}{
		{"example.org", "example.org."},
		{"example.org.", "example.org."},
		{"8.8.8.8", "8.8.8.8"},
	} {
		output := EnsureTrailingDot(tc.input)

		if output != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, output)
		}
	}
}

func TestNewSoftError(t *testing.T) {
	cause := errors.New("something failed")
	err := NewSoftError(cause)
	require.ErrorIs(t, err, SoftError)
	require.ErrorIs(t, err, cause)
}

func TestNewSoftErrorf(t *testing.T) {
	err := NewSoftErrorf("failed with code %d", 42)
	require.ErrorIs(t, err, SoftError)
	assert.Contains(t, err.Error(), "failed with code 42")
}

func TestBaseProvider_AdjustEndpoints(t *testing.T) {
	b := BaseProvider{}
	eps := []*endpoint.Endpoint{{DNSName: "example.com"}}
	got, err := b.AdjustEndpoints(eps)
	require.NoError(t, err)
	assert.Equal(t, eps, got)
}

func TestBaseProvider_GetDomainFilter(t *testing.T) {
	b := BaseProvider{}
	f := b.GetDomainFilter()
	require.NotNil(t, f)
	assert.True(t, f.Match("example.com"))
}

func TestZoneApexerOf(t *testing.T) {
	stub := &apexStub{names: []string{"example.com", "example.org"}}

	apexer, ok := ZoneApexerOf(stub)
	require.True(t, ok)
	names, err := apexer.ZoneApexes(t.Context())
	require.NoError(t, err)
	assert.Equal(t, []string{"example.com", "example.org"}, names)

	apexer, ok = ZoneApexerOf(NewCachedProvider(stub, 0))
	require.True(t, ok)
	names, err = apexer.ZoneApexes(t.Context())
	require.NoError(t, err)
	assert.Equal(t, stub.names, names)

	_, ok = ZoneApexerOf(&plainStub{})
	assert.False(t, ok)

	_, ok = ZoneApexerOf(nil)
	assert.False(t, ok)

	_, ok = ZoneApexerOf(&loopUnwrapper{})
	assert.False(t, ok)
}

func TestDifference(t *testing.T) {
	current := []string{"foo", "bar"}
	desired := []string{"bar", "baz"}
	add, remove, leave := Difference(current, desired)
	assert.Equal(t, []string{"baz"}, add)
	assert.Equal(t, []string{"foo"}, remove)
	assert.Equal(t, []string{"bar"}, leave)
}
