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

func setup(opts ...Option) {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Cloudflare client configured to use test server
	client, _ = New("deadbeef", "cloudflare@example.org", opts...)
	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestClient_Headers(t *testing.T) {
	// it should set default headers
	setup()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "cloudflare@example.org", r.Header.Get("X-Auth-Email"))
		assert.Equal(t, "deadbeef", r.Header.Get("X-Auth-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	client.UserDetails()
	teardown()

	// it should override appropriate default headers when custom headers given
	headers := make(http.Header)
	headers.Set("Content-Type", "application/xhtml+xml")
	headers.Add("X-Random", "a random header")
	setup(Headers(headers))
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Equal(t, "cloudflare@example.org", r.Header.Get("X-Auth-Email"))
		assert.Equal(t, "deadbeef", r.Header.Get("X-Auth-Key"))
		assert.Equal(t, "application/xhtml+xml", r.Header.Get("Content-Type"))
		assert.Equal(t, "a random header", r.Header.Get("X-Random"))
	})
	client.UserDetails()
	teardown()

	// it should set X-Auth-User-Service-Key and omit X-Auth-Email and X-Auth-Key when client.authType is AuthUserService
	setup()
	client.SetAuthType(AuthUserService)
	client.APIUserServiceKey = "userservicekey"
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "Expected method 'GET', got %s", r.Method)
		assert.Empty(t, r.Header.Get("X-Auth-Email"))
		assert.Empty(t, r.Header.Get("X-Auth-Key"))
		assert.Equal(t, "userservicekey", r.Header.Get("X-Auth-User-Service-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	})
	client.UserDetails()
	teardown()
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
