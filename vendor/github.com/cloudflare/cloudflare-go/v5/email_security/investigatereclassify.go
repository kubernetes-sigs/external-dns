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

// InvestigateReclassifyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateReclassifyService] method instead.
type InvestigateReclassifyService struct {
	Options []option.RequestOption
}

// NewInvestigateReclassifyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateReclassifyService(opts ...option.RequestOption) (r *InvestigateReclassifyService) {
	r = &InvestigateReclassifyService{}
	r.Options = opts
	return
}

// Change email classfication
func (r *InvestigateReclassifyService) New(ctx context.Context, postfixID string, params InvestigateReclassifyNewParams, opts ...option.RequestOption) (res *InvestigateReclassifyNewResponse, err error) {
	var env InvestigateReclassifyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/reclassify", params.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateReclassifyNewResponse = interface{}

type InvestigateReclassifyNewParams struct {
	// Account Identifier
	AccountID           param.Field[string]                                            `path:"account_id,required"`
	ExpectedDisposition param.Field[InvestigateReclassifyNewParamsExpectedDisposition] `json:"expected_disposition,required"`
	// Base64 encoded content of the EML file
	EmlContent            param.Field[string] `json:"eml_content"`
	EscalatedSubmissionID param.Field[string] `json:"escalated_submission_id"`
}

func (r InvestigateReclassifyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigateReclassifyNewParamsExpectedDisposition string

const (
	InvestigateReclassifyNewParamsExpectedDispositionNone       InvestigateReclassifyNewParamsExpectedDisposition = "NONE"
	InvestigateReclassifyNewParamsExpectedDispositionBulk       InvestigateReclassifyNewParamsExpectedDisposition = "BULK"
	InvestigateReclassifyNewParamsExpectedDispositionMalicious  InvestigateReclassifyNewParamsExpectedDisposition = "MALICIOUS"
	InvestigateReclassifyNewParamsExpectedDispositionSpam       InvestigateReclassifyNewParamsExpectedDisposition = "SPAM"
	InvestigateReclassifyNewParamsExpectedDispositionSpoof      InvestigateReclassifyNewParamsExpectedDisposition = "SPOOF"
	InvestigateReclassifyNewParamsExpectedDispositionSuspicious InvestigateReclassifyNewParamsExpectedDisposition = "SUSPICIOUS"
)

func (r InvestigateReclassifyNewParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateReclassifyNewParamsExpectedDispositionNone, InvestigateReclassifyNewParamsExpectedDispositionBulk, InvestigateReclassifyNewParamsExpectedDispositionMalicious, InvestigateReclassifyNewParamsExpectedDispositionSpam, InvestigateReclassifyNewParamsExpectedDispositionSpoof, InvestigateReclassifyNewParamsExpectedDispositionSuspicious:
		return true
	}
	return false
}

type InvestigateReclassifyNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                        `json:"errors,required"`
	Messages []shared.ResponseInfo                        `json:"messages,required"`
	Result   InvestigateReclassifyNewResponse             `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	JSON     investigateReclassifyNewResponseEnvelopeJSON `json:"-"`
}

// investigateReclassifyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateReclassifyNewResponseEnvelope]
type investigateReclassifyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateReclassifyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReclassifyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
