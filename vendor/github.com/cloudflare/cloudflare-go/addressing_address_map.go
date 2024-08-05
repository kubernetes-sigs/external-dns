package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// AddressMap contains information about an address map.
type AddressMap struct {
	ID           string                 `json:"id"`
	Description  *string                `json:"description,omitempty"`
	DefaultSNI   *string                `json:"default_sni"`
	Enabled      *bool                  `json:"enabled"`
	Deletable    *bool                  `json:"can_delete"`
	CanModifyIPs *bool                  `json:"can_modify_ips"`
	Memberships  []AddressMapMembership `json:"memberships"`
	IPs          []AddressMapIP         `json:"ips"`
	CreatedAt    time.Time              `json:"created_at"`
	ModifiedAt   time.Time              `json:"modified_at"`
}

type AddressMapIP struct {
	IP        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
}

type AddressMapMembershipContainer struct {
	Identifier string                   `json:"identifier"`
	Kind       AddressMapMembershipKind `json:"kind"`
}

type AddressMapMembership struct {
	Identifier string                   `json:"identifier"`
	Kind       AddressMapMembershipKind `json:"kind"`
	Deletable  *bool                    `json:"can_delete"`
	CreatedAt  time.Time                `json:"created_at"`
}

func (ammb *AddressMapMembershipContainer) URLFragment() string {
	switch ammb.Kind {
	case AddressMapMembershipAccount:
		return fmt.Sprintf("accounts/%s", ammb.Identifier)
	case AddressMapMembershipZone:
		return fmt.Sprintf("zones/%s", ammb.Identifier)
	default:
		return fmt.Sprintf("%s/%s", ammb.Kind, ammb.Identifier)
	}
}

type AddressMapMembershipKind string

const (
	AddressMapMembershipZone    AddressMapMembershipKind = "zone"
	AddressMapMembershipAccount AddressMapMembershipKind = "account"
)

// ListAddressMapResponse contains a slice of address maps.
type ListAddressMapResponse struct {
	Response
	Result []AddressMap `json:"result"`
}

// GetAddressMapResponse contains a specific address map's API Response.
type GetAddressMapResponse struct {
	Response
	Result AddressMap `json:"result"`
}

// CreateAddressMapParams contains information about an address map to be created.
type CreateAddressMapParams struct {
	Description *string                         `json:"description"`
	Enabled     *bool                           `json:"enabled"`
	IPs         []string                        `json:"ips"`
	Memberships []AddressMapMembershipContainer `json:"memberships"`
}

// UpdateAddressMapParams contains information about an address map to be updated.
type UpdateAddressMapParams struct {
	ID          string  `json:"-"`
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	DefaultSNI  *string `json:"default_sni,omitempty"`
}

// AddressMapFilter contains filter parameters for finding a list of address maps.
type ListAddressMapsParams struct {
	IP   *string `url:"ip,omitempty"`
	CIDR *string `url:"cidr,omitempty"`
}

// CreateIPAddressToAddressMapParams contains request parameters to add/remove IP address to/from an address map.
type CreateIPAddressToAddressMapParams struct {
	// ID represents the target address map for adding the IP address.
	ID string
	// The IP address.
	IP string
}

// CreateMembershipToAddressMapParams contains request parameters to add/remove membership from an address map.
type CreateMembershipToAddressMapParams struct {
	// ID represents the target address map for adding the membershp.
	ID         string
	Membership AddressMapMembershipContainer
}

type DeleteMembershipFromAddressMapParams struct {
	// ID represents the target address map for removing the IP address.
	ID         string
	Membership AddressMapMembershipContainer
}

type DeleteIPAddressFromAddressMapParams struct {
	// ID represents the target address map for adding the membershp.
	ID string
	IP string
}

