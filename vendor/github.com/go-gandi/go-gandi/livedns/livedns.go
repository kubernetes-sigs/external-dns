package livedns

import (
<<<<<<< HEAD
	"github.com/go-gandi/go-gandi/internal/client"
)

// LiveDNS is the API client to the Gandi v5 LiveDNS API
type LiveDNS struct {
	client client.Gandi
}

// New returns an instance of the LiveDNS API client
func New(apikey string, sharingid string, debug bool, dryRun bool) *LiveDNS {
	client := client.New(apikey, sharingid, debug, dryRun)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the LiveDNS API client
func New(config config.Config) *LiveDNS {
	client := client.New(config.APIKey, config.PersonalAccessToken, config.APIURL, config.SharingID, config.Debug, config.DryRun, config.Timeout)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	client.SetEndpoint("livedns/")
	return &LiveDNS{client: *client}
}

// NewFromClient returns an instance of the LiveDNS API client
func NewFromClient(g client.Gandi) *LiveDNS {
	g.SetEndpoint("livedns/")
	return &LiveDNS{client: g}
}
