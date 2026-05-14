// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/intel"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// InsightService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightService] method instead.
type InsightService struct {
	Options  []option.RequestOption
	Class    *InsightClassService
	Severity *InsightSeverityService
	Type     *InsightTypeService
}

// NewInsightService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInsightService(opts ...option.RequestOption) (r *InsightService) {
	r = &InsightService{}
	r.Options = opts
	r.Class = NewInsightClassService(opts...)
	r.Severity = NewInsightSeverityService(opts...)
	r.Type = NewInsightTypeService(opts...)
	return
}

// Get Security Center Insights
func (r *InsightService) List(ctx context.Context, params InsightListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[InsightListResponse], err error) {
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
	path := fmt.Sprintf("%s/%s/security-center/insights", accountOrZone, accountOrZoneID)
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

// Get Security Center Insights
func (r *InsightService) ListAutoPaging(ctx context.Context, params InsightListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[InsightListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Archive Security Center Insight
func (r *InsightService) Dismiss(ctx context.Context, issueID string, params InsightDismissParams, opts ...option.RequestOption) (res *InsightDismissResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/security-center/insights/%s/dismiss", accountOrZone, accountOrZoneID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type InsightListResponse struct {
	// Total number of results
	Count  int64                      `json:"count"`
	Issues []InsightListResponseIssue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                   `json:"per_page"`
	JSON    insightListResponseJSON `json:"-"`
}

// insightListResponseJSON contains the JSON metadata for the struct
// [InsightListResponse]
type insightListResponseJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightListResponseJSON) RawJSON() string {
	return r.raw
}

type InsightListResponseIssue struct {
	ID          string                            `json:"id"`
	Dismissed   bool                              `json:"dismissed"`
	IssueClass  string                            `json:"issue_class"`
	IssueType   intel.IssueType                   `json:"issue_type"`
	Payload     interface{}                       `json:"payload"`
	ResolveLink string                            `json:"resolve_link"`
	ResolveText string                            `json:"resolve_text"`
	Severity    InsightListResponseIssuesSeverity `json:"severity"`
	Since       time.Time                         `json:"since" format:"date-time"`
	Subject     string                            `json:"subject"`
	Timestamp   time.Time                         `json:"timestamp" format:"date-time"`
	JSON        insightListResponseIssueJSON      `json:"-"`
}

// insightListResponseIssueJSON contains the JSON metadata for the struct
// [InsightListResponseIssue]
type insightListResponseIssueJSON struct {
	ID          apijson.Field
	Dismissed   apijson.Field
	IssueClass  apijson.Field
	IssueType   apijson.Field
	Payload     apijson.Field
	ResolveLink apijson.Field
	ResolveText apijson.Field
	Severity    apijson.Field
	Since       apijson.Field
	Subject     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightListResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightListResponseIssueJSON) RawJSON() string {
	return r.raw
}

type InsightListResponseIssuesSeverity string

const (
	InsightListResponseIssuesSeverityLow      InsightListResponseIssuesSeverity = "Low"
	InsightListResponseIssuesSeverityModerate InsightListResponseIssuesSeverity = "Moderate"
	InsightListResponseIssuesSeverityCritical InsightListResponseIssuesSeverity = "Critical"
)

func (r InsightListResponseIssuesSeverity) IsKnown() bool {
	switch r {
	case InsightListResponseIssuesSeverityLow, InsightListResponseIssuesSeverityModerate, InsightListResponseIssuesSeverityCritical:
		return true
	}
	return false
}

type InsightDismissResponse struct {
	Errors   []InsightDismissResponseError   `json:"errors,required"`
	Messages []InsightDismissResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success InsightDismissResponseSuccess `json:"success,required"`
	JSON    insightDismissResponseJSON    `json:"-"`
}

// insightDismissResponseJSON contains the JSON metadata for the struct
// [InsightDismissResponse]
type insightDismissResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightDismissResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightDismissResponseJSON) RawJSON() string {
	return r.raw
}

type InsightDismissResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           InsightDismissResponseErrorsSource `json:"source"`
	JSON             insightDismissResponseErrorJSON    `json:"-"`
}

// insightDismissResponseErrorJSON contains the JSON metadata for the struct
// [InsightDismissResponseError]
type insightDismissResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightDismissResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightDismissResponseErrorJSON) RawJSON() string {
	return r.raw
}

type InsightDismissResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    insightDismissResponseErrorsSourceJSON `json:"-"`
}

// insightDismissResponseErrorsSourceJSON contains the JSON metadata for the struct
// [InsightDismissResponseErrorsSource]
type insightDismissResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightDismissResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightDismissResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InsightDismissResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           InsightDismissResponseMessagesSource `json:"source"`
	JSON             insightDismissResponseMessageJSON    `json:"-"`
}

// insightDismissResponseMessageJSON contains the JSON metadata for the struct
// [InsightDismissResponseMessage]
type insightDismissResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightDismissResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightDismissResponseMessageJSON) RawJSON() string {
	return r.raw
}

type InsightDismissResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    insightDismissResponseMessagesSourceJSON `json:"-"`
}

// insightDismissResponseMessagesSourceJSON contains the JSON metadata for the
// struct [InsightDismissResponseMessagesSource]
type insightDismissResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightDismissResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightDismissResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InsightDismissResponseSuccess bool

const (
	InsightDismissResponseSuccessTrue InsightDismissResponseSuccess = true
)

func (r InsightDismissResponseSuccess) IsKnown() bool {
	switch r {
	case InsightDismissResponseSuccessTrue:
		return true
	}
	return false
}

type InsightListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID        param.Field[string]            `path:"zone_id"`
	Dismissed     param.Field[bool]              `query:"dismissed"`
	IssueClass    param.Field[[]string]          `query:"issue_class"`
	IssueClassNeq param.Field[[]string]          `query:"issue_class~neq"`
	IssueType     param.Field[[]intel.IssueType] `query:"issue_type"`
	IssueTypeNeq  param.Field[[]intel.IssueType] `query:"issue_type~neq"`
	// Current page within paginated list of results
	Page param.Field[int64] `query:"page"`
	// Number of results per page of results
	PerPage     param.Field[int64]                      `query:"per_page"`
	Product     param.Field[[]string]                   `query:"product"`
	ProductNeq  param.Field[[]string]                   `query:"product~neq"`
	Severity    param.Field[[]intel.SeverityQueryParam] `query:"severity"`
	SeverityNeq param.Field[[]intel.SeverityQueryParam] `query:"severity~neq"`
	Subject     param.Field[[]string]                   `query:"subject"`
	SubjectNeq  param.Field[[]string]                   `query:"subject~neq"`
}

// URLQuery serializes [InsightListParams]'s query parameters as `url.Values`.
func (r InsightListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InsightDismissParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID  param.Field[string] `path:"zone_id"`
	Dismiss param.Field[bool]   `json:"dismiss"`
}

func (r InsightDismissParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
