package authenticator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/transip/gotransip/v6/jwt"
	"github.com/transip/gotransip/v6/rest"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// this is the header key we will add the signature to
	signatureHeader = "Signature"
	// this prefix will be used to name tokens we requested
	// customers are able to see this in their control panel
	labelPrefix = "gotransip-client"
	// authenticationPath is the endpoint that the authenticator
	// will communicate with
	authenticationPath = "/auth"
	// a requested Token expires after a day by default
	// will be used if Authenticator.TokenExpiration is not set
	defaultTokenExpiration = "1 day"
	// DemoToken can be used to test with the api without using your own account
	DemoToken = `eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImN3MiFSbDU2eDNoUnkjelM4YmdOIn0.` +
		`eyJpc3MiOiJhcGkudHJhbnNpcC5ubCIsImF1ZCI6ImFwaS50cmFuc2lwLm5sIiwianRpIjoiY3cy` +
		`IVJsNTZ4M2hSeSN6UzhiZ04iLCJpYXQiOjE1ODIyMDE1NTAsIm5iZiI6MTU4MjIwMTU1MCwiZXhw` +
		`IjoyMTE4NzQ1NTUwLCJjaWQiOiI2MDQ0OSIsInJvIjpmYWxzZSwiZ2siOmZhbHNlLCJrdiI6dHJ1` +
		`ZX0.fYBWV4O5WPXxGuWG-vcrFWqmRHBm9yp0PHiYh_oAWxWxCaZX2Rf6WJfc13AxEeZ67-lY0TA2` +
		`kSaOCp0PggBb_MGj73t4cH8gdwDJzANVxkiPL1Saqiw2NgZ3IHASJnisUWNnZp8HnrhLLe5ficvb` +
		`1D9WOUOItmFC2ZgfGObNhlL2y-AMNLT4X7oNgrNTGm-mespo0jD_qH9dK5_evSzS3K8o03gu6p19` +
		`jxfsnIh8TIVRvNdluYC2wo4qDl5EW5BEZ8OSuJ121ncOT1oRpzXB0cVZ9e5_UVAEr9X3f26_Eomg` +
		`52-PjrgcRJ_jPIUYbrlo06KjjX2h0fzMr21ZE023Gw`
)

var (
	// ErrTokenExpired will be throwed when the static token that has been set by the client is expired
	// and we cannot request a new one
	ErrTokenExpired = errors.New("token expired and no private key is set")
)

// Authenticator is used to store,retrieve and request new tokens on every request.
// It checks the expiry date of a Token and if it is expired it will request a new one
type Authenticator struct {
	// this contains a []byte representation of the the private key of the customer
	// this key will be used to sign a new Token request
	PrivateKeyBody []byte
	// this is Token, that is filled with a static Token that a customer provides
	// or a Token that we got from a Token request
	Token jwt.Token
	// this is the http client to do auth requests with
	HTTPClient *http.Client
	// this would be the auth path, thus where we will get new tokens from
	BasePath string
	// this would be the account name of customer
	Login string
	// When this is set to true the requested tokens can only be used with the 'ip' we are requesting with
	Whitelisted bool
	// Whether or not we want to request read only Tokens, that can only only be used to retrieve information
	// not to create, modify or delete it
	ReadOnly bool
	// TokenCache is used to retrieve previously acquired tokens and saving new ones
	// If not set we do not use a cache to store the tokens
	TokenCache TokenCache
	// TokenExpiration defines lifetime of generated tokens.
	// If unspecified, the default is 1 day.
	// Has no effect for tokens provided via the Token field
	TokenExpiration time.Duration
}

// AuthRequest will be transformed and send in order to request a new Token
// for more information, see: https://api.transip.nl/rest/docs.html#header-authentication
type AuthRequest struct {
	// Account name
	Login string `json:"login"`
	// Unique number for this request
	Nonce string `json:"nonce"`
	// Custom name to give this Token, you can see your tokens in the transip control panel
	Label string `json:"label,omitempty"`
	// Enable read only mode
	ReadOnly bool `json:"read_only"`
	// Unix time stamp of when this Token should expire
	ExpirationTime string `json:"expiration_time"`
	// Whether this key can be used from everywhere, e.g should not be whitelisted to the current requesting ip
	GlobalKey bool `json:"global_key"`
}

