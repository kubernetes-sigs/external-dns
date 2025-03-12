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

type zoneIDFilterTest struct {
	zoneIDFilter []string
	zone         string
	expected     bool
}

func TestZoneIDFilterMatch(t *testing.T) {
	zone := "/hostedzone/ZTST1"

	for _, tt := range []zoneIDFilterTest{
		{
			[]string{},
			zone,
			true,
		},
		{
			[]string{"/hostedzone/ZTST1"},
			zone,
			true,
		},
		{
			[]string{"/hostedzone/ZTST2"},
			zone,
			false,
		},
		{
			[]string{"ZTST1"},
			zone,
			true,
		},
		{
			[]string{"ZTST2"},
			zone,
			false,
		},
		{
			[]string{"/hostedzone/ZTST1", "/hostedzone/ZTST2"},
			zone,
			true,
		},
		{
			[]string{"/hostedzone/ZTST2", "/hostedzone/ZTST3"},
			zone,
			false,
		},
		{
			[]string{"/hostedzone/ZTST2", "/hostedzone/ZTST1"},
			zone,
			true,
		},
	} {
		zoneIDFilter := NewZoneIDFilter(tt.zoneIDFilter)
		assert.Equal(t, tt.expected, zoneIDFilter.Match(tt.zone))
	}
}
