// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_interconnects

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// InterconnectService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInterconnectService] method instead.
type InterconnectService struct {
	Options []option.RequestOption
}

// NewInterconnectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInterconnectService(opts ...option.RequestOption) (r *InterconnectService) {
	r = &InterconnectService{}
	r.Options = opts
	return
}

// Create a new interconnect
func (r *InterconnectService) New(ctx context.Context, params InterconnectNewParams, opts ...option.RequestOption) (res *InterconnectNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List existing interconnects
func (r *InterconnectService) List(ctx context.Context, params InterconnectListParams, opts ...option.RequestOption) (res *InterconnectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete an interconnect object
func (r *InterconnectService) Delete(ctx context.Context, icon string, body InterconnectDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s", body.AccountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about an interconnect object
func (r *InterconnectService) Get(ctx context.Context, icon string, query InterconnectGetParams, opts ...option.RequestOption) (res *InterconnectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s", query.AccountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate the Letter of Authorization (LOA) for a given interconnect
func (r *InterconnectService) LOA(ctx context.Context, icon string, query InterconnectLOAParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s/loa", query.AccountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Get the current status of an interconnect object
func (r *InterconnectService) Status(ctx context.Context, icon string, query InterconnectStatusParams, opts ...option.RequestOption) (res *InterconnectStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if icon == "" {
		err = errors.New("missing required icon parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/interconnects/%s/status", query.AccountID, icon)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InterconnectNewResponse struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Type    string `json:"type,required"`
	// This field can have the runtime type of
	// [InterconnectNewResponseNscInterconnectPhysicalBodyFacility].
	Facility interface{} `json:"facility"`
	Owner    string      `json:"owner"`
	Region   string      `json:"region"`
	// A Cloudflare site name.
	Site   string                      `json:"site"`
	SlotID string                      `json:"slot_id" format:"uuid"`
	Speed  string                      `json:"speed"`
	JSON   interconnectNewResponseJSON `json:"-"`
	union  InterconnectNewResponseUnion
}

// interconnectNewResponseJSON contains the JSON metadata for the struct
// [InterconnectNewResponse]
type interconnectNewResponseJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Facility    apijson.Field
	Owner       apijson.Field
	Region      apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r interconnectNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *InterconnectNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = InterconnectNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InterconnectNewResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [InterconnectNewResponseNscInterconnectPhysicalBody],
// [InterconnectNewResponseNscInterconnectGcpPartnerBody].
func (r InterconnectNewResponse) AsUnion() InterconnectNewResponseUnion {
	return r.union
}

// Union satisfied by [InterconnectNewResponseNscInterconnectPhysicalBody] or
// [InterconnectNewResponseNscInterconnectGcpPartnerBody].
type InterconnectNewResponseUnion interface {
	implementsInterconnectNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InterconnectNewResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InterconnectNewResponseNscInterconnectPhysicalBody{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InterconnectNewResponseNscInterconnectGcpPartnerBody{}),
		},
	)
}

type InterconnectNewResponseNscInterconnectPhysicalBody struct {
	Account  string                                                     `json:"account,required"`
	Facility InterconnectNewResponseNscInterconnectPhysicalBodyFacility `json:"facility,required"`
	Name     string                                                     `json:"name,required"`
	// A Cloudflare site name.
	Site   string                                                 `json:"site,required"`
	SlotID string                                                 `json:"slot_id,required" format:"uuid"`
	Speed  string                                                 `json:"speed,required"`
	Type   string                                                 `json:"type,required"`
	Owner  string                                                 `json:"owner"`
	JSON   interconnectNewResponseNscInterconnectPhysicalBodyJSON `json:"-"`
}

// interconnectNewResponseNscInterconnectPhysicalBodyJSON contains the JSON
// metadata for the struct [InterconnectNewResponseNscInterconnectPhysicalBody]
type interconnectNewResponseNscInterconnectPhysicalBodyJSON struct {
	Account     apijson.Field
	Facility    apijson.Field
	Name        apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectNewResponseNscInterconnectPhysicalBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectNewResponseNscInterconnectPhysicalBodyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectNewResponseNscInterconnectPhysicalBody) implementsInterconnectNewResponse() {}

type InterconnectNewResponseNscInterconnectPhysicalBodyFacility struct {
	Address []string                                                       `json:"address,required"`
	Name    string                                                         `json:"name,required"`
	JSON    interconnectNewResponseNscInterconnectPhysicalBodyFacilityJSON `json:"-"`
}

// interconnectNewResponseNscInterconnectPhysicalBodyFacilityJSON contains the JSON
// metadata for the struct
// [InterconnectNewResponseNscInterconnectPhysicalBodyFacility]
type interconnectNewResponseNscInterconnectPhysicalBodyFacilityJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectNewResponseNscInterconnectPhysicalBodyFacility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectNewResponseNscInterconnectPhysicalBodyFacilityJSON) RawJSON() string {
	return r.raw
}

type InterconnectNewResponseNscInterconnectGcpPartnerBody struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Region  string `json:"region,required"`
	Type    string `json:"type,required"`
	Owner   string `json:"owner"`
	// Bandwidth structure as visible through the customer-facing API.
	Speed InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed `json:"speed"`
	JSON  interconnectNewResponseNscInterconnectGcpPartnerBodyJSON  `json:"-"`
}

// interconnectNewResponseNscInterconnectGcpPartnerBodyJSON contains the JSON
// metadata for the struct [InterconnectNewResponseNscInterconnectGcpPartnerBody]
type interconnectNewResponseNscInterconnectGcpPartnerBodyJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectNewResponseNscInterconnectGcpPartnerBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectNewResponseNscInterconnectGcpPartnerBodyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectNewResponseNscInterconnectGcpPartnerBody) implementsInterconnectNewResponse() {}

// Bandwidth structure as visible through the customer-facing API.
type InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed string

const (
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed50M  InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "50M"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed100M InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "100M"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed200M InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "200M"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed300M InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "300M"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed400M InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "400M"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed500M InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "500M"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed1G   InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "1G"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed2G   InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "2G"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed5G   InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "5G"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed10G  InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "10G"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed20G  InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "20G"
	InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed50G  InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed = "50G"
)

func (r InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed) IsKnown() bool {
	switch r {
	case InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed50M, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed100M, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed200M, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed300M, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed400M, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed500M, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed1G, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed2G, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed5G, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed10G, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed20G, InterconnectNewResponseNscInterconnectGcpPartnerBodySpeed50G:
		return true
	}
	return false
}

type InterconnectListResponse struct {
	Items []InterconnectListResponseItem `json:"items,required"`
	Next  int64                          `json:"next,nullable"`
	JSON  interconnectListResponseJSON   `json:"-"`
}

// interconnectListResponseJSON contains the JSON metadata for the struct
// [InterconnectListResponse]
type interconnectListResponseJSON struct {
	Items       apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectListResponseJSON) RawJSON() string {
	return r.raw
}

type InterconnectListResponseItem struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Type    string `json:"type,required"`
	// This field can have the runtime type of
	// [InterconnectListResponseItemsNscInterconnectPhysicalBodyFacility].
	Facility interface{} `json:"facility"`
	Owner    string      `json:"owner"`
	Region   string      `json:"region"`
	// A Cloudflare site name.
	Site   string                           `json:"site"`
	SlotID string                           `json:"slot_id" format:"uuid"`
	Speed  string                           `json:"speed"`
	JSON   interconnectListResponseItemJSON `json:"-"`
	union  InterconnectListResponseItemsUnion
}

