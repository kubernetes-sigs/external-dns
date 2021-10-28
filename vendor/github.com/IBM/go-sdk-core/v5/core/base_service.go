package core

// (C) Copyright IBM Corp. 2019, 2021.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

const (
	headerNameUserAgent = "User-Agent"
	sdkName             = "ibm-go-sdk-core"
)

// ServiceOptions is a struct of configuration values for a service.
type ServiceOptions struct {
	// This is the base URL associated with the service instance. This value will
	// be combined with the paths for each operation to form the request URL
	// [required].
	URL string

	// Authenticator holds the authenticator implementation to be used by the
	// service instance to authenticate outbound requests, typically by adding the
	// HTTP "Authorization" header.
	Authenticator Authenticator

	// EnableGzipCompression indicates whether or not request bodies
	// should be gzip-compressed.
	// This field has no effect on response bodies.
	// If enabled, the Body field will be gzip-compressed and
	// the "Content-Encoding" header will be added to the request with the
	// value "gzip".
	EnableGzipCompression bool
}

// BaseService implements the common functionality shared by generated services
// to manage requests and responses, authenticate outbound requests, etc.
type BaseService struct {

	// Configuration values for a service.
	Options *ServiceOptions

	// A set of "default" http headers to be included with each outbound request.
	DefaultHeaders http.Header

	// The HTTP Client used to send requests and receive responses.
	Client *http.Client

	// The value to be used for the "User-Agent" HTTP header that is added to each
	// outbound request. If this value is not set, then a default value will be
	// used for the header.
	UserAgent string
}

// NewBaseService constructs a new instance of BaseService. Validation on input
// parameters and service options will be performed before instance creation.
func NewBaseService(options *ServiceOptions) (*BaseService, error) {
	if HasBadFirstOrLastChar(options.URL) {
		return nil, fmt.Errorf(ERRORMSG_PROP_INVALID, "URL")
	}

	if IsNil(options.Authenticator) {
		return nil, fmt.Errorf(ERRORMSG_NO_AUTHENTICATOR)
	}

	if err := options.Authenticator.Validate(); err != nil {
		return nil, err
	}

	service := BaseService{
		Options: options,

		Client: DefaultHTTPClient(),
	}

	// Set a default value for the User-Agent http header.
	service.SetUserAgent(service.buildUserAgent())

	return &service, nil
}

// Clone will return a copy of "service" suitable for use by a
// generated service instance to process requests.
func (service *BaseService) Clone() *BaseService {
	if IsNil(service) {
		return nil
	}

	// First, copy the service options struct.
	serviceOptions := *service.Options

	// Next, make a copy the service struct, then use the copy of the service options.
	// Note, we'll re-use the "Client" instance from the original BaseService instance.
	clone := *service
	clone.Options = &serviceOptions

	return &clone
}

// ConfigureService updates the service with external configuration values.
func (service *BaseService) ConfigureService(serviceName string) error {
	// Try to load service properties from external config.
	serviceProps, err := getServiceProperties(serviceName)
	if err != nil {
		return err
	}

	// If we were able to load any properties for this service, then check to see if the
	// service-level properties were present and set them on the service if so.
	if serviceProps != nil {

		// URL
		if url, ok := serviceProps[PROPNAME_SVC_URL]; ok && url != "" {
			err := service.SetURL(url)
			if err != nil {
				return err
			}
		}

		// DISABLE_SSL
		if disableSSL, ok := serviceProps[PROPNAME_SVC_DISABLE_SSL]; ok && disableSSL != "" {
			// Convert the config string to bool.
			boolValue, err := strconv.ParseBool(disableSSL)
			if err != nil {
				boolValue = false
			}

			// If requested, disable SSL.
			if boolValue {
				service.DisableSSLVerification()
			}
		}

		// ENABLE_GZIP
		if enableGzip, ok := serviceProps[PROPNAME_SVC_ENABLE_GZIP]; ok && enableGzip != "" {
			// Convert the config string to bool.
			boolValue, err := strconv.ParseBool(enableGzip)
			if err == nil {
				service.SetEnableGzipCompression(boolValue)
			}
		}

		// ENABLE_RETRIES
		// If "ENABLE_RETRIES" is set to true, then we'll also try to retrieve "MAX_RETRIES" and
		// "RETRY_INTERVAL".  If those are not specified, we'll use 0 to trigger a default value for each.
		if enableRetries, ok := serviceProps[PROPNAME_SVC_ENABLE_RETRIES]; ok && enableRetries != "" {
			boolValue, err := strconv.ParseBool(enableRetries)
			if boolValue && err == nil {
				var maxRetries int = 0
				var retryInterval time.Duration = 0

				var s string
				var ok bool
				if s, ok = serviceProps[PROPNAME_SVC_MAX_RETRIES]; ok && s != "" {
					n, err := strconv.ParseInt(s, 10, 32)
					if err == nil {
						maxRetries = int(n)
					}
				}

				if s, ok = serviceProps[PROPNAME_SVC_RETRY_INTERVAL]; ok && s != "" {
					n, err := strconv.ParseInt(s, 10, 32)
					if err == nil {
						retryInterval = time.Duration(n) * time.Second
					}
				}

				service.EnableRetries(maxRetries, retryInterval)
			}
		}
	}
	return nil
}

