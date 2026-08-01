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

// StatusReporter is implemented by sources that write the outcome of a reconcile
// back onto the Kubernetes objects their endpoints came from. A source only
// learns the provider's verdict after the plan has run, so this cannot live in
// Source.Endpoints; the controller calls it after ApplyChanges.
type StatusReporter interface {
	// ReportStatus records that the changes derived from refs were applied
	// (applyErr == nil) or rejected (applyErr != nil) by the provider. refs may
	// hold references from other sources; implementations must ignore those,
	// typically by matching ObjectReference.Source().
	ReportStatus(ctx context.Context, refs []*events.ObjectReference, applyErr error)
}

// StatusReporters returns the sources built from this Config that report status
// back to Kubernetes. Populated by ByNames.
func (c *Config) StatusReporters() []StatusReporter {
	return c.statusReporters
}
