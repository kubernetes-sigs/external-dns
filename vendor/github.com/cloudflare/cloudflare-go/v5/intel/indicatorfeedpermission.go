// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

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

// IndicatorFeedPermissionService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndicatorFeedPermissionService] method instead.
type IndicatorFeedPermissionService struct {
	Options []option.RequestOption
}

// NewIndicatorFeedPermissionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndicatorFeedPermissionService(opts ...option.RequestOption) (r *IndicatorFeedPermissionService) {
	r = &IndicatorFeedPermissionService{}
	r.Options = opts
	return
}

// Grant permission to indicator feed
func (r *IndicatorFeedPermissionService) New(ctx context.Context, params IndicatorFeedPermissionNewParams, opts ...option.RequestOption) (res *IndicatorFeedPermissionNewResponse, err error) {
	var env IndicatorFeedPermissionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/add", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List indicator feed permissions
func (r *IndicatorFeedPermissionService) List(ctx context.Context, query IndicatorFeedPermissionListParams, opts ...option.RequestOption) (res *[]IndicatorFeedPermissionListResponse, err error) {
	var env IndicatorFeedPermissionListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/view", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revoke permission to indicator feed
func (r *IndicatorFeedPermissionService) Delete(ctx context.Context, params IndicatorFeedPermissionDeleteParams, opts ...option.RequestOption) (res *IndicatorFeedPermissionDeleteResponse, err error) {
	var env IndicatorFeedPermissionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/permissions/remove", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndicatorFeedPermissionNewResponse struct {
	// Whether the update succeeded or not
	Success bool                                   `json:"success"`
	JSON    indicatorFeedPermissionNewResponseJSON `json:"-"`
}

// indicatorFeedPermissionNewResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedPermissionNewResponse]
type indicatorFeedPermissionNewResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponse struct {
	// The unique identifier for the indicator feed
	ID int64 `json:"id"`
	// The description of the example test
	Description string `json:"description"`
	// Whether the indicator feed can be attributed to a provider
	IsAttributable bool `json:"is_attributable"`
	// Whether the indicator feed can be downloaded
	IsDownloadable bool `json:"is_downloadable"`
	// Whether the indicator feed is exposed to customers
	IsPublic bool `json:"is_public"`
	// The name of the indicator feed
	Name string                                  `json:"name"`
	JSON indicatorFeedPermissionListResponseJSON `json:"-"`
}

// indicatorFeedPermissionListResponseJSON contains the JSON metadata for the
// struct [IndicatorFeedPermissionListResponse]
type indicatorFeedPermissionListResponseJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	IsAttributable apijson.Field
	IsDownloadable apijson.Field
	IsPublic       apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponse struct {
	// Whether the update succeeded or not
	Success bool                                      `json:"success"`
	JSON    indicatorFeedPermissionDeleteResponseJSON `json:"-"`
}

// indicatorFeedPermissionDeleteResponseJSON contains the JSON metadata for the
// struct [IndicatorFeedPermissionDeleteResponse]
type indicatorFeedPermissionDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IndicatorFeedPermissionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedPermissionNewResponseEnvelope struct {
	Errors   []IndicatorFeedPermissionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedPermissionNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedPermissionNewResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedPermissionNewResponse                `json:"result"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionNewResponseEnvelope]
type indicatorFeedPermissionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           IndicatorFeedPermissionNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedPermissionNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IndicatorFeedPermissionNewResponseEnvelopeErrors]
type indicatorFeedPermissionNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionNewResponseEnvelopeErrorsSource]
type indicatorFeedPermissionNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           IndicatorFeedPermissionNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedPermissionNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionNewResponseEnvelopeMessages]
type indicatorFeedPermissionNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    indicatorFeedPermissionNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedPermissionNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [IndicatorFeedPermissionNewResponseEnvelopeMessagesSource]
type indicatorFeedPermissionNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedPermissionNewResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionNewResponseEnvelopeSuccessTrue IndicatorFeedPermissionNewResponseEnvelopeSuccess = true
)

func (r IndicatorFeedPermissionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedPermissionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedPermissionListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedPermissionListResponseEnvelope struct {
	Errors   []IndicatorFeedPermissionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedPermissionListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedPermissionListResponseEnvelopeSuccess `json:"success,required"`
	Result  []IndicatorFeedPermissionListResponse              `json:"result"`
	JSON    indicatorFeedPermissionListResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionListResponseEnvelope]
type indicatorFeedPermissionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           IndicatorFeedPermissionListResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedPermissionListResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IndicatorFeedPermissionListResponseEnvelopeErrors]
type indicatorFeedPermissionListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    indicatorFeedPermissionListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [IndicatorFeedPermissionListResponseEnvelopeErrorsSource]
type indicatorFeedPermissionListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           IndicatorFeedPermissionListResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedPermissionListResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionListResponseEnvelopeMessages]
type indicatorFeedPermissionListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionListResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    indicatorFeedPermissionListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedPermissionListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [IndicatorFeedPermissionListResponseEnvelopeMessagesSource]
type indicatorFeedPermissionListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedPermissionListResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionListResponseEnvelopeSuccessTrue IndicatorFeedPermissionListResponseEnvelopeSuccess = true
)

func (r IndicatorFeedPermissionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedPermissionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndicatorFeedPermissionDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The Cloudflare account tag of the account to change permissions on
	AccountTag param.Field[string] `json:"account_tag"`
	// The ID of the feed to add/remove permissions on
	FeedID param.Field[int64] `json:"feed_id"`
}

func (r IndicatorFeedPermissionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndicatorFeedPermissionDeleteResponseEnvelope struct {
	Errors   []IndicatorFeedPermissionDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndicatorFeedPermissionDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IndicatorFeedPermissionDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedPermissionDeleteResponse                `json:"result"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedPermissionDeleteResponseEnvelope]
type indicatorFeedPermissionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           IndicatorFeedPermissionDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionDeleteResponseEnvelopeErrors]
type indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [IndicatorFeedPermissionDeleteResponseEnvelopeErrorsSource]
type indicatorFeedPermissionDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           IndicatorFeedPermissionDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IndicatorFeedPermissionDeleteResponseEnvelopeMessages]
type indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedPermissionDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    indicatorFeedPermissionDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// indicatorFeedPermissionDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [IndicatorFeedPermissionDeleteResponseEnvelopeMessagesSource]
type indicatorFeedPermissionDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedPermissionDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedPermissionDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IndicatorFeedPermissionDeleteResponseEnvelopeSuccess bool

const (
	IndicatorFeedPermissionDeleteResponseEnvelopeSuccessTrue IndicatorFeedPermissionDeleteResponseEnvelopeSuccess = true
)

func (r IndicatorFeedPermissionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedPermissionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
