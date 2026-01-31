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
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/tests/integration/toolkit"
)

func TestParseResources(t *testing.T) {
	dir, _ := os.Getwd()
	scenarios, err := toolkit.LoadScenarios(dir)
	require.NoError(t, err, "failed to load scenarios")
	require.NotEmpty(t, scenarios.Scenarios, "no scenarios found")

	assert.Len(t, scenarios.Scenarios, 2, "unexpected number of scenarios")
	scenario := scenarios.Scenarios[0]

	_, err = toolkit.ParseResources(scenario.Resources)
	require.NoError(t, err, "failed to parse resources")
}

func TestSourceIntegration(t *testing.T) {
	// TODO: this is required to ensure annotation parsing works as expected. Ideally, should be set differently.
	annotations.SetAnnotationPrefix(annotations.DefaultAnnotationPrefix)

	dir, _ := os.Getwd()
	scenarios, err := toolkit.LoadScenarios(dir)
	require.NoError(t, err, "failed to load scenarios")
	require.NotEmpty(t, scenarios.Scenarios, "no scenarios found")

	for _, scenario := range scenarios.Scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

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
