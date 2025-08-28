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
	"fmt"
	"time"

	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"

	extdnshttp "sigs.k8s.io/external-dns/pkg/http"
	"sigs.k8s.io/external-dns/pkg/metrics"
)

type requestMetrics struct {
	StartTime time.Time
}

type requestMetricsKey struct{}

func getRequestMetric(ctx context.Context) requestMetrics {
	requestMetrics, _ := middleware.GetStackValue(ctx, requestMetricsKey{}).(requestMetrics)
	return requestMetrics
}

func setRequestMetric(ctx context.Context, requestMetrics requestMetrics) context.Context {
	return middleware.WithStackValue(ctx, requestMetricsKey{}, requestMetrics)
}

var initializeTimedOperationMiddleware = middleware.InitializeMiddlewareFunc("timedOperation", func(
	ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler,
) (middleware.InitializeOutput, middleware.Metadata, error) {
	requestMetrics := requestMetrics{}
	requestMetrics.StartTime = time.Now()
	ctx = setRequestMetric(ctx, requestMetrics)

	return next.HandleInitialize(ctx, in)
})

var extractAWSRequestParameters = middleware.DeserializeMiddlewareFunc("extractAWSRequestParameters", func(
	ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler,
) (middleware.DeserializeOutput, middleware.Metadata, error) {
	// Call the next middleware first to get the response
	out, metadata, err := next.HandleDeserialize(ctx, in)

	requestMetrics := getRequestMetric(ctx)

	labels := metrics.Labels{}

	if req, ok := in.Request.(*smithyhttp.Request); ok && req != nil {
		labels[metrics.LabelScheme] = req.URL.Scheme
		labels[metrics.LabelHost] = req.URL.Host
		labels[metrics.LabelPath] = metrics.PathProcessor(req.URL.Path)
		labels[metrics.LabelMethod] = req.Method
		labels[metrics.LabelStatus] = "unknown"
	}

	// Try to access HTTP response and status code
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil {
		labels[metrics.LabelStatus] = fmt.Sprintf("%d", resp.StatusCode)
	}

	extdnshttp.RequestDurationMetric.SetWithLabels(time.Since(requestMetrics.StartTime).Seconds(), labels)

	return out, metadata, err
})

func GetInstrumentationMiddlewares() []func(*middleware.Stack) error {
	return []func(s *middleware.Stack) error{
		func(s *middleware.Stack) error {
			if err := s.Initialize.Add(initializeTimedOperationMiddleware, middleware.Before); err != nil {
				return fmt.Errorf("error adding timedOperationMiddleware: %w", err)
			}

			if err := s.Deserialize.Add(extractAWSRequestParameters, middleware.After); err != nil {
				return fmt.Errorf("error adding extractAWSRequestParameters: %w", err)
			}

			return nil
		},
	}
}
