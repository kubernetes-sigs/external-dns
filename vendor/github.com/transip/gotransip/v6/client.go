package gotransip

import (
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/authenticator"
	"github.com/transip/gotransip/v6/jwt"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// client manages communication with the TransIP API
// In most cases there should be only one, shared, client.
type client struct {
	// client configuration file, allows you to:
	// - setting a custom useragent
	// - enable test mode
	// - use the demo token
	// - enable debugging
	config ClientConfiguration
	// authenticator wraps all authentication logic
	// - checking if the token is not expired yet
	// - creating an authentication request
	// - requesting and setting a new token
	authenticator *authenticator.Authenticator
}

// httpBodyLimit provides a maximum byte limit around the http body reader.
// If a request somehow ends up having a huge response body you load all of that data into memory.
// We do not expect to hit this extreme high number even when serving things like PDFs
const httpBodyLimit = 1024 * 1024 * 4

// NewClient creates a new API client.
// optionally you could put a custom http.client in the configuration struct
// to allow for advanced features such as caching.
func NewClient(config ClientConfiguration) (repository.Client, error) {
	return newClient(config)
}

// newClient method is used internally for testing,
// the NewClient method is exported as it follows the repository.Client interface
// which is so that we don't have to bind to this specific implementation
func newClient(config ClientConfiguration) (*client, error) {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}
	var privateKeyBody []byte
	var token jwt.Token

	// check account name
	if len(config.AccountName) == 0 && len(config.Token) == 0 {
		return &client{}, errors.New("AccountName is required")
	}

	// if a private key path is specified and a private key reader is not we
	// fill the private key reader with a opened file on the given PrivateKeyPath
	if len(config.PrivateKeyPath) > 0 && config.PrivateKeyReader == nil {
		privateKeyFile, err := os.Open(config.PrivateKeyPath)
		config.PrivateKeyReader = privateKeyFile

		if err != nil {
			return &client{}, fmt.Errorf("error while opening private key file: %w", err)
		}
	}

	// check if token or private key is set
	if len(config.Token) == 0 && config.PrivateKeyReader == nil {
		return &client{}, errors.New("PrivateKeyReader, token or PrivateKeyReader is required")
	}

	if config.PrivateKeyReader != nil {
		var err error
		privateKeyBody, err = ioutil.ReadAll(config.PrivateKeyReader)

		if err != nil {
			return &client{}, fmt.Errorf("error while reading private key: %w", err)
		}
	}

	if len(config.Token) > 0 {
		var err error
		token, err = jwt.New(config.Token)

		if err != nil {
			return &client{}, err
		}
	}

	// default to APIMode read/write
	if len(config.Mode) == 0 {
		config.Mode = APIModeReadWrite
	}

	// set defaultBasePath by default
	if len(config.URL) == 0 {
		config.URL = defaultBasePath
	}

	return &client{
		authenticator: &authenticator.Authenticator{
			Login:           config.AccountName,
			PrivateKeyBody:  privateKeyBody,
			Token:           token,
			HTTPClient:      config.HTTPClient,
			TokenCache:      config.TokenCache,
			BasePath:        config.URL,
			ReadOnly:        config.Mode == APIModeReadOnly,
			TokenExpiration: config.TokenExpiration,
			Whitelisted:     config.TokenWhitelisted,
		},
		config: config,
	}, nil
}

// This method is used by all rest client methods, thus: 'get','post','put','delete'
// It uses the authenticator to get a token, either statically provided by the user or requested from the authentication server
// Then decodes the json response to a supplied interface
func (c *client) call(method rest.Method, request rest.Request, result interface{}) error {
	token, err := c.authenticator.GetToken()
	if err != nil {
		return fmt.Errorf("could not get token from authenticator: %w", err)
	}

	// if test mode is enabled we always want to change rest requests to add a HTTP test=1 query string
	// to a HTTP request
	if c.config.TestMode {
		request.TestMode = true
	}

	httpRequest, err := request.GetHTTPRequest(c.config.URL, method.Method)
	if err != nil {
		return fmt.Errorf("error during request creation: %w", err)
	}

	httpRequest.Header.Add("Authorization", token.GetAuthenticationHeaderValue())
	httpRequest.Header.Set("User-Agent", userAgent)
	client := c.config.HTTPClient
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	defer httpResponse.Body.Close()

	bodyReader := io.LimitReader(httpResponse.Body, httpBodyLimit)

	// read entire httpResponse body
	b, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return fmt.Errorf("error reading http response body: %w", err)
	}

	restResponse := rest.Response{
		Body:       b,
		StatusCode: httpResponse.StatusCode,
		Method:     method,
	}

	return restResponse.ParseResponse(&result)
}

// ChangeBasePath changes base path to allow switching to mocks
func (c *client) ChangeBasePath(path string) {
	c.config.URL = path
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *client) GetConfig() ClientConfiguration {
	return c.config
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *client) GetAuthenticator() *authenticator.Authenticator {
	return c.authenticator
}

// This method will create and execute a http Get request
func (c *client) Get(request rest.Request, responseObject interface{}) error {
	return c.call(rest.GetMethod, request, responseObject)
}

// This method will create and execute a http Post request
// It expects no response, that is why it does not ask for a responseObject
func (c *client) Post(request rest.Request) error {
	return c.call(rest.PostMethod, request, nil)
}

// This method will create and execute a http Put request
// It expects no response, that is why it does not ask for a responseObject
func (c *client) Put(request rest.Request) error {
	return c.call(rest.PutMethod, request, nil)
}

// This method will create and execute a http Delete request
// It expects no response, that is why it does not ask for a responseObject
func (c *client) Delete(request rest.Request) error {
	return c.call(rest.DeleteMethod, request, nil)
}

// This method will create and execute a http Patch request
// It expects no response, that is why it does not ask for a responseObject
func (c *client) Patch(request rest.Request) error {
	return c.call(rest.PatchMethod, request, nil)
}