// interconnectListResponseItemJSON contains the JSON metadata for the struct
// [InterconnectListResponseItem]
type interconnectListResponseItemJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Facility    apijson.Field
	Owner       apijson.Field
	Region      apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r interconnectListResponseItemJSON) RawJSON() string {
	return r.raw
}

func (r *InterconnectListResponseItem) UnmarshalJSON(data []byte) (err error) {
	*r = InterconnectListResponseItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InterconnectListResponseItemsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InterconnectListResponseItemsNscInterconnectPhysicalBody],
// [InterconnectListResponseItemsNscInterconnectGcpPartnerBody].
func (r InterconnectListResponseItem) AsUnion() InterconnectListResponseItemsUnion {
	return r.union
}

// Union satisfied by [InterconnectListResponseItemsNscInterconnectPhysicalBody] or
// [InterconnectListResponseItemsNscInterconnectGcpPartnerBody].
type InterconnectListResponseItemsUnion interface {
	implementsInterconnectListResponseItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InterconnectListResponseItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InterconnectListResponseItemsNscInterconnectPhysicalBody{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InterconnectListResponseItemsNscInterconnectGcpPartnerBody{}),
		},
	)
}

type InterconnectListResponseItemsNscInterconnectPhysicalBody struct {
	Account  string                                                           `json:"account,required"`
	Facility InterconnectListResponseItemsNscInterconnectPhysicalBodyFacility `json:"facility,required"`
	Name     string                                                           `json:"name,required"`
	// A Cloudflare site name.
	Site   string                                                       `json:"site,required"`
	SlotID string                                                       `json:"slot_id,required" format:"uuid"`
	Speed  string                                                       `json:"speed,required"`
	Type   string                                                       `json:"type,required"`
	Owner  string                                                       `json:"owner"`
	JSON   interconnectListResponseItemsNscInterconnectPhysicalBodyJSON `json:"-"`
}

