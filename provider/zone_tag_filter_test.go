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

func TestZoneTagFilterMatch(t *testing.T) {
	for _, tc := range []struct {
		name          string
		zoneTagFilter []string
		zoneTags      map[string]string
		matches       bool
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
	} {
		zoneTagFilter := NewZoneTagFilter(tc.zoneTagFilter)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.matches, zoneTagFilter.Match(tc.zoneTags))
		})
	}
}
