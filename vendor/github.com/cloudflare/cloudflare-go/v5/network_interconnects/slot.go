// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_interconnects

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
)

// SlotService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSlotService] method instead.
type SlotService struct {
	Options []option.RequestOption
}

// NewSlotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSlotService(opts ...option.RequestOption) (r *SlotService) {
	r = &SlotService{}
	r.Options = opts
	return
}

// Retrieve a list of all slots matching the specified parameters
func (r *SlotService) List(ctx context.Context, params SlotListParams, opts ...option.RequestOption) (res *SlotListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/slots", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get information about the specified slot
func (r *SlotService) Get(ctx context.Context, slot string, query SlotGetParams, opts ...option.RequestOption) (res *SlotGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if slot == "" {
		err = errors.New("missing required slot parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/slots/%s", query.AccountID, slot)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SlotListResponse struct {
	Items []SlotListResponseItem `json:"items,required"`
	Next  int64                  `json:"next,nullable"`
	JSON  slotListResponseJSON   `json:"-"`
}

// slotListResponseJSON contains the JSON metadata for the struct
// [SlotListResponse]
type slotListResponseJSON struct {
	Items       apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slotListResponseJSON) RawJSON() string {
	return r.raw
}

type SlotListResponseItem struct {
	// Slot ID
	ID       string                        `json:"id,required" format:"uuid"`
	Facility SlotListResponseItemsFacility `json:"facility,required"`
	// Whether the slot is occupied or not
	Occupied bool   `json:"occupied,required"`
	Site     string `json:"site,required"`
	Speed    string `json:"speed,required"`
	// Customer account tag
	Account string                   `json:"account"`
	JSON    slotListResponseItemJSON `json:"-"`
}

// slotListResponseItemJSON contains the JSON metadata for the struct
// [SlotListResponseItem]
type slotListResponseItemJSON struct {
	ID          apijson.Field
	Facility    apijson.Field
	Occupied    apijson.Field
	Site        apijson.Field
	Speed       apijson.Field
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlotListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slotListResponseItemJSON) RawJSON() string {
	return r.raw
}

type SlotListResponseItemsFacility struct {
	Address []string                          `json:"address,required"`
	Name    string                            `json:"name,required"`
	JSON    slotListResponseItemsFacilityJSON `json:"-"`
}

// slotListResponseItemsFacilityJSON contains the JSON metadata for the struct
// [SlotListResponseItemsFacility]
type slotListResponseItemsFacilityJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlotListResponseItemsFacility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slotListResponseItemsFacilityJSON) RawJSON() string {
	return r.raw
}

type SlotGetResponse struct {
	// Slot ID
	ID       string                  `json:"id,required" format:"uuid"`
	Facility SlotGetResponseFacility `json:"facility,required"`
	// Whether the slot is occupied or not
	Occupied bool   `json:"occupied,required"`
	Site     string `json:"site,required"`
	Speed    string `json:"speed,required"`
	// Customer account tag
	Account string              `json:"account"`
	JSON    slotGetResponseJSON `json:"-"`
}

// slotGetResponseJSON contains the JSON metadata for the struct [SlotGetResponse]
type slotGetResponseJSON struct {
	ID          apijson.Field
	Facility    apijson.Field
	Occupied    apijson.Field
	Site        apijson.Field
	Speed       apijson.Field
	Account     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlotGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slotGetResponseJSON) RawJSON() string {
	return r.raw
}

type SlotGetResponseFacility struct {
	Address []string                    `json:"address,required"`
	Name    string                      `json:"name,required"`
	JSON    slotGetResponseFacilityJSON `json:"-"`
}

// slotGetResponseFacilityJSON contains the JSON metadata for the struct
// [SlotGetResponseFacility]
type slotGetResponseFacilityJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SlotGetResponseFacility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r slotGetResponseFacilityJSON) RawJSON() string {
	return r.raw
}

type SlotListParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
	// If specified, only show slots with the given text in their address field
	AddressContains param.Field[string] `query:"address_contains"`
	Cursor          param.Field[int64]  `query:"cursor"`
	Limit           param.Field[int64]  `query:"limit"`
	// If specified, only show slots with a specific occupied/unoccupied state
	Occupied param.Field[bool] `query:"occupied"`
	// If specified, only show slots located at the given site
	Site param.Field[string] `query:"site"`
	// If specified, only show slots that support the given speed
	Speed param.Field[string] `query:"speed"`
}

// URLQuery serializes [SlotListParams]'s query parameters as `url.Values`.
func (r SlotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SlotGetParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}
