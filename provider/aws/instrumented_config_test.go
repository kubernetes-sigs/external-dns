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

package aws

import (
	"testing"

	"github.com/aws/smithy-go/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetInstrumentationMiddlewares(t *testing.T) {
	t.Run("adds expected middlewares", func(t *testing.T) {
		stack := middleware.NewStack("test-stack", nil)

		for _, mw := range GetInstrumentationMiddlewares() {
			err := mw(stack)
			require.NoError(t, err)
		}

		// Check Initialize stage
		timedOperationMiddleware, found := stack.Initialize.Get("timedOperation")
		assert.True(t, found, "timedOperation middleware should be present in Initialize stage")
		assert.NotNil(t, timedOperationMiddleware)

		// Check Deserialize stage
		extractAWSRequestParametersMiddleware, found := stack.Deserialize.Get("extractAWSRequestParameters")
		assert.True(t, found, "extractAWSRequestParameters middleware should be present in Deserialize stage")
		assert.NotNil(t, extractAWSRequestParametersMiddleware)
	})
}
