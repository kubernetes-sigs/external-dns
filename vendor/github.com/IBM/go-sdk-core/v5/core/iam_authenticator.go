package core

// (C) Copyright IBM Corp. 2019.
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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// IamAuthenticator uses an apikey to obtain an IAM access token,
// and adds the access token to requests via an Authorization header
// of the form:
//
// 		Authorization: Bearer <access-token>
//
type IamAuthenticator struct {

	// The apikey used to fetch the bearer token from the IAM token server
	// [required].
	ApiKey string

	// The URL representing the IAM token server's endpoint; If not specified,
	// a suitable default value will be used [optional].
	URL string

	// The ClientId and ClientSecret fields are used to form a "basic auth"
	// Authorization header for interactions with the IAM token server

	// If neither field is specified, then no Authorization header will be sent
	// with token server requests [optional]. These fields are optional, but must
	// be specified together.
	ClientId string

	// If neither field is specified, then no Authorization header will be sent
	// with token server requests [optional]. These fields are optional, but must
	// be specified together.
	ClientSecret string

	// A flag that indicates whether verification of the server's SSL certificate
	// should be disabled; defaults to false [optional].
	DisableSSLVerification bool

	// [Optional] The "scope" to use when fetching the bearer token from the
	// IAM token server.   This can be used to obtain an access token
	// with a specific scope.
	Scope string

	// [Optional] A set of key/value pairs that will be sent as HTTP headers in requests
	// made to the token server.
	Headers map[string]string

	// [Optional] The http.Client object used to invoke token server requests.
	// If not specified by the user, a suitable default Client will be constructed.
	Client *http.Client

	// The cached token and expiration time.
	tokenData *iamTokenData

	// Mutex to make the tokenData field thread safe.
	tokenDataMutex sync.Mutex
}

var iamRequestTokenMutex sync.Mutex
var iamNeedsRefreshMutex sync.Mutex

const (
	// The default (prod) IAM token server base endpoint address.
	defaultIamTokenServerEndpoint = "https://iam.cloud.ibm.com" // #nosec G101
	iamGrantTypeApiKey            = "urn:ibm:params:oauth:grant-type:apikey" // #nosec G101
)

// NewIamAuthenticator constructs a new IamAuthenticator instance.
func NewIamAuthenticator(apikey string, url string, clientId string, clientSecret string,
	disableSSLVerification bool, headers map[string]string) (*IamAuthenticator, error) {
	authenticator := &IamAuthenticator{
		ApiKey:                 apikey,
		URL:                    url,
		ClientId:               clientId,
		ClientSecret:           clientSecret,
		DisableSSLVerification: disableSSLVerification,
		Headers:                headers,
	}

	// Make sure the config is valid.
	err := authenticator.Validate()
	if err != nil {
		return nil, err
	}

	return authenticator, nil
}

// NewIamAuthenticatorFromMap constructs a new IamAuthenticator instance from a
// map.
func newIamAuthenticatorFromMap(properties map[string]string) (authenticator *IamAuthenticator, err error) {
	if properties == nil {
		return nil, fmt.Errorf(ERRORMSG_PROPS_MAP_NIL)
	}

	disableSSL, err := strconv.ParseBool(properties[PROPNAME_AUTH_DISABLE_SSL])
	if err != nil {
		disableSSL = false
	}

	authenticator, err = NewIamAuthenticator(properties[PROPNAME_APIKEY], properties[PROPNAME_AUTH_URL],
		properties[PROPNAME_CLIENT_ID], properties[PROPNAME_CLIENT_SECRET],
		disableSSL, nil)
	if authenticator != nil {
		authenticator.Scope = properties[PROPNAME_SCOPE]
	}
	return
}

// AuthenticationType returns the authentication type for this authenticator.
func (*IamAuthenticator) AuthenticationType() string {
	return AUTHTYPE_IAM
}

// Authenticate adds IAM authentication information to the request.
//
// The IAM bearer token will be added to the request's headers in the form:
//
// 		Authorization: Bearer <bearer-token>
//
func (authenticator *IamAuthenticator) Authenticate(request *http.Request) error {
	token, err := authenticator.GetToken()
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	return nil
}

