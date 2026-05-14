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

// InvestigateRawService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateRawService] method instead.
type InvestigateRawService struct {
	Options []option.RequestOption
}

// NewInvestigateRawService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateRawService(opts ...option.RequestOption) (r *InvestigateRawService) {
	r = &InvestigateRawService{}
	r.Options = opts
	return
}

// Returns the raw eml of any non-benign message.
func (r *InvestigateRawService) Get(ctx context.Context, postfixID string, query InvestigateRawGetParams, opts ...option.RequestOption) (res *InvestigateRawGetResponse, err error) {
	var env InvestigateRawGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/raw", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateRawGetResponse struct {
	// A UTF-8 encoded eml file of the email.
	Raw  string                        `json:"raw,required"`
	JSON investigateRawGetResponseJSON `json:"-"`
}

// investigateRawGetResponseJSON contains the JSON metadata for the struct
// [InvestigateRawGetResponse]
type investigateRawGetResponseJSON struct {
	Raw         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateRawGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []shared.ResponseInfo                 `json:"messages,required"`
	Result   InvestigateRawGetResponse             `json:"result,required"`
	Success  bool                                  `json:"success,required"`
	JSON     investigateRawGetResponseEnvelopeJSON `json:"-"`
}

// investigateRawGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateRawGetResponseEnvelope]
type investigateRawGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
