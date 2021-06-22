package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const netPath = "/v2/private-networks"

// NetworkService is the interface to interact with the network endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/private-Networks
type NetworkService interface {
	Create(ctx context.Context, createReq *NetworkReq) (*Network, error)
	Get(ctx context.Context, networkID string) (*Network, error)
	Update(ctx context.Context, networkID string, description string) error
	Delete(ctx context.Context, networkID string) error
	List(ctx context.Context, options *ListOptions) ([]Network, *Meta, error)
}

// NetworkServiceHandler handles interaction with the network methods for the Vultr API
type NetworkServiceHandler struct {
	client *Client
}

// Network represents a Vultr private network
type Network struct {
	NetworkID    string `json:"id"`
	Region       string `json:"region"`
	Description  string `json:"description"`
	V4Subnet     string `json:"v4_subnet"`
	V4SubnetMask int    `json:"v4_subnet_mask"`
	DateCreated  string `json:"date_created"`
}

// NetworkReq represents parameters to create or update Network resource
type NetworkReq struct {
	Region       string `json:"region"`
	Description  string `json:"description"`
	V4Subnet     string `json:"v4_subnet"`
	V4SubnetMask int    `json:"v4_subnet_mask"`
}

type networksBase struct {
	Networks []Network `json:"networks"`
	Meta     *Meta     `json:"meta"`
}

type networkBase struct {
	Network *Network `json:"network"`
}

// Create a new private network. A private network can only be used at the location for which it was created.
func (n *NetworkServiceHandler) Create(ctx context.Context, createReq *NetworkReq) (*Network, error) {
	req, err := n.client.NewRequest(ctx, http.MethodPost, netPath, createReq)
	if err != nil {
		return nil, err
	}

	network := new(networkBase)
	if err = n.client.DoWithContext(ctx, req, network); err != nil {
		return nil, err
	}

	return network.Network, nil
}

// Get gets the private networks of the requested ID
func (n *NetworkServiceHandler) Get(ctx context.Context, networkID string) (*Network, error) {
	uri := fmt.Sprintf("%s/%s", netPath, networkID)
	req, err := n.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	network := new(networkBase)
	if err = n.client.DoWithContext(ctx, req, network); err != nil {
		return nil, err
	}

	return network.Network, nil
}

// Update updates a private network
func (n *NetworkServiceHandler) Update(ctx context.Context, networkID string, description string) error {
	uri := fmt.Sprintf("%s/%s", netPath, networkID)

	netReq := RequestBody{"description": description}
	req, err := n.client.NewRequest(ctx, http.MethodPut, uri, netReq)
	if err != nil {
		return err
	}

	return n.client.DoWithContext(ctx, req, nil)
}

// Delete a private network. Before deleting, a network must be disabled from all instances
func (n *NetworkServiceHandler) Delete(ctx context.Context, networkID string) error {
	uri := fmt.Sprintf("%s/%s", netPath, networkID)
	req, err := n.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return n.client.DoWithContext(ctx, req, nil)
}

// List lists all private networks on the current account
func (n *NetworkServiceHandler) List(ctx context.Context, options *ListOptions) ([]Network, *Meta, error) {
	req, err := n.client.NewRequest(ctx, http.MethodGet, netPath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	networks := new(networksBase)
	if err = n.client.DoWithContext(ctx, req, networks); err != nil {
		return nil, nil, err
	}

	return networks.Networks, networks.Meta, nil
}
