package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

const (
	clientVersion = "2.10.0"

	defaultEndpoint               = "https://api.nsone.net/v1/"
	defaultShouldFollowPagination = true
	defaultUserAgent              = "go-ns1/" + clientVersion

	headerAuth          = "X-NSONE-Key"
	headerRateLimit     = "X-Ratelimit-Limit"
	headerRateRemaining = "X-Ratelimit-Remaining"
	headerRatePeriod    = "X-Ratelimit-Period"

	defaultRateLimitWaitTime = time.Millisecond * 100
)

// Doer is a single method interface that allows a user to extend/augment an http.Client instance.
// Note: http.Client satisfies the Doer interface.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Client manages communication with the NS1 Rest API.
type Client struct {
	// httpClient handles all rest api communication,
	// and expects an *http.Client.
	httpClient Doer

	// NS1 rest endpoint, overrides default if given.
	Endpoint *url.URL

	// NS1 api key (value for http request header 'X-NSONE-Key').
	APIKey string

	// NS1 go rest user agent (value for http request header 'User-Agent').
	UserAgent string

	// Func to call after response is returned in Do
	RateLimitFunc func(RateLimit)

	// Whether the client should handle paginated responses automatically.
	FollowPagination bool

	// Enables permissions compatibility with the DDI API.
	DDI bool

	// From the excellent github-go client.
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for communicating with different components of the NS1 API.
	APIKeys           *APIKeysService
	DataFeeds         *DataFeedsService
	DataSources       *DataSourcesService
	Jobs              *JobsService
	MonitorRegions    *MonitorRegionsService
	PulsarJobs        *PulsarJobsService
	Notifications     *NotificationsService
	Records           *RecordsService
	Applications      *ApplicationsService
	RecordSearch      *RecordSearchService
	ZoneSearch        *ZoneSearchService
	Settings          *SettingsService
	Stats             *StatsService
	Teams             *TeamsService
	Users             *UsersService
	Warnings          *WarningsService
	Zones             *ZonesService
	Versions          *VersionsService
	DNSSEC            *DNSSECService
	IPAM              *IPAMService
	ScopeGroup        *ScopeGroupService
	Scope             *ScopeService
	Reservation       *ReservationService
	OptionDef         *OptionDefService
	TSIG              *TsigService
	View              *DNSViewService
	Network           *NetworkService
	GlobalIPWhitelist *GlobalIPWhitelistService
	Datasets          *DatasetsService
	Activity          *ActivityService
}

// NewClient constructs and returns a reference to an instantiated Client.
func NewClient(httpClient Doer, options ...func(*Client)) *Client {
	endpoint, _ := url.Parse(defaultEndpoint)

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient:       httpClient,
		Endpoint:         endpoint,
		RateLimitFunc:    defaultRateLimitFunc,
		UserAgent:        defaultUserAgent,
		FollowPagination: defaultShouldFollowPagination,
	}

	c.common.client = c
	c.APIKeys = (*APIKeysService)(&c.common)
	c.DataFeeds = (*DataFeedsService)(&c.common)
	c.DataSources = (*DataSourcesService)(&c.common)
	c.Jobs = (*JobsService)(&c.common)
	c.MonitorRegions = (*MonitorRegionsService)(&c.common)
	c.PulsarJobs = (*PulsarJobsService)(&c.common)
	c.Notifications = (*NotificationsService)(&c.common)
	c.Records = (*RecordsService)(&c.common)
	c.Applications = (*ApplicationsService)(&c.common)
	c.RecordSearch = (*RecordSearchService)(&c.common)
	c.ZoneSearch = (*ZoneSearchService)(&c.common)
	c.Settings = (*SettingsService)(&c.common)
	c.Stats = (*StatsService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Warnings = (*WarningsService)(&c.common)
	c.Zones = (*ZonesService)(&c.common)
	c.Versions = (*VersionsService)(&c.common)
	c.DNSSEC = (*DNSSECService)(&c.common)
	c.IPAM = (*IPAMService)(&c.common)
	c.ScopeGroup = (*ScopeGroupService)(&c.common)
	c.Scope = (*ScopeService)(&c.common)
	c.Reservation = (*ReservationService)(&c.common)
	c.OptionDef = (*OptionDefService)(&c.common)
	c.TSIG = (*TsigService)(&c.common)
	c.View = (*DNSViewService)(&c.common)
	c.Network = (*NetworkService)(&c.common)
	c.GlobalIPWhitelist = (*GlobalIPWhitelistService)(&c.common)
	c.Datasets = (*DatasetsService)(&c.common)
	c.Activity = (*ActivityService)(&c.common)

	for _, option := range options {
		option(c)
	}
	return c
}

