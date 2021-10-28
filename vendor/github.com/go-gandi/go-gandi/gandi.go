package gandi

import (
	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/livedns"
)

// Config manages common config for all Gandi API types
type Config struct {
	// SharingID is the Organization ID, available from the Organization API
	SharingID string
	// Debug enables verbose debugging of HTTP calls
	Debug bool
	// DryRun prevents the API from making changes. Only certain API calls support it.
	DryRun bool
}

// NewDomainClient returns a client to the Gandi Domains API
// It expects an API key, available from https://account.gandi.net/en/
func NewDomainClient(apikey string, config Config) *domain.Domain {
	return domain.New(apikey, config.SharingID, config.Debug, config.DryRun)
}

// NewLiveDNSClient returns a client to the Gandi Domains API
// It expects an API key, available from https://account.gandi.net/en/
func NewLiveDNSClient(apikey string, config Config) *livedns.LiveDNS {
	return livedns.New(apikey, config.SharingID, config.Debug, config.DryRun)
}
