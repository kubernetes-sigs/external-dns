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

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestZoneIDName(t *testing.T) {
	z := ZoneIDName{}
	z.Add("123456", "foo.bar")
	z.Add("123456", "qux.baz")
	z.Add("654321", "foo.qux.baz")
	z.Add("987654", "エイミー.みんな")
	z.Add("123123", "_metadata.example.com")
	z.Add("456456", "_metadata.エイミー.みんな")

	assert.Equal(t, ZoneIDName{
		"123456": "qux.baz",
		"654321": "foo.qux.baz",
		"987654": "エイミー.みんな",
		"123123": "_metadata.example.com",
		"456456": "_metadata.エイミー.みんな",
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
	assert.Empty(t, zoneName)
	assert.Empty(t, zoneID)

	// no possible zone for entry of a substring to valid a zone
	zoneID, zoneName = z.FindZone("nomatch-foo.bar")
	assert.Empty(t, zoneName)
	assert.Empty(t, zoneID)

	// entry's suffix matches a subdomain but doesn't belong there
	zoneID, zoneName = z.FindZone("name-foo.qux.baz")
	assert.Equal(t, "qux.baz", zoneName)
	assert.Equal(t, "123456", zoneID)

	// entry is an exact match of the domain (e.g. azure provider)
	zoneID, zoneName = z.FindZone("foo.qux.baz")
	assert.Equal(t, "foo.qux.baz", zoneName)
	assert.Equal(t, "654321", zoneID)

	// entry gets normalized before finding
	zoneID, zoneName = z.FindZone("xn--eckh0ome.xn--q9jyb4c")
	assert.Equal(t, "エイミー.みんな", zoneName)
	assert.Equal(t, "987654", zoneID)

	hook := testutils.LogsUnderTestWithLogLevel(log.WarnLevel, t)
	_, _ = z.FindZone("???")

	testutils.TestHelperLogContains("Failed to convert label '???' of hostname '???' to its Unicode form: idna: disallowed rune U+003F", hook, t)
}
