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

package zonetagfilter

import (
	"strings"
)

// For AWS every zone could have up to 50 tags, and it could be up-to 500 zones in a region

// ZoneTagFilter holds a list of zone tags to filter by
type ZoneTagFilter struct {
	zoneTags    []string
	zoneTagsMap map[string]string
}

// NewZoneTagFilter returns a new ZoneTagFilter given a list of zone tags
func NewZoneTagFilter(tags []string) ZoneTagFilter {
	if len(tags) == 1 && len(tags[0]) == 0 {
		tags = []string{}
	}
	z := ZoneTagFilter{}
	z.zoneTags = tags
	z.zoneTagsMap = make(map[string]string, len(tags))
	// tags pre-processing, to make sure the pre-processing is not happening at the time of filtering
	for _, tag := range z.zoneTags {
		parts := strings.SplitN(tag, "=", 2)
		key := strings.TrimSpace(parts[0])
		if key == "" {
			continue
		}
		if len(parts) == 2 {
			value := strings.TrimSpace(parts[1])
			z.zoneTagsMap[key] = value
		} else {
			z.zoneTagsMap[key] = ""
		}
	}
	return z
}

// Match checks whether a zone's set of tags matches the provided tag values
func (f ZoneTagFilter) Match(tagsMap map[string]string) bool {
	for key, v := range f.zoneTagsMap {
		if value, hasTag := tagsMap[key]; !hasTag || (v != "" && value != v) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if there are no tags for the filter
func (f ZoneTagFilter) IsEmpty() bool {
	return len(f.zoneTags) == 0
}