// interconnectListResponseItemsNscInterconnectPhysicalBodyJSON contains the JSON
// metadata for the struct
// [InterconnectListResponseItemsNscInterconnectPhysicalBody]
type interconnectListResponseItemsNscInterconnectPhysicalBodyJSON struct {
	Account     apijson.Field
	Facility    apijson.Field
	Name        apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectListResponseItemsNscInterconnectPhysicalBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectListResponseItemsNscInterconnectPhysicalBodyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectListResponseItemsNscInterconnectPhysicalBody) implementsInterconnectListResponseItem() {
}

type InterconnectListResponseItemsNscInterconnectPhysicalBodyFacility struct {
	Address []string                                                             `json:"address,required"`
	Name    string                                                               `json:"name,required"`
	JSON    interconnectListResponseItemsNscInterconnectPhysicalBodyFacilityJSON `json:"-"`
}

// interconnectListResponseItemsNscInterconnectPhysicalBodyFacilityJSON contains
// the JSON metadata for the struct
// [InterconnectListResponseItemsNscInterconnectPhysicalBodyFacility]
type interconnectListResponseItemsNscInterconnectPhysicalBodyFacilityJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectListResponseItemsNscInterconnectPhysicalBodyFacility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectListResponseItemsNscInterconnectPhysicalBodyFacilityJSON) RawJSON() string {
	return r.raw
}

type InterconnectListResponseItemsNscInterconnectGcpPartnerBody struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Region  string `json:"region,required"`
	Type    string `json:"type,required"`
	Owner   string `json:"owner"`
	// Bandwidth structure as visible through the customer-facing API.
	Speed InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed `json:"speed"`
	JSON  interconnectListResponseItemsNscInterconnectGcpPartnerBodyJSON  `json:"-"`
}

// interconnectListResponseItemsNscInterconnectGcpPartnerBodyJSON contains the JSON
// metadata for the struct
// [InterconnectListResponseItemsNscInterconnectGcpPartnerBody]
type interconnectListResponseItemsNscInterconnectGcpPartnerBodyJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectListResponseItemsNscInterconnectGcpPartnerBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectListResponseItemsNscInterconnectGcpPartnerBodyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectListResponseItemsNscInterconnectGcpPartnerBody) implementsInterconnectListResponseItem() {
}

// Bandwidth structure as visible through the customer-facing API.
type InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed string

const (
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed50M  InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "50M"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed100M InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "100M"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed200M InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "200M"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed300M InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "300M"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed400M InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "400M"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed500M InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "500M"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed1G   InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "1G"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed2G   InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "2G"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed5G   InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "5G"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed10G  InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "10G"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed20G  InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "20G"
	InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed50G  InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed = "50G"
)

func (r InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed) IsKnown() bool {
	switch r {
	case InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed50M, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed100M, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed200M, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed300M, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed400M, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed500M, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed1G, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed2G, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed5G, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed10G, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed20G, InterconnectListResponseItemsNscInterconnectGcpPartnerBodySpeed50G:
		return true
	}
	return false
}

