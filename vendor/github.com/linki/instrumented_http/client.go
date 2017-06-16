// Package instrumented_http provides a drop-in metrics-enabled replacement for
// any http.Client or http.RoundTripper.
package instrumented_http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// Transport is a http.RoundTripper that collects Prometheus metrics of every
// request it processes. It allows to be configured with callbacks that process
// request path and query into a suitable label value.
type Transport struct {
	next http.RoundTripper
	cbs  *Callbacks
}

// Callbacks is a collection of callbacks passed to Transport.
type Callbacks struct {
	PathProcessor  func(string) string
	QueryProcessor func(string) string
}

const (
	// Metrics created can be identified by this label value.
	handlerName = "instrumented_http"
)

var (
	// RequestDurationSeconds is a Prometheus summary to collect request times.
	RequestDurationSeconds = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        "request_duration_seconds",
			Help:        "The HTTP request latencies in seconds.",
			Subsystem:   "http",
			ConstLabels: prometheus.Labels{"handler": handlerName},
			Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"scheme", "host", "path", "query", "method", "status"},
	)

	// EliminatingProcessor is a callback that returns a blank string on any input.
	EliminatingProcessor = func(_ string) string { return "" }
	// IdentityProcessor is callback that returns whatever is passed to it.
	IdentityProcessor = func(input string) string { return input }
	// LastPathElementProcessor is callback that returns the last element of a URL path.
	LastPathElementProcessor = func(path string) string {
		parts := strings.Split(path, "/")
		return parts[len(parts)-1]
	}
)

// init registers the Prometheus metric globally when the package is loaded.
func init() {
	prometheus.MustRegister(RequestDurationSeconds)
}

// RoundTrip implements http.RoundTripper. It forwards the request to the
// next RoundTripper and measures the time it took in Prometheus summary.
func (it *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var statusCode int

	// Remember the current time.
	now := time.Now()

	// Make the request using the next RoundTripper.
	resp, err := it.next.RoundTrip(req)
	if resp != nil {
		statusCode = resp.StatusCode
	}

	// Observe the time it took to make the request.
	RequestDurationSeconds.WithLabelValues(
		req.URL.Scheme,
		req.URL.Host,
		it.cbs.PathProcessor(req.URL.Path),
		it.cbs.QueryProcessor(req.URL.RawQuery),
		req.Method,
		fmt.Sprintf("%d", statusCode),
	).Observe(time.Since(now).Seconds())

	// return the response and error reported from the next RoundTripper.
	return resp, err
}

// NewClient takes a *http.Client and returns a *http.Client that has its
// RoundTripper wrapped with instrumentation. Optionally, It can receive a
// collection of callbacks that process request path and query into a suitable
// label value.
func NewClient(next *http.Client, cbs *Callbacks) *http.Client {
	// If next client is not defined we'll use http.DefaultClient.
	if next == nil {
		next = http.DefaultClient
	}

	next.Transport = NewTransport(next.Transport, cbs)

	return next
}

// NewTransport takes a http.RoundTripper, wraps it with instrumentation and
// returns it as a new http.RoundTripper. Optionally, It can receive a
// collection of callbacks that process request path and query into a suitable
// label value.
func NewTransport(next http.RoundTripper, cbs *Callbacks) http.RoundTripper {
	// If next RoundTripper is not defined we'll use http.DefaultTransport.
	if next == nil {
		next = http.DefaultTransport
	}

	// If cbs is not defined we'll initilialize it with defaults.
	if cbs == nil {
		cbs = &Callbacks{}
	}
	// By default, path and query will be ignored.
	if cbs.PathProcessor == nil {
		cbs.PathProcessor = EliminatingProcessor
	}
	if cbs.QueryProcessor == nil {
		cbs.QueryProcessor = EliminatingProcessor
	}

	return &Transport{next: next, cbs: cbs}
}
