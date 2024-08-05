package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type AccessTag struct {
	Name     string `json:"name,omitempty"`
	AppCount int    `json:"app_count,omitempty"`
}

type AccessTagListResponse struct {
	Response
	Result     []AccessTag `json:"result"`
	ResultInfo `json:"result_info"`
}

type AccessTagResponse struct {
	Response
	Result AccessTag `json:"result"`
}

type ListAccessTagsParams struct{}

type CreateAccessTagParams struct {
	Name string `json:"name,omitempty"`
}

func (api *API) ListAccessTags(ctx context.Context, rc *ResourceContainer, params ListAccessTagsParams) ([]AccessTag, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/access/tags", rc.Level, rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessTag{}, err
	}

	var TagsResponse AccessTagListResponse
	err = json.Unmarshal(res, &TagsResponse)
	if err != nil {
		return []AccessTag{}, err
	}
	return TagsResponse.Result, nil
}

func (api *API) GetAccessTag(ctx context.Context, rc *ResourceContainer, tagName string) (AccessTag, error) {
	uri := fmt.Sprintf("/%s/%s/access/tags/%s", rc.Level, rc.Identifier, tagName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessTag{}, err
	}

	var TagResponse AccessTagResponse
	err = json.Unmarshal(res, &TagResponse)
	if err != nil {
		return AccessTag{}, err
	}
	return TagResponse.Result, nil
}

func (api *API) CreateAccessTag(ctx context.Context, rc *ResourceContainer, params CreateAccessTagParams) (AccessTag, error) {
	uri := fmt.Sprintf("/%s/%s/access/tags", rc.Level, rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return AccessTag{}, err
	}

	var TagResponse AccessTagResponse
	err = json.Unmarshal(res, &TagResponse)
	if err != nil {
		return AccessTag{}, err
	}
	return TagResponse.Result, nil
}

func (api *API) DeleteAccessTag(ctx context.Context, rc *ResourceContainer, tagName string) error {
	uri := fmt.Sprintf("/%s/%s/access/tags/%s", rc.Level, rc.Identifier, tagName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}
