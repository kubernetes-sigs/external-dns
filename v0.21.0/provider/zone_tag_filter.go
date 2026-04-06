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
	"strings"
)

// ZoneTagFilter holds a list of zone tags to filter by
type ZoneTagFilter struct {
	tagsMap map[string]string
}

// NewZoneTagFilter returns a new ZoneTagFilter given a list of zone tags
func NewZoneTagFilter(tags []string) ZoneTagFilter {
	if len(tags) == 1 && len(tags[0]) == 0 {
		tags = []string{}
	}
	tagsMap := make(map[string]string)
	// tags pre-processing, to make sure the pre-processing is not happening at the time of filtering
	for _, tag := range tags {
		parts := strings.SplitN(tag, "=", 2)
		key := strings.TrimSpace(parts[0])
		if key == "" {
			continue
		}
		if len(parts) == 2 {
			value := strings.TrimSpace(parts[1])
			tagsMap[key] = value
		} else {
			tagsMap[key] = ""
		}
	}
	return ZoneTagFilter{tagsMap: tagsMap}
}

// Match checks whether a zone's set of tags matches the provided tag values
func (f ZoneTagFilter) Match(tagsMap map[string]string) bool {
	for key, v := range f.tagsMap {
		if value, hasTag := tagsMap[key]; !hasTag || (v != "" && value != v) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if there are no tags for the filter
func (f ZoneTagFilter) IsEmpty() bool {
	return len(f.tagsMap) == 0
}
