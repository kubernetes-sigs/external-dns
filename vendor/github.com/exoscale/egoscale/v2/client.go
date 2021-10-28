package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"time"

	"github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
	"github.com/exoscale/egoscale/version"
)

const (
	defaultTimeout      = 60 * time.Second
	defaultPollInterval = oapi.DefaultPollingInterval
)

// UserAgent is the "User-Agent" HTTP request header added to outgoing HTTP requests.
var UserAgent = fmt.Sprintf("egoscale/%s (%s; %s/%s)",
	version.Version,
	runtime.Version(),
	runtime.GOOS,
	runtime.GOARCH)

// defaultTransport is the default HTTP client transport.
type defaultTransport struct {
	next http.RoundTripper
}

// RoundTrip executes a single HTTP transaction, returning a Response for the provided Request.
func (t *defaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", UserAgent)

	resp, err := t.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ClientOpt represents a function setting Exoscale API client option.
type ClientOpt func(*Client) error

// ClientOptWithAPIEndpoint returns a ClientOpt overriding the default Exoscale
// API endpoint.
func ClientOptWithAPIEndpoint(v string) ClientOpt {
	return func(c *Client) error {
		endpointURL, err := url.Parse(v)
		if err != nil {
			return fmt.Errorf("failed to parse URL: %s", err)
		}

		endpointURL = endpointURL.ResolveReference(&url.URL{Path: api.Prefix})
		c.apiEndpoint = endpointURL.String()

		return nil
	}
}

// ClientOptWithTimeout returns a ClientOpt overriding the default client timeout.
func ClientOptWithTimeout(v time.Duration) ClientOpt {
	return func(c *Client) error {
		if v <= 0 {
			return errors.New("timeout value must be greater than 0")
		}

		c.timeout = v

		return nil
	}
}

// ClientOptWithPollInterval returns a ClientOpt overriding the default client async operation polling interval.
func ClientOptWithPollInterval(v time.Duration) ClientOpt {
	return func(c *Client) error {
		if v <= 0 {
			return errors.New("poll interval value must be greater than 0")
		}

		c.pollInterval = v

		return nil
	}
}

// ClientOptWithTrace returns a ClientOpt enabling HTTP request/reponse tracing.
func ClientOptWithTrace() ClientOpt {
	return func(c *Client) error {
		c.trace = true
		return nil
	}
}

// ClientOptCond returns the specified ClientOpt if the fc function bool result
// evaluates to true, otherwise returns a no-op ClientOpt.
func ClientOptCond(fc func() bool, opt ClientOpt) ClientOpt {
	if fc() {
		return opt
	}

	return func(*Client) error { return nil }
}

// ClientOptWithHTTPClient returns a ClientOpt overriding the default http.Client.
// Note: the Exoscale API client will chain additional middleware
// (http.RoundTripper) on the HTTP client internally, which can alter the HTTP
// requests and responses. If you don't want any other middleware than the ones
// currently set to your HTTP client, you should duplicate it and pass a copy
// instead.
func ClientOptWithHTTPClient(v *http.Client) ClientOpt {
	return func(c *Client) error {
		c.httpClient = v

		return nil
	}
}

// Client represents an Exoscale API client.
type Client struct {
	apiKey       string
	apiSecret    string
	apiEndpoint  string
	timeout      time.Duration
	pollInterval time.Duration
	trace        bool
	httpClient   *http.Client

	*oapi.ClientWithResponses
}

// NewClient returns a new Exoscale API client, or an error if one couldn't be initialized.
func NewClient(apiKey, apiSecret string, opts ...ClientOpt) (*Client, error) {
	client := Client{
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		apiEndpoint:  api.EndpointURL,
		httpClient:   &http.Client{Transport: &defaultTransport{http.DefaultTransport}},
		timeout:      defaultTimeout,
		pollInterval: defaultPollInterval,
	}

	if client.apiKey == "" || client.apiSecret == "" {
		return nil, fmt.Errorf("%w: missing or incomplete API credentials", ErrClientConfig)
	}

	for _, opt := range opts {
		if err := opt(&client); err != nil {
			return nil, fmt.Errorf("%w: %s", ErrClientConfig, err)
		}
	}

	apiSecurityProvider, err := api.NewSecurityProvider(client.apiKey, client.apiSecret)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize API security provider: %s", err)
	}

	apiURL, err := url.Parse(client.apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize API client: %s", err)
	}
	apiURL = apiURL.ResolveReference(&url.URL{Path: api.Prefix})

	// Tracing must be performed before API error handling in the middleware chain,
	// otherwise the response won't be dumped in case of an API error.
	if client.trace {
		client.httpClient.Transport = api.NewTraceMiddleware(client.httpClient.Transport)
	}

	client.httpClient.Transport = api.NewAPIErrorHandlerMiddleware(client.httpClient.Transport)

	oapiOpts := []oapi.ClientOption{
		oapi.WithHTTPClient(client.httpClient),
		oapi.WithRequestEditorFn(
			oapi.MultiRequestsEditor(
				apiSecurityProvider.Intercept,
				setEndpointFromContext,
			),
		),
	}

	if client.ClientWithResponses, err = oapi.NewClientWithResponses(apiURL.String(), oapiOpts...); err != nil {
		return nil, fmt.Errorf("unable to initialize API client: %s", err)
	}

	return &client, nil
}

