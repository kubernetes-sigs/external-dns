package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EmailRoutingDestinationAddress struct {
	Tag      string     `json:"tag,omitempty"`
	Email    string     `json:"email,omitempty"`
	Verified *time.Time `json:"verified,omitempty"`
	Created  *time.Time `json:"created,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
}

type ListEmailRoutingAddressParameters struct {
	ResultInfo
	Direction string `url:"direction,omitempty"`
	Verified  *bool  `url:"verified,omitempty"`
}

type ListEmailRoutingAddressResponse struct {
	Result     []EmailRoutingDestinationAddress `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

type CreateEmailRoutingAddressParameters struct {
	Email string `json:"email,omitempty"`
}

type CreateEmailRoutingAddressResponse struct {
	Result EmailRoutingDestinationAddress `json:"result,omitempty"`
	Response
}

// ListEmailRoutingDestinationAddresses Lists existing destination addresses.
//
// API reference: https://api.cloudflare.com/#email-routing-destination-addresses-list-destination-addresses
func (api *API) ListEmailRoutingDestinationAddresses(ctx context.Context, rc *ResourceContainer, params ListEmailRoutingAddressParameters) ([]EmailRoutingDestinationAddress, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []EmailRoutingDestinationAddress{}, &ResultInfo{}, ErrMissingAccountID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 50
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var addresses []EmailRoutingDestinationAddress
	var eResponse ListEmailRoutingAddressResponse
	for {
		uri := buildURI(fmt.Sprintf("/accounts/%s/email/routing/addresses", rc.Identifier), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []EmailRoutingDestinationAddress{}, &ResultInfo{}, err
		}
		err = json.Unmarshal(res, &eResponse)
		if err != nil {
			return []EmailRoutingDestinationAddress{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		addresses = append(addresses, eResponse.Result...)
		params.ResultInfo = eResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return addresses, &eResponse.ResultInfo, nil
}

// CreateEmailRoutingDestinationAddress Create a destination address to forward your emails to.
// Destination addresses need to be verified before they become active.
//
// API reference: https://api.cloudflare.com/#email-routing-destination-addresses-create-a-destination-address
func (api *API) CreateEmailRoutingDestinationAddress(ctx context.Context, rc *ResourceContainer, params CreateEmailRoutingAddressParameters) (EmailRoutingDestinationAddress, error) {
	if rc.Identifier == "" {
		return EmailRoutingDestinationAddress{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/email/routing/addresses", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return EmailRoutingDestinationAddress{}, ErrMissingAccountID
	}

	var r CreateEmailRoutingAddressResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingDestinationAddress{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// GetEmailRoutingDestinationAddress Gets information for a specific destination email already created.
//
// API reference: https://api.cloudflare.com/#email-routing-destination-addresses-get-a-destination-address
func (api *API) GetEmailRoutingDestinationAddress(ctx context.Context, rc *ResourceContainer, addressID string) (EmailRoutingDestinationAddress, error) {
	if rc.Identifier == "" {
		return EmailRoutingDestinationAddress{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/email/routing/addresses/%s", rc.Identifier, addressID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return EmailRoutingDestinationAddress{}, ErrMissingAccountID
	}

	var r CreateEmailRoutingAddressResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingDestinationAddress{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DeleteEmailRoutingDestinationAddress Deletes a specific destination address.
//
// API reference: https://api.cloudflare.com/#email-routing-destination-addresses-delete-destination-address
func (api *API) DeleteEmailRoutingDestinationAddress(ctx context.Context, rc *ResourceContainer, addressID string) (EmailRoutingDestinationAddress, error) {
	if rc.Identifier == "" {
		return EmailRoutingDestinationAddress{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/email/routing/addresses/%s", rc.Identifier, addressID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return EmailRoutingDestinationAddress{}, ErrMissingAccountID
	}

	var r CreateEmailRoutingAddressResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingDestinationAddress{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}