type service struct {
	client *Client
}

// SetHTTPClient sets a Client instances' httpClient.
func SetHTTPClient(httpClient Doer) func(*Client) {
	return func(c *Client) { c.httpClient = httpClient }
}

// SetAPIKey sets a Client instances' APIKey.
func SetAPIKey(key string) func(*Client) {
	return func(c *Client) { c.APIKey = key }
}

// SetEndpoint sets a Client instances' Endpoint.
func SetEndpoint(endpoint string) func(*Client) {
	return func(c *Client) { c.Endpoint, _ = url.Parse(endpoint) }
}

// SetUserAgent sets a Client instances' user agent.
func SetUserAgent(ua string) func(*Client) {
	return func(c *Client) { c.UserAgent = ua }
}

// SetRateLimitFunc sets a Client instances' RateLimitFunc.
func SetRateLimitFunc(ratefunc func(rl RateLimit)) func(*Client) {
	return func(c *Client) { c.RateLimitFunc = ratefunc }
}

// SetFollowPagination sets a Client instances' FollowPagination attribute.
func SetFollowPagination(shouldFollow bool) func(*Client) {
	return func(c *Client) { c.FollowPagination = shouldFollow }
}

// SetDDIAPI configures the client to use permissions compatible with the DDI API.
func SetDDIAPI() func(*Client) {
	return func(c *Client) { c.DDI = true }
}

// Param is a container struct which holds a `Key` and `Value` field corresponding to the values of a URL parameter. 
type Param struct {
	Key, Value string
}

// Do satisfies the Doer interface. resp will be nil if a non-HTTP error
// occurs, otherwise it is available for inspection when the error reflects a
// non-2XX response. It accepts a variadic number of optional URL parameters to
// supply to the request. URL parameters are of type `rest.Param`.
func (c Client) Do(req *http.Request, v interface{}, params ...Param) (*http.Response, error) {
	q := req.URL.Query()
	for _, p := range params {
		q.Set(p.Key, p.Value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rl := parseRate(resp)
	c.RateLimitFunc(rl)

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		// For non-JSON responses, the desired destination might be a bytes buffer
		if buf, ok := v.(*bytes.Buffer); ok {
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return nil, err
			}
			return resp, err
		}

		// Try to unmarshal body into given type using streaming decoder.
		if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return nil, err
		}
	}

	return resp, err
}

// NextFunc knows how to get and parse additional info from uri into v.
type NextFunc func(v *interface{}, uri string) (*http.Response, error)

// DoWithPagination Does, and follows Link headers for pagination. The returned
// Response is from the last URI visited - either the last page, or one that
// responded with a non-2XX status. If a non-HTTP error occurs, resp will be
// nil. It accepts a variadic number of optional URL parameters to supply to
// the underlying `.Do()` method request(s). URL parameters are of type
// `rest.Param`.
func (c Client) DoWithPagination(req *http.Request, v interface{}, f NextFunc, params ...Param) (*http.Response, error) {
	resp, err := c.Do(req, v, params...)
	if err != nil {
		return resp, err
	}

	// See PLAT-188
	forceHTTPS := c.Endpoint.Scheme == "https"

	nextURI := ParseLink(resp.Header.Get("Link"), forceHTTPS).Next()
	for nextURI != "" {
		resp, err = f(&v, nextURI)
		if err != nil {
			return resp, err
		}
		nextURI = ParseLink(resp.Header.Get("Link"), forceHTTPS).Next()
	}
	return resp, nil
}

// NewRequest constructs and returns a http.Request.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	uri := c.Endpoint.ResolveReference(rel)

	// Encode body as json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAuth, c.APIKey)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Response wraps stdlib http response.
type Response struct {
	*http.Response
}

// Error contains all http responses outside the 2xx range.
type Error struct {
	Resp    *http.Response
	Message string
}

