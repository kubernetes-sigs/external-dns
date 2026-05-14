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

// ThreatEventIndicatorTypeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventIndicatorTypeService] method instead.
type ThreatEventIndicatorTypeService struct {
	Options []option.RequestOption
}

// NewThreatEventIndicatorTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventIndicatorTypeService(opts ...option.RequestOption) (r *ThreatEventIndicatorTypeService) {
	r = &ThreatEventIndicatorTypeService{}
	r.Options = opts
	return
}

// Lists all indicator types
func (r *ThreatEventIndicatorTypeService) List(ctx context.Context, query ThreatEventIndicatorTypeListParams, opts ...option.RequestOption) (res *ThreatEventIndicatorTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/indicatorTypes", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventIndicatorTypeListResponse struct {
	Items ThreatEventIndicatorTypeListResponseItems `json:"items,required"`
	Type  string                                    `json:"type,required"`
	JSON  threatEventIndicatorTypeListResponseJSON  `json:"-"`
}

// threatEventIndicatorTypeListResponseJSON contains the JSON metadata for the
// struct [ThreatEventIndicatorTypeListResponse]
type threatEventIndicatorTypeListResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorTypeListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorTypeListResponseItems struct {
	Type string                                        `json:"type,required"`
	JSON threatEventIndicatorTypeListResponseItemsJSON `json:"-"`
}

// threatEventIndicatorTypeListResponseItemsJSON contains the JSON metadata for the
// struct [ThreatEventIndicatorTypeListResponseItems]
type threatEventIndicatorTypeListResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorTypeListResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorTypeListResponseItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorTypeListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
