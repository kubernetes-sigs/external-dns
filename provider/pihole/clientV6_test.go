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
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

	// Test new client with no password. Should create the
	// client cleanly.
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

	// Create a test server for auth tests
	srvr := newTestServerV6(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/auth" && r.Method == "POST" {
			var requestData map[string]string
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			err = json.Unmarshal(body, &requestData)
			if err != nil {
				http.Error(w, "Error parsing JSON", http.StatusBadRequest)
				return
			}

			pw := requestData["password"]
			if pw != "correct" {
				// Return unsuccessful authentication response
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"session":{"valid":false,"totp":false,"sid":null,"validity":-1,"message":"password incorrect"},"took":0.2}`))
				return
			}

			w.Header().Set("Content-Type", "application/json")
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
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Form.Get("action") != "get" {
			t.Error("Expected 'get' action in form from client")
		}
		if strings.Contains(r.URL.Path, "cname") {
			w.Write([]byte(`
			{
				"data": [
					["test4.example.com", "cname.example.com"],
					["test5.example.com", "cname.example.com"],
					["test6.match.com", "cname.example.com"]
				]
			}
			`))
			return
		}
		// Pihole makes no distinction between A and AAAA records
		w.Write([]byte(`
		{
			"data": [
				["test1.example.com", "192.168.1.1"],
				["test2.example.com", "192.168.1.2"],
				["test3.match.com", "192.168.1.3"],
				["test1.example.com", "fc00::1:192:168:1:1"],
				["test2.example.com", "fc00::1:192:168:1:2"],
				["test3.match.com", "fc00::1:192:168:1:3"]
			]
		}
		`))
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
		{"test1.example.com", "192.168.1.1"},
		{"test2.example.com", "192.168.1.2"},
		{"test3.match.com", "192.168.1.3"},
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
		{"test1.example.com", "fc00::1:192:168:1:1"},
		{"test2.example.com", "fc00::1:192:168:1:2"},
		{"test3.match.com", "fc00::1:192:168:1:3"},
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
		{"test4.example.com", "cname.example.com"},
		{"test5.example.com", "cname.example.com"},
		{"test6.match.com", "cname.example.com"},
	}
	for idx, rec := range cnamerecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
	}

	// Same tests but with a domain filter

	cfg.DomainFilter = endpoint.NewDomainFilter([]string{"match.com"})
	cl, err = newPiholeClientV6(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Test retrieve A records filtered
	arecs, err = cl.listRecords(context.Background(), endpoint.RecordTypeA)
	if err != nil {
		t.Fatal(err)
	}
	if len(arecs) != 1 {
		t.Fatal("Expected 1 A record returned, got:", len(arecs))
	}
	// Ensure records were parsed correctly
	expected = [][]string{
		{"test3.match.com", "192.168.1.3"},
	}
	for idx, rec := range arecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
	}

	// Test retrieve AAAA records filtered
	arecs, err = cl.listRecords(context.Background(), endpoint.RecordTypeAAAA)
	if err != nil {
		t.Fatal(err)
	}
	if len(arecs) != 1 {
		t.Fatal("Expected 1 AAAA record returned, got:", len(arecs))
	}
	// Ensure records were parsed correctly
	expected = [][]string{
		{"test3.match.com", "fc00::1:192:168:1:3"},
	}
	for idx, rec := range arecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
	}

	// Test retrieve CNAME records filtered
	cnamerecs, err = cl.listRecords(context.Background(), endpoint.RecordTypeCNAME)
	if err != nil {
		t.Fatal(err)
	}
	if len(cnamerecs) != 1 {
		t.Fatal("Expected 1 CNAME record returned, got:", len(cnamerecs))
	}
	// Ensure records were parsed correctly
	expected = [][]string{
		{"test6.match.com", "cname.example.com"},
	}
	for idx, rec := range cnamerecs {
		if rec.DNSName != expected[idx][0] {
			t.Error("Got invalid DNS Name:", rec.DNSName, "expected:", expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Error("Got invalid target:", rec.Targets[0], "expected:", expected[idx][1])
		}
	}
}

func TestCreateRecordV6(t *testing.T) {
	var ep *endpoint.Endpoint
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Form.Get("action") != "add" {
			t.Error("Expected 'add' action in form from client")
		}
		if r.Form.Get("domain") != ep.DNSName {
			t.Error("Invalid domain in form:", r.Form.Get("domain"), "Expected:", ep.DNSName)
		}
		switch ep.RecordType {
		case endpoint.RecordTypeA:
			if r.Form.Get("ip") != ep.Targets[0] {
				t.Error("Invalid ip in form:", r.Form.Get("ip"), "Expected:", ep.Targets[0])
			}
		// Pihole makes no distinction between A and AAAA records
		case endpoint.RecordTypeAAAA:
			if r.Form.Get("ip") != ep.Targets[0] {
				t.Error("Invalid ip in form:", r.Form.Get("ip"), "Expected:", ep.Targets[0])
			}
		case endpoint.RecordTypeCNAME:
			if r.Form.Get("target") != ep.Targets[0] {
				t.Error("Invalid target in form:", r.Form.Get("target"), "Expected:", ep.Targets[0])
			}
		}
		out, err := json.Marshal(actionResponse{
			Success: true,
			Message: "",
		})
		if err != nil {
			t.Fatal(err)
		}
		w.Write(out)
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
		DNSName:    "test.example.com",
		Targets:    []string{"test.cname.com"},
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
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Form.Get("action") != "delete" {
			t.Error("Expected 'delete' action in form from client")
		}
		if r.Form.Get("domain") != ep.DNSName {
			t.Error("Invalid domain in form:", r.Form.Get("domain"), "Expected:", ep.DNSName)
		}
		switch ep.RecordType {
		case endpoint.RecordTypeA:
			if r.Form.Get("ip") != ep.Targets[0] {
				t.Error("Invalid ip in form:", r.Form.Get("ip"), "Expected:", ep.Targets[0])
			}
		// Pihole makes no distinction between A and AAAA records
		case endpoint.RecordTypeAAAA:
			if r.Form.Get("ip") != ep.Targets[0] {
				t.Error("Invalid ip in form:", r.Form.Get("ip"), "Expected:", ep.Targets[0])
			}
		case endpoint.RecordTypeCNAME:
			if r.Form.Get("target") != ep.Targets[0] {
				t.Error("Invalid target in form:", r.Form.Get("target"), "Expected:", ep.Targets[0])
			}
		}
		out, err := json.Marshal(actionResponse{
			Success: true,
			Message: "",
		})
		if err != nil {
			t.Fatal(err)
		}
		w.Write(out)
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
		DNSName:    "test.example.com",
		Targets:    []string{"test.cname.com"},
		RecordType: endpoint.RecordTypeCNAME,
	}
	if err := cl.deleteRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}
}
