// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// RiskScoringIntegrationReferenceService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringIntegrationReferenceService] method instead.
type RiskScoringIntegrationReferenceService struct {
	Options []option.RequestOption
}

// NewRiskScoringIntegrationReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRiskScoringIntegrationReferenceService(opts ...option.RequestOption) (r *RiskScoringIntegrationReferenceService) {
	r = &RiskScoringIntegrationReferenceService{}
	r.Options = opts
	return
}

// Get risk score integration by reference id.
func (r *RiskScoringIntegrationReferenceService) Get(ctx context.Context, referenceID string, query RiskScoringIntegrationReferenceGetParams, opts ...option.RequestOption) (res *RiskScoringIntegrationReferenceGetResponse, err error) {
	var env RiskScoringIntegrationReferenceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if referenceID == "" {
		err = errors.New("missing required reference_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/integrations/reference_id/%s", query.AccountID, referenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringIntegrationReferenceGetResponse struct {
	// The id of the integration, a UUIDv4.
	ID string `json:"id,required" format:"uuid"`
	// The Cloudflare account tag.
	AccountTag string `json:"account_tag,required"`
	// Whether this integration is enabled and should export changes in risk score.
	Active bool `json:"active,required"`
	// When the integration was created in RFC3339 format.
	CreatedAt       time.Time                                                 `json:"created_at,required" format:"date-time"`
	IntegrationType RiskScoringIntegrationReferenceGetResponseIntegrationType `json:"integration_type,required"`
	// A reference ID defined by the client. Should be set to the Access-Okta IDP
	// integration ID. Useful when the risk-score integration needs to be associated
	// with a secondary asset and recalled using that ID.
	ReferenceID string `json:"reference_id,required"`
	// The base URL for the tenant. E.g. "https://tenant.okta.com".
	TenantURL string `json:"tenant_url,required"`
	// The URL for the Shared Signals Framework configuration, e.g.
	// "/.well-known/sse-configuration/{integration_uuid}/".
	// https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1.
	WellKnownURL string                                         `json:"well_known_url,required"`
	JSON         riskScoringIntegrationReferenceGetResponseJSON `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseJSON contains the JSON metadata for
// the struct [RiskScoringIntegrationReferenceGetResponse]
type riskScoringIntegrationReferenceGetResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Active          apijson.Field
	CreatedAt       apijson.Field
	IntegrationType apijson.Field
	ReferenceID     apijson.Field
	TenantURL       apijson.Field
	WellKnownURL    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseIntegrationType string

const (
	RiskScoringIntegrationReferenceGetResponseIntegrationTypeOkta RiskScoringIntegrationReferenceGetResponseIntegrationType = "Okta"
)

func (r RiskScoringIntegrationReferenceGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationReferenceGetResponseIntegrationTypeOkta:
		return true
	}
	return false
}

type RiskScoringIntegrationReferenceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type RiskScoringIntegrationReferenceGetResponseEnvelope struct {
	Errors   []RiskScoringIntegrationReferenceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RiskScoringIntegrationReferenceGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RiskScoringIntegrationReferenceGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RiskScoringIntegrationReferenceGetResponse                `json:"result"`
	JSON    riskScoringIntegrationReferenceGetResponseEnvelopeJSON    `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RiskScoringIntegrationReferenceGetResponseEnvelope]
type riskScoringIntegrationReferenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           RiskScoringIntegrationReferenceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             riskScoringIntegrationReferenceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [RiskScoringIntegrationReferenceGetResponseEnvelopeErrors]
type riskScoringIntegrationReferenceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    riskScoringIntegrationReferenceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [RiskScoringIntegrationReferenceGetResponseEnvelopeErrorsSource]
type riskScoringIntegrationReferenceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           RiskScoringIntegrationReferenceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             riskScoringIntegrationReferenceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [RiskScoringIntegrationReferenceGetResponseEnvelopeMessages]
type riskScoringIntegrationReferenceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RiskScoringIntegrationReferenceGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    riskScoringIntegrationReferenceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// riskScoringIntegrationReferenceGetResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [RiskScoringIntegrationReferenceGetResponseEnvelopeMessagesSource]
type riskScoringIntegrationReferenceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringIntegrationReferenceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringIntegrationReferenceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RiskScoringIntegrationReferenceGetResponseEnvelopeSuccess bool

const (
	RiskScoringIntegrationReferenceGetResponseEnvelopeSuccessTrue RiskScoringIntegrationReferenceGetResponseEnvelopeSuccess = true
)

func (r RiskScoringIntegrationReferenceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringIntegrationReferenceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
