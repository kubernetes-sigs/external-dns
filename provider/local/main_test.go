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

package main

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miekg/dns"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider/webhook/api"
)

func TestNewDNSRecordStore(t *testing.T) {
	store := NewDNSRecordStore()
	if store == nil {
		t.Fatal("NewDNSRecordStore returned nil")
	}
	if store.domains == nil {
		t.Fatal("domains map not initialized")
	}
}

func TestDNSRecordStore_AddRecord(t *testing.T) {
	store := NewDNSRecordStore()

	// Test adding a simple A record
	store.AddRecord("example.com", "A", []string{"192.168.1.1"})

	targets := store.GetRecord("example.com", "A")
	if len(targets) != 1 || targets[0] != "192.168.1.1" {
		t.Errorf("Expected [192.168.1.1], got %v", targets)
	}

	// Test adding multiple targets
	store.AddRecord("example.com", "A", []string{"192.168.1.1", "192.168.1.2"})
	targets = store.GetRecord("example.com", "A")
	if len(targets) != 2 {
		t.Errorf("Expected 2 targets, got %d", len(targets))
	}

	// Test adding different record type for same domain
	store.AddRecord("example.com", "AAAA", []string{"::1"})
	ipv6Targets := store.GetRecord("example.com", "AAAA")
	if len(ipv6Targets) != 1 || ipv6Targets[0] != "::1" {
		t.Errorf("Expected [::1], got %v", ipv6Targets)
	}

	// Ensure A records still exist
	targets = store.GetRecord("example.com", "A")
	if len(targets) != 2 {
		t.Errorf("A records should still exist, got %d", len(targets))
	}
}

func TestDNSRecordStore_RemoveRecord(t *testing.T) {
	store := NewDNSRecordStore()

	// Add some records
	store.AddRecord("example.com", "A", []string{"192.168.1.1"})
	store.AddRecord("example.com", "AAAA", []string{"::1"})
	store.AddRecord("test.com", "A", []string{"10.0.0.1"})

	// Remove one record type
	store.RemoveRecord("example.com", "A")
	targets := store.GetRecord("example.com", "A")
	if len(targets) != 0 {
		t.Errorf("A record should be removed, got %v", targets)
	}

	// Ensure other records still exist
	ipv6Targets := store.GetRecord("example.com", "AAAA")
	if len(ipv6Targets) != 1 {
		t.Errorf("AAAA record should still exist, got %v", ipv6Targets)
	}

	testTargets := store.GetRecord("test.com", "A")
	if len(testTargets) != 1 {
		t.Errorf("test.com A record should still exist, got %v", testTargets)
	}

	// Remove last record for domain
	store.RemoveRecord("example.com", "AAAA")

	// Domain should be completely removed
	if store.domains["example.com"] != nil {
		t.Error("Domain should be removed when no records remain")
	}
}

func TestDNSRecordStore_GetRecord_NotFound(t *testing.T) {
	store := NewDNSRecordStore()

	// Test getting non-existent record
	targets := store.GetRecord("nonexistent.com", "A")
	if targets != nil {
		t.Errorf("Expected nil for non-existent record, got %v", targets)
	}

	// Test getting non-existent record type for existing domain
	store.AddRecord("example.com", "A", []string{"192.168.1.1"})
	targets = store.GetRecord("example.com", "CNAME")
	if targets != nil {
		t.Errorf("Expected nil for non-existent record type, got %v", targets)
	}
}

func TestDNSRecordStore_GetAllRecords(t *testing.T) {
	store := NewDNSRecordStore()

	// Test empty store
	endpoints := store.GetAllRecords()
	if len(endpoints) != 0 {
		t.Errorf("Expected 0 endpoints for empty store, got %d", len(endpoints))
	}

	// Add some records
	store.AddRecord("example.com", "A", []string{"192.168.1.1", "192.168.1.2"})
	store.AddRecord("example.com", "AAAA", []string{"::1"})
	store.AddRecord("test.com", "CNAME", []string{"example.com"})

	endpoints = store.GetAllRecords()
	if len(endpoints) != 3 {
		t.Errorf("Expected 3 endpoints, got %d", len(endpoints))
	}

	// Verify endpoint content
	endpointMap := make(map[string]endpoint.Endpoint)
	for _, ep := range endpoints {
		key := ep.DNSName + ":" + ep.RecordType
		endpointMap[key] = ep
	}

	if ep, exists := endpointMap["example.com:A"]; !exists {
		t.Error("Missing example.com A record")
	} else if len(ep.Targets) != 2 {
		t.Errorf("Expected 2 targets for A record, got %d", len(ep.Targets))
	}

	if ep, exists := endpointMap["example.com:AAAA"]; !exists {
		t.Error("Missing example.com AAAA record")
	} else if len(ep.Targets) != 1 || ep.Targets[0] != "::1" {
		t.Errorf("Expected [::1] for AAAA record, got %v", ep.Targets)
	}

	if ep, exists := endpointMap["test.com:CNAME"]; !exists {
		t.Error("Missing test.com CNAME record")
	} else if len(ep.Targets) != 1 || ep.Targets[0] != "example.com" {
		t.Errorf("Expected [example.com] for CNAME record, got %v", ep.Targets)
	}
}

func TestNegotiateHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedHeader string
	}{
		{
			name:           "Valid GET request",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedHeader: api.MediaTypeFormatAndVersion,
		},
		{
			name:           "Invalid POST request",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedHeader: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()

			negotiateHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			if tt.expectedHeader != "" {
				contentType := res.Header.Get("Content-Type")
				if contentType != tt.expectedHeader {
					t.Errorf("Expected header %s, got %s", tt.expectedHeader, contentType)
				}

				defer res.Body.Close()
				var domainFilter endpoint.DomainFilter
				err := json.NewDecoder(res.Body).Decode(&domainFilter)
				if err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
			}
		})
	}
}

func TestRecordsHandler_GET(t *testing.T) {
	store := NewDNSRecordStore()
	store.AddRecord("example.com", "A", []string{"192.168.1.1"})
	store.AddRecord("test.com", "CNAME", []string{"example.com"})

	req := httptest.NewRequest("GET", "/records", nil)
	w := httptest.NewRecorder()

	recordsHandler(w, req, store)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var endpoints []endpoint.Endpoint
	if err := json.Unmarshal(w.Body.Bytes(), &endpoints); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(endpoints) != 2 {
		t.Errorf("Expected 2 endpoints, got %d", len(endpoints))
	}
}

func TestRecordsHandler_POST_Create(t *testing.T) {
	store := NewDNSRecordStore()

	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: "A",
				Targets:    []string{"192.168.1.1"},
			},
			{
				DNSName:    "test.com",
				RecordType: "CNAME",
				Targets:    []string{"example.com"},
			},
		},
	}

	body, _ := json.Marshal(changes)
	req := httptest.NewRequest("POST", "/records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	recordsHandler(w, req, store)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify records were created
	targets := store.GetRecord("example.com", "A")
	if len(targets) != 1 || targets[0] != "192.168.1.1" {
		t.Errorf("A record not created correctly, got %v", targets)
	}

	targets = store.GetRecord("test.com", "CNAME")
	if len(targets) != 1 || targets[0] != "example.com" {
		t.Errorf("CNAME record not created correctly, got %v", targets)
	}
}

func TestRecordsHandler_POST_Delete(t *testing.T) {
	store := NewDNSRecordStore()
	store.AddRecord("example.com", "A", []string{"192.168.1.1"})
	store.AddRecord("test.com", "CNAME", []string{"example.com"})

	changes := plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: "A",
			},
		},
	}

	body, _ := json.Marshal(changes)
	req := httptest.NewRequest("POST", "/records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	recordsHandler(w, req, store)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify record was deleted
	targets := store.GetRecord("example.com", "A")
	if len(targets) != 0 {
		t.Errorf("A record should be deleted, got %v", targets)
	}

	// Verify other record still exists
	targets = store.GetRecord("test.com", "CNAME")
	if len(targets) != 1 {
		t.Errorf("CNAME record should still exist, got %v", targets)
	}
}

func TestRecordsHandler_POST_Update(t *testing.T) {
	store := NewDNSRecordStore()
	store.AddRecord("example.com", "A", []string{"192.168.1.1"})

	changes := plan.Changes{
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: "A",
				Targets:    []string{"192.168.1.1"},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: "A",
				Targets:    []string{"192.168.1.2", "192.168.1.3"},
			},
		},
	}

	body, _ := json.Marshal(changes)
	req := httptest.NewRequest("POST", "/records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	recordsHandler(w, req, store)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", w.Code)
	}

	// Verify record was updated
	targets := store.GetRecord("example.com", "A")
	if len(targets) != 2 {
		t.Errorf("Expected 2 targets after update, got %d", len(targets))
	}

	expectedTargets := map[string]bool{"192.168.1.2": true, "192.168.1.3": true}
	for _, target := range targets {
		if !expectedTargets[target] {
			t.Errorf("Unexpected target after update: %s", target)
		}
	}
}

func TestRecordsHandler_InvalidMethod(t *testing.T) {
	store := NewDNSRecordStore()

	req := httptest.NewRequest("PUT", "/records", nil)
	w := httptest.NewRecorder()

	recordsHandler(w, req, store)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestHandleDNSQuery_ARecord(t *testing.T) {
	store := NewDNSRecordStore()
	store.AddRecord("example.com", "A", []string{"192.168.1.1", "192.168.1.2"})

	// Create DNS query
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn("example.com"), dns.TypeA)

	// Create test response writer
	responseWriter := &testResponseWriter{}

	// Handle the query
	handleDNSQuery(store, responseWriter, m, 300)

	// Verify response
	response := responseWriter.msg
	if response == nil {
		t.Fatal("No response received")
	}

	if len(response.Answer) != 2 {
		t.Errorf("Expected 2 answers, got %d", len(response.Answer))
	}

	if response.Authoritative != true {
		t.Error("Response should be authoritative")
	}

	// Check answer records
	ips := make(map[string]bool)
	for _, rr := range response.Answer {
		if a, ok := rr.(*dns.A); ok {
			ips[a.A.String()] = true
			if a.Hdr.Ttl != 300 {
				t.Errorf("Expected TTL 300, got %d", a.Hdr.Ttl)
			}
		} else {
			t.Errorf("Expected A record, got %T", rr)
		}
	}

	if !ips["192.168.1.1"] || !ips["192.168.1.2"] {
		t.Error("Expected both IP addresses in response")
	}
}

