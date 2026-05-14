package linodego

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/go-resty/resty/v2"
)

const (
	// MonitorAPIHost is the default monitor-api host
	MonitorAPIHost = "monitor-api.linode.com"
	// MonitorAPIHostVar is the env var to check for the alternate Monitor API URL
	MonitorAPIHostVar = "MONITOR_API_URL"
	// MonitorAPIVersion is the default API version to use
	MonitorAPIVersion = "v2beta"
	// MonitorAPIVersionVar is the env var to check for the alternate Monitor API version
	MonitorAPIVersionVar = "MONITOR_API_VERSION"
	// MonitorAPIEnvVar is the env var to check for Monitor API token
	MonitorAPIEnvVar = "MONITOR_API_TOKEN"
)

// MonitorClient is a wrapper around the Resty client
type MonitorClient struct {
	resty       *resty.Client
	debug       bool
	apiBaseURL  string
	apiProtocol string
	apiVersion  string
	userAgent   string
}

// NewMonitorClient is the entry point for user to create a new MonitorClient
// It utilizes default values and looks for environment variables to initialize a MonitorClient.
func NewMonitorClient(hc *http.Client) (mClient MonitorClient) {
	if hc != nil {
		mClient.resty = resty.NewWithClient(hc)
	} else {
		mClient.resty = resty.New()
	}

	mClient.SetUserAgent(DefaultUserAgent)

	baseURL, baseURLExists := os.LookupEnv(MonitorAPIHostVar)
	if baseURLExists {
		mClient.SetBaseURL(baseURL)
	} else {
		mClient.SetBaseURL(MonitorAPIHost)
	}

	apiVersion, apiVersionExists := os.LookupEnv(MonitorAPIVersionVar)
	if apiVersionExists {
		mClient.SetAPIVersion(apiVersion)
	} else {
		mClient.SetAPIVersion(MonitorAPIVersion)
	}

	token, apiTokenExists := os.LookupEnv(MonitorAPIEnvVar)
	if apiTokenExists {
		mClient.SetToken(token)
	}

	mClient.SetDebug(envDebug)

	return mClient
}

// SetUserAgent sets a custom user-agent for HTTP requests
func (mc *MonitorClient) SetUserAgent(ua string) *MonitorClient {
	mc.userAgent = ua
	mc.resty.SetHeader("User-Agent", mc.userAgent)

	return mc
}

// R wraps resty's R method
func (mc *MonitorClient) R(ctx context.Context) *resty.Request {
	return mc.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

// SetDebug sets the debug on resty's client
func (mc *MonitorClient) SetDebug(debug bool) *MonitorClient {
	mc.debug = debug
	mc.resty.SetDebug(debug)

	return mc
}

// SetLogger allows the user to override the output
// logger for debug logs.
func (mc *MonitorClient) SetLogger(logger Logger) *MonitorClient {
	mc.resty.SetLogger(logger)

	return mc
}

// SetBaseURL is the helper function to set base url
func (mc *MonitorClient) SetBaseURL(baseURL string) *MonitorClient {
	baseURLPath, _ := url.Parse(baseURL)

	mc.apiBaseURL = path.Join(baseURLPath.Host, baseURLPath.Path)
	mc.apiProtocol = baseURLPath.Scheme

	mc.updateMonitorHostURL()

	return mc
}

// SetAPIVersion is the helper function to set api version
func (mc *MonitorClient) SetAPIVersion(apiVersion string) *MonitorClient {
	mc.apiVersion = apiVersion

	mc.updateMonitorHostURL()

	return mc
}

// SetRootCertificate adds a root certificate to the underlying TLS client config
func (mc *MonitorClient) SetRootCertificate(path string) *MonitorClient {
	mc.resty.SetRootCertificate(path)
	return mc
}

// SetToken sets the API token for all requests from this client
func (mc *MonitorClient) SetToken(token string) *MonitorClient {
	mc.resty.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	return mc
}

// SetHeader sets a custom header to be used in all API requests made with the current client.
// NOTE: Some headers may be overridden by the individual request functions.
func (mc *MonitorClient) SetHeader(name, value string) {
	mc.resty.SetHeader(name, value)
}

func (mc *MonitorClient) updateMonitorHostURL() {
	apiProto := APIProto
	baseURL := MonitorAPIHost
	apiVersion := MonitorAPIVersion

	if mc.apiBaseURL != "" {
		baseURL = mc.apiBaseURL
	}

	if mc.apiVersion != "" {
		apiVersion = mc.apiVersion
	}

	if mc.apiProtocol != "" {
		apiProto = mc.apiProtocol
	}

	mc.resty.SetBaseURL(
		fmt.Sprintf(
			"%s://%s/%s",
			apiProto,
			baseURL,
			url.PathEscape(apiVersion),
		),
	)
}
