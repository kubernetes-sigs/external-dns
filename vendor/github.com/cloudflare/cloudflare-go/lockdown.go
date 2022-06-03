package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ZoneLockdown represents a Zone Lockdown rule. A rule only permits access to
// the provided URL pattern(s) from the given IP address(es) or subnet(s).
type ZoneLockdown struct {
	ID             string               `json:"id"`
	Description    string               `json:"description"`
	URLs           []string             `json:"urls"`
	Configurations []ZoneLockdownConfig `json:"configurations"`
	Paused         bool                 `json:"paused"`
	Priority       int                  `json:"priority,omitempty"`
	CreatedOn      *time.Time           `json:"created_on,omitempty"`
	ModifiedOn     *time.Time           `json:"modified_on,omitempty"`
}

// ZoneLockdownConfig represents a Zone Lockdown config, which comprises
// a Target ("ip" or "ip_range") and a Value (an IP address or IP+mask,
// respectively.)
type ZoneLockdownConfig struct {
	Target string `json:"target"`
	Value  string `json:"value"`
}

// ZoneLockdownResponse represents a response from the Zone Lockdown endpoint.
type ZoneLockdownResponse struct {
	Result ZoneLockdown `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// ZoneLockdownListResponse represents a response from the List Zone Lockdown
// endpoint.
type ZoneLockdownListResponse struct {
	Result []ZoneLockdown `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// ZoneLockdownCreateParams contains required and optional params
// for creating a zone lockdown.
type ZoneLockdownCreateParams struct {
	Description    string               `json:"description"`
	URLs           []string             `json:"urls"`
	Configurations []ZoneLockdownConfig `json:"configurations"`
	Paused         bool                 `json:"paused"`
	Priority       int                  `json:"priority,omitempty"`
}

// ZoneLockdownUpdateParams contains required and optional params
// for updating a zone lockdown.
type ZoneLockdownUpdateParams struct {
	ID             string               `json:"id"`
	Description    string               `json:"description"`
	URLs           []string             `json:"urls"`
	Configurations []ZoneLockdownConfig `json:"configurations"`
	Paused         bool                 `json:"paused"`
	Priority       int                  `json:"priority,omitempty"`
}

type LockdownListParams struct {
	ResultInfo
}

// CreateZoneLockdown creates a Zone ZoneLockdown rule for the given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-create-a-ZoneLockdown-rule
func (api *API) CreateZoneLockdown(ctx context.Context, rc *ResourceContainer, params ZoneLockdownCreateParams) (ZoneLockdown, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return ZoneLockdown{}, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneLockdown{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// UpdateZoneLockdown updates a Zone ZoneLockdown rule (based on the ID) for the given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-update-ZoneLockdown-rule
func (api *API) UpdateZoneLockdown(ctx context.Context, rc *ResourceContainer, params ZoneLockdownUpdateParams) (ZoneLockdown, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return ZoneLockdown{}, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneLockdown{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// DeleteZoneLockdown deletes a Zone ZoneLockdown rule (based on the ID) for the given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-delete-ZoneLockdown-rule
func (api *API) DeleteZoneLockdown(ctx context.Context, rc *ResourceContainer, id string) (ZoneLockdown, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns/%s", rc.Identifier, id)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return ZoneLockdown{}, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneLockdown{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// ZoneLockdown retrieves a Zone ZoneLockdown rule (based on the ID) for the given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-ZoneLockdown-rule-details
func (api *API) ZoneLockdown(ctx context.Context, rc *ResourceContainer, id string) (ZoneLockdown, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns/%s", rc.Identifier, id)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneLockdown{}, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneLockdown{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// ListZoneLockdowns retrieves every Zone ZoneLockdown rules for a given zone ID.
//
// Automatically paginates all results unless `params.PerPage` and `params.Page`
// is set.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-list-ZoneLockdown-rules
func (api *API) ListZoneLockdowns(ctx context.Context, rc *ResourceContainer, params LockdownListParams) ([]ZoneLockdown, *ResultInfo, error) {
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

	var zoneLockdowns []ZoneLockdown
	var zResponse ZoneLockdownListResponse
	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/firewall/lockdowns", rc.Identifier), params)

<<<<<<< HEAD
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns?%s", zoneID, v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// ZoneLockdown represents a Zone Lockdown rule. A rule only permits access to
// the provided URL pattern(s) from the given IP address(es) or subnet(s).
type ZoneLockdown struct {
	ID             string               `json:"id"`
	Description    string               `json:"description"`
	URLs           []string             `json:"urls"`
	Configurations []ZoneLockdownConfig `json:"configurations"`
	Paused         bool                 `json:"paused"`
	Priority       int                  `json:"priority,omitempty"`
	CreatedOn      *time.Time           `json:"created_on,omitempty"`
	ModifiedOn     *time.Time           `json:"modified_on,omitempty"`
}

// ZoneLockdownConfig represents a Zone Lockdown config, which comprises
// a Target ("ip" or "ip_range") and a Value (an IP address or IP+mask,
// respectively.)
type ZoneLockdownConfig struct {
	Target string `json:"target"`
	Value  string `json:"value"`
}

// ZoneLockdownResponse represents a response from the Zone Lockdown endpoint.
type ZoneLockdownResponse struct {
	Result ZoneLockdown `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// ZoneLockdownListResponse represents a response from the List Zone Lockdown
// endpoint.
type ZoneLockdownListResponse struct {
	Result []ZoneLockdown `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// CreateZoneLockdown creates a Zone ZoneLockdown rule for the given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-create-a-ZoneLockdown-rule
func (api *API) CreateZoneLockdown(ctx context.Context, zoneID string, ld ZoneLockdown) (*ZoneLockdownResponse, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, ld)
	if err != nil {
		return nil, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response, nil
}

// UpdateZoneLockdown updates a Zone ZoneLockdown rule (based on the ID) for the
// given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-update-ZoneLockdown-rule
func (api *API) UpdateZoneLockdown(ctx context.Context, zoneID string, id string, ld ZoneLockdown) (*ZoneLockdownResponse, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns/%s", zoneID, id)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, ld)
	if err != nil {
		return nil, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response, nil
}

// DeleteZoneLockdown deletes a Zone ZoneLockdown rule (based on the ID) for the
// given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-delete-ZoneLockdown-rule
func (api *API) DeleteZoneLockdown(ctx context.Context, zoneID string, id string) (*ZoneLockdownResponse, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns/%s", zoneID, id)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response, nil
}

// ZoneLockdown retrieves a Zone ZoneLockdown rule (based on the ID) for the
// given zone ID.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-ZoneLockdown-rule-details
func (api *API) ZoneLockdown(ctx context.Context, zoneID string, id string) (*ZoneLockdownResponse, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns/%s", zoneID, id)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &ZoneLockdownResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response, nil
}

// ListZoneLockdowns retrieves a list of Zone ZoneLockdown rules for a given
// zone ID by page number.
//
// API reference: https://api.cloudflare.com/#zone-ZoneLockdown-list-ZoneLockdown-rules
func (api *API) ListZoneLockdowns(ctx context.Context, zoneID string, page int) (*ZoneLockdownListResponse, error) {
	v := url.Values{}
	if page <= 0 {
		page = 1
	}

	v.Set("page", strconv.Itoa(page))
	v.Set("per_page", strconv.Itoa(100))

	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns?%s", zoneID, v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return nil, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return nil, errors.Wrap(err, errMakeRequestError)
=======
		return nil, err
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	uri := fmt.Sprintf("/zones/%s/firewall/lockdowns?%s", zoneID, v.Encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
=======
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ZoneLockdown{}, &ResultInfo{}, err
		}

		err = json.Unmarshal(res, &zResponse)
		if err != nil {
			return []ZoneLockdown{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal filters JSON data: %w", err)
		}

		zoneLockdowns = append(zoneLockdowns, zResponse.Result...)
		params.ResultInfo = zResponse.ResultInfo.Next()

		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	}

	return zoneLockdowns, &zResponse.ResultInfo, nil
}
