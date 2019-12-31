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

package source

import (
	"net"
	"regexp"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

func generateTestEndpoints() []*endpoint.Endpoint {
	sc, _ := NewFakeSource("")

	endpoints, _ := sc.Endpoints()

	return endpoints
}

func TestFakeSourceReturnsTenEndpoints(t *testing.T) {
	endpoints := generateTestEndpoints()

	count := len(endpoints)

	if count != 10 {
		t.Error(count)
	}
}

func TestFakeEndpointsBelongToDomain(t *testing.T) {
	validRecord := regexp.MustCompile(`^[a-z]{4}\.example\.com$`)

	endpoints := generateTestEndpoints()

	for _, e := range endpoints {
		valid := validRecord.MatchString(e.DNSName)

		if !valid {
			t.Error(e.DNSName)
		}
	}
}

func TestFakeEndpointsResolveToIPAddresses(t *testing.T) {
	endpoints := generateTestEndpoints()

	for _, e := range endpoints {
		ip := net.ParseIP(e.Targets[0])

		if ip == nil {
			t.Error(e)
		}
	}
}

// Validate that FakeSource is a source
var _ Source = &fakeSource{}
