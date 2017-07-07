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
			map[string]string{"zone.com": "1234567"},
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
