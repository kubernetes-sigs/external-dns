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

package cloudflare

import (
	"sync"
	"time"
)

// defaultZonePlanCacheTTL is how long a cached "zone has a paid plan" answer
// is trusted. Zone plan changes are rare (billing action by the account
// owner), so 15 minutes is generous without introducing meaningful staleness.
const defaultZonePlanCacheTTL = 15 * time.Minute

// zonePlanCache memoises ZoneHasPaidPlan answers. It exists to prevent any
// call site that fires per DNS record change (notably the long-comment
// truncation path) from amplifying into O(n) ListZones+GetZone calls against
// the operator's Cloudflare account every sync cycle.
type zonePlanCache struct {
	mu      sync.Mutex
	entries map[string]zonePlanCacheEntry
	ttl     time.Duration
	now     func() time.Time // overridable for tests
}

type zonePlanCacheEntry struct {
	paid     bool
	cachedAt time.Time
}

func newZonePlanCache(ttl time.Duration) *zonePlanCache {
	return &zonePlanCache{
		entries: make(map[string]zonePlanCacheEntry),
		ttl:     ttl,
		now:     time.Now,
	}
}

// get returns the cached "paid" value for zone, or (_, false) if there is no
// unexpired entry. Expired entries are evicted lazily on lookup.
func (c *zonePlanCache) get(zone string) (bool, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[zone]
	if !ok {
		return false, false
	}
	if c.now().Sub(entry.cachedAt) > c.ttl {
		delete(c.entries, zone)
		return false, false
	}
	return entry.paid, true
}

// set stores a fresh answer for zone.
func (c *zonePlanCache) set(zone string, paid bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[zone] = zonePlanCacheEntry{
		paid:     paid,
		cachedAt: c.now(),
	}
}
