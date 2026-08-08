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

	route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"

	"github.com/stretchr/testify/assert"
)

func TestZoneTypeFilterMatch(t *testing.T) {
	publicZoneStr := "public"
	privateZoneStr := "private"
	publicZoneAWS := route53types.HostedZone{Config: &route53types.HostedZoneConfig{PrivateZone: false}}
	privateZoneAWS := route53types.HostedZone{Config: &route53types.HostedZoneConfig{PrivateZone: true}}

	for _, tc := range []struct {
		zoneTypeFilter string
		matches        bool
		zones          []any
	}{
		{
			"", true, []any{publicZoneStr, privateZoneStr, route53types.HostedZone{}},
		},
		{
			"public", true, []any{publicZoneStr, publicZoneAWS, route53types.HostedZone{}},
		},
		{
			"public", false, []any{privateZoneStr, privateZoneAWS},
		},
		{
			"private", true, []any{privateZoneStr, privateZoneAWS},
		},
		{
			"private", false, []any{publicZoneStr, publicZoneAWS, route53types.HostedZone{}},
		},
		{
			"unknown", false, []any{publicZoneStr},
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
