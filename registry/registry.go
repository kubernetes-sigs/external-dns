/*
Copyright 2017 The Kubernetes Authors.

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

package registry

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// FinishFunc completes an apply after the provider call. It receives the provider error and
// returns it, or an error of its own. A provider error must stay reachable with errors.Is,
// so that a provider.SoftError stays soft.
type FinishFunc func(ctx context.Context, providerErr error) error

// Registry tracks ownership of DNS records managed by external-dns.
type Registry interface {
	// Records returns all DNS records known to the registry, including ownership metadata.
	Records(ctx context.Context) ([]*endpoint.Endpoint, error)
	// Plan prepares the changes for the provider: it drops changes to records owned by
	// others, sets the owner label and adds ownership records. It does not change any
	// state outside the given changes, so it also runs in dry-run.
	Plan(changes *plan.Changes) *plan.Changes
	// PreApply writes the registry state that must exist before the provider call.
	// It returns the changes to send to the provider, which can be fewer than given,
	// and the function that completes the apply. It does not run in dry-run.
	PreApply(ctx context.Context, changes *plan.Changes) (*plan.Changes, FinishFunc, error)
	// Provider returns the DNS provider the registry wraps.
	Provider() provider.Provider
	// AdjustEndpoints normalises endpoints before they are processed by the planner.
	AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error)
	// GetDomainFilter returns the domain filter configured for the underlying provider.
	GetDomainFilter() endpoint.DomainFilterInterface
	// OwnerID returns the owner identifier used to claim DNS records.
	OwnerID() string
}

// NoFinish is the FinishFunc of registries with nothing to do after the provider call.
func NoFinish(_ context.Context, providerErr error) error {
	return providerErr
}

// Apply plans the changes with the registry and applies them to its provider.
// In dry-run only Plan and the provider run: the provider logs what it would change,
// and the registry state is left untouched.
func Apply(ctx context.Context, r Registry, changes *plan.Changes, dryRun bool) error {
	changes = r.Plan(changes)
	if dryRun {
		return r.Provider().ApplyChanges(ctx, changes)
	}
	changes, finish, err := r.PreApply(ctx, changes)
	if err != nil {
		return err
	}
	return finish(ctx, r.Provider().ApplyChanges(ctx, changes))
}
