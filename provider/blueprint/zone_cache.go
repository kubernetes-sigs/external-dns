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
	"sync"
	"time"
)

// ZoneCache is a generic cache for DNS zones with TTL-based expiration.
// It can store any type of zone data and provides thread-safe access.
type ZoneCache[T any] struct {
	mu       sync.RWMutex
	age      time.Time
	duration time.Duration
	data     T
	isEmpty  func(T) bool
}

// NewZoneCache creates a new ZoneCache with the specified TTL duration.
// The isEmpty function is used to determine if the cached data is empty
// (which forces a refresh on first access).
func NewZoneCache[T any](duration time.Duration, isEmpty func(T) bool) *ZoneCache[T] {
	return &ZoneCache[T]{
		duration: duration,
		isEmpty:  isEmpty,
	}
}

// NewSliceZoneCache creates a ZoneCache for slice types.
func NewSliceZoneCache[T any](duration time.Duration) *ZoneCache[[]T] {
	return NewZoneCache(duration, func(data []T) bool {
		return len(data) == 0
	})
}

// NewMapZoneCache creates a ZoneCache for map types.
func NewMapZoneCache[K comparable, V any](duration time.Duration) *ZoneCache[map[K]V] {
	return NewZoneCache(duration, func(data map[K]V) bool {
		return len(data) == 0
	})
}

// Get returns the cached data. Returns the zero value if cache is empty.
func (c *ZoneCache[T]) Get() T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}

// Reset updates the cached data and refreshes the age timestamp.
// Only updates if caching is enabled (duration > 0).
func (c *ZoneCache[T]) Reset(data T) {
	if c.duration <= 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = data
	c.age = time.Now()
}

// Expired returns true if the cache has expired or is empty.
func (c *ZoneCache[T]) Expired() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.isEmpty(c.data) || time.Since(c.age) > c.duration
}

// Duration returns the cache TTL duration.
func (c *ZoneCache[T]) Duration() time.Duration {
	return c.duration
}

// ZoneFetcher is a function type that fetches zones from a provider.
// It's used by CachedZoneProvider to fetch zones when the cache is expired.
type ZoneFetcher[T any] func(ctx context.Context) (T, error)
