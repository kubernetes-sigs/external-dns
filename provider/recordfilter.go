package provider

// recordTypeFilter returns true only for supported record types.
// Currently only A, CNAME and TXT record types are supported.
func recordTypeFilter(recordType string) bool {
	switch recordType {
	case "A", "CNAME", "TXT":
		return true
	default:
		return false
	}
}
