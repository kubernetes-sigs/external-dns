package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

const (
	// ListTypeIP specifies a list containing IP addresses.
	ListTypeIP = "ip"
	// ListTypeRedirect specifies a list containing redirects.
	ListTypeRedirect = "redirect"
)

// ListBulkOperation contains information about a Bulk Operation.
type ListBulkOperation struct {
	ID        string     `json:"id"`
	Status    string     `json:"status"`
	Error     string     `json:"error"`
	Completed *time.Time `json:"completed"`
}

// List contains information about a List.
type List struct {
	ID                    string     `json:"id"`
	Name                  string     `json:"name"`
	Description           string     `json:"description"`
	Kind                  string     `json:"kind"`
	NumItems              int        `json:"num_items"`
	NumReferencingFilters int        `json:"num_referencing_filters"`
	CreatedOn             *time.Time `json:"created_on"`
	ModifiedOn            *time.Time `json:"modified_on"`
}

// Redirect represents a redirect item in a List.
type Redirect struct {
	SourceUrl           string `json:"source_url"`
	IncludeSubdomains   *bool  `json:"include_subdomains,omitempty"`
	TargetUrl           string `json:"target_url"`
	StatusCode          *int   `json:"status_code,omitempty"`
	PreserveQueryString *bool  `json:"preserve_query_string,omitempty"`
	SubpathMatching     *bool  `json:"subpath_matching,omitempty"`
	PreservePathSuffix  *bool  `json:"preserve_path_suffix,omitempty"`
}

// ListItem contains information about a single List Item.
type ListItem struct {
	ID         string     `json:"id"`
	IP         *string    `json:"ip,omitempty"`
	Redirect   *Redirect  `json:"redirect,omitempty"`
	Comment    string     `json:"comment"`
	CreatedOn  *time.Time `json:"created_on"`
	ModifiedOn *time.Time `json:"modified_on"`
}

// ListCreateRequest contains data for a new List.
type ListCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
}

// ListItemCreateRequest contains data for a new List Item.
type ListItemCreateRequest struct {
	IP       *string   `json:"ip,omitempty"`
	Redirect *Redirect `json:"redirect,omitempty"`
	Comment  string    `json:"comment"`
}

// ListItemDeleteRequest wraps List Items that shall be deleted.
type ListItemDeleteRequest struct {
	Items []ListItemDeleteItemRequest `json:"items"`
}

// ListItemDeleteItemRequest contains single List Items that shall be deleted.
type ListItemDeleteItemRequest struct {
	ID string `json:"id"`
}

// ListUpdateRequest contains data for a List update.
type ListUpdateRequest struct {
	Description string `json:"description"`
}

// ListResponse contains a single List.
type ListResponse struct {
	Response
	Result List `json:"result"`
}

// ListItemCreateResponse contains information about the creation of a List Item.
type ListItemCreateResponse struct {
	Response
	Result struct {
		OperationID string `json:"operation_id"`
	} `json:"result"`
}

// ListListResponse contains a slice of Lists.
type ListListResponse struct {
	Response
	Result []List `json:"result"`
}

// ListBulkOperationResponse contains information about a Bulk Operation.
type ListBulkOperationResponse struct {
	Response
	Result ListBulkOperation `json:"result"`
}

// ListDeleteResponse contains information about the deletion of a List.
type ListDeleteResponse struct {
	Response
	Result struct {
		ID string `json:"id"`
	} `json:"result"`
}

// ListItemsListResponse contains information about List Items.
type ListItemsListResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []ListItem `json:"result"`
}

// ListItemDeleteResponse contains information about the deletion of a List Item.
type ListItemDeleteResponse struct {
	Response
	Result struct {
		OperationID string `json:"operation_id"`
	} `json:"result"`
}

// ListItemsGetResponse contains information about a single List Item.
type ListItemsGetResponse struct {
	Response
	Result ListItem `json:"result"`
}

type ListListsParams struct {
}

type ListCreateParams struct {
	Name        string
	Description string
	Kind        string
}

type ListGetParams struct {
	ID string
}

type ListUpdateParams struct {
	ID          string
	Description string
}

type ListDeleteParams struct {
	ID string
}

type ListListItemsParams struct {
	ID string
}

type ListCreateItemsParams struct {
	ID    string
	Items []ListItemCreateRequest
}

type ListCreateItemParams struct {
	ID   string
	Item ListItemCreateRequest
}

type ListReplaceItemsParams struct {
	ID    string
	Items []ListItemCreateRequest
}

type ListDeleteItemsParams struct {
	ID    string
	Items ListItemDeleteRequest
}

type ListGetItemParams struct {
	ListID string
	ID     string
}

type ListGetBulkOperationParams struct {
	ID string
}