// GetToken will return the current Token if it is not expired.
// If it is expired it will try to request a new Token, set and return that.
func (a *Authenticator) GetToken() (jwt.Token, error) {
	// If token is not set, and we have a token cache,
	// try to retrieve it from the token cache
	if a.Token.ExpiryDate == 0 && a.TokenCache != nil {
		if err := a.retrieveTokenFromCache(); err != nil {
			return jwt.Token{}, err
		}
	}

	if a.Token.Expired() && a.PrivateKeyBody == nil {
		return jwt.Token{}, ErrTokenExpired
	}
	if a.Token.Expired() {
		var err error
		a.Token, err = a.requestNewToken()

		if err != nil {
			return jwt.Token{}, err
		}

		// if a TokenCache is set we want to write acquired tokens to the cache
		if a.TokenCache != nil {
			if err = a.TokenCache.Set(a.getTokenCacheKey(), a.Token); err != nil {
				return jwt.Token{}, fmt.Errorf("error writing token to cache: %w", err)
			}
		}
	}

	return a.Token, nil
}

// retrieveTokenFromCache gets the token from the cache
func (a *Authenticator) retrieveTokenFromCache() error {
	var err error
	a.Token, err = a.TokenCache.Get(a.getTokenCacheKey())
	if err != nil {
		return fmt.Errorf("error getting token from cache: %w", err)
	}

	return nil
}

// requestNewToken will request a new Token using the http client
// creating a new AuthRequest, converting it to json and sending that to the api auth url
// on error it will pass this back
func (a *Authenticator) requestNewToken() (jwt.Token, error) {
	restRequest, err := a.getAuthRequest()
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error during auth request creation: %w", err)
	}

	getMethod := rest.PostMethod

	httpRequest, err := restRequest.GetHTTPRequest(a.BasePath, getMethod.Method)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error constructing token http request: %w", err)
	}
	bodyToSign, err := restRequest.GetJSONBody()
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error marshalling token request: %w", err)
	}
	signature, err := signWithKey(bodyToSign, a.PrivateKeyBody)
	if err != nil {
		return jwt.Token{}, err
	}
	httpRequest.Header.Add(signatureHeader, signature)

	httpResponse, err := a.HTTPClient.Do(httpRequest)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error requesting token: %w", err)
	}

	defer httpResponse.Body.Close()

	// read entire response body
	b, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error requesting token: %w", err)
	}

	restResponse := rest.Response{
		Body:       b,
		StatusCode: httpResponse.StatusCode,
		Method:     getMethod,
	}

	var tokenToReturn tokenResponse
	err = restResponse.ParseResponse(&tokenToReturn)
	if err != nil {
		return jwt.Token{}, fmt.Errorf("error requesting token: %w", err)
	}

	return jwt.New(tokenToReturn.Token)
}

// tokenResponse is used to extract a Token from the api server response
type tokenResponse struct {
	Token string `json:"Token"`
}

// getNonce returns a random 16 character length string nonce
// each time it is called
func (a *Authenticator) getNonce() (string, error) {
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", fmt.Errorf("error when getting random data for new nonce: %w", err)
	}

	// convert to hex
	return fmt.Sprintf("%02x", randomBytes), nil
}

// getAuthRequest returns a rest.Request filled with a new AuthRequest
func (a *Authenticator) getAuthRequest() (rest.Request, error) {
	labelPostFix := time.Now().UnixNano()

	nonce, err := a.getNonce()
	if err != nil {
		return rest.Request{}, err
	}

	authRequest := AuthRequest{
		Login:          a.Login,
		Nonce:          nonce,
		Label:          fmt.Sprintf("%s-%d", labelPrefix, labelPostFix),
		ReadOnly:       a.ReadOnly,
		ExpirationTime: a.getTokenExpirationString(),
		GlobalKey:      !a.Whitelisted,
	}

	return rest.Request{
		Endpoint: authenticationPath,
		Body:     authRequest,
	}, nil
}

// getTokenCacheKey returns a name for the given Login and our authenticator name
func (a *Authenticator) getTokenCacheKey() string {
	return fmt.Sprintf("%s-%s-token", labelPrefix, a.Login)
}

// getTokenExpirationString returns the requested or default expiration in string format for the API
func (a *Authenticator) getTokenExpirationString() string {
	if a.TokenExpiration != time.Duration(0) {
		return fmt.Sprintf("%0.0f seconds", a.TokenExpiration.Seconds())
	}
	return defaultTokenExpiration
}