// SetURL sets the service URL.
//
// Deprecated: use SetServiceURL instead.
func (service *BaseService) SetURL(url string) error {
	return service.SetServiceURL(url)
}

// SetServiceURL sets the service URL.
func (service *BaseService) SetServiceURL(url string) error {
	if HasBadFirstOrLastChar(url) {
		return fmt.Errorf(ERRORMSG_PROP_INVALID, "URL")
	}

	service.Options.URL = url
	return nil
}

// GetServiceURL returns the service URL.
func (service *BaseService) GetServiceURL() string {
	return service.Options.URL
}

// SetDefaultHeaders sets HTTP headers to be sent in every request.
func (service *BaseService) SetDefaultHeaders(headers http.Header) {
	service.DefaultHeaders = headers
}

// SetHTTPClient updates the client handling the requests.
func (service *BaseService) SetHTTPClient(client *http.Client) {
	service.Client = client
}

// DisableSSLVerification skips SSL verification.
// This function sets a new http.Client instance on the service
// and configures it to bypass verification of server certificates
// and host names, making the client susceptible to "man-in-the-middle"
// attacks.  This should be used only for testing.
func (service *BaseService) DisableSSLVerification() {
	client := DefaultHTTPClient()
	tr, ok := client.Transport.(*http.Transport)
	if tr != nil && ok {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402
	}

	service.SetHTTPClient(client)
}

