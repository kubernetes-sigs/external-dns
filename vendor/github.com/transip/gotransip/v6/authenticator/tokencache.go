package authenticator

import "github.com/transip/gotransip/v6/jwt"

// TokenCache asks for two methods,
// one to save a token by a name
// and one to get a previously acquired token by name returned as jwt.Token
type TokenCache interface {
	// Set will save a token by name as byte array
	Set(key string, token jwt.Token) error
	// Get a previously acquired token by name returned as byte array
	Get(key string) (jwt.Token, error)
}
