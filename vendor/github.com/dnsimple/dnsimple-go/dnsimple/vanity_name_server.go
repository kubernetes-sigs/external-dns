package dnsimple

import (
	"context"
	"fmt"
)

// VanityNameServersService handles communication with Vanity Name Servers
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/vanity/
type VanityNameServersService struct {
	client *Client
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// VanityNameServer represents data for a single vanity name server.
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// VanityNameServer represents data for a single vanity name server
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// VanityNameServer represents data for a single vanity name server
=======
// VanityNameServer represents data for a single vanity name server.
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// VanityNameServer represents data for a single vanity name server
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
type VanityNameServer struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IPv4      string `json:"ipv4,omitempty"`
	IPv6      string `json:"ipv6,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func vanityNameServerPath(accountID string, domainIdentifier string) string {
	return fmt.Sprintf("/%v/vanity/%v", accountID, domainIdentifier)
}

// VanityNameServerResponse represents a response for vanity name server enable and disable operations.
type VanityNameServerResponse struct {
	Response
	Data []VanityNameServer `json:"data"`
}

// EnableVanityNameServers Vanity Name Servers for the given domain
//
// See https://developer.dnsimple.com/v2/vanity/#enableVanityNameServers
func (s *VanityNameServersService) EnableVanityNameServers(ctx context.Context, accountID string, domainIdentifier string) (*VanityNameServerResponse, error) {
	path := versioned(vanityNameServerPath(accountID, domainIdentifier))
	vanityNameServerResponse := &VanityNameServerResponse{}

	resp, err := s.client.put(ctx, path, nil, vanityNameServerResponse)
	if err != nil {
		return nil, err
	}

	vanityNameServerResponse.HTTPResponse = resp
	return vanityNameServerResponse, nil
}

// DisableVanityNameServers Vanity Name Servers for the given domain
//
// See https://developer.dnsimple.com/v2/vanity/#disableVanityNameServers
func (s *VanityNameServersService) DisableVanityNameServers(ctx context.Context, accountID string, domainIdentifier string) (*VanityNameServerResponse, error) {
	path := versioned(vanityNameServerPath(accountID, domainIdentifier))
	vanityNameServerResponse := &VanityNameServerResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	vanityNameServerResponse.HTTPResponse = resp
	return vanityNameServerResponse, nil
}
