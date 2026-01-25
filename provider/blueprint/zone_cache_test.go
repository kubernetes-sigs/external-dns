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
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZoneCache_SliceCache(t *testing.T) {
	cache := NewZoneCache[[]string](time.Hour)

	// Initially expired (empty)
	assert.True(t, cache.Expired())

	// After reset, not expired
	cache.Reset([]string{"zone1", "zone2"})
	assert.False(t, cache.Expired())
	assert.Equal(t, []string{"zone1", "zone2"}, cache.Get())
}

func TestZoneCache_MapCache(t *testing.T) {
	cache := NewZoneCache[map[string]int](time.Hour)

	// Initially expired (empty)
	assert.True(t, cache.Expired())

	// After reset, not expired
	cache.Reset(map[string]int{"a": 1, "b": 2})
	assert.False(t, cache.Expired())
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, cache.Get())
}

func TestZoneCache_Expiration(t *testing.T) {
	// Very short duration for testing
	cache := NewZoneCache[[]string](10 * time.Millisecond)

	cache.Reset([]string{"zone1"})
	assert.False(t, cache.Expired())

	// Wait for expiration
	time.Sleep(20 * time.Millisecond)
	assert.True(t, cache.Expired())
}

func TestZoneCache_CachingDisabled(t *testing.T) {
	cache := NewZoneCache[[]string](0)

	cache.Reset([]string{"zone1"})
	// Should still be expired because caching is disabled
	assert.True(t, cache.Expired())
	// Data should not be stored
	assert.Nil(t, cache.Get())
}

func TestZoneCache_Expiration_Synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		cache := NewZoneCache[[]string](5 * time.Minute)

		cache.Reset([]string{"zone1", "zone2"})
		assert.False(t, cache.Expired(), "should not be expired immediately after reset")
		assert.Equal(t, []string{"zone1", "zone2"}, cache.Get())

		// Advance time but not past expiration
		time.Sleep(3 * time.Minute)
		assert.False(t, cache.Expired(), "should not be expired before duration")

		// Advance time past expiration
		time.Sleep(3 * time.Minute) // Total: 6 minutes > 5 minute duration
		assert.True(t, cache.Expired(), "should be expired after duration")

		// Data is still accessible even when expired
		assert.Equal(t, []string{"zone1", "zone2"}, cache.Get())

		// Reset refreshes the cache
		cache.Reset([]string{"zone3"})
		assert.False(t, cache.Expired(), "should not be expired after fresh reset")
		assert.Equal(t, []string{"zone3"}, cache.Get())
	})
}

func TestZoneCache_ThreadSafety(t *testing.T) {
	cache := NewZoneCache[[]int](time.Hour)

	var wg sync.WaitGroup
	const numWriters = 3
	const numReaders = 5
	const iterations = 100

	var validReads atomic.Int64

	// Writer goroutines
	for w := range numWriters {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			for i := range iterations {
				cache.Reset([]int{writerID, i})
			}
		}(w)
	}

	// Reader goroutines
	for range numReaders {
		wg.Go(func() {
			for range iterations {
				data := cache.Get()
				expired := cache.Expired()

				// Verify data consistency: if we got data, it should be valid
				if data != nil {
					assert.Len(t, data, 2, "cached slice should always have exactly 2 elements")
					assert.GreaterOrEqual(t, data[0], 0)
					assert.Less(t, data[0], numWriters)
					assert.GreaterOrEqual(t, data[1], 0)
					assert.Less(t, data[1], iterations)
					validReads.Add(1)
				}

				// Expired is a valid boolean - just verify it doesn't panic
				_ = expired
			}
		})
	}

	wg.Wait()

	// After all writes complete, cache should have valid final state
	finalData := cache.Get()
	assert.NotNil(t, finalData, "cache should have data after writes")
	assert.Len(t, finalData, 2, "final data should have 2 elements")
	assert.False(t, cache.Expired(), "cache should not be expired")
}
