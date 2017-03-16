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

	"github.com/kubernetes-incubator/external-dns/dnsprovider"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// initInMemoryDNSProvider initialize the state for in memory dns dnsprovider
func initInMemoryDNSProvider() (*dnsprovider.InMemoryProvider, string) {
	zone := "org"
	registry := dnsprovider.NewInMemoryProvider()
	registry.CreateZone(zone)
	registry.ApplyChanges(zone, &plan.Changes{
		Create: []endpoint.Endpoint{
			endpoint.Endpoint{
				DNSName: "foo.org",
				Target:  "foo-lb.org",
			},
			endpoint.Endpoint{
				DNSName: "bar.org",
				Target:  "bar-lb.org",
			},
			endpoint.Endpoint{
				DNSName: "baz.org",
				Target:  "baz-lb.org",
			},
			endpoint.Endpoint{
				DNSName: "qux.org",
				Target:  "qux-lb.org",
			},
		},
	})
	return registry, zone
}

func TestUpdatedCache(t *testing.T) {
	for _, ti := range []struct {
		title        string
		records      []endpoint.Endpoint
		cacheRecords []*SharedEndpoint
		expected     []*SharedEndpoint
	}{
		{
			title:        "all empty",
			records:      []endpoint.Endpoint{},
			cacheRecords: []*SharedEndpoint{},
			expected:     []*SharedEndpoint{},
		},
		{
			title:   "no records, should produce empty cache",
			records: []endpoint.Endpoint{},
			cacheRecords: []*SharedEndpoint{
				&SharedEndpoint{},
			},
			expected: []*SharedEndpoint{},
		},
		{
			title: "new records, empty cache",
			records: []endpoint.Endpoint{
				endpoint.Endpoint{
					DNSName: "foo.org",
					Target:  "elb.com",
				},
				endpoint.Endpoint{
					DNSName: "bar.org",
					Target:  "alb.com",
				},
			},
			cacheRecords: []*SharedEndpoint{
				&SharedEndpoint{},
			},
			expected: []*SharedEndpoint{
				&SharedEndpoint{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				&SharedEndpoint{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "alb.com",
					},
				},
			},
		},
		{
			title: "new records, non-empty cache",
			records: []endpoint.Endpoint{
				endpoint.Endpoint{
					DNSName: "foo.org",
					Target:  "elb.com",
				},
				endpoint.Endpoint{
					DNSName: "bar.org",
					Target:  "alb.com",
				},
				endpoint.Endpoint{
					DNSName: "owned.org",
					Target:  "8.8.8.8",
				},
			},
			cacheRecords: []*SharedEndpoint{
				&SharedEndpoint{
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "owned.org",
						Target:  "8.8.8.8",
					},
				},
				&SharedEndpoint{
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "to-be-deleted.org",
						Target:  "52.53.54.55",
					},
				},
			},
			expected: []*SharedEndpoint{
				&SharedEndpoint{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "elb.com",
					},
				},
				&SharedEndpoint{
					Owner: "",
					Endpoint: endpoint.Endpoint{
						DNSName: "bar.org",
						Target:  "alb.com",
					},
				},
				&SharedEndpoint{
					Owner: "me",
					Endpoint: endpoint.Endpoint{
						DNSName: "owned.org",
						Target:  "8.8.8.8",
					},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			if !sameSharedEndpoints(updatedCache(ti.records, ti.cacheRecords), ti.expected) {
				t.Errorf("incorrect result produced by updatedCache")
			}
		})
	}
}

//helper functions
func sameSharedEndpoints(a, b []*SharedEndpoint) bool {
	if len(a) != len(b) {
		return false
	}
	for _, recordA := range a {
		found := false
		for _, recordB := range b {
			if recordA.DNSName == recordB.DNSName && recordA.Target == recordB.Target && recordA.Owner == recordB.Owner {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	for _, recordB := range b {
		found := false
		for _, recordA := range a {
			if recordB.DNSName == recordA.DNSName && recordB.Target == recordA.Target && recordA.Owner == recordB.Owner {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
