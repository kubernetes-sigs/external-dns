// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// AccessApplicationCAService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationCAService] method instead.
type AccessApplicationCAService struct {
	Options []option.RequestOption
}

// NewAccessApplicationCAService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationCAService(opts ...option.RequestOption) (r *AccessApplicationCAService) {
	r = &AccessApplicationCAService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccessApplicationCAService) New(ctx context.Context, appID string, body AccessApplicationCANewParams, opts ...option.RequestOption) (res *CA, err error) {
	var env AccessApplicationCANewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessApplicationCAService) List(ctx context.Context, params AccessApplicationCAListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CA], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/ca", accountOrZone, accountOrZoneID)
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

// Lists short-lived certificate CAs and their public keys.
func (r *AccessApplicationCAService) ListAutoPaging(ctx context.Context, params AccessApplicationCAListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CA] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a short-lived certificate CA.
func (r *AccessApplicationCAService) Delete(ctx context.Context, appID string, body AccessApplicationCADeleteParams, opts ...option.RequestOption) (res *AccessApplicationCADeleteResponse, err error) {
	var env AccessApplicationCADeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessApplicationCAService) Get(ctx context.Context, appID string, query AccessApplicationCAGetParams, opts ...option.RequestOption) (res *CA, err error) {
	var env AccessApplicationCAGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CA struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	AUD string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string `json:"public_key"`
	JSON      caJSON `json:"-"`
}

// caJSON contains the JSON metadata for the struct [CA]
type caJSON struct {
	ID          apijson.Field
	AUD         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponse struct {
	// The ID of the CA.
	ID   string                                `json:"id"`
	JSON accessApplicationCADeleteResponseJSON `json:"-"`
}

// accessApplicationCADeleteResponseJSON contains the JSON metadata for the struct
// [AccessApplicationCADeleteResponse]
type accessApplicationCADeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCANewResponseEnvelope struct {
	Errors   []AccessApplicationCANewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCANewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationCANewResponseEnvelopeSuccess `json:"success,required"`
	Result  CA                                            `json:"result"`
	JSON    accessApplicationCANewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCANewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCANewResponseEnvelope]
type accessApplicationCANewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccessApplicationCANewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationCANewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationCANewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCANewResponseEnvelopeErrors]
type accessApplicationCANewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accessApplicationCANewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationCANewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessApplicationCANewResponseEnvelopeErrorsSource]
type accessApplicationCANewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessApplicationCANewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationCANewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationCANewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCANewResponseEnvelopeMessages]
type accessApplicationCANewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessApplicationCANewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationCANewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessApplicationCANewResponseEnvelopeMessagesSource]
type accessApplicationCANewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationCANewResponseEnvelopeSuccess bool

const (
	AccessApplicationCANewResponseEnvelopeSuccessTrue AccessApplicationCANewResponseEnvelopeSuccess = true
)

func (r AccessApplicationCANewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCANewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCAListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccessApplicationCAListParams]'s query parameters as
// `url.Values`.
func (r AccessApplicationCAListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessApplicationCADeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCADeleteResponseEnvelope struct {
	Errors   []AccessApplicationCADeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCADeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationCADeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationCADeleteResponse                `json:"result"`
	JSON    accessApplicationCADeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCADeleteResponseEnvelope]
type accessApplicationCADeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccessApplicationCADeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationCADeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationCADeleteResponseEnvelopeErrors]
type accessApplicationCADeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accessApplicationCADeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessApplicationCADeleteResponseEnvelopeErrorsSource]
type accessApplicationCADeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessApplicationCADeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationCADeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCADeleteResponseEnvelopeMessages]
type accessApplicationCADeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessApplicationCADeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationCADeleteResponseEnvelopeMessagesSource]
type accessApplicationCADeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationCADeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationCADeleteResponseEnvelopeSuccessTrue AccessApplicationCADeleteResponseEnvelopeSuccess = true
)

func (r AccessApplicationCADeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCADeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCAGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCAGetResponseEnvelope struct {
	Errors   []AccessApplicationCAGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCAGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationCAGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CA                                            `json:"result"`
	JSON    accessApplicationCAGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCAGetResponseEnvelope]
type accessApplicationCAGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccessApplicationCAGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationCAGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCAGetResponseEnvelopeErrors]
type accessApplicationCAGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accessApplicationCAGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessApplicationCAGetResponseEnvelopeErrorsSource]
type accessApplicationCAGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessApplicationCAGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationCAGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCAGetResponseEnvelopeMessages]
type accessApplicationCAGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessApplicationCAGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessApplicationCAGetResponseEnvelopeMessagesSource]
type accessApplicationCAGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationCAGetResponseEnvelopeSuccess bool

const (
	AccessApplicationCAGetResponseEnvelopeSuccessTrue AccessApplicationCAGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationCAGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCAGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
