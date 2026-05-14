package credentials

import (
	"errors"
	"fmt"
)

var (
	ErrNoValidCredentialProviders = errors.New("no valid credential providers")
)

// A ChainProvider will search for a provider which returns credentials
// and cache that provider until Retrieve is called again.
type ChainProvider struct {
	Providers []Provider
	current   Provider
}

// NewChainCredentials returns a pointer to a new Credentials object
// wrapping a chain of providers.
func NewChainCredentials(providers []Provider) *Credentials {
	return NewCredentials(&ChainProvider{
		Providers: append([]Provider{}, providers...),
	})
}

// Retrieve returns the first provider in the chain that succeeds,
// or error if no provider returned.
//
// If a provider is found it will be cached and any calls to IsExpired()
// will return the expired state of the cached provider.
func (c *ChainProvider) Retrieve() (Value, error) {
	var errs = ErrNoValidCredentialProviders

	for _, p := range c.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			c.current = p
			return creds, nil
		}

		errs = fmt.Errorf("%v: %w", errs, err)
	}
	c.current = nil

	return Value{}, fmt.Errorf("chain provider: %w", errs)
}

// IsExpired will returned the expired state of the currently cached provider
// if there is one.  If there is no current provider, true will be returned.
func (c *ChainProvider) IsExpired() bool {
	if c.current != nil {
		return c.current.IsExpired()
	}

	return true
}