type InterconnectGetResponse struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Type    string `json:"type,required"`
	// This field can have the runtime type of
	// [InterconnectGetResponseNscInterconnectPhysicalBodyFacility].
	Facility interface{} `json:"facility"`
	Owner    string      `json:"owner"`
	Region   string      `json:"region"`
	// A Cloudflare site name.
	Site   string                      `json:"site"`
	SlotID string                      `json:"slot_id" format:"uuid"`
	Speed  string                      `json:"speed"`
	JSON   interconnectGetResponseJSON `json:"-"`
	union  InterconnectGetResponseUnion
}

// interconnectGetResponseJSON contains the JSON metadata for the struct
// [InterconnectGetResponse]
type interconnectGetResponseJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Facility    apijson.Field
	Owner       apijson.Field
	Region      apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r interconnectGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *InterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = InterconnectGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InterconnectGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [InterconnectGetResponseNscInterconnectPhysicalBody],
// [InterconnectGetResponseNscInterconnectGcpPartnerBody].
func (r InterconnectGetResponse) AsUnion() InterconnectGetResponseUnion {
	return r.union
}

// Union satisfied by [InterconnectGetResponseNscInterconnectPhysicalBody] or
// [InterconnectGetResponseNscInterconnectGcpPartnerBody].
type InterconnectGetResponseUnion interface {
	implementsInterconnectGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InterconnectGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InterconnectGetResponseNscInterconnectPhysicalBody{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InterconnectGetResponseNscInterconnectGcpPartnerBody{}),
		},
	)
}

type InterconnectGetResponseNscInterconnectPhysicalBody struct {
	Account  string                                                     `json:"account,required"`
	Facility InterconnectGetResponseNscInterconnectPhysicalBodyFacility `json:"facility,required"`
	Name     string                                                     `json:"name,required"`
	// A Cloudflare site name.
	Site   string                                                 `json:"site,required"`
	SlotID string                                                 `json:"slot_id,required" format:"uuid"`
	Speed  string                                                 `json:"speed,required"`
	Type   string                                                 `json:"type,required"`
	Owner  string                                                 `json:"owner"`
	JSON   interconnectGetResponseNscInterconnectPhysicalBodyJSON `json:"-"`
}

// interconnectGetResponseNscInterconnectPhysicalBodyJSON contains the JSON
// metadata for the struct [InterconnectGetResponseNscInterconnectPhysicalBody]
type interconnectGetResponseNscInterconnectPhysicalBodyJSON struct {
	Account     apijson.Field
	Facility    apijson.Field
	Name        apijson.Field
	Site        apijson.Field
	SlotID      apijson.Field
	Speed       apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectGetResponseNscInterconnectPhysicalBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectGetResponseNscInterconnectPhysicalBodyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectGetResponseNscInterconnectPhysicalBody) implementsInterconnectGetResponse() {}

type InterconnectGetResponseNscInterconnectPhysicalBodyFacility struct {
	Address []string                                                       `json:"address,required"`
	Name    string                                                         `json:"name,required"`
	JSON    interconnectGetResponseNscInterconnectPhysicalBodyFacilityJSON `json:"-"`
}

// interconnectGetResponseNscInterconnectPhysicalBodyFacilityJSON contains the JSON
// metadata for the struct
// [InterconnectGetResponseNscInterconnectPhysicalBodyFacility]
type interconnectGetResponseNscInterconnectPhysicalBodyFacilityJSON struct {
	Address     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectGetResponseNscInterconnectPhysicalBodyFacility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectGetResponseNscInterconnectPhysicalBodyFacilityJSON) RawJSON() string {
	return r.raw
}

type InterconnectGetResponseNscInterconnectGcpPartnerBody struct {
	Account string `json:"account,required"`
	Name    string `json:"name,required"`
	Region  string `json:"region,required"`
	Type    string `json:"type,required"`
	Owner   string `json:"owner"`
	// Bandwidth structure as visible through the customer-facing API.
	Speed InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed `json:"speed"`
	JSON  interconnectGetResponseNscInterconnectGcpPartnerBodyJSON  `json:"-"`
}

// interconnectGetResponseNscInterconnectGcpPartnerBodyJSON contains the JSON
// metadata for the struct [InterconnectGetResponseNscInterconnectGcpPartnerBody]
type interconnectGetResponseNscInterconnectGcpPartnerBodyJSON struct {
	Account     apijson.Field
	Name        apijson.Field
	Region      apijson.Field
	Type        apijson.Field
	Owner       apijson.Field
	Speed       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectGetResponseNscInterconnectGcpPartnerBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectGetResponseNscInterconnectGcpPartnerBodyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectGetResponseNscInterconnectGcpPartnerBody) implementsInterconnectGetResponse() {}

