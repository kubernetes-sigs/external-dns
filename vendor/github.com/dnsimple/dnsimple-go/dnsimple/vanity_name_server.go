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

// VanityNameServer represents data for a single vanity name server
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
