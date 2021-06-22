package gotransip

import (
	"github.com/transip/gotransip/v6/authenticator"
	"io"
	"net/http"
	"time"
)

const (
	libraryVersion  = "6.6.0"
	defaultBasePath = "https://api.transip.nl/v6"
	userAgent       = "go-client-gotransip/" + libraryVersion
)

// APIMode specifies in which mode the API is used. Currently this is only
// supports either readonly or readwrite
type APIMode string

var (
	// APIModeReadOnly specifies that no changes can be made from API calls.
	// If you do try to order a product or change some data, the api will return an error.
	APIModeReadOnly APIMode = "readonly"
	// APIModeReadWrite specifies that changes can be made from API calls
	APIModeReadWrite APIMode = "readwrite"
)

// DemoClientConfiguration is the default configuration to use when testing the demo mode of the transip api.
// Demo mode allows users to test without authenticating with their own credentials.
var DemoClientConfiguration = ClientConfiguration{Token: authenticator.DemoToken}

// ClientConfiguration stores the configuration of the API client
type ClientConfiguration struct {
	// AccountName is the name of the account of the user, this is used in combination with a private key.
	// When requesting a new token, the account name will be part of the token request body
	AccountName string
	// URL is set by default to the transip api server
	// this is mainly used in tests to point to a mock server
	URL string
	// PrivateKeyPath is the filesystem location to the private key
	PrivateKeyPath string
	// For users that want the possibility to store their key elsewhere,
	// not on a filesystem but on X datastore
	PrivateKeyReader io.Reader
	// Token field gives users the option of providing their own acquired token,
	// for example when generated in the transip control panel
	Token string
	// TestMode is used when users want to tinker with the api without touching their real data.
	// So you can view your own data, order new products, but the actual order never happens.
	TestMode bool
	// optionally you can set your own HTTPClient
	// to set extra non default settings
	HTTPClient *http.Client
	// APIMode specifies in which mode the API is used. Currently this is only
	// supports either readonly or readwrite
	Mode APIMode
	// TokenCache is used to retrieve previously acquired tokens and saving new ones
	// If not set we do not use a cache to store the new acquired tokens
	TokenCache authenticator.TokenCache
	// TokenExpiration defines the lifetime of new tokens requested by the authenticator.
	// If unspecified, the default is 1 day.
	// This has no effect for tokens provided via the Token field.
	TokenExpiration time.Duration
	// TokenWhitelisted is used to indicate only whitelisted IP's may use the new tokens requested by the authenticator.
	// This has no effect for tokens provided via the Token field.
	TokenWhitelisted bool
}
