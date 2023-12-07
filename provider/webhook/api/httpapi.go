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
)

type WebhookServer struct {
	Provider provider.Provider
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
		if err := json.NewDecoder(req.Body).Decode(&changes); err != nil {
			log.Errorf("Failed to decode changes: %v", err)
			w.WriteHeader(http.StatusBadRequest)
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

	pve := []*endpoint.Endpoint{}
	if err := json.NewDecoder(req.Body).Decode(&pve); err != nil {
		log.Errorf("Failed to decode in adjustEndpointsHandler: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set(ContentTypeHeader, MediaTypeFormatAndVersion)
	pve, err := p.Provider.AdjustEndpoints(pve)
	if err != nil {
		log.Errorf("Failed to call adjust endpoints: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(&pve); err != nil {
		log.Errorf("Failed to encode in adjustEndpointsHandler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *WebhookServer) NegotiateHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(ContentTypeHeader, MediaTypeFormatAndVersion)
	json.NewEncoder(w).Encode(p.Provider.GetDomainFilter())
}

// StartHTTPApi starts a HTTP server given any provider.
// the function takes an optional channel as input which is used to signal that the server has started.
// The server will listen on port `providerPort`.
// The server will respond to the following endpoints:
// - / (GET): initialization, negotiates headers and returns the domain filter
// - /records (GET): returns the current records
// - /records (POST): applies the changes
// - /adjustendpoints (POST): executes the AdjustEndpoints method
func StartHTTPApi(provider provider.Provider, startedChan chan struct{}, readTimeout, writeTimeout time.Duration, providerPort string) {
	p := WebhookServer{
		Provider: provider,
	}

	m := http.NewServeMux()
	m.HandleFunc("/", p.NegotiateHandler)
	m.HandleFunc("/records", p.RecordsHandler)
	m.HandleFunc("/adjustendpoints", p.AdjustEndpointsHandler)

	s := &http.Server{
		Addr:         providerPort,
		Handler:      m,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	l, err := net.Listen("tcp", providerPort)
	if err != nil {
		log.Fatal(err)
	}

	if startedChan != nil {
		startedChan <- struct{}{}
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
