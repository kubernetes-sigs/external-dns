// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

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

// AttackSurfaceReportIssueService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackSurfaceReportIssueService] method instead.
type AttackSurfaceReportIssueService struct {
	Options []option.RequestOption
}

// NewAttackSurfaceReportIssueService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackSurfaceReportIssueService(opts ...option.RequestOption) (r *AttackSurfaceReportIssueService) {
	r = &AttackSurfaceReportIssueService{}
	r.Options = opts
	return
}

// Get Security Center Issues
//
// Deprecated: deprecated
func (r *AttackSurfaceReportIssueService) List(ctx context.Context, params AttackSurfaceReportIssueListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[AttackSurfaceReportIssueListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues", params.AccountID)
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

// Get Security Center Issues
//
// Deprecated: deprecated
func (r *AttackSurfaceReportIssueService) ListAutoPaging(ctx context.Context, params AttackSurfaceReportIssueListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[AttackSurfaceReportIssueListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Get Security Center Issue Counts by Class
//
// Deprecated: deprecated
func (r *AttackSurfaceReportIssueService) Class(ctx context.Context, params AttackSurfaceReportIssueClassParams, opts ...option.RequestOption) (res *[]AttackSurfaceReportIssueClassResponse, err error) {
	var env AttackSurfaceReportIssueClassResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/class", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Archive Security Center Insight
//
// Deprecated: deprecated
func (r *AttackSurfaceReportIssueService) Dismiss(ctx context.Context, issueID string, params AttackSurfaceReportIssueDismissParams, opts ...option.RequestOption) (res *AttackSurfaceReportIssueDismissResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/%s/dismiss", params.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get Security Center Issue Counts by Severity
//
// Deprecated: deprecated
func (r *AttackSurfaceReportIssueService) Severity(ctx context.Context, params AttackSurfaceReportIssueSeverityParams, opts ...option.RequestOption) (res *[]AttackSurfaceReportIssueSeverityResponse, err error) {
	var env AttackSurfaceReportIssueSeverityResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/severity", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Security Center Issue Counts by Type
//
// Deprecated: deprecated
func (r *AttackSurfaceReportIssueService) Type(ctx context.Context, params AttackSurfaceReportIssueTypeParams, opts ...option.RequestOption) (res *[]AttackSurfaceReportIssueTypeResponse, err error) {
	var env AttackSurfaceReportIssueTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/type", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IssueType string

const (
	IssueTypeComplianceViolation   IssueType = "compliance_violation"
	IssueTypeEmailSecurity         IssueType = "email_security"
	IssueTypeExposedInfrastructure IssueType = "exposed_infrastructure"
	IssueTypeInsecureConfiguration IssueType = "insecure_configuration"
	IssueTypeWeakAuthentication    IssueType = "weak_authentication"
)

func (r IssueType) IsKnown() bool {
	switch r {
	case IssueTypeComplianceViolation, IssueTypeEmailSecurity, IssueTypeExposedInfrastructure, IssueTypeInsecureConfiguration, IssueTypeWeakAuthentication:
		return true
	}
	return false
}

type SeverityQueryParam string

const (
	SeverityQueryParamLow      SeverityQueryParam = "low"
	SeverityQueryParamModerate SeverityQueryParam = "moderate"
	SeverityQueryParamCritical SeverityQueryParam = "critical"
)

func (r SeverityQueryParam) IsKnown() bool {
	switch r {
	case SeverityQueryParamLow, SeverityQueryParamModerate, SeverityQueryParamCritical:
		return true
	}
	return false
}

type AttackSurfaceReportIssueListResponse struct {
	// Total number of results
	Count  int64                                       `json:"count"`
	Issues []AttackSurfaceReportIssueListResponseIssue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                                    `json:"per_page"`
	JSON    attackSurfaceReportIssueListResponseJSON `json:"-"`
}

// attackSurfaceReportIssueListResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueListResponse]
type attackSurfaceReportIssueListResponseJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueListResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListResponseIssue struct {
	ID          string                                             `json:"id"`
	Dismissed   bool                                               `json:"dismissed"`
	IssueClass  string                                             `json:"issue_class"`
	IssueType   IssueType                                          `json:"issue_type"`
	Payload     interface{}                                        `json:"payload"`
	ResolveLink string                                             `json:"resolve_link"`
	ResolveText string                                             `json:"resolve_text"`
	Severity    AttackSurfaceReportIssueListResponseIssuesSeverity `json:"severity"`
	Since       time.Time                                          `json:"since" format:"date-time"`
	Subject     string                                             `json:"subject"`
	Timestamp   time.Time                                          `json:"timestamp" format:"date-time"`
	JSON        attackSurfaceReportIssueListResponseIssueJSON      `json:"-"`
}

// attackSurfaceReportIssueListResponseIssueJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueListResponseIssue]
type attackSurfaceReportIssueListResponseIssueJSON struct {
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

func (r *AttackSurfaceReportIssueListResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueListResponseIssueJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListResponseIssuesSeverity string

const (
	AttackSurfaceReportIssueListResponseIssuesSeverityLow      AttackSurfaceReportIssueListResponseIssuesSeverity = "Low"
	AttackSurfaceReportIssueListResponseIssuesSeverityModerate AttackSurfaceReportIssueListResponseIssuesSeverity = "Moderate"
	AttackSurfaceReportIssueListResponseIssuesSeverityCritical AttackSurfaceReportIssueListResponseIssuesSeverity = "Critical"
)

func (r AttackSurfaceReportIssueListResponseIssuesSeverity) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueListResponseIssuesSeverityLow, AttackSurfaceReportIssueListResponseIssuesSeverityModerate, AttackSurfaceReportIssueListResponseIssuesSeverityCritical:
		return true
	}
	return false
}

type AttackSurfaceReportIssueClassResponse struct {
	Count int64                                     `json:"count"`
	Value string                                    `json:"value"`
	JSON  attackSurfaceReportIssueClassResponseJSON `json:"-"`
}

// attackSurfaceReportIssueClassResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueClassResponse]
type attackSurfaceReportIssueClassResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueDismissResponse struct {
	Errors   []AttackSurfaceReportIssueDismissResponseError   `json:"errors,required"`
	Messages []AttackSurfaceReportIssueDismissResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AttackSurfaceReportIssueDismissResponseSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueDismissResponseJSON    `json:"-"`
}

// attackSurfaceReportIssueDismissResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueDismissResponse]
type attackSurfaceReportIssueDismissResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueDismissResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueDismissResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueDismissResponseError struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AttackSurfaceReportIssueDismissResponseErrorsSource `json:"source"`
	JSON             attackSurfaceReportIssueDismissResponseErrorJSON    `json:"-"`
}

// attackSurfaceReportIssueDismissResponseErrorJSON contains the JSON metadata for
// the struct [AttackSurfaceReportIssueDismissResponseError]
type attackSurfaceReportIssueDismissResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueDismissResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueDismissResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueDismissResponseErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    attackSurfaceReportIssueDismissResponseErrorsSourceJSON `json:"-"`
}

// attackSurfaceReportIssueDismissResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueDismissResponseErrorsSource]
type attackSurfaceReportIssueDismissResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueDismissResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueDismissResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueDismissResponseMessage struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AttackSurfaceReportIssueDismissResponseMessagesSource `json:"source"`
	JSON             attackSurfaceReportIssueDismissResponseMessageJSON    `json:"-"`
}

// attackSurfaceReportIssueDismissResponseMessageJSON contains the JSON metadata
// for the struct [AttackSurfaceReportIssueDismissResponseMessage]
type attackSurfaceReportIssueDismissResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueDismissResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueDismissResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueDismissResponseMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    attackSurfaceReportIssueDismissResponseMessagesSourceJSON `json:"-"`
}

// attackSurfaceReportIssueDismissResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueDismissResponseMessagesSource]
type attackSurfaceReportIssueDismissResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueDismissResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueDismissResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AttackSurfaceReportIssueDismissResponseSuccess bool

const (
	AttackSurfaceReportIssueDismissResponseSuccessTrue AttackSurfaceReportIssueDismissResponseSuccess = true
)

func (r AttackSurfaceReportIssueDismissResponseSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueDismissResponseSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueSeverityResponse struct {
	Count int64                                        `json:"count"`
	Value string                                       `json:"value"`
	JSON  attackSurfaceReportIssueSeverityResponseJSON `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueSeverityResponse]
type attackSurfaceReportIssueSeverityResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueTypeResponse struct {
	Count int64                                    `json:"count"`
	Value string                                   `json:"value"`
	JSON  attackSurfaceReportIssueTypeResponseJSON `json:"-"`
}

// attackSurfaceReportIssueTypeResponseJSON contains the JSON metadata for the
// struct [AttackSurfaceReportIssueTypeResponse]
type attackSurfaceReportIssueTypeResponseJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueListParams struct {
	// Identifier.
	AccountID     param.Field[string]      `path:"account_id,required"`
	Dismissed     param.Field[bool]        `query:"dismissed"`
	IssueClass    param.Field[[]string]    `query:"issue_class"`
	IssueClassNeq param.Field[[]string]    `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType] `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType] `query:"issue_type~neq"`
	// Current page within paginated list of results
	Page param.Field[int64] `query:"page"`
	// Number of results per page of results
	PerPage     param.Field[int64]                `query:"per_page"`
	Product     param.Field[[]string]             `query:"product"`
	ProductNeq  param.Field[[]string]             `query:"product~neq"`
	Severity    param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject     param.Field[[]string]             `query:"subject"`
	SubjectNeq  param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueListParams]'s query parameters as
// `url.Values`.
func (r AttackSurfaceReportIssueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AttackSurfaceReportIssueClassParams struct {
	// Identifier.
	AccountID     param.Field[string]               `path:"account_id,required"`
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueClassParams]'s query parameters as
// `url.Values`.
func (r AttackSurfaceReportIssueClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AttackSurfaceReportIssueClassResponseEnvelope struct {
	Errors   []AttackSurfaceReportIssueClassResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AttackSurfaceReportIssueClassResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AttackSurfaceReportIssueClassResponseEnvelopeSuccess `json:"success,required"`
	Result  []AttackSurfaceReportIssueClassResponse              `json:"result"`
	JSON    attackSurfaceReportIssueClassResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueClassResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackSurfaceReportIssueClassResponseEnvelope]
type attackSurfaceReportIssueClassResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueClassResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AttackSurfaceReportIssueClassResponseEnvelopeErrorsSource `json:"source"`
	JSON             attackSurfaceReportIssueClassResponseEnvelopeErrorsJSON   `json:"-"`
}

// attackSurfaceReportIssueClassResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueClassResponseEnvelopeErrors]
type attackSurfaceReportIssueClassResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueClassResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    attackSurfaceReportIssueClassResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// attackSurfaceReportIssueClassResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AttackSurfaceReportIssueClassResponseEnvelopeErrorsSource]
type attackSurfaceReportIssueClassResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueClassResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AttackSurfaceReportIssueClassResponseEnvelopeMessagesSource `json:"source"`
	JSON             attackSurfaceReportIssueClassResponseEnvelopeMessagesJSON   `json:"-"`
}

// attackSurfaceReportIssueClassResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueClassResponseEnvelopeMessages]
type attackSurfaceReportIssueClassResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueClassResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    attackSurfaceReportIssueClassResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// attackSurfaceReportIssueClassResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AttackSurfaceReportIssueClassResponseEnvelopeMessagesSource]
type attackSurfaceReportIssueClassResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueClassResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueClassResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AttackSurfaceReportIssueClassResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueClassResponseEnvelopeSuccessTrue AttackSurfaceReportIssueClassResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueClassResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueClassResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueDismissParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Dismiss   param.Field[bool]   `json:"dismiss"`
}

func (r AttackSurfaceReportIssueDismissParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AttackSurfaceReportIssueSeverityParams struct {
	// Identifier.
	AccountID     param.Field[string]               `path:"account_id,required"`
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueSeverityParams]'s query parameters
// as `url.Values`.
func (r AttackSurfaceReportIssueSeverityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AttackSurfaceReportIssueSeverityResponseEnvelope struct {
	Errors   []AttackSurfaceReportIssueSeverityResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AttackSurfaceReportIssueSeverityResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess `json:"success,required"`
	Result  []AttackSurfaceReportIssueSeverityResponse              `json:"result"`
	JSON    attackSurfaceReportIssueSeverityResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackSurfaceReportIssueSeverityResponseEnvelope]
type attackSurfaceReportIssueSeverityResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueSeverityResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AttackSurfaceReportIssueSeverityResponseEnvelopeErrorsSource `json:"source"`
	JSON             attackSurfaceReportIssueSeverityResponseEnvelopeErrorsJSON   `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueSeverityResponseEnvelopeErrors]
type attackSurfaceReportIssueSeverityResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueSeverityResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    attackSurfaceReportIssueSeverityResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AttackSurfaceReportIssueSeverityResponseEnvelopeErrorsSource]
type attackSurfaceReportIssueSeverityResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueSeverityResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AttackSurfaceReportIssueSeverityResponseEnvelopeMessagesSource `json:"source"`
	JSON             attackSurfaceReportIssueSeverityResponseEnvelopeMessagesJSON   `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AttackSurfaceReportIssueSeverityResponseEnvelopeMessages]
type attackSurfaceReportIssueSeverityResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueSeverityResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    attackSurfaceReportIssueSeverityResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// attackSurfaceReportIssueSeverityResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AttackSurfaceReportIssueSeverityResponseEnvelopeMessagesSource]
type attackSurfaceReportIssueSeverityResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueSeverityResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueSeverityResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueSeverityResponseEnvelopeSuccessTrue AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueSeverityResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueSeverityResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AttackSurfaceReportIssueTypeParams struct {
	// Identifier.
	AccountID     param.Field[string]               `path:"account_id,required"`
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AttackSurfaceReportIssueTypeParams]'s query parameters as
// `url.Values`.
func (r AttackSurfaceReportIssueTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AttackSurfaceReportIssueTypeResponseEnvelope struct {
	Errors   []AttackSurfaceReportIssueTypeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AttackSurfaceReportIssueTypeResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AttackSurfaceReportIssueTypeResponseEnvelopeSuccess `json:"success,required"`
	Result  []AttackSurfaceReportIssueTypeResponse              `json:"result"`
	JSON    attackSurfaceReportIssueTypeResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackSurfaceReportIssueTypeResponseEnvelope]
type attackSurfaceReportIssueTypeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueTypeResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AttackSurfaceReportIssueTypeResponseEnvelopeErrorsSource `json:"source"`
	JSON             attackSurfaceReportIssueTypeResponseEnvelopeErrorsJSON   `json:"-"`
}

// attackSurfaceReportIssueTypeResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueTypeResponseEnvelopeErrors]
type attackSurfaceReportIssueTypeResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueTypeResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    attackSurfaceReportIssueTypeResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// attackSurfaceReportIssueTypeResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AttackSurfaceReportIssueTypeResponseEnvelopeErrorsSource]
type attackSurfaceReportIssueTypeResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueTypeResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AttackSurfaceReportIssueTypeResponseEnvelopeMessagesSource `json:"source"`
	JSON             attackSurfaceReportIssueTypeResponseEnvelopeMessagesJSON   `json:"-"`
}

// attackSurfaceReportIssueTypeResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AttackSurfaceReportIssueTypeResponseEnvelopeMessages]
type attackSurfaceReportIssueTypeResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AttackSurfaceReportIssueTypeResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    attackSurfaceReportIssueTypeResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// attackSurfaceReportIssueTypeResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AttackSurfaceReportIssueTypeResponseEnvelopeMessagesSource]
type attackSurfaceReportIssueTypeResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AttackSurfaceReportIssueTypeResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueTypeResponseEnvelopeSuccessTrue AttackSurfaceReportIssueTypeResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueTypeResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueTypeResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
