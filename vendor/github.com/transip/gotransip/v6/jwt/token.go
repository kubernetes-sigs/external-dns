package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Token is the jwt that will be used by the client in the Authorization header
// to send in every request to the api server, every request except for the auth request
// which is used to request a new token.
//
// For more information see: https://jwt.io/
type Token struct {
	ExpiryDate int64
	RawToken   string
}

// Define a hard expiration skew in seconds,
// so we will retrieve a new token way before the moment of expiration
const expirationSkew = 30

// Expired returns true when the token expiry date is reached
func (t *Token) Expired() bool {
	return time.Now().Unix()+expirationSkew > t.ExpiryDate
}

// GetAuthenticationHeaderValue returns the authentication header value value
// including the Bearer prefix
func (t *Token) GetAuthenticationHeaderValue() string {
	return fmt.Sprintf("Bearer %s", t.RawToken)
}

// String returns the string representation of the Token struct
func (t *Token) String() string {
	return t.RawToken
}

// tokenPayload is used to unpack the payload from the jwt
type tokenPayload struct {
	// This ExpirationTime is a 64 bit epoch that will be put into the token struct
	// that will later be used to validate if the token is expired or not
	// once expired we need to request a new token
	ExpirationTime int64 `json:"exp"`
}

// New expects a raw token as string.
// It will try to decode it and return an error on error.
// Once decoded it will retrieve the expiry date and
// return a Token struct with the RawToken and ExpiryDate set.
func New(token string) (Token, error) {
	if len(token) == 0 {
		return Token{}, errors.New("no token given, a token should be set")
	}

	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return Token{}, fmt.Errorf("invalid token '%s' given, token should exist at least of 3 parts", token)
	}

	jsonBody, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return Token{}, errors.New("could not decode token, invalid base64")
	}

	var tokenRequest tokenPayload
	err = json.Unmarshal(jsonBody, &tokenRequest)
	if err != nil {
		return Token{}, errors.New("could not read token body, invalid json")
	}

	return Token{
		RawToken:   token,
		ExpiryDate: tokenRequest.ExpirationTime,
	}, nil
}
