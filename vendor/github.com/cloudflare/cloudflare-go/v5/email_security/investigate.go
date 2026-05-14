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

// InvestigateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateService] method instead.
type InvestigateService struct {
	Options    []option.RequestOption
	Detections *InvestigateDetectionService
	Preview    *InvestigatePreviewService
	Raw        *InvestigateRawService
	Trace      *InvestigateTraceService
	Move       *InvestigateMoveService
	Reclassify *InvestigateReclassifyService
	Release    *InvestigateReleaseService
}

// NewInvestigateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateService(opts ...option.RequestOption) (r *InvestigateService) {
	r = &InvestigateService{}
	r.Options = opts
	r.Detections = NewInvestigateDetectionService(opts...)
	r.Preview = NewInvestigatePreviewService(opts...)
	r.Raw = NewInvestigateRawService(opts...)
	r.Trace = NewInvestigateTraceService(opts...)
	r.Move = NewInvestigateMoveService(opts...)
	r.Reclassify = NewInvestigateReclassifyService(opts...)
	r.Release = NewInvestigateReleaseService(opts...)
	return
}

// Returns information for each email that matches the search parameter(s).
func (r *InvestigateService) List(ctx context.Context, params InvestigateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InvestigateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate", params.AccountID)
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

// Returns information for each email that matches the search parameter(s).
func (r *InvestigateService) ListAutoPaging(ctx context.Context, params InvestigateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InvestigateListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Get message details
func (r *InvestigateService) Get(ctx context.Context, postfixID string, query InvestigateGetParams, opts ...option.RequestOption) (res *InvestigateGetResponse, err error) {
	var env InvestigateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateListResponse struct {
	ID                string      `json:"id,required"`
	ActionLog         interface{} `json:"action_log,required"`
	ClientRecipients  []string    `json:"client_recipients,required"`
	DetectionReasons  []string    `json:"detection_reasons,required"`
	IsPhishSubmission bool        `json:"is_phish_submission,required"`
	IsQuarantined     bool        `json:"is_quarantined,required"`
	// The identifier of the message.
	PostfixID        string                                  `json:"postfix_id,required"`
	Properties       InvestigateListResponseProperties       `json:"properties,required"`
	Ts               string                                  `json:"ts,required"`
	AlertID          string                                  `json:"alert_id,nullable"`
	DeliveryMode     InvestigateListResponseDeliveryMode     `json:"delivery_mode,nullable"`
	EdfHash          string                                  `json:"edf_hash,nullable"`
	FinalDisposition InvestigateListResponseFinalDisposition `json:"final_disposition,nullable"`
	Findings         []InvestigateListResponseFinding        `json:"findings,nullable"`
	From             string                                  `json:"from,nullable"`
	FromName         string                                  `json:"from_name,nullable"`
	MessageID        string                                  `json:"message_id,nullable"`
	SentDate         string                                  `json:"sent_date,nullable"`
	Subject          string                                  `json:"subject,nullable"`
	ThreatCategories []string                                `json:"threat_categories,nullable"`
	To               []string                                `json:"to,nullable"`
	ToName           []string                                `json:"to_name,nullable"`
	Validation       InvestigateListResponseValidation       `json:"validation,nullable"`
	JSON             investigateListResponseJSON             `json:"-"`
}

// investigateListResponseJSON contains the JSON metadata for the struct
// [InvestigateListResponse]
type investigateListResponseJSON struct {
	ID                apijson.Field
	ActionLog         apijson.Field
	ClientRecipients  apijson.Field
	DetectionReasons  apijson.Field
	IsPhishSubmission apijson.Field
	IsQuarantined     apijson.Field
	PostfixID         apijson.Field
	Properties        apijson.Field
	Ts                apijson.Field
	AlertID           apijson.Field
	DeliveryMode      apijson.Field
	EdfHash           apijson.Field
	FinalDisposition  apijson.Field
	Findings          apijson.Field
	From              apijson.Field
	FromName          apijson.Field
	MessageID         apijson.Field
	SentDate          apijson.Field
	Subject           apijson.Field
	ThreatCategories  apijson.Field
	To                apijson.Field
	ToName            apijson.Field
	Validation        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvestigateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponseProperties struct {
	AllowlistedPattern     string                                                  `json:"allowlisted_pattern"`
	AllowlistedPatternType InvestigateListResponsePropertiesAllowlistedPatternType `json:"allowlisted_pattern_type"`
	BlocklistedMessage     bool                                                    `json:"blocklisted_message"`
	BlocklistedPattern     string                                                  `json:"blocklisted_pattern"`
	WhitelistedPatternType InvestigateListResponsePropertiesWhitelistedPatternType `json:"whitelisted_pattern_type"`
	JSON                   investigateListResponsePropertiesJSON                   `json:"-"`
}

// investigateListResponsePropertiesJSON contains the JSON metadata for the struct
// [InvestigateListResponseProperties]
type investigateListResponsePropertiesJSON struct {
	AllowlistedPattern     apijson.Field
	AllowlistedPatternType apijson.Field
	BlocklistedMessage     apijson.Field
	BlocklistedPattern     apijson.Field
	WhitelistedPatternType apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateListResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponsePropertiesAllowlistedPatternType string

const (
	InvestigateListResponsePropertiesAllowlistedPatternTypeQuarantineRelease       InvestigateListResponsePropertiesAllowlistedPatternType = "quarantine_release"
	InvestigateListResponsePropertiesAllowlistedPatternTypeAcceptableSender        InvestigateListResponsePropertiesAllowlistedPatternType = "acceptable_sender"
	InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedSender           InvestigateListResponsePropertiesAllowlistedPatternType = "allowed_sender"
	InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedRecipient        InvestigateListResponsePropertiesAllowlistedPatternType = "allowed_recipient"
	InvestigateListResponsePropertiesAllowlistedPatternTypeDomainSimilarity        InvestigateListResponsePropertiesAllowlistedPatternType = "domain_similarity"
	InvestigateListResponsePropertiesAllowlistedPatternTypeDomainRecency           InvestigateListResponsePropertiesAllowlistedPatternType = "domain_recency"
	InvestigateListResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender InvestigateListResponsePropertiesAllowlistedPatternType = "managed_acceptable_sender"
)

func (r InvestigateListResponsePropertiesAllowlistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateListResponsePropertiesAllowlistedPatternTypeQuarantineRelease, InvestigateListResponsePropertiesAllowlistedPatternTypeAcceptableSender, InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedSender, InvestigateListResponsePropertiesAllowlistedPatternTypeAllowedRecipient, InvestigateListResponsePropertiesAllowlistedPatternTypeDomainSimilarity, InvestigateListResponsePropertiesAllowlistedPatternTypeDomainRecency, InvestigateListResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender:
		return true
	}
	return false
}

type InvestigateListResponsePropertiesWhitelistedPatternType string

const (
	InvestigateListResponsePropertiesWhitelistedPatternTypeQuarantineRelease       InvestigateListResponsePropertiesWhitelistedPatternType = "quarantine_release"
	InvestigateListResponsePropertiesWhitelistedPatternTypeAcceptableSender        InvestigateListResponsePropertiesWhitelistedPatternType = "acceptable_sender"
	InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedSender           InvestigateListResponsePropertiesWhitelistedPatternType = "allowed_sender"
	InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedRecipient        InvestigateListResponsePropertiesWhitelistedPatternType = "allowed_recipient"
	InvestigateListResponsePropertiesWhitelistedPatternTypeDomainSimilarity        InvestigateListResponsePropertiesWhitelistedPatternType = "domain_similarity"
	InvestigateListResponsePropertiesWhitelistedPatternTypeDomainRecency           InvestigateListResponsePropertiesWhitelistedPatternType = "domain_recency"
	InvestigateListResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender InvestigateListResponsePropertiesWhitelistedPatternType = "managed_acceptable_sender"
)

func (r InvestigateListResponsePropertiesWhitelistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateListResponsePropertiesWhitelistedPatternTypeQuarantineRelease, InvestigateListResponsePropertiesWhitelistedPatternTypeAcceptableSender, InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedSender, InvestigateListResponsePropertiesWhitelistedPatternTypeAllowedRecipient, InvestigateListResponsePropertiesWhitelistedPatternTypeDomainSimilarity, InvestigateListResponsePropertiesWhitelistedPatternTypeDomainRecency, InvestigateListResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender:
		return true
	}
	return false
}

type InvestigateListResponseDeliveryMode string

const (
	InvestigateListResponseDeliveryModeDirect                InvestigateListResponseDeliveryMode = "DIRECT"
	InvestigateListResponseDeliveryModeBcc                   InvestigateListResponseDeliveryMode = "BCC"
	InvestigateListResponseDeliveryModeJournal               InvestigateListResponseDeliveryMode = "JOURNAL"
	InvestigateListResponseDeliveryModeReviewSubmission      InvestigateListResponseDeliveryMode = "REVIEW_SUBMISSION"
	InvestigateListResponseDeliveryModeDMARCUnverified       InvestigateListResponseDeliveryMode = "DMARC_UNVERIFIED"
	InvestigateListResponseDeliveryModeDMARCFailureReport    InvestigateListResponseDeliveryMode = "DMARC_FAILURE_REPORT"
	InvestigateListResponseDeliveryModeDMARCAggregateReport  InvestigateListResponseDeliveryMode = "DMARC_AGGREGATE_REPORT"
	InvestigateListResponseDeliveryModeThreatIntelSubmission InvestigateListResponseDeliveryMode = "THREAT_INTEL_SUBMISSION"
	InvestigateListResponseDeliveryModeSimulationSubmission  InvestigateListResponseDeliveryMode = "SIMULATION_SUBMISSION"
	InvestigateListResponseDeliveryModeAPI                   InvestigateListResponseDeliveryMode = "API"
	InvestigateListResponseDeliveryModeRetroScan             InvestigateListResponseDeliveryMode = "RETRO_SCAN"
)

func (r InvestigateListResponseDeliveryMode) IsKnown() bool {
	switch r {
	case InvestigateListResponseDeliveryModeDirect, InvestigateListResponseDeliveryModeBcc, InvestigateListResponseDeliveryModeJournal, InvestigateListResponseDeliveryModeReviewSubmission, InvestigateListResponseDeliveryModeDMARCUnverified, InvestigateListResponseDeliveryModeDMARCFailureReport, InvestigateListResponseDeliveryModeDMARCAggregateReport, InvestigateListResponseDeliveryModeThreatIntelSubmission, InvestigateListResponseDeliveryModeSimulationSubmission, InvestigateListResponseDeliveryModeAPI, InvestigateListResponseDeliveryModeRetroScan:
		return true
	}
	return false
}

type InvestigateListResponseFinalDisposition string

const (
	InvestigateListResponseFinalDispositionMalicious    InvestigateListResponseFinalDisposition = "MALICIOUS"
	InvestigateListResponseFinalDispositionMaliciousBec InvestigateListResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateListResponseFinalDispositionSuspicious   InvestigateListResponseFinalDisposition = "SUSPICIOUS"
	InvestigateListResponseFinalDispositionSpoof        InvestigateListResponseFinalDisposition = "SPOOF"
	InvestigateListResponseFinalDispositionSpam         InvestigateListResponseFinalDisposition = "SPAM"
	InvestigateListResponseFinalDispositionBulk         InvestigateListResponseFinalDisposition = "BULK"
	InvestigateListResponseFinalDispositionEncrypted    InvestigateListResponseFinalDisposition = "ENCRYPTED"
	InvestigateListResponseFinalDispositionExternal     InvestigateListResponseFinalDisposition = "EXTERNAL"
	InvestigateListResponseFinalDispositionUnknown      InvestigateListResponseFinalDisposition = "UNKNOWN"
	InvestigateListResponseFinalDispositionNone         InvestigateListResponseFinalDisposition = "NONE"
)

func (r InvestigateListResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateListResponseFinalDispositionMalicious, InvestigateListResponseFinalDispositionMaliciousBec, InvestigateListResponseFinalDispositionSuspicious, InvestigateListResponseFinalDispositionSpoof, InvestigateListResponseFinalDispositionSpam, InvestigateListResponseFinalDispositionBulk, InvestigateListResponseFinalDispositionEncrypted, InvestigateListResponseFinalDispositionExternal, InvestigateListResponseFinalDispositionUnknown, InvestigateListResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateListResponseFinding struct {
	Detail string                             `json:"detail,nullable"`
	Name   string                             `json:"name,nullable"`
	Value  string                             `json:"value,nullable"`
	JSON   investigateListResponseFindingJSON `json:"-"`
}

// investigateListResponseFindingJSON contains the JSON metadata for the struct
// [InvestigateListResponseFinding]
type investigateListResponseFindingJSON struct {
	Detail      apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateListResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseFindingJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponseValidation struct {
	Comment string                                 `json:"comment,nullable"`
	DKIM    InvestigateListResponseValidationDKIM  `json:"dkim,nullable"`
	DMARC   InvestigateListResponseValidationDMARC `json:"dmarc,nullable"`
	SPF     InvestigateListResponseValidationSPF   `json:"spf,nullable"`
	JSON    investigateListResponseValidationJSON  `json:"-"`
}

// investigateListResponseValidationJSON contains the JSON metadata for the struct
// [InvestigateListResponseValidation]
type investigateListResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateListResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateListResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateListResponseValidationDKIM string

const (
	InvestigateListResponseValidationDKIMPass    InvestigateListResponseValidationDKIM = "pass"
	InvestigateListResponseValidationDKIMNeutral InvestigateListResponseValidationDKIM = "neutral"
	InvestigateListResponseValidationDKIMFail    InvestigateListResponseValidationDKIM = "fail"
	InvestigateListResponseValidationDKIMError   InvestigateListResponseValidationDKIM = "error"
	InvestigateListResponseValidationDKIMNone    InvestigateListResponseValidationDKIM = "none"
)

func (r InvestigateListResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateListResponseValidationDKIMPass, InvestigateListResponseValidationDKIMNeutral, InvestigateListResponseValidationDKIMFail, InvestigateListResponseValidationDKIMError, InvestigateListResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateListResponseValidationDMARC string

const (
	InvestigateListResponseValidationDMARCPass    InvestigateListResponseValidationDMARC = "pass"
	InvestigateListResponseValidationDMARCNeutral InvestigateListResponseValidationDMARC = "neutral"
	InvestigateListResponseValidationDMARCFail    InvestigateListResponseValidationDMARC = "fail"
	InvestigateListResponseValidationDMARCError   InvestigateListResponseValidationDMARC = "error"
	InvestigateListResponseValidationDMARCNone    InvestigateListResponseValidationDMARC = "none"
)

func (r InvestigateListResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateListResponseValidationDMARCPass, InvestigateListResponseValidationDMARCNeutral, InvestigateListResponseValidationDMARCFail, InvestigateListResponseValidationDMARCError, InvestigateListResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateListResponseValidationSPF string

const (
	InvestigateListResponseValidationSPFPass    InvestigateListResponseValidationSPF = "pass"
	InvestigateListResponseValidationSPFNeutral InvestigateListResponseValidationSPF = "neutral"
	InvestigateListResponseValidationSPFFail    InvestigateListResponseValidationSPF = "fail"
	InvestigateListResponseValidationSPFError   InvestigateListResponseValidationSPF = "error"
	InvestigateListResponseValidationSPFNone    InvestigateListResponseValidationSPF = "none"
)

func (r InvestigateListResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateListResponseValidationSPFPass, InvestigateListResponseValidationSPFNeutral, InvestigateListResponseValidationSPFFail, InvestigateListResponseValidationSPFError, InvestigateListResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateGetResponse struct {
	ID                string      `json:"id,required"`
	ActionLog         interface{} `json:"action_log,required"`
	ClientRecipients  []string    `json:"client_recipients,required"`
	DetectionReasons  []string    `json:"detection_reasons,required"`
	IsPhishSubmission bool        `json:"is_phish_submission,required"`
	IsQuarantined     bool        `json:"is_quarantined,required"`
	// The identifier of the message.
	PostfixID        string                                 `json:"postfix_id,required"`
	Properties       InvestigateGetResponseProperties       `json:"properties,required"`
	Ts               string                                 `json:"ts,required"`
	AlertID          string                                 `json:"alert_id,nullable"`
	DeliveryMode     InvestigateGetResponseDeliveryMode     `json:"delivery_mode,nullable"`
	EdfHash          string                                 `json:"edf_hash,nullable"`
	FinalDisposition InvestigateGetResponseFinalDisposition `json:"final_disposition,nullable"`
	Findings         []InvestigateGetResponseFinding        `json:"findings,nullable"`
	From             string                                 `json:"from,nullable"`
	FromName         string                                 `json:"from_name,nullable"`
	MessageID        string                                 `json:"message_id,nullable"`
	SentDate         string                                 `json:"sent_date,nullable"`
	Subject          string                                 `json:"subject,nullable"`
	ThreatCategories []string                               `json:"threat_categories,nullable"`
	To               []string                               `json:"to,nullable"`
	ToName           []string                               `json:"to_name,nullable"`
	Validation       InvestigateGetResponseValidation       `json:"validation,nullable"`
	JSON             investigateGetResponseJSON             `json:"-"`
}

// investigateGetResponseJSON contains the JSON metadata for the struct
// [InvestigateGetResponse]
type investigateGetResponseJSON struct {
	ID                apijson.Field
	ActionLog         apijson.Field
	ClientRecipients  apijson.Field
	DetectionReasons  apijson.Field
	IsPhishSubmission apijson.Field
	IsQuarantined     apijson.Field
	PostfixID         apijson.Field
	Properties        apijson.Field
	Ts                apijson.Field
	AlertID           apijson.Field
	DeliveryMode      apijson.Field
	EdfHash           apijson.Field
	FinalDisposition  apijson.Field
	Findings          apijson.Field
	From              apijson.Field
	FromName          apijson.Field
	MessageID         apijson.Field
	SentDate          apijson.Field
	Subject           apijson.Field
	ThreatCategories  apijson.Field
	To                apijson.Field
	ToName            apijson.Field
	Validation        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvestigateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseProperties struct {
	AllowlistedPattern     string                                                 `json:"allowlisted_pattern"`
	AllowlistedPatternType InvestigateGetResponsePropertiesAllowlistedPatternType `json:"allowlisted_pattern_type"`
	BlocklistedMessage     bool                                                   `json:"blocklisted_message"`
	BlocklistedPattern     string                                                 `json:"blocklisted_pattern"`
	WhitelistedPatternType InvestigateGetResponsePropertiesWhitelistedPatternType `json:"whitelisted_pattern_type"`
	JSON                   investigateGetResponsePropertiesJSON                   `json:"-"`
}

// investigateGetResponsePropertiesJSON contains the JSON metadata for the struct
// [InvestigateGetResponseProperties]
type investigateGetResponsePropertiesJSON struct {
	AllowlistedPattern     apijson.Field
	AllowlistedPatternType apijson.Field
	BlocklistedMessage     apijson.Field
	BlocklistedPattern     apijson.Field
	WhitelistedPatternType apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *InvestigateGetResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponsePropertiesAllowlistedPatternType string

const (
	InvestigateGetResponsePropertiesAllowlistedPatternTypeQuarantineRelease       InvestigateGetResponsePropertiesAllowlistedPatternType = "quarantine_release"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeAcceptableSender        InvestigateGetResponsePropertiesAllowlistedPatternType = "acceptable_sender"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedSender           InvestigateGetResponsePropertiesAllowlistedPatternType = "allowed_sender"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedRecipient        InvestigateGetResponsePropertiesAllowlistedPatternType = "allowed_recipient"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainSimilarity        InvestigateGetResponsePropertiesAllowlistedPatternType = "domain_similarity"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainRecency           InvestigateGetResponsePropertiesAllowlistedPatternType = "domain_recency"
	InvestigateGetResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender InvestigateGetResponsePropertiesAllowlistedPatternType = "managed_acceptable_sender"
)

func (r InvestigateGetResponsePropertiesAllowlistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateGetResponsePropertiesAllowlistedPatternTypeQuarantineRelease, InvestigateGetResponsePropertiesAllowlistedPatternTypeAcceptableSender, InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedSender, InvestigateGetResponsePropertiesAllowlistedPatternTypeAllowedRecipient, InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainSimilarity, InvestigateGetResponsePropertiesAllowlistedPatternTypeDomainRecency, InvestigateGetResponsePropertiesAllowlistedPatternTypeManagedAcceptableSender:
		return true
	}
	return false
}

type InvestigateGetResponsePropertiesWhitelistedPatternType string

const (
	InvestigateGetResponsePropertiesWhitelistedPatternTypeQuarantineRelease       InvestigateGetResponsePropertiesWhitelistedPatternType = "quarantine_release"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeAcceptableSender        InvestigateGetResponsePropertiesWhitelistedPatternType = "acceptable_sender"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedSender           InvestigateGetResponsePropertiesWhitelistedPatternType = "allowed_sender"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedRecipient        InvestigateGetResponsePropertiesWhitelistedPatternType = "allowed_recipient"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainSimilarity        InvestigateGetResponsePropertiesWhitelistedPatternType = "domain_similarity"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainRecency           InvestigateGetResponsePropertiesWhitelistedPatternType = "domain_recency"
	InvestigateGetResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender InvestigateGetResponsePropertiesWhitelistedPatternType = "managed_acceptable_sender"
)

func (r InvestigateGetResponsePropertiesWhitelistedPatternType) IsKnown() bool {
	switch r {
	case InvestigateGetResponsePropertiesWhitelistedPatternTypeQuarantineRelease, InvestigateGetResponsePropertiesWhitelistedPatternTypeAcceptableSender, InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedSender, InvestigateGetResponsePropertiesWhitelistedPatternTypeAllowedRecipient, InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainSimilarity, InvestigateGetResponsePropertiesWhitelistedPatternTypeDomainRecency, InvestigateGetResponsePropertiesWhitelistedPatternTypeManagedAcceptableSender:
		return true
	}
	return false
}

type InvestigateGetResponseDeliveryMode string

const (
	InvestigateGetResponseDeliveryModeDirect                InvestigateGetResponseDeliveryMode = "DIRECT"
	InvestigateGetResponseDeliveryModeBcc                   InvestigateGetResponseDeliveryMode = "BCC"
	InvestigateGetResponseDeliveryModeJournal               InvestigateGetResponseDeliveryMode = "JOURNAL"
	InvestigateGetResponseDeliveryModeReviewSubmission      InvestigateGetResponseDeliveryMode = "REVIEW_SUBMISSION"
	InvestigateGetResponseDeliveryModeDMARCUnverified       InvestigateGetResponseDeliveryMode = "DMARC_UNVERIFIED"
	InvestigateGetResponseDeliveryModeDMARCFailureReport    InvestigateGetResponseDeliveryMode = "DMARC_FAILURE_REPORT"
	InvestigateGetResponseDeliveryModeDMARCAggregateReport  InvestigateGetResponseDeliveryMode = "DMARC_AGGREGATE_REPORT"
	InvestigateGetResponseDeliveryModeThreatIntelSubmission InvestigateGetResponseDeliveryMode = "THREAT_INTEL_SUBMISSION"
	InvestigateGetResponseDeliveryModeSimulationSubmission  InvestigateGetResponseDeliveryMode = "SIMULATION_SUBMISSION"
	InvestigateGetResponseDeliveryModeAPI                   InvestigateGetResponseDeliveryMode = "API"
	InvestigateGetResponseDeliveryModeRetroScan             InvestigateGetResponseDeliveryMode = "RETRO_SCAN"
)

func (r InvestigateGetResponseDeliveryMode) IsKnown() bool {
	switch r {
	case InvestigateGetResponseDeliveryModeDirect, InvestigateGetResponseDeliveryModeBcc, InvestigateGetResponseDeliveryModeJournal, InvestigateGetResponseDeliveryModeReviewSubmission, InvestigateGetResponseDeliveryModeDMARCUnverified, InvestigateGetResponseDeliveryModeDMARCFailureReport, InvestigateGetResponseDeliveryModeDMARCAggregateReport, InvestigateGetResponseDeliveryModeThreatIntelSubmission, InvestigateGetResponseDeliveryModeSimulationSubmission, InvestigateGetResponseDeliveryModeAPI, InvestigateGetResponseDeliveryModeRetroScan:
		return true
	}
	return false
}

type InvestigateGetResponseFinalDisposition string

const (
	InvestigateGetResponseFinalDispositionMalicious    InvestigateGetResponseFinalDisposition = "MALICIOUS"
	InvestigateGetResponseFinalDispositionMaliciousBec InvestigateGetResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateGetResponseFinalDispositionSuspicious   InvestigateGetResponseFinalDisposition = "SUSPICIOUS"
	InvestigateGetResponseFinalDispositionSpoof        InvestigateGetResponseFinalDisposition = "SPOOF"
	InvestigateGetResponseFinalDispositionSpam         InvestigateGetResponseFinalDisposition = "SPAM"
	InvestigateGetResponseFinalDispositionBulk         InvestigateGetResponseFinalDisposition = "BULK"
	InvestigateGetResponseFinalDispositionEncrypted    InvestigateGetResponseFinalDisposition = "ENCRYPTED"
	InvestigateGetResponseFinalDispositionExternal     InvestigateGetResponseFinalDisposition = "EXTERNAL"
	InvestigateGetResponseFinalDispositionUnknown      InvestigateGetResponseFinalDisposition = "UNKNOWN"
	InvestigateGetResponseFinalDispositionNone         InvestigateGetResponseFinalDisposition = "NONE"
)

func (r InvestigateGetResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateGetResponseFinalDispositionMalicious, InvestigateGetResponseFinalDispositionMaliciousBec, InvestigateGetResponseFinalDispositionSuspicious, InvestigateGetResponseFinalDispositionSpoof, InvestigateGetResponseFinalDispositionSpam, InvestigateGetResponseFinalDispositionBulk, InvestigateGetResponseFinalDispositionEncrypted, InvestigateGetResponseFinalDispositionExternal, InvestigateGetResponseFinalDispositionUnknown, InvestigateGetResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateGetResponseFinding struct {
	Detail string                            `json:"detail,nullable"`
	Name   string                            `json:"name,nullable"`
	Value  string                            `json:"value,nullable"`
	JSON   investigateGetResponseFindingJSON `json:"-"`
}

// investigateGetResponseFindingJSON contains the JSON metadata for the struct
// [InvestigateGetResponseFinding]
type investigateGetResponseFindingJSON struct {
	Detail      apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseFinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseFindingJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseValidation struct {
	Comment string                                `json:"comment,nullable"`
	DKIM    InvestigateGetResponseValidationDKIM  `json:"dkim,nullable"`
	DMARC   InvestigateGetResponseValidationDMARC `json:"dmarc,nullable"`
	SPF     InvestigateGetResponseValidationSPF   `json:"spf,nullable"`
	JSON    investigateGetResponseValidationJSON  `json:"-"`
}

// investigateGetResponseValidationJSON contains the JSON metadata for the struct
// [InvestigateGetResponseValidation]
type investigateGetResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateGetResponseValidationDKIM string

const (
	InvestigateGetResponseValidationDKIMPass    InvestigateGetResponseValidationDKIM = "pass"
	InvestigateGetResponseValidationDKIMNeutral InvestigateGetResponseValidationDKIM = "neutral"
	InvestigateGetResponseValidationDKIMFail    InvestigateGetResponseValidationDKIM = "fail"
	InvestigateGetResponseValidationDKIMError   InvestigateGetResponseValidationDKIM = "error"
	InvestigateGetResponseValidationDKIMNone    InvestigateGetResponseValidationDKIM = "none"
)

func (r InvestigateGetResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateGetResponseValidationDKIMPass, InvestigateGetResponseValidationDKIMNeutral, InvestigateGetResponseValidationDKIMFail, InvestigateGetResponseValidationDKIMError, InvestigateGetResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateGetResponseValidationDMARC string

const (
	InvestigateGetResponseValidationDMARCPass    InvestigateGetResponseValidationDMARC = "pass"
	InvestigateGetResponseValidationDMARCNeutral InvestigateGetResponseValidationDMARC = "neutral"
	InvestigateGetResponseValidationDMARCFail    InvestigateGetResponseValidationDMARC = "fail"
	InvestigateGetResponseValidationDMARCError   InvestigateGetResponseValidationDMARC = "error"
	InvestigateGetResponseValidationDMARCNone    InvestigateGetResponseValidationDMARC = "none"
)

func (r InvestigateGetResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateGetResponseValidationDMARCPass, InvestigateGetResponseValidationDMARCNeutral, InvestigateGetResponseValidationDMARCFail, InvestigateGetResponseValidationDMARCError, InvestigateGetResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateGetResponseValidationSPF string

const (
	InvestigateGetResponseValidationSPFPass    InvestigateGetResponseValidationSPF = "pass"
	InvestigateGetResponseValidationSPFNeutral InvestigateGetResponseValidationSPF = "neutral"
	InvestigateGetResponseValidationSPFFail    InvestigateGetResponseValidationSPF = "fail"
	InvestigateGetResponseValidationSPFError   InvestigateGetResponseValidationSPF = "error"
	InvestigateGetResponseValidationSPFNone    InvestigateGetResponseValidationSPF = "none"
)

func (r InvestigateGetResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateGetResponseValidationSPFPass, InvestigateGetResponseValidationSPFNeutral, InvestigateGetResponseValidationSPFFail, InvestigateGetResponseValidationSPFError, InvestigateGetResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Determines if the message action log is included in the response.
	ActionLog param.Field[bool]   `query:"action_log"`
	AlertID   param.Field[string] `query:"alert_id"`
	// Determines if the search results will include detections or not.
	DetectionsOnly param.Field[bool] `query:"detections_only"`
	// The sender domains the search filters by.
	Domain param.Field[string] `query:"domain"`
	// The end of the search date range. Defaults to `now`.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// The dispositions the search filters by.
	FinalDisposition param.Field[InvestigateListParamsFinalDisposition] `query:"final_disposition"`
	// The message actions the search filters by.
	MessageAction param.Field[InvestigateListParamsMessageAction] `query:"message_action"`
	MessageID     param.Field[string]                             `query:"message_id"`
	Metric        param.Field[string]                             `query:"metric"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The space-delimited term used in the query. The search is case-insensitive.
	//
	// The content of the following email metadata fields are searched:
	//
	// - alert_id
	// - CC
	// - From (envelope_from)
	// - From Name
	// - final_disposition
	// - md5 hash (of any attachment)
	// - sha1 hash (of any attachment)
	// - sha256 hash (of any attachment)
	// - name (of any attachment)
	// - Reason
	// - Received DateTime (yyyy-mm-ddThh:mm:ss)
	// - Sent DateTime (yyyy-mm-ddThh:mm:ss)
	// - ReplyTo
	// - To (envelope_to)
	// - To Name
	// - Message-ID
	// - smtp_helo_server_ip
	// - smtp_previous_hop_ip
	// - x_originating_ip
	// - Subject
	Query     param.Field[string] `query:"query"`
	Recipient param.Field[string] `query:"recipient"`
	Sender    param.Field[string] `query:"sender"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start   param.Field[time.Time] `query:"start" format:"date-time"`
	Subject param.Field[string]    `query:"subject"`
}

// URLQuery serializes [InvestigateListParams]'s query parameters as `url.Values`.
func (r InvestigateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The dispositions the search filters by.
type InvestigateListParamsFinalDisposition string

const (
	InvestigateListParamsFinalDispositionMalicious  InvestigateListParamsFinalDisposition = "MALICIOUS"
	InvestigateListParamsFinalDispositionSuspicious InvestigateListParamsFinalDisposition = "SUSPICIOUS"
	InvestigateListParamsFinalDispositionSpoof      InvestigateListParamsFinalDisposition = "SPOOF"
	InvestigateListParamsFinalDispositionSpam       InvestigateListParamsFinalDisposition = "SPAM"
	InvestigateListParamsFinalDispositionBulk       InvestigateListParamsFinalDisposition = "BULK"
	InvestigateListParamsFinalDispositionNone       InvestigateListParamsFinalDisposition = "NONE"
)

func (r InvestigateListParamsFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateListParamsFinalDispositionMalicious, InvestigateListParamsFinalDispositionSuspicious, InvestigateListParamsFinalDispositionSpoof, InvestigateListParamsFinalDispositionSpam, InvestigateListParamsFinalDispositionBulk, InvestigateListParamsFinalDispositionNone:
		return true
	}
	return false
}

// The message actions the search filters by.
type InvestigateListParamsMessageAction string

const (
	InvestigateListParamsMessageActionPreview            InvestigateListParamsMessageAction = "PREVIEW"
	InvestigateListParamsMessageActionQuarantineReleased InvestigateListParamsMessageAction = "QUARANTINE_RELEASED"
	InvestigateListParamsMessageActionMoved              InvestigateListParamsMessageAction = "MOVED"
)

func (r InvestigateListParamsMessageAction) IsKnown() bool {
	switch r {
	case InvestigateListParamsMessageActionPreview, InvestigateListParamsMessageActionQuarantineReleased, InvestigateListParamsMessageActionMoved:
		return true
	}
	return false
}

type InvestigateGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   InvestigateGetResponse             `json:"result,required"`
	Success  bool                               `json:"success,required"`
	JSON     investigateGetResponseEnvelopeJSON `json:"-"`
}

// investigateGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateGetResponseEnvelope]
type investigateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
