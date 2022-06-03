package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

// Magic Transit GRE Tunnel Error messages.
const (
	errMagicTransitGRETunnelNotModified = "When trying to modify GRE tunnel, API returned modified: false"
	errMagicTransitGRETunnelNotDeleted  = "When trying to delete GRE tunnel, API returned deleted: false"
)

// MagicTransitGRETunnel contains information about a GRE tunnel.
type MagicTransitGRETunnel struct {
	ID                    string                            `json:"id,omitempty"`
	CreatedOn             *time.Time                        `json:"created_on,omitempty"`
	ModifiedOn            *time.Time                        `json:"modified_on,omitempty"`
	Name                  string                            `json:"name"`
	CustomerGREEndpoint   string                            `json:"customer_gre_endpoint"`
	CloudflareGREEndpoint string                            `json:"cloudflare_gre_endpoint"`
	InterfaceAddress      string                            `json:"interface_address"`
	Description           string                            `json:"description,omitempty"`
	TTL                   uint8                             `json:"ttl,omitempty"`
	MTU                   uint16                            `json:"mtu,omitempty"`
	HealthCheck           *MagicTransitGRETunnelHealthcheck `json:"health_check,omitempty"`
}

// MagicTransitGRETunnelHealthcheck contains information about a GRE tunnel health check.
type MagicTransitGRETunnelHealthcheck struct {
	Enabled bool   `json:"enabled"`
	Target  string `json:"target,omitempty"`
	Type    string `json:"type,omitempty"`
}

// ListMagicTransitGRETunnelsResponse contains a response including GRE tunnels.
type ListMagicTransitGRETunnelsResponse struct {
	Response
	Result struct {
		GRETunnels []MagicTransitGRETunnel `json:"gre_tunnels"`
	} `json:"result"`
}

// GetMagicTransitGRETunnelResponse contains a response including zero or one GRE tunnels.
type GetMagicTransitGRETunnelResponse struct {
	Response
	Result struct {
		GRETunnel MagicTransitGRETunnel `json:"gre_tunnel"`
	} `json:"result"`
}

// CreateMagicTransitGRETunnelsRequest is an array of GRE tunnels to create.
type CreateMagicTransitGRETunnelsRequest struct {
	GRETunnels []MagicTransitGRETunnel `json:"gre_tunnels"`
}

// UpdateMagicTransitGRETunnelResponse contains a response after updating a GRE Tunnel.
type UpdateMagicTransitGRETunnelResponse struct {
	Response
	Result struct {
		Modified          bool                  `json:"modified"`
		ModifiedGRETunnel MagicTransitGRETunnel `json:"modified_gre_tunnel"`
	} `json:"result"`
}

// DeleteMagicTransitGRETunnelResponse contains a response after deleting a GRE Tunnel.
type DeleteMagicTransitGRETunnelResponse struct {
	Response
	Result struct {
		Deleted          bool                  `json:"deleted"`
		DeletedGRETunnel MagicTransitGRETunnel `json:"deleted_gre_tunnel"`
	} `json:"result"`
}

// ListMagicTransitGRETunnels lists all GRE tunnels for a given account.
//
// API reference: https://api.cloudflare.com/#magic-gre-tunnels-list-gre-tunnels
func (api *API) ListMagicTransitGRETunnels(ctx context.Context, accountID string) ([]MagicTransitGRETunnel, error) {
	uri := fmt.Sprintf("/accounts/%s/magic/gre_tunnels", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MagicTransitGRETunnel{}, err
	}

	result := ListMagicTransitGRETunnelsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []MagicTransitGRETunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result.GRETunnels, nil
}

// GetMagicTransitGRETunnel returns zero or one GRE tunnel.
//
// API reference: https://api.cloudflare.com/#magic-gre-tunnels-gre-tunnel-details
func (api *API) GetMagicTransitGRETunnel(ctx context.Context, accountID string, id string) (MagicTransitGRETunnel, error) {
	uri := fmt.Sprintf("/accounts/%s/magic/gre_tunnels/%s", accountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return MagicTransitGRETunnel{}, err
	}

	result := GetMagicTransitGRETunnelResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicTransitGRETunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result.GRETunnel, nil
}

// CreateMagicTransitGRETunnels creates one or more GRE tunnels.
//
// API reference: https://api.cloudflare.com/#magic-gre-tunnels-create-gre-tunnels
func (api *API) CreateMagicTransitGRETunnels(ctx context.Context, accountID string, tunnels []MagicTransitGRETunnel) ([]MagicTransitGRETunnel, error) {
	uri := fmt.Sprintf("/accounts/%s/magic/gre_tunnels", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, CreateMagicTransitGRETunnelsRequest{
		GRETunnels: tunnels,
	})

	if err != nil {
		return []MagicTransitGRETunnel{}, err
	}

	result := ListMagicTransitGRETunnelsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []MagicTransitGRETunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result.GRETunnels, nil
}

// UpdateMagicTransitGRETunnel updates a GRE tunnel.
//
// API reference: https://api.cloudflare.com/#magic-gre-tunnels-update-gre-tunnel
func (api *API) UpdateMagicTransitGRETunnel(ctx context.Context, accountID string, id string, tunnel MagicTransitGRETunnel) (MagicTransitGRETunnel, error) {
	uri := fmt.Sprintf("/accounts/%s/magic/gre_tunnels/%s", accountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, tunnel)

	if err != nil {
		return MagicTransitGRETunnel{}, err
	}

	result := UpdateMagicTransitGRETunnelResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicTransitGRETunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if !result.Result.Modified {
		return MagicTransitGRETunnel{}, errors.New(errMagicTransitGRETunnelNotModified)
	}

	return result.Result.ModifiedGRETunnel, nil
}

// DeleteMagicTransitGRETunnel deletes a GRE tunnel.
//
// API reference: https://api.cloudflare.com/#magic-gre-tunnels-delete-gre-tunnel
func (api *API) DeleteMagicTransitGRETunnel(ctx context.Context, accountID string, id string) (MagicTransitGRETunnel, error) {
	uri := fmt.Sprintf("/accounts/%s/magic/gre_tunnels/%s", accountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return MagicTransitGRETunnel{}, err
	}

	result := DeleteMagicTransitGRETunnelResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicTransitGRETunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if !result.Result.Deleted {
		return MagicTransitGRETunnel{}, errors.New(errMagicTransitGRETunnelNotDeleted)
	}

	return result.Result.DeletedGRETunnel, nil
}
