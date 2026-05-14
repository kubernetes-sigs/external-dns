// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// BillingHistoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingHistoryService] method instead.
type BillingHistoryService struct {
	Options []option.RequestOption
}

// NewBillingHistoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingHistoryService(opts ...option.RequestOption) (r *BillingHistoryService) {
	r = &BillingHistoryService{}
	r.Options = opts
	return
}

// Accesses your billing history object.
//
// Deprecated: deprecated
func (r *BillingHistoryService) List(ctx context.Context, query BillingHistoryListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[BillingHistory], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/billing/history"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Accesses your billing history object.
//
// Deprecated: deprecated
func (r *BillingHistoryService) ListAutoPaging(ctx context.Context, query BillingHistoryListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[BillingHistory] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

type BillingHistory struct {
	// Billing item identifier tag.
	ID string `json:"id,required"`
	// The billing item action.
	Action string `json:"action,required"`
	// The amount associated with this billing item.
	Amount float64 `json:"amount,required"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency,required"`
	// The billing item description.
	Description string `json:"description,required"`
	// When the billing item was created.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The billing item type.
	Type string             `json:"type,required"`
	Zone BillingHistoryZone `json:"zone,required"`
	JSON billingHistoryJSON `json:"-"`
}

// billingHistoryJSON contains the JSON metadata for the struct [BillingHistory]
type billingHistoryJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryJSON) RawJSON() string {
	return r.raw
}

type BillingHistoryZone struct {
	Name string                 `json:"name"`
	JSON billingHistoryZoneJSON `json:"-"`
}

// billingHistoryZoneJSON contains the JSON metadata for the struct
// [BillingHistoryZone]
type billingHistoryZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryZoneJSON) RawJSON() string {
	return r.raw
}

type BillingHistoryListParams struct {
	// The billing item action.
	Action param.Field[string] `query:"action"`
	// When the billing item was created.
	OccurredAt param.Field[time.Time] `query:"occurred_at" format:"date-time"`
	// Field to order billing history by.
	Order param.Field[BillingHistoryListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The billing item type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [BillingHistoryListParams]'s query parameters as
// `url.Values`.
func (r BillingHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order billing history by.
type BillingHistoryListParamsOrder string

const (
	BillingHistoryListParamsOrderType       BillingHistoryListParamsOrder = "type"
	BillingHistoryListParamsOrderOccurredAt BillingHistoryListParamsOrder = "occurred_at"
	BillingHistoryListParamsOrderAction     BillingHistoryListParamsOrder = "action"
)

func (r BillingHistoryListParamsOrder) IsKnown() bool {
	switch r {
	case BillingHistoryListParamsOrderType, BillingHistoryListParamsOrderOccurredAt, BillingHistoryListParamsOrderAction:
		return true
	}
	return false
}
