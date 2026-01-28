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
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/smithy-go/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	smithyhttp "github.com/aws/smithy-go/transport/http"
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

type MockInitializeHandler struct {
	CapturedContext context.Context
}

func (mock *MockInitializeHandler) HandleInitialize(ctx context.Context, _ middleware.InitializeInput) (middleware.InitializeOutput, middleware.Metadata, error) {
	mock.CapturedContext = ctx

	return middleware.InitializeOutput{}, middleware.Metadata{}, nil
}

func Test_InitializedTimedOperationMiddleware(t *testing.T) {
	testContext := context.Background()
	mockInitializeHandler := &MockInitializeHandler{}

	_, _, err := initializeTimedOperationMiddleware.HandleInitialize(testContext, middleware.InitializeInput{}, mockInitializeHandler)
	require.NoError(t, err)

	requestMetrics := middleware.GetStackValue(mockInitializeHandler.CapturedContext, requestMetricsKey{}).(requestMetrics)
	assert.NotNil(t, requestMetrics.StartTime)
}

type MockDeserializeHandler struct {
}

func (mock *MockDeserializeHandler) HandleDeserialize(_ context.Context, _ middleware.DeserializeInput) (middleware.DeserializeOutput, middleware.Metadata, error) {
	return middleware.DeserializeOutput{}, middleware.Metadata{}, nil
}

func Test_ExtractAWSRequestParameters(t *testing.T) {
	testContext := context.Background()
	middleware.WithStackValue(testContext, requestMetricsKey{}, requestMetrics{StartTime: time.Now()})

	mockDeserializeHandler := &MockDeserializeHandler{}

	deserializeInput := middleware.DeserializeInput{
		Request: &smithyhttp.Request{
			Request: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Host:   "example.com",
					Scheme: "HTTPS",
					Path:   "/testPath",
				},
			},
		},
	}
	_, _, err := extractAWSRequestParameters.HandleDeserialize(testContext, deserializeInput, mockDeserializeHandler)
	require.NoError(t, err)
}
