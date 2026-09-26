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
	// Endpoints left after the domain and record-type filters; 0 if none.
	Endpoints int
}

// StatusReporter is implemented by sources that report a sync's outcome on the
// objects they read. The outcome exists only after ApplyChanges, not in Endpoints.
type StatusReporter interface {
	// ReportStatus gets every object behind a desired endpoint, changed or not and
	// from any source; skip foreign ones via ObjectReference.Source(). applyErr is
	// nil on success.
	ReportStatus(ctx context.Context, objects []PlannedObject, applyErr error)
}

// StatusReporters returns the sources built by ByNames that implement StatusReporter.
func (c *Config) StatusReporters() []StatusReporter {
	return c.statusReporters
}
