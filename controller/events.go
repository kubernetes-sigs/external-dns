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
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
)

// This function emits events for each change in the provided plan.Changes object using the given EventEmitter.
// It handles create, update, and delete changes, assigning appropriate actions and reasons to each event.
// If the emitter is nil, it does nothing.
func emitChangeEvent(e events.EventEmitter, ch plan.Changes, reason events.Reason) {
	if e == nil {
		return
	}
	for _, ep := range ch.Create {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionCreate, reason))
	}
	for _, ep := range ch.UpdateNew {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionUpdate, reason))
	}
	for _, ep := range ch.Delete {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionDelete, events.RecordDeleted))
	}
}
