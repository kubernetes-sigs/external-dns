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

package adguardhome

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

func newTestServer(t *testing.T, hdlr http.HandlerFunc) *httptest.Server {
	t.Helper()
	svr := httptest.NewServer(hdlr)
	return svr
}

func TestNewClient(t *testing.T) {
	// Test correct error on no server provided
	_, err := newClient(AdGuardHomeConfig{})
	if err == nil {
		t.Error("Expected error from config with no server")
	} else if err != ErrNoServer {
		t.Error("Expected ErrNoServer, got", err)
	}

	// Create a test server for auth tests
	noAuthSrvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		// This is a subset of the status response
		w.Write([]byte("{}"))
	})
	defer noAuthSrvr.Close()

	// Test new client with no password. Should create the
	// client cleanly.
	cl, err := newClient(AdGuardHomeConfig{Server: noAuthSrvr.URL})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cl.(*client); !ok {
		t.Error("Did not create a new adguardhome client")
	}

	// Create a test server for auth tests
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/status" {
			t.Errorf("Invalid request URL: %s, Expected: %s", r.URL.Path, "/control/status")
		}
		authorization := r.Header.Get("Authorization")
		b64 := strings.TrimPrefix(authorization, "Basic ")
		userPass, _ := base64.StdEncoding.DecodeString(b64)
		user, pass, _ := strings.Cut(string(userPass), ":")
		if user != "correct" || pass != "correct" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		}
		// This is a subset of what happens on successful login
		w.Write([]byte("{}"))
	})
	defer srvr.Close()

	// Test invalid password
	_, err = newClient(
		AdGuardHomeConfig{Server: srvr.URL, Username: "correct", Password: "wrong"},
	)
	if err == nil {
		t.Error("Expected error for creating client with invalid password")
	}

	// Test correct password
	cl, err = newClient(
		AdGuardHomeConfig{Server: srvr.URL, Username: "correct", Password: "correct"},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Test invalid username
	_, err = newClient(
		AdGuardHomeConfig{Server: srvr.URL, Username: "wrong", Password: "correct"},
	)
	if err == nil {
		t.Error("Expected error for creating client with invalid username")
	}

	// Test correct username
	cl, err = newClient(
		AdGuardHomeConfig{Server: srvr.URL, Username: "correct", Password: "correct"},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Base64 for `correct:correct`
	if cl.(*client).token != "Y29ycmVjdDpjb3JyZWN0" {
		t.Error("Parsed invalid token from login response:", cl.(*client).token)
	}
}

func TestListRecords(t *testing.T) {
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/control/status" {
			w.Write([]byte("{}"))
			return
		}

		if r.Method != http.MethodGet {
			t.Errorf("Invalid request: %s, Expected: %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != "/control/rewrite/list" {
			t.Errorf("Invalid request URL: %s, Expected: %s", r.URL.Path, "/control/rewrite/list")
		}
		w.Write([]byte(`
		[
			{"domain": "test1.example.com", "answer": "192.168.1.1"},
			{"domain": "test2.example.com", "answer": "192.168.1.2"},
			{"domain": "test3.match.com",   "answer": "192.168.1.3"},
			{"domain": "test4.example.com", "answer": "cname.example.com"},
			{"domain": "test5.example.com", "answer": "cname.example.com"},
			{"domain": "test6.match.com",   "answer": "cname.example.com"}
		]
		`))
	})
	defer srvr.Close()

	// Create a client
	cfg := AdGuardHomeConfig{
		Server: srvr.URL,
	}
	cl, err := newClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Test retrieve records unfiltered
	arecs, err := cl.listRecords(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(arecs) != 6 {
		t.Fatal("Expected 6 records returned, got:", len(arecs))
	}
	// Ensure records were parsed correctly
	expected := [][]string{
		{"test1.example.com", "192.168.1.1"},
		{"test2.example.com", "192.168.1.2"},
		{"test3.match.com", "192.168.1.3"},
		{"test4.example.com", "cname.example.com"},
		{"test5.example.com", "cname.example.com"},
		{"test6.match.com", "cname.example.com"},
	}
	for idx, rec := range arecs {
		if rec.DNSName != expected[idx][0] {
			t.Errorf("Got invalid DNS Name: %s, Expected %s", rec.DNSName, expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Errorf("Got invalid target: %s, Expected: %s", rec.Targets[0], expected[idx][1])
		}
	}

	// Same tests but with a domain filter

	cfg.DomainFilter = endpoint.NewDomainFilter([]string{"match.com"})
	cl, err = newClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Test retrieve records
	arecs, err = cl.listRecords(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(arecs) != 2 {
		t.Fatal("Expected 2 records returned, got:", len(arecs))
	}
	// Ensure records were parsed correctly
	expected = [][]string{
		{"test3.match.com", "192.168.1.3"},
		{"test6.match.com", "cname.example.com"},
	}
	for idx, rec := range arecs {
		if rec.DNSName != expected[idx][0] {
			t.Errorf("Got invalid DNS Name: %s, Expected: %s", rec.DNSName, expected[idx][0])
		}
		if rec.Targets[0] != expected[idx][1] {
			t.Errorf("Got invalid target: %s, Expected: %s", rec.Targets[0], expected[idx][1])
		}
	}
}

func TestCreateRecord(t *testing.T) {
	var ep *endpoint.Endpoint
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/control/status" {
			w.Write([]byte("{}"))
			return
		}

		if r.URL.Path != "/control/rewrite/add" {
			t.Errorf("Invalid request URL: %s, Expected: %s", r.URL.Path, "/control/rewrite/add")
		}
		if r.Method != http.MethodPost {
			t.Errorf("Invalid request: %s, Expected: %s", r.Method, http.MethodPost)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			w.Write([]byte("only content-type application/json is allowed"))
		}
		var reqBody rewriteModel
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			t.Fatal(err)
		}
		if reqBody.Domain != ep.DNSName {
			t.Errorf("Invalid domain in body: %s, Expected: %s", reqBody.Domain, ep.DNSName)
		}
		if reqBody.Answer != ep.Targets[0] {
			t.Errorf("Invalid target in body: %s, Expected: %s", reqBody.Answer, ep.Targets[0])
		}
		out, err := json.Marshal(rewriteModel{
			Domain: ep.DNSName,
			Answer: ep.Targets[0],
		})
		if err != nil {
			t.Fatal(err)
		}
		w.Write(out)
	})
	defer srvr.Close()

	// Create a client
	cfg := AdGuardHomeConfig{
		Server: srvr.URL,
	}
	cl, err := newClient(cfg)
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

	// Test create CNAME record
	ep = &endpoint.Endpoint{
		DNSName:    "test.example.com",
		Targets:    []string{"test.cname.com"},
		RecordType: endpoint.RecordTypeCNAME,
	}
	if err := cl.createRecord(context.Background(), ep); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteRecord(t *testing.T) {
	var ep *endpoint.Endpoint
	srvr := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/control/status" {
			w.Write([]byte("{}"))
			return
		}

		if r.URL.Path != "/control/rewrite/delete" {
			t.Errorf("Invalid request URL: %s, Expected: %s", r.URL.Path, "/control/rewrite/delete")
		}
		if r.Method != http.MethodPost {
			t.Errorf("Invalid request: %s, Expected: %s", r.Method, http.MethodPost)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			w.Write([]byte("only content-type application/json is allowed"))
		}
		var reqBody rewriteModel
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			t.Fatal(err)
		}
		if reqBody.Domain != ep.DNSName {
			t.Errorf("Invalid domain in body: %s, Expected: %s", reqBody.Domain, ep.DNSName)
		}
		if reqBody.Answer != ep.Targets[0] {
			t.Errorf("Invalid target in body: %s, Expected: %s", reqBody.Answer, ep.Targets[0])
		}
		out, err := json.Marshal(rewriteModel{
			Domain: ep.DNSName,
			Answer: ep.Targets[0],
		})
		if err != nil {
			t.Fatal(err)
		}
		w.Write(out)
	})
	defer srvr.Close()

	// Create a client
	cfg := AdGuardHomeConfig{
		Server: srvr.URL,
	}
	cl, err := newClient(cfg)
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
