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

package storage

import (
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

func TestInMemory(t *testing.T) {
	t.Run("TestAssign", testInMemoryAssign)
	t.Run("TestWaitForSync", testInMemoryWaitForSync)
	t.Run("TestOwnRecords", testInMemoryOwnRecords)
	t.Run("TestRecords", testInMemoryRecords)
	t.Run("TestNewInMemoryStorage", testNewInMemoryStorage)
}

func testInMemoryAssign(t *testing.T) {
	initRegistry := []endpoint.Endpoint{
		{
			DNSName: "foo.org",
			Target:  "foo-lb.org",
		},
		{
			DNSName: "bar.org",
			Target:  "bar-lb.org",
		},
		{
			DNSName: "baz.org",
			Target:  "baz-lb.org",
		},
		{
			DNSName: "qux.org",
			Target:  "qux-lb.org",
		},
	}
	for _, ti := range []struct {
		title         string
		owner         string
		cache         map[string]*endpoint.SharedEndpoint
		assign        []endpoint.Endpoint
		initRegistry  []endpoint.Endpoint
		expectedCache map[string]*endpoint.SharedEndpoint
	}{
		{
			title:         "empty cache, empty assign",
			owner:         "me",
			cache:         map[string]*endpoint.SharedEndpoint{},
			expectedCache: map[string]*endpoint.SharedEndpoint{},
			assign:        []endpoint.Endpoint{},
			initRegistry:  initRegistry,
		},
		{
			title: "non-empty cache, empty assign",
			owner: "me",
			cache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "foo-lb.org",
					},
					Owner: "",
				},
				"bar.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "bar-lb.org",
					},
					Owner: "",
				},
				"baz.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "baz.org",
						Target:  "baz-lb.org",
					},
					Owner: "",
				},
				"qux.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "qux.org",
						Target:  "qux-lb.org",
					},
					Owner: "",
				},
			},
			expectedCache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "foo-lb.org",
					},
					Owner: "",
				},
				"bar.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "bar-lb.org",
					},
					Owner: "",
				},
				"baz.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "baz.org",
						Target:  "baz-lb.org",
					},
					Owner: "",
				},
				"qux.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "qux.org",
						Target:  "qux-lb.org",
					},
					Owner: "",
				},
			},
			assign:       []endpoint.Endpoint{},
			initRegistry: initRegistry,
		},
		{
			title: "non-empty cache, with new assign",
			owner: "me",
			cache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "foo-lb.org",
					},
					Owner: "",
				},
				"bar.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "bar-lb.org",
					},
					Owner: "",
				},
				"baz.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "baz.org",
						Target:  "baz-lb.org",
					},
					Owner: "",
				},
				"qux.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "qux.org",
						Target:  "qux-lb.org",
					},
					Owner: "",
				},
			},
			expectedCache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "foo-lb.org",
					},
					Owner: "",
				},
				"bar.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "bar-lb.org",
					},
					Owner: "",
				},
				"baz.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "baz.org",
						Target:  "baz-lb.org",
					},
					Owner: "",
				},
				"qux.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "qux.org",
						Target:  "qux-lb.org",
					},
					Owner: "",
				},
				"new.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "new.org",
						Target:  "new-lb.org",
					},
					Owner: "me",
				},
				"another-new.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "another-new.org",
						Target:  "another-new-lb.org",
					},
					Owner: "me",
				},
			},
			assign: []endpoint.Endpoint{
				{
					DNSName: "new.org",
					Target:  "new-lb.org",
				},
				{
					DNSName: "another-new.org",
					Target:  "another-new-lb.org",
				},
			},
			initRegistry: initRegistry,
		},
		{
			title: "non-empty cache, with old and new assign",
			owner: "me",
			cache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "foo-lb.org",
					},
					Owner: "another",
				},
				"bar.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "bar-lb.org",
					},
					Owner: "",
				},
				"baz.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "baz.org",
						Target:  "baz-lb.org",
					},
					Owner: "",
				},
				"qux.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "qux.org",
						Target:  "qux-lb.org",
					},
					Owner: "",
				},
			},
			expectedCache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "foo-lb.org",
					},
					Owner: "me",
				},
				"bar.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "bar-lb.org",
					},
					Owner: "",
				},
				"baz.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "baz.org",
						Target:  "baz-lb.org",
					},
					Owner: "",
				},
				"qux.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "qux.org",
						Target:  "qux-lb.org",
					},
					Owner: "",
				},
				"new.org": {
					Endpoint: endpoint.Endpoint{
						DNSName: "new.org",
						Target:  "new-lb.org",
					},
					Owner: "me",
				},
			},
			assign: []endpoint.Endpoint{
				{
					DNSName: "new.org",
					Target:  "new-lb.org",
				},
				{
					DNSName: "foo.org",
					Target:  "foo-lb.org",
				},
			},
			initRegistry: initRegistry,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			registry := provider.NewInMemoryProvider()
			zone := "org"
			owner := "me"
			im, _ := NewInMemoryStorage(registry, owner, zone)
			im.cache = ti.cache
			registry.CreateZone(zone)
			registry.ApplyChanges(zone, &plan.Changes{
				Create: ti.initRegistry,
			})
			err := im.Assign(ti.assign)
			if err != nil {
				t.Error(err)
			}
			if len(im.cache) != len(ti.expectedCache) {
				t.Error("provided expected cache does not match the cache")
				return
			}
			for dns := range im.cache {
				if !testutils.SameSharedEndpoint(*im.cache[dns], *ti.expectedCache[dns]) {
					t.Error("provided expected cache does not match the cache")
				}
			}
		})
	}
}

