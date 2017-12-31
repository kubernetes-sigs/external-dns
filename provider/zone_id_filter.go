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

import "strings"

// ZoneIDFilter holds a list of zone ids to filter by
type ZoneIDFilter struct {
	zoneIDs []string
}

// NewZoneIDFilter returns a new ZoneIDFilter given a list of zone ids
func NewZoneIDFilter(zoneIDs []string) ZoneIDFilter {
	return ZoneIDFilter{zoneIDs}
}

// Match checks whether a zone matches one of the provided zone ids
func (f ZoneIDFilter) Match(zoneID string) bool {
	// An empty filter includes all zones.
	if len(f.zoneIDs) == 0 {
		return true
	}

	for _, id := range f.zoneIDs {
		if strings.HasSuffix(zoneID, id) {
			return true
		}
	}

	return false
}