// Bandwidth structure as visible through the customer-facing API.
type InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed string

const (
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed50M  InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "50M"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed100M InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "100M"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed200M InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "200M"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed300M InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "300M"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed400M InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "400M"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed500M InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "500M"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed1G   InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "1G"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed2G   InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "2G"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed5G   InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "5G"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed10G  InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "10G"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed20G  InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "20G"
	InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed50G  InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed = "50G"
)

func (r InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed) IsKnown() bool {
	switch r {
	case InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed50M, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed100M, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed200M, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed300M, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed400M, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed500M, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed1G, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed2G, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed5G, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed10G, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed20G, InterconnectGetResponseNscInterconnectGcpPartnerBodySpeed50G:
		return true
	}
	return false
}

type InterconnectStatusResponse struct {
	State InterconnectStatusResponseState `json:"state,required"`
	// Diagnostic information, if available
	Reason string                         `json:"reason,nullable"`
	JSON   interconnectStatusResponseJSON `json:"-"`
	union  InterconnectStatusResponseUnion
}

// interconnectStatusResponseJSON contains the JSON metadata for the struct
// [InterconnectStatusResponse]
type interconnectStatusResponseJSON struct {
	State       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r interconnectStatusResponseJSON) RawJSON() string {
	return r.raw
}

func (r *InterconnectStatusResponse) UnmarshalJSON(data []byte) (err error) {
	*r = InterconnectStatusResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InterconnectStatusResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [InterconnectStatusResponsePending],
// [InterconnectStatusResponseDown], [InterconnectStatusResponseUnhealthy],
// [InterconnectStatusResponseHealthy].
func (r InterconnectStatusResponse) AsUnion() InterconnectStatusResponseUnion {
	return r.union
}

// Union satisfied by [InterconnectStatusResponsePending],
// [InterconnectStatusResponseDown], [InterconnectStatusResponseUnhealthy] or
// [InterconnectStatusResponseHealthy].
type InterconnectStatusResponseUnion interface {
	implementsInterconnectStatusResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InterconnectStatusResponseUnion)(nil)).Elem(),
		"state",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InterconnectStatusResponsePending{}),
			DiscriminatorValue: "Pending",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InterconnectStatusResponseDown{}),
			DiscriminatorValue: "Down",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InterconnectStatusResponseUnhealthy{}),
			DiscriminatorValue: "Unhealthy",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InterconnectStatusResponseHealthy{}),
			DiscriminatorValue: "Healthy",
		},
	)
}

type InterconnectStatusResponsePending struct {
	State InterconnectStatusResponsePendingState `json:"state,required"`
	JSON  interconnectStatusResponsePendingJSON  `json:"-"`
}

// interconnectStatusResponsePendingJSON contains the JSON metadata for the struct
// [InterconnectStatusResponsePending]
type interconnectStatusResponsePendingJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectStatusResponsePending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectStatusResponsePendingJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectStatusResponsePending) implementsInterconnectStatusResponse() {}

type InterconnectStatusResponsePendingState string

const (
	InterconnectStatusResponsePendingStatePending InterconnectStatusResponsePendingState = "Pending"
)

func (r InterconnectStatusResponsePendingState) IsKnown() bool {
	switch r {
	case InterconnectStatusResponsePendingStatePending:
		return true
	}
	return false
}

type InterconnectStatusResponseDown struct {
	State InterconnectStatusResponseDownState `json:"state,required"`
	// Diagnostic information, if available
	Reason string                             `json:"reason,nullable"`
	JSON   interconnectStatusResponseDownJSON `json:"-"`
}

// interconnectStatusResponseDownJSON contains the JSON metadata for the struct
// [InterconnectStatusResponseDown]
type interconnectStatusResponseDownJSON struct {
	State       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectStatusResponseDown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectStatusResponseDownJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectStatusResponseDown) implementsInterconnectStatusResponse() {}

type InterconnectStatusResponseDownState string

const (
	InterconnectStatusResponseDownStateDown InterconnectStatusResponseDownState = "Down"
)

