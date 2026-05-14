// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// QuotaService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options []option.RequestOption
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r *QuotaService) {
	r = &QuotaService{}
	r.Options = opts
	return
}

// Lists the number of secrets used in the account.
func (r *QuotaService) Get(ctx context.Context, query QuotaGetParams, opts ...option.RequestOption) (res *QuotaGetResponse, err error) {
	var env QuotaGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secrets_store/quota", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QuotaGetResponse struct {
	Secrets QuotaGetResponseSecrets `json:"secrets,required"`
	JSON    quotaGetResponseJSON    `json:"-"`
}

// quotaGetResponseJSON contains the JSON metadata for the struct
// [QuotaGetResponse]
type quotaGetResponseJSON struct {
	Secrets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseSecrets struct {
	// The number of secrets the account is entitlted to use
	Quota float64 `json:"quota,required"`
	// The number of secrets the account is currently using
	Usage float64                     `json:"usage,required"`
	JSON  quotaGetResponseSecretsJSON `json:"-"`
}

// quotaGetResponseSecretsJSON contains the JSON metadata for the struct
// [QuotaGetResponseSecrets]
type quotaGetResponseSecretsJSON struct {
	Quota       apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseSecrets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseSecretsJSON) RawJSON() string {
	return r.raw
}

type QuotaGetParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QuotaGetResponseEnvelope struct {
	Errors   []QuotaGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QuotaGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    QuotaGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     QuotaGetResponse                   `json:"result"`
	ResultInfo QuotaGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       quotaGetResponseEnvelopeJSON       `json:"-"`
}

// quotaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [QuotaGetResponseEnvelope]
type quotaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           QuotaGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             quotaGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// quotaGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QuotaGetResponseEnvelopeErrors]
type quotaGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    quotaGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// quotaGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [QuotaGetResponseEnvelopeErrorsSource]
type quotaGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           QuotaGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             quotaGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// quotaGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [QuotaGetResponseEnvelopeMessages]
type quotaGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    quotaGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// quotaGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [QuotaGetResponseEnvelopeMessagesSource]
type quotaGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type QuotaGetResponseEnvelopeSuccess bool

const (
	QuotaGetResponseEnvelopeSuccessTrue QuotaGetResponseEnvelopeSuccess = true
)

func (r QuotaGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QuotaGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type QuotaGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                `json:"total_count"`
	JSON       quotaGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// quotaGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [QuotaGetResponseEnvelopeResultInfo]
type quotaGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
