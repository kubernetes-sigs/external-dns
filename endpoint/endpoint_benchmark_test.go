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

package endpoint

import (
	"fmt"
	"testing"
)

// benchmarkEndpointConfig holds configuration for generating benchmark test endpoints.
type benchmarkEndpointConfig struct {
	// Count specifies how many endpoints to create.
	Count int
	// ProviderSpecificCount specifies how many provider-specific properties each endpoint should have.
	ProviderSpecificCount int
}

// generateBenchmarkEndpoints creates test endpoints based on the provided configuration.
// Populates both ProviderSpecific (slice) and ProviderSpecificM (map) for comparison benchmarks.
func generateBenchmarkEndpoints(cfg benchmarkEndpointConfig) []*Endpoint {
	endpoints := make([]*Endpoint, cfg.Count)
	for i := range cfg.Count {
		endpoints[i] = &Endpoint{
			DNSName:           fmt.Sprintf("endpoint-%d.example.com", i),
			RecordType:        RecordTypeA,
			Targets:           Targets{fmt.Sprintf("192.0.2.%d", i%256)},
			RecordTTL:         TTL(300),
			Labels:            NewLabels(),
			ProviderSpecific:  generateBenchmarkProviderSpecificSlice(cfg.ProviderSpecificCount),
			ProviderSpecificM: generateBenchmarkProviderSpecificMap(cfg.ProviderSpecificCount),
		}
	}
	return endpoints
}

const benchmarkPropertyPrefix = "external-dns.alpha.kubernetes.io/"

// generateBenchmarkProviderSpecificSlice creates a slice of ProviderSpecificProperty with the given count.
func generateBenchmarkProviderSpecificSlice(count int) ProviderSpecific {
	if count <= 0 {
		return nil
	}
	props := make(ProviderSpecific, count)
	for i := range count {
		props[i] = ProviderSpecificProperty{
			Name:  fmt.Sprintf("%sproperty-%d", benchmarkPropertyPrefix, i),
			Value: fmt.Sprintf("value-%d", i),
		}
	}
	return props
}

// generateBenchmarkProviderSpecificMap creates a map of provider-specific properties with the given count.
func generateBenchmarkProviderSpecificMap(count int) ProviderSpecificMap {
	if count <= 0 {
		return nil
	}
	props := make(ProviderSpecificMap, count)
	for i := range count {
		props[fmt.Sprintf("%sproperty-%d", benchmarkPropertyPrefix, i)] = fmt.Sprintf("value-%d", i)
	}
	return props
}

func BenchmarkGetProviderSpecificProperty(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			name := fmt.Sprintf("props=%d/endpoints=%d", propCount, epCount)
			b.Run(name, func(b *testing.B) {
				endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
					Count:                 epCount,
					ProviderSpecificCount: propCount,
				})
				// Search for a property that exists (middle of the list) or doesn't exist (if no props)
				searchKey := benchmarkPropertyPrefix + "property-2"
				if propCount == 0 {
					searchKey = benchmarkPropertyPrefix + "nonexistent"
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for _, ep := range endpoints {
						ep.GetProviderSpecificProperty(searchKey)
					}
				}
			})
		}
	}
}

func BenchmarkDeleteProviderSpecificProperty(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			name := fmt.Sprintf("props=%d/endpoints=%d", propCount, epCount)
			b.Run(name, func(b *testing.B) {
				// Delete a property that exists (middle of the list) or doesn't exist (if no props)
				deleteKey := benchmarkPropertyPrefix + "property-2"
				if propCount == 0 {
					deleteKey = benchmarkPropertyPrefix + "nonexistent"
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
						Count:                 epCount,
						ProviderSpecificCount: propCount,
					})
					b.StartTimer()

					for _, ep := range endpoints {
						ep.DeleteProviderSpecificProperty(deleteKey)
					}
				}
			})
		}
	}
}

// Map-based benchmarks for comparison

func BenchmarkGetProviderSpecificPropertyM(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			name := fmt.Sprintf("props=%d/endpoints=%d", propCount, epCount)
			b.Run(name, func(b *testing.B) {
				endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
					Count:                 epCount,
					ProviderSpecificCount: propCount,
				})
				searchKey := benchmarkPropertyPrefix + "property-2"
				if propCount == 0 {
					searchKey = benchmarkPropertyPrefix + "nonexistent"
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for _, ep := range endpoints {
						ep.GetProviderSpecificPropertyM(searchKey)
					}
				}
			})
		}
	}
}

