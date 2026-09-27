/*
Copyright 2026 The Kubernetes Authors.

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

package plan

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
)

func countFor(objects []PlannedObject, name string) int {
	for _, obj := range objects {
		if obj.Ref.Name() == name {
			return obj.Endpoints
		}
	}
	return -1
}

func TestPlannedObjects(t *testing.T) {
	first := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "first", "", "crd")
	second := events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", "second", "", "crd")

	desired := []*endpoint.Endpoint{
		endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(first),
		endpoint.NewEndpoint("b.example.com", endpoint.RecordTypeA, "10.0.0.2").WithRefObject(first),
		endpoint.NewEndpoint("c.example.com", endpoint.RecordTypeA, "10.0.0.3").WithRefObject(second),
		endpoint.NewEndpoint("d.example.com", endpoint.RecordTypeA, "10.0.0.4"),
	}

	t.Run("one entry per object, counting its planned endpoints", func(t *testing.T) {
		objects := (&Plan{Desired: desired, Planned: desired}).PlannedObjects()

		require.Len(t, objects, 2)
		assert.Equal(t, 2, countFor(objects, "first"))
		assert.Equal(t, 1, countFor(objects, "second"))
	})

	t.Run("filtered-out objects are kept at zero", func(t *testing.T) {
		objects := (&Plan{Desired: desired, Planned: desired[:2]}).PlannedObjects()

		require.Len(t, objects, 2)
		assert.Equal(t, 2, countFor(objects, "first"))
		assert.Equal(t, 0, countFor(objects, "second"))
	})

	t.Run("Calculate fills Planned from the domain filter", func(t *testing.T) {
		p := (&Plan{
			Desired:        desired,
			DomainFilter:   endpoint.MatchAllDomainFilters{endpoint.NewDomainFilter([]string{"a.example.com"})},
			ManagedRecords: []string{endpoint.RecordTypeA},
		}).Calculate()

		objects := p.PlannedObjects()
		assert.Equal(t, 1, countFor(objects, "first"))
		assert.Equal(t, 0, countFor(objects, "second"))
	})

	t.Run("Calculate drops endpoints another owner holds", func(t *testing.T) {
		owned := func(name, recordType, target, owner string) *endpoint.Endpoint {
			return endpoint.NewEndpoint(name, recordType, target).WithLabel(endpoint.OwnerLabelKey, owner)
		}
		current := []*endpoint.Endpoint{
			owned("mine.example.com", endpoint.RecordTypeA, "10.0.0.9", "me"),
			owned("theirs.example.com", endpoint.RecordTypeA, "10.0.0.9", "them"),
			owned("steady.example.com", endpoint.RecordTypeA, "10.0.0.5", "them"),
			owned("shared.example.com", endpoint.RecordTypeA, "10.0.0.6", "them"),
		}
		ref := func(name string) *events.ObjectReference {
			return events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", name, "", "crd")
		}
		desired := []*endpoint.Endpoint{
			endpoint.NewEndpoint("mine.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(ref("update-mine")),
			endpoint.NewEndpoint("free.example.com", endpoint.RecordTypeA, "10.0.0.2").WithRefObject(ref("create-free")),
			endpoint.NewEndpoint("theirs.example.com", endpoint.RecordTypeA, "10.0.0.3").WithRefObject(ref("update-theirs")),
			endpoint.NewEndpoint("steady.example.com", endpoint.RecordTypeA, "10.0.0.5").WithRefObject(ref("in-sync-theirs")),
			endpoint.NewEndpoint("shared.example.com", endpoint.RecordTypeAAAA, "::1").WithRefObject(ref("create-next-to-theirs")),
		}

		p := (&Plan{
			Current:        current,
			Desired:        desired,
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA},
			OwnerID:        "me",
		}).Calculate()

		objects := p.PlannedObjects()
		require.Len(t, objects, 5)
		assert.Equal(t, 1, countFor(objects, "update-mine"))
		assert.Equal(t, 1, countFor(objects, "create-free"))
		assert.Equal(t, 0, countFor(objects, "update-theirs"))
		assert.Equal(t, 0, countFor(objects, "in-sync-theirs"))
		assert.Equal(t, 0, countFor(objects, "create-next-to-theirs"))
		// Cross-check against what the planner actually does.
		assert.Len(t, p.Changes.Create, 1)
		assert.Len(t, p.Changes.UpdateNew, 1)
	})
}
