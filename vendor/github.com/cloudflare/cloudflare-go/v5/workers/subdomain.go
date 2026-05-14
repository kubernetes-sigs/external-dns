// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// SubdomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubdomainService] method instead.
type SubdomainService struct {
	Options []option.RequestOption
}

// NewSubdomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubdomainService(opts ...option.RequestOption) (r *SubdomainService) {
	r = &SubdomainService{}
	r.Options = opts
	return
}

// Creates a Workers subdomain for an account.
func (r *SubdomainService) Update(ctx context.Context, params SubdomainUpdateParams, opts ...option.RequestOption) (res *SubdomainUpdateResponse, err error) {
	var env SubdomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a Workers subdomain for an account.
func (r *SubdomainService) Get(ctx context.Context, query SubdomainGetParams, opts ...option.RequestOption) (res *SubdomainGetResponse, err error) {
	var env SubdomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/subdomain", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SubdomainUpdateResponse struct {
	Subdomain string                      `json:"subdomain,required"`
	JSON      subdomainUpdateResponseJSON `json:"-"`
}

// subdomainUpdateResponseJSON contains the JSON metadata for the struct
// [SubdomainUpdateResponse]
type subdomainUpdateResponseJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponse struct {
	Subdomain string                   `json:"subdomain,required"`
	JSON      subdomainGetResponseJSON `json:"-"`
}

// subdomainGetResponseJSON contains the JSON metadata for the struct
// [SubdomainGetResponse]
type subdomainGetResponseJSON struct {
	Subdomain   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubdomainUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Subdomain param.Field[string] `json:"subdomain,required"`
}

func (r SubdomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubdomainUpdateResponseEnvelope struct {
	Errors   []SubdomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubdomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SubdomainUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success SubdomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    subdomainUpdateResponseEnvelopeJSON    `json:"-"`
}

// subdomainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubdomainUpdateResponseEnvelope]
type subdomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubdomainUpdateResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           SubdomainUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             subdomainUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// subdomainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubdomainUpdateResponseEnvelopeErrors]
type subdomainUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubdomainUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    subdomainUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// subdomainUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SubdomainUpdateResponseEnvelopeErrorsSource]
type subdomainUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SubdomainUpdateResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           SubdomainUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             subdomainUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// subdomainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubdomainUpdateResponseEnvelopeMessages]
type subdomainUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SubdomainUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    subdomainUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// subdomainUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SubdomainUpdateResponseEnvelopeMessagesSource]
type subdomainUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SubdomainUpdateResponseEnvelopeSuccess bool

const (
	SubdomainUpdateResponseEnvelopeSuccessTrue SubdomainUpdateResponseEnvelopeSuccess = true
)

func (r SubdomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubdomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubdomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SubdomainGetResponseEnvelope struct {
	Errors   []SubdomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubdomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SubdomainGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success SubdomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    subdomainGetResponseEnvelopeJSON    `json:"-"`
}

// subdomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubdomainGetResponseEnvelope]
type subdomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           SubdomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             subdomainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// subdomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SubdomainGetResponseEnvelopeErrors]
type subdomainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    subdomainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// subdomainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SubdomainGetResponseEnvelopeErrorsSource]
type subdomainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SubdomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             subdomainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// subdomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubdomainGetResponseEnvelopeMessages]
type subdomainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SubdomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    subdomainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// subdomainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SubdomainGetResponseEnvelopeMessagesSource]
type subdomainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subdomainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SubdomainGetResponseEnvelopeSuccess bool

const (
	SubdomainGetResponseEnvelopeSuccessTrue SubdomainGetResponseEnvelopeSuccess = true
)

func (r SubdomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubdomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
