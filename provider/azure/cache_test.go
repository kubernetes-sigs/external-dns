/*
Copyright 2024 The Kubernetes Authors.

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

package azure

import (
	"testing"
	"time"

	dns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	"github.com/stretchr/testify/assert"
)

func TestzonesCache(t *testing.T) {
	now := time.Now()
	zoneName := "example.com"
	var testCases = map[string]struct {
		z       *zonesCache[dns.Zone]
		expired bool
	}{
		"inactive-zone-cache": {
			&zonesCache[dns.Zone]{
				duration: 0 * time.Second,
			},
			true,
		},
		"empty-active-zone-cache": {
			&zonesCache[dns.Zone]{
				duration: 30 * time.Second,
			},
			true,
		},
		"expired-zone-cache": {
			&zonesCache[dns.Zone]{
				age:      now.Add(-300 * time.Second),
				duration: 30 * time.Second,
			},
			true,
		},
		"active-zone-cache": {
			&zonesCache[dns.Zone]{
				zones: []dns.Zone{{
					Name: &zoneName,
				}},
				duration: 30 * time.Second,
				age:      now,
			},
			false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expired, testCase.z.Expired())
			var resetZoneLength = 1
			if testCase.z.duration == 0 {
				resetZoneLength = 0
			}
			testCase.z.Reset([]dns.Zone{{
				Name: &zoneName,
			}})
			assert.Len(t, testCase.z.Get(), resetZoneLength)
		})
	}
}
