// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// IAMService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIAMService] method instead.
type IAMService struct {
	Options          []option.RequestOption
	PermissionGroups *PermissionGroupService
	ResourceGroups   *ResourceGroupService
	UserGroups       *UserGroupService
}

// NewIAMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIAMService(opts ...option.RequestOption) (r *IAMService) {
	r = &IAMService{}
	r.Options = opts
	r.PermissionGroups = NewPermissionGroupService(opts...)
	r.ResourceGroups = NewResourceGroupService(opts...)
	r.UserGroups = NewUserGroupService(opts...)
	return
}
