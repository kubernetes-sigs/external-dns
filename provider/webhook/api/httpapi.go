/*
Copyright 2023 The Kubernetes Authors.

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

package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"

	log "github.com/sirupsen/logrus"
)

const (
	MediaTypeFormatAndVersion = "application/external.dns.webhook+json;version=1"
	ContentTypeHeader         = "Content-Type"
	UrlAdjustEndpoints        = "/adjustendpoints"
	UrlApplyChanges           = "/applychanges"
	UrlRecords                = "/records"
	defaultShutdownTimeout    = 30 * time.Second
)

type WebhookServer struct {
	Provider provider.Provider
	// MaxBodySize caps decoded request body bytes; <= 0 disables the limit.
	MaxBodySize int64
}

// ServerOptions configures the webhook HTTP server started by StartHTTPApi.
type ServerOptions struct {
	Provider provider.Provider
	// StartedChan, if non-nil, is signaled once the server is listening.
	StartedChan chan struct{}
	// ProviderPort is the listen address, e.g. "127.0.0.1:8888".
	ProviderPort      string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
	IdleTimeout       time.Duration
	// ShutdownTimeout bounds request draining; <= 0 defaults to 30 seconds.
	ShutdownTimeout time.Duration
	// MaxBodySize caps request body bytes; <= 0 disables the limit.
	MaxBodySize int64
}

func (p *WebhookServer) RecordsHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		records, err := p.Provider.Records(context.Background())
		if err != nil {
			log.Errorf("Failed to get Records: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set(ContentTypeHeader, MediaTypeFormatAndVersion)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(records); err != nil {
			log.Errorf("Failed to encode records: %v", err)
		}
		return
	case http.MethodPost:
		var changes plan.Changes
		p.limitRequestBody(w, req)
		if err := json.NewDecoder(req.Body).Decode(&changes); err != nil {
			log.Errorf("Failed to decode changes: %v", err)
			w.WriteHeader(requestBodyStatus(err))
			return
		}
		err := p.Provider.ApplyChanges(context.Background(), &changes)
		if err != nil {
			log.Errorf("Failed to apply changes: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	default:
		log.Errorf("Unsupported method %s", req.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (p *WebhookServer) AdjustEndpointsHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Errorf("Unsupported method %s", req.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var pve []*endpoint.Endpoint
	p.limitRequestBody(w, req)
	if err := json.NewDecoder(req.Body).Decode(&pve); err != nil {
		log.Errorf("Failed to decode in adjustEndpointsHandler: %v", err)
		w.WriteHeader(requestBodyStatus(err))
		return
	}
	w.Header().Set(ContentTypeHeader, MediaTypeFormatAndVersion)
	pve, err := p.Provider.AdjustEndpoints(pve)
	if err != nil {
		log.Errorf("Failed to call adjust endpoints: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(&pve); err != nil {
		log.Errorf("Failed to encode in adjustEndpointsHandler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// limitRequestBody caps the request body with MaxBytesReader so oversized
// bodies fail with 413; a non-positive cap leaves the body unchanged.
func (p *WebhookServer) limitRequestBody(w http.ResponseWriter, req *http.Request) {
	if p.MaxBodySize > 0 {
		req.Body = http.MaxBytesReader(w, req.Body, p.MaxBodySize)
	}
}

// requestBodyStatus maps a request-body decode failure to the right client
// error: 413 when the MaxBytesReader cap was hit, 400 for malformed JSON.
func requestBodyStatus(err error) int {
	if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
		return http.StatusRequestEntityTooLarge
	}
	return http.StatusBadRequest
}

func (p *WebhookServer) NegotiateHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(ContentTypeHeader, MediaTypeFormatAndVersion)
	err := json.NewEncoder(w).Encode(p.Provider.GetDomainFilter())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// StartHTTPApi serves the webhook API until a fatal server error occurs.
// Use StartHTTPApiWithContext to support graceful shutdown and handle errors.
func StartHTTPApi(opts ServerOptions) {
	if err := StartHTTPApiWithContext(context.Background(), opts); err != nil {
		log.Fatal(err)
	}
}

// StartHTTPApiWithContext serves the webhook API until ctx is canceled or the
// server fails. Cancellation drains active requests for up to ShutdownTimeout.
// It returns nil on graceful cancellation and an error on startup, serving, or
// shutdown failure. After a shutdown timeout, remaining connections are closed.
func StartHTTPApiWithContext(ctx context.Context, opts ServerOptions) error {
	if ctx.Err() != nil {
		return nil
	}
	l, err := net.Listen("tcp", opts.ProviderPort)
	if err != nil {
		return fmt.Errorf("listen for webhook HTTP API: %w", err)
	}
	return serveHTTPAPI(ctx, opts, l)
}

func serveHTTPAPI(ctx context.Context, opts ServerOptions, l net.Listener) error {
	p := WebhookServer{
		Provider:    opts.Provider,
		MaxBodySize: opts.MaxBodySize,
	}

	m := http.NewServeMux()
	m.HandleFunc("/", p.NegotiateHandler)
	m.HandleFunc(UrlRecords, p.RecordsHandler)
	m.HandleFunc(UrlAdjustEndpoints, p.AdjustEndpointsHandler)

	s := &http.Server{
		Addr:         opts.ProviderPort,
		Handler:      m,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		// Separate knobs from ReadTimeout, which an operator may raise for slow
		// or large request bodies without loosening these bounds.
		ReadHeaderTimeout: opts.ReadHeaderTimeout,
		IdleTimeout:       opts.IdleTimeout,
	}

	// Also close active connections if graceful draining exceeds its deadline.
	defer func() { _ = s.Close() }()
	serveErr := make(chan error, 1)
	go func() { serveErr <- s.Serve(l) }()

	started := opts.StartedChan
	for {
		select {
		case started <- struct{}{}:
			started = nil
		case err := <-serveErr:
			return httpServeError(err)
		case <-ctx.Done():
			timeout := opts.ShutdownTimeout
			if timeout <= 0 {
				timeout = defaultShutdownTimeout
			}
			// The lifecycle context is already canceled; draining needs a fresh deadline.
			shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), timeout)
			defer cancel()
			if err := s.Shutdown(shutdownCtx); err != nil {
				<-serveErr
				return fmt.Errorf("shut down webhook HTTP API: %w", err)
			}
			return httpServeError(<-serveErr)
		}
	}
}

func httpServeError(err error) error {
	if err == nil || errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return fmt.Errorf("serve webhook HTTP API: %w", err)
}
