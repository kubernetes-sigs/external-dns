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

package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Controller is the interface that wraps the basic event handling methods.
// mock GetRestConfig for testing
var origGetRestConfig = GetRestConfig

func TestNewEventController_Success(t *testing.T) {
	// GetRestConfig = func(kubeConfig, apiServerURL string) (*rest.Config, error) {
	// 	return &rest.Config{}, nil
	// }
	// defer func() { GetRestConfig = origGetRestConfig }()

	cfg := NewConfig()
	ctrl, err := NewEventController(cfg)
	require.NoError(t, err)
	require.NotNil(t, ctrl)
	require.True(t, ctrl.dryRun)
}
