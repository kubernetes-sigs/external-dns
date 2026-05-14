// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// InvestigateDetectionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateDetectionService] method instead.
type InvestigateDetectionService struct {
	Options []option.RequestOption
}

// NewInvestigateDetectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateDetectionService(opts ...option.RequestOption) (r *InvestigateDetectionService) {
	r = &InvestigateDetectionService{}
	r.Options = opts
	return
}

// Returns detection details such as threat categories and sender information for
// non-benign messages.
func (r *InvestigateDetectionService) Get(ctx context.Context, postfixID string, query InvestigateDetectionGetParams, opts ...option.RequestOption) (res *InvestigateDetectionGetResponse, err error) {
	var env InvestigateDetectionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/detections", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateDetectionGetResponse struct {
	Action           string                                          `json:"action,required"`
	Attachments      []InvestigateDetectionGetResponseAttachment     `json:"attachments,required"`
	Headers          []InvestigateDetectionGetResponseHeader         `json:"headers,required"`
	Links            []InvestigateDetectionGetResponseLink           `json:"links,required"`
	SenderInfo       InvestigateDetectionGetResponseSenderInfo       `json:"sender_info,required"`
	ThreatCategories []InvestigateDetectionGetResponseThreatCategory `json:"threat_categories,required"`
	Validation       InvestigateDetectionGetResponseValidation       `json:"validation,required"`
	FinalDisposition InvestigateDetectionGetResponseFinalDisposition `json:"final_disposition,nullable"`
	JSON             investigateDetectionGetResponseJSON             `json:"-"`
}