func (r InterconnectStatusResponseDownState) IsKnown() bool {
	switch r {
	case InterconnectStatusResponseDownStateDown:
		return true
	}
	return false
}

type InterconnectStatusResponseUnhealthy struct {
	State InterconnectStatusResponseUnhealthyState `json:"state,required"`
	// Diagnostic information, if available
	Reason string                                  `json:"reason,nullable"`
	JSON   interconnectStatusResponseUnhealthyJSON `json:"-"`
}

// interconnectStatusResponseUnhealthyJSON contains the JSON metadata for the
// struct [InterconnectStatusResponseUnhealthy]
type interconnectStatusResponseUnhealthyJSON struct {
	State       apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectStatusResponseUnhealthy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectStatusResponseUnhealthyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectStatusResponseUnhealthy) implementsInterconnectStatusResponse() {}

type InterconnectStatusResponseUnhealthyState string

const (
	InterconnectStatusResponseUnhealthyStateUnhealthy InterconnectStatusResponseUnhealthyState = "Unhealthy"
)

func (r InterconnectStatusResponseUnhealthyState) IsKnown() bool {
	switch r {
	case InterconnectStatusResponseUnhealthyStateUnhealthy:
		return true
	}
	return false
}

type InterconnectStatusResponseHealthy struct {
	State InterconnectStatusResponseHealthyState `json:"state,required"`
	JSON  interconnectStatusResponseHealthyJSON  `json:"-"`
}

// interconnectStatusResponseHealthyJSON contains the JSON metadata for the struct
// [InterconnectStatusResponseHealthy]
type interconnectStatusResponseHealthyJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InterconnectStatusResponseHealthy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r interconnectStatusResponseHealthyJSON) RawJSON() string {
	return r.raw
}

func (r InterconnectStatusResponseHealthy) implementsInterconnectStatusResponse() {}

type InterconnectStatusResponseHealthyState string

const (
	InterconnectStatusResponseHealthyStateHealthy InterconnectStatusResponseHealthyState = "Healthy"
)

func (r InterconnectStatusResponseHealthyState) IsKnown() bool {
	switch r {
	case InterconnectStatusResponseHealthyStateHealthy:
		return true
	}
	return false
}

type InterconnectStatusResponseState string

const (
	InterconnectStatusResponseStatePending   InterconnectStatusResponseState = "Pending"
	InterconnectStatusResponseStateDown      InterconnectStatusResponseState = "Down"
	InterconnectStatusResponseStateUnhealthy InterconnectStatusResponseState = "Unhealthy"
	InterconnectStatusResponseStateHealthy   InterconnectStatusResponseState = "Healthy"
)

func (r InterconnectStatusResponseState) IsKnown() bool {
	switch r {
	case InterconnectStatusResponseStatePending, InterconnectStatusResponseStateDown, InterconnectStatusResponseStateUnhealthy, InterconnectStatusResponseStateHealthy:
		return true
	}
	return false
}

type InterconnectNewParams struct {
	// Customer account tag
	AccountID param.Field[string]            `path:"account_id,required"`
	Body      InterconnectNewParamsBodyUnion `json:"body,required"`
}

