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

// InvestigatePreviewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigatePreviewService] method instead.
type InvestigatePreviewService struct {
	Options []option.RequestOption
}

// NewInvestigatePreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigatePreviewService(opts ...option.RequestOption) (r *InvestigatePreviewService) {
	r = &InvestigatePreviewService{}
	r.Options = opts
	return
}

// Preview for non-detection messages
func (r *InvestigatePreviewService) New(ctx context.Context, params InvestigatePreviewNewParams, opts ...option.RequestOption) (res *InvestigatePreviewNewResponse, err error) {
	var env InvestigatePreviewNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/preview", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a preview of the message body as a base64 encoded PNG image for
// non-benign messages.
func (r *InvestigatePreviewService) Get(ctx context.Context, postfixID string, query InvestigatePreviewGetParams, opts ...option.RequestOption) (res *InvestigatePreviewGetResponse, err error) {
	var env InvestigatePreviewGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/preview", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigatePreviewNewResponse struct {
	// A base64 encoded PNG image of the email.
	Screenshot string                            `json:"screenshot,required"`
	JSON       investigatePreviewNewResponseJSON `json:"-"`
}

// investigatePreviewNewResponseJSON contains the JSON metadata for the struct
// [InvestigatePreviewNewResponse]
type investigatePreviewNewResponseJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetResponse struct {
	// A base64 encoded PNG image of the email.
	Screenshot string                            `json:"screenshot,required"`
	JSON       investigatePreviewGetResponseJSON `json:"-"`
}

// investigatePreviewGetResponseJSON contains the JSON metadata for the struct
// [InvestigatePreviewGetResponse]
type investigatePreviewGetResponseJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewNewParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The identifier of the message.
	PostfixID param.Field[string] `json:"postfix_id,required"`
}

func (r InvestigatePreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigatePreviewNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   InvestigatePreviewNewResponse             `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     investigatePreviewNewResponseEnvelopeJSON `json:"-"`
}

// investigatePreviewNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigatePreviewNewResponseEnvelope]
type investigatePreviewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigatePreviewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   InvestigatePreviewGetResponse             `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     investigatePreviewGetResponseEnvelopeJSON `json:"-"`
}

// investigatePreviewGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigatePreviewGetResponseEnvelope]
type investigatePreviewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
