package cloudflare

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested
	client *API

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Cloudflare client configured to use test server
	client, _ = New("cloudflare@example.org", "deadbeef")
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestClient_Auth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ips", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "cloudflare@example.com", r.Header.Get("X-Auth-Email"))
		assert.Equal(t, "deadbeef", r.Header.Get("X-Auth-Token"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
  "success": true,
  "errors": [],
  "messages": [],
  "response": {
    "ipv4_cidrs": ["199.27.128.0/21"],
    "ipv6_cidrs": ["199.27.128.0/21"]
  }
}`)
	})

	_, err := IPs()

	assert.NoError(t, err)
}
