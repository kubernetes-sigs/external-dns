/*
Copyright 2025 The Kubernetes Authors.

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

package aws

import (
	"context"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestAWSRecordsV1(t *testing.T) {
	var zones HostedZones
	unmarshalTestHelper("/fixtures/160-plus-zones.yaml", &zones, t)

	stub := NewRoute53APIFixtureStub(&zones)
	provider := providerFilters(stub,
		WithZoneIDFilters(
			"Z10242883PKPS38KA4S6C", "Z10295763LSQ170JCTR78",
			"Z102957NOTEXISTS", "Z09418121E8V6WT4FASZE",
		),
		WithDomainFilters("w2.w1.ex.com", "ex.com"),
	)

	ctx := context.Background()
	z, err := provider.Zones(ctx)
	assert.NoError(t, err)
	assert.EqualValues(t, 3, len(z))
}

func TestAWSZonesFilterWithTags(t *testing.T) {
	var zones HostedZones
	unmarshalTestHelper("/fixtures/160-plus-zones.yaml", &zones, t)

	stub := NewRoute53APIFixtureStub(&zones)
	provider := providerFilters(stub,
		WithZoneTagFilters([]string{"level=5", "owner=ext-dns"}),
	)

	ctx := context.Background()
	z, err := provider.Zones(ctx)
	assert.NoError(t, err)
	assert.EqualValues(t, 24, len(z))
	assert.Equal(t, 169, stub.calls["listtagsforresource"])
}

func TestAWSZonesSecondRequestHitsTheCache(t *testing.T) {
	var zones HostedZones
	unmarshalTestHelper("/fixtures/160-plus-zones.yaml", &zones, t)

	stub := NewRoute53APIFixtureStub(&zones)
	provider := providerFilters(stub)

	ctx := context.Background()
	_, err := provider.Zones(ctx)
	assert.NoError(t, err)
	b := testutils.LogsToBuffer(log.DebugLevel, t)
	_, _ = provider.Zones(ctx)
	assert.Contains(t, b.String(), "level=debug msg=\"Using cached zones list\"")
}
