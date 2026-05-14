package livedns

import (
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the LiveDNS API client
func New(config config.Config) *LiveDNS {
	client := client.New(config.APIKey, config.PersonalAccessToken, config.APIURL, config.SharingID, config.Debug, config.DryRun, config.Timeout)
	client.SetEndpoint("livedns/")
	return &LiveDNS{client: *client}
}

// NewFromClient returns an instance of the LiveDNS API client
func NewFromClient(g client.Gandi) *LiveDNS {
	g.SetEndpoint("livedns/")
	return &LiveDNS{client: g}
}
