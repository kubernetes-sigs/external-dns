package civogo

import (
	"bytes"
	"encoding/json"
)

// ExchangeAuthTokenRequest contains data that can be passed to ExchangeAuthToken
type ExchangeAuthTokenRequest struct {
	Scope string `json:"scope"`
}

// ExchangeAuthTokenResponse is the response returned by a successful ExchangeAuthToken
type ExchangeAuthTokenResponse struct {
	// The access token issued by the authorization server
	AccessToken string `json:"access_token"`
	// The type of the token issued
	TokenType string `json:"token_type"`
	// The refresh token, which can be used to obtain new access tokens
	RefreshToken string `json:"refresh_token"`
	// The lifetime in seconds of the access token
	ExpiresIn int32 `json:"expires_in"`
	// The ID token, which is a JWT that contains user profile information
	IDToken string `json:"id_token"`
	// The id of the account associated with the provided API Key
	AccountID string `json:"account_id"`
}

// ExchangeAuthToken exchanges an apikey with a new civo JWT
func (c *Client) ExchangeAuthToken(er *ExchangeAuthTokenRequest) (*ExchangeAuthTokenResponse, error) {
	body, err := c.SendPostRequest("/v2/auth/exchange", er)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &ExchangeAuthTokenResponse{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
