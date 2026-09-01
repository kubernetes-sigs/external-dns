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
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

func TestZonePlanCache_HitAndMiss(t *testing.T) {
	c := newZonePlanCache(time.Minute)

	if _, ok := c.get("foo.com"); ok {
		t.Fatalf("expected miss on empty cache")
	}

	c.set("foo.com", true)

	paid, ok := c.get("foo.com")
	if !ok || !paid {
		t.Fatalf("expected hit with paid=true, got paid=%v ok=%v", paid, ok)
	}
}

func TestZonePlanCache_TTLEvicts(t *testing.T) {
	var now atomic.Pointer[time.Time]
	t0 := time.Unix(0, 0)
	now.Store(&t0)

	c := newZonePlanCache(time.Minute)
	c.now = func() time.Time { return *now.Load() }

	c.set("foo.com", true)
	if _, ok := c.get("foo.com"); !ok {
		t.Fatalf("expected hit immediately after set")
	}

	// Advance past TTL.
	later := t0.Add(2 * time.Minute)
	now.Store(&later)

	if _, ok := c.get("foo.com"); ok {
		t.Fatalf("expected miss after TTL expiry")
	}
}

// TestZoneHasPaidPlan_Cached verifies the provider answers repeat calls from
// cache rather than hitting ZoneIDByName + GetZone on every call. This is the
// fix for the per-change API amplification flagged in #6391.
func TestZoneHasPaidPlan_Cached(t *testing.T) {
	client := NewMockCloudFlareClient()
	cfProvider := &CloudFlareProvider{
		Client:        client,
		domainFilter:  endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter:  provider.NewZoneIDFilter([]string{""}),
		zonePlanCache: newZonePlanCache(time.Minute),
	}

	// First call populates the cache; subsequent calls must not re-fetch.
	assert.True(t, cfProvider.ZoneHasPaidPlan("subdomain.bar.com"))

	// Break the API: if cache isn't consulted, ZoneHasPaidPlan falls through
	// to ZoneIDByName/GetZone and returns false on error.
	client.getZoneError = errors.New("should not be called — cached answer expected")

	for range 10 {
		assert.True(t, cfProvider.ZoneHasPaidPlan("subdomain.bar.com"),
			"expected cached paid=true; cache miss re-hit the broken API")
	}

	// Different hostname in a different zone still uses the API.
	assert.False(t, cfProvider.ZoneHasPaidPlan("subdomain.foo.com"),
		"expected miss path to surface the API error as paid=false")
}

// TestZoneHasPaidPlan_NilCache confirms the provider is resilient to a nil
// cache (e.g. hand-constructed in tests without calling newProvider).
func TestZoneHasPaidPlan_NilCache(t *testing.T) {
	client := NewMockCloudFlareClient()
	cfProvider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	assert.True(t, cfProvider.ZoneHasPaidPlan("subdomain.bar.com"))
	assert.False(t, cfProvider.ZoneHasPaidPlan("subdomain.foo.com"))
}
