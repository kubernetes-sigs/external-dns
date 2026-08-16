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
	"context"

	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source"
)

// emitChangeEvent emits a Kubernetes event for each DNS record change.
// Deletes use RecordDeleted on success and RecordError on failure.
func emitChangeEvent(e events.EventEmitter, ch *plan.Changes, reason events.Reason) {
	// events.Discard is the default (--events-emit selected nothing). Building an
	// Event per change formats a message that is then thrown away.
	if e == nil || e == events.Discard {
		return
	}
	for _, ep := range ch.Create {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionCreate, reason))
	}
	for _, ep := range ch.UpdateNew {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionUpdate, reason))
	}
	deleteReason := events.RecordDeleted
	if reason == events.RecordError {
		deleteReason = events.RecordError
	}
	for _, ep := range ch.Delete {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionDelete, deleteReason))
	}
}

// reportSyncStatus tells every status-reporting source the outcome of the sync.
// It covers every object that contributed a desired endpoint, not just the ones
// whose records changed, so an already-in-sync object still gets a status.
// applyErr is nil on success.
func reportSyncStatus(ctx context.Context, reporters []source.StatusReporter, p *plan.Plan, applyErr error) {
	if len(reporters) == 0 {
		return
	}

	// Endpoints merged from several objects carry several refs, and one object
	// usually contributes several endpoints, so dedupe by ObjectReference.Key.
	// Map order is unstable, hence the separate key list.
	byKey := map[string]*source.PlannedObject{}
	var keys []string
	for _, ep := range p.Desired {
		for _, ref := range ep.RefObjects() {
			if ref == nil {
				continue
			}
			if _, ok := byKey[ref.Key()]; !ok {
				byKey[ref.Key()] = &source.PlannedObject{Ref: ref}
				keys = append(keys, ref.Key())
			}
		}
	}
	if len(byKey) == 0 {
		return
	}

	// Filtered-out endpoints never reach p.Planned, so their objects keep a count
	// of zero and their source reports them as filtered rather than programmed.
	for _, ep := range p.Planned {
		for _, ref := range ep.RefObjects() {
			if ref == nil {
				continue
			}
			if obj, ok := byKey[ref.Key()]; ok {
				obj.Endpoints++
			}
		}
	}

	objects := make([]source.PlannedObject, 0, len(keys))
	for _, key := range keys {
		objects = append(objects, *byKey[key])
	}

	for _, reporter := range reporters {
		reporter.ReportStatus(ctx, objects, applyErr)
	}
}