// ListLists lists all Lists.
//
// API reference: https://api.cloudflare.com/#rules-lists-list-lists
func (api *API) ListLists(ctx context.Context, rc *ResourceContainer, params ListListsParams) ([]List, error) {
	if rc.Identifier == "" {
		return []List{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []List{}, err
	}

	result := ListListResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []List{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// CreateList creates a new List.
//
// API reference: https://api.cloudflare.com/#rules-lists-create-list
func (api *API) CreateList(ctx context.Context, rc *ResourceContainer, params ListCreateParams) (List, error) {
	if rc.Identifier == "" {
		return List{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, ListCreateRequest{Name: params.Name, Description: params.Description, Kind: params.Kind})
	if err != nil {
		return List{}, err
	}

	result := ListResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return List{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// GetList returns a single List.
//
// API reference: https://api.cloudflare.com/#rules-lists-get-list
func (api *API) GetList(ctx context.Context, rc *ResourceContainer, listID string) (List, error) {
	if rc.Identifier == "" {
		return List{}, ErrMissingAccountID
	}

	if listID == "" {
		return List{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s", rc.Identifier, listID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return List{}, err
	}

	result := ListResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return List{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// UpdateList updates the description of an existing List.
//
// API reference: https://api.cloudflare.com/#rules-lists-update-list
func (api *API) UpdateList(ctx context.Context, rc *ResourceContainer, params ListUpdateParams) (List, error) {
	if rc.Identifier == "" {
		return List{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return List{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, ListUpdateRequest{Description: params.Description})
	if err != nil {
		return List{}, err
	}

	result := ListResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return List{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// DeleteList deletes a List.
//
// API reference: https://api.cloudflare.com/#rules-lists-delete-list
func (api *API) DeleteList(ctx context.Context, rc *ResourceContainer, listID string) (ListDeleteResponse, error) {
	if rc.Identifier == "" {
		return ListDeleteResponse{}, ErrMissingAccountID
	}

	if listID == "" {
		return ListDeleteResponse{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s", rc.Identifier, listID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return ListDeleteResponse{}, err
	}

	result := ListDeleteResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListDeleteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, nil
}

// ListListItems returns a list with all items in a List.
//
// API reference: https://api.cloudflare.com/#rules-lists-list-list-items
func (api *API) ListListItems(ctx context.Context, rc *ResourceContainer, params ListListItemsParams) ([]ListItem, error) {
	var list []ListItem
	var cursor string
	var cursorQuery string

	for {
		if len(cursor) > 0 {
			cursorQuery = fmt.Sprintf("?cursor=%s", cursor)
		}
		uri := fmt.Sprintf("/accounts/%s/rules/lists/%s/items%s", rc.Identifier, params.ID, cursorQuery)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ListItem{}, err
		}

		result := ListItemsListResponse{}
		if err := json.Unmarshal(res, &result); err != nil {
			return []ListItem{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}

		list = append(list, result.Result...)
		if cursor = result.ResultInfo.Cursors.After; cursor == "" {
			break
		}
	}

	return list, nil
}

// CreateListItemAsync creates a new List Item asynchronously. Users have to poll the operation status by
// using the operation_id returned by this function.
//
// API reference: https://api.cloudflare.com/#rules-lists-create-list-items
func (api *API) CreateListItemAsync(ctx context.Context, rc *ResourceContainer, params ListCreateItemParams) (ListItemCreateResponse, error) {
	if rc.Identifier == "" {
		return ListItemCreateResponse{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return ListItemCreateResponse{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s/items", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, []ListItemCreateRequest{params.Item})
	if err != nil {
		return ListItemCreateResponse{}, err
	}

	result := ListItemCreateResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListItemCreateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, nil
}

// CreateListItem creates a new List Item synchronously and returns the current set of List Items.
func (api *API) CreateListItem(ctx context.Context, rc *ResourceContainer, params ListCreateItemParams) ([]ListItem, error) {
	result, err := api.CreateListItemAsync(ctx, rc, params)

	if err != nil {
		return []ListItem{}, err
	}

	err = api.pollListBulkOperation(ctx, rc, result.Result.OperationID)
	if err != nil {
		return []ListItem{}, err
	}

	return api.ListListItems(ctx, rc, ListListItemsParams{ID: params.ID})
}

// CreateListItemsAsync bulk creates multiple List Items asynchronously. Users
// have to poll the operation status by using the operation_id returned by this
// function.
//
// API reference: https://api.cloudflare.com/#rules-lists-create-list-items
func (api *API) CreateListItemsAsync(ctx context.Context, rc *ResourceContainer, params ListCreateItemsParams) (ListItemCreateResponse, error) {
	if rc.Identifier == "" {
		return ListItemCreateResponse{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return ListItemCreateResponse{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s/items", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Items)
	if err != nil {
		return ListItemCreateResponse{}, err
	}

	result := ListItemCreateResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListItemCreateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, nil
}

// CreateListItems bulk creates multiple List Items synchronously and returns
// the current set of List Items.
func (api *API) CreateListItems(ctx context.Context, rc *ResourceContainer, params ListCreateItemsParams) ([]ListItem, error) {
	result, err := api.CreateListItemsAsync(ctx, rc, params)
	if err != nil {
		return []ListItem{}, err
	}

	err = api.pollListBulkOperation(ctx, rc, result.Result.OperationID)
	if err != nil {
		return []ListItem{}, err
	}

	return api.ListListItems(ctx, rc, ListListItemsParams{ID: params.ID})
}

// ReplaceListItemsAsync replaces all List Items asynchronously. Users have to
// poll the operation status by using the operation_id returned by this
// function.
//
// API reference: https://api.cloudflare.com/#rules-lists-replace-list-items
func (api *API) ReplaceListItemsAsync(ctx context.Context, rc *ResourceContainer, params ListReplaceItemsParams) (ListItemCreateResponse, error) {
	if rc.Identifier == "" {
		return ListItemCreateResponse{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return ListItemCreateResponse{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s/items", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.Items)
	if err != nil {
		return ListItemCreateResponse{}, err
	}

	result := ListItemCreateResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListItemCreateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, nil
}

// ReplaceListItems replaces all List Items synchronously and returns the
// current set of List Items.
func (api *API) ReplaceListItems(ctx context.Context, rc *ResourceContainer, params ListReplaceItemsParams) (
	[]ListItem, error) {
	result, err := api.ReplaceListItemsAsync(ctx, rc, params)
	if err != nil {
		return []ListItem{}, err
	}

	err = api.pollListBulkOperation(ctx, rc, result.Result.OperationID)
	if err != nil {
		return []ListItem{}, err
	}

	return api.ListListItems(ctx, rc, ListListItemsParams{ID: params.ID})
}

// DeleteListItemsAsync removes specific Items of a List by their ID
// asynchronously. Users have to poll the operation status by using the
// operation_id returned by this function.
//
// API reference: https://api.cloudflare.com/#rules-lists-delete-list-items
func (api *API) DeleteListItemsAsync(ctx context.Context, rc *ResourceContainer, params ListDeleteItemsParams) (ListItemDeleteResponse, error) {
	if rc.Identifier == "" {
		return ListItemDeleteResponse{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return ListItemDeleteResponse{}, ErrMissingListID
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s/items", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, params.Items)
	if err != nil {
		return ListItemDeleteResponse{}, err
	}

	result := ListItemDeleteResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListItemDeleteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, nil
}

// DeleteListItems removes specific Items of a List by their ID synchronously
// and returns the current set of List Items.
func (api *API) DeleteListItems(ctx context.Context, rc *ResourceContainer, params ListDeleteItemsParams) ([]ListItem, error) {
	result, err := api.DeleteListItemsAsync(ctx, rc, params)
	if err != nil {
		return []ListItem{}, err
	}

	err = api.pollListBulkOperation(ctx, rc, result.Result.OperationID)
	if err != nil {
		return []ListItem{}, err
	}

	return api.ListListItems(ctx, AccountIdentifier(rc.Identifier), ListListItemsParams{ID: params.ID})
}

// GetListItem returns a single List Item.
//
// API reference: https://api.cloudflare.com/#rules-lists-get-list-item
func (api *API) GetListItem(ctx context.Context, rc *ResourceContainer, listID, itemID string) (ListItem, error) {
	if rc.Identifier == "" {
		return ListItem{}, ErrMissingAccountID
	}

	if listID == "" {
		return ListItem{}, ErrMissingListID
	}

	if itemID == "" {
		return ListItem{}, ErrMissingResourceIdentifier
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/%s/items/%s", rc.Identifier, listID, itemID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListItem{}, err
	}

	result := ListItemsGetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListItem{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// GetListBulkOperation returns the status of a bulk operation.
//
// API reference: https://api.cloudflare.com/#rules-lists-get-bulk-operation
func (api *API) GetListBulkOperation(ctx context.Context, rc *ResourceContainer, ID string) (ListBulkOperation, error) {
	if rc.Identifier == "" {
		return ListBulkOperation{}, ErrMissingAccountID
	}

	if ID == "" {
		return ListBulkOperation{}, ErrMissingResourceIdentifier
	}

	uri := fmt.Sprintf("/accounts/%s/rules/lists/bulk_operations/%s", rc.Identifier, ID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ListBulkOperation{}, err
	}

	result := ListBulkOperationResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ListBulkOperation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// pollListBulkOperation implements synchronous behaviour for some asynchronous
// endpoints. bulk-operation status can be either pending, running, failed or
// completed.
func (api *API) pollListBulkOperation(ctx context.Context, rc *ResourceContainer, ID string) error {
	for i := uint8(0); i < 16; i++ {
		sleepDuration := 1 << (i / 2) * time.Second
		select {
		case <-time.After(sleepDuration):
		case <-ctx.Done():
			return fmt.Errorf("operation aborted during backoff: %w", ctx.Err())
		}

		bulkResult, err := api.GetListBulkOperation(ctx, rc, ID)
		if err != nil {
			return err
		}

		switch bulkResult.Status {
		case "failed":
			return errors.New(bulkResult.Error)
		case "pending", "running":
			continue
		case "completed":
			return nil
		default:
			return fmt.Errorf("%s: %s", errOperationUnexpectedStatus, bulkResult.Status)
		}
	}

	return errors.New(errOperationStillRunning)
}