func BenchmarkDeleteProviderSpecificPropertyM(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			name := fmt.Sprintf("props=%d/endpoints=%d", propCount, epCount)
			b.Run(name, func(b *testing.B) {
				deleteKey := benchmarkPropertyPrefix + "property-2"
				if propCount == 0 {
					deleteKey = benchmarkPropertyPrefix + "nonexistent"
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
						Count:                 epCount,
						ProviderSpecificCount: propCount,
					})
					b.StartTimer()

					for _, ep := range endpoints {
						ep.DeleteProviderSpecificPropertyM(deleteKey)
					}
				}
			})
		}
	}
}

// Comparison benchmarks - slice vs map side by side

func BenchmarkProviderSpecificGetComparison(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000, 10000, 50000}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
				Count:                 epCount,
				ProviderSpecificCount: propCount,
			})
			searchKey := benchmarkPropertyPrefix + "property-2"
			if propCount == 0 {
				searchKey = benchmarkPropertyPrefix + "nonexistent"
			}

			b.Run(fmt.Sprintf("slice/props=%d/endpoints=%d", propCount, epCount), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, ep := range endpoints {
						ep.GetProviderSpecificProperty(searchKey)
					}
				}
			})

			b.Run(fmt.Sprintf("map/props=%d/endpoints=%d", propCount, epCount), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, ep := range endpoints {
						ep.GetProviderSpecificPropertyM(searchKey)
					}
				}
			})
		}
	}
}

func BenchmarkProviderSpecificDeleteComparison(b *testing.B) {
	propertyCounts := []int{0, 5, 10}
	endpointCounts := []int{100, 300, 1000, 10000, 50000}

	for _, propCount := range propertyCounts {
		for _, epCount := range endpointCounts {
			deleteKey := benchmarkPropertyPrefix + "property-2"
			if propCount == 0 {
				deleteKey = benchmarkPropertyPrefix + "nonexistent"
			}

			b.Run(fmt.Sprintf("slice/props=%d/endpoints=%d", propCount, epCount), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
						Count:                 epCount,
						ProviderSpecificCount: propCount,
					})
					b.StartTimer()

					for _, ep := range endpoints {
						ep.DeleteProviderSpecificProperty(deleteKey)
					}
				}
			})

			b.Run(fmt.Sprintf("map/props=%d/endpoints=%d", propCount, epCount), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					endpoints := generateBenchmarkEndpoints(benchmarkEndpointConfig{
						Count:                 epCount,
						ProviderSpecificCount: propCount,
					})
					b.StartTimer()

					for _, ep := range endpoints {
						ep.DeleteProviderSpecificPropertyM(deleteKey)
					}
				}
			})
		}
	}
}

// BenchmarkStringCompareVsMapLookup demonstrates why slice linear search
// can be faster than map lookup for small slices with long keys.
func BenchmarkStringCompareVsMapLookup(b *testing.B) {
	// Setup: 3 properties set, searching for various keys
	slice := ProviderSpecific{
		{Name: "alias", Value: "true"},
		{Name: "aws/target-hosted-zone", Value: "Z123"},
		{Name: "aws/evaluate-target-health", Value: "true"},
	}
	m := ProviderSpecificMap{
		"alias":                       "true",
		"aws/target-hosted-zone":      "Z123",
		"aws/evaluate-target-health": "true",
	}

	// Keys to search - mix of hits and misses with varying prefix lengths
	searchKeys := []string{
		"alias",                             // hit, short key
		"aws/weight",                        // miss, fails at position 4 (after "aws/")
		"aws/geolocation-continent-code",    // miss, fails at position 4
		"aws/geolocation-subdivision-code",  // miss, fails at position 4, long key (33 chars)
		"aws/evaluate-target-health",        // hit, long key
		"same-zone",                         // miss, fails at position 0
	}

	b.Run("slice-linear-search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, key := range searchKeys {
				for _, prop := range slice {
					if prop.Name == key {
						_ = prop.Value
						break
					}
				}
			}
		}
	})

	b.Run("map-lookup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, key := range searchKeys {
				_ = m[key]
			}
		}
	})
}