func TestHandleDNSQuery_AAAARecord(t *testing.T) {
	store := NewDNSRecordStore()
	store.AddRecord("example.com", "AAAA", []string{"2001:db8::1"})

	// Create DNS query
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn("example.com"), dns.TypeAAAA)

	// Create test response writer
	responseWriter := &testResponseWriter{}

	// Handle the query
	handleDNSQuery(store, responseWriter, m, 300)

	// Verify response
	response := responseWriter.msg
	if response == nil {
		t.Fatal("No response received")
	}

	if len(response.Answer) != 1 {
		t.Errorf("Expected 1 answer, got %d", len(response.Answer))
	}

	if aaaa, ok := response.Answer[0].(*dns.AAAA); ok {
		if aaaa.AAAA.String() != "2001:db8::1" {
			t.Errorf("Expected 2001:db8::1, got %s", aaaa.AAAA.String())
		}
	} else {
		t.Errorf("Expected AAAA record, got %T", response.Answer[0])
	}
}

func TestHandleDNSQuery_CNAMERecord(t *testing.T) {
	store := NewDNSRecordStore()
	store.AddRecord("www.example.com", "CNAME", []string{"example.com"})

	// Create DNS query
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn("www.example.com"), dns.TypeCNAME)

	// Create test response writer
	responseWriter := &testResponseWriter{}

	// Handle the query
	handleDNSQuery(store, responseWriter, m, 300)

	// Verify response
	response := responseWriter.msg
	if response == nil {
		t.Fatal("No response received")
	}

	if len(response.Answer) != 1 {
		t.Errorf("Expected 1 answer, got %d", len(response.Answer))
	}

	if cname, ok := response.Answer[0].(*dns.CNAME); ok {
		if cname.Target != "example.com." {
			t.Errorf("Expected example.com., got %s", cname.Target)
		}
	} else {
		t.Errorf("Expected CNAME record, got %T", response.Answer[0])
	}
}

func TestHandleDNSQuery_NXDOMAIN(t *testing.T) {
	store := NewDNSRecordStore()

	// Create DNS query for non-existent domain
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn("nonexistent.com"), dns.TypeA)

	// Create test response writer
	responseWriter := &testResponseWriter{}

	// Handle the query
	handleDNSQuery(store, responseWriter, m, 300)

	// Verify response
	response := responseWriter.msg
	if response == nil {
		t.Fatal("No response received")
	}

	if response.Rcode != dns.RcodeNameError {
		t.Errorf("Expected NXDOMAIN (rcode %d), got %d", dns.RcodeNameError, response.Rcode)
	}

	if len(response.Answer) != 0 {
		t.Errorf("Expected 0 answers for NXDOMAIN, got %d", len(response.Answer))
	}
}

func TestHealthzHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	healthzHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	if w.Body.String() != "ok" {
		t.Errorf("Expected 'ok', got '%s'", w.Body.String())
	}
}

func TestAdjustEndpointsHandler(t *testing.T) {
	endpoints := []endpoint.Endpoint{
		{
			DNSName:    "test.example.com",
			RecordType: "A",
			Targets:    []string{"192.168.1.1", "192.168.1.2"},
		},
	}

	endpointsJSON, err := json.Marshal(endpoints)
	if err != nil {
		t.Fatalf("Failed to marshal endpoints: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/adjustendpoints", bytes.NewReader(endpointsJSON))
	w := httptest.NewRecorder()

	adjustEndpointsHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	var returnedEndpoints []endpoint.Endpoint
	if err := json.Unmarshal(w.Body.Bytes(), &returnedEndpoints); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if len(returnedEndpoints) != 1 {
		t.Errorf("Expected 1 endpoint returned, got %d", len(returnedEndpoints))
	}
}

// testResponseWriter implements dns.ResponseWriter for testing
type testResponseWriter struct {
	msg *dns.Msg
}

func (w *testResponseWriter) LocalAddr() net.Addr {
	return nil
}

func (w *testResponseWriter) RemoteAddr() net.Addr {
	return nil
}

func (w *testResponseWriter) WriteMsg(m *dns.Msg) error {
	w.msg = m
	return nil
}

func (w *testResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w *testResponseWriter) Close() error {
	return nil
}

func (w *testResponseWriter) TsigStatus() error {
	return nil
}

func (w *testResponseWriter) TsigTimersOnly(bool) {}

func (w *testResponseWriter) Hijack() {}
