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

package registry

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/fakes"
)

// recordingProvider adds its ApplyChanges call to the registry's call log.
type recordingProvider struct {
	fakes.MockProvider
	calls *[]string
	got   *plan.Changes
}

func (p *recordingProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	*p.calls = append(*p.calls, "provider")
	p.got = changes
	return p.ApplyChangesErr
}

// fakeRegistry logs the phases in the order Apply runs them.
type fakeRegistry struct {
	provider    *recordingProvider
	planned     *plan.Changes
	calls       []string
	preApplyErr error
	finishErr   error
	finishGot   error
}

func newFakeRegistry() *fakeRegistry {
	r := &fakeRegistry{
		planned: &plan.Changes{
			Create: []*endpoint.Endpoint{endpoint.NewEndpoint("planned.example.com", endpoint.RecordTypeA, "1.2.3.4")},
		},
	}
	r.provider = &recordingProvider{calls: &r.calls}
	return r
}

func (r *fakeRegistry) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (r *fakeRegistry) Plan(_ *plan.Changes) *plan.Changes {
	r.calls = append(r.calls, "plan")
	return r.planned
}

func (r *fakeRegistry) PreApply(_ context.Context, changes *plan.Changes) (*plan.Changes, FinishFunc, error) {
	r.calls = append(r.calls, "preApply")
	if r.preApplyErr != nil {
		return nil, nil, r.preApplyErr
	}
	finish := func(_ context.Context, providerErr error) error {
		r.calls = append(r.calls, "finish")
		r.finishGot = providerErr
		if r.finishErr != nil {
			return r.finishErr
		}
		return providerErr
	}
	return changes, finish, nil
}

func (r *fakeRegistry) Provider() provider.Provider {
	return r.provider
}

func (r *fakeRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return endpoints, nil
}

func (r *fakeRegistry) GetDomainFilter() endpoint.DomainFilterInterface {
	return r.provider.GetDomainFilter()
}

func (r *fakeRegistry) OwnerID() string {
	return "owner"
}

func TestApplyRunsPhasesInOrder(t *testing.T) {
	r := newFakeRegistry()

	require.NoError(t, Apply(t.Context(), r, &plan.Changes{}, false))

	assert.Equal(t, []string{"plan", "preApply", "provider", "finish"}, r.calls)
	assert.Same(t, r.planned, r.provider.got, "the provider gets the planned changes")
	require.NoError(t, r.finishGot)
}

func TestApplyDryRunOnlyPlansAndCallsProvider(t *testing.T) {
	r := newFakeRegistry()

	require.NoError(t, Apply(t.Context(), r, &plan.Changes{}, true))

	assert.Equal(t, []string{"plan", "provider"}, r.calls)
	assert.Same(t, r.planned, r.provider.got, "the provider gets the planned changes in dry-run too")
}

func TestApplyPreApplyErrorSkipsProvider(t *testing.T) {
	r := newFakeRegistry()
	r.preApplyErr = errors.New("pre-apply failed")

	err := Apply(t.Context(), r, &plan.Changes{}, false)

	require.ErrorIs(t, err, r.preApplyErr)
	assert.Equal(t, []string{"plan", "preApply"}, r.calls)
}

func TestApplyProviderErrorReachesFinishAndStaysSoft(t *testing.T) {
	r := newFakeRegistry()
	r.provider.ApplyChangesErr = provider.NewSoftError(errors.New("provider failed"))

	err := Apply(t.Context(), r, &plan.Changes{}, false)

	require.ErrorIs(t, err, provider.SoftError)
	assert.Equal(t, []string{"plan", "preApply", "provider", "finish"}, r.calls)
	require.ErrorIs(t, r.finishGot, r.provider.ApplyChangesErr)
}

func TestApplyReturnsFinishError(t *testing.T) {
	r := newFakeRegistry()
	r.finishErr = errors.New("finish failed")

	err := Apply(t.Context(), r, &plan.Changes{}, false)

	require.ErrorIs(t, err, r.finishErr)
}

func TestNoFinishReturnsProviderError(t *testing.T) {
	providerErr := errors.New("provider failed")

	require.NoError(t, NoFinish(t.Context(), nil))
	require.ErrorIs(t, NoFinish(t.Context(), providerErr), providerErr)
}
