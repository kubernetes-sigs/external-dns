// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ThreatEventCategoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventCategoryService] method instead.
type ThreatEventCategoryService struct {
	Options []option.RequestOption
}

// NewThreatEventCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventCategoryService(opts ...option.RequestOption) (r *ThreatEventCategoryService) {
	r = &ThreatEventCategoryService{}
	r.Options = opts
	return
}

// Creates a new category
func (r *ThreatEventCategoryService) New(ctx context.Context, params ThreatEventCategoryNewParams, opts ...option.RequestOption) (res *ThreatEventCategoryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/categories/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists categories
func (r *ThreatEventCategoryService) List(ctx context.Context, query ThreatEventCategoryListParams, opts ...option.RequestOption) (res *[]ThreatEventCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/categories", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a category
func (r *ThreatEventCategoryService) Delete(ctx context.Context, categoryID string, body ThreatEventCategoryDeleteParams, opts ...option.RequestOption) (res *ThreatEventCategoryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/categories/%s", body.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates a category
func (r *ThreatEventCategoryService) Edit(ctx context.Context, categoryID string, params ThreatEventCategoryEditParams, opts ...option.RequestOption) (res *ThreatEventCategoryEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/categories/%s", params.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Reads a category
func (r *ThreatEventCategoryService) Get(ctx context.Context, categoryID string, query ThreatEventCategoryGetParams, opts ...option.RequestOption) (res *ThreatEventCategoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/categories/%s", query.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventCategoryNewResponse struct {
	KillChain   float64                            `json:"killChain,required"`
	Name        string                             `json:"name,required"`
	UUID        string                             `json:"uuid,required"`
	MitreAttack []string                           `json:"mitreAttack"`
	Shortname   string                             `json:"shortname"`
	JSON        threatEventCategoryNewResponseJSON `json:"-"`
}

// threatEventCategoryNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventCategoryNewResponse]
type threatEventCategoryNewResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCategoryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCategoryNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCategoryListResponse struct {
	KillChain   float64                             `json:"killChain,required"`
	Name        string                              `json:"name,required"`
	UUID        string                              `json:"uuid,required"`
	MitreAttack []string                            `json:"mitreAttack"`
	Shortname   string                              `json:"shortname"`
	JSON        threatEventCategoryListResponseJSON `json:"-"`
}

// threatEventCategoryListResponseJSON contains the JSON metadata for the struct
// [ThreatEventCategoryListResponse]
type threatEventCategoryListResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCategoryDeleteResponse struct {
	UUID string                                `json:"uuid,required"`
	JSON threatEventCategoryDeleteResponseJSON `json:"-"`
}

// threatEventCategoryDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatEventCategoryDeleteResponse]
type threatEventCategoryDeleteResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCategoryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCategoryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCategoryEditResponse struct {
	KillChain   float64                             `json:"killChain,required"`
	Name        string                              `json:"name,required"`
	UUID        string                              `json:"uuid,required"`
	MitreAttack []string                            `json:"mitreAttack"`
	Shortname   string                              `json:"shortname"`
	JSON        threatEventCategoryEditResponseJSON `json:"-"`
}

// threatEventCategoryEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventCategoryEditResponse]
type threatEventCategoryEditResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCategoryEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCategoryEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCategoryGetResponse struct {
	KillChain   float64                            `json:"killChain,required"`
	Name        string                             `json:"name,required"`
	UUID        string                             `json:"uuid,required"`
	MitreAttack []string                           `json:"mitreAttack"`
	Shortname   string                             `json:"shortname"`
	JSON        threatEventCategoryGetResponseJSON `json:"-"`
}

// threatEventCategoryGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventCategoryGetResponse]
type threatEventCategoryGetResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	MitreAttack apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCategoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCategoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCategoryNewParams struct {
	// Account ID.
	AccountID   param.Field[string]   `path:"account_id,required"`
	KillChain   param.Field[float64]  `json:"killChain,required"`
	Name        param.Field[string]   `json:"name,required"`
	MitreAttack param.Field[[]string] `json:"mitreAttack"`
	Shortname   param.Field[string]   `json:"shortname"`
}

func (r ThreatEventCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventCategoryListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventCategoryDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ThreatEventCategoryEditParams struct {
	// Account ID.
	AccountID   param.Field[string]   `path:"account_id,required"`
	KillChain   param.Field[float64]  `json:"killChain"`
	MitreAttack param.Field[[]string] `json:"mitreAttack"`
	Name        param.Field[string]   `json:"name"`
	Shortname   param.Field[string]   `json:"shortname"`
}

func (r ThreatEventCategoryEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventCategoryGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
