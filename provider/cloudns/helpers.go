package cloudns

import (
	"fmt"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
)

// mergeEndpointsByNameType takes a slice of endpoints and returns a new slice of endpoints
// with the endpoints merged based on their DNS name and record type. If no merge occurs,
// the original slice of endpoints is returned.
// From pkg/digitalocean/provider.go
func mergeEndpointsByNameType(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpointsByNameType := map[string][]*endpoint.Endpoint{}

	for _, e := range endpoints {
		key := fmt.Sprintf("%s-%s", e.DNSName, e.RecordType)
		endpointsByNameType[key] = append(endpointsByNameType[key], e)
	}

	// If no merge occurred, just return the existing endpoints.
	if len(endpointsByNameType) == len(endpoints) {
		return endpoints
	}

	// Otherwise, construct a new list of endpoints with the endpoints merged.
	var result []*endpoint.Endpoint
	for _, endpoints := range endpointsByNameType {
		dnsName := endpoints[0].DNSName
		recordType := endpoints[0].RecordType
		ttl := endpoints[0].RecordTTL

		targets := make([]string, len(endpoints))
		for i, ep := range endpoints {
			targets[i] = ep.Targets[0]
		}

		e := endpoint.NewEndpoint(dnsName, recordType, targets...)
		e.RecordTTL = ttl
		result = append(result, e)
	}

	return result
}

// isValidTTL checks if the given time-to-live (TTL) value is valid.
// A valid TTL value is a string representation of a positive integer that is one of the following values:
// "60", "300", "900", "1800", "3600", "21600", "43200", "86400", "172800", "259200", "604800", "1209600", "2592000".
// The function returns true if the given TTL value is valid and false otherwise.
func isValidTTL(ttl string) bool {
	validTTLs := []string{"60", "300", "900", "1800", "3600", "21600", "43200", "86400", "172800", "259200", "604800", "1209600", "2592000"}

	for _, validTTL := range validTTLs {
		if ttl == validTTL {
			return true
		}
	}

	return false
}

// rootZone returns the root zone of a domain name.
// A root zone is the last two parts of a domain name, separated by a "." character.
// For example, the root zone of "test.this.program.com" is "program.com" and
// the root zone of "easy.com" is "easy.com".
// If the domain name has less than two parts, the domain name is returned as-is.
func rootZone(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return domain
	}
	return strings.Join(parts[len(parts)-2:], ".")
}

// Returns the domain name with the root zone and any trailing periods removed.
// domain is the domain name to be modified.
// rootZone is the root zone to be removed from the domain name.
func removeRootZone(domain string, rootZone string) string {

	if strings.LastIndex(domain, rootZone) == -1 {
		return domain
	}

	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return domain
	}
	rootZoneIndex := len(parts) - len(strings.Split(rootZone, "."))
	return strings.TrimSuffix(strings.Join(parts[:rootZoneIndex], "."), ".")
}

// removeLastOccurrence removes the last occurrence of the given substring from the given string.
// If the substring is not present, the original string is returned.
func removeLastOccurrance(str, subStr string) string {
	i := strings.LastIndex(str, subStr)

	if i == -1 {
		return str
	}

	return strings.Join([]string{str[:i], str[i+len(subStr):]}, "")
}
