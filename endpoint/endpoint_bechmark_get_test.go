package endpoint

import (
	"fmt"
	"testing"
)

func BenchmarkEndpointGet100With3Properties(b *testing.B) {
	endpoints := buildEndpointsWithProperties(100, 3)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificProperty("prop1")
			e.GetProviderSpecificProperty("prop2")
			e.GetProviderSpecificProperty("prop3")
		}
	}
}

func BenchmarkEndpointGet100With3PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithProperties(100, 3)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificPropertyHashMap("prop1")
			e.GetProviderSpecificPropertyHashMap("prop2")
			e.GetProviderSpecificPropertyHashMap("prop3")
		}
	}
}

func BenchmarkEndpointGet1000With3Properties(b *testing.B) {
	endpoints := buildEndpointsWithProperties(1000, 3)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificProperty("prop1")
			e.GetProviderSpecificProperty("prop2")
			e.GetProviderSpecificProperty("prop3")
		}
	}
}

func BenchmarkEndpointGet1000With3PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithProperties(1000, 3)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificPropertyHashMap("prop1")
			e.GetProviderSpecificPropertyHashMap("prop2")
			e.GetProviderSpecificPropertyHashMap("prop3")
		}
	}
}

func BenchmarkEndpointGet100With5Properties(b *testing.B) {
	endpoints := buildEndpointsWithProperties(100, 5)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificProperty("prop1")
			e.GetProviderSpecificProperty("prop2")
			e.GetProviderSpecificProperty("prop3")
			e.GetProviderSpecificProperty("prop4")
			e.GetProviderSpecificProperty("prop5")
		}
	}
}

func BenchmarkEndpointGet100With5PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithProperties(100, 5)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificPropertyHashMap("prop1")
			e.GetProviderSpecificPropertyHashMap("prop2")
			e.GetProviderSpecificPropertyHashMap("prop3")
			e.GetProviderSpecificPropertyHashMap("prop4")
			e.GetProviderSpecificPropertyHashMap("prop5")
		}
	}
}

func BenchmarkEndpointGet1000With5Properties(b *testing.B) {
	endpoints := buildEndpointsWithProperties(1000, 5)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificProperty("prop1")
			e.GetProviderSpecificProperty("prop2")
			e.GetProviderSpecificProperty("prop3")
			e.GetProviderSpecificProperty("prop4")
			e.GetProviderSpecificProperty("prop5")
		}
	}
}

func BenchmarkEndpointGet1000With5PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithProperties(1000, 5)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificPropertyHashMap("prop1")
			e.GetProviderSpecificPropertyHashMap("prop2")
			e.GetProviderSpecificPropertyHashMap("prop3")
			e.GetProviderSpecificPropertyHashMap("prop4")
			e.GetProviderSpecificPropertyHashMap("prop5")
		}
	}
}

func BenchmarkEndpointGet100With10Properties(b *testing.B) {
	endpoints := buildEndpointsWithProperties(100, 10)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificProperty("prop1")
			e.GetProviderSpecificProperty("prop2")
			e.GetProviderSpecificProperty("prop3")
			e.GetProviderSpecificProperty("prop4")
			e.GetProviderSpecificProperty("prop5")
			e.GetProviderSpecificProperty("prop6")
			e.GetProviderSpecificProperty("prop7")
			e.GetProviderSpecificProperty("prop8")
			e.GetProviderSpecificProperty("prop9")
			e.GetProviderSpecificProperty("prop10")
		}
	}
}

func BenchmarkEndpointGet100With10PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithProperties(100, 10)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificPropertyHashMap("prop1")
			e.GetProviderSpecificPropertyHashMap("prop2")
			e.GetProviderSpecificPropertyHashMap("prop3")
			e.GetProviderSpecificPropertyHashMap("prop4")
			e.GetProviderSpecificPropertyHashMap("prop5")
			e.GetProviderSpecificPropertyHashMap("prop6")
			e.GetProviderSpecificPropertyHashMap("prop7")
			e.GetProviderSpecificPropertyHashMap("prop8")
			e.GetProviderSpecificPropertyHashMap("prop9")
			e.GetProviderSpecificPropertyHashMap("prop10")
		}
	}
}

func BenchmarkEndpointGet1000With10Properties(b *testing.B) {
	endpoints := buildEndpointsWithProperties(1000, 10)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificProperty("prop1")
			e.GetProviderSpecificProperty("prop2")
			e.GetProviderSpecificProperty("prop3")
			e.GetProviderSpecificProperty("prop4")
			e.GetProviderSpecificProperty("prop5")
			e.GetProviderSpecificProperty("prop6")
			e.GetProviderSpecificProperty("prop7")
			e.GetProviderSpecificProperty("prop8")
			e.GetProviderSpecificProperty("prop9")
			e.GetProviderSpecificProperty("prop10")
		}
	}
}

func BenchmarkEndpointGet1000With10PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithProperties(1000, 10)

	for b.Loop() {
		for _, e := range endpoints {
			e.GetProviderSpecificPropertyHashMap("prop1")
			e.GetProviderSpecificPropertyHashMap("prop2")
			e.GetProviderSpecificPropertyHashMap("prop3")
			e.GetProviderSpecificPropertyHashMap("prop4")
			e.GetProviderSpecificPropertyHashMap("prop5")
			e.GetProviderSpecificPropertyHashMap("prop6")
			e.GetProviderSpecificPropertyHashMap("prop7")
			e.GetProviderSpecificPropertyHashMap("prop8")
			e.GetProviderSpecificPropertyHashMap("prop9")
			e.GetProviderSpecificPropertyHashMap("prop10")
		}
	}
}

func buildEndpointsWithProperties(nEndpoints int, properties int) []Endpoint {
	endpoints := make([]Endpoint, nEndpoints)
	for i := 0; i < nEndpoints; i++ {
		endpoints[i] = *NewEndpoint(fmt.Sprintf("index-%d.example.com", i), RecordTypeA)
		for n := 0; n < properties; n++ {
			endpoints[i].SetProviderSpecificProperty(fmt.Sprintf("prop%d", n+1), fmt.Sprintf("value%d", n+1))
		}
	}
	return endpoints
}
