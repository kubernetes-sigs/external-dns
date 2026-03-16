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

package integration

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/tests/integration/toolkit"
)

var (
	//go:embed scenarios/tests.yaml
	testsYAML []byte
)

func mustLoadScenarios(t *testing.T) *toolkit.TestScenarios {
	t.Helper()
	testScenarios, err := toolkit.LoadScenarios(testsYAML)
	require.NoError(t, err, "failed to load scenarios")
	require.NotEmpty(t, testScenarios.Scenarios, "no scenarios found")
	return testScenarios
}

func TestParseResources(t *testing.T) {
	scenarios := mustLoadScenarios(t)
	for _, scenario := range scenarios.Scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			parsed, err := toolkit.ParseResources(scenario.Resources)
			require.NoError(t, err, "failed to parse resources")

			totalParsed := len(parsed.Services) + len(parsed.Ingresses) + len(parsed.Pods) + len(parsed.EndpointSlices)
			// Pods and EndpointSlices may be auto-generated from dependencies, so count
			// only the explicitly declared resources when checking nothing was silently dropped.
			explicitResources := 0
			for _, r := range scenario.Resources {
				explicitResources++
				if r.Dependencies != nil && r.Dependencies.Pods != nil {
					// Each Service with pod dependencies generates Pods + one EndpointSlice.
					explicitResources += r.Dependencies.Pods.Replicas + 1
				}
			}
			assert.Equal(t, explicitResources, totalParsed, "parsed resource count does not match declared resources")
		})
	}
}

func TestSourceIntegration(t *testing.T) {
	scenarios := mustLoadScenarios(t)
	for _, scenario := range scenarios.Scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			ctx := t.Context()

			client, err := toolkit.LoadResources(ctx, scenario)
			require.NoError(t, err, "failed to populate resources")

			// Create wrapped source
			wrappedSource, err := toolkit.CreateWrappedSource(ctx, client, scenario.Config)
			require.NoError(t, err, "failed to create wrapped source")

			// Get endpoints
			endpoints, err := wrappedSource.Endpoints(ctx)
			require.NoError(t, err)
			toolkit.ValidateEndpoints(t, endpoints, scenario.Expected)
		})
	}
}
