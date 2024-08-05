package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type PermissionGroup struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Meta        map[string]string `json:"meta"`
	Permissions []Permission      `json:"permissions"`
}

type Permission struct {
	ID         string            `json:"id"`
	Key        string            `json:"key"`
	Attributes map[string]string `json:"attributes,omitempty"` // same as Meta in other structs
}

type PermissionGroupListResponse struct {
	Success  bool              `json:"success"`
	Errors   []string          `json:"errors"`
	Messages []string          `json:"messages"`
	Result   []PermissionGroup `json:"result"`
}

type PermissionGroupDetailResponse struct {
	Success  bool            `json:"success"`
	Errors   []string        `json:"errors"`
	Messages []string        `json:"messages"`
	Result   PermissionGroup `json:"result"`
}

type ListPermissionGroupParams struct {
	Depth    int    `url:"depth,omitempty"`
	RoleName string `url:"name,omitempty"`
}

const errMissingPermissionGroupID = "missing required permission group ID"

var ErrMissingPermissionGroupID = errors.New(errMissingPermissionGroupID)

// GetPermissionGroup returns a specific permission group from the API given
// the account ID and permission group ID.
func (api *API) GetPermissionGroup(ctx context.Context, rc *ResourceContainer, permissionGroupId string) (PermissionGroup, error) {
	if rc.Level != AccountRouteLevel {
		return PermissionGroup{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return PermissionGroup{}, ErrMissingAccountID
	}

	if permissionGroupId == "" {
		return PermissionGroup{}, ErrMissingPermissionGroupID
	}

	uri := fmt.Sprintf("/accounts/%s/iam/permission_groups/%s?depth=2", rc.Identifier, permissionGroupId)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PermissionGroup{}, err
	}

	var permissionGroupResponse PermissionGroupDetailResponse
	err = json.Unmarshal(res, &permissionGroupResponse)
	if err != nil {
		return PermissionGroup{}, err
	}

	return permissionGroupResponse.Result, nil
}

// ListPermissionGroups returns all valid permission groups for the provided
// parameters.
func (api *API) ListPermissionGroups(ctx context.Context, rc *ResourceContainer, params ListPermissionGroupParams) ([]PermissionGroup, error) {
	if rc.Level != AccountRouteLevel {
		return []PermissionGroup{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	params.Depth = 2
	uri := buildURI(fmt.Sprintf("/accounts/%s/iam/permission_groups", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []PermissionGroup{}, err
	}

	var permissionGroupResponse PermissionGroupListResponse
	err = json.Unmarshal(res, &permissionGroupResponse)
	if err != nil {
		return []PermissionGroup{}, err
	}

	return permissionGroupResponse.Result, nil
}
