/*
Copyright 2017 The Kubernetes Authors.

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

package pihole

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

func newTestServerV6(t *testing.T, hdlr http.HandlerFunc) *httptest.Server {
	t.Helper()
	svr := httptest.NewServer(hdlr)
	return svr
}

func TestNewPiholeClientV6(t *testing.T) {
	// Test correct error on no server provided
	_, err := newPiholeClientV6(PiholeConfig{APIVersion: "6"})
	if err == nil {
		t.Error("Expected error from config with no server")
	} else if err != ErrNoPiholeServer {
		t.Error("Expected ErrNoPiholeServer, got", err)
	}

	// Test new client with no password. Should create the client cleanly.
	cl, err := newPiholeClientV6(PiholeConfig{
		Server:     "test",
		APIVersion: "6",
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cl.(*piholeClientV6); !ok {
		t.Error("Did not create a new pihole client")
	}

	// Create a test server
	srvr := newTestServerV6(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/auth" && r.Method == "POST" {
			var requestData map[string]string
			json.NewDecoder(r.Body).Decode(&requestData)
			defer r.Body.Close()

			w.Header().Set("Content-Type", "application/json")

			if requestData["password"] != "correct" {
				// Return unsuccessful authentication response
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{
				"session": {
					"valid": false,
					"totp": false,
					"sid": null,
					"validity": -1,
					"message": "password incorrect"
				},
				"took": 0.2
			}`))
				return
			}

			// Return successful authentication response
			w.Write([]byte(`{
			"session": {
				"valid": true,
				"totp": false,
				"sid": "supersecret",
				"csrf": "csrfvalue",
				"validity": 1800,
				"message": "password correct"
			},
			"took": 0.23066902160644531
		}`))
		} else {
			http.NotFound(w, r)
		}
	})
	defer srvr.Close()

	// Test invalid password
	_, err = newPiholeClientV6(
		PiholeConfig{Server: srvr.URL, APIVersion: "6", Password: "wrong"},
	)
	if err == nil {
		t.Error("Expected error for creating client with invalid password")
	}

	// Test correct password
	cl, err = newPiholeClientV6(
		PiholeConfig{Server: srvr.URL, APIVersion: "6", Password: "correct"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if cl.(*piholeClientV6).token != "supersecret" {
		t.Error("Parsed invalid token from login response:", cl.(*piholeClient).token)
	}
}

func TestListRecordsV6(t *testing.T) {
	// Create a test server
	srvr := newTestServerV6(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/config/dns/hosts" && r.Method == "GET" {
			var requestData map[string]string
			json.NewDecoder(r.Body).Decode(&requestData)
			defer r.Body.Close()

			w.Header().Set("Content-Type", "application/json")

			// Return A records
			w.Write([]byte(`{
				"config": {
					"dns": {
						"hosts": [
							"192.168.178.33 service1.example.com",
							"192.168.178.34 service2.example.com",
							"192.168.178.34 service3.example.com",
							"fc00::1:192:168:1:1 service4.example.com",
							"fc00::1:192:168:1:2 service5.example.com",
							"fc00::1:192:168:1:3 service6.example.com"
						]
					}
				},
				"took": 5
			}`))
			w.WriteHeader(http.StatusOK)
		} else if r.URL.Path == "/api/config/dns/cnameRecords" && r.Method == "GET" {
			var requestData map[string]string
			json.NewDecoder(r.Body).Decode(&requestData)
			defer r.Body.Close()

			w.Header().Set("Content-Type", "application/json")

			// Return A records
			w.Write([]byte(`{
				"config": {
					"dns": {
						"cnameRecords": [
							"source1.example.com,target1.domain.com,1000",
							"source2.example.com,target2.domain.com,50",
							"source3.example.com,target3.domain.com"
						]
					}
				},
				"took": 5
			}`))
			w.WriteHeader(http.StatusOK)
		} else {
			http.NotFound(w, r)
		}
	})
	defer srvr.Close()

	// Create a client
	cfg := PiholeConfig{
		Server:     srvr.URL,
		APIVersion: "6",
	}
	cl, err := newPiholeClientV6(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Test retrieve A records unfiltered
	arecs, err := cl.listRecords(context.Background(), endpoint.RecordTypeA)
	if err != nil {
		t.Fatal(err)
	}
	if len(arecs) != 3 {
		t.Fatal("Expected 3 A records returned, got:", len(arecs))
	}
	// Ensure records were parsed correctly
	expected := [][]string{
		{"service1.example.com", "192.168.178.33"},
		{"service2.example.com", "192.168.178.34"},
		{"service3.example.com", "192.168.178.34"},
	}
	for idx, rec := range arecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
	}

	// Test retrieve AAAA records unfiltered
	arecs, err = cl.listRecords(context.Background(), endpoint.RecordTypeAAAA)
	if err != nil {
		t.Fatal(err)
	}
	if len(arecs) != 3 {
		t.Fatal("Expected 3 AAAA records returned, got:", len(arecs))
	}
	// Ensure records were parsed correctly
	expected = [][]string{
		{"service4.example.com", "fc00::1:192:168:1:1"},
		{"service5.example.com", "fc00::1:192:168:1:2"},
		{"service6.example.com", "fc00::1:192:168:1:3"},
	}
	for idx, rec := range arecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
	}

	// Test retrieve CNAME records unfiltered
	cnamerecs, err := cl.listRecords(context.Background(), endpoint.RecordTypeCNAME)
	if err != nil {
		t.Fatal(err)
	}
	if len(cnamerecs) != 3 {
		t.Fatal("Expected 3 CAME records returned, got:", len(cnamerecs))
	}
	// Ensure records were parsed correctly
	expected = [][]string{
		{"source1.example.com", "target1.domain.com", "1000"},
		{"source2.example.com", "target2.domain.com", "50"},
		{"source3.example.com", "target3.domain.com"},
	}
	for idx, rec := range cnamerecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
		if len(expected[idx]) == 3 {
			expectedTTL, _ := strconv.ParseInt(expected[idx][2], 10, 64)
			if int64(rec.RecordTTL) != expectedTTL {
				t.Error("Got invalid TTL:", rec.RecordTTL, "expected:", expected[idx][2])
			}
		}
	}

	// Note: filtered tests are not needed since A/AAAA records are tested filtered already
	// and cnameRecords have their own element
}

func TestCreateRecordV6(t *testing.T) {
	var ep *endpoint.Endpoint
	srvr := newTestServerV6(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" && (r.URL.Path == "/api/config/dns/hosts/192.168.1.1 test.example.com" ||
			r.URL.Path == "/api/config/dns/hosts/fc00::1:192:168:1:1 test.example.com" ||
			r.URL.Path == "/api/config/dns/cnameRecords/source1.example.com,target1.domain.com" ||
			r.URL.Path == "/api/config/dns/cnameRecords/source2.example.com,target2.domain.com,500") {

			// Return A records
			w.WriteHeader(http.StatusCreated)
		} else {
			http.NotFound(w, r)
		}
	})
	defer srvr.Close()

	// Create a client
	cfg := PiholeConfig{
		Server:     srvr.URL,
		APIVersion: "6",
	}
	cl, err := newPiholeClientV6(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Test create A record
	ep = &endpoint.Endpoint{
		DNSName:    "test.example.com",
		Targets:    []string{"192.168.1.1"},
		RecordType: endpoint.RecordTypeA,
	}
	if err := cl.createRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test create AAAA record
	ep = &endpoint.Endpoint{
		DNSName:    "test.example.com",
		Targets:    []string{"fc00::1:192:168:1:1"},
		RecordType: endpoint.RecordTypeAAAA,
	}
	if err := cl.createRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test create CNAME record
	ep = &endpoint.Endpoint{
		DNSName:    "source1.example.com",
		Targets:    []string{"target1.domain.com"},
		RecordType: endpoint.RecordTypeCNAME,
	}
	if err := cl.createRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test create CNAME record with TTL
	ep = &endpoint.Endpoint{
		DNSName:    "source2.example.com",
		Targets:    []string{"target2.domain.com"},
		RecordTTL:  endpoint.TTL(500),
		RecordType: endpoint.RecordTypeCNAME,
	}
	if err := cl.createRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test create a wildcard record and ensure it fails
	ep = &endpoint.Endpoint{
		DNSName:    "*.example.com",
		Targets:    []string{"192.168.1.1"},
		RecordType: endpoint.RecordTypeA,
	}
	if err := cl.createRecord(context.Background(), ep); err == nil {
		t.Fatal(err)
	}
}

func TestDeleteRecordV6(t *testing.T) {
	var ep *endpoint.Endpoint
	srvr := newTestServerV6(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" && (r.URL.Path == "/api/config/dns/hosts/192.168.1.1 test.example.com" ||
			r.URL.Path == "/api/config/dns/hosts/fc00::1:192:168:1:1 test.example.com" ||
			r.URL.Path == "/api/config/dns/cnameRecords/source1.example.com,target1.domain.com" ||
			r.URL.Path == "/api/config/dns/cnameRecords/source2.example.com,target2.domain.com,500") {

			// Return A records
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.NotFound(w, r)
		}
	})
	defer srvr.Close()

	// Create a client
	cfg := PiholeConfig{
		Server:     srvr.URL,
		APIVersion: "6",
	}
	cl, err := newPiholeClientV6(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Test delete A record
	ep = &endpoint.Endpoint{
		DNSName:    "test.example.com",
		Targets:    []string{"192.168.1.1"},
		RecordType: endpoint.RecordTypeA,
	}
	if err := cl.deleteRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test delete AAAA record
	ep = &endpoint.Endpoint{
		DNSName:    "test.example.com",
		Targets:    []string{"fc00::1:192:168:1:1"},
		RecordType: endpoint.RecordTypeAAAA,
	}
	if err := cl.deleteRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test delete CNAME record
	ep = &endpoint.Endpoint{
		DNSName:    "source1.example.com",
		Targets:    []string{"target1.domain.com"},
		RecordType: endpoint.RecordTypeCNAME,
	}
	if err := cl.deleteRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}

	// Test delete CNAME record with TTL
	ep = &endpoint.Endpoint{
		DNSName:    "source2.example.com",
		Targets:    []string{"target2.domain.com"},
		RecordTTL:  endpoint.TTL(500),
		RecordType: endpoint.RecordTypeCNAME,
	}
	if err := cl.deleteRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}
}
