package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// setup creates a test HTTP server and a client configured to use it
func setup(t *testing.T) (*Client, *httptest.Server, *http.ServeMux) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := NewClient(http.DefaultClient,
		SetEndpoint(server.URL),
		SetAPIKey("test-key"))

	return client, server, mux
}
