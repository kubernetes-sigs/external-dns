package dnsimple

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// BasicAuthHTTPClient returns a client that authenticates via HTTP Basic Auth with given username and password.
func BasicAuthHTTPClient(_ context.Context, username, password string) *http.Client {
	tp := BasicAuthTransport{Username: username, Password: password}
	return tp.Client()
}

// StaticTokenHTTPClient returns a client that authenticates with a static OAuth token.
func StaticTokenHTTPClient(ctx context.Context, token string) *http.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	return oauth2.NewClient(ctx, ts)
}

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and password.
type BasicAuthTransport struct {
	Username string
	Password string

	// Transport is the transport RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.  We just add the
// basic auth and return the RoundTripper for this transport type.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract

	req2.SetBasicAuth(t.Username, t.Password)
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that uses the BasicAuthTransport transport
// to authenticate the request via HTTP Basic Auth.
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
