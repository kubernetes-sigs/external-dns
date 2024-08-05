package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

var ErrMissingListID = errors.New("required missing list ID")

// TeamsList represents a Teams List.
type TeamsList struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Description string          `json:"description,omitempty"`
	Items       []TeamsListItem `json:"items,omitempty"`
	Count       uint64          `json:"count,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
}

// TeamsListItem represents a single list item.
type TeamsListItem struct {
	Value     string     `json:"value"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// PatchTeamsList represents a patch request for appending/removing list items.
type PatchTeamsList struct {
	ID     string          `json:"id"`
	Append []TeamsListItem `json:"append"`
	Remove []string        `json:"remove"`
}

// TeamsListListResponse represents the response from the list
// teams lists endpoint.
type TeamsListListResponse struct {
	Result []TeamsList `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// TeamsListItemsListResponse represents the response from the list
// teams list items endpoint.
type TeamsListItemsListResponse struct {
	Result []TeamsListItem `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// TeamsListDetailResponse is the API response, containing a single
// teams list.
type TeamsListDetailResponse struct {
	Response
	Result TeamsList `json:"result"`
}

type TeamsListItemsParams struct {
	AccountID string `url:"-"`
	ListID    string `url:"-"`

	PaginationOptions
}

// TeamsLists returns all lists within an account.
//
// API reference: https://api.cloudflare.com/#teams-lists-list-teams-lists
func (api *API) TeamsLists(ctx context.Context, accountID string) ([]TeamsList, ResultInfo, error) {
	uri := fmt.Sprintf("/%s/%s/gateway/lists", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsList{}, ResultInfo{}, err
	}

	var teamsListListResponse TeamsListListResponse
	err = json.Unmarshal(res, &teamsListListResponse)
	if err != nil {
		return []TeamsList{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListListResponse.Result, teamsListListResponse.ResultInfo, nil
}

// TeamsList returns a single list based on the list ID.
//
// API reference: https://api.cloudflare.com/#teams-lists-teams-list-details
func (api *API) TeamsList(ctx context.Context, accountID, listID string) (TeamsList, error) {
	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		AccountRouteRoot,
		accountID,
		listID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// TeamsListItems returns all list items for a list.
//
// API reference: https://api.cloudflare.com/#teams-lists-teams-list-items
func (api *API) TeamsListItems(ctx context.Context, params TeamsListItemsParams) ([]TeamsListItem, ResultInfo, error) {
	if params.AccountID == "" {
		return []TeamsListItem{}, ResultInfo{}, ErrMissingAccountID
	}

	if params.ListID == "" {
		return []TeamsListItem{}, ResultInfo{}, ErrMissingListID
	}

	uri := buildURI(
		fmt.Sprintf("/%s/%s/gateway/lists/%s/items", AccountRouteRoot, params.AccountID, params.ListID),
		params,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsListItem{}, ResultInfo{}, err
	}

	var teamsListItemsListResponse TeamsListItemsListResponse
	err = json.Unmarshal(res, &teamsListItemsListResponse)
	if err != nil {
		return []TeamsListItem{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListItemsListResponse.Result, teamsListItemsListResponse.ResultInfo, nil
}

// CreateTeamsList creates a new teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-create-teams-list
func (api *API) CreateTeamsList(ctx context.Context, accountID string, teamsList TeamsList) (TeamsList, error) {
	uri := fmt.Sprintf("/%s/%s/gateway/lists", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, teamsList)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// UpdateTeamsList updates an existing teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-update-teams-list
func (api *API) UpdateTeamsList(ctx context.Context, accountID string, teamsList TeamsList) (TeamsList, error) {
	if teamsList.ID == "" {
		return TeamsList{}, fmt.Errorf("teams list ID cannot be empty")
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		AccountRouteRoot,
		accountID,
		teamsList.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, teamsList)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// PatchTeamsList updates the items in an existing teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-patch-teams-list
func (api *API) PatchTeamsList(ctx context.Context, accountID string, listPatch PatchTeamsList) (TeamsList, error) {
	if listPatch.ID == "" {
		return TeamsList{}, fmt.Errorf("teams list ID cannot be empty")
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		AccountRouteRoot,
		accountID,
		listPatch.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, listPatch)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// DeleteTeamsList deletes a teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-delete-teams-list
func (api *API) DeleteTeamsList(ctx context.Context, accountID, teamsListID string) error {
	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		AccountRouteRoot,
		accountID,
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var ErrMissingListID = errors.New("required missing list ID")

// TeamsList represents a Teams List.
type TeamsList struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Description string          `json:"description,omitempty"`
	Items       []TeamsListItem `json:"items,omitempty"`
	Count       uint64          `json:"count,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
}

// TeamsListItem represents a single list item.
type TeamsListItem struct {
	Value     string     `json:"value"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// PatchTeamsList represents a patch request for appending/removing list items.
type PatchTeamsList struct {
	ID     string          `json:"id"`
	Append []TeamsListItem `json:"append"`
	Remove []string        `json:"remove"`
}

// TeamsListListResponse represents the response from the list
// teams lists endpoint.
type TeamsListListResponse struct {
	Result []TeamsList `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// TeamsListItemsListResponse represents the response from the list
// teams list items endpoint.
type TeamsListItemsListResponse struct {
	Result []TeamsListItem `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// TeamsListDetailResponse is the API response, containing a single
// teams list.
type TeamsListDetailResponse struct {
	Response
	Result TeamsList `json:"result"`
}

type ListTeamsListItemsParams struct {
	ListID string `url:"-"`

	ResultInfo
}

type ListTeamListsParams struct{}

type CreateTeamsListParams struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Description string          `json:"description,omitempty"`
	Items       []TeamsListItem `json:"items,omitempty"`
	Count       uint64          `json:"count,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
}

type UpdateTeamsListParams struct {
	ID          string          `json:"id,omitempty"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Description string          `json:"description,omitempty"`
	Items       []TeamsListItem `json:"items,omitempty"`
	Count       uint64          `json:"count,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
}

type PatchTeamsListParams struct {
	ID     string          `json:"id"`
	Append []TeamsListItem `json:"append"`
	Remove []string        `json:"remove"`
}

// ListTeamsLists returns all lists within an account.
//
// API reference: https://api.cloudflare.com/#teams-lists-list-teams-lists
func (api *API) ListTeamsLists(ctx context.Context, rc *ResourceContainer, params ListTeamListsParams) ([]TeamsList, ResultInfo, error) {
	uri := fmt.Sprintf("/%s/%s/gateway/lists", AccountRouteRoot, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsList{}, ResultInfo{}, err
	}

	var teamsListListResponse TeamsListListResponse
	err = json.Unmarshal(res, &teamsListListResponse)
	if err != nil {
		return []TeamsList{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListListResponse.Result, teamsListListResponse.ResultInfo, nil
}

// GetTeamsList returns a single list based on the list ID.
//
// API reference: https://api.cloudflare.com/#teams-lists-teams-list-details
func (api *API) GetTeamsList(ctx context.Context, rc *ResourceContainer, listID string) (TeamsList, error) {
	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		rc.Level,
		rc.Identifier,
		listID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// ListTeamsListItems returns all list items for a list.
//
// API reference: https://api.cloudflare.com/#teams-lists-teams-list-items
func (api *API) ListTeamsListItems(ctx context.Context, rc *ResourceContainer, params ListTeamsListItemsParams) ([]TeamsListItem, ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []TeamsListItem{}, ResultInfo{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return []TeamsListItem{}, ResultInfo{}, ErrMissingAccountID
	}

	if params.ListID == "" {
		return []TeamsListItem{}, ResultInfo{}, ErrMissingListID
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

	var teamListItems []TeamsListItem
	var lResponse TeamsListItemsListResponse
	for {
		lResponse = TeamsListItemsListResponse{}
		uri := buildURI(
			fmt.Sprintf("/%s/%s/gateway/lists/%s/items", rc.Level, rc.Identifier, params.ListID),
			params,
		)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []TeamsListItem{}, ResultInfo{}, err
		}

		err = json.Unmarshal(res, &lResponse)
		if err != nil {
			return []TeamsListItem{}, ResultInfo{}, fmt.Errorf("failed to unmarshal teams list JSON data: %w", err)
		}

		teamListItems = append(teamListItems, lResponse.Result...)
		params.ResultInfo = lResponse.ResultInfo.Next()

		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return teamListItems, lResponse.ResultInfo, nil
}

// CreateTeamsList creates a new teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-create-teams-list
func (api *API) CreateTeamsList(ctx context.Context, rc *ResourceContainer, params CreateTeamsListParams) (TeamsList, error) {
	if rc.Level != AccountRouteLevel {
		return TeamsList{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return TeamsList{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/gateway/lists", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// UpdateTeamsList updates an existing teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-update-teams-list
func (api *API) UpdateTeamsList(ctx context.Context, rc *ResourceContainer, params UpdateTeamsListParams) (TeamsList, error) {
	if rc.Level != AccountRouteLevel {
		return TeamsList{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return TeamsList{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return TeamsList{}, fmt.Errorf("teams list ID cannot be empty")
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		rc.Level,
		rc.Identifier,
		params.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// PatchTeamsList updates the items in an existing teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-patch-teams-list
func (api *API) PatchTeamsList(ctx context.Context, rc *ResourceContainer, listPatch PatchTeamsListParams) (TeamsList, error) {
	if rc.Level != AccountRouteLevel {
		return TeamsList{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return TeamsList{}, ErrMissingAccountID
	}

	if listPatch.ID == "" {
		return TeamsList{}, fmt.Errorf("teams list ID cannot be empty")
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		AccountRouteRoot,
		rc.Identifier,
		listPatch.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, listPatch)
	if err != nil {
		return TeamsList{}, err
	}

	var teamsListDetailResponse TeamsListDetailResponse
	err = json.Unmarshal(res, &teamsListDetailResponse)
	if err != nil {
		return TeamsList{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsListDetailResponse.Result, nil
}

// DeleteTeamsList deletes a teams list.
//
// API reference: https://api.cloudflare.com/#teams-lists-delete-teams-list
func (api *API) DeleteTeamsList(ctx context.Context, rc *ResourceContainer, teamsListID string) error {
	if rc.Level != AccountRouteLevel {
		return fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/lists/%s",
		AccountRouteRoot,
		rc.Identifier,
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		teamsListID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
