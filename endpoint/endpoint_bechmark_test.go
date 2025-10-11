package endpoint

import (
	"fmt"
	"testing"
)

func BenchmarkEndpoint100With3Properties(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(100)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificProperty("prop1", "value1")
			e.SetProviderSpecificProperty("prop2", "value2")
			e.SetProviderSpecificProperty("prop3", "value3")
		}
	}
}

func BenchmarkEndpoint1000With3Properties(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(1000)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificProperty("prop1", "value1")
			e.SetProviderSpecificProperty("prop2", "value2")
			e.SetProviderSpecificProperty("prop3", "value3")
		}
	}
}

func BenchmarkEndpoint100With3PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(100)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificPropertyHashMap("prop1", "value1")
			e.SetProviderSpecificPropertyHashMap("prop2", "value2")
			e.SetProviderSpecificPropertyHashMap("prop3", "value3")
		}
	}
}

func BenchmarkEndpoint1000With3PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(1000)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificPropertyHashMap("prop1", "value1")
			e.SetProviderSpecificPropertyHashMap("prop2", "value2")
			e.SetProviderSpecificPropertyHashMap("prop3", "value3")
		}
	}
}

func BenchmarkEndpoint100With5Properties(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(100)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificProperty("prop1", "value1")
			e.SetProviderSpecificProperty("prop2", "value2")
			e.SetProviderSpecificProperty("prop3", "value3")
			e.SetProviderSpecificProperty("prop4", "value3")
			e.SetProviderSpecificProperty("prop5", "value3")
		}
	}
}

func BenchmarkEndpoint1000With5Properties(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(1000)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificProperty("prop1", "value1")
			e.SetProviderSpecificProperty("prop2", "value2")
			e.SetProviderSpecificProperty("prop3", "value3")
			e.SetProviderSpecificProperty("prop4", "value3")
			e.SetProviderSpecificProperty("prop5", "value3")
		}
	}
}

func BenchmarkEndpoint100With5PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(100)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificPropertyHashMap("prop1", "value1")
			e.SetProviderSpecificPropertyHashMap("prop2", "value2")
			e.SetProviderSpecificPropertyHashMap("prop3", "value3")
			e.SetProviderSpecificPropertyHashMap("prop4", "value3")
			e.SetProviderSpecificPropertyHashMap("prop5", "value3")
		}
	}
}

func BenchmarkEndpoint1000With5PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(1000)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificPropertyHashMap("prop1", "value1")
			e.SetProviderSpecificPropertyHashMap("prop2", "value2")
			e.SetProviderSpecificPropertyHashMap("prop3", "value3")
			e.SetProviderSpecificPropertyHashMap("prop4", "value3")
			e.SetProviderSpecificPropertyHashMap("prop5", "value3")
		}
	}
}

func BenchmarkEndpoint100With10Properties(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(100)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificProperty("prop1", "value1")
			e.SetProviderSpecificProperty("prop2", "value2")
			e.SetProviderSpecificProperty("prop3", "value3")
			e.SetProviderSpecificProperty("prop4", "value3")
			e.SetProviderSpecificProperty("prop5", "value3")
			e.SetProviderSpecificProperty("prop6", "value3")
			e.SetProviderSpecificProperty("prop7", "value3")
			e.SetProviderSpecificProperty("prop8", "value3")
			e.SetProviderSpecificProperty("prop9", "value3")
			e.SetProviderSpecificProperty("prop10", "value3")
		}
	}
}

func BenchmarkEndpoint1000With10Properties(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(1000)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificProperty("prop1", "value1")
			e.SetProviderSpecificProperty("prop2", "value2")
			e.SetProviderSpecificProperty("prop3", "value3")
			e.SetProviderSpecificProperty("prop4", "value3")
			e.SetProviderSpecificProperty("prop5", "value3")
			e.SetProviderSpecificProperty("prop6", "value3")
			e.SetProviderSpecificProperty("prop7", "value3")
			e.SetProviderSpecificProperty("prop8", "value3")
			e.SetProviderSpecificProperty("prop9", "value3")
			e.SetProviderSpecificProperty("prop10", "value3")
		}
	}
}

func BenchmarkEndpoint100With10PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(100)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificPropertyHashMap("prop1", "value1")
			e.SetProviderSpecificPropertyHashMap("prop2", "value2")
			e.SetProviderSpecificPropertyHashMap("prop3", "value3")
			e.SetProviderSpecificPropertyHashMap("prop4", "value3")
			e.SetProviderSpecificPropertyHashMap("prop5", "value3")
			e.SetProviderSpecificPropertyHashMap("prop6", "value3")
			e.SetProviderSpecificPropertyHashMap("prop7", "value3")
			e.SetProviderSpecificPropertyHashMap("prop8", "value3")
			e.SetProviderSpecificPropertyHashMap("prop9", "value3")
			e.SetProviderSpecificPropertyHashMap("prop10", "value3")
		}
	}
}

func BenchmarkEndpoint1000With10PropertiesHashMap(b *testing.B) {
	endpoints := buildEndpointsWithoutProperties(1000)

	for b.Loop() {
		for _, e := range endpoints {
			e.SetProviderSpecificPropertyHashMap("prop1", "value1")
			e.SetProviderSpecificPropertyHashMap("prop2", "value2")
			e.SetProviderSpecificPropertyHashMap("prop3", "value3")
			e.SetProviderSpecificPropertyHashMap("prop4", "value3")
			e.SetProviderSpecificPropertyHashMap("prop5", "value3")
			e.SetProviderSpecificPropertyHashMap("prop6", "value3")
			e.SetProviderSpecificPropertyHashMap("prop7", "value3")
			e.SetProviderSpecificPropertyHashMap("prop8", "value3")
			e.SetProviderSpecificPropertyHashMap("prop9", "value3")
			e.SetProviderSpecificPropertyHashMap("prop10", "value3")
		}
	}
}

func buildEndpointsWithoutProperties(nEndpoints int) []Endpoint {
	endpoints := make([]Endpoint, nEndpoints)
	for i := 0; i < nEndpoints; i++ {
		endpoints[i] = *NewEndpoint(fmt.Sprintf("index-%d.example.com", i), RecordTypeA)
	}
	return endpoints
}
