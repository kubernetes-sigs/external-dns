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

func TestZoneIDName(t *testing.T) {
	z := zoneIDName{}
	z.Add("123456", "foo.bar")
	z.Add("123456", "qux.baz")
	z.Add("654321", "foo.qux.baz")

	assert.Equal(t, zoneIDName{
		"123456": "qux.baz",
		"654321": "foo.qux.baz",
	}, z)

	// simple entry in a domain
	zoneID, zoneName := z.FindZone("name.qux.baz")
	assert.Equal(t, "qux.baz", zoneName)
	assert.Equal(t, "123456", zoneID)

	// simple entry in a domain's subdomain.
	zoneID, zoneName = z.FindZone("name.foo.qux.baz")
	assert.Equal(t, "foo.qux.baz", zoneName)
	assert.Equal(t, "654321", zoneID)

	// no possible zone for entry
	zoneID, zoneName = z.FindZone("name.qux.foo")
	assert.Equal(t, "", zoneName)
	assert.Equal(t, "", zoneID)

	// no possible zone for entry of a substring to valid a zone
	zoneID, zoneName = z.FindZone("nomatch-foo.bar")
	assert.Equal(t, "", zoneName)
	assert.Equal(t, "", zoneID)

	// entry's suffix matches a subdomain but doesn't belong there
	zoneID, zoneName = z.FindZone("name-foo.qux.baz")
	assert.Equal(t, "qux.baz", zoneName)
	assert.Equal(t, "123456", zoneID)

	// entry is an exact match of the domain (e.g. azure provider)
	zoneID, zoneName = z.FindZone("foo.qux.baz")
	assert.Equal(t, "foo.qux.baz", zoneName)
	assert.Equal(t, "654321", zoneID)
}
