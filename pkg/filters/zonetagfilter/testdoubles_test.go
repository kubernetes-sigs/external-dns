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
	"fmt"
)

type filterZoneTags struct {
	filterTags []string
	zoneTags   map[string]string
}

func generateTagFilterAndZoneTagsForMatch(filter, zone int) filterZoneTags {
	validate(filter, zone)

	filterTags := make([]string, 0, filter)
	for i := filter - 1; i >= 0; i-- {
		filterTags = append(filterTags, fmt.Sprintf("tag-%d=value-%d", i, i))
	}
	zoneTags := make(map[string]string, zone)
	for i := 0; i < zone; i++ {
		zoneTags[fmt.Sprintf("tag-%d", i)] = fmt.Sprintf("value-%d", i)
	}
	return filterZoneTags{filterTags, zoneTags}
}

// TODO: explain what is doing this function
func generateTagFilterAndZoneTagsForNotMatch(filter, zone int) filterZoneTags {
	validate(filter, zone)

	filterTags := make([]string, 0, filter)
	for i := filter - 1; i >= 0; i-- {
		filterTags = append(filterTags, fmt.Sprintf("tag-%d=value-%d", i+50, i))
	}
	zoneTags := make(map[string]string, zone)
	for i := 0; i < zone; i++ {
		zoneTags[fmt.Sprintf("tag-%d", i)] = fmt.Sprintf("value-%d", i+2)
	}
	return filterZoneTags{filterTags, zoneTags}
}

func validate(filter int, zone int) {
	if zone > 50 {
		panic("zone number is too high")
	}
	if filter > zone {
		panic("filter number is too high")
	}
}
