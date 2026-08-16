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

package source

import (
	"context"

	"sigs.k8s.io/external-dns/pkg/events"
)

// PlannedObject is an object that contributed endpoints to a sync.
type PlannedObject struct {
	Ref *events.ObjectReference
	// Endpoints counts this object's endpoints that survived --domain-filter and
	// the managed record types. Zero means none reached the provider.
	Endpoints int
}

// StatusReporter is implemented by sources that write the outcome of a reconcile
// back onto the objects their endpoints came from. The provider's verdict is only
// known after the plan has run, so this cannot live in Source.Endpoints.
type StatusReporter interface {
	// ReportStatus records what became of each object's endpoints: applied
	// (applyErr == nil), rejected (applyErr != nil), or never offered (Endpoints == 0).
	//
	// objects covers every object that produced an endpoint this sync, not only the
	// changed ones, and may hold objects from other sources — implementations must
	// ignore those by matching ObjectReference.Source().
	ReportStatus(ctx context.Context, objects []PlannedObject, applyErr error)
}

// StatusReporters returns the sources built from this Config that report status.
// Populated by ByNames.
func (c *Config) StatusReporters() []StatusReporter {
	return c.statusReporters
}
