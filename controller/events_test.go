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
	"github.com/stretchr/testify/require"

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
			name: "create, update and labelled delete emit events",
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
					endpoint.NewEndpoint("five.example.com", endpoint.RecordTypeA, "192.10.10.0").
						WithLabel(endpoint.ResourceLabelKey, "ingress/default/five"),
					endpoint.NewEndpoint("six.example.com", endpoint.RecordTypeA, "192.10.10.1"),
				},
			},
			asserts: func(em *fake.EventEmitter, ch plan.Changes) {
				for _, ep := range ch.Create {
					em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ep, events.ActionCreate, events.RecordReady))
				}
				for _, ep := range ch.UpdateNew {
					em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ep, events.ActionUpdate, events.RecordReady))
				}
				deleted := deletedEndpoint{Endpoint: ch.Delete[0], ref: refFromResourceLabel("ingress/default/five")}
				em.AssertCalled(t, "Add", events.NewEventFromEndpoint(deleted, events.ActionDelete, events.RecordDeleted))
				em.AssertNotCalled(t, "Add", mock.MatchedBy(func(e events.Event) bool {
					return e.EventType() == events.EventTypeWarning
				}))
				em.AssertNumberOfCalls(t, "Add", 5)
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

	tests := []struct {
		name    string
		changes plan.Changes
		asserts func(em *fake.EventEmitter, ch plan.Changes)
	}{
		{
			name: "failed changes emit RecordError, including deletes",
			changes: plan.Changes{
				Create: []*endpoint.Endpoint{
					endpoint.NewEndpoint("one.example.com", endpoint.RecordTypeA, "10.10.10.0").WithRefObject(refObj),
				},
				UpdateNew: []*endpoint.Endpoint{
					endpoint.NewEndpoint("two.example.com", endpoint.RecordTypeA, "10.10.10.1").WithRefObject(refObj),
				},
				Delete: []*endpoint.Endpoint{
					endpoint.NewEndpoint("three.example.com", endpoint.RecordTypeA, "10.10.10.2").
						WithLabel(endpoint.ResourceLabelKey, "service/default/three"),
				},
			},
			asserts: func(em *fake.EventEmitter, ch plan.Changes) {
				em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ch.Create[0], events.ActionCreate, events.RecordError))
				em.AssertCalled(t, "Add", events.NewEventFromEndpoint(ch.UpdateNew[0], events.ActionUpdate, events.RecordError))
				deleted := deletedEndpoint{Endpoint: ch.Delete[0], ref: refFromResourceLabel("service/default/three")}
				em.AssertCalled(t, "Add", events.NewEventFromEndpoint(deleted, events.ActionDelete, events.RecordError))
				em.AssertNotCalled(t, "Add", mock.MatchedBy(func(e events.Event) bool {
					return e.Reason() == events.RecordDeleted
				}))
				em.AssertNumberOfCalls(t, "Add", 3)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emitter := fake.NewFakeEventEmitter()

			emitChangeEvent(emitter, &tt.changes, events.RecordError)

			tt.asserts(emitter, tt.changes)
			mock.AssertExpectationsForObjects(t, emitter)
		})
	}
}

func TestRefFromResourceLabel(t *testing.T) {
	tests := []struct {
		resource  string
		kind      string
		namespace string
		name      string
	}{
		{resource: "ingress/default/web", kind: "Ingress", namespace: "default", name: "web"},
		{resource: "crd/team-a/records", kind: "DNSEndpoint", namespace: "team-a", name: "records"},
		{resource: "httproute/gw/api", kind: "HTTPRoute", namespace: "gw", name: "api"},
		{resource: "node/worker-1", kind: "Node", name: "worker-1"},
		{resource: "HTTPProxy/default/proxy", kind: "HTTPProxy", namespace: "default", name: "proxy"},
		{resource: "virtualmachineinstance/vms/vm1", kind: "virtualmachineinstance", namespace: "vms", name: "vm1"},
	}
	for _, tt := range tests {
		t.Run(tt.resource, func(t *testing.T) {
			ref := refFromResourceLabel(tt.resource)
			require.NotNil(t, ref)
			assert.Equal(t, tt.kind, ref.Kind())
			assert.Equal(t, tt.namespace, ref.Namespace())
			assert.Equal(t, tt.name, ref.Name())
			assert.Empty(t, ref.UID())
			assert.Equal(t, registrySource, ref.Source())
		})
	}

	for _, resource := range []string{"", "ingress", "ingress/", "/default/web", "a/b/c/d"} {
		t.Run("invalid "+resource, func(t *testing.T) {
			assert.Nil(t, refFromResourceLabel(resource))
		})
	}
}
