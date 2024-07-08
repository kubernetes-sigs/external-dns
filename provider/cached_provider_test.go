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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type testProviderFunc struct {
	records             func(ctx context.Context) ([]*endpoint.Endpoint, error)
	applyChanges        func(ctx context.Context, changes *plan.Changes) error
	propertyValuesEqual func(name string, previous string, current string) bool
	adjustEndpoints     func(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error)
	getDomainFilter     func() endpoint.DomainFilterInterface
}

func (p *testProviderFunc) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return p.records(ctx)
}

func (p *testProviderFunc) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	return p.applyChanges(ctx, changes)
}

func (p *testProviderFunc) PropertyValuesEqual(name string, previous string, current string) bool {
	return p.propertyValuesEqual(name, previous, current)
}

func (p *testProviderFunc) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return p.adjustEndpoints(endpoints)
}

func (p *testProviderFunc) GetDomainFilter() endpoint.DomainFilterInterface {
	return p.getDomainFilter()
}

func recordsNotCalled(t *testing.T) func(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return func(ctx context.Context) ([]*endpoint.Endpoint, error) {
		t.Errorf("unexpected call to Records")
		return nil, nil
	}
}

func applyChangesNotCalled(t *testing.T) func(ctx context.Context, changes *plan.Changes) error {
	return func(ctx context.Context, changes *plan.Changes) error {
		t.Errorf("unexpected call to ApplyChanges")
		return nil
	}
}

func propertyValuesEqualNotCalled(t *testing.T) func(name string, previous string, current string) bool {
	return func(name string, previous string, current string) bool {
		t.Errorf("unexpected call to PropertyValuesEqual")
		return false
	}
}

func adjustEndpointsNotCalled(t *testing.T) func(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return func(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
		t.Errorf("unexpected call to AdjustEndpoints")
		return endpoints, errors.New("unexpected call to AdjustEndpoints")
	}
}

func newTestProviderFunc(t *testing.T) *testProviderFunc {
	return &testProviderFunc{
		records:             recordsNotCalled(t),
		applyChanges:        applyChangesNotCalled(t),
		propertyValuesEqual: propertyValuesEqualNotCalled(t),
		adjustEndpoints:     adjustEndpointsNotCalled(t),
	}
}

func TestCachedProviderCallsProviderOnFirstCall(t *testing.T) {
	testProvider := newTestProviderFunc(t)
	testProvider.records = func(ctx context.Context) ([]*endpoint.Endpoint, error) {
		return []*endpoint.Endpoint{{DNSName: "domain.fqdn"}}, nil
	}
	provider := CachedProvider{
		Provider: testProvider,
	}
	endpoints, err := provider.Records(context.Background())
	assert.NoError(t, err)
	require.NotNil(t, endpoints)
	require.Len(t, endpoints, 1)
	require.NotNil(t, endpoints[0])
	assert.Equal(t, "domain.fqdn", endpoints[0].DNSName)
}

func TestCachedProviderUsesCacheWhileValid(t *testing.T) {
	testProvider := newTestProviderFunc(t)
	testProvider.records = func(ctx context.Context) ([]*endpoint.Endpoint, error) {
		return []*endpoint.Endpoint{{DNSName: "domain.fqdn"}}, nil
	}
	provider := CachedProvider{
		RefreshDelay: 30 * time.Second,
		Provider:     testProvider,
	}
	_, err := provider.Records(context.Background())
	require.NoError(t, err)

	t.Run("With consecutive calls within the caching time frame", func(t *testing.T) {
		testProvider.records = recordsNotCalled(t)
		endpoints, err := provider.Records(context.Background())
		assert.NoError(t, err)
		require.NotNil(t, endpoints)
		require.Len(t, endpoints, 1)
		require.NotNil(t, endpoints[0])
		assert.Equal(t, "domain.fqdn", endpoints[0].DNSName)
	})

	t.Run("When the caching time frame is exceeded", func(t *testing.T) {
		testProvider.records = func(ctx context.Context) ([]*endpoint.Endpoint, error) {
			return []*endpoint.Endpoint{{DNSName: "new.domain.fqdn"}}, nil
		}
		provider.lastRead = time.Now().Add(-20 * time.Minute)
		endpoints, err := provider.Records(context.Background())
		assert.NoError(t, err)
		require.NotNil(t, endpoints)
		require.Len(t, endpoints, 1)
		require.NotNil(t, endpoints[0])
		assert.Equal(t, "new.domain.fqdn", endpoints[0].DNSName)
	})
}

func TestCachedProviderForcesCacheRefreshOnUpdate(t *testing.T) {
	testProvider := newTestProviderFunc(t)
	testProvider.records = func(ctx context.Context) ([]*endpoint.Endpoint, error) {
		return []*endpoint.Endpoint{{DNSName: "domain.fqdn"}}, nil
	}
	provider := CachedProvider{
		RefreshDelay: 30 * time.Second,
		Provider:     testProvider,
	}
	_, err := provider.Records(context.Background())
	require.NoError(t, err)

	t.Run("When empty changes are applied", func(t *testing.T) {
		testProvider.records = recordsNotCalled(t)
		testProvider.applyChanges = func(ctx context.Context, changes *plan.Changes) error {
			return nil
		}
		err := provider.ApplyChanges(context.Background(), &plan.Changes{})
		assert.NoError(t, err)
		t.Run("Next call to Records is cached", func(t *testing.T) {
			testProvider.applyChanges = applyChangesNotCalled(t)
			testProvider.records = func(ctx context.Context) ([]*endpoint.Endpoint, error) {
				return []*endpoint.Endpoint{{DNSName: "new.domain.fqdn"}}, nil
			}
			endpoints, err := provider.Records(context.Background())

			assert.NoError(t, err)
			require.NotNil(t, endpoints)
			require.Len(t, endpoints, 1)
			require.NotNil(t, endpoints[0])
			assert.Equal(t, "domain.fqdn", endpoints[0].DNSName)
		})
	})

	t.Run("When changes are applied", func(t *testing.T) {
		testProvider.records = recordsNotCalled(t)
		testProvider.applyChanges = func(ctx context.Context, changes *plan.Changes) error {
			return nil
		}
		err := provider.ApplyChanges(context.Background(), &plan.Changes{
			Create: []*endpoint.Endpoint{
				{DNSName: "hello.world"},
			},
		})
		assert.NoError(t, err)
		t.Run("Next call to Records is not cached", func(t *testing.T) {
			testProvider.applyChanges = applyChangesNotCalled(t)
			testProvider.records = func(ctx context.Context) ([]*endpoint.Endpoint, error) {
				return []*endpoint.Endpoint{{DNSName: "new.domain.fqdn"}}, nil
			}
			endpoints, err := provider.Records(context.Background())

			assert.NoError(t, err)
			require.NotNil(t, endpoints)
			require.Len(t, endpoints, 1)
			require.NotNil(t, endpoints[0])
			assert.Equal(t, "new.domain.fqdn", endpoints[0].DNSName)
		})
	})
}
