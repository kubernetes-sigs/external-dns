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

package txt

import (
	"fmt"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider/inmemory"
)

// BenchmarkTXTRegistryRecords measures Records() in the steady state of a reconcile loop:
// every record is owned by this registry and its ownership TXT record already exists.
func BenchmarkTXTRegistryRecords(b *testing.B) {
	for _, encrypted := range []bool{false, true} {
		for _, count := range []int{1000, 10000} {
			b.Run(fmt.Sprintf("encryption=%t/endpoints=%d", encrypted, count), func(b *testing.B) {
				r := newBenchmarkRegistry(b, count, encrypted)

				for b.Loop() {
					if _, err := r.Records(b.Context()); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

// newBenchmarkRegistry returns a registry whose provider holds count A and CNAME records
// together with the ownership TXT records the registry wrote for them.
func newBenchmarkRegistry(b *testing.B, count int, encrypted bool) *TXTRegistry {
	b.Helper()

	p := inmemory.NewInMemoryProvider()
	if err := p.CreateZone(testZone); err != nil {
		b.Fatal(err)
	}

	var aesKey []byte
	if encrypted {
		aesKey = []byte("passphrasewhichneedstobe32bytes!")
	}
	managed := []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME}
	r, err := newRegistry(p, "", "", "owner", 0, "", managed, []string{}, encrypted, aesKey, "")
	if err != nil {
		b.Fatal(err)
	}

	create := make([]*endpoint.Endpoint, 0, count)
	for i := range count {
		name := fmt.Sprintf("host-%d.%s", i, testZone)
		if i%2 == 0 {
			create = append(create, endpoint.NewEndpoint(name, endpoint.RecordTypeA, "192.0.2.1"))
		} else {
			create = append(create, endpoint.NewEndpoint(name, endpoint.RecordTypeCNAME, "lb.example.com"))
		}
	}
	if err := r.ApplyChanges(b.Context(), &plan.Changes{Create: create}); err != nil {
		b.Fatal(err)
	}

	// Fail fast if the fixture stops exercising the code under test.
	records, err := r.Records(b.Context())
	if err != nil {
		b.Fatal(err)
	}
	if len(records) != count {
		b.Fatalf("expected %d records, got %d", count, len(records))
	}
	for _, ep := range records {
		if ep.Labels[endpoint.OwnerLabelKey] != "owner" {
			b.Fatalf("record %s is not owned by the registry", ep.DNSName)
		}
		if _, ok := ep.GetProviderSpecificProperty(providerSpecificForceUpdate); ok {
			b.Fatalf("record %s unexpectedly needs a TXT update", ep.DNSName)
		}
	}
	return r
}