// SetHTTPClient overrides the current HTTP client.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// SetTimeout overrides the current client timeout value.
func (c *Client) SetTimeout(v time.Duration) {
	c.timeout = v
}

// SetTrace enables or disables HTTP request/reponse tracing.
func (c *Client) SetTrace(enabled bool) {
	c.trace = enabled
}

// setEndpointFromContext is an HTTP client request interceptor that overrides the "Host" header
// with information from a request endpoint optionally set in the context instance. If none is
// found, the request is left untouched.
func setEndpointFromContext(ctx context.Context, req *http.Request) error {
	if v, ok := ctx.Value(api.ReqEndpoint{}).(api.ReqEndpoint); ok {
		req.Host = v.Host()
		req.URL.Host = v.Host()
	}

	return nil
}

// fetchFromIDs returns a list of API resources fetched from the specified list of IDs.
// It is meant to be used with API resources implementing the getter interface, e.g.:
//
//     func (i Instance) get(ctx context.Context, client *Client, zone, id string) (interface{}, error) {
//         return client.GetInstance(ctx, zone, id)
//     }
//
//     func (i *InstancePool) Instances(ctx context.Context) ([]*Instance, error) {
//         res, err := i.c.fetchFromIDs(ctx, i.zone, i.InstanceIDs, new(Instance))
//         return res.([]*Instance), err
//     }
//
func (c *Client) fetchFromIDs(ctx context.Context, zone string, ids []string, rt interface{}) (interface{}, error) {
	if rt == nil {
		return nil, errors.New("resource type must not be <nil>")
	}

	resType := reflect.ValueOf(rt).Type()
	if kind := resType.Kind(); kind != reflect.Ptr {
		return nil, fmt.Errorf("expected resource type to be a pointer, got %s", kind)
	}

	// Base type identification is necessary as it is not possible to call
	// the Getter.Get() method on a nil pointer, so we create a new value
	// using the base type and call the Get() method on it. The corollary is
	// that the Get() method must be implemented on the type directly,
	// not as a pointer receiver.
	baseType := resType.Elem()

	if !resType.Implements(reflect.TypeOf(new(getter)).Elem()) {
		return nil, fmt.Errorf("resource type %s does not implement the Getter interface", resType)
	}

	// As a convenience to the caller, even if the list of IDs passed as
	// parameter is empty we always allocate a slice of <rt> and return
	// it to them, this way they can confidently convert the returned
	// interface{} into a []<rt> without having to perform type assertion.
	collector := reflect.MakeSlice(reflect.SliceOf(resType), 0, 0)

	for _, id := range ids {
		res, err := reflect.New(baseType).Elem().Interface().(getter).get(ctx, c, zone, id)
		if err != nil {
			return nil, err
		}
		collector = reflect.Append(collector, reflect.ValueOf(res))
	}

	return collector.Interface(), nil
}
