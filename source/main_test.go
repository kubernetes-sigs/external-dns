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
	"os"
	"testing"

	"k8s.io/client-go/features"
)

func TestMain(m *testing.M) {
	// Disable WatchListClient to prevent 10s timeouts when using fake clients.
	// Since client-go v0.35, WatchListClient is enabled by default, but fake
	// clients don't emit the required bookmark events, causing reflectors to
	// stall for 10 seconds before falling back to the legacy list/watch path.
	// Only disable if it is actually on; pre-v0.35 client-go defaults it to
	// false so this is a no-op there, but makes the intent explicit.
	if features.FeatureGates().Enabled(features.WatchListClient) {
		type featureGatesSetter interface {
			features.Gates
			Set(features.Feature, bool) error
		}
		if gates, ok := features.FeatureGates().(featureGatesSetter); ok {
			_ = gates.Set(features.WatchListClient, false)
		}
	}
	os.Exit(m.Run())
}
