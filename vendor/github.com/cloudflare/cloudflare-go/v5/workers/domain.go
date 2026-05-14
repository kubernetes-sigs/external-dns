// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	return
}

// Attaches a Worker to a zone and hostname.
func (r *DomainService) Update(ctx context.Context, params DomainUpdateParams, opts ...option.RequestOption) (res *Domain, err error) {
	var env DomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Worker Domains for an account.
func (r *DomainService) List(ctx context.Context, params DomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Domain], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains", params.AccountID)
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

// Lists all Worker Domains for an account.
func (r *DomainService) ListAutoPaging(ctx context.Context, params DomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Domain] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Detaches a Worker from a zone and hostname.
func (r *DomainService) Delete(ctx context.Context, domainID string, body DomainDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", body.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Gets a Worker domain.
func (r *DomainService) Get(ctx context.Context, domainID string, query DomainGetParams, opts ...option.RequestOption) (res *Domain, err error) {
	var env DomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainID == "" {
		err = errors.New("missing required domain_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/domains/%s", query.AccountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Domain struct {
	// Identifer of the Worker Domain.
	ID string `json:"id"`
	// Worker environment associated with the zone and hostname.
	Environment string `json:"environment"`
	// Hostname of the Worker Domain.
	Hostname string `json:"hostname"`
	// Worker service associated with the zone and hostname.
	Service string `json:"service"`
	// Identifier of the zone.
	ZoneID string `json:"zone_id"`
	// Name of the zone.
	ZoneName string     `json:"zone_name"`
	JSON     domainJSON `json:"-"`
}

// domainJSON contains the JSON metadata for the struct [Domain]
type domainJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Domain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `json:"environment,required"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `json:"hostname,required"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `json:"service,required"`
	// Identifier of the zone.
	ZoneID param.Field[string] `json:"zone_id,required"`
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainUpdateResponseEnvelope struct {
	Errors   []DomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Domain                              `json:"result"`
	JSON    domainUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelope]
type domainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           DomainUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             domainUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// domainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelopeErrors]
type domainUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    domainUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// domainUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DomainUpdateResponseEnvelopeErrorsSource]
type domainUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DomainUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             domainUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// domainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DomainUpdateResponseEnvelopeMessages]
type domainUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DomainUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    domainUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// domainUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DomainUpdateResponseEnvelopeMessagesSource]
type domainUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DomainUpdateResponseEnvelopeSuccess bool

const (
	DomainUpdateResponseEnvelopeSuccessTrue DomainUpdateResponseEnvelopeSuccess = true
)

func (r DomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainListParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `query:"environment"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `query:"hostname"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `query:"service"`
	// Identifier of the zone.
	ZoneID param.Field[string] `query:"zone_id"`
	// Name of the zone.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [DomainListParams]'s query parameters as `url.Values`.
func (r DomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DomainDeleteParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetParams struct {
	// Identifer of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetResponseEnvelope struct {
	Errors   []DomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DomainGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Domain                           `json:"result"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           DomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             domainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// domainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeErrors]
type domainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    domainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// domainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DomainGetResponseEnvelopeErrorsSource]
type domainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             domainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// domainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeMessages]
type domainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    domainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// domainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DomainGetResponseEnvelopeMessagesSource]
type domainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DomainGetResponseEnvelopeSuccess bool

const (
	DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

func (r DomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
