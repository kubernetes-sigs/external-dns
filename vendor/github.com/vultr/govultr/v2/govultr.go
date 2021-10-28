package govultr

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	version     = "2.9.0"
	defaultBase = "https://api.vultr.com"
	userAgent   = "govultr/" + version
	rateLimit   = 500 * time.Millisecond
	retryLimit  = 3
)

// APIKey contains a users API Key for interacting with the API
type APIKey struct {
	// API Key
	key string
}

// RequestBody is used to create JSON bodies for one off calls
type RequestBody map[string]interface{}

// Client manages interaction with the Vultr V1 API
type Client struct {
	// Http Client used to interact with the Vultr V1 API
	client *retryablehttp.Client

	// BASE URL for APIs
	BaseURL *url.URL

	// User Agent for the client
	UserAgent string

	// Services used to interact with the API
	Account         AccountService
	Application     ApplicationService
	Backup          BackupService
	BareMetalServer BareMetalServerService
	BlockStorage    BlockStorageService
	Domain          DomainService
	DomainRecord    DomainRecordService
	FirewallGroup   FirewallGroupService
	FirewallRule    FireWallRuleService
	Instance        InstanceService
	ISO             ISOService
	Kubernetes      KubernetesService
	LoadBalancer    LoadBalancerService
	Network         NetworkService
	ObjectStorage   ObjectStorageService
	OS              OSService
	Plan            PlanService
	Region          RegionService
	ReservedIP      ReservedIPService
	Snapshot        SnapshotService
	SSHKey          SSHKeyService
	StartupScript   StartupScriptService
	User            UserService

	// Optional function called after every successful request made to the Vultr API
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// NewClient returns a Vultr API Client
func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBase)

	client := &Client{
		client:    retryablehttp.NewClient(),
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	client.client.HTTPClient = httpClient
	client.client.Logger = nil
	client.client.ErrorHandler = client.vultrErrorHandler
	client.SetRetryLimit(retryLimit)
	client.SetRateLimit(rateLimit)

	client.Account = &AccountServiceHandler{client}
	client.Application = &ApplicationServiceHandler{client}
	client.Backup = &BackupServiceHandler{client}
	client.BareMetalServer = &BareMetalServerServiceHandler{client}
	client.BlockStorage = &BlockStorageServiceHandler{client}
	client.Domain = &DomainServiceHandler{client}
	client.DomainRecord = &DomainRecordsServiceHandler{client}
	client.FirewallGroup = &FireWallGroupServiceHandler{client}
	client.FirewallRule = &FireWallRuleServiceHandler{client}
	client.Instance = &InstanceServiceHandler{client}
	client.ISO = &ISOServiceHandler{client}
	client.Kubernetes = &KubernetesHandler{client}
	client.LoadBalancer = &LoadBalancerHandler{client}
	client.Network = &NetworkServiceHandler{client}
	client.ObjectStorage = &ObjectStorageServiceHandler{client}
	client.OS = &OSServiceHandler{client}
	client.Plan = &PlanServiceHandler{client}
	client.Region = &RegionServiceHandler{client}
	client.ReservedIP = &ReservedIPServiceHandler{client}
	client.Snapshot = &SnapshotServiceHandler{client}
	client.SSHKey = &SSHKeyServiceHandler{client}
	client.StartupScript = &StartupScriptServiceHandler{client}
	client.User = &UserServiceHandler{client}

	return client
}

// NewRequest creates an API Request
func (c *Client) NewRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	resolvedURL, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, resolvedURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// DoWithContext sends an API Request and returns back the response. The API response is checked  to see if it was
// a successful call. A successful call is then checked to see if we need to unmarshal since some resources
// have their own implements of unmarshal.
func (c *Client) DoWithContext(ctx context.Context, r *http.Request, data interface{}) error {
	rreq, err := retryablehttp.FromRequest(r)
	if err != nil {
		return err
	}

	rreq = rreq.WithContext(ctx)

	res, err := c.client.Do(rreq)

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(r, res)
	}

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode <= http.StatusNoContent {
		if data != nil {
			if err := json.Unmarshal(body, data); err != nil {
				return err
			}
		}
		return nil
	}

	return errors.New(string(body))
}

// SetBaseURL Overrides the default BaseUrl
func (c *Client) SetBaseURL(baseURL string) error {
	updatedURL, err := url.Parse(baseURL)

	if err != nil {
		return err
	}

	c.BaseURL = updatedURL
	return nil
}

// SetRateLimit Overrides the default rateLimit. For performance, exponential
// backoff is used with the minimum wait being 2/3rds the time provided.
func (c *Client) SetRateLimit(time time.Duration) {
	c.client.RetryWaitMin = time / 3 * 2
	c.client.RetryWaitMax = time
}

// SetUserAgent Overrides the default UserAgent
func (c *Client) SetUserAgent(ua string) {
	c.UserAgent = ua
}

// OnRequestCompleted sets the API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// SetRetryLimit overrides the default RetryLimit
func (c *Client) SetRetryLimit(n int) {
	c.client.RetryMax = n
}

func (c *Client) vultrErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	if resp == nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (resp == nil)", c.client.RetryMax+1)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gave up after %d attempts, last error unavailable (error reading response body: %v)", c.client.RetryMax+1, err)
	}
	return nil, fmt.Errorf("gave up after %d attempts, last error: %#v", c.client.RetryMax+1, strings.TrimSpace(string(buf)))
}

// BoolToBoolPtr helper function that returns a pointer from your bool value
func BoolToBoolPtr(value bool) *bool {
	return &value
}

// StringToStringPtr helper function that returns a pointer from your string value
func StringToStringPtr(value string) *string {
	return &value
}

// IntToIntPtr helper function that returns a pointer from your string value
func IntToIntPtr(value int) *int {
	return &value
}
