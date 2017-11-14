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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

const (
	zoneTypePublic  = "public"
	zoneTypePrivate = "private"
)

// ZoneTypeFilter holds a zone type to filter for.
type ZoneTypeFilter struct {
	zoneType string
}

// NewZoneTypeFilter returns a new ZoneTypeFilter given a zone type to filter for.
func NewZoneTypeFilter(zoneType string) ZoneTypeFilter {
	return ZoneTypeFilter{zoneType: zoneType}
}

// Match checks whether a zone matches the zone type that's filtered for.
func (f ZoneTypeFilter) Match(zone *route53.HostedZone) bool {
	// An empty zone filter includes all hosted zones.
	if f.zoneType == "" {
		return true
	}

	// If the zone has no config we assume it's a public zone since the config's field
	// `PrivateZone` is false by default in go.
	if zone.Config == nil {
		return f.zoneType == zoneTypePublic
	}

	// Given a zone type we return true if the given zone matches this type.
	switch f.zoneType {
	case zoneTypePublic:
		return !aws.BoolValue(zone.Config.PrivateZone)
	case zoneTypePrivate:
		return aws.BoolValue(zone.Config.PrivateZone)
	}

	// We return false on any other path, e.g. unknown zone type filter value.
	return false
}
