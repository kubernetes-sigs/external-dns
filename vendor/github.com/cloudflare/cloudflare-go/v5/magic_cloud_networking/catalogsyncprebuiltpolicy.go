// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// CatalogSyncPrebuiltPolicyService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogSyncPrebuiltPolicyService] method instead.
type CatalogSyncPrebuiltPolicyService struct {
	Options []option.RequestOption
}

// NewCatalogSyncPrebuiltPolicyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCatalogSyncPrebuiltPolicyService(opts ...option.RequestOption) (r *CatalogSyncPrebuiltPolicyService) {
	r = &CatalogSyncPrebuiltPolicyService{}
	r.Options = opts
	return
}

// List prebuilt catalog sync policies (Closed Beta).
func (r *CatalogSyncPrebuiltPolicyService) List(ctx context.Context, params CatalogSyncPrebuiltPolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CatalogSyncPrebuiltPolicyListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/catalog-syncs/prebuilt-policies", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List prebuilt catalog sync policies (Closed Beta).
func (r *CatalogSyncPrebuiltPolicyService) ListAutoPaging(ctx context.Context, params CatalogSyncPrebuiltPolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CatalogSyncPrebuiltPolicyListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type CatalogSyncPrebuiltPolicyListResponse struct {
	ApplicableDestinations []CatalogSyncPrebuiltPolicyListResponseApplicableDestination `json:"applicable_destinations,required"`
	PolicyDescription      string                                                       `json:"policy_description,required"`
	PolicyName             string                                                       `json:"policy_name,required"`
	PolicyString           string                                                       `json:"policy_string,required"`
	JSON                   catalogSyncPrebuiltPolicyListResponseJSON                    `json:"-"`
}

// catalogSyncPrebuiltPolicyListResponseJSON contains the JSON metadata for the
// struct [CatalogSyncPrebuiltPolicyListResponse]
type catalogSyncPrebuiltPolicyListResponseJSON struct {
	ApplicableDestinations apijson.Field
	PolicyDescription      apijson.Field
	PolicyName             apijson.Field
	PolicyString           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CatalogSyncPrebuiltPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catalogSyncPrebuiltPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type CatalogSyncPrebuiltPolicyListResponseApplicableDestination string

const (
	CatalogSyncPrebuiltPolicyListResponseApplicableDestinationNone          CatalogSyncPrebuiltPolicyListResponseApplicableDestination = "NONE"
	CatalogSyncPrebuiltPolicyListResponseApplicableDestinationZeroTrustList CatalogSyncPrebuiltPolicyListResponseApplicableDestination = "ZERO_TRUST_LIST"
)

func (r CatalogSyncPrebuiltPolicyListResponseApplicableDestination) IsKnown() bool {
	switch r {
	case CatalogSyncPrebuiltPolicyListResponseApplicableDestinationNone, CatalogSyncPrebuiltPolicyListResponseApplicableDestinationZeroTrustList:
		return true
	}
	return false
}

type CatalogSyncPrebuiltPolicyListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Specify type of destination, omit to return all.
	DestinationType param.Field[CatalogSyncPrebuiltPolicyListParamsDestinationType] `query:"destination_type"`
}

// URLQuery serializes [CatalogSyncPrebuiltPolicyListParams]'s query parameters as
// `url.Values`.
func (r CatalogSyncPrebuiltPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specify type of destination, omit to return all.
type CatalogSyncPrebuiltPolicyListParamsDestinationType string

const (
	CatalogSyncPrebuiltPolicyListParamsDestinationTypeNone          CatalogSyncPrebuiltPolicyListParamsDestinationType = "NONE"
	CatalogSyncPrebuiltPolicyListParamsDestinationTypeZeroTrustList CatalogSyncPrebuiltPolicyListParamsDestinationType = "ZERO_TRUST_LIST"
)

func (r CatalogSyncPrebuiltPolicyListParamsDestinationType) IsKnown() bool {
	switch r {
	case CatalogSyncPrebuiltPolicyListParamsDestinationTypeNone, CatalogSyncPrebuiltPolicyListParamsDestinationTypeZeroTrustList:
		return true
	}
	return false
}
