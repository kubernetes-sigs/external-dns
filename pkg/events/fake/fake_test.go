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

package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/pkg/events"
)

func TestNewFakeEventEmitter(t *testing.T) {
	emitter := NewFakeEventEmitter()

	require.NotNil(t, emitter)
	assert.IsType(t, &EventEmitter{}, emitter)
}

func TestEventEmitter_Add_SingleEvent(t *testing.T) {
	emitter := NewFakeEventEmitter()

	event := events.NewEvent(nil, "test message", events.ActionCreate, events.RecordReady)

	emitter.Add(event)

	emitter.AssertExpectations(t)
}

func TestEventEmitter_Add_MultipleEvents(t *testing.T) {
	emitter := NewFakeEventEmitter()

	event1 := events.NewEvent(nil, "test message 1", events.ActionCreate, events.RecordReady)
	event2 := events.NewEvent(nil, "test message 2", events.ActionUpdate, events.RecordReady)

	// Note: The implementation only processes events[0], so we test that behavior
	emitter.Add(event1, event2)

	emitter.AssertExpectations(t)
}

func TestEventEmitter_Add_WithDifferentEventTypes(t *testing.T) {
	tests := []struct {
		name   string
		action events.Action
		reason events.Reason
	}{
		{
			name:   "create action with RecordReady",
			action: events.ActionCreate,
			reason: events.RecordReady,
		},
		{
			name:   "update action with RecordReady",
			action: events.ActionUpdate,
			reason: events.RecordReady,
		},
		{
			name:   "delete action with RecordDeleted",
			action: events.ActionDelete,
			reason: events.RecordDeleted,
		},
		{
			name:   "failed action with RecordError",
			action: events.ActionFailed,
			reason: events.RecordError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emitter := NewFakeEventEmitter()
			event := events.NewEvent(nil, "test message", tt.action, tt.reason)

			emitter.Add(event)

			emitter.AssertExpectations(t)
		})
	}
}

func TestEventEmitter_Add_VerifyMockCalled(t *testing.T) {
	emitter := &EventEmitter{}

	event := events.NewEvent(nil, "test message", events.ActionCreate, events.RecordReady)

	emitter.On("Add", event).Return()

	emitter.Add(event)

	emitter.AssertExpectations(t)
}

func TestEventEmitter_Add_VerifyMockCalledWithAnyEvent(t *testing.T) {
	emitter := NewFakeEventEmitter()

	event := events.NewEvent(nil, "test message", events.ActionCreate, events.RecordReady)

	emitter.Add(event)

	// NewFakeEventEmitter sets up mock.AnythingOfType, so this should pass
	emitter.AssertExpectations(t)
}

func TestEventEmitter_Add_EmptyEventsPanics(t *testing.T) {
	emitter := NewFakeEventEmitter()

	// The Add method accesses events[0] without checking if events is empty
	// This will panic if called with no arguments
	assert.Panics(t, func() {
		emitter.Add()
	}, "Add() should panic when called with no events")
}