func testInMemoryWaitForSync(t *testing.T) {
	registry := provider.NewInMemoryProvider()
	zone := "org"
	owner := "me"
	im, _ := NewInMemoryStorage(registry, owner, zone)
	err := im.WaitForSync()
	if err == nil {
		t.Errorf("should fail, because zone does not exist")
	}

	registry.CreateZone(zone)
	im.WaitForSync()
	if len(im.cache) != 0 {
		t.Errorf("cache should be empty!")
	}

	registry.ApplyChanges(zone, &plan.Changes{
		Create: []endpoint.Endpoint{
			{
				DNSName: "foo.org",
				Target:  "foo-lb.org",
			},
			{
				DNSName: "bar.org",
				Target:  "bar-lb.org",
			},
			{
				DNSName: "baz.org",
				Target:  "baz-lb.org",
			},
			{
				DNSName: "qux.org",
				Target:  "qux-lb.org",
			},
		},
	})

	expectedCache := []*endpoint.SharedEndpoint{
		{
			Endpoint: endpoint.Endpoint{
				DNSName: "foo.org",
				Target:  "foo-lb.org",
			},
			Owner: "",
		},
		{
			Endpoint: endpoint.Endpoint{
				DNSName: "bar.org",
				Target:  "bar-lb.org",
			},
			Owner: "",
		},
		{
			Endpoint: endpoint.Endpoint{
				DNSName: "baz.org",
				Target:  "baz-lb.org",
			},
			Owner: "",
		},
		{
			Endpoint: endpoint.Endpoint{
				DNSName: "qux.org",
				Target:  "qux-lb.org",
			},
			Owner: "",
		},
	}

	im.WaitForSync()

	flatCache := []*endpoint.SharedEndpoint{}
	for _, record := range im.cache {
		flatCache = append(flatCache, record)
	}

	if !testutils.SameSharedEndpoints(expectedCache, flatCache) {
		t.Errorf("cache is incorrectly populated")
	}

	im.Assign([]endpoint.Endpoint{
		{
			DNSName: "foo.org",
			Target:  "foo-lb.org",
		},
	})
	im.WaitForSync()
	expectedCache[0].Owner = "me"
	flatCache = []*endpoint.SharedEndpoint{}

	for _, record := range im.cache {
		flatCache = append(flatCache, record)
	}

	if !testutils.SameSharedEndpoints(expectedCache, flatCache) {
		t.Errorf("cache is incorrectly populated")
	}
}

func testInMemoryOwnRecords(t *testing.T) {
	for _, ti := range []struct {
		title    string
		owner    string
		cache    map[string]*endpoint.SharedEndpoint
		expected []endpoint.Endpoint
	}{
		{
			title:    "empty cache",
			owner:    "me",
			cache:    map[string]*endpoint.SharedEndpoint{},
			expected: []endpoint.Endpoint{},
		},
		{
			title: "non-empty cache, empty result",
			owner: "me",
			cache: map[string]*endpoint.SharedEndpoint{
				"bar.org": {
					Owner: "you",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "elb.com",
					},
				},
			},
			expected: []endpoint.Endpoint{},
		},
		{
			title: "non-empty cache, filter owned records",
			owner: "me",
			cache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				"bar.org": {
					Owner: "you",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "elb.com",
					},
				},
			},
			expected: []endpoint.Endpoint{
				{
					DNSName: "foo.org",
					Target:  "elb.com",
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			im := &InMemoryStorage{
				cache: ti.cache,
				owner: ti.owner,
			}
			if !testutils.SameEndpoints(im.OwnRecords(), ti.expected) {
				t.Errorf("unexpected result")
			}
		})
	}
}

func testInMemoryRecords(t *testing.T) {
	for _, ti := range []struct {
		title    string
		cache    map[string]*endpoint.SharedEndpoint
		expected []*endpoint.SharedEndpoint
	}{
		{
			title:    "empty cache",
			cache:    map[string]*endpoint.SharedEndpoint{},
			expected: []*endpoint.SharedEndpoint{},
		},
		{
			title: "non-empty cache",
			cache: map[string]*endpoint.SharedEndpoint{
				"foo.org": {
					Owner: "instance-id",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				"bar.org": {
					Owner: "another-id",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "elb.com",
					},
				},
			},
			expected: []*endpoint.SharedEndpoint{
				{
					Owner: "instance-id",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				{
					Owner: "another-id",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "elb.com",
					},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			im := &InMemoryStorage{
				cache: ti.cache,
			}
			if !testutils.SameSharedEndpoints(im.Records(), ti.expected) {
				t.Errorf("unexpected result")
			}
		})
	}
}

func testNewInMemoryStorage(t *testing.T) {
	var owner, zone string
	var dnsProvider provider.InMemoryProvider
	if _, err := NewInMemoryStorage(&dnsProvider, owner, zone); err == nil {
		t.Errorf("should fail when owner/zone is empty")
	}
	owner = "test-1"
	zone = ""
	if _, err := NewInMemoryStorage(&dnsProvider, owner, zone); err == nil {
		t.Errorf("should fail when owner/zone is empty")
	}
	owner = ""
	zone = "zone-1"
	if _, err := NewInMemoryStorage(&dnsProvider, owner, zone); err == nil {
		t.Errorf("should fail when owner/zone is empty")
	}
	owner = "test-1"
	zone = "zone-1"
	storage, err := NewInMemoryStorage(&dnsProvider, owner, zone)
	if err != nil {
		t.Error(err)
	}
	if storage.zone != zone || storage.owner != owner || storage.cache == nil {
		t.Errorf("incorrectly initialized in memory storage dnsProvider")
	}
	if _, ok := storage.registry.(*provider.InMemoryProvider); !ok {
		t.Errorf("incorrect dns provider is used for registry")
	}
}