// investigateDetectionGetResponseJSON contains the JSON metadata for the struct
// [InvestigateDetectionGetResponse]
type investigateDetectionGetResponseJSON struct {
	Action           apijson.Field
	Attachments      apijson.Field
	Headers          apijson.Field
	Links            apijson.Field
	SenderInfo       apijson.Field
	ThreatCategories apijson.Field
	Validation       apijson.Field
	FinalDisposition apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseAttachment struct {
	Size        int64                                               `json:"size,required"`
	ContentType string                                              `json:"content_type,nullable"`
	Detection   InvestigateDetectionGetResponseAttachmentsDetection `json:"detection,nullable"`
	Encrypted   bool                                                `json:"encrypted,nullable"`
	Name        string                                              `json:"name,nullable"`
	JSON        investigateDetectionGetResponseAttachmentJSON       `json:"-"`
}

// investigateDetectionGetResponseAttachmentJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseAttachment]
type investigateDetectionGetResponseAttachmentJSON struct {
	Size        apijson.Field
	ContentType apijson.Field
	Detection   apijson.Field
	Encrypted   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseAttachmentJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseAttachmentsDetection string

const (
	InvestigateDetectionGetResponseAttachmentsDetectionMalicious    InvestigateDetectionGetResponseAttachmentsDetection = "MALICIOUS"
	InvestigateDetectionGetResponseAttachmentsDetectionMaliciousBec InvestigateDetectionGetResponseAttachmentsDetection = "MALICIOUS-BEC"
	InvestigateDetectionGetResponseAttachmentsDetectionSuspicious   InvestigateDetectionGetResponseAttachmentsDetection = "SUSPICIOUS"
	InvestigateDetectionGetResponseAttachmentsDetectionSpoof        InvestigateDetectionGetResponseAttachmentsDetection = "SPOOF"
	InvestigateDetectionGetResponseAttachmentsDetectionSpam         InvestigateDetectionGetResponseAttachmentsDetection = "SPAM"
	InvestigateDetectionGetResponseAttachmentsDetectionBulk         InvestigateDetectionGetResponseAttachmentsDetection = "BULK"
	InvestigateDetectionGetResponseAttachmentsDetectionEncrypted    InvestigateDetectionGetResponseAttachmentsDetection = "ENCRYPTED"
	InvestigateDetectionGetResponseAttachmentsDetectionExternal     InvestigateDetectionGetResponseAttachmentsDetection = "EXTERNAL"
	InvestigateDetectionGetResponseAttachmentsDetectionUnknown      InvestigateDetectionGetResponseAttachmentsDetection = "UNKNOWN"
	InvestigateDetectionGetResponseAttachmentsDetectionNone         InvestigateDetectionGetResponseAttachmentsDetection = "NONE"
)

func (r InvestigateDetectionGetResponseAttachmentsDetection) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseAttachmentsDetectionMalicious, InvestigateDetectionGetResponseAttachmentsDetectionMaliciousBec, InvestigateDetectionGetResponseAttachmentsDetectionSuspicious, InvestigateDetectionGetResponseAttachmentsDetectionSpoof, InvestigateDetectionGetResponseAttachmentsDetectionSpam, InvestigateDetectionGetResponseAttachmentsDetectionBulk, InvestigateDetectionGetResponseAttachmentsDetectionEncrypted, InvestigateDetectionGetResponseAttachmentsDetectionExternal, InvestigateDetectionGetResponseAttachmentsDetectionUnknown, InvestigateDetectionGetResponseAttachmentsDetectionNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseHeader struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  investigateDetectionGetResponseHeaderJSON `json:"-"`
}

// investigateDetectionGetResponseHeaderJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseHeader]
type investigateDetectionGetResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseLink struct {
	Href string                                  `json:"href,required"`
	Text string                                  `json:"text,nullable"`
	JSON investigateDetectionGetResponseLinkJSON `json:"-"`
}

// investigateDetectionGetResponseLinkJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseLink]
type investigateDetectionGetResponseLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseLinkJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseSenderInfo struct {
	// The name of the autonomous system.
	AsName string `json:"as_name,nullable"`
	// The number of the autonomous system.
	AsNumber int64                                         `json:"as_number,nullable"`
	Geo      string                                        `json:"geo,nullable"`
	IP       string                                        `json:"ip,nullable"`
	Pld      string                                        `json:"pld,nullable"`
	JSON     investigateDetectionGetResponseSenderInfoJSON `json:"-"`
}

// investigateDetectionGetResponseSenderInfoJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseSenderInfo]
type investigateDetectionGetResponseSenderInfoJSON struct {
	AsName      apijson.Field
	AsNumber    apijson.Field
	Geo         apijson.Field
	IP          apijson.Field
	Pld         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseSenderInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseSenderInfoJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseThreatCategory struct {
	ID          int64                                             `json:"id,required"`
	Description string                                            `json:"description,nullable"`
	Name        string                                            `json:"name,nullable"`
	JSON        investigateDetectionGetResponseThreatCategoryJSON `json:"-"`
}

// investigateDetectionGetResponseThreatCategoryJSON contains the JSON metadata for
// the struct [InvestigateDetectionGetResponseThreatCategory]
type investigateDetectionGetResponseThreatCategoryJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseThreatCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseThreatCategoryJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseValidation struct {
	Comment string                                         `json:"comment,nullable"`
	DKIM    InvestigateDetectionGetResponseValidationDKIM  `json:"dkim,nullable"`
	DMARC   InvestigateDetectionGetResponseValidationDMARC `json:"dmarc,nullable"`
	SPF     InvestigateDetectionGetResponseValidationSPF   `json:"spf,nullable"`
	JSON    investigateDetectionGetResponseValidationJSON  `json:"-"`
}

// investigateDetectionGetResponseValidationJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseValidation]
type investigateDetectionGetResponseValidationJSON struct {
	Comment     apijson.Field
	DKIM        apijson.Field
	DMARC       apijson.Field
	SPF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseValidationJSON) RawJSON() string {
	return r.raw
}

type InvestigateDetectionGetResponseValidationDKIM string

const (
	InvestigateDetectionGetResponseValidationDKIMPass    InvestigateDetectionGetResponseValidationDKIM = "pass"
	InvestigateDetectionGetResponseValidationDKIMNeutral InvestigateDetectionGetResponseValidationDKIM = "neutral"
	InvestigateDetectionGetResponseValidationDKIMFail    InvestigateDetectionGetResponseValidationDKIM = "fail"
	InvestigateDetectionGetResponseValidationDKIMError   InvestigateDetectionGetResponseValidationDKIM = "error"
	InvestigateDetectionGetResponseValidationDKIMNone    InvestigateDetectionGetResponseValidationDKIM = "none"
)

func (r InvestigateDetectionGetResponseValidationDKIM) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseValidationDKIMPass, InvestigateDetectionGetResponseValidationDKIMNeutral, InvestigateDetectionGetResponseValidationDKIMFail, InvestigateDetectionGetResponseValidationDKIMError, InvestigateDetectionGetResponseValidationDKIMNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseValidationDMARC string

const (
	InvestigateDetectionGetResponseValidationDMARCPass    InvestigateDetectionGetResponseValidationDMARC = "pass"
	InvestigateDetectionGetResponseValidationDMARCNeutral InvestigateDetectionGetResponseValidationDMARC = "neutral"
	InvestigateDetectionGetResponseValidationDMARCFail    InvestigateDetectionGetResponseValidationDMARC = "fail"
	InvestigateDetectionGetResponseValidationDMARCError   InvestigateDetectionGetResponseValidationDMARC = "error"
	InvestigateDetectionGetResponseValidationDMARCNone    InvestigateDetectionGetResponseValidationDMARC = "none"
)

func (r InvestigateDetectionGetResponseValidationDMARC) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseValidationDMARCPass, InvestigateDetectionGetResponseValidationDMARCNeutral, InvestigateDetectionGetResponseValidationDMARCFail, InvestigateDetectionGetResponseValidationDMARCError, InvestigateDetectionGetResponseValidationDMARCNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseValidationSPF string

const (
	InvestigateDetectionGetResponseValidationSPFPass    InvestigateDetectionGetResponseValidationSPF = "pass"
	InvestigateDetectionGetResponseValidationSPFNeutral InvestigateDetectionGetResponseValidationSPF = "neutral"
	InvestigateDetectionGetResponseValidationSPFFail    InvestigateDetectionGetResponseValidationSPF = "fail"
	InvestigateDetectionGetResponseValidationSPFError   InvestigateDetectionGetResponseValidationSPF = "error"
	InvestigateDetectionGetResponseValidationSPFNone    InvestigateDetectionGetResponseValidationSPF = "none"
)

func (r InvestigateDetectionGetResponseValidationSPF) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseValidationSPFPass, InvestigateDetectionGetResponseValidationSPFNeutral, InvestigateDetectionGetResponseValidationSPFFail, InvestigateDetectionGetResponseValidationSPFError, InvestigateDetectionGetResponseValidationSPFNone:
		return true
	}
	return false
}