// IsSSLDisabled returns true if and only if the service's http.Client instance
// is configured to skip verification of server SSL certificates.
func (service *BaseService) IsSSLDisabled() bool {
	if service.Client != nil {
		if tr, ok := service.Client.Transport.(*http.Transport); tr != nil && ok {
			if tr.TLSClientConfig != nil {
				return tr.TLSClientConfig.InsecureSkipVerify
			}
		}
	}
	return false
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (service *BaseService) SetEnableGzipCompression(enableGzip bool) {
	service.Options.EnableGzipCompression = enableGzip
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (service *BaseService) GetEnableGzipCompression() bool {
	return service.Options.EnableGzipCompression
}

// buildUserAgent builds the user agent string.
func (service *BaseService) buildUserAgent() string {
	return fmt.Sprintf("%s-%s %s", sdkName, __VERSION__, SystemInfo())
}

// SetUserAgent sets the user agent value.
func (service *BaseService) SetUserAgent(userAgentString string) {
	if userAgentString == "" {
		service.UserAgent = service.buildUserAgent()
	}
	service.UserAgent = userAgentString
}

//
// Request invokes the specified HTTP request and returns the response.
//
// Parameters:
// req: the http.Request object that holds the request information
//
// result: a pointer to the operation result.  This should be one of:
//   - *io.ReadCloser (for a byte-stream type response)
//   - *<primitive>, *[]<primitive>, *map[string]<primitive>
//   - *map[string]json.RawMessage, *[]json.RawMessage
//
// Return values:
// detailedResponse: a DetailedResponse instance containing the status code, headers, etc.
//
// err: a non-nil error object if an error occurred
//
func (service *BaseService) Request(req *http.Request, result interface{}) (detailedResponse *DetailedResponse, err error) {
	// Add default headers.
	if service.DefaultHeaders != nil {
		for k, v := range service.DefaultHeaders {
			req.Header.Add(k, strings.Join(v, ""))
		}

		// After adding the default headers, make one final check to see if the user
		// specified the "Host" header within the default headers.
		// This needs to be handled separately because it will be ignored by
		// the Request.Write() method.
		host := service.DefaultHeaders.Get("Host")
		if host != "" {
			req.Host = host
		}
	}

	// Add the default User-Agent header if not already present.
	userAgent := req.Header.Get(headerNameUserAgent)
	if userAgent == "" {
		req.Header.Add(headerNameUserAgent, service.UserAgent)
	}

	// Add authentication to the outbound request.
	if IsNil(service.Options.Authenticator) {
		err = fmt.Errorf(ERRORMSG_NO_AUTHENTICATOR)
		return
	}

	authError := service.Options.Authenticator.Authenticate(req)
	if authError != nil {
		err = fmt.Errorf(ERRORMSG_AUTHENTICATE_ERROR, authError.Error())
		castErr, ok := authError.(*AuthenticationError)
		if ok {
			detailedResponse = castErr.Response
		}
		return
	}

	// If debug is enabled, then dump the request.
	if GetLogger().IsLogLevelEnabled(LevelDebug) {
		buf, dumpErr := httputil.DumpRequestOut(req, req.Body != nil)
		if dumpErr == nil {
			GetLogger().Debug("Request:\n%s\n", string(buf))
		} else {
			GetLogger().Debug("error while attempting to log outbound request: %s", dumpErr.Error())
		}
	}

	var httpResponse *http.Response

	// Try to get the retryable Client hidden inside service.Client
	retryableClient := getRetryableHTTPClient(service.Client)
	if retryableClient != nil {
		retryableRequest, retryableErr := retryablehttp.FromRequest(req)
		if retryableErr != nil {
			err = fmt.Errorf(ERRORMSG_CREATE_RETRYABLE_REQ, retryableErr.Error())
			return
		}

		// Invoke the retryable request.
		httpResponse, err = retryableClient.Do(retryableRequest)
	} else {
		// Invoke the normal (non-retryable) request.
		httpResponse, err = service.Client.Do(req)
	}

	// Check for errors during the invocation.
	if err != nil {
		if strings.Contains(err.Error(), SSL_CERTIFICATION_ERROR) {
			err = fmt.Errorf(ERRORMSG_SSL_VERIFICATION_FAILED + "\n" + err.Error())
		}
		return
	}

	// If debug is enabled, then dump the response.
	if GetLogger().IsLogLevelEnabled(LevelDebug) {
		buf, dumpErr := httputil.DumpResponse(httpResponse, httpResponse.Body != nil)
		if err == nil {
			GetLogger().Debug("Response:\n%s\n", string(buf))
		} else {
			GetLogger().Debug("error while attempting to log inbound response: %s", dumpErr.Error())
		}
	}

	// Start to populate the DetailedResponse.
	detailedResponse = &DetailedResponse{
		StatusCode: httpResponse.StatusCode,
		Headers:    httpResponse.Header,
	}

	contentType := httpResponse.Header.Get(CONTENT_TYPE)

	// If the operation was unsuccessful, then set up the DetailedResponse
	// and error objects appropriately.
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {

		var responseBody []byte

		// First, read the response body into a byte array.
		if httpResponse.Body != nil {
			var readErr error

			defer httpResponse.Body.Close()
			responseBody, readErr = ioutil.ReadAll(httpResponse.Body)
			if readErr != nil {
				err = fmt.Errorf(ERRORMSG_READ_RESPONSE_BODY, readErr.Error())
				return
			}
		}

		// If the responseBody is empty, then just return a generic error based on the status code.
		if len(responseBody) == 0 {
			err = fmt.Errorf(http.StatusText(httpResponse.StatusCode))
			return
		}

		// For a JSON-based error response body, decode it into a map (generic JSON object).
		if IsJSONMimeType(contentType) {
			// Return the error response body as a map, along with an
			// error object containing our best guess at an error message.
			responseMap, decodeErr := decodeAsMap(responseBody)
			if decodeErr == nil {
				detailedResponse.Result = responseMap
				err = fmt.Errorf(getErrorMessage(responseMap, detailedResponse.StatusCode))
				return
			}
		}

		// For a non-JSON response or if we tripped while decoding the JSON response,
		// just return the response body byte array in the RawResult field along with
		// an error object that contains the generic error message for the status code.
		detailedResponse.RawResult = responseBody
		err = fmt.Errorf(http.StatusText(httpResponse.StatusCode))
		return
	}

	// Operation was successful and we are expecting a response, so process the response.
	if !IsNil(result) {
		resultType := reflect.TypeOf(result).String()

		// If 'result' is a io.ReadCloser, then pass the response body back reflectively via 'result'
		// and bypass any further unmarshalling of the response.
		if resultType == "*io.ReadCloser" {
			rResult := reflect.ValueOf(result).Elem()
			rResult.Set(reflect.ValueOf(httpResponse.Body))
			detailedResponse.Result = httpResponse.Body
		} else {

			// First, read the response body into a byte array.
			defer httpResponse.Body.Close()
			responseBody, readErr := ioutil.ReadAll(httpResponse.Body)
			if readErr != nil {
				err = fmt.Errorf(ERRORMSG_READ_RESPONSE_BODY, readErr.Error())
				return
			}

			// If the response body is empty, then skip any attempt to deserialize and just return
			if len(responseBody) == 0 {
				return
			}

			// If the content-type indicates JSON, then unmarshal the response body as JSON.
			if IsJSONMimeType(contentType) {
				// Decode the byte array as JSON.
				decodeErr := json.NewDecoder(bytes.NewReader(responseBody)).Decode(result)
				if decodeErr != nil {
					// Error decoding the response body.
					// Return the response body in RawResult, along with an error.
					err = fmt.Errorf(ERRORMSG_UNMARSHAL_RESPONSE_BODY, decodeErr.Error())
					detailedResponse.RawResult = responseBody
					return
				}

				// Decode step was successful. Return the decoded response object in the Result field.
				detailedResponse.Result = reflect.ValueOf(result).Elem().Interface()
				return
			}

			// Check to see if the caller wanted the response body as a string.
			// If the caller passed in 'result' as the address of *string,
			// then we'll reflectively set result to point to it.
			if resultType == "**string" {
				responseString := string(responseBody)
				rResult := reflect.ValueOf(result).Elem()
				rResult.Set(reflect.ValueOf(&responseString))

				// And set the string in the Result field.
				detailedResponse.Result = &responseString
			} else if resultType == "*[]uint8" { // byte is an alias for uint8
				rResult := reflect.ValueOf(result).Elem()
				rResult.Set(reflect.ValueOf(responseBody))

				// And set the byte slice in the Result field.
				detailedResponse.Result = responseBody
			} else {
				// At this point, we don't know how to set the result field, so we have to return an error.
				// But make sure we save the bytes we read in the DetailedResponse for debugging purposes
				detailedResponse.Result = responseBody
				err = fmt.Errorf(ERRORMSG_UNEXPECTED_RESPONSE, contentType, resultType)
				return
			}
		}
	}

	return
}

// Errors is a struct used to hold an array of errors received in an operation
// response.
type Errors struct {
	Errors []Error `json:"errors,omitempty"`
}

// Error is a struct used to represent a single error received in an operation
// response.
type Error struct {
	Message string `json:"message,omitempty"`
}

// decodeAsMap: Decode the specified JSON byte-stream into a map (akin to a generic JSON object).
// Notes:
// 1) This function will return the map (result of decoding the byte-stream) as well as the raw
// byte buffer.  We return the byte buffer in addition to the decoded map so that the caller can
// re-use (if necessary) the stream of bytes after we've consumed them via the JSON decode step.
// 2) The primary return value of this function will be:
//    a) an instance of map[string]interface{} if the specified byte-stream was successfully
//       decoded as JSON.
//    b) the string form of the byte-stream if the byte-stream could not be successfully
//       decoded as JSON.
// 3) This function will close the io.ReadCloser before returning.
func decodeAsMap(byteBuffer []byte) (result map[string]interface{}, err error) {
	err = json.NewDecoder(bytes.NewReader(byteBuffer)).Decode(&result)
	return
}

// getErrorMessage: try to retrieve an error message from the decoded response body (map).
func getErrorMessage(responseMap map[string]interface{}, statusCode int) string {

	// If the response contained the "errors" field, then try to deserialize responseMap
	// into an array of Error structs, then return the first entry's "Message" field.
	if _, ok := responseMap["errors"]; ok {
		var errors Errors
		responseBuffer, _ := json.Marshal(responseMap)
		if err := json.Unmarshal(responseBuffer, &errors); err == nil {
			return errors.Errors[0].Message
		}
	}

	// Return the "error" field if present and is a string.
	if val, ok := responseMap["error"]; ok {
		errorMsg, ok := val.(string)
		if ok {
			return errorMsg
		}
	}

	// Return the "message" field if present and is a string.
	if val, ok := responseMap["message"]; ok {
		errorMsg, ok := val.(string)
		if ok {
			return errorMsg
		}
	}

	// Finally, return the "errorMessage" field if present and is a string.
	if val, ok := responseMap["errorMessage"]; ok {
		errorMsg, ok := val.(string)
		if ok {
			return errorMsg
		}
	}

	// If we couldn't find an error message above, just return the generic text
	// for the status code.
	return http.StatusText(statusCode)
}

// EnableRetries will construct a "retryable" HTTP Client with the specified
// configuration, and then set it on the service instance.
// If maxRetries and/or maxRetryInterval are specified as 0, then default values
// are used instead.
func (service *BaseService) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	client := NewRetryableHTTPClient()
	if maxRetries > 0 {
		client.RetryMax = maxRetries
	}
	if maxRetryInterval > 0 {
		client.RetryWaitMax = maxRetryInterval
	}

	service.SetHTTPClient(client.StandardClient())
}

// DisableRetries will disable automatic retries by constructing a new
// default (non-retryable) HTTP Client instance and setting it on the service.
func (service *BaseService) DisableRetries() {
	service.SetHTTPClient(DefaultHTTPClient())
}

// DefaultHTTPClient returns a non-retryable http client with default configuration.
func DefaultHTTPClient() *http.Client {
	return cleanhttp.DefaultPooledClient()
}

// httpLogger is a shim layer used to allow the Go core's logger to be used with the retryablehttp interfaces.
type httpLogger struct {
}

func (l *httpLogger) Printf(format string, inserts ...interface{}) {
	GetLogger().Log(LevelDebug, format, inserts...)
}

// NewRetryableHTTPClient returns a new instance of go-retryablehttp.Client
// with a default configuration that supports Go SDK usage.
func NewRetryableHTTPClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.Logger = &httpLogger{}
	client.CheckRetry = IBMCloudSDKRetryPolicy
	client.Backoff = IBMCloudSDKBackoffPolicy
	client.ErrorHandler = retryablehttp.PassthroughErrorHandler
	return client
}

