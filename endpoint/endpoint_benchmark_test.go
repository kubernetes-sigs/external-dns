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

package endpoint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	providerSpecificKeys = []string{
		"alias",
		"provider/target-hosted-zone",
		"provider/evaluate-target-health",
		"provider/weight",
		"provider/region",
		"provider/failover",
		"provider/geolocation-continent-code",
		"provider/geolocation-country-code",
		"provider/geolocation-subdivision-code",
		"provider/geoproximity-region",
		"provider/geoproximity-bias",
		"provider/geoproximity-coordinates",
		"provider/geoproximity-local-zone-group",
		"provider/multi-value-answer",
		"provider/health-check-id",
		"same-zone",
	}
)

// TestEndpointGeneration validates that generateBenchmarkEndpoints
// creates correct data for both slice implementations.
func TestEndpointGeneration(t *testing.T) {
	for _, setProps := range []int{0, 1, 3, 5, 16} {
		t.Run(fmt.Sprintf("set=%d", setProps), func(t *testing.T) {
			endpoints := generateBenchmarkEndpoints(10, setProps)
			assert.Len(t, endpoints, 10)
			for _, ep := range endpoints {
				assert.Len(t, ep.ProviderSpecific, setProps)
			}
		})
	}
}

// BenchmarkProviderSpecificRealisticAccess simulates realistic provider behavior:
// The provider checks ALL its supported properties on each endpoint,
// even though only a few (setProps) are actually configured.
func BenchmarkProviderSpecificRandomAccess(b *testing.B) {
	// setProps: how many properties are actually set on the endpoint
	// The provider will still check all 16 keys
	setPropsOptions := []int{0, 1, 5, 9, 16}
	endpointCounts := []int{100, 1000, 10000, 50000, 100000, 200000}

	keys := []string{
		"provider/weight",
		"nonexistent",
		"provider/geoproximity-region",
		"same-zone",
	}

	for _, setProps := range setPropsOptions {
		for _, epCount := range endpointCounts {
			endpoints := generateBenchmarkEndpoints(epCount, setProps)
			b.Run(fmt.Sprintf("slice/set=%d/endpoints=%d", setProps, epCount), func(b *testing.B) {
				for b.Loop() {
					for _, ep := range endpoints {
						// Provider checks random supported properties
						for _, key := range keys {
							ep.GetProviderSpecificProperty(key)
						}
					}
				}
			})
		}
	}
}

func BenchmarkProviderSpecificDelete(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000, 10000, 50000}

	keys := []string{
		"provider/weight",
		"nonexistent",
	}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			b.Run(fmt.Sprintf("slice/props=%d/endpoints=%d", propCount, epCount), func(b *testing.B) {
				template := generateBenchmarkEndpoints(epCount, propCount)
				b.ResetTimer()
				for b.Loop() {
					// Shallow copy is enough if we only care about the slice structure
					endpoints := make([]*Endpoint, len(template))
					copy(endpoints, template)
					for _, ep := range endpoints {
						for _, key := range keys {
							ep.DeleteProviderSpecificProperty(key)
						}
					}
				}
			})
		}
	}
}

// generateBenchmarkEndpoints creates endpoints with realistic AWS provider-specific properties.
// setPropsCount determines how many of the providerSpecificKeys are actually set on each endpoint.
func generateBenchmarkEndpoints(count, setPropsCount int) []*Endpoint {
	endpoints := make([]*Endpoint, count)
	for i := range count {
		ep := &Endpoint{
			DNSName:    fmt.Sprintf("endpoint-%d.example.com", i),
			RecordType: RecordTypeA,
			Targets:    Targets{fmt.Sprintf("192.0.2.%d", i%256)},
			RecordTTL:  TTL(300),
			Labels:     NewLabels(),
		}

		// Set only the first setPropsCount properties
		if setPropsCount > 0 {
			ep.ProviderSpecific = make(ProviderSpecific, setPropsCount)
			for j := range setPropsCount {
				key := providerSpecificKeys[j%len(providerSpecificKeys)]
				ep.ProviderSpecific[j] = ProviderSpecificProperty{
					Name:  key,
					Value: fmt.Sprintf("value-%d", j),
				}
			}
		}

		endpoints[i] = ep
	}
	return endpoints
}