// Satisfy std lib error interface.
func (re *Error) Error() string {
	return fmt.Sprintf("%v %v: %d %v", re.Resp.Request.Method, re.Resp.Request.URL, re.Resp.StatusCode, re.Message)
}

// CheckResponse handles parsing of rest api errors. Returns nil if no error.
func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	restErr := &Error{Resp: resp}

	msgBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(msgBody) == 0 {
		return restErr
	}

	err = json.Unmarshal(msgBody, restErr)
	if err != nil {
		restErr.Message = string(msgBody)
		return restErr
	}

	return restErr
}

// Helper function for parsing API responses for a specific error.
// Ideally this would take place in CheckResponse above rather than
// in each caller.
var resourceMissingMatch = regexp.MustCompile(` not found`).MatchString

// RateLimitFunc is rate limiting strategy for the Client instance.
type RateLimitFunc func(RateLimit)

// RateLimit stores X-Ratelimit-* headers
type RateLimit struct {
	Limit     int
	Remaining int
	Period    int
}

var defaultRateLimitFunc = func(rl RateLimit) {}

// PercentageLeft returns the ratio of Remaining to Limit as a percentage
func (rl RateLimit) PercentageLeft() int {
	return rl.Remaining * 100 / rl.Limit
}

// WaitTime returns the time.Duration ratio of Period to Limit
func (rl RateLimit) WaitTime() time.Duration {
	if rl.Limit == 0 || rl.Period == 0 {
		// rate-limit headers missing or corrupt, punt
		return defaultRateLimitWaitTime
	}
	return (time.Second * time.Duration(rl.Period)) / time.Duration(rl.Limit)
}

// WaitTimeRemaining returns the time.Duration ratio of Period to Remaining
func (rl RateLimit) WaitTimeRemaining() time.Duration {
	if rl.Remaining < 2 {
		return time.Second * time.Duration(rl.Period)
	}
	return (time.Second * time.Duration(rl.Period)) / time.Duration(rl.Remaining)
}

// RateLimitStrategySleep sets RateLimitFunc to sleep by WaitTimeRemaining
func (c *Client) RateLimitStrategySleep() {
	c.RateLimitFunc = func(rl RateLimit) {
		remaining := rl.WaitTimeRemaining()
		time.Sleep(remaining)
	}
}

// RateLimitStrategyConcurrent sleeps for WaitTime * parallelism when
// remaining is less than or equal to parallelism.
func (c *Client) RateLimitStrategyConcurrent(parallelism int) {
	c.RateLimitFunc = func(rl RateLimit) {
		if rl.Remaining <= parallelism {
			wait := rl.WaitTime() * time.Duration(parallelism)
			time.Sleep(wait)
		}
	}
}

// parseRate parses rate related headers from http response.
func parseRate(resp *http.Response) RateLimit {
	var rl RateLimit

	if limit := resp.Header.Get(headerRateLimit); limit != "" {
		rl.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := resp.Header.Get(headerRateRemaining); remaining != "" {
		rl.Remaining, _ = strconv.Atoi(remaining)
	}
	if period := resp.Header.Get(headerRatePeriod); period != "" {
		rl.Period, _ = strconv.Atoi(period)
	}

	return rl
}

// SetTimeParam sets a url timestamp query param given the parameters name.
func SetTimeParam(key string, t time.Time) func(*url.Values) {
	return func(v *url.Values) { v.Set(key, strconv.Itoa(int(t.Unix()))) }
}

// SetBoolParam sets a url boolean query param given the parameters name.
func SetBoolParam(key string, b bool) func(*url.Values) {
	return func(v *url.Values) { v.Set(key, strconv.FormatBool(b)) }
}

// SetStringParam sets a url string query param given the parameters name.
func SetStringParam(key, val string) func(*url.Values) {
	return func(v *url.Values) { v.Set(key, val) }
}

// SetIntParam sets a url integer query param given the parameters name.
func SetIntParam(key string, val int) func(*url.Values) {
	return func(v *url.Values) { v.Set(key, strconv.Itoa(val)) }
}

func (c *Client) getURI(v interface{}, uri string) (*http.Response, error) {
	req, err := c.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	// For non-2XX responses, Do returns the response as well as an error, for
	// other errs, resp will be nil. Caller's responsibility to sort that out.
	return c.Do(req, v)
}
