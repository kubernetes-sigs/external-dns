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

package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/pkg/events/fake"
	"sigs.k8s.io/external-dns/plan"
)

func TestEmit_RecordReady(t *testing.T) {
	refObj := &events.ObjectReference{}

	tests := []struct {
		name    string
		changes plan.Changes
		asserts func(em *fake.EventEmitter, ch plan.Changes)
	}{
		{
			name: "create, update and delete endpoints",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("one.example.com", endpoint.RecordTypeA, "10.10.10.0").WithRefObject(refObj),
					endpoint.NewEndpoint("two.example.com", endpoint.RecordTypeA, "10.10.10.1").WithRefObject(refObj),
				},
				UpdateNew: []*endpoint.Endpoint{
					endpoint.NewEndpoint("three.example.com", endpoint.RecordTypeA, "10.10.10.2").WithRefObject(refObj),
					endpoint.NewEndpoint("four.example.com", endpoint.RecordTypeA, "10.10.10.3").WithRefObject(refObj),
				},
				Delete: []*endpoint.Endpoint{
					endpoint.NewEndpoint("five.example.com", endpoint.RecordTypeA, "192.10.10.0").WithRefObject(refObj),
				},
			},
			asserts: func(em *fake.EventEmitter, ch plan.Changes) {
				for _, ep := range ch.Create {
					em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ep, events.ActionCreate, events.RecordReady))
				}
				for _, ep := range ch.Delete {
					em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ep, events.ActionDelete, events.RecordDeleted))
				}
				em.AssertNotCalled(t, "Add", mock.MatchedBy(func(e events.Event) bool {
					return e.EventType() == events.EventTypeWarning
				}))
				em.AssertNumberOfCalls(t, "Add", 5)
			},
		},
		{
			name: "delete endpoints",
			changes: plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					endpoint.NewEndpoint("five.example.com", endpoint.RecordTypeA, "192.10.10.0").WithRefObject(refObj),
				},
			},
			asserts: func(em *fake.EventEmitter, ch plan.Changes) {
				for _, ep := range ch.Delete {
					em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ep, events.ActionDelete, events.RecordDeleted))
				}
				em.AssertCalled(t, "Add", mock.MatchedBy(func(e events.Event) bool {
					return e.EventType() == events.EventTypeNormal &&
						e.Action() == events.ActionDelete &&
						e.Reason() == events.RecordDeleted
				}))

				em.AssertNumberOfCalls(t, "Add", 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emitter := fake.NewFakeEventEmitter()

			emitChangeEvent(emitter, &tt.changes, events.RecordReady)

			tt.asserts(emitter, tt.changes)
			mock.AssertExpectationsForObjects(t, emitter)
		})
	}
}

func TestEmit_NilEmitter(t *testing.T) {
	assert.NotPanics(t, func() {
		emitChangeEvent(nil, &plan.Changes{}, events.RecordError)
	})
}

func TestEmit_RecordError(t *testing.T) {
	refObj := &events.ObjectReference{}
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("one.example.com", endpoint.RecordTypeA, "10.10.10.0").
				WithRefObject(refObj),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("two.example.com", endpoint.RecordTypeA, "10.10.10.1").
				WithRefObject(refObj),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("three.example.com", endpoint.RecordTypeA, "10.10.10.2").
				WithRefObject(refObj),
		},
	}
	emitter := fake.NewFakeEventEmitter()

	emitChangeEvent(emitter, &changes, events.RecordError)

	emitter.AssertCalled(t, "Add", events.NewEventFromEndpoint(changes.Create[0], events.ActionCreate, events.RecordError))
	emitter.AssertCalled(t, "Add", events.NewEventFromEndpoint(changes.UpdateNew[0], events.ActionUpdate, events.RecordError))
	emitter.AssertCalled(t, "Add", events.NewEventFromEndpoint(changes.Delete[0], events.ActionDelete, events.RecordError))
	emitter.AssertNumberOfCalls(t, "Add", 3)
}
