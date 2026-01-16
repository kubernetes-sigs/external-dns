/*
Copyright 2026 The Kubernetes Authors.

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

package blueprint

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZoneCache_SliceCache(t *testing.T) {
	cache := NewSliceZoneCache[string](time.Hour)

	// Initially expired (empty)
	assert.True(t, cache.Expired())

	// After reset, not expired
	cache.Reset([]string{"zone1", "zone2"})
	assert.False(t, cache.Expired())
	assert.Equal(t, []string{"zone1", "zone2"}, cache.Get())
}

func TestZoneCache_MapCache(t *testing.T) {
	cache := NewMapZoneCache[string, int](time.Hour)

	// Initially expired (empty)
	assert.True(t, cache.Expired())

	// After reset, not expired
	cache.Reset(map[string]int{"a": 1, "b": 2})
	assert.False(t, cache.Expired())
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, cache.Get())
}

func TestZoneCache_Expiration(t *testing.T) {
	// Very short duration for testing
	cache := NewSliceZoneCache[string](10 * time.Millisecond)

	cache.Reset([]string{"zone1"})
	assert.False(t, cache.Expired())

	// Wait for expiration
	time.Sleep(20 * time.Millisecond)
	assert.True(t, cache.Expired())
}

func TestZoneCache_ZeroDuration(t *testing.T) {
	// Zero duration means caching is disabled
	cache := NewSliceZoneCache[string](0)

	cache.Reset([]string{"zone1"})
	// Should still be expired because caching is disabled
	assert.True(t, cache.Expired())
	// Data should not be stored
	assert.Nil(t, cache.Get())
}

func TestZoneCache_ThreadSafety(t *testing.T) {
	cache := NewSliceZoneCache[int](time.Hour)

	done := make(chan bool)

	// Writer goroutine
	go func() {
		for i := range 100 {
			cache.Reset([]int{i})
		}
		done <- true
	}()

	// Reader goroutine
	go func() {
		for range 100 {
			_ = cache.Get()
			_ = cache.Expired()
		}
		done <- true
	}()

	<-done
	<-done
}

func TestCachedZoneProvider_UsesCacheWhenNotExpired(t *testing.T) {
	cache := NewSliceZoneCache[string](time.Hour)
	fetchCount := 0

	fetcher := func(ctx context.Context) ([]string, error) {
		fetchCount++
		return []string{"zone1", "zone2"}, nil
	}

	provider := NewCachedZoneProvider(cache, fetcher, "test")

	// First call fetches
	zones, err := provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []string{"zone1", "zone2"}, zones)
	assert.Equal(t, 1, fetchCount)

	// Second call uses cache
	zones, err = provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Equal(t, []string{"zone1", "zone2"}, zones)
	assert.Equal(t, 1, fetchCount) // Still 1, no new fetch
}

func TestCachedZoneProvider_RefreshesOnExpiration(t *testing.T) {
	cache := NewSliceZoneCache[string](10 * time.Millisecond)
	fetchCount := 0

	fetcher := func(ctx context.Context) ([]string, error) {
		fetchCount++
		return []string{"zone1"}, nil
	}

	provider := NewCachedZoneProvider(cache, fetcher, "test")

	// First call
	_, err := provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, fetchCount)

	// Wait for expiration
	time.Sleep(20 * time.Millisecond)

	// Should fetch again
	_, err = provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 2, fetchCount)
}

func TestCachedZoneProvider_PropagatesErrors(t *testing.T) {
	cache := NewSliceZoneCache[string](time.Hour)
	expectedErr := errors.New("fetch error")

	fetcher := func(ctx context.Context) ([]string, error) {
		return nil, expectedErr
	}

	provider := NewCachedZoneProvider(cache, fetcher, "test")

	zones, err := provider.Zones(context.Background())
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, zones)
}

func TestCachedZoneProvider_Invalidate(t *testing.T) {
	cache := NewSliceZoneCache[string](time.Hour)
	fetchCount := 0

	fetcher := func(ctx context.Context) ([]string, error) {
		fetchCount++
		return []string{"zone1"}, nil
	}

	provider := NewCachedZoneProvider(cache, fetcher, "test")

	// First call
	_, err := provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, fetchCount)

	// Invalidate cache
	provider.Invalidate()

	// Should fetch again
	_, err = provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 2, fetchCount)
}

type testZone struct {
	ID   string
	Name string
}

func TestCachedZoneProvider_WithStructSlice(t *testing.T) {
	cache := NewSliceZoneCache[testZone](time.Hour)

	fetcher := func(ctx context.Context) ([]testZone, error) {
		return []testZone{
			{ID: "1", Name: "zone1.example.com"},
			{ID: "2", Name: "zone2.example.com"},
		}, nil
	}

	provider := NewCachedZoneProvider(cache, fetcher, "test")

	zones, err := provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Len(t, zones, 2)
	assert.Equal(t, "zone1.example.com", zones[0].Name)
}

func TestCachedZoneProvider_WithMap(t *testing.T) {
	cache := NewMapZoneCache[string, *testZone](time.Hour)

	fetcher := func(ctx context.Context) (map[string]*testZone, error) {
		return map[string]*testZone{
			"zone1": {ID: "1", Name: "zone1.example.com"},
			"zone2": {ID: "2", Name: "zone2.example.com"},
		}, nil
	}

	provider := NewCachedZoneProvider(cache, fetcher, "test")

	zones, err := provider.Zones(context.Background())
	require.NoError(t, err)
	assert.Len(t, zones, 2)
	assert.Equal(t, "zone1.example.com", zones["zone1"].Name)
}
