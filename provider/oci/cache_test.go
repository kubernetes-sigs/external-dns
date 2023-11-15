/*
Copyright 2023 The Kubernetes Authors.

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

package oci

import (
	"github.com/oracle/oci-go-sdk/v65/dns"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestZoneCache(t *testing.T) {
	now := time.Now()
	var testCases = map[string]struct {
		z       *zoneCache
		expired bool
	}{
		"inactive-zone-cache": {
			&zoneCache{
				duration: 0 * time.Second,
			},
			true,
		},
		"empty-active-zone-cache": {
			&zoneCache{
				duration: 30 * time.Second,
			},
			true,
		},
		"expired-zone-cache": {
			&zoneCache{
				age:      now.Add(300 * time.Second),
				duration: 30 * time.Second,
			},
			true,
		},
		"active-zone-cache": {
			&zoneCache{
				zones: map[string]dns.ZoneSummary{
					zoneIdBaz: testPrivateZoneSummaryBaz,
				},
				duration: 30 * time.Second,
			},
			true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expired, testCase.z.Expired())
			var resetZoneLength = 1
			if testCase.z.duration == 0 {
				resetZoneLength = 0
			}
			testCase.z.Reset(map[string]dns.ZoneSummary{
				zoneIdQux: testPrivateZoneSummaryQux,
			})
			assert.Len(t, testCase.z.Get(), resetZoneLength)
		})
	}
}
