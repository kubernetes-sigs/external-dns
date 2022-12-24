package cloudns

import (
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

var testEndpoints = []*endpoint.Endpoint{
	endpoint.NewEndpoint("example.com", "A", "1.2.3.4"),
	endpoint.NewEndpoint("example.com", "A", "5.6.7.8"),
	endpoint.NewEndpoint("example.com", "A", "9.10.11.12"),
	endpoint.NewEndpoint("example.com", "CNAME", "target.com"),
	endpoint.NewEndpoint("example.com", "CNAME", "target1.com"),
	endpoint.NewEndpoint("example.com", "CNAME", "target2.com"),
	endpoint.NewEndpoint("example.com", "CNAME", "target3.com"),
	endpoint.NewEndpoint("test.com", "A", "1.2.3.4"),
}

func TestMergeEndpointsByNameType(t *testing.T) {
	// Test case with no merge
	endpoints1 := append(testEndpoints[6:], testEndpoints[0])
	mergedEndpoints1 := mergeEndpointsByNameType(endpoints1)
	if len(mergedEndpoints1) != len(endpoints1) {
		t.Errorf("Expected mergeEndpointsByNameType to return %d endpoints, got %d", len(endpoints1), len(mergedEndpoints1))
	}

	// Test case with merge of A records
	endpoints2 := testEndpoints[:4]
	mergedEndpoints2 := mergeEndpointsByNameType(endpoints2)
	if len(mergedEndpoints2) != 2 {
		t.Errorf("Expected mergeEndpointsByNameType to return 2 endpoints, got %d", len(mergedEndpoints2))
	}
	if mergedEndpoints2[0].DNSName != "example.com" || mergedEndpoints2[0].RecordType != "A" || len(mergedEndpoints2[0].Targets) != 3 {
		t.Error("Unexpected endpoint returned after merge")
	}
	if mergedEndpoints2[1].DNSName != "example.com" || mergedEndpoints2[1].RecordType != "CNAME" || len(mergedEndpoints2[1].Targets) != 1 {
		t.Error("Unexpected endpoint returned after merge")
	}

	// Test case with merge of CNAME records
	endpoints3 := testEndpoints[4:]
	mergedEndpoints3 := mergeEndpointsByNameType(endpoints3)
	if len(mergedEndpoints3) != 2 {
		t.Errorf("Expected mergeEndpointsByNameType to return 2 endpoints, got %d", len(mergedEndpoints3))
	}
	if mergedEndpoints3[0].DNSName != "example.com" || mergedEndpoints3[0].RecordType != "CNAME" || len(mergedEndpoints3[0].Targets) != 3 {
		t.Error("Unexpected endpoint returned after merge")
	}
	if mergedEndpoints3[1].DNSName != "test.com" || mergedEndpoints3[1].RecordType != "A" || len(mergedEndpoints3[1].Targets) != 1 {
		t.Error("Unexpected endpoint returned after merge")
	}

}

// isValidTTL checks if a given Time to Live (TTL) value is valid.
// Valid TTL values are strings representing a number of seconds.
//
// The function returns true if the given TTL value is valid, and false otherwise.
func TestIsValidTTL(t *testing.T) {
	// Test valid TTLs
	if !isValidTTL("60") {
		t.Error("Expected isValidTTL to return true for TTL 60, got false")
	}
	if !isValidTTL("3600") {
		t.Error("Expected isValidTTL to return true for TTL 3600, got false")
	}
	if !isValidTTL("1209600") {
		t.Error("Expected isValidTTL to return true for TTL 1209600, got false")
	}

	// Test invalid TTLs
	if isValidTTL("0") {
		t.Error("Expected isValidTTL to return false for TTL 0, got true")
	}
	if isValidTTL("300000") {
		t.Error("Expected isValidTTL to return false for TTL 300000, got true")
	}
	if isValidTTL("abc") {
		t.Error("Expected isValidTTL to return false for TTL abc, got true")
	}
}

// rootZone returns the root zone of a domain name. A root zone is the last two parts
// of a domain name, separated by a "." character. If the domain name has less than two parts,
// the domain name is returned as-is.
// domain is the domain name to be checked.
func TestRootZone(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected string
	}{
		{
			name:     "two-part domain",
			domain:   "easy.com",
			expected: "easy.com",
		},
		{
			name:     "three-part domain",
			domain:   "test.this.program.com",
			expected: "program.com",
		},
		{
			name:     "four-part domain",
			domain:   "something.really.long.com",
			expected: "long.com",
		},
		{
			name:     "one-part domain",
			domain:   "root",
			expected: "root",
		},
		{
			name:     "empty domain",
			domain:   "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := rootZone(test.domain)
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}

// removeRootZone removes the given root zone and any trailing periods from the domain name.
// If the root zone is not present in the domain, the domain is returned unmodified.
// domain is the domain name to be modified.
// rootZone is the root zone to be removed from the domain name.
func TestRemoveRootZone(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		rootZone string
		expected string
	}{
		{
			name:     "root zone not present",
			domain:   "www.example.com",
			rootZone: "foo",
			expected: "www.example.com",
		},
		{
			name:     "root zone at end of domain",
			domain:   "www.example.com",
			rootZone: "com",
			expected: "www.example",
		},
		{
			name:     "root zone in middle of domain",
			domain:   "www.example.co.uk",
			rootZone: "co.uk",
			expected: "www.example",
		},
		{
			name:     "root zone spans multiple levels",
			domain:   "www.example.co.uk",
			rootZone: "example.co.uk",
			expected: "www",
		},
		{
			name:     "root zone is entire domain",
			domain:   "example.co.uk",
			rootZone: "example.co.uk",
			expected: "",
		},
		{
			name:     "root zone is root",
			domain:   ".",
			rootZone: ".",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeRootZone(test.domain, test.rootZone)
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}

// TestRemoveLastOccurrance tests the removeLastOccurrance function.
// It verifies that the function correctly removes the last occurrence of the given
// substring from the input string, and returns the modified string.
// The test cases cover various scenarios including when the substring is not present
// in the input string, when it is at the start or end of the input string, and when
// there are multiple occurrences of the substring in the input string.
func TestRemoveLastOccurrance(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		subStr   string
		expected string
	}{
		{
			name:     "subdomain not present",
			str:      "www.example.com",
			subStr:   "foo",
			expected: "www.example.com",
		},
		{
			name:     "subdomain at start of domain",
			str:      "www.example.com",
			subStr:   "www",
			expected: ".example.com",
		},
		{
			name:     "subdomain at end of domain",
			str:      "www.example.com",
			subStr:   "com",
			expected: "www.example.",
		},
		{
			name:     "subdomain in middle of domain",
			str:      "www.example.com",
			subStr:   ".e",
			expected: "wwwxample.com",
		},
		{
			name:     "multiple occurrences of subdomain",
			str:      "www.example.com.example.com",
			subStr:   "example",
			expected: "www.example.com..com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeLastOccurrance(test.str, test.subStr)
			if actual != test.expected {
				t.Errorf("got %q, want %q", actual, test.expected)
			}
		})
	}
}
