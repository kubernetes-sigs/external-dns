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

	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
)

type mockEventEmitter struct {
	events []events.Event
}

func (m *mockEventEmitter) Add(events ...events.Event) {
	for _, event := range events {
		m.events = append(m.events, event)
	}

}

func TestEmit(t *testing.T) {
	// emitter := &mockEventEmitter{}
	// obj := &struct{}{} // dummy object

	// change := plan.Change{
	// 	Ref: obj,
	// }
	// changes := plan.Changes{
	// 	Create:    []plan.Change{change},
	// 	UpdateNew: []plan.Change{change},
	// 	Delete:    []plan.Change{change},
	// }
	//
	// emitChangeEvent(emitter, changes, events.Reason("TestReason"))
	//
	// require.Len(t, emitter.events, 3)
	// require.Equal(t, events.ActionCreate, emitter.events[0].Action)
	// require.Equal(t, events.ActionUpdate, emitter.events[1].Action)
	// require.Equal(t, events.ActionDelete, emitter.events[2].Action)
}

func TestEmit_NilEmitter(t *testing.T) {
	emitChangeEvent(nil, plan.Changes{}, events.Reason("TestReason"))
	// Should not panic or do anything
}
