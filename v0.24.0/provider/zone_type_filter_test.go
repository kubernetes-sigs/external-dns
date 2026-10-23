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

	"github.com/stretchr/testify/assert"
)

func TestZoneTypeFilterMatch(t *testing.T) {
	publicZoneStr := "public"
	privateZoneStr := "private"

	for _, tc := range []struct {
		zoneTypeFilter string
		matches        bool
		zones          []string
	}{
		{
			"", true, []string{publicZoneStr, privateZoneStr},
		},
		{
			"public", true, []string{publicZoneStr},
		},
		{
			"public", false, []string{privateZoneStr},
		},
		{
			"private", true, []string{privateZoneStr},
		},
		{
			"private", false, []string{publicZoneStr},
		},
		{
			"unknown", false, []string{publicZoneStr},
		},
	} {
		t.Run(tc.zoneTypeFilter, func(t *testing.T) {
			zoneTypeFilter := NewZoneTypeFilter(tc.zoneTypeFilter)
			for _, zone := range tc.zones {
				assert.Equal(t, tc.matches, zoneTypeFilter.Match(zone))
			}
		})
	}
}
