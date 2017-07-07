package provider

import "strings"

// zoneFinder returns the most suitable zone for a given hostname
// and a set of zones.
func zoneFinder(hostname string, zones map[string]string) string {
	var suitableZone string
	for zoneName := range zones {
		if strings.HasSuffix(hostname, zoneName) {
			if suitableZone == "" || len(zoneName) > len(suitableZone) {
				suitableZone = zoneName
			}
		}
	}
	return suitableZone
}