// getRetryableHTTPClient returns the "retryable" Client hidden inside the specified http.Client instance
// or nil if "client" is not hiding a retryable Client instance.
func getRetryableHTTPClient(client *http.Client) *retryablehttp.Client {
	if client != nil {
		if client.Transport != nil {
			// A retryable client will have its Transport field set to an
			// instance of retryablehttp.RoundTripper.
			if rt, ok := client.Transport.(*retryablehttp.RoundTripper); ok {
				return rt.Client
			}
		}
	}
	return nil
}

var (
	// A regular expression to match the error returned by net/http when the
	// configured number of redirects is exhausted. This error isn't typed
	// specifically so we resort to matching on the error string.
	redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

	// A regular expression to match the error returned by net/http when the
	// scheme specified in the URL is invalid. This error isn't typed
	// specifically so we resort to matching on the error string.
	schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)
)

// IBMCloudSDKRetryPolicy provides a default implementation of the CheckRetry interface
// associated with a retryablehttp.Client.
// This function will return true if the specified request/response should be retried.
func IBMCloudSDKRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// This logic was adapted from go-relyablehttp.ErrorPropagatedRetryPolicy().

	// Do not retry on a Context-related error (Canceled or DeadlineExceeded).
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// Next, check for a few non-retryable errors.
	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// Don't retry if the error was due to too many redirects.
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid protocol scheme.
			if schemeErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to TLS cert verification failure.
			if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
				return false, v
			}
		}

		// The error is likely recoverable so retry.
		return true, nil
	}

	// Now check the status code.

	// A 429 should be retryable.
	if resp.StatusCode == 429 {
		return true, nil
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
		return true, fmt.Errorf(ERRORMSG_UNEXPECTED_STATUS_CODE, resp.StatusCode, resp.Status)
	}

	return false, nil
}

// IBMCloudSDKBackoffPolicy provides a default implementation of the Backoff interface
// associated with a retryablehttp.Client.
// This function will return the wait time to be associated with the next retry attempt.
func IBMCloudSDKBackoffPolicy(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	// Check for a Retry-After header.
	if resp != nil {
		if s, ok := resp.Header["Retry-After"]; ok {
			// First, try to parse the value as an integer (number of seconds to wait)
			if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
				return time.Second * time.Duration(sleep)
			}

			// Otherwise, try to parse the value as an HTTP Time value.
			if retryTime, err := http.ParseTime(s[0]); err == nil {
				sleep := time.Until(retryTime)
				if sleep > max {
					sleep = max
				}
				return sleep
			}

		}
	}

	// If no header-based wait time can be determined, then ask DefaultBackoff()
	// to compute an exponential backoff.
	return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
}
