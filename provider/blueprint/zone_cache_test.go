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
	"testing"
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

func TestZoneCache_ThreadSafety(t *testing.T) {
	cache := NewZoneCache[[]int](time.Hour)

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
