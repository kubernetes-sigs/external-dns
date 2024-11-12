package godaddy

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

// Tests that
func TestClient_DoWhenQuotaExceeded(t *testing.T) {
	assert := assert.New(t)

	// Mock server to return 429 with a JSON payload
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		_, err := w.Write([]byte(`{"code": "QUOTA_EXCEEDED", "message": "rate limit exceeded"}`))
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}))
	defer mockServer.Close()

	client := Client{
		APIKey:      "",
		APISecret:   "",
		APIEndPoint: mockServer.URL,
		Client:      &http.Client{},
		// Add one token every second
		Ratelimiter: rate.NewLimiter(rate.Every(time.Second), 60),
		Timeout:     DefaultTimeout,
	}

	req, err := client.NewRequest("GET", "/v1/domains/example.net/records", nil, false)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	assert.Nil(err, "A CODE_EXCEEDED response should not return an error")
	assert.Equal(http.StatusTooManyRequests, resp.StatusCode, "Expected a 429 response")

	respContents := GDErrorResponse{}
	err = client.UnmarshalResponse(resp, &respContents)
	if assert.NotNil(err) {
		var apiErr *APIError
		errors.As(err, &apiErr)
		assert.Equal("QUOTA_EXCEEDED", apiErr.Code)
		assert.Equal("rate limit exceeded", apiErr.Message)
	}
}
