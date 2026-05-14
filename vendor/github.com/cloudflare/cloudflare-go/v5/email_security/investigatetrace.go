// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// InvestigateTraceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateTraceService] method instead.
type InvestigateTraceService struct {
	Options []option.RequestOption
}

// NewInvestigateTraceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateTraceService(opts ...option.RequestOption) (r *InvestigateTraceService) {
	r = &InvestigateTraceService{}
	r.Options = opts
	return
}

// Get email trace
func (r *InvestigateTraceService) Get(ctx context.Context, postfixID string, query InvestigateTraceGetParams, opts ...option.RequestOption) (res *InvestigateTraceGetResponse, err error) {
	var env InvestigateTraceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/trace", query.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InvestigateTraceGetResponse struct {
	Inbound  InvestigateTraceGetResponseInbound  `json:"inbound,required"`
	Outbound InvestigateTraceGetResponseOutbound `json:"outbound,required"`
	JSON     investigateTraceGetResponseJSON     `json:"-"`
}

// investigateTraceGetResponseJSON contains the JSON metadata for the struct
// [InvestigateTraceGetResponse]
type investigateTraceGetResponseJSON struct {
	Inbound     apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseInbound struct {
	Lines   []InvestigateTraceGetResponseInboundLine `json:"lines,nullable"`
	Pending bool                                     `json:"pending,nullable"`
	JSON    investigateTraceGetResponseInboundJSON   `json:"-"`
}

// investigateTraceGetResponseInboundJSON contains the JSON metadata for the struct
// [InvestigateTraceGetResponseInbound]
type investigateTraceGetResponseInboundJSON struct {
	Lines       apijson.Field
	Pending     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseInbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseInboundJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseInboundLine struct {
	Lineno  int64                                      `json:"lineno,required"`
	Message string                                     `json:"message,required"`
	Ts      time.Time                                  `json:"ts,required" format:"date-time"`
	JSON    investigateTraceGetResponseInboundLineJSON `json:"-"`
}

// investigateTraceGetResponseInboundLineJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseInboundLine]
type investigateTraceGetResponseInboundLineJSON struct {
	Lineno      apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseInboundLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseInboundLineJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseOutbound struct {
	Lines   []InvestigateTraceGetResponseOutboundLine `json:"lines,nullable"`
	Pending bool                                      `json:"pending,nullable"`
	JSON    investigateTraceGetResponseOutboundJSON   `json:"-"`
}

// investigateTraceGetResponseOutboundJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseOutbound]
type investigateTraceGetResponseOutboundJSON struct {
	Lines       apijson.Field
	Pending     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseOutboundJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseOutboundLine struct {
	Lineno  int64                                       `json:"lineno,required"`
	Message string                                      `json:"message,required"`
	Ts      time.Time                                   `json:"ts,required" format:"date-time"`
	JSON    investigateTraceGetResponseOutboundLineJSON `json:"-"`
}

// investigateTraceGetResponseOutboundLineJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseOutboundLine]
type investigateTraceGetResponseOutboundLineJSON struct {
	Lineno      apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseOutboundLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseOutboundLineJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type InvestigateTraceGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors,required"`
	Messages []shared.ResponseInfo                   `json:"messages,required"`
	Result   InvestigateTraceGetResponse             `json:"result,required"`
	Success  bool                                    `json:"success,required"`
	JSON     investigateTraceGetResponseEnvelopeJSON `json:"-"`
}

// investigateTraceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseEnvelope]
type investigateTraceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