// getTokenData returns the tokenData field from the authenticator.
func (authenticator *IamAuthenticator) getTokenData() *iamTokenData {
	authenticator.tokenDataMutex.Lock()
	defer authenticator.tokenDataMutex.Unlock()

	return authenticator.tokenData
}

// setTokenData sets the given iamTokenData to the tokenData field of the authenticator.
func (authenticator *IamAuthenticator) setTokenData(tokenData *iamTokenData) {
	authenticator.tokenDataMutex.Lock()
	defer authenticator.tokenDataMutex.Unlock()

	authenticator.tokenData = tokenData
}

// Validate the authenticator's configuration.
//
// Ensures the ApiKey is valid, and the ClientId and ClientSecret pair are
// mutually inclusive.
func (this *IamAuthenticator) Validate() error {
	if this.ApiKey == "" {
		return fmt.Errorf(ERRORMSG_PROP_MISSING, "ApiKey")
	}

	if HasBadFirstOrLastChar(this.ApiKey) {
		return fmt.Errorf(ERRORMSG_PROP_INVALID, "ApiKey")
	}

	// Validate ClientId and ClientSecret.  They must both be specified togther or neither should be specified.
	if this.ClientId == "" && this.ClientSecret == "" {
		// Do nothing as this is the valid scenario
	} else {
		// Since it is NOT the case that both properties are empty, make sure BOTH are specified.
		if this.ClientId == "" {
			return fmt.Errorf(ERRORMSG_PROP_MISSING, "ClientId")
		}

		if this.ClientSecret == "" {
			return fmt.Errorf(ERRORMSG_PROP_MISSING, "ClientSecret")
		}
	}

	return nil
}

// GetToken: returns an access token to be used in an Authorization header.
// Whenever a new token is needed (when a token doesn't yet exist, needs to be refreshed,
// or the existing token has expired), a new access token is fetched from the token server.
func (authenticator *IamAuthenticator) GetToken() (string, error) {
	if authenticator.getTokenData() == nil || !authenticator.getTokenData().isTokenValid() {
		// synchronously request the token
		err := authenticator.synchronizedRequestToken()
		if err != nil {
			return "", err
		}
	} else if authenticator.getTokenData().needsRefresh() {
		// If refresh needed, kick off a go routine in the background to get a new token
		//nolint: errcheck
		go authenticator.invokeRequestTokenData()
	}

	// return an error if the access token is not valid or was not fetched
	if authenticator.getTokenData() == nil || authenticator.getTokenData().AccessToken == "" {
		return "", fmt.Errorf("Error while trying to get access token")
	}

	return authenticator.getTokenData().AccessToken, nil
}

// synchronizedRequestToken: synchronously checks if the current token in cache
// is valid. If token is not valid or does not exist, it will fetch a new token
// and set the tokenRefreshTime
func (authenticator *IamAuthenticator) synchronizedRequestToken() error {
	iamRequestTokenMutex.Lock()
	defer iamRequestTokenMutex.Unlock()
	// if cached token is still valid, then just continue to use it
	if authenticator.getTokenData() != nil && authenticator.getTokenData().isTokenValid() {
		return nil
	}

	return authenticator.invokeRequestTokenData()
}

// invokeRequestTokenData: requests a new token from the access server and
// unmarshals the token information to the tokenData cache. Returns
// an error if the token was unable to be fetched, otherwise returns nil
func (authenticator *IamAuthenticator) invokeRequestTokenData() error {
	tokenResponse, err := authenticator.RequestToken()
	if err != nil {
		return err
	}

	if tokenData, err := newIamTokenData(tokenResponse); err != nil {
		return err
	} else {
		authenticator.setTokenData(tokenData)
	}

	return nil
}

