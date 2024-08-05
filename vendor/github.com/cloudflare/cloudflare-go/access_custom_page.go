package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

var ErrMissingUID = errors.New("required UID missing")

type AccessCustomPageType string

const (
	Forbidden      AccessCustomPageType = "forbidden"
	IdentityDenied AccessCustomPageType = "identity_denied"
)

type AccessCustomPage struct {
	// The HTML content of the custom page.
	CustomHTML string               `json:"custom_html,omitempty"`
	Name       string               `json:"name,omitempty"`
	AppCount   int                  `json:"app_count,omitempty"`
	Type       AccessCustomPageType `json:"type,omitempty"`
	UID        string               `json:"uid,omitempty"`
}

type AccessCustomPageListResponse struct {
	Response
	Result     []AccessCustomPage `json:"result"`
	ResultInfo `json:"result_info"`
}

type AccessCustomPageResponse struct {
	Response
	Result AccessCustomPage `json:"result"`
}

type ListAccessCustomPagesParams struct{}

type CreateAccessCustomPageParams struct {
	CustomHTML string               `json:"custom_html,omitempty"`
	Name       string               `json:"name,omitempty"`
	Type       AccessCustomPageType `json:"type,omitempty"`
}

type UpdateAccessCustomPageParams struct {
	CustomHTML string               `json:"custom_html,omitempty"`
	Name       string               `json:"name,omitempty"`
	Type       AccessCustomPageType `json:"type,omitempty"`
	UID        string               `json:"uid,omitempty"`
}

func (api *API) ListAccessCustomPages(ctx context.Context, rc *ResourceContainer, params ListAccessCustomPagesParams) ([]AccessCustomPage, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/access/custom_pages", rc.Level, rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessCustomPage{}, err
	}

	var customPagesResponse AccessCustomPageListResponse
	err = json.Unmarshal(res, &customPagesResponse)
	if err != nil {
		return []AccessCustomPage{}, err
	}
	return customPagesResponse.Result, nil
}

func (api *API) GetAccessCustomPage(ctx context.Context, rc *ResourceContainer, id string) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages/%s", rc.Level, rc.Identifier, id)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessCustomPage{}, err
	}

	var customPageResponse AccessCustomPageResponse
	err = json.Unmarshal(res, &customPageResponse)
	if err != nil {
		return AccessCustomPage{}, err
	}
	return customPageResponse.Result, nil
}

func (api *API) CreateAccessCustomPage(ctx context.Context, rc *ResourceContainer, params CreateAccessCustomPageParams) (AccessCustomPage, error) {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages", rc.Level, rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return AccessCustomPage{}, err
	}

	var customPageResponse AccessCustomPageResponse
	err = json.Unmarshal(res, &customPageResponse)
	if err != nil {
		return AccessCustomPage{}, err
	}
	return customPageResponse.Result, nil
}

func (api *API) DeleteAccessCustomPage(ctx context.Context, rc *ResourceContainer, id string) error {
	uri := fmt.Sprintf("/%s/%s/access/custom_pages/%s", rc.Level, rc.Identifier, id)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) UpdateAccessCustomPage(ctx context.Context, rc *ResourceContainer, params UpdateAccessCustomPageParams) (AccessCustomPage, error) {
	if params.UID == "" {
		return AccessCustomPage{}, ErrMissingUID
	}

	uri := fmt.Sprintf("/%s/%s/access/custom_pages/%s", rc.Level, rc.Identifier, params.UID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return AccessCustomPage{}, err
	}

	var customPageResponse AccessCustomPageResponse
	err = json.Unmarshal(res, &customPageResponse)
	if err != nil {
		return AccessCustomPage{}, err
	}
	return customPageResponse.Result, nil
}