func (r InterconnectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type InterconnectNewParamsBody struct {
	Account param.Field[string] `json:"account,required"`
	Type    param.Field[string] `json:"type,required"`
	// Bandwidth structure as visible through the customer-facing API.
	Bandwidth param.Field[InterconnectNewParamsBodyBandwidth] `json:"bandwidth"`
	// Pairing key provided by GCP
	PairingKey param.Field[string] `json:"pairing_key"`
	SlotID     param.Field[string] `json:"slot_id" format:"uuid"`
	Speed      param.Field[string] `json:"speed"`
}

func (r InterconnectNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InterconnectNewParamsBody) implementsInterconnectNewParamsBodyUnion() {}

// Satisfied by
// [network_interconnects.InterconnectNewParamsBodyNscInterconnectCreatePhysicalBody],
// [network_interconnects.InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody],
// [InterconnectNewParamsBody].
type InterconnectNewParamsBodyUnion interface {
	implementsInterconnectNewParamsBodyUnion()
}

type InterconnectNewParamsBodyNscInterconnectCreatePhysicalBody struct {
	Account param.Field[string] `json:"account,required"`
	SlotID  param.Field[string] `json:"slot_id,required" format:"uuid"`
	Type    param.Field[string] `json:"type,required"`
	Speed   param.Field[string] `json:"speed"`
}

func (r InterconnectNewParamsBodyNscInterconnectCreatePhysicalBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InterconnectNewParamsBodyNscInterconnectCreatePhysicalBody) implementsInterconnectNewParamsBodyUnion() {
}

type InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody struct {
	Account param.Field[string] `json:"account,required"`
	// Bandwidth structure as visible through the customer-facing API.
	Bandwidth param.Field[InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth] `json:"bandwidth,required"`
	// Pairing key provided by GCP
	PairingKey param.Field[string] `json:"pairing_key,required"`
	Type       param.Field[string] `json:"type,required"`
}

func (r InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBody) implementsInterconnectNewParamsBodyUnion() {
}

// Bandwidth structure as visible through the customer-facing API.
type InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth string

const (
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50M  InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "50M"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth100M InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "100M"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth200M InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "200M"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth300M InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "300M"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth400M InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "400M"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth500M InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "500M"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth1G   InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "1G"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth2G   InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "2G"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth5G   InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "5G"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth10G  InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "10G"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth20G  InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "20G"
	InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50G  InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth = "50G"
)

func (r InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth) IsKnown() bool {
	switch r {
	case InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50M, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth100M, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth200M, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth300M, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth400M, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth500M, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth1G, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth2G, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth5G, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth10G, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth20G, InterconnectNewParamsBodyNscInterconnectCreateGcpPartnerBodyBandwidth50G:
		return true
	}
	return false
}

// Bandwidth structure as visible through the customer-facing API.
type InterconnectNewParamsBodyBandwidth string

const (
	InterconnectNewParamsBodyBandwidth50M  InterconnectNewParamsBodyBandwidth = "50M"
	InterconnectNewParamsBodyBandwidth100M InterconnectNewParamsBodyBandwidth = "100M"
	InterconnectNewParamsBodyBandwidth200M InterconnectNewParamsBodyBandwidth = "200M"
	InterconnectNewParamsBodyBandwidth300M InterconnectNewParamsBodyBandwidth = "300M"
	InterconnectNewParamsBodyBandwidth400M InterconnectNewParamsBodyBandwidth = "400M"
	InterconnectNewParamsBodyBandwidth500M InterconnectNewParamsBodyBandwidth = "500M"
	InterconnectNewParamsBodyBandwidth1G   InterconnectNewParamsBodyBandwidth = "1G"
	InterconnectNewParamsBodyBandwidth2G   InterconnectNewParamsBodyBandwidth = "2G"
	InterconnectNewParamsBodyBandwidth5G   InterconnectNewParamsBodyBandwidth = "5G"
	InterconnectNewParamsBodyBandwidth10G  InterconnectNewParamsBodyBandwidth = "10G"
	InterconnectNewParamsBodyBandwidth20G  InterconnectNewParamsBodyBandwidth = "20G"
	InterconnectNewParamsBodyBandwidth50G  InterconnectNewParamsBodyBandwidth = "50G"
)

func (r InterconnectNewParamsBodyBandwidth) IsKnown() bool {
	switch r {
	case InterconnectNewParamsBodyBandwidth50M, InterconnectNewParamsBodyBandwidth100M, InterconnectNewParamsBodyBandwidth200M, InterconnectNewParamsBodyBandwidth300M, InterconnectNewParamsBodyBandwidth400M, InterconnectNewParamsBodyBandwidth500M, InterconnectNewParamsBodyBandwidth1G, InterconnectNewParamsBodyBandwidth2G, InterconnectNewParamsBodyBandwidth5G, InterconnectNewParamsBodyBandwidth10G, InterconnectNewParamsBodyBandwidth20G, InterconnectNewParamsBodyBandwidth50G:
		return true
	}
	return false
}

type InterconnectListParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
	Cursor    param.Field[int64]  `query:"cursor"`
	Limit     param.Field[int64]  `query:"limit"`
	// If specified, only show interconnects located at the given site
	Site param.Field[string] `query:"site"`
	// If specified, only show interconnects of the given type
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [InterconnectListParams]'s query parameters as `url.Values`.
func (r InterconnectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InterconnectDeleteParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}

type InterconnectGetParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}

type InterconnectLOAParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}

type InterconnectStatusParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}
