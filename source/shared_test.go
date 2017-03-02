package source

import (
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// test helper functions

func validateEndpoints(t *testing.T, endpoints, expected []endpoint.Endpoint) {
	if len(endpoints) != len(expected) {
		t.Fatalf("expected %d endpoints, got %d", len(expected), len(endpoints))
	}

	for i := range endpoints {
		validateEndpoint(t, endpoints[i], expected[i])
	}
}

func validateEndpoint(t *testing.T, endpoint, expected endpoint.Endpoint) {
	if endpoint.DNSName != expected.DNSName {
		t.Errorf("expected %s, got %s", expected.DNSName, endpoint.DNSName)
	}

	if endpoint.Target != expected.Target {
		t.Errorf("expected %s, got %s", expected.Target, endpoint.Target)
	}
}