// TestRealisticEndpointGeneration validates that generateRealisticEndpoints
// creates identical data for both slice and map implementations.
func TestRealisticEndpointGeneration(t *testing.T) {
	for _, setProps := range []int{0, 1, 3, 5, 16} {
		t.Run(fmt.Sprintf("set=%d", setProps), func(t *testing.T) {
			endpoints := generateRealisticEndpoints(10, setProps)

			for i, ep := range endpoints {
				// Verify slice and map have same number of entries
				if setProps == 0 {
					if ep.ProviderSpecific != nil {
						t.Errorf("endpoint %d: expected nil slice, got %d elements", i, len(ep.ProviderSpecific))
					}
					if ep.ProviderSpecificM != nil {
						t.Errorf("endpoint %d: expected nil map, got %d elements", i, len(ep.ProviderSpecificM))
					}
					continue
				}

				if len(ep.ProviderSpecific) != setProps {
					t.Errorf("endpoint %d: slice has %d elements, expected %d", i, len(ep.ProviderSpecific), setProps)
				}
				if len(ep.ProviderSpecificM) != setProps {
					t.Errorf("endpoint %d: map has %d elements, expected %d", i, len(ep.ProviderSpecificM), setProps)
				}

				// Verify all 16 keys return same result from both implementations
				for _, key := range awsProviderSpecificKeys {
					sliceVal, sliceOk := ep.GetProviderSpecificProperty(key)
					mapVal, mapOk := ep.GetProviderSpecificPropertyM(key)

					if sliceOk != mapOk {
						t.Errorf("endpoint %d, key %q: slice found=%v, map found=%v", i, key, sliceOk, mapOk)
					}
					if sliceVal != mapVal {
						t.Errorf("endpoint %d, key %q: slice value=%q, map value=%q", i, key, sliceVal, mapVal)
					}
				}
			}
		})
	}
}

// Realistic AWS provider-specific property keys
var awsProviderSpecificKeys = []string{
	"alias",
	"aws/target-hosted-zone",
	"aws/evaluate-target-health",
	"aws/weight",
	"aws/region",
	"aws/failover",
	"aws/geolocation-continent-code",
	"aws/geolocation-country-code",
	"aws/geolocation-subdivision-code",
	"aws/geoproximity-region",
	"aws/geoproximity-bias",
	"aws/geoproximity-coordinates",
	"aws/geoproximity-local-zone-group",
	"aws/multi-value-answer",
	"aws/health-check-id",
	"same-zone",
}

// generateRealisticEndpoints creates endpoints with realistic AWS provider-specific properties.
// setPropsCount determines how many of the awsProviderSpecificKeys are actually set on each endpoint.
func generateRealisticEndpoints(count, setPropsCount int) []*Endpoint {
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
			ep.ProviderSpecificM = make(ProviderSpecificMap, setPropsCount)
			for j := range setPropsCount {
				key := awsProviderSpecificKeys[j%len(awsProviderSpecificKeys)]
				ep.ProviderSpecific[j] = ProviderSpecificProperty{
					Name:  key,
					Value: fmt.Sprintf("value-%d", j),
				}
				ep.ProviderSpecificM[key] = fmt.Sprintf("value-%d", j)
			}
		}

		endpoints[i] = ep
	}
	return endpoints
}

// BenchmarkProviderSpecificRealisticAccess simulates realistic provider behavior:
// The provider checks ALL its supported properties on each endpoint,
// even though only a few (setProps) are actually configured.
func BenchmarkProviderSpecificRealisticAccess(b *testing.B) {
	// setProps: how many properties are actually set on the endpoint
	// The provider will still check all 16 AWS keys
	setPropsOptions := []int{0, 1, 2, 3, 5, 9, 16}
	endpointCounts := []int{100, 1000, 10000, 50000, 100000, 200000}

	for _, setProps := range setPropsOptions {
		for _, epCount := range endpointCounts {
			endpoints := generateRealisticEndpoints(epCount, setProps)

			b.Run(fmt.Sprintf("slice/set=%d/endpoints=%d", setProps, epCount), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, ep := range endpoints {
						// Provider checks ALL supported properties
						for _, key := range awsProviderSpecificKeys {
							ep.GetProviderSpecificProperty(key)
						}
					}
				}
			})

			b.Run(fmt.Sprintf("map/set=%d/endpoints=%d", setProps, epCount), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, ep := range endpoints {
						// Provider checks ALL supported properties
						for _, key := range awsProviderSpecificKeys {
							ep.GetProviderSpecificPropertyM(key)
						}
					}
				}
			})
		}
	}
}
