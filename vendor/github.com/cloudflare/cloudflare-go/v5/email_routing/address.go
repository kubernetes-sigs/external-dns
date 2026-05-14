// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
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

// AddressService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r *AddressService) {
	r = &AddressService{}
	r.Options = opts
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *AddressService) New(ctx context.Context, params AddressNewParams, opts ...option.RequestOption) (res *Address, err error) {
	var env AddressNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing destination addresses.
func (r *AddressService) List(ctx context.Context, params AddressListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Address], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", params.AccountID)
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

// Lists existing destination addresses.
func (r *AddressService) ListAutoPaging(ctx context.Context, params AddressListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Address] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a specific destination address.
func (r *AddressService) Delete(ctx context.Context, destinationAddressIdentifier string, body AddressDeleteParams, opts ...option.RequestOption) (res *Address, err error) {
	var env AddressDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if destinationAddressIdentifier == "" {
		err = errors.New("missing required destination_address_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", body.AccountID, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information for a specific destination email already created.
func (r *AddressService) Get(ctx context.Context, destinationAddressIdentifier string, query AddressGetParams, opts ...option.RequestOption) (res *Address, err error) {
	var env AddressGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if destinationAddressIdentifier == "" {
		err = errors.New("missing required destination_address_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", query.AccountID, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Address struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	//
	// Deprecated: deprecated
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time   `json:"verified" format:"date-time"`
	JSON     addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r AddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressNewResponseEnvelope struct {
	Errors   []AddressNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AddressNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Address                           `json:"result"`
	JSON    addressNewResponseEnvelopeJSON    `json:"-"`
}

// addressNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressNewResponseEnvelope]
type addressNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressNewResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AddressNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             addressNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// addressNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AddressNewResponseEnvelopeErrors]
type addressNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressNewResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    addressNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// addressNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AddressNewResponseEnvelopeErrorsSource]
type addressNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressNewResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AddressNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             addressNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// addressNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AddressNewResponseEnvelopeMessages]
type addressNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AddressNewResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    addressNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// addressNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AddressNewResponseEnvelopeMessagesSource]
type addressNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressNewResponseEnvelopeSuccess bool

const (
	AddressNewResponseEnvelopeSuccessTrue AddressNewResponseEnvelopeSuccess = true
)

func (r AddressNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Sorts results in an ascending or descending order.
	Direction param.Field[AddressListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[AddressListParamsVerified] `query:"verified"`
}

// URLQuery serializes [AddressListParams]'s query parameters as `url.Values`.
func (r AddressListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sorts results in an ascending or descending order.
type AddressListParamsDirection string

const (
	AddressListParamsDirectionAsc  AddressListParamsDirection = "asc"
	AddressListParamsDirectionDesc AddressListParamsDirection = "desc"
)

func (r AddressListParamsDirection) IsKnown() bool {
	switch r {
	case AddressListParamsDirectionAsc, AddressListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter by verified destination addresses.
type AddressListParamsVerified bool

const (
	AddressListParamsVerifiedTrue  AddressListParamsVerified = true
	AddressListParamsVerifiedFalse AddressListParamsVerified = false
)

func (r AddressListParamsVerified) IsKnown() bool {
	switch r {
	case AddressListParamsVerifiedTrue, AddressListParamsVerifiedFalse:
		return true
	}
	return false
}

type AddressDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressDeleteResponseEnvelope struct {
	Errors   []AddressDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AddressDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  Address                              `json:"result"`
	JSON    addressDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressDeleteResponseEnvelope]
type addressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressDeleteResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AddressDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             addressDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// addressDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressDeleteResponseEnvelopeErrors]
type addressDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    addressDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// addressDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AddressDeleteResponseEnvelopeErrorsSource]
type addressDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressDeleteResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AddressDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             addressDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// addressDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressDeleteResponseEnvelopeMessages]
type addressDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AddressDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    addressDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// addressDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AddressDeleteResponseEnvelopeMessagesSource]
type addressDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressDeleteResponseEnvelopeSuccess bool

const (
	AddressDeleteResponseEnvelopeSuccessTrue AddressDeleteResponseEnvelopeSuccess = true
)

func (r AddressDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressGetResponseEnvelope struct {
	Errors   []AddressGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AddressGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Address                           `json:"result"`
	JSON    addressGetResponseEnvelopeJSON    `json:"-"`
}

// addressGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressGetResponseEnvelope]
type addressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AddressGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             addressGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// addressGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AddressGetResponseEnvelopeErrors]
type addressGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    addressGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// addressGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AddressGetResponseEnvelopeErrorsSource]
type addressGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AddressGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             addressGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// addressGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AddressGetResponseEnvelopeMessages]
type addressGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AddressGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    addressGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// addressGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AddressGetResponseEnvelopeMessagesSource]
type addressGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressGetResponseEnvelopeSuccess bool

const (
	AddressGetResponseEnvelopeSuccessTrue AddressGetResponseEnvelopeSuccess = true
)

func (r AddressGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
