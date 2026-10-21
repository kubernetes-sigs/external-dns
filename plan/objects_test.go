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

func find(objects []PlannedObject, name string) PlannedObject {
	for _, obj := range objects {
		if obj.Ref.Name() == name {
			return obj
		}
	}
	return PlannedObject{Endpoints: -1}
}

func crdRef(name string) *events.ObjectReference {
	return events.NewObjectReferenceFromParts("DNSEndpoint", "externaldns.k8s.io/v1alpha1", "ns", name, "", "crd")
}

func ownedA(name, target, owner string) *endpoint.Endpoint {
	return endpoint.NewEndpoint(name, endpoint.RecordTypeA, target).WithLabel(endpoint.OwnerLabelKey, owner)
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
		current := []*endpoint.Endpoint{
			ownedA("mine.example.com", "10.0.0.9", "me"),
			ownedA("calm.example.com", "10.0.0.4", "me"),
			ownedA("theirs.example.com", "10.0.0.9", "them"),
			ownedA("steady.example.com", "10.0.0.5", "them"),
			ownedA("shared.example.com", "10.0.0.6", "them"),
		}
		desired := []*endpoint.Endpoint{
			endpoint.NewEndpoint("mine.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(crdRef("update-mine")),
			endpoint.NewEndpoint("calm.example.com", endpoint.RecordTypeA, "10.0.0.4").WithRefObject(crdRef("in-sync-mine")),
			endpoint.NewEndpoint("free.example.com", endpoint.RecordTypeA, "10.0.0.2").WithRefObject(crdRef("create-free")),
			endpoint.NewEndpoint("theirs.example.com", endpoint.RecordTypeA, "10.0.0.3").WithRefObject(crdRef("update-theirs")),
			endpoint.NewEndpoint("steady.example.com", endpoint.RecordTypeA, "10.0.0.5").WithRefObject(crdRef("in-sync-theirs")),
			endpoint.NewEndpoint("shared.example.com", endpoint.RecordTypeAAAA, "::1").WithRefObject(crdRef("create-next-to-theirs")),
		}

		p := (&Plan{
			Current:        current,
			Desired:        desired,
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA},
			OwnerID:        "me",
		}).Calculate()

		objects := p.PlannedObjects()
		require.Len(t, objects, 6)
		assert.Equal(t, PlannedObject{Ref: crdRef("update-mine"), Endpoints: 1, Changed: true}, find(objects, "update-mine"))
		assert.Equal(t, PlannedObject{Ref: crdRef("create-free"), Endpoints: 1, Changed: true}, find(objects, "create-free"))
		assert.Equal(t, PlannedObject{Ref: crdRef("in-sync-mine"), Endpoints: 1}, find(objects, "in-sync-mine"))
		assert.Equal(t, 0, countFor(objects, "update-theirs"))
		assert.Equal(t, 0, countFor(objects, "in-sync-theirs"))
		assert.Equal(t, 0, countFor(objects, "create-next-to-theirs"))
		// Cross-check against what the planner actually does.
		assert.Len(t, p.Changes.Create, 1)
		assert.Len(t, p.Changes.UpdateNew, 1)
	})

	t.Run("Calculate drops candidates that lose a conflict", func(t *testing.T) {
		desired := []*endpoint.Endpoint{
			endpoint.NewEndpoint("mixed.example.com", endpoint.RecordTypeCNAME, "target.example.com").WithRefObject(crdRef("cname-loser")),
			endpoint.NewEndpoint("mixed.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(crdRef("a-winner")),
			endpoint.NewEndpoint("dup.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(crdRef("dup-winner")),
			endpoint.NewEndpoint("dup.example.com", endpoint.RecordTypeA, "10.0.0.2").WithRefObject(crdRef("dup-loser")),
		}

		p := (&Plan{
			Desired:        desired,
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
		}).Calculate()

		objects := p.PlannedObjects()
		assert.Equal(t, 0, countFor(objects, "cname-loser"))
		assert.Equal(t, 1, countFor(objects, "a-winner"))
		assert.Equal(t, 1, countFor(objects, "dup-winner"))
		assert.Equal(t, 0, countFor(objects, "dup-loser"))
	})

	t.Run("Calculate drops updates the policy forbids", func(t *testing.T) {
		current := []*endpoint.Endpoint{
			ownedA("stale.example.com", "10.0.0.9", "me"),
			ownedA("calm.example.com", "10.0.0.4", "me"),
		}
		desired := []*endpoint.Endpoint{
			endpoint.NewEndpoint("stale.example.com", endpoint.RecordTypeA, "10.0.0.1").WithRefObject(crdRef("update-blocked")),
			endpoint.NewEndpoint("calm.example.com", endpoint.RecordTypeA, "10.0.0.4").WithRefObject(crdRef("in-sync")),
		}

		p := (&Plan{
			Policies:       []Policy{&CreateOnlyPolicy{}},
			Current:        current,
			Desired:        desired,
			ManagedRecords: []string{endpoint.RecordTypeA},
			OwnerID:        "me",
		}).Calculate()

		objects := p.PlannedObjects()
		assert.Equal(t, PlannedObject{Ref: crdRef("update-blocked")}, find(objects, "update-blocked"))
		assert.Equal(t, PlannedObject{Ref: crdRef("in-sync"), Endpoints: 1}, find(objects, "in-sync"))
	})
}
