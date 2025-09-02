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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var basicZoneTags = []struct {
	name       string
	tagsFilter []string
	zoneTags   map[string]string
	matches    bool
}{
	{
		"single tag no match", []string{"tag1=value1"}, map[string]string{"tag0": "value0"}, false,
	},
	{
		"single tag matches", []string{"tag1=value1"}, map[string]string{"tag1": "value1"}, true,
	},
	{
		"multiple tags no value match", []string{"tag1=value1"}, map[string]string{"tag0": "value0", "tag1": "value2"}, false,
	},
	{
		"multiple tags matches", []string{"tag1=value1"}, map[string]string{"tag0": "value0", "tag1": "value1"}, true,
	},
	{
		"tag name no match", []string{"tag1"}, map[string]string{"tag0": "value0"}, false,
	},
	{
		"tag name matches", []string{"tag1"}, map[string]string{"tag1": "value1"}, true,
	},
	{
		"multiple filter no match", []string{"tag1=value1", "tag2=value2"}, map[string]string{"tag1": "value1"}, false,
	},
	{
		"multiple filter matches", []string{"tag1=value1", "tag2=value2"}, map[string]string{"tag2": "value2", "tag1": "value1", "tag3": "value3"}, true,
	},
	{
		"empty tag filter matches all", []string{""}, map[string]string{"tag0": "value0"}, true,
	},
	{
		"tag filter without key and equal sign", []string{"tag1=value1", "=haha"}, map[string]string{"tag1": "value1"}, true,
	},
}

func TestZoneTagFilterMatch(t *testing.T) {
	for _, tc := range basicZoneTags {
		zoneTagFilter := NewZoneTagFilter(tc.tagsFilter)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.matches, zoneTagFilter.Match(tc.zoneTags))
		})
	}
}

func TestZoneTagFilterNotSupportedFormat(t *testing.T) {
	tests := []struct {
		desc string
		tags []string
		want map[string]string
	}{
		{desc: "multiple or separate values with commas", tags: []string{"key1=val1,key2=val2"}, want: map[string]string{"key1": "val1,key2=val2"}},
		{desc: "exclude tag", tags: []string{"!key1"}, want: map[string]string{"!key1": ""}},
		{desc: "exclude tags", tags: []string{"!key1=val"}, want: map[string]string{"!key1": "val"}},
		{desc: "key is empty", tags: []string{"=val"}, want: map[string]string{}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.desc), func(t *testing.T) {
			got := NewZoneTagFilter(tc.tags)
			assert.Equal(t, tc.want, got.tagsMap)
		})
	}
}

func TestZoneTagFilterMatchGeneratedValues(t *testing.T) {
	tests := []struct {
		filters int
		zones   int
		source  filterZoneTags
	}{
		{10, 30, generateTagFilterAndZoneTagsForMatch(10, 30)},
		{5, 40, generateTagFilterAndZoneTagsForMatch(5, 40)},
		{30, 50, generateTagFilterAndZoneTagsForMatch(30, 50)},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("filters:%d zones:%d", tc.filters, tc.zones), func(t *testing.T) {
			assert.True(t, tc.source.ZoneTagFilter.Match(tc.source.inputTags))
		})
	}
}

func TestZoneTagFilterNotMatchGeneratedValues(t *testing.T) {
	tests := []struct {
		filters int
		zones   int
		source  filterZoneTags
	}{
		{10, 30, generateTagFilterAndZoneTagsForNotMatch(10, 30)},
		{5, 40, generateTagFilterAndZoneTagsForNotMatch(5, 40)},
		{30, 50, generateTagFilterAndZoneTagsForNotMatch(30, 50)},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("filters:%d zones:%d", tc.filters, tc.zones), func(t *testing.T) {
			assert.False(t, tc.source.ZoneTagFilter.Match(tc.source.inputTags))
		})
	}
}

// benchmarks
func BenchmarkZoneTagFilterMatchBasic(b *testing.B) {
	for _, tc := range basicZoneTags {
		zoneTagFilter := NewZoneTagFilter(tc.tagsFilter)
		for range b.N {
			zoneTagFilter.Match(tc.zoneTags)
		}
	}
}

var benchFixtures = []struct {
	source filterZoneTags
}{
	// match
	{generateTagFilterAndZoneTagsForMatch(10, 30)},
	{generateTagFilterAndZoneTagsForMatch(5, 40)},
	{generateTagFilterAndZoneTagsForMatch(30, 50)},
	// 	no match
	{generateTagFilterAndZoneTagsForNotMatch(10, 30)},
	{generateTagFilterAndZoneTagsForNotMatch(5, 40)},
	{generateTagFilterAndZoneTagsForNotMatch(30, 50)},
}

func BenchmarkZoneTagFilterComplex(b *testing.B) {
	for _, tc := range benchFixtures {
		for range b.N {
			tc.source.ZoneTagFilter.Match(tc.source.inputTags)
		}
	}
}

// test doubles
type filterZoneTags struct {
	ZoneTagFilter
	inputTags map[string]string
}

// generateTagFilterAndZoneTagsForMatch generates filter tags and zone tags that do match.
func generateTagFilterAndZoneTagsForMatch(filter, zone int) filterZoneTags {
	return generateTagFilterAndZoneTags(filter, zone, true)
}

// generateTagFilterAndZoneTagsForNotMatch generates filter tags and zone tags that do not match.
func generateTagFilterAndZoneTagsForNotMatch(filter, zone int) filterZoneTags {
	return generateTagFilterAndZoneTags(filter, zone, false)
}

// generateTagFilterAndZoneTags generates filter tags and zone tags based on the match parameter.
func generateTagFilterAndZoneTags(filter, zone int, match bool) filterZoneTags {
	validate(filter, zone)
	toFilterTags := make([]string, 0, filter)
	inputTags := make(map[string]string, zone)

	for i := 0; i < filter; i++ {
		tagIndex := i
		if !match {
			tagIndex += 50
		}
		toFilterTags = append(toFilterTags, fmt.Sprintf("tag-%d=value-%d", tagIndex, i))
	}

	for i := 0; i < zone; i++ {
		tagIndex := i
		if !match {
			// Make sure the input tags are different from the filter tags
			tagIndex += 2
		}
		inputTags[fmt.Sprintf("tag-%d", i)] = fmt.Sprintf("value-%d", tagIndex)
	}

	return filterZoneTags{NewZoneTagFilter(toFilterTags), inputTags}
}

func validate(filter int, zone int) {
	if zone > 50 {
		panic("zone number is too high")
	}
	if filter > zone {
		panic("filter number is too high")
	}
}
