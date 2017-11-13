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

	zoneID, zoneName := z.FindZone("name.qux.baz")
	assert.Equal(t, "qux.baz", zoneName)
	assert.Equal(t, "123456", zoneID)

	zoneID, zoneName = z.FindZone("name.foo.qux.baz")
	assert.Equal(t, "foo.qux.baz", zoneName)
	assert.Equal(t, "654321", zoneID)

	zoneID, zoneName = z.FindZone("name.qux.foo")
	assert.Equal(t, "", zoneName)
	assert.Equal(t, "", zoneID)
}
