// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// SettingAllowPolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAllowPolicyService] method instead.
type SettingAllowPolicyService struct {
	Options []option.RequestOption
}

// NewSettingAllowPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAllowPolicyService(opts ...option.RequestOption) (r *SettingAllowPolicyService) {
	r = &SettingAllowPolicyService{}
	r.Options = opts
	return
}

// Create an email allow policy
func (r *SettingAllowPolicyService) New(ctx context.Context, params SettingAllowPolicyNewParams, opts ...option.RequestOption) (res *SettingAllowPolicyNewResponse, err error) {
	var env SettingAllowPolicyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists, searches, and sorts an account’s email allow policies.
func (r *SettingAllowPolicyService) List(ctx context.Context, params SettingAllowPolicyListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingAllowPolicyListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies", params.AccountID)
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

// Lists, searches, and sorts an account’s email allow policies.
func (r *SettingAllowPolicyService) ListAutoPaging(ctx context.Context, params SettingAllowPolicyListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingAllowPolicyListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an email allow policy
func (r *SettingAllowPolicyService) Delete(ctx context.Context, policyID int64, body SettingAllowPolicyDeleteParams, opts ...option.RequestOption) (res *SettingAllowPolicyDeleteResponse, err error) {
	var env SettingAllowPolicyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%v", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an email allow policy
func (r *SettingAllowPolicyService) Edit(ctx context.Context, policyID int64, params SettingAllowPolicyEditParams, opts ...option.RequestOption) (res *SettingAllowPolicyEditResponse, err error) {
	var env SettingAllowPolicyEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%v", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an email allow policy
func (r *SettingAllowPolicyService) Get(ctx context.Context, policyID int64, query SettingAllowPolicyGetParams, opts ...option.RequestOption) (res *SettingAllowPolicyGetResponse, err error) {
	var env SettingAllowPolicyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/allow_policies/%v", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingAllowPolicyNewResponse struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool                                     `json:"is_trusted_sender,required"`
	LastModified    time.Time                                `json:"last_modified,required" format:"date-time"`
	Pattern         string                                   `json:"pattern,required"`
	PatternType     SettingAllowPolicyNewResponsePatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                              `json:"is_spoof"`
	JSON    settingAllowPolicyNewResponseJSON `json:"-"`
}

// settingAllowPolicyNewResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyNewResponse]
type settingAllowPolicyNewResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyNewResponsePatternType string

const (
	SettingAllowPolicyNewResponsePatternTypeEmail   SettingAllowPolicyNewResponsePatternType = "EMAIL"
	SettingAllowPolicyNewResponsePatternTypeDomain  SettingAllowPolicyNewResponsePatternType = "DOMAIN"
	SettingAllowPolicyNewResponsePatternTypeIP      SettingAllowPolicyNewResponsePatternType = "IP"
	SettingAllowPolicyNewResponsePatternTypeUnknown SettingAllowPolicyNewResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyNewResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyNewResponsePatternTypeEmail, SettingAllowPolicyNewResponsePatternTypeDomain, SettingAllowPolicyNewResponsePatternTypeIP, SettingAllowPolicyNewResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyListResponse struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool                                      `json:"is_trusted_sender,required"`
	LastModified    time.Time                                 `json:"last_modified,required" format:"date-time"`
	Pattern         string                                    `json:"pattern,required"`
	PatternType     SettingAllowPolicyListResponsePatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                               `json:"is_spoof"`
	JSON    settingAllowPolicyListResponseJSON `json:"-"`
}

// settingAllowPolicyListResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyListResponse]
type settingAllowPolicyListResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyListResponsePatternType string

const (
	SettingAllowPolicyListResponsePatternTypeEmail   SettingAllowPolicyListResponsePatternType = "EMAIL"
	SettingAllowPolicyListResponsePatternTypeDomain  SettingAllowPolicyListResponsePatternType = "DOMAIN"
	SettingAllowPolicyListResponsePatternTypeIP      SettingAllowPolicyListResponsePatternType = "IP"
	SettingAllowPolicyListResponsePatternTypeUnknown SettingAllowPolicyListResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyListResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListResponsePatternTypeEmail, SettingAllowPolicyListResponsePatternTypeDomain, SettingAllowPolicyListResponsePatternTypeIP, SettingAllowPolicyListResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyDeleteResponse struct {
	// The unique identifier for the allow policy.
	ID   int64                                `json:"id,required"`
	JSON settingAllowPolicyDeleteResponseJSON `json:"-"`
}

// settingAllowPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyDeleteResponse]
type settingAllowPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditResponse struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool                                      `json:"is_trusted_sender,required"`
	LastModified    time.Time                                 `json:"last_modified,required" format:"date-time"`
	Pattern         string                                    `json:"pattern,required"`
	PatternType     SettingAllowPolicyEditResponsePatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                               `json:"is_spoof"`
	JSON    settingAllowPolicyEditResponseJSON `json:"-"`
}

// settingAllowPolicyEditResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyEditResponse]
type settingAllowPolicyEditResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditResponsePatternType string

const (
	SettingAllowPolicyEditResponsePatternTypeEmail   SettingAllowPolicyEditResponsePatternType = "EMAIL"
	SettingAllowPolicyEditResponsePatternTypeDomain  SettingAllowPolicyEditResponsePatternType = "DOMAIN"
	SettingAllowPolicyEditResponsePatternTypeIP      SettingAllowPolicyEditResponsePatternType = "IP"
	SettingAllowPolicyEditResponsePatternTypeUnknown SettingAllowPolicyEditResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyEditResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyEditResponsePatternTypeEmail, SettingAllowPolicyEditResponsePatternTypeDomain, SettingAllowPolicyEditResponsePatternTypeIP, SettingAllowPolicyEditResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyGetResponse struct {
	// The unique identifier for the allow policy.
	ID        int64     `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender bool `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient bool `json:"is_exempt_recipient,required"`
	IsRegex           bool `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender bool                                     `json:"is_trusted_sender,required"`
	LastModified    time.Time                                `json:"last_modified,required" format:"date-time"`
	Pattern         string                                   `json:"pattern,required"`
	PatternType     SettingAllowPolicyGetResponsePatternType `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender bool   `json:"verify_sender,required"`
	Comments     string `json:"comments,nullable"`
	// Deprecated: deprecated
	IsRecipient bool `json:"is_recipient"`
	// Deprecated: deprecated
	IsSender bool `json:"is_sender"`
	// Deprecated: deprecated
	IsSpoof bool                              `json:"is_spoof"`
	JSON    settingAllowPolicyGetResponseJSON `json:"-"`
}

// settingAllowPolicyGetResponseJSON contains the JSON metadata for the struct
// [SettingAllowPolicyGetResponse]
type settingAllowPolicyGetResponseJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	IsAcceptableSender apijson.Field
	IsExemptRecipient  apijson.Field
	IsRegex            apijson.Field
	IsTrustedSender    apijson.Field
	LastModified       apijson.Field
	Pattern            apijson.Field
	PatternType        apijson.Field
	VerifySender       apijson.Field
	Comments           apijson.Field
	IsRecipient        apijson.Field
	IsSender           apijson.Field
	IsSpoof            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyGetResponsePatternType string

const (
	SettingAllowPolicyGetResponsePatternTypeEmail   SettingAllowPolicyGetResponsePatternType = "EMAIL"
	SettingAllowPolicyGetResponsePatternTypeDomain  SettingAllowPolicyGetResponsePatternType = "DOMAIN"
	SettingAllowPolicyGetResponsePatternTypeIP      SettingAllowPolicyGetResponsePatternType = "IP"
	SettingAllowPolicyGetResponsePatternTypeUnknown SettingAllowPolicyGetResponsePatternType = "UNKNOWN"
)

func (r SettingAllowPolicyGetResponsePatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyGetResponsePatternTypeEmail, SettingAllowPolicyGetResponsePatternTypeDomain, SettingAllowPolicyGetResponsePatternTypeIP, SettingAllowPolicyGetResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyNewParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender,required"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient,required"`
	IsRegex           param.Field[bool] `json:"is_regex,required"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender param.Field[bool]                                   `json:"is_trusted_sender,required"`
	Pattern         param.Field[string]                                 `json:"pattern,required"`
	PatternType     param.Field[SettingAllowPolicyNewParamsPatternType] `json:"pattern_type,required"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool]   `json:"verify_sender,required"`
	Comments     param.Field[string] `json:"comments"`
	IsRecipient  param.Field[bool]   `json:"is_recipient"`
	IsSender     param.Field[bool]   `json:"is_sender"`
	IsSpoof      param.Field[bool]   `json:"is_spoof"`
}

func (r SettingAllowPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAllowPolicyNewParamsPatternType string

const (
	SettingAllowPolicyNewParamsPatternTypeEmail   SettingAllowPolicyNewParamsPatternType = "EMAIL"
	SettingAllowPolicyNewParamsPatternTypeDomain  SettingAllowPolicyNewParamsPatternType = "DOMAIN"
	SettingAllowPolicyNewParamsPatternTypeIP      SettingAllowPolicyNewParamsPatternType = "IP"
	SettingAllowPolicyNewParamsPatternTypeUnknown SettingAllowPolicyNewParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyNewParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyNewParamsPatternTypeEmail, SettingAllowPolicyNewParamsPatternTypeDomain, SettingAllowPolicyNewParamsPatternTypeIP, SettingAllowPolicyNewParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   SettingAllowPolicyNewResponse             `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     settingAllowPolicyNewResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyNewResponseEnvelope]
type settingAllowPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The sorting direction.
	Direction          param.Field[SettingAllowPolicyListParamsDirection] `query:"direction"`
	IsAcceptableSender param.Field[bool]                                  `query:"is_acceptable_sender"`
	IsExemptRecipient  param.Field[bool]                                  `query:"is_exempt_recipient"`
	IsRecipient        param.Field[bool]                                  `query:"is_recipient"`
	IsSender           param.Field[bool]                                  `query:"is_sender"`
	IsSpoof            param.Field[bool]                                  `query:"is_spoof"`
	IsTrustedSender    param.Field[bool]                                  `query:"is_trusted_sender"`
	// The field to sort by.
	Order param.Field[SettingAllowPolicyListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page        param.Field[int64]                                   `query:"page"`
	Pattern     param.Field[string]                                  `query:"pattern"`
	PatternType param.Field[SettingAllowPolicyListParamsPatternType] `query:"pattern_type"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search       param.Field[string] `query:"search"`
	VerifySender param.Field[bool]   `query:"verify_sender"`
}

// URLQuery serializes [SettingAllowPolicyListParams]'s query parameters as
// `url.Values`.
func (r SettingAllowPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingAllowPolicyListParamsDirection string

const (
	SettingAllowPolicyListParamsDirectionAsc  SettingAllowPolicyListParamsDirection = "asc"
	SettingAllowPolicyListParamsDirectionDesc SettingAllowPolicyListParamsDirection = "desc"
)

func (r SettingAllowPolicyListParamsDirection) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListParamsDirectionAsc, SettingAllowPolicyListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type SettingAllowPolicyListParamsOrder string

const (
	SettingAllowPolicyListParamsOrderPattern   SettingAllowPolicyListParamsOrder = "pattern"
	SettingAllowPolicyListParamsOrderCreatedAt SettingAllowPolicyListParamsOrder = "created_at"
)

func (r SettingAllowPolicyListParamsOrder) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListParamsOrderPattern, SettingAllowPolicyListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingAllowPolicyListParamsPatternType string

const (
	SettingAllowPolicyListParamsPatternTypeEmail   SettingAllowPolicyListParamsPatternType = "EMAIL"
	SettingAllowPolicyListParamsPatternTypeDomain  SettingAllowPolicyListParamsPatternType = "DOMAIN"
	SettingAllowPolicyListParamsPatternTypeIP      SettingAllowPolicyListParamsPatternType = "IP"
	SettingAllowPolicyListParamsPatternTypeUnknown SettingAllowPolicyListParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyListParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyListParamsPatternTypeEmail, SettingAllowPolicyListParamsPatternTypeDomain, SettingAllowPolicyListParamsPatternTypeIP, SettingAllowPolicyListParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyDeleteParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAllowPolicyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                        `json:"errors,required"`
	Messages []shared.ResponseInfo                        `json:"messages,required"`
	Result   SettingAllowPolicyDeleteResponse             `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	JSON     settingAllowPolicyDeleteResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyDeleteResponseEnvelope]
type settingAllowPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyEditParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Comments  param.Field[string] `json:"comments"`
	// Messages from this sender will be exempted from Spam, Spoof and Bulk
	// dispositions. Note: This will not exempt messages with Malicious or Suspicious
	// dispositions.
	IsAcceptableSender param.Field[bool] `json:"is_acceptable_sender"`
	// Messages to this recipient will bypass all detections.
	IsExemptRecipient param.Field[bool] `json:"is_exempt_recipient"`
	IsRegex           param.Field[bool] `json:"is_regex"`
	// Messages from this sender will bypass all detections and link following.
	IsTrustedSender param.Field[bool]                                    `json:"is_trusted_sender"`
	Pattern         param.Field[string]                                  `json:"pattern"`
	PatternType     param.Field[SettingAllowPolicyEditParamsPatternType] `json:"pattern_type"`
	// Enforce DMARC, SPF or DKIM authentication. When on, Email Security only honors
	// policies that pass authentication.
	VerifySender param.Field[bool] `json:"verify_sender"`
}

func (r SettingAllowPolicyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAllowPolicyEditParamsPatternType string

const (
	SettingAllowPolicyEditParamsPatternTypeEmail   SettingAllowPolicyEditParamsPatternType = "EMAIL"
	SettingAllowPolicyEditParamsPatternTypeDomain  SettingAllowPolicyEditParamsPatternType = "DOMAIN"
	SettingAllowPolicyEditParamsPatternTypeIP      SettingAllowPolicyEditParamsPatternType = "IP"
	SettingAllowPolicyEditParamsPatternTypeUnknown SettingAllowPolicyEditParamsPatternType = "UNKNOWN"
)

func (r SettingAllowPolicyEditParamsPatternType) IsKnown() bool {
	switch r {
	case SettingAllowPolicyEditParamsPatternTypeEmail, SettingAllowPolicyEditParamsPatternTypeDomain, SettingAllowPolicyEditParamsPatternTypeIP, SettingAllowPolicyEditParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingAllowPolicyEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                      `json:"errors,required"`
	Messages []shared.ResponseInfo                      `json:"messages,required"`
	Result   SettingAllowPolicyEditResponse             `json:"result,required"`
	Success  bool                                       `json:"success,required"`
	JSON     settingAllowPolicyEditResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyEditResponseEnvelope]
type settingAllowPolicyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAllowPolicyGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAllowPolicyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   SettingAllowPolicyGetResponse             `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     settingAllowPolicyGetResponseEnvelopeJSON `json:"-"`
}

// settingAllowPolicyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAllowPolicyGetResponseEnvelope]
type settingAllowPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAllowPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAllowPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