// RequestToken fetches a new access token from the token server.
func (authenticator *IamAuthenticator) RequestToken() (*IamTokenServerResponse, error) {
	var operationPath = "/identity/token"

	// Use the default IAM URL if one was not specified by the user.
	url := authenticator.URL
	if url == "" {
		url = defaultIamTokenServerEndpoint
	} else {
		// Canonicalize the URL by removing the operation path if it was specified by the user.
		url = strings.TrimSuffix(url, operationPath)
	}

	builder := NewRequestBuilder(POST)
	_, err := builder.ResolveRequestURL(url, operationPath, nil)
	if err != nil {
		return nil, err
	}

	builder.AddHeader(CONTENT_TYPE, "application/x-www-form-urlencoded").
		AddHeader(Accept, APPLICATION_JSON).
		AddFormData("grant_type", "", "", iamGrantTypeApiKey).
		AddFormData("apikey", "", "", authenticator.ApiKey).
		AddFormData("response_type", "", "", "cloud_iam")

	// Add any optional parameters to the request.
	if authenticator.Scope != "" {
		builder.AddFormData("scope", "", "", authenticator.Scope)
	}

	// Add user-defined headers to request.
	for headerName, headerValue := range authenticator.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	req, err := builder.Build()
	if err != nil {
		return nil, err
	}

	// If client id and secret were configured by the user, then set them on the request
	// as a basic auth header.
	if authenticator.ClientId != "" && authenticator.ClientSecret != "" {
		req.SetBasicAuth(authenticator.ClientId, authenticator.ClientSecret)
	}

	// If the authenticator does not have a Client, create one now.
	if authenticator.Client == nil {
		authenticator.Client = &http.Client{
			Timeout: time.Second * 30,
		}

		// If the user told us to disable SSL verification, then do it now.
		if authenticator.DisableSSLVerification {
			transport := &http.Transport{
				/* #nosec G402 */
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			authenticator.Client.Transport = transport
		}
	}

	resp, err := authenticator.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		buff := new(bytes.Buffer)
		_, _ = buff.ReadFrom(resp.Body)

		// Create a DetailedResponse to be included in the error below.
		detailedResponse := &DetailedResponse{
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
			RawResult:  buff.Bytes(),
		}

		iamErrorMsg := string(detailedResponse.RawResult)
		if iamErrorMsg == "" {
			iamErrorMsg =
				fmt.Sprintf("unexpected status code %d received from IAM token server %s", detailedResponse.StatusCode, builder.URL)
		}
		return nil, NewAuthenticationError(detailedResponse, fmt.Errorf(iamErrorMsg))
	}

	tokenResponse := &IamTokenServerResponse{}
	_ = json.NewDecoder(resp.Body).Decode(tokenResponse)
	defer resp.Body.Close()
	return tokenResponse, nil
}

// IamTokenServerResponse : This struct models a response received from the token server.
type IamTokenServerResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Expiration   int64  `json:"expiration"`
}

// iamTokenData : This struct represents the cached information related to a fetched access token.
type iamTokenData struct {
	AccessToken string
	RefreshTime int64
	Expiration  int64
}

// newIamTokenData: constructs a new IamTokenData instance from the specified IamTokenServerResponse instance.
func newIamTokenData(tokenResponse *IamTokenServerResponse) (*iamTokenData, error) {

	if tokenResponse == nil {
		return nil, fmt.Errorf("Error while trying to parse access token!")
	}
	// Compute the adjusted refresh time (expiration time - 20% of timeToLive)
	timeToLive := tokenResponse.ExpiresIn
	expireTime := tokenResponse.Expiration
	refreshTime := expireTime - int64(float64(timeToLive)*0.2)

	tokenData := &iamTokenData{
		AccessToken: tokenResponse.AccessToken,
		Expiration:  expireTime,
		RefreshTime: refreshTime,
	}

	return tokenData, nil
}

// isTokenValid: returns true iff the IamTokenData instance represents a valid (non-expired) access token.
func (this *iamTokenData) isTokenValid() bool {
	if this.AccessToken != "" && GetCurrentTime() < this.Expiration {
		return true
	}
	return false
}

// needsRefresh: synchronously returns true iff the currently stored access token should be refreshed. This method also
// updates the refresh time if it determines the token needs refreshed to prevent other threads from
// making multiple refresh calls.
func (this *iamTokenData) needsRefresh() bool {
	iamNeedsRefreshMutex.Lock()
	defer iamNeedsRefreshMutex.Unlock()

	// Advance refresh by one minute
	if this.RefreshTime >= 0 && GetCurrentTime() > this.RefreshTime {
		this.RefreshTime = GetCurrentTime() + 60
		return true
	}

	return false

}
