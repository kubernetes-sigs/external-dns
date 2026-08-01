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

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source"
)

// emitChangeEvent emits a Kubernetes event for each DNS record change.
// Deletes use RecordDeleted on success and RecordError on failure.
func emitChangeEvent(e events.EventEmitter, ch *plan.Changes, reason events.Reason) {
	if e == nil {
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

// reportSyncStatus tells every status-reporting source the outcome of the apply,
// so the Kubernetes objects that produced the plan can show whether the provider
// programmed them. applyErr is nil on success.
func reportSyncStatus(ctx context.Context, reporters []source.StatusReporter, ch *plan.Changes, applyErr error) {
	if len(reporters) == 0 {
		return
	}

	// Endpoints merged from several objects carry several refs, and one object
	// usually contributes several endpoints, so dedupe by ObjectReference.Key.
	seen := map[string]*events.ObjectReference{}
	for _, endpoints := range [][]*endpoint.Endpoint{ch.Create, ch.UpdateNew, ch.Delete} {
		for _, ep := range endpoints {
			for _, ref := range ep.RefObjects() {
				if ref != nil {
					seen[ref.Key()] = ref
				}
			}
		}
	}
	if len(seen) == 0 {
		return
	}

	refs := make([]*events.ObjectReference, 0, len(seen))
	for _, ref := range seen {
		refs = append(refs, ref)
	}

	for _, reporter := range reporters {
		reporter.ReportStatus(ctx, refs, applyErr)
	}
}