// ListAddressMaps lists all address maps owned by the account.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-list-address-maps
func (api *API) ListAddressMaps(ctx context.Context, rc *ResourceContainer, params ListAddressMapsParams) ([]AddressMap, error) {
	if rc.Level != AccountRouteLevel {
		return []AddressMap{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := buildURI(fmt.Sprintf("/%s/addressing/address_maps", rc.URLFragment()), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AddressMap{}, err
	}

	result := ListAddressMapResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []AddressMap{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// CreateAddressMap creates a new address map under the account.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-create-address-map
func (api *API) CreateAddressMap(ctx context.Context, rc *ResourceContainer, params CreateAddressMapParams) (AddressMap, error) {
	if rc.Level != AccountRouteLevel {
		return AddressMap{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps", rc.URLFragment())
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return AddressMap{}, err
	}

	result := GetAddressMapResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return AddressMap{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// GetAddressMap returns a specific address map.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-address-map-details
func (api *API) GetAddressMap(ctx context.Context, rc *ResourceContainer, id string) (AddressMap, error) {
	if rc.Level != AccountRouteLevel {
		return AddressMap{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s", rc.URLFragment(), id)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AddressMap{}, err
	}

	result := GetAddressMapResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return AddressMap{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// UpdateAddressMap modifies properties of an address map owned by the account.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-update-address-map
func (api *API) UpdateAddressMap(ctx context.Context, rc *ResourceContainer, params UpdateAddressMapParams) (AddressMap, error) {
	if rc.Level != AccountRouteLevel {
		return AddressMap{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s", rc.URLFragment(), params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AddressMap{}, err
	}

	result := GetAddressMapResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return AddressMap{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// DeleteAddressMap deletes a particular address map owned by the account.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-delete-address-map
func (api *API) DeleteAddressMap(ctx context.Context, rc *ResourceContainer, id string) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s", rc.URLFragment(), id)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

// CreateIPAddressToAddressMap adds an IP address from a prefix owned by the account to a particular address map.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-add-an-ip-to-an-address-map
func (api *API) CreateIPAddressToAddressMap(ctx context.Context, rc *ResourceContainer, params CreateIPAddressToAddressMapParams) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s/ips/%s", rc.URLFragment(), params.ID, params.IP)
	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, nil)
	return err
}

// DeleteIPAddressFromAddressMap removes an IP address from a particular address map.
//
// API reference: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-remove-an-ip-from-an-address-map
func (api *API) DeleteIPAddressFromAddressMap(ctx context.Context, rc *ResourceContainer, params DeleteIPAddressFromAddressMapParams) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s/ips/%s", rc.URLFragment(), params.ID, params.IP)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

// CreateMembershipToAddressMap adds a zone/account as a member of a particular address map.
//
// API reference:
//   - account: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-add-an-account-membership-to-an-address-map
//   - zone: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-add-a-zone-membership-to-an-address-map
func (api *API) CreateMembershipToAddressMap(ctx context.Context, rc *ResourceContainer, params CreateMembershipToAddressMapParams) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	if params.Membership.Kind != AddressMapMembershipZone && params.Membership.Kind != AddressMapMembershipAccount {
		return fmt.Errorf("requested membershp kind (%q) is not supported", params.Membership.Kind)
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s/%s", rc.URLFragment(), params.ID, params.Membership.URLFragment())
	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, nil)
	return err
}

// DeleteMembershipFromAddressMap removes a zone/account as a member of a particular address map.
//
// API reference:
//   - account: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-remove-an-account-membership-from-an-address-map
//   - zone: https://developers.cloudflare.com/api/operations/ip-address-management-address-maps-remove-a-zone-membership-from-an-address-map
func (api *API) DeleteMembershipFromAddressMap(ctx context.Context, rc *ResourceContainer, params DeleteMembershipFromAddressMapParams) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	if params.Membership.Kind != AddressMapMembershipZone && params.Membership.Kind != AddressMapMembershipAccount {
		return fmt.Errorf("requested membershp kind (%q) is not supported", params.Membership.Kind)
	}

	uri := fmt.Sprintf("/%s/addressing/address_maps/%s/%s", rc.URLFragment(), params.ID, params.Membership.URLFragment())
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}
