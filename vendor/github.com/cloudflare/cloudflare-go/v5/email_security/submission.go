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
)

// SubmissionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubmissionService] method instead.
type SubmissionService struct {
	Options []option.RequestOption
}

// NewSubmissionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubmissionService(opts ...option.RequestOption) (r *SubmissionService) {
	r = &SubmissionService{}
	r.Options = opts
	return
}

// This endpoint returns information for submissions to made to reclassify emails.
func (r *SubmissionService) List(ctx context.Context, params SubmissionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SubmissionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/submissions", params.AccountID)
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

// This endpoint returns information for submissions to made to reclassify emails.
func (r *SubmissionService) ListAutoPaging(ctx context.Context, params SubmissionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SubmissionListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type SubmissionListResponse struct {
	RequestedTs          time.Time                                  `json:"requested_ts,required" format:"date-time"`
	SubmissionID         string                                     `json:"submission_id,required"`
	OriginalDisposition  SubmissionListResponseOriginalDisposition  `json:"original_disposition,nullable"`
	OriginalEdfHash      string                                     `json:"original_edf_hash,nullable"`
	Outcome              string                                     `json:"outcome,nullable"`
	OutcomeDisposition   SubmissionListResponseOutcomeDisposition   `json:"outcome_disposition,nullable"`
	RequestedBy          string                                     `json:"requested_by,nullable"`
	RequestedDisposition SubmissionListResponseRequestedDisposition `json:"requested_disposition,nullable"`
	Status               string                                     `json:"status,nullable"`
	Subject              string                                     `json:"subject,nullable"`
	Type                 string                                     `json:"type,nullable"`
	JSON                 submissionListResponseJSON                 `json:"-"`
}

// submissionListResponseJSON contains the JSON metadata for the struct
// [SubmissionListResponse]
type submissionListResponseJSON struct {
	RequestedTs          apijson.Field
	SubmissionID         apijson.Field
	OriginalDisposition  apijson.Field
	OriginalEdfHash      apijson.Field
	Outcome              apijson.Field
	OutcomeDisposition   apijson.Field
	RequestedBy          apijson.Field
	RequestedDisposition apijson.Field
	Status               apijson.Field
	Subject              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SubmissionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submissionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubmissionListResponseOriginalDisposition string

const (
	SubmissionListResponseOriginalDispositionMalicious    SubmissionListResponseOriginalDisposition = "MALICIOUS"
	SubmissionListResponseOriginalDispositionMaliciousBec SubmissionListResponseOriginalDisposition = "MALICIOUS-BEC"
	SubmissionListResponseOriginalDispositionSuspicious   SubmissionListResponseOriginalDisposition = "SUSPICIOUS"
	SubmissionListResponseOriginalDispositionSpoof        SubmissionListResponseOriginalDisposition = "SPOOF"
	SubmissionListResponseOriginalDispositionSpam         SubmissionListResponseOriginalDisposition = "SPAM"
	SubmissionListResponseOriginalDispositionBulk         SubmissionListResponseOriginalDisposition = "BULK"
	SubmissionListResponseOriginalDispositionEncrypted    SubmissionListResponseOriginalDisposition = "ENCRYPTED"
	SubmissionListResponseOriginalDispositionExternal     SubmissionListResponseOriginalDisposition = "EXTERNAL"
	SubmissionListResponseOriginalDispositionUnknown      SubmissionListResponseOriginalDisposition = "UNKNOWN"
	SubmissionListResponseOriginalDispositionNone         SubmissionListResponseOriginalDisposition = "NONE"
)

func (r SubmissionListResponseOriginalDisposition) IsKnown() bool {
	switch r {
	case SubmissionListResponseOriginalDispositionMalicious, SubmissionListResponseOriginalDispositionMaliciousBec, SubmissionListResponseOriginalDispositionSuspicious, SubmissionListResponseOriginalDispositionSpoof, SubmissionListResponseOriginalDispositionSpam, SubmissionListResponseOriginalDispositionBulk, SubmissionListResponseOriginalDispositionEncrypted, SubmissionListResponseOriginalDispositionExternal, SubmissionListResponseOriginalDispositionUnknown, SubmissionListResponseOriginalDispositionNone:
		return true
	}
	return false
}

type SubmissionListResponseOutcomeDisposition string

const (
	SubmissionListResponseOutcomeDispositionMalicious    SubmissionListResponseOutcomeDisposition = "MALICIOUS"
	SubmissionListResponseOutcomeDispositionMaliciousBec SubmissionListResponseOutcomeDisposition = "MALICIOUS-BEC"
	SubmissionListResponseOutcomeDispositionSuspicious   SubmissionListResponseOutcomeDisposition = "SUSPICIOUS"
	SubmissionListResponseOutcomeDispositionSpoof        SubmissionListResponseOutcomeDisposition = "SPOOF"
	SubmissionListResponseOutcomeDispositionSpam         SubmissionListResponseOutcomeDisposition = "SPAM"
	SubmissionListResponseOutcomeDispositionBulk         SubmissionListResponseOutcomeDisposition = "BULK"
	SubmissionListResponseOutcomeDispositionEncrypted    SubmissionListResponseOutcomeDisposition = "ENCRYPTED"
	SubmissionListResponseOutcomeDispositionExternal     SubmissionListResponseOutcomeDisposition = "EXTERNAL"
	SubmissionListResponseOutcomeDispositionUnknown      SubmissionListResponseOutcomeDisposition = "UNKNOWN"
	SubmissionListResponseOutcomeDispositionNone         SubmissionListResponseOutcomeDisposition = "NONE"
)

func (r SubmissionListResponseOutcomeDisposition) IsKnown() bool {
	switch r {
	case SubmissionListResponseOutcomeDispositionMalicious, SubmissionListResponseOutcomeDispositionMaliciousBec, SubmissionListResponseOutcomeDispositionSuspicious, SubmissionListResponseOutcomeDispositionSpoof, SubmissionListResponseOutcomeDispositionSpam, SubmissionListResponseOutcomeDispositionBulk, SubmissionListResponseOutcomeDispositionEncrypted, SubmissionListResponseOutcomeDispositionExternal, SubmissionListResponseOutcomeDispositionUnknown, SubmissionListResponseOutcomeDispositionNone:
		return true
	}
	return false
}

type SubmissionListResponseRequestedDisposition string

const (
	SubmissionListResponseRequestedDispositionMalicious    SubmissionListResponseRequestedDisposition = "MALICIOUS"
	SubmissionListResponseRequestedDispositionMaliciousBec SubmissionListResponseRequestedDisposition = "MALICIOUS-BEC"
	SubmissionListResponseRequestedDispositionSuspicious   SubmissionListResponseRequestedDisposition = "SUSPICIOUS"
	SubmissionListResponseRequestedDispositionSpoof        SubmissionListResponseRequestedDisposition = "SPOOF"
	SubmissionListResponseRequestedDispositionSpam         SubmissionListResponseRequestedDisposition = "SPAM"
	SubmissionListResponseRequestedDispositionBulk         SubmissionListResponseRequestedDisposition = "BULK"
	SubmissionListResponseRequestedDispositionEncrypted    SubmissionListResponseRequestedDisposition = "ENCRYPTED"
	SubmissionListResponseRequestedDispositionExternal     SubmissionListResponseRequestedDisposition = "EXTERNAL"
	SubmissionListResponseRequestedDispositionUnknown      SubmissionListResponseRequestedDisposition = "UNKNOWN"
	SubmissionListResponseRequestedDispositionNone         SubmissionListResponseRequestedDisposition = "NONE"
)

func (r SubmissionListResponseRequestedDisposition) IsKnown() bool {
	switch r {
	case SubmissionListResponseRequestedDispositionMalicious, SubmissionListResponseRequestedDispositionMaliciousBec, SubmissionListResponseRequestedDispositionSuspicious, SubmissionListResponseRequestedDispositionSpoof, SubmissionListResponseRequestedDispositionSpam, SubmissionListResponseRequestedDispositionBulk, SubmissionListResponseRequestedDispositionEncrypted, SubmissionListResponseRequestedDispositionExternal, SubmissionListResponseRequestedDispositionUnknown, SubmissionListResponseRequestedDispositionNone:
		return true
	}
	return false
}

type SubmissionListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The end of the search date range. Defaults to `now`.
	End                 param.Field[time.Time]                               `query:"end" format:"date-time"`
	OriginalDisposition param.Field[SubmissionListParamsOriginalDisposition] `query:"original_disposition"`
	OutcomeDisposition  param.Field[SubmissionListParamsOutcomeDisposition]  `query:"outcome_disposition"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage              param.Field[int64]                                    `query:"per_page"`
	Query                param.Field[string]                                   `query:"query"`
	RequestedDisposition param.Field[SubmissionListParamsRequestedDisposition] `query:"requested_disposition"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start        param.Field[time.Time]                `query:"start" format:"date-time"`
	Status       param.Field[string]                   `query:"status"`
	SubmissionID param.Field[string]                   `query:"submission_id"`
	Type         param.Field[SubmissionListParamsType] `query:"type"`
}

// URLQuery serializes [SubmissionListParams]'s query parameters as `url.Values`.
func (r SubmissionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SubmissionListParamsOriginalDisposition string

const (
	SubmissionListParamsOriginalDispositionMalicious  SubmissionListParamsOriginalDisposition = "MALICIOUS"
	SubmissionListParamsOriginalDispositionSuspicious SubmissionListParamsOriginalDisposition = "SUSPICIOUS"
	SubmissionListParamsOriginalDispositionSpoof      SubmissionListParamsOriginalDisposition = "SPOOF"
	SubmissionListParamsOriginalDispositionSpam       SubmissionListParamsOriginalDisposition = "SPAM"
	SubmissionListParamsOriginalDispositionBulk       SubmissionListParamsOriginalDisposition = "BULK"
	SubmissionListParamsOriginalDispositionNone       SubmissionListParamsOriginalDisposition = "NONE"
)

func (r SubmissionListParamsOriginalDisposition) IsKnown() bool {
	switch r {
	case SubmissionListParamsOriginalDispositionMalicious, SubmissionListParamsOriginalDispositionSuspicious, SubmissionListParamsOriginalDispositionSpoof, SubmissionListParamsOriginalDispositionSpam, SubmissionListParamsOriginalDispositionBulk, SubmissionListParamsOriginalDispositionNone:
		return true
	}
	return false
}

type SubmissionListParamsOutcomeDisposition string

const (
	SubmissionListParamsOutcomeDispositionMalicious  SubmissionListParamsOutcomeDisposition = "MALICIOUS"
	SubmissionListParamsOutcomeDispositionSuspicious SubmissionListParamsOutcomeDisposition = "SUSPICIOUS"
	SubmissionListParamsOutcomeDispositionSpoof      SubmissionListParamsOutcomeDisposition = "SPOOF"
	SubmissionListParamsOutcomeDispositionSpam       SubmissionListParamsOutcomeDisposition = "SPAM"
	SubmissionListParamsOutcomeDispositionBulk       SubmissionListParamsOutcomeDisposition = "BULK"
	SubmissionListParamsOutcomeDispositionNone       SubmissionListParamsOutcomeDisposition = "NONE"
)

func (r SubmissionListParamsOutcomeDisposition) IsKnown() bool {
	switch r {
	case SubmissionListParamsOutcomeDispositionMalicious, SubmissionListParamsOutcomeDispositionSuspicious, SubmissionListParamsOutcomeDispositionSpoof, SubmissionListParamsOutcomeDispositionSpam, SubmissionListParamsOutcomeDispositionBulk, SubmissionListParamsOutcomeDispositionNone:
		return true
	}
	return false
}

type SubmissionListParamsRequestedDisposition string

const (
	SubmissionListParamsRequestedDispositionMalicious  SubmissionListParamsRequestedDisposition = "MALICIOUS"
	SubmissionListParamsRequestedDispositionSuspicious SubmissionListParamsRequestedDisposition = "SUSPICIOUS"
	SubmissionListParamsRequestedDispositionSpoof      SubmissionListParamsRequestedDisposition = "SPOOF"
	SubmissionListParamsRequestedDispositionSpam       SubmissionListParamsRequestedDisposition = "SPAM"
	SubmissionListParamsRequestedDispositionBulk       SubmissionListParamsRequestedDisposition = "BULK"
	SubmissionListParamsRequestedDispositionNone       SubmissionListParamsRequestedDisposition = "NONE"
)

func (r SubmissionListParamsRequestedDisposition) IsKnown() bool {
	switch r {
	case SubmissionListParamsRequestedDispositionMalicious, SubmissionListParamsRequestedDispositionSuspicious, SubmissionListParamsRequestedDispositionSpoof, SubmissionListParamsRequestedDispositionSpam, SubmissionListParamsRequestedDispositionBulk, SubmissionListParamsRequestedDispositionNone:
		return true
	}
	return false
}

type SubmissionListParamsType string

const (
	SubmissionListParamsTypeTeam SubmissionListParamsType = "TEAM"
	SubmissionListParamsTypeUser SubmissionListParamsType = "USER"
)

func (r SubmissionListParamsType) IsKnown() bool {
	switch r {
	case SubmissionListParamsTypeTeam, SubmissionListParamsTypeUser:
		return true
	}
	return false
}
