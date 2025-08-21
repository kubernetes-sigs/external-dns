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

// ref: https://github.com/linki/instrumented_http/blob/master/client.go

package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/external-dns/pkg/metrics"
)

var (
	RequestDurationMetric = metrics.NewSummaryVecWithOpts(
		prometheus.SummaryOpts{
			Name:        "request_duration_seconds",
			Help:        "The HTTP request latencies in seconds.",
			Subsystem:   "http",
			ConstLabels: prometheus.Labels{"handler": "instrumented_http"},
			Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{metrics.LabelScheme, metrics.LabelHost, metrics.LabelPath, metrics.LabelMethod, metrics.LabelStatus},
	)
)

func init() {
	metrics.RegisterMetric.MustRegister(RequestDurationMetric)
}

type CustomRoundTripper struct {
	next http.RoundTripper
}

// CancelRequest is a no-op to satisfy interfaces that require it.
// https://github.com/kubernetes/client-go/blob/34f52c14eaed7e50c845cc14e85e1c4c91e77470/transport/transport.go#L346
func (r *CustomRoundTripper) CancelRequest(_ *http.Request) {
}

func (r *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := r.next.RoundTrip(req)

	status := ""
	if resp != nil {
		status = fmt.Sprintf("%d", resp.StatusCode)
	}

	RequestDurationMetric.SetWithLabels(time.Since(start).Seconds(), metrics.Labels{
		metrics.LabelScheme: req.URL.Scheme,
		metrics.LabelHost:   req.URL.Host,
		metrics.LabelPath:   metrics.PathProcessor(req.URL.Path),
		metrics.LabelMethod: req.Method,
		metrics.LabelStatus: status,
	})

	return resp, err
}

func NewInstrumentedClient(next *http.Client) *http.Client {
	if next == nil {
		next = http.DefaultClient
	}

	next.Transport = NewInstrumentedTransport(next.Transport)

	return next
}

func NewInstrumentedTransport(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}

	return &CustomRoundTripper{next: next}
}
