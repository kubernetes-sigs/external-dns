package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AccessBookmark represents an Access bookmark application.
type AccessBookmark struct {
	ID                 string     `json:"id,omitempty"`
	Domain             string     `json:"domain"`
	Name               string     `json:"name"`
	LogoURL            string     `json:"logo_url,omitempty"`
	AppLauncherVisible *bool      `json:"app_launcher_visible,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// AccessBookmarkListResponse represents the response from the list
// access bookmarks endpoint.
type AccessBookmarkListResponse struct {
	Result []AccessBookmark `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccessBookmarkDetailResponse is the API response, containing a single
// access bookmark.
type AccessBookmarkDetailResponse struct {
	Response
	Result AccessBookmark `json:"result"`
}

// AccessBookmarks returns all bookmarks within an account.
//
// API reference: https://api.cloudflare.com/#access-bookmarks-list-access-bookmarks
func (api *API) AccessBookmarks(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]AccessBookmark, ResultInfo, error) {
	return api.accessBookmarks(ctx, accountID, pageOpts, AccountRouteRoot)
}

// ZoneLevelAccessBookmarks returns all bookmarks within a zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-bookmarks-list-access-bookmarks
func (api *API) ZoneLevelAccessBookmarks(ctx context.Context, zoneID string, pageOpts PaginationOptions) ([]AccessBookmark, ResultInfo, error) {
	return api.accessBookmarks(ctx, zoneID, pageOpts, ZoneRouteRoot)
}

func (api *API) accessBookmarks(ctx context.Context, id string, pageOpts PaginationOptions, routeRoot RouteRoot) ([]AccessBookmark, ResultInfo, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/access/bookmarks", routeRoot, id), pageOpts)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessBookmark{}, ResultInfo{}, err
	}

	var accessBookmarkListResponse AccessBookmarkListResponse
	err = json.Unmarshal(res, &accessBookmarkListResponse)
	if err != nil {
		return []AccessBookmark{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessBookmarkListResponse.Result, accessBookmarkListResponse.ResultInfo, nil
}

// AccessBookmark returns a single bookmark based on the
// bookmark ID.
//
// API reference: https://api.cloudflare.com/#access-bookmarks-access-bookmarks-details
func (api *API) AccessBookmark(ctx context.Context, accountID, bookmarkID string) (AccessBookmark, error) {
	return api.accessBookmark(ctx, accountID, bookmarkID, AccountRouteRoot)
}

// ZoneLevelAccessBookmark returns a single zone level bookmark based on the
// bookmark ID.
//
// API reference: https://api.cloudflare.com/#zone-level-access-bookmarks-access-bookmarks-details
func (api *API) ZoneLevelAccessBookmark(ctx context.Context, zoneID, bookmarkID string) (AccessBookmark, error) {
	return api.accessBookmark(ctx, zoneID, bookmarkID, ZoneRouteRoot)
}

func (api *API) accessBookmark(ctx context.Context, id, bookmarkID string, routeRoot RouteRoot) (AccessBookmark, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/bookmarks/%s",
		routeRoot,
		id,
		bookmarkID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessBookmark{}, err
	}

	var accessBookmarkDetailResponse AccessBookmarkDetailResponse
	err = json.Unmarshal(res, &accessBookmarkDetailResponse)
	if err != nil {
		return AccessBookmark{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessBookmarkDetailResponse.Result, nil
}

// CreateAccessBookmark creates a new access bookmark.
//
// API reference: https://api.cloudflare.com/#access-bookmarks-create-access-bookmark
func (api *API) CreateAccessBookmark(ctx context.Context, accountID string, accessBookmark AccessBookmark) (AccessBookmark, error) {
	return api.createAccessBookmark(ctx, accountID, accessBookmark, AccountRouteRoot)
}

// CreateZoneLevelAccessBookmark creates a new zone level access bookmark.
//
// API reference: https://api.cloudflare.com/#zone-level-access-bookmarks-create-access-bookmark
func (api *API) CreateZoneLevelAccessBookmark(ctx context.Context, zoneID string, accessBookmark AccessBookmark) (AccessBookmark, error) {
	return api.createAccessBookmark(ctx, zoneID, accessBookmark, ZoneRouteRoot)
}

func (api *API) createAccessBookmark(ctx context.Context, id string, accessBookmark AccessBookmark, routeRoot RouteRoot) (AccessBookmark, error) {
	uri := fmt.Sprintf("/%s/%s/access/bookmarks", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, accessBookmark)
	if err != nil {
		return AccessBookmark{}, err
	}

	var accessBookmarkDetailResponse AccessBookmarkDetailResponse
	err = json.Unmarshal(res, &accessBookmarkDetailResponse)
	if err != nil {
		return AccessBookmark{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessBookmarkDetailResponse.Result, nil
}

// UpdateAccessBookmark updates an existing access bookmark.
//
// API reference: https://api.cloudflare.com/#access-bookmarks-update-access-bookmark
func (api *API) UpdateAccessBookmark(ctx context.Context, accountID string, accessBookmark AccessBookmark) (AccessBookmark, error) {
	return api.updateAccessBookmark(ctx, accountID, accessBookmark, AccountRouteRoot)
}

// UpdateZoneLevelAccessBookmark updates an existing zone level access bookmark.
//
// API reference: https://api.cloudflare.com/#zone-level-access-bookmarks-update-access-bookmark
func (api *API) UpdateZoneLevelAccessBookmark(ctx context.Context, zoneID string, accessBookmark AccessBookmark) (AccessBookmark, error) {
	return api.updateAccessBookmark(ctx, zoneID, accessBookmark, ZoneRouteRoot)
}

func (api *API) updateAccessBookmark(ctx context.Context, id string, accessBookmark AccessBookmark, routeRoot RouteRoot) (AccessBookmark, error) {
	if accessBookmark.ID == "" {
		return AccessBookmark{}, fmt.Errorf("access bookmark ID cannot be empty")
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/bookmarks/%s",
		routeRoot,
		id,
		accessBookmark.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, accessBookmark)
	if err != nil {
		return AccessBookmark{}, err
	}

	var accessBookmarkDetailResponse AccessBookmarkDetailResponse
	err = json.Unmarshal(res, &accessBookmarkDetailResponse)
	if err != nil {
		return AccessBookmark{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessBookmarkDetailResponse.Result, nil
}

// DeleteAccessBookmark deletes an access bookmark.
//
// API reference: https://api.cloudflare.com/#access-bookmarks-delete-access-bookmark
func (api *API) DeleteAccessBookmark(ctx context.Context, accountID, bookmarkID string) error {
	return api.deleteAccessBookmark(ctx, accountID, bookmarkID, AccountRouteRoot)
}

// DeleteZoneLevelAccessBookmark deletes a zone level access bookmark.
//
// API reference: https://api.cloudflare.com/#zone-level-access-bookmarks-delete-access-bookmark
func (api *API) DeleteZoneLevelAccessBookmark(ctx context.Context, zoneID, bookmarkID string) error {
	return api.deleteAccessBookmark(ctx, zoneID, bookmarkID, ZoneRouteRoot)
}

func (api *API) deleteAccessBookmark(ctx context.Context, id, bookmarkID string, routeRoot RouteRoot) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/bookmarks/%s",
		routeRoot,
		id,
		bookmarkID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
