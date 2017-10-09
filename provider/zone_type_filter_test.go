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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/stretchr/testify/assert"
)

func TestZoneTypeFilterMatch(t *testing.T) {
	publicZone := &route53.HostedZone{Config: &route53.HostedZoneConfig{PrivateZone: aws.Bool(false)}}
	privateZone := &route53.HostedZone{Config: &route53.HostedZoneConfig{PrivateZone: aws.Bool(true)}}

	for _, tc := range []struct {
		zoneTypeFilter string
		zone           *route53.HostedZone
		matches        bool
	}{
		{
			"", publicZone, true,
		},
		{
			"", privateZone, true,
		},
		{
			"public", publicZone, true,
		},
		{
			"public", privateZone, false,
		},
		{
			"private", publicZone, false,
		},
		{
			"private", privateZone, true,
		},
		{
			"unknown", publicZone, false,
		},
		{
			"", &route53.HostedZone{}, true,
		},
		{
			"public", &route53.HostedZone{}, true,
		},
		{
			"private", &route53.HostedZone{}, false,
		},
	} {
		zoneTypeFilter := NewZoneTypeFilter(tc.zoneTypeFilter)
		assert.Equal(t, tc.matches, zoneTypeFilter.Match(tc.zone))
	}
}
