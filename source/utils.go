package source

import (
	"fmt"
	"net/netip"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
)

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A/AAAA for IPs and type CNAME for everything else.
func suitableType(target string) string {
	netIP, err := netip.ParseAddr(target)
	if err == nil && netIP.Is4() {
		return endpoint.RecordTypeA
	} else if err == nil && netIP.Is6() {
		return endpoint.RecordTypeAAAA
	}
	return endpoint.RecordTypeCNAME
}

// ParseIngress parses an ingress string in the format "namespace/name" or "name".
// It returns the namespace and name extracted from the string, or an error if the format is invalid.
// If the namespace is not provided, it defaults to an empty string.
func ParseIngress(ingress string) (namespace, name string, err error) {
	parts := strings.Split(ingress, "/")
	if len(parts) == 2 {
		namespace, name = parts[0], parts[1]
	} else if len(parts) == 1 {
		name = parts[0]
	} else {
		err = fmt.Errorf("invalid ingress name (name or namespace/name) found %q", ingress)
	}

	return
}

// SelectorMatchesServiceSelector checks if all key-value pairs in the selector map
// are present and match the corresponding key-value pairs in the svcSelector map.
// It returns true if all pairs match, otherwise it returns false.
func SelectorMatchesServiceSelector(selector, svcSelector map[string]string) bool {
	for k, v := range selector {
		if lbl, ok := svcSelector[k]; !ok || lbl != v {
			return false
		}
	}
	return true
}
