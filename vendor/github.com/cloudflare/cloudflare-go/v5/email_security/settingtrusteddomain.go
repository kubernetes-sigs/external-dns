// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// SettingTrustedDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingTrustedDomainService] method instead.
type SettingTrustedDomainService struct {
	Options []option.RequestOption
}

// NewSettingTrustedDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingTrustedDomainService(opts ...option.RequestOption) (r *SettingTrustedDomainService) {
	r = &SettingTrustedDomainService{}
	r.Options = opts
	return
}

// Create a trusted email domain
func (r *SettingTrustedDomainService) New(ctx context.Context, params SettingTrustedDomainNewParams, opts ...option.RequestOption) (res *SettingTrustedDomainNewResponseUnion, err error) {
	var env SettingTrustedDomainNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists, searches, and sorts an account’s trusted email domains.
func (r *SettingTrustedDomainService) List(ctx context.Context, params SettingTrustedDomainListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingTrustedDomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains", params.AccountID)
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

// Lists, searches, and sorts an account’s trusted email domains.
func (r *SettingTrustedDomainService) ListAutoPaging(ctx context.Context, params SettingTrustedDomainListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingTrustedDomainListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a trusted email domain
func (r *SettingTrustedDomainService) Delete(ctx context.Context, trustedDomainID int64, body SettingTrustedDomainDeleteParams, opts ...option.RequestOption) (res *SettingTrustedDomainDeleteResponse, err error) {
	var env SettingTrustedDomainDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%v", body.AccountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a trusted email domain
func (r *SettingTrustedDomainService) Edit(ctx context.Context, trustedDomainID int64, params SettingTrustedDomainEditParams, opts ...option.RequestOption) (res *SettingTrustedDomainEditResponse, err error) {
	var env SettingTrustedDomainEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%v", params.AccountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a trusted email domain
func (r *SettingTrustedDomainService) Get(ctx context.Context, trustedDomainID int64, query SettingTrustedDomainGetParams, opts ...option.RequestOption) (res *SettingTrustedDomainGetResponse, err error) {
	var env SettingTrustedDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%v", query.AccountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [SettingTrustedDomainNewResponseEmailSecurityTrustedDomain]
// or [SettingTrustedDomainNewResponseArray].
type SettingTrustedDomainNewResponseUnion interface {
	implementsSettingTrustedDomainNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingTrustedDomainNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingTrustedDomainNewResponseEmailSecurityTrustedDomain{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingTrustedDomainNewResponseArray{}),
		},
	)
}

type SettingTrustedDomainNewResponseEmailSecurityTrustedDomain struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                                          `json:"is_similarity,required"`
	LastModified time.Time                                                     `json:"last_modified,required" format:"date-time"`
	Pattern      string                                                        `json:"pattern,required"`
	Comments     string                                                        `json:"comments,nullable"`
	JSON         settingTrustedDomainNewResponseEmailSecurityTrustedDomainJSON `json:"-"`
}

// settingTrustedDomainNewResponseEmailSecurityTrustedDomainJSON contains the JSON
// metadata for the struct
// [SettingTrustedDomainNewResponseEmailSecurityTrustedDomain]
type settingTrustedDomainNewResponseEmailSecurityTrustedDomainJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEmailSecurityTrustedDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEmailSecurityTrustedDomainJSON) RawJSON() string {
	return r.raw
}

func (r SettingTrustedDomainNewResponseEmailSecurityTrustedDomain) implementsSettingTrustedDomainNewResponseUnion() {
}

type SettingTrustedDomainNewResponseArray []SettingTrustedDomainNewResponseArrayItem

func (r SettingTrustedDomainNewResponseArray) implementsSettingTrustedDomainNewResponseUnion() {}

type SettingTrustedDomainNewResponseArrayItem struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                         `json:"is_similarity,required"`
	LastModified time.Time                                    `json:"last_modified,required" format:"date-time"`
	Pattern      string                                       `json:"pattern,required"`
	Comments     string                                       `json:"comments,nullable"`
	JSON         settingTrustedDomainNewResponseArrayItemJSON `json:"-"`
}

// settingTrustedDomainNewResponseArrayItemJSON contains the JSON metadata for the
// struct [SettingTrustedDomainNewResponseArrayItem]
type settingTrustedDomainNewResponseArrayItemJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseArrayItemJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainListResponse struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                 `json:"is_similarity,required"`
	LastModified time.Time                            `json:"last_modified,required" format:"date-time"`
	Pattern      string                               `json:"pattern,required"`
	Comments     string                               `json:"comments,nullable"`
	JSON         settingTrustedDomainListResponseJSON `json:"-"`
}

// settingTrustedDomainListResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainListResponse]
type settingTrustedDomainListResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainDeleteResponse struct {
	// The unique identifier for the trusted domain.
	ID   int64                                  `json:"id,required"`
	JSON settingTrustedDomainDeleteResponseJSON `json:"-"`
}

// settingTrustedDomainDeleteResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainDeleteResponse]
type settingTrustedDomainDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainEditResponse struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                 `json:"is_similarity,required"`
	LastModified time.Time                            `json:"last_modified,required" format:"date-time"`
	Pattern      string                               `json:"pattern,required"`
	Comments     string                               `json:"comments,nullable"`
	JSON         settingTrustedDomainEditResponseJSON `json:"-"`
}

// settingTrustedDomainEditResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainEditResponse]
type settingTrustedDomainEditResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainGetResponse struct {
	// The unique identifier for the trusted domain.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent,required"`
	IsRegex  bool `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool                                `json:"is_similarity,required"`
	LastModified time.Time                           `json:"last_modified,required" format:"date-time"`
	Pattern      string                              `json:"pattern,required"`
	Comments     string                              `json:"comments,nullable"`
	JSON         settingTrustedDomainGetResponseJSON `json:"-"`
}

// settingTrustedDomainGetResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainGetResponse]
type settingTrustedDomainGetResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainNewParams struct {
	// Account Identifier
	AccountID param.Field[string]                    `path:"account_id,required"`
	Body      SettingTrustedDomainNewParamsBodyUnion `json:"body,required"`
}

func (r SettingTrustedDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Satisfied by
// [email_security.SettingTrustedDomainNewParamsBodyEmailSecurityCreateTrustedDomain],
// [email_security.SettingTrustedDomainNewParamsBodyArray].
type SettingTrustedDomainNewParamsBodyUnion interface {
	implementsSettingTrustedDomainNewParamsBodyUnion()
}

type SettingTrustedDomainNewParamsBodyEmailSecurityCreateTrustedDomain struct {
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent,required"`
	IsRegex  param.Field[bool] `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity,required"`
	Pattern      param.Field[string] `json:"pattern,required"`
	Comments     param.Field[string] `json:"comments"`
}

func (r SettingTrustedDomainNewParamsBodyEmailSecurityCreateTrustedDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingTrustedDomainNewParamsBodyEmailSecurityCreateTrustedDomain) implementsSettingTrustedDomainNewParamsBodyUnion() {
}

type SettingTrustedDomainNewParamsBodyArray []SettingTrustedDomainNewParamsBodyArrayItem

func (r SettingTrustedDomainNewParamsBodyArray) implementsSettingTrustedDomainNewParamsBodyUnion() {}

type SettingTrustedDomainNewParamsBodyArrayItem struct {
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent,required"`
	IsRegex  param.Field[bool] `json:"is_regex,required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity,required"`
	Pattern      param.Field[string] `json:"pattern,required"`
	Comments     param.Field[string] `json:"comments"`
}

func (r SettingTrustedDomainNewParamsBodyArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingTrustedDomainNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   SettingTrustedDomainNewResponseUnion        `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     settingTrustedDomainNewResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTrustedDomainNewResponseEnvelope]
type settingTrustedDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The sorting direction.
	Direction    param.Field[SettingTrustedDomainListParamsDirection] `query:"direction"`
	IsRecent     param.Field[bool]                                    `query:"is_recent"`
	IsSimilarity param.Field[bool]                                    `query:"is_similarity"`
	// The field to sort by.
	Order param.Field[SettingTrustedDomainListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page    param.Field[int64]  `query:"page"`
	Pattern param.Field[string] `query:"pattern"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingTrustedDomainListParams]'s query parameters as
// `url.Values`.
func (r SettingTrustedDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingTrustedDomainListParamsDirection string

const (
	SettingTrustedDomainListParamsDirectionAsc  SettingTrustedDomainListParamsDirection = "asc"
	SettingTrustedDomainListParamsDirectionDesc SettingTrustedDomainListParamsDirection = "desc"
)

func (r SettingTrustedDomainListParamsDirection) IsKnown() bool {
	switch r {
	case SettingTrustedDomainListParamsDirectionAsc, SettingTrustedDomainListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type SettingTrustedDomainListParamsOrder string

const (
	SettingTrustedDomainListParamsOrderPattern   SettingTrustedDomainListParamsOrder = "pattern"
	SettingTrustedDomainListParamsOrderCreatedAt SettingTrustedDomainListParamsOrder = "created_at"
)

func (r SettingTrustedDomainListParamsOrder) IsKnown() bool {
	switch r {
	case SettingTrustedDomainListParamsOrderPattern, SettingTrustedDomainListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingTrustedDomainDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingTrustedDomainDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                          `json:"errors,required"`
	Messages []shared.ResponseInfo                          `json:"messages,required"`
	Result   SettingTrustedDomainDeleteResponse             `json:"result,required"`
	Success  bool                                           `json:"success,required"`
	JSON     settingTrustedDomainDeleteResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrustedDomainDeleteResponseEnvelope]
type settingTrustedDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainEditParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Comments  param.Field[string] `json:"comments"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent"`
	IsRegex  param.Field[bool] `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity"`
	Pattern      param.Field[string] `json:"pattern"`
}

func (r SettingTrustedDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingTrustedDomainEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                        `json:"errors,required"`
	Messages []shared.ResponseInfo                        `json:"messages,required"`
	Result   SettingTrustedDomainEditResponse             `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	JSON     settingTrustedDomainEditResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTrustedDomainEditResponseEnvelope]
type settingTrustedDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingTrustedDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   SettingTrustedDomainGetResponse             `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     settingTrustedDomainGetResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTrustedDomainGetResponseEnvelope]
type settingTrustedDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
