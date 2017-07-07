package provider

import (
	"testing"
)

func TestZoneFinder(t *testing.T) {
	var test = []struct {
		hostname string
		zones    map[string]string
		expect   string
	}{
		{
			"foobar.zone.com",
			map[string]string{"zone.com": "123456"},
			"zone.com",
		},
	}
	for _, v := range test {
		suitableZone := zoneFinder(v.hostname, v.zones)
		if suitableZone != v.expect {
			t.Fatalf("expect %v, but got %v", v.expect, suitableZone)
		}
	}
}