type InvestigateDetectionGetResponseFinalDisposition string

const (
	InvestigateDetectionGetResponseFinalDispositionMalicious    InvestigateDetectionGetResponseFinalDisposition = "MALICIOUS"
	InvestigateDetectionGetResponseFinalDispositionMaliciousBec InvestigateDetectionGetResponseFinalDisposition = "MALICIOUS-BEC"
	InvestigateDetectionGetResponseFinalDispositionSuspicious   InvestigateDetectionGetResponseFinalDisposition = "SUSPICIOUS"
	InvestigateDetectionGetResponseFinalDispositionSpoof        InvestigateDetectionGetResponseFinalDisposition = "SPOOF"
	InvestigateDetectionGetResponseFinalDispositionSpam         InvestigateDetectionGetResponseFinalDisposition = "SPAM"
	InvestigateDetectionGetResponseFinalDispositionBulk         InvestigateDetectionGetResponseFinalDisposition = "BULK"
	InvestigateDetectionGetResponseFinalDispositionEncrypted    InvestigateDetectionGetResponseFinalDisposition = "ENCRYPTED"
	InvestigateDetectionGetResponseFinalDispositionExternal     InvestigateDetectionGetResponseFinalDisposition = "EXTERNAL"
	InvestigateDetectionGetResponseFinalDispositionUnknown      InvestigateDetectionGetResponseFinalDisposition = "UNKNOWN"
	InvestigateDetectionGetResponseFinalDispositionNone         InvestigateDetectionGetResponseFinalDisposition = "NONE"
)

func (r InvestigateDetectionGetResponseFinalDisposition) IsKnown() bool {
	switch r {
	case InvestigateDetectionGetResponseFinalDispositionMalicious, InvestigateDetectionGetResponseFinalDispositionMaliciousBec, InvestigateDetectionGetResponseFinalDispositionSuspicious, InvestigateDetectionGetResponseFinalDispositionSpoof, InvestigateDetectionGetResponseFinalDispositionSpam, InvestigateDetectionGetResponseFinalDispositionBulk, InvestigateDetectionGetResponseFinalDispositionEncrypted, InvestigateDetectionGetResponseFinalDispositionExternal, InvestigateDetectionGetResponseFinalDispositionUnknown, InvestigateDetectionGetResponseFinalDispositionNone:
		return true
	}
	return false
}

type InvestigateDetectionGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateDetectionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   InvestigateDetectionGetResponse             `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     investigateDetectionGetResponseEnvelopeJSON `json:"-"`
}

// investigateDetectionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateDetectionGetResponseEnvelope]
type investigateDetectionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateDetectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateDetectionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
